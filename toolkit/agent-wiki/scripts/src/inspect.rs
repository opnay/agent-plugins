use std::collections::BTreeSet;
use std::ffi::OsString;
use std::fmt::Write as _;
use std::path::{Path, PathBuf};

const PREVIEW_CHARS: usize = 60;

#[derive(Clone, Debug, PartialEq, Eq)]
pub enum Kind {
    Document,
    FrontMatter,
    Heading(u8),
    Paragraph,
    List { ordered: bool },
    Item,
    Blockquote,
    Code,
    Table,
    ThematicBreak,
    Html,
    Link,
    Image,
}

#[derive(Clone, Debug)]
pub struct Node {
    pub kind: Kind,
    pub details: String,
    pub text: String,
    pub slug: Option<String>,
    pub flags: Vec<String>,
    pub children: Vec<Node>,
    pub id: String,
    /// 원본 문서의 0-based 시작·끝 줄 번호(양쪽 포함). 인라인 노드는 None.
    pub span: Option<(usize, usize)>,
}

fn node(kind: Kind) -> Node {
    Node {
        kind,
        details: String::new(),
        text: String::new(),
        slug: None,
        flags: Vec::new(),
        children: Vec::new(),
        id: String::new(),
        span: None,
    }
}

fn preview(text: &str) -> String {
    let first = text.trim().lines().next().unwrap_or_default().trim();
    let chars: Vec<char> = first.chars().collect();
    if chars.len() > PREVIEW_CHARS {
        let mut shortened: String = chars[..PREVIEW_CHARS].iter().collect();
        shortened.push('…');
        shortened
    } else {
        first.to_string()
    }
}

pub fn slugify(text: &str) -> String {
    let mut slug = String::new();
    let mut pending_dash = false;
    for character in text.chars() {
        if character.is_alphanumeric() || character == '-' || character == '_' {
            if pending_dash && !slug.is_empty() {
                slug.push('-');
            }
            pending_dash = false;
            slug.extend(character.to_lowercase());
        } else {
            pending_dash = true;
        }
    }
    slug
}

pub fn parse_options(args: &[OsString]) -> Result<(PathBuf, Option<String>), String> {
    let mut target: Option<PathBuf> = None;
    let mut id: Option<String> = None;
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
        } else if argument.starts_with("--") {
            return Err(format!("알 수 없는 옵션입니다: {argument}"));
        } else {
            if target.is_some() {
                return Err("inspect에는 Markdown 파일 하나만 지정합니다.".into());
            }
            target = Some(PathBuf::from(&args[index]));
            index += 1;
        }
    }
    let target = target.ok_or("inspect에는 Markdown 파일 경로를 지정합니다.")?;
    Ok((target, id))
}

struct Parser<'a> {
    lines: Vec<&'a str>,
    index: usize,
    /// `lines[0]`이 원본 문서의 몇 번째 줄인지 가리키는 오프셋.
    offset: usize,
}

pub(crate) fn fence_marker(line: &str) -> Option<(char, usize, String)> {
    let trimmed = line.trim_start();
    let fence = trimmed.chars().next()?;
    if fence != '`' && fence != '~' {
        return None;
    }
    let run = trimmed.chars().take_while(|c| *c == fence).count();
    if run < 3 {
        return None;
    }
    Some((fence, run, trimmed[run..].trim().to_string()))
}

fn atx_heading(line: &str) -> Option<(u8, String)> {
    let trimmed = line.trim_start();
    let level = trimmed.chars().take_while(|c| *c == '#').count();
    if !(1..=6).contains(&level) {
        return None;
    }
    let rest = &trimmed[level..];
    if !rest.is_empty() && !rest.starts_with(' ') && !rest.starts_with('\t') {
        return None;
    }
    let mut text = rest.trim().to_string();
    let hashes = text.chars().rev().take_while(|c| *c == '#').count();
    if hashes > 0 {
        let before = &text[..text.len() - hashes];
        if before.is_empty() || before.ends_with(' ') || before.ends_with('\t') {
            text = before.trim_end().to_string();
        }
    }
    Some((level as u8, text))
}

fn is_thematic_break(line: &str) -> bool {
    let trimmed = line.trim();
    let Some(first) = trimmed.chars().next() else {
        return false;
    };
    if !matches!(first, '-' | '*' | '_') {
        return false;
    }
    trimmed.chars().filter(|c| *c == first).count() >= 3
        && trimmed.chars().all(|c| c == first || c == ' ')
}

