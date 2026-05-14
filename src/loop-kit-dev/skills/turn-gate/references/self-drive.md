# Self-Drive Overlay

Use this reference only when the user explicitly asks for autonomous continuation across a prepared planned flow sequence.

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
- current flow label;
- allowed autonomous actions;
- prohibited autonomous actions;
- approval-sensitive checkpoints;
- endpoint;
- blocker return conditions;
- progress note and progress ledger.

Do not add self-drive-only sequence fields to general templates by default. If self-drive is not active, do not create `000-self-drive.md`.

Each active flow record should only mirror the flow-local sequence position, local progress note, next handoff, and blocker return condition. Put that note in the most natural existing section, such as `Flow Contract`, `Execution Log`, `Report`, or `Next Flow Options`. Do not duplicate the full sequence contract in every flow record.

If `000-plan.md` says self-drive is inactive but still points at `000-self-drive.md`, or if a sidecar file remains after self-drive ended, treat it as stale sidecar state. Do not use the leftover sidecar as active continuation authority. Clear or correct the pointer/status when the current routing state is clear, or ask a user-gated clarification before autonomous continuation. You may read the old sidecar only as historical context after marking it stale.

## Priority

While self-drive applies, its continuation rules take priority over default next-flow questioning and ordinary phase protocol selection. Each planned flow still runs through preparation, work, verification, reporting, and next-flow internally.

Treat `active_flow_index` as a 0-based machine field. Always pair it with a human-readable current flow label, such as flow number, name, file, or slug. If an existing numeric index conflicts with numbered planned flows, or if the index base is unclear, inspect the flow names/files and reconcile the record before continuing. If the current flow still cannot be identified unambiguously, return to user-gated question routing instead of advancing silently.

Before reporting and before moving to the next planned flow, refresh `000-self-drive.md` so the active flow index, current flow label, sequence progress note, next handoff, and blocker state match the current result. Refresh `000-plan.md` only if the self-drive active status, sidecar pointer, active flow pointer, required next action, or date-level index changed. Refresh the active flow record separately with only the local progress note, next handoff, and blocker return condition.

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

## Ending

When the prepared sequence ends, do not close by default. Continue to the recorded endpoint, commit-readiness reporting handoff, blocker decision, or next-flow reopening. Terminal closure still requires a source-recorded explicit stop.
