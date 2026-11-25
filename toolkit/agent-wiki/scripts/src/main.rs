use std::collections::BTreeMap;
use std::env;
use std::error::Error;
use std::ffi::{OsStr, OsString};
use std::fmt::Write as _;
use std::fs;
use std::io::{self, Read, Write};
use std::os::unix::fs::PermissionsExt;
use std::path::{Path, PathBuf};
use std::process::ExitCode;

type Result<T> = std::result::Result<T, Box<dyn Error>>;

mod edit;
mod inspect;

const COMMAND: &str = "wdown";

const HELP: &str = "wdown: 위키 저장 위치를 설정·조회하고 Markdown 구조를 점검·편집합니다.

사용법:
  wdown install [--force] [--dir <dir>]   <dir>/wdown에 현재 실행 파일을 복사합니다.
  wdown uninstall [--dir <dir>]           <dir>/wdown 실행 파일을 제거합니다.
  wdown doctor [--dir <dir>]              설치·PATH·설정 상태를 점검합니다.
  wdown list                              설정된 root를 조회합니다.
  wdown set <name> <dir>                  named root 경로를 설정합니다.
  wdown default <name>                    default root를 지정합니다.
  wdown path [name]                       default 또는 named root 경로를 출력합니다.
  wdown inspect <file> [--id <id>]        Markdown 구조 트리를 출력합니다.
  wdown edit <file> --id <id> [--write]   표준 입력으로 해당 노드를 대체합니다. 기본은 dry-run입니다.
  wdown --help";

#[derive(Clone, Debug, PartialEq, Eq)]
struct RootConfig {
    path: PathBuf,
    description: Option<String>,
}

#[derive(Clone, Debug, PartialEq, Eq)]
struct WikiConfig {
    roots: BTreeMap<String, RootConfig>,
    default: String,
}

fn valid_root_name(name: &str) -> bool {
    !name.is_empty()
        && name.chars().all(|character| {
            character.is_ascii_alphanumeric() || character == '-' || character == '_'
        })
}

fn read_config(config: &Path) -> Result<WikiConfig> {
    let data = fs::read_to_string(config)?.parse::<toml::Table>()?;
    if data.len() != 2 || !data.contains_key("default") || !data.contains_key("roots") {
        return Err("지원하지 않는 설정 형식입니다.".into());
    }
    let default = data
        .get("default")
        .and_then(toml::Value::as_str)
        .ok_or("설정에는 default root 이름이 필요합니다.")?;
    let root_table = data
        .get("roots")
        .and_then(toml::Value::as_table)
        .ok_or("설정에는 roots table이 필요합니다.")?;
    if root_table.is_empty() {
        return Err("설정에는 하나 이상의 root가 필요합니다.".into());
    }

    let mut roots = BTreeMap::new();
    for (name, value) in root_table {
        if !valid_root_name(name) {
            return Err("root 이름이 잘못되었습니다.".into());
        }
        let root = value.as_table().ok_or("root 설정은 table이어야 합니다.")?;
        if root.is_empty()
            || root.len() > 2
            || !root.contains_key("path")
            || root.keys().any(|key| key != "path" && key != "description")
        {
            return Err("root 설정에는 path와 선택적인 description만 포함할 수 있습니다.".into());
        }
        let path = root
            .get("path")
            .and_then(toml::Value::as_str)
            .ok_or("root path가 필요합니다.")?;
        let path = PathBuf::from(path);
        if !path.is_absolute() {
            return Err("root path는 절대 경로여야 합니다.".into());
        }
        let description = match root.get("description") {
            Some(value) => Some(
                value
                    .as_str()
                    .ok_or("root description은 문자열이어야 합니다.")?
                    .into(),
            ),
            None => None,
        };
        roots.insert(name.clone(), RootConfig { path, description });
    }
    if !roots.contains_key(default) {
        return Err("default root가 없습니다.".into());
    }
    Ok(WikiConfig {
        roots,
        default: default.into(),
    })
}

fn atomic_write(target: &Path, content: &[u8], mode: u32, overwrite: bool) -> Result<()> {
    let parent = target.parent().ok_or("저장할 상위 폴더가 없습니다.")?;
    fs::create_dir_all(parent)?;
    let mut temporary = tempfile::NamedTempFile::new_in(parent)?;
    temporary.write_all(content)?;
    temporary
        .as_file()
        .set_permissions(fs::Permissions::from_mode(mode))?;
    temporary.as_file().sync_all()?;
    if overwrite {
        temporary.persist(target)?;
    } else {
        temporary.persist_noclobber(target)?;
    }
    Ok(())
}

