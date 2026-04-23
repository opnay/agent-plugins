---
name: turn-gate
description: Loop gate for repositories where one turn must continue until the user explicitly ends the work. Keep analysis, plan, work, result reporting, and user-response-based next-flow selection explicit inside the same turn.
---

# Turn Gate

## Overview

Use this skill when the repository or working agreement requires one user turn to continue until the user explicitly ends the work.
Its job is not to replace downstream workflow skills.
Its job is to keep the turn loop explicit:

1. analyze the user's message
2. state the plan
3. do the work
4. report the result or commit-ready state
5. open the next flow through a user response with explicit choices
6. continue unless the user explicitly ends the work

This skill is a loop gate.
It owns turn continuity and next-flow reopening, not the domain work inside each phase.

## Use When

- the repository requires one turn to stay open until the user explicitly ends the work
- the repository requires `analysis -> plan -> work -> result reporting / commit-ready -> user response` style progression
- result reporting must be followed by a next-flow choice surface instead of a soft closing
- the user should be given explicit choices for the next flow
- the main risk is premature turn termination rather than lack of a phase-specific workflow

## Do Not Use When

- the task is a normal single-phase request that can end cleanly after one answer
- the repository does not require turn continuity until explicit user stop
- the main blocker is still choosing between clarification, planning, or execution rather than managing turn continuity itself

## Scope Boundary

This skill owns:

- turn-level phase classification
- downstream workflow selection for the current phase work
- explicit analysis / plan / work / result reporting structure
- next-flow reopening after every phase result unless the user explicitly ends the work
- choice-granting user-response surface for the next flow

This skill does not own:

- requirements interviewing itself
- read-only planning itself
- implementation itself
- review-loop handling itself
- commit finalization itself

## Core Policy

- Treat each incoming message as the start or continuation of one loop-gated turn.
- Treat the user's next-flow response as the next user message inside the same turn.
- Choose the narrowest downstream workflow that owns the current phase work.
- Make `analysis`, `plan`, `work`, and `result reporting` visible in the response shape.
- Do not let result reporting become a soft stop.
- Report results as prior explanation for the user's response into the next flow, not as a terminal message.
- Reopen the next flow through a question tool that gives the user explicit choices.
- Treat termination judgment as the user's choice, not the assistant's shortcut.
- Treat "no next flow" as an exception that must be justified by explicit user stop or confirmed closure.
- Prefer the structured user-input tool for the next-flow step.
- Keep the loop moving; do not reopen broad framing once the next phase is already clear.

## Phase Loop

### Phase 0: Analyze

1. State what the user is asking for in direct terms.
2. Decide what the current phase work actually is.
3. Choose the downstream workflow that owns that work.

Output:

- `Analysis`
- `Chosen downstream owner`

### Phase 1: Plan

1. State the smallest useful plan for the current phase.
2. Include fallback or verification steps when they matter.
3. Keep the plan narrow enough to finish before reopening the next flow.

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

Output:

- `Plan`

### Phase 2: Work

1. Hand off to the selected downstream workflow.
2. Keep the current work bounded.
3. Do not replace work with meta commentary.

Output:

- `Work`
- `Phase result`

### Phase 3: Report Result Or Commit-Ready State

1. Report what changed, what was decided, or what remains blocked.
2. If the work reached a readiness boundary, report that state explicitly.
3. Treat the report as prior explanation for the user's next response.
4. Do not treat the report as the end of the turn.

Output:

- `Result report`
- `Commit-ready state` when relevant

### Phase 4: Open The Next Flow Through User Response

1. Ask what next flow the user wants to proceed with.
2. Use a question tool that grants explicit choices.
3. Offer the narrowest next-flow options that fit the current result.
4. Treat the user's response as the next user message and route it back into Phase 0 instead of ending the turn.

Output:

- `User-response question`
- `Next-flow choices`
- `Planned next-flow continuation`

## Output Contract

- `Analysis`
- `Chosen downstream owner`
- `Plan`
- `Work`
- `Phase result`
- `Result report`
- `User-response question`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Response Pattern

Preferred turn shape:

1. analyze the user's message
2. state the plan
3. describe the work
4. report the result briefly
5. ask for the user's next-flow response through explicit choices

Bad ending shape:

- summary-only closing such as "완료했습니다", "필요하면 더 말씀해주세요", or option lists without a concrete next-flow response surface
- freeform next-step prompting without giving the user explicit choices
- blocked-state closing such as "여기까지 확인했습니다" without a next-flow response surface

Good turn-flow example:

- "`workflow-kit`의 기본 시작점을 확인하는 요청으로 보고, `README.md`와 `plugin-spec.md`를 확인했습니다. 현재 저장소 기준 기본 시작점은 `workflow-kit-guide`입니다. 다음 플로우는 어떤걸 진행하시나요?"
- "1. `workflow-kit-guide` 역할 점검 2. `turn-gate` 동작 점검 3. `plugin-spec` 라우팅 규칙 확인"

## Guardrails

- Do not replace a phase-specific workflow with vague meta commentary.
- Do not emit a terminal summary unless the user explicitly ends the work.
- Do not decide on behalf of the user that the turn should terminate.
- Do not skip the user-response step merely because the next phase seems obvious.
- Do not ask the next-flow question without giving the user explicit choices.
- Do not treat the user's next-flow response as a new independent turn when the loop gate is still active.
- Do not treat temporary blocking states as permission to close the turn.
- Do not let this skill absorb domain execution, planning, or review detail.
- Do not confuse turn continuity with endless conversation.