pub(crate) fn list_marker(line: &str) -> Option<(usize, usize, bool, char)> {
    let indent = line.len() - line.trim_start().len();
    if indent >= 4 {
        return None;
    }
    let rest = line.trim_start();
    if let Some(after) = rest.strip_prefix(['-', '*', '+']) {
        if after.is_empty() || after.starts_with(' ') || after.starts_with('\t') {
            let marker = rest.chars().next().unwrap_or('-');
            return Some((indent, line.len() - after.trim_start().len(), false, marker));
        }
        return None;
    }
    let digits = rest.chars().take_while(|c| c.is_ascii_digit()).count();
    if (1..=9).contains(&digits)
        && let Some(after) = rest[digits..].strip_prefix(['.', ')'])
        && (after.is_empty() || after.starts_with(' ') || after.starts_with('\t'))
    {
        let marker = rest[digits..].chars().next().unwrap_or('.');
        return Some((indent, line.len() - after.trim_start().len(), true, marker));
    }
    None
}

fn setext_level(line: &str) -> Option<u8> {
    if line.len() - line.trim_start().len() >= 4 {
        return None;
    }
    let trimmed = line.trim();
    if trimmed.is_empty() {
        return None;
    }
    if trimmed.chars().all(|c| c == '=') {
        Some(1)
    } else if trimmed.chars().all(|c| c == '-') {
        Some(2)
    } else {
        None
    }
}

fn is_table_start(lines: &[&str], index: usize) -> bool {
    let Some(header) = lines.get(index) else {
        return false;
    };
    header.contains('|')
        && lines
            .get(index + 1)
            .is_some_and(|row| is_delimiter_row(row))
}

fn is_delimiter_row(line: &str) -> bool {
    let trimmed = line.trim();
    trimmed.contains('-')
        && trimmed.contains('|')
        && trimmed.chars().all(|c| matches!(c, '-' | ':' | '|' | ' '))
}

fn cell_count(line: &str) -> usize {
    let trimmed = line.trim().trim_start_matches('|').trim_end_matches('|');
    let mut cells = 1;
    let mut escaped = false;
    for character in trimmed.chars() {
        if escaped {
            escaped = false;
        } else if character == '\\' {
            escaped = true;
        } else if character == '|' {
            cells += 1;
        }
    }
    cells
}

fn is_html_block(line: &str) -> bool {
    let trimmed = line.trim_start();
    let mut characters = trimmed.chars();
    if characters.next() != Some('<') {
        return false;
    }
    match characters.next() {
        Some('/') | Some('!') => true,
        Some(character) => character.is_ascii_alphabetic(),
        None => false,
    }
}

fn starts_block(lines: &[&str], index: usize) -> bool {
    let Some(line) = lines.get(index) else {
        return false;
    };
    fence_marker(line).is_some()
        || atx_heading(line).is_some()
        || is_thematic_break(line)
        || line.trim_start().starts_with('>')
        || list_marker(line).is_some()
        || is_table_start(lines, index)
        || is_html_block(line)
}

fn task_checkbox(text: &mut String) -> Option<bool> {
    let trimmed = text.trim_start();
    let rest = trimmed.strip_prefix('[')?;
    let mark = rest.as_bytes().first().copied()?;
    if !matches!(mark, b' ' | b'x' | b'X') || rest.as_bytes().get(1) != Some(&b']') {
        return None;
    }
    *text = rest[2..].trim_start().to_string();
    Some(mark == b'x' || mark == b'X')
}

impl<'a> Parser<'a> {
    fn new(lines: Vec<&'a str>, offset: usize) -> Self {
        Self {
            lines,
            index: 0,
            offset,
        }
    }

    fn parse(&mut self) -> Vec<Node> {
        let mut blocks = Vec::new();
        while self.index < self.lines.len() {
            if self.lines[self.index].trim().is_empty() {
                self.index += 1;
                continue;
            }
            blocks.push(self.block());
        }
        blocks
    }

    fn block(&mut self) -> Node {
        let start = self.index;
        let mut block = self.block_raw();
        block.span = Some((
            start + self.offset,
            self.index.saturating_sub(1) + self.offset,
        ));
        block
    }

