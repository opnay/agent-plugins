"""Non-blocking main orchestrator loop for teammate workflow.

This script intentionally avoids agent wait-style blocking and drives work via
filesystem watch with polling fallback plus subprocess health polling.
"""

from __future__ import annotations

import argparse
import json
import subprocess
import sys
import threading
import time
from pathlib import Path
from typing import Any, Callable

import team_bus


class _PollingWaiter:
    backend = "poll"

    def wait(self, timeout_seconds: float) -> bool:
        if timeout_seconds > 0:
            time.sleep(timeout_seconds)
        return False

    def close(self) -> None:
        return None


class _ThreadedEventWaiter:
    def __init__(
        self,
        backend: str,
        changed_event: threading.Event,
        stop_callback: Callable[[], None],
    ) -> None:
        self.backend = backend
        self._changed_event = changed_event
        self._stop_callback = stop_callback

    def wait(self, timeout_seconds: float) -> bool:
        waited = self._changed_event.wait(timeout=max(0.0, timeout_seconds))
        if waited:
            self._changed_event.clear()
        return waited

    def close(self) -> None:
        self._stop_callback()


def _create_watchfiles_waiter(
    bus_path: Path,
    poll_seconds: float,
) -> _ThreadedEventWaiter | None:
    try:
        from watchfiles import watch
    except Exception:
        return None

    changed_event = threading.Event()
    stop_event = threading.Event()
    bus_name = bus_path.name
    target_path = bus_path.resolve()
    watch_dir = bus_path.parent.resolve()
    step_ms = max(50, int(max(0.05, poll_seconds) * 1000))

    def _watch() -> None:
        try:
            for changes in watch(
                str(watch_dir),
                recursive=False,
                stop_event=stop_event,
                yield_on_timeout=True,
                debounce=0,
                step=step_ms,
            ):
                if stop_event.is_set():
                    break
                if not changes:
                    continue
                for _change, raw_path in changes:
                    candidate = Path(str(raw_path))
                    if candidate.name == bus_name:
                        changed_event.set()
                        break
                    try:
                        if candidate.resolve() == target_path:
                            changed_event.set()
                            break
                    except OSError:
                        continue
        except Exception:
            return

    thread = threading.Thread(
        target=_watch,
        name="teammate-watchfiles",
        daemon=True,
    )
    thread.start()

    def _stop() -> None:
        stop_event.set()
        thread.join(timeout=2.0)

    return _ThreadedEventWaiter(
        backend="watchfiles",
        changed_event=changed_event,
        stop_callback=_stop,
    )


def _build_bus_waiter(
    bus_path: Path,
    poll_seconds: float,
    watch_mode: str,
) -> _PollingWaiter | _ThreadedEventWaiter:
    mode = watch_mode.strip().lower()
    if mode not in {"auto", "watch", "poll"}:
        mode = "auto"

    if mode == "poll":
        return _PollingWaiter()

    watcher = _create_watchfiles_waiter(bus_path, poll_seconds=poll_seconds)
    if watcher is not None:
        return watcher

    return _PollingWaiter()


def _coerce_float(value: Any, default: float) -> float:
    try:
        return float(value)
    except (TypeError, ValueError):
        return default


def _coerce_int(value: Any, default: int) -> int:
    try:
        return int(value)
    except (TypeError, ValueError):
        return default


def _format_event(event: dict[str, Any]) -> str:
    seq = _coerce_int(event.get("seq"), default=0)
    event_type = str(event.get("type", ""))
    sender = str(event.get("from", ""))
    recipient = str(event.get("to", ""))
    payload = event.get("payload") or {}
    return json.dumps(
        {
            "seq": seq,
            "from": sender,
            "to": recipient,
            "type": event_type,
            "payload": payload,
        },
        ensure_ascii=True,
    )


def _collect_new_events(
    bus_path: Path,
    last_seq: int,
    last_offset: int,
) -> tuple[list[dict[str, Any]], int, int]:
    events = []
    highest_seq = last_seq

    messages, next_offset = team_bus.load_messages_since_offset(bus_path, last_offset)
    for message in messages:
        seq = _coerce_int(message.get("seq"), default=0)
        if seq <= last_seq:
            continue
        events.append(message)
        if seq > highest_seq:
            highest_seq = seq

    events.sort(key=lambda event: _coerce_int(event.get("seq"), default=0))
    return events, highest_seq, next_offset


