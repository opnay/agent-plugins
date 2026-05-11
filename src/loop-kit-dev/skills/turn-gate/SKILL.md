---
name: turn-gate
description: Keep a Codex turn open across preparation, work, verification, reporting, and explicit next-flow reopening until the user explicitly stops the turn.
---

# turn-gate

## Important

When this skill is active, `turn-gate` is a conversation-level operating rule for the current turn.

- Do not close the turn with a terminal summary unless the user explicitly asks to stop this turn.
- Treat every non-stop user message as continuation input for the same turn.
- End each completed flow in one of these states: next-flow reopening, user-gated blocker, or source-recorded explicit turn stop.
- After reporting, reopen the next flow with `request_user_input` when structured choices are possible and the tool is available.
- Maintain session records throughout the turn, including the plan, flow records, `Continuity Guard`, verification status, and `Next Flow Options`.
- Read only installed runtime materials during execution: local `references/` files for selected internal modes and local `templates/` files for session records.

## Purpose

Use `turn-gate` when work must continue inside one turn until the user explicitly stops it. The skill keeps the loop structured as:

1. preparation
2. work
3. verification
4. reporting
5. continuation through next-flow reopening

Reporting is continuity context for the next flow. It is not permission to stop.

## Phase Start Messages

When you tell the user that a phase is starting, or when a progress update begins with the start of a phase, the user-facing message must start with the canonical phase label.

Canonical labels:

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[continuation]`

This prefix is an operational marker for phase-start messages only. Do not mechanically add it to flow records, artifact bodies, command output summaries, or every sentence in question choices. If you are giving a general explanation without starting a phase, do not force a prefix.

## Core Loop

### 1. Preparation

Before work begins, shape the active flow and lock enough context to proceed.

- Classify the incoming user message first: explicit turn stop or continuation input.
- If it is not an explicit stop, keep the turn open.
- Identify intent, scope, non-goals, acceptance signal, verification expectation, and approval boundary.
- Distinguish user-message-based preparation from preparation for an already selected flow.
- Use deep-interview, flow-list design, meaning resolution, and current-state inspection as preparation techniques.
- If the scope is empty, too broad, able to produce multiple materially different results, or able to change success criteria or verification paths, ask before work.
- If you infer the scope without a question, record the work boundary and non-goals in the flow record.
- If operation or target meaning is structurally ambiguous, lock meaning before execution.
- If user-message interpretation and planned flow design are themselves the work, treat them as an `operational-preparation` flow.
- Planned execution flows produced by preparation should be reviewable or commit-sized `change-unit` flows.
- Keep follow-up candidates separate from the active execution flow.
- For a prepared planned flow sequence that can continue without more user input, read `references/self-drive.md` and apply its limits.
- Once meaningful work starts, use the plan tool to track phase state.

### 2. Work

Before entering work, choose one internal mode for the current-phase work and read its local runtime reference.

Internal mode choices:

- `deep-interview`: requirements, unclear intent, missing scope, or unresolved approval lines block work.
- `review-loop`: review feedback, QA findings, or self-review findings are the main issue.
- `ralph-loop`: one small fix-verify-reassess cycle is the right unit.
- `autopilot`: broad end-to-end delivery is the current unit.
- `self-drive`: a prepared planned flow sequence can continue by bounded subagent decision.
- `commit-readiness-gate`: the change unit is nearly complete and readiness judgment is the main task.

If modes overlap, prefer the earliest blocker in this order: `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `self-drive`, `commit-readiness-gate`.

Read the matching file under `references/` before applying that mode. External actions such as commit execution, push, PR creation, publish, release, or version bump are not internal modes; they are user-gated handoffs.

### 3. Verification

After work and before reporting, perform clean-context verification.

- Use a read-only bounded verifier subagent for clean-context verification.
- Do not fork the full conversation history as the verifier context.
- Provide only the bounded packet: verifier identity or request id, target, user intent, changed files or artifacts, checks or commands to run, pass/fail criteria, no-edit instruction, and stop condition.
- The verifier must not edit files, expand scope, perform destructive or external actions, or commit, push, open PRs, publish, release, or bump versions.
- Integrate the verifier result as `pass`, `fail`, `blocked`, or `insufficient`.
- Treat `fail`, `blocked`, and `insufficient` as non-pass.
- For `fail` or `insufficient`, return to the earliest safe phase before result reporting.
- For `blocked`, or if verifier tooling is unavailable, open user-gated question routing. Do not report blocked verification as successful completion.

