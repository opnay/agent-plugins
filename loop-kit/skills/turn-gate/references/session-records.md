# Session Records

Maintain records under `.agents/sessions/{YYYYMMDD}/` for active turn-gated tasks.

Use:

- `000-plan.md` as the date-level current recovery snapshot, compact flow table, and current routing pointer;
- `000-self-drive.md` as the optional same-level self-drive sidecar record, created only while self-drive is active;
- `{count-pad3}-{eng-lower-slug}.md` as the canonical detail record for one flow.

Use `templates/plan-template.md`, `templates/self-drive-template.md`, and `templates/flow-record-template.md` as starting structures.

## Plan Ownership

`000-plan.md` owns:

- latest user request and decision snapshot;
- active flow pointer and required next action;
- current state needed to recover the turn;
- compact one-line flow table;
- current and future planned flow sequence only when it is still relevant;
- self-drive active status and `000-self-drive.md` pointer when self-drive is active;
- active date-level open risks;
- date-level note that the user can explicitly end the turn.

It does not own detailed flow scope, non-goals, approval boundary, evidence, verification detail, per-flow residual risk, canonical Continuity Guard state, or self-drive sequence-level state.

When self-drive is active, `000-plan.md` owns only `self_drive_status` and `self_drive_record` as self-drive-specific fields. The general `Planned Flow Sequence` may be kept as a date-level routing snapshot, but sequence-level state belongs to `000-self-drive.md`.

Keep the flow table compact, one line per flow. Keep any planned sequence limited to current or future selected flows; completed flows should stay only as compact recovery entries. Keep `Open Risks` limited to active date-level risks.

Do not let `000-plan.md` become the full all-day transcript. Long chronological request history, duplicate completed summaries, historical next-flow choices, per-flow evidence, and stale planned sequence detail belong in the relevant `001+` flow record, not in active recovery context.

## Self-Drive Sidecar Ownership

`000-self-drive.md` owns sequence-level state only while self-drive is active:

- sequence objective;
- planned flow list;
- active flow index as a 0-based machine field;
- current flow label with human-readable number/name/file or slug;
- current sequence progress note, mirrored in frontmatter as `progress_note`;
- allowed autonomous actions;
- prohibited autonomous actions;
- approval-sensitive checkpoints;
- endpoint;
- blocker return conditions;
- append-only progress ledger.

It is not a replacement for `000-plan.md`, does not own the date-level flow index, does not own each flow's canonical scope, evidence, report, or closure authority, and does not grant commit/push/PR/publish/release/version-bump approval.

If `000-plan.md` says self-drive is inactive but still points at `000-self-drive.md`, or if an old sidecar file still exists, treat that as stale sidecar state. Do not use the leftover sidecar as active continuation authority. Clear or correct the pointer/status when the routing state is clear, or ask/record a user-gated clarification before autonomous continuation. You may read the old sidecar only as historical context after marking it stale.

## Flow Record Ownership

Each `001+` flow record owns:

- original user request raw text when it matters, kept separate from summary or interpretation;
- task, flow type, scope, and the path-derived parent-plan relation;
- current phase;
- Continuity Guard;
- flow contract;
- optional self-drive flow-local sequence snapshot when self-drive is active;
- optional risky actions;
- execution log;
- verification detail;
- report;
- next-flow options;
- residual risk.

Completed flow records may retain historical audit detail, but active recovery state should stay compact. After a flow has moved on, stale `Next Flow Options`, pending or superseded question state, empty closure fields, and default/no-op template fields should not read like current authority. Prefer a compact final snapshot unless unresolved blockers, non-pass verification, approval-sensitive actions, or uncommitted changed surfaces require fuller evidence.

The parent-plan relation means the same-date `.agents/sessions/{YYYYMMDD}/000-plan.md` indexes the flow record. Do not add a duplicate `parent_plan` frontmatter field during normal runtime; if a future export or pathless aggregation tool needs explicit parent metadata, that metadata belongs to the exporter rather than the base flow template.

Update the flow record after each phase instead of waiting for final completion.

When self-drive is active, the flow record mirrors only flow-local sequence position, local progress note, next handoff, and blocker return condition in the most natural existing section. Do not add a self-drive-only top-level template section by default, and do not duplicate the full sequence objective, planned flow list, active flow index, autonomous boundaries, approval checkpoints, or endpoint in every flow record.

