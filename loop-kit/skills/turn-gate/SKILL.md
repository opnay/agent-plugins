---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# Turn Gate

## Important

When this skill is active, treat it as a conversation-level first-class operating rule for the current session. Do not close the turn with a terminal summary after reporting results unless the user explicitly asks to stop this turn.

Required ending states are:

- active work continues inside the current flow;
- active question-routing is open for clarification, blocker handling, or next-flow selection;
- a blocker is reported and the required user decision is recorded;
- terminal close is allowed only by a source-recorded explicit stop message.

After every result report without explicit stop, reopen the next flow with `request_user_input` when structured choices are possible. If that tool is unavailable, use a plain-text fallback that clearly remains active question-routing.

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` records for active turn-gated work. Update the flow record's Continuity Guard before reporting and before next-flow reopening.

## Purpose

Use `turn-gate` to preserve turn continuity while applying the phase protocol needed for the current work. The default operating state is implicit: do not expose or record a standalone mode unless a separate explicit overlay, such as self-drive, has been requested and prepared.

The core cycle is:

1. `preparation`
2. `work`
3. `verification`
4. `reporting`
5. `next-flow`

Activation-only requests, such as "Use turn-gate", should not become completion summaries. Open scope setup or next-flow selection instead.

## Phase Start Messages

When you tell the user a phase is starting, or provide a progress update at the start of a phase, begin that user-facing message with one of these prefixes:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Apply the prefix to phase-start conversation messages only. Do not mechanically add it to flow records, output artifacts, command summaries, or every sentence in question options.

Priority when labels could overlap:

- activation-only with no concrete task: start with `[preparation]` for scope setup; use `[next-flow]` only when opening actual next-flow choices.
- mid-work status: use the current active phase, usually `[work]`.
- record access blocker: use the phase where the blocker is discovered, usually `[reporting]` or `[next-flow]`.
- report-only evaluation: if there is no explicit stop, reporting still leads to `[next-flow]`.

## Preparation

Before work, lock enough context to make the current flow executable:

- intent;
- scope and non-goals;
- acceptance signal;
- verification expectation;
- approval boundary.

Distinguish initial preparation from pre-work preparation for an already selected flow. If the user request needs interpretation before execution, record that as an `operational-preparation` flow. Its output may be a planned list of `change-unit` flows or follow-up candidates; do not treat candidates as active execution flows until selected.

Use meaning resolution before protocol selection or execution when operation or target wording could change the file scope, deletion behavior, routing rule, approval boundary, or commit scope. Ask a narrow user-gated question when needed.

If scope is empty, too broad, can produce different outputs, or can change the verification path, lock it with `request_user_input` before work. If you infer scope without a question, record the work boundary and non-goals in the flow record.

For approval-sensitive actions, record exact target, expected effect, risk, rollback or recovery possibility, included and excluded scope, and stopping point. Readiness reporting is not execution authority. Commit, push, PR, publish, release, and version bump require explicit approval or a handoff that records the missing approval.

When meaningful work begins, keep the visible task plan updated with the planning tool.

## Work

Perform only work that fits the active flow boundary. A single task completion does not decide flow completion or turn closure.

Use the implicit default state for normal turn-gated work. Select a phase protocol only as a way to perform the current phase:

- `deep-interview`: requirement discovery or scope lock blocks work.
- `review-loop`: a review, QA, or self-review finding is the current blocker.
- `ralph-loop`: one narrow fix-verify-reassess cycle is appropriate.
- `autopilot`: locked scope needs end-to-end execution.
- `commit-readiness-gate`: the change unit needs readiness judgment, not commit execution.

If multiple protocols seem possible, handle the earliest blocker first, usually in this order: `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `commit-readiness-gate`.

For protocol details, read [phase-protocols.md](references/phase-protocols.md). If the user explicitly asks for self-drive over a prepared sequence, read [self-drive.md](references/self-drive.md).

## Verification

After work and before result reporting, run a clean-context verification step. Clean-context verification is not a full-history fork: send a bounded verification packet with only the needed intent, files or artifacts, changed surfaces, checks, pass/fail criteria, and stop conditions.

The verifier must be read-only:

- no edits;
- no scope expansion;
- no destructive or external work;
- no commit, push, PR, publish, release, or version bump.

The main agent builds the packet and integrates the result; it does not replace clean-context verification by rechecking in the same context. Treat verification as one of `pass`, `fail`, `blocked`, or `insufficient`.

Do not report `fail`, `blocked`, or `insufficient` as successful completion. For `fail` or `insufficient`, return to the earliest safe phase before result reporting. For `blocked`, open user-gated blocker routing.

## Reporting

Reporting is continuity context, not a terminal close. Include what was prepared, changed or inspected, how it was verified, remaining uncertainty, blockers, and material judgment calls that affected routing.

Before reporting:

1. read or reconstruct the active flow record only if the record is missing;
2. do not silently bypass record access failures;
3. update the Continuity Guard;
4. confirm whether the current incoming user message contains an explicit stop.

Only a source-recorded explicit stop, such as "stop the turn" or "turn end", permits a terminal summary. Stale closure records or source-less `terminal_summary_allowed: yes` values are invalid and must be repaired in the flow record.

For session record structure, read [session-records.md](references/session-records.md). Use the templates in [plan-template.md](templates/plan-template.md) and [flow-record-template.md](templates/flow-record-template.md).

## Next Flow

After reporting, if there is no explicit stop, enter `next-flow`.

Use `request_user_input` when available and when choices can be structured. Offer narrow choices connected to the result just reported. If the visible choices cannot include a turn-end option, still tell the user they can explicitly stop the turn, and record a turn-end option in the flow record's `Next Flow Options`.

Plain-text follow-up phrases are not a substitute for next-flow reopening. If `request_user_input` is unavailable, state that, list the open choices, and leave the required next action in the session record.

Do not infer approval for commit, push, PR, publish, release, version bump, destructive action, or external action from a next-flow choice unless that exact action boundary was explicitly approved.
