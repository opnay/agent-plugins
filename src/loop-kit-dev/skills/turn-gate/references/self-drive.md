# Self-Drive

Self-drive is a prepared sequence overlay. It narrows when to continue autonomously, but it does not replace preparation, work, verification, reporting, or next-flow routing.

## Activation Requirements

Self-drive is active only when records contain:

- sequence objective
- prepared flow sequence
- active flow index
- current flow label
- progress note
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
- active flow index
- current flow label
- progress note
- planned flow count
- endpoint
- required next action
- acceptance signal
- blocker state

If identity or index state is missing, conflicting, or out of range, stop autonomous advancement and route to the user. Do not wrap the index or advance from memory.

`000-plan.md` stores only self-drive status and the sidecar pointer. `000-self-drive.md` owns sequence-level state. The active flow record owns flow-local snapshots.

## Interruption Handling

When a user message arrives during self-drive, interpret it inside the active sequence unless it explicitly stops the turn.

Priority order:

1. source-recorded explicit stop
2. destructive, external, commit, push, PR, publish, release, version bump, or approval-boundary-expanding request
3. scope, non-goal, endpoint, target, prepared order, or acceptance-signal change
4. blocker or repeated failure
5. status or progress question
6. ordinary note inside the recorded boundary

If the user changes endpoint, order, target, approval boundary, or acceptance signal, stop autonomous advancement and return to preparation.

If the user asks for status, report current flow, sequence position, verification state, and next required action, then continue only if the recorded sequence still permits it.

## Endpoint

Every self-drive flow needs verification before endpoint exhaustion handling. Non-pass verification routes through verification recovery first.

Before and after reporting, update the sidecar with active flow index, current flow label, progress note, next handoff, and blocker state. Keep the progress ledger as history; do not overwrite it with only the current summary.

Open-ended self-drive still needs a finite current cycle and an explicit repeat policy.

Sequence completion is not terminal closure. Report completion, update records, then reopen next-flow routing unless the user explicitly stops.
