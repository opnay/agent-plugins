# Self-Drive Overlay

Use this reference when the user explicitly asks for autonomous continuation across a prepared planned flow sequence, or when an already-active self-drive sequence receives another user message.

Self-drive is an overlay on `turn-gate`, not a separate installed skill entrypoint. It can apply only after preparation has recorded:

- planned flow sequence;
- self-drive sequence objective;
- active flow index;
- current flow label;
- scope and non-goals;
- acceptance signal;
- verification expectation;
- approval boundary;
- expected risky actions;
- allowed and prohibited autonomous actions;
- user-gated checkpoints, endpoint, blocker return conditions, and sequence progress note.

## Record Ownership

When self-drive is active, create or maintain `.agents/sessions/{YYYYMMDD}/000-self-drive.md` from `templates/self-drive-template.md`.

`000-plan.md` remains the date-level plan, index, and routing snapshot. For self-drive-specific fields, it owns only:

- `self_drive_status`;
- `self_drive_record`, pointing to `000-self-drive.md`.

The general `Planned Flow Sequence` in `000-plan.md` may remain a date-level routing snapshot. It is not the canonical self-drive sequence-level state.

Keep sequence-level state in `000-self-drive.md`:

- sequence objective;
- planned flow list;
- active flow index;
- current flow label, mirrored in frontmatter as `current_flow_label`;
- current sequence progress note, mirrored in frontmatter as `progress_note`;
- allowed autonomous actions;
- prohibited autonomous actions;
- approval-sensitive checkpoints;
- endpoint;
- blocker return conditions;
- progress note and append-only progress ledger.

Do not add self-drive-only sequence fields to general templates by default. If self-drive is not active, do not create `000-self-drive.md`.

Each active flow record should only mirror the flow-local sequence position, local progress note, next handoff, and blocker return condition. Put that note in the most natural existing section, such as `Flow Contract`, `Execution Log`, `Report`, or `Next Flow Options`. Do not duplicate the full sequence contract in every flow record.

If `000-plan.md` says self-drive is inactive but still points at `000-self-drive.md`, or if a sidecar file remains after self-drive ended, treat it as stale sidecar state. Do not use the leftover sidecar as active continuation authority. Clear or correct the pointer/status when the current routing state is clear, or ask a user-gated clarification before autonomous continuation. You may read the old sidecar only as historical context after marking it stale.

## Priority

While self-drive applies, its continuation rules take priority over default next-flow questioning and ordinary phase protocol selection. Each planned flow still runs through preparation, work, verification, reporting, and next-flow internally.

Self-drive does not remove the `next-flow` phase. It changes the `next-flow` outcome to recorded loop continuation when the prepared sequence is still valid and the next planned flow is identifiable. If continuation identity, scope, endpoint, approval boundary, or blocker state is unclear, pause autonomous continuation and return to user-gated routing.

Treat `active_flow_index` as a 0-based machine field. Always pair it with `current_flow_label`, a human-readable flow number, name, file, or slug. If the frontmatter `current_flow_label`, body `Current flow`, numeric index, or planned-flow list conflict, inspect the flow names/files and reconcile the record before continuing. If the current flow still cannot be identified unambiguously, return to user-gated question routing instead of advancing silently.

Before reporting and before moving to the next planned flow, refresh `000-self-drive.md` so the active flow index, current flow label, `progress_note`, next handoff, and blocker state match the current result. Treat `progress_note` as the current overwriteable sequence summary, and treat `Progress Ledger` as append-only history. Refresh `000-plan.md` only if the self-drive active status, sidecar pointer, active flow pointer, required next action, or date-level index changed. Refresh the active flow record separately with only the local progress note, next handoff, and blocker return condition.

## Mid-Sequence User Input

When a user message arrives while self-drive is executing a prepared sequence, treat it as self-drive mid-sequence input even if the message does not repeat the word "self-drive". Handle it inside the active flow before continuing.

