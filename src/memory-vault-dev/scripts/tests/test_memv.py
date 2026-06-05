from __future__ import annotations

import subprocess
import sys
import tempfile
import unittest
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

    def test_creates_root_and_category_indexes(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()

            result = self.run_memv(target, "--category", "Programming/React")

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertTrue((target / "INDEX.md").exists())
            self.assertTrue((target / "Programming" / "INDEX.md").exists())
            self.assertTrue((target / "Programming" / "React" / "INDEX.md").exists())
            agents = (target / "AGENTS.md").read_text(encoding="utf-8")
            self.assertIn("`Programming/INDEX.md`", agents)
            self.assertIn("`Programming/React/INDEX.md`", agents)

    def test_second_run_preserves_existing_files(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()
            self.assertEqual(self.run_memv(target, "--category", "Programming/React").returncode, 0)

            result = self.run_memv(target, "--category", "Programming/React", "--dry-run")

            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertIn(f"preserve: {target / 'INDEX.md'}", result.stdout)
            self.assertIn(f"preserve: {target / 'AGENTS.md'}", result.stdout)

    def test_rejects_category_deeper_than_two_levels(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            target = Path(tmp).resolve()

            result = self.run_memv(target, "--category", "A/B/C", "--dry-run")

            self.assertNotEqual(result.returncode, 0)
            self.assertIn("one or two path parts", result.stderr)


if __name__ == "__main__":
    unittest.main()
