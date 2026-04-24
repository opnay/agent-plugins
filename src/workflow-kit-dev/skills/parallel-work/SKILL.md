---
name: parallel-work
description: Parallel execution workflow for a small set of clearly independent tasks. Use when several bounded work items can proceed at the same time and the main challenge is safe split, coordination, and integration rather than long multi-phase delivery.
---

# Parallel Work

## Overview

Use this skill when the task should move through a few parallel lanes instead of one sequential queue.
This is a user-facing workflow mode for bounded independent work, not an internal execution engine.
Its job is to decide what can safely run in parallel, launch only the lanes that are truly independent, and then integrate the results with one shared verification pass.

## Use When

- the work contains two or more bounded tasks with clear independence
- sequential execution would mostly waste time rather than reduce risk
- each task has a clear owner, output, and verification path
- the user explicitly wants parallel execution or faster turnaround on independent work

## Do Not Use When

- the work depends on a shared prerequisite or ordered decisions
- several tasks touch the same files, responsibilities, or fragile boundary
- the task needs a full end-to-end delivery pipeline with phase gates
- the main input is review findings that should be triaged by blocking severity first

## Scope Boundary

- Use this skill for a small number of independent lanes, usually 2 to 4.
- Treat lane independence as the main decision, not the desire for speed by itself.
- Stop using this skill when lane coordination becomes the dominant problem.

## Core Policy

- Split first, parallelize second.
- Default to `not independent yet` until the boundary is explicit.
- Do not run two lanes in parallel when they are likely to touch the same file set or the same decision surface.
- Keep each lane narrowly scoped enough that completion can be verified quickly.
- After lane work completes, run one integration pass before declaring success.
- If one lane blocks, do not pretend the whole parallel batch succeeded.

## Workflow

### Phase 0: Split The Work

1. List the candidate tasks.
2. Group them into lanes with one primary objective each.
3. Record why each lane is independent.
4. Record any shared files, shared dependencies, or shared decisions.

Output:

- Candidate lanes
- Independence claim
- Shared risk map

### Phase 1: Decide Safe Parallelism

1. Mark each lane as:
   - parallel-safe
   - sequential-after-prerequisite
   - not worth parallelizing
2. Launch only the lanes that are clearly parallel-safe.
3. Keep uncertain lanes out of the parallel batch.

Output:

- Parallel lanes
- Deferred lanes
- Why blocked lanes were not parallelized

### Phase 2: Execute The Lanes

1. Run each parallel-safe lane with its own bounded objective.
2. Keep lane scope stable while execution is in flight.
3. If new overlap appears, stop widening the lane and move the unresolved part to integration.

Output:

- Lane result per task

### Phase 3: Collect And Integrate

1. Gather lane outputs.
2. Check for overlap, drift, or missing glue work.
3. Apply only the minimum integration changes needed to make the lanes work together.

Output:

- Integration summary
- Remaining conflicts or merges

### Phase 4: Verify The Whole Result

1. Re-run lane-local checks when needed.
2. Run one shared verification pass for the integrated result.
3. Report complete, partial, and blocked lanes separately.

Output:

- Verification result
- Completed lanes
- Blocked lanes

## Independence Rules

Treat lanes as independent only when all of these are reasonably true:

- the tasks do not require the same unanswered design decision
- the tasks do not depend on each other's output to start
- the tasks are unlikely to edit the same files or tightly coupled module boundary
- each lane can be described and verified on its own

If any of those points is weak, prefer sequential execution.

## Verification Rules

- Verify each lane proportionally to its size.
- Always verify the integrated result after collection, even if all lanes passed individually.
- If lane success and integrated success differ, report the integration failure as the real result.
- Keep the verification path explicit before launching the parallel batch.

## Stop Conditions

- All parallel-safe lanes completed and the integrated result passed verification.
- The remaining lanes are no longer independent enough to justify this mode.
- One or more lanes reveal a shared dependency that requires sequential planning.
- Integration risk becomes larger than the time saved by parallel execution.
- The user says stop, pause, or switch modes.

When stopping, report:

1. Which lanes completed
2. Which lanes were deferred or blocked
3. Whether the next step should switch to another workflow mode

## Output Contract

- `Parallel lanes`
- `Why they are independent`
- `Shared risk`
- `Verification plan`
- `Completed lanes`
- `Integration result`
- `Stop reason`
- `Residual risk`

## Guardrails

- Do not use parallel execution as a default speed hack.
- Do not treat lane separation as valid unless file overlap and decision overlap were checked.
- Do not let one lane silently absorb shared integration work that changes its scope.
- Do not skip the integrated verification pass just because lane-local checks passed.

## Example Triggers

- "서로 독립적인 작업 몇 개를 병렬로 처리해줘"
- "이건 순서보다 병렬이 맞는지 보고 진행해줘"
- "작은 작업 여러 개를 동시에 돌리고 마지막에 합쳐줘"
