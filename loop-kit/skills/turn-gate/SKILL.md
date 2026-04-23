---
name: turn-gate
description: Main loop controller for `loop-kit`. Keep one turn alive until explicit user stop and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Overview

Use this skill as the main operational surface of `loop-kit`.
Keep one turn alive until the user explicitly ends the work.
Keep the turn shape explicit:

1. analyze the user's message
2. state the plan
3. do the work
4. report the result or readiness state
5. reopen the next flow with explicit choices
6. continue unless the user explicitly stops

In this plugin, users do not call `ralph-loop`, `review-loop`, or readiness loops directly.
Instead, this skill selects the right internal loop mode for the current phase.
Read the needed local reference under `references/` before running a mode.
Those references absorb the operational loop contracts into this skill while staying aligned with `workflow-kit` as upstream SSOT.

## Use When

- the repository or working agreement requires non-terminal turns until explicit user stop
- the active work mode may shift between refinement, review handling, and readiness inside the same turn
- the main risk is losing turn continuity or exposing too many direct loop entrypoints

## Do Not Use When

- the task can end cleanly after one response
- the real blocker is still broad workflow selection rather than loop execution
- the work does not need explicit next-flow reopening

## Core Policy

- Treat each incoming message as the current state of the same loop-gated turn.
- Keep `analysis`, `plan`, `work`, `result reporting`, and next-flow reopening visible.
- Always use the question tool `request_user_input` when opening user choices, scope locks, or next-flow decisions.
- Always use the plan tool `update_plan` once meaningful work begins and keep the active step current as the turn progresses.
- Before `work`, choose one internal loop mode that best owns the current phase.
- Use `request_user_input` whenever mode selection, criteria, or scope is still unclear.
- Reopen the next flow with explicit choices after each result unless the user explicitly stops.
- Do not expose direct loop entrypoints from this plugin surface.

## Internal Loop Modes

- Read `references/ralph-loop.md` and use that contract when the work is a bounded issue that benefits from small fix-verify-reassess cycles.
- Read `references/review-loop.md` and use that contract when the work is driven by review findings and only material issues should be fixed.
- Read `references/commit-readiness-gate.md` and use that contract when implementation is largely done and the current question is readiness for commit.
- If none of those modes is clearly selected yet, narrow the choice before continuing the work phase.

## Output

- `Analysis`
- `Plan`
- `Chosen internal mode`
- `Work`
- `Result report`
- `User-response question`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Guardrails

- Do not end the turn by default.
- Do not ask freeform textual choice questions when `request_user_input` can carry the decision.
- Do not skip `update_plan` after moving past initial orientation into real work.
- Do not skip the next-flow question after reporting a result.
- Do not let result reporting collapse into a soft closing.
- Do not expose direct user entrypoints for internal loop modes in this plugin.
- Do not drift away from the canonical loop-mode contracts owned by `workflow-kit`.
