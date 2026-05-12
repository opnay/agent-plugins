# Ralph Loop Reference

Use this phase protocol when the current phase is one bounded issue that should be improved through a short fix-verify-reassess cycle.

## Core Contract

- Keep one loop focused on one primary issue.
- Prefer the smallest useful fix that can validate the current hypothesis.
- Verify immediately after the fix.
- Reassess whether another loop is justified before continuing.
- If the issue boundary expands enough to change success criteria, non-goals, verification, expected risky actions, or approval boundaries, stop the loop and return to preparation or question-routing.
- Do not execute destructive, irreversible, external, commit, push, PR, publish, release, version bump, or similar sensitive work unless the exact boundary was already approved in initial preparation or is approved through user-gated question-routing.

## Protocol Boundary

- Good fit: UI polish, bounded refactor stabilization, flaky issue reduction, or one narrow quality problem.
- Not a fit: broad delivery, multi-phase planning, review triage across many findings, missing scope boundaries, unresolved risky-action approval, or final commit execution.

## Handoff

- After one cycle, report verification results and residual risk or justify the next bounded cycle.
- If the loop expands, split it into a new flow or return to user-gated question-routing.

## Review Questions

- Did this loop stay on one primary issue?
- Was the fix small enough to test the hypothesis quickly?
- Is another loop actually justified?
- Did any expanded risk or approval boundary trigger question-routing instead of silent continuation?
