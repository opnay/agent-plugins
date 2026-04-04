from __future__ import annotations

import subprocess
import sys
import tempfile
import unittest
from pathlib import Path

SCRIPTS_DIR = Path(__file__).resolve().parents[1] / "scripts"
if str(SCRIPTS_DIR) not in sys.path:
    sys.path.insert(0, str(SCRIPTS_DIR))

import orchestrator
import team_bus


class OrchestratorStabilityTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp_dir = tempfile.TemporaryDirectory()
        base = Path(self.temp_dir.name)
        self.bus_path = base / "chat.jsonl"
        self.dlq_path = base / "chat.dlq.jsonl"
        self.checkpoint_path = base / "chat.checkpoints.json"

    def tearDown(self) -> None:
        self.temp_dir.cleanup()

    def test_dispatch_route_request_emits_target_message(self) -> None:
        event = {
            "id": "evt-1",
            "from": "main",
            "type": "route.request",
            "payload": {
                "text": "@alpha summarize",
                "team": ["alpha", "beta"],
            },
        }

        generated = orchestrator._dispatch_route_request(self.bus_path, event)

        self.assertEqual(generated, 2)
        messages = team_bus.load_messages(self.bus_path)
        self.assertEqual(len(messages), 2)

        chat_events = [m for m in messages if m.get("type") == "chat.message"]
        self.assertEqual(len(chat_events), 1)
        self.assertEqual(chat_events[0].get("to"), "alpha")
        self.assertEqual(chat_events[0]["payload"].get("text"), "summarize")

        dispatch_events = [m for m in messages if m.get("type") == "orchestrator.dispatched"]
        self.assertEqual(len(dispatch_events), 1)
        self.assertEqual(dispatch_events[0]["payload"].get("target"), "alpha")

    def test_run_orchestrator_processes_route_and_writes_checkpoint(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="main",
            message_type="route.request",
            payload={"text": "@alpha ping", "team": ["alpha", "beta"]},
        )

        summary = orchestrator.run_orchestrator(
            bus_path=self.bus_path,
            dlq_path=self.dlq_path,
            checkpoint_path=self.checkpoint_path,
            duration_seconds=0.2,
            poll_seconds=0.01,
            watch_mode="poll",
            rotate_max_bytes=0,
            rotate_archive_dir=self.bus_path.parent / "archives",
            spawn_workers=False,
            workers=["alpha", "beta"],
            worker_duration=0.2,
            worker_poll=0.01,
            max_hops=6,
            pending_timeout=2.0,
            max_attempts=2,
            retry_base=0.1,
            verbose=False,
        )

        self.assertEqual(summary.get("status"), "completed")
        self.assertEqual(summary.get("wait_backend"), "poll")
        checkpoint = team_bus.get_consumer_checkpoint(self.checkpoint_path, "orchestrator:main")
        self.assertGreaterEqual(int(checkpoint.get("last_seq", 0)), 1)

        message_types = [m.get("type") for m in team_bus.load_messages(self.bus_path)]
        self.assertIn("orchestrator.started", message_types)
        self.assertIn("chat.message", message_types)
        self.assertIn("orchestrator.stopped", message_types)

    def test_build_bus_waiter_prefers_watchfiles_when_available(self) -> None:
        class DummyWaiter:
            def __init__(self, backend: str) -> None:
                self.backend = backend

            def wait(self, timeout_seconds: float) -> bool:
                return False

            def close(self) -> None:
                return None

        original_watchfiles = orchestrator._create_watchfiles_waiter
        try:
            orchestrator._create_watchfiles_waiter = (
                lambda _bus_path, poll_seconds: DummyWaiter("watchfiles")
            )

            waiter = orchestrator._build_bus_waiter(
                bus_path=self.bus_path,
                poll_seconds=0.01,
                watch_mode="auto",
            )
            self.assertEqual(waiter.backend, "watchfiles")
        finally:
            orchestrator._create_watchfiles_waiter = original_watchfiles

    def test_build_bus_waiter_falls_back_to_poll(self) -> None:
        original_watchfiles = orchestrator._create_watchfiles_waiter
        try:
            orchestrator._create_watchfiles_waiter = lambda _bus_path, poll_seconds: None

            waiter = orchestrator._build_bus_waiter(
                bus_path=self.bus_path,
                poll_seconds=0.01,
                watch_mode="watch",
            )
            self.assertEqual(waiter.backend, "poll")
        finally:
            orchestrator._create_watchfiles_waiter = original_watchfiles

    def test_terminate_collect_handles_repeated_timeout(self) -> None:
        class DummyProc:
            def __init__(self) -> None:
                self.terminated = False
                self.killed = False
                self.calls = 0

            def terminate(self) -> None:
                self.terminated = True

            def kill(self) -> None:
                self.killed = True

            def communicate(self, timeout: float) -> tuple[str, str]:
                self.calls += 1
                raise subprocess.TimeoutExpired(cmd="dummy", timeout=timeout)

        dummy = DummyProc()
        stdout_text, stderr_text = orchestrator._terminate_collect_output(dummy, timeout_seconds=0.01)
        self.assertEqual(stdout_text, "")
        self.assertIn("timed out", stderr_text)
        self.assertTrue(dummy.terminated)
        self.assertTrue(dummy.killed)
        self.assertEqual(dummy.calls, 2)

    def test_collect_new_events_uses_incremental_reader(self) -> None:
        team_bus.append_message(
            bus_path=self.bus_path,
            sender="main",
            recipient="main",
            message_type="route.request",
            payload={"text": "@alpha hi", "team": ["alpha", "beta"]},
        )

        original_load_messages = team_bus.load_messages
        original_load_since = team_bus.load_messages_since_offset
        called = {"value": False}

        def _raise_full(*args: object, **kwargs: object) -> list[dict[str, object]]:
            raise AssertionError("full scan should not be used")

        def _load_since(path: Path, offset: int) -> tuple[list[dict[str, object]], int]:
            called["value"] = True
            return original_load_since(path, offset)

        team_bus.load_messages = _raise_full
        team_bus.load_messages_since_offset = _load_since
        try:
            events, highest_seq, next_offset = orchestrator._collect_new_events(
                bus_path=self.bus_path,
                last_seq=0,
                last_offset=0,
            )
        finally:
            team_bus.load_messages = original_load_messages
            team_bus.load_messages_since_offset = original_load_since

        self.assertEqual(len(events), 1)
        self.assertGreaterEqual(highest_seq, 1)
        self.assertGreaterEqual(next_offset, 1)
        self.assertTrue(called["value"])

    def test_run_orchestrator_throttles_rotation_calls(self) -> None:
        original_rotate = orchestrator.team_bus.rotate_bus_log
        calls = {"count": 0}

        def _rotate(*args: object, **kwargs: object) -> dict[str, object]:
            calls["count"] += 1
            return {"rotated": True, "archive_file": "dummy"}

        orchestrator.team_bus.rotate_bus_log = _rotate
        try:
            orchestrator.run_orchestrator(
                bus_path=self.bus_path,
                dlq_path=self.dlq_path,
                checkpoint_path=self.checkpoint_path,
                duration_seconds=0.08,
                poll_seconds=0.01,
                watch_mode="poll",
                rotate_max_bytes=1,
                rotate_archive_dir=self.bus_path.parent / "archives",
                spawn_workers=False,
                workers=["alpha", "beta"],
                worker_duration=0.2,
                worker_poll=0.01,
                max_hops=6,
                pending_timeout=2.0,
                max_attempts=2,
                retry_base=0.1,
                verbose=False,
            )
        finally:
            orchestrator.team_bus.rotate_bus_log = original_rotate

        self.assertLessEqual(calls["count"], 2)


if __name__ == "__main__":
    unittest.main()
