---
name: frontend-design-guide
description: Guide frontend UI decisions with designer-level implementation rules. Use when a task involves hierarchy, layout, spacing, typography, states, responsiveness, accessibility, design tokens, or turning product intent into implementable screen, component, and interaction guidance for frontend code.
---

# Frontend Design Guide

## Overview

Use this skill when a frontend task changes how the interface looks, reads, or behaves.
The goal is to convert design intent into concrete implementation guidance for hierarchy, interaction, and consistency.

## Workflow

1. Clarify the screen goal, user task, and primary action.
2. Review hierarchy, layout, spacing, typography, states, and responsive behavior.
3. Check accessibility and token usage before cosmetic polish.
4. Recommend the smallest set of changes with the highest UX payoff.
5. End with guidance an implementer can apply directly in code.

## Output Contract

- `Screen goal`
- `UX findings`
- `Priority design changes`
- `Token and state guidance`
- `Responsive and accessibility notes`
- `Implementation handoff`

## Guardrails

- Do not reduce design feedback to subjective taste.
- Start from user task clarity before visual polish.
- Prefer reusable tokens and rules over one-off fixes.
- Call out when a recommendation depends on product context that is not yet known.

## References

- Read [references/review-rubric.md](references/review-rubric.md) for evaluation order.
