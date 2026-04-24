---
name: turn-gate
description: Loop gate for repositories where one turn must continue until the user asks to end the turn. Keep analysis, plan, work, verification, result reporting, and question-routing-based next-flow selection explicit inside the same turn.
---

# Turn Gate

## Overview

Use this skill when the repository or working agreement requires one user turn to continue until the user asks to end the turn.
Using this skill means treating `turn-gate` as a first-class rule for the rest of the current session.
Its job is not to replace downstream workflow skills.
Its job is to keep the turn loop explicit:

1. analyze the user's message
2. state the plan
3. do the work
4. verify the work
5. report the result or commit-ready state
6. open the next flow through a question-routing response with explicit choices
7. continue unless the user asks to end the turn

This skill is a loop gate.
It owns turn continuity and next-flow reopening, not the domain work inside each phase.
It also owns the question-routing axis: `user-gated` by default, or `self-drive` when questions should go to subagents so the loop can continue without user intervention.

## Use When

- the repository requires one turn to stay open until the user asks to end the turn
- the repository requires `analysis -> plan -> work -> result reporting / commit-ready -> user response` style progression
- result reporting must be followed by a next-flow choice surface instead of a soft closing
- the user should be given explicit choices for the next flow
- the main risk is premature turn termination rather than lack of a phase-specific workflow

## Do Not Use When

- the task is a normal single-phase request that can end cleanly after one answer
- the repository does not require turn continuity until the user asks to end the turn
- the main blocker is still choosing between clarification, planning, or execution rather than managing turn continuity itself

## Scope Boundary

This skill owns:

- turn-level phase classification
- downstream workflow selection for the current phase work
- explicit analysis / plan / work / verification / result reporting structure
- next-flow reopening after every phase result unless the user asks to end the turn
- choice-granting question-routing surface for the next flow

This skill does not own:

- requirements interviewing itself
- read-only planning itself
- implementation itself
- review-loop handling itself
- commit finalization itself

## Core Policy

- Treat invocation of this skill as activation of a session-level first-class loop gate.
- Treat each incoming message as the start or continuation of one loop-gated turn.
- Treat the user's next-flow response or the `self-drive` subagent answer as the next message inside the same turn.
- When a user message arrives while `self-drive` is active, treat it as authoritative loop input rather than a reason to stop.
- Classify mid-self-drive user input as explicit turn stop, current-flow correction, current-flow priority change, or next-flow priority request.
- Adjust the current flow immediately when the message changes active work; otherwise register the message as the highest-priority next-flow candidate.
- Supersede any pending or returned self-drive subagent answer that conflicts with newer user input.
- Choose the narrowest downstream workflow that owns the current phase work.
- Make `analysis`, `plan`, `work`, `verification`, and `result reporting` visible in the response shape.
- Maintain running turn-gate records under `.agents/sessions/{YYYYMMDD}/`.
- Maintain a compact `Continuity Guard` in every flow record and refresh it before result reporting and next-flow reopening.
- The `Continuity Guard` must state whether `turn-gate` is active, the question-routing mode, whether the user explicitly stopped the turn, whether a terminal summary is allowed, and the required next action.
- Use `000-plan.md` for the higher-level multi-flow plan and `001+` files for per-flow records.
- Keep `000-plan.md` incrementally updated beyond one user request when the larger task continues.
- Use `analysis` to structure the user's message into requested intent and requested action, and to decide whether future flows or phases need forward design or redesign.
- Use `plan` to prepare the detailed next steps needed to fulfill the analyzed request and, when useful, a provisional design for later flows or phases.
- Use `work` to execute the prepared plan.
- Use `verification` to confirm the work outcome before result reporting and to surface whether later flow/phase redesign is needed.
- Use `result reporting` to report the completed work outcome.
- Do not let result reporting become a soft stop.
- Before result reporting, read or reconstruct the `Continuity Guard`; if the user has not explicitly stopped the turn, a terminal summary is invalid.
- Report results as prior explanation for the user's response into the next flow, not as a terminal message.
- Reopen the next flow through the active question-routing mode with explicit choices.
- In `self-drive`, replace phase questions that would normally go to the user with self-drive question packets sent to subagents, require the self-drive answer contract, and continue from the subagent answer.
- Allow questions during `analysis` and `plan` when clarifying intent, criteria, or scope is necessary.
- Treat termination judgment as the user's choice, not the assistant's shortcut.
- Treat "no next flow" as an exception that must be justified by the user asking to end the turn or by confirmed closure.
- When later loops return to `analysis` or `plan`, revise future flow/phase design only when new evidence, changed intent, or a revealed blocker makes redesign necessary.
- Prefer the structured user-input tool for the next-flow step unless `self-drive` is active.
- Keep the loop moving; do not reopen broad framing once the next phase is already clear.