def _resolve_mentions(text: str, payload_mentions: Any) -> list[str]:
    mentions = team_bus._resolve_mentions(text=text, payload_mentions=payload_mentions)
    return [mention for mention in mentions if mention]


def _strip_all_mentions(text: str) -> str:
    return team_bus.MENTION_PATTERN.sub("", text).strip()


def _resolve_targets(
    mentions: list[str],
    team: list[str],
    sender: str,
) -> list[str]:
    sender_key = sender.lower()
    candidates: list[str] = []

    for mention in mentions:
        if mention == "all":
            for member in team:
                if member in {sender_key, "main", "orchestrator"}:
                    continue
                if member not in candidates:
                    candidates.append(member)
            continue
        if mention in team and mention not in {sender_key, "main", "orchestrator"}:
            if mention not in candidates:
                candidates.append(mention)

    return candidates


def _dispatch_route_request(
    bus_path: Path,
    event: dict[str, Any],
) -> int:
    payload = event.get("payload") or {}
    text = str(payload.get("text", "")).strip()
    cleaned_text = _strip_all_mentions(text)
    team = team_bus._resolve_team(payload)
    mentions = _resolve_mentions(text=text, payload_mentions=payload.get("mentions"))
    sender = str(event.get("from", "main"))

    targets = _resolve_targets(mentions=mentions, team=team, sender=sender)
    if not targets:
        team_bus.append_message(
            bus_path=bus_path,
            sender="orchestrator",
            recipient=sender,
            message_type="orchestrator.no_target",
            payload={"text": text, "team": team},
        )
        return 1

    dispatch_count = 0
    for target in targets:
        team_bus.append_message(
            bus_path=bus_path,
            sender="orchestrator",
            recipient=target,
            message_type="chat.message",
            payload={
                "text": cleaned_text if cleaned_text else text,
                "origin": sender,
                "hops": 0,
                "team": team,
            },
        )
        team_bus.append_message(
            bus_path=bus_path,
            sender="orchestrator",
            recipient="main",
            message_type="orchestrator.dispatched",
            payload={
                "source_event_id": str(event.get("id", "")),
                "target": target,
                "text": text,
            },
        )
        dispatch_count += 2

    return dispatch_count


def _handle_main_events(bus_path: Path, events: list[dict[str, Any]], verbose: bool) -> int:
    generated = 0

    for event in events:
        if verbose:
            print(_format_event(event), flush=True)

        if event.get("to") != "main":
            continue

        if event.get("type") == "route.request":
            generated += _dispatch_route_request(bus_path=bus_path, event=event)

    return generated


def _spawn_workers(
    script_path: Path,
    bus_path: Path,
    dlq_path: Path,
    checkpoint_path: Path,
    workers: list[str],
    worker_duration: float,
    worker_poll: float,
    max_hops: int,
    pending_timeout: float,
    max_attempts: int,
    retry_base: float,
) -> dict[str, subprocess.Popen[str]]:
    procs: dict[str, subprocess.Popen[str]] = {}

    for worker in workers:
        cmd = [
            sys.executable,
            str(script_path),
            "worker",
            "--name",
            worker,
            "--bus",
            str(bus_path),
            "--duration",
            str(worker_duration),
            "--poll",
            str(worker_poll),
            "--max-hops",
            str(max_hops),
            "--pending-timeout",
            str(pending_timeout),
            "--max-attempts",
            str(max_attempts),
            "--retry-base",
            str(retry_base),
            "--dlq",
            str(dlq_path),
            "--checkpoint",
            str(checkpoint_path),
            "--checkpoint-consumer",
            f"worker:{worker}",
        ]

        procs[worker] = subprocess.Popen(
            cmd,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
        )

    return procs


def _poll_worker_processes(
    bus_path: Path,
    processes: dict[str, subprocess.Popen[str]],
    verbose: bool,
) -> int:
    closed = 0

    for worker, proc in list(processes.items()):
        if proc.poll() is None:
            continue

        try:
            stdout_text, stderr_text = proc.communicate(timeout=1)
        except subprocess.TimeoutExpired:
            stdout_text, stderr_text = "", ""
        team_bus.append_message(
            bus_path=bus_path,
            sender="orchestrator",
            recipient="main",
            message_type="orchestrator.worker_exit",
            payload={
                "worker": worker,
                "returncode": proc.returncode,
                "stdout": stdout_text.strip(),
                "stderr": stderr_text.strip(),
            },
        )
        if verbose:
            print(
                json.dumps(
                    {
                        "worker": worker,
                        "returncode": proc.returncode,
                        "stdout": stdout_text.strip(),
                        "stderr": stderr_text.strip(),
                    },
                    ensure_ascii=True,
                ),
                flush=True,
            )

        del processes[worker]
        closed += 1

    return closed


