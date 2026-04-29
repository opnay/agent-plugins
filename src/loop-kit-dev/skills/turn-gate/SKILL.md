---
name: turn-gate
description: Main loop controller for `loop-kit-dev`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Important

- Once invoked, `turn-gate` is a session-level first-class operating rule for the assistant response lifecycle, not an internal checklist.
- While `user_explicit_stop` is false, do not close with a terminal summary; result reporting must continue into next-flow reopening or active question-routing.
- If the user only asks to activate `turn-gate`, activate it, update or create the session record, keep `user_explicit_stop=false`, and open a scope or next-flow question instead of choosing work mode prematurely.
- Every response must end in one of these states:
  - loop continuation
  - active question-routing
  - explicit user stop handling
- Use `request_user_input` for next-flow decisions, scope locks, mode narrowing, and user-gated approval boundaries.
- Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and the active flow record, including `Continuity Guard` and `Next Flow Options`.
- On explicit user stop, record confirmed closure before sending a terminal summary and do not reopen next-flow choices.

## Purpose

Use this skill as the main operational surface of `loop-kit-dev`. Keep one turn alive until the user explicitly asks to end it.

While `user_explicit_stop` is false, valid response endings are:

- loop continuation
- active question-routing
- explicit user stop handling

Result reporting, readiness reporting, and the normal final-answer channel are not turn termination. A summary-only close is invalid unless the user explicitly stops the turn or a confirmed closure is already recorded.

Activation-only requests are still meaningful work: mark `turn-gate` active, initialize or refresh the session records, and ask the user to choose the first scope or next flow with user-gated routing.

## Boundary

`loop-kit-dev` exposes `turn-gate` as the single user-facing loop controller. Do not reopen direct user entrypoints for `ralph-loop`, `review-loop`, readiness gates, or other internal modes.

This skill owns turn continuity, current-phase internal mode selection, local reference loading, session continuity records, and next-flow reopening. It does not own the canonical broad workflow taxonomy or domain-specific implementation detail; those contracts stay aligned with `workflow-kit`.

## Core Loop

Treat every incoming user message as authoritative input inside the same loop-gated turn. Classify in-turn input as:

- explicit turn stop
- status or progress check
- current-flow correction
- current-flow priority change
- next-flow priority request

For status or progress checks, answer with the current phase, blocker or progress, and next concrete action, then continue the active flow. Do not treat a status check as a stop or as permission to close the turn.

For corrections or priority changes, re-read affected files, records, or state, reconcile them against the correction, adjust the current analysis/plan immediately, and resume from the earliest safe phase. For next-flow priority requests, record the request as the highest-priority next-flow candidate and continue to the next safe handoff point.

For explicit turn stop, update the active flow record and `Continuity Guard` with confirmed closure, set terminal summary allowed, then give the closure summary without reopening next-flow choices.

Keep this runtime phase shape visible:

1. `analysis`: identify requested intent, requested action, current blocker, likely internal mode, whether meaning resolution is needed, and any approval boundary. Note future flow/phase candidates only when they help the current decision.
2. `plan`: set the active steps for the current flow; use `update_plan` once meaningful work begins and keep the active step current.
3. `work`: before working, choose one internal mode and read its local `references/` contract. Execute only after the mode and relevant contract are clear.
4. `verification`: verify the work before reporting it, surface residual uncertainty, and state whether later flow/phase redesign is needed.
5. `result reporting`: report the outcome, readiness state, or blocker as context for the next choice. This is not a terminal response while `user_explicit_stop` is false.
6. `question-routing reopening`: read the `Continuity Guard`, reconstruct only if unavailable, confirm whether terminal summary is allowed, and reopen the next flow with visible `request_user_input` choices unless the user explicitly stopped the turn.

Analysis and planning may include provisional future flow/phase design. Revisit that design only when new evidence, changed intent, or a revealed blocker makes redesign useful.

## Meaning Resolution

Before selecting an internal mode or acting on a user correction, check whether the user's operation wording or contextual reference can mean more than one thing. Terms such as merge, absorb, remove, delete, split, route, phase, surface, skill, spec, or contract often change meaning depending on whether the target is a file, workflow phase, routing rule, plugin surface, or destructive action. Referential wording such as this, that, below, above, current one, `그`, `그 밑`, `그건`, or `그거` also needs locking when multiple nearby targets are plausible.

Treat provenance, source URLs, and user-intent/spec-intent blocks as meaningful targets, not disposable conversation context. If "source", "original", "intent", or "below that" could point to either a provenance note, a user intent block, or the normative spec body, lock that target before editing.

When the interpretation would change the work, do not pick one meaning silently. Record the literal wording, the interpreted operation, the operation target, plausible alternate interpretations, and the impact of ambiguity. Then use user-gated question routing to lock the meaning before work starts.

This is not `deep-interview`. It is current-flow meaning resolution for the user's instruction itself. Ask a narrow structural question, such as whether "merge" means combining skill/spec surfaces or absorbing behavior into a `turn-gate` phase.

