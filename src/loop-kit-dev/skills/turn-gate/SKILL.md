---
name: turn-gate
description: Keep an active Codex turn open until explicit stop; apply flow decisions instead of redefining them; maintain session records, verification routing, interruption recovery, reporting-as-pre-intake next-flow routing, and prepared self-drive.
---

# Turn Gate

## Active Turn

Keep the active Codex turn open until the user explicitly stops it and that stop source is recorded.
Completion, commits, passing checks, status answers, reports, interrupted questions, final-looking summaries, and final responses are not closure authority.

Explicit stop closes only the current turn. When the next user message arrives, reactivate `turn-gate` as a new active turn regardless of whether the previous stop was intentional, accidental, interrupted, or terminal-looking.

Every active flow ends in exactly one recorded state:

- `next-flow`: reporting is done, records are updated, and the next user decision is open
- `blocked`: input, approval, access, or external state is required
- `explicit-stop`: the current user message stops the turn and the stop source is recorded

Maintain session records throughout the turn.
Use `references/session-records.md` for `000-plan.md`, flow records, phase checklists, compact metadata, and recovery rules.

If the user only activates `turn-gate`, record the operating state and open scope or next-flow routing.
Do not answer with a terminal activation summary.

If the previous record shows `explicit-stop` or stale closure, treat the new user message as fresh intake, reset closure authority, and record the reactivation before work.

## Flow Dependency

`turn-gate` applies `flow`; it does not define flow taxonomy, readiness, discovery, ambiguity, contract impact, phase checkpoint expectations, verification expectation, flow-local strategy, or handoff.
Use `flow` for those decisions, then record and route the result.

`turn-gate` owns active-turn continuity, session record maintenance, question routing, verification method/status routing, approval-sensitive guardrails, self-drive sidecar gate, and explicit-stop handling.

## Lifecycle

Run each active flow through:

1. `intake`: reread needed skills and apply `flow` intake.
2. `framing`: apply `flow` classification and selected-vs-candidate distinction.
3. `preparation`: apply `flow` readiness and ambiguity decisions before work.
4. `work`: act only inside the recorded flow boundary.
5. `verification`: choose method, run or justify it, and record result status.
6. `reporting`: update records, report continuity context, and create the next decision surface.
7. `next-flow`: route a next action, blocker, self-drive continuation, or explicit stop.

Use phase prefixes for visible phase-start or meaningful progress messages: `[intake]`, `[framing]`, `[preparation]`, `[work]`, `[verification]`, `[reporting]`, `[next-flow]`.
Do not copy prefixes into artifacts, records, command summaries, or question option labels.

Use `flow` checkpoint expectations to decide whether `000-plan.md`, the active flow record, or both need updates.
For meaningful multi-step work, use the available plan tool to keep current phase or task state visible.

At a new flow start, reread the skills needed for that flow.
When a user message moves toward preparation, reread `turn-gate` and `flow`; keep both in `000-plan.md` `active_skills`.

## Reporting As Pre-Intake

When `turn-gate` is active, reporting is not terminal closeout.
Reporting is the pre-intake transition for the next user decision.

Before reporting, update the active flow record and any required `000-plan.md` fields.
Then report the result, verification status, material judgment calls, residual risk, and required next action.
After that, create the next decision surface:

- Use `request_user_input` when it is available and the choices are narrow.
- If the question tool is unavailable, keep the turn open with an active plain-text question.
- If a valid self-drive continuation is already prepared, route through that continuation gate.
- If an explicit stop is present, record the stop source before closure.

If `turn_gate_active` is true and no explicit stop is recorded, do not use a final/terminal closeout as the last action.
Compression, status-only reporting, or summary wording cannot remove the next-flow question, blocker question, or valid self-drive handoff.

Use `references/question-routing.md` for next-flow choices, post-flow continue, fallback text, and question recovery.

## Dates

Interpret relative dates from the current system date and timezone by default.
If a relative date affects result, target, verification path, reporting scope, or record reconstruction, write the absolute date.

