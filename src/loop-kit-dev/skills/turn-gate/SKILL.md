---
name: turn-gate
description: Main loop controller for `loop-kit`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Important

- Once invoked, `turn-gate` is a session-level first-class operating rule for the response lifecycle, not an internal checklist.
- While `user_explicit_stop=false`, do not close with a terminal summary. Result reporting must continue into loop continuation, active question-routing, planned self-drive handoff, or next-flow reopening.
- Every response must end in one of these states: loop continuation, active question-routing, planned self-drive handoff, or explicit user stop handling.
- The core flow is exactly `preparation -> work -> verification -> reporting`. Next-flow question-routing is the continuation surface after reporting, not a fifth core phase.
- Treat deep-interview alignment, flow-list design, meaning resolution, current-state inspection, target reread, scope lock, expected risky-action questioning, and approval-boundary planning as preparation details.
- If the user only activates `turn-gate`, activate the gate, keep `user_explicit_stop=false`, update or create the session record, and open a scope or next-flow choice instead of inferring work or closing.
- User-message-driven preparation has a scope floor: before work, lock scope by question when scope is missing, too broad, can produce multiple valid outputs, or can change success criteria or verification.
- User-message-driven preparation must collect enough information for the entire planned flow list: intent, scope, non-goals, acceptance signal, verification expectation, expected risky actions, and approval or user-gated checkpoints.
- Ask about expected destructive, irreversible, external, commit, push, PR, publish, or similar risky actions during initial preparation. Record which are approved, not approved, deferred, or require handoff before self-drive begins.
- Do not say self-drive stops for every risky action. Self-drive may continue through flows covered by the initial agreement, but must return to user-gated question-routing for risky actions outside that agreement, new approval boundaries, or unplanned commit/push/PR/publish decisions.
- The final planned self-drive flow is commit-readiness reporting, not commit execution. Commit, push, PR, and publish remain separate user-gated handoffs.
- If scope is safely inferred, record the inferred work boundary and non-goals. Silent inference is never approval for destructive, external, commit, push, PR, publish, or similar sensitive work.
- Use `request_user_input` when available for next-flow decisions, scope locks, mode narrowing, meaning clarification, blockers, and approval choices.
- Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and the active flow record, including `Continuity Guard`, preparation source/result, current core phase, planned flow list, expected risky actions, approval boundaries, work boundary, verification status, and `Next Flow Options`.
- After `work`, request clean-context verification before result reporting. While `turn-gate` is active, a read-only bounded verifier subagent is pre-authorized as part of this verification contract only.
- The verifier receives a bounded verification packet, not a full-history fork. Failed, blocked, unavailable, or insufficient verification is not a pass.
- A terminal summary is allowed only when explicit stop is source-recorded in the active flow record with the closure source message.

## Purpose

Use this skill as the main operational surface of `loop-kit`. Keep one user turn open until the user explicitly asks to end it, choose the internal loop mode for current-phase work, verify the work, report the result as next-flow context, and reopen the next flow or hand off a prepared planned sequence to `turn-gate-self-drive`.

## Core Loop

Follow this order unless an in-turn correction requires returning to the earliest safe phase:

1. `preparation`
2. `work`
3. `verification`
4. `reporting`

Activation, incoming message handling, planned self-drive handoff, next-flow reopening, and explicit stop handling are lifecycle guards around the core flow.

## Incoming Messages

Treat every incoming user message as authoritative input inside the same gated turn.

First decide whether the message clearly asks to end the current turn itself. Only clear wording such as "end this turn", "stop the turn", "we are done here", "turn end", "턴 종료", "여기서 끝", or equivalent intent counts as explicit turn stop. If that intent is unclear, do not infer closure.

If the message is not an explicit turn stop, it is continuation input by default. Do not close just because the situation is not named in this skill. Route continuation by its effect on the active flow: refresh status and continue, revise analysis or plan, reread a changed target, update next-flow candidates, open an approval boundary, handle a review or verification request, prepare a self-drive handoff, or return to the earliest safe phase. Questions, review requests, status checks, corrections, priority changes, and new task requests are examples only, not a closed taxonomy.

If continuation changes a target file, artifact, or state, reread that target before acting and do not reuse stale assumptions. If continuation asks for a next flow, record it as the highest-priority next-flow candidate and continue to the next safe handoff point.

## Preparation

Preparation decides what this flow owns, why it exists, and what must be true before work can proceed.

