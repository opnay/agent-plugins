---
name: turn-gate
description: Keep an active Codex turn open until explicit stop; apply the flow skill as-is; route flow handoff through next-flow questions or prepared self-drive.
---

# Turn Gate

## Wrapper Loop

Keep the active Codex turn open until the user explicitly stops it and that stop source is recorded.
Completion, commits, passing checks, status answers, reports, interrupted questions, final-looking summaries, and final responses are not closure authority.
Treat every new user message as open input inside the active turn by default. Questions, status checks, task changes, direction changes, corrections, and follow-up requests are not stop signals unless the user explicitly says to end the active turn.

Apply `flow` as-is.
Do not define or restate flow taxonomy, lifecycle, readiness, discovery, ambiguity, contract impact, flow-local strategy, template meaning, or handoff meaning.

Use this wrapper loop:

1. Route the user message into the `flow skill` group inside the `turn-gate` wrapper.
2. Treat the internal path from interview to handoff as `생략...`; do not model it inside `turn-gate`.
3. From `flow skill: handoff`, run `next-flow gate`:
   - run the `skill reconfigure` group: identify the full session active skill list, reread each skill body, and accept them as the new active skill set
   - select the next flow through the question tool or a prepared self-drive gate
   - if self-drive is active, update `000-self-drive.md`
   - update `000-plan.md`
4. Use the same interview flow as `flow: deep-interview` to clarify the next flow input.
5. Reenter `flow skill: interview` with the clarified input and prioritize flow-design questions.

If `request_user_input` is available and choices are narrow, use it for question-tool next-flow selection.
If unavailable, keep the turn open with an active plain-text question.
Update `000-plan.md` every time `next-flow gate` runs, including active skills, selected or pending question state, next action, and self-drive pointer when relevant.

Self-drive is not a graph node.
Use it only when an explicit prepared sequence gate can replace the question tool.

## Phase Prefixes

Start user-facing phase-start or meaningful progress messages with the current phase prefix.
Apply phase prefixes produced by `flow`; do not redefine flow-owned phase labels inside `turn-gate`.
Use `[next-flow]` for post-handoff next-flow questions, question tool opening, or self-drive continuation.

Do not mechanically copy phase prefixes into artifact bodies, records, raw command output, command summaries, or question option labels.
Do not prefix every sentence or bullet inside an already prefixed user-facing message.

## Records

Use `references/session-records.md` only to keep active-turn routing recoverable.
Records support the wrapper loop; they are not extra graph nodes.
Use `turn-gate/templates/self-drive-template.md` only for self-drive sidecar state.

Do not treat readiness, verification, generated release surface, previous context, self-drive, or subagent output as authority for commit, push, PR, publish, release, version bump, destructive work, or external effects.

## Questions

After `flow skill: handoff`, run `skill reconfigure`, reopen routing unless explicit stop is recorded, update `000-self-drive.md` when self-drive is active, then update `000-plan.md`.
The next routing surface is not terminal closeout.

If a question tool call is aborted, canceled, or interrupted, preserve the pending question and treat the next user message as one of:

- answer to the pending question
- superseding new flow request
- status/progress question
- explicit-stop

The non-stop cases above keep the turn open. If the new input is not concrete enough to continue, return to the same interview flow to clarify it instead of closing the turn.

If the user says "continue", "계속", or "이어가", continue only when the next flow input is already concrete enough to reenter `flow skill: interview`.

## Interruption

When a new user message arrives during an active turn, preserve pending question, approval boundary, verification status, and required next action.
Apply `flow` contract-impact and route the result as `active-flow`, `current-flow-revision`, `background-current-flow`, `reserve-later-analysis`, `supersede-current-flow`, `blocker-question`, or `explicit-stop`.
Use `explicit-stop` only when the current message clearly asks to end the active turn.
Interruption never authorizes work outside the active contract or approval-sensitive execution.

## Verification

Record verification method and result separately.

Methods: `clean-context`, `normal`, `not-required`.
Results: `pass`, `fail`, `blocked`, `insufficient`.

`not-required` is a method, not a pass.
`not-started` and `requested` are progress states, not success evidence.
Route non-pass results before question routing, self-drive continuation, release readiness, commit-readiness, or handoff routing.
Use `references/verification.md` for verifier packet boundaries and detailed status routing.

## Self-Drive

Self-drive is an explicit prepared sequence overlay, not the default turn state.
Use `references/self-drive.md` only when records define objective, current flow or loop, allowed actions, approval checkpoints, endpoint, blockers, acceptance signal, and verification expectation.

At each self-drive step, read `000-plan.md` and `000-self-drive.md`.
Replace the question tool only after pass verification, non-blocked `flow` handoff, known next identity, matching approval boundary, and a passing sidecar gate.

Non-pass verification, blockers, approval need, stale sidecar state, endpoint changes, or scope changes stop autonomous advancement before continuation.
