---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# Turn Gate

## Important

When this skill is active, treat it as a conversation-level operating rule for the whole session. `turn-gate` controls the response loop itself; it is not an internal checklist that can be dropped after a task appears complete.

Do not close with a terminal summary after result reporting unless the current user message, or a source-recorded explicit stop in the active flow record that matches the current message, explicitly ends the turn.

Every assistant response while this skill is active must end in one of these states:

- loop continuation into the next phase or next flow;
- user-gated question routing for scope, target, approval boundary, blocker, or next-flow selection;
- explicit stop closure, only when a source-recorded user stop allows terminal summary.

Use the question tool for bounded choices, scope locks, approval checkpoints, blocker decisions, and next-flow selection whenever it is available. Use the plan tool once meaningful work starts and keep it aligned with the current phase.

Maintain session records under `.agents/sessions/{YYYYMMDD}/` for active turn-gated work. Refresh the active flow record's Continuity Guard before reporting and before next-flow reopening.

Use only runtime files bundled with this skill, such as `references/*` and `templates/*`. Do not depend on development specs or repository-only spec paths being available at runtime.

## Purpose

Use this skill to keep a turn open across preparation, work, verification, reporting, and next-flow selection. The default state is implicit: phase protocols shape the current phase, but they are not separate modes.

If the user explicitly asks for autonomous continuation across a prepared planned flow sequence, or if an already-active autonomous sequence receives another user message, read `references/self-drive.md`. That reference owns the self-drive overlay's continuation and interruption rules.

## Operating Cycle

Run each active flow in this order:

1. preparation
2. work
3. verification
4. reporting
5. next-flow

Activation without a concrete task opens scope setup or next-flow selection. It does not end with an activation summary.

Individual task completion does not complete the flow, skip reporting, skip next-flow, or permit turn closure by itself. A flow is complete only when its completion criteria, verification expectation, reporting, and continuation state are handled.

## Phase Start Prefix

When a user-facing message starts a phase or announces phase progress, begin it with `[<phase-name>(/<phase-protocol>)]`.

- Canonical phase labels are `preparation`, `work`, `verification`, `reporting`, and `next-flow`.
- The `(/<phase-protocol>)` segment is notation only. In actual output, use a slash suffix when a phase protocol is active and do not print literal parentheses.
- Valid examples include `[preparation]`, `[work]`, `[verification]`, `[reporting]`, `[next-flow]`, `[preparation/deep-interview]`, `[work/ralph-loop]`, `[verification/review-loop]`, and `[reporting/commit-readiness-gate]`.
- Apply the prefix to phase-start messages, not to every command summary, artifact body, flow record line, command output summary, or question option.
- For activation-only requests with no concrete task, start with `[preparation]` for scope setup. Use `[next-flow]` only when opening actual next-flow choices.
- For status questions, use the current active phase. During work, this is usually `[work]`; use `[reporting]` only when intentionally summarizing flow context.
- For self-drive continuation, prefix user-facing status, verification, reporting, and automatic next-flow handoff messages with the phase being announced. Do not add phase prefixes inside the self-drive record, flow record, generated artifact body, or option labels.
- For session-record access blockers, use the phase where the blocker was found, usually `[reporting]` or `[next-flow]`.
- For report-only evaluation, gather evidence and report without edits if appropriate, then continue to `[next-flow]` unless explicit stop is confirmed.

## Preparation

Before work, lock the active flow enough to execute safely:

- intent, scope, non-goals, acceptance signal, and verification expectation;
- approval boundary for destructive, irreversible, external, commit, push, PR, publish, release, or version-bump actions;
- operation and target meaning when user wording can point to multiple files, surfaces, phases, routing rules, ownership changes, or provenance/intent blocks;
- whether this is an `operational-preparation` flow or a `change-unit` flow.

