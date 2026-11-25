use super::edit;
use std::path::Path;

const DOC: &str = "---\ntitle: note\n---\n\n# 제목\n\n본문 첫 줄\n이어지는 줄\n\n- 항목\n  - 중첩 항목\n\n> 인용 첫 줄\n> 인용 둘째 줄\n";

fn plan(source: &str, id: &str, input: &str) -> edit::Plan {
    edit::plan(source, Path::new("note.md"), id, input).expect("계획 생성")
}

#[test]
fn heading_keeps_marker_and_replaces_text() {
    let made = plan(DOC, "2", "새 제목\n");
    assert_eq!(made.start, 4);
    assert_eq!(made.end, 4);
    assert_eq!(made.new_lines, vec!["# 새 제목".to_string()]);
    assert!(made.changed());
}

#[test]
fn setext_heading_becomes_atx() {
    let source = "제목\n===\n\n본문\n";
    let made = plan(source, "1", "새 제목\n");
    assert_eq!(made.start, 0);
    assert_eq!(made.end, 1);
    assert_eq!(made.new_lines, vec!["# 새 제목".to_string()]);
}

#[test]
fn paragraph_replaces_every_spanned_line() {
    let made = plan(DOC, "2.1", "새 본문\n둘째 줄\n");
    assert_eq!(made.start, 6);
    assert_eq!(made.end, 7);
    assert_eq!(
        made.new_lines,
        vec!["새 본문".to_string(), "둘째 줄".to_string()]
    );
}

#[test]
fn blockquote_paragraph_keeps_quote_prefix() {
    let made = plan(DOC, "2.3.1", "새 인용\n");
    assert_eq!(made.start, 12);
    assert_eq!(made.end, 13);
    assert_eq!(made.new_lines, vec!["> 새 인용".to_string()]);
}

#[test]
fn blockquote_heading_item_and_code_keep_quote_prefix() {
    let source = "> # 인용 제목\n>\n> - 인용 항목\n>\n> ```sh\n> echo hi\n> ```\n";
    let heading = plan(source, "1.1", "새 인용 제목\n");
    assert_eq!(heading.new_lines, vec!["> # 새 인용 제목".to_string()]);
    let item = plan(source, "1.2.1", "새 인용 항목\n");
    assert_eq!(item.new_lines, vec!["> - 새 인용 항목".to_string()]);
    let code = plan(source, "1.3", "echo bye\n");
    assert_eq!(
        code.new_lines,
        vec![
            "> ```sh".to_string(),
            "> echo bye".to_string(),
            "> ```".to_string()
        ]
    );
}

#[test]
fn item_keeps_marker_and_task_checkbox() {
    let source = "- [x] 완료\n- [ ] 미완료\n";
    let done = plan(source, "1.1", "고친 완료\n");
    assert_eq!(done.new_lines, vec!["- [x] 고친 완료".to_string()]);
    let open = plan(source, "1.2", "고친 미완료\n");
    assert_eq!(open.new_lines, vec!["- [ ] 고친 미완료".to_string()]);
}

#[test]
fn item_replaces_only_its_own_line_and_keeps_children() {
    let source = "- 부모 항목\n  - 자식 항목\n    - 손주 항목\n";
    let made = plan(source, "1.1", "고친 부모\n");
    assert_eq!(made.start, 0);
    assert_eq!(made.end, 0);
    assert_eq!(made.new_lines, vec!["- 고친 부모".to_string()]);
    assert_eq!(
        edit::apply(source, &made),
        "- 고친 부모\n  - 자식 항목\n    - 손주 항목\n"
    );
}

#[test]
fn nested_item_span_uses_original_line() {
    let made = plan(DOC, "2.2.1.1.1", "고친 중첩\n");
    assert_eq!(made.start, 10);
    assert_eq!(made.end, 10);
    assert_eq!(made.new_lines, vec!["  - 고친 중첩".to_string()]);
}

#[test]
fn code_body_keeps_fences_and_info() {
    let source = "```sh\nls\n```\n";
    let made = plan(source, "1", "ls -la\npwd\n");
    assert_eq!(
        made.new_lines,
        vec![
            "```sh".to_string(),
            "ls -la".to_string(),
            "pwd".to_string(),
            "```".to_string()
        ]
    );
}

