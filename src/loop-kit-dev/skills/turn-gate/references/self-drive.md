# Self-Drive Overlay

Self-drive is an overlay for an already prepared flow sequence. It is not a separate installed skill entrypoint and not a default runtime mode.

Use this reference only when the user explicitly asks for self-drive or when current session records show an active self-drive sequence.

## Preconditions

Self-drive can run only after records include:

- sequence objective
- finite prepared flow sequence
- active flow index as a 0-based machine field
- current flow label with human-readable number, name, file, or slug
- allowed autonomous actions
- prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- verification expectation
- progress note

If any of these are missing or contradictory, return to user-gated preparation or next-flow routing.

## Record Ownership

`000-plan.md` owns only self-drive active status and the `000-self-drive.md` pointer.

`000-self-drive.md` owns sequence-level state: objective, prepared sequence, active index, current label, autonomous boundaries, checkpoints, endpoint, blocker return conditions, progress note, and progress ledger.

Each active flow record owns only flow-local self-drive snapshots: sequence position, local progress note, next handoff, and blocker return condition. Do not repeat the full prepared sequence in every flow record.

Use `templates/self-drive-template.md` when creating `000-self-drive.md`.

## Continuation Rule

Each flow still runs through `preparation -> work -> verification -> reporting -> next-flow`.

Self-drive does not remove `next-flow`. It changes the result of `next-flow`: when the prepared sequence is still valid and the next flow is identifiable, continuation can be record-driven instead of a user question.

Stop autonomous continuation and return to user-gated routing when continuation identity, scope, endpoint, approval boundary, blocker state, or current-flow identity is unclear.

Finite sequence exhaustion does not create new work automatically. Follow the recorded endpoint: stop/handoff, commit-readiness reporting handoff, bounded repeat policy, blocker decision, or next-flow reopening.

## Incoming User Messages During Self-Drive

If a user message arrives while self-drive is active, treat it as input inside the active self-drive sequence unless it explicitly stops the turn.

Priority:

1. Source-recorded explicit stop: record closure and end only after reporting permits it.
2. Destructive, external, commit, push, PR, publish, release, version-bump, or other approval-sensitive action outside recorded approval boundary: stop and ask.
3. Scope, non-goal, endpoint, target, flow order, or acceptance signal change: return to preparation or next-flow routing and re-lock the sequence.
4. Blocker or repeated critical failure: route to earliest safe phase or user-gated blocker decision.
5. Status/progress question only: report current phase, active flow, verification state, and next action, then continue if no higher rule applies.
6. Ordinary continuation note inside recorded boundary: record material context and continue.

This priority list is self-drive interruption handling, not a general user-message taxonomy.

## Approval Boundary

Self-drive may execute approval-sensitive actions only when the initial prepared sequence recorded exact action, target, expected effect, risk, recovery path, include/exclude scope, and endpoint.

Commit, push, PR, publish, release, and version bump are approval-sensitive execution steps. If they are not explicitly included with clear endpoint and risk, return to user-gated routing.

Subagents cannot grant approval, expand scope, change endpoint, or authorize external/destructive work.
