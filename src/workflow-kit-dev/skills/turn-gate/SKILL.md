---
name: turn-gate
description: Loop gate for repositories where one turn must continue until the user asks to end the turn. Keep preparation, work, verification, reporting, and question-routing-based next-flow selection explicit inside the same turn.
---

# Turn Gate

## Purpose

Use this skill when a repository or working agreement requires one user turn to stay open until the user explicitly asks to end it. Once invoked, `turn-gate` is a conversation-level first-class loop gate rule for the current session, not a checklist that applies only inside this file.

While `user_explicit_stop` is false, every response must end in one of these states:

- loop continuation
- active question-routing
- explicit user stop handling

Result reporting, readiness reporting, and the normal final-answer channel are not turn termination. A summary-only close is invalid unless the user explicitly stops the turn or a confirmed closure is already source-recorded.

## Boundary

This skill owns turn continuity, phase classification, downstream owner selection for the current phase, session continuity records, and next-flow reopening.

This skill does not own requirements interviewing, planning, implementation, review handling, or commit execution. Delegate those details to the narrowest downstream workflow once the phase is clear.

## Core Loop

Treat each incoming user message as authoritative input inside the same loop-gated turn. Classify in-turn input as:

- explicit turn stop
- status or progress check
- current-flow correction
- current-flow priority change
- next-flow priority request

For status checks, answer with the current phase, blocker or progress, and next concrete action, then continue the active flow. For corrections or priority changes, adjust analysis/plan immediately and resume from the earliest safe phase. If a correction changes a target file, artifact, or state, re-read that target and do not reuse stale assumptions. For next-flow priority requests, record the request as the highest-priority next-flow candidate and continue to the next safe handoff point.

Only a clear request to end the current turn counts as explicit stop.

Keep the core flow as `preparation -> work -> verification -> reporting`. Deep-interview alignment, flow-list design, meaning resolution, state inspection, target reread, scope lock, approval-boundary checks, and planning are preparation details.

## Phase Loop

Keep the core flow separate enough to inspect while working:

1. `preparation`: decide what this flow owns, why it exists, and what must be true before work can proceed. For user-message-driven preparation, use deep-interview alignment and convert the result into a planned flow list. For existing-flow or non-user-message preparation, inspect required scope, current state, target files, stale assumptions, evidence, and execution conditions.
2. `work`: execute the user's requested work through the narrowest downstream workflow. Work can be editing, investigation, verification execution, review handling, planning, or another task shape selected during preparation.
3. `verification`: verify this flow's work before reporting it. For file edits, check intended file changes and relevant type/test/lint/build/parse checks. For investigation or reasoning work, criticize the logic from multiple angles and check contrary evidence.
4. `reporting`: report the completed outcome, readiness state, or blocker as context for the next flow. This is not a terminal response while `user_explicit_stop` is false.

After reporting, read or reconstruct the `Continuity Guard`, confirm whether terminal summary is allowed, and open explicit next-flow choices through user-gated question routing unless the user explicitly stopped the turn. This next-flow response is the continuation surface for the same gated turn, not an additional core phase.

Future flow/phase design is provisional. Revisit it only when new evidence, changed intent, or a revealed blocker makes redesign useful.

## Meaning Resolution

Meaning resolution happens before routing, planning, or editing. Its job is to verify that the interpreted operation matches the user's intended operation.

Use it when a user instruction contains an overloaded operation, structural target, or contextual reference whose interpretation would change the work. Common examples include merging a skill versus absorbing behavior into a phase, removing a user-facing surface versus removing a routing entry, changing a spec contract versus changing only runtime wording, or editing "the section below" when multiple nearby sections could match.

Treat provenance, source URLs, and user-intent/spec-intent blocks as meaningful targets. If "source", "original", "intent", or "below that" could point to either a provenance note, a user intent block, or the normative spec body, lock that target before editing.

During preparation, keep these fields explicit when relevant:

- `Literal wording`: the user's exact ambiguous wording
- `Interpreted operation`: the operation you are about to act on
- `Operation target`: the affected unit, such as skill, spec, plugin, phase, routing rule, release surface, or commit scope
- `Alternate interpretations`: other plausible meanings that would change the work
- `Impact of ambiguity`: files, behavior, deletion, migration, or commit scope that would differ

If alternate interpretations would materially change the work, open user-gated question routing before selecting a downstream owner. The question should lock the ambiguous structure directly.

Do not route this to `deep-interview` merely because meaning is unclear. `deep-interview` is for broader requirement discovery; meaning resolution is the current-flow check that locks the operation or target of the user's instruction. If target ambiguity would change the planned flow list, lock the target before flow-list design.

## Question Routing

