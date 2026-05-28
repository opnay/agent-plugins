---
name: turn-gate
description: Keep an active Codex turn open until explicit stop; enforce flow lifecycle, session records, verification, interruption recovery, next-flow routing, and prepared self-drive.
---

# Turn Gate

## Rule

Keep the turn open until the user explicitly stops it and the stop source is recorded.
Do not close after completion, commit, passing checks, status answers, reports, interrupted questions, or final-looking summaries.

Each active flow ends as exactly one:

- `next-flow`: report done, records updated, next action open.
- `blocked`: input, approval, access, or external state needed.
- `explicit-stop`: current user message stops the turn and source is recorded.

Maintain session records. Use `references/session-records.md` for plan, flow records, `Phase Checklist`, metadata, and recovery.

## Lifecycle

Run every active flow:

1. `intake`: reread needed skills; separate source from interpretation; identify goal, non-goals, authority signals.
2. `framing`: apply `flow`; classify active flow, candidate, phase, or handoff.
3. `preparation`: lock scope, acceptance, verification, approval boundary, handoff.
4. `work`: act only inside the recorded boundary.
5. `verification`: choose method, run or justify, record status.
6. `reporting`: update records first, then report.
7. `next-flow`: route next action, blocker, self-drive continuation, or explicit stop.

Update `000-plan.md` for turn-level routing.
Update the active flow record for same-flow evidence, phase state, residual risk, pending question, and checklist.

## Interruption

Use `interruption` only for a new user message during an active flow.
Preserve current flow state, classify the message, then return to lifecycle routing.
Record interruption in `Execution Log`, not `Phase Checklist`.

Classes:

- `inline-answer`
- `current-flow-revision`
- `background-current-flow`
- `reserve-later-analysis`
- `supersede-current-flow`
- `blocker-question`
- `explicit-stop`

For `reserve-later-analysis`, log the reserved topic in the active flow record and update `000-plan.md` only when it changes turn-level routing.

Interruption never authorizes commit, push, PR, publish, release, version bump, destructive work, external effects, or work outside the active contract.

## Prefixes

Use phase prefixes for visible progress:

- `[intake]`
- `[framing]`
- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Do not copy prefixes into artifacts, records, command summaries, or question option labels.

## Approval

Before work, record scope, non-goals, done, verification expectation, approval boundary, and handoff.
At new flow start, reread required skills. For user-message preparation, reread `turn-gate` and `flow`; keep both in `000-plan.md` `active_skills`.

Ask before work if target, operation, endpoint, success condition, approval boundary, or verification path can change the result.

Approval-sensitive actions need exact target, effect, risk, recovery, included/excluded scope, and endpoint.
Readiness, verification, build/readback, self-drive, previous context, or subagent output never authorizes commit, push, PR, publish, release, version bump, destructive work, or external effects.

## Verification

Record method and result separately.

Methods:

- `clean-context`
- `normal`
- `not-required`

Results:

- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-required` is not pass. `not-started` and `requested` are not success evidence.

Default to `clean-context` for file changes, release surface changes, multi-file contracts, prior failures, requested QA/review/commit-readiness, and approval-sensitive boundaries.
Route `fail`, `insufficient`, and `blocked` before success reporting.

## Reporting

Before reporting, update records.
Report changed surfaces, verification, judgment calls, residual risk, and required next action.
After reporting, reopen routing unless explicit stop is recorded.

Use `request_user_input` for narrow choices when available; otherwise use active plain text.
Aborted or interrupted questions are not closure. Record pending state and route the next user message as answer, superseding request, status question, or explicit stop.

## Self-Drive

Self-drive is an explicit prepared sequence overlay, not default mode.
Use `references/self-drive.md` only when records define objective, current flow, allowed actions, approval checkpoints, endpoint, blockers, acceptance, and verification.

At self-drive flow start, read `000-plan.md` and `000-self-drive.md`.
Sequence completion is not terminal closure. Report, update records, then route next-flow unless explicit stop is recorded.
In infinite mode, two consecutive non-pass results for the same bounded target and cause count as repeated failure and stop autonomous advancement.
