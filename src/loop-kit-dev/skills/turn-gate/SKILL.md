---
name: turn-gate
description: Main loop controller for `loop-kit`. Keep one turn alive until the user asks to end the turn and select the right internal loop mode for the current phase of work.
---

# turn-gate

## Important

When this skill is active, `turn-gate` is a session-level first-class operating rule for the conversation itself.
Do not close the turn with a terminal summary unless the current user message explicitly asks to end the current turn.

Every non-stop user message is continuation input.
Treat questions, corrections, status checks, review requests, priority changes, handoff results, and next-task requests as input that can update the active flow, reopen preparation, refresh a target, or route to a user-gated decision.
If the user intent to stop is unclear, do not infer closure.

Required ending states are:

- active work continues into the next safe phase;
- active question-routing is open;
- a blocker is reported with the next required user decision;
- a planned self-drive handoff is prepared;
- explicit turn stop is confirmed from the current user message.

After each result report, reopen the next flow with `request_user_input` when structured choices are possible and the tool is available.
Plain follow-up wording does not replace active next-flow reopening.
If visible choices cannot include a turn-end choice, still record a turn-end option in the session record.

Maintain session records for turn-gated work.
Before result reporting and next-flow reopening, refresh the active flow record's `Continuity Guard`.
Terminal summary is allowed only when the guard matches a source user message that explicitly ends the current turn.
Stale or source-less closure notes do not authorize terminal summary.

## Purpose

Use `turn-gate` as the main loop controller for work that should continue within one turn until the user explicitly stops it.
The skill keeps continuity, passes each transition through internal gates, chooses one internal mode for current-phase work, records the flow, verifies work before reporting, and routes the user to the next flow instead of ending by default.

## Core Loop

Run each active flow through this order:

1. preparation
2. work
3. verification
4. reporting

Next-flow reopening is the continuation surface after reporting, not a core phase.
Activation, incoming message handling, explicit stop handling, self-drive handoff, and session record checks wrap the loop.

Do not promote phase names such as analysis, work, verification, reporting, or readiness checking into separate planned flows.
A flow is a cohesive reviewable or commit-sized unit; it does not have to be only a direct end-user value unit.
Final QA, integration checking, consistency checking, verification result reporting, and commit-readiness reporting are not separate planned flows unless they create or change a reviewable artifact of their own.

## Internal Gates

`turn-gate` uses internal gates as transition checks.
These gates are not separate user-facing skills and do not replace the core loop.
Each gate owns only its routing condition and may not infer terminal closure without an explicit stop message.

### Message Intake Gate

Classify the current incoming user message.

- Decide whether the message explicitly stops the current turn.
- If it is not explicit stop, treat it as continuation input.
- Identify whether it is a new task, correction, review, status check, approval-sensitive request, handoff result, next-task request, or another continuation shape.
- Mark operation or target ambiguity, scope gaps, and possible approval boundaries.

This gate does not execute work, finalize flow completion criteria, or decide turn closure.

### Flow Shaping Gate

Create or update the active flow from the message intake result.

- Decide whether to create a new flow, update the current flow, continue reporting, or open question-routing.
- Classify the flow as `operational-preparation`, `change-unit`, user-gated handoff, reporting context, or question-routing state.
- Record the flow boundary, completion criteria, verification expectation, and next-flow reopening condition.
- Keep follow-up `change-unit` candidates separate from active execution flows.

This gate prevents task lists, phase checklists, and direct-user-value-only slices from becoming invalid planned flows.

### Task Policy Gate

Choose the execution policy inside the selected flow.
Task policy is internal to a flow; it is not a separate layer or independent planned flow.

- Choose the task sequence, local references, target rereads, commands, edits, builds, checks, or handoffs needed for this flow.
- Confirm approval-sensitive work such as commit execution is inside the recorded approval boundary.
- Use the plan tool once meaningful work begins.

Individual task completion cannot decide flow completion or turn closure.
For example, a successful `git commit` still needs the commit flow's recorded confirmation, reporting, and next-flow reopening conditions.

### Verification Gate

Verify the current flow's work before reporting.

- Build a bounded clean-context verification packet when work was performed.
- Classify verification as `pass`, `fail`, `blocked`, or `insufficient`.
- For `fail` or `insufficient`, return to the earliest safe gate before successful reporting.
- For `blocked`, report the blocker through active question-routing.

This gate does not approve risky actions, choose the next flow, or close the turn.

### Reporting Gate

Report the flow result as continuity context.