If a request points to session records, prior flows, last record, yesterday's work, or record-based resume, do not let record dates silently override system date.
If date source changes the target or verification path, treat it as `flow` ambiguity/readiness and ask before work.

Use compact explicit wording such as `today (2026-05-28)`, `yesterday (2026-05-27, system-date basis)`, or `needs clarification: last record date or system-date yesterday`.

## Preparation And Approval

Before work, record the `flow` contract: scope, non-goals, completion or acceptance signal, verification expectation, approval boundary, and handoff condition.
Ask before work if target, operation, endpoint, success condition, approval boundary, date source, or verification path can change the result.

Approval-sensitive actions require exact target, expected effect, risk, recovery path, included and excluded scope, and endpoint.
Readiness, verification, build/readback, self-drive, previous context, and subagent output never authorize commit, push, PR, publish, release, version bump, destructive work, or external effects.

## Interruption

Use `interruption` only when a new user message arrives during an active flow.
It is an entry-only routing event, not a lifecycle phase.

Preserve current phase, scope, non-goals, approval boundary, verification status, and required next action.
Then apply `flow` contract-impact and route exactly one result:

- `inline-answer`
- `current-flow-revision`
- `background-current-flow`
- `reserve-later-analysis`
- `supersede-current-flow`
- `blocker-question`
- `explicit-stop`

Interruption never authorizes work outside the active contract or approval-sensitive execution.

## Verification

Record method and result separately.

Methods: `clean-context`, `normal`, `not-required`.
Results: `pass`, `fail`, `blocked`, `insufficient`.

`not-required` is a method, not a pass.
`not-started` and `requested` are progress states, not success evidence.

If a verifier, subagent, or tool returns `partial`, `mixed`, or `inconclusive`, reconcile it before success reporting:

- `pass`: remaining evidence supports the flow acceptance signal
- `insufficient`: evidence gaps or ambiguity remain
- `fail`: clear failure evidence exists
- `blocked`: input, access, approval, or external state is required

Default to `clean-context` for file changes, generated release surface changes, multi-file contracts, prior failures, requested QA/review/commit-readiness, and approval-sensitive boundaries.
Route `fail`, `insufficient`, and `blocked` before success reporting, self-drive continuation, release readiness, commit-readiness, or next-flow continuation.
Use `references/verification.md` for verifier packet boundaries and detailed status routing.

## Questions And Next Flow

After reporting, reopen routing unless explicit stop is recorded.
`next-flow` is an open routing state, not a final answer.
Use `request_user_input` for narrow choices when available; otherwise use active plain text.
Keep explicit turn-end available in record routing even when the visible question UI cannot show it.

If the user says "continue", "계속", or "이어가" after a report, inspect the recorded next action first.
Continue only when the next identity, target, scope, endpoint, approval boundary, and verification expectation are known.
Otherwise keep next-flow routing open and ask.
Post-flow continue does not activate self-drive and does not authorize approval-sensitive work.

An aborted, canceled, or interrupted question is not closure.
Record pending question state and route the next user message as answer, superseding request, status question, or explicit stop.
Use `references/question-routing.md` for fallback, blocker, and question recovery details.

## Self-Drive

Self-drive is an explicit prepared sequence overlay, not the default turn state.
Use `references/self-drive.md` only when records define objective, current flow or loop, allowed actions, approval checkpoints, endpoint, blockers, acceptance signal, and verification expectation.

At each self-drive flow or loop start, read `000-plan.md` and `000-self-drive.md`.
Do not advance from memory.

Finite mode advances only after pass verification, non-blocked `flow` handoff, known next `flow` identity, matching approval boundary, and a passing plan plus sidecar gate.
Infinite mode is counted bounded iteration, not an unlimited todo list.
Increment `loop_count` only after continuation remains valid.

Non-pass verification, repeated failure, blockers, approval need, stale sidecar state, and endpoint or scope changes stop autonomous advancement before continuation.
If endpoint, target, order, acceptance, approval, or scope changes, apply `flow` readiness/ambiguity to choose the earliest phase to relock.
Sequence completion is not terminal closure; report, update records, then route next-flow unless explicit stop is recorded.