#[test]
fn apply_rewrites_only_the_selected_node() {
    let made = plan(DOC, "2.1", "새 본문\n");
    let after = edit::apply(DOC, &made);
    assert_eq!(
        after,
        "---\ntitle: note\n---\n\n# 제목\n\n새 본문\n\n- 항목\n  - 중첩 항목\n\n> 인용 첫 줄\n> 인용 둘째 줄\n"
    );
}

#[test]
fn apply_keeps_missing_trailing_newline() {
    let source = "본문";
    let made = plan(source, "1", "새 본문");
    assert_eq!(edit::apply(source, &made), "새 본문");
}

#[test]
fn same_content_reports_no_change() {
    let made = plan(DOC, "2.1", "본문 첫 줄\n이어지는 줄\n");
    assert!(!made.changed());
    let rendered = edit::render_plan(&made, Path::new("note.md"), false);
    assert!(rendered.contains("변경 없음"));
}

#[test]
fn dry_run_output_shows_diff_and_hint() {
    let made = plan(DOC, "2", "새 제목\n");
    let rendered = edit::render_plan(&made, Path::new("note.md"), false);
    assert_eq!(
        rendered,
        "edit note.md\n2 heading 1 \"제목\" #제목\n- # 제목\n+ # 새 제목\ndry-run: 파일을 변경하지 않았습니다. 저장하려면 --write 를 사용하세요.\n"
    );
}

#[test]
fn non_editable_kinds_are_rejected() {
    let source = "| 가 | 나 |\n| --- | --- |\n| 1 | 2 |\n";
    let error = edit::plan(source, Path::new("note.md"), "1", "본문\n").unwrap_err();
    assert!(error.contains("table는 편집하지 않습니다"), "{error}");
    let link = edit::plan(
        "본문 [링크](./a.md)\n",
        Path::new("note.md"),
        "1.1",
        "본문\n",
    )
    .unwrap_err();
    assert!(link.contains("인라인"), "{link}");
}

#[test]
fn missing_and_ambiguous_ids_are_rejected() {
    let missing = edit::plan(DOC, Path::new("note.md"), "9", "본문\n").unwrap_err();
    assert!(missing.contains("id를 찾을 수 없습니다"), "{missing}");
    let ambiguous = edit::plan("# A\n\n## A\n", Path::new("note.md"), "a", "본문\n").unwrap_err();
    assert!(ambiguous.contains("자리 번호를 사용하세요"), "{ambiguous}");
}

#[test]
fn invalid_input_is_rejected() {
    let empty = edit::plan(DOC, Path::new("note.md"), "2.1", "  \n").unwrap_err();
    assert!(empty.contains("대체 본문이 비어 있습니다"), "{empty}");
    let crlf = edit::plan(
        "# 제목\r\n\r\n본문\r\n",
        Path::new("note.md"),
        "1",
        "새 제목\n",
    )
    .unwrap_err();
    assert!(crlf.contains("CRLF"), "{crlf}");
    let multiline = edit::plan(DOC, Path::new("note.md"), "2", "두\n줄\n").unwrap_err();
    assert!(multiline.contains("heading에는 한 줄만"), "{multiline}");
    let blank = edit::plan(DOC, Path::new("note.md"), "2.1", "첫\n\n둘\n").unwrap_err();
    assert!(blank.contains("paragraph에는 빈 줄"), "{blank}");
    let unclosed = edit::plan("```sh\nls\n", Path::new("note.md"), "1", "pwd\n").unwrap_err();
    assert!(unclosed.contains("닫지 않은 code block"), "{unclosed}");
}

#[test]
fn options_require_id_and_accept_write() {
    let args = |values: &[&str]| -> Vec<std::ffi::OsString> {
        values.iter().map(std::ffi::OsString::from).collect()
    };
    let (target, id, write) =
        edit::parse_options(&args(&["note.md", "--id", "2.1", "--write"])).expect("옵션 해석");
    assert_eq!(target, std::path::PathBuf::from("note.md"));
    assert_eq!(id, "2.1");
    assert!(write);
    assert!(edit::parse_options(&args(&["note.md"])).is_err());
    assert!(edit::parse_options(&args(&["note.md", "--id", "1", "--out"])).is_err());
    assert!(edit::parse_options(&args(&["a.md", "b.md", "--id", "1"])).is_err());
}
