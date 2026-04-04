"""Teammate relay with file-backed bus, mention routing, and reliability guards."""

from __future__ import annotations

import argparse
import gzip
import json
import os
import re
import shutil
import sys
import time
import uuid
from datetime import datetime, timezone
from pathlib import Path
from typing import Any

try:
    import fcntl
except ImportError as exc:  # pragma: no cover - platform specific guard
    raise RuntimeError(
        "team_bus requires Unix-style file locking (fcntl). Linux/macOS only."
    ) from exc

MENTION_PATTERN = re.compile(r"@([a-zA-Z0-9_-]+)")
DEFAULT_TEAM = ["alpha", "beta"]


def _now_iso() -> str:
    return datetime.now(timezone.utc).isoformat()


def _ensure_parent(path: Path) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)


def _default_meta_path(bus_path: Path) -> Path:
    return bus_path.with_name(f"{bus_path.stem}.meta.json")


def _default_checkpoint_path(bus_path: Path) -> Path:
    return bus_path.with_name(f"{bus_path.stem}.checkpoints.json")


def _default_dlq_path(bus_path: Path) -> Path:
    if bus_path.suffix:
        return bus_path.with_name(f"{bus_path.stem}.dlq{bus_path.suffix}")
    return bus_path.with_name(f"{bus_path.name}.dlq")


def _default_archive_dir(bus_path: Path) -> Path:
    return bus_path.with_name(f"{bus_path.stem}.archives")


def _scan_max_seq(bus_path: Path) -> int:
    max_seq = 0
    if not bus_path.exists():
        return max_seq

    with bus_path.open("r", encoding="utf-8") as handle:
        for raw_line in handle:
            line = raw_line.strip()
            if not line:
                continue
            try:
                parsed = json.loads(line)
            except json.JSONDecodeError:
                continue
            if not isinstance(parsed, dict):
                continue
            seq = _coerce_int(parsed.get("seq"), default=0)
            if seq > max_seq:
                max_seq = seq
    return max_seq


def _next_seq(bus_path: Path) -> int:
    meta_path = _default_meta_path(bus_path)
    _ensure_parent(meta_path)

    with meta_path.open("a+", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_EX)
        handle.seek(0)
        raw = handle.read().strip()

        if raw:
            try:
                state = json.loads(raw)
            except json.JSONDecodeError:
                state = {}
            next_seq = _coerce_int(state.get("next_seq"), default=0)
            if next_seq <= 0:
                next_seq = _scan_max_seq(bus_path) + 1
        else:
            next_seq = _scan_max_seq(bus_path) + 1

        if next_seq <= 0:
            next_seq = 1

        new_state = {"next_seq": next_seq + 1, "updated_at": _now_iso()}
        handle.seek(0)
        handle.truncate()
        handle.write(json.dumps(new_state, ensure_ascii=True))
        handle.flush()
        os.fsync(handle.fileno())
        fcntl.flock(handle.fileno(), fcntl.LOCK_UN)

    return next_seq


def _append_jsonl_record(path: Path, record: dict[str, Any]) -> dict[str, Any]:
    _ensure_parent(path)
    with path.open("a+", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_EX)
        handle.write(json.dumps(record, ensure_ascii=True) + "\n")
        handle.flush()
        os.fsync(handle.fileno())
        fcntl.flock(handle.fileno(), fcntl.LOCK_UN)
    return record


def append_message(
    bus_path: Path,
    sender: str,
    recipient: str,
    message_type: str,
    payload: dict[str, Any],
) -> dict[str, Any]:
    message = {
        "id": str(uuid.uuid4()),
        "seq": _next_seq(bus_path),
        "ts": _now_iso(),
        "from": sender,
        "to": recipient,
        "type": message_type,
        "payload": payload,
    }
    return _append_jsonl_record(bus_path, message)


def append_dlq_entry(
    dlq_path: Path,
    reason: str,
    worker: str,
    failed_message: dict[str, Any],
    context: dict[str, Any],
) -> dict[str, Any]:
    entry = {
        "id": str(uuid.uuid4()),
        "ts": _now_iso(),
        "reason": reason,
        "worker": worker,
        "failed_message": failed_message,
        "context": context,
    }
    return _append_jsonl_record(dlq_path, entry)


