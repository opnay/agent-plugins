---
name: turn-gate
description: Main loop controller for `loop-kit`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Important

- Once invoked, `turn-gate` is a session-level first-class operating rule for the response lifecycle, not an internal checklist.
- While `user_explicit_stop=false`, do not close with a terminal summary. Result reporting must continue into next-flow reopening, loop continuation, or active question-routing.
- Every response must end in one of these states: loop continuation, active question-routing, or explicit user stop handling.
- Keep the core flow as `preparation -> work -> verification -> reporting`. Treat deep-interview, flow-list design, meaning resolution, state inspection, target reread, scope lock, and approval-boundary checks as preparation details.
- If the user only activates `turn-gate`, do not infer work. Activate the gate, keep `user_explicit_stop=false`, update or create the session record, and open a scope or next-flow choice.
- Use `request_user_input` when available for next-flow decisions, scope locks, mode narrowing, meaning clarification, blockers, and approval choices.
- Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and the active flow record, including `Continuity Guard`, preparation source/result, current core phase, work boundary, verification status, and `Next Flow Options`.
- After `work`, request clean-context verification before result reporting. The verifier receives a bounded packet, not a full-history fork.
- Failed, blocked, unavailable, or insufficient verification is not a pass.
- A terminal summary is allowed only when explicit stop is source-recorded in the active flow record with the closure source message.

## Purpose

Use this skill as the main operational surface of `loop-kit`. Keep one user turn open until the user explicitly asks to end it, choose the internal loop mode for current-phase work, verify the work, report the result as next-flow context, and reopen the next flow.

## Core Loop

Follow this order unless an in-turn correction requires returning to the earliest safe phase:

1. `preparation`
2. `work`
3. `verification`
4. `reporting`

Activation, incoming message classification, next-flow reopening, and explicit stop handling are lifecycle guards around the core flow. Next-flow question-routing is the continuation surface after reporting, not an additional core phase.

## Incoming Messages

Treat every new user message as authoritative input inside the same gated turn. Classify it as one of:

- explicit turn stop
- status or progress check
- current-flow correction
- current-flow priority change
- next-flow priority request

For a status check, report current phase, blocker or progress, and next concrete action, then continue the active flow. For corrections or priority changes, update analysis/plan immediately and resume from the earliest safe phase. If a correction changes a target file, artifact, or state, re-read that target and do not reuse stale assumptions. For next-flow priority requests, record the request as the highest-priority next-flow candidate and continue to the next safe handoff point.

Only a clear request to end the current turn counts as explicit stop.

## Preparation

Preparation decides what this flow owns, why it exists, and what must be true before work can proceed.

- User-message-driven preparation uses deep-interview alignment to lock or infer intent, scope, non-goals, success criteria, approval boundary, and verification signal. Convert that result into a planned flow list.
- Existing-flow or non-user-message preparation prepares an already selected flow. Inspect required change scope, current state, target files or artifacts, stale assumptions, available evidence, and execution conditions.
- If operation or target ambiguity can change the flow list or work result, lock it through meaning resolution before flow-list design, mode selection, or editing.
- Identify requested intent, requested action, current blocker, likely internal mode, approval boundary, preparation result, planned flow list, work boundary, and verification expectation.
- Use the planning tool once meaningful work begins. For multi-flow work, keep `000-plan.md` as the flow sequence and put detailed task steps in the active `001+` flow record.

## Meaning Resolution

Run meaning resolution before routing, planning, or editing when wording can map to more than one operation or target.

- Ambiguous structural terms include merge, absorb, remove, delete, split, route, phase, surface, skill, spec, and contract.
- Ambiguous references include this, that, above, below, current one, `그`, `그 밑`, `그건`, and `그거` when nearby targets differ.
- Treat provenance notes, source URLs, user-spec intent blocks, and spec intent text as possible work targets.
- Ask a narrow question that locks the structure directly. Prefer `request_user_input` with bounded choices.
- Record literal wording, interpreted operation, operation target, alternate interpretations, and ambiguity impact in the flow record.
- Meaning resolution is not approval. If a locked target still involves destructive, irreversible, external, commit, push, PR, or publish work, request explicit user approval separately.

## Work

Work is where the requested task is actually performed. It may be file edits, investigation, verification execution, review handling, planning artifacts, or another task shape selected during preparation.

Before work begins, select exactly one internal mode for current-phase work. If the mode is unclear, use narrow analysis or active question-routing first. Read and apply the selected local reference from `skills/turn-gate/references/`:

