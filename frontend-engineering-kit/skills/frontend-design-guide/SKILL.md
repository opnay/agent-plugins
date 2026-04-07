---
name: frontend-design-guide
description: Guide frontend UI decisions with designer-level implementation rules. Use when a task involves hierarchy, layout, spacing, typography, states, responsiveness, accessibility, design tokens, reusable design-system patterns, or turning product intent into implementable screen, component, and interaction guidance for frontend code.
---

# Frontend Design Guide

## Overview

Use this skill when a frontend task changes how the interface looks, reads, or behaves.
The goal is to convert design intent into concrete implementation guidance for hierarchy, interaction, and consistency.

## Workflow

1. Clarify the screen goal, user task, and primary action.
2. Detect whether the main problem is design-system level or product level before reviewing details.
3. Decide whether the current hierarchy makes the primary task, primary information, and primary CTA obvious enough.
4. Decide whether UI states are communicated clearly enough for the user to know the current status and next action.
5. Review layout, spacing, typography, and responsive behavior with level-appropriate criteria.
6. Check accessibility and token usage before cosmetic polish.
7. Recommend the smallest set of changes with the highest UX payoff.
8. End with guidance an implementer can apply directly in code.

## Design Level Rules

- `Design-system level`
  - Use when the main problem is about shared component consistency, token discipline, variant control, or reusable UI pattern quality.
  - Prioritize token correctness, spacing and typography scale consistency, variant discipline, and baseline accessibility.
- `Product level`
  - Use when the main problem is tied to a specific screen, section, feature flow, CTA priority, or state communication for a concrete user task.
  - Prioritize task clarity, information hierarchy, CTA emphasis, local state visibility, and workflow-fit responsiveness.
- `Cross-level smell`
  - Do not treat a product-level workflow-fit issue as only a token cleanup problem.
  - Do not treat a design-system consistency problem as only a one-screen polish problem.

## Hierarchy And State Decision Rules

- `Hierarchy decision`
  - Decide what the user must notice first, what they should understand second, and what action they should be able to take next.
  - Give the primary task, primary information, and primary CTA clearly more visual weight than secondary content.
  - Use grouping and spacing to support decision-making order, not just visual neatness.
- `State communication decision`
  - Make loading, empty, error, success, disabled, active, and blocked states meaningfully distinct.
  - Do not rely on color alone to communicate state.
  - When a state changes what the user can do next, communicate both the state and the next expected action.
- `Hierarchy and state cross-check`
  - If state changes, re-check hierarchy and CTA emphasis instead of assuming the existing visual priority still fits.
  - The most important current status should be easier to notice than lower-priority supporting content.

## Output Contract

- `Screen goal`
- `Design level detection`
- `Hierarchy decision`
- `State communication decision`
- `UX findings`
- `Priority design changes`
- `Token and state guidance`
- `Responsive and accessibility notes`
- `Implementation handoff`

## Guardrails

- Do not reduce design feedback to subjective taste.
- Start from user task clarity before visual polish.
- Prefer reusable tokens and rules over one-off fixes.
- Do not mix design-system consistency guidance with product-level task guidance without stating the level first.
- Do not flatten hierarchy by giving primary and secondary actions the same visual force.
- Do not treat loading, empty, error, and disabled states as interchangeable placeholders.
- Call out when a recommendation depends on product context that is not yet known.

## References

- Read [references/design-level-rules.md](references/design-level-rules.md) to decide whether the work is design-system level or product level.
- Read [references/hierarchy-and-state-rules.md](references/hierarchy-and-state-rules.md) for hierarchy and state communication decisions.
- Read [references/review-rubric.md](references/review-rubric.md) for evaluation order.
