---
name: turn-gate-self-drive
description: Self-drive overlay for `turn-gate`. Use when a turn-gated task should keep progressing without waiting for user answers by routing blocked decisions to subagents while preserving the `turn-gate` continuity contract.
---

# Turn Gate Self Drive

## Important

Use this skill only as an overlay on the same-plugin `turn-gate` skill. First apply `turn-gate` as the base loop contract: keep the same turn alive, preserve the `Continuity Guard`, maintain flow records, run verification before reporting, and reopen the next flow unless the user explicitly ends the turn.

Initial preparation is owned by `turn-gate`, not by this skill. Before self-drive begins, `turn-gate` must have collected the planned flow list, intent, scope, non-goals, acceptance signal, verification expectations, expected risky actions, approval boundaries, and user-gated checkpoints. This skill executes that prepared plan and routes bounded decisions; it does not invent missing scope or replace user approval.

## Purpose

Use `turn-gate-self-drive` when a `turn-gate` task should continue through a prepared planned flow list without repeatedly asking the user. The overlay sends bounded decision packets to subagents, consumes their answers as decision input, and returns to `turn-gate` user-gated question routing when the prepared agreement no longer covers the next action.

This skill owns self-drive packet structure, answer handling, stale-answer handling, context-gap recovery, planned-flow sequence execution, deterministic flow-end rereads, and final commit-readiness reporting. It does not own `turn-gate` phase continuity, current-phase mode selection, direct user-gated next-flow operation, runtime approval policy, or commit/push/PR/publish execution.

## Use When

- `turn-gate` is the governing loop.
- The user requested autonomous continuation, self-driving progress, or no extra user questions between prepared flows.
- `turn-gate` preparation already recorded the planned flow list and per-flow boundaries.
- The next decision is bounded by recorded evidence, constraints, non-goals, verification expectations, and approval boundaries.

Do not use self-drive for a flow whose scope, acceptance signal, verification expectation, risky-action treatment, or approval boundary is missing and would affect the result.

## Planned Flow Execution

Execute prepared flows in order. Before each flow starts:

1. Re-read the active flow record and the prepared planned flow list.
2. Confirm the flow's work boundary, non-goals, acceptance signal, verification expectation, expected risky actions, approval boundary, and user-gated checkpoints.
3. Check whether the next action is inside the recorded agreement.
4. Continue without asking the user only when the decision is bounded and evidence is sufficient.
5. If a bounded decision is needed, create a self-drive packet and ask a subagent.
6. Run the work, verification, reporting, and next-flow reopening under the base `turn-gate` contract.

Expected risky actions are not automatic stop points. If the initial `turn-gate` preparation explicitly recorded a risky action as approved, not approved, or assigned to a handoff checkpoint, follow that recorded boundary. Stop self-drive only when the risk is outside the initial agreement, the evidence no longer matches the recorded boundary, or a new approval boundary appears.

If the flow list reaches its last planned flow, run commit-readiness reporting as the final planned flow. Report whether the work is ready for a commit and what remains uncertain. Do not commit, stage for commit, push, open a PR, publish, or ask for execution approval as part of this self-drive flow. Commit/push/PR/publish are separate explicit user-gated handoffs.

## Boundary Recheck

At each flow start, before each subagent question, and before acting on each subagent answer, recheck:

- current `Continuity Guard`
- active flow record
- planned flow list position
- recorded scope and non-goals
- verification expectation
- expected risky actions and their recorded treatment
- approval boundaries and user-gated checkpoints
- incoming user messages since the packet was created

Return to same-plugin `turn-gate` user-gated question routing with `request_user_input` when you find:

- missing scope that changes the result
- a risky action outside the initial agreement
- a new destructive, irreversible, external-action, safety, platform, or tool approval boundary
- an unplanned commit, push, PR, publish, or release decision
- conflicting or superseding user input
- evidence too weak to choose among materially different outcomes

This is a routing fallback, not a terminal summary. Preserve the turn and ask the smallest question needed to continue.

## Question Packet

Every self-drive subagent packet must be bounded enough to answer without hidden context. Include:

