---
name: teammate-kit-guide
description: Entrypoint skill for the `teammate-kit` plugin. Use when a task involves teammate-style collaboration and Codex should first decide whether the job needs event-bus orchestration or a single bounded teammate role such as research, implementation, or review.
---

# Teammate Kit Guide

## Overview

Use this skill as the default entrypoint for `teammate-kit`.
Its job is to classify the collaboration shape before work starts and route to the narrowest teammate skill that fits.
Do not default to full orchestration when one bounded teammate role is enough.

## Workflow

1. Identify the collaboration shape:
   - event-bus orchestration across several teammates
   - one bounded research handoff
   - one bounded implementation slice
   - one independent review pass
2. Decide whether the task needs:
   - durable coordination across several workers
   - one teammate owning one bounded job
3. Route to the narrowest teammate skill that owns the work.
4. If the task spans several teammate roles, define the order before execution starts.

## Routing Rules

- Choose `teammate-orchestrator` when the work needs main-agent relay, append-only event history, checkpoints, retries, or long-running coordination between several teammates.
- Choose `teammate-researcher` when the main need is evidence gathering, codebase investigation, option comparison, or a decision-ready handoff without implementation.
- Choose `teammate-implementer` when one teammate should make a bounded code or content change and report concrete verification.
- Choose `teammate-reviewer` when one teammate should do an independent pass on correctness, regressions, scope creep, or missing tests.
- If the task needs several of these roles, define the handoff sequence explicitly instead of blending them into one prompt.

## Decision Rules

- Treat orchestration as the top-level mode only when durability and teammate-to-teammate routing are part of the problem.
- Treat a single teammate role as the top-level mode when one bounded worker can finish the job directly.
- Prefer direct role routing over orchestration when the extra runtime would add complexity without reducing risk.

## Output Contract

- `Task shape`
- `Chosen teammate mode`
- `Why this mode fits`
- `Planned handoff order`
- `Main risk`
- `Residual risk`

## Guardrails

- Do not start with `teammate-orchestrator` unless the task actually needs multi-teammate runtime coordination.
- Do not hide several teammate roles inside one vague assignment.
- Do not route a bounded research, implementation, or review task through orchestration just because several people are mentioned.
