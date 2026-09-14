use std::env;
use std::error::Error;
use std::ffi::OsString;
use std::fs;
use std::io::{self, Write};
use std::os::unix::fs::PermissionsExt;
use std::path::{Path, PathBuf};
use std::process::ExitCode;

type Result<T> = std::result::Result<T, Box<dyn Error>>;

const HELP: &str = "agent-wiki: 위키 저장 위치를 설정·조회합니다.

사용법:
  agent-wiki install [--force]  ~/.local/bin/agent-wiki에 설치합니다.
  agent-wiki set <directory>    기존 위키 폴더를 지정합니다.
  agent-wiki path               설정된 절대 경로를 출력합니다.
  agent-wiki --help";

fn read_root(config: &Path) -> Result<PathBuf> {
    let data = fs::read_to_string(config)?.parse::<toml::Table>()?;
    let root = data.get("root").and_then(toml::Value::as_str);
    match root {
        Some(root) if data.len() == 1 && Path::new(root).is_absolute() => Ok(root.into()),
        _ => Err(format!(
            "설정은 절대 경로 root 문자열 하나만 지원합니다: {}",
            config.display()
        )
        .into()),
    }
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

fn set_root(directory: &Path, home: &Path, config: &Path) -> Result<PathBuf> {
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
    if config.is_symlink() {
        return Err(format!("설정 심볼릭 링크를 변경하지 않습니다: {}", config.display()).into());
    }
    if config.try_exists()? {
        read_root(config)?;
    }
    let mut data = toml::Table::new();
    let value = root.to_str().ok_or("설정 경로는 UTF-8이어야 합니다.")?;
    data.insert("root".into(), toml::Value::String(value.into()));
    atomic_write(config, toml::to_string(&data)?.as_bytes(), 0o600, true)?;
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
        [command, directory] if command == "set" => set_root(Path::new(directory), home, &config),
        [command] if command == "path" => {
            if !config.try_exists()? {
                return Err(
                    "저장 위치가 없습니다. agent-wiki set <directory>를 실행하세요.".into(),
                );
            }
            let root = read_root(&config)?;
            if !root.is_dir() {
                return Err(format!(
                    "저장 폴더가 존재하지 않거나 디렉터리가 아닙니다: {}",
                    root.display()
                )
                .into());
            }
            Ok(root)
        }
        _ => Err(format!("명령 또는 인자가 잘못되었습니다.\n{HELP}").into()),
    }
}

fn main() -> ExitCode {
    let args: Vec<_> = env::args_os().skip(1).collect();
    if matches!(args.as_slice(), [arg] if arg == "--help" || arg == "-h")
        || matches!(args.as_slice(), [command, option] if (command == "install" || command == "set" || command == "path") && (option == "--help" || option == "-h"))
    {
        println!("{HELP}");
        return ExitCode::SUCCESS;
    }
    let result = env::home_dir()
        .ok_or_else(|| "홈 디렉터리를 찾을 수 없습니다.".into())
        .and_then(|home| run(&args, &home));
    match result {
        Ok(path) => {
            println!("{}", path.display());
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