fn expand_home(directory: &Path, home: &Path) -> PathBuf {
    if directory == Path::new("~") {
        home.to_path_buf()
    } else if let Ok(relative) = directory.strip_prefix("~/") {
        home.join(relative)
    } else {
        directory.to_path_buf()
    }
}

fn canonical_directory(directory: &Path, home: &Path) -> Result<PathBuf> {
    if directory.as_os_str().is_empty() {
        return Err("저장 폴더를 지정하세요.".into());
    }
    let root = expand_home(directory, home).canonicalize()?;
    if !root.is_dir() {
        return Err(format!("디렉터리가 아닙니다: {}", root.display()).into());
    }
    Ok(root)
}

fn canonical_root(path: &Path) -> Result<PathBuf> {
    let root = path.canonicalize()?;
    if !root.is_dir() {
        return Err(format!(
            "root가 존재하지 않거나 디렉터리가 아닙니다: {}",
            path.display()
        )
        .into());
    }
    Ok(root)
}

fn validated_roots(roots: &BTreeMap<String, RootConfig>) -> Result<BTreeMap<String, PathBuf>> {
    let canonical_roots: BTreeMap<_, _> = roots
        .iter()
        .map(|(name, root)| Ok((name.clone(), canonical_root(&root.path)?)))
        .collect::<Result<_>>()?;
    let paths: Vec<_> = canonical_roots.values().collect();
    for (index, first) in paths.iter().enumerate() {
        for second in paths.iter().skip(index + 1) {
            if first == second || first.starts_with(second) || second.starts_with(first) {
                return Err("root 경로는 중복되거나 중첩될 수 없습니다.".into());
            }
        }
    }
    Ok(canonical_roots)
}

fn format_root_list(value: &WikiConfig) -> Result<String> {
    let canonical_roots = validated_roots(&value.roots)?;
    let mut output = String::new();
    for (name, root) in &value.roots {
        let description = root
            .description
            .as_deref()
            .map(|description| format!("{description:?}"))
            .unwrap_or_else(|| "-".into());
        writeln!(
            output,
            "{}\t{}\tdefault={}\tdescription={}",
            name,
            canonical_roots[name].display(),
            name == &value.default,
            description
        )?;
    }
    Ok(output)
}

fn write_config(config: &Path, value: &WikiConfig) -> Result<()> {
    let target = if config.is_symlink() {
        config.canonicalize()?
    } else {
        config.to_path_buf()
    };
    let canonical_roots = validated_roots(&value.roots)?;
    let mut data = toml::Table::new();
    data.insert("default".into(), toml::Value::String(value.default.clone()));
    let mut roots = toml::Table::new();
    for (name, root) in &value.roots {
        let mut item = toml::Table::new();
        item.insert(
            "path".into(),
            toml::Value::String(
                canonical_roots[name]
                    .to_str()
                    .ok_or("설정 경로는 UTF-8이어야 합니다.")?
                    .into(),
            ),
        );
        if let Some(description) = &root.description {
            item.insert(
                "description".into(),
                toml::Value::String(description.clone()),
            );
        }
        roots.insert(name.clone(), toml::Value::Table(item));
    }
    data.insert("roots".into(), toml::Value::Table(roots));
    atomic_write(&target, toml::to_string(&data)?.as_bytes(), 0o600, true)
}

fn set_root(name: &str, directory: &Path, home: &Path, config: &Path) -> Result<PathBuf> {
    if !valid_root_name(name) {
        return Err("root 이름이 잘못되었습니다.".into());
    }
    let root = canonical_directory(directory, home)?;
    let mut value = if config.try_exists()? {
        read_config(config)?
    } else {
        WikiConfig {
            roots: BTreeMap::new(),
            default: name.into(),
        }
    };
    let description = value
        .roots
        .get(name)
        .and_then(|existing| existing.description.clone());
    value.roots.insert(
        name.into(),
        RootConfig {
            path: root.clone(),
            description,
        },
    );
    validated_roots(&value.roots)?;
    write_config(config, &value)?;
    Ok(root)
}

