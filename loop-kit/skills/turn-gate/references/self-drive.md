# Self-Drive

Self-drive lets Codex continue inside a recorded boundary without asking after every flow.
It does not replace `flow`, records, verification, approval checkpoints, or explicit-stop handling.

Use self-drive only when the user explicitly requests it or selects it as a next-flow mode.
Do not infer it from a long task list, passing checks, available subagents, or a plain "continue" request.

## Modes

- `finite`: run a prepared bounded flow sequence.
- `infinite`: continue until the user stops, using counted bounded iterations.

Infinite mode is not unlimited permission.
Each iteration still needs a bounded target, verification, approval check, and recorded continuation gate.

## Required State

Before autonomous continuation, `000-plan.md` must point to `000-self-drive.md`.

The sidecar must record:

- status and mode
- source-backed goal or objective
- current flow or loop identity
- active flow record identity
- next action
- endpoint or stop condition
- acceptance signal
- verification expectation
- allowed autonomous actions
- approval checkpoints
- blocker return conditions
- ledger

Finite mode also records sequence position.
Infinite mode also records `loop_count`, current loop label, and next bounded iteration.

Flow labels, next identity, handoff, readiness, ambiguity, and affected phase decisions come from `flow`.

## Gate

At every self-drive flow or loop start, read `000-plan.md`, `000-self-drive.md`, and the active flow record.

Continue only when all are true:

- plan says self-drive is active
- sidecar pointer matches
- mode and identity match the active flow
- endpoint or stop condition still applies
- verification expectation is known
- approval boundary still matches
- blocker state is clear
- next action is known

If any item is missing, stale, conflicting, or inaccessible, stop autonomous advancement and route to the user.
Do not advance from memory.

## Advance

Advance only after:

- current flow verification is `pass`
- `flow` handoff is not blocked
- next identity is known
- approval boundary still matches
- plan and sidecar gate pass again

Repair and recheck stay in the current flow until it passes.
Sequence completion or endpoint exhaustion stops autonomous advancement, reports status, updates records, and reopens next-flow routing unless the user explicitly stops.

## Updates

If the user asks for another bounded batch, inventory cycle, or changed endpoint, refresh the sidecar before work.
Preserve ledger history, but relock objective, target, endpoint, acceptance, verification, approval, and next action from current sources.

Scope, target, order, acceptance, endpoint, approval, blocker, or record conflicts stop autonomous advancement before continuation.
Return to `flow` for the affected decision.

Ordinary notes inside the recorded boundary can be logged without stopping self-drive.
Update `000-plan.md` only if turn-level routing changes.

## Interruption

During active self-drive, interpret new user input in this order:

1. source-recorded explicit stop
2. approval-sensitive action or approval boundary expansion
3. scope, non-goal, endpoint, target, order, or acceptance change
4. blocker or repeated failure
5. status or progress question
6. ordinary note inside the recorded boundary

For status questions, report current identity, finite index or loop count, verification state, and next required action.
Continue only if the sidecar still permits it.

## Approval And Stop

Self-drive cannot approve commit, push, PR, publish, release, version bump, destructive action, external effect, or scope expansion.
Those actions need exact target, expected effect, risk, recovery path, included/excluded scope, and explicit approval.

If writes or record creation are forbidden, report in-memory continuity only.
Do not use unrecorded state for later autonomous advance.

Terminal closure requires a source-recorded explicit stop stored in the active flow record.