def _load_jsonl_records_with_offset(
    path: Path,
    start_offset: int = 0,
) -> tuple[list[dict[str, Any]], int]:
    if not path.exists():
        return [], 0

    records: list[dict[str, Any]] = []
    next_offset = 0
    with path.open("r", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_SH)
        try:
            handle.seek(0, os.SEEK_END)
            file_size = handle.tell()
            seek_offset = start_offset if 0 < start_offset <= file_size else 0
            handle.seek(seek_offset)
            next_offset = seek_offset

            while True:
                line_start = handle.tell()
                raw_line = handle.readline()
                if raw_line == "":
                    next_offset = handle.tell()
                    break

                line = raw_line.strip()
                if not line:
                    next_offset = handle.tell()
                    continue

                try:
                    parsed = json.loads(line)
                except json.JSONDecodeError:
                    # Treat malformed JSON as incomplete tail and retry from this byte offset later.
                    next_offset = line_start
                    break

                if isinstance(parsed, dict) and "id" in parsed:
                    records.append(parsed)
                next_offset = handle.tell()
        finally:
            fcntl.flock(handle.fileno(), fcntl.LOCK_UN)

    return records, next_offset


def _load_jsonl_records(path: Path) -> list[dict[str, Any]]:
    records, _ = _load_jsonl_records_with_offset(path, 0)
    return records


def load_messages(bus_path: Path) -> list[dict[str, Any]]:
    return _load_jsonl_records(bus_path)


def load_messages_since_offset(
    bus_path: Path,
    start_offset: int,
) -> tuple[list[dict[str, Any]], int]:
    return _load_jsonl_records_with_offset(bus_path, start_offset)


def load_dlq_entries(dlq_path: Path) -> list[dict[str, Any]]:
    return _load_jsonl_records(dlq_path)


def _load_checkpoint_state(checkpoint_path: Path) -> dict[str, Any]:
    _ensure_parent(checkpoint_path)
    with checkpoint_path.open("a+", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_EX)
        handle.seek(0)
        raw = handle.read().strip()
        if raw:
            try:
                state = json.loads(raw)
            except json.JSONDecodeError:
                state = {}
        else:
            state = {}

        if not isinstance(state, dict):
            state = {}

        fcntl.flock(handle.fileno(), fcntl.LOCK_UN)
    return state


def _save_checkpoint_state(checkpoint_path: Path, state: dict[str, Any]) -> None:
    _ensure_parent(checkpoint_path)
    with checkpoint_path.open("a+", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_EX)
        handle.seek(0)
        handle.truncate()
        handle.write(json.dumps(state, ensure_ascii=True, indent=2))
        handle.flush()
        os.fsync(handle.fileno())
        fcntl.flock(handle.fileno(), fcntl.LOCK_UN)


def get_consumer_checkpoint(checkpoint_path: Path, consumer: str) -> dict[str, Any]:
    state = _load_checkpoint_state(checkpoint_path)
    entry = state.get(consumer)
    if isinstance(entry, dict):
        return {
            "last_seq": _coerce_int(entry.get("last_seq"), default=0),
            "last_event_id": str(entry.get("last_event_id", "")),
            "last_offset": _coerce_int(entry.get("last_offset"), default=0),
            "updated_at": str(entry.get("updated_at", "")),
        }

    return {"last_seq": 0, "last_event_id": "", "last_offset": 0, "updated_at": ""}


def update_consumer_checkpoint(
    checkpoint_path: Path,
    consumer: str,
    last_seq: int,
    last_event_id: str,
    last_offset: int | None = None,
) -> None:
    _ensure_parent(checkpoint_path)
    with checkpoint_path.open("a+", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_EX)
        handle.seek(0)
        raw = handle.read().strip()
        if raw:
            try:
                state = json.loads(raw)
            except json.JSONDecodeError:
                state = {}
        else:
            state = {}

        if not isinstance(state, dict):
            state = {}

        current_entry = state.get(consumer)
        persisted_offset = 0
        if isinstance(current_entry, dict):
            persisted_offset = _coerce_int(current_entry.get("last_offset"), default=0)

        state[consumer] = {
            "last_seq": max(0, last_seq),
            "last_event_id": last_event_id,
            "last_offset": max(0, last_offset) if last_offset is not None else persisted_offset,
            "updated_at": _now_iso(),
        }

        handle.seek(0)
        handle.truncate()
        handle.write(json.dumps(state, ensure_ascii=True, indent=2))
        handle.flush()
        os.fsync(handle.fileno())
        fcntl.flock(handle.fileno(), fcntl.LOCK_UN)