fn install(source: &Path, target: &Path, force: bool) -> Result<PathBuf> {
    let content = fs::read(source)?;
    match fs::symlink_metadata(target) {
        Ok(metadata) => {
            if !metadata.file_type().is_file() {
                return Err(format!(
                    "일반 파일 이외의 대상은 교체하지 않습니다: {}",
                    target.display()
                )
                .into());
            }
            if fs::read(target)? == content && metadata.permissions().mode() & 0o111 != 0 {
                return Ok(target.into());
            }
            if !force {
                return Err(format!(
                    "기존 파일을 보존합니다. 교체하려면 install --force: {}",
                    target.display()
                )
                .into());
            }
        }
        Err(error) if error.kind() == io::ErrorKind::NotFound => {}
        Err(error) => return Err(error.into()),
    }
    atomic_write(target, &content, 0o755, force)?;
    Ok(target.into())
}

struct MaintenanceOptions {
    force: bool,
    directory: PathBuf,
}

fn parse_maintenance_options(
    args: &[OsString],
    home: &Path,
    force_allowed: bool,
) -> Result<MaintenanceOptions> {
    let mut force = false;
    let mut directory: Option<PathBuf> = None;
    let mut index = 0;
    while index < args.len() {
        match args[index].to_string_lossy().as_ref() {
            "--force" => {
                if !force_allowed {
                    return Err("--force는 install 명령에서만 사용할 수 있습니다.".into());
                }
                if force {
                    return Err("--force는 한 번만 지정할 수 있습니다.".into());
                }
                force = true;
                index += 1;
            }
            "--dir" => {
                if directory.is_some() {
                    return Err("--dir는 한 번만 지정할 수 있습니다.".into());
                }
                let value = args.get(index + 1).ok_or("--dir에는 폴더를 지정하세요.")?;
                if value.is_empty() {
                    return Err("--dir에는 폴더를 지정하세요.".into());
                }
                directory = Some(expand_home(Path::new(value), home));
                index += 2;
            }
            value => return Err(format!("옵션이 잘못되었습니다: {value}").into()),
        }
    }
    Ok(MaintenanceOptions {
        force,
        directory: directory.unwrap_or_else(|| home.join(".local/bin")),
    })
}

fn uninstall(target: &Path) -> Result<Option<PathBuf>> {
    match fs::symlink_metadata(target) {
        Ok(metadata) => {
            if !metadata.file_type().is_file() {
                return Err(format!(
                    "일반 파일 이외의 대상은 제거하지 않습니다: {}",
                    target.display()
                )
                .into());
            }
            if metadata.permissions().mode() & 0o111 == 0 {
                return Err(format!(
                    "실행 권한이 없는 파일은 제거하지 않습니다: {}",
                    target.display()
                )
                .into());
            }
            fs::remove_file(target)?;
            Ok(Some(target.into()))
        }
        Err(error) if error.kind() == io::ErrorKind::NotFound => Ok(None),
        Err(error) => Err(error.into()),
    }
}

#[derive(Clone, Copy, PartialEq, Eq)]
enum Status {
    Ok,
    Fail,
    Skip,
}

fn record(
    report: &mut String,
    counts: &mut [usize; 3],
    status: Status,
    name: &str,
    detail: String,
) {
    let label = match status {
        Status::Ok => "OK",
        Status::Fail => "FAIL",
        Status::Skip => "SKIP",
    };
    counts[status as usize] += 1;
    let _ = writeln!(report, "{label:<4} {name:<6} {detail}");
}

