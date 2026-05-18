#!/usr/bin/env python3
"""Codex SessionStart context loader for loop-kit turn-gate sessions."""

import datetime as _dt
import json
import re
import sys
from pathlib import Path
from typing import Optional, Tuple


FRONTMATTER_RE = re.compile(r"\A---\n(.*?)\n---\n", re.DOTALL)
MAX_CONTEXT_CHARS = 1500


def main() -> int:
    try:
        payload = json.load(sys.stdin)
    except Exception:
        return 0

    if payload.get("hook_event_name") != "SessionStart":
        return 0
    if payload.get("source") not in ("startup", "resume"):
        return 0

    session_id = _payload_session_id(payload)
    cwd = Path(payload.get("cwd") or ".").resolve()
    plan_path, plan_date = _find_plan(cwd)
    if plan_path is None:
        _emit_context(
            "No turn-gate session plan was found under .agents/sessions. "
            "Start normal preparation and ask for scope if the user request is ambiguous."
        )
        return 0

    plan = _read_frontmatter(plan_path)
    active_flow = str(plan.get("active_flow") or "").strip()
    flow_path = plan_path.parent / f"{active_flow}.md" if active_flow else None
    flow_source = "active flow"

    if flow_path is None or not flow_path.is_file():
        flow_path = _find_latest_flow(plan_path.parent)
        flow_source = "latest flow"

    flow = _read_frontmatter(flow_path) if flow_path else {}
    context = _build_context(plan_date, session_id, plan, flow_path, flow_source, flow)
    _emit_context(context)
    return 0


def _find_plan(cwd: Path) -> Tuple[Optional[Path], str]:
    sessions_dir = None
    for root in [cwd, *cwd.parents]:
        candidate = root / ".agents" / "sessions"
        if candidate.is_dir():
            sessions_dir = candidate
            break
    if sessions_dir is None:
        return None, ""

    today = _dt.datetime.now().strftime("%Y%m%d")
    today_plan = sessions_dir / today / "000-plan.md"
    if today_plan.is_file():
        return today_plan, today

    dated_dirs = sorted(
        path for path in sessions_dir.iterdir() if path.is_dir() and path.name.isdigit()
    )
    for dated_dir in reversed(dated_dirs):
        plan_path = dated_dir / "000-plan.md"
        if plan_path.is_file():
            return plan_path, dated_dir.name
    return None, ""


def _find_latest_flow(day_dir: Path) -> Optional[Path]:
    flows = sorted(
        path
        for path in day_dir.glob("[0-9][0-9][0-9]-*.md")
        if path.name != "000-plan.md"
    )
    return flows[-1] if flows else None


def _read_frontmatter(path: Optional[Path]) -> dict:
    if path is None:
        return {}
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


def _build_context(
    plan_date: str,
    session_id: str,
    plan: dict,
    flow_path: Optional[Path],
    flow_source: str,
    flow: dict,
) -> str:
    active_flow = _clean(plan.get("active_flow"))
    latest_user_request = _clean(plan.get("latest_user_request"))
    latest_decision = _clean(plan.get("latest_decision"))
    required_next_action = _clean(plan.get("required_next_action"))
    pending_question_state = _clean(plan.get("pending_question_state"))
    verification_status = _clean(plan.get("verification_status"))

    flow_name = flow_path.name if flow_path else "none"
    closure = _flow_closure_summary(flow)

    lines = [
        f"Turn-gate session context found for {plan_date}.",
        _line("Current Codex session id", session_id or "not provided"),
        _line("Plan active_flow", active_flow or "none"),
        _line(flow_source.title(), flow_name),
        _line("Latest user request", latest_user_request or "not recorded"),
        _line("Latest decision", latest_decision or "not recorded"),
        _line("Verification status", verification_status or "not recorded"),
        _line("Required next action", required_next_action or "not recorded"),
        _line("Pending question", pending_question_state or "not recorded"),
        _line("Closure state", closure),
        (
            "Use this as startup context only. Do not treat it as approval for "
            "commit, push, release, deletion, external actions, or terminal closure."
        ),
    ]
    return "\n".join(lines)[:MAX_CONTEXT_CHARS]


def _flow_closure_summary(flow: dict) -> str:
    if not flow:
        return "flow record not found"
    if flow.get("user_explicit_stop") is True:
        return "user explicit stop recorded"
    if flow.get("confirmed_closure") is True:
        return "confirmed closure recorded"
    if flow.get("terminal_summary_allowed") is True:
        return "terminal summary allowed"
    if flow.get("turn_gate_active") is True:
        return "turn-gate active"
    return "not recorded"


def _clean(value) -> str:
    text = str(value or "").replace("\n", " ").strip()
    return re.sub(r"\s+", " ", text)


def _line(label: str, value: str) -> str:
    text = value.rstrip(".")
    return f"{label}: {text}."


def _payload_session_id(payload: dict) -> str:
    for key in ("session_id", "sessionId", "conversation_id", "conversationId"):
        value = payload.get(key)
        if isinstance(value, str) and value.strip():
            return value.strip()

    for key in ("session", "conversation"):
        value = payload.get(key)
        if isinstance(value, dict):
            nested_id = value.get("id")
            if isinstance(nested_id, str) and nested_id.strip():
                return nested_id.strip()

    return ""


def _emit_context(context: str) -> None:
    output = {
        "hookSpecificOutput": {
            "hookEventName": "SessionStart",
            "additionalContext": context,
        }
    }
    print(json.dumps(output, ensure_ascii=False))


if __name__ == "__main__":
    raise SystemExit(main())
