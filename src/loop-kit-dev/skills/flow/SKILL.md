---
name: flow
description: Interpret a user message through message interview, flow design, main-flow lifecycle, and handoff; decide active flow vs parent flow vs sub-flow candidate vs phase vs handoff; lock scope, non-goals, verification, approval boundary, next intake condition, and candidate execution boundaries.
---

# Flow

Shape: `message interview -> flow design -> main flow -> handoff condition`.

A flow is a reviewable work unit with scope, non-goals, completion criteria, verification expectation, approval boundary, and handoff condition. Phase checklists, QA, reports, repair, blockers, and commit-readiness stay inside their owning phase or handoff surface.
Record template meaning belongs to flow; turn-gate applies and updates records during an active turn.

## Message Interview

Use the internal deep-interview loop when intent, scope, tradeoff, acceptance, approval boundary, or decision boundary needs alignment:

1. Capture `initial alignment snapshot`.
2. Pick the highest alignment risk.
3. Ask one high-leverage question when the answer changes the flow contract.
4. Prefer bounded choices with tradeoffs when they lock the contract.
5. Pressure-test the answer with example, counterexample, explicit non-goal, or rejected tradeoff.
6. Stay on the same risk while it remains vague.
7. Produce `locked execution brief`.

Clear low-risk messages can collapse to snapshot plus brief.

## Flow Design

Classify the current item and create a list-up result:

- `active flow`: selected main flow
- `parent flow`: flow configuration or candidate producer
- `sub-flow candidate`: pending option, not an active flow
- `phase`: internal active-flow step
- `handoff`: readiness, blocker, next intake condition, or routing result

For selected flows and candidates, lock identity, type, scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition, and unresolved blocker.

`operational-preparation` creates briefs, flow configuration, candidates, or next-main-flow contracts.
`change-unit` owns reviewable artifact changes.
Candidates stay pending until selected as a main flow.
`000-plan.md` and active flow records are record surfaces attached to the current phase.
Purpose chains live in the `000-plan.md` purpose section when they affect scope, acceptance, verification, approval, or handoff.

## Record Templates

When defining or checking record surfaces:

- Use exact bundled templates from `templates/plan.md`, `templates/flow-record.md`, and `templates/review.md` when creating those record types.
- `plan`: keep a compact flow routing card with active flow, next action, handoff condition, approval boundary, verification expectation, active skills, current request, purpose, flow index, and continuity note.
- `flow record`: keep one reviewable work unit with contract, phase checklist, execution log, result, metadata, optional pending question, and approval-sensitive action section only when needed.
- `000-review.md`: keep retrospective notes as a flat tagged list; do not use it for active routing, raw logs, verification authority, commit/release authority, or closure authority.
- Do not let readiness, verification, build output, previous context, or generated surfaces authorize commit, push, PR, publish, release, version bump, destructive, or external actions.

## Main Flow

Each selected main flow runs:

`intake -> framing -> preparation -> work -> verification -> reporting`

- `intake`: confirm input, interpretation boundary, missing fields.
- `framing`: confirm classification, ownership, candidate-vs-selected state.
- `preparation`: lock scope, non-goals, completion, verification, approval, handoff.
- `work`: act inside the active-flow boundary.
- `verification`: verify or record missing evidence.
- `reporting`: report result, verification, residual risk, handoff.

If target, operation, scope, verification path, approval boundary, or acceptance changes, return to the earliest safe message-interview, flow-design, or preparation point.

## Strategy And Handoff

Inside one selected active flow, choose:

- `review-loop`: one bounded blocking review/QA/self-review finding
- `fix-verify-loop`: smallest useful fix or check, then immediate verification
- `broad-execution`: one locked active flow end to end

Commit, push, PR, publish, release, version bump, and destructive work require separate approval authority.

For non-pass evidence, route to verification, reconciliation, preparation/design relock, or blocked handoff.
Reporting may produce the next main-flow intake condition.
Commit-readiness is a handoff judgment over intended change unit, diff scope, unrelated changes, verification evidence, and residual risk. Execution needs separate authority.

## Contract Impact

For a new message during an active flow, decide whether it is inline answer, current-flow revision, new foreground flow, future candidate, supersede, or blocker question.
Return the contract-impact decision as flow output.