- Include what was prepared, what work was done, how it was verified, what remains uncertain, and what approval boundaries remain.
- Surface material routing judgment calls.
- Refresh the session record and `Continuity Guard` before reporting.

Reporting is not terminal closure.

### Continuation Gate

Route the conversation after reporting.

- Check whether the current user message source-recorded explicit stop.
- If explicit stop is absent, continue into active question-routing, loop continuation, planned self-drive handoff, or blocker decision.
- Use `request_user_input` when structured choices are possible and the tool is available.
- Record a turn-end option in `Next Flow Options`, even when the visible prompt omits it.

This gate does not infer commit, push, PR, publish, or turn-stop approval from task results, prior wording, stale records, or source-less closure notes.

## Incoming Messages

Treat every incoming user message as authoritative input inside the same gated turn.
Only clear wording such as "end this turn", "stop the turn", "we are done here", "턴 종료", "여기서 끝", or equivalent intent counts as explicit turn stop.
If that intent is unclear, route the message as continuation input or ask a narrow clarification.

If continuation changes a target file, artifact, or state, reread that target before acting and do not reuse stale assumptions.
If continuation asks for a next flow, record it as the highest-priority next-flow candidate and continue to the next safe handoff point.
If the message is only activation without a concrete task, do not pick a work mode; open scope selection or next-flow choice instead.

## Preparation

Preparation decides what this flow is, why it exists, and when work may begin.

For user-message-based preparation, use deep-interview style alignment to collect the intent, included scope, excluded scope, success signal, approval boundary, and verification expectation needed for the planned work.
If the scope is empty, too broad, capable of producing multiple outcomes, or able to change success criteria or verification paths, lock scope with user-gated question-routing before work.

If scope is inferred without a question, record the inferred work boundary and non-goals in the flow record.
Do not proceed without a question when a wrong inference would make the result hard to reverse or would cross a risky approval boundary.

User-message interpretation and planned-flow design may be an `operational-preparation` flow.
That flow owns plan and session artifacts such as scope notes, approval boundaries, and proposed flow structure.
Its output may be follow-up `change-unit` candidates and their verification expectations.
Those candidates are not active execution flows.
Start a separate `change-unit` flow only after the user approves, selects, or asks to execute one.

When the user asks only for judgment, design, or scope confirmation, the `operational-preparation` flow may record follow-up candidates and stop at reporting plus handoff conditions.
Do not pull target file edits, release builds, or commit-readiness reporting into that same preparation flow unless the user also requested execution and the approval boundary allows it.

For non-user-message-based preparation, confirm the existing flow's target, current state, stale assumptions, preconditions, and exact work boundary before continuing.

If operation or target ambiguity can change the flow list or work result, resolve meaning before flow design or work.
Meaning resolution locks what the user meant.
Approval boundary separately decides whether a risky action may run.
Ambiguous verbs such as merge, absorb, move, promote, formalize, remove, delete, exclude, and equivalent wording in the user's language need meaning resolution when they can change artifact ownership, deletion behavior, output inclusion, migration scope, or approval boundaries.

Risky actions remain user-gated.
Destructive, irreversible, external, commit, push, PR, and publish actions require exact target, expected effect, risk, recovery possibility, included scope, and excluded scope before execution or handoff.
Prior wording, inferred intent, or subagent output does not grant approval.
If a new approval boundary appears after initial preparation, stop self-driven execution and ask.

Self-driven planned flow execution may hand off to the `turn-gate-self-drive` overlay when autonomous continuation is appropriate.
When that sequence is done, report commit readiness rather than executing a commit.
Commit, push, PR, and publish are separate user-gated handoffs.

## Work

Before work starts, pass through the task policy gate and select exactly one internal mode for the current phase.
If the mode is not clear, narrow it through short analysis or active question-routing.

Choose the earliest blocking mode that applies:

- `deep-interview`: unclear requirements, missing scope, unresolved non-goals, or unresolved approval lines block work.
- `review-loop`: review feedback, QA findings, or self-review findings are the material issue.
- `ralph-loop`: one small fix-verify-reassess cycle is the right unit.
- `autopilot`: broad end-to-end delivery is the current need.
- `commit-readiness-gate`: the change unit is nearly complete and readiness judgment is the current need.

After selecting a mode, read the matching local file in `references/` and apply that contract before doing the work.
External actions such as commit execution, push, PR, or publish are not internal modes.
They are user-gated handoff workflows.

Commit execution is task policy inside an approved commit handoff flow.
Commit completion does not end the turn.
After commit confirmation, report the commit result and reopen next-flow routing unless the user explicitly stopped the turn.

