# Self-Drive Overlay

Use this reference only when the user explicitly asks for autonomous continuation across a prepared planned flow sequence.

Self-drive is an overlay on `turn-gate`, not a separate installed skill entrypoint. It can apply only after preparation has recorded:

- planned flow sequence;
- self-drive sequence objective;
- active flow index;
- scope and non-goals;
- acceptance signal;
- verification expectation;
- approval boundary;
- expected risky actions;
- allowed and prohibited autonomous actions;
- user-gated checkpoints, endpoint, blocker return conditions, and sequence progress note.

## Sequence Record

When self-drive is active, keep the sequence-level state in `.agents/sessions/{YYYYMMDD}/000-plan.md`:

- sequence objective;
- planned flow list;
- active flow index;
- allowed autonomous actions;
- prohibited autonomous actions;
- approval-sensitive checkpoints;
- endpoint;
- blocker return conditions;
- progress note.

Do not add self-drive-only fields to every general template by default. Instead, when self-drive is active, add the sequence-level note to `000-plan.md` near `Planned Flow Sequence` or the active routing snapshot.

Each active flow record should only mirror the flow-local sequence position, local progress note, next handoff, and blocker return condition. Put that note in the most natural existing section, such as `Flow Contract`, `Execution Log`, `Report`, or `Next Flow Options`. Do not duplicate the full sequence contract in every flow record.

## Priority

While self-drive applies, its continuation rules take priority over default next-flow questioning and ordinary phase protocol selection. Each planned flow still runs through preparation, work, verification, reporting, and next-flow internally.

Before reporting and before moving to the next planned flow, refresh `000-plan.md` so the active flow index, sequence progress note, and blocker state match the current result. Refresh the active flow record separately with only the local progress note, next handoff, and blocker return condition.

## Execution Authority

Self-drive may continue without asking only inside the prepared sequence and recorded approval boundary.

It may execute approval-sensitive actions only when the initial agreement records exact action, target, expected effect, risk, recovery path, included/excluded scope, and end point.

Return to user-gated question routing when:

- a new risky action appears;
- scope or non-goal changes;
- the endpoint is unclear;
- repeated critical failure suggests a root blocker;
- commit, push, PR, publish, release, or version bump was not explicitly approved with exact boundaries.

## Ending

When the prepared sequence ends, do not close by default. Continue to the recorded endpoint, commit-readiness reporting handoff, blocker decision, or next-flow reopening. Terminal closure still requires a source-recorded explicit stop.
