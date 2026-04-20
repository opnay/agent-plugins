from __future__ import annotations

import os
import subprocess
import sys
import tempfile
import time
import unittest
from pathlib import Path
from unittest.mock import patch

SCRIPTS_DIR = Path(__file__).resolve().parents[1]
if str(SCRIPTS_DIR) not in sys.path:
    sys.path.insert(0, str(SCRIPTS_DIR))

import controller


class SessionManagerControllerTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp_dir = tempfile.TemporaryDirectory()
        self.base = Path(self.temp_dir.name)
        self.root = self.base / ".agents" / "sessions"
        self.controller_path = SCRIPTS_DIR / "controller.py"

    def tearDown(self) -> None:
        self.temp_dir.cleanup()

    def write_text(self, relative_path: str, text: str) -> Path:
        path = self.root / relative_path
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text, encoding="utf-8")
        return path

    def init_artifact_with_thread_id(
        self,
        *,
        thread_id: str,
        artifact: str,
        goal: str | None,
        scope: str | None,
        constraints: str | None,
        force: bool = False,
    ) -> tuple[str, Path]:
        with patch.dict(os.environ, {"CODEX_THREAD_ID": thread_id}, clear=False):
            return controller.init_artifact(
                root=self.root,
                artifact=artifact,
                goal=goal,
                scope=scope,
                constraints=constraints,
                force=force,
            )

    def ensure_artifact_with_thread_id(
        self,
        *,
        thread_id: str,
        artifact: str,
        goal: str | None,
        scope: str | None,
        constraints: str | None,
    ) -> tuple[str, Path, bool]:
        with patch.dict(os.environ, {"CODEX_THREAD_ID": thread_id}, clear=False):
            return controller.ensure_artifact(
                root=self.root,
                artifact=artifact,
                goal=goal,
                scope=scope,
                constraints=constraints,
            )

    def test_init_uses_codex_thread_id_when_uuid_missing(self) -> None:
        env = os.environ.copy()
        env["CODEX_THREAD_ID"] = "thread-from-env"
        result = subprocess.run(
            [
                sys.executable,
                str(self.controller_path),
                "init",
                "--root",
                str(self.root),
                "--goal",
                "Goal",
                "--scope",
                "Scope",
                "--constraints",
                "Constraints",
            ],
            check=True,
            capture_output=True,
            text=True,
            env=env,
        )

        self.assertIn("session_uuid=thread-from-env", result.stdout)
        record_path = self.root / "thread-from-env" / "session_record.md"
        self.assertTrue(record_path.exists())

    def test_init_requires_codex_thread_id_when_no_record_path_override_exists(self) -> None:
        env = os.environ.copy()
        env.pop("CODEX_THREAD_ID", None)
        result = subprocess.run(
            [
                sys.executable,
                str(self.controller_path),
                "init",
                "--root",
                str(self.root),
                "--goal",
                "Goal",
                "--scope",
                "Scope",
                "--constraints",
                "Constraints",
            ],
            capture_output=True,
            text=True,
            env=env,
        )

        self.assertEqual(result.returncode, 1)
        self.assertIn("CODEX_THREAD_ID", result.stderr)

    def test_init_rejects_uuid_cli_option(self) -> None:
        env = os.environ.copy()
        env["CODEX_THREAD_ID"] = "thread-from-env"
        result = subprocess.run(
            [
                sys.executable,
                str(self.controller_path),
                "init",
                "--root",
                str(self.root),
                "--uuid",
                "explicit-uuid",
            ],
            capture_output=True,
            text=True,
            env=env,
        )

        self.assertEqual(result.returncode, 2)
        self.assertIn("unrecognized arguments: --uuid explicit-uuid", result.stderr)

    def test_init_change_record_creates_markdown_sections(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="change-target",
            artifact="change-record",
            goal=None,
            scope=None,
            constraints=None,
            force=False,
        )

        self.assertEqual(record_path.name, "change_record.md")
        text = record_path.read_text(encoding="utf-8")
        self.assertIn("# Purpose", text)
        self.assertIn("# Review Gates", text)

    def test_init_retrospective_record_creates_markdown_sections(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="retrospective-target",
            artifact="retrospective-record",
            goal=None,
            scope=None,
            constraints=None,
            force=False,
        )

        self.assertEqual(record_path.name, "retrospective_record.md")
        text = record_path.read_text(encoding="utf-8")
        self.assertIn("# Goal Recap", text)
        self.assertIn("# Follow-Up Guardrails", text)

    def test_ensure_returns_existing_record_without_rewriting(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="ensure-target",
            artifact="session-record",
            goal="Goal",
            scope="Scope",
            constraints="Constraints",
            force=False,
        )
        initial_text = record_path.read_text(encoding="utf-8")
        time.sleep(0.01)

        _, ensured_path, created = self.ensure_artifact_with_thread_id(
            thread_id="ensure-target",
            artifact="session-record",
            goal="Changed",
            scope="Changed",
            constraints="Changed",
        )

        self.assertFalse(created)
        self.assertEqual(ensured_path, record_path)
        self.assertEqual(record_path.read_text(encoding="utf-8"), initial_text)

    def test_write_updates_current_session_record(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="write-target",
            artifact="session-record",
            goal="Old goal",
            scope="Old scope",
            constraints="Old constraints",
            force=False,
        )

        updated_path = controller.update_record_field(
            record_path=record_path,
            field="goal",
            value="New goal",
            artifact="session-record",
        )

        text = updated_path.read_text(encoding="utf-8")
        self.assertIn("3. Goal: New goal", text)
        self.assertIn("4. Scope: Old scope", text)

    def test_write_updates_change_record_from_value_file(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="change-write-target",
            artifact="change-record",
            goal=None,
            scope=None,
            constraints=None,
            force=False,
        )
        value_path = self.base / "purpose.txt"
        value_path.write_text("Summarize the redesign\nfor reviewers.", encoding="utf-8")

        updated_path = controller.update_record_field(
            record_path=record_path,
            field="purpose",
            value=value_path.read_text(encoding="utf-8"),
            artifact="change-record",
        )

        text = updated_path.read_text(encoding="utf-8")
        self.assertIn("Summarize the redesign", text)
        self.assertIn("for reviewers.", text)

    def test_change_record_allows_nested_markdown_headings_in_section_body(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="change-nested-heading",
            artifact="change-record",
            goal=None,
            scope=None,
            constraints=None,
            force=False,
        )

        controller.update_record_field(
            record_path=record_path,
            field="reviewer_notes",
            value="Line one\n# Review Gates\nLine two",
            artifact="change-record",
        )

        validation = controller.validate_record(record_path=record_path, artifact="change-record")
        self.assertTrue(validation.valid)
        self.assertEqual(
            controller.show_record(record_path=record_path, field="reviewer_notes", artifact="change-record"),
            "Line one\n# Review Gates\nLine two\n",
        )

    def test_write_updates_retrospective_record_from_value_file(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="retrospective-write-target",
            artifact="retrospective-record",
            goal=None,
            scope=None,
            constraints=None,
            force=False,
        )
        value_path = self.base / "lessons.txt"
        value_path.write_text("Prefer deterministic fixtures\nfor harness work.", encoding="utf-8")

        updated_path = controller.update_record_field(
            record_path=record_path,
            field="lessons",
            value=value_path.read_text(encoding="utf-8"),
            artifact="retrospective-record",
        )

        text = updated_path.read_text(encoding="utf-8")
        self.assertIn("Prefer deterministic fixtures", text)
        self.assertIn("for harness work.", text)

    def test_retrospective_record_allows_nested_markdown_headings_in_section_body(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="retrospective-nested-heading",
            artifact="retrospective-record",
            goal=None,
            scope=None,
            constraints=None,
            force=False,
        )

        controller.update_record_field(
            record_path=record_path,
            field="lessons",
            value="Line one\n# Outcome Summary\nLine two",
            artifact="retrospective-record",
        )

        validation = controller.validate_record(record_path=record_path, artifact="retrospective-record")
        self.assertTrue(validation.valid)
        self.assertEqual(
            controller.show_record(record_path=record_path, field="lessons", artifact="retrospective-record"),
            "Line one\n# Outcome Summary\nLine two\n",
        )

    def test_write_requires_resolvable_record_target(self) -> None:
        with patch.dict(os.environ, {}, clear=True):
            with self.assertRaises(SystemExit) as error:
                controller.resolve_record_path(
                    record_path=None,
                    root=self.root,
                    artifact="session-record",
                )
        self.assertIn("CODEX_THREAD_ID", str(error.exception))

    def test_show_returns_specific_field_value_from_current_record(self) -> None:
        _, record_path = self.init_artifact_with_thread_id(
            thread_id="show-target",
            artifact="session-record",
            goal="Visible goal",
            scope="Scope",
            constraints="Constraints",
            force=False,
        )

        output = controller.show_record(record_path=record_path, field="goal", artifact="session-record")
        self.assertEqual(output, "Visible goal\n")

    def test_validate_reports_malformed_session_record(self) -> None:
        record_path = self.write_text(
            "bad-record/session_record.md",
            "\n".join(
                [
                    "1. Session UUID: bad-record",
                    "2. Goal: Goal",
                    "3. Scope: Scope",
                    "",
                ]
            ),
        )

        report = controller.validate_record(record_path=record_path, artifact="session-record")
        self.assertFalse(report.valid)
        self.assertEqual(report.detected_schema, None)
        self.assertIn("numbered sections do not match", report.errors[0])

    def test_list_returns_sessions_sorted_by_recent_update_with_schema(self) -> None:
        _, first_path = self.init_artifact_with_thread_id(
            thread_id="first",
            artifact="session-record",
            goal="Goal",
            scope="Scope",
            constraints="Constraints",
            force=False,
        )
        time.sleep(0.01)
        _, second_path = self.init_artifact_with_thread_id(
            thread_id="second",
            artifact="session-record",
            goal="Goal",
            scope="Scope",
            constraints="Constraints",
            force=False,
        )

        sessions = controller.list_session_records(self.root, artifact="session-record")
        self.assertEqual(sessions[0][0], "second")
        self.assertEqual(Path(sessions[0][2]), second_path)
        self.assertEqual(Path(sessions[1][2]), first_path)

        detailed_sessions = controller.list_artifact_records(self.root, artifact="session-record")
        self.assertEqual(detailed_sessions[0][2], "session-record")
        self.assertEqual(Path(detailed_sessions[0][3]), second_path)

    def test_validate_cli_returns_nonzero_for_malformed_record(self) -> None:
        self.write_text(
            "bad-cli/session_record.md",
            "\n".join(
                [
                    "1. Session UUID: bad-cli",
                    "2. Goal: Goal",
                    "3. Scope: Scope",
                    "",
                ]
            ),
        )

        result = subprocess.run(
            [
                sys.executable,
                str(self.controller_path),
                "validate",
                "--artifact",
                "session-record",
                "--record-path",
                str(self.root / "bad-cli" / "session_record.md"),
            ],
            capture_output=True,
            text=True,
        )

        self.assertEqual(result.returncode, 1)
        self.assertIn("valid=false", result.stdout)
        self.assertIn("detected_schema=None", result.stdout)


if __name__ == "__main__":
    unittest.main()
