use super::*;

fn args(values: &[&str]) -> Vec<OsString> {
    values.iter().map(OsString::from).collect()
}

fn tree(source: &str) -> String {
    let document = inspect::parse_markdown(source);
    inspect::render(&document, Path::new("doc.md"), None).unwrap()
}

fn subtree(source: &str, id: &str) -> String {
    let document = inspect::parse_markdown(source);
    inspect::render(&document, Path::new("doc.md"), Some(id)).unwrap()
}

#[test]
fn front_matter_and_setext_headings_become_nodes() {
    let output = tree(
        "---\ntype: record\ntags:\n  - alpha\n---\n\nSetext H2\n--------\n\nSetext H1\n====\n",
    );
    assert!(output.starts_with("document doc.md nodes=4 flags=0"));
    assert!(output.contains("1 front matter keys=2"));
    assert!(output.contains("2 heading 2 \"Setext H2\" #setext-h2"));
    assert!(output.contains("3 heading 1 \"Setext H1\" #setext-h1"));
}

#[test]
fn heading_sections_nest_until_the_next_heading() {
    let output = tree("# A\n\n## B\n\n### C\n\n# D\n");
    assert!(output.contains("1 heading 1 \"A\" #a"));
    assert!(output.contains("  1.1 heading 2 \"B\" #b"));
    assert!(output.contains("    1.1.1 heading 3 \"C\" #c"));
    assert!(output.contains("2 heading 1 \"D\" #d"));
    assert!(tree("# A\n\n### C\n").contains("flag=heading-skip"));
    assert!(!tree("# A\n\n## B\n").contains("flag=heading-skip"));
}

#[test]
fn lists_split_by_marker_and_report_tasks() {
    let output = tree("- [ ] open\n- [x] done\n  - nested\n\n1. one\n2. two\n");
    assert!(output.contains("1 list unordered items=2"));
    assert!(output.contains("1.1 item task=open \"open\""));
    assert!(output.contains("1.2 item task=done \"done\""));
    assert!(output.contains("1.2.1 list unordered items=1"));
    assert!(output.contains("2 list ordered items=2"));
    assert!(output.contains("2.1 item \"one\""));
}

#[test]
fn fences_and_blockquotes_keep_nested_blocks() {
    let output = tree("> quote\n>\n> - item\n\n```rust\nfn main() {}\n```\n");
    assert!(output.contains("1 blockquote blocks=2 \"quote\""));
    assert!(output.contains("  1.1 paragraph \"quote\""));
    assert!(output.contains("  1.2 list unordered items=1"));
    assert!(output.contains("    1.2.1 item \"item\""));
    assert!(output.contains("2 code lines=1 info=rust \"fn main() {}\""));
    assert!(tree("~~~\ntilde\n~~~\n").contains("code lines=1 \"tilde\""));
}

#[test]
fn links_images_and_wikilinks_become_nodes() {
    let output =
        tree("[text](./a.md)\n\n![alt](b.png)\n\n![](c.png)\n\n[[Note]]\n\n![[Note.png]]\n");
    assert!(output.contains("link dest=./a.md \"text\""));
    assert!(output.contains("image dest=b.png \"alt\""));
    assert!(output.contains("image dest=c.png flag=missing-alt"));
    assert!(output.contains("link target=Note \"Note\""));
    assert!(output.contains("image target=Note.png \"Note.png\""));
    assert!(tree("[](  )\n").contains("link dest= flag=empty-link-text"));
}

#[test]
fn structure_flags_are_reported_and_counted() {
    let output = tree("# A\n\n# A\n\n| a | b |\n| - | - |\n| 1 |\n");
    assert!(output.starts_with("document doc.md nodes=4 flags=2"));
    assert!(output.contains("1 heading 1 \"A\" #a"));
    assert!(output.contains("2 heading 1 \"A\" #a flag=duplicate-slug"));
    assert!(output.contains("2.1 table rows=1 cols=2 flag=table-column-mismatch"));
    assert!(tree("| a | b |\n| - | - |\n| 1 | 2 |\n").contains("table rows=1 cols=2\n"));
}

#[test]
fn code_spans_count_on_the_parent_node() {
    let output = tree("# `A`\n\nwith `one` and `two`\n");
    assert!(output.contains("paragraph code=2 \"with `one` and `two`\""));
    assert!(output.contains("heading 1 code=1 \"`A`\" #a"));
    assert!(!output.contains("code span"));
}

#[test]
fn id_selects_positional_number_and_slug() {
    let source = "# A\n\n## B\n\ntext\n";
    assert!(subtree(source, "1.1").starts_with("subtree 1.1 nodes=2 flags=0"));
    assert!(subtree(source, "b").starts_with("subtree 1.1 nodes=2 flags=0"));
    assert!(subtree(source, "#b").starts_with("subtree 1.1 nodes=2 flags=0"));
    assert!(subtree(source, "1.1.1").contains("paragraph \"text\""));
}

#[test]
fn ambiguous_and_missing_ids_fail() {
    let document = inspect::parse_markdown("# A\n\n# A\n");
    assert_eq!(
        inspect::render(&document, Path::new("doc.md"), Some("a")).unwrap_err(),
        "id가 2개 노드와 일치합니다. 자리 번호를 사용하세요: 1 2"
    );
    let document = inspect::parse_markdown("# A\n\n## A\n");
    assert_eq!(
        inspect::render(&document, Path::new("doc.md"), Some("a")).unwrap_err(),
        "id가 2개 노드와 일치합니다. 자리 번호를 사용하세요: 1 1.1"
    );
    let document = inspect::parse_markdown("# A\n");
    assert_eq!(
        inspect::render(&document, Path::new("doc.md"), Some("zz")).unwrap_err(),
        "id를 찾을 수 없습니다: zz"
    );
}

#[test]
fn indented_lines_start_blocks_but_never_code() {
    let output = tree("    # indented\n\n    ---\n");
    assert!(output.contains("1 heading 1 \"indented\" #indented"));
    assert!(output.contains("  1.1 thematic break"));
    assert!(!tree("    plain text\n").contains("code"));
}

#[test]
fn inspect_cli_reads_a_file_and_rejects_bad_arguments() {
    let home = tempfile::tempdir().unwrap();
    let file = home.path().join("doc.md");
    fs::write(&file, "# A\n\ntext\n").unwrap();
    let path = file.to_str().unwrap();
    assert!(run(&args(&["inspect", path]), home.path()).is_ok());

    let missing = home.path().join("missing.md");
    assert!(run(&args(&["inspect", missing.to_str().unwrap()]), home.path()).is_err());
    assert!(!missing.exists());

    for options in [
        args(&["inspect"]),
        args(&["inspect", path, "--id"]),
        args(&["inspect", path, "--id", "a", "--id", "b"]),
        args(&["inspect", path, "--outline"]),
        args(&["inspect", path, path]),
    ] {
        assert!(run(&options, home.path()).is_err(), "{options:?}");
    }

    let binary = home.path().join("binary.md");
    fs::write(&binary, [0xff_u8, 0xfe, b'a']).unwrap();
    assert!(run(&args(&["inspect", binary.to_str().unwrap()]), home.path()).is_err());
}
