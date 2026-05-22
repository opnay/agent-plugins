---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and next-flow routing until the user explicitly stops it. Use when a request needs continuity records, explicit stop handling, post-report next-flow reopening, question abort recovery, verification status discipline, or a prepared self-drive sequence. active turn, next-flow, continuity guard, explicit stop, session records, question recovery, self-drive
---

# Turn Gate

## Active-Turn Rule

When this skill is active, keep the current Codex turn operationally open until the latest user message explicitly stops it and that stop source is recorded.

Do not treat task completion, successful verification, a status answer, an answered question, an interrupted question tool call, or final-sounding wording as permission to close the turn. A report is continuity context, not a terminal summary. After reporting, route to exactly one of these states:

- `next-flow`: records are updated and the required next action is open.
- `blocked`: user input, approval, access, or an external state change is required.
- `explicit-stop`: the current user message explicitly ends the turn and the closure source is recorded.

If the user only activates `turn-gate`, record the active state and open scope or next-flow routing. Do not answer with only an activation summary.

For meaningful multi-step work, maintain the available plan or task-state tool with the current phase or task status.

## Lifecycle

Run each active flow through five phases:

1. `preparation`: lock intent, scope, non-goals, acceptance signal, verification expectation, approval boundary, and handoff condition. Apply the sibling `flow` skill or record an equivalent flow contract before work begins.
2. `work`: execute only inside the active flow boundary.
3. `verification`: choose a method, run or justify it, and record the result status separately from the method.
4. `reporting`: update records first, then report changed surfaces, verification status, material judgment calls, residual risk, and required next action.
5. `next-flow`: reopen the next flow, route a blocker, continue a valid self-drive sequence, or record explicit stop.

User-facing phase-start and phase-progress messages start with one canonical prefix: `[preparation]`, `[work]`, `[verification]`, `[reporting]`, or `[next-flow]`. Do not copy these prefixes into generated artifacts, record headings, command summaries, or question option labels.

## Preparation Gates

Do not start work until the active flow contract covers:

- scope and non-goals
- completion or acceptance signal
- verification expectation
- approval boundary
- handoff condition
- exact target and operation when those affect success or verification

If scope, target, endpoint, acceptance signal, approval boundary, or verification path is unclear, ask before work. Use structured `request_user_input` when it is available and the choices are narrow. Otherwise use plain text and keep the routing active.

Approval-sensitive actions require exact target, expected effect, risk, recovery path, included and excluded scope, and endpoint. Readiness, verification, self-drive state, prior context, or subagent output never authorizes commit, push, PR, publish, release, version bump, destructive history rewrite, or external side effects.

## Session Records

Maintain runtime continuity records unless the user explicitly forbids recording. Use the templates in `templates/` when creating records:

- `templates/plan-template.md` for `.agents/sessions/{YYYYMMDD}/000-plan.md`
- `templates/flow-record-template.md` for `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`
- `templates/self-drive-template.md` for `.agents/sessions/{YYYYMMDD}/000-self-drive.md` when self-drive is active

Read `references/session-records.md` when creating, recovering, or updating records. Records are updated incrementally after each phase. Refresh the Continuity Guard before reporting and before next-flow reopening.

Keep raw user text separate from summaries or interpretations. If a record that should exist is unexpectedly missing or inaccessible, do not silently reconstruct it. Route blocker recovery or ask the user how to proceed.

If `no-record`, `기록 남기지 마`, or similar wording could mean either "do not read records" or "do not write records", clarify before reading existing records.

## Verification

Verification has two separate fields:

- method: `clean-context`, `normal`, or `not-required`
- result status: `pass`, `fail`, `blocked`, or `insufficient`

`not-required` is a method, not a pass. Record the reason and residual uncertainty when using it.

Default to `clean-context` for file changes, release surface changes, multi-file contract changes, prior check failures, user-requested review or QA, commit-readiness, and approval-sensitive actions. Use `normal` or `not-required` only when the flow record explains why that is sufficient for the actual risk.

Route non-pass results before success reporting:

- `fail`: return to the earliest safe repair point.
- `insufficient`: collect more evidence or strengthen verification.
- `blocked`: open user-gated blocker routing.

Read `references/verification.md` for method selection and verifier packet boundaries.

## Question Routing

After reporting, reopen next-flow routing unless the latest user message explicitly stopped the turn. Read `references/question-routing.md` for structured choice and recovery details.

Use `request_user_input` when available for narrow choices. If tool constraints prevent a visible stop option, still state that the user can explicitly end the turn, and include that option in the record. If the tool is unavailable, say so and use active plain-text routing.

If a question tool call is aborted, canceled, or interrupted, do not close the turn and do not immediately repeat the same question tool call. Record the pending question as `aborted`, `interrupted`, or `superseded`, then interpret the next user message as an answer, a new flow request, a status request, or explicit stop.

## Self-Drive

Self-drive is only an overlay on the normal lifecycle. It is valid only when records contain a finite prepared sequence, current index and label, progress note, allowed and prohibited autonomous actions, approval-sensitive checkpoints, endpoint, blocker return conditions, acceptance signal, and verification expectation.

Do not infer self-drive from enthusiasm, a long task list, verification success, or subagent availability. Read `references/self-drive.md` before running or continuing a self-drive sequence.

Sequence completion is not turn closure. After endpoint handling and verification, update records, report completion, and reopen next-flow routing unless the user explicitly stops.
