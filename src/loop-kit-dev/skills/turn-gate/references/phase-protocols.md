# Phase Protocols

`turn-gate` normally runs in the implicit default operating state. The names below are phase protocols, not modes. Use them to decide how to perform the current phase.

## Routing

Choose the earliest blocker:

1. `deep-interview`
2. `review-loop`
3. `ralph-loop`
4. `autopilot`
5. `commit-readiness-gate`

External actions, destructive actions, commit, push, PR, publish, release, and version bump are approval-sensitive execution or handoff steps. Do not treat protocol selection as approval.

## Deep Interview

Use when requirement discovery or scope lock blocks work:

- intent, scope, non-goals, acceptance signal, verification expectation, or approval boundary is missing;
- the same request can produce several valid artifacts;
- the answer can change output shape or verification path;
- a planned flow list cannot yet be built safely.

Prefer bounded `request_user_input` choices. If you proceed by inference, record the work boundary and non-goals. When enough is locked, return to routing.

Do not use this protocol for simple operation/target ambiguity; use meaning resolution inside preparation instead.

## Review Loop

Use when review feedback, QA findings, or self-review findings are the current blocker.

Handle one bounded blocking finding at a time. Prioritize correctness, regression, reliability, and delivery risk. Low-value notes become follow-up candidates unless they fit the current flow boundary.

If a finding creates new broad scope or a new approval boundary, return to preparation or question-routing.

## Ralph Loop

Use when one narrow issue can be improved by a small fix-verify-reassess cycle.

Keep each cycle focused on one primary issue. Make the smallest useful change that tests the current hypothesis, verify immediately, then reassess whether another cycle is justified.

Return to preparation or question-routing if success criteria, non-goals, verification expectation, expected risky actions, or approval boundary change.

## Autopilot

Use when the scope, non-goals, verification expectation, and approval boundary are locked and the task needs end-to-end execution.

Continue within the recorded boundary until blocked. Treat new approval-sensitive actions as blockers unless they were explicitly included in the prepared boundary. Handle QA issues with bounded loops, and report repeated critical failure as a root blocker.

Autopilot does not imply approval for destructive, external, commit, push, PR, publish, release, or version bump actions.

## Commit Readiness Gate

Use when the intended change unit is mostly complete and needs readiness judgment.

Assess only the intended change unit, not the whole repository. Report:

- intended diff scope;
- unrelated changes to exclude;
- verification status and evidence;
- residual risk;
- likely commit-message scope;
- minimum review recommendation.

Readiness is evidence, not execution authority. Staging, committing, pushing, opening PRs, publishing, releasing, and version bumping require explicit approval or a separate handoff.
