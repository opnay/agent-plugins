#!/usr/bin/env python3
from __future__ import annotations

import argparse
import re
from pathlib import Path


VAULT_FILES = {
    "README.md": """# Memory Vault

이 폴더는 에이전트 사용 전반에서 반복적으로 참고할 장기 기억을 보관합니다.
작업 전 `INDEX.md`를 확인하고, 재사용 가능한 사용자 선호, 결정, 환경, workflow, 용어, 질문만 갱신합니다.
""",
    "preferences.md": """# Preferences

- 아직 기록된 사용자 선호가 없습니다.
""",
    "decisions.md": """# Decisions

- 아직 기록된 결정이 없습니다.
""",
    "environment.md": """# Environment

- 아직 기록된 환경 정보가 없습니다.
""",
    "workflows.md": """# Workflows

- 아직 기록된 workflow가 없습니다.
""",
    "glossary.md": """# Glossary

- 아직 기록된 용어가 없습니다.
""",
    "open-questions.md": """# Open Questions

- 아직 기록된 질문이 없습니다.
""",
}

DOC_NAME_RE = re.compile(r"^(?P<index>\d{3})-[a-z0-9]+(?:-[a-z0-9]+)*\.md$")

BASIC_DOCS = [
    ("preferences.md", "사용자 선호와 상호작용 기본값"),
    ("decisions.md", "오래 유지해야 하는 결정과 정책"),
    ("environment.md", "경로, 도구, 런타임, 로컬 환경"),
    ("workflows.md", "반복 작업 절차와 검증 방식"),
    ("glossary.md", "용어와 약어"),
    ("open-questions.md", "아직 풀리지 않은 기억 후보"),
]

KNOWLEDGE_DOC = """# {title}

- 아직 기록된 내용이 없습니다.
"""


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Initialize or update a personal agent memory vault folder.")
    parser.add_argument(
        "target",
        nargs="?",
        default="~/Workspace/Memory-vault",
        help="Target folder that should itself be managed as the memory vault. Defaults to ~/Workspace/Memory-vault.",
    )
    parser.add_argument(
        "--category",
        action="append",
        default=[],
        help="Category path to index, up to two levels, for example Agents or Agents/Prompting.",
    )
    parser.add_argument(
        "--knowledge",
        action="append",
        default=[],
        help="Knowledge document path to create, for example Programming/some-knowledge or Programming/React/hook-rules.",
    )
    parser.add_argument("--dry-run", action="store_true", help="Print planned changes without writing files.")
    return parser.parse_args()


def ensure_target(path: Path) -> Path:
    target = path.expanduser().resolve()
    if not target.exists():
        raise SystemExit(f"Target folder does not exist: {target}")
    if not target.is_dir():
        raise SystemExit(f"Target is not a folder: {target}")
    return target


def normalize_category(value: str) -> Path:
    raw_parts = [part.strip() for part in value.replace("\\", "/").split("/") if part.strip()]
    if not raw_parts or len(raw_parts) > 2:
        raise SystemExit(f"Category must have one or two path parts: {value}")
    if any(part in {".", ".."} for part in raw_parts):
        raise SystemExit(f"Category cannot use relative traversal: {value}")
    if Path(value).is_absolute():
        raise SystemExit(f"Category must be relative: {value}")
    return Path(*raw_parts)


def slugify(value: str) -> str:
    slug = re.sub(r"[^a-z0-9]+", "-", value.strip().lower()).strip("-")
    if not slug:
        raise SystemExit(f"Knowledge slug is empty: {value}")
    return slug


def normalize_knowledge(value: str) -> tuple[Path, str]:
    raw_parts = [part.strip() for part in value.replace("\\", "/").split("/") if part.strip()]
    if len(raw_parts) not in {2, 3}:
        raise SystemExit(f"Knowledge path must be category/slug or category/subcategory/slug: {value}")
    if any(part in {".", ".."} for part in raw_parts):
        raise SystemExit(f"Knowledge path cannot use relative traversal: {value}")
    if Path(value).is_absolute():
        raise SystemExit(f"Knowledge path must be relative: {value}")
    return Path(*raw_parts[:-1]), slugify(raw_parts[-1])


