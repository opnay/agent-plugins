---
name: workflow-kit-guide
description: Default first-read skill for incoming requests. Use to choose which `workflow-kit` skill should own the current workflow and what handoff order should follow.
---

# Workflow Kit Guide

## Overview

Use this skill as the default entrypoint for `workflow-kit`.
Treat this skill as the normal first routing layer for incoming requests.
Choose the starting skill in this plugin that most directly advances the user's result.
Prefer one clear starting skill plus an explicit handoff.
Do not route to a specialist plugin first from the global layer; pick the workflow first.
Skip this guide only when a narrower `workflow-kit` skill is already clearly the right starting point.

## Workflow

1. Identify which bundled skill is the best starting point:
   - `structured-thinking`
   - `deep-interview`
   - `planner`
   - `autopilot`
   - `parallel-work`
   - `ralph-loop`
   - `review-loop`
   - `commit-readiness-gate`
2. Decide whether the work is:
   - still active execution
   - review-driven correction
   - final readiness checking
3. If the work belongs in execution, estimate:
   - scope size
   - verification style
   - whether the brief is actually locked yet
4. Route to the narrowest bundled skill that owns the current bottleneck.
5. If the task spans several bundled skills, choose the starting skill and handoff point explicitly.

## Routing Rules

- Choose `structured-thinking` when the task is still too unstable to choose the next workflow safely and the first job is to isolate ambiguity, assumptions, and the most plausible next path.
- Choose `deep-interview` when the main job is to understand the user's real intent, boundaries, tradeoffs, approval lines, or success criteria through questions, pressure-testing, or direction evaluation.
- Choose `planner` when implementation should stay deferred until a read-only investigation, tradeoff analysis, verification path, and execution-ready plan are complete.
- Choose `autopilot` when the user wants broad end-to-end delivery from brief to verified result.
- Choose `parallel-work` when the work contains several bounded lanes with a clear integration path.
- Choose `ralph-loop` when the work is best handled as repeated bounded fix-verify-reassess cycles.
- Choose `review-loop` when the input is review feedback or findings and the job is to fix only the material issues one at a time.
- Choose `commit-readiness-gate` when implementation is largely done and the main question is whether the current change is isolated, verified, risk-classified, and ready to move toward commit.
- If the task clearly begins as execution and ends with a commit-readiness pass, start with the fitting execution workflow and hand off to `commit-readiness-gate` only after the intended change unit is complete.

## Decision Rules

- Treat this guide as the default first stop for routing.
- Pick the skill that best addresses the current bottleneck.
- Choose `deep-interview` over `structured-thinking` when the request is a proposal, direction check, greenfield setup, or "괜찮을까?" style evaluation and the missing information is about what the user really wants, what constraints matter most, or what success should mean.
- Keep `structured-thinking` for workflow ambiguity, not for requirement discovery that should continue as an actual interview.
- Use `deep-interview` when the blocker is still understanding what the user actually wants, where the scope should stop, or how to evaluate a direction before committing to a plan or implementation.
- Use clarification skills when intent, scope, tradeoff, or approval boundaries are still the blocker.
- Use planning when execution should remain deferred.
- Use execution workflows when meaningful implementation or refinement still remains.
- Use `commit-readiness-gate` only when the main question is readiness.
- Prefer one clear starting skill plus an explicit handoff.

## Execution Mode Rules

Use these rules when the task already belongs in execution.

Estimate scope size as:

- `single bounded issue`
- `several related issues`
- `multi-phase feature or product request`

Estimate verification style as:

- `immediate narrow verification after each fix`
- `broader phased verification`
- `review-driven verification against findings`

Choose execution mode with these rules:

- Hand off to `deep-interview` when execution mode selection is premature because the implementation brief is still materially misaligned or underspecified.
- Choose `autopilot` for broad end-to-end delivery that spans requirements, implementation, testing, and validation.
- Choose `parallel-work` when the work contains several clearly independent lanes with an explicit integration path.
- Choose `ralph-loop` when the task is bounded and benefits from repeated fix-verify-reassess cycles.
- Choose `review-loop` when the input is a set of findings and only material items should block progress.
- If the work starts as parallel lanes but shared dependencies appear, collapse back to a sequential or broader workflow.
- If the work starts as bounded refinement but expands into several coordinated areas, escalate to a broader workflow.
- If the work starts broad but enters a short local polish cycle, temporarily use a bounded loop without losing the broader scope frame.

Execution heuristics:

- Treat the task as broad end-to-end delivery when:
  - the user wants a brief idea taken all the way to working code
  - the work clearly spans requirements, implementation, testing, and validation
  - several phases are needed before delivery is complete
- Treat the task as bounded iterative refinement when:
  - the work can be improved one bounded issue at a time
  - repeated fix-verify-reassess cycles are more useful than one large pass
  - the user wants polish, stabilization, or gradual improvement
- Treat the task as parallel independent work when:
  - the work contains several bounded tasks that can proceed without waiting on each other
  - the main challenge is safe split and integration rather than long phased delivery
  - running the tasks sequentially would waste time more than it would reduce risk
- Treat the task as blocking-first review handling when:
  - the input is review feedback, QA findings, or self-review findings
  - the goal is to fix only material issues without stalling delivery for low-value polish

## Escalation Signals

- The current issue can no longer be solved as one bounded unit.
- The work keeps collapsing back into requirement clarification rather than execution decisions.
- The supposedly independent lanes now share files, decisions, or prerequisites.
- The work begins to require explicit multi-phase planning.
- Review findings reveal a deeper structural problem than the current mode can safely own.
- Repeated loops stop producing meaningful improvement.
- Verification cost grows beyond the selected mode's intended scope.

## Output Contract

- `Current bottleneck`
- `Scope size` when execution is selected
- `Verification style` when execution is selected
- `Chosen skill`
- `Planned handoff`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not treat another plugin as the default global entrypoint.
- Do not skip an obvious alignment pass when the user's intent is still materially underspecified.
- Do not skip a framing pass when workflow selection itself is still unstable.
- Do not skip a planning pass when execution should remain deferred until read-only investigation and handoff are complete.
- Do not treat a work-defining prompt as a simple opinion request when the likely next step is real task execution.
- Do not stay in execution-mode selection when the real blocker is unclear user intent or an unlocked scope boundary.
- Do not send in-progress implementation work straight to the commit gate.
- Do not use execution workflows when the real need is a final readiness decision.
- Do not bury plugin-level routing guidance inside sibling skills.
- Do not blend clarification, execution, and gate responsibilities without naming which one starts and when the handoff happens.
- Do not choose a broad workflow for a narrow bounded issue just because the task sounds important.
- Do not choose a parallel workflow unless independence and integration risk are both explicit.
- Do not choose a bounded loop when the work clearly needs planning across several phases.
- Do not let review handling turn into open-ended polish by default.
- Do not keep the initial execution mode if the task shape has obviously changed.