fn doctor_report(directory: &Path, home: &Path, path: Option<&OsStr>) -> Result<(String, usize)> {
    let target = directory.join(COMMAND);
    let mut report = String::new();
    let mut counts = [0; 3];

    match fs::symlink_metadata(&target) {
        Ok(metadata) if metadata.file_type().is_file() => {
            let mode = metadata.permissions().mode() & 0o777;
            if fs::read(&target)? != fs::read(env::current_exe()?)? {
                record(
                    &mut report,
                    &mut counts,
                    Status::Fail,
                    "install",
                    format!(
                        "현재 실행 파일과 다릅니다. wdown install --force --dir {}: {}",
                        directory.display(),
                        target.display()
                    ),
                );
            } else if mode & 0o111 == 0 {
                record(
                    &mut report,
                    &mut counts,
                    Status::Fail,
                    "install",
                    format!("실행 권한이 없습니다: {}", target.display()),
                );
            } else {
                record(
                    &mut report,
                    &mut counts,
                    Status::Ok,
                    "install",
                    format!("{} ({mode:o})", target.display()),
                );
            }
        }
        Ok(_) => record(
            &mut report,
            &mut counts,
            Status::Fail,
            "install",
            format!("일반 파일이 아닙니다: {}", target.display()),
        ),
        Err(error) if error.kind() == io::ErrorKind::NotFound => record(
            &mut report,
            &mut counts,
            Status::Skip,
            "install",
            format!(
                "설치 파일이 없습니다. wdown install --dir {}",
                directory.display()
            ),
        ),
        Err(error) => return Err(error.into()),
    }

    let selected = path.and_then(|value| {
        env::split_paths(value)
            .map(|entry| entry.join(COMMAND))
            .find(|candidate| candidate.is_file())
    });
    match selected {
        Some(selected) => {
            let same = fs::canonicalize(&selected)
                .ok()
                .zip(fs::canonicalize(&target).ok())
                .is_some_and(|(selected, target)| selected == target);
            if target.is_file() && !same {
                record(
                    &mut report,
                    &mut counts,
                    Status::Fail,
                    "path",
                    format!(
                        "PATH에서 다른 실행 파일이 먼저 선택됩니다: {}",
                        selected.display()
                    ),
                );
            } else {
                record(
                    &mut report,
                    &mut counts,
                    Status::Ok,
                    "path",
                    selected.display().to_string(),
                );
            }
        }
        None => record(
            &mut report,
            &mut counts,
            Status::Fail,
            "path",
            format!(
                "PATH에 {COMMAND}이 없습니다. export PATH=\"{}:$PATH\"",
                directory.display()
            ),
        ),
    }

    let config = home.join(".agents/config.wiki.toml");
    match fs::symlink_metadata(&config) {
        Err(error) if error.kind() == io::ErrorKind::NotFound => record(
            &mut report,
            &mut counts,
            Status::Skip,
            "config",
            format!("저장 위치가 없습니다. {COMMAND} set <name> <directory>"),
        ),
        Err(error) => return Err(error.into()),
        Ok(_) => {
            let mode = fs::metadata(&config)
                .or_else(|_| fs::symlink_metadata(&config))
                .map(|metadata| format!(" ({:o})", metadata.permissions().mode() & 0o777))
                .unwrap_or_default();
            match read_config(&config).and_then(|value| validated_roots(&value.roots).map(|_| ())) {
                Ok(()) => record(
                    &mut report,
                    &mut counts,
                    Status::Ok,
                    "config",
                    format!("{}{mode}", config.display()),
                ),
                Err(error) => record(
                    &mut report,
                    &mut counts,
                    Status::Fail,
                    "config",
                    format!("{}{mode}: {error}", config.display()),
                ),
            }
        }
    }

    let _ = writeln!(
        report,
        "합계: OK={} FAIL={} SKIP={}",
        counts[Status::Ok as usize],
        counts[Status::Fail as usize],
        counts[Status::Skip as usize]
    );
    Ok((report, counts[Status::Fail as usize]))
}

fn doctor(directory: &Path, home: &Path) -> Result<()> {
    let (report, failures) = doctor_report(directory, home, env::var_os("PATH").as_deref())?;
    io::stdout().write_all(report.as_bytes())?;
    if failures > 0 {
        return Err(format!("점검 실패: FAIL {failures}건").into());
    }
    Ok(())
}

fn configured_root(config: &Path, name: Option<&OsString>) -> Result<PathBuf> {
    if !config.try_exists()? {
        return Err("저장 위치가 없습니다. wdown set <name> <directory>를 실행하세요.".into());
    }
    let value = read_config(config)?;
    let canonical_roots = validated_roots(&value.roots)?;
    let name = name
        .map(|value| value.to_string_lossy().into_owned())
        .unwrap_or(value.default);
    let root = canonical_roots
        .get(&name)
        .ok_or("존재하지 않는 root입니다.")?
        .clone();
    Ok(root)
}

