---
name: workflow-kit-guide
description: Entrypoint skill for the `workflow-kit` plugin. Use when a task needs the right clarification workflow, execution mode, or final gate chosen first, such as deciding whether work should begin with deep intent alignment, active execution, or a commit-readiness pass.
---

# Workflow Kit Guide

## Overview

Use this skill as the default entrypoint for `workflow-kit`.
Its job is to classify whether the task is primarily about clarifying implementation intent, doing the work, or deciding whether the work is ready to move to commit.
Do not jump straight into execution when the real need is a deep alignment pass, and do not run the commit gate when the work is still clearly in progress.

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
3. Route to the narrowest bundled skill that owns the main concern.
4. If the task spans execution and gate concerns, pick the starting skill and explicit handoff point.

## Routing Rules

- Choose `deep-interview` when the user already wants implementation, but the main risk is misreading intent, boundaries, tradeoffs, or approval lines before execution begins.
- Choose `workflow-guide` when the task is active implementation work and the main question is which execution mode should run first.
- Choose `autopilot` when the user wants broad end-to-end delivery from brief to verified result.
- Choose `parallel-work` when the work contains several bounded lanes with a clear integration path.
- Choose `ralph-loop` when the work is best handled as repeated bounded fix-verify-reassess cycles.
- Choose `review-loop` when the input is review feedback or findings and the job is to fix only the material issues one at a time.
- Choose `commit-readiness-gate` when implementation is largely done and the main question is whether the current change is isolated, verified, risk-classified, and ready to move toward commit.
- If the task clearly begins as execution and ends with a commit-readiness pass, start with the fitting execution workflow and hand off to `commit-readiness-gate` only after the intended change unit is complete.

## Decision Rules

- Choose `deep-interview` when alignment quality is the main blocker and implementation should follow a locked brief.
- Treat `workflow-guide` as the execution-mode selector, not as the whole plugin entrypoint.
- Choose a gate only when the main uncertainty is readiness, not implementation approach.
- Choose an execution workflow when meaningful implementation, refinement, or findings work still remains.
- Do not use `commit-readiness-gate` as a substitute for actual review handling or incomplete implementation.
- Prefer one clear starting skill plus an explicit handoff over mixing clarification, execution, and gate responsibilities into one prompt.

## Output Contract

- `Task shape`
- `Clarification vs execution vs gate decision`
- `Chosen skill`
- `Why this route fits`
- `Planned handoff`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not skip an obvious alignment pass when the user's intent is still materially underspecified.
- Do not send in-progress implementation work straight to the commit gate.
- Do not use execution workflows when the real need is a final readiness decision.
- Do not bury plugin-level routing guidance inside sibling skills.
- Do not blend clarification, execution, and gate responsibilities without naming which one starts and when the handoff happens.
