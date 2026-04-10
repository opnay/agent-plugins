---
name: frontend-engineering-kit-guide
description: Entrypoint skill for the `frontend-engineering-kit` plugin. Use when a task involves frontend engineering work and Codex should first decide whether the job is primarily about overall frontend workflow, frontend architecture-pattern choice, React structure and rendering boundaries, local component boundaries, domain modeling, UI implementation quality, or test-driven execution.
---

# Frontend Engineering Kit Guide

## Overview

Use this skill as the default entrypoint for `frontend-engineering-kit`.
Its job is to classify the frontend task before deeper execution starts and route to the narrowest bundled skill that fits.
Do not jump straight into a specialized skill when the task is broad, mixed, or still needs a first structural read.

## Workflow

1. Identify the frontend task shape:
   - broad implementation, refactor, fix, or review work
   - frontend architecture-pattern decision
   - React structure, rendering boundary, or design-system boundary decision
   - component boundary or ownership decision
   - domain, feature, or layering decision
   - UI implementation quality or design guidance
   - test-driven execution strategy
2. Decide whether the task is:
   - one dominant concern
   - several frontend concerns that need ordering
3. Route to the narrowest bundled skill that owns the main concern.
4. If the task spans several concerns, choose the starting skill and execution order explicitly instead of blending them together.

## Routing Rules

- Choose `frontend-workflow-guide` when the task is broad, mixed, or needs an initial project-structure read before deciding the dominant concern.
- Choose `frontend-architecture-patterns` when the main problem is Atomic/FSD/custom structure fit, architecture-pattern choice, migration direction, design-system versus product-app structure, or reviewing whether the frontend needs UI taxonomy, feature slicing, or both.
- Choose `react-architecture` when the main problem is React hierarchy inside the chosen pattern, Template/Page or widget/shared boundaries, prop contract rules, context/provider scope, or rerender spread caused by poor structure inside the accepted architecture.
- Choose `component-architecture` when the main problem is local component boundaries, ownership, extraction, composition, hook placement, or split direction inside an existing feature or screen.
- Choose `frontend-domain-modeling` when the main problem is business concepts, feature boundaries, layering, domain rules, or where business logic should live.
- Choose `frontend-design-guide` when the main problem is hierarchy, layout, spacing, typography, states, responsiveness, accessibility, or token-driven UI quality.
- Choose `frontend-tdd-rgb` when the main problem is how to drive the change through the right failing test, test level, and Red-Green-Refactor loop.
- If the task clearly crosses several of these concerns, start with `frontend-workflow-guide` and let it choose the first deeper path.

## Decision Rules

- Treat `frontend-workflow-guide` as the default mode when the user asks for a general frontend change rather than a narrow architectural or design decision.
- Route directly to a specialized skill only when the dominant concern is already clear from the request.
- Prefer one clear starting skill plus an explicit handoff order over mixing several skill responsibilities into one ambiguous prompt.
- Choose `frontend-architecture-patterns` before `react-architecture` when the real question is whether the repository should be Atomic, FSD, or hybrid at all.
- Choose `react-architecture` before `component-architecture` when the real question is about React layer policy, provider scope, or rerender spread across already accepted architectural boundaries.
- If the task changes both structure and visuals, decide structure first unless the user explicitly asked for design critique only.
- If the task changes observable behavior and test protection is part of the risk, keep `frontend-tdd-rgb` visible even when another skill owns the first step.

## Output Contract

- `Task shape`
- `Dominant concern`
- `Chosen skill`
- `Why this route fits`
- `Planned handoff order`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not route broad frontend work straight into a narrow specialist skill before reading the local project shape.
- Do not treat hook extraction, folder moves, or token cleanup as default answers before classifying the real problem.
- Do not bury plugin-level routing guidance inside sibling skills.
- Do not send a mixed frontend task into several skills at once without naming the starting point and handoff order.
