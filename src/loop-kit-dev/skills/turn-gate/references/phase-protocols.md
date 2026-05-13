# Phase Protocols

`turn-gate` normally runs in an implicit default operating state. The default state always preserves:

1. preparation
2. work
3. verification
4. reporting
5. next-flow

Phase protocols are not modes. They are local contracts for how to perform the current phase.

## Selection

Choose the earliest blocker:

1. `deep-interview`
2. `review-loop`
3. `ralph-loop`
4. `autopilot`
5. `commit-readiness-gate`

If none applies, stay in the default operating state without a protocol suffix.

## Protocols

### deep-interview

Use in preparation when requirement discovery or scope lock blocks work.

Apply when intent, included scope, non-goals, acceptance signal, verification expectation, planned flow list, or approval boundary is insufficient. Ask bounded questions with `request_user_input` when possible. Do not use this for already-locked implementation work or for a single narrow review finding.

Handoff back to routing after the scope and approval boundary are locked. If they remain unclear, keep active question routing open.

### review-loop

Use when review, QA, or self-review findings are the current blocker.

Focus one loop on one bounded blocking finding. Fix and verify that finding inside the active flow boundary. Keep low-value notes as follow-up candidates. If a finding creates broader scope or a new approval boundary, return to preparation or question routing.

### ralph-loop

Use for one narrow fix-verify-reassess cycle.

Keep one primary issue, make the smallest useful change that tests the current hypothesis, verify immediately, then reassess whether another cycle is justified. If the loop grows enough to change success criteria, non-goals, verification, or approval boundary, return to preparation or question routing.

### autopilot

Use for locked-scope end-to-end execution.

Proceed autonomously only inside recorded scope, non-goals, verification expectation, and approval boundary. Run meaningful verification after changes. Treat QA issues as bounded loops. Do not treat autonomous execution as approval for destructive, external, commit, push, PR, publish, release, or version-bump actions.

### commit-readiness-gate

Use when judging whether the intended change unit is ready to commit.

Evaluate only the intended change unit. Report readiness, residual risk, intended diff scope, excluded unrelated changes, verification evidence, likely commit-message scope, and minimum review recommendation. Keep readiness judgment separate from staging, commit, push, PR, publish, release, and version-bump authority.

## External Actions

Commit execution, push, PR, publish, release, version bump, and other external or irreversible actions require an explicit approval-sensitive checkpoint. Record exact target, expected effect, risk, recovery path, included/excluded scope, and end point before execution.
