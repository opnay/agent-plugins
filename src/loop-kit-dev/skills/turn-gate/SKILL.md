---
name: turn-gate
description: Main loop controller for `loop-kit-dev`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Overview

Use this skill as the main operational surface of `loop-kit-dev`.
Using this skill means treating it as a first-class rule for the rest of the current session.
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
It also selects the question-routing mode for how blocked decisions are answered.
Read the needed local reference under `references/` before running a mode.
Those references absorb the operational loop contracts into this skill while staying aligned with `workflow-kit` as upstream SSOT.

## Use When

- the repository or working agreement requires non-terminal turns until the user asks to end the turn
- the active work may need a discovery round before execution or refinement can safely continue
- the active work mode may shift between autonomous execution, refinement, review handling, and readiness inside the same turn
- the main risk is losing turn continuity or exposing too many direct loop entrypoints

## Do Not Use When

- the task can end cleanly after one response
- the real blocker is still broad workflow selection rather than loop execution
- the work does not need explicit next-flow reopening

## Core Policy

- Treat invocation of this skill as activation of a session-level first-class operating rule.
- Treat each incoming message as the current state of the same loop-gated turn.
- Keep `analysis`, `plan`, `work`, `verification`, `result reporting`, and next-flow reopening visible.
- Maintain turn-gate records under `.agents/sessions/{YYYYMMDD}/`.
- Maintain a compact `Continuity Guard` in every flow record and refresh it before result reporting and next-flow reopening.
- The `Continuity Guard` must state whether `turn-gate` is active, the question-routing mode, whether the user explicitly stopped the turn, whether a terminal summary is allowed, and the required next action.
- Use `000-plan.md` for the higher-level plan when the task spans several flows.
- Keep `000-plan.md` incrementally updated even if one user request ends and later follow-up flows continue the same larger task.
- Use `analysis` and `plan` to design not only the current flow but also likely next flows or phases when forward design is useful.
- Treat that forward flow/phase design as provisional and revise it in a later `analysis` or `plan` step only when new evidence, changed intent, or a revealed blocker makes redesign necessary.
- Update `001+` flow records incrementally at each completed phase instead of batching them at the end of the flow.
- Use the question-routing axis when opening choices, scope locks, or next-flow decisions.
- In the default user-gated mode, use `request_user_input` for those questions.
- In `self-drive` mode, ask subagents instead of the user and continue from their answer.
- Always use the plan tool `update_plan` once meaningful work begins and keep the active step current as the turn progresses.
- Before `work`, choose one internal loop mode that best owns the current phase.
- Also choose one question-routing mode: default `user-gated` or `self-drive`.
- Use the active question-routing mode whenever mode selection, criteria, scope, or next-flow choice is still unclear.
- After `work`, run an explicit verification step before result reporting, and use that verification to surface whether later flow/phase redesign is needed.
- Before result reporting, read or reconstruct the `Continuity Guard`; if the user has not explicitly stopped the turn, a terminal summary is invalid.
- Reopen the next flow with explicit choices after each result unless the user asks to end the turn; in `self-drive`, route that choice to a subagent and continue automatically.
- Do not expose direct loop entrypoints from this plugin surface.

## Internal Loop Modes

- Read `references/deep-interview.md` and use that contract when the current phase is still requirement discovery and the next step depends on an actual question round.
- Read `references/autopilot.md` and use that contract when the current phase is broad end-to-end delivery from a brief request to a verified result.
- Read `references/ralph-loop.md` and use that contract when the work is a bounded issue that benefits from small fix-verify-reassess cycles.
- Read `references/review-loop.md` and use that contract when the work is driven by review findings and only material issues should be fixed.
- Read `references/commit-readiness-gate.md` and use that contract when implementation is largely done and the current question is readiness for commit.
- If none of those modes is clearly selected yet, narrow the choice before continuing the work phase.

## Mode Selection Matrix

