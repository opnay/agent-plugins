---
name: turn-gate-self-drive
description: Self-drive overlay for `turn-gate`. Use when a turn-gated task should keep progressing without waiting for user answers by routing blocked decisions to subagents while preserving the `turn-gate` continuity contract.
---

# Turn Gate Self Drive

## Overview

Use this plugin skill as a self-drive overlay that runs beside `turn-gate` in `loop-kit-dev`.
First apply the base contract from the `turn-gate` skill in the same plugin: keep the same turn alive, maintain `.agents/sessions/{YYYYMMDD}/` records, preserve the `Continuity Guard`, verify before result reporting, and reopen the next flow.
Then route bounded decisions into self-drive question packets sent to subagents.

This skill owns self-drive packet, answer, pause, recovery, and stale-answer handling.
It is not a surface owned by `turn-gate`; it is a `loop-kit-dev` plugin skill that depends on `turn-gate` for base loop continuity.
It does not replace `turn-gate`, and it does not own the current-phase loop modes.

## Use When

- The user explicitly wants autonomous continuation, self-driving progress, or no user intervention between phases.
- The work is already governed by `turn-gate`, but blocked decisions can be answered from available evidence.
- Mode selection, criteria, scope assumptions, verification choices, or next-flow choices can be delegated to a subagent.

## Do Not Use When

- The user wants manual approval at each step.
- A decision requires explicit user, tool, platform, safety, destructive, irreversible, or external-action approval.
- The task does not need `turn-gate` continuity.

## Core Contract

- Use the `turn-gate` skill in the same plugin as the base loop contract for every phase.
- Build a self-drive question packet before asking a subagent.
- Ask subagents to resolve bounded decisions from packet evidence.
- Continue the loop using the subagent answer as the current decision input.
- Record the subagent question, answer, confidence, assumptions, and stale-answer handling in the flow record.
- Carry the current `Continuity Guard` into every self-drive packet.
- Reject any answer that reports completion but does not preserve next-flow continuation when no explicit user stop or hard approval boundary exists.
- Treat any user message that arrives during self-drive as higher-priority loop input than pending or returned subagent answers.
- If explicit approval is required, pause self-drive only, switch back to user-gated routing from the `turn-gate` skill in the same plugin, and call `request_user_input`.

## User Intervention

When a user message arrives while self-drive work is in progress:

1. Treat the message as authoritative current loop input.
2. Refresh the `Continuity Guard` and preserve `user_explicit_stop: false` unless the message explicitly asks to end the turn.
3. Classify the message as `explicit-turn-stop`, `current-flow-correction`, `current-flow-priority-change`, or `next-flow-priority-request`.
4. For current-flow correction or priority change, revise analysis and plan immediately, then continue from the earliest safe phase.
5. For next-flow priority request, record it as the highest-priority next-flow candidate and continue to the next safe handoff point.
6. Ignore or supersede any subagent answer that conflicts with the newer user message.

Do not call user intervention a pause, stop, or completion unless the user explicitly asks to end the turn.

## Question Packet

Every self-drive subagent question must include:

- `question_id`
- `phase`
- `current_mode`
- `question_type`
- `decision_needed`
- `options`
- `context`
- `user_interventions`
- `continuity_guard`
- `constraints`
- `fallback`
- `expected_answer`

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
- `superseded_by_user_input`
- `continuity_check`
- `next_action`

If `continuity_guard.user_explicit_stop` is false and no approval boundary is present, the answer must set `continuity_check.terminal_summary_allowed: false` and provide a concrete continuing `next_action`.
If an approval boundary is present, the answer must set `next_action` to `switch-to-user-gated-question` and the main agent must call `request_user_input` rather than end the turn.
An answer that only summarizes completion without a next action is invalid.

## Confidence Rules

- `high`: evidence supports the decision, assumptions are minor, and no approval boundary is present.
- `medium`: the decision has recoverable gaps or assumptions; continue by recording assumptions or running the recovery action.
- `low`: the decision requires explicit approval, destructive/irreversible/external action approval, or a platform/tool/safety boundary; pause self-drive and switch back to user-gated `turn-gate`.

Do not classify a missing context item as `low` if the main agent can recover it through repo search, file reads, logs, deterministic checks, or policy-allowed research.

## Output

- `Base turn-gate contract`
- `Self-drive question packet`
- `Subagent answer`
- `Continuity check`
- `Recorded assumptions`
- `Next action`
