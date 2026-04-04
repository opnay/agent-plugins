from __future__ import annotations

import subprocess
import sys
import tempfile
import unittest
from datetime import datetime, timezone
from pathlib import Path

SCRIPTS_DIR = Path(__file__).resolve().parents[1] / "scripts"
if str(SCRIPTS_DIR) not in sys.path:
    sys.path.insert(0, str(SCRIPTS_DIR))

import team_bus


class TeamBusStabilityTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp_dir = tempfile.TemporaryDirectory()
        base = Path(self.temp_dir.name)
        self.bus_path = base / "chat.jsonl"
        self.dlq_path = base / "chat.dlq.jsonl"
        self.checkpoint_path = base / "chat.checkpoints.json"
        self.archive_dir = base / "chat.archives"

    def tearDown(self) -> None:
        self.temp_dir.cleanup()

    def test_unknown_target_routes_to_guard_and_dlq(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="alpha",
            message_type="chat.message",
            payload={
                "text": "@ghost investigate",
                "origin": "main",
                "hops": 0,
                "team": ["alpha", "beta"],
            },
        )

        handled = team_bus.process_once(
            worker_name="alpha",
            bus_path=self.bus_path,
            processed_ids=set(),
            pending={},
            dlq_path=self.dlq_path,
        )

        self.assertEqual(handled, 1)
        messages = team_bus.load_messages(self.bus_path)
        guard_events = [m for m in messages if m.get("type") == "guard.unknown_target"]
        self.assertEqual(len(guard_events), 1)
        self.assertEqual(guard_events[0].get("to"), "main")

        dlq_entries = team_bus.load_dlq_entries(self.dlq_path)
        self.assertEqual(len(dlq_entries), 1)
        self.assertEqual(dlq_entries[0].get("reason"), "chat.unknown_target")

    def test_retry_exhausted_generates_timeout_and_dlq(self) -> None:
        pending = {
            "corr-1": {
                "requester": "main",
                "a": 4,
                "b": 5,
                "started_at_monotonic": 0.0,
                "attempts": 2,
                "retry_due_at_monotonic": 0.0,
                "max_attempts": 2,
            }
        }

        generated = team_bus._maintain_pending_for_alpha(
            bus_path=self.bus_path,
            pending=pending,
            now_monotonic=10.0,
            pending_timeout_seconds=0.0,
            max_attempts=2,
            retry_base_seconds=0.01,
            dlq_path=self.dlq_path,
        )

        self.assertEqual(generated, 1)
        self.assertEqual(pending, {})

        messages = team_bus.load_messages(self.bus_path)
        timeout_events = [m for m in messages if m.get("type") == "result.timeout"]
        self.assertEqual(len(timeout_events), 1)
        self.assertEqual(timeout_events[0]["payload"]["reason"], "beta.retry_exhausted")

        dlq_entries = team_bus.load_dlq_entries(self.dlq_path)
        self.assertEqual(len(dlq_entries), 1)
        self.assertEqual(dlq_entries[0].get("reason"), "beta.retry_exhausted")

    def test_rotate_bus_log_compacts_and_keeps_summary(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="alpha",
            message_type="chat.message",
            payload={"text": "hello"},
        )

        result = team_bus.rotate_bus_log(
            bus_path=self.bus_path,
            max_bytes=1,
            archive_dir=self.archive_dir,
            compress=True,
        )

        self.assertTrue(result.get("rotated"))
        archive_file = Path(str(result.get("archive_file", "")))
        self.assertTrue(archive_file.exists())

        messages = team_bus.load_messages(self.bus_path)
        self.assertEqual(len(messages), 1)
        self.assertEqual(messages[0].get("type"), "system.log_compacted")

    def test_process_once_updates_checkpoint(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="alpha",
            message_type="chat.message",
            payload={"text": "status", "origin": "main", "hops": 0, "team": ["alpha", "beta"]},
        )

        handled = team_bus.process_once(
            worker_name="alpha",
            bus_path=self.bus_path,
            processed_ids=set(),
            pending={},
            checkpoint_path=self.checkpoint_path,
            checkpoint_consumer="worker:alpha",
            dlq_path=self.dlq_path,
        )

        self.assertEqual(handled, 1)
        checkpoint = team_bus.get_consumer_checkpoint(self.checkpoint_path, "worker:alpha")
        self.assertGreaterEqual(int(checkpoint.get("last_seq", 0)), 1)

    def test_process_once_handles_worker_exception_with_guard_and_dlq(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="alpha",
            message_type="task.add",
            payload={"a": "bad", "b": 2},
        )

        original_handle_message = team_bus._handle_message

        def _raise(*args: object, **kwargs: object) -> None:
            raise ValueError("boom")

        team_bus._handle_message = _raise
        processed_ids: set[str] = set()
        try:
            handled = team_bus.process_once(
                worker_name="alpha",
                bus_path=self.bus_path,
                processed_ids=processed_ids,
                pending={},
                dlq_path=self.dlq_path,
            )
        finally:
            team_bus._handle_message = original_handle_message

        self.assertEqual(handled, 0)
        messages = team_bus.load_messages(self.bus_path)
        processing_errors = [m for m in messages if m.get("type") == "guard.processing_error"]
        self.assertEqual(len(processing_errors), 1)
        self.assertIn("boom", str(processing_errors[0].get("payload", {}).get("error", "")))

        dlq_entries = team_bus.load_dlq_entries(self.dlq_path)
        self.assertEqual(len(dlq_entries), 1)
        self.assertEqual(dlq_entries[0].get("reason"), "worker.processing_error")
        self.assertEqual(len(processed_ids), 1)

    def test_rotate_bus_log_uses_unique_archive_name_even_same_timestamp(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="alpha",
            message_type="chat.message",
            payload={"text": "hello"},
        )

        original_datetime = team_bus.datetime

        class _FixedDatetime:
            @staticmethod
            def now(_tz: timezone) -> datetime:
                return datetime(2026, 2, 24, 12, 0, 0, tzinfo=timezone.utc)

        team_bus.datetime = _FixedDatetime
        try:
            first = team_bus.rotate_bus_log(
                bus_path=self.bus_path,
                max_bytes=1,
                archive_dir=self.archive_dir,
                compress=False,
            )
            team_bus.append_message(
                bus_path=self.bus_path,
                sender="main",
                recipient="alpha",
                message_type="chat.message",
                payload={"text": "after-rotate"},
            )
            second = team_bus.rotate_bus_log(
                bus_path=self.bus_path,
                max_bytes=1,
                archive_dir=self.archive_dir,
                compress=False,
            )
        finally:
            team_bus.datetime = original_datetime

        self.assertTrue(first.get("rotated"))
        self.assertTrue(second.get("rotated"))
        self.assertNotEqual(first.get("archive_file"), second.get("archive_file"))

    def test_cli_send_reports_invalid_payload_without_traceback(self) -> None:
        cmd = [
            sys.executable,
            str(SCRIPTS_DIR / "team_bus.py"),
            "send",
            "--bus",
            str(self.bus_path),
            "--from",
            "main",
            "--to",
            "alpha",
            "--type",
            "chat.message",
            "--payload",
            "not-json",
        ]
        completed = subprocess.run(
            cmd,
            capture_output=True,
            text=True,
            check=False,
        )

        self.assertEqual(completed.returncode, 1)
        self.assertIn("payload must be a JSON object", completed.stderr)
        self.assertNotIn("Traceback", completed.stderr)

    def test_load_records_uses_shared_lock_and_tail_recovery(self) -> None:
        first = {
            "id": "msg-1",
            "seq": 1,
            "ts": "2023-02-24T00:00:00+00:00",
            "from": "main",
            "to": "alpha",
            "type": "chat.message",
            "payload": {"text": "one"},
        }
        self.bus_path.write_text(
            f"{team_bus.json.dumps(first, ensure_ascii=True)}\n"
            "{\"id\":\"msg-2\",\"seq\":2",
            encoding="utf-8",
        )

        original_flock = team_bus.fcntl.flock
        lock_ops: list[int] = []

        def _spy(fd: int, op: int) -> None:
            lock_ops.append(op)
            original_flock(fd, op)

        team_bus.fcntl.flock = _spy
        try:
            records, offset = team_bus.load_messages_since_offset(self.bus_path, 0)
        finally:
            team_bus.fcntl.flock = original_flock

        self.assertEqual([m["id"] for m in records], ["msg-1"])
        self.assertIn(team_bus.fcntl.LOCK_SH, lock_ops)

        with self.bus_path.open("a", encoding="utf-8") as handle:
            handle.write(",\"ts\":\"2023-02-24T00:00:01+00:00\",\"from\":\"main\",\"to\":\"alpha\",\"type\":\"chat.message\",\"payload\":{\"text\":\"two\"}}\n")

        records_after, _ = team_bus.load_messages_since_offset(self.bus_path, offset)
        self.assertEqual([m["id"] for m in records_after], ["msg-2"])

    def test_update_consumer_checkpoint_does_not_use_split_read_write_helpers(self) -> None:
        original_load = team_bus._load_checkpoint_state
        original_save = team_bus._save_checkpoint_state

        def _raise_load(*args: object, **kwargs: object) -> dict[str, object]:
            raise AssertionError("split read helper must not be used")

        def _raise_save(*args: object, **kwargs: object) -> None:
            raise AssertionError("split write helper must not be used")

        team_bus._load_checkpoint_state = _raise_load
        team_bus._save_checkpoint_state = _raise_save
        try:
            team_bus.update_consumer_checkpoint(
                checkpoint_path=self.checkpoint_path,
                consumer="worker:alpha",
                last_seq=3,
                last_event_id="evt-3",
                last_offset=128,
            )
        finally:
            team_bus._load_checkpoint_state = original_load
            team_bus._save_checkpoint_state = original_save

        checkpoint = team_bus.get_consumer_checkpoint(self.checkpoint_path, "worker:alpha")
        self.assertEqual(int(checkpoint.get("last_seq", 0)), 3)
        self.assertEqual(int(checkpoint.get("last_offset", 0)), 128)

    def test_process_once_reads_incrementally_with_checkpoint_offset(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="alpha",
            message_type="chat.message",
            payload={"text": "hello"},
        )

        original_load_messages = team_bus.load_messages
        original_load_since = team_bus.load_messages_since_offset
        load_since_called = {"value": False}

        def _raise_full(*args: object, **kwargs: object) -> list[dict[str, object]]:
            raise AssertionError("full scan should not be used when checkpoint is enabled")

        def _load_since(path: Path, offset: int) -> tuple[list[dict[str, object]], int]:
            load_since_called["value"] = True
            return original_load_since(path, offset)

        team_bus.load_messages = _raise_full
        team_bus.load_messages_since_offset = _load_since
        try:
            handled = team_bus.process_once(
                worker_name="alpha",
                bus_path=self.bus_path,
                processed_ids=set(),
                pending={},
                checkpoint_path=self.checkpoint_path,
                checkpoint_consumer="worker:alpha",
                dlq_path=self.dlq_path,
            )
        finally:
            team_bus.load_messages = original_load_messages
            team_bus.load_messages_since_offset = original_load_since

        self.assertEqual(handled, 1)
        self.assertTrue(load_since_called["value"])


if __name__ == "__main__":
    unittest.main()
