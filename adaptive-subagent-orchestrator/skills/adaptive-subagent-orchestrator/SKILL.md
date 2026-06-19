---
name: adaptive-subagent-orchestrator
description: >
  Automatically orchestrate parallel subagents for complex software-engineering
  tasks with two or more independent workstreams: multi-module exploration,
  cross-cutting review, debugging across components, independent test failure
  analysis, migration impact analysis, and technical comparison. Invoke
  implicitly when a task can be safely split into independent investigation,
  review, testing, comparison, or implementation workstreams, even when the
  user does not mention subagents. When activated, evaluate and spawn the
  minimum necessary bounded subagents if delegation is worthwhile; otherwise
  work directly. Do not use for trivial, tightly sequential, single-scope, or
  overlapping-write tasks. automatic delegation, parallel subagents,
  independent workstreams, multi-module exploration, cross-cutting review,
  자동 위임, 병렬 서브에이전트, 독립 작업, 다중 모듈 분석
---

# Adaptive Subagent Orchestrator

## Core Rule
When this skill is activated, explicitly evaluate the task for subagent delegation. If the delegation gate passes, spawn the minimum necessary bounded subagents, wait for all required results, validate their evidence, and integrate the final outcome. Do not merely recommend using subagents.

If the gate fails, handle the task directly. If the user explicitly says not to use subagents, use DIRECT.

## Owned Job
Coordinate bounded subagent use for complex software-engineering work with independent workstreams. Keep final interpretation, integration, verification, and user response with the main agent.

Non-goals: trivial edits, single-file questions, tightly sequential work, overlapping-write delegation, nested subagent creation, custom agent configuration, model or reasoning-effort hardcoding, and external settings changes.

## First Decision
Separate skill activation from subagent spawning:

1. Skill activation: the request matched explicit subagent wording or an implicit multi-workstream trigger.
2. Subagent spawning: the delegation gate passed.

Skill activation alone never requires spawning.

## Delegation Gate
Spawn subagents only when all conditions are true:

- At least two meaningful workstreams can proceed independently.
- Each workstream has a clear scope and deliverable.
- Each workstream can start without another workstream's unfinished result.
- The main agent can compare, verify, or integrate results.
- Parallel execution improves time, exploration quality, or context separation more than it adds coordination cost.
- File, runtime, data, and shared-state conflicts can be controlled.

Do not delegate because there are many files. Delegate only when independent reasoning or verification is needed.

## Execution Modes
Use DIRECT when there is one workstream, the task is small or local, steps are strongly sequential, split/integration cost is larger than the work, or all work depends on one shared cause or state.

Use PARALLEL_READ when independent investigations or reviews can start safely, final edited files might overlap, implementation needs evidence from several areas, or parallel writing is risky. When uncertain, choose PARALLEL_READ over PARALLEL_WRITE.

Use PARALLEL_WRITE only when all conditions are true:

- Each workstream owns a fully disjoint file set.
- Shared interfaces and change contracts are already fixed.
- One unfinished change does not block another workstream.
- No two agents edit the same file.
- No shared config, schema, lockfile, generated file, or mutable shared state is edited concurrently.
- The main agent can review final diffs and run integrated verification.

If any condition is uncertain, use PARALLEL_READ or DIRECT.

## Do Not Spawn For
- documentation typos
- obvious one- or two-line fixes
- small refactors inside one function
- local rename work
- one known-cause test fix
- questions answerable from one file
- strongly sequential tasks
- concurrent edits to the same file
- one shared config or shared state causing every issue
- tests sharing a database, port, temp directory, build output, emulator, mutable fixture, or external account
- cases where integration cost exceeds direct execution
- explicit user instruction to use one agent

If the skill was implicit and DIRECT is chosen, do not over-explain that choice unless asked.

## Agent Count And Roles
Use the minimum necessary agents:

- small parallel work: 2
- typical composite work: 2-3
- large work with clear independent areas: up to 4

Do not exceed 4 unless the user explicitly requests more and the session allows it. Do not split work just to fill slots. Do not duplicate the same assignment unless independent review perspectives are the goal.

Prefer built-in roles:

- `explorer`: codebase exploration, path tracing, impact analysis, logs, test-failure cause analysis, code review, option research, read-only security or performance review
- `worker`: owned implementation, isolated tests, independent fixes, reproducible bug fixes, bounded validation
- `default`: only when explorer/worker does not fit

Tell every subagent not to spawn more subagents.

## Main Agent Responsibilities
Do not delegate final user-request interpretation, delegation decision, workstream decomposition, agent count/role selection, overlap prevention, file ownership, shared interface decisions, conflict adjudication, final implementation direction, code integration, full verification, or final response.

Treat subagent conclusions as evidence. Verify important claims against code, tests, logs, or direct command output.

## Subagent Instructions
Every subagent prompt must include:

- Objective
- Scope, including out-of-scope items
- Access mode: `read-only` or `write-enabled`
- Ownership: `none` for read-only; explicit writable file set for write-enabled
- Inputs
- Constraints, including no nested subagents and no scope expansion
- Deliverable
- Evidence requirements
- Completion criteria

Use [task-contract.md](references/task-contract.md) before spawning subagents or requiring the return template.

## Parallel Read
Parallelize module exploration, service incident analysis, test-bundle failure analysis, perspective-based review, platform compatibility checks, option comparison, log-slice analysis, and migration impact analysis when scopes or review lenses are distinct.

If implementation is needed after PARALLEL_READ, the main agent integrates findings, fixes the shared contract, then either edits directly or assigns one writer. Use PARALLEL_WRITE only after ownership is disjoint.

## Parallel Write Safety
Never allow two agents to edit the same file. Do not parallelize writes to shared config, common schemas, generated files, lockfiles, or shared mutable fixtures. Do not split frontend/API/data implementation until the common interface is fixed.

Subagents must not commit or push. The main agent reviews final diffs and runs integrated verification.

## Result Collection
Wait for every required subagent result. Then separate completed, blocked, and inconclusive results; merge duplicates; compare evidence per claim; identify conflicts; resolve conflicts with code, tests, or direct execution; reject unsupported claims; send at most one narrow follow-up only when a critical scope is missing; integrate final changes; run whole-result verification.

If one subagent fails, do not restart everything. Confirm the missing part directly, retry once with a narrower scope, or state the unverified scope.

## Reference Routing
- Read [delegation-rubric.md](references/delegation-rubric.md) when mode choice, agent count, shared-state risk, or a boundary case is not obvious.
- Read [task-contract.md](references/task-contract.md) before spawning subagents or reviewing their returned structure.
- Read [examples.md](references/examples.md) when checking implicit-trigger behavior or validating a proposed split.

## Final Response
If subagents were used, summarize number/roles, integrated conclusion, changed files, validation result, and residual risk. Do not paste raw subagent transcripts. If no subagents were used, report the work normally.
