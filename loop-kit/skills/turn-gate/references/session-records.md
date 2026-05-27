# Session Records

Use this reference whenever `turn-gate` is active and records must be created, refreshed, recovered, or inspected after interruption or compaction.

## Files

- `.agents/sessions/{YYYYMMDD}/000-plan.md`: date-level routing context, active flow pointer, required next action, request history, compact flow index, planned current or future sequence, current/planned flow skill list, completed summaries, explicit turn-end availability, active date-level risks, and self-drive sidecar pointer.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`: one active flow's contract, raw request when needed, interpretation, scope, non-goals, approval boundary, execution log, verification, report, next-flow options, and residual risk.
- `.agents/sessions/{YYYYMMDD}/000-self-drive.md`: optional sequence-level self-drive state, used only when self-drive is active.

Use the bundled templates in `templates/` when creating records:

- `templates/plan-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`

Flow filenames use a zero-padded counter and a lowercase English slug. Create a new flow record when the active flow boundary changes. Update an existing flow record only when the same flow remains active or when correcting that flow's own Continuity Guard before reporting.

## Phase Checkpoints

Refresh records incrementally. Do not wait for a completed flow.

At each active flow phase start, record enough state to reconstruct:

- current phase
- scope boundary
- required next action
- required skills for the current flow, when skill use is part of the flow contract
- pending question or blocker state
- whether the change belongs in `000-plan.md` or the active flow record

At each active flow phase end, record enough state to reconstruct:

- phase result
- next phase or next required action
- verification status change
- residual risk
- handoff or next-flow condition
- whether the change belongs in `000-plan.md` or the active flow record

Update `000-plan.md` when the active flow pointer, date-level required next action, planned/current sequence, current/planned flow skill list, self-drive status pointer, or turn-level routing changes.

Update the active flow record when the same flow's current phase, execution log, verification evidence, report outcome, residual risk, handoff condition, pending question, or blocker state changes.

Phase checkpoints are record maintenance checkpoints. They do not turn `preparation`, `work`, `verification`, or `reporting` into separate flows.

If a trivial read-only judgment appears to need no record mutation, record the reason in the active flow record or report so a later agent can reconstruct why no change was needed.

## Minimum Flow Record Sections

Even a compact flow record must include:

- `Flow Contract`
- `Optional Risky Actions`
- `Execution Log`
- `Verification`
- `Report`
- `Next Flow Options`
- `Residual Risk`

The flow record frontmatter or Continuity Guard must expose current phase, required next action, closure fields, pending or superseded question state, verification status, and continuity note.

## Avoid Duplication

Keep `000-plan.md` compact. Do not repeat detailed scope, evidence, verification, residual risk, or self-drive sequence detail there. Store that detail in the active flow record or `000-self-drive.md`.

Keep `Flow Index` and `Completed Flow Summaries` as one compact line per flow. Do not delete completed summaries.

Keep `Planned Flow Sequence` limited to selected current or future flows. Handoff candidates from discovery or planning are not active or completed flows until selected.

Keep `Flow Skill List` compact: list only skill names and usage points for active or selected future flows. Do not copy full skill text, candidate-only skills, or completed-flow detail into `000-plan.md`.

When self-drive is active, `000-plan.md` stores only status and sidecar pointer. `000-self-drive.md` owns sequence-level state.

## Raw Requests

When storing raw user request text, keep it separate from interpretation or summary. Do not normalize, translate, correct, soften, merge, or infer missing words inside the raw request field.

You may write a separate compact interpretation, but it must not replace the raw source when raw text matters.

## Continuity Guard

Every active flow record must maintain a Continuity Guard with:

- `turn_gate_active`
- `question_routing_mode`
- `user_explicit_stop`
- `terminal_summary_allowed`
- `required_next_action`
- `last_refreshed_phase`
- `confirmed_closure`
- `closure_source_message`
- `closure_recorded_phase`
- `pending_question_state`
- `pending_question_id_or_summary`
- `superseded_question_id_or_summary`
- `verification_status`
- `continuity_note`

Refresh the guard at each phase start and end, before reporting, and before next-flow reopening.

Only a current source-recorded explicit stop can set terminal closure authority. Reset stale or source-less closure state to open continuity and record the recovery.

The guard must allow recovery after compaction or interruption. At minimum, it must show whether `turn-gate` is active, whether a pending question exists, verification status, whether closure is allowed, and the next required action.

Verification status may include progress states:

- `not-started`: verification has not begun.
- `requested`: a verifier or check has been requested but no result exists yet.
- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-started` and `requested` are not success evidence and cannot authorize terminal closure.

## Recovery

Distinguish record states carefully:

- not-yet-created plan: create the first plan from `templates/plan-template.md`.
- not-yet-created flow: create the selected flow record from `templates/flow-record-template.md`.
- unexpectedly missing active record: route to blocker recovery or ask the user how to recover.
- inaccessible active record: report the access blocker until access is restored or the user chooses recovery.
- stale closure state: reset closure authority and record the recovery.
- stale self-drive sidecar: if `000-plan.md` says self-drive is inactive, treat the sidecar as historical.
- stale routing mismatch: reconcile from the latest source record or ask a clarification.

Do not silently reconstruct an active record that should already exist.

At a recovered flow start, reread the skills named in `000-plan.md` for the current flow. If the list is stale or missing, reconstruct only the minimum current-flow skill list from source-of-truth records before work resumes.

Read-only requests usually restrict target/source changes, not session records. Do not write session records only when the user explicitly forbids all writes or record creation.

If wording such as `no-record`, `do not record`, or `without session records` is ambiguous about whether reading existing session records is also forbidden, ask for clarification before reading them. Until clarified, maintain only minimal in-memory continuity.
