---
name: loop-kit-guide
description: Entrypoint skill for `loop-kit`. Use it to decide whether the task should start in the narrow turn-continuity package and, if so, route into `turn-gate`.
---

# Loop Kit Guide

## Overview

Use this skill as the default entrypoint for `loop-kit`.
Its job is to decide whether the current task should begin in this narrow loop package.
If the answer is yes, route into `turn-gate`.
Do not expose direct loop entrypoints such as `ralph-loop` or `review-loop` to the user from this plugin surface.

## Routing Rules

- Start in `turn-gate` when the repository or task requires the turn to stay open until the user asks to end the turn.
- Start in `turn-gate` when the main value is keeping the turn alive while the active work mode changes across autonomous execution, refinement, review, or readiness.
- Start in `turn-gate` when the user wants self-driving progress where blocked questions are answered by subagents instead of by the user.
- Do not start in `loop-kit` when the main problem is still broad workflow selection, requirement discovery, or planning outside a loop-gated turn.
- When `loop-kit` is not the best starting point, say so explicitly and prefer the broader `workflow-kit` surface.

## Output

- `Loop fit`
- `Why loop-kit or not`
- `Chosen entrypoint`
- `Residual risk`

## Guardrails

- Do not offer direct loop skills as user-facing entrypoints in this plugin.
- Do not treat loop continuity as required when the task can end cleanly in one response.
- Do not hide a broader workflow-selection problem inside `loop-kit`.
