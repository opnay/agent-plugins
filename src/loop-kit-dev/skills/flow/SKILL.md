---
name: flow
description: Route every user message through message interview, flow design, main-flow lifecycle, main-flow review, and handoff condition; classify active flows, parent flows, candidates, phases, and handoffs; preserve flow records and handoff boundaries.
---

# Flow

Use this shape for every user message:

```text
message interview -> flow design -> main flow -> main-flow review -> handoff condition
```

Questions, status checks, explanations, short answers, and work requests all enter the same shape.
Do not create an inline answer path outside the graph.
If a user question would not change the flow contract, continue without asking the user, but still produce the interviewed brief and flow design.

`flow` owns message interpretation, flow design, selected main-flow lifecycle, main-flow review, handoff condition, and the meaning of `000-plan.md`, flow record, and `000-review.md` update points.
`flow` does not own question-tool execution, active-turn continuity, next-flow question routing, self-drive control, or approval-sensitive execution such as commit, push, pull request, release, version bump, or destructive action.

## Message Interview

Run message interview for every user message:

1. Capture the initial intent snapshot.
2. Identify the highest alignment risk.
3. Ask one high-leverage question only when the answer can change the flow contract.
4. Apply the answer to the same risk.
5. Pressure-test the brief with an example, counterexample, explicit non-goal, or tradeoff.
6. If the pressure test fails, narrow the same alignment risk again.
7. Produce the locked execution brief.

If no user question is needed, record that the brief is locked enough and move to flow design.
The no-question path is still an interview path.

Interview output includes the snapshot, risk, question or no-question decision, answer effect, pressure-test result, locked execution brief, and whether `000-plan.md` should be updated.

## Flow Design

Convert the locked execution brief into the flow configuration:

1. Classify items as `active flow`, `parent flow`, `sub-flow candidate`, `phase`, or `handoff`.
2. Decide whether one main flow is enough or multiple main flows are needed.
3. Write the contract for each flow.
4. Order the next main flow and later candidates.
5. Select the main flow to enter.

Questions, status checks, and explanation requests can be selected as active flows.
Candidates remain pending until selected.
Phases stay inside the active flow.
Handoffs are post-main-flow conditions and do not create execution authority.

Each flow contract includes identity, scope, non-goals, completion criteria, verification expectation, approval boundary, and handoff condition.
If a purpose chain affects the contract, absorb it into the purpose section of `000-plan.md`.

## Main Flow

Run each selected active flow in this order:

```text
intake -> framing -> preparation -> work -> verification -> reporting
```

- `intake`: confirm the locked execution brief and current active-flow input.
- `framing`: confirm classification, ownership, selected-vs-candidate state, phase, and handoff boundaries.
- `preparation`: lock scope, non-goals, completion criteria, verification expectation, approval boundary, and handoff condition before work.
- `work`: act inside the active-flow contract or produce the answer, explanation, summary, comparison, or status result.
- `verification`: verify against the locked execution brief and active-flow contract, or record insufficient evidence.
- `reporting`: report result, verification, residual risk, and the next intake or handoff condition.

If another flow follows, route from `reporting` to the next `intake`.
If target, scope, operation, verification expectation, approval boundary, or acceptance changes, return to the earliest safe interview, design, or preparation point.
If an artifact change becomes an independent reviewable unit, route it back through flow design as a candidate or selected flow.

## Records

Use records to make routing recoverable; do not treat records as approval.

- `000-plan.md`: update from message interview and flow design with active flow, candidates, order, purpose chain, and next action.
- Flow record: update at main-flow phases with input, classification, evidence, verification, reporting, and handoff condition.
- `000-review.md`: update after the main-flow group and before handoff condition.

When creating record files, use the bundled templates:

- `templates/plan.md`
- `templates/flow-record.md`
- `templates/review.md`

Records, generated surfaces, verification notes, and previous context do not authorize commit, push, pull request, release, version bump, destructive action, or other approval-sensitive execution.

## Main-Flow Review

After a main-flow group completes, run main-flow review before handoff condition.
Update `000-review.md`.
Record findings, or record a short no-finding result when there are none.

Main-flow review is not active routing.
It is not handoff authority, raw log storage, verification authority, commit authority, release authority, or closure authority.

## Handoff Condition

After main-flow review, produce the handoff condition.

The handoff condition includes:

- completion state
- verification state
- residual risk
- next intake condition
- blocker or insufficient evidence
- approval-sensitive action status

Commit-readiness can be a handoff judgment, but commit execution requires separate authority.
Push, pull request, release, version bump, and destructive action also require separate authority.

## Contract Impact

When a new user message arrives during an active flow, route it through message interview and flow design first.
Then return the contract-impact result as one of the flow outputs:

- revise the current flow
- start a new foreground active flow
- keep a future candidate
- supersede the current flow
- answer a blocker question
- select a new active flow
- produce an explicit-stop handoff signal when the stop source is clear

Do not answer inline outside the flow graph.
