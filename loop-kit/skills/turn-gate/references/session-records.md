# Session Records

`turn-gate` records preserve continuity for `.agents/sessions/{YYYYMMDD}/`. They store active turn state, sibling `flow` decision fields, verification state, questions, and next action. They do not define how to judge flow boundaries, candidates, phases, or completion.

## Files

- `000-plan.md`: date-level routing snapshot and bounded index
- `000-self-drive.md`: optional self-drive sidecar, only when self-drive is active
- `{count-pad3}-{eng-lower-slug}.md`: per-flow canonical detail record

Use runtime templates:

- `templates/plan-template.md`
- `templates/self-drive-template.md`
- `templates/flow-record-template.md`

## Plan Record

`000-plan.md` owns:

- latest user request and decision snapshot
- active flow pointer and required next action
- recent user requests
- compact flow index
- `Planned Flow Sequence` as a date-level routing snapshot
- completed flow summaries
- open date-level risks
- explicit turn-end availability snapshot
- self-drive active status and sidecar pointer

It does not own detailed scope, non-goals, approval boundary, command output, evidence, or flow-local residual risk. Keep those in the `001+` flow record.

## Flow Record

Each `001+` flow record owns the canonical detail for one active flow:

- user request raw text when exact wording matters
- user request summary or interpretation
- task and flow type from the sibling `flow` decision
- scope, non-goals, acceptance signal, approval boundary, verification expectation
- current phase
- Continuity Guard
- execution log
- verification method and result status
- report
- next-flow options
- residual risk

Update the flow record incrementally after each phase. Do not wait until the flow is complete.

## Continuity Guard

Before reporting and before next-flow reopening, refresh:

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

Only a source-recorded explicit stop can allow terminal summary. If closure source is missing or stale, reset stop/summary fields to `no` and record the stale-state note.

## Verification Fields

Record method and status separately.

Methods:

- `clean-context`
- `normal`
- `not-required`

Statuses:

- `not-started`
- `requested`
- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-required` is a method, not a passing status. Include reason and residual uncertainty.

## Recovery

Separate these states:

- `not-yet-created plan`
- `not-yet-created flow`
- `unexpectedly missing active record`
- `inaccessible active record`
- `stale closure state`
- `stale self-drive sidecar`
- `stale routing mismatch`

Create from template only for first creation. If an active pointer expects a record and it is missing or inaccessible, route as blocker or ask the user. Do not silently reconstruct and proceed as complete.

If plan, self-drive sidecar, active flow record, and handoff disagree, reconcile from the newest reliable source. If unresolved, return to user-gated clarification.

## Read-Only Boundary

`read-only`, `no-edit`, or source-change bans usually prevent target/source edits, not operational session records. If the user says no file writes, no session records, or no records, do not write session files.

If the boundary is unclear, ask before writing records.
