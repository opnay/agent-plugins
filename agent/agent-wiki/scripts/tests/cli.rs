use super::*;
use std::os::unix::fs::symlink;

fn args(values: &[&str]) -> Vec<OsString> {
    values.iter().map(OsString::from).collect()
}

fn config(home: &Path) -> PathBuf {
    home.join(".agents/config.wiki.toml")
}

#[test]
fn unset_path_does_not_create_config() {
    let home = tempfile::tempdir().unwrap();
    assert!(run(&args(&["path"]), home.path()).is_err());
    assert!(!home.path().join(".agents").exists());
}

#[test]
fn legacy_config_is_rejected_without_overwriting_bytes() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("wiki");
    fs::create_dir(&root).unwrap();
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    let content = format!("root = {:?}\n", root);
    fs::write(&target, &content).unwrap();

    for input in [
        args(&["list"]),
        args(&["path"]),
        args(&["path", "default"]),
        args(&["default", "default"]),
        args(&["set", "knowledge", root.to_str().unwrap()]),
    ] {
        assert!(run(&input, home.path()).is_err());
        assert_eq!(fs::read_to_string(&target).unwrap(), content);
    }
}

#[test]
fn unnamed_set_is_rejected_without_creating_or_changing_config() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("wiki");
    fs::create_dir(&root).unwrap();
    let command = args(&["set", root.to_str().unwrap()]);
    assert!(run(&command, home.path()).is_err());
    assert!(!config(home.path()).exists());

    run(
        &args(&["set", "knowledge", root.to_str().unwrap()]),
        home.path(),
    )
    .unwrap();
    let before = fs::read(config(home.path())).unwrap();
    assert!(run(&command, home.path()).is_err());
    assert_eq!(fs::read(config(home.path())).unwrap(), before);
}

#[test]
fn named_roots_support_path_and_default_selection() {
    let home = tempfile::tempdir().unwrap();
    let knowledge = home.path().join("knowledge");
    let personality = home.path().join("personality");
    fs::create_dir(&knowledge).unwrap();
    fs::create_dir(&personality).unwrap();

    run(
        &args(&["set", "knowledge", knowledge.to_str().unwrap()]),
        home.path(),
    )
    .unwrap();
    run(
        &args(&["set", "personality", personality.to_str().unwrap()]),
        home.path(),
    )
    .unwrap();
    assert_eq!(
        run(&args(&["path"]), home.path()).unwrap(),
        knowledge.canonicalize().unwrap()
    );
    assert_eq!(
        run(&args(&["path", "personality"]), home.path()).unwrap(),
        personality.canonicalize().unwrap()
    );
    assert_eq!(
        run(&args(&["default", "personality"]), home.path()).unwrap(),
        personality.canonicalize().unwrap()
    );
    assert_eq!(
        run(&args(&["path"]), home.path()).unwrap(),
        personality.canonicalize().unwrap()
    );
}

#[test]
fn named_config_preserves_optional_description() {
    let home = tempfile::tempdir().unwrap();
    let knowledge = home.path().join("knowledge");
    let work = home.path().join("work");
    let configured_work = work.join("..").join("work");
    let refreshed_knowledge = home.path().join("refreshed-knowledge");
    fs::create_dir(&knowledge).unwrap();
    fs::create_dir(&work).unwrap();
    fs::create_dir(&refreshed_knowledge).unwrap();
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    fs::write(
        &target,
        format!(
            "default = \"knowledge\"\n\n[roots.knowledge]\npath = {:?}\ndescription = \"shared\"\n\n[roots.work]\npath = {:?}\n",
            knowledge, configured_work
        ),
    )
    .unwrap();

    run(
        &args(&["set", "knowledge", refreshed_knowledge.to_str().unwrap()]),
        home.path(),
    )
    .unwrap();
    run(&args(&["default", "work"]), home.path()).unwrap();
    let parsed = read_config(&target).unwrap();
    assert_eq!(
        parsed.roots["knowledge"].description.as_deref(),
        Some("shared")
    );
    assert_eq!(parsed.default, "work");
    assert_eq!(parsed.roots["work"].path, work.canonicalize().unwrap());
    assert_eq!(
        format_root_list(&parsed).unwrap(),
        format!(
            "knowledge\t{}\tdefault=false\tdescription=\"shared\"\nwork\t{}\tdefault=true\tdescription=-\n",
            refreshed_knowledge.canonicalize().unwrap().display(),
            work.canonicalize().unwrap().display()
        )
    );
}