This implicit self-drive context does not override explicit stop, approval boundary, scope/non-goal/endpoint locks, or user-gated routing. Apply this priority order:

1. If it is a source-recorded explicit stop, record closure state and stop after reporting.
2. If it requests a destructive, external, commit, push, PR, publish, release, version-bump, or other approval-sensitive action outside the exact recorded approval boundary, pause self-drive and return to user-gated approval routing.
3. If it changes scope, non-goals, endpoint, target, planned flow order, or acceptance signal, pause self-drive and return to preparation or next-flow routing to relock the updated sequence before continuing.
4. If it reports or reveals a blocker, route to the earliest safe phase or user-gated blocker decision.
5. If it only asks for status or progress, report the current phase, active flow, verification state, and next action, then continue self-drive unless the message also matches an earlier rule. Do not convert a status-only input into next-flow selection or terminal closure.
6. If it is an ordinary continuation note within the recorded boundary, record the note if material and continue.

This is self-drive interruption handling, not a general user-message taxonomy. Do not add a separate routing layer outside self-drive; use the existing meaning-resolution, approval-boundary, explicit-stop, and question-routing contracts.

Subagent packets may help with read-only evidence checks, status/progress synthesis, or low-risk local decisions inside the recorded boundary. They cannot replace user approval for approval-sensitive actions, scope changes, endpoint changes, planned-flow-order changes, or new external/destructive execution.

## Execution Authority

Self-drive may continue without asking only inside the prepared sequence and recorded approval boundary.

It may execute approval-sensitive actions only when the initial agreement records exact action, target, expected effect, risk, recovery path, included scope, excluded scope, and end point.

Commit, push, PR, publish, release, and version bump are approval-sensitive execution steps. If their exact boundary and end point were not explicitly recorded, return to user-gated question routing before executing them.

Return to user-gated question routing when:

- a new risky action appears;
- scope or non-goal changes;
- the endpoint is unclear;
- repeated critical failure suggests a root blocker;
- commit, push, PR, publish, release, or version bump is outside the exact approved boundary.

## Question Tool Boundary

Self-drive narrows question-tool use; it does not disable it. Do not ask before every recorded sequence transition, but do ask when user input is needed to keep the sequence safe and valid.

| Situation | Question tool routing |
| --- | --- |
| Clear next planned flow inside the recorded boundary | Do not ask; refresh records and continue. |
| Status/progress-only input | Usually do not ask; report current state and continue. |
| Scope, target, endpoint, order, non-goal, or acceptance signal changes | Ask or relock the sequence before work. |
| Approval-sensitive, destructive, external, commit, push, PR, publish, release, or version-bump action outside exact recorded boundary | Ask before execution. |
| Blocker, record access failure, repeated critical failure, or unclear current-flow identity | Ask or open blocker routing before continuing. |
| Non-self-drive result reporting with no explicit stop | Use default next-flow question routing. |

## Ending

When the prepared sequence ends, do not close by default. Continue to the recorded endpoint, commit-readiness reporting handoff, blocker decision, or next-flow reopening. Terminal closure still requires a source-recorded explicit stop.

Record even open-ended self-drive as bounded cycles. Do not use vague endpoints such as "forever" or "until stopped" without a current finite planned flow list, cycle exhaustion behavior, blocker return conditions, and approval boundary.

| Endpoint pattern | Required behavior at sequence exhaustion |
| --- | --- |
| Finite list exhaustion | Stop self-drive or hand off exactly as recorded; do not create new work silently. |
| Repeat inventory loop | Create the next bounded inventory cycle only if the endpoint explicitly says to repeat; refresh planned flow count, active flow index, current flow label, and progress note. |
| User-stop-only unbounded request | Convert into finite cycles with a recorded repeat policy and user-gated blockers; terminal closure still needs source-recorded explicit stop. |
| Unclear endpoint | Pause autonomous continuation and ask or open endpoint clarification. |
