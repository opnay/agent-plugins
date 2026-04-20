---
name: screen-spec
description: Turn an aligned design direction into a pre-code screen specification. Use when the task is to define a screen's sections, content intent, interaction expectations, hierarchy, and scope boundaries before Figma or implementation work begins.
---

# Screen Spec

## Overview

Use this skill when the design direction is already aligned enough that detailed screen planning will be useful.
Its job is to convert a design brief or aligned direction into a pre-code screen specification that makes the screen structure, section purpose, content intent, and interaction expectations explicit.
This skill stays before implementation. It does not own frontend architecture, code contracts, or Figma execution.

## Use When

- the high-level design direction is already aligned
- the next useful step is planning one screen in concrete, pre-code terms
- the task needs section structure, hierarchy, content intent, and interaction expectations
- a designer or planner needs a stronger handoff artifact before visual execution

## Do Not Use When

- the direction is still too vague and should become a design brief first
- the main need is critique of an existing weak direction
- the user is asking for code-ready tickets, component props, or engineering specs

## Core Policy

- Specify the screen through structure and intent, not code or pixel instructions.
- Give each section one clear job.
- Keep hierarchy and user progression explicit.
- State what interactions are expected and what should remain simple or deferred.
- Preserve scope boundaries so the screen does not silently absorb nearby product problems.

## Workflow

### 1. Confirm Screen Goal

Clarify:

- primary user task
- screen success condition
- critical user questions the screen must answer

### 2. Define Screen Structure

Specify:

- section order
- section purpose
- content intent per section
- primary and secondary actions

### 3. Define Interaction Expectations

Make explicit:

- expected user flow
- important interaction states
- what needs emphasis, reassurance, or simplification
- what is intentionally deferred

### 4. Produce The Spec

Output a pre-code screen specification with:

- `Screen goal`
- `Primary user task`
- `Success condition`
- `Section map`
- `Section purpose and content intent`
- `Action hierarchy`
- `Interaction expectations`
- `Edge states or unanswered questions`
- `Scope boundaries`
- `Recommended next step`

## Output Contract

- `Screen goal`
- `Primary user task`
- `Success condition`
- `Section map`
- `Section purpose and content intent`
- `Action hierarchy`
- `Interaction expectations`
- `Edge states or unanswered questions`
- `Scope boundaries`
- `Recommended next step`

## Guardrails

- Do not write code or component implementation guidance.
- Do not skip screen goal clarification and jump straight to section lists.
- Do not produce vague section names without explaining their job.
- Do not silently expand one screen into a whole flow unless the request explicitly needs that.

## Example Triggers

- "이 브리프를 바탕으로 화면 스펙을 짜줘"
- "구현 전에 이 화면의 섹션 구조와 상호작용 기대치를 정리해줘"
- "디자인 방향은 맞았으니 이제 pre-code screen spec으로 내려줘"
