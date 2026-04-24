# Self-Drive Reference

Use this question-routing mode when the turn should keep progressing without waiting for user answers.

## Core Contract

- Replace phase questions that would normally go to the user with questions to subagents.
- Ask subagents to resolve mode selection, criteria, scope assumptions, verification choices, and next-flow decisions from available evidence.
- Continue the loop using the subagent answer as the current decision input.
- Record the subagent question, answer, confidence, and any assumptions in the flow record.
- If a platform, tool, or safety policy requires explicit user approval, stop at that approval boundary instead of simulating consent.

## Mode Boundary

- Good fit: user explicitly wants autonomous continuation, self-driving loops, or no user intervention between phases.
- Not a fit: the user wants manual approval at each step, the next step requires real-world/destructive approval, or the decision depends on private user preference that no subagent can infer from available context.

## Review Questions

- Did every non-approval question route to a subagent instead of the user?
- Did the subagent answer include enough confidence and assumptions to continue responsibly?
- Did the loop preserve hard approval boundaries required by the runtime or tool policy?
