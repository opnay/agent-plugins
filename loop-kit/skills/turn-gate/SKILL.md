---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# Turn Gate

## Important

When this skill is active, treat it as a conversation-level operating rule for the whole session. Do not close with a terminal summary after result reporting unless the current user message, or a flow record tied to that exact source message, explicitly ends the turn.

Required ending states are one of:

- continue into the next flow through active question routing;
- return to preparation because scope, target, approval boundary, or verification expectation is not locked;
- report a blocker and ask for a user-gated decision;
- close only when a source-recorded explicit stop says to end this turn.

After result reporting, enter `next-flow` unless explicit stop is confirmed. Use `request_user_input` for next-flow reopening whenever bounded choices are available. If the tool is unavailable, leave an active plain-text question and record the required next action.

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` while a turn-gated task is active. Refresh the active flow record's Continuity Guard before reporting and before next-flow reopening.

Use only runtime files bundled with this skill during execution, such as `references/*` and `templates/*`. Do not depend on development specs being available at runtime.

## Operating Cycle

Run each active flow through:

1. preparation
2. work
3. verification
4. reporting
5. next-flow

Activation without a concrete task opens scope setup or next-flow selection; it does not end with an activation summary. Individual task completion does not complete the flow or permit turn closure by itself.

`turn-gate` normally runs in an implicit default operating state. Phase protocols shape the current phase; they are not standalone modes. Before meaningful work, choose the current phase protocol from `references/phase-protocols.md`. If the user explicitly asks for autonomous continuation across a prepared flow sequence, read `references/self-drive.md`; that reference owns the self-drive overlay's continuation rules.

## Phase Start Prefix

When you tell the user that a phase is starting, start that user-facing message with `[<phase-name>(/<phase-protocol>)]`.

- Canonical phase labels are `preparation`, `work`, `verification`, `reporting`, and `next-flow`.
- The `(/<phase-protocol>)` segment is optional notation. In actual output, use a slash suffix when a phase protocol is active and do not print literal parentheses.
- Valid examples include `[preparation]`, `[work]`, `[verification]`, `[reporting]`, `[next-flow]`, `[preparation/deep-interview]`, `[work/ralph-loop]`, `[verification/review-loop]`, and `[reporting/commit-readiness-gate]`.
- Apply the prefix to phase-start messages, not to every command summary, artifact body, flow record line, command output summary, or question option.
- For activation-only requests with no concrete task, start with `[preparation]` for scope setup. Use `[next-flow]` only when opening actual next-flow choices.
- For status questions, use the current active phase. During work, this is usually `[work]`.
- For session-record access blockers, use the phase where the blocker was found, usually `[reporting]` or `[next-flow]`.
- For report-only evaluation, gather evidence and report without edits if appropriate, then continue to `[next-flow]` unless explicit stop is confirmed.

## Preparation

Before work, lock the active flow enough to execute safely:

- intent, scope, non-goals, acceptance signal, and verification expectation;
- approval boundary for destructive, irreversible, external, commit, push, PR, publish, release, or version-bump actions;
- operation and target meaning when user wording can point to multiple files, surfaces, phases, routing rules, or ownership changes;
- whether this is an `operational-preparation` flow or a `change-unit` flow.

Use deep-interview, flow list design, meaning resolution, current-state inspection, target reread, and scope lock as preparation techniques. Ask before work when scope is empty, too broad, ambiguous, likely to create multiple outputs, or likely to change the verification path. Prefer `request_user_input` for bounded choices.

If you infer scope without asking, still record the work boundary and non-goals in the flow record. If the current work is interpreting a request, designing a planned flow list, or collecting approvals, treat that as an `operational-preparation` flow and keep follow-up `change-unit` candidates separate from active execution.

For approval-sensitive actions, record exact target, expected effect, risk, rollback or recovery path, included and excluded scope, and end point before execution. Readiness reporting, self-drive, closure wording, and next-flow routing are not authority to stage, commit, push, open a PR, publish, release, bump a version, or run any other destructive or external action.

When meaningful work starts, keep the visible task plan current with `update_plan`.

## Work

Work only inside the active flow boundary. A flow is a cohesive reviewable or commit-sized unit, not a checklist of phases such as analysis, implementation, verification, and reporting.

Before work, select the needed phase protocol from `references/phase-protocols.md`. Use the earliest blocker as the routing basis. If no protocol applies, stay in the default operating state without a protocol suffix.

Keep follow-up candidates, broader refactors, unrelated plugins, and new approval-sensitive work out of the active flow unless the user explicitly selects or approves them.

## Verification Method

After work and before reporting, choose one verification method and keep it separate from result status.

- `clean-context`: a bounded read-only verifier subagent checks the flow from an independent packet. This is not a full-history fork.
- `normal`: the main agent verifies in the same context using commands/checks, source readback, evidence checklists, log review, and logical counterexample review.
- `not-required`: no separate verification action is needed because there is no work output to verify. This is a method judgment, not a successful result.

Use `clean-context` by default for file changes, release surfaces, manifests, templates, scenario fixtures, build output, multi-file contracts, previous failed checks, user-requested verification/review/QA/commit-readiness, and destructive, irreversible, external, commit, push, PR, publish, release, or version-bump preparation or execution.

Use `normal` for lower-risk no-edit or read-only work, single state checks, narrow explanations, already-evidenced work, or cases where command/check/source readback and logical review are enough. Record why `clean-context` was not needed and what uncertainty remains.

Use `not-required` only for blocker-before-work, activation-only, next-flow selection, scope routing, or no-output routing cases. Record the omission reason, already-known evidence or no-output rationale, and residual uncertainty. Do not use `not-required` for file changes, release surfaces, multi-file contracts, failed checks, user-requested verification, or approval-sensitive actions.

Result status is separate from method. Use `pass`, `fail`, `blocked`, or `insufficient` after verification; lifecycle records may also use `not-started` or `requested` before final result. Never treat `not-required`, `blocked`, `fail`, or `insufficient` as automatic pass.

## Clean-Context Packet

Clean-context verification is pre-authorized only within a read-only verification boundary. The verifier packet must include:

- verifier identity or request id;
- verification target and expected user intent;
- changed files or produced artifacts;
- already captured evidence, including commands/checks run during work when that evidence is still valid;
- commands or checks to run or review, scoped to gaps in the evidence;
- pass/fail criteria;
- no edit permission, no scope expansion, no destructive or external actions, and no commit/push/PR/publish/release/version-bump authority;
- stop condition.

Do not automatically ask the verifier to rerun the same command or check when the same evidence was already captured during work and is complete, current, and tied to the final changed state. Ask the verifier to read back recorded evidence first and request rerun or blocker reporting only when evidence is stale, incomplete, suspicious, or misses a changed path.

If a verifier would need edit permission, scope expansion, destructive/external action, or commit/push/PR/publish/release/version-bump authority, stop and route to a user-gated question.

## Result Handling

Do not report `fail`, `blocked`, or `insufficient` as successful completion. Return to the earliest safe phase for repair, or open a user-gated blocker when verification cannot be completed.

For report-only work with no file changes, `normal` verification may focus on source/evidence readback, logical counterexamples, user-intent fit, and missing-risk checks instead of unnecessary command execution.

## Reporting

Reporting is continuity context for the next flow, not a terminal close. Report:

- what was prepared, changed, checked, or decided;
- verification method, status, and evidence;
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
