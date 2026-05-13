# Self-Drive Overlay

Self-drive is an explicit overlay for a prepared planned flow sequence. It is not a separate installed skill entrypoint and not the default behavior.

Apply self-drive only after preparation has recorded:

- planned flow sequence;
- scope and non-goals;
- acceptance signal;
- approval boundary;
- verification expectation;
- stopping point or handoff condition.

While self-drive applies, its continuation rule controls movement through the prepared sequence. Each flow still runs `preparation -> work -> verification -> reporting -> next-flow`.

## Execution Authority

Self-drive can continue without asking after each flow only inside the prepared sequence and recorded approval boundary.

It may execute approval-sensitive actions only when the initial agreement recorded exact action, target, expected effect, risk, recovery path, included and excluded scope, and stopping point.

Commit, push, PR, publish, release, and version bump are approval-sensitive execution steps. If they were not explicitly included with a clear endpoint, return to user-gated question-routing.

## Return To Question-Routing

Return to the default user-gated path when:

- a new risky action appears;
- scope expands beyond the prepared sequence;
- the endpoint becomes unclear;
- verification is `blocked` or `insufficient`;
- a user decision is needed to choose among valid next flows.

When the prepared sequence ends, do not close by default. Continue to the recorded endpoint, commit-readiness handoff, or next-flow reopening.
