# Self-Drive Reference

Use this question-routing mode when the turn should keep progressing without waiting for user answers.

## Core Contract

- Replace phase questions that would normally go to the user with questions to subagents.
- Build a self-drive question packet before asking a subagent.
- Ask subagents to resolve mode selection, criteria, scope assumptions, verification choices, and next-flow decisions from available evidence in that packet.
- Continue the loop using the subagent answer as the current decision input.
- Record the subagent question, answer, confidence, and any assumptions in the flow record.
- If a platform, tool, or safety policy requires explicit user approval, stop at that approval boundary instead of simulating consent.

## Question Packet

Every self-drive subagent question must include:

- `question_id`: stable short id for linking the answer back to the flow record.
- `phase`: current turn-gate phase, such as analysis, plan, work, verification, result-reporting, or next-flow.
- `current_mode`: selected current-phase mode, or `undecided`.
- `question_type`: mode-selection, criteria, scope-assumption, verification-choice, next-flow, or blocker-triage.
- `decision_needed`: the exact decision the subagent must make.
- `options`: explicit options when the decision can be bounded.
- `context`: concise evidence from the current request, files, diffs, logs, plans, or prior flow records.
- `constraints`: non-goals, approval boundaries, safety/tool limits, and anything the subagent must not assume.
- `fallback`: what to do if confidence is too low or evidence conflicts.
- `expected_answer`: required answer shape for the subagent response.

The packet should be small enough for the subagent to answer directly.
If the packet needs broad research before the question is answerable, split it into a discovery subtask first.

## Mode Boundary

- Good fit: user explicitly wants autonomous continuation, self-driving loops, or no user intervention between phases.
- Not a fit: the user wants manual approval at each step, the next step requires real-world/destructive approval, or the decision depends on private user preference that no subagent can infer from available context.

## Review Questions

- Did every non-approval question route to a subagent instead of the user?
- Did each subagent question include a complete self-drive question packet?
- Did the subagent answer include enough confidence and assumptions to continue responsibly?
- Did the loop preserve hard approval boundaries required by the runtime or tool policy?
