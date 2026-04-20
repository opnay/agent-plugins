#!/usr/bin/env python3
from __future__ import annotations

import argparse
import os
import re
import sys
from dataclasses import dataclass
from datetime import datetime
from pathlib import Path

NUMBERED_HEADER_RE = re.compile(r"^(?P<number>\d+)\.\s(?P<label>[^:]+):(?:\s(?P<value>.*))?$")
MARKDOWN_MARKER_RE = re.compile(r"^<!-- session-manager:(?P<key>[a-z_]+) -->$")


@dataclass(frozen=True)
class SectionSpec:
    key: str
    label: str
    placeholder: str
    number: int | None = None


@dataclass(frozen=True)
class RecordSchema:
    name: str
    artifact: str
    filename: str
    style: str
    sections: tuple[SectionSpec, ...]
    read_only_fields: frozenset[str] = frozenset()

    @property
    def writable_fields(self) -> tuple[str, ...]:
        return tuple(section.key for section in self.sections if section.key not in self.read_only_fields)


@dataclass(frozen=True)
class ValidationReport:
    valid: bool
    artifact: str
    detected_schema: str | None
    record_path: Path
    errors: tuple[str, ...]


SESSION_RECORD = RecordSchema(
    name="session-record",
    artifact="session-record",
    filename="session_record.md",
    style="numbered",
    sections=(
        SectionSpec("session_uuid", "Session UUID", "[TODO: Capture the session UUID]", number=1),
        SectionSpec("timestamp", "Timestamp", "[TODO: Capture the session timestamp]", number=2),
        SectionSpec("goal", "Goal", "[TODO: Describe the task goal]", number=3),
        SectionSpec("scope", "Scope", "[TODO: Define what is in scope]", number=4),
        SectionSpec("constraints", "Constraints", "[TODO: List constraints and non-goals]", number=5),
        SectionSpec("working_assumptions", "Working assumptions", "[TODO: Capture key assumptions]", number=6),
        SectionSpec("runbook", "Runbook", "[TODO: Describe the expected operating steps]", number=7),
        SectionSpec(
            "verification_evidence",
            "Verification evidence",
            "[TODO: Record checks and outcomes]",
            number=8,
        ),
        SectionSpec("mistake_notes", "Mistake notes", "[TODO: Record mistakes or leave none]", number=9),
        SectionSpec("residual_risks", "Residual risks", "[TODO: Capture remaining risks]", number=10),
    ),
    read_only_fields=frozenset({"session_uuid", "timestamp"}),
)

CHANGE_RECORD = RecordSchema(
    name="change-record",
    artifact="change-record",
    filename="change_record.md",
    style="markdown",
    sections=(
        SectionSpec("purpose", "Purpose", "[TODO: Summarize the purpose of the change.]"),
        SectionSpec("key_changes", "Key Changes", "[TODO: Capture the key changes for reviewers.]"),
        SectionSpec("verification", "Verification", "[TODO: Record validation commands and outcomes.]"),
        SectionSpec("review_gates", "Review Gates", "[TODO: Record review lenses and findings.]"),
        SectionSpec("reviewer_notes", "Reviewer Notes", "[TODO: Add reviewer-relevant notes or leave none.]"),
        SectionSpec("residual_risks", "Residual Risks", "[TODO: Capture remaining risks.]"),
    ),
)

SCHEMAS_BY_ARTIFACT = {
    "session-record": SESSION_RECORD,
    "change-record": CHANGE_RECORD,
}
SCHEMAS_BY_FILENAME = {
    SESSION_RECORD.filename: SESSION_RECORD,
    CHANGE_RECORD.filename: CHANGE_RECORD,
}


def supported_artifact_names() -> list[str]:
    return sorted(SCHEMAS_BY_ARTIFACT)


def schema_for_artifact(artifact: str) -> RecordSchema:
    try:
        return SCHEMAS_BY_ARTIFACT[artifact]
    except KeyError as error:
        raise SystemExit(f"unsupported artifact: {artifact}") from error


