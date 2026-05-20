# turn-gate session records

Use session records as operational continuity artifacts for active turn-gated work.

## Files

For each active date, maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` from `templates/plan-template.md`.

For each active flow, maintain `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` from `templates/flow-record-template.md`.

When self-drive is active, also maintain `.agents/sessions/{YYYYMMDD}/000-self-drive.md` from `templates/self-drive-template.md`.

Use `001`, `002`, `003` style zero-padded counters. Slugs use lowercase English words and hyphens only.

## Ownership

`000-plan.md` owns date-level routing context: latest request, active flow pointer, required next action, request history, compact flow index, planned current/future sequence, completed flow summaries, explicit turn-end availability, and active date-level risks.

The `001+` flow record owns canonical detail for one flow: raw request when needed, interpretation, scope, non-goals, approval boundary, execution log, verification detail, report, next-flow options, and residual risk.

When self-drive is active, `000-plan.md` stores only self-drive status and the sidecar pointer. `000-self-drive.md` owns sequence-level state. Active flow records store only flow-local self-drive snapshots.

Do not repeat detailed scope, evidence, verification, or residual risk in both plan and flow record. Keep plan entries compact and link or point to flow records for detail.

## Continuity Guard

Each flow record must keep a Continuity Guard and refresh it before result reporting and next-flow reopening. It must cover:

- turn-gate active state
- question-routing mode
- user explicit stop
- terminal summary allowed
- required next action
- last refreshed phase
- confirmed closure
- closure source message
- closure recorded phase
- pending question state
- pending question id or summary
- superseded question id or summary
- verification status
- continuity note

Only a source-recorded explicit stop can allow terminal close. If a closure is source-less or stale, reset user explicit stop and terminal summary allowed to `no`, and record that recovery in the continuity note.

## Recovery

Separate these cases:

- `not-yet-created plan`: create the first `000-plan.md` for the date from the template.
- `not-yet-created flow`: create the selected new flow record from the template.
- `unexpectedly missing active record`: report a blocker or ask for recovery; do not silently reconstruct.
- `inaccessible active record`: report a blocker until access is restored or the user chooses recovery.
- `stale closure state`: reset closure authority and note the stale state.
- `stale self-drive sidecar`: treat the sidecar as historical if plan says self-drive is inactive.
- `stale routing mismatch`: reconcile from the latest source/handoff or ask; do not choose the more closed state.

Record access blockers use the phase where they are found. Before result reporting use `[reporting]`; before next-flow reopening use `[next-flow]`.

## Read-Only Boundaries

Read-only or no-edit requests usually prohibit target/source changes, not session records. Continue maintaining session records unless the user explicitly says not to write any file, not to create any file, not to leave session records, or to answer without records.

If the user forbids all writes or records, do not create or update session files. Route to clarification or blocker with the minimum in-memory continuity needed.
