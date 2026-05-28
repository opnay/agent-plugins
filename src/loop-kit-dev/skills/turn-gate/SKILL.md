---
name: turn-gate
description: Keep an active Codex turn open until explicit stop; enforce flow lifecycle, session records, verification, interruption recovery, next-flow routing, and prepared self-drive.
---

# Turn Gate

## Active Turn

Keep the turn open until the user explicitly stops it and the stop source is recorded.
Completion, commits, passing checks, status answers, reports, interrupted questions, and final-looking summaries are not closure authority.

Every active flow ends in exactly one recorded state:

- `next-flow`: reporting is done, records are updated, and the next action is open.
- `blocked`: input, approval, access, or external state is required.
- `explicit-stop`: the current user message stops the turn and the stop source is recorded.

Maintain session records throughout the turn.
Use `references/session-records.md` for `000-plan.md`, flow records, phase checklists, compact metadata, and recovery rules.

If the user only activates turn-gate, record the operating state and open scope or next-flow routing.
Do not answer with a terminal activation summary.

## Flow Lifecycle

Run each active flow through this lifecycle:

1. `intake`: reread needed skills, separate source wording from interpretation, identify goal, non-goals, authority-sensitive signals, date assumptions, and discovery topics.
2. `framing`: apply `flow`; classify the item as active flow, parent flow, candidate, phase, or handoff.
3. `preparation`: lock scope, acceptance, verification expectation, approval boundary, and handoff.
4. `work`: act only inside the recorded flow boundary.
5. `verification`: choose method, run or justify it, and record result status.
6. `reporting`: update records first, then report continuity context.
7. `next-flow`: route a next action, blocker, self-drive continuation, or explicit stop.

Use phase prefixes for visible phase-start or meaningful progress messages:

- `[intake]`
- `[framing]`
- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Do not copy prefixes into artifacts, records, command summaries, or question option labels.

Update `000-plan.md` for turn-level routing changes.
Update the active flow record for same-flow phase state, evidence, residual risk, pending question state, and checklist changes.
For meaningful multi-step work, use the available plan tool to keep the current phase or task state visible.

## Intake Details

At a new flow start, reread the skills needed for that flow.
When a user message moves toward preparation, reread `turn-gate` and `flow`; keep both in `000-plan.md` `active_skills`.

Interpret relative dates such as today, tomorrow, yesterday, this week, last record, and previous flow from the current system date and timezone by default.
If a relative date affects the result, target, verification path, reporting scope, or later record reconstruction, write the absolute date.

If the request points to session records, prior flows, the last record, yesterday's work, or record-based resume, do not let record dates silently override the system date.
If the intended date source can change target, verification path, reporting scope, or record reconstruction, ask a user-gated clarification before work.
If the conflict does not affect the result, proceed with the system-date basis and record that basis briefly.

Use compact explicit wording such as `today (2026-05-28)`, `yesterday (2026-05-27, system-date basis)`, or `needs clarification: last record date or system-date yesterday`.

## Preparation And Approval

Before work, record:

- scope and non-goals
- completion or acceptance signal
- verification expectation
- approval boundary
- handoff condition

Ask before work if target, operation, endpoint, success condition, approval boundary, date source, or verification path can change the result.

Approval-sensitive actions require exact target, expected effect, risk, recovery path, included and excluded scope, and endpoint.
Readiness, verification, build/readback, self-drive, previous context, and subagent output never authorize commit, push, PR, publish, release, version bump, destructive work, or external effects.

## Interruption

Use `interruption` only when a new user message arrives during an active flow.
It is an entry-only routing event, not a lifecycle phase.

When interruption starts, preserve the current flow phase, scope, non-goals, approval boundary, verification status, and required next action.
Then choose one result:

- `inline-answer`: answer without changing the flow contract, then return to the preserved phase.
- `current-flow-revision`: update changed scope, non-goals, completion criteria, verification expectation, approval boundary, or handoff; return to `framing` or `preparation`.
- `background-current-flow`: keep the current flow resumable and start a new foreground flow.
- `reserve-later-analysis`: record a future topic, then return to the preserved phase.
- `supersede-current-flow`: mark the current flow superseded and start the replacement flow.
- `blocker-question`: mark the flow blocked and ask for the needed decision, access, approval, or scope.
- `explicit-stop`: record the stop source before terminal closure.

Classify short natural-language messages by effect:

- "summary only", "status only", and "why stopped?" usually stay `inline-answer`.
- "later" or "remember this" stays `reserve-later-analysis` when it does not change the current contract.
- "continue" may allow work inside the current contract, but it does not activate self-drive.
- If the same wording changes scope, endpoint, acceptance, verification, date basis, or approval boundary, use `current-flow-revision`.
- If it replaces the active work, use `supersede-current-flow` or `background-current-flow`.

Interruption never authorizes work outside the active contract or approval-sensitive execution.

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

`not-required` is a method, not a pass; record the reason and residual uncertainty.
`not-started` and `requested` are progress states, not success evidence.

If a verifier, subagent, or tool returns `partial`, `mixed`, or `inconclusive`, reconcile it before success reporting:

- `pass`: remaining evidence supports the acceptance signal
- `insufficient`: evidence gaps or ambiguity remain
- `fail`: clear failure evidence exists
- `blocked`: input, access, approval, or external state is required

Default to `clean-context` for file changes, generated release surface changes, multi-file contracts, prior failures, requested QA/review/commit-readiness, and approval-sensitive boundaries.
Route `fail`, `insufficient`, and `blocked` before success reporting, self-drive continuation, release readiness, commit-readiness, or next-flow continuation.

Use `references/verification.md` for verifier packet boundaries and detailed status routing.

## Reporting And Questions

Before reporting, update the active flow record and any required `000-plan.md` fields.
Report changed surfaces, verification status, material judgment calls, residual risk, and required next action.

After reporting, reopen routing unless explicit stop is recorded.
Use `request_user_input` for narrow choices when available; otherwise use active plain text.
Keep explicit turn-end available in record routing even when the visible question UI cannot show it.

An aborted, canceled, or interrupted question is not closure.
Record pending question state and route the next user message as:

- answer to the pending question
- superseding request
- status question
- explicit stop

If a free-form answer clearly gives a new task, mark the pending question `superseded` and prepare that flow.
If it selects an option and adds a note, record both.
Use `references/question-routing.md` for fallback, blocker, and question recovery details.

## Self-Drive

Self-drive is an explicit prepared sequence overlay, not the default turn state.
Use `references/self-drive.md` only when records define objective, current flow or loop, allowed actions, approval checkpoints, endpoint, blockers, acceptance signal, and verification expectation.

At each self-drive flow or loop start, read `000-plan.md` and `000-self-drive.md`.
Do not advance from memory.

Finite mode advances only after current verification is pass, handoff is not blocked, next identity is known, approval boundary matches, and plan plus sidecar gate pass again.
Infinite mode is counted bounded iteration, not an unlimited todo list.
Increment `loop_count` only after continuation remains valid.

Non-pass verification, repeated failure, blockers, approval need, stale sidecar state, and endpoint or scope changes stop autonomous advancement before continuation.
Sequence completion is not terminal closure; report, update records, then route next-flow unless explicit stop is recorded.
