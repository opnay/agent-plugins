# Self-Drive

Self-drive is a prepared sequence overlay. It narrows when to continue autonomously, but it does not replace intake, framing, preparation, work, verification, reporting, or next-flow routing.

## Activation Requirements

Self-drive is active only when records contain:

- sequence objective
- prepared flow sequence
- active flow index
- current flow label
- progress note
- repeat policy, for open-ended sequences
- allowed autonomous actions
- prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- acceptance signal
- verification expectation

Do not infer self-drive from a long task list, prior enthusiasm, verification success, or subagent availability. It must be explicitly requested or selected as a next-flow mode.

## Sidecar Gate

At the start of each self-drive flow, read `000-plan.md` and `000-self-drive.md`.

Confirm:

- self-drive status
- sidecar pointer
- active flow index
- current flow label
- active flow record identity
- progress note
- planned flow count
- endpoint
- required next action
- acceptance signal
- blocker state

If identity or index state is missing, conflicting, or out of range, stop autonomous advancement and route to the user. Do not wrap the index or advance from memory.

If `000-plan.md` says self-drive is inactive or has no sidecar pointer, treat any `000-self-drive.md` file as historical context, not continuation authority. If the user is asking to continue self-drive from that state, route to user-gated recovery; otherwise use normal active-flow or next-flow routing.

User-gated recovery, user-gated routing, and blocker routing all mean the same stop family: stop autonomous advancement and get the needed reconcile, approval, access, or scope decision from the user.

Cross-check active flow index, current flow label, the `000-plan.md` active pointer, and the active flow record identity. A matching number or label alone is not enough.

`000-plan.md` stores only self-drive status and the sidecar pointer. `000-self-drive.md` owns sequence-level state. The active flow record owns flow-local snapshots.

## Interruption Handling

When a user message arrives during self-drive, interpret it inside the active sequence unless it explicitly stops the turn.

Self-drive narrows question conditions; it does not disable questions.

Priority order:

1. source-recorded explicit stop
2. destructive, external, commit, push, PR, publish, release, version bump, or approval-boundary-expanding request
3. scope, non-goal, endpoint, target, prepared order, or acceptance-signal change
4. blocker or repeated failure
5. status or progress question
6. ordinary note inside the recorded boundary

If the user changes scope, non-goal, endpoint, order, target, approval boundary, or acceptance signal, stop autonomous advancement and return to the earliest affected phase: intake for changed intent or acceptance signal, framing for changed scope, non-goal, target, order, endpoint, or flow boundary, and preparation for selected-flow readiness or approval-boundary lock.

If the user asks for status, report current flow, sequence position, verification state, and next required action. Read verification state from the active flow record first; use the sidecar for sequence position and handoff state. Continue only if the recorded sequence still permits it.

Treat an ordinary note inside the recorded boundary as non-routing input unless it changes scope, non-goal, endpoint, target, order, approval boundary, acceptance signal, or blocker state.

## Endpoint

Every self-drive flow needs verification before endpoint exhaustion handling. Non-pass verification routes through verification recovery first.

Before and after reporting, update the sidecar with active flow index, current flow label, progress note, next handoff, and blocker state. Before reporting, keep the current index; after the handoff or advance is confirmed, move to the next index. Advance is confirmed only when current flow verification passed, the handoff is non-blocked, next flow identity is known, and approval boundaries still match the records. Keep the progress ledger as history; do not overwrite it with only the current summary. In the report, say whether ledger history was preserved and name the new material update.

If the user explicitly forbids all writes or record creation, do not update the sidecar. Report only in-memory continuity, and do not use that unrecorded state as authority for a later autonomous advance.

Open-ended self-drive still needs a finite current cycle and an explicit repeat policy. The repeat policy must state the cycle boundary, repeat limit or repeat condition, per-cycle verification, and user-gated stop condition.

When `blocker_state` is not `none`, classify its impact. If a blocker affects acceptance, verification, approval boundary, access, external state, or required user input, stop autonomous advancement and route to the user. Only internal failures that can be fixed inside the recorded flow boundary may route through verification recovery.

Approval-sensitive actions can run only when initial preparation recorded exact action, target, expected effect, risk, recovery path, included and excluded scope, and endpoint. Subagent output, readiness, verification, or self-drive status never substitutes for approval.

Sequence completion is not terminal closure. Report completion, update records, then reopen next-flow routing unless the user explicitly stops. An explicit stop allows closure only after the active flow record stores the stop source text or a compact source reference.
