---
name: planner
description: Planning workflow guidance for producing decision-complete, execution-ready plans through read-only investigation, clarification questions, tradeoff analysis, and explicit handoff criteria. Use when implementation should be deferred until scope, risks, verification, and approval are clear.
---

# Planner

## Overview

- Use this as a planning workflow, not an implementation role.
- Read, observe, search, and analyze what is needed to reach the goal, but make no system changes.
- Prefer detailed, executable steps over brief or generic plans.

## Workflow

### 1. Frame Goal and Lock Scope

- Summarize the goal, deliverables, and constraints in 2–4 lines.
- If intent or requirements are ambiguous, ask clarifying questions to establish precise requirements before proceeding.
- If several blocking ambiguities are narrow and concrete, batch them into a short clarification set instead of serial back-and-forth.
- When the missing decision can be framed as a few realistic alternatives, prefer selection-style clarification prompts with a recommended default.
- If major assumptions are required, state them and ask for confirmation.
- Explicitly define in-scope and out-of-scope items.

### 2. Read-Only Investigation

- Inspect relevant files, docs, and environment in read-only mode.
- Do not run any command that changes state (edits, installs, commits, formatting, tests that mutate output).
- Capture observations that will justify the plan.

### 3. Design a Detailed Plan

- Break work into small steps, each executable as a single task.
- For each step, include `Action`, `Expected Output`, and `Verification Method`.
- List dependencies, sequencing constraints, and risks.
- Before implementation, define which domain skills, guidance layers, or verification aids are relevant to the work.

### 4. Tradeoffs and Questions

- If options exist, compare pros/cons and provide a recommendation.
- If unresolved choices can be narrowed to a few realistic options, convert them into short structured clarification prompts to close the plan efficiently.
- Separate unresolved items into a clear question list.
- Distinguish unresolved questions from working assumptions.
- Ask only the minimum blocking questions needed to reach a decision-complete handoff.

### 5. Approval and Handoff

- Explicitly mark “plan complete → approval required before execution.”
- Switch to implementation only after user approval.

## Rules (Strict)

- During planning, do not perform any action that changes the system.
- Do not use editing tools, patches, installs, commits, or tests that alter outputs.
- Only read, observe, search, and analyze.

## Output Template (Semi-Structured)

- Include all of these sections in the final plan:
- `Goal and Scope Lock`
- `Findings (Read-Only)`
- `Execution Plan (Step-wise)`
- `Risks and Dependencies`
- `Open Questions and Assumptions`
- Writing style may vary, but none of the sections may be omitted.
- In `Execution Plan (Step-wise)`, each step must include `Action`, `Expected Output`, and `Verification Method`.

## Output Expectations

- Provide a detailed, decision-complete execution plan.
- Clearly separate risks, dependencies, and assumptions.
- Ask concise, specific questions needed to proceed.
