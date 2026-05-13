---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# turn-gate

## Important

When this skill is active, treat it as a conversation-level first-class operating rule for the current session. Do not close the turn with a terminal summary unless the user explicitly asks to end this turn.

The required flow is `preparation -> work -> verification -> reporting -> next-flow`. After reporting, reopen the next flow with `request_user_input` when that tool is available. Keep `.agents/sessions/{YYYYMMDD}/` records current, including the active flow record and its `Continuity Guard`.

Valid ending states are limited to a source-recorded explicit user stop, an active user-gated next-flow question, or a blocker that prevents the loop from safely continuing. A plain follow-up sentence or generic closing phrase is not a substitute for next-flow reopening.

## Purpose

Use `turn-gate` as an independent loop controller. It keeps one turn alive while work moves through scoped preparation, active work, clean-context verification, continuity reporting, and explicit next-flow selection.

The default state is implicit. Do not ask the user to choose a named mode. Treat `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, and `commit-readiness-gate` as phase protocols inside the current flow, not standalone modes.

## Phase Messages

When a user-facing message starts a phase or reports progress while starting a phase, begin it with one canonical prefix:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Use the prefix only for phase-start user messages. Do not mechanically add it to flow records, output artifacts, command summaries, or every sentence in a question option.

For activation-only requests, start with `[preparation]` to set scope; use `[next-flow]` only when opening an actual next-flow choice. For status questions, use the current active phase label unless you are deliberately summarizing flow context. For session-record blockers, use the phase where the blocker is found.

## Core Loop

1. Preparation
   - Classify incoming user messages first. Only a clear request to stop this turn allows closure; all other messages are continuation input.
   - Lock intent, scope, non-goals, acceptance signal, verification expectation, and approval boundary before work.
   - Use user-gated questions when scope is empty, too broad, structurally ambiguous, able to produce multiple artifacts, or likely to change success criteria or verification.
   - If you infer scope without a question, record the work boundary and non-goals before handoff.
   - Treat requirement discovery, flow list design, meaning resolution, current-state inspection, target reread, and scope lock as preparation work.
   - User-message interpretation or planned-flow design can be an `operational-preparation` flow when it owns session plan, flow list, scope, or approval-boundary artifacts.
   - Execution-ready planned flows must be reviewable or commit-sized `change-unit` flows, not phase checklists.

2. Work
   - Execute only inside the active flow boundary.
   - Reread any target file, artifact, or state changed by a new user message before continuing.
   - Keep follow-up candidates separate from the active execution flow until the user or recorded plan selects them.
   - Use the relevant local reference only when a phase protocol is needed:
     - `references/deep-interview.md` for requirement discovery and scope-lock blockers.
     - `references/review-loop.md` for review, QA, or self-review findings.
     - `references/ralph-loop.md` for a bounded fix-verify-reassess cycle.
     - `references/autopilot.md` for broad locked-scope execution.
     - `references/commit-readiness-gate.md` for commit readiness judgment.
   - If the user explicitly requests self-drive or a prepared autonomous sequence, read `references/self-drive.md`; that overlay owns continuation decisions for the prepared sequence.

3. Verification
   - After work and before successful reporting, run clean-context verification with a bounded verifier subagent.
   - Do not fork the full conversation. Send only the target, user intent, changed files or artifacts, checks to run, pass/fail criteria, and stop conditions.
   - The verifier is read-only: no edits, no scope expansion, no destructive or external action, and no commit, push, PR, publish, release, or version bump.
   - Integrate the verifier result as `pass`, `fail`, `blocked`, or `insufficient`.
   - Treat `fail`, `blocked`, and `insufficient` as non-pass. Return to the earliest safe phase or open user-gated blocker routing; do not describe non-pass verification as completion.

4. Reporting
   - Report as continuity context, not as turn closure.
   - Include material judgment calls that affected routing, phase selection, approval, or verification.
   - Update and reread the active flow record's `Continuity Guard` before reporting and before next-flow reopening.
   - A terminal summary is allowed only when the current user message clearly asks to end this turn or the flow record has a source-recorded explicit stop matching the current message.

5. Next-flow
   - If there is no explicit stop, reopen the loop after reporting.
   - Use `request_user_input` when available and offer narrow choices tied to the just-reported result.
   - If the tool is unavailable, state that fact, list the choices in plain text, and leave the record in active question-routing state.
   - Always record a turn-end option in `Next Flow Options`, even when it is not visible among the displayed choices.

## Internal Gates

- Message intake gate classifies the incoming user message and explicit stop status. It does not execute work.
- Flow shaping gate creates or updates the active flow, completion criteria, verification expectation, and next-flow reopening condition.
- Task policy gate controls command, edit, build, test, local reference, and handoff decisions inside the current flow only.
- Verification gate decides whether verification passed, failed, was blocked, or was insufficient.
- Reporting gate packages the result as continuity context and leads into `next-flow`.

Individual task completion cannot decide flow completion or turn closure. Reporting must lead to `next-flow` unless explicit stop is source-recorded.

## Meaning And Approval

Resolve operation and target ambiguity before phase protocol selection or work execution when user wording could point to different files, surfaces, phases, specs, routing rules, or release actions. Use user-gated choices when the interpretation changes the work.

Separate meaning resolution from execution approval. A locked target or inferred intent does not authorize destructive, irreversible, external, publish, push, PR, commit, release, or version-bump actions.

Before approval-sensitive work, record the exact target, expected effect, risk, rollback or recovery path, included and excluded scope, and intended stop point. Readiness reporting is evidence only; it is not execution authority.

If approval is missing, ask a user-gated question before work. If target ambiguity and approval-sensitive action appear together, lock the target first, then ask for action approval.

## Session Records

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for active turn-gated work.

Use `templates/plan-template.md` for the plan record and `templates/flow-record-template.md` for flow records when creating new records. Keep `000-plan.md` as the date-level flow-sequence snapshot and index. Keep detailed scope, non-goals, approval boundary, evidence, verification detail, report, residual risk, `Continuity Guard`, and `Next Flow Options` in the `001+` flow record.

`000-plan.md` is a bounded date-level index and active snapshot, not the canonical detail record for each flow. It keeps recent routing context, active flow pointers, compact flow index entries, selected current or future planned flows, one-line completed flow summaries, and active date-level risks. Detailed scope, evidence, verification, flow-local risks, and completion rationale belong in the linked `001+` flow record.

Update the active flow record at the end of each phase. The `Continuity Guard` must show whether `turn-gate` is active, question-routing state, explicit stop status, terminal summary permission, required next action, last refreshed phase, closure source when present, pending or superseded question state, verification status, and continuity notes.

If records are inaccessible, do not silently reconstruct and close. Report the blocker with the phase prefix appropriate to where it was found and keep terminal summary disallowed.

## Stop Rule

Only the user can explicitly end the turn. Treat questions, corrections, status checks, review requests, priority changes, and new task requests as continuation input unless the user clearly says to stop the current turn.

Examples of explicit stop intent include "여기서 끝", "턴 종료", "이 turn은 그만", and "stop the turn". If stop intent is unclear, keep the loop active or ask a user-gated clarification.
