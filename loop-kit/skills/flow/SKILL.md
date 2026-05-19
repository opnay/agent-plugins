---
name: flow
description: Use when a request, action, plan item, review finding, or handoff must be interpreted as a cohesive flow, parent flow, finite sub-flow candidate, readiness/discovery/ambiguity decision, flow-local strategy, or handoff condition before work proceeds.
---

# flow

## Purpose

Use this skill to decide what the current work flow is, whether it is executable now, and whether it should be split into finite `sub-flow candidates`.

A flow is not a phase checklist. A flow is a cohesive unit that can be understood, reviewed, verified, and, when relevant, committed together. Each active flow has its own `preparation -> work -> verification -> reporting` lifecycle. Finishing a flow does not end the turn; turn continuation and next-flow questions belong to `turn-gate`.

## Ownership

This skill owns:

- flow boundary and flow-vs-phase judgment
- parent flow and finite `sub-flow candidate` design
- `operational-preparation flow` and `change-unit flow` distinction
- active flow versus follow-up candidate distinction
- readiness, discovery, and operation/target ambiguity decisions before work
- flow-local strategy selection: review-loop, fix-verify-loop, broad-execution
- flow completion criteria, verification expectation, and handoff condition
- commit-readiness as a readiness/handoff condition, not commit execution

This skill does not own:

- turn activation, explicit stop, or next-flow question routing
- session continuity or terminal closure guard
- user question tool mechanics
- commit, push, PR, publish, release, or version-bump execution details
- self-drive sequence-level continuation

## Flow Model

Classify the current item before treating it as executable work.

- `flow`: a cohesive work stream with preparation, work, verification, and reporting.
- `parent flow`: a flow whose output is a finite list of sub-flow candidates and their contracts.
- `sub-flow candidate`: a possible future flow. Creating a candidate is not execution.
- `active flow`: the flow selected for current work.
- `operational-preparation flow`: a flow that locks intent, scope, non-goals, success signals, verification expectations, approval boundaries, and sub-flow candidates or selected sequence.
- `change-unit flow`: a flow that owns concrete reviewable changes such as code, docs, fixtures, configuration, or release surface.

Do not treat `analysis`, `work`, `verification`, `reporting`, or `commit readiness` as flows by name alone. Pure QA, final consistency checks, verification reporting, and commit-readiness reporting are not separate change-unit flows unless they own a distinct reviewable artifact change.

## Required Flow Output

When designing a flow or candidate list, make these fields explicit enough for a session record or handoff:

- flow label or slug
- flow type: `operational-preparation` or `change-unit`
- scope and non-goals
- completion criteria
- verification expectation
- readiness status or missing contract fields
- recommended question topics or unresolved ambiguity
- recommended flow-local strategy
- approval-sensitive checkpoint, if expected
- handoff condition
- unresolved questions or blocker
- whether this is the active flow or only a sub-flow candidate

Keep completion criteria separate from handoff condition. A flow can be complete while still needing a next-flow choice or a commit-readiness handoff.

## Preparation Decisions

Before work, decide whether the flow contract is ready.

Work can start only when scope, non-goals, completion criteria, verification expectation, and handoff condition are sufficient. If any of these are missing, output missing fields and recommended question topics instead of executing.

Use discovery when user intent, included scope, non-goals, success criteria, verification expectation, output shape, decomposition, or approval-sensitive checkpoints could change based on the answer.

Use ambiguity resolution when operation or target wording could change the flow contract. Examples include `merge`, `absorb`, `move`, `promote`, `remove`, `delete`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, `contract`, or pronouns such as "that", "above", "below", and "current" when multiple targets are possible.

Approval-sensitive execution is separate from ambiguity resolution. This skill may identify that approval is needed, but the approval boundary and execution authority are owned outside `flow`.

## Flow-Local Strategy

Choose a strategy inside the active flow only after the flow boundary is locked.

- Use `review-loop` only for one bounded blocking review, QA, or self-review finding inside the current active flow. A finding is blocking when it directly affects correctness, regression risk, reliability, or delivery risk for the active flow.
- Use `fix-verify-loop` for one narrow problem where a small fix or check can test the current hypothesis. Reassess after each loop.
- Use `broad-execution` for a single locked active flow whose implementation, QA, and validation all remain inside the same boundary.

These strategies do not authorize multiple flows to run automatically. Sequence-level continuation belongs to a prepared self-drive sequence, not to `flow`.

### Review-Loop Contract

Do not use `review-loop` as a broad execution plan for all review comments, QA findings, or self-review notes. It is a flow-local strategy for handling one material finding that is already bounded, blocking, and inside the active flow.

When there are multiple findings:

- choose the highest-priority bounded blocking finding that belongs to the active flow, then run one `review-loop` for that finding only
- use discovery if the active flow scope, success criteria, or verification expectation must be clarified before choosing a finding
- return to the parent flow when the findings imply a larger decomposition decision
- create finite follow-up candidates for non-blocking findings, speculative polish, or findings outside the active flow

After handling the selected finding, verify the expectation directly tied to that finding. If the finding requires new scope, a new approval boundary, destructive action, or external action, stop expanding execution and route back to preparation or handoff instead.

## Handoff Conditions

Use handoff when the current flow reaches a boundary that another controller or workflow must handle.

Commit-readiness is a handoff condition. It checks intended change unit, diff scope, unrelated-change exclusion, verification evidence, and residual risk. It does not stage, commit, push, open PRs, publish, release, or bump versions.

If verification evidence is missing, do not report readiness as successful. Route the active flow back to the earliest safe preparation, work, or verification point.

## Relationship To turn-gate

`turn-gate` applies this skill inside an active turn. It prevents work without a source-recorded active flow or a `flow` decision, records this skill's output contract in session records, and opens next-flow routing after reporting.

`turn-gate` may use this skill's missing fields and question topics for user-gated routing, but it must not redefine discovery, readiness, flow-local strategy, or handoff conditions. This skill never opens the next flow by itself.
