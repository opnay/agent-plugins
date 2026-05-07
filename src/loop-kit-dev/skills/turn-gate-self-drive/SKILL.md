---
name: turn-gate-self-drive
description: Self-drive overlay for `turn-gate`. Use when a turn-gated task should keep progressing without waiting for user answers by routing bounded decisions to subagents while preserving the `turn-gate` continuity contract.
---

## Important

Use this skill only as an overlay on the same plugin's `turn-gate` skill.

- Apply `turn-gate` first. This skill does not replace the base loop, core phases, session records, verification rules, or user-gated approval routing.
- Use self-drive only after user-message-driven preparation has produced a planned flow list, work boundaries, non-goals, verification expectations, expected risk boundaries, and user-gated checkpoints.
- The planned flow list is a list of cohesive change units, not a phase list. Do not create self-drive flows for analysis, work, verification, reporting, final QA, consistency checking, or commit-readiness reporting unless that item creates or modifies a distinct reviewable artifact or change unit.
- Continue through prepared flows without asking the user again only when the next decision is bounded enough for a subagent question packet.
- Stop self-drive and return to `turn-gate` user-gated question routing for unprepared scope gaps, new approval boundaries, destructive or irreversible action, external effects, commit, push, PR, publish, or safety/tool approval.
- After the last planned change-unit flow ends, leave a commit-readiness gate report. This report is not commit execution approval and is not a new planned flow unless it owns a distinct artifact change.
- Commit, push, PR, and publish are separate user-gated handoffs.
- At every flow-end question-routing boundary, reread `turn-gate` and `turn-gate-self-drive` before forming the next flow question, scope question, self-drive packet, or user-gated approval question. This refresh is deterministic, not conditional on detecting context cleanup.

## Purpose

Use `turn-gate-self-drive` when a turn-gated task should keep progressing through a prepared sequence without waiting for user answers at every bounded decision.

This skill owns self-drive question routing: how to package bounded questions for subagents, how to read their answers, how to recover after context gaps, how to treat stale answers, and how to return to user-gated routing at approval boundaries. The base `turn-gate` skill continues to own turn continuity, phase order, flow records, `Continuity Guard`, clean-context verification, result reporting, and next-flow reopening.

## Entry Conditions

Before self-drive starts, confirm the base `turn-gate` preparation has already recorded:

- user intent and active task scope
- planned flow list made of cohesive reviewable change units
- work boundary and non-goals for each planned flow
- acceptance signals and verification expectations
- expected risky work and approval boundaries
- user-gated checkpoints
- current `Continuity Guard`

If this information is missing and the gap can change the result, do not invent a self-drive plan. Return to `turn-gate` preparation or user-gated question routing.

## Planned Flow Boundaries

A self-driven planned flow is a cohesive change unit that can be understood, reviewed, verified, and, if needed, committed together. It does not need to be direct user-visible value. UI component scaffolding, logic layer work, integration assembly, fixture changes, documentation changes, and validator output changes can all be planned flows when each owns a distinct reviewable change unit.

Do not promote core phases or status-only checks into planned flows. Pure final QA, consistency checking, verification-result reporting, and commit-readiness reporting belong in the last change-unit flow's verification/reporting or in a user-gated handoff, unless they create or modify a distinct artifact such as a regression fixture, snapshot baseline, documentation page, operator report output, or validator diagnostic output.

When the last planned change-unit flow is complete, provide a commit-readiness gate report. State readiness, blockers, verification status, residual risk, and any unrelated changes that must stay out of commit scope. Do not ask for or imply commit execution approval inside that report. Commit, push, PR, and publish require a separate user-gated handoff through the base `turn-gate` approval boundary.

## Self-Drive Packet

Use a subagent only for a bounded decision that can be answered without broad user judgment or privileged action.

Each self-drive packet must include:

- the current `Continuity Guard`
- active flow record path or summary
- planned flow name and change-unit boundary
- included scope and non-goals
- relevant files, artifacts, or commands to inspect
- the exact bounded question to answer
- allowed actions and forbidden actions
- verification expectation or pass/fail criteria when relevant
- approval boundaries that must not be crossed
- required output shape
- stale-answer rule: user intervention after packet dispatch supersedes the answer unless the main agent revalidates it