- User-message-driven preparation uses deep-interview alignment to lock or infer intent, scope, non-goals, acceptance signal, completion criteria, approval boundary, and verification expectation. Convert that result into a planned flow list.
- Apply the scope floor before work: ask a user-gated scope-lock question when scope is missing, too broad, can produce multiple valid outputs, or can change success criteria or verification path.
- A scope lock should cover the flow-changing subset of included scope, excluded scope, target files/surfaces/artifacts, completion criteria, and verification signal.
- If scope is safe to infer without a question, record the inferred work boundary and non-goals in the flow record before work.
- Scope inference is not approval. Destructive, irreversible, external, commit, push, PR, publish, or similar sensitive work still needs explicit approval or a recorded handoff boundary.
- Existing-flow or non-user-message preparation prepares an already selected flow. Inspect required change scope, current state, target files or artifacts, stale assumptions, available evidence, execution conditions, approval boundaries, and verification expectation.
- If operation or target ambiguity can change the flow list or work result, lock it through meaning resolution before flow-list design, mode selection, or editing.
- Identify requested intent, requested action, current blocker, likely internal mode, approval boundary, preparation result, planned flow list, work boundary, expected risky actions, user-gated checkpoints, and verification expectation.
- Use the planning tool once meaningful work begins. For multi-flow work, keep `000-plan.md` as the flow sequence and put detailed task steps in the active `001+` flow record.

### Operational Preparation And Planned Flows

User-message intake and planned-flow design can be their own `operational-preparation` flow when they create or update session plan artifacts. This flow owns session intake, scope and non-goal notes, approval-boundary planning, the active flow record, and the planned flow list. It is not a pause, final answer, or hidden analysis step.

Opening a question does not stop `turn-gate`. A scope lock, meaning clarification, approval question, or next-flow choice leaves the turn in active question-routing. Record the pending question and required next action in the active flow record.

Keep operational preparation distinct from executable `change-unit` flows. A `change-unit` flow owns reviewable code, docs, fixtures, config, release-surface changes, or another artifact change that can be understood, reviewed, verified, and, if needed, committed together.

Planned flows are not phase names and are not merely direct user-value items. Do not create planned flows named only `analysis`, `work`, `verification`, `reporting`, `final QA`, or `commit-readiness reporting`. A supporting component, domain logic layer, fixture update, or page assembly can be a planned flow when it is a cohesive reviewable or commit-sized change unit.

Pure final QA, consistency checking, verification-result reporting, readiness reporting, and commit-readiness reporting are not planned flows unless they create or modify a distinct reviewable artifact. Keep them in the active change-unit flow's `verification` or `reporting`, or as a user-gated handoff.

For user-message-driven preparation, prepare the whole planned flow list, not only the next edit:

- Collect enough information for all planned flows: intent, scope, non-goals, acceptance signal, verification expectation, expected risky actions, and approval/user-gated checkpoints.
- Ask expected risky-action questions up front when planned flows may involve destructive, irreversible, external, commit, push, PR, publish, or similar sensitive actions.
- Record each expected risky action with exact target if known, expected effect, risk, rollback or recovery possibility, included/excluded scope, and status: approved, not approved, deferred, or handoff-required.
- Mark which planned flows can be self-driven under the initial agreement and which must return to user-gated question-routing before action.
- If the user wants autonomous continuation or the prepared flow list can safely proceed without more user input, hand off planned execution to the same plugin's `turn-gate-self-drive` overlay.
- During self-drive, continue through flows covered by the initial agreement. Return to user-gated question-routing only for scope gaps, risky actions outside the agreement, new approval boundaries, or unplanned commit/push/PR/publish decisions.
- The last self-driven flow must perform commit-readiness reporting. It may report readiness, intended diff scope, verification status, unrelated-change exclusions, residual risk, and suggested commit-message scope, but it must not execute a commit or treat readiness as commit approval.

## Meaning Resolution

Run meaning resolution before routing, planning, or editing when wording can map to more than one operation or target.

- Ambiguous structural terms include merge, absorb, remove, delete, split, route, phase, surface, skill, spec, and contract, plus equivalent wording in the user's language.
- Ambiguous references include this, that, above, below, current one, `그`, `그 밑`, `그건`, and `그거` when nearby targets differ.
- Treat provenance notes, source URLs, user-spec intent blocks, and spec intent text as possible work targets.
- Ask a narrow question that locks the structure directly. Prefer `request_user_input` with bounded choices.
- Record literal wording, interpreted operation, operation target, alternate interpretations, and ambiguity impact in the flow record.
- Meaning resolution is not approval. If a locked target still involves destructive, irreversible, external, commit, push, PR, or publish work, request explicit user approval separately or record the handoff checkpoint.

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

