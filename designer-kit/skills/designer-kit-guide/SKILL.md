---
name: designer-kit-guide
description: Entrypoint skill for the `designer-kit` plugin. Use when a design-only task needs the right design workflow chosen first, such as deciding whether the job is to define a design brief, critique an existing direction, or produce a pre-code screen specification.
---

# Designer Kit Guide

## Overview

Use this skill as the default entrypoint for `designer-kit`.
Its job is to classify the design task before deeper execution starts and route to the narrowest bundled skill that fits.
Do not jump straight into a specialist skill when the request is still broad, mixed, or unclear about whether it needs invention, critique, or specification.

## Workflow

1. Identify the design task shape:
   - early direction setting from a vague idea
   - critique of an existing screen, flow, or direction
   - pre-code specification of a screen that is already directionally aligned
2. Decide whether the task is:
   - one dominant design concern
   - several design concerns that need ordering
3. Route to the narrowest bundled skill that owns the main concern.
4. If the task spans several concerns, choose the starting skill and handoff order explicitly instead of blending them together.

## Routing Rules

- Choose `design-brief` when the main need is to define the problem, audience, user outcome, hierarchy, tone, non-goals, and directional design intent before a screen is specified.
- Choose `ui-critique` when the main need is to review an existing concept, screen, structure, or flow and identify the biggest UX, hierarchy, clarity, or tone problems.
- Choose `screen-spec` when the main need is to translate an already aligned direction into a pre-code screen-level specification with sections, content intent, interaction expectations, and scope boundaries.
- If the task starts vague and then needs concrete screen-level structure, start with `design-brief` and hand off to `screen-spec`.
- If the task begins from an existing design idea but the real problem is that the direction itself is weak, start with `ui-critique` and then decide whether a revised brief or a screen spec should follow.

## Decision Rules

- Treat `design-brief` as the starting point when the request is still about what should be designed and why.
- Treat `ui-critique` as the starting point when something already exists and the user wants designer-style feedback rather than invention from scratch.
- Treat `screen-spec` as the starting point only when the underlying direction is already aligned enough that detailed screen planning will not be wasted.
- Prefer one clear starting skill plus an explicit handoff order over mixing invention, critique, and specification into one ambiguous prompt.
- Keep this plugin design-first: if the user is really asking for code, implementation details, or frontend architecture, hand off to another plugin rather than stretching `designer-kit`.

## Output Contract

- `Task shape`
- `Dominant concern`
- `Chosen skill`
- `Why this route fits`
- `Planned handoff order`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not treat pre-code design work as frontend implementation guidance.
- Do not bury plugin-level routing guidance inside sibling skills.
- Do not route broad design ambiguity straight into `screen-spec`.
- Do not route invention work into `ui-critique` just because the request mentions a screen.
- Do not send mixed design tasks into several skills at once without naming the starting point and handoff order.

## Example Triggers

- "이 디자인 작업은 브리프부터 해야 할지, critique부터 해야 할지 먼저 골라줘"
- "designer-kit으로 이 요청을 어떤 디자인 스킬로 처리할지 판단해줘"
- "화면을 짜기 전에 먼저 어떤 디자인 workflow가 맞는지 정해줘"
