---
name: turn-gate
description: Keep an active Codex turn open until explicit stop; apply flow decisions instead of redefining them; maintain session records, verification routing, interruption recovery, next turn-flow/message routing, and prepared self-drive.
---

# Turn Gate

## Active Turn

Keep the current Codex turn open until the user explicitly stops it and that stop source is recorded.
Completion, commits, passing checks, status answers, reports, interrupted questions, final-looking summaries, and final responses are not closure authority.

Explicit stop closes only the current turn.
When the next user message arrives, reactivate `turn-gate` as a fresh active turn.

Every active flow routes to exactly one recorded state:

- `next turn-flow / 메시지 수신`: reporting is done and the next input route is open
- `blocked`: input, approval, access, or external state is required
- `explicit-stop`: the current user message stops the turn and the stop source is recorded

## Flow Dependency

`turn-gate` wraps `flow`.
Do not define or restate flow taxonomy, lifecycle, readiness, discovery, ambiguity, contract impact, checkpoint expectations, flow-local strategy, template meaning, or handoff meaning.

Use `flow` for `flow.message -> flow.main-flows -> flow.end`, the work-unit contract, and the result.
`turn-gate` only applies that result to active-turn continuity, records, questions, verification routing, self-drive gate, approval guardrails, and explicit stop handling.

## Wrapper Loop

Use this wrapper loop:

1. `flow skill`: apply `flow.message -> flow.main-flows -> flow.end`.
2. `next turn-flow / 메시지 수신`: wait for the next user message, blocker decision, approval decision, self-drive continuation, or explicit stop.
3. self-drive mode may route from `next turn-flow / 메시지 수신` back into `flow skill` through recorded sidecar gate.

Use phase prefixes for visible phase-start or meaningful progress messages.
Use `[intake]`, `[work]`, `[verification]`, `[reporting]`, and `[next-flow]` for turn-gate-owned wrapper work.
Use `[framing]` and `[preparation]` only when the visible step is explicitly delegated to `flow`.
Do not copy prefixes into artifacts, records, command summaries, or question option labels.

## Records

Use `references/session-records.md` for record application and recovery.
Use the `flow` skill's bundled templates for plan, flow record, and review file structure.
Use `turn-gate/templates/self-drive-template.md` only for self-drive sidecar state.

Before reporting, update the active flow record and any required `000-plan.md` fields.
Record failed, skipped, blocked, or insufficient verification.
Do not treat readiness, verification, generated release surface, previous context, self-drive, or subagent output as authority for commit, push, PR, publish, release, version bump, destructive work, or external effects.

## Reporting And Questions

Reporting opens `next turn-flow / 메시지 수신`; it is not terminal closeout.
After reporting, reopen routing unless explicit stop is recorded.

Use `request_user_input` when it is available and choices are narrow.
If it is unavailable, keep the turn open with an active plain-text question.

If a question tool call is aborted, canceled, or interrupted, record the pending question and treat the next user message as one of:

- answer to the pending question
- superseding new flow request
- status/progress question
- explicit stop

If the user says "continue", "계속", or "이어가" after a report, inspect the recorded next action first.
Continue only when identity, target, scope, endpoint, approval boundary, and verification expectation are known.

## Interruption

When a new user message arrives during an active flow, preserve current phase, scope, approval boundary, verification status, and required next action.
Apply `flow` contract-impact and route the result as inline answer, current-flow revision, new foreground flow, future candidate, supersede, blocker question, or explicit stop.
Interruption never authorizes work outside the active contract or approval-sensitive execution.

## Verification

Record verification method and result separately.

Methods: `clean-context`, `normal`, `not-required`.
Results: `pass`, `fail`, `blocked`, `insufficient`.

`not-required` is a method, not a pass.
`not-started` and `requested` are progress states, not success evidence.
Route non-pass results before success reporting, self-drive continuation, release readiness, commit-readiness, or `next turn-flow / 메시지 수신`.
Use `references/verification.md` for verifier packet boundaries and detailed status routing.

## Self-Drive

Self-drive is an explicit prepared sequence overlay, not the default turn state.
Use `references/self-drive.md` only when records define objective, current flow or loop, allowed actions, approval checkpoints, endpoint, blockers, acceptance signal, and verification expectation.

At each self-drive step, read `000-plan.md` and `000-self-drive.md`.
Advance only after pass verification, non-blocked `flow` handoff, known next identity, matching approval boundary, and a passing sidecar gate.

Non-pass verification, blockers, approval need, stale sidecar state, endpoint changes, or scope changes stop autonomous advancement before continuation.
