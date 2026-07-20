---
name: dispatch-subagents
description: >
  Validate and dispatch bounded subagents for a software-engineering request
  or RouteDecision. Use when explicitly invoked or routed from an
  orchestration entrypoint to choose DIRECT, PARALLEL_READ, or PARALLEL_WRITE,
  create complete task packets, and spawn only the minimum necessary agents.
---

# Dispatch Subagents

## Owned Job

Own the full delegation gate, execution mode, workstream boundaries, agent count, task packets, actual spawn, and `DispatchManifest`. Do not wait for results, integrate evidence, choose the final implementation direction, or produce the final response.

Accept either the raw user request or a `RouteDecision`. A direct call may bypass the entry skill, but never bypasses the gate. Re-evaluate incomplete or stale handoffs against the current request and state.

## Delegation Gate

Spawn only when every condition is true:

- At least two meaningful workstreams are independent.
- Each workstream has a clear scope, deliverable, and evidence source.
- Each can start without another unfinished result.
- The main agent can compare, verify, and integrate the outputs.
- Parallel benefit exceeds coordination and re-check cost.
- File, runtime, data, and shared-state conflicts are controlled.

If the user forbids subagents or any condition fails, return `DIRECT` without spawning. Do not delegate because the task is complex or touches many files.

Read [delegation-rubric.md](references/delegation-rubric.md) when independence, mode, agent count, or shared-state safety is not obvious.

## Execution Mode

- Use `DIRECT` for one workstream, small or sequential work, a shared root cause, or excessive coordination cost.
- Use `PARALLEL_READ` for independent investigation or review, uncertain ownership, overlapping edits, or work needed before a shared contract is fixed.
- Use `PARALLEL_WRITE` only when writable file sets are fully disjoint, shared contracts are fixed, and no shared config, schema, lockfile, generated file, or mutable state is edited concurrently.

When uncertain, use `PARALLEL_READ`.

Do not run tests concurrently when they share a database, port, temporary directory, build output, emulator, mutable fixture, or external account.

## Agent Selection

Use the minimum meaningful count: normally 2, typically 2-3, and up to 4 within session capacity for clearly separate large work. Never exceed the number of independent lanes or available slots.

Prefer:

- `explorer` for read-only exploration, review, logs, tests, causes, security, or performance.
- `worker` for isolated implementation with explicit file ownership and focused validation.
- `default` only when neither role fits.

Do not let subagents spawn subagents, expand scope, commit, push, open PRs, release, or edit outside ownership.

## Task Packets And Spawn

Read [task-contract.md](references/task-contract.md) before spawning. Give every agent Objective, Scope, Access mode, Ownership, Inputs, Constraints, Deliverable, Evidence, and Completion criteria. Use `Ownership: none` for read-only work and one writer per writable file.

Spawn the bounded agents. Record every returned agent ID and exact packet. Do not claim dispatch succeeded without tool evidence.

## DispatchManifest

After successful spawn, hand the main agent a complete manifest containing:

- `mode`
- `assignments`: agent ID, role, and task packet
- `ownership`: agent-owned and main-owned writable surfaces
- `required_results`: agent IDs required before integration
- `main_owned_work`: shared contracts, files, and decisions
- `follow_up_used: false`: lifecycle-wide recovery state
- `whole_result_verification`

Do not load the integrator if the manifest is incomplete. Otherwise tell the main agent to apply `$adaptive-subagent-orchestrator:integrate-subagent-results` with the manifest. Use [examples.md](references/examples.md) only for boundary-case comparison or validation.
