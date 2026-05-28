# Self-Drive

Self-drive lets Codex continue inside a recorded boundary without asking after every flow.
It does not replace turn-gate lifecycle, records, verification, approval checkpoints, or explicit-stop handling.

## Modes

Use self-drive only when the user explicitly requests it or selects it as a next-flow mode.

- `finite`: run a prepared flow sequence.
- `infinite`: continue until the user stops, but only through counted bounded iterations.

Do not infer self-drive from a long task list, passing verification, available subagents, or a plain "continue" request.

## Required State

Before autonomous continuation, `000-plan.md` must point to `000-self-drive.md`.

The sidecar must contain:

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

For `finite`, also keep:

- `active_flow_index`
- `current_flow_label`
- `planned_flow_count`
- prepared sequence

Flow labels, next flow identity, and handoff conditions come from `flow` output.

For `infinite`, keep:

- `loop_count`
- `current_loop_label`
- next bounded iteration

Start a new infinite sidecar with `loop_count: 1`.
If no bounded target is known, route to target selection or blocker handling before work.

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

If identity, index, loop count, pointer, endpoint, approval boundary, or active record conflicts, stop autonomous advancement and route to the user.
If `000-plan.md` says self-drive is inactive, any sidecar is historical.
If records cannot be read while deciding endpoint or continuation, report an access blocker.
Do not advance from memory.

`000-plan.md` stores only self-drive status and sidecar pointer.
The sidecar owns sequence or loop state.
The active flow record owns flow-local evidence.

## Finite Mode

Advance only after:

- current flow verification passed
- `flow` handoff is not blocked
- `flow` output identifies the next flow
- approval boundary still matches the sidecar
- plan and sidecar gate pass again

Keep the current index before reporting.
Even when the other advance conditions look satisfied, recheck the plan and sidecar gate immediately before advancing.
After confirmed advance, update the next index, current flow label, progress note, `next_action`, and ledger.
Repair or recheck work does not count as a sequence advance; it only makes the current flow eligible to pass.

Sequence completion is not terminal closure.
Report completion, update records, then reopen next-flow routing unless the user explicitly stops.

If the endpoint says to stop self-drive after sequence exhaustion, stop only autonomous advancement.

If the user asks for another bounded batch or inventory cycle after finite completion, do not continue from the old endpoint.
Refresh the sidecar as a new bounded finite cycle before work:

- source-backed batch objective
- first flow identity
- planned flow count
- acceptance signal
- verification expectation
- approval checkpoints
- endpoint or stop condition
- `next_action`
- ledger continuation note

Order, endpoint, scope, target, acceptance, or approval changes stop autonomous advancement.
Use `flow` readiness or ambiguity to choose the earliest affected phase.
Update sidecar and affected flow record, then continue only if the gate passes again.

If an external blocker or access blocker recovers, do not continue from memory.
Reread `000-plan.md`, `000-self-drive.md`, and the active flow record.
Relock endpoint, approval boundary, next identity, and verification expectation before advancing.

Ordinary notes inside the recorded boundary do not stop self-drive.
Record them in the active flow record `Execution Log` or sidecar ledger.
Update `000-plan.md` only if turn-level routing changes.

## Infinite Mode

Use infinite mode for requests like:

```text
내가 강제로 종료할 때까지 무한히 작업해줘.
```

Treat this as counted bounded iterations, not unlimited permission.

Each iteration:

1. Choose one bounded target inside the recorded scope.
2. Execute the smallest complete change or check.
3. Verify it against the `flow` verification expectation.
4. Report and append the ledger.
5. Check approval boundary, blocker state, and stop condition.
6. If continuation is still valid, increment `loop_count` and refresh `next_action`.

Increment `loop_count` only after continuation remains valid.
If gate or record state conflicts before the next loop, keep the previous count and route recovery.
If the endpoint says the current batch is complete, report endpoint completion instead of incrementing a next-loop count.

Stop autonomous advancement for:

- missing target
- non-pass or incomplete verification
- two consecutive non-pass results for the same bounded target and cause
- approval need
- access or external blocker
- user input need
- no useful bounded work left
- sidecar or active record conflict
- endpoint, scope, target, order, acceptance, or approval boundary change

Non-pass verification (`fail`, `blocked`, `insufficient`, `not-started`, or `requested`) comes before endpoint exhaustion, loop advance, and next-flow continuation.
Even if work evidence looks successful, do not advance while the active flow record or sidecar metadata still has `verification_status: requested`, `not-started`, `fail`, `blocked`, or `insufficient`; update it to pass first.
If work evidence and verification metadata conflict, route it as a verification mismatch.
Do not continue or advance until the mismatch is resolved.
For `insufficient`, gather evidence and recheck the same bounded target.
Continue only after it becomes pass.

If an external blocker recovers, do not continue from memory.
Reread records, recheck sidecar gate, relock endpoint and approval boundary, then decide continuation.

If the user asks for another bounded batch or inventory cycle, treat it as endpoint/order update inside the current infinite sidecar.
Preserve ledger history, set the new bounded objective or stop condition, and choose the next bounded target before continuing.
If it is unclear whether the count includes the current loop or adds more loops, ask.

If the endpoint changes from infinite continuation to a finite stop condition, stop current infinite advancement.
Record the endpoint update in the sidecar and ledger.
If a finite sequence is needed, prepare a new finite sequence before work.

## Interruption

During active self-drive, interpret new user input in this order:

1. source-recorded explicit stop
2. approval-sensitive action or expanded approval boundary
3. scope, non-goal, endpoint, target, order, or acceptance change
4. blocker or repeated failure
5. status or progress question
6. ordinary note inside the recorded boundary

For status questions, report current identity, finite index or loop count, verification state, and next required action.
Continue only if the recorded sidecar still permits it.

Scope, endpoint, target, order, acceptance, or approval changes stop autonomous advancement and return to the earliest affected phase.
Endpoint changes are relock/update events: update sidecar endpoint, ledger, and affected flow record next action before continuing.

## Approval And Stop

Self-drive cannot approve commit, push, PR, publish, release, version bump, destructive action, external effect, or scope expansion.
Those actions need exact target, expected effect, risk, recovery path, included/excluded scope, and explicit approval.

If writes or record creation are forbidden, report in-memory continuity only.
Do not use unrecorded state for later autonomous advance.

Terminal closure requires a source-recorded explicit stop stored in the active flow record.