- Do not execute destructive, irreversible, external, publish, push, PR, commit, or similar actions without explicit user approval for the exact action, or without a previously recorded initial agreement that covers the exact action boundary.
- Initial preparation may pre-plan expected risky actions as approved, not approved, deferred, or handoff-required. This planning narrows later self-drive behavior; it does not approve unplanned risky work.
- Before approval-sensitive work, present exact target, expected effect, risk, rollback or recovery possibility, and included/excluded scope.
- Do not use subagent output, inferred intent, nearby wording, readiness requests, or verifier pre-authorization as approval.
- If target ambiguity and action approval are both present, lock meaning first, then request action approval with exact risk and scope. A combined question is allowed only if those two decisions remain visibly separate.
- Commit approval requires staged/final status, intended diff, unrelated-change exclusion, and commit message scope before handoff to the commit workflow.
- Commit-readiness reporting is not commit approval and must not imply that commit execution may proceed.

## Verification

- Verify after `work` and before result reporting.
- Verification checks this flow's work, not the whole turn by default.
- For file edits, check intended file changes and relevant type, test, lint, build, or parse checks.
- For investigation or reasoning work, criticize the logic from multiple angles and check contrary evidence.
- Clean context means the verifier receives a bounded packet, not a full-history fork. Include verifier identity or request id, target, expected user intent, changed files or artifacts, commands/checks, pass/fail criteria, no-edit permission, and stop condition.
- While `turn-gate` is active, launching a read-only bounded verifier subagent is pre-authorized for this verification step. This does not authorize edits, implementation, scope expansion, destructive or external actions, approval decisions, commits, pushes, PRs, publishing, or bypassing user-gated boundaries.
- If the verifier would need to exceed that read-only bounded scope, stop and return to user-gated question-routing.
- Do not substitute same-context self-review as passing clean-context verification.
- Integrate verification as exactly one of `pass`, `fail`, `blocked`, or `insufficient`.
- If verification is `fail` or `insufficient`, return to the earliest safe phase before result reporting.
- If verification is `blocked` or unavailable, open a user-gated blocker through question-routing. Do not call it a pass.

## Reporting

Reporting is context for the next flow decision, not terminal closure.

- Include what was prepared, what work was done, verification result, residual uncertainty, blocker if any, material routing judgment calls, approval-boundary status, and the next concrete decision.
- Refresh the active flow record and read/update `Continuity Guard` before reporting.
- While `user_explicit_stop=false`, reporting requires no terminal summary and must continue into active question-routing, loop continuation, or a planned self-drive handoff.
- If a planned self-drive sequence is complete, report commit-readiness state instead of executing a commit. Include readiness, verification status, intended scope, unrelated-change exclusion, residual risk, and suggested handoff question if commit/push/PR/publish is desired.
- If planned flows are exhausted, use `request_user_input` when available to ask the user for the next flow or task.
- A stale `terminal summary allowed: yes` or source-less confirmed closure cannot justify terminal close.

## Question Routing

- Use user-gated question-routing for clarifications, choices, scope locks, mode narrowing, approval boundaries, blockers, and next-flow decisions.
- Use `request_user_input` when available and structural choices can be offered.
- Next-flow choices after result reporting must be narrow and directly connected to the reported result.
- Plain text follow-up or a generic closing phrase does not replace next-flow reopening.
- If `request_user_input` is unavailable, state that the tool is unavailable, list the open choices, and record the required next action. The turn remains in active question-routing.
- If visible choices cannot include a turn-end option, still tell the user they can explicitly stop the turn.
- Always record an explicit turn-end option in the flow record `Next Flow Options`.

## Session Records

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for active turn-gated work.

- Use `skills/turn-gate/templates/plan-template.md` for `000-plan.md`.
- Use `skills/turn-gate/templates/flow-record-template.md` for flow records.
- `000-plan.md` owns date-level history, user requests, flow index, planned flow sequence, transition criteria, and completed flow summaries.
- `000-plan.md` planned flow sequence must distinguish `operational-preparation` flows from executable `change-unit` flows.
- Each `001+` file owns one flow's details: user request message, task, flow type, flow scope, current mode, question-routing mode, current core phase, `Continuity Guard`, preparation source/result, planned flow list, work boundary, non-goals, expected risky actions, approval-boundary status, verification expectation, work, verification, report, next-flow options, and residual risk.
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
