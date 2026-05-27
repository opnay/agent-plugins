# Session Records

Use this reference whenever `turn-gate` is active and records must be created, refreshed, recovered, or inspected after interruption or compaction.

## Files

- `.agents/sessions/{YYYYMMDD}/000-plan.md`: compact date-level routing card. Frontmatter owns active flow pointer, next action, closure flags, self-drive status and sidecar pointer, unapproved actions, and active skill list. Body owns only compact recent requests, active/recent/archive flow index, and continuity note.
- `.agents/sessions/{YYYYMMDD}/000-review.md`: optional date-level retrospective notes. Use it for reusable lessons, process corrections, and follow-up candidates. It must not own active routing state, raw flow logs, verification authority, or closure authority.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`: one active flow's compact `Contract`, `Execution Log`, and `Result`, with raw request and `Risky Action` added only when needed.
- `.agents/sessions/{YYYYMMDD}/000-self-drive.md`: optional sequence-level self-drive state, used only when self-drive is active.

Use the bundled templates in `templates/` when creating records:

- `templates/plan-template.md`
- `templates/review-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`

Flow filenames use a zero-padded counter and a lowercase English slug. Create a new flow record when the active flow boundary changes. Update an existing flow record only when the same flow remains active or when correcting that flow's own continuity metadata before reporting.

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

Update `000-plan.md` frontmatter when the active flow pointer, date-level required next action, active skill list, self-drive status, unapproved action state, or turn-level routing changes.

Update the active flow record when the same flow's current phase, execution log, verification evidence, report outcome, residual risk, handoff condition, pending question, or blocker state changes.

Phase checkpoints are record maintenance checkpoints. They do not turn `preparation`, `work`, `verification`, or `reporting` into separate flows.

If a trivial read-only judgment appears to need no record mutation, record the reason in the active flow record or report so a later agent can reconstruct why no change was needed.

## Compact Flow Records

Use compact formal style for flow records by default. The required sections are:

- `Contract`
- `Execution Log`
- `Result`

Keep `Execution Log`. It is the main recovery trail. Compress other sections into field-like bullets.

Add `Risky Action` only when approval-sensitive action exists. Readiness, verification, build, and generated release surface updates do not by themselves authorize commit, publish, release, version bump, destructive work, or external side effects. Add raw request only when exact source wording affects interpretation. Add non-pass routing only for `fail`, `blocked`, or `insufficient` results.

The flow record frontmatter must expose phase, verification status, next action, flags, question state, and continuity note. Use `answered_question` by default, add `pending_question` when a question is waiting, and do not invent synonyms such as `question_state`. Prefer compact metadata such as `flags: [turn_gate_active, terminal_summary_blocked]` over long boolean lists.

## Avoid Duplication

Keep `000-plan.md` as a compact routing card. Do not repeat git-reconstructable branch/latest-commit state, detailed scope, evidence, verification, residual risk, or self-drive sequence detail there. Store that detail in the active flow record, `000-self-drive.md`, or fresh tool readback.

Store skill reread requirements in frontmatter `active_skills` as names only. Do not add usage explanations, candidate-only skills, completed-flow skill lists, or the full available skill catalog.

Keep `Flow Index` to active, recent, and archive entries only. `active` is the current flow, `recent` is the immediately previous handoff flow, and `archive` points to older flow ranges recoverable from individual flow records. Do not accumulate completed flow summaries or old planned flows in `000-plan.md`.

Keep `Recent Requests` short. Prefer `[current] compact note` and one or two routing-relevant previous entries over sentence-form request history. Older raw requests and interpretations belong in individual flow records.

When self-drive is active, `000-plan.md` stores only status and sidecar pointer. `000-self-drive.md` owns sequence-level state.

Use `000-review.md` only for retrospective notes that are useful after the current routing problem is solved. Keep it as a flat tagged list, not a flow-by-flow log or section-per-category document. Each item starts with one bracketed axis tag such as `[conversation]`, `[records]`, `[docs]`, `[code-structure]`, `[verification]`, `[git]`, `[release]`, or a task-specific tag. Tags are open-ended. Add compact sub-bullets only when they clarify invalid/correct examples, evidence, or follow-up candidates.

Do not use `000-review.md` to reconstruct the active turn. Active flow pointer, required next action, pending question, verification status, and closure state belong in `000-plan.md` and active flow records.

## Raw Requests

When storing raw user request text, keep it separate from interpretation or summary. Do not normalize, translate, correct, soften, merge, or infer missing words inside the raw request field.

You may write a separate compact interpretation, but it must not replace the raw source when raw text matters.

## Continuity Metadata

Every active flow record must maintain compact continuity metadata with:

- `phase`
- `verification_status`
- `next_action`
- `flags`
- `answered_question`, and `pending_question` when a question is waiting
- `continuity`

Use `flags` for recovery-relevant states such as `turn_gate_active`, `terminal_summary_blocked`, `question_pending`, `blocked`, `approval_required`, or `explicit_stop_recorded`. If a source-recorded explicit stop exists, keep the closure source in `continuity` or another compact field.

Refresh metadata at each phase start and end, before reporting, and before next-flow reopening.

Only a current source-recorded explicit stop can set terminal closure authority. Reset stale or source-less closure state to open continuity and record the recovery.

The metadata must allow recovery after compaction or interruption. At minimum, it must show whether `turn-gate` is active, whether a pending question exists, verification status, whether closure is allowed, and the next required action.

Verification status may include progress states:

- `not-started`: verification has not begun.
- `requested`: a verifier or check has been requested but no result exists yet.
- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-started` and `requested` are not success evidence and cannot authorize terminal closure.

When preserving prior flow state, keep the existing `verification_status` value and describe preservation in `continuity`. Do not write `verification_status: preserved`.

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
