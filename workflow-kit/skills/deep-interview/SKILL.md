---
name: deep-interview
description: Use when the main job is to understand the user's real intent through questions, boundary checks, tradeoff probing, and direction evaluation before locking requirements or moving into planning or execution.
---

# Deep Interview

## Overview

Use this skill when implementation quality depends on understanding what the user really means, not just what they first said.
Its job is to surface intent, hidden constraints, non-goals, tradeoffs, decision boundaries, and success criteria before requirements harden around the wrong assumptions.
Use it for requirement discovery, scope locking, direction evaluation, and any question-led pass where the real deliverable is a correct understanding of what should happen.
This skill does not replace execution. It produces a sharper brief, clearer judgment, or a better-defined direction for whatever comes next.

## Use When

- the user wants something implemented, but the request is still underspecified
- the user has a mental picture of the result and wants the implementation to align closely with it
- the user explicitly asks for a deep interview, alignment pass, or "don't assume" style clarification
- the task definition is still too loose for confident planning or execution
- the user wants questions first so the real requirement becomes explicit
- the user wants a direction, fix, or proposal evaluated before turning it into a concrete plan
- the user asks for feedback on a direction, fix, or approach, but the real risk is hidden assumptions and premature solutioning
- the request sounds lightweight on the surface, yet the likely next step is real work and major scope or success decisions remain unlocked
- even a small task such as a bug fix, behavior tweak, or typo-level change still needs intent, scope, or acceptance boundaries clarified first
- a premature implementation pass would likely cause misalignment, churn, or rework
- a tentative implementation plan exists, but its assumptions still need to be pressure-tested against user intent

## Do Not Use When

- the user already gave clear targets, boundaries, and acceptance criteria
- the user explicitly wants fast execution with reasonable assumptions
- the user is only asking for lightweight opinion, inspiration, or taste-level feedback with no implied implementation commitment
- the user explicitly wants rough idea generation and is comfortable with broad assumptions
- the task is only about choosing a workflow mode
- the task is already implemented and only needs final verification or commit-readiness judgment

## Scope Boundary

Use this skill to make the execution brief operationally explicit before implementation.
Use it also to make the requirement itself explicit before it becomes a plan or build request.
It owns clarification of:

- why the user wants the change
- what outcome they actually care about
- what requirement or direction they are really endorsing
- where the scope should stop
- what should explicitly stay out
- which tradeoffs are acceptable or unacceptable
- what Codex may decide autonomously
- what still requires confirmation

This skill does not own:

- full implementation
- broad architecture planning
- final commit-readiness review

## Core Policy

- Ask one high-leverage question at a time.
- Prioritize intent, boundaries, and tradeoffs over solution detail.
- Do not ask the user for repository facts that can be discovered directly.
- Prefer the structured user-input tool when the question fits 1-3 short bounded choices.
- Use normal conversational questioning only when the missing input cannot be captured cleanly in a short bounded choice.
- Treat each answer as incomplete until it is pressure-tested by example, counterexample, tradeoff, or explicit exclusion.
- Stay on the same thread when the answer is still vague instead of rotating topics too early.
- Prefer alignment over coverage; the goal is not to ask many questions, but to reduce implementation risk caused by misunderstanding.

## Alignment Dimensions

Clarify these dimensions before handoff:

- `Intent`
  - why the user wants this change
- `Target outcome`
  - what "done correctly" looks like to them
- `Scope edge`
  - how far the change should go
- `Non-goals`
  - what should explicitly stay out
- `Tradeoffs`
  - what is acceptable to sacrifice and what is not
- `Decision boundaries`
  - what Codex may decide without asking again
- `Constraints`
  - technical, product, process, or style limits
- `Acceptance signals`
  - what would make the user say the implementation matches their intent

## Workflow

### Phase 0: Context Intake

1. Read the request and any existing implementation plan or draft.
2. Inspect local codebase facts when relevant.
3. Form a provisional picture of:
   - stated request
   - likely intent
   - obvious assumptions
   - unresolved ambiguity