def load_all_checkpoints(checkpoint_path: Path) -> dict[str, Any]:
    return _load_checkpoint_state(checkpoint_path)


def rotate_bus_log(
    bus_path: Path,
    max_bytes: int,
    archive_dir: Path | None = None,
    compress: bool = True,
) -> dict[str, Any]:
    if max_bytes <= 0:
        raise ValueError("max_bytes must be positive")

    if not bus_path.exists():
        return {"rotated": False, "reason": "bus_missing", "bus": str(bus_path)}

    size_bytes = bus_path.stat().st_size
    if size_bytes < max_bytes:
        return {
            "rotated": False,
            "reason": "below_threshold",
            "bus": str(bus_path),
            "size_bytes": size_bytes,
            "max_bytes": max_bytes,
        }

    target_archive_dir = archive_dir or _default_archive_dir(bus_path)
    _ensure_parent(target_archive_dir / ".keep")

    timestamp = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ")
    archive_plain = target_archive_dir / f"{bus_path.stem}-{timestamp}-{uuid.uuid4().hex}.jsonl"

    archived_events: list[dict[str, Any]] = []
    with bus_path.open("a+", encoding="utf-8") as handle:
        fcntl.flock(handle.fileno(), fcntl.LOCK_EX)
        handle.seek(0)
        archived_raw = handle.read()

        for raw_line in archived_raw.splitlines():
            line = raw_line.strip()
            if not line:
                continue
            try:
                parsed = json.loads(line)
            except json.JSONDecodeError:
                continue
            if isinstance(parsed, dict) and "id" in parsed:
                archived_events.append(parsed)

        if len(archived_events) == 1 and archived_events[0].get("type") == "system.log_compacted":
            fcntl.flock(handle.fileno(), fcntl.LOCK_UN)
            return {
                "rotated": False,
                "reason": "compaction_summary_only",
                "bus": str(bus_path),
                "size_bytes": size_bytes,
                "max_bytes": max_bytes,
            }

        archive_plain.write_text(archived_raw, encoding="utf-8")
        handle.seek(0)
        handle.truncate()
        handle.flush()
        os.fsync(handle.fileno())
        fcntl.flock(handle.fileno(), fcntl.LOCK_UN)

    event_count = len(archived_events)

    archive_path: Path
    if compress:
        archive_path = archive_plain.with_suffix(f"{archive_plain.suffix}.gz")
        with archive_plain.open("rb") as src, gzip.open(archive_path, "wb") as dst:
            shutil.copyfileobj(src, dst)
        archive_plain.unlink()
    else:
        archive_path = archive_plain

    summary_event = append_message(
        bus_path=bus_path,
        sender="system",
        recipient="main",
        message_type="system.log_compacted",
        payload={
            "archived_file": str(archive_path),
            "events_archived": event_count,
            "compressed": compress,
            "previous_size_bytes": size_bytes,
            "max_bytes": max_bytes,
        },
    )

    return {
        "rotated": True,
        "bus": str(bus_path),
        "archive_file": str(archive_path),
        "events_archived": event_count,
        "summary_event_id": summary_event["id"],
    }