Do not give a self-drive packet authority to edit outside its assigned scope, approve destructive or external action, bypass safety or tool approval, commit, push, create a PR, publish, or end the turn.

## Subagent Answer Contract

A useful answer must provide the next action, not a terminal summary.

Require the answer to include:

- direct answer to the bounded question
- evidence or files inspected
- assumptions and uncertainty
- whether the answer still fits the current `Continuity Guard`
- recommended next action for the main flow
- any condition that requires returning to user-gated question routing

Treat missing evidence, broad speculation, or an answer that crosses the packet boundary as insufficient. Integrate the answer yourself; do not let it replace base `turn-gate` verification, approval, or reporting.

## Stale Answers And User Intervention

The latest user message is authoritative.

If the user intervenes while a self-drive packet is pending:

- pause use of pending subagent answers
- classify the user message through base `turn-gate` incoming-message handling
- reread any changed target files or artifacts before relying on old assumptions
- mark pending answers as superseded when the user message changes scope, priority, approval state, target, or next-flow candidates
- use a pending answer only after revalidating it against the updated `Continuity Guard`, active flow record, and user message

Never use a stale subagent answer to justify terminal closure, approval-sensitive action, or continuation that conflicts with the user's latest input.

## Flow-End Refresh

At every boundary after a flow result is reported and before the next routing question is formed, deterministically refresh the runtime contracts:

1. Reread the same plugin's `turn-gate` skill.
2. Reread this `turn-gate-self-drive` skill.
3. Recheck the active flow record's `Continuity Guard`.
4. Recheck active flow boundary, pending next-flow candidates, pending or superseded questions, and approval boundaries.
5. Only then form the next flow question, scope question, self-drive packet, or user-gated approval question.

This is not a context-cleanup detector. You cannot reliably know whether compaction or cleanup happened, so the refresh is unconditional at every flow-end question-routing boundary. Its purpose is to reload the base loop contract and this overlay contract before the next routing decision.

The refresh does not grant new permissions. Approval, destructive or irreversible action, external effects, safety/tool decisions, commit, push, PR, and publish still require the base `turn-gate` user-gated handoff.

## Approval Boundary Return

When a decision is not bounded enough for self-drive, return to the base `turn-gate` user-gated question-routing path. Do not end the turn just because self-drive cannot continue.

Return immediately for:

- missing scope that can change the result
- new work outside the prepared planned flow list
- destructive or irreversible action
- external effects or network-facing publication
- commit, push, PR, or publish
- tool/runtime approval requirements
- safety-sensitive decisions
- stale or conflicting subagent answers
- inaccessible session records or invalid `Continuity Guard`

The handoff question should state the current flow, the blocked decision, included scope, excluded scope, risk, and the next safe options.

## Recovery From Context Gaps

If context is incomplete, recover from durable records instead of guessing.

Read the active `.agents/sessions/{YYYYMMDD}/000-plan.md`, active `001+` flow record, and the relevant target files. Reconstruct only the minimum state needed to continue: active flow, flow boundary, non-goals, verification expectation, pending next-flow candidates, approval boundaries, and `Continuity Guard`.

If records are missing, stale, inaccessible, or contradictory, report the blocker through base `turn-gate` question routing. Do not silently continue self-drive from memory.

## Review Checklist

Before continuing self-drive, verify:

- base `turn-gate` is active and still governs the turn
- the planned flow list is made of cohesive change units, not phases
- pure QA, consistency checking, verification-result reporting, or commit-readiness reporting has not become a planned flow without owning a distinct artifact change
- the active flow has recorded work boundary, non-goals, verification expectation, and approval boundaries
- the next decision is bounded enough for a self-drive packet
- the packet includes `Continuity Guard`, allowed actions, forbidden actions, stale-answer handling, and required output shape
- user intervention has been checked before using any pending answer
- approval-sensitive decisions return to `turn-gate` user-gated question routing
- flow-end deterministic refresh has reread both `turn-gate` and `turn-gate-self-drive`
- the final change-unit flow ends with commit-readiness gate reporting, not commit execution approval
