---
name: turn-gate
description: Loop gate for repositories where one turn must continue until the user asks to end the turn. Keep analysis, plan, work, verification, result reporting, and question-routing-based next-flow selection explicit inside the same turn.
---

# Turn Gate

## Purpose

Use this skill when a repository or working agreement requires one user turn to stay open until the user explicitly asks to end it. Once invoked, `turn-gate` is a conversation-level first-class loop gate rule for the current session, not a checklist that applies only inside this file.

While `user_explicit_stop` is false, every response must end in one of these states:

- loop continuation
- active question-routing
- explicit user stop handling

Result reporting, readiness reporting, and the normal final-answer channel are not turn termination. A summary-only close is invalid unless the user explicitly stops the turn or a confirmed closure is already recorded.

## Boundary

This skill owns turn continuity, phase classification, downstream owner selection for the current phase, session continuity records, and next-flow reopening.

This skill does not own requirements interviewing, planning, implementation, review handling, or commit execution. Delegate those details to the narrowest downstream workflow once the phase is clear.

## Core Loop

Treat each incoming user message as authoritative input inside the same loop-gated turn. Classify in-turn input as:

- explicit turn stop
- current-flow correction
- current-flow priority change
- next-flow priority request

For corrections or priority changes, adjust the current analysis/plan immediately and resume from the earliest safe phase. For next-flow priority requests, record the request as the highest-priority next-flow candidate and continue to the next safe handoff point.

## Phase Loop

Keep the loop phases separate enough to inspect while working:

1. `analysis`: structure the user message into requested intent and requested action. Identify current blockers, approval boundaries, whether a downstream workflow owns the next phase, and any useful future flow/phase candidates.
2. `plan`: prepare the current-flow plan from the analysis. Include the selected downstream owner, the next concrete steps, required records, verification target, and any provisional future flow/phase design.
3. `work`: execute the plan through the narrowest downstream workflow that owns the current phase. Keep the meta loop here and leave phase-specific detail to the downstream workflow.
4. `verification`: verify the work before reporting it. State what was checked, what failed or remains uncertain, and whether new evidence means the later flow/phase design should be revised.
5. `result reporting`: report the completed outcome, readiness state, or blocker as context for the next choice. This is not a terminal response while `user_explicit_stop` is false.
6. `next-flow question-routing`: read or reconstruct the `Continuity Guard`, confirm whether terminal summary is allowed, and open explicit next-flow choices through user-gated question routing unless the user explicitly stopped the turn.

Future flow/phase design is provisional. Revisit it only when new evidence, changed intent, or a revealed blocker makes redesign useful.

## Question Routing

Use `user-gated` question routing by default. Use the user-input question tool for clarification, choices, scope locks, next-flow decisions, and any explicit user/tool/platform/safety/destructive/irreversible/external-action approval boundary.

Next-flow choices should be narrow and directly connected to the result just reported. Do not reopen broad framing when the next phase is already clear. If three visible choices are already needed and a turn-end option cannot be shown, still record an explicit turn-end option in the flow record's `Next Flow Options`.

## Session Records

Maintain turn-gated work under `.agents/sessions/{YYYYMMDD}/`.

- `000-plan.md` is the date-scoped plan artifact. It owns the day's turn-gated history, user request list, flow index, current plan, and completed-flow summaries. Update it incrementally and keep completed flow references.
- `{count-pad3}-{eng-lower-slug}.md` files are detailed flow reports, not phase notes. Use `001`, `002`, `003` style numbering and English lower-case `-` delimited slugs.
- Prefer `templates/plan-template.md` for `000-plan.md`.
- Prefer `templates/flow-record-template.md` for flow records.

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

## Downstream Owner Signals

Choose the narrowest downstream workflow that owns the current phase:

- `deep-interview` when requirement discovery, intent ambiguity, scope boundaries, or approval lines are the blocker
- `review-loop` when the input is review or QA findings and the work must stay bounded to material issues
- `ralph-loop` when a small fix-verify-reassess cycle is the right current move
- `autopilot` when the current phase is broad end-to-end delivery from brief to verified result
- `commit-readiness-gate` when the intended change unit is nearly done and readiness is the core question
- another specialist workflow only after the current phase is clear

Keep the meta loop here. Keep phase-specific detail in the downstream workflow.

## Response Contract

Use the labels that fit the situation, but preserve this information shape:

- `Analysis`
- `Requested intent`
- `Requested action`
- `Chosen downstream owner`
- `Question-routing mode`
- `Plan`
- `Work`
- `Verification`
- `Result report`
- `Continuity guard`
- `Question-routing prompt`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Guardrails

- Do not treat result reporting, readiness reporting, or a final summary as turn termination while `user_explicit_stop` is false.
- Do not decide on behalf of the user that the turn should terminate.
- Do not end with generic follow-up phrasing such as "let me know if you need anything else".
- Do not use summary-only closing as a valid ending shape. Bad ending shapes include: reporting only "done", giving a broad final summary without next-flow choices, or ending with a generic follow-up phrase instead of question-routing.
- Do not skip explicit verification between work and result reporting.
- Do not skip the next-flow question because the next phase seems obvious.
- Do not let session records lag behind phase boundaries.
- Do not ask the next-flow question without giving the user explicit choices.
- Do not omit the session-recorded turn-end option from `Next Flow Options`, even when the visible question options are full.
- Do not route user-gated questions to subagents from this skill.
- Do not simulate user approval where the runtime or tool policy requires explicit approval.
- Do not treat in-turn user intervention as completion, stop, or an approval-boundary pause unless the user explicitly asks to end the turn or creates a real approval boundary.
- Do not treat the user's next-flow response as a new independent turn when the loop gate is still active.
- Do not treat temporary blocking states as permission to close the turn.
- Do not let this skill absorb domain execution, planning, or review detail.
- Do not confuse turn continuity with endless conversation.
