---
name: turn-gate
description: Main loop controller for `loop-kit`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Purpose

Use this skill as the main operational surface of `loop-kit`. Once invoked, `turn-gate` is a conversation-level first-class operating rule for the current session, not a checklist that applies only inside this file.

Keep one turn alive until the user explicitly asks to end it. While `user_explicit_stop` is false, every response must end in one of these states:

- loop continuation
- active question-routing
- explicit user stop handling

Result reporting, readiness reporting, and the normal final-answer channel are not turn termination. A summary-only close is invalid unless the user explicitly stops the turn or a confirmed closure is already recorded.

## Boundary

`loop-kit` exposes `turn-gate` as the single user-facing loop controller. Do not reopen direct user entrypoints for `ralph-loop`, `review-loop`, readiness gates, or other internal modes.

This skill owns turn continuity, current-phase internal mode selection, local reference loading, session continuity records, and next-flow reopening. It does not own the canonical broad workflow taxonomy or domain-specific implementation detail; those contracts stay aligned with `workflow-kit`.

## Core Loop

Treat every incoming user message as authoritative input inside the same loop-gated turn. Classify in-turn input as:

- explicit turn stop
- current-flow correction
- current-flow priority change
- next-flow priority request

For corrections or priority changes, adjust the current analysis/plan immediately and resume from the earliest safe phase. For next-flow priority requests, record the request as the highest-priority next-flow candidate and continue to the next safe handoff point.

Keep this runtime phase shape visible:

1. `analysis`: identify requested intent, requested action, current blocker, likely internal mode, and any approval boundary. Note future flow/phase candidates only when they help the current decision.
2. `plan`: set the active steps for the current flow; use `update_plan` once meaningful work begins and keep the active step current.
3. `work`: before working, choose one internal mode and read its local `references/` contract. Execute only after the mode and relevant contract are clear.
4. `verification`: verify the work before reporting it, surface residual uncertainty, and state whether later flow/phase redesign is needed.
5. `result reporting`: report the outcome, readiness state, or blocker as context for the next choice. This is not a terminal response while `user_explicit_stop` is false.
6. `question-routing reopening`: read or reconstruct the `Continuity Guard`, confirm whether terminal summary is allowed, and reopen the next flow with explicit choices unless the user explicitly stopped the turn.

Analysis and planning may include provisional future flow/phase design. Revisit that design only when new evidence, changed intent, or a revealed blocker makes redesign useful.

## Internal Mode Selection

Before `work`, select exactly one internal mode for the current phase. If the mode is unclear, use active question-routing or a narrow analysis pass before continuing.

- `deep-interview`: read `references/deep-interview.md` when requirement discovery, unclear intent, missing scope boundaries, or unresolved approval lines block progress.
- `review-loop`: read `references/review-loop.md` when review feedback, QA findings, or self-review findings should be handled as bounded material issues.
- `ralph-loop`: read `references/ralph-loop.md` when one small fix-verify-reassess cycle is the right current move.
- `autopilot`: read `references/autopilot.md` when the phase is broad end-to-end delivery from brief to verified result.
- `commit-readiness-gate`: read `references/commit-readiness-gate.md` when implementation is largely done and readiness for commit is the core question.

When multiple modes seem plausible, prefer the earliest blocker in this order: `deep-interview` -> `review-loop` -> `ralph-loop` -> `autopilot` -> `commit-readiness-gate`.

## Question Routing

Use `user-gated` question routing for clarification, choices, scope locks, mode narrowing, and next-flow decisions. Use `request_user_input` for those questions.

Explicit user, tool, platform, safety, destructive, irreversible, or external-action approval boundaries must stay user-gated. Do not simulate approval and do not route user-gated questions to subagents.

Next-flow choices should be narrow and directly connected to the result just reported. If three visible choices are already needed and a turn-end option cannot be shown, still record an explicit turn-end option in the flow record's `Next Flow Options`.

## Session Records

Maintain active turn-gated work under `.agents/sessions/{YYYYMMDD}/`.

- `000-plan.md` is the date-scoped plan artifact. It owns the day's turn-gated history, user request list, flow index, current plan, and completed-flow summaries. Update it incrementally and keep completed flow references.
- `{count-pad3}-{eng-lower-slug}.md` files are detailed flow reports, not phase notes. Use `001`, `002`, `003` style numbering and English lower-case `-` delimited slugs.
- Use `templates/plan-template.md` as the default `000-plan.md` template.
- Use `templates/flow-record-template.md` as the default flow-record template.

Each flow record must include at least: user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, and residual risk.

Update the active flow record after each completed phase. Do not wait for the flow to finish.

## Continuity Guard

Every flow record must contain a compact `Continuity Guard`. Refresh it before result reporting and before next-flow reopening.

It must state:

- whether `turn-gate` is active
- question-routing mode
- whether the user explicitly stopped the turn
- whether terminal summary is allowed
- required next action

Before result reporting, read or reconstruct this guard. If `user_explicit_stop` is false, terminal summary is not allowed and the response must proceed to next-flow reopening or active question-routing.

## Output Shape

Use the labels that fit the situation, but preserve this information shape:

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
- Do not treat result reporting, readiness reporting, or a final summary as turn termination while `user_explicit_stop` is false.
- Do not end with generic follow-up phrasing such as "let me know if you need anything else".
- Do not use summary-only closing as a valid ending shape. Bad ending shapes include: reporting only "done", giving a broad final summary without next-flow choices, or ending with a generic follow-up phrase instead of question-routing.
- Do not ask freeform textual choice questions when the active question-routing mode can carry the decision.
- Do not treat in-turn user intervention as a stop, completion, or approval-boundary pause unless the user explicitly asks to stop or creates a real approval boundary.
- Do not skip `update_plan` after meaningful work begins.
- Do not skip explicit verification between work and result reporting.
- Do not emit result reporting until the `Continuity Guard` says whether next-flow reopening is still required.
- Do not skip the next-flow question after reporting a result.
- Do not let session records lag behind phase boundaries.
- Do not omit the session-recorded turn-end option from `Next Flow Options`, even when visible question options are full.
- Do not drift away from the canonical loop-mode contracts owned by `workflow-kit`.
