---
name: turn-gate
description: Use to keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn, while enforcing the sibling flow contract and maintaining session records.
---

# turn-gate

## Important

When this skill is active, it is a conversation-level operating rule for the current session. Do not close with a terminal summary unless the user has explicitly stopped the turn and that stop is source-recorded.

Every active flow must end in one of these states:

- `next-flow`: result reported, records refreshed, and a next-flow choice or self-drive continuation opened.
- `blocked`: a blocker or missing approval is reported through user-gated routing.
- `explicit-stop`: the user explicitly ended the turn, with the closure source recorded.

After reporting, use `request_user_input` for next-flow reopening when structured choices are possible and the tool is available. Keep `.agents/sessions/{YYYYMMDD}/000-plan.md`, the active `001+` flow record, and any active self-drive sidecar current enough to explain the next required action.

## Purpose

Use `turn-gate` to keep one turn continuous while work moves through explicit flow records. This skill does not define flow boundaries itself. It applies the sibling `flow` contract before work starts, records the selected active flow, and reopens the next flow after reporting.

Runtime support files live beside this skill:

- `references/session-records.md`
- `references/self-drive.md`
- `templates/plan-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`

## Operating Cycle

Run every active flow in this order:

1. `preparation`: lock the active flow before work. Apply the sibling `flow` decision for flow label, type, scope, non-goals, completion criteria, verification expectation, handoff condition, missing fields, question topics, and flow-local strategy. If scope is empty, too broad, multi-output, ambiguous, or could change the success or verification path, ask before work.
2. `work`: execute only inside the active flow boundary. A task finishing is not enough to finish the flow or close the turn.
3. `verification`: choose and record a method and a result status before reporting.
4. `reporting`: report continuity context, not a terminal close. Refresh the active flow record and Continuity Guard first.
5. `next-flow`: reopen the next flow through `request_user_input`, fallback active question-routing, or a valid self-drive continuation.

## Phase Messages

When telling the user a phase is starting or giving phase-start progress, begin the user-facing message with the canonical phase prefix:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Use the literal bracketed label, not parenthesized text. Do not append flow-local strategy names as turn-gate protocol labels. The prefix marks phase-start/progress messages; do not mechanically add it inside flow records, generated artifacts, command summaries, or question option labels.

Activation-only messages start with `[preparation]` for scope setup unless immediately opening choices with `[next-flow]`. Mid-work status normally uses the current active phase. Record access blockers use the phase where the blocker is discovered: `[reporting]` before result reporting, `[next-flow]` before reopening. Report-only evaluation still proceeds to `[next-flow]` unless the user explicitly stops.

For self-drive continuation, prefix user-facing status, verification, reporting, and automatic handoff messages. Do not propagate prefixes into `000-self-drive.md`, flow records, generated artifact bodies, or question choices.

## Preparation Rules

Before work, record the active flow or candidate decision from `flow`. Keep the flow contract separate from phase notes. If the request could mean different operations or targets, especially words like `merge`, `absorb`, `move`, `promote`, `remove`, `delete`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, `contract`, or pronouns with multiple possible targets, resolve that ambiguity before work.

If you infer scope without asking, still record the work boundary and non-goals. Preparation must also capture intent, scope, non-goal, acceptance signal, verification expectation, approval boundary, and handoff condition for the active flow.

Approval-sensitive actions require exact target, expected effect, risk, rollback or recovery path, included/excluded scope, and endpoint. Readiness reporting is not execution authority. Commit, push, PR, publish, release, and version bump steps require an explicit recorded approval boundary.

If the user explicitly requests self-drive over a prepared sequence, read `references/self-drive.md` and maintain `000-self-drive.md` with `templates/self-drive-template.md`.

## Verification Rules

Record verification method separately from result status.

Methods:

- `clean-context`: a bounded read-only verifier packet, not a full-history fork.
- `normal`: main-thread checks, readback, evidence review, or logical counterexample review.
- `not-required`: no separate verification action is needed; record the reason and residual uncertainty.

Statuses:

- `pass`
- `fail`
- `blocked`
- `insufficient`

Use `clean-context` by default when files changed, release surfaces changed, multiple-file contracts changed, there is prior check failure, the user requested verification/review/QA/commit-readiness, or approval-sensitive action is involved.

Never treat `not-required` as a successful result. Do not use it for file changes, release surface changes, multiple-file contracts, prior failures, user-requested verification, or approval-sensitive action.

A verifier packet must include target, user intent, changed files or artifacts, checks or evidence to inspect, pass/fail criteria, no edit permission, no scope expansion, no destructive/external work, and no commit/push/PR/publish/release/version-bump action.

Before endpoint exhaustion or self-drive sequence completion, handle non-pass verification first. `fail` returns to the earliest safe repair/work point, `insufficient` returns to evidence repair, and `blocked` opens user-gated blocker routing. Non-pass states are not successful completion or next-flow authority.

## Reporting And Next Flow

Reporting summarizes what matters for continuation: changed surfaces, verification status, material judgment calls, residual risk, and required next action. Before reporting or reopening next-flow, read and refresh the active flow record's Continuity Guard.

Only a source-recorded explicit stop allows terminal close. Stale `terminal_summary_allowed: yes`, source-less closure, or closure that does not match the current incoming user message is not enough.

After reporting, reopen the next flow. Use `request_user_input` when available. If it is unavailable, state that the tool is unavailable, list the active choices in text, and record the required next action. Even when the visible choices omit a turn-end option, record an explicit turn-end option in `Next Flow Options`.
