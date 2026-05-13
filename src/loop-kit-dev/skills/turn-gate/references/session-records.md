# Session Records

Maintain date-scoped session records for active turn-gated work:

- `.agents/sessions/{YYYYMMDD}/000-plan.md`
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`

Use [plan-template.md](../templates/plan-template.md) and [flow-record-template.md](../templates/flow-record-template.md).

## Plan Record

`000-plan.md` is a bounded date-level index and active snapshot. It owns:

- latest user request and decision;
- active flow pointer;
- required next action;
- user request history;
- flow index;
- planned flow sequence;
- completed flow summaries;
- active date-level risks;
- explicit turn-end option availability.

Do not copy detailed flow contract, approval boundary, command evidence, or verification details into the plan. Those belong in the flow record.

`Planned Flow Sequence` must list cohesive flow units, not phase checklists. Completed flows move to compact index and summary entries.

## Flow Record

Each `001+` record is the canonical detail artifact for one flow. Create or update it incrementally after each phase.

It owns:

- original user request;
- task, flow type, scope, and parent plan;
- current phase;
- Continuity Guard;
- Flow Contract;
- Optional Risky Actions;
- Execution Log;
- Verification;
- Report;
- Next Flow Options;
- Residual Risk.

Flow type is `operational-preparation` or `change-unit`. An `operational-preparation` flow may end with follow-up `change-unit` candidates rather than active execution flows.

## Continuity Guard

Keep the Continuity Guard current, especially before reporting and next-flow reopening. It must show:

- turn-gate active;
- question-routing mode;
- user explicit stop;
- terminal summary allowed;
- required next action;
- last refreshed phase;
- confirmed closure;
- closure source message;
- closure recorded phase;
- pending question state;
- pending or superseded question summary;
- verification status;
- continuity note.

Valid verification statuses are `not-started`, `requested`, `pass`, `fail`, `blocked`, and `insufficient`.

Only a source-recorded explicit stop can allow terminal close. Source-less closure or stale `terminal_summary_allowed: yes` must be repaired to `user explicit stop: no` and `terminal summary allowed: no`, with a note.

## Next Flow Options

The flow record's `Next Flow Options` must include an explicit turn-end option even when the visible user prompt does not show one. The plan may store only the selected result or active next-flow pointer.