## Internal Mode Selection

Before `work`, select exactly one internal mode for the current phase. If the mode is unclear, use active question-routing or a narrow analysis pass before continuing.

- `deep-interview`: read `references/deep-interview.md` when requirement discovery, unclear intent, missing scope boundaries, or unresolved approval lines block progress.
- `review-loop`: read `references/review-loop.md` when review feedback, QA findings, or self-review findings should be handled as bounded material issues.
- `ralph-loop`: read `references/ralph-loop.md` when one small fix-verify-reassess cycle is the right current move.
- `autopilot`: read `references/autopilot.md` when the phase is broad end-to-end delivery from brief to verified result.
- `commit-readiness-gate`: read `references/commit-readiness-gate.md` when implementation is largely done and readiness for commit is the core question.
- commit execution workflow: when the user explicitly asks to commit and the intended change unit has passed scope, staged/final status, and readiness checks, hand off to the appropriate commit workflow. `turn-gate` keeps the loop; it does not own commit execution details.
- external publish workflow: when the user asks to publish, push, open a pull request, or otherwise affect an external system, first verify branch, remote, scope, and risk, then get user approval and hand off to the GitHub or matching external-action workflow. `turn-gate` keeps the loop; it does not own the external execution details.

When multiple modes seem plausible, prefer the earliest blocker in this order: `deep-interview` -> `review-loop` -> `ralph-loop` -> `autopilot` -> `commit-readiness-gate`.

## Question Routing

Use `user-gated` question routing for clarification, choices, scope locks, mode narrowing, and next-flow decisions. Use `request_user_input` for those questions.

Explicit user, tool, platform, safety, destructive, irreversible, or external-action approval boundaries must stay user-gated. Do not simulate approval and do not route user-gated questions to subagents.

Before destructive, irreversible, or external actions, inspect the current state closely enough to state the exact target and risk. Treat wording like delete, remove, wipe, reset, overwrite, discard, publish, push, or open PR as approval-sensitive when it can change local state irreversibly or affect external systems. If the user asks to use subagents to keep moving, hand off autonomous question routing to `turn-gate-self-drive`; approval, destructive, irreversible, external-action, and safety decisions remain with the user and should be recorded as such.

Next-flow choices should be narrow, visible, tool-backed, and directly connected to the result just reported. Use `request_user_input` whenever available. If three visible choices are already needed and a turn-end option cannot be shown, still record an explicit turn-end option in the flow record's `Next Flow Options`.

## Session Records

Maintain active turn-gated work under `.agents/sessions/{YYYYMMDD}/`.

- `000-plan.md` is the date-scoped plan artifact. It owns the day's turn-gated history, user request list, flow index, current plan, and completed-flow summaries. Update it incrementally and keep completed flow references.
- `{count-pad3}-{eng-lower-slug}.md` files are detailed flow reports, not phase notes. Use `001`, `002`, `003` style numbering and English lower-case `-` delimited slugs.
- Use `templates/plan-template.md` as the default `000-plan.md` template.
- Use `templates/flow-record-template.md` as the default flow-record template.

Each flow record must include at least: user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, and residual risk.

Update the active flow record after each completed phase. Also update it after status/progress checks when the phase, blocker, progress, or required next action changed. Do not wait for the flow to finish.

## Continuity Guard

Every flow record must contain a compact `Continuity Guard`. Refresh it before result reporting and before next-flow reopening.

It must state:

- whether `turn-gate` is active
- question-routing mode
- whether the user explicitly stopped the turn
- whether terminal summary is allowed
- required next action

Before result reporting, read the recorded guard; reconstruct it only if it is missing or inaccessible. If reconstructed, write it back to the flow record as soon as possible. If `user_explicit_stop` is false, terminal summary is not allowed and the response must proceed to next-flow reopening or active question-routing. If explicit stop is confirmed, record closure and terminal summary allowance before closing.

## Output Shape

Use the labels that fit the situation, but preserve this information shape:

- `Analysis`
- `Meaning resolution` when relevant
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
- Do not collapse overloaded user wording into one concrete action when different interpretations would change files, phases, routing, or commit scope.
- Do not treat a readiness request as commit approval, and do not treat commit approval as permission to include unrelated changes.
- Do not skip `update_plan` after meaningful work begins.
- Do not skip explicit verification between work and result reporting.
- Do not emit result reporting until the `Continuity Guard` says whether next-flow reopening is still required.
- Do not skip the next-flow question after reporting a result.
- Do not let session records lag behind phase boundaries.
- Do not let status/progress answers change phase, blocker, or next action without updating the active flow record.
- Do not omit the session-recorded turn-end option from `Next Flow Options`, even when visible question options are full.
- Do not perform publish, push, PR, destructive, irreversible, or external actions without target/risk inspection and user-gated approval.
- Do not drift away from the canonical loop-mode contracts owned by `workflow-kit`.
