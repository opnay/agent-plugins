---
name: turn-gate
description: Main loop controller for `loop-kit`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Overview

Use this skill as the main operational surface of `loop-kit`.
Keep one turn alive until the user asks to end the turn.
Keep the turn shape explicit:

1. analyze the user's message
2. state the plan
3. do the work
4. verify the work
5. report the result or readiness state
6. reopen the next flow with explicit choices
7. continue unless the user asks to end the turn

In this plugin, users do not call `ralph-loop`, `review-loop`, or readiness loops directly.
Instead, this skill selects the right internal loop mode for the current phase.
Read the needed local reference under `references/` before running a mode.
Those references absorb the operational loop contracts into this skill while staying aligned with `workflow-kit` as upstream SSOT.

## Use When

- the repository or working agreement requires non-terminal turns until the user asks to end the turn
- the active work may need a discovery round before execution or refinement can safely continue
- the active work mode may shift between refinement, review handling, and readiness inside the same turn
- the main risk is losing turn continuity or exposing too many direct loop entrypoints

## Do Not Use When

- the task can end cleanly after one response
- the real blocker is still broad workflow selection rather than loop execution
- the work does not need explicit next-flow reopening

## Core Policy

- Treat each incoming message as the current state of the same loop-gated turn.
- Keep `analysis`, `plan`, `work`, `verification`, `result reporting`, and next-flow reopening visible.
- Maintain turn-gate records under `.agents/sessions/{YYYYMMDD}/`.
- Use `000-plan.md` for the higher-level plan when the task spans several flows.
- Keep `000-plan.md` incrementally updated even if one user request ends and later follow-up flows continue the same larger task.
- Update `001+` flow records after each completed flow instead of waiting until the end of the turn.
- Always use the question tool `request_user_input` when opening user choices, scope locks, or next-flow decisions.
- Always use the plan tool `update_plan` once meaningful work begins and keep the active step current as the turn progresses.
- Before `work`, choose one internal loop mode that best owns the current phase.
- Use `request_user_input` whenever mode selection, criteria, or scope is still unclear.
- After `work`, run an explicit verification step before result reporting.
- Reopen the next flow with explicit choices after each result unless the user asks to end the turn.
- Do not expose direct loop entrypoints from this plugin surface.

## Internal Loop Modes

- Read `references/deep-interview.md` and use that contract when the current phase is still requirement discovery and the next step depends on an actual question round.
- Read `references/ralph-loop.md` and use that contract when the work is a bounded issue that benefits from small fix-verify-reassess cycles.
- Read `references/review-loop.md` and use that contract when the work is driven by review findings and only material issues should be fixed.
- Read `references/commit-readiness-gate.md` and use that contract when implementation is largely done and the current question is readiness for commit.
- If none of those modes is clearly selected yet, narrow the choice before continuing the work phase.

## Mode Selection Matrix

- Choose `deep-interview` when the current blocker is requirement discovery, unclear intent, missing scope boundaries, or unresolved approval lines.
- Choose `review-loop` when the input is review feedback, QA findings, or self-review findings and the work should stay bounded to one material issue at a time.
- Choose `ralph-loop` when the work is one bounded improvement cycle and the best next move is a small fix followed by immediate verification.
- Choose `commit-readiness-gate` when implementation is largely done and the current question is whether the intended change unit is ready to move toward commit.
- If more than one mode seems plausible, prefer the earliest blocker in this order: `deep-interview` -> `review-loop` -> `ralph-loop` -> `commit-readiness-gate`.
- If the blocker is still broader than any one internal mode, use `request_user_input` to narrow the mode choice before continuing.

## Session Record

- Use `.agents/sessions/{YYYYMMDD}/000-plan.md` for the higher-level plan across several flows.
- Use the filename shape `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for flow records.
- Keep `count-pad3` zero-padded like `001`, `002`, `003`.
- Keep the slug in English lower-case words joined by `-`.
- Use `.agents/sessions/_turn-gate-flow-template.md` as the default flow-record template.
- Use `.agents/sessions/_turn-gate-plan-template.md` as the default `000-plan.md` template.
- Record at least: user request message, task, flow scope, current mode, analysis, plan, work, verification, result report, next-flow options, residual risk.
- Treat `000-plan.md` as a long-lived incremental plan artifact, not a single-turn note.
- Treat each `001+` file as one flow record, not one phase record.

## Output

- `Analysis`
- `Plan`
- `Chosen internal mode`
- `Work`
- `Verification`
- `Result report`
- `User-response question`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Guardrails

- Do not end the turn by default.
- Do not ask freeform textual choice questions when `request_user_input` can carry the decision.
- Do not skip `update_plan` after moving past initial orientation into real work.
- Do not skip the `000-plan.md` update when the higher-level plan changes across flows.
- Do not skip the `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` flow-record update for a completed flow.
- Do not skip explicit verification between work and result reporting.
- Do not skip the next-flow question after reporting a result.
- Do not let result reporting collapse into a soft closing.
- Do not expose direct user entrypoints for internal loop modes in this plugin.
- Do not drift away from the canonical loop-mode contracts owned by `workflow-kit`.