def resolve_existing_artifact_path(session_dir: Path, artifact: str) -> Path | None:
    candidate = session_dir / schema_for_artifact(artifact).filename
    if candidate.exists():
        return candidate
    return None


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Manage session artifacts under .sessions/<uuid>/.")
    subparsers = parser.add_subparsers(dest="command", required=True)

    def add_target_args(target_parser: argparse.ArgumentParser) -> None:
        target_parser.add_argument(
            "--artifact",
            choices=supported_artifact_names(),
            default="session-record",
            help="Target artifact kind (default: session-record)",
        )
        target_parser.add_argument("--record-path", type=Path, help="Direct path to the target record file")
        target_parser.add_argument(
            "--root",
            type=Path,
            default=Path(".sessions"),
            help="Base directory for session folders when resolving by UUID (default: .sessions)",
        )

    init_parser = subparsers.add_parser("init", help="Create a new artifact template")
    init_parser.add_argument(
        "--artifact",
        choices=supported_artifact_names(),
        default="session-record",
        help="Artifact kind to initialize (default: session-record)",
    )
    init_parser.add_argument("--root", type=Path, default=Path(".sessions"), help="Base directory for generated session folders (default: .sessions)")
    init_parser.add_argument("--goal", help="Optional session-record goal text")
    init_parser.add_argument("--scope", help="Optional session-record scope text")
    init_parser.add_argument("--constraints", help="Optional session-record constraints text")
    init_parser.add_argument("--force", action="store_true", help="Overwrite an existing artifact file")

    ensure_parser = subparsers.add_parser("ensure", help="Return an artifact path, creating it when missing")
    ensure_parser.add_argument(
        "--artifact",
        choices=supported_artifact_names(),
        default="session-record",
        help="Artifact kind to ensure (default: session-record)",
    )
    ensure_parser.add_argument("--root", type=Path, default=Path(".sessions"), help="Base directory for generated session folders (default: .sessions)")
    ensure_parser.add_argument("--goal", help="Optional session-record goal text")
    ensure_parser.add_argument("--scope", help="Optional session-record scope text")
    ensure_parser.add_argument("--constraints", help="Optional session-record constraints text")

    write_parser = subparsers.add_parser("write", help="Update a field in an artifact")
    add_target_args(write_parser)
    write_parser.add_argument("--field", required=True, help="Logical field name to update")
    value_group = write_parser.add_mutually_exclusive_group(required=True)
    value_group.add_argument("--value", help="New value for the field")
    value_group.add_argument("--value-file", help="Path to a UTF-8 file containing the field value, or '-' to read stdin")

    show_parser = subparsers.add_parser("show", help="Show an artifact or one field value")
    add_target_args(show_parser)
    show_parser.add_argument("--field", help="Optional field name to print instead of the full record")

    validate_parser = subparsers.add_parser("validate", help="Validate an artifact structure")
    add_target_args(validate_parser)

    list_parser = subparsers.add_parser("list", help="List artifact files under the root directory")
    list_parser.add_argument(
        "--artifact",
        choices=supported_artifact_names(),
        default="session-record",
        help="Artifact kind to list (default: session-record)",
    )
    list_parser.add_argument("--root", type=Path, default=Path(".sessions"), help="Base directory for session folders (default: .sessions)")
    list_parser.add_argument("--include-schema", action="store_true", help="Include schema information in the list output")

    return parser


def resolve_session_uuid() -> str:
    resolved_uuid = os.environ.get("CODEX_THREAD_ID")
    if resolved_uuid:
        return resolved_uuid
    raise SystemExit("CODEX_THREAD_ID is required; automatic session UUID generation is disabled")


def resolve_record_path(*, artifact: str, record_path: Path | None, root: Path) -> Path:
    if record_path is not None:
        return record_path

    resolved_uuid = os.environ.get("CODEX_THREAD_ID")
    if not resolved_uuid:
        raise SystemExit("command requires --record-path or CODEX_THREAD_ID")
    session_dir = root / resolved_uuid
    existing_path = resolve_existing_artifact_path(session_dir, artifact)
    if existing_path is not None:
        return existing_path
    return session_dir / schema_for_artifact(artifact).filename


def build_initial_values(*, schema: RecordSchema, session_uuid: str, goal: str | None, scope: str | None, constraints: str | None) -> dict[str, str]:
    values = {section.key: section.placeholder for section in schema.sections}
    if schema.artifact == "session-record":
        values.update(
            {
                "session_uuid": session_uuid,
                "timestamp": datetime.now().astimezone().isoformat(timespec="seconds"),
                "goal": goal or values["goal"],
                "scope": scope or values["scope"],
                "constraints": constraints or values["constraints"],
            }
        )
    return values


