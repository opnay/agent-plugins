---
name: workflow-guide
description: Default workflow selector for execution-heavy tasks. Use when a task needs the right execution mode chosen before work begins, such as deciding between end-to-end delivery, bounded iterative refinement, or blocking-first review handling based on scope, risk, and verification style.
---

# Workflow Guide

## Overview

Use this skill as the default entrypoint for execution-heavy work in this kit.
Its job is not to do the whole task by itself. Its job is to classify the task shape, choose the right execution mode, and make sure the work starts with the right rhythm and stop conditions.

## Workflow

1. Identify the task shape:
   - broad end-to-end delivery
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
4. Select the execution mode that best matches the task shape.
5. State what should trigger escalation or workflow switching if the initial mode stops fitting.

## Task Shape Rules

Treat the task as broad end-to-end delivery when:

- the user wants a brief idea taken all the way to working code
- the work clearly spans requirements, implementation, testing, and validation
- several phases are needed before delivery is complete

Treat the task as bounded iterative refinement when:

- the work can be improved one bounded issue at a time
- repeated fix-verify-reassess cycles are more useful than one large pass
- the user wants polish, stabilization, or gradual improvement

Treat the task as blocking-first review handling when:

- the input is review feedback, QA findings, or self-review findings
- the goal is to fix only material issues without stalling delivery for low-value polish

## Selection Rules

- Choose a phase-driven workflow when the task is broad and end-to-end.
- Choose a loop-driven workflow when the task is bounded and benefits from repeated fix-verify-reassess cycles.
- Choose a blocking-first review workflow when the input is a set of findings and only material items should block progress.
- If the work starts as bounded refinement but expands into several coordinated areas, escalate to a broader workflow.
- If the work starts broad but enters a short local polish cycle, temporarily use a bounded loop without losing the broader scope frame.

## Escalation Signals

- The current issue can no longer be solved as one bounded unit.
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

- Do not choose a broad workflow for a narrow bounded issue just because the task sounds important.
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
- "이거 끝까지 미는 게 맞아, 아니면 루프로 다듬는 게 맞아?"
- "리뷰 반영인지, 구현 workflow인지 먼저 구분해줘"
