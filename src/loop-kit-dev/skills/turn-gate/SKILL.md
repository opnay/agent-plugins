---
name: turn-gate
description: Main loop controller for `loop-kit-dev`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# Turn Gate

## Important

- Once invoked, `turn-gate` is a session-level first-class operating rule for the response lifecycle, not an internal checklist.
- While `user_explicit_stop=false`, do not close with a terminal summary. Result reporting must continue into next-flow reopening, loop continuation, or active question-routing.
- If the user only activates `turn-gate`, do not choose a work mode. Activate the gate, keep `user_explicit_stop=false`, update or create the session record, and open a scope or next-flow choice.
- Every response must end in one of these states: loop continuation, active question-routing, or explicit user stop handling.
- Use `request_user_input` when available for next-flow decisions, scope locks, mode narrowing, meaning clarification, and user-gated approval choices.
- Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and the active flow record, including `Continuity Guard`, pending question state, verification status, and `Next Flow Options`.
- After `work`, request clean-context verification before result reporting. The verifier receives only a bounded packet, not a full-history fork.
- Failed, blocked, unavailable, or insufficient verification is not a pass.
- A terminal summary is allowed only when an explicit stop is source-recorded in the active flow record with the closure source message.

## Purpose

Use this skill as the main operational surface of `loop-kit-dev`. Keep one turn open until the user explicitly asks to end it, choose the internal loop mode for the current phase of work, verify completed work, report results, and reopen the next flow.

## Core Loop

Follow this phase order unless a correction requires returning to the earliest safe phase:

1. activation
2. incoming message classification
3. analysis
4. plan
5. work
6. verification
7. result reporting
8. next-flow reopening
9. explicit stop handling

### Activation

- Treat invocation as activating a conversation-level operating rule.
- If there is no concrete task beyond activation, do not infer one. Open a scope or next-flow choice with `request_user_input` when available.
- Create or update the session records before substantial work starts.

### Incoming Message Classification

Classify every new user message as authoritative input for the same gated turn:

- explicit turn stop
- status or progress check
- current-flow correction
- current-flow priority change
- next-flow priority request

Status checks get a short status, blocker if any, and next concrete action before work continues. Corrections and priority changes update analysis/plan immediately and resume from the earliest safe phase. If a correction changes the target file, artifact, or state, re-read that target before continuing and do not reuse stale assumptions. A next-flow priority request is recorded as a candidate for the next safe handoff point. Do not treat status checks, corrections, or "do X next" requests as explicit stop.

Only a clear request to end the current turn counts as explicit stop.

### Analysis

- Identify requested intent, requested action, current blocker, likely internal mode, meaning-resolution needs, and approval boundaries.
- Resolve operation or target ambiguity before mode selection or work when different readings change files, deletion behavior, routing, phase design, migration meaning, or commit scope.
- Ambiguous terms include structural verbs such as merge, absorb, remove, delete, split, route, phase, surface, skill, spec, and contract, plus pronouns or positional references that could point to multiple targets.
- Treat provenance notes, source URLs, user-spec intent blocks, and spec intent text as possible work targets when wording such as source, original, intent, above, or below could refer to them.
- For structural ambiguity, ask a narrow question that locks the target or operation. Prefer `request_user_input` with choices and record literal wording, interpreted operation, operation target, alternate interpretations, and ambiguity impact.
- Keep meaning resolution separate from approval. A locked meaning does not grant approval for risky execution.
- If one user message has both target ambiguity and approval-sensitive action, lock the target/meaning first, then request action approval with exact risk and scope. A combined question is acceptable only when it clearly separates those two decisions.

### Plan

- Define the active steps for the current flow and use the planning tool once meaningful work begins.
- For multi-flow work, keep `.agents/sessions/{YYYYMMDD}/000-plan.md` as the flow sequence, not a low-level checklist.
- Each planned flow should state purpose, why it exists, completion criteria, and transition criteria.
- Put detailed task steps in the active `001+` flow record.

### Work

