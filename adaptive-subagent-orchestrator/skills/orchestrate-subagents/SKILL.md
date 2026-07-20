---
name: orchestrate-subagents
description: >
  Lightweight entrypoint that routes software-engineering requests to DIRECT
  or bounded subagent dispatch. Use for explicit subagent, delegation, or
  parallel-agent requests, or implicitly only when at least two clearly
  independent workstreams are visible. Do not use for complexity, many files,
  implementation, review, testing, or debugging alone. subagent orchestration,
  parallel subagents, independent workstreams, 자동 위임, 병렬 서브에이전트
---

# Orchestrate Subagents

## Owned Job

Route the request to `DIRECT` or `DISPATCH` at minimal context cost. Do not run the full delegation gate, choose an execution mode, write task packets, spawn agents, or integrate results.

## Decision

1. Use `DIRECT` when the user forbids subagents.
2. Use `DISPATCH` when the user explicitly requests subagents, delegation, or parallel agents. Dispatch still decides whether spawning is worthwhile.
3. For implicit use, choose `DISPATCH` only when at least two meaningful workstreams are clearly independent and can start now.
4. Otherwise use `DIRECT`. Complexity, file count, or broad engineering wording is insufficient.

## Handoff

For `DIRECT`, record `route: DIRECT` with a short reason in the current context, then continue the task without loading another orchestration skill or explaining the routing decision unless asked.

For `DISPATCH`, pass a compact `RouteDecision` containing:

- `route: DISPATCH`
- `trigger_basis`
- `goal`
- `candidate_workstreams`
- `user_constraints`
- `shared_state_flags`

Then use `$adaptive-subagent-orchestrator:dispatch-subagents`. Preserve the original request and explicit constraints; do not pre-decide mode, ownership, role, or agent count.
