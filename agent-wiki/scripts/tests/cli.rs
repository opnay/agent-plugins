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
fn legacy_config_is_read_and_normalized_on_successful_set() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("legacy");
    fs::create_dir(&root).unwrap();
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    fs::write(&target, format!("root = {:?}\n", root)).unwrap();

    assert_eq!(read_config(&target).unwrap().roots["default"].path, root);
    let replacement = home.path().join("replacement");
    fs::create_dir(&replacement).unwrap();
    assert_eq!(
        run(&args(&["set", replacement.to_str().unwrap()]), home.path()).unwrap(),
        replacement.canonicalize().unwrap()
    );

    let parsed = read_config(&target).unwrap();
    assert_eq!(parsed.default, "default");
    assert_eq!(parsed.roots.len(), 1);
    assert_eq!(
        parsed.roots["default"].path,
        replacement.canonicalize().unwrap()
    );
    assert!(fs::read_to_string(target).unwrap().contains("version = 2"));
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
    let parsed = read_config(&config(home.path())).unwrap();
    assert_eq!(
        format_root_list(&parsed).unwrap(),
        format!(
            "knowledge\t{}\tdefault=false\tdescription=-\npersonality\t{}\tdefault=true\tdescription=-\n",
            knowledge.canonicalize().unwrap().display(),
            personality.canonicalize().unwrap().display()
        )
    );
}

#[test]
fn v2_config_preserves_optional_description() {
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
            "version = 2\ndefault = \"knowledge\"\n\n[roots.knowledge]\npath = {:?}\ndescription = \"shared\"\n\n[roots.work]\npath = {:?}\n",
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
fn query_commands_reject_duplicate_v2_roots_without_overwriting_config() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("wiki");
    fs::create_dir(&root).unwrap();
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    fs::write(
        &target,
        format!(
            "version = 2\ndefault = \"knowledge\"\n\n[roots.knowledge]\npath = {:?}\n\n[roots.personality]\npath = {:?}\n",
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
        "version = 3\ndefault = \"default\"\n[roots.default]\npath = \"/tmp\"\n",
        "version = 2\ndefault = \"default\"\n[roots.default]\npath = \"relative\"\n",
        "root = \"/tmp\"\nother = true\n",
    ] {
        fs::write(&target, content).unwrap();
        assert!(set_root("default", home.path(), home.path(), &target).is_err());
        assert_eq!(fs::read_to_string(&target).unwrap(), content);
    }

    fs::remove_file(&target).unwrap();
    set_root("default", home.path(), home.path(), &target).unwrap();
    let before = fs::read(&target).unwrap();
    assert!(set_root("invalid name", home.path(), home.path(), &target).is_err());
    assert!(set_root("other", &home.path().join("missing"), home.path(), &target).is_err());
    assert_eq!(fs::read(&target).unwrap(), before);
}

#[test]
fn deleted_root_and_unknown_name_are_reported() {
    let home = tempfile::tempdir().unwrap();
    let root = home.path().join("wiki");
    fs::create_dir(&root).unwrap();
    run(&args(&["set", root.to_str().unwrap()]), home.path()).unwrap();
    assert!(run(&args(&["path", "missing"]), home.path()).is_err());
    fs::remove_dir(&root).unwrap();
    assert!(run(&args(&["path"]), home.path()).is_err());
}

#[test]
fn symlinked_config_is_not_overwritten() {
    let home = tempfile::tempdir().unwrap();
    let original = home.path().join("original.toml");
    let target = config(home.path());
    fs::create_dir(target.parent().unwrap()).unwrap();
    fs::write(&original, "root = \"/tmp\"\n").unwrap();
    symlink(&original, &target).unwrap();
    assert!(set_root("default", home.path(), home.path(), &target).is_err());
    assert_eq!(fs::read_to_string(&original).unwrap(), "root = \"/tmp\"\n");
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