def _normalize_process_output(value: Any) -> str:
    if value is None:
        return ""
    if isinstance(value, bytes):
        return value.decode("utf-8", "replace")
    return str(value)


def _terminate_collect_output(
    proc: Any,
    timeout_seconds: float,
) -> tuple[str, str]:
    proc.terminate()
    try:
        stdout_text, stderr_text = proc.communicate(timeout=timeout_seconds)
    except subprocess.TimeoutExpired:
        proc.kill()
        try:
            stdout_text, stderr_text = proc.communicate(timeout=timeout_seconds)
        except subprocess.TimeoutExpired:
            return "", "process output unavailable: terminate+kill timed out"

    return _normalize_process_output(stdout_text), _normalize_process_output(stderr_text)


def run_orchestrator(
    bus_path: Path,
    dlq_path: Path,
    checkpoint_path: Path,
    duration_seconds: float,
    poll_seconds: float,
    watch_mode: str,
    rotate_max_bytes: int,
    rotate_archive_dir: Path,
    spawn_workers: bool,
    workers: list[str],
    worker_duration: float,
    worker_poll: float,
    max_hops: int,
    pending_timeout: float,
    max_attempts: int,
    retry_base: float,
    verbose: bool,
) -> dict[str, Any]:
    start = time.monotonic()
    deadline = start + duration_seconds

    team_bus._ensure_parent(bus_path)
    team_bus._ensure_parent(dlq_path)
    team_bus._ensure_parent(checkpoint_path)

    team_bus.append_message(
        bus_path=bus_path,
        sender="orchestrator",
        recipient="main",
        message_type="orchestrator.started",
        payload={"workers": workers, "duration_seconds": duration_seconds},
    )

    script_path = Path(team_bus.__file__).resolve()
    processes: dict[str, subprocess.Popen[str]] = {}
    if spawn_workers:
        processes = _spawn_workers(
            script_path=script_path,
            bus_path=bus_path,
            dlq_path=dlq_path,
            checkpoint_path=checkpoint_path,
            workers=workers,
            worker_duration=worker_duration,
            worker_poll=worker_poll,
            max_hops=max_hops,
            pending_timeout=pending_timeout,
            max_attempts=max_attempts,
            retry_base=retry_base,
        )

    consumer = "orchestrator:main"
    checkpoint_entry = team_bus.get_consumer_checkpoint(checkpoint_path, consumer)
    last_seq = _coerce_int(checkpoint_entry.get("last_seq"), default=0)
    last_event_id = str(checkpoint_entry.get("last_event_id", ""))
    last_offset = _coerce_int(checkpoint_entry.get("last_offset"), default=0)
    loops = 0
    generated_events = 0
    last_rotation_monotonic = 0.0
    rotation_min_interval_seconds = max(1.0, poll_seconds)
    waiter = _build_bus_waiter(
        bus_path=bus_path,
        poll_seconds=poll_seconds,
        watch_mode=watch_mode,
    )

    try:
        while time.monotonic() < deadline:
            loops += 1

            previous_offset = last_offset
            events, highest_seq, last_offset = _collect_new_events(
                bus_path=bus_path,
                last_seq=last_seq,
                last_offset=last_offset,
            )
            if events:
                generated_events += _handle_main_events(
                    bus_path=bus_path,
                    events=events,
                    verbose=verbose,
                )
                last_seq = highest_seq
                last_event_id = str(events[-1].get("id", ""))
                team_bus.update_consumer_checkpoint(
                    checkpoint_path=checkpoint_path,
                    consumer=consumer,
                    last_seq=last_seq,
                    last_event_id=last_event_id,
                    last_offset=last_offset,
                )
            elif last_offset != previous_offset:
                team_bus.update_consumer_checkpoint(
                    checkpoint_path=checkpoint_path,
                    consumer=consumer,
                    last_seq=last_seq,
                    last_event_id=last_event_id,
                    last_offset=last_offset,
                )

            _poll_worker_processes(bus_path=bus_path, processes=processes, verbose=verbose)

            now_monotonic = time.monotonic()
            if (
                rotate_max_bytes > 0
                and now_monotonic - last_rotation_monotonic >= rotation_min_interval_seconds
            ):
                team_bus.rotate_bus_log(
                    bus_path=bus_path,
                    max_bytes=rotate_max_bytes,
                    archive_dir=rotate_archive_dir,
                    compress=True,
                )
                last_rotation_monotonic = now_monotonic

            if not processes and spawn_workers and time.monotonic() - start > 1.0:
                break

            remaining_seconds = max(0.0, deadline - time.monotonic())
            if remaining_seconds == 0:
                break
            waiter.wait(min(poll_seconds, remaining_seconds))
    finally:
        waiter.close()

    for worker, proc in list(processes.items()):
        if proc.poll() is None:
            stdout_text, stderr_text = _terminate_collect_output(proc, timeout_seconds=2.0)
            team_bus.append_message(
                bus_path=bus_path,
                sender="orchestrator",
                recipient="main",
                message_type="orchestrator.worker_terminated",
                payload={
                    "worker": worker,
                    "stdout": stdout_text.strip(),
                    "stderr": stderr_text.strip(),
                },
            )

    summary = {
        "status": "completed",
        "loops": loops,
        "generated_events": generated_events,
        "consumer": consumer,
        "last_seq": last_seq,
        "wait_backend": waiter.backend,
        "watch_mode": watch_mode.strip().lower(),
    }

    team_bus.append_message(
        bus_path=bus_path,
        sender="orchestrator",
        recipient="main",
        message_type="orchestrator.stopped",
        payload=summary,
    )

    return summary


