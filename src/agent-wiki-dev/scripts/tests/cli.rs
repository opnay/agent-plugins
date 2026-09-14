use super::*;
use std::os::unix::fs::symlink;

fn args(values: &[&str]) -> Vec<OsString> {
    values.iter().map(OsString::from).collect()
}

#[test]
fn unset_path_does_not_create_config() {
    let home = tempfile::tempdir().unwrap();
    assert!(run(&args(&["path"]), home.path()).is_err());
    assert!(!home.path().join(".agents").exists());
}

#[test]
fn special_paths_round_trip_without_wiki_writes() {
    let home = tempfile::tempdir().unwrap();
    let wiki = home.path().join("자료 \"React\" \\ Wiki\n19");
    fs::create_dir(&wiki).unwrap();
    let config = home.path().join(".agents/config.wiki.toml");
    let expected = wiki.canonicalize().unwrap();
    assert_eq!(set_root(&wiki, home.path(), &config).unwrap(), expected);
    assert_eq!(run(&args(&["path"]), home.path()).unwrap(), expected);
    assert_eq!(fs::read_dir(wiki).unwrap().count(), 0);
}

#[test]
fn relative_and_tilde_paths_are_normalized() {
    let home = tempfile::tempdir().unwrap();
    let config = home.path().join(".agents/config.wiki.toml");
    let result = set_root(Path::new("."), home.path(), &config).unwrap();
    assert_eq!(result, env::current_dir().unwrap().canonicalize().unwrap());
    let wiki = home.path().join("Wiki");
    fs::create_dir(&wiki).unwrap();
    assert_eq!(
        set_root(Path::new("~/Wiki"), home.path(), &config).unwrap(),
        wiki.canonicalize().unwrap()
    );
    assert_eq!(
        set_root(Path::new("~"), home.path(), &config).unwrap(),
        home.path().canonicalize().unwrap()
    );
}

#[test]
fn invalid_directory_preserves_config() {
    let home = tempfile::tempdir().unwrap();
    let config = home.path().join(".agents/config.wiki.toml");
    set_root(home.path(), home.path(), &config).unwrap();
    let before = fs::read(&config).unwrap();
    for directory in [PathBuf::new(), home.path().join("missing"), config.clone()] {
        assert!(set_root(&directory, home.path(), &config).is_err());
        assert_eq!(fs::read(&config).unwrap(), before);
    }
}

#[test]
fn invalid_config_is_not_overwritten() {
    let home = tempfile::tempdir().unwrap();
    let config = home.path().join("config.wiki.toml");
    for content in [
        "root = \"unterminated",
        "root = 4\n",
        "root = 'relative'\n",
        "root = '/tmp'\nother = true\n",
    ] {
        fs::write(&config, content).unwrap();
        assert!(set_root(home.path(), home.path(), &config).is_err());
        assert_eq!(fs::read_to_string(&config).unwrap(), content);
    }
}

#[test]
fn existing_python_config_remains_compatible() {
    let home = tempfile::tempdir().unwrap();
    let config = home.path().join(".agents/config.wiki.toml");
    fs::create_dir(config.parent().unwrap()).unwrap();
    fs::write(&config, "# existing config\nroot = \"/tmp\"\n").unwrap();
    assert_eq!(read_root(&config).unwrap(), Path::new("/tmp"));
    set_root(home.path(), home.path(), &config).unwrap();
    assert_eq!(
        read_root(&config).unwrap(),
        home.path().canonicalize().unwrap()
    );
}

#[test]
fn deleted_wiki_is_reported() {
    let home = tempfile::tempdir().unwrap();
    let wiki = home.path().join("Wiki");
    fs::create_dir(&wiki).unwrap();
    let config = home.path().join(".agents/config.wiki.toml");
    set_root(&wiki, home.path(), &config).unwrap();
    fs::remove_dir(wiki).unwrap();
    assert!(run(&args(&["path"]), home.path()).is_err());
}

#[test]
fn failed_persist_cleans_temporary_file() {
    let directory = tempfile::tempdir().unwrap();
    let target = directory.path().join("existing");
    fs::create_dir(&target).unwrap();
    let sentinel = target.join("keep");
    fs::write(&sentinel, "keep").unwrap();
    assert!(atomic_write(&target, b"replace", 0o600, true).is_err());
    assert_eq!(fs::read_to_string(sentinel).unwrap(), "keep");
    assert_eq!(fs::read_dir(directory.path()).unwrap().count(), 1);
}

#[test]
fn install_copies_a_standalone_executable_and_is_idempotent() {
    let home = tempfile::tempdir().unwrap();
    let source = env::current_exe().unwrap();
    let target = home.path().join(".local/bin/agent-wiki");
    install(&source, &target, false).unwrap();
    assert_eq!(fs::read(&source).unwrap(), fs::read(&target).unwrap());
    assert_eq!(
        fs::metadata(&target).unwrap().permissions().mode() & 0o777,
        0o755
    );
    let before = fs::metadata(&target).unwrap().modified().unwrap();
    install(&source, &target, false).unwrap();
    assert_eq!(fs::metadata(&target).unwrap().modified().unwrap(), before);
    // The copied native test executable runs without PATH or interpreter variables.
    let result = std::process::Command::new(&target)
        .arg("--list")
        .env_clear()
        .output()
        .unwrap();
    assert!(result.status.success());
    assert!(!home.path().join(".agents").exists());
}

#[test]
fn install_preserves_collisions_unless_forced() {
    let home = tempfile::tempdir().unwrap();
    let source = env::current_exe().unwrap();
    let target = home.path().join("agent-wiki");
    fs::write(&target, "existing command").unwrap();
    assert!(install(&source, &target, false).is_err());
    assert_eq!(fs::read_to_string(&target).unwrap(), "existing command");
    install(&source, &target, true).unwrap();
    assert_eq!(fs::read(source).unwrap(), fs::read(target).unwrap());
}

#[test]
fn symlinks_and_directories_are_preserved() {
    let home = tempfile::tempdir().unwrap();
    let original = home.path().join("original");
    fs::write(&original, "root = '/tmp'\n").unwrap();
    let target = home.path().join("link");
    symlink(&original, &target).unwrap();
    assert!(set_root(home.path(), home.path(), &target).is_err());
    assert!(install(&env::current_exe().unwrap(), &target, true).is_err());
    assert!(target.is_symlink());
    assert!(install(&env::current_exe().unwrap(), home.path(), true).is_err());
    assert_eq!(fs::read_to_string(original).unwrap(), "root = '/tmp'\n");
}

#[test]
fn invalid_arguments_do_not_write_files() {
    let home = tempfile::tempdir().unwrap();
    for input in [
        args(&[]),
        args(&["set"]),
        args(&["install", "--unknown"]),
        args(&["path", "extra"]),
    ] {
        assert!(run(&input, home.path()).is_err());
    }
    assert_eq!(fs::read_dir(home.path()).unwrap().count(), 0);
}