- Choose `deep-interview` when the current blocker is requirement discovery, unclear intent, missing scope boundaries, or unresolved approval lines.
- Choose `review-loop` when the input is review feedback, QA findings, or self-review findings and the work should stay bounded to one material issue at a time.
- Choose `ralph-loop` when the work is one bounded improvement cycle and the best next move is a small fix followed by immediate verification.
- Choose `autopilot` when the work is broad end-to-end delivery that spans scope assumptions, implementation, QA, validation, and delivery.
- Choose `commit-readiness-gate` when implementation is largely done and the current question is whether the intended change unit is ready to move toward commit.
- If more than one mode seems plausible, prefer the earliest blocker in this order: `deep-interview` -> `review-loop` -> `ralph-loop` -> `autopilot` -> `commit-readiness-gate`.
- If the blocker is still broader than any one internal mode, use the active question-routing mode to narrow the mode choice before continuing.

## Question Routing Axis

- Use `user-gated` by default: ask the user through `request_user_input` for choices, scope locks, and next-flow decisions.
- Use `self-drive` when the user wants the loop to continue without user intervention.
- In `self-drive`, read `references/self-drive.md`, send subagents a self-drive question packet for questions that would otherwise go to the user, require the self-drive answer contract, record the answer and assumptions, then continue the loop from that answer.
- Every `self-drive` packet must carry the current `Continuity Guard`, and every answer must include a continuity check that preserves next-flow continuation unless a hard approval boundary or explicit user stop exists.
- `self-drive` may answer mode selection, criteria, scope assumptions, verification choices, and next-flow decisions through subagents.
- In `self-drive`, recover subagent `context_gap` results through main-agent discovery when the missing evidence can be found without an explicit approval boundary.
- Treat missing user preference as a reversible assumption to record and continue unless the user explicitly requested manual preference locking.
- Treat `low` confidence as an approval-boundary pause only when the missing decision requires explicit approval, destructive/irreversible/external action approval, or a platform/tool/safety boundary.
- In `self-drive`, an approval-boundary pause stops autonomous routing only; switch to `user-gated` and use `request_user_input` instead of ending the turn.
- `self-drive` does not override platform, tool, or safety policies that require explicit user approval.

## Session Record

- Use `.agents/sessions/{YYYYMMDD}/000-plan.md` for the higher-level plan across several flows.
- Use the filename shape `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for flow records.
- Keep `count-pad3` zero-padded like `001`, `002`, `003`.
- Keep the slug in English lower-case words joined by `-`.
- Use `templates/flow-record-template.md` as the default flow-record template.
- Use `templates/plan-template.md` as the default `000-plan.md` template.
- Record at least: user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, residual risk.
- Treat `000-plan.md` as a long-lived incremental plan artifact, not a single-turn note.
- Treat each `001+` file as one flow record, not one phase record.
- Keep the current flow record current after `analysis`, `plan`, `work`, `verification`, and `result reporting`.

## Output

- `Analysis`
- `Plan`
- `Chosen internal mode`
- `Question-routing mode`
- `Work`
- `Verification`
- `Result report`
- `Continuity guard`
- `Question-routing prompt`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Guardrails

- Do not end the turn by default.
- Do not ask freeform textual choice questions when the active question-routing mode can carry the decision.
- Do not route user questions to subagents unless `self-drive` is active.
- Do not let `self-drive` simulate user approval where the runtime or tool policy requires explicit approval.
- Do not end the turn when `self-drive` reaches an approval boundary; pause self-drive, switch to `user-gated`, and ask through `request_user_input`.
- Do not skip `update_plan` after moving past initial orientation into real work.
- Do not skip the `000-plan.md` update when the higher-level plan changes across flows.
- Do not defer `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` updates until the end of the flow; write them at each completed phase boundary.
- Do not skip explicit verification between work and result reporting.
- Do not emit result reporting until the `Continuity Guard` says whether next-flow reopening is still required.
- Do not skip the next-flow question after reporting a result.
- Do not let result reporting collapse into a soft closing.
- Do not expose direct user entrypoints for internal loop modes in this plugin.
- Do not drift away from the canonical loop-mode contracts owned by `workflow-kit`.
