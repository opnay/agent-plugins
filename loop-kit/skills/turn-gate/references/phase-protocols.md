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

Before selecting a protocol, resolve operation/target ambiguity and approval-sensitive boundaries that would change the work surface, verification path, or execution authority.

## deep-interview

Use in preparation when requirement discovery or scope lock blocks work.

Apply when intent, included scope, non-goals, acceptance signal, verification expectation, planned flow list, or approval boundary is insufficient. Ask bounded questions with the question tool when possible.

Do not use this for already-locked implementation work, a single narrow review finding, or final commit readiness judgment.

Handoff back to phase routing after the scope and approval boundary are locked. If they remain unclear, keep active question routing open and do not enter work.

## review-loop

Use when review, QA, or self-review findings are the current blocker.

Focus one loop on one bounded blocking finding. Fix and verify that finding inside the active flow boundary. Keep low-value notes as follow-up candidates. If a finding creates broader scope or a new approval boundary, return to preparation or question routing.

Review feedback is not approval for destructive, external, commit, push, PR, publish, release, or version-bump actions.

## ralph-loop

Use for one narrow fix-verify-reassess cycle.

Keep one primary issue, make the smallest useful change that tests the current hypothesis, verify immediately, then reassess whether another cycle is justified. If the loop grows enough to change success criteria, non-goals, verification, or approval boundary, return to preparation or question routing.

Do not execute destructive, irreversible, external, commit, push, PR, publish, release, or version-bump actions unless the exact boundary is already approved.

## autopilot

Use for locked-scope end-to-end execution.

Proceed autonomously only inside recorded scope, non-goals, verification expectation, and approval boundary. Run meaningful verification after changes. Treat QA issues as bounded loops. Do not treat autonomous execution as approval for destructive, external, commit, push, PR, publish, release, or version-bump actions.

If scope floor is not met, return to `deep-interview` or question routing. If an approval boundary appears, stop autonomous work and route to a user-gated checkpoint.

## commit-readiness-gate

Use when judging whether the intended change unit is ready to commit.

Evaluate only the intended change unit. Report readiness, residual risk, intended diff scope, excluded unrelated changes, verification evidence, likely commit-message scope, and minimum review recommendation.

Readiness is not execution authority. Keep staging, commit execution, push, PR, publish, release, and version bump behind a separate approval-sensitive boundary.

## External And Approval-Sensitive Actions

Commit execution, push, PR, publish, release, version bump, destructive changes, irreversible changes, and other external actions require an explicit approval-sensitive checkpoint. Record exact action, target, expected effect, risk, recovery path, included scope, excluded scope, and end point before execution.

Meaning resolution, phase protocol choice, readiness reporting, self-drive, or next-flow selection must not be interpreted as approval to execute these actions.
