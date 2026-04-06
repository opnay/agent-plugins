---
name: component-architecture
description: Design or review frontend component boundaries. Use when a task involves splitting large components, choosing presentational vs container responsibilities, extracting hooks, shaping component APIs, colocating state, or deciding composition, ownership, and folder structure for React or UI code.
---

# Component Architecture

## Overview

Use this skill when a frontend task is primarily about component shape and responsibility.
The goal is to decide where rendering, state, effects, domain calls, and reusable abstractions should live before implementation sprawls.

## Workflow

1. Map the current component responsibilities and data flow.
2. Identify mixed concerns: rendering, business logic, side effects, state orchestration, and styling.
3. Recommend the smallest architectural split that improves clarity without creating abstraction theater.
4. Define public APIs, ownership boundaries, and state placement.
5. End with an implementation-ready boundary recommendation.

## Output Contract

- `Current responsibility map`
- `Boundary problems`
- `Recommended split`
- `Public API shape`
- `State and hook placement`
- `Migration risk`

## Guardrails

- Do not split components just because they are large.
- Prefer composition over inheritance-style abstraction.
- Prefer existing project patterns unless they are actively causing confusion.
- Keep boundaries understandable to the next maintainer.

## References

- Read [references/boundary-checklist.md](references/boundary-checklist.md) for split criteria.
