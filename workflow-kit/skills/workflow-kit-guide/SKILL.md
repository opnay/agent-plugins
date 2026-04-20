---
name: workflow-kit-guide
description: Entrypoint skill for the `workflow-kit` plugin. Use when a task needs the right clarification workflow, execution mode, or final gate chosen first, such as deciding whether work should begin with deep intent alignment, active execution, or a commit-readiness pass.
---

# Workflow Kit Guide

## Overview

Use this skill as the default entrypoint for `workflow-kit`.
Its job is to classify whether the task is primarily about clarifying implementation intent, choosing the right execution mode, or deciding whether the work is ready to move to commit.
Do not jump straight into execution when the real need is a deep alignment pass, and do not run the commit gate when the work is still clearly in progress.
When execution is the right branch, this guide should also choose the concrete execution mode instead of delegating that choice to another sibling guide.

## Workflow

1. Identify the task shape:
   - pre-execution alignment and intent clarification
   - broad end-to-end delivery
   - parallel independent execution
   - bounded iterative refinement
   - blocking-first review handling
   - commit-readiness decision before review or commit
2. Decide whether the change is:
   - still actively being implemented
   - nearly done and ready for final gate checks
3. If the work belongs in execution, estimate:
   - scope size
   - verification style
   - whether the execution brief is actually locked yet
4. Route to the narrowest bundled skill that owns the main concern.
5. If the task spans clarification, execution, and gate concerns, pick the starting skill and explicit handoff point.

## Routing Rules

- Choose `deep-interview` when the user already wants implementation, but the main risk is misreading intent, boundaries, tradeoffs, or approval lines before execution begins.
- Choose `autopilot` when the user wants broad end-to-end delivery from brief to verified result.
- Choose `parallel-work` when the work contains several bounded lanes with a clear integration path.
- Choose `ralph-loop` when the work is best handled as repeated bounded fix-verify-reassess cycles.
- Choose `review-loop` when the input is review feedback or findings and the job is to fix only the material issues one at a time.
- Choose `commit-readiness-gate` when implementation is largely done and the main question is whether the current change is isolated, verified, risk-classified, and ready to move toward commit.
- If the task clearly begins as execution and ends with a commit-readiness pass, start with the fitting execution workflow and hand off to `commit-readiness-gate` only after the intended change unit is complete.

## Decision Rules

- Choose `deep-interview` when alignment quality is the main blocker and implementation should follow a locked brief.
- Choose a gate only when the main uncertainty is readiness, not implementation approach.
- Choose an execution workflow when meaningful implementation, refinement, or findings work still remains.
- Do not use `commit-readiness-gate` as a substitute for actual review handling or incomplete implementation.
- Prefer one clear starting skill plus an explicit handoff over mixing clarification, execution, and gate responsibilities into one prompt.

## Execution Mode Rules

Use these rules when the task already belongs in execution rather than a clarification pass or a final gate.

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

Execution-shape heuristics:

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

- `Task shape`
- `Scope size` when execution is selected
- `Verification style` when execution is selected
- `Clarification vs execution vs gate decision`
- `Chosen skill`
- `Why this route fits`
- `Escalation signal`
- `Planned handoff`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not skip an obvious alignment pass when the user's intent is still materially underspecified.
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

## Example Triggers

- "이 작업에 어떤 workflow로 들어가는 게 맞는지 먼저 판단해줘"
- "이거 병렬로 나눠서 처리하는 게 맞는지 먼저 판단해줘"
- "이거 끝까지 미는 게 맞아, 아니면 루프로 다듬는 게 맞아?"
- "리뷰 반영인지, 구현 workflow인지 먼저 구분해줘"
