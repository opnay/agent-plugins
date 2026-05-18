---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# turn-gate

## Important

When this skill is active, it is a conversation-level first-class operating rule for the current session. Keep the turn open across preparation, work, verification, reporting, and next-flow routing until the user gives an explicit turn stop.

Do not end with a terminal summary after reporting unless the current user message, or a source-recorded closure in the active flow record that matches the current user message, explicitly stops this turn.

After reporting, reopen the next flow with `request_user_input` whenever available unless there is a source-recorded explicit stop or a recorded self-drive continuation whose next planned flow is still valid and identifiable. If the question tool is unavailable, ask a visible next-flow question and record the options, including an explicit turn-end option, in the active flow record.

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and the active `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` flow record unless the user explicitly forbids all file writes, file creation, or session records. Refresh the active flow record at each phase boundary, and refresh the `Continuity Guard` before result reporting and before next-flow reopening.

Trusted bundled Stop and SessionStart hooks can help as optional backstops, but they do not replace this operating rule. SessionStart context is advisory startup context only. Stop hook blocking is a signal to refresh records and return to the required next action, not permission to skip reporting, record refresh, or next-flow reopening.

## Purpose

Use `turn-gate` to keep one Codex turn alive while a task moves through:

1. preparation
2. work
3. verification
4. reporting
5. next-flow

Treat `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, and `commit-readiness-gate` as phase protocols, not standalone modes. Choose the protocol that fits the active phase and current task. For local protocol details, read `references/phase-protocols.md`. For self-drive sequences explicitly requested by the user, read `references/self-drive.md`.

## Phase Messages

When a user-facing message starts a phase or gives phase-start progress, begin it with `[<phase-name>(/<phase-protocol>)]`.

Canonical phase labels are `preparation`, `work`, `verification`, `reporting`, and `next-flow`.

The `(/<phase-protocol>)` segment is optional notation. In actual output, use a slash suffix only when a protocol is active, and do not print literal parentheses.

Examples:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`
- `[preparation/deep-interview]`
- `[work/ralph-loop]`
- `[verification/review-loop]`
- `[reporting/commit-readiness-gate]`

Apply this prefix to phase-start user messages and status/progress messages. Do not mechanically copy it into flow records, generated artifacts, command summaries, or question option labels.

If activation has no concrete task, start with `[preparation]` to set scope, then use `[next-flow]` only when opening the next-flow choice. For mid-work status, use the current active phase unless intentionally summarizing flow context. For self-drive continuation, use prefixes for visible status, verification, reporting, and automatic handoff messages, but do not copy prefixes into `000-self-drive.md`, flow records, generated artifacts, or question option labels. For record access blockers, use the phase where the blocker is discovered.

## Operating Cycle

### Preparation

Before work, align intent, scope, non-goals, acceptance signal, verification expectation, and approval boundary.

Use preparation for deep interview, flow list design, meaning resolution, current-state inspection, target reread, and scope lock. If the request asks for interpretation, planning, or target selection before execution, the first flow can be an `operational-preparation` flow whose output is a set of follow-up `change-unit` flow candidates.

If the target, operation, output shape, success criteria, or verification path can materially change the work, ask before acting. Ambiguous operation triggers include phrases like "정리해줘", "반영해줘", "이 기준으로 고쳐줘", or references to multiple possible files, skills, specs, plugins, branches, or records.

If you infer scope without asking, record the chosen scope, non-goals, and why a question was not needed.

For approval-sensitive work, record the exact target, expected effect, risk, rollback or recovery path, included and excluded scope, and stopping point. Readiness reporting is not execution authority. Commit, push, PR, publish, release, version bump, destructive action, deletion, and external side effects require explicit user approval.

### Work

Do only the active flow's work. A single task completion does not complete the flow and does not authorize turn closure.

Keep active execution separate from follow-up candidates. If a new issue appears, decide whether it belongs inside the current flow boundary or should be recorded as a later next-flow option.

Use the simplest complete implementation first, then improve within the same turn when the flow requires it.

### Verification

Verification method is one of `clean-context`, `normal`, or `not-required`. Do not mix method with result status.

