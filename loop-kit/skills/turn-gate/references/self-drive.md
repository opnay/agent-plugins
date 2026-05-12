# Self-Drive Reference

Use this reference when a turn-gated task should keep progressing through a prepared sequence without waiting for user answers at every bounded decision.

Self-drive is not a separate user-facing skill. It is a runtime overlay contract for a prepared planned flow sequence. Once this reference is selected, it takes priority over base `turn-gate` routing for that prepared sequence and governs continuation until it hands control back to user-gated routing, reporting, or explicit stop handling.

## Core Contract

- Start self-drive only after user-message-driven preparation has recorded a planned flow list, work boundaries, non-goals, verification expectations, expected risk boundaries, and user-gated checkpoints.
- Treat the planned flow list as cohesive change units, not phases. Do not create self-drive flows for analysis, work, verification, reporting, final QA, consistency checking, or commit-readiness reporting unless that item owns a distinct reviewable artifact or change unit.
- Continue through prepared flows without asking the user again when the next decision is covered by this self-drive contract.
- Continue through approval-sensitive work, including commit, push, PR, publish, release, or version bump, only when the initial preparation explicitly recorded the exact action, target, expected effect, risk, recovery path, included/excluded scope, and stop boundary.
- Stop self-drive and return to user-gated question routing for unprepared scope gaps, new approval boundaries, approval-sensitive work not covered by the initial agreement, missing stop boundaries, safety/tool approval, or inaccessible session records.
- After the last planned change-unit flow ends, leave a commit-readiness gate report or another recorded handoff. Treat execution authority and new planned-flow creation as separate recorded decisions.
- Base `turn-gate` still supplies explicit stop handling, safety/tool approval limits, session records, clean-context verification, reporting shape, next-flow reopening mechanics, and approval-boundary tracking.

## Entry Conditions

Before self-drive starts, confirm preparation has recorded:

- user intent and active task scope
- planned flow list made of cohesive reviewable change units
- work boundary and non-goals for each planned flow
- acceptance signals and verification expectations
- expected risky work and approval boundaries
- explicit stop boundary for any approval-sensitive execution that self-drive may perform
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
- approval boundaries and stop boundaries that must not be crossed
- required output shape
- stale-answer rule: user intervention after packet dispatch supersedes the answer unless the main agent revalidates it

Do not give a self-drive packet authority to edit outside its assigned scope, create new approval, bypass safety or tool approval, execute approval-sensitive action outside the recorded initial agreement, or end the turn.

## Subagent Answer Contract

A useful answer must provide the next action, not a terminal summary.

Require the answer to include:

- direct answer to the bounded question
- evidence or files inspected
- assumptions and uncertainty
- whether the answer still fits the current `Continuity Guard`
- recommended next action for the main flow
- any condition that requires returning to user-gated question routing

Treat missing evidence, broad speculation, or an answer that crosses the packet boundary as insufficient. Integrate the answer yourself; do not let it replace self-drive verification, approval, or reporting requirements.

## User Intervention And Refresh

The latest user message is authoritative. If the user intervenes while a self-drive packet is pending, pause use of pending answers, classify the user message through message intake, reread changed targets, and mark answers superseded when the message changes scope, priority, approval state, target, or next-flow candidates.

At every boundary after a flow result is reported and before the next routing decision is formed, refresh this reference, the base `turn-gate` mechanics needed for records and safety, and the active flow record's `Continuity Guard`. This refresh does not grant new permissions.

## Review Questions

- Is self-drive still the active overlay for this prepared sequence?
- Is the planned flow list made of cohesive change units, not phases?
- Does the active flow have recorded work boundary, non-goals, verification expectation, and approval boundaries?
- Is the next decision covered by the self-drive packet boundary?
- Has user intervention been checked before using any pending answer?
- Do approval-sensitive decisions either fit the recorded initial agreement and stop boundary or return to user-gated question routing?
