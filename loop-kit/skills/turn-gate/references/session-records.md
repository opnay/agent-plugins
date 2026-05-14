# Session Records

Maintain records under `.agents/sessions/{YYYYMMDD}/` for active turn-gated tasks.

Use:

- `000-plan.md` as the date-level bounded plan, flow index, and current routing snapshot;
- `{count-pad3}-{eng-lower-slug}.md` as the canonical detail record for one flow.

Use `templates/plan-template.md` and `templates/flow-record-template.md` as starting structures.

## Plan Ownership

`000-plan.md` owns:

- latest user request and decision snapshot;
- active flow pointer and required next action;
- recent user request history;
- one-line flow index;
- current and future planned flow sequence;
- one-line completed flow summaries;
- active date-level open risks;
- date-level note that the user can explicitly end the turn.

It does not own detailed flow scope, non-goals, approval boundary, evidence, verification detail, per-flow residual risk, or canonical Continuity Guard state.

## Flow Record Ownership

Each `001+` flow record owns:

- original user request;
- task, flow type, scope, and parent plan;
- current phase;
- Continuity Guard;
- flow contract;
- optional risky actions;
- execution log;
- verification detail;
- report;
- next-flow options;
- residual risk.

Update the flow record after each phase instead of waiting for final completion.

## Flow Types

Use `operational-preparation` when the flow interprets a request, locks scope, designs a planned flow list, or records approvals without starting product/code/document execution.

Use `change-unit` when the flow owns a cohesive reviewable or commit-sized artifact change.

Keep follow-up change-unit candidates separate from active execution until the user selects or approves them.

## Continuity Guard

The Continuity Guard must track:

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
- pending question id or summary;
- superseded question id or summary;
- verification status;
- continuity note.

`verification status` is the result or lifecycle status, not the method. Use `not-started` or `requested` before verification completes, then `pass`, `fail`, `blocked`, or `insufficient`.

Record the verification method in the `Verification` section as one of `clean-context`, `normal`, or `not-required`. The method is separate from status:

- `clean-context`: bounded read-only verifier subagent.
- `normal`: same-context verification by command/check, source readback, evidence checklist, log review, or logical counterexample review.
- `not-required`: no work output required separate verification; record the reason, no-output rationale or existing evidence, and residual uncertainty.

Do not treat `not-required` as a pass. Do not use it for file changes, release surfaces, multi-file contracts, previous failed checks, user-requested verification, or approval-sensitive actions.

Only a source-recorded explicit stop can make terminal summary allowed. If closure source is missing or stale, reset `user explicit stop` to `no`, reset `terminal summary allowed` to `no`, and note the stale closure.

If records are inaccessible, report a blocker. Do not silently reconstruct records or treat missing state as permission to close.

## Next Flow Options

The flow record owns detailed next-flow options. Even when visible choices omit a turn-end option, record an explicit turn-end option. Reflect only the selected result or active next-flow pointer in `000-plan.md`.