def process_once(
    worker_name: str,
    bus_path: Path,
    processed_ids: set[str],
    pending: dict[str, dict[str, Any]],
    max_hops: int = 8,
    pending_timeout_seconds: float = 5.0,
    max_attempts: int = 3,
    retry_base_seconds: float = 0.5,
    dlq_path: Path | None = None,
    checkpoint_path: Path | None = None,
    checkpoint_consumer: str | None = None,
    now_monotonic: float | None = None,
) -> int:
    current_monotonic = time.monotonic() if now_monotonic is None else now_monotonic
    effective_max_attempts = max(1, max_attempts)
    effective_retry_base = max(0.0, retry_base_seconds)
    effective_dlq_path = _default_dlq_path(bus_path) if dlq_path is None else dlq_path

    consumer_name = checkpoint_consumer or worker_name
    effective_checkpoint_path = checkpoint_path
    checkpoint_last_seq = 0
    checkpoint_last_event_id = ""
    checkpoint_last_offset = 0

    if effective_checkpoint_path is not None:
        checkpoint_entry = get_consumer_checkpoint(effective_checkpoint_path, consumer_name)
        checkpoint_last_seq = _coerce_int(checkpoint_entry.get("last_seq"), default=0)
        checkpoint_last_event_id = str(checkpoint_entry.get("last_event_id", ""))
        checkpoint_last_offset = _coerce_int(checkpoint_entry.get("last_offset"), default=0)

    handled = 0

    if worker_name == "alpha":
        handled += _maintain_pending_for_alpha(
            bus_path=bus_path,
            pending=pending,
            now_monotonic=current_monotonic,
            pending_timeout_seconds=pending_timeout_seconds,
            max_attempts=effective_max_attempts,
            retry_base_seconds=effective_retry_base,
            dlq_path=effective_dlq_path,
        )

    highest_seen_seq = checkpoint_last_seq
    latest_event_id = checkpoint_last_event_id

    messages: list[dict[str, Any]]
    next_offset = checkpoint_last_offset
    if effective_checkpoint_path is not None:
        messages, next_offset = load_messages_since_offset(bus_path, checkpoint_last_offset)
    else:
        messages = load_messages(bus_path)

    for message in messages:
        message_id = str(message.get("id", ""))
        message_seq = _coerce_int(message.get("seq"), default=0)

        if not message_id or message_id in processed_ids:
            continue

        if message_seq > 0 and message_seq <= checkpoint_last_seq:
            continue

        if message.get("to") != worker_name:
            if message_seq > highest_seen_seq:
                highest_seen_seq = message_seq
                latest_event_id = message_id
            continue

        try:
            _handle_message(
                worker_name=worker_name,
                bus_path=bus_path,
                message=message,
                pending=pending,
                max_hops=max_hops,
                now_monotonic=current_monotonic,
                max_attempts=effective_max_attempts,
                retry_base_seconds=effective_retry_base,
                dlq_path=effective_dlq_path,
            )
        except Exception as exc:
            append_message(
                bus_path=bus_path,
                sender=worker_name,
                recipient="main",
                message_type="guard.processing_error",
                payload={
                    "worker": worker_name,
                    "message_id": message_id,
                    "message_type": str(message.get("type", "")),
                    "error": str(exc),
                },
            )
            append_dlq_entry(
                dlq_path=effective_dlq_path,
                reason="worker.processing_error",
                worker=worker_name,
                failed_message={
                    "id": message_id,
                    "from": str(message.get("from", "")),
                    "to": str(message.get("to", "")),
                    "type": str(message.get("type", "")),
                    "payload": message.get("payload") if isinstance(message.get("payload"), dict) else {},
                },
                context={
                    "error": str(exc),
                    "seq": message_seq,
                },
            )
            processed_ids.add(message_id)
            if message_seq > highest_seen_seq:
                highest_seen_seq = message_seq
                latest_event_id = message_id
            continue

        processed_ids.add(message_id)
        handled += 1

        if message_seq > highest_seen_seq:
            highest_seen_seq = message_seq
            latest_event_id = message_id

    should_update_checkpoint = (
        effective_checkpoint_path is not None
        and (highest_seen_seq > checkpoint_last_seq or next_offset != checkpoint_last_offset)
    )
    if should_update_checkpoint:
        update_consumer_checkpoint(
            checkpoint_path=effective_checkpoint_path,
            consumer=consumer_name,
            last_seq=highest_seen_seq,
            last_event_id=latest_event_id,
            last_offset=next_offset,
        )

    return handled


def _handle_message(
    worker_name: str,
    bus_path: Path,
    message: dict[str, Any],
    pending: dict[str, dict[str, Any]],
    max_hops: int,
    now_monotonic: float,
    max_attempts: int,
    retry_base_seconds: float,
    dlq_path: Path,
) -> None:
    if message.get("type") == "chat.message":
        _handle_chat_message(
            worker_name=worker_name,
            bus_path=bus_path,
            message=message,
            max_hops=max_hops,
            dlq_path=dlq_path,
        )
        return

    if worker_name == "alpha":
        _handle_alpha(
            bus_path=bus_path,
            message=message,
            pending=pending,
            now_monotonic=now_monotonic,
            max_attempts=max_attempts,
            retry_base_seconds=retry_base_seconds,
        )
        return

    if worker_name == "beta":
        _handle_beta(bus_path, message)
        return

    append_message(
        bus_path=bus_path,
        sender=worker_name,
        recipient=str(message.get("from", "main")),
        message_type="echo",
        payload={"seen_type": message.get("type")},
    )