Treat structurally ambiguous operation words as meaning-resolution triggers before editing. High-signal examples include `merge`, `absorb`, `promote`, `remove`, `split`, `route`, `병합`, `흡수`, `녹여`, `승격`, `정식 규칙화`, `삭제`, `빼`, `그 밑`, `출처`, and `의도` when they could point to different files, sections, ownership boundaries, or destructive effects.

Use deep-interview, flow list design, meaning resolution, current-state inspection, target reread, and scope lock as preparation techniques. Ask before work when scope is empty, too broad, ambiguous, likely to create multiple outputs, or likely to change the verification path. Prefer the question tool for bounded choices.

If you infer scope without asking, record the work boundary and non-goals in the flow record. If the current work is interpreting a request, designing a planned flow list, or collecting approvals, treat it as an `operational-preparation` flow and keep follow-up `change-unit` candidates separate from active execution.

For approval-sensitive actions, record exact action, target, expected effect, risk, rollback or recovery path, included scope, excluded scope, and end point before execution. Meaning resolution, readiness reporting, self-drive, closure wording, and next-flow routing are not authority to stage, commit, push, open a PR, publish, release, bump a version, or run any destructive or external action.

Approval-sensitive execution is allowed only when the exact boundary is already source-recorded or the user grants it through a user-gated approval checkpoint. If the boundary is missing or ambiguous, do not proceed to work; route to a question.

## Work

Work only inside the active flow boundary. A flow is a cohesive reviewable or commit-sized unit, not a checklist of phases such as analysis, implementation, verification, and reporting.

Use these common misclassification checks before shaping work:

| Input shape | Correct handling |
| --- | --- |
| `analysis`, `work`, `verification`, `reporting`, or `commit readiness` listed as separate planned flows | Treat them as phases or handoffs unless they create a separate reviewable artifact. |
| Commit completed | Report and reopen next-flow; it is not a turn stop. |
| Status/progress question during self-drive | Report status and continue the recorded sequence unless it also changes scope, endpoint, approval, or blocker state. |
| Future stop such as "stop when the list is exhausted" | Update the endpoint; do not treat it as immediate terminal closure. |
| Any file, release surface, manifest, template, or multi-file contract change | Default to `clean-context` verification. |

For planned flow examples, prefer cohesive artifacts such as `login-ui-components`, `login-logic`, and `login-page-assembly`. Avoid phase-only planned flows such as `analysis`, `implementation`, `verification`, and `reporting`.

Before work, select the needed phase protocol from `references/phase-protocols.md`. Use the earliest blocker as the routing basis. If no protocol applies, stay in the default operating state without a protocol suffix.

Keep follow-up candidates, broader refactors, unrelated plugins, and new approval-sensitive work out of the active flow unless the user explicitly selects or approves them.

Task policy is flow-local. It can choose commands, edits, builds, tests, local references, target rereads, and handoffs inside the selected flow, but it cannot redefine the flow boundary, skip verification, skip reporting, skip next-flow reopening, or decide turn closure.

## Verification

After work and before reporting, choose one verification method and keep it separate from result status.

- `clean-context`: a bounded read-only verifier subagent checks the flow from an independent packet. This is not a full-history fork.
- `normal`: the main agent verifies in the same context using commands/checks, source readback, evidence checklists, log review, and logical counterexample review.
- `not-required`: no separate verification action is needed because there is no work output to verify. This is a method judgment, not a successful result.

Use `clean-context` by default for file changes, release surfaces, manifests, templates, scenario fixtures, build output, multi-file contracts, previous failed checks, user-requested verification/review/QA/commit-readiness, and destructive, irreversible, external, commit, push, PR, publish, release, or version-bump preparation or execution.

Use `normal` for lower-risk no-edit or read-only work, single state checks, narrow explanations, already-evidenced work, or cases where command/check/source readback and logical review are enough. Record why `clean-context` was not needed and what uncertainty remains.

