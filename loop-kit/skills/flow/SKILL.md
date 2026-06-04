---
name: flow
description: Interpret every user message through message interview, flow design, main-flow lifecycle, main-flow review, and handoff condition; route questions, status checks, explanations, work requests, phases, candidates, and handoffs through the same flow graph; lock scope, non-goals, verification, approval boundary, next intake condition, and candidate execution boundaries.
---

# Flow

Shape: `message interview -> flow design -> main flow -> main-flow review -> handoff condition`.

Every user message enters this shape.
If a high-leverage question would not change the contract, continue without asking the user.
Still produce an alignment snapshot, risk check, pressure test, locked brief, flow design, selected main flow, review, and handoff condition.

A flow is a reviewable work unit with scope, non-goals, completion criteria, verification expectation, approval boundary, and handoff condition.
Questions, status checks, explanations, and short answers are selected `active flow` items.
Their work output may be a deeper answer, explanation, summary, comparison, or status result.
Phase checklists, QA, reports, repair, blockers, main-flow review, and commit-readiness stay inside their owning surface.
Record template meaning belongs to flow; turn-gate applies and updates records during an active turn.

## Message Interview

Run the internal deep-interview loop for every user message:

1. Capture `initial alignment snapshot`.
2. Pick the highest alignment risk.
3. Ask one high-leverage question only when the answer can change the flow contract.
4. If no question is needed, record why the contract is already locked enough.
5. Prefer bounded choices with tradeoffs when they lock the contract.
6. Pressure-test the brief with an example, counterexample, explicit non-goal, or rejected tradeoff.
7. Stay on the same risk while it remains vague.
8. Produce `locked execution brief`.

Interview risks include intent, target outcome, scope edge, non-goal, tradeoff, decision boundary, constraint, acceptance, and approval boundary.
The no-question path is not a skip; it is an interviewed message with a locked brief and residual ambiguity recorded.

## Flow Design

Classify the locked brief and create a list-up result:

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

- `intake`: confirm input, interpretation boundary, missing fields, and locked brief source.
- `framing`: confirm classification, ownership, candidate-vs-selected state, and active-flow output.
- `preparation`: lock scope, non-goals, completion, verification, approval, and handoff.
- `work`: act inside the active-flow boundary.
- `verification`: verify or record missing evidence against the locked brief and scope.
- `reporting`: report result, verification, residual risk, and handoff.

For user-facing phase-start or meaningful progress messages, produce the current phase label: `[intake]`, `[framing]`, `[preparation]`, `[work]`, `[verification]`, or `[reporting]`.
Do not mechanically copy phase labels into artifact bodies, records, command summaries, or question option labels.

If target, operation, scope, verification path, approval boundary, or acceptance changes, return to the earliest safe message-interview, flow-design, or preparation point.

## Main-Flow Review

After the main flow group completes, run main-flow review and update `000-review.md`.
Record a compact tagged review result for the completed main flow, including a no-finding result when no retrospective finding exists.
Do this before the handoff condition.
The review record is not active routing, raw flow log, verification authority, commit/release authority, or closure authority.

## Strategy And Handoff

Inside one selected active flow, choose:

- `review-loop`: one bounded blocking review/QA/self-review finding
- `fix-verify-loop`: smallest useful fix or check, then immediate verification
- `broad-execution`: one locked active flow end to end

Commit, push, PR, publish, release, version bump, and destructive work require separate approval authority.

For non-pass evidence, route to verification, reconciliation, preparation/design relock, or blocked handoff.
Reporting may produce the next main-flow intake condition.
After main-flow review, completion produces the handoff condition.
Commit-readiness is a handoff judgment over intended change unit, diff scope, unrelated changes, verification evidence, and residual risk. Execution needs separate authority.

## Contract Impact

For a new message during an active flow, first route the message through message interview and flow design.
Then decide whether it revises the current flow, starts a new foreground flow, becomes a future candidate, supersedes the current flow, answers a blocker question, or selects a new active flow.
Return the contract-impact decision as flow output.
Do not answer inline outside the flow graph.
