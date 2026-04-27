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
- Treat every user message as authoritative loop input, not as a reason to stop.
- Classify an in-turn user message as one of: explicit turn stop, current-flow correction, current-flow priority change, or next-flow priority request.
- If the message changes the active work, adjust the current flow immediately from the earliest safe phase.
- If the message is a valid follow-up that does not need to interrupt the active work, register it as the highest-priority next flow in the flow record and continue to the next safe handoff point.
- Keep `analysis`, `plan`, `work`, `verification`, `result reporting`, and next-flow reopening visible.
- Maintain turn-gate records under `.agents/sessions/{YYYYMMDD}/`.
- Maintain a compact `Continuity Guard` in every flow record and refresh it before result reporting and next-flow reopening.
- The `Continuity Guard` must state whether `turn-gate` is active, the question-routing mode, whether the user explicitly stopped the turn, whether a terminal summary is allowed, and the required next action.
- Record an explicit turn-end option in the flow record's `Next Flow Options` even when the user-facing question already has three visible choices and cannot display that option.
- Use `000-plan.md` as the date-scoped plan: the durable history, user-request list, flow index, current plan, and completed-flow summary for `.agents/sessions/{YYYYMMDD}/`.
- Keep `000-plan.md` incremental. Do not delete completed work from it; summarize completed flows and keep their references.
- Use `001+` files as detailed flow reports for each user-request-driven work flow.
- Use `analysis` and `plan` to design not only the current flow but also likely next flows or phases when forward design is useful.
- Treat that forward flow/phase design as provisional and revise it in a later `analysis` or `plan` step only when new evidence, changed intent, or a revealed blocker makes redesign necessary.
- Update `001+` flow records incrementally at each completed phase instead of batching them at the end of the flow.
- Use user-gated question routing when opening choices, scope locks, or next-flow decisions.
- Use `request_user_input` for those questions.
- Always use the plan tool `update_plan` once meaningful work begins and keep the active step current as the turn progresses.
- Before `work`, choose one internal loop mode that best owns the current phase.
- Use user-gated question routing whenever mode selection, criteria, scope, or next-flow choice is still unclear.
- After `work`, run an explicit verification step before result reporting, and use that verification to surface whether later flow/phase redesign is needed.
- Before result reporting, read or reconstruct the `Continuity Guard`; if the user has not explicitly stopped the turn, a terminal summary is invalid.
- Reopen the next flow with explicit choices after each result unless the user asks to end the turn.
- If three or more user-facing next-flow choices are already needed, keep those visible choices narrow and record a separate turn-end option in the session flow record even if it is not displayed.
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

## Question Routing

- Use `user-gated` question routing: ask the user through `request_user_input` for choices, scope locks, and next-flow decisions.
- Treat missing user preference as a reversible assumption only when the active internal mode allows proceeding without a scope lock.
- Explicit user, tool, platform, safety, destructive, irreversible, or external-action approval boundaries must be asked through `request_user_input`.

## Session Record

- Use `.agents/sessions/{YYYYMMDD}/000-plan.md` as the date-scoped plan, history, and flow index for that day's turn-gated work.
- Use the filename shape `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for detailed flow reports.
- Keep `count-pad3` zero-padded like `001`, `002`, `003`.
- Keep the slug in English lower-case words joined by `-`.
- Each user-request-driven flow should get its own `001+` flow report when it needs detailed tracking.
- The plan should list the day's user requests and the corresponding `001+` flow reports.
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
- Do not route user-gated questions to subagents from this skill.
- Do not simulate user approval where the runtime or tool policy requires explicit approval.
- Do not treat in-turn user intervention as a stop, completion, or approval-boundary pause unless the user explicitly asks to stop the turn or the message creates a real approval boundary.
- Do not skip `update_plan` after moving past initial orientation into real work.
- Do not skip the `000-plan.md` update when the higher-level plan changes across flows.
- Do not defer `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` updates until the end of the flow; write them at each completed phase boundary.
- Do not skip explicit verification between work and result reporting.
- Do not emit result reporting until the `Continuity Guard` says whether next-flow reopening is still required.
- Do not skip the next-flow question after reporting a result.
- Do not omit the session-recorded turn-end option from `Next Flow Options`, even when the visible question options are full.
- Do not let result reporting collapse into a soft closing.
- Do not expose direct user entrypoints for internal loop modes in this plugin.
- Do not drift away from the canonical loop-mode contracts owned by `workflow-kit`.