    fn block_raw(&mut self) -> Node {
        let line = self.lines[self.index];
        if let Some((fence, run, info)) = fence_marker(line) {
            return self.code_block(fence, run, info);
        }
        if let Some((level, text)) = atx_heading(line) {
            self.index += 1;
            let mut heading = node(Kind::Heading(level));
            heading.text = preview(&text);
            heading.slug = Some(slugify(&text));
            let (inline, code_spans) = inline_nodes(&text);
            heading.children = inline;
            append_code(&mut heading.details, code_spans);
            return heading;
        }
        if is_thematic_break(line) {
            self.index += 1;
            return node(Kind::ThematicBreak);
        }
        if line.trim_start().starts_with('>') {
            return self.blockquote();
        }
        if list_marker(line).is_some() {
            return self.list();
        }
        if is_table_start(&self.lines, self.index) {
            return self.table();
        }
        if is_html_block(line) {
            return self.html_block();
        }
        self.paragraph()
    }

    fn code_block(&mut self, fence: char, run: usize, info: String) -> Node {
        let opened = self.index;
        self.index += 1;
        let mut body: Vec<&str> = Vec::new();
        while self.index < self.lines.len() {
            let line = self.lines[self.index];
            if fence_marker(line)
                .is_some_and(|(c, length, rest)| c == fence && length >= run && rest.is_empty())
            {
                break;
            }
            body.push(line);
            self.index += 1;
        }
        if self.index >= self.lines.len() {
            self.index = opened + body.len() + 1;
        } else {
            self.index += 1;
        }
        let mut code = node(Kind::Code);
        code.details = if info.is_empty() {
            format!("lines={}", body.len())
        } else {
            format!("lines={} info={info}", body.len())
        };
        code.text = body.first().map(|line| preview(line)).unwrap_or_default();
        code
    }

    fn blockquote(&mut self) -> Node {
        let start = self.index;
        let mut inner: Vec<&str> = Vec::new();
        while self.index < self.lines.len() {
            let line = self.lines[self.index];
            let Some(rest) = line.trim_start().strip_prefix('>') else {
                break;
            };
            inner.push(rest.strip_prefix([' ', '\t']).unwrap_or(rest));
            self.index += 1;
        }
        let blocks = Parser::new(inner, self.offset + start).parse();
        let mut quote = node(Kind::Blockquote);
        quote.details = format!("blocks={}", blocks.len());
        quote.text = blocks
            .first()
            .map(|block| block.text.clone())
            .unwrap_or_default();
        quote.children = blocks;
        quote
    }

    fn table(&mut self) -> Node {
        let columns = cell_count(self.lines[self.index + 1]);
        let mut rows: usize = 0;
        let mut mismatch = false;
        while self.index < self.lines.len() {
            let line = self.lines[self.index];
            if line.trim().is_empty() || !line.contains('|') {
                break;
            }
            if cell_count(line) != columns {
                mismatch = true;
            }
            rows += 1;
            self.index += 1;
        }
        let mut table = node(Kind::Table);
        table.details = format!("rows={} cols={columns}", rows.saturating_sub(2));
        if mismatch {
            table.flags.push("table-column-mismatch".into());
        }
        table
    }

    fn html_block(&mut self) -> Node {
        let mut lines = 0;
        while self.index < self.lines.len() && !self.lines[self.index].trim().is_empty() {
            lines += 1;
            self.index += 1;
        }
        let mut html = node(Kind::Html);
        html.details = format!("lines={lines}");
        html
    }