def discover_categories(target: Path) -> set[Path]:
    categories: set[Path] = set()
    for first in sorted(target.iterdir()):
        if not first.is_dir() or first.name.startswith("."):
            continue
        try:
            has_first_index = (first / "INDEX.md").exists()
            children = sorted(first.iterdir())
        except PermissionError:
            continue
        if has_first_index:
            categories.add(Path(first.name))
        if discover_docs(first):
            categories.add(Path(first.name))
        for second in children:
            try:
                has_second_index = second.is_dir() and not second.name.startswith(".") and (second / "INDEX.md").exists()
            except PermissionError:
                continue
            if has_second_index:
                categories.add(Path(first.name, second.name))
            if second.is_dir() and not second.name.startswith(".") and discover_docs(second):
                categories.add(Path(first.name, second.name))
    return categories


def expand_categories(categories: set[Path]) -> set[Path]:
    expanded = set(categories)
    for category in categories:
        if len(category.parts) == 2:
            expanded.add(Path(category.parts[0]))
    return expanded


def category_title(category: Path) -> str:
    return " - ".join(category.parts)


def doc_title(slug: str) -> str:
    return " ".join(part.capitalize() for part in slug.split("-"))


def discover_docs(folder: Path) -> list[Path]:
    if not folder.exists() or not folder.is_dir():
        return []
    return sorted(path for path in folder.iterdir() if path.is_file() and DOC_NAME_RE.match(path.name))


def next_knowledge_doc(folder: Path, slug: str) -> Path:
    existing = discover_docs(folder)
    for path in existing:
        if path.name[4:-3] == slug:
            return path
    max_index = max((int(path.name[:3]) for path in existing), default=0)
    return folder / f"{max_index + 1:03d}-{slug}.md"


def markdown_list(lines: list[str], empty: str) -> str:
    return "\n".join(lines) if lines else f"- {empty}"


def render_root_index(categories: set[Path]) -> str:
    basic = [f"- [{name}]({name}): {description}" for name, description in BASIC_DOCS]
    return f"""# Memory Vault Index

## Core Documents

{markdown_list(basic, "아직 등록된 기본 문서가 없습니다.")}

## Categories

{category_map(categories)}
"""


def planned_action(file_path: Path, content: str) -> tuple[str, Path, str]:
    if not file_path.exists():
        return "create", file_path, content
    if file_path.read_text(encoding="utf-8") == content:
        return "preserve", file_path, ""
    return "updated", file_path, content


def render_category_index(target: Path, category: Path, categories: set[Path], planned_docs: dict[Path, list[Path]]) -> str:
    folder = target / category
    docs = sorted({path.name for path in discover_docs(folder)} | {path.name for path in planned_docs.get(category, [])})
    doc_links = [f"- [{name}]({name})" for name in docs]
    children = [
        f"- [{child.parts[1]}/INDEX.md]({child.parts[1]}/INDEX.md)"
        for child in sorted(categories)
        if len(category.parts) == 1 and len(child.parts) == 2 and child.parts[0] == category.parts[0]
    ]
    return f"""# {category_title(category)} Index

## Documents

{markdown_list(doc_links, "아직 등록된 지식 문서가 없습니다.")}

## Child Categories

{markdown_list(children, "아직 등록된 하위 카테고리가 없습니다.")}
"""


def category_map(categories: set[Path]) -> str:
    if not categories:
        return "- 아직 등록된 카테고리가 없습니다."

    lines: list[str] = []
    first_level = sorted({category.parts[0] for category in categories})
    for first in first_level:
        lines.append(f"- [{first}/INDEX.md]({first}/INDEX.md)")
        children = sorted(category.parts[1] for category in categories if len(category.parts) == 2 and category.parts[0] == first)
        for child in children:
            lines.append(f"  - [{first}/{child}/INDEX.md]({first}/{child}/INDEX.md)")
    return "\n".join(lines)