Use `user-gated` question routing by default. Use the user-input question tool for clarification, choices, scope locks, next-flow decisions, and any explicit user/tool/platform/safety/destructive/irreversible/external-action approval boundary.

Before destructive, irreversible, or external actions, inspect the current state closely enough to state the exact target and risk. Keep approval user-gated. If the user asks to use subagents to keep moving, hand off autonomous question routing to `turn-gate-self-drive`; approval, destructive, irreversible, external-action, and safety decisions remain with the user and should be recorded as such.

Next-flow choices should be narrow and directly connected to the result just reported. Do not reopen broad framing when the next phase is already clear. If three visible choices are already needed and a turn-end option cannot be shown, still record an explicit turn-end option in the flow record's `Next Flow Options`.

## Session Records

Maintain turn-gated work under `.agents/sessions/{YYYYMMDD}/`.

- `000-plan.md` is the date-scoped plan artifact. It owns the day's turn-gated history, user request list, flow index, current plan, and completed-flow summaries. Update it incrementally and keep completed flow references.
- `{count-pad3}-{eng-lower-slug}.md` files are detailed flow reports, not phase notes. Use `001`, `002`, `003` style numbering and English lower-case `-` delimited slugs.
- Prefer `templates/plan-template.md` for `000-plan.md`.
- Prefer `templates/flow-record-template.md` for flow records.

Each flow record must include at least: user request message, task, flow scope, current mode, question-routing mode, current core phase, continuity guard, preparation source/result, planned flow list, work boundary, verification expectation, work, verification, report, next-flow options, and residual risk.

Update the active flow record after each completed phase. Do not wait for the flow to finish.

## Continuity Guard

Every flow record must contain a compact `Continuity Guard`. Refresh it before result reporting and before next-flow reopening.

It must state:

- whether `turn-gate` is active
- question-routing mode
- whether the user explicitly stopped the turn
- whether terminal summary is allowed
- required next action
- last refreshed phase
- pending or superseded question state
- verification status

Before result reporting, read or reconstruct this guard. If `user_explicit_stop` is false, terminal summary is not allowed and the response must proceed to next-flow reopening or active question-routing. A stale source-less closure record does not allow terminal summary.

## Downstream Owner Signals

Choose the narrowest downstream workflow that owns the current phase:

- `deep-interview` when requirement discovery, intent ambiguity, scope boundaries, or approval lines are the blocker
- `review-loop` when the input is review or QA findings and the work must stay bounded to material issues
- `ralph-loop` when a small fix-verify-reassess cycle is the right current move
- `autopilot` when the current phase is broad end-to-end delivery from brief to verified result
- `commit-readiness-gate` when the intended change unit is nearly done and readiness is the core question
- a commit execution workflow only when the user explicitly asks to commit and the intended change unit has passed scope, staged/final status, and readiness checks
- another specialist workflow only after the current phase is clear

Keep the meta loop here. Keep phase-specific detail in the downstream workflow.

## Response Contract

Use the labels that fit the situation, but preserve this information shape:

- `Requested intent`
- `Requested action`
- `Chosen downstream owner`
- `Preparation source`
- `Preparation result`
- `Planned flow list`
- `Question-routing mode`
- `Current core phase`
- `Work boundary`
- `Verification expectation`
- `Work`
- `Verification`
- `Report`
- `Continuity guard`
- `Question-routing prompt`
- `Next-flow choices`
- `Loop state`
- `Residual risk`

## Guardrails

- Do not treat result reporting, readiness reporting, or a final summary as turn termination while `user_explicit_stop` is false.
- Do not decide on behalf of the user that the turn should terminate.
- Do not end with generic follow-up phrasing.
- Do not use summary-only closing as a valid ending shape.
- Do not skip explicit verification between work and result reporting.
- Do not turn ambiguous user wording into a concrete action until meaning resolution has ruled out materially different interpretations.
- Do not skip the next-flow question because the next phase seems obvious.
- Do not let session records lag behind phase boundaries.
- Do not ask the next-flow question without giving the user explicit choices.
- Do not omit the session-recorded turn-end option from `Next Flow Options`, even when the visible question options are full.
- Do not route user-gated questions to subagents from this skill.
- Do not simulate user approval where the runtime or tool policy requires explicit approval.
- Do not treat a readiness request as commit approval, and do not treat commit approval as permission to include unrelated changes.
- Do not treat in-turn user intervention as completion, stop, or an approval-boundary pause unless the user explicitly asks to end the turn or creates a real approval boundary.
- Do not treat the user's next-flow response as a new independent turn when the loop gate is still active.
- Do not treat temporary blocking states as permission to close the turn.
- Do not let this skill absorb domain execution, planning, or review detail.
