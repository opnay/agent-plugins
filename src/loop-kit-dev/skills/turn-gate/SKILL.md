---
name: turn-gate
description: Keep an active Codex turn open across preparation, work, verification, reporting, and next-flow routing until the user explicitly stops the turn; apply sibling flow contracts, maintain session records, route questions, verify before reporting, and support prepared self-drive sequences.
---

# Turn Gate

## Active Turn Rule

When this skill is active, keep the current turn open until the user explicitly asks to end it and that stop source is recorded. Task completion, a passing check, a status answer, a completed report, an interrupted question tool call, or a final-looking summary is not terminal closure authority.

Do not use a terminal/final closeout as the normal report. Reporting and next-flow reopening stay in the ongoing conversation channel unless a source-recorded explicit stop allows closure.

Every active flow must end in exactly one recorded state:

- `next-flow`: reporting is complete, records are updated, and the next required action is open.
- `blocked`: user input, approval, access, or external state is required before continuing.
- `explicit-stop`: the current user message explicitly ends the turn and the closure source is recorded.

Maintain session records while the turn is active. Use `references/session-records.md` for the record model, templates, Continuity Guard, recovery cases, and the split between `000-plan.md`, active flow records, and optional self-drive sidecars.

## Five-Phase Lifecycle

Run each active flow through:

1. `preparation`: lock intent, scope, non-goals, acceptance signal, verification expectation, approval boundary, and handoff condition. Apply the sibling `flow` contract for flow boundary, readiness, ambiguity, and flow-local strategy; do not redefine those rules here.
2. `work`: execute only inside the recorded active flow boundary.
3. `verification`: choose a verification method, run or justify it, and record result status.
4. `reporting`: update records first, then report continuity context rather than closing the turn.
5. `next-flow`: route to a next-flow choice, blocker decision, valid self-drive continuation, or source-recorded explicit stop.

At the start and end of each active flow phase, apply the sibling `flow` phase record checkpoint expectation. Decide which surface changed:

- Update `000-plan.md` when the active flow pointer, date-level required next action, planned/current sequence, or turn-level routing changes.
- Update the active flow record when the same flow's current phase, execution log, verification evidence, report outcome, residual risk, or handoff condition changes.

Phase checkpoints do not make phases separate flows. `preparation`, `work`, `verification`, and `reporting` remain phases inside the same active flow.

Use the phase prefix at the start of user-facing phase-start or phase-progress messages:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Do not copy the prefix into generated artifacts, session records, command summaries, or question option labels.

For meaningful multi-step work, use the available planning tool to keep the current phase or task state visible.

## Preparation And Approval Boundary

Before work begins, ensure the active flow contract is recorded with:

- scope
- non-goals
- completion or acceptance signal
- verification expectation
- approval boundary
- handoff condition

If any missing target, operation, endpoint, success condition, approval boundary, or verification path could change the work, route through a user-gated clarification before proceeding.

Approval-sensitive actions require exact target, expected effect, risk, recovery path, included and excluded scope, and endpoint before the execution checkpoint. Readiness, verification, self-drive, previous context, or subagent output cannot authorize commit, push, PR, publish, release, version bump, destructive history rewrite, or external side effects.

## Verification

Before reporting, select and record a verification method and reason:

- `clean-context`: bounded read-only verifier packet, not a full-history fork.
- `normal`: main-thread checks, readback, evidence review, or logical counterexample review.
- `not-required`: no separate verification action is justified; record the reason and residual uncertainty.

Record result status separately from method:

- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-required` is not a pass. Progress states such as `not-started` and `requested` are not success evidence.

Default to `clean-context` for file changes, release surface changes, multi-file contract changes, prior check failures, user-requested verification/review/QA/commit-readiness, and approval-sensitive action boundaries. Use `references/verification.md` for method details, verifier packet boundaries, and non-pass routing.

Route non-pass results before success reporting:

- `fail`: return to the earliest safe repair or work point.
- `insufficient`: collect more evidence or strengthen verification.
- `blocked`: open user-gated blocker routing.

## Reporting And Next Flow

Before reporting, update the active flow record and any required `000-plan.md` routing fields. The report must include:

- changed surfaces
- verification status
- material judgment calls
- residual risk
- required next action

After reporting, reopen routing unless the current user message explicitly stopped the turn and that source is recorded.

Use `request_user_input` when available for narrow next-flow choices, clarifications, blocker recovery, approval-boundary decisions, or pending question recovery. If the tool is unavailable, use an active plain-text fallback and record the required next action.

Always keep explicit turn-end available in the flow record's `Next Flow Options`, even when the visible question UI cannot show a stop option. Use `references/question-routing.md` for structured question use, fallback behavior, blocker routing, and interrupted question recovery.

An aborted, canceled, or interrupted question tool call is not flow completion and is not terminal closure. Record the pending question state, keep `terminal_summary_allowed: no`, and interpret the next user message as a pending answer, superseding flow request, status question, or explicit stop.

## Self-Drive

Self-drive is an explicit prepared sequence overlay; it is not the default turn state and does not replace the five-phase lifecycle. Use `references/self-drive.md` when records show a prepared sequence objective, active flow index, allowed and prohibited autonomous actions, approval-sensitive checkpoints, endpoint, blocker return conditions, acceptance signal, and verification expectation.

At each self-drive flow start, read `000-plan.md` and `000-self-drive.md`. `000-plan.md` stores only self-drive status and sidecar pointer; `000-self-drive.md` owns sequence-level state; the active flow record owns flow-local state.

Sequence completion is not terminal closure. Report completion, update records, then route to next-flow unless the user explicitly stops the turn.
