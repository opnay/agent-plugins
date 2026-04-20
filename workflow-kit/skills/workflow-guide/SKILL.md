---
name: workflow-guide
description: Execution workflow selector for active implementation work. Use when a task already belongs in workflow execution rather than final gating and needs the right mode chosen before work begins, such as deciding between end-to-end delivery, parallel independent work, bounded iterative refinement, or blocking-first review handling based on scope, risk, and verification style.
---

# Workflow Guide

## Overview

Use this skill when the work is still actively being executed and the main question is which execution mode fits best.
Its job is not to do the whole task by itself. Its job is to classify the execution shape, choose the right mode, and make sure the work starts with the right rhythm and stop conditions.
Do not use this skill as the plugin-wide entrypoint when the real question is whether the change first needs a deep alignment pass or is already ready for a commit gate.

## Workflow

1. Identify the task shape:
   - broad end-to-end delivery
   - parallel independent work
   - bounded iterative refinement
   - blocking-first review handling
2. Estimate the scope size:
   - single bounded issue
   - several related issues
   - multi-phase feature or product request
3. Identify the verification style:
   - immediate narrow verification after each fix
   - broader phased verification
   - review-driven verification against findings
4. Decide whether execution mode selection is actually ready, or whether the work first needs a locked execution brief.
5. If user intent, boundaries, or tradeoffs are still unclear, hand off to `deep-interview` before choosing an execution mode.
6. Otherwise, select the execution mode that best matches the task shape.
7. State what should trigger escalation or workflow switching if the initial mode stops fitting.
8. If the work becomes nearly done and the next question is readiness rather than execution, hand off to `commit-readiness-gate`.

## Task Shape Rules

Treat the task as broad end-to-end delivery when:

- the user wants a brief idea taken all the way to working code
- the work clearly spans requirements, implementation, testing, and validation
- several phases are needed before delivery is complete

Treat the task as bounded iterative refinement when:

- the work can be improved one bounded issue at a time
- repeated fix-verify-reassess cycles are more useful than one large pass
- the user wants polish, stabilization, or gradual improvement

Treat the task as parallel independent work when:

- the work contains several bounded tasks that can proceed without waiting on each other
- the main challenge is safe split and integration rather than long phased delivery
- running the tasks sequentially would waste time more than it would reduce risk

Treat the task as blocking-first review handling when:

- the input is review feedback, QA findings, or self-review findings
- the goal is to fix only material issues without stalling delivery for low-value polish

## Selection Rules

- Hand off to `deep-interview` when execution mode selection is premature because the implementation brief is still materially misaligned or underspecified.
- Choose a phase-driven workflow when the task is broad and end-to-end.
- Choose a parallel workflow when the task contains several clearly independent lanes with an explicit integration path.
- Choose a loop-driven workflow when the task is bounded and benefits from repeated fix-verify-reassess cycles.
- Choose a blocking-first review workflow when the input is a set of findings and only material items should block progress.
- If the work starts as parallel lanes but shared dependencies appear, collapse back to a sequential or broader workflow.
- If the work starts as bounded refinement but expands into several coordinated areas, escalate to a broader workflow.
- If the work starts broad but enters a short local polish cycle, temporarily use a bounded loop without losing the broader scope frame.

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
- `Scope size`
- `Primary risk`
- `Verification style`
- `Selected workflow mode`
- `Why this mode fits`
- `Escalation signal`
- `Residual risk`

## Guardrails

- Do not use this skill when the task is primarily a final commit-readiness judgment.
- Do not stay in execution-mode selection when the real blocker is unclear user intent or an unlocked scope boundary.
- Do not choose a broad workflow for a narrow bounded issue just because the task sounds important.
- Do not choose a parallel workflow unless independence and integration risk are both explicit.
- Do not choose a bounded loop when the work clearly needs planning across several phases.
- Do not let review handling turn into open-ended polish by default.
- Do not keep the initial workflow mode if the task shape has obviously changed.

## Concept Lenses

Use these lenses implicitly as needed:

- Scope lens:
  - how wide the task really is
- Verification lens:
  - how often and how narrowly the work should be checked
- Delivery lens:
  - whether the task is about shipping, refining, or burning down findings
- Escalation lens:
  - when the current execution mode has stopped fitting the task

## Example Triggers

- "이 작업에 어떤 workflow로 들어가는 게 맞는지 먼저 판단해줘"
- "이거 병렬로 나눠서 처리하는 게 맞는지 먼저 판단해줘"
- "이거 끝까지 미는 게 맞아, 아니면 루프로 다듬는 게 맞아?"
- "리뷰 반영인지, 구현 workflow인지 먼저 구분해줘"