def agents_section(categories: set[Path]) -> str:
    return f"""<!-- memory-vault:start -->
## Memory Vault

- 관련 작업 전 `INDEX.md`를 확인하고, 필요한 장기 기억 문서를 읽습니다.
- 사용자 선호는 `preferences.md`에 기록합니다.
- 오래 유지해야 하는 결정과 정책은 `decisions.md`에 기록합니다.
- 경로, 도구, 런타임, 로컬 환경은 `environment.md`에 기록합니다.
- 반복 작업 절차와 검증 방식은 `workflows.md`에 기록합니다.
- 용어와 약어는 `glossary.md`에 기록합니다.
- 아직 확정되지 않은 기억 후보는 `open-questions.md`에 기록합니다.
- 1-2차 카테고리는 각 폴더의 `INDEX.md`에 관리합니다.
- 지식 문서는 `<index>-<lowercase-hyphen-slug>.md` 형식으로 만들고, 같은 폴더의 `INDEX.md`에 링크합니다.
- 일회성 진행 로그, 임시 오류, 추측, 취소된 방향, 민감 정보는 저장하지 않습니다.
- 기존 vault 문서는 삭제하지 말고, 필요한 경우 작은 단위로 갱신합니다.

### Category Map

{category_map(categories)}
<!-- memory-vault:end -->"""


def update_agents_text(existing: str, section: str) -> tuple[str, str]:
    start = "<!-- memory-vault:start -->"
    end = "<!-- memory-vault:end -->"
    start_count = existing.count(start)
    end_count = existing.count(end)
    if start_count != end_count:
        raise SystemExit("AGENTS.md has an incomplete memory-vault marker section.")
    if start_count > 1:
        raise SystemExit("AGENTS.md has multiple memory-vault marker sections.")
    if start_count == 1 and existing.index(start) > existing.index(end):
        raise SystemExit("AGENTS.md has memory-vault markers in the wrong order.")

    if start_count == 1:
        before = existing.split(start, 1)[0].rstrip()
        after = existing.split(end, 1)[1].lstrip()
        text = f"{before}\n\n{section}\n"
        if after:
            text += f"\n{after}"
        if text == existing:
            return existing, "preserve"
        return text, "updated"

    text = existing.rstrip()
    if text:
        text += "\n\n"
    text += f"{section}\n"
    if text == existing:
        return existing, "preserve"
    return text, "updated"


def main() -> int:
    args = parse_args()
    target = ensure_target(Path(args.target))
    requested_categories = {normalize_category(category) for category in args.category}
    requested_knowledge = [normalize_knowledge(knowledge) for knowledge in args.knowledge]
    knowledge_categories = {category for category, _ in requested_knowledge}
    categories = expand_categories(discover_categories(target) | requested_categories | knowledge_categories)
    actions: list[tuple[str, Path, str]] = []
    planned_docs: dict[Path, list[Path]] = {}

    for filename, content in VAULT_FILES.items():
        file_path = target / filename
        if file_path.exists():
            actions.append(("preserve", file_path, ""))
        else:
            actions.append(("create", file_path, content))

    for category, slug in requested_knowledge:
        doc_path = next_knowledge_doc(target / category, slug)
        if doc_path.exists():
            actions.append(("preserve", doc_path, ""))
        else:
            actions.append(("create", doc_path, KNOWLEDGE_DOC.format(title=doc_title(slug))))
        planned_docs.setdefault(category, []).append(doc_path)

    actions.append(planned_action(target / "INDEX.md", render_root_index(categories)))

    for category in sorted(categories):
        index_path = target / category / "INDEX.md"
        content = render_category_index(target, category, categories, planned_docs)
        actions.append(planned_action(index_path, content))

    agents_path = target / "AGENTS.md"
    section = agents_section(categories)
    if agents_path.exists():
        next_text, status = update_agents_text(agents_path.read_text(encoding="utf-8"), section)
        actions.append((status, agents_path, next_text))
    else:
        actions.append(("create", agents_path, f"# AGENTS.md\n\n{section}\n"))

    for action, file_path, _ in actions:
        print(f"{action}: {file_path}")

    if args.dry_run:
        return 0

    for action, file_path, content in actions:
        if action in {"create", "updated"}:
            file_path.parent.mkdir(parents=True, exist_ok=True)
            file_path.write_text(content, encoding="utf-8")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
