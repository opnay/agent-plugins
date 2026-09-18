use std::collections::BTreeMap;
use std::env;
use std::error::Error;
use std::ffi::OsString;
use std::fmt::Write as _;
use std::fs;
use std::io::{self, Write};
use std::os::unix::fs::PermissionsExt;
use std::path::{Path, PathBuf};
use std::process::ExitCode;

type Result<T> = std::result::Result<T, Box<dyn Error>>;

const HELP: &str = "agent-wiki: 위키 저장 위치를 설정·조회합니다.

사용법:
  agent-wiki install [--force]  ~/.local/bin/agent-wiki에 설치합니다.
  agent-wiki list               설정된 root를 조회합니다.
  agent-wiki set <directory>    default root 경로를 설정합니다.
  agent-wiki set <name> <dir>   named root 경로를 설정합니다.
  agent-wiki default <name>     default root를 지정합니다.
  agent-wiki path [name]        default 또는 named root 경로를 출력합니다.
  agent-wiki --help";

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

fn parse_legacy(data: &toml::Table) -> Result<Option<WikiConfig>> {
    let Some(root) = data.get("root") else {
        return Ok(None);
    };
    if data.len() != 1 {
        return Err("legacy 설정에는 root만 포함할 수 있습니다.".into());
    }
    let root = root.as_str().ok_or("legacy root는 문자열이어야 합니다.")?;
    let path = PathBuf::from(root);
    if !path.is_absolute() {
        return Err("legacy root는 절대 경로여야 합니다.".into());
    }
    Ok(Some(WikiConfig {
        roots: [(
            "default".into(),
            RootConfig {
                path,
                description: None,
            },
        )]
        .into(),
        default: "default".into(),
    }))
}

fn read_config(config: &Path) -> Result<WikiConfig> {
    let data = fs::read_to_string(config)?.parse::<toml::Table>()?;
    if let Some(value) = parse_legacy(&data)? {
        return Ok(value);
    }
    if data.len() != 3
        || !data.contains_key("version")
        || !data.contains_key("default")
        || !data.contains_key("roots")
    {
        return Err("지원하지 않는 설정 형식입니다.".into());
    }
    if data.get("version").and_then(toml::Value::as_integer) != Some(2) {
        return Err("지원하지 않는 설정 버전입니다.".into());
    }
    let default = data
        .get("default")
        .and_then(toml::Value::as_str)
        .ok_or("v2 설정에는 default root 이름이 필요합니다.")?;
    let root_table = data
        .get("roots")
        .and_then(toml::Value::as_table)
        .ok_or("v2 설정에는 roots table이 필요합니다.")?;
    if root_table.is_empty() {
        return Err("v2 설정에는 하나 이상의 root가 필요합니다.".into());
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

fn canonical_directory(directory: &Path, home: &Path) -> Result<PathBuf> {
    if directory.as_os_str().is_empty() {
        return Err("저장 폴더를 지정하세요.".into());
    }
    let expanded = if directory == Path::new("~") {
        home.to_path_buf()
    } else if let Ok(relative) = directory.strip_prefix("~/") {
        home.join(relative)
    } else {
        directory.to_path_buf()
    };
    let root = expanded.canonicalize()?;
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
    if config.is_symlink() {
        return Err(format!("설정 심볼릭 링크를 변경하지 않습니다: {}", config.display()).into());
    }
    let canonical_roots = validated_roots(&value.roots)?;
    let mut data = toml::Table::new();
    data.insert("version".into(), toml::Value::Integer(2));
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
    atomic_write(config, toml::to_string(&data)?.as_bytes(), 0o600, true)
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

fn configured_root(config: &Path, name: Option<&OsString>) -> Result<PathBuf> {
    if !config.try_exists()? {
        return Err("저장 위치가 없습니다. agent-wiki set <directory>를 실행하세요.".into());
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
        [command] if command == "install" => install(
            &env::current_exe()?,
            &home.join(".local/bin/agent-wiki"),
            false,
        ),
        [command, option] if command == "install" && option == "--force" => install(
            &env::current_exe()?,
            &home.join(".local/bin/agent-wiki"),
            true,
        ),
        [command, directory] if command == "set" => {
            let name = if config.try_exists()? {
                read_config(&config)?.default
            } else {
                "default".into()
            };
            set_root(&name, Path::new(directory), home, &config)
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
        _ => Err(format!("명령 또는 인자가 잘못되었습니다.\n{HELP}").into()),
    }
}

fn main() -> ExitCode {
    let args: Vec<_> = env::args_os().skip(1).collect();
    if matches!(args.as_slice(), [arg] if arg == "--help" || arg == "-h")
        || matches!(args.as_slice(), [command, option] if (command == "install" || command == "set" || command == "path" || command == "list" || command == "default") && (option == "--help" || option == "-h"))
    {
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
            eprintln!("agent-wiki: {error}");
            ExitCode::FAILURE
        }
    }
}

#[cfg(test)]
#[path = "../tests/cli.rs"]
mod tests;
