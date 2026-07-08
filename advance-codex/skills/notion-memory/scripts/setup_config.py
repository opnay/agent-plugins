#!/usr/bin/env python3
"""Create or validate a notion-memory TOML config."""

from __future__ import annotations

import argparse
import sys
from pathlib import Path


DEFAULT_PATH = Path("~/.agents/configs/notion-memory.toml").expanduser()
DEFAULTS = {
    "general": {
        "timezone": "Asia/Seoul",
        "language": "ko",
        "locale": "ko-KR",
        "register": "polite",
    },
    "notion": {
        "db_url": "",
        "db_id": "",
        "data_source_id": "",
        "view_id": "",
        "schema_status": "unverified",
    },
}
SECRET_KEY_PARTS = ("token", "secret", "credential", "cookie", "auth", "password")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--path", type=Path, default=DEFAULT_PATH, help="Config path")
    parser.add_argument("--db-url", default=None, help="Notion memory database URL")
    parser.add_argument("--db-id", default=None, help="Notion memory database ID")
    parser.add_argument("--data-source-id", default=None, help="Optional Notion data source ID")
    parser.add_argument("--view-id", default=None, help="Optional Notion view ID")
    parser.add_argument("--schema-status", default=None, help="Schema status")
    parser.add_argument("--timezone", default=None, help="IANA timezone")
    parser.add_argument("--language", default=None, help="Memory document language")
    parser.add_argument("--locale", default=None, help="Memory document locale")
    parser.add_argument("--register", default=None, help="Memory document register")
    parser.add_argument("--print-template", action="store_true", help="Print a blank config template")
    parser.add_argument("--dry-run", action="store_true", help="Print the config without writing")
    parser.add_argument("--check", action="store_true", help="Validate the config and exit")
    parser.add_argument(
        "--allow-missing-db",
        action="store_true",
        help="Allow template output without notion.db_url or notion.db_id",
    )
    return parser.parse_args()


def read_config(path: Path) -> dict[str, dict[str, str]]:
    config = clone_defaults()
    if not path.exists():
        return config

    section = None
    for raw_line in path.read_text(encoding="utf-8").splitlines():
        line = raw_line.strip()
        if not line or line.startswith("#"):
            continue
        if line.startswith("[") and line.endswith("]"):
            section = line[1:-1].strip()
            if section not in config:
                config[section] = {}
            continue
        if "=" not in line or section is None:
            raise ValueError(f"Unsupported TOML line: {raw_line}")
        key, value = line.split("=", 1)
        key = key.strip()
        if looks_secret(key):
            raise ValueError(f"Refusing secret-like key: {section}.{key}")
        config.setdefault(section, {})[key] = parse_value(value.strip())
    return config


def clone_defaults() -> dict[str, dict[str, str]]:
    return {section: values.copy() for section, values in DEFAULTS.items()}


def parse_value(value: str) -> str:
    if len(value) >= 2 and value[0] == '"' and value[-1] == '"':
        return value[1:-1].replace('\\"', '"').replace("\\\\", "\\")
    return value


def apply_args(config: dict[str, dict[str, str]], args: argparse.Namespace) -> None:
    updates = {
        ("general", "timezone"): args.timezone,
        ("general", "language"): args.language,
        ("general", "locale"): args.locale,
        ("general", "register"): args.register,
        ("notion", "db_url"): args.db_url,
        ("notion", "db_id"): args.db_id,
        ("notion", "data_source_id"): args.data_source_id,
        ("notion", "view_id"): args.view_id,
        ("notion", "schema_status"): args.schema_status,
    }
    for (section, key), value in updates.items():
        if value is not None:
            config.setdefault(section, {})[key] = value


def validate(config: dict[str, dict[str, str]], allow_missing_db: bool) -> list[str]:
    errors = []
    for section, values in DEFAULTS.items():
        if section not in config:
            errors.append(f"missing section: {section}")
            continue
        for key in values:
            if key not in config[section]:
                errors.append(f"missing key: {section}.{key}")

    for section, values in config.items():
        for key in values:
            if looks_secret(key):
                errors.append(f"secret-like key is not allowed: {section}.{key}")

    notion = config.get("notion", {})
    if not allow_missing_db and not (notion.get("db_url") or notion.get("db_id")):
        errors.append("missing notion.db_url or notion.db_id")
    return errors


def looks_secret(key: str) -> bool:
    lowered = key.lower()
    return any(part in lowered for part in SECRET_KEY_PARTS)


def render(config: dict[str, dict[str, str]]) -> str:
    lines = []
    for section in ("general", "notion"):
        lines.append(f"[{section}]")
        for key in DEFAULTS[section]:
            lines.append(f'{key} = "{escape(config.get(section, {}).get(key, ""))}"')
        lines.append("")
    return "\n".join(lines).rstrip() + "\n"


def escape(value: str) -> str:
    return str(value).replace("\\", "\\\\").replace('"', '\\"')


def main() -> int:
    args = parse_args()
    path = args.path.expanduser()
    config = clone_defaults() if args.print_template else read_config(path)
    apply_args(config, args)

    errors = validate(config, allow_missing_db=args.allow_missing_db or args.print_template)
    if args.check:
        if errors:
            for error in errors:
                print(f"ERROR: {error}", file=sys.stderr)
            return 1
        print(f"WORKED: config valid: {path}")
        return 0

    output = render(config)
    if args.print_template or args.dry_run:
        print(output, end="")
        return 0 if not errors else 1

    if errors:
        for error in errors:
            print(f"ERROR: {error}", file=sys.stderr)
        return 1

    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(output, encoding="utf-8")
    print(f"WORKED: wrote config: {path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