def _handle_alpha(
    bus_path: Path,
    message: dict[str, Any],
    pending: dict[str, dict[str, Any]],
    now_monotonic: float,
    max_attempts: int,
    retry_base_seconds: float,
) -> None:
    message_type = message.get("type")
    payload = message.get("payload") or {}

    if message_type == "task.add":
        a = int(payload.get("a", 0))
        b = int(payload.get("b", 0))
        correlation_id = str(message["id"])

        pending[correlation_id] = {
            "requester": str(message.get("from", "main")),
            "a": a,
            "b": b,
            "started_at_monotonic": now_monotonic,
            "attempts": 1,
            "retry_due_at_monotonic": now_monotonic
            + _retry_delay_seconds(1, retry_base_seconds),
            "max_attempts": max_attempts,
        }

        _send_verify_add(
            bus_path=bus_path,
            correlation_id=correlation_id,
            a=a,
            b=b,
            attempt=1,
        )
        return

    if message_type == "verify.result":
        correlation_id = str(payload.get("correlation_id", ""))
        if not correlation_id or correlation_id not in pending:
            return

        request_context = pending.pop(correlation_id)
        append_message(
            bus_path=bus_path,
            sender="alpha",
            recipient=str(request_context["requester"]),
            message_type="result.add",
            payload={
                "a": request_context["a"],
                "b": request_context["b"],
                "sum": int(payload.get("sum", 0)),
                "verified_by": "beta",
                "attempts": int(request_context.get("attempts", 1)),
            },
        )


def _send_verify_add(
    bus_path: Path,
    correlation_id: str,
    a: int,
    b: int,
    attempt: int,
) -> None:
    append_message(
        bus_path=bus_path,
        sender="alpha",
        recipient="beta",
        message_type="verify.add",
        payload={
            "correlation_id": correlation_id,
            "a": a,
            "b": b,
            "attempt": attempt,
        },
    )


def _handle_beta(bus_path: Path, message: dict[str, Any]) -> None:
    if message.get("type") != "verify.add":
        return

    payload = message.get("payload") or {}
    a = int(payload.get("a", 0))
    b = int(payload.get("b", 0))
    append_message(
        bus_path=bus_path,
        sender="beta",
        recipient="alpha",
        message_type="verify.result",
        payload={
            "correlation_id": str(payload.get("correlation_id", "")),
            "sum": a + b,
        },
    )


def _handle_chat_message(
    worker_name: str,
    bus_path: Path,
    message: dict[str, Any],
    max_hops: int,
    dlq_path: Path,
) -> None:
    payload = message.get("payload") or {}
    text = str(payload.get("text", "")).strip()
    origin = str(payload.get("origin") or message.get("from", "main")).strip()
    origin_key = origin.lower()
    worker_key = worker_name.lower()
    hops = _coerce_int(payload.get("hops"), default=0)
    team = _resolve_team(payload)
    mentions = _resolve_mentions(text=text, payload_mentions=payload.get("mentions"))

    if mentions:
        next_target = mentions[0]
        remaining_mentions = mentions[1:]

        if hops >= max_hops:
            append_message(
                bus_path=bus_path,
                sender=worker_name,
                recipient=origin,
                message_type="guard.max_hops",
                payload={
                    "stopped_by": worker_name,
                    "text": text,
                    "hops": hops,
                    "max_hops": max_hops,
                },
            )
            return

        if next_target == "all":
            forwarded_text = _strip_first_mention(text)
            targets = [
                member
                for member in team
                if member not in {worker_key, origin_key, "main"}
            ]

            if not targets:
                append_message(
                    bus_path=bus_path,
                    sender=worker_name,
                    recipient=origin,
                    message_type="chat.reply",
                    payload={
                        "text": f"{worker_name} handled: no targets for @all",
                        "hops": hops,
                    },
                )
                return

            for target in targets:
                next_payload = {
                    "text": forwarded_text,
                    "origin": origin,
                    "hops": hops + 1,
                    "team": team,
                }
                if remaining_mentions:
                    next_payload["mentions"] = remaining_mentions

                append_message(
                    bus_path=bus_path,
                    sender=worker_name,
                    recipient=target,
                    message_type="chat.message",
                    payload=next_payload,
                )
            return

        if next_target not in team:
            append_message(
                bus_path=bus_path,
                sender=worker_name,
                recipient=origin,
                message_type="guard.unknown_target",
                payload={
                    "target": next_target,
                    "text": text,
                    "team": team,
                },
            )
            append_dlq_entry(
                dlq_path=dlq_path,
                reason="chat.unknown_target",
                worker=worker_name,
                failed_message={
                    "id": message.get("id"),
                    "from": message.get("from"),
                    "to": message.get("to"),
                    "type": message.get("type"),
                    "payload": payload,
                },
                context={"target": next_target, "team": team},
            )
            return

        next_payload = {
            "text": _strip_first_mention(text),
            "origin": origin,
            "hops": hops + 1,
            "team": team,
        }
        if remaining_mentions:
            next_payload["mentions"] = remaining_mentions

        append_message(
            bus_path=bus_path,
            sender=worker_name,
            recipient=next_target,
            message_type="chat.message",
            payload=next_payload,
        )
        return

    append_message(
        bus_path=bus_path,
        sender=worker_name,
        recipient=origin,
        message_type="chat.reply",
        payload={
            "text": f"{worker_name} handled: {text if text else '(empty)'}",
            "hops": hops,
        },
    )


