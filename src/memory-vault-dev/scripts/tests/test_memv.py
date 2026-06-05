from __future__ import annotations

import subprocess
import sys
import tempfile
import unittest
import os
from pathlib import Path


SCRIPT = Path(__file__).resolve().parents[1] / "memv.py"


class MemvTests(unittest.TestCase):
    def run_memv(self, target: Path, *args: str) -> subprocess.CompletedProcess[str]:
        return subprocess.run(
            [sys.executable, str(SCRIPT), str(target), *args],
            check=False,
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )

    def test_creates_root_memory_docs_and_category_indexes(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()

            result = self.run_memv(target, "--category", "Agents/Prompting")

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertTrue((target / "INDEX.md").exists())
            self.assertTrue((target / "preferences.md").exists())
            self.assertTrue((target / "environment.md").exists())
            self.assertTrue((target / "workflows.md").exists())
            self.assertTrue((target / "Agents" / "INDEX.md").exists())
            self.assertTrue((target / "Agents" / "Prompting" / "INDEX.md").exists())
            agents = (target / "AGENTS.md").read_text(encoding="utf-8")
            self.assertIn("[Agents/INDEX.md](Agents/INDEX.md)", agents)
            self.assertIn("[Agents/Prompting/INDEX.md](Agents/Prompting/INDEX.md)", agents)
            self.assertIn("preferences.md", agents)
            self.assertIn("일회성 진행 로그", agents)

    def test_creates_indexed_knowledge_docs_and_updates_indexes(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()

            result = self.run_memv(
                target,
                "--knowledge",
                "Programming/some-knowledge",
                "--knowledge",
                "Programming/React/hook-rules",
            )

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertTrue((target / "Programming" / "001-some-knowledge.md").exists())
            self.assertTrue((target / "Programming" / "React" / "001-hook-rules.md").exists())
            root_index = (target / "INDEX.md").read_text(encoding="utf-8")
            programming_index = (target / "Programming" / "INDEX.md").read_text(encoding="utf-8")
            react_index = (target / "Programming" / "React" / "INDEX.md").read_text(encoding="utf-8")
            self.assertIn("[Programming/INDEX.md](Programming/INDEX.md)", root_index)
            self.assertIn("[001-some-knowledge.md](001-some-knowledge.md)", programming_index)
            self.assertIn("[React/INDEX.md](React/INDEX.md)", programming_index)
            self.assertIn("[001-hook-rules.md](001-hook-rules.md)", react_index)

    def test_existing_knowledge_slug_is_preserved(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()
            self.assertEqual(self.run_memv(target, "--knowledge", "Programming/some-knowledge").returncode, 0)

            result = self.run_memv(target, "--knowledge", "Programming/some-knowledge", "--dry-run")

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertIn(f"preserve: {target / 'Programming' / '001-some-knowledge.md'}", result.stdout)

    def test_second_run_preserves_existing_files(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()
            self.assertEqual(self.run_memv(target, "--category", "Agents/Prompting").returncode, 0)

            result = self.run_memv(target, "--category", "Agents/Prompting", "--dry-run")

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertIn(f"preserve: {target / 'INDEX.md'}", result.stdout)
            self.assertIn(f"preserve: {target / 'AGENTS.md'}", result.stdout)

    def test_rejects_category_deeper_than_two_levels(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()

            result = self.run_memv(target, "--category", "A/B/C", "--dry-run")

            self.assertNotEqual(result.returncode, 0)
            self.assertIn("one or two path parts", result.stderr)

    def test_rejects_knowledge_deeper_than_two_category_levels(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()

            result = self.run_memv(target, "--knowledge", "A/B/C/D", "--dry-run")

            self.assertNotEqual(result.returncode, 0)
            self.assertIn("category/slug or category/subcategory/slug", result.stderr)

    def test_target_defaults_to_workspace_memory_vault(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            home = Path(tmp)
            target = home / "Workspace" / "Memory-vault"
            target.mkdir(parents=True)
            env = os.environ.copy()
            env["HOME"] = str(home)

            result = subprocess.run(
                [sys.executable, str(SCRIPT), "--dry-run"],
                check=False,
                text=True,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                env=env,
            )

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertIn(str(target / "INDEX.md"), result.stdout)


if __name__ == "__main__":
    unittest.main()