Use `clean-context` by default for file changes, release surfaces, multi-file contracts, prior failure history, user-requested verification, and approval-sensitive actions. Clean-context verification is a bounded verification packet, not a full-history fork.

A verifier packet should include the target, expected behavior, changed surface, relevant files or commands, constraints, and explicit prohibitions. Verifiers must not edit unless asked. They must not expand scope, perform destructive or external work, or commit, push, open PRs, publish, release, or bump versions.

`not-required` is a method judgment, not a pass status. Record the reason and residual uncertainty.

Do not treat failed, blocked, or insufficient verification as passed. Route non-pass results back to the appropriate phase or report the blocker with concrete residual risk.

### Reporting

Reporting is continuity context for the next flow, not terminal closure.

Report what changed, what was verified, what remains uncertain, and what the next required action is. Refresh the active flow record and its `Continuity Guard` before reporting.

Only a source-recorded explicit stop can allow terminal summary. "여기서 끝", "턴 종료", "이 turn은 그만", or "stop the turn" can count when they clearly refer to ending the current turn. If the stop intent is unclear, treat it as continuation input or ask.

### Next-flow

After reporting, move to `[next-flow]` unless there is a confirmed explicit stop. Refresh `Continuity Guard` again before opening the next-flow choice.

Use `request_user_input` with 2-3 concise choices when possible. The recorded `Next Flow Options` must include a turn-end option even if the visible choices do not show one.

If `request_user_input` is unavailable, ask a direct next-flow question in the response and record the fallback.

If active self-drive has a prepared sequence that is still valid and the next planned flow is identifiable, `[next-flow]` may continue from records instead of asking. If continuation identity, endpoint, approval boundary, blocker state, or current-flow identity is unclear, return to user-gated routing.

## Session Records

Use runtime templates when creating records:

- `templates/plan-template.md`
- `templates/flow-record-template.md`
- `templates/self-drive-template.md`

Use `references/session-records.md` for record details and recovery rules.

The plan record is a compact active recovery snapshot and flow index. The flow record owns flow-local scope, non-goals, approval boundary, execution log, verification evidence, report, residual risk, `Continuity Guard`, and `Next Flow Options`.

Do not turn phase steps into separate flow records. A flow usually contains preparation, work, verification, reporting, and next-flow phases.

If a user forbids target/source edits but does not forbid session records, you may still maintain session records as operational continuity artifacts. If the user forbids all writes, all file creation, or session records specifically, do not write records; route to clarification or blocker handling.

## Optional Hooks

### Stop Hook

When a trusted bundled Stop hook is enabled, treat it as an optional runtime backstop for preventing source-unrecorded terminal closure. It does not replace the `turn-gate` conversation-level rule.

If the Stop hook blocks, the main agent must refresh `000-plan.md` and the active flow record, then return to the hook reason's `required_next_action`. The main agent still owns record correction, reporting, and next-flow reopening.

The Stop hook does not approve destructive actions, external actions, commits, pushes, PRs, publishing, releases, version bumps, global Codex config changes, plugin hook activation, or record writes by the hook itself.

Detailed hook block and quiet-exit behavior belongs to the bundled plugin hook implementation and the plugin-level Stop hook spec, not this runtime skill body.

### SessionStart Hook

When a trusted bundled SessionStart hook is enabled, treat its output as advisory startup context from `.agents/sessions/`.

Context with `required_next_action` is not automatic continuation authority. Recheck the current user message, `000-plan.md`, and the active flow record before acting.

SessionStart context does not approve commit, push, PR, release, version bump, deletion, external action, or terminal closure. If the previous flow is closed, treat it as historical context only. If no plan or flow exists, begin normal preparation.

## Common Misclassifications

- A completed command, edit, or test is not flow completion by itself.
- A final-looking report is not terminal closure without source-recorded explicit stop.
- Commit readiness is a judgment protocol; it is not approval to commit.
- A future endpoint like "끝나면 알려줘" is not an immediate turn stop.
- File changes usually require `clean-context` verification unless risk is clearly low and the reason is recorded.
- A status question during self-drive is still answered within the current phase and does not reset the flow.
- `not-required` is a verification method, not a successful verification result.
- A stale closure field, stale routing mismatch, or stale self-drive sidecar is not authority to close, skip next-flow, or continue autonomously.
