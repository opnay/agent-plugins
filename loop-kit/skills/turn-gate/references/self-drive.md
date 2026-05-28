# Self-Drive

Self-drive lets Codex continue inside a recorded boundary without asking after every flow. It never replaces the normal turn-gate lifecycle, verification, records, or approval checkpoints.

## Modes

Use self-drive only when the user explicitly requests it or selects it as a next-flow mode.

- `finite`: run a prepared flow sequence.
- `infinite`: keep working until the user stops, but only through one bounded iteration at a time.

Do not infer self-drive from a long task list, successful verification, or available subagents.

## Required State

Before any autonomous continuation, `000-plan.md` must point to `000-self-drive.md`. The sidecar must contain:

- `status`
- `mode`
- source-backed goal or objective
- current flow or loop identity
- active flow record identity
- `next_action`
- endpoint or stop condition
- acceptance signal
- verification expectation
- allowed autonomous actions
- approval checkpoints
- blocker return conditions
- ledger

For `finite`, also keep `active_flow_index`, `current_flow_label`, and `planned_flow_count`.

For `infinite`, keep `loop_count` and a current loop label. Do not create a large todo list just to represent infinity.

Start a new infinite sidecar with `loop_count: 1`. If no bounded target is known, route to target selection or blocker handling before work.

## Sidecar Gate

At every self-drive flow or loop start, read `000-plan.md` and `000-self-drive.md`.

Confirm:

- plan status and sidecar pointer
- mode
- current identity
- active flow record identity
- endpoint or stop condition
- verification expectation
- approval checkpoint
- blocker state
- required next action

If identity, index, loop count, pointer, endpoint, approval boundary, or active record conflicts, stop autonomous advancement and route to the user. In finite mode, `active_flow_index > planned_flow_count` is stale/corrupt sidecar state and needs reconcile. If records cannot be read while deciding endpoint or continuation, report an access blocker. Do not advance from memory.

`000-plan.md` stores only self-drive status and the sidecar pointer. The sidecar owns sequence or loop state. The active flow record owns flow-local evidence.

## Finite Mode

Advance only after:

- current flow verification passed
- handoff is not blocked
- next flow identity is known
- approval boundary still matches the sidecar

Keep the current index before reporting. After confirmed advance, update the next index, current flow label, progress note, `next_action`, and ledger.

Sequence completion is not terminal closure. Report completion, update records, then reopen next-flow routing unless the user explicitly stops.

If the endpoint says to stop self-drive after sequence exhaustion, stop autonomous advancement and report the handoff; do not terminal-close the turn. If the endpoint says to create another inventory cycle, refresh the sidecar as a new bounded finite cycle before work: first flow identity, planned count, acceptance, verification, `next_action`, and ledger.

## Infinite Mode

Use infinite mode for requests like:

```text
내가 강제로 종료할 때까지 무한히 작업해줘.
```

Treat this as counted bounded iterations, not unlimited permission.

Each iteration:

1. Choose one bounded target inside the recorded scope.
2. Execute the smallest complete change or check.
3. Verify it.
4. Report and append the ledger.
5. Check approval boundary, blocker state, and stop condition.
6. If continuation is still valid, increment `loop_count` and refresh `next_action`.

Stop autonomous advancement for missing target, insufficient verification, repeated failure, approval need, access or external blocker, user input need, or no useful bounded work left.

Non-pass verification (`fail`, `blocked`, `insufficient`, `not-started`, or `requested`) comes before endpoint exhaustion, loop advance, and next-flow continuation. Repair, gather evidence, or route the blocker first.

## Interruption

During active self-drive, interpret new user input in this order:

1. source-recorded explicit stop
2. approval-sensitive action or expanded approval boundary
3. scope, non-goal, endpoint, target, order, or acceptance change
4. blocker or repeated failure
5. status or progress question
6. ordinary note inside the recorded boundary

Scope, endpoint, target, order, acceptance, or approval changes stop autonomous advancement and return to the earliest affected phase. Endpoint changes are relock/update events: update the sidecar endpoint, ledger, and affected flow record next action before continuing.

For status questions, report current identity, finite index or loop count, verification state, and next required action. Continue only if the recorded sidecar still permits it.

## Approval And Stop

Self-drive cannot approve commit, push, PR, publish, release, version bump, destructive action, external effect, or scope expansion. Those actions need exact target, expected effect, risk, recovery path, included/excluded scope, and explicit approval.

If writes or record creation are forbidden, report in-memory continuity only. Do not use unrecorded state for later autonomous advance.

Terminal closure requires a source-recorded explicit stop stored in the active flow record.
