use crate::inspect;
use crate::inspect::{Kind, Node};
use std::ffi::OsString;
use std::fmt::Write as _;
use std::path::{Path, PathBuf};

const EDITABLE_KINDS: &str = "heading·paragraph·item·code";

/// `--id`로 고른 노드 하나를 대체하는 계획. 줄 번호는 0-based이며 양쪽 포함.
#[derive(Clone, Debug)]
pub struct Plan {
    pub node_line: String,
    pub start: usize,
    pub end: usize,
    pub old_lines: Vec<String>,
    pub new_lines: Vec<String>,
}

impl Plan {
    pub fn changed(&self) -> bool {
        self.old_lines != self.new_lines
    }
}

pub fn parse_options(args: &[OsString]) -> Result<(PathBuf, String, bool), String> {
    let mut target: Option<PathBuf> = None;
    let mut id: Option<String> = None;
    let mut write = false;
    let mut index = 0;
    while index < args.len() {
        let argument = args[index].to_string_lossy().into_owned();
        if argument == "--id" {
            if id.is_some() {
                return Err("--id는 한 번만 지정할 수 있습니다.".into());
            }
            let value = args
                .get(index + 1)
                .ok_or("--id 값을 지정하세요.")?
                .to_string_lossy()
                .trim()
                .to_string();
            if value.is_empty() {
                return Err("--id 값을 지정하세요.".into());
            }
            id = Some(value);
            index += 2;
        } else if argument == "--write" {
            if write {
                return Err("--write는 한 번만 지정할 수 있습니다.".into());
            }
            write = true;
            index += 1;
        } else if argument.starts_with("--") {
            return Err(format!("알 수 없는 옵션입니다: {argument}"));
        } else {
            if target.is_some() {
                return Err("edit에는 Markdown 파일 하나만 지정합니다.".into());
            }
            target = Some(PathBuf::from(&args[index]));
            index += 1;
        }
    }
    let target = target.ok_or("edit에는 Markdown 파일 경로를 지정합니다.")?;
    let id = id.ok_or("edit에는 --id로 대체할 노드를 지정합니다.")?;
    Ok((target, id, write))
}

pub fn plan(source: &str, path: &Path, id: &str, input: &str) -> Result<Plan, String> {
    if source.contains('\r') {
        return Err(format!(
            "CRLF 줄 바꿈 파일은 편집하지 않습니다: {}",
            path.display()
        ));
    }
    if input.trim().is_empty() {
        return Err("대체 본문이 비어 있습니다. 노드 삭제는 지원하지 않습니다.".into());
    }
    let lines: Vec<&str> = source.lines().collect();
    let document = inspect::parse_markdown(source);
    let node = inspect::find_node(&document, id)?;
    let (span_start, span_end) = node
        .span
        .ok_or("link·image 인라인 노드는 편집하지 않습니다.")?;
    if span_end >= lines.len() {
        return Err(format!(
            "문서와 구조가 맞지 않습니다: {} 줄에 id {id}가 없습니다.",
            lines.len()
        ));
    }
    // item의 span은 하위 항목·중첩 블록까지 포함하므로, 항목 자기 줄만 대체한다.
    let (start, end) = if node.kind == Kind::Item {
        (span_start, span_start)
    } else {
        (span_start, span_end)
    };
    let new_lines = replacement(node, &lines, start, end, input)?;
    Ok(Plan {
        node_line: inspect::node_line(node),
        start,
        end,
        old_lines: lines[start..=end]
            .iter()
            .map(|line| (*line).to_string())
            .collect(),
        new_lines,
    })
}

