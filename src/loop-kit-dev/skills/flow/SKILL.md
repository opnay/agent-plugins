---
name: flow
description: Interpret a message, action, plan item, review finding, or handoff as a cohesive flow; decide whether it is an active flow, parent flow, finite sub-flow candidate, phase, or handoff; and define readiness, discovery, ambiguity, flow-local strategy, phase record checkpoints, verification expectation, and commit-readiness handoff.
---

# Flow

Use this skill to turn ambiguous work into a bounded flow contract before execution. A flow is a cohesive work unit that can be understood, reviewed, verified, and handed off. It is not just a phase label such as `analysis`, `work`, `verification`, `reporting`, or `commit readiness`.

Each active flow has internal phases: `preparation -> work -> verification -> reporting`. These phases remain inside the same active flow; do not split them into separate flows only because a phase starts or ends.

## Flow Boundary

Classify the current item before acting:

- `active flow`: the selected flow currently being prepared, executed, verified, or reported.
- `parent flow`: a flow whose result is a finite list of `sub-flow candidates`.
- `sub-flow candidate`: a proposed later flow with its own scope, non-goals, completion criteria, verification expectation, and handoff condition. Creating candidates is not execution.
- `phase`: an internal step of the active flow, not a flow by itself.
- `handoff`: a readiness or routing result, not execution authority.

Pure final QA, consistency checks, verification reporting, and commit-readiness reporting are not separate change-unit flows unless they create or modify a reviewable artifact. A fixture, snapshot baseline, operator report, validator output, code change, document change, config change, or release surface change can be a flow when it owns a reviewable artifact.

## Flow Types

Use two default flow types:

- `operational-preparation`: locks intent, scope, non-goals, success signal, verification expectation, approval boundary, and planned flow candidates or sequence.
- `change-unit`: owns a reviewable change to code, docs, fixtures, config, release surface, or another artifact.

An operational-preparation flow may produce change-unit candidates. Those candidates stay inactive until selected by the surrounding routing or an explicitly prepared sequence.

## Output Contract

When designing a flow or sub-flow candidate, make these fields visible:

- flow label or slug
- flow type: `operational-preparation` or `change-unit`
- scope
- non-goals
- completion criteria
- verification expectation
- phase start/end record checkpoint expectation
- readiness status or missing contract fields
- recommended question topics or unresolved ambiguity
- recommended flow-local strategy
- approval-sensitive checkpoint, if any
- handoff condition
- unresolved questions or blocker
- whether the item is an active flow or a sub-flow candidate

Keep completion criteria separate from handoff condition. Keep missing fields, question topics, and recommended strategies separate from execution authority.

## Preparation

Before work begins, check readiness. The flow contract must cover scope, non-goals, completion criteria, verification expectation, and handoff condition. If the operation, target, endpoint, approval boundary, or verification path changes the result, lock it before work.

Use discovery when user intent, included scope, non-goals, success criteria, verification expectation, or candidate decomposition is missing. Produce bounded question topics and missing fields; do not decide how the question is routed to the user.

Use ambiguity handling when an operation or target can point to multiple structures and the interpretation changes scope, output, verification, approval sensitivity, or handoff. Record the interpreted operation, target, alternate interpretations, and impact of ambiguity when needed.

## Phase Record Checkpoints

At the start and end of each active-flow phase, decide whether `000-plan.md` or the active flow record needs an update. `flow` defines the checkpoint expectation; the active turn controller applies record updates when it owns that runtime surface.

Use these checkpoint rules:

- `preparation` start: expose the current flow label, type, scope boundary, pending contract fields, and next preparation action.
- `preparation` end: expose readiness status, locked scope and non-goals, missing questions or blockers, selected strategy, and whether work may begin.
- `work` start: expose the active flow boundary, required next action, approval-sensitive checkpoint status, and expected work artifact.
- `work` end: expose changed artifact or work result, issue found, next phase, and whether verification expectation changed.
- `verification` start: expose verification method expectation, evidence needed, target surface, and any known limitation.
- `verification` end: expose pass/fail/blocked/insufficient status, evidence gathered, residual risk, and earliest safe next phase if verification did not pass.
- `reporting` start: expose reportable result, verification status, handoff condition, and unresolved question or blocker.
- `reporting` end: expose reported outcome, residual risk, handoff condition, and any next-flow candidate without executing it.

Update `000-plan.md` expectation when the active flow changes or the turn-level required next action changes. Update active flow record expectation when flow-local phase state, execution evidence, verification evidence, report outcome, or residual risk changes. If a trivial read-only judgment needs no record change, make the reason visible in the active flow record or report.

## Flow-Local Strategies

Choose a strategy inside the active flow only after readiness is sufficient:

- `review-loop`: handle one bounded blocking review, QA, or self-review finding that materially affects correctness, regression risk, reliability, or delivery. If there are multiple findings, choose the highest-priority bounded finding or produce finite follow-up candidates.
- `fix-verify-loop`: use one small fix or confirmation action to test one primary issue, verify immediately, then reassess whether another loop is justified.
- `broad-execution`: execute a single locked active flow end to end when scope, non-goals, completion criteria, verification expectation, and approval boundary are clear.

Do not use a flow-local strategy as authority to continue through multiple flows. If the issue introduces new scope, a new approval boundary, destructive or external effects, or changed completion criteria, return to preparation or handoff.

## Verification And Handoff

Set verification expectation from the flow risk and changed surfaces. After work, verify against that expectation or mark what evidence is missing. A verification failure or insufficient evidence sends the flow back to the earliest safe phase rather than success reporting.

For commit-readiness, judge only whether handoff conditions are ready: intended change unit, diff scope, unrelated-change exclusion, verification evidence, and residual risk. Commit-readiness is not commit execution, staging, pushing, PR creation, publishing, release, or version bump authority.

Flow completion means the active flow met its completion criteria and produced its reporting or handoff condition. It does not mean the surrounding turn is closed, and it does not automatically execute the next flow.