def format_record(schema: RecordSchema, values: dict[str, str]) -> str:
    lines: list[str] = []
    if schema.style == "numbered":
        for section in schema.sections:
            assert section.number is not None
            raw_lines = (values.get(section.key) or section.placeholder).splitlines() or [""]
            lines.append(f"{section.number}. {section.label}: {raw_lines[0]}")
            for raw_line in raw_lines[1:]:
                lines.append(f"   {raw_line}" if raw_line else "   ")
        lines.append("")
        return "\n".join(lines)

    for section in schema.sections:
        lines.append(f"<!-- session-manager:{section.key} -->")
        lines.append(f"# {section.label}")
        content = values.get(section.key) or section.placeholder
        lines.extend(content.splitlines() or [""])
        lines.append("")
    return "\n".join(lines)


def init_record(
    *,
    artifact: str,
    root: Path,
    force: bool,
    goal: str | None = None,
    scope: str | None = None,
    constraints: str | None = None,
) -> tuple[str, Path]:
    schema = schema_for_artifact(artifact)
    resolved_uuid = resolve_session_uuid()
    session_dir = root / resolved_uuid
    session_dir.mkdir(parents=True, exist_ok=True)
    record_path = session_dir / schema.filename
    if record_path.exists() and not force:
        raise SystemExit(f"record already exists: {record_path}")

    values = build_initial_values(
        schema=schema,
        session_uuid=resolved_uuid,
        goal=goal,
        scope=scope,
        constraints=constraints,
    )
    record_path.write_text(format_record(schema, values), encoding="utf-8")
    return resolved_uuid, record_path


def ensure_record(
    *,
    artifact: str,
    root: Path,
    goal: str | None = None,
    scope: str | None = None,
    constraints: str | None = None,
) -> tuple[str, Path, bool]:
    resolved_uuid = resolve_session_uuid()
    session_dir = root / resolved_uuid
    existing_path = resolve_existing_artifact_path(session_dir, artifact)
    if existing_path is not None:
        return resolved_uuid, existing_path, False

    _, created_path = init_record(
        artifact=artifact,
        root=root,
        goal=goal,
        scope=scope,
        constraints=constraints,
        force=False,
    )
    return resolved_uuid, created_path, True


def collect_numbered_headers(lines: list[str]) -> list[tuple[int, str, int]]:
    headers: list[tuple[int, str, int]] = []
    for index, line in enumerate(lines):
        match = NUMBERED_HEADER_RE.match(line)
        if match:
            headers.append((int(match.group("number")), match.group("label"), index))
    return headers


def collect_markdown_markers(lines: list[str]) -> list[tuple[str, int]]:
    markers: list[tuple[str, int]] = []
    for index, line in enumerate(lines):
        match = MARKDOWN_MARKER_RE.match(line)
        if match:
            markers.append((match.group("key"), index))
    return markers


def match_schema(lines: list[str], schema: RecordSchema) -> bool:
    if schema.style == "numbered":
        expected = [(section.number, section.label) for section in schema.sections]
        actual = [(number, label) for number, label, _ in collect_numbered_headers(lines)]
        return actual == expected

    expected_keys = [section.key for section in schema.sections]
    actual_keys = [key for key, _ in collect_markdown_markers(lines)]
    if actual_keys != expected_keys:
        return False
    for section, (_, marker_index) in zip(schema.sections, collect_markdown_markers(lines)):
        header_index = marker_index + 1
        if header_index >= len(lines) or lines[header_index] != f"# {section.label}":
            return False
    return True


def infer_artifact_from_path(record_path: Path) -> str | None:
    schema = SCHEMAS_BY_FILENAME.get(record_path.name)
    if schema is None:
        return None
    return schema.artifact


def detect_schema(record_path: Path, artifact: str | None = None) -> RecordSchema | None:
    artifact_name = artifact or infer_artifact_from_path(record_path)
    if artifact_name is None or not record_path.exists():
        return None
    schema = schema_for_artifact(artifact_name)
    lines = record_path.read_text(encoding="utf-8").splitlines()
    if match_schema(lines, schema):
        return schema
    return None


