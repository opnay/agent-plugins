---
name: frontend-tdd-rgb
description: Apply test-driven frontend workflow. Use when a task should start from a failing test, choose between component, hook, integration, or state-level tests, run Red-Green-Refactor, improve testability, or add coverage for UI behavior, rendering logic, state transitions, and user interactions without brittle assertions.
---

# Frontend TDD RGB

## Overview

Use this skill when the best path is to drive frontend work through a failing test first.
The goal is not to maximize test count. The goal is to choose the right test level, make the behavior fail for the right reason, then implement and refactor without weakening confidence.

## Workflow

1. Define the exact user-visible or contract-visible behavior under test.
2. Pick the narrowest test level that can prove that behavior.
3. Write the smallest failing test with a meaningful assertion.
4. Make the minimal production change to turn the test green.
5. Refactor only after the behavior is protected.

## Output Contract

- `Behavior under test`
- `Chosen test level`
- `Smallest failing test`
- `Production change`
- `Refactor boundary`
- `Residual risk`

## Guardrails

- Do not start from implementation details when user-observable behavior is testable.
- Do not write brittle tests around markup trivia or incidental timing.
- Do not widen production scope to satisfy a poorly chosen test.
- Prefer confidence over coverage theater.

## References

- Read [references/test-level-guide.md](references/test-level-guide.md) when choosing the test boundary.