    fn list(&mut self) -> Node {
        let Some((base_indent, content_offset, ordered, delimiter)) =
            list_marker(self.lines[self.index])
        else {
            self.index += 1;
            return node(Kind::List { ordered: false });
        };
        let mut list = node(Kind::List { ordered });
        while self.index < self.lines.len() {
            let line = self.lines[self.index];
            if line.trim().is_empty() {
                let mut peek = self.index + 1;
                while peek < self.lines.len() && self.lines[peek].trim().is_empty() {
                    peek += 1;
                }
                if peek < self.lines.len()
                    && list_marker(self.lines[peek])
                        .is_some_and(|(_, _, _, marker)| marker == delimiter)
                {
                    self.index = peek;
                    continue;
                }
                break;
            }
            let Some((indent, offset, _, marker)) = list_marker(line) else {
                break;
            };
            if indent < base_indent || marker != delimiter {
                break;
            }
            let mut item = node(Kind::Item);
            let item_start = self.index;
            let mut text = line[offset..].to_string();
            if let Some(done) = task_checkbox(&mut text) {
                item.details = if done { "task=done" } else { "task=open" }.into();
            }
            item.text = preview(&text);
            let (inline, code_spans) = inline_nodes(&text);
            item.children = inline;
            append_code(&mut item.details, code_spans);
            self.index += 1;
            let inner_start = self.index;
            let mut inner: Vec<&str> = Vec::new();
            while self.index < self.lines.len() {
                let next = self.lines[self.index];
                if next.trim().is_empty() {
                    break;
                }
                if next.len() - next.trim_start().len() > base_indent {
                    inner.push(next);
                    self.index += 1;
                } else {
                    break;
                }
            }
            item.span = Some((
                item_start + self.offset,
                self.index.saturating_sub(1) + self.offset,
            ));
            if !inner.is_empty() {
                let dedented: Vec<&str> = inner
                    .iter()
                    .map(|line| &line[content_offset.min(line.len() - line.trim_start().len())..])
                    .collect();
                item.children
                    .extend(Parser::new(dedented, self.offset + inner_start).parse());
            }
            list.children.push(item);
        }
        list.details = format!("items={}", list.children.len());
        list
    }

    fn paragraph(&mut self) -> Node {
        let start = self.index;
        let mut collected: Vec<&str> = Vec::new();
        while self.index < self.lines.len() {
            let line = self.lines[self.index];
            if line.trim().is_empty() {
                break;
            }
            if self.index > start
                && (starts_block(&self.lines, self.index) || setext_level(line).is_some())
            {
                break;
            }
            collected.push(line);
            self.index += 1;
        }
        let text = collected
            .iter()
            .map(|line| line.trim())
            .collect::<Vec<_>>()
            .join(" ");
        if self.index < self.lines.len() {
            let level = setext_level(self.lines[self.index]);
            if let Some(level) = level {
                self.index += 1;
                let mut heading = node(Kind::Heading(level));
                heading.text = preview(&text);
                heading.slug = Some(slugify(&text));
                let (inline, code_spans) = inline_nodes(&text);
                heading.children = inline;
                append_code(&mut heading.details, code_spans);
                return heading;
            }
        }
        let mut paragraph = node(Kind::Paragraph);
        paragraph.text = preview(&text);
        let (inline, code_spans) = inline_nodes(&text);
        paragraph.children = inline;
        append_code(&mut paragraph.details, code_spans);
        paragraph
    }
}

fn find_bracket_end(chars: &[char], start: usize) -> Option<usize> {
    let mut depth = 0;
    let mut index = start;
    while index < chars.len() {
        match chars[index] {
            '[' => depth += 1,
            ']' => {
                depth -= 1;
                if depth == 0 {
                    return Some(index);
                }
            }
            _ => {}
        }
        index += 1;
    }
    None
}

fn find_paren_end(chars: &[char], start: usize) -> Option<usize> {
    let mut depth = 0;
    let mut index = start;
    while index < chars.len() {
        match chars[index] {
            '(' => depth += 1,
            ')' => {
                depth -= 1;
                if depth == 0 {
                    return Some(index);
                }
            }
            _ => {}
        }
        index += 1;
    }
    None
}

fn link_node(image: bool, text: String, details: String) -> Node {
    let mut link = node(if image { Kind::Image } else { Kind::Link });
    link.text = preview(&text);
    link.details = details;
    if image && text.trim().is_empty() {
        link.flags.push("missing-alt".into());
    }
    if !image && text.trim().is_empty() {
        link.flags.push("empty-link-text".into());
    }
    link
}

fn parse_link(chars: &[char], start: usize, image: bool) -> Option<(Node, usize)> {
    let mut after;
    let inner: String;
    if chars.get(start + 1) == Some(&'[') {
        let inner_close = find_bracket_end(chars, start + 1)?;
        if chars.get(inner_close + 1) != Some(&']') {
            return None;
        }
        inner = chars[start + 2..inner_close].iter().collect();
        after = inner_close + 2;
        let (target, alias) = match inner.split_once('|') {
            Some((target, alias)) => (target.trim(), alias.trim()),
            None => (inner.trim(), ""),
        };
        let label = if alias.is_empty() { target } else { alias };
        let node = link_node(image, label.to_string(), format!("target={target}"));
        return Some((node, after));
    }
    let close = find_bracket_end(chars, start)?;
    inner = chars[start + 1..close].iter().collect();
    after = close + 1;
    let destination = if chars.get(after) == Some(&'(') {
        let end = find_paren_end(chars, after)?;
        let raw: String = chars[after + 1..end].iter().collect();
        after = end + 1;
        raw.split_whitespace()
            .next()
            .unwrap_or_default()
            .to_string()
    } else if chars.get(after) == Some(&'[') {
        let end = find_bracket_end(chars, after)?;
        let raw: String = chars[after + 1..end].iter().collect();
        after = end + 1;
        raw
    } else {
        return None;
    };
    let node = link_node(image, inner, format!("dest={destination}"));
    Some((node, after))
}

