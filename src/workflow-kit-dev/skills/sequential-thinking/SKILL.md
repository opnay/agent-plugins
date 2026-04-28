---
name: sequential-thinking
description: Use for complex problem solving, analysis, planning, design, or debugging that needs a self-contained sequential thinking workflow with revision, branching, hypothesis generation, and verification before a final answer.
---

# Sequential Thinking

## Purpose

Use this skill when the task needs a deliberate multi-step thinking workflow rather than a one-shot answer.
It is a self-contained skill conversion of the Sequential Thinking MCP server's reflective problem-solving behavior.
Do not call or require an MCP tool to use this skill.
Reference model: https://github.com/modelcontextprotocol/servers/tree/main/src/sequentialthinking

This is not a router.
Do not use it merely to decide which workflow should run next.
Use it when the current work itself needs sequential analysis, course correction, or hypothesis verification.

## Entry Decision

Use this skill for:

- complex analysis where the answer is not obvious at the start
- planning or design work that may need revision as constraints become clearer
- debugging or diagnosis that needs hypotheses and verification
- comparison of multiple approaches where one branch may need to be explored or rejected
- problems whose scope may change while reasoning progresses
- tasks where irrelevant information must be filtered out step by step

Do not use this skill when:

- the blocker is unclear user intent, scope, tradeoffs, or acceptance criteria; hand off to `deep-interview` and state the smallest question or handoff needed
- the requested deliverable is an execution-ready plan artifact with steps, expected outputs, and verification methods; hand off to `planner`
- the task is a straightforward implementation request; hand off to the fitting execution workflow
- the user asked only for a final readiness decision; hand off to `commit-readiness-gate`
- the problem is already simple enough for a direct answer

## Working Notes

Maintain compact working notes while using this skill.
These notes replace the state a tool would otherwise track for you.
Keep them concise and task-facing; do not expose raw private chain-of-thought.
Step estimates are for steering the work, not a required user-facing output.

Track:

- current step and estimated remaining steps
- whether another reasoning step is needed before conclusion
- any revised earlier assumption or conclusion
- any alternate branch that was explored and why it was kept or rejected
- constraints, non-goals, evidence boundaries, or resource limits that shape the answer
- the current hypothesis when one exists
- the verification that supports or rejects that hypothesis

Working notes are an internal aid for maintaining flow state.
Report only the reasoning checkpoints, verification state, and decision-relevant constraints the user needs.

## Sequential Progress

Start with an initial estimate of the needed steps, then adjust that estimate as the problem becomes clearer.
The estimate is not a promise and is not the stopping condition.
Continue until the answer is sufficiently supported, not merely until the original estimate has been consumed.

For each step:

- advance the problem with new information, a refined assumption, a constraint, or a decision point
- avoid repeating earlier conclusions unless they are being revised or compared
- filter out irrelevant information instead of carrying it forward
- keep the next-step-needed state current

If excluded information materially affects how the user should understand the conclusion, mention the exclusion and why it did not drive the answer.

## Revision and Branching

When new evidence shows an earlier assumption, interpretation, or conclusion was wrong or incomplete, revise it explicitly.
Do not present a revised assumption as if it had always been established fact.

Use branching only when a meaningful competing path needs separate treatment.
For each material branch, track:

- what alternative it represents
- what evidence or constraint supports it
- why it was kept, rejected, or deferred

Do not expand branches just because possibilities exist.
Branching should reduce uncertainty, compare realistic options, or preserve a material alternative for handoff.

## Hypothesis and Verification

When a plausible answer emerges, treat it as a hypothesis until it has been checked.
Choose verification that matches the task: code inspection, tests, docs, user-provided facts, external constraints, counterexamples, or consistency checks.

Report verification honestly:

- completed checks and their results
- checks that were unavailable or intentionally skipped
- proposed checks that would be needed before stronger confidence

Do not imply verification was completed when it was only proposed.
For debugging, diagnosis, and high-impact recommendations, do not stop at a plausible narrative if a direct check is available.

## Constraints and Multi-Cause Analysis

Keep decision-shaping constraints visible throughout the flow.
This includes non-goals, resource limits, ownership boundaries, process requirements, evidence gaps, external dependencies, and operational constraints.

When a problem has multiple causes, do not collapse it to the easiest technical cause if process, ownership, handoff, or external evidence materially affects the answer.
When a recommendation depends on non-code evidence or broader constraints, carry those constraints into the conclusion rather than reducing the answer to the most convenient technical path.

Do not treat one narrow source of evidence as sufficient when the decision depends on missing negative cases, broader constraints, or unverified assumptions.

## Reporting and Handoff

When reporting, include only the task-relevant summary:

- `Conclusion`
- `Key reasoning checkpoints`
- `Revisions or branches considered`
- `Ignored or deferred information` when filtering materially affected the answer
- `Verification`, distinguishing completed checks from proposed checks when evidence is unavailable
- `Constraints that shaped the answer` when non-goals, evidence limits, ownership, or resource limits materially affect the recommendation
- `Residual risk`

If more user input is required, ask the smallest question that would change the reasoning path.
Use `request_user_input` when the question can be expressed as bounded choices.

Hand off instead of continuing in this skill when the bottleneck changes:

- `deep-interview`: the next useful step is discovering user intent, scope, tradeoffs, acceptance criteria, or missing requirements
- `planner`: the next useful step is producing a read-only execution plan with steps, expected outputs, and verification methods
- `commit-readiness-gate`: the next useful step is a final readiness decision for a commit or similar completion gate
- execution workflow: the next useful step is implementing, editing files, running tests, or otherwise performing the chosen work

When handing off, state why the current sequential problem-solving flow is no longer the bottleneck and identify the next surface.
If only a small clarification is needed before continuing, ask that question rather than converting the task into a broader interview.

## Guardrails

- Do not expose raw hidden chain-of-thought, full thought logs, or exhaustive internal traces.
- Do not keep branching after the important uncertainty is resolved.
- Do not force a fixed number of steps when the task no longer needs them.
- Do not skip hypothesis verification for debugging, diagnosis, or high-impact recommendations.
- Do not treat one narrow evidence source as sufficient when the decision depends on broader constraints or missing negative cases.
- Do not reduce a multi-cause analysis to the easiest technical cause when process, ownership, handoff, or external evidence is material.
- Do not present a revised assumption as fact without saying it changed.
- Do not use this skill as a replacement for `deep-interview`, `planner`, `commit-readiness-gate`, or execution workflows.
- Do not require Sequential Thinking MCP server installation or MCP tool calls.
