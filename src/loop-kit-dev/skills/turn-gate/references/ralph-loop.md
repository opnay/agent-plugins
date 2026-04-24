# Ralph Loop Reference

Use this mode when the current phase is one bounded issue that should be improved through a short fix-verify-reassess cycle.

## Core Contract

- Keep one loop focused on one primary issue.
- Prefer the smallest useful fix that can validate the current hypothesis.
- Verify immediately after the fix.
- Reassess whether another loop is justified before continuing.

## Mode Boundary

- Good fit: UI polish, bounded refactor stabilization, flaky issue reduction, one narrow quality problem.
- Not a fit: broad delivery, multi-phase planning, or review triage across many findings.

## Review Questions

- Did this loop stay on one primary issue?
- Was the fix small enough to test the hypothesis quickly?
- Is another loop actually justified?