fn run(args: &[OsString], home: &Path) -> Result<PathBuf> {
    let config = home.join(".agents/config.wiki.toml");
    match args {
        [command, options @ ..] if command == "install" => {
            let options = parse_maintenance_options(options, home, true)?;
            install(
                &env::current_exe()?,
                &options.directory.join(COMMAND),
                options.force,
            )
        }
        [command, options @ ..] if command == "uninstall" => {
            let options = parse_maintenance_options(options, home, false)?;
            Ok(uninstall(&options.directory.join(COMMAND))?.unwrap_or_default())
        }
        [command, options @ ..] if command == "doctor" => {
            let options = parse_maintenance_options(options, home, false)?;
            doctor(&options.directory, home)?;
            Ok(PathBuf::new())
        }
        [command, name, directory] if command == "set" => {
            set_root(&name.to_string_lossy(), Path::new(directory), home, &config)
        }
        [command, name] if command == "default" => {
            let mut value = read_config(&config)?;
            let name = name.to_string_lossy().into_owned();
            if !value.roots.contains_key(&name) {
                return Err("존재하지 않는 root입니다.".into());
            }
            let canonical_roots = validated_roots(&value.roots)?;
            value.default = name;
            let root = canonical_roots[&value.default].clone();
            write_config(&config, &value)?;
            Ok(root)
        }
        [command] if command == "list" => {
            let value = read_config(&config)?;
            io::stdout().write_all(format_root_list(&value)?.as_bytes())?;
            Ok(PathBuf::new())
        }
        [command] if command == "path" => configured_root(&config, None),
        [command, name] if command == "path" => configured_root(&config, Some(name)),
        [command, options @ ..] if command == "inspect" => {
            let (target, selection) = inspect::parse_options(options)?;
            let source = fs::read_to_string(&target).map_err(|error| {
                format!(
                    "Markdown을 읽을 수 없습니다: {} ({error})",
                    target.display()
                )
            })?;
            let document = inspect::parse_markdown(&source);
            let output = inspect::render(&document, &target, selection.as_deref())?;
            io::stdout().write_all(output.as_bytes())?;
            Ok(PathBuf::new())
        }
        [command, options @ ..] if command == "edit" => {
            let (target, id, write) = edit::parse_options(options)?;
            let mut input = String::new();
            io::stdin().read_to_string(&mut input)?;
            let output = edit_target(&target, &id, write, &input)?;
            io::stdout().write_all(output.as_bytes())?;
            Ok(PathBuf::new())
        }
        _ => Err(format!("명령 또는 인자가 잘못되었습니다.\n{HELP}").into()),
    }
}

/// `--id`로 고른 노드를 stdin 본문으로 대체한다. `write`가 아니고 변경이 있을 때만 저장한다.
fn edit_target(target: &Path, id: &str, write: bool, input: &str) -> Result<String> {
    let source = fs::read_to_string(target).map_err(|error| {
        format!(
            "Markdown을 읽을 수 없습니다: {} ({error})",
            target.display()
        )
    })?;
    let plan = edit::plan(&source, target, id, input)?;
    let saved = if write && plan.changed() {
        let metadata = fs::symlink_metadata(target).map_err(|error| {
            format!(
                "Markdown 상태를 확인할 수 없습니다: {} ({error})",
                target.display()
            )
        })?;
        if metadata.file_type().is_symlink() {
            return Err(
                format!("심볼릭 링크 파일은 편집하지 않습니다: {}", target.display()).into(),
            );
        }
        let mode = metadata.permissions().mode() & 0o777;
        atomic_write(target, edit::apply(&source, &plan).as_bytes(), mode, true)?;
        true
    } else {
        false
    };
    Ok(edit::render_plan(&plan, target, saved))
}

fn help_requested(args: &[OsString]) -> bool {
    let is_help = |value: &OsString| value == "--help" || value == "-h";
    match args {
        [first] => is_help(first),
        [command, second, ..] => {
            is_help(second)
                && [
                    COMMAND,
                    "install",
                    "uninstall",
                    "doctor",
                    "list",
                    "set",
                    "default",
                    "path",
                    "inspect",
                    "edit",
                ]
                .contains(&command.to_string_lossy().as_ref())
        }
        _ => false,
    }
}

fn main() -> ExitCode {
    let args: Vec<_> = env::args_os().skip(1).collect();
    if help_requested(&args) {
        println!("{HELP}");
        return ExitCode::SUCCESS;
    }
    let result = env::home_dir()
        .ok_or_else(|| "홈 디렉터리를 찾을 수 없습니다.".into())
        .and_then(|home| run(&args, &home));
    match result {
        Ok(path) => {
            if !path.as_os_str().is_empty() {
                println!("{}", path.display());
            }
            ExitCode::SUCCESS
        }
        Err(error) => {
            eprintln!("{COMMAND}: {error}");
            ExitCode::FAILURE
        }
    }
}

#[cfg(test)]
#[path = "../tests/cli.rs"]
mod tests;

#[cfg(test)]
#[path = "../tests/inspect.rs"]
mod inspect_tests;

#[cfg(test)]
#[path = "../tests/edit.rs"]
mod edit_tests;
