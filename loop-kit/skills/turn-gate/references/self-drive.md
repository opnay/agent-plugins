# Self-Drive Overlay

Use this reference only when the user explicitly asks for autonomous continuation across a prepared planned flow sequence.

Self-drive is an overlay on `turn-gate`, not a separate installed skill entrypoint. It can apply only after preparation has recorded:

- planned flow sequence;
- scope and non-goals;
- acceptance signal;
- verification expectation;
- approval boundary;
- expected risky actions;
- user-gated checkpoint and endpoint conditions.

## Priority

While self-drive applies, its continuation rules take priority over default next-flow questioning and ordinary phase protocol selection. Each planned flow still runs through preparation, work, verification, reporting, and next-flow internally.

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
