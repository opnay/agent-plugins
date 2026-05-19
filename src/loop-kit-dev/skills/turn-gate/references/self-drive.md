# Self-Drive Overlay

Self-drive is a prepared sequence overlay for `turn-gate`. It is not a separate skill and not a selectable runtime mode. Apply it only when the user explicitly asks for autonomous continuation and the sequence is already finite, scoped, and recorded.

## Required Starting State

Before self-drive can continue, records must show:

- sequence objective
- prepared flow sequence
- active flow index as a 0-based machine field
- current flow label with human-readable number, name, file, or slug
- allowed autonomous actions
- prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- progress note

`000-plan.md` stores only self-drive active status and the `000-self-drive.md` pointer. `000-self-drive.md` owns sequence-level state. Each `001+` flow record owns only flow-local progress, next handoff, and blocker return condition.

## Continuation Rule

Each flow still runs through `preparation -> work -> verification -> reporting -> next-flow`.

Self-drive does not remove `next-flow`. If the prepared sequence is still valid and the next flow is identifiable, `next-flow` may continue by record-based handoff instead of asking the user.

Do not continue autonomously when continuation identity, scope, endpoint, approval boundary, blocker state, or current-flow identity is unclear.

## User Messages During Self-Drive

If a user message arrives while self-drive is active, interpret it inside the active sequence first. The user does not need to repeat `self-drive`.

Priority:

1. explicit turn stop: record closure source and stop after reporting
2. approval-sensitive action outside the recorded boundary: return to user-gated approval routing
3. scope, non-goal, endpoint, target, order, or acceptance change: return to preparation or next-flow routing
4. blocker or repeated critical failure: return to earliest safe phase or user-gated blocker decision
5. status-only question: report current phase, active flow, verification state, and next action, then continue if safe
6. ordinary continuation note inside recorded boundary: record the material note and continue

This priority is only for self-drive interruption handling. It does not replace meaning resolution, approval boundary, explicit stop, or question routing.

## Question Tool

Self-drive narrows questions; it does not turn them off.

Use user-gated routing for:

- approval-sensitive execution
- scope, target, endpoint, order, non-goal, or acceptance changes
- blocker state
- record access failure
- repeated critical failure
- current-flow identity ambiguity
- unclear sequence exhaustion behavior

Status-only input and clear prepared transitions should not be over-interrupted with questions.

## Approval Boundary

Self-drive may execute approval-sensitive actions only when initial preparation recorded exact action, target, expected effect, risk, recovery path, included/excluded scope, and endpoint.

Commit, push, PR, publish, release, and version bump remain approval-sensitive execution steps. If not explicitly covered, return to user-gated routing.

## Endpoint

The endpoint must be finite. `forever` or `until stopped` alone is not enough.

When the finite sequence is exhausted, do not invent new work. Continue to the recorded endpoint, commit-readiness handoff, blocker decision, or next-flow reopening.

If repeat cycles are allowed, the endpoint must say how to start the next bounded cycle and how to refresh count, active index, current label, and progress note.
