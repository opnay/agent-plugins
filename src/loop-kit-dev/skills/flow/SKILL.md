---
name: flow
description: Interpret a user message or action as an active flow, parent flow, finite sub-flow candidates, operational-preparation flow, or change-unit flow; use when deciding flow boundaries, flow-vs-phase, completion criteria, verification expectation, and handoff shape before turn-gate applies the decision.
---

# Flow

Use this skill to decide what the current work unit is. A flow is not a phase checklist. It is a cohesive unit that can be understood, reviewed, verified, and, when relevant, committed together.

## Ownership

`flow` owns:

- flow boundary and flow-vs-phase judgment
- parent flow and finite `sub-flow candidates`
- active flow versus candidate status
- `operational-preparation` versus `change-unit` type
- flow-local `preparation -> work -> verification -> reporting`
- completion criteria, verification expectation, and handoff condition

`flow` does not own turn activation, explicit stop handling, next-flow question routing, session continuity, self-drive continuation, commit execution, push, PR, or publish rules.

## Core Model

Treat each flow as one bounded work stream with these internal phases:

1. `preparation`
2. `work`
3. `verification`
4. `reporting`

If the request is too broad or creates multiple reviewable outputs, model it as a parent flow that produces finite `sub-flow candidates`. Candidate creation is not execution. A candidate becomes the active flow only after a turn controller selects it, or after a prepared self-drive sequence advances to it.

An active flow is the unit currently being prepared, worked, verified, or reported. A sub-flow candidate is a possible later unit with enough detail to choose, hand off, or defer.

## Flow Types

Use one of these types unless the runtime contract has been explicitly extended:

- `operational-preparation`: locks intent, scope, non-goals, success signals, verification expectation, approval boundaries, and candidate handoff. Its output is a plan/session artifact or bounded candidate set, not product work.
- `change-unit`: changes or creates reviewable artifacts such as code, docs, fixtures, configuration, templates, or release surfaces.

An `operational-preparation` flow may produce `change-unit` candidates. Those candidates remain candidates until selected.

## Boundary Checks

Do not treat these labels as flows by themselves:

- `analysis`
- `work`
- `verification`
- `reporting`
- `commit readiness`
- final QA, consistency checks, or result reporting with no separate artifact change

An item can be a flow when it owns a reviewable artifact or a bounded operational artifact, such as a session plan, candidate handoff, fixture, snapshot baseline, diagnostic output, or validator report.

Use these checks:

- Does this unit have a coherent scope and non-goals?
- Can it be reviewed and verified on its own?
- Would it make sense as a commit-sized or handoff-sized unit?
- Is it merely a phase inside another flow?
- If it produces candidates, are those finite and not being executed yet?

## Output Contract

When you return or record a flow decision, include:

- `flow_label` or `slug`
- `flow_type`: `operational-preparation` or `change-unit`
- `status`: `active_flow` or `sub_flow_candidate`
- `scope`
- `non_goals`
- `completion_criteria`
- `verification_expectation`
- `approval_sensitive_checkpoint`
- `handoff_condition`
- `unresolved_questions_or_blockers`

For a parent flow, include the parent label and a finite `sub-flow candidates` section. For each candidate, include the same contract fields and make clear that it is not active execution.

## Relationship To Turn Gate

`turn-gate` applies this decision inside an active turn. It records the flow fields, prevents work without a source-recorded active flow or new flow decision, and reopens next-flow routing after reporting.

Do not use flow completion as turn completion. A completed flow can still require next-flow selection, blocker routing, commit-readiness handoff, or explicit turn stop handling.