## Session Record

- Use `.agents/sessions/{YYYYMMDD}/000-plan.md` for the higher-level plan when the task spans several flows.
- Use `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for flow records.
- Keep `count-pad3` zero-padded like `001`, `002`, `003`.
- Keep the slug English lower-case and `-` delimited.
- Record at least: user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, residual risk.
- Keep the current flow record current after `analysis`, `plan`, `work`, `verification`, and `result reporting`.
- Update the flow record incrementally after each completed phase.
- Prefer `templates/flow-record-template.md` as the default flow-record layout.
- Prefer `templates/plan-template.md` as the default `000-plan.md` layout.

## Phase Loop

### Phase 0: Analyze

1. Structure the user's message into requested intent and requested action.
2. State what is already clear and what still needs clarification.
3. Decide what the current phase work actually is.
4. Ask through the active question-routing mode when clarification is necessary before safe planning.
5. Choose the downstream workflow that owns that work.

Output:

- `Analysis`
- `Requested intent`
- `Requested action`
- `Chosen downstream owner`

### Phase 1: Plan

1. Prepare the detailed plan needed to fulfill the analyzed request.
2. Ask through the active question-routing mode when planning is blocked by missing criteria, scope, or approval.
3. Include fallback or verification steps when they matter.
4. Keep the plan narrow enough to finish before reopening the next flow.

Typical downstream owners:

- `structured-thinking`
- `deep-interview`
- `planner`
- `autopilot`
- `parallel-work`
- `ralph-loop`
- `review-loop`
- `commit-readiness-gate`
- a specialist plugin after the workflow phase is clear

Selection signals:

- `deep-interview` when the blocker is real requirement discovery
- `review-loop` when the input is review findings and only material issues should be fixed
- `ralph-loop` when one bounded fix-verify-reassess cycle is the best next step
- `autopilot` when the current phase is broad end-to-end delivery from a brief request through implementation, QA, and validation
- `commit-readiness-gate` when implementation is largely done and the current question is readiness for commit

Question-routing signals:

- `user-gated` by default, using the user-input question tool for choices, scope locks, and next-flow decisions
- `self-drive` when the user wants questions answered by subagents so work can continue without user intervention
- `self-drive` can answer mode selection, criteria, scope assumptions, verification choices, and next-flow decisions through self-drive question packets sent to subagents
- self-drive packets must carry the current `Continuity Guard`
- self-drive subagent answers must include the chosen option or no-option result, decision, rationale, evidence, assumptions, confidence, blockers, approval boundary, continuity check, and next action
- self-drive answers must reject terminal summaries unless an explicit user stop or hard approval boundary exists
- real user messages outrank pending or returned self-drive subagent answers
- mid-self-drive user intervention should adjust the current flow or become the highest-priority next flow, not stop the turn
- self-drive should recover subagent `context_gap` results through main-agent discovery when evidence can be found without an explicit approval boundary
- missing user preference should become a recorded reversible assumption unless the user explicitly requested manual preference locking
- `low` confidence means an approval-boundary pause only when the decision requires explicit approval, destructive/irreversible/external action approval, or a platform/tool/safety boundary
- in `self-drive`, an approval-boundary pause stops autonomous routing only; switch to `user-gated` and use `request_user_input` instead of ending the turn
- `self-drive` must still pause at explicit user-approval boundaries required by platform, tool, or safety policy

Output:

- `Plan`

### Phase 2: Work

1. Execute the prepared plan through the selected downstream workflow.
2. Keep the current work bounded.
3. Do not replace work with meta commentary.

Output:

- `Work`
- `Phase result`

### Phase 3: Verification

1. Run the narrowest meaningful verification for the work just performed.
2. State what was verified, what passed, and what remains uncertain.
3. Treat missing verification as an explicit residual risk, not an implicit omission.

Output:

- `Verification`

### Phase 4: Report Result Or Commit-Ready State

1. Report what changed, what was decided, or what remains blocked.
2. If the work reached a readiness boundary, report that state explicitly.
3. Treat the report as prior explanation for the user's next response.
4. Do not treat the report as the end of the turn.

Output:

- `Result report`
- `Commit-ready state` when relevant

### Phase 5: Open The Next Flow Through Question Routing

1. Ask what next flow should proceed.
2. Use the active question-routing mode with explicit choices.
3. Offer the narrowest next-flow options that fit the current result.
4. In `user-gated`, treat the user's response as the next user message and route it back into Phase 0 instead of ending the turn.
5. In `self-drive`, ask a subagent to choose the next flow from the explicit choices, record its answer, and route that answer back into Phase 0 instead of ending the turn.
6. If the user sends a message while self-drive is in progress, route it back into Phase 0 immediately as either a current-flow adjustment or a priority next-flow registration.

Output:

- `Question-routing prompt`
- `Next-flow choices`
- `Planned next-flow continuation`

## Output Contract

- `Analysis`
- `Requested intent`
- `Requested action`
- `Chosen downstream owner`
- `Question-routing mode`
- `Plan`
- `Work`
- `Verification`
- `Phase result`
- `Result report`
- `Continuity guard`
- `Question-routing prompt`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Response Pattern

Preferred turn shape:

1. analyze the user's message
2. state the plan
3. describe the work
4. state the verification briefly
5. report the result briefly
6. ask for the next-flow response through the active question-routing mode

Bad ending shape:

- summary-only closing such as "완료했습니다", "필요하면 더 말씀해주세요", or option lists without a concrete next-flow response surface
- freeform next-step prompting without giving the user explicit choices
- blocked-state closing such as "여기까지 확인했습니다" without a next-flow response surface

Good turn-flow example:

- "`workflow-kit-dev`의 기본 시작점을 찾아달라는 요청이지만, 시작점을 어떤 기준으로 볼지 먼저 맞춰야 한다고 판단했습니다. 지금은 기준 선택이 먼저 필요합니다. 시작점을 어떤 기준으로 볼까요?"
- "1. 플러그인 관점 2. `AGENTS.md` 관점 3. 스킬 관점"

## Guardrails

- Do not replace a phase-specific workflow with vague meta commentary.
- Do not emit a terminal summary unless the user asks to end the turn.
- Do not decide on behalf of the user that the turn should terminate.
- Do not skip the `000-plan.md` update when the higher-level plan changes across flows.
- Do not defer the flow-record update at `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` until the end of the flow; write it at each completed phase boundary.
- Do not skip explicit verification between work and result reporting.
- Do not emit result reporting until the `Continuity Guard` says whether next-flow reopening is still required.
- Do not skip the next-flow response step merely because the next phase seems obvious.
- Do not ask the next-flow question without giving the user explicit choices.
- Do not use `self-drive` unless that question-routing mode is active.
- Do not let `self-drive` simulate user approval where the runtime or tool policy requires explicit approval.
- Do not treat mid-self-drive user intervention as completion, stop, or an approval-boundary pause unless the user explicitly asks to end the turn or creates a real approval boundary.
- Do not continue from a stale self-drive subagent answer after newer user input changes the active flow.
- Do not end the turn when `self-drive` reaches an approval boundary; pause self-drive, switch to `user-gated`, and ask through `request_user_input`.
- Do not treat the user's next-flow response as a new independent turn when the loop gate is still active.
- Do not treat temporary blocking states as permission to close the turn.
- Do not let this skill absorb domain execution, planning, or review detail.
- Do not confuse turn continuity with endless conversation.
