# Self-Drive Reference

Use this reference when a turn-gated task should keep progressing through a prepared sequence without waiting for user answers at every bounded decision.

Self-drive is not a separate user-facing skill. It is an internal `turn-gate` execution reference for bounded autonomous continuation.

## Core Contract

- Apply base `turn-gate` first. Self-drive does not replace the core phases, session records, clean-context verification, reporting, next-flow reopening, or user-gated approval routing.
- Start self-drive only after user-message-driven preparation has recorded a planned flow list, work boundaries, non-goals, verification expectations, expected risk boundaries, and user-gated checkpoints.
- Treat the planned flow list as cohesive change units, not phases. Do not create self-drive flows for analysis, work, verification, reporting, final QA, consistency checking, or commit-readiness reporting unless that item owns a distinct reviewable artifact or change unit.
- Continue through prepared flows without asking the user again only when the next decision is bounded enough for a subagent question packet.
- Stop self-drive and return to user-gated question routing for unprepared scope gaps, new approval boundaries, destructive or irreversible action, external effects, commit, push, PR, publish, or safety/tool approval.
- After the last planned change-unit flow ends, leave a commit-readiness gate report. This report is not commit execution approval and is not a new planned flow unless it owns a distinct artifact change.
- Commit, push, PR, and publish are separate user-gated handoffs.
- At every flow-end question-routing boundary, reread `turn-gate` and this reference before forming the next flow question, scope question, self-drive packet, or user-gated approval question. This refresh is deterministic, not conditional on detecting context cleanup.

## Entry Conditions

Before self-drive starts, confirm preparation has recorded:

- user intent and active task scope
- planned flow list made of cohesive reviewable change units
- work boundary and non-goals for each planned flow
- acceptance signals and verification expectations
- expected risky work and approval boundaries
- user-gated checkpoints
- current `Continuity Guard`

If this information is missing and the gap can change the result, do not invent a self-drive plan. Return to preparation or user-gated question routing.

## Self-Drive Packet

Use a subagent only for a bounded decision that can be answered without broad user judgment or privileged action.

Each self-drive packet must include:

- current `Continuity Guard`
- active flow record path or summary
- planned flow name and change-unit boundary
- included scope and non-goals
- relevant files, artifacts, or commands to inspect
- exact bounded question to answer
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
- classify the user message through `turn-gate` incoming-message handling
- reread any changed target files or artifacts before relying on old assumptions
- mark pending answers as superseded when the user message changes scope, priority, approval state, target, or next-flow candidates
- use a pending answer only after revalidating it against the updated `Continuity Guard`, active flow record, and user message

Never use a stale subagent answer to justify terminal closure, approval-sensitive action, or continuation that conflicts with the user's latest input.

## Flow-End Refresh

At every boundary after a flow result is reported and before the next routing question is formed, deterministically refresh the runtime contracts:

1. Reread the active `turn-gate` skill.
2. Reread this `references/self-drive.md` file.
3. Recheck the active flow record's `Continuity Guard`.
4. Recheck active flow boundary, pending next-flow candidates, pending or superseded questions, and approval boundaries.
5. Only then form the next flow question, scope question, self-drive packet, or user-gated approval question.

This is not a context-cleanup detector. Its purpose is to reload the base loop contract and self-drive contract before the next routing decision.

The refresh does not grant new permissions. Approval, destructive or irreversible action, external effects, safety/tool decisions, commit, push, PR, and publish still require user-gated handoff.

## Mode Boundary

- Good fit: prepared planned flow sequences where the next decisions are bounded, reversible, and answerable by subagent packets.
- Not a fit: missing scope that can change the result, new work outside the prepared flow list, destructive or irreversible decisions, external effects, commit/push/PR/publish, tool/runtime approvals, safety-sensitive decisions, stale/conflicting subagent answers, or inaccessible session records.

## Review Questions

- Is base `turn-gate` active and still governing the turn?
- Is the planned flow list made of cohesive change units, not phases?
- Does the active flow have recorded work boundary, non-goals, verification expectation, and approval boundaries?
- Is the next decision bounded enough for a self-drive packet?
- Does the packet include `Continuity Guard`, allowed actions, forbidden actions, stale-answer handling, and required output shape?
- Has user intervention been checked before using any pending answer?
- Do approval-sensitive decisions return to user-gated question routing?
- Has flow-end deterministic refresh reread both `turn-gate` and `references/self-drive.md`?
- Does the final change-unit flow end with commit-readiness gate reporting, not commit execution approval?