def validate_record(record_path: Path, artifact: str | None = None) -> ValidationReport:
    artifact_name = artifact or infer_artifact_from_path(record_path) or "session-record"
    if not record_path.exists():
        return ValidationReport(
            valid=False,
            artifact=artifact_name,
            detected_schema=None,
            record_path=record_path,
            errors=(f"record does not exist: {record_path}",),
        )

    schema = detect_schema(record_path, artifact_name)
    if schema is not None:
        return ValidationReport(
            valid=True,
            artifact=artifact_name,
            detected_schema=schema.name,
            record_path=record_path,
            errors=(),
        )

    lines = record_path.read_text(encoding="utf-8").splitlines()
    if artifact_name == "session-record":
        actual = [(number, label) for number, label, _ in collect_numbered_headers(lines)]
        errors = (
            "numbered sections do not match the supported schema",
            f"actual={actual}",
        )
    else:
        actual = [key for key, _ in collect_markdown_markers(lines)]
        errors = (
            "markdown sections do not match the supported schema",
            f"actual={actual}",
        )
    return ValidationReport(
        valid=False,
        artifact=artifact_name,
        detected_schema=None,
        record_path=record_path,
        errors=errors,
    )


def parse_sections(record_path: Path, schema: RecordSchema) -> dict[str, str]:
    lines = record_path.read_text(encoding="utf-8").splitlines()
    sections: dict[str, str] = {}

    if schema.style == "numbered":
        headers = collect_numbered_headers(lines)
        for index, (_, label, start) in enumerate(headers):
            end = headers[index + 1][2] if index + 1 < len(headers) else len(lines)
            block = lines[start:end]
            match = NUMBERED_HEADER_RE.match(block[0])
            if match is None:
                raise SystemExit(f"failed to parse numbered header: {block[0]}")
            values = [match.group("value") or ""]
            for line in block[1:]:
                values.append(line[3:] if line.startswith("   ") else line)
            section = next(section for section in schema.sections if section.label == label)
            sections[section.key] = "\n".join(values).rstrip()
        return sections

    markers = collect_markdown_markers(lines)
    for index, (key, start) in enumerate(markers):
        end = markers[index + 1][1] if index + 1 < len(markers) else len(lines)
        section = next(section for section in schema.sections if section.key == key)
        header_index = start + 1
        if header_index >= len(lines) or lines[header_index] != f"# {section.label}":
            raise SystemExit(f"failed to parse markdown header for key '{key}'")
        block = lines[start + 2 : end]
        while block and block[-1] == "":
            block = block[:-1]
        sections[section.key] = "\n".join(block).rstrip()
    return sections


def show_record(record_path: Path, field: str | None = None, artifact: str | None = None) -> str:
    if field is None:
        return record_path.read_text(encoding="utf-8")

    report = validate_record(record_path=record_path, artifact=artifact)
    if not report.valid or report.detected_schema is None:
        raise SystemExit("; ".join(report.errors))
    schema = schema_for_artifact(report.artifact)
    sections = parse_sections(record_path, schema)
    if field not in sections:
        raise SystemExit(f"field '{field}' is not available for schema '{schema.name}'")
    return sections[field] + "\n"


def update_record_field(*, record_path: Path, artifact: str, field: str, value: str) -> Path:
    report = validate_record(record_path=record_path, artifact=artifact)
    if not report.valid:
        raise SystemExit("; ".join(report.errors))

    schema = schema_for_artifact(artifact)
    if field not in schema.writable_fields:
        raise SystemExit(f"field '{field}' is not writable for schema '{schema.name}'")

    sections = parse_sections(record_path, schema)
    sections[field] = value
    record_path.write_text(format_record(schema, sections), encoding="utf-8")
    return record_path


def read_value_from_args(value: str | None, value_file: str | None) -> str:
    if value is not None:
        return value
    if value_file is None:
        raise SystemExit("write requires --value or --value-file")
    if value_file == "-":
        return sys.stdin.read()
    return Path(value_file).read_text(encoding="utf-8")


def list_artifact_records(root: Path, artifact: str = "session-record") -> list[tuple[str, str, str | None, Path]]:
    if not root.exists():
        return []

    schema = schema_for_artifact(artifact)
    records: list[tuple[str, float, str | None, Path]] = []
    for record_path in root.glob(f"*/{schema.filename}"):
        detected = detect_schema(record_path, artifact)
        records.append(
            (
                record_path.parent.name,
                record_path.stat().st_mtime,
                detected.name if detected is not None else None,
                record_path,
            )
        )

    records.sort(key=lambda item: item[1], reverse=True)
    return [
        (
            session_uuid,
            datetime.fromtimestamp(updated_at).astimezone().isoformat(timespec="seconds"),
            detected_schema,
            record_path,
        )
        for session_uuid, updated_at, detected_schema, record_path in records
    ]