def _maintain_pending_for_alpha(
    bus_path: Path,
    pending: dict[str, dict[str, Any]],
    now_monotonic: float,
    pending_timeout_seconds: float,
    max_attempts: int,
    retry_base_seconds: float,
    dlq_path: Path,
) -> int:
    if not pending:
        return 0

    generated = 0
    failures: list[tuple[str, dict[str, Any], str]] = []

    for correlation_id, request_context in list(pending.items()):
        started = float(request_context.get("started_at_monotonic", now_monotonic))

        if pending_timeout_seconds > 0 and now_monotonic - started >= pending_timeout_seconds:
            failures.append((correlation_id, request_context, "beta.timeout"))
            continue

        attempts = _coerce_int(request_context.get("attempts"), default=1)
        retry_due = float(
            request_context.get(
                "retry_due_at_monotonic",
                now_monotonic + _retry_delay_seconds(attempts, retry_base_seconds),
            )
        )

        if now_monotonic >= retry_due:
            if attempts < max_attempts:
                next_attempt = attempts + 1
                request_context["attempts"] = next_attempt
                request_context["retry_due_at_monotonic"] = now_monotonic + _retry_delay_seconds(
                    next_attempt, retry_base_seconds
                )
                _send_verify_add(
                    bus_path=bus_path,
                    correlation_id=correlation_id,
                    a=int(request_context.get("a", 0)),
                    b=int(request_context.get("b", 0)),
                    attempt=next_attempt,
                )
                generated += 1
            else:
                failures.append((correlation_id, request_context, "beta.retry_exhausted"))

    for correlation_id, request_context, reason in failures:
        if correlation_id not in pending:
            continue
        pending.pop(correlation_id)
        _emit_alpha_failure(
            bus_path=bus_path,
            dlq_path=dlq_path,
            correlation_id=correlation_id,
            request_context=request_context,
            reason=reason,
            max_attempts=max_attempts,
        )
        generated += 1

    return generated


def _emit_alpha_failure(
    bus_path: Path,
    dlq_path: Path,
    correlation_id: str,
    request_context: dict[str, Any],
    reason: str,
    max_attempts: int,
) -> None:
    requester = str(request_context.get("requester", "main"))
    a = int(request_context.get("a", 0))
    b = int(request_context.get("b", 0))
    attempts = _coerce_int(request_context.get("attempts"), default=1)

    append_message(
        bus_path=bus_path,
        sender="alpha",
        recipient=requester,
        message_type="result.timeout",
        payload={
            "correlation_id": correlation_id,
            "a": a,
            "b": b,
            "reason": reason,
            "attempts": attempts,
            "max_attempts": max_attempts,
        },
    )

    append_dlq_entry(
        dlq_path=dlq_path,
        reason=reason,
        worker="alpha",
        failed_message={
            "id": correlation_id,
            "from": requester,
            "to": "alpha",
            "type": "task.add",
            "payload": {"a": a, "b": b},
        },
        context={
            "requester": requester,
            "attempts": attempts,
            "max_attempts": max_attempts,
        },
    )


