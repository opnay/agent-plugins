---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and next-flow reopening until the user explicitly stops; require sibling flow decisions before work, maintain session records, route questions with request_user_input, and apply self-drive only as a prepared sequence overlay.
---

# Turn Gate

## Important

When this skill is active, it is a conversation-level operating rule. Do not close the response as a terminal summary unless the user explicitly asks to end the current turn.

Required ending states are:

- reporting followed by `next-flow` reopening
- active question routing
- blocker routing
- prepared self-drive continuation
- source-recorded explicit turn stop

Use `request_user_input` for bounded choices when available. Keep `.agents/sessions/{YYYYMMDD}/` records current unless the user has forbidden all file writes or session records.

## Purpose

`turn-gate` keeps one active turn moving through:

1. `preparation`
2. `work`
3. `verification`
4. `reporting`
5. `next-flow`

It does not define flow boundaries, flow types, candidates, or flow completion logic. Those decisions belong to the sibling `flow` skill. `turn-gate` requires that decision, applies it, records it, and prevents work from starting without it.

## Phase Prefix

When telling the user that a phase is starting or reporting phase progress, start the message with `[<phase-name>(/<phase-protocol>)]`.

Use canonical phase labels:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

When a phase protocol is active, add it as a slash suffix, for example `[preparation/deep-interview]`, `[work/ralph-loop]`, or `[reporting/commit-readiness-gate]`. Do not include literal parentheses in real output.

The prefix applies to user-facing phase/progress messages. Do not mechanically add it to session record bodies, generated artifacts, command summaries, or question option labels.

## Preparation

Before work starts:

- confirm there is a source-recorded active flow, or request/apply a sibling `flow` decision
- record the chosen flow fields in the session record
- lock intent, scope, non-goals, acceptance signal, verification expectation, approval boundary, and expected risky actions enough for the current flow
- ask a user-gated question if scope, target, endpoint, approval, or current-flow identity changes the work
- use `update_plan` once meaningful work begins

If user wording is structurally ambiguous, resolve the operation and target before choosing a phase protocol or editing files. Examples include `merge`, `move`, `promote`, `delete`, `exclude`, `split`, `surface`, `skill`, `spec`, `contract`, or pronouns like `that` and `current one` when they could point to different artifacts.

Approval-sensitive actions need explicit target, expected effect, risk, recovery path, included/excluded scope, and endpoint. Readiness reporting is not approval to stage, commit, push, publish, release, or bump versions.

## Work

Work stays inside the active flow boundary and selected phase protocol. Read [references/phase-protocols.md](references/phase-protocols.md) when choosing or applying `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, or `commit-readiness-gate`.

Individual task completion does not complete the flow or close the turn. After work, continue to verification unless the flow is blocked before any verifiable output exists.

## Verification

Choose a method before reporting:

- `clean-context`: bounded read-only verifier packet. Default for file changes, release surfaces, templates, multi-file contracts, previous check failures, user-requested verification, and approval-sensitive work.
- `normal`: main-thread evidence review with commands, source readback, checklist, and counterexample review when risk is low.
- `not-required`: only when there is no work output to verify, such as activation-only, blocker-before-work, or next-flow selection.

Method is not status. Record result status separately as `pass`, `fail`, `blocked`, or `insufficient`. Do not treat `not-required`, `blocked`, `fail`, or `insufficient` as success.

A clean-context verifier receives only a bounded packet: target files or artifacts, user intent, changed surfaces, checks to inspect, pass/fail criteria, no edit permission, no scope expansion, no destructive or external actions, and no commit/push/PR/publish/release/version-bump authority.

## Reporting

Reporting is continuity context, not terminal closure. Before reporting, refresh the active flow record and its `Continuity Guard`.

Report:

- what was prepared, changed, checked, or blocked
- verification method and result status
- residual uncertainty
- material routing or approval judgments
- required next action

If verification is non-pass, route to the earliest safe repair phase or blocker question. Do not frame it as completed work.

## Next-Flow

After reporting, if there is no source-recorded explicit turn stop, enter `[next-flow]`.

Use `request_user_input` when a bounded choice is possible. Choices should be narrow and connected to the just-reported result. If the tool is unavailable, state that active question routing remains open and record the required next action.

Always record an explicit turn-end option in `Next Flow Options`, even if it is not visible in the user-facing choices.

Only these count as explicit turn stop: messages clearly ending the current turn, such as `turn 종료`, `여기서 끝`, `stop the turn`, or equivalent. Stale closure records or source-less `terminal_summary_allowed: yes` are not enough.

## Session Records

Use [references/session-records.md](references/session-records.md) for `.agents/sessions/{YYYYMMDD}/` structure, recovery rules, and templates. Runtime start files are:

- [templates/plan-template.md](templates/plan-template.md)
- [templates/flow-record-template.md](templates/flow-record-template.md)
- [templates/self-drive-template.md](templates/self-drive-template.md)

Records store the sibling `flow` decision fields and turn continuity state. They do not redefine how to judge whether an item is a flow, candidate, phase, or completed unit.

## Self-Drive

Self-drive is not a separate skill or mode. It is an overlay for an explicitly prepared finite sequence. Read [references/self-drive.md](references/self-drive.md) before applying it.

Self-drive narrows when to ask; it does not disable questions. Return to user-gated routing for approval, scope, endpoint, blocker, current-flow identity, record access, repeated critical failure, or sequence ambiguity.