## Flow Types

Use `operational-preparation` when the flow interprets a request, locks scope, designs a planned flow list, or records approvals without starting product/code/document execution.

Use `change-unit` when the flow owns a cohesive reviewable or commit-sized artifact change.

Keep follow-up change-unit candidates separate from active execution until the user selects or approves them. Final QA, consistency checks, verification result reporting, and commit-readiness reporting are not separate planned flows unless they create or change a reviewable artifact.

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

Record closure source and closure phase only when explicit stop has actually been source-recorded. If there is no closure, avoid treating empty closure values as meaningful active context.

`verification status` is the result or lifecycle status, not the method. Use `not-started` or `requested` before verification completes, then `pass`, `fail`, `blocked`, or `insufficient`.

Record the verification method in the `Verification` section as one of `clean-context`, `normal`, or `not-required`. The method is separate from status:

- `clean-context`: bounded read-only verifier subagent.
- `normal`: same-context verification by command/check, source readback, evidence checklist, log review, or logical counterexample review.
- `not-required`: no work output required separate verification; record the reason, no-output rationale or existing evidence, and residual uncertainty.

Do not treat `not-required` as a pass. Do not use it for file changes, release surfaces, multi-file contracts, previous failed checks, user-requested verification, or approval-sensitive actions.

Only a source-recorded explicit stop can make terminal summary allowed. If closure source is missing or stale, reset `user explicit stop` to `no`, reset `terminal summary allowed` to `no`, and note the stale closure.

## Record Recovery Boundary

Separate record recovery states before reporting or reopening next-flow:

| State | Use when | Action |
| --- | --- | --- |
| `not-yet-created plan` | today's plan does not exist because this is the first turn-gated work for the day | create from the plan template and record the current message as source |
| `not-yet-created flow` | a new flow was just selected and its `001+` detail record has not been created yet | create from the flow template before work/reporting |
| `unexpectedly missing active record` | `000-plan.md`, `000-self-drive.md`, current phase state, or previous handoff points to a specific flow record that is absent | report a blocker or ask for a user-gated recovery decision |
| `inaccessible active record` | the active record exists but cannot be trusted because of permission error, lock, parse failure, partial write, encoding failure, or similar access failure | report a blocker; retry only after access is restored or the user decides |
| `stale closure state` | closure source is missing, stale, or does not match the current incoming message | reset closure/terminal summary fields to `no` and note the stale state |
| `stale self-drive sidecar` | `000-plan.md` says self-drive is inactive but a sidecar pointer or file remains | treat the sidecar as historical context only |
| `stale routing mismatch` | `000-plan.md`, `000-self-drive.md`, active flow phase, or previous handoff point to conflicting phases or flows | reconcile from the latest source/handoff or ask for clarification |

Do not silently reconstruct unexpectedly missing or inaccessible active records. Do not treat missing state, pass verification, stale closure, stale routing mismatch, or a leftover self-drive sidecar as permission to close the turn.

## Read-Only Write Boundary

Ordinary `read-only`, `no-edit`, `only read files`, `do not change source`, or `do not touch code` requests normally restrict target/source/spec/runtime/release-surface changes. They do not, by default, forbid `.agents/sessions/{YYYYMMDD}/` operational records. Record that split in the flow boundary.

If the user explicitly forbids all file writes, file creation, session records, artifacts, or asks for a no-record answer, session records are also forbidden. Do not create or update session records without clarification. Keep any in-memory continuity only long enough to ask the user or report the blocker.

If the user forbids leaving session records but it is unclear whether reading existing records is also forbidden, ask before reading them. If the user asks for status and does not forbid record reads, read-only record inspection is allowed; writing remains forbidden.

If the user says "stop if edits are needed", treat target/source edits as the stopping condition. If that wording could also forbid operational records, ask before writing records.

Read-only verifier or subagent packets forbid edits by that verifier/subagent and changes to the verification target. They do not automatically forbid main-thread session records unless the user also gave an all-files/no-record constraint.

## Next Flow Options

The flow record owns detailed next-flow options. Even when visible choices omit a turn-end option, record an explicit turn-end option. Reflect only the selected result or active next-flow pointer in `000-plan.md`.