def _resolve_team(payload: dict[str, Any]) -> list[str]:
    raw_team = payload.get("team")
    if isinstance(raw_team, list):
        candidates = [str(member).strip().lower() for member in raw_team]
    else:
        candidates = list(DEFAULT_TEAM)

    team: list[str] = []
    seen: set[str] = set()
    for candidate in candidates:
        if not candidate or candidate in seen:
            continue
        seen.add(candidate)
        team.append(candidate)

    return team if team else list(DEFAULT_TEAM)


def _retry_delay_seconds(attempt: int, retry_base_seconds: float) -> float:
    exponent = max(0, attempt - 1)
    return retry_base_seconds * (2**exponent)


def _parse_mentions(text: str) -> list[str]:
    return [match.group(1).lower() for match in MENTION_PATTERN.finditer(text)]


def _resolve_mentions(text: str, payload_mentions: Any) -> list[str]:
    mentions: list[str] = []
    seen: set[str] = set()

    for mention in _parse_mentions(text):
        if mention and mention not in seen:
            seen.add(mention)
            mentions.append(mention)

    if isinstance(payload_mentions, list):
        for item in payload_mentions:
            mention = str(item).strip().lower()
            if mention and mention not in seen:
                seen.add(mention)
                mentions.append(mention)

    return mentions


def _strip_first_mention(text: str) -> str:
    return MENTION_PATTERN.sub("", text, count=1).strip()


def _coerce_int(value: Any, default: int) -> int:
    try:
        return int(value)
    except (TypeError, ValueError):
        return default


def run_worker(
    worker_name: str,
    bus_path: Path,
    duration_seconds: float,
    poll_seconds: float,
    max_hops: int = 8,
    pending_timeout_seconds: float = 5.0,
    max_attempts: int = 3,
    retry_base_seconds: float = 0.5,
    dlq_path: Path | None = None,
    checkpoint_path: Path | None = None,
    checkpoint_consumer: str | None = None,
    emit_status_events: bool = True,
) -> dict[str, Any]:
    processed_ids: set[str] = set()
    pending: dict[str, dict[str, Any]] = {}
    deadline = time.monotonic() + duration_seconds
    handled_total = 0

    effective_checkpoint_path = (
        _default_checkpoint_path(bus_path) if checkpoint_path is None else checkpoint_path
    )
    effective_consumer = checkpoint_consumer or f"worker:{worker_name}"

    if emit_status_events:
        append_message(
            bus_path=bus_path,
            sender=worker_name,
            recipient="main",
            message_type="worker.status",
            payload={"state": "started", "consumer": effective_consumer},
        )

    while time.monotonic() < deadline:
        handled = process_once(
            worker_name=worker_name,
            bus_path=bus_path,
            processed_ids=processed_ids,
            pending=pending,
            max_hops=max_hops,
            pending_timeout_seconds=pending_timeout_seconds,
            max_attempts=max_attempts,
            retry_base_seconds=retry_base_seconds,
            dlq_path=dlq_path,
            checkpoint_path=effective_checkpoint_path,
            checkpoint_consumer=effective_consumer,
        )
        handled_total += handled
        if handled == 0:
            time.sleep(poll_seconds)

    summary = {
        "worker": worker_name,
        "handled_messages": handled_total,
        "pending": len(pending),
        "consumer": effective_consumer,
    }

    if emit_status_events:
        append_message(
            bus_path=bus_path,
            sender=worker_name,
            recipient="main",
            message_type="worker.status",
            payload={"state": "stopped", **summary},
        )

    return summary


def _parse_payload(raw_payload: str) -> dict[str, Any]:
    if not raw_payload:
        return {}
    try:
        parsed = json.loads(raw_payload)
    except json.JSONDecodeError as exc:
        raise ValueError("payload must be a JSON object") from exc
    if not isinstance(parsed, dict):
        raise ValueError("payload must be a JSON object")
    return parsed


