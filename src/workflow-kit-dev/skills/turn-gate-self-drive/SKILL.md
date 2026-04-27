---
name: turn-gate-self-drive
description: Self-drive overlay for `turn-gate`. Use when a turn-gated workflow should keep moving without waiting for user answers by routing bounded decisions to subagents while preserving the base `turn-gate` continuity contract.
---

# Turn Gate Self Drive

## Overview

Use this plugin skill as a self-drive overlay that runs beside `turn-gate` in `workflow-kit-dev`.
First apply the base loop contract from the `turn-gate` skill in the same plugin: analyze, plan, work, verify, report, maintain session records, refresh the `Continuity Guard`, and reopen the next flow.
Then route bounded decisions into subagent question packets.

This skill owns self-drive packet, answer, stale-answer, pause, and recovery rules.
It is not a surface owned by `turn-gate`; it is a `workflow-kit-dev` plugin skill that depends on `turn-gate` for base loop continuity.
It does not own the base loop gate or the current-phase workflow.

## Use When

- The user explicitly wants autonomous continuation, self-driving progress, or no user intervention between phases.
- The task is already governed by `turn-gate`.
- A blocked decision can be answered from a bounded packet of available evidence.

## Do Not Use When

- The user wants manual approval at each step.
- The next decision requires explicit user, tool, platform, safety, destructive, irreversible, or external-action approval.
- The work does not need `turn-gate` continuity.

## Core Contract

- Treat the `turn-gate` skill in the same plugin as the base contract.
- Build a self-drive question packet before asking a subagent.
- Ask subagents to resolve mode selection, criteria, scope assumptions, verification choices, and next-flow decisions from packet evidence.
- Continue the loop using the subagent answer as the current decision input.
- Record the subagent question, answer, confidence, assumptions, and stale-answer handling in the flow record.
- Carry the current `Continuity Guard` into every packet.
- Reject any answer that reports completion without preserving next-flow continuation when no explicit user stop or hard approval boundary exists.
- Treat newer user messages as higher-priority loop input than pending or returned subagent answers.
- At explicit approval boundaries, pause self-drive only, switch back to user-gated routing from the `turn-gate` skill in the same plugin, and call `request_user_input`.

## Question Packet

Every packet must include: `question_id`, `phase`, `current_mode`, `question_type`, `decision_needed`, `options`, `context`, `user_interventions`, `continuity_guard`, `constraints`, `fallback`, and `expected_answer`.

## Answer Contract

Every answer must include: `question_id`, `selected_option`, `decision`, `rationale`, `evidence`, `assumptions`, `confidence`, `context_gap`, `blockers`, `approval_boundary`, `superseded_by_user_input`, `continuity_check`, and `next_action`.

If no explicit user stop or approval boundary exists, `continuity_check.terminal_summary_allowed` must be false and `next_action` must continue the loop.
If an approval boundary exists, `next_action` must be `switch-to-user-gated-question`.
Answers that only summarize completion are invalid.

## User Intervention

When a user message arrives during self-drive:

1. Treat it as authoritative loop input.
2. Keep `user_explicit_stop: false` unless the message explicitly asks to end the turn.
3. Classify it as explicit turn stop, current-flow correction, current-flow priority change, or next-flow priority request.
4. Supersede any pending or returned subagent answer that conflicts with the newer message.
5. Continue from the earliest safe phase, or switch to user-gated routing from the `turn-gate` skill in the same plugin if a real approval boundary appears.

## Output

- `Base turn-gate contract`
- `Self-drive question packet`
- `Subagent answer`
- `Continuity check`
- `Recorded assumptions`
- `Next action`