Use `not-required` only for blocker-before-work, activation-only, next-flow selection, scope routing, or no-output routing cases. Record the omission reason, already-known evidence or no-output rationale, and residual uncertainty. Do not use `not-required` for file changes, release surfaces, multi-file contracts, failed checks, user-requested verification, or approval-sensitive actions.

Result status is separate from method. Use `pass`, `fail`, `blocked`, or `insufficient` after verification; lifecycle records may also use `not-started` or `requested` before final result. Never treat `not-required`, `blocked`, `fail`, or `insufficient` as automatic pass.

### Clean-Context Packet

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

| Verification status | Successful completion report | Expected routing |
| --- | --- | --- |
| `pass` | allowed | reporting, then next-flow reopening or self-drive continuation |
| `fail` | forbidden | return to the earliest safe work or repair phase |
| `insufficient` | forbidden | return to verification or work to gather enough evidence |
| `blocked` | forbidden | open a user-gated blocker question |

Never use a non-pass status as terminal summary, next-flow selection, or commit-readiness authority.

For report-only work with no file changes, `normal` verification may focus on source/evidence readback, logical counterexamples, user-intent fit, and missing-risk checks instead of unnecessary command execution.

For documentation-only research artifacts such as `.agents/researches/**/topic.md`, keep `clean-context` when files changed, but keep the verifier packet narrow: changed topic/session files, required headings, conclusion-to-evidence consistency, absence of implementation claims, coherent verification status, and no terminal closure without explicit stop. Do not ask the verifier to redo broad source research unless the recorded evidence is stale, incomplete, suspicious, or misses a changed path.

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

Use the question tool for narrow choices connected to the result just reported. Include the next useful flow options, blocker choices, or scope decisions. If visible choices cannot include a turn-end option, still state that the user can explicitly stop the turn and record an explicit turn-end option in the flow record.

When active self-drive is recorded, read `references/self-drive.md` before opening a default next-flow question. A valid prepared self-drive sequence may route `next-flow` to recorded loop continuation instead of a user question; blockers, scope or endpoint changes, approval-sensitive actions, and ambiguous current-flow identity still return to user-gated routing.

If the question tool is unavailable, ask a plain-text active question, state that the tool was unavailable, and record the open choices and required next action.

Next-flow terminal closure is valid only when explicit stop is source-recorded. A stale `confirmed_closure`, a source-less closure, a completed task, or an inactive self-drive sidecar is not terminal closure authority.

## Records And Templates

Use these bundled resources when session records are needed:

- `references/session-records.md` for record ownership, Continuity Guard, stale sidecar handling, and next-flow option rules.
- `references/self-drive.md` for autonomous continuation and interruption handling across a prepared flow sequence.
- `templates/plan-template.md` for `.agents/sessions/{YYYYMMDD}/000-plan.md`.
- `templates/self-drive-template.md` for `.agents/sessions/{YYYYMMDD}/000-self-drive.md` when self-drive is active.
- `templates/flow-record-template.md` for `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`.

`000-plan.md` is a date-level bounded plan and routing snapshot. It may contain a general `Planned Flow Sequence` as a date-level routing snapshot, but it does not own self-drive sequence-level state. When self-drive is active, `000-plan.md` owns only `self_drive_status` and `self_drive_record` as self-drive-specific fields.

`000-self-drive.md` owns self-drive sequence-level state: sequence objective, planned flow list, 0-based `active_flow_index`, human-readable current flow label, autonomous boundaries, approval-sensitive checkpoints, endpoint, blocker return conditions, and progress ledger.

Each `001+` flow record owns flow-local detail, including exact user request raw text when needed and its separate summary or interpretation. When self-drive is active, the flow record records only flow-local sequence position, local progress note, next handoff, and blocker return condition in an existing section. It must not repeat the full self-drive sequence contract.

Do not silently reconstruct inaccessible records. Report record access failure as a blocker and do not use missing or stale closure state as a reason to close the turn.
