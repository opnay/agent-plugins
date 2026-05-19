---
name: turn-gate
description: Use to keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn, while enforcing the sibling flow contract and maintaining session records.
---

# turn-gate

## Important

When this skill is active, it is a conversation-level first-class rule. Do not close the response as a terminal summary unless the user explicitly stops the turn and that stop is source-recorded.

Every concrete task must stay inside an active flow. Apply the sibling `flow` contract for flow boundary, readiness, discovery, ambiguity, flow-local strategy, and handoff condition. Do not redefine those rules in `turn-gate`.

After reporting, move into `next-flow` unless a source-recorded explicit stop permits terminal closure. Use `request_user_input` for structured next-flow choices when available. If it is unavailable, leave the turn in an active question-routing state and report the fallback clearly.

Maintain `.agents/sessions/{YYYYMMDD}/` records for turn continuity unless the user explicitly forbids all writes or session records. Read `references/session-records.md` when creating or updating records, and use `templates/` files for new records.

## Purpose

`turn-gate` keeps the current turn open across one or more flows. It owns turn continuity, active flow enforcement, reporting as continuity context, next-flow reopening, explicit stop handling, session records, verification method selection, and approval boundary handling.

It does not own flow-local strategy. Requirement discovery, review-loop, fix-verify-loop, broad-execution, and commit-readiness are `flow` decisions that `turn-gate` applies inside the active flow.

## Operating Cycle

Run each active flow through this order:

1. `preparation`
2. `work`
3. `verification`
4. `reporting`
5. `next-flow`

Individual task completion does not prove flow completion or turn closure. Reporting is context for continuation, not a clean stop.

## Phase Prefixes

When a user-facing message starts a phase or reports progress at phase start, begin it with the canonical phase prefix:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Use the literal labels above without parenthesized placeholders. Do not add flow-local strategy suffixes such as review-loop names to the prefix. The prefix applies to phase-start/progress messages, not mechanically to session records, generated artifacts, command summaries, or every sentence in a question option.

For activation-only requests, start with `[preparation]` to set scope, then use `[next-flow]` only when opening concrete choices. For mid-work status, use the current active phase, usually `[work]`. For record blockers, use the phase where the blocker is found, commonly `[reporting]` before result reporting or `[next-flow]` before reopening.

Self-drive status, verification, reporting, and automatic handoff messages still use the current phase prefix, but do not copy prefixes into `000-self-drive.md`, flow records, generated artifact bodies, or question option labels.

## Preparation

Before work, ensure there is a source-recorded active flow or obtain and apply a sibling `flow` decision.

Lock or record:

- intent, scope, and non-goals
- acceptance signal and completion criteria
- verification expectation and likely verification method
- approval boundary and expected approval-sensitive checkpoints
- active flow versus sub-flow candidate status
- missing contract fields, recommended question topics, and unresolved ambiguity from `flow`
- interpreted operation and target when wording is ambiguous

If scope is empty, too broad, able to produce multiple valid outputs, or able to change success criteria or verification path, do not proceed to work. Use user-gated question routing.

Meaning resolution and approval are different. First lock what the user means by an operation or target. Then, for destructive, irreversible, external, commit, push, PR, publish, release, or version-bump actions, confirm exact target, expected effect, risk, recovery path, include/exclude scope, and endpoint before execution.

Use `update_plan` once meaningful multi-step work starts. Keep the current phase state accurate.

## Work

Perform only the work inside the active flow boundary. Apply the sibling `flow` strategy as execution policy, but do not treat a strategy as a turn mode or next-flow authority.

If work uncovers new scope, changed non-goals, a new approval boundary, a destructive/external action, or a blocker that changes completion criteria, return to preparation or user-gated routing instead of widening the flow silently.

## Verification

Before reporting, select and perform a verification method:

- `clean-context`: bounded read-only verifier subagent. Default for file changes, release surfaces, manifest/template/scenario fixture/build output changes, multi-file contracts, prior failed checks, user-requested QA/review/commit-readiness, and approval-sensitive work.
- `normal`: main-agent evidence such as command/check output, source readback, checklist, and logical counterexample review. Use for low-risk no-edit/read-only work or when already gathered evidence is sufficient.
- `not-required`: only when there is no work output to verify, such as activation-only, blocker-before-work, next-flow selection, or routing-only results. Record the reason and residual uncertainty.

Verification method is not result status. Result status is `pass`, `fail`, `blocked`, or `insufficient`. Do not report `fail`, `blocked`, or `insufficient` as successful completion. Route to the earliest safe repair phase or user-gated blocker decision.

A `clean-context` verifier packet is not a full-history fork. Send only bounded facts: target files or artifacts, user intent, change summary, checks or readback to inspect, pass/fail criteria, no edit permission, no scope expansion, no destructive/external work, and no commit/push/PR/publish/release/version-bump authority.

## Reporting

Reporting summarizes the active flow so the turn can continue. Include what was prepared, worked, and verified; changed surfaces; blockers; residual uncertainty; and verification method/status.

Before reporting and before next-flow reopening, refresh the active flow record's Continuity Guard. If records are missing, inaccessible, stale, or contradictory, do not infer terminal closure. Use the record recovery rules in `references/session-records.md`.

## Next-Flow

After reporting, check for a source-recorded explicit stop. Only clear requests such as "end this turn", "stop the turn", or equivalent current-turn closure language permit terminal closure.

If there is no explicit stop, open the next flow:

- Prefer `request_user_input` with narrow options tied to the report.
- Include or record a turn-end option so the user can explicitly stop the turn.
- If tool-based questions are unavailable, state that fallback, list the open choices, and leave required next action active.
- If self-drive is active and the prepared sequence still identifies the next flow, the `next-flow` result may be recorded loop continuation instead of a user question.

Plain closing phrases do not replace next-flow reopening.

## Session Records

Use `.agents/sessions/{YYYYMMDD}/000-plan.md` for the date-level index and active snapshot. Use `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for each flow record. Use `.agents/sessions/{YYYYMMDD}/000-self-drive.md` only when self-drive is active.

For new records, use:

- `templates/plan-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`

For detailed record ownership, read `references/session-records.md`.

## Self-Drive Overlay

Self-drive is not a separate installed entrypoint and not a default mode label. It is an overlay applied only to a prepared flow sequence with recorded scope, non-goals, acceptance signal, approval boundary, verification expectation, endpoint, and blocker return conditions.

When the user explicitly asks for self-drive or the current prepared sequence already has active self-drive state, read `references/self-drive.md`. That reference owns sequence continuation decisions while reusing this skill's explicit stop, approval boundary, session record, verification, and reporting rules.

## Runtime Resources

Use only installed runtime resources from this skill folder:

- `references/session-records.md`
- `references/self-drive.md`
- `templates/plan-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`

Do not require runtime users to read development specs or spec-side fixtures.