## Verification

After work and before reporting, run explicit verification.
For file changes, verify the applied diff and use the narrowest meaningful checks such as parsing, linting, tests, build, or command output inspection.
For judgment or design work, verify by checking the reasoning against alternatives, contradictions, and the user's stated scope.

Clean-context verification is required after work.
Use a read-only verifier subagent with a bounded packet, not a full-history fork.
The packet should include only the needed paths, user intent, change summary, checks to run, pass/fail criteria, no-edit permission, and stop conditions.

The verifier may not edit files, expand scope, perform destructive or external actions, or approve commit/push/PR/publish work.
If verification needs any of those, stop and route back to the user.

Classify verification as `pass`, `fail`, `blocked`, or `insufficient`.
Do not treat `fail`, `blocked`, or `insufficient` as pass.
For `fail` or `insufficient`, return to the earliest safe gate before reporting successful completion.
For `blocked`, report the blocker through active question-routing.

## Reporting

Reporting is continuity context, not terminal closure.
Report what was prepared, what work was done, how it was verified, what remains uncertain, and what decision or flow should happen next.

Before reporting:

- update the active flow record through the current phase;
- refresh `Continuity Guard`;
- confirm whether the current user message explicitly stopped the turn;
- confirm verification status is not being overstated.

If explicit stop is absent, pass through the continuation gate after the report.

## Next-Flow Reopening

Use `request_user_input` for clarification, scope locks, mode narrowing, approval decisions, and next-flow selection when structured choices are possible.
The choices should be narrow and directly connected to the result just reported.

If `request_user_input` is unavailable, use a plain-text fallback that states:

- the tool was unavailable;
- the open choices;
- the required next action recorded in the session record.

Fallback is still active question-routing.
It is not a terminal summary.

## Session Records

For turn-gated tasks, keep `.agents/sessions/{YYYYMMDD}/000-plan.md` and per-flow records under `.agents/sessions/{YYYYMMDD}/`.

`000-plan.md` owns the date-scoped task history, user requests, flow index, current planned flow sequence, completed flow summaries, and transition conditions.
The plan is a sequence of cohesive flows, not a checklist of phases.
Do not delete completed flow summaries; append and update incrementally.

Each per-flow record should be named with a three-digit counter and an English lowercase slug.
Update it after preparation, work, verification, and reporting rather than waiting until the end.

Record at least:

- user request message;
- task;
- flow type: `operational-preparation` or `change-unit`;
- flow scope;
- current mode;
- question-routing mode;
- current core phase;
- preparation source and result;
- planned flow list or follow-up `change-unit` candidates;
- `Continuity Guard`;
- work;
- verification;
- report;
- next-flow options;
- residual risk.

For an `operational-preparation` flow that ends after judgment, design, or scope confirmation, record follow-up `change-unit` candidates instead of treating them as active planned flows.
Each candidate should include the expected artifact, why it is separate or combined, expected verification, and the user-gated handoff condition.

`Continuity Guard` must include:

- whether `turn-gate` is active;
- question-routing mode;
- whether the user explicitly stopped the turn;
- whether terminal summary is allowed;
- required next action;
- last refreshed phase;
- verification status when verification is needed.

Use verification statuses `not-started`, `requested`, `pass`, `fail`, `blocked`, or `insufficient`.
Confirmed closure is valid only with the source user message and recorded phase.
If closure state is stale or source-less, reset terminal summary permission to no and record why.

`Next Flow Options` must always include a recorded turn-end option, even when the visible prompt omits it.

## Quick Checks

- Is `turn-gate` still a first-class rule for this conversation?
- Did the current user message explicitly stop the turn?
- Did message intake classify the message without executing work?
- Did flow shaping define or update the active flow and keep candidates separate from active execution flows?
- Did task policy stay inside the flow instead of deciding flow completion or turn closure?
- Is the current response ending in active work, active question-routing, planned self-drive handoff, a blocker decision, or confirmed closure?
- Did preparation distinguish execution candidates from active execution flows?
- Did the flow sequence avoid phase-only entries and direct-user-value-only filtering?
- Did final QA or readiness-only work stay inside verification/reporting or a user-gated handoff unless it changed a reviewable artifact?
- Was exactly one internal mode selected before work?
- Was the selected local `references/` contract read?
- Was clean-context verification requested with a bounded read-only packet?
- Were non-pass verification outcomes kept out of successful completion language?
- Was `Continuity Guard` refreshed before reporting and next-flow reopening?
- Did `request_user_input` or active fallback reopen the next flow unless explicit stop was confirmed?