- `question_id`
- `phase`
- `current_mode`
- `flow_id`
- `planned_flow_position`
- `decision_needed`
- `options`
- `context`
- `work_boundary`
- `non_goals`
- `acceptance_signal`
- `verification_expectation`
- `expected_risky_actions`
- `approval_boundaries`
- `user_gated_checkpoints`
- `user_interventions`
- `continuity_guard`
- `constraints`
- `fallback`
- `expected_answer`

The packet must tell the subagent to answer the bounded decision only, preserve turn continuity, flag context gaps, and provide a concrete next action rather than a completion summary.

## Answer Contract

Every self-drive subagent answer must include:

- `question_id`
- `selected_option`
- `decision`
- `rationale`
- `evidence`
- `assumptions`
- `confidence`
- `context_gap`
- `blockers`
- `approval_boundary`
- `expected_risky_action_match`
- `superseded_by_user_input`
- `continuity_check`
- `next_action`

Accept the answer only if it matches the active packet, is not superseded by newer user input, respects the recorded flow boundary, and includes a continuing `next_action`.

If `continuity_guard.user_explicit_stop` is false and no approval boundary is present, `continuity_check.terminal_summary_allowed` must be false. An answer that only reports completion is invalid. If an approval boundary is present, `next_action` must be `switch-to-user-gated-question`, and the main flow must return to same-plugin `turn-gate` question routing.

## Context Gap Recovery

Do not treat every missing detail as a user question. If the gap can be closed by repo search, file reads, logs, deterministic checks, or policy-allowed research inside the recorded boundary, recover the context and continue.

Ask the user only when the missing detail changes scope, non-goals, acceptance, verification, approval, external action, or risky-action treatment. Record the gap, recovery action, and result in the flow record.

## User Intervention

Any user message during self-drive is authoritative current loop input and outranks pending or returned subagent answers.

When user input arrives:

1. Refresh the `Continuity Guard`; keep `user_explicit_stop: false` unless the user explicitly ends the turn.
2. Classify the input as explicit turn stop, current-flow correction, current-flow priority change, next-flow priority request, approval answer, or scope update.
3. Mark any incompatible pending subagent answer as stale.
4. Re-read affected targets before continuing if the input changes files, artifacts, scope, or state.
5. Continue from the earliest safe phase under `turn-gate`, or route to user-gated questioning if the input creates a new boundary.

Do not call an intervention a pause, stop, or completion unless the user explicitly asks to end the turn.

## Flow-End Deterministic Reread

After each self-drive flow result report, and before asking or routing any next question, always refresh the contracts:

1. Re-read the same-plugin `turn-gate` skill.
2. Re-read this `turn-gate-self-drive` skill.
3. Recheck the current `Continuity Guard`, active flow record, pending next-flow candidates, planned flow position, and approval boundaries.
4. Route the next step through the refreshed base `turn-gate` contract and this overlay contract.

This reread is unconditional. Do not try to detect whether context cleanup or compaction happened. The purpose is to reload the base loop contract and overlay contract before next-flow questions, scope questions, self-drive packets, or user-gated approval questions.

The reread does not authorize self-drive to bypass approval, destructive, irreversible, external-action, safety, commit, push, PR, publish, or release boundaries.

## Confidence Rules

- `high`: evidence supports the decision, assumptions are minor, expected risky actions match the recorded preparation, and no new approval boundary is present.
- `medium`: gaps are recoverable inside the recorded boundary; continue after recording assumptions or running recovery.
- `low`: the decision changes scope, crosses an unprepared risky-action boundary, needs explicit approval, or lacks enough evidence for a bounded choice; switch to same-plugin `turn-gate` user-gated question routing.

## Output

When using this skill in a self-drive flow, keep these items available in the flow record and final flow report:

- `Base turn-gate contract`
- `Prepared planned-flow boundary`
- `Self-drive question packet`
- `Subagent answer`
- `Boundary recheck`
- `Context gap recovery`
- `User intervention handling`
- `Continuity check`
- `Recorded assumptions`
- `Commit-readiness report` for the last planned flow
- `Next action`
