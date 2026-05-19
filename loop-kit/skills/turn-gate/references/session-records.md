# Session Records

Use session records to preserve turn continuity and prevent terminal closure from stale or missing state.

## Paths

For each active turn-gated task, maintain records under:

```text
.agents/sessions/{YYYYMMDD}/
├── 000-plan.md
├── 000-self-drive.md        # only when self-drive is active
└── {count-pad3}-{eng-lower-slug}.md
```

Use `001`, `002`, `003` style zero-padded flow numbers. Slugs use lowercase English words and hyphens.

## Ownership

`000-plan.md` is the date-level index and snapshot. It owns recent user requests, active flow pointer, required next action, compact flow index, selected or future planned flow sequence, completed flow summaries, explicit turn-end availability, and date-level open risks.

Each `001+` flow record is the canonical detail artifact for one flow. It owns the flow contract, work boundary, non-goals, approval boundary, material judgment calls, execution log, verification detail, report, next-flow options, and flow-local residual risk.

`000-self-drive.md` exists only when self-drive is active. It owns sequence-level state. `000-plan.md` only keeps self-drive active status and pointer; active flow records keep only flow-local self-drive snapshots.

Do not repeat detailed scope, evidence, verification logs, or flow-local risk in `000-plan.md`. Keep plan entries compact and link or point to the flow record.

## Creating Records

Use these installed templates for first creation:

- `templates/plan-template.md` for `000-plan.md`
- `templates/flow-record-template.md` for `001+` flow records
- `templates/self-drive-template.md` for `000-self-drive.md`

Only create `000-self-drive.md` when self-drive is active.

If a new flow record is not yet created because the flow was just selected, template-based first creation is allowed. If an active plan or handoff points to a missing existing flow record, treat it as unexpectedly missing and do not silently reconstruct it.

## Continuity Guard

Every flow record must expose these Continuity Guard fields in frontmatter or equivalent visible state:

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

Refresh the Continuity Guard before result reporting and before next-flow reopening.

`terminal_summary_allowed: true` is valid only when the current incoming message or a source-recorded explicit stop supports it. A source-less closure or stale closure state is not terminal authority.

## Recovery Rules

Separate these cases:

- `not-yet-created plan`: first turn-gated work for the date; create from template.
- `not-yet-created flow`: newly selected flow before first record write; create from template.
- `unexpectedly missing active record`: plan or handoff points to a missing `001+` record; report blocker or ask for user-gated recovery.
- `inaccessible active record`: read failure, permissions, lock, parse failure, encoding failure, or partial write; report blocker and do not use it as closure authority.
- `stale closure state`: closure has no source message or does not match the current incoming message; reset explicit stop and terminal summary permission to false and record a note.
- `stale self-drive sidecar`: plan says self-drive inactive but sidecar remains; treat sidecar as historical context only.
- `stale routing mismatch`: plan, sidecar, flow record, and handoff disagree; reconcile from the latest reliable source or ask the user.

Never use stale closure, inaccessible records, missing records, or routing mismatch as a reason to close the turn.

## Read-Only And No-Write Requests

`read-only`, `no-edit`, `do not modify source`, or similar requests usually forbid target/source changes, not operational session records. Record that boundary separately.

If the user says not to write any files, not to create files, not to keep session records, or asks for no-record operation, do not write session records. If record reads or writes are ambiguous, ask before accessing or writing them when needed for continuity.

Clean-context verifier read-only means the verifier cannot edit or expand scope. It does not automatically forbid main turn-gate session records unless the user also forbids records.

## Next Flow Options

Each flow record owns `Next Flow Options`. Even if the visible `request_user_input` choices cannot show a turn-end option, record that explicit turn-end remains available.

`000-plan.md` should only snapshot the selected result or active next-flow pointer. It should not duplicate the full option list.