- `deep-interview.md`: requirement discovery, unclear intent, missing scope boundaries, unresolved approval lines
- `review-loop.md`: review feedback, QA finding, or self-review finding is the material issue
- `ralph-loop.md`: one small fix-verify-reassess cycle is the right unit
- `autopilot.md`: broad end-to-end delivery is the current phase
- `commit-readiness-gate.md`: the change unit is nearly complete and readiness judgment is the current phase

If several modes fit, prefer the earliest blocker in this order: `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `commit-readiness-gate`.

Commit execution, push, PR, publish, and similar external actions are not internal modes. Treat them as user-gated handoff workflows.

## Approval Boundaries

- Do not execute destructive, irreversible, external, publish, push, PR, or commit actions without explicit user approval.
- Before approval-sensitive work, present exact target, expected effect, risk, rollback or recovery possibility, and included/excluded scope.
- Do not use subagent output, inferred intent, nearby wording, or readiness requests as approval.
- Commit approval requires staged/final status, intended diff, unrelated-change exclusion, and commit message scope before handoff to the commit workflow.

## Verification

- Verify after `work` and before result reporting.
- Verification checks this flow's work, not the whole turn by default.
- For file edits, check intended file changes and relevant type, test, lint, build, or parse checks.
- For investigation or reasoning work, criticize the logic from multiple angles and check contrary evidence.
- Clean context means the verifier receives a bounded packet: verifier id or request id, target, expected user intent, changed files or artifacts, commands/checks to run, pass/fail criteria, no-edit permission, and stop condition.
- Do not substitute same-context self-review as passing clean-context verification.
- Integrate verification as exactly one of `pass`, `fail`, `blocked`, or `insufficient`.
- If verification is `fail` or `insufficient`, return to the earliest safe phase before result reporting.
- If verification is `blocked` or unavailable, open a user-gated blocker through question-routing. Do not call it a pass.

## Reporting

Reporting is context for the next flow decision, not terminal closure.

- Include what was prepared, what work was done, verification result, residual uncertainty, blocker if any, material routing judgment calls, and the next concrete decision.
- Refresh the active flow record and read/update `Continuity Guard` before reporting.
- If planned flows are exhausted, use `request_user_input` when available to ask the user for the next flow or task.
- A stale `terminal summary allowed: yes` or source-less confirmed closure cannot justify terminal close.

## Question Routing

- Use user-gated question-routing for clarifications, choices, scope locks, mode narrowing, approval boundaries, blockers, and next-flow decisions.
- Use `request_user_input` when available and structural choices can be offered.
- Next-flow choices after result reporting must be narrow and directly connected to the reported result.
- If `request_user_input` is unavailable, state that the tool is unavailable, list the open choices, and record the required next action. The turn remains in active question-routing.
- If visible choices cannot include a turn-end option, still tell the user they can explicitly stop the turn.
- Always record an explicit turn-end option in the flow record `Next Flow Options`.

## Session Records

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for active turn-gated work.

- Use `skills/turn-gate/templates/plan-template.md` for `000-plan.md`.
- Use `skills/turn-gate/templates/flow-record-template.md` for flow records.
- `000-plan.md` owns date-level history, user requests, flow index, planned flow sequence, transition criteria, and completed flow summaries.
- Each `001+` file owns one flow's details: user request message, task, flow scope, current mode, question-routing mode, current core phase, `Continuity Guard`, preparation source/result, planned flow list, work boundary, verification expectation, work, verification, report, next-flow options, and residual risk.
- Update the flow record incrementally after preparation, work, verification, and reporting.

### Continuity Guard

Before result reporting and next-flow reopening, update the active flow record's `Continuity Guard` with:

- `turn-gate` activation state
- question-routing mode
- `user_explicit_stop`
- whether terminal summary is allowed
- required next action
- last refreshed phase
- pending or superseded question state
- verification status: `not-started`, `requested`, `pass`, `fail`, `blocked`, or `insufficient`

Record confirmed closure only after explicit stop. Include closure source message and closure recorded phase. If records are inaccessible, report that as a blocker and do not reconstruct silent terminal-close permission. When stale terminal-summary permission is detected, refresh the guard to `user_explicit_stop=false`, `terminal summary allowed=no`, and record that the prior closure state was stale or source-less.

## Stop Handling

- Do not send a terminal summary unless the current incoming message clearly asks to end the turn or a valid confirmed closure with source message matches the current state.
- On explicit stop, write confirmed closure, closure source message, and closure recorded phase into the active flow record.
- After source-recorded closure, terminal summary is allowed and next-flow reopening is not required.
