---
name: flow
description: Route every user message through message interview, flow design, main-flow lifecycle, main-flow review, and handoff condition; create locked execution briefs, flow configurations, records, and handoff conditions from the intent graph.
---

# Flow

Use this path for every user message:

```text
message interview -> flow design -> main flow -> main-flow review -> handoff condition
```

`flow` owns message interview, flow design, the selected main-flow lifecycle, main-flow review, handoff condition, and the meaning of `000-plan.md`, flow record, and `000-review.md` update points.
`flow` does not own question-tool execution, active-turn continuity, next-flow question routing, self-drive control, or approval-sensitive execution such as commit, push, pull request, release, version bump, or destructive action.

## Message Interview

Create a locked execution brief from the user message:

1. Capture the initial intent snapshot.
2. Identify the alignment risk.
3. Ask one high-leverage question when needed.
4. Apply the answer.
5. Pressure-test with an example, counterexample, explicit non-goal, or tradeoff.
6. If the pressure test fails, narrow the same alignment risk again.
7. Produce the locked execution brief.

The initial intent snapshot shows desired result, target, scope, and constraints.
The alignment risk is the largest uncertainty that makes the locked execution brief hard to use as execution input.
A high-leverage question narrows one alignment risk at a time.
The locked execution brief records purpose, target or targets, scope, non-goals, completion criteria, verification expectation, approval boundary, evidence basis, resolved alignment risk, and residual ambiguity as the current settled state.

If the brief is already settled enough, continue to flow design without asking the user.

## Flow Design

Convert the locked execution brief into the flow configuration:

1. Classify items as `active flow`, `parent flow`, `sub-flow candidate`, `phase`, or `handoff`.
2. Decompose flow when multiple main flows are needed.
3. Write each flow contract.
4. Order the next main flow and later candidates.
5. Select the main flow to enter.

Candidates remain pending until selected as the next main flow; pending candidates do not authorize execution.
Each flow contract includes scope, non-goals, completion criteria, verification expectation, approval boundary, and handoff condition.
If a purpose chain affects the contract, absorb it into the purpose section of `000-plan.md`.

## Main Flow

Run the selected active flow in this order:

```text
intake -> framing -> preparation -> work -> verification -> reporting
```

- `intake`: confirm the locked execution brief and main-flow input.
- `framing`: confirm the current step's frame and ownership boundary.
- `preparation`: lock the contract before work entry.
- `work`: produce the contracted output.
- `verification`: check output against the contract.
- `reporting`: record result, verification, residual risk, next intake, or handoff condition.

These names are fixed phases inside the active flow. A phase shows the active flow's current position, next action, and record update point. User-facing messages from a phase must start with the `[<phase-name>]` label pattern. Phase labels are progress markers; do not mechanically copy them into artifact bodies, record bodies, command summaries, or question option labels.

If a phase starts to own a reviewable artifact, completion criteria, approval boundary, or handoff condition, route it through flow design as a new flow or sub-flow candidate instead of treating the phase label as the flow identity.

The contracted output may be an artifact change, answer, explanation, summary, status report, verification result, or other requested result.
If another flow follows, route from `reporting` to the next `intake`.

## Main-Flow Review

After a main-flow group completes, run main-flow review before handoff condition.
Update `000-review.md`.
Record findings, or record a short no-finding result when there are none.
Do not use `000-review.md` as active routing or handoff authority.

## Handoff Condition

After main-flow review, produce the handoff condition.
The handoff condition shows result, verification, residual risk, and next intake condition.

If blockers or insufficient evidence remain, expose them in the handoff condition.
Commit-readiness can be a handoff judgment, but commit execution requires separate authority.
Commit, push, pull request, release, version bump, and destructive action are not executed from handoff without separate approval.

## Records

Use records to keep routing recoverable.
Records are not execution authority.

- `000-plan.md`: may be updated from message interview and flow design.
- Flow record: may be updated during main-flow phases and handoff condition.
- `000-review.md`: update after the main-flow group and before handoff condition.

When creating record files, use the bundled templates:

- `templates/plan.md`
- `templates/flow-record.md`
- `templates/review.md`