Output:

- `Initial alignment snapshot`

### Phase 1: Intent Interview

1. Ask the single question that will most reduce alignment risk.
2. If that question can be expressed as 1-3 short bounded choices, use the structured user-input tool.
3. Otherwise ask it directly in conversation.
4. After each answer, decide whether the real gap is:
   - unclear intent
   - blurry scope edge
   - missing non-goals
   - hidden tradeoff
   - unclear approval boundary
5. Keep probing until the answer becomes operationally useful.

First-question rule:

- Default the first question toward success criteria, intended scope, or what must be true for the user to consider the task correctly done before discussing solution details.
- When the first question uses the structured user-input tool for a proposal, direction check, or greenfield setup, make the answer choices represent success criteria, scope priorities, or decision boundaries.
- Do not use the first structured question only to triage what kind of help the user wants if the deeper blocker is still success criteria, intended scope, or approval boundary.

Pressure patterns:

- ask for an example
- ask for a counterexample
- ask what should explicitly not happen
- ask what tradeoff they would reject
- ask what they would want confirmed before proceeding

Output:

- `Alignment progress`
- `Current unresolved edge`

### Phase 2: Lock The Execution Brief

Summarize the clarified brief in execution-ready form:

- `Intent`
- `Desired outcome`
- `In scope`
- `Out of scope`
- `Acceptable tradeoffs`
- `Unacceptable tradeoffs`
- `Decision boundaries`
- `Constraints`
- `Acceptance criteria`
- `Open questions requiring confirmation`

Output:

- `Locked execution brief`

### Phase 3: Handoff To Execution

Hand the locked brief to the active execution workflow or recommend one if none has been chosen yet.

Typical downstream consumers:

- `autopilot`
- `ralph-loop`
- `parallel-work`
- `review-loop`

The downstream workflow should treat the locked brief as the current source of truth for implementation intent unless the user explicitly reopens alignment.

Output:

- `Execution handoff`
- `Residual ambiguity`

## Stop Conditions

- the implementation brief is specific enough that downstream execution can proceed without major intent risk
- the user explicitly chooses to proceed with remaining ambiguity
- further questions stop producing meaningful clarification
- the user says stop, pause, or switch to execution

## Output Contract

- `Initial alignment snapshot`
- `Main alignment risks`
- `Locked execution brief`
- `What may be decided autonomously`
- `What still needs confirmation`
- `Recommended execution handoff`
- `Residual risk`

## Guardrails

- Do not turn this into generic brainstorming.
- Do not jump into implementation from an obviously misaligned brief.
- Do not stop at a high-level advisory answer when the skill has already determined that more user input is required.
- Do not over-question once intent and boundaries are operationally clear.
- Do not treat vague user phrasing as permission to fill in major product decisions silently.
- Do not overwrite the user's intent with a cleaner but different implementation idea.
- Do not mistake a work-defining prompt for a simple opinion request when the user is implicitly trying to define what should happen next.
- Do not reduce this skill to a mere pre-execution checklist; it also owns question-led requirement discovery and direction evaluation when those are the real task.
- Before recommending solution details, first clarify success criteria, scope edges, and non-goals when they are still materially ambiguous.

## Example Triggers

- "구현 전에 내 의도를 인터뷰로 더 명확히 해줘"
- "내가 생각하는 범위를 질문으로 먼저 잠가줘"
- "이 계획을 바로 코드로 옮기지 말고 alignment부터 맞추자"
- "어디까지 구현하고 어디서 멈출지 먼저 깊게 정리해줘"
- "이 버그 수정, 진짜 뭘 고치면 끝난 건지 먼저 맞춰보자"
- "이 오타 수정도 범위랑 기준을 먼저 잠그고 가자"
- "신규 프로젝트 시작하려는데 어때?"
- "이 방향으로 가려는데 뭐부터 정해야 해?"