#[test]
fn query_commands_reject_duplicate_named_roots_without_overwriting_config() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("wiki");
    fs::create_dir(&root).unwrap();
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    fs::write(
        &target,
        format!(
            "default = \"knowledge\"\n\n[roots.knowledge]\npath = {:?}\n\n[roots.personality]\npath = {:?}\n",
            root, root
        ),
    )
    .unwrap();
    let before = fs::read(&target).unwrap();

    assert!(run(&args(&["list"]), home.path()).is_err());
    assert!(run(&args(&["path", "knowledge"]), home.path()).is_err());
    assert_eq!(fs::read(target).unwrap(), before);
}

#[test]
fn duplicate_and_nested_roots_are_rejected_without_overwriting_config() {
    let home = tempfile::tempdir().unwrap();
    let first = home.path().join("first");
    let nested = first.join("nested");
    let sibling = home.path().join("sibling");
    fs::create_dir(&first).unwrap();
    fs::create_dir(&nested).unwrap();
    fs::create_dir(&sibling).unwrap();
    let target = config(home.path());

    set_root("first", &first, home.path(), &target).unwrap();
    let before = fs::read(&target).unwrap();
    assert!(set_root("same", &first, home.path(), &target).is_err());
    assert_eq!(fs::read(&target).unwrap(), before);
    assert!(set_root("nested", &nested, home.path(), &target).is_err());
    assert_eq!(fs::read(&target).unwrap(), before);
    set_root("sibling", &sibling, home.path(), &target).unwrap();
}

#[test]
fn invalid_config_and_invalid_target_preserve_prior_bytes() {
    let home = tempfile::tempdir().unwrap();
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    for content in [
        "root = \"unterminated",
        "unexpected = true\ndefault = \"default\"\n[roots.default]\npath = \"/tmp\"\n",
        "default = \"default\"\n[roots.default]\npath = \"relative\"\n",
        "root = \"/tmp\"\nother = true\n",
    ] {
        fs::write(&target, content).unwrap();
        assert!(set_root("default", home.path(), home.path(), &target).is_err());
        assert_eq!(fs::read_to_string(&target).unwrap(), content);
    }

    fs::remove_file(&target).unwrap();
    let root = home.path().join("wiki");
    let other = home.path().join("other");
    fs::create_dir(&root).unwrap();
    fs::create_dir(&other).unwrap();
    set_root("default", &root, home.path(), &target).unwrap();
    let before = fs::read(&target).unwrap();
    assert!(set_root("invalid name", &other, home.path(), &target).is_err());
    assert_eq!(fs::read(&target).unwrap(), before);
    assert!(set_root("other", &home.path().join("missing"), home.path(), &target).is_err());
    assert_eq!(fs::read(&target).unwrap(), before);
}

#[test]
fn deleted_root_and_unknown_name_are_reported() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("wiki");
    fs::create_dir(&root).unwrap();
    run(
        &args(&["set", "knowledge", root.to_str().unwrap()]),
        home.path(),
    )
    .unwrap();
    assert!(run(&args(&["path", "missing"]), home.path()).is_err());
    fs::remove_dir(&root).unwrap();
    assert!(run(&args(&["path"]), home.path()).is_err());
}