### 4. Reporting

Reporting summarizes the active flow for continuity, not closure.

- State what was prepared, changed or completed, and verified.
- Include residual uncertainty and blockers if they exist.
- Update and read the active flow record's `Continuity Guard` before reporting.
- Only allow a terminal summary when the current user message or a source-recorded confirmed closure explicitly stops the turn.
- Ignore stale closure state, source-less `confirmed closure`, or old `terminal summary allowed: yes` records.
- If stale closure state is found, reset the guard to continuation state and record why.

### 5. Continuation

After reporting, reopen the next flow unless there is a source-recorded explicit stop.

- Use `request_user_input` for structured next-flow choices when available.
- Choices should be narrow and directly connected to the just-reported result.
- If `request_user_input` is unavailable, use plain text fallback, state that the tool is unavailable, list the open choices, and keep the state as active question routing.
- A generic closing phrase is not next-flow reopening.
- If visible choices cannot include a turn-end option, still tell the user they can explicitly stop the turn.
- Always record a turn-end option in the flow record's `Next Flow Options`.

## Internal Gates

Use internal gates to decide transitions. Do not expose them as separate user-facing skills.

- Message intake gate: classify explicit stop vs continuation input. It does not execute work.
- Flow shaping gate: create or update the active flow, follow-up candidates, and completion criteria.
- Task policy gate: decide execution policy inside the selected flow only.
- Verification gate: integrate work results as pass, fail, blocked, or insufficient.
- Reporting gate: convert flow results into continuity context.
- Continuation gate: confirm no explicit stop exists and reopen the next flow.

No gate may infer terminal closure without an explicit user stop. Individual task completion cannot decide flow completion or turn closure by itself.

## Approval Boundary

Meaning resolution and action approval are different decisions.

- Lock operation and target meaning before asking for risky action approval when both are uncertain.
- Destructive, irreversible, external, commit, push, PR, publish, release, and version-bump actions require user-gated approval before execution.
- Approval-sensitive prompts must include exact target, expected effect, risk, rollback or recovery possibility, and included/excluded scope.
- Subagent judgment, inferred intent, and nearby prior wording cannot substitute for approval.
- A readiness request is not commit approval.
- Commit, push, PR, publish, release, and version-bump execution belong to separately approved handoff workflows.

## Session Records

Maintain `.agents/sessions/{YYYYMMDD}/` records for active turn-gated work.

- Keep `000-plan.md` as the date-scoped flow sequence, user request history, current plan, and completed-flow summary.
- Keep `001+` flow records as detailed reports for individual flows.
- Use `templates/plan-template.md` and `templates/flow-record-template.md` when creating records.
- Flow record filenames use `{count-pad3}-{eng-lower-slug}.md`, for example `001-update-parser.md`.
- Record flow type as `operational-preparation` or `change-unit`.
- Each flow normally carries preparation, work, verification, and reporting.
- Update flow records incrementally at phase boundaries instead of waiting for the flow to finish.
- Keep completed work in the plan; summarize it and retain the flow reference.

The `Continuity Guard` must include turn-gate active state, question-routing mode, user explicit stop status, terminal summary permission, required next action, last refreshed phase, and verification status when relevant.

Before reporting and next-flow reopening, refresh the `Continuity Guard`. Missing records may be reconstructed, then written back. Inaccessible records are blockers, not permission to close the turn.

## Runtime References

During execution, only read installed runtime files that exist with this skill:

- `references/deep-interview.md`
- `references/review-loop.md`
- `references/ralph-loop.md`
- `references/autopilot.md`
- `references/self-drive.md`
- `references/commit-readiness-gate.md`
- `templates/plan-template.md`
- `templates/flow-record-template.md`

Do not instruct runtime users to read development-only specification files. If detailed behavior is needed at runtime, it must be present in this skill body or in the installed `references/` or `templates/` files.
