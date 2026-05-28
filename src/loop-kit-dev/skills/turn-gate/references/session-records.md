# Session Records

Use this reference when `turn-gate` is active and session records must be created, updated, recovered, or checked after interruption or compaction.

## Files

- `.agents/sessions/{YYYYMMDD}/000-plan.md`: compact date-level routing card.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`: one active flow record.
- `.agents/sessions/{YYYYMMDD}/000-self-drive.md`: optional self-drive sequence state, only when self-drive is active.
- `.agents/sessions/{YYYYMMDD}/000-review.md`: optional retrospective notes.

Use bundled templates in `templates/`:

- `templates/plan-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`
- `templates/review-template.md`

Do not make users read dev-only specs at runtime. This reference and the templates are the runtime contract.

## 000-plan.md

Keep `000-plan.md` small. It owns routing state, not detailed history.

Frontmatter owns:

- `turn_gate_active`
- `active_flow`
- `next_action`
- closure state
- `self_drive` and `self_drive_sidecar`
- `unapproved_actions`
- `active_skills`

Body owns only:

- compact recent requests
- active/recent/archive flow index
- continuity note

Do not accumulate completed flow summaries, full history, detailed verification evidence, branch/latest-commit state, or self-drive sequence detail in `000-plan.md`.

## Flow Records

Create a new flow record when the active flow boundary changes.
Use a zero-padded counter and lowercase English slug.
Update an existing flow record only while the same flow remains active or when correcting that flow's own continuity metadata before reporting.

Every flow record keeps these default sections:

- `Contract`
- `Phase Checklist`
- `Execution Log`
- `Result`

Add `Risky Action` only for approval-sensitive actions.
Add raw request text only when exact source wording affects interpretation.
Add non-pass routing under `Result` only for `fail`, `blocked`, or `insufficient` results.

`Contract` must keep:

- `scope`
- `exclude`
- `done`
- `boundary`

Do not treat readiness, verification, build output, generated release surface, self-drive state, previous context, or subagent output as authority to commit, push, publish, release, version bump, or perform destructive/external work.

## Phase Checklist

`Phase Checklist` shows which required lifecycle phases have passed their end checkpoint.
Frontmatter `phase` shows the current location; checklist shows completed lifecycle steps.

Default checklist:

- `intake`
- `framing`
- `preparation`
- `work`
- `verification`
- `reporting`
- `next-flow`

Check a phase only after its end checkpoint is recorded.
Do not check a phase just because it started.

`interruption` is not a checklist item. It is an entry-only event while another flow is active. Record it in `Execution Log`, then continue with the lifecycle phase it returns to.

## Execution Log

Use `Execution Log` as the main recovery trail.
Keep entries factual and compact.

Record:

- phase start or end results
- pending and answered questions
- approval-sensitive checkpoint state
- edits, builds, checks, verifier results
- interruption classification and routing
- reporting outcome and next-flow reopening

Do not hide failed checks, skipped verification, insufficient evidence, blockers, or approval gaps.

## Continuity Metadata

Every active flow record frontmatter must expose:

- `phase`
- `verification_status`
- `next_action`
- `flags`
- `answered_question`
- `pending_question`, when a question is waiting
- `continuity`

Use `flags` only for recovery-relevant states, such as:

- `turn_gate_active`
- `terminal_summary_blocked`
- `question_pending`
- `blocked`
- `approval_required`
- `explicit_stop_recorded`

Use only these `verification_status` values:

- `not-started`
- `requested`
- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-started` and `requested` are not success evidence and cannot authorize terminal closure.
When preserving a previous flow state, keep the existing status and record the preservation in `continuity`; do not create a `preserved` status.

Refresh metadata at phase start, phase end, before reporting, and before next-flow reopening.
Only a current source-recorded explicit stop can authorize terminal closure.
Reset stale or source-less closure state to open continuity and record the recovery.

## Questions And Next Flow

After reporting, reopen routing unless the current user message explicitly stopped the turn and that stop source is recorded.

Use `request_user_input` when available for narrow next-flow choices, clarifications, blocker recovery, approval-boundary decisions, and pending question recovery.
If it is unavailable, ask with a concise plain-text fallback.

An aborted, canceled, or interrupted question tool call is not flow completion and is not terminal closure.
Record the pending question, keep `terminal_summary_blocked`, and interpret the next user message as one of:

- pending question answer
- superseding new flow request
- status/progress question
- explicit stop

Use `answered_question` and `pending_question` for question state.
Do not invent `question_state` or similar synonyms.

## Recovery

Distinguish missing states:

- not-yet-created plan: create the first plan from the runtime template.
- not-yet-created flow: create the selected flow record from the runtime template.
- unexpectedly missing active record: report blocker or ask for recovery choice.
- inaccessible active record: keep blocker until access is restored or the user chooses recovery.
- stale closure state: reset closure authority and record recovery.
- stale self-drive sidecar: if plan says self-drive is inactive, treat the sidecar as historical.
- stale routing mismatch: reconcile from the latest source or ask a clarification.

Do not silently reconstruct an active record that should already exist.

Read-only requests usually restrict target/source changes, not session records.
Do not write session records only when the user explicitly forbids all writes or record creation.
If `no-record`, `do not record`, or similar wording is ambiguous about reading existing records, ask before reading them.