def main() -> int:
    parser = argparse.ArgumentParser(description="Non-blocking orchestrator loop")
    parser.add_argument("--bus", required=True)
    parser.add_argument("--dlq", required=True)
    parser.add_argument("--checkpoint", required=True)
    parser.add_argument("--duration", type=float, default=30.0)
    parser.add_argument("--poll", type=float, default=0.5)
    parser.add_argument("--watch-mode", choices=["auto", "watch", "poll"], default="auto")
    parser.add_argument("--rotate-max-bytes", type=int, default=0)
    parser.add_argument("--rotate-archive-dir", default="")
    parser.add_argument("--workers", default="alpha,beta,gamma")
    parser.add_argument("--spawn-workers", action="store_true")
    parser.add_argument("--worker-duration", type=float, default=20.0)
    parser.add_argument("--worker-poll", type=float, default=0.2)
    parser.add_argument("--max-hops", type=int, default=8)
    parser.add_argument("--pending-timeout", type=float, default=5.0)
    parser.add_argument("--max-attempts", type=int, default=3)
    parser.add_argument("--retry-base", type=float, default=0.5)
    parser.add_argument("--verbose", action="store_true")
    args = parser.parse_args()

    bus_path = Path(args.bus)
    dlq_path = Path(args.dlq)
    checkpoint_path = Path(args.checkpoint)
    archive_dir = (
        Path(args.rotate_archive_dir)
        if args.rotate_archive_dir
        else team_bus._default_archive_dir(bus_path)
    )

    workers = [w.strip().lower() for w in args.workers.split(",") if w.strip()]
    if not workers:
        workers = ["alpha", "beta"]

    summary = run_orchestrator(
        bus_path=bus_path,
        dlq_path=dlq_path,
        checkpoint_path=checkpoint_path,
        duration_seconds=max(1.0, _coerce_float(args.duration, 30.0)),
        poll_seconds=max(0.05, _coerce_float(args.poll, 0.5)),
        watch_mode=args.watch_mode,
        rotate_max_bytes=max(0, _coerce_int(args.rotate_max_bytes, 0)),
        rotate_archive_dir=archive_dir,
        spawn_workers=bool(args.spawn_workers),
        workers=workers,
        worker_duration=max(1.0, _coerce_float(args.worker_duration, 20.0)),
        worker_poll=max(0.05, _coerce_float(args.worker_poll, 0.2)),
        max_hops=max(1, _coerce_int(args.max_hops, 8)),
        pending_timeout=max(0.1, _coerce_float(args.pending_timeout, 5.0)),
        max_attempts=max(1, _coerce_int(args.max_attempts, 3)),
        retry_base=max(0.0, _coerce_float(args.retry_base, 0.5)),
        verbose=bool(args.verbose),
    )

    print(json.dumps(summary, ensure_ascii=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