- Before work begins, select exactly one internal mode for current-phase work.
- If mode is not clear, use narrow analysis or active question-routing before work.
- Read and apply the selected local reference from `skills/turn-gate/references/`:
  - `deep-interview.md`: requirement discovery, unclear intent, missing scope boundaries, unresolved approval lines
  - `review-loop.md`: review feedback, QA finding, or self-review finding is the material issue
  - `ralph-loop.md`: one small fix-verify-reassess cycle is the right unit
  - `autopilot.md`: broad end-to-end delivery is the current phase
  - `commit-readiness-gate.md`: the change unit is nearly complete and readiness judgment is the current phase
- If several modes fit, prefer the earliest blocker in this order: `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `commit-readiness-gate`.
- Commit execution, push, PR, publish, and similar external actions are not internal modes. Treat them as user-gated handoff workflows.

## Approval Boundaries

- Do not execute destructive, irreversible, external, publish, push, PR, or commit actions without explicit user approval.
- Before approval-sensitive work, present exact target, expected effect, risk, rollback or recovery possibility, and included/excluded scope.
- Do not use subagent output, inferred intent, nearby wording, or readiness requests as approval.
- Commit approval requires staged/final status, intended diff, unrelated-change exclusion, and commit message scope before handoff to the commit workflow.
- Push, PR, publish, and other external effects require branch, remote or target, external effect, and risk before handoff to the matching workflow.

## Verification

- After `work` and before result reporting, request clean-context verification.
- Clean context means the verifier receives a bounded packet only: verifier id or request id, target, expected user intent, changed files or artifacts, commands/checks to run, pass/fail criteria, no-edit permission, and stop condition.
- Do not substitute same-context self-review as passing verification.
- Integrate the verifier result as exactly one of `pass`, `fail`, `blocked`, or `insufficient`.
- If verification is `fail` or `insufficient`, return to the earliest safe phase before result reporting.
- If verification is `blocked` or the verifier/tool is unavailable, open a user-gated blocker through question-routing. Do not call it a pass.
- A non-pass status report may be shown only as a blocker report with active question-routing; it must not be framed as successful completion.

## Result Reporting

- Result reporting is context for the next decision, not terminal closure.
- Include what changed, verification result, residual uncertainty, blocker if any, material judgment calls that affected routing or phase selection, and the next concrete decision.
- Refresh the active flow record and read/update `Continuity Guard` before reporting.
- A stale `terminal summary allowed: yes` or source-less confirmed closure cannot justify terminal close.

## Question Routing

- Use user-gated question-routing for clarifications, choices, scope locks, mode narrowing, approval boundaries, blockers, and next-flow decisions.
- Use `request_user_input` when available and structural choices can be offered.
- Next-flow choices after result reporting must be narrow and directly connected to the reported result.
- If `request_user_input` is unavailable, state that the tool is unavailable, list the open choices in the report, and record the required next action. The turn remains in active question-routing.
- If visible choices cannot include a turn-end option, still tell the user they can explicitly stop the turn.
- Always record an explicit turn-end option in the flow record `Next Flow Options`.

## Session Records

- Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for active turn-gated work.
- Use `skills/turn-gate/templates/plan-template.md` for `000-plan.md` and `skills/turn-gate/templates/flow-record-template.md` for flow records.
- `000-plan.md` owns date-level history, user requests, flow index, planned flow sequence, transition criteria, and completed flow summaries.
- Each `001+` file owns one flow's details: user request message, task, flow scope, current mode, question-routing mode, `Continuity Guard`, analysis, plan, work, verification, result report, next-flow options, and residual risk.
- Update the flow record incrementally after analysis, plan, work, verification, and result reporting.

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

Record confirmed closure only after explicit stop. Include closure source message and closure recorded phase. If records are inaccessible, report that as a blocker and do not reconstruct silent terminal-close permission.
When stale terminal-summary permission is detected, refresh the guard to `user_explicit_stop=false`, `terminal summary allowed=no`, and record that the prior closure state was stale or source-less.

## Stop Handling

- Do not send a terminal summary unless the current incoming message clearly asks to end the turn or a valid confirmed closure with source message matches the current state.
- On explicit stop, write confirmed closure, closure source message, and closure recorded phase into the active flow record.
- After source-recorded closure, terminal summary is allowed and next-flow reopening is not required.
