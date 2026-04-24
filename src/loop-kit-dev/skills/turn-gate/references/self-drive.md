# Self-Drive Reference

Use this question-routing mode when the turn should keep progressing without waiting for user answers.

## Core Contract

- Replace phase questions that would normally go to the user with questions to subagents.
- Build a self-drive question packet before asking a subagent.
- Ask subagents to resolve mode selection, criteria, scope assumptions, verification choices, and next-flow decisions from available evidence in that packet.
- Continue the loop using the subagent answer as the current decision input.
- Record the subagent question, answer, confidence, and any assumptions in the flow record.
- Carry the current `Continuity Guard` into every self-drive packet so long contexts preserve the active loop rule.
- Reject any self-drive answer that reports completion but does not preserve next-flow continuation when no explicit user stop or hard approval boundary exists.
- If a platform, tool, or safety policy requires explicit user approval, pause self-drive at that boundary, switch to `user-gated`, and open `request_user_input` with explicit choices instead of ending the turn.

## Question Packet

Every self-drive subagent question must include:

- `question_id`: stable short id for linking the answer back to the flow record.
- `phase`: current turn-gate phase, such as analysis, plan, work, verification, result-reporting, or next-flow.
- `current_mode`: selected current-phase mode, or `undecided`.
- `question_type`: mode-selection, criteria, scope-assumption, verification-choice, next-flow, or blocker-triage.
- `decision_needed`: the exact decision the subagent must make.
- `options`: explicit options when the decision can be bounded.
- `context`: concise evidence from the current request, files, diffs, logs, plans, or prior flow records.
- `continuity_guard`: compact current-state reminder with `turn_gate_active`, `question_routing_mode`, `user_explicit_stop`, `terminal_summary_allowed`, and `required_next_action`.
- `constraints`: non-goals, approval boundaries, safety/tool limits, and anything the subagent must not assume.
- `fallback`: what to do if confidence is too low or evidence conflicts.
- `expected_answer`: required answer shape for the subagent response.

The packet should be small enough for the subagent to answer directly.
If the packet needs broad research before the question is answerable, split it into a discovery subtask first.

## Answer Contract

Every self-drive subagent answer must include:

- `question_id`: the packet id being answered.
- `selected_option`: chosen option when the packet provided options, or `none` if no option can be selected.
- `decision`: the concrete decision to use as the next loop input.
- `rationale`: short reason for the decision.
- `evidence`: specific evidence from the packet context or discovered facts.
- `assumptions`: assumptions the loop will carry forward if it continues.
- `confidence`: high, medium, or low.
- `context_gap`: missing information the subagent needs before it can answer better, if any.
- `blockers`: unresolved blockers or conflicts, if any.
- `approval_boundary`: whether explicit user/tool approval is required before continuing.
- `continuity_check`: whether `turn-gate` remains active, whether terminal summary is allowed, and what next-flow action is required.
- `next_action`: the next phase or action the loop should take.

If `continuity_guard.user_explicit_stop` is false and no approval boundary is present, the answer must set `continuity_check.terminal_summary_allowed: false` and provide a concrete continuing `next_action`.
If an approval boundary is present, the answer must set `next_action` to `switch-to-user-gated-question` and the main agent must call `request_user_input` rather than end the turn.
An answer that only summarizes completion without a next action is invalid in self-drive.
If the subagent cannot answer because recoverable context is missing, it must return `selected_option: none`, `confidence: medium`, and `context_gap`.
If a missing user preference is not recoverable, choose the safest reversible default, record it as an assumption, and continue.
Only return `selected_option: none`, `confidence: low`, and a blocker when the missing decision requires explicit approval, destructive/irreversible/external action approval, or a platform/tool/safety boundary.

## Confidence Rules

- `high`: evidence in the packet supports the decision, assumptions are minor, and no approval boundary is present. Continue automatically.
- `medium`: the decision is plausible but has recoverable gaps or assumptions. Continue only by recording carried assumptions or by running the recovery action named in `next_action`.
- `low`: self-drive cannot continue because the decision requires explicit approval, destructive/irreversible/external action approval, or a platform/tool/safety boundary. Pause self-drive at that boundary, switch to `user-gated`, and open `request_user_input` with explicit choices.

Do not classify a missing context item as `low` merely because it was absent from the packet.
If the main agent can recover the gap through repo search, file reads, logs, deterministic checks, or policy-allowed web research, treat it as recoverable.

## Context-Gap Recovery

When a subagent returns `context_gap`:

1. Classify whether the gap is recoverable by the main agent.
2. If recoverable, run the needed discovery and build a new self-drive question packet with the added evidence.
3. Ask the subagent again with the revised packet.
4. If the gap is an unrecoverable user preference, choose the safest reversible default, record it as an assumption, and continue.
5. If unrecoverable because it requires explicit approval, destructive/irreversible/external action approval, or a policy boundary, pause self-drive with `confidence: low`, switch to `user-gated`, and open `request_user_input`.

## Mode Boundary

- Good fit: user explicitly wants autonomous continuation, self-driving loops, or no user intervention between phases.
- Not a fit: the user explicitly wants manual approval at each step, or the next step requires real-world/destructive/irreversible approval that the runtime or tool policy requires the user to approve.

## Review Questions

- Did every non-approval question route to a subagent instead of the user?
- Did each subagent question include a complete self-drive question packet?
- Did the subagent answer follow the self-drive answer contract?
- Did the packet include the current `Continuity Guard`, and did the answer include a continuity check?
- Did the subagent answer include enough evidence, confidence, and assumptions to continue responsibly?
- Were recoverable context gaps routed back through main-agent discovery instead of treated as approval-boundary pauses?
- Were ordinary preference gaps converted into recorded reversible assumptions instead of stop conditions?
- Did any completion report still provide a next-flow action unless an explicit stop or hard approval boundary existed?
- When self-drive paused at a hard boundary, did the main agent switch to `user-gated` and use `request_user_input` instead of ending the turn?
- Did the loop preserve hard approval boundaries required by the runtime or tool policy?