fn replacement(
    node: &Node,
    lines: &[&str],
    start: usize,
    end: usize,
    input: &str,
) -> Result<Vec<String>, String> {
    let first = lines[start];
    let prefix = container_prefix(first);
    let content = &first[prefix.len()..];
    match &node.kind {
        Kind::Heading(level) => {
            let text = single_line(input, "heading")?;
            let marker = if content.trim_start().starts_with('#') {
                let run = content
                    .trim_start()
                    .chars()
                    .take_while(|c| *c == '#')
                    .count();
                format!("{} ", "#".repeat(run))
            } else {
                format!("{} ", "#".repeat(usize::from(*level)))
            };
            Ok(vec![format!("{prefix}{marker}{text}")])
        }
        Kind::Paragraph => {
            let mut replaced = Vec::new();
            for line in input.lines() {
                if line.trim().is_empty() {
                    return Err("paragraph에는 빈 줄을 넣을 수 없습니다.".into());
                }
                replaced.push(format!("{prefix}{}", line.trim_start()));
            }
            Ok(replaced)
        }
        Kind::Item => {
            let text = single_line(input, "item")?;
            let Some((_, offset, _, _)) = inspect::list_marker(content) else {
                return Err("list 항목의 시작 줄을 인식할 수 없습니다.".into());
            };
            let base = prefix.len();
            let mut cut = base + offset;
            if let Some(stop) = checkbox_end(&content[offset..]) {
                cut = base + offset + stop;
            }
            Ok(vec![format!("{}{text}", &first[..cut])])
        }
        Kind::Code => {
            let Some((open_fence, _, _info)) = inspect::fence_marker(content) else {
                return Err("code block의 시작 fence를 인식할 수 없습니다.".into());
            };
            let close_prefix = container_prefix(lines[end]);
            let closed = inspect::fence_marker(&lines[end][close_prefix.len()..])
                .is_some_and(|(c, _, rest)| c == open_fence && rest.is_empty() && end > start);
            if !closed {
                return Err(format!(
                    "닫지 않은 code block은 편집하지 않습니다. id {}를 다시 확인하세요",
                    node.id
                ));
            }
            let mut replaced = vec![first.to_string()];
            for line in input.lines() {
                replaced.push(if line.trim().is_empty() {
                    String::new()
                } else {
                    format!("{prefix}{}", line.trim_start())
                });
            }
            replaced.push(lines[end].to_string());
            Ok(replaced)
        }
        kind => Err(format!(
            "{}는 편집하지 않습니다. 편집할 수 있는 kind: {EDITABLE_KINDS}",
            inspect::kind_name(kind)
        )),
    }
}

pub fn render_plan(plan: &Plan, path: &Path, saved: bool) -> String {
    let mut output = format!("edit {}\n{}\n", path.display(), plan.node_line);
    if !plan.changed() {
        output.push_str("변경 없음: 이미 같은 내용입니다.\n");
        return output;
    }
    for line in &plan.old_lines {
        let _ = writeln!(output, "- {line}");
    }
    for line in &plan.new_lines {
        let _ = writeln!(output, "+ {line}");
    }
    if saved {
        let _ = writeln!(output, "저장했습니다: {}", path.display());
    } else {
        output.push_str("dry-run: 파일을 변경하지 않았습니다. 저장하려면 --write 를 사용하세요.\n");
    }
    output
}

pub fn apply(source: &str, plan: &Plan) -> String {
    let lines: Vec<&str> = source.lines().collect();
    let mut parts: Vec<String> = Vec::new();
    parts.extend(lines[..plan.start].iter().map(|line| (*line).to_string()));
    parts.extend(plan.new_lines.clone());
    parts.extend(lines[plan.end + 1..].iter().map(|line| (*line).to_string()));
    let mut content = parts.join("\n");
    if source.ends_with('\n') {
        content.push('\n');
    }
    content
}

fn single_line(input: &str, kind: &str) -> Result<String, String> {
    let text = input.trim();
    if text.lines().count() > 1 {
        return Err(format!("{kind}에는 한 줄만 넣을 수 있습니다."));
    }
    Ok(text.to_string())
}

/// 줄 앞쪽의 들여쓰기와 blockquote 마커만 돌려줍니다. 대체 줄에 다시 붙입니다.
fn container_prefix(line: &str) -> String {
    let mut prefix = String::new();
    let mut rest = line;
    while let Some(stripped) = rest.strip_prefix(' ') {
        prefix.push(' ');
        rest = stripped;
    }
    while let Some(stripped) = rest.strip_prefix('>') {
        prefix.push('>');
        rest = stripped;
        if let Some(spaced) = stripped.strip_prefix(' ') {
            prefix.push(' ');
            rest = spaced;
        }
    }
    prefix
}

/// 항목 본문 앞의 task checkbox(`[ ]`, `[x]`, `[X]`) 끝 byte 위치. 뒤의 공백도 포함합니다.
fn checkbox_end(text: &str) -> Option<usize> {
    let bytes = text.as_bytes();
    let mut index = 0;
    while bytes.get(index) == Some(&b' ') {
        index += 1;
    }
    if bytes.get(index) != Some(&b'[') {
        return None;
    }
    if !matches!(bytes.get(index + 1), Some(b' ' | b'x' | b'X'))
        || bytes.get(index + 2) != Some(&b']')
    {
        return None;
    }
    let mut stop = index + 3;
    while bytes.get(stop) == Some(&b' ') {
        stop += 1;
    }
    Some(stop)
}
