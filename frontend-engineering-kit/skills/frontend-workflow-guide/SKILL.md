---
name: frontend-workflow-guide
description: Default workflow for frontend implementation, refactor, fix, and review tasks. Use when a task involves building or changing components, hooks, pages, screens, interactions, UI state, rendering behavior, or frontend feature logic and Codex should first inspect project structure, choose fitting boundaries, apply UI implementation rules, and prefer a failing-test-first Red-Green-Refactor path when practical.
---

# Frontend Workflow Guide

## Overview

Use this skill as the default entrypoint for frontend work in this kit.
When the task is frontend-facing, do not jump straight into code. First decide how the existing project structure should shape the work, what responsibility candidates actually need evaluation, what level of UI and domain guidance is relevant, and whether the change should be driven by a failing test first.

## Workflow

1. Confirm the frontend surface being changed: component, hook, page, screen, flow, or feature slice.
2. Inspect the local project structure before proposing new folders, abstractions, shared moves, or ownership changes.
3. Extract the responsibility candidates that actually need a decision.
4. Decide the dominant concern for the task:
   - component boundary and ownership problem
   - domain or feature boundary problem
   - UI implementation or design quality problem
   - test strategy or Red-Green-Refactor problem
5. For architecture-heavy work, decide:
   - who owns the responsibility
   - which layer it belongs to
   - whether it should stay local, move, or be extracted
   - whether the component should be treated as library-level or product-level
6. Load only the guidance needed for the dominant concern, but keep the other concerns visible enough to avoid local optimization.
7. When practical, define the behavior through the narrowest failing test before production edits.
8. Make the smallest change that respects project structure, ownership clarity, UI quality, and testability.

## Decision Rules

- If the task changes ownership of rendering, state, effects, hooks, reusable APIs, or refactor direction, reason about component boundaries before implementation.
- If the task changes business concepts, feature slices, view models, or where business rules live, reason about domain and layering before implementation.
- If the task changes hierarchy, layout, spacing, states, responsiveness, accessibility, or tokens, reason about UI implementation quality before implementation.
- If the task changes observable behavior and can be protected by tests, choose a test level and use Red-Green-Refactor before widening production changes.
- If the task looks like a component problem, extract the responsibility candidate before deciding whether it belongs in component-local code, utility, api, domain, or feature space.
- If the task suggests hook extraction, decide the candidate and target layer first; do not treat hooks as the default cleanup move.
- If the task changes a component API, decide whether that component is library-level or product-level before judging props, callbacks, or composition.
- If multiple concerns apply, prioritize them in this order:
  1. Wrong local structure fit or wrong boundary
  2. Wrong responsibility ownership or wrong layer
  3. Wrong behavior or missing test protection
  4. Wrong UI clarity or implementation quality

## Output Contract

- `Task shape`
- `Project structure read`
- `Responsibility candidates`
- `Primary concern`
- `Layer decision`
- `Hook or local decision`
- `Component level`
- `Chosen workflow`
- `Verification path`
- `Residual risk`

## Guardrails

- Do not invent architecture before reading the current project layout.
- Do not introduce a new structural pattern if the repo already has a workable one nearby.
- Do not skip test strategy thinking just because the change looks small.
- Do not treat design as optional when the task changes what users see or interact with.
- Do not load every guide by default; load the narrowest relevant guidance after the first structural read.
- Do not let a failing test force a bad boundary decision.
- Do not choose hook extraction before deciding whether the responsibility should stay local or move to another layer.

## Concept Lenses

Use these lenses implicitly as needed:

- Boundary lens:
  - who owns rendering, state, effects, reusable APIs, and refactor direction
- Domain lens:
  - what is business logic, what is orchestration, what is presentation, and what belongs to a feature boundary
- UI lens:
  - what improves hierarchy, state clarity, responsiveness, accessibility, token consistency, and component level fit
- Test lens:
  - what is the narrowest failing test that protects the intended behavior
