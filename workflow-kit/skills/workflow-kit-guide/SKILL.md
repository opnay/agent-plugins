---
name: workflow-kit-guide
description: Entrypoint skill for the `workflow-kit` plugin. Use when a task needs the right workflow or gate chosen first, such as deciding between active execution modes and a commit-readiness gate for near-finished changes.
---

# Workflow Kit Guide

## Overview

Use this skill as the default entrypoint for `workflow-kit`.
Its job is to classify whether the task is primarily about doing the work or deciding whether the work is ready to move to commit.
Do not jump straight into an execution workflow when the real need is a commit-readiness decision, and do not run the commit gate when the work is still clearly in progress.

## Workflow

1. Identify the task shape:
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

- Choose `workflow-guide` when the task is active implementation work and the main question is which execution mode should run first.
- Choose `autopilot` when the user wants broad end-to-end delivery from brief to verified result.
- Choose `parallel-work` when the work contains several bounded lanes with a clear integration path.
- Choose `ralph-loop` when the work is best handled as repeated bounded fix-verify-reassess cycles.
- Choose `review-loop` when the input is review feedback or findings and the job is to fix only the material issues one at a time.
- Choose `commit-readiness-gate` when implementation is largely done and the main question is whether the current change is isolated, verified, risk-classified, and ready to move toward commit.
- If the task clearly begins as execution and ends with a commit-readiness pass, start with the fitting execution workflow and hand off to `commit-readiness-gate` only after the intended change unit is complete.

## Decision Rules

- Treat `workflow-guide` as the execution-mode selector, not as the whole plugin entrypoint.
- Choose a gate only when the main uncertainty is readiness, not implementation approach.
- Choose an execution workflow when meaningful implementation, refinement, or findings work still remains.
- Do not use `commit-readiness-gate` as a substitute for actual review handling or incomplete implementation.
- Prefer one clear starting skill plus an explicit handoff over mixing execution and gate responsibilities into one prompt.

## Output Contract

- `Task shape`
- `Execution vs gate decision`
- `Chosen skill`
- `Why this route fits`
- `Planned handoff`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not send in-progress implementation work straight to the commit gate.
- Do not use execution workflows when the real need is a final readiness decision.
- Do not bury plugin-level routing guidance inside sibling skills.
- Do not blend execution and gate responsibilities without naming which one starts and when the handoff happens.