def main() -> int:
    parser = argparse.ArgumentParser(description="Teammate relay bus")
    subparsers = parser.add_subparsers(dest="command", required=True)

    send_parser = subparsers.add_parser("send")
    send_parser.add_argument("--bus", required=True)
    send_parser.add_argument("--from", dest="sender", required=True)
    send_parser.add_argument("--to", required=True)
    send_parser.add_argument("--type", required=True)
    send_parser.add_argument("--payload", default="{}")

    list_parser = subparsers.add_parser("list")
    list_parser.add_argument("--bus", required=True)

    dlq_list_parser = subparsers.add_parser("dlq-list")
    dlq_list_parser.add_argument("--dlq", required=True)

    checkpoint_list_parser = subparsers.add_parser("checkpoint-list")
    checkpoint_list_parser.add_argument("--checkpoint", required=True)

    reset_parser = subparsers.add_parser("reset")
    reset_parser.add_argument("--bus", required=True)

    rotate_parser = subparsers.add_parser("rotate")
    rotate_parser.add_argument("--bus", required=True)
    rotate_parser.add_argument("--max-bytes", required=True, type=int)
    rotate_parser.add_argument("--archive-dir", default="")
    rotate_parser.add_argument("--no-compress", action="store_true")

    worker_parser = subparsers.add_parser("worker")
    worker_parser.add_argument("--name", required=True)
    worker_parser.add_argument("--bus", required=True)
    worker_parser.add_argument("--duration", type=float, default=20.0)
    worker_parser.add_argument("--poll", type=float, default=0.2)
    worker_parser.add_argument("--max-hops", type=int, default=8)
    worker_parser.add_argument("--pending-timeout", type=float, default=5.0)
    worker_parser.add_argument("--max-attempts", type=int, default=3)
    worker_parser.add_argument("--retry-base", type=float, default=0.5)
    worker_parser.add_argument("--dlq", default="")
    worker_parser.add_argument("--checkpoint", default="")
    worker_parser.add_argument("--checkpoint-consumer", default="")
    worker_parser.add_argument("--no-status-events", action="store_true")

    args = parser.parse_args()

    if args.command == "send":
        bus_path = Path(args.bus)
        try:
            payload = _parse_payload(args.payload)
        except ValueError as exc:
            print(str(exc), file=sys.stderr)
            return 1
        message = append_message(
            bus_path=bus_path,
            sender=args.sender,
            recipient=args.to,
            message_type=args.type,
            payload=payload,
        )
        print(json.dumps(message, ensure_ascii=True))
        return 0

    if args.command == "list":
        bus_path = Path(args.bus)
        print(json.dumps(load_messages(bus_path), ensure_ascii=True, indent=2))
        return 0

    if args.command == "dlq-list":
        dlq_path = Path(args.dlq)
        print(json.dumps(load_dlq_entries(dlq_path), ensure_ascii=True, indent=2))
        return 0

    if args.command == "checkpoint-list":
        checkpoint_path = Path(args.checkpoint)
        print(json.dumps(load_all_checkpoints(checkpoint_path), ensure_ascii=True, indent=2))
        return 0

    if args.command == "reset":
        bus_path = Path(args.bus)
        _ensure_parent(bus_path)
        bus_path.write_text("", encoding="utf-8")
        print(json.dumps({"bus": str(bus_path), "status": "reset"}, ensure_ascii=True))
        return 0

    if args.command == "rotate":
        bus_path = Path(args.bus)
        archive_dir = Path(args.archive_dir) if args.archive_dir else None
        result = rotate_bus_log(
            bus_path=bus_path,
            max_bytes=args.max_bytes,
            archive_dir=archive_dir,
            compress=not args.no_compress,
        )
        print(json.dumps(result, ensure_ascii=True))
        return 0

    if args.command == "worker":
        bus_path = Path(args.bus)
        dlq_path = Path(args.dlq) if args.dlq else None
        checkpoint_path = Path(args.checkpoint) if args.checkpoint else None
        checkpoint_consumer = args.checkpoint_consumer if args.checkpoint_consumer else None

        summary = run_worker(
            worker_name=args.name,
            bus_path=bus_path,
            duration_seconds=args.duration,
            poll_seconds=args.poll,
            max_hops=args.max_hops,
            pending_timeout_seconds=args.pending_timeout,
            max_attempts=args.max_attempts,
            retry_base_seconds=args.retry_base,
            dlq_path=dlq_path,
            checkpoint_path=checkpoint_path,
            checkpoint_consumer=checkpoint_consumer,
            emit_status_events=not args.no_status_events,
        )
        print(json.dumps(summary, ensure_ascii=True))
        return 0

    return 1


if __name__ == "__main__":
    raise SystemExit(main())
