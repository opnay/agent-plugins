# Session Records

Use this reference when `turn-gate` is active and records are needed to recover handoff question routing, pending question state, explicit stop state, or self-drive gate state.

## Ownership

`flow` owns shared record template meaning and file rules for:

- `plan.md`
- `flow-record.md`
- `review.md`

`turn-gate` owns active-turn recovery:

- active flow pointer
- required next action
- pending or answered question state
- verification status routing
- explicit stop state
- self-drive pointer
- unapproved action state
- active skill list for the next-flow question

`turn-gate` owns only one bundled template:

- `templates/self-drive-template.md`

Do not invent sibling filesystem paths between installed skills.
Do not make users read dev-only specs at runtime.

## Record Application

Use the `flow` skill's bundled templates when a plan, flow record, or review file must be created.
Use the turn-gate self-drive template only when self-drive is active.

Keep `000-plan.md` small.
It is a date-level routing card, not full history.

Create a new flow record only when the active flow boundary changes.
Update the current flow record while the same flow remains active, or when correcting that flow's own continuity metadata before reporting.

Use `000-review.md` only for retrospective notes.
It is not active routing, raw flow log, verification authority, closure authority, or commit/release authority.

## Recovery

Record enough state to recover after compaction or interruption:

- current active flow
- current phase or wrapper state
- verification status
- pending question, if any
- next required action
- explicit stop source, if any
- self-drive status and sidecar pointer, if active

If an expected active record is missing or inaccessible, route to blocker recovery.
Do not silently reconstruct a record that should already exist.

If closure state is stale or source-less, reset closure authority and record recovery.
Only a current source-recorded explicit stop can authorize terminal closure.

Read-only source work usually does not forbid session record writes.
If the user forbids all writes or record creation, do not write records and do not use that in-memory state as autonomous continuation authority.

## Questions

After `flow skill: handoff`, enter `<gate:next-flow>`: identify the full session active skill list, reread each skill body, accept the refreshed active skill set, reopen handoff question routing unless explicit stop is recorded, and update `000-plan.md`.
Do not record handoff, final-looking reporting, status-only output, or verification pass as closure authority.
Use `answered_question` and `pending_question` for question recovery.
Do not invent alternate question-state fields.

An aborted, canceled, or interrupted question tool call is not flow completion and is not terminal closure.