fn inline_nodes(text: &str) -> (Vec<Node>, usize) {
    let chars: Vec<char> = text.chars().collect();
    let mut index = 0;
    let mut nodes = Vec::new();
    let mut code_spans = 0;
    while index < chars.len() {
        let character = chars[index];
        if character == '`' {
            let run = chars[index..].iter().take_while(|c| **c == '`').count();
            let mut cursor = index + run;
            let mut closed = None;
            while cursor + run <= chars.len() {
                if chars[cursor..].iter().take_while(|c| **c == '`').count() >= run {
                    closed = Some(cursor + run);
                    break;
                }
                cursor += 1;
            }
            if closed.is_some() {
                code_spans += 1;
                index = closed.unwrap_or(index + run);
            } else {
                index += run;
            }
            continue;
        }
        let image = character == '!' && chars.get(index + 1) == Some(&'[');
        let bracket = if image { index + 1 } else { index };
        if chars.get(bracket) == Some(&'[')
            && let Some((node, end)) = parse_link(&chars, bracket, image)
        {
            nodes.push(node);
            index = end;
            continue;
        }
        if character == '<'
            && let Some(end) = chars[index + 1..].iter().position(|c| *c == '>')
        {
            let inner: String = chars[index + 1..index + 1 + end].iter().collect();
            if inner.contains("://") {
                let link = link_node(false, inner.clone(), format!("dest={inner}"));
                nodes.push(link);
                index += end + 2;
                continue;
            }
        }
        index += 1;
    }
    (nodes, code_spans)
}

fn append_code(details: &mut String, code_spans: usize) {
    if code_spans == 0 {
        return;
    }
    if details.is_empty() {
        *details = format!("code={code_spans}");
    } else {
        let _ = write!(details, " code={code_spans}");
    }
}

fn nest_sections(blocks: Vec<Node>) -> Vec<Node> {
    let mut nested: Vec<Node> = Vec::new();
    let mut index = 0;
    while index < blocks.len() {
        let mut current = blocks[index].clone();
        if let Kind::Heading(level) = current.kind {
            let mut section = Vec::new();
            let mut cursor = index + 1;
            while cursor < blocks.len() {
                if let Kind::Heading(next) = blocks[cursor].kind
                    && next <= level
                {
                    break;
                }
                section.push(blocks[cursor].clone());
                cursor += 1;
            }
            current.children.extend(nest_sections(section));
            index = cursor;
        } else {
            index += 1;
        }
        nested.push(current);
    }
    nested
}

struct FlagState {
    previous: Option<u8>,
    seen: BTreeSet<String>,
}

fn apply_flags(node: &mut Node, state: &mut FlagState) {
    for child in node.children.iter_mut() {
        if let Kind::Heading(level) = child.kind {
            if let Some(previous) = state.previous
                && level > previous + 1
            {
                child.flags.push("heading-skip".into());
            }
            state.previous = Some(level);
            if let Some(slug) = child.slug.clone()
                && !state.seen.insert(slug)
            {
                child.flags.push("duplicate-slug".into());
            }
        }
        apply_flags(child, state);
    }
}

pub fn parse_markdown(source: &str) -> Node {
    let lines: Vec<&str> = source.lines().collect();
    let mut blocks: Vec<Node> = Vec::new();
    let mut start = 0;
    if lines.first().is_some_and(|line| line.trim_end() == "---") {
        let mut closing = None;
        let mut keys = 0;
        for (offset, line) in lines[1..].iter().enumerate() {
            let trimmed = line.trim_end();
            if trimmed == "---" || trimmed == "..." {
                closing = Some(offset + 1);
                break;
            }
            if line.contains(':') && !line.starts_with([' ', '\t', '-']) {
                keys += 1;
            }
        }
        if let Some(closing) = closing {
            let mut front = node(Kind::FrontMatter);
            front.details = format!("keys={keys}");
            front.span = Some((0, closing));
            blocks.push(front);
            start = closing + 1;
        }
    }
    blocks.extend(Parser::new(lines[start..].to_vec(), start).parse());
    let mut document = node(Kind::Document);
    document.children = nest_sections(blocks);
    assign_ids(&mut document, "");
    let mut state = FlagState {
        previous: None,
        seen: BTreeSet::new(),
    };
    apply_flags(&mut document, &mut state);
    document
}

