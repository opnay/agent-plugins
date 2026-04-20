---
name: ui-critique
description: Critique an existing screen, concept, flow, or design direction from a designer and UX perspective. Use when the task is to identify the biggest hierarchy, clarity, flow, tone, or usability issues and explain which improvements matter most before implementation.
---

# UI Critique

## Overview

Use this skill when something already exists and the main need is designer-style critique rather than invention from scratch.
Its job is to identify the highest-value problems in a screen, flow, or design direction, explain why they matter, and propose a clearer next direction without drifting into implementation detail.
This skill is for judgment and prioritization. It does not own detailed pre-code screen specification unless the critique clearly hands off to that next step.

## Use When

- the user has an existing screen idea, draft, wireframe, or flow that needs review
- the main problem is not "what should we build?" but "what is wrong or weak about this direction?"
- the user wants hierarchy, UX, tone, or clarity feedback from a designer perspective
- the work needs prioritization of issues instead of a full redesign spec

## Do Not Use When

- the design direction is still too vague and needs a brief first
- the user already knows the direction and now needs a concrete screen spec
- the user is asking for code-level implementation critique

## Core Policy

- Prioritize the few issues that most change user understanding or task success.
- Critique in terms of hierarchy, flow, clarity, trust, and interaction burden rather than taste alone.
- Explain impact, not just preference.
- Separate blocking issues from polish notes.
- Keep the critique actionable enough to guide a next design pass.

## Critique Lenses

Use these lenses as needed:

- information hierarchy
- action clarity
- user flow and friction
- content density and cognitive load
- tone, trust, and emotional signal
- consistency between intention and visible UI

## Workflow

### 1. Read The Intended Outcome

Clarify what the screen or flow is trying to help the user do.

### 2. Identify The Highest-Impact Problems

Separate:

- blocking issues
- important but non-blocking issues
- low-value polish notes

### 3. Explain Why They Matter

For each important issue, explain:

- what the problem is
- how it affects the user
- why it undermines the design goal

### 4. Recommend A Better Direction

Produce:

- `Strongest problems`
- `Why they matter`
- `Priority order`
- `Suggested design direction`
- `What to preserve`
- `Recommended next step`

## Output Contract

- `Intended outcome`
- `Strongest problems`
- `Why they matter`
- `Priority order`
- `What to preserve`
- `Suggested design direction`
- `Recommended next step`
- `Residual risk`

## Guardrails

- Do not reduce critique to taste or aesthetics alone.
- Do not list minor polish items before identifying the structural issues.
- Do not turn critique into a full redesign spec unless the next step explicitly requires it.
- Do not pretend every issue is equally important.

## Example Triggers

- "이 화면 아이디어를 디자이너 관점에서 critique해줘"
- "이 플로우의 UX 문제가 뭔지 우선순위로 정리해줘"
- "무엇이 가장 어색하고 왜 그런지 디자인 리뷰해줘"