def list_session_records(root: Path, artifact: str = "session-record") -> list[tuple[str, str, Path]]:
    return [
        (session_uuid, updated_at, record_path)
        for session_uuid, updated_at, _, record_path in list_artifact_records(root, artifact=artifact)
    ]


def init_artifact(*, root: Path, artifact: str, goal: str | None, scope: str | None, constraints: str | None, force: bool) -> tuple[str, Path]:
    return init_record(
        artifact=artifact,
        root=root,
        force=force,
        goal=goal,
        scope=scope,
        constraints=constraints,
    )


def ensure_artifact(*, root: Path, artifact: str, goal: str | None, scope: str | None, constraints: str | None) -> tuple[str, Path, bool]:
    return ensure_record(
        artifact=artifact,
        root=root,
        goal=goal,
        scope=scope,
        constraints=constraints,
    )


def validate_record_structure(record_path: Path, artifact: str | None = None) -> list[str]:
    report = validate_record(record_path=record_path, artifact=artifact)
    if report.valid:
        return []
    return list(report.errors)


def main() -> int:
    parser = build_parser()
    args = parser.parse_args()

    if args.command == "init":
        session_uuid, record_path = init_record(
            artifact=args.artifact,
            root=args.root,
            force=args.force,
            goal=args.goal,
            scope=args.scope,
            constraints=args.constraints,
        )
        print(f"artifact={args.artifact}")
        print(f"schema={schema_for_artifact(args.artifact).name}")
        print(f"session_uuid={session_uuid}")
        print(f"record_path={record_path}")
        return 0

    if args.command == "ensure":
        session_uuid, record_path, created = ensure_record(
            artifact=args.artifact,
            root=args.root,
            goal=args.goal,
            scope=args.scope,
            constraints=args.constraints,
        )
        print(f"artifact={args.artifact}")
        print(f"schema={schema_for_artifact(args.artifact).name}")
        print(f"created={'true' if created else 'false'}")
        print(f"session_uuid={session_uuid}")
        print(f"record_path={record_path}")
        return 0

    if args.command == "write":
        record_path = resolve_record_path(
            artifact=args.artifact,
            record_path=args.record_path,
            root=args.root,
        )
        updated_path = update_record_field(
            record_path=record_path,
            artifact=args.artifact,
            field=args.field,
            value=read_value_from_args(args.value, args.value_file),
        )
        print(f"artifact={args.artifact}")
        print(f"field={args.field}")
        print(f"schema={schema_for_artifact(args.artifact).name}")
        print(f"record_path={updated_path}")
        return 0

    if args.command == "show":
        record_path = resolve_record_path(
            artifact=args.artifact,
            record_path=args.record_path,
            root=args.root,
        )
        print(show_record(record_path=record_path, field=args.field, artifact=args.artifact), end="")
        return 0

    if args.command == "validate":
        record_path = resolve_record_path(
            artifact=args.artifact,
            record_path=args.record_path,
            root=args.root,
        )
        report = validate_record(record_path=record_path, artifact=args.artifact)
        print(f"valid={'true' if report.valid else 'false'}")
        print(f"artifact={report.artifact}")
        print(f"detected_schema={report.detected_schema}")
        print(f"record_path={report.record_path}")
        for error in report.errors:
            print(f"error={error}")
        return 0 if report.valid else 1

    if args.command == "list":
        if args.include_schema:
            print("session_uuid\tupdated_at\tdetected_schema\trecord_path")
            for session_uuid, updated_at, detected_schema, record_path in list_artifact_records(args.root, artifact=args.artifact):
                print(f"{session_uuid}\t{updated_at}\t{detected_schema}\t{record_path}")
        else:
            print("session_uuid\tupdated_at\trecord_path")
            for session_uuid, updated_at, record_path in list_session_records(args.root, artifact=args.artifact):
                print(f"{session_uuid}\t{updated_at}\t{record_path}")
        return 0

    raise SystemExit(f"unsupported command: {args.command}")


if __name__ == "__main__":
    raise SystemExit(main())