fn assign_ids(node: &mut Node, prefix: &str) {
    for (position, child) in node.children.iter_mut().enumerate() {
        child.id = if prefix.is_empty() {
            format!("{}", position + 1)
        } else {
            format!("{prefix}.{}", position + 1)
        };
        assign_ids(child, &child.id.clone());
    }
}

fn collect_matches<'a>(node: &'a Node, id: &str, found: &mut Vec<&'a Node>) {
    let slug = id.strip_prefix('#').unwrap_or(id);
    for child in &node.children {
        if child.id == id || child.slug.as_deref() == Some(slug) {
            found.push(child);
        }
        collect_matches(child, id, found);
    }
}

fn count_nodes(node: &Node) -> usize {
    1 + node.children.iter().map(count_nodes).sum::<usize>()
}

fn count_flags(node: &Node) -> usize {
    node.flags.len() + node.children.iter().map(count_flags).sum::<usize>()
}

pub fn kind_name(kind: &Kind) -> String {
    match kind {
        Kind::Document => "document".to_string(),
        Kind::FrontMatter => "front matter".into(),
        Kind::Heading(level) => format!("heading {level}"),
        Kind::Paragraph => "paragraph".into(),
        Kind::List { ordered } => {
            format!("list {}", if *ordered { "ordered" } else { "unordered" })
        }
        Kind::Item => "item".into(),
        Kind::Blockquote => "blockquote".into(),
        Kind::Code => "code".into(),
        Kind::Table => "table".into(),
        Kind::ThematicBreak => "thematic break".into(),
        Kind::Html => "html".into(),
        Kind::Link => "link".into(),
        Kind::Image => "image".into(),
    }
}

pub fn node_line(node: &Node) -> String {
    let kind = kind_name(&node.kind);
    let mut line = format!("{} {kind}", node.id);
    if !node.details.is_empty() {
        let _ = write!(line, " {}", node.details);
    }
    if !node.text.is_empty() {
        let _ = write!(line, " \"{}\"", node.text);
    }
    if let Some(slug) = &node.slug {
        let _ = write!(line, " #{slug}");
    }
    for flag in &node.flags {
        let _ = write!(line, " flag={flag}");
    }
    line
}

fn write_tree(node: &Node, depth: usize, output: &mut String) {
    for child in &node.children {
        let _ = writeln!(output, "{}{}", "  ".repeat(depth), node_line(child));
        write_tree(child, depth + 1, output);
    }
}

pub fn find_node<'a>(document: &'a Node, id: &str) -> Result<&'a Node, String> {
    let mut found = Vec::new();
    collect_matches(document, id, &mut found);
    match found.len() {
        0 => Err(format!("id를 찾을 수 없습니다: {id}")),
        1 => Ok(found[0]),
        matches => {
            let ids = found
                .iter()
                .map(|node| node.id.as_str())
                .collect::<Vec<_>>()
                .join(" ");
            Err(format!(
                "id가 {matches}개 노드와 일치합니다. 자리 번호를 사용하세요: {ids}"
            ))
        }
    }
}

pub fn render(document: &Node, path: &Path, id: Option<&str>) -> Result<String, String> {
    let root = match id {
        None => document,
        Some(value) => find_node(document, value)?,
    };
    let mut output = String::new();
    if id.is_some() {
        let _ = writeln!(
            output,
            "subtree {} nodes={} flags={}",
            root.id,
            count_nodes(root),
            count_flags(root)
        );
        let _ = writeln!(output, "{}", node_line(root));
        write_tree(root, 1, &mut output);
    } else {
        let _ = writeln!(
            output,
            "document {} nodes={} flags={}",
            path.display(),
            count_nodes(document),
            count_flags(document)
        );
        write_tree(document, 0, &mut output);
    }
    Ok(output)
}
