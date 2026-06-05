#!/usr/bin/env python3
from __future__ import annotations

import argparse
from pathlib import Path


VAULT_FILES = {
    "README.md": """# Memory Vault

이 폴더는 현재 프로젝트의 지속 지식을 보관합니다.
작업 전 `INDEX.md`를 확인하고, 오래 유지해야 하는 결정과 용어와 질문만 갱신합니다.
""",
    "INDEX.md": """# Memory Vault Index

## 현재 요약

- 아직 기록된 프로젝트 요약이 없습니다.

## 주요 문서

- `decisions.md`: 오래 유지해야 하는 결정
- `glossary.md`: 프로젝트 용어와 약어
- `open-questions.md`: 아직 풀리지 않은 질문

## 카테고리

- 아직 등록된 카테고리가 없습니다.
""",
    "decisions.md": """# Decisions

- 아직 기록된 결정이 없습니다.
""",
    "glossary.md": """# Glossary

- 아직 기록된 용어가 없습니다.
""",
    "open-questions.md": """# Open Questions

- 아직 기록된 질문이 없습니다.
""",
}

CATEGORY_INDEX = """# {title} Index

## 범위

- 아직 기록된 범위가 없습니다.

## 내용

- 아직 등록된 내용이 없습니다.
"""


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Initialize or update a user-provided memory vault folder.")
    parser.add_argument("target", help="Target folder that should itself be managed as the memory vault.")
    parser.add_argument(
        "--category",
        action="append",
        default=[],
        help="Category path to index, up to two levels, for example Programming or Programming/React.",
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


def discover_categories(target: Path) -> set[Path]:
    categories: set[Path] = set()
    for first in sorted(target.iterdir()):
        if not first.is_dir() or first.name.startswith("."):
            continue
        if (first / "INDEX.md").exists():
            categories.add(Path(first.name))
        for second in sorted(first.iterdir()):
            if second.is_dir() and not second.name.startswith(".") and (second / "INDEX.md").exists():
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


def category_map(categories: set[Path]) -> str:
    if not categories:
        return "- 아직 등록된 카테고리가 없습니다."

    lines: list[str] = []
    first_level = sorted({category.parts[0] for category in categories})
    for first in first_level:
        lines.append(f"- `{first}/INDEX.md`")
        children = sorted(category.parts[1] for category in categories if len(category.parts) == 2 and category.parts[0] == first)
        for child in children:
            lines.append(f"  - `{first}/{child}/INDEX.md`")
    return "\n".join(lines)


def agents_section(categories: set[Path]) -> str:
    return f"""<!-- memory-vault:start -->
## Memory Vault

- 작업 전 `INDEX.md`를 확인하고, 이 폴더의 지속 지식을 기준으로 삼습니다.
- 오래 유지해야 하는 프로젝트 결정은 `decisions.md`에 기록합니다.
- 프로젝트 용어와 약어는 `glossary.md`에 기록합니다.
- 미해결 질문은 `open-questions.md`에 기록합니다.
- 1-2차 카테고리는 각 폴더의 `INDEX.md`에 관리합니다.
- 임시 작업 로그나 추측은 지속 지식으로 확정되기 전까지 memory vault에 남기지 않습니다.
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
    categories = expand_categories(discover_categories(target) | requested_categories)
    actions: list[tuple[str, Path, str]] = []

    for filename, content in VAULT_FILES.items():
        file_path = target / filename
        if file_path.exists():
            actions.append(("preserve", file_path, ""))
        else:
            actions.append(("create", file_path, content))

    for category in sorted(categories):
        index_path = target / category / "INDEX.md"
        if index_path.exists():
            actions.append(("preserve", index_path, ""))
        else:
            actions.append(("create", index_path, CATEGORY_INDEX.format(title=category_title(category))))

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
