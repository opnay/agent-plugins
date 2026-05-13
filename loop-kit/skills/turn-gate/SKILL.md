---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# Turn Gate

## Important

When this skill is active, treat it as a conversation-level operating rule for the whole session. Do not close with a terminal summary after reporting results unless the current user message explicitly ends the turn.

Required ending states are one of:

- continue into the next flow through active question routing;
- return to preparation because scope, target, approval, or verification is not locked;
- report a blocker and ask for a user-gated decision;
- close only when a source-recorded explicit stop says to end this turn.

After result reporting, use `request_user_input` for next-flow reopening whenever structured choices are available. If the tool is unavailable, leave an active plain-text question and record the required next action.

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` while a turn-gated task is active. Refresh the flow record's Continuity Guard before reporting and before next-flow reopening.

## Purpose

Use `turn-gate` as the main controller for turn-level continuity. Each active flow moves through:

1. preparation
2. work
3. verification
4. reporting
5. next-flow

Activation without a concrete task opens scope setup or next-flow selection; it does not end with an activation summary.

## Phase Start Prefix

When you tell the user that a phase is starting, start that user-facing message with `[<phase-name>(/<phase-protocol>)]`.

- Canonical phase labels are `preparation`, `work`, `verification`, `reporting`, and `next-flow`.
- The `(/<phase-protocol>)` segment is optional notation only. In actual output, use a slash suffix without literal parentheses.
- Valid examples include `[preparation]`, `[work]`, `[verification]`, `[reporting]`, `[next-flow]`, `[preparation/deep-interview]`, `[work/ralph-loop]`, and `[reporting/commit-readiness-gate]`.
- Apply the prefix to phase-start messages, not to every command summary, artifact body, flow record line, or question option.
- For activation-only requests with no concrete task, start with `[preparation]` for scope setup. Use `[next-flow]` only when opening actual next-flow choices.
- For status questions, use the current active phase. During work, this is usually `[work]`.
- For session-record access blockers, use the phase where the blocker was found, usually `[reporting]` or `[next-flow]`.
- For report-only evaluation, gather evidence and report without edits if appropriate, then continue to `[next-flow]` unless explicit stop is confirmed.

## Preparation

Before work, lock the active flow enough to execute safely:

- intent, scope, non-goals, acceptance signal, and verification expectation;
- approval boundary for destructive, irreversible, external, commit, push, PR, publish, release, or version-bump actions;
- operation and target meaning when user wording can point to multiple files, surfaces, phases, or ownership changes;
- whether this is an `operational-preparation` flow or a `change-unit` flow.

Use user-gated question routing when scope is empty, too broad, ambiguous, or likely to change the output or verification path. Prefer `request_user_input` for bounded choices.

If you infer scope without asking, still record the work boundary and non-goals in the flow record. If the current work is only interpreting a request, designing a planned flow list, or collecting approvals, treat that as an `operational-preparation` flow and keep follow-up `change-unit` candidates separate from active execution.

For approval-sensitive actions, record exact target, expected effect, risk, rollback or recovery path, included and excluded scope, and end point before execution. Readiness reporting is evidence only; it is not authority to stage, commit, push, open a PR, publish, release, bump a version, or run any other external action.

When meaningful work starts, keep the visible task plan current with `update_plan`.

## Work

Work only inside the active flow boundary. A flow is a cohesive reviewable or commit-sized unit, not a checklist of phases such as analysis, implementation, verification, and reporting.

Choose the current phase protocol from `references/phase-protocols.md` before meaningful work. Phase protocols are not standalone modes; they shape how the current phase runs inside the default operating state.

If the user explicitly asks for autonomous continuation across a prepared flow sequence, read `references/self-drive.md`. The self-drive reference owns that overlay's continuation rules; do not repeat or improvise them from memory.

Individual task completion does not by itself complete the flow or allow turn closure.

## Verification

After work and before reporting, perform clean-context verification with a bounded read-only verifier subagent unless the task is blocked before verification.

The verifier packet must include:

- verifier identity or request id;
- verification target and expected user intent;
- changed files or produced artifacts;
- commands or checks to run;
- pass/fail criteria;
- no edit permission, no scope expansion, no destructive or external actions, and no commit/push/PR/publish/release/version-bump authority;
- stop condition.

Clean context means a bounded verification packet, not a full-history fork. Treat results as `pass`, `fail`, `blocked`, or `insufficient`. Do not report `fail`, `blocked`, or `insufficient` as successful completion. Route non-pass results to the earliest safe phase or to user-gated blocker handling.

## Reporting

Reporting is continuity context for the next flow, not a terminal close. Report:

- what was prepared, changed, checked, or decided;
- verification status and evidence;
- material judgment calls that affected routing or phase selection;
- residual uncertainty, blockers, and risks;
- changed surfaces when applicable.

Before reporting, refresh the active flow record and Continuity Guard. Terminal summary is allowed only when the current incoming message or a source-recorded closure in the flow record explicitly says to stop the turn.

## Next Flow

After reporting, enter `next-flow` unless explicit stop is confirmed.

Use `request_user_input` for narrow choices connected to the result just reported. Include the next useful flow options, blocker choices, or scope decisions. If visible choices cannot include a turn-end option, still state that the user can explicitly stop the turn and record an explicit turn-end option in the flow record.

If `request_user_input` is unavailable, ask a plain-text active question, state that the tool was unavailable, and record the open choices and required next action.

## Records And Templates

Use these bundled resources when session records are needed:

- `references/session-records.md` for record ownership, Continuity Guard, and next-flow option rules.
- `templates/plan-template.md` for `.agents/sessions/{YYYYMMDD}/000-plan.md`.
- `templates/flow-record-template.md` for `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`.

Do not silently reconstruct inaccessible records. Report record access failure as a blocker and do not use missing or stale closure state as a reason to close the turn.