#[test]
fn symlinked_config_updates_target_and_preserves_links() {
    let home = tempfile::tempdir().unwrap();
    let original = home.path().join("original.toml");
    let intermediate = home.path().join("config-link.toml");
    let target = config(home.path());
    let root = home.path().join("wiki");
    let work = home.path().join("work");
    fs::create_dir(&root).unwrap();
    fs::create_dir(&work).unwrap();
    set_root("knowledge", &root, home.path(), &target).unwrap();
    fs::rename(&target, &original).unwrap();
    symlink("original.toml", &intermediate).unwrap();
    symlink("../config-link.toml", &target).unwrap();

    assert_eq!(
        run(&args(&["path"]), home.path()).unwrap(),
        root.canonicalize().unwrap()
    );
    assert!(run(&args(&["list"]), home.path()).is_ok());
    run(&args(&["set", "work", work.to_str().unwrap()]), home.path()).unwrap();
    run(&args(&["default", "work"]), home.path()).unwrap();
    assert_eq!(read_config(&original).unwrap().default, "work");
    assert_eq!(read_config(&original).unwrap().roots.len(), 2);
    assert_eq!(
        run(&args(&["path"]), home.path()).unwrap(),
        work.canonicalize().unwrap()
    );
    assert_eq!(
        fs::read_link(&target).unwrap(),
        Path::new("../config-link.toml")
    );
    assert_eq!(
        fs::read_link(&intermediate).unwrap(),
        Path::new("original.toml")
    );
}

#[test]
fn broken_config_symlink_is_not_replaced_or_created() {
    let home = tempfile::tempdir().unwrap();
    let target = config(home.path());
    let missing = home.path().join("missing.toml");
    fs::create_dir(target.parent().unwrap()).unwrap();
    symlink(&missing, &target).unwrap();
    assert!(
        run(
            &args(&["set", "knowledge", home.path().to_str().unwrap()]),
            home.path()
        )
        .is_err()
    );
    assert!(run(&args(&["path"]), home.path()).is_err());
    assert!(target.is_symlink());
    assert_eq!(fs::read_link(&target).unwrap(), missing);
    assert!(!missing.exists());
}

#[test]
fn install_copies_a_standalone_executable_and_preserves_collisions() {
    let home = tempfile::tempdir().unwrap();
    let source = env::current_exe().unwrap();
    let target = home.path().join(".local/bin/agent-wiki");
    install(&source, &target, false).unwrap();
    assert_eq!(fs::read(&source).unwrap(), fs::read(&target).unwrap());
    assert_eq!(
        fs::metadata(&target).unwrap().permissions().mode() & 0o777,
        0o755
    );
    install(&source, &target, false).unwrap();

    let collision = home.path().join("collision");
    fs::write(&collision, "existing command").unwrap();
    assert!(install(&source, &collision, false).is_err());
    assert_eq!(fs::read_to_string(&collision).unwrap(), "existing command");
    install(&source, &collision, true).unwrap();
    assert!(
        fs::read(&source).unwrap() == fs::read(&collision).unwrap(),
        "forced installation did not copy source bytes"
    );
    assert_eq!(
        fs::metadata(&collision).unwrap().permissions().mode() & 0o777,
        0o755
    );
}

#[test]
fn atomic_write_and_invalid_arguments_do_not_damage_files() {
    let directory = tempfile::tempdir().unwrap();
    let target = directory.path().join("existing");
    fs::create_dir(&target).unwrap();
    fs::write(target.join("keep"), "keep").unwrap();
    assert!(atomic_write(&target, b"replace", 0o600, true).is_err());
    assert_eq!(fs::read_to_string(target.join("keep")).unwrap(), "keep");

    let home = tempfile::tempdir().unwrap();
    for input in [
        args(&[]),
        args(&["set"]),
        args(&["default"]),
        args(&["path", "a", "b"]),
    ] {
        assert!(run(&input, home.path()).is_err());
    }
    assert_eq!(fs::read_dir(home.path()).unwrap().count(), 0);
}
