#!/usr/bin/env python3
"""Codex Stop hook guard for loop-kit turn-gate sessions."""

import datetime as _dt
import json
import re
import sys
from pathlib import Path
from typing import Optional


FRONTMATTER_RE = re.compile(r"\A---\n(.*?)\n---\n", re.DOTALL)


def main() -> int:
    try:
        payload = json.load(sys.stdin)
    except Exception:
        return 0

    if payload.get("hook_event_name") != "Stop":
        return 0
    if payload.get("stop_hook_active") is True:
        return 0

    cwd = Path(payload.get("cwd") or ".").resolve()
    plan_path = _find_today_plan(cwd)
    if plan_path is None:
        return 0

    plan = _read_frontmatter(plan_path)
    active_flow = plan.get("active_flow", "").strip()
    if not active_flow:
        return 0

    flow_path = plan_path.parent / f"{active_flow}.md"
    flow = _read_frontmatter(flow_path)
    if not flow:
        return 0

    if flow.get("turn_gate_active") is not True:
        return 0
    if flow.get("terminal_summary_allowed") is True:
        return 0
    if flow.get("user_explicit_stop") is True:
        return 0
    if flow.get("confirmed_closure") is True:
        return 0

    required_next_action = str(flow.get("required_next_action") or "").strip()
    if not required_next_action:
        return 0

    last_message = str(payload.get("last_assistant_message") or "")
    if _already_routes_forward(last_message):
        return 0

    reason = (
        "turn-gate is active and terminal closure is not allowed. "
        "Refresh the active flow record, then continue to the required next "
        f"action: {required_next_action}"
    )
    print(json.dumps({"decision": "block", "reason": reason}, ensure_ascii=False))
    return 0


def _find_today_plan(cwd: Path) -> Optional[Path]:
    for root in [cwd, *cwd.parents]:
        sessions_dir = root / ".agents" / "sessions"
        if sessions_dir.is_dir():
            today = _dt.datetime.now().strftime("%Y%m%d")
            plan_path = sessions_dir / today / "000-plan.md"
            return plan_path if plan_path.is_file() else None
    return None


def _read_frontmatter(path: Path) -> dict:
    try:
        text = path.read_text(encoding="utf-8")
    except OSError:
        return {}

    match = FRONTMATTER_RE.match(text)
    if not match:
        return {}

    result = {}
    for raw_line in match.group(1).splitlines():
        if not raw_line.strip() or raw_line.lstrip().startswith("#"):
            continue
        if ":" not in raw_line:
            continue
        key, raw_value = raw_line.split(":", 1)
        result[key.strip()] = _parse_scalar(raw_value.strip())
    return result


def _parse_scalar(value: str):
    if value == "true":
        return True
    if value == "false":
        return False
    if value == '""':
        return ""
    if len(value) >= 2 and value[0] == value[-1] == '"':
        return value[1:-1]
    return value


def _already_routes_forward(message: str) -> bool:
    lowered = message.lower()
    routing_markers = [
        "[next-flow]",
        "request_user_input",
        "다음 흐름",
        "다음으로",
        "무엇을 진행할까요",
        "어떻게 이어갈까요",
        "명시적으로 중지",
        "explicitly ask to stop",
        "await user",
    ]
    return any(marker in lowered for marker in routing_markers)


if __name__ == "__main__":
    raise SystemExit(main())
