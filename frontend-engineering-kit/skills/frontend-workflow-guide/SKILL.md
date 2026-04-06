---
name: frontend-workflow-guide
description: Default workflow for frontend implementation, refactor, fix, and review tasks. Use when a task involves building or changing components, hooks, pages, screens, interactions, UI state, rendering behavior, or frontend feature logic and Codex should first inspect project structure, choose fitting boundaries, apply UI implementation rules, and prefer a failing-test-first Red-Green-Refactor path when practical.
---

# Frontend Workflow Guide

## Overview

Use this skill as the default entrypoint for frontend work in this kit.
When the task is frontend-facing, do not jump straight into code. First decide how the existing project structure should shape the work, what level of UI and domain guidance is relevant, and whether the change should be driven by a failing test first.

## Workflow

1. Confirm the frontend surface being changed: component, hook, page, screen, flow, or feature slice.
2. Inspect the local project structure before proposing new folders, abstractions, or state ownership.
3. Decide the architectural shape of the change:
   - component boundary problem
   - domain or feature boundary problem
   - UI implementation or design quality problem
   - test strategy or Red-Green-Refactor problem
4. Load only the guidance needed for the dominant concern, but keep the other concerns visible enough to avoid local optimization.
5. When practical, define the behavior through the narrowest failing test before production edits.
6. Make the smallest change that respects project structure, UI quality, and testability.

## Decision Rules

- If the task changes ownership of rendering, state, effects, hooks, or reusable APIs, reason about component boundaries before implementation.
- If the task changes business concepts, feature slices, view models, or where business rules live, reason about domain and layering before implementation.
- If the task changes hierarchy, layout, spacing, states, responsiveness, accessibility, or tokens, reason about UI implementation quality before implementation.
- If the task changes observable behavior and can be protected by tests, choose a test level and use Red-Green-Refactor before widening production changes.
- If multiple concerns apply, prioritize them in this order:
  1. Wrong architecture or boundary
  2. Wrong behavior or missing test protection
  3. Wrong UI clarity or implementation quality

## Output Contract

- `Task shape`
- `Project structure read`
- `Primary concern`
- `Chosen workflow`
- `Verification path`
- `Residual risk`

## Guardrails

- Do not invent architecture before reading the current project layout.
- Do not skip test strategy thinking just because the change looks small.
- Do not treat design as optional when the task changes what users see or interact with.
- Do not load every guide by default; load the narrowest relevant guidance after the first structural read.
- Do not let a failing test force a bad boundary decision.

## Concept Lenses

Use these lenses implicitly as needed:

- Boundary lens:
  - who owns rendering, state, effects, and reusable APIs
- Domain lens:
  - what is business logic, what is orchestration, and what is presentation
- UI lens:
  - what improves hierarchy, state clarity, responsiveness, accessibility, and token consistency
- Test lens:
  - what is the narrowest failing test that protects the intended behavior
