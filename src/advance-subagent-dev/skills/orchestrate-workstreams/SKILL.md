---
name: orchestrate-workstreams
description: >
  Orchestrate bounded subagents across software-engineering, cross-domain, or
  mixed investigation-and-action work. Use for explicit subagent requests or
  goals with at least two meaningful, independently startable workstreams such
  as separate modules, execution paths, review lenses, research,
  implementation, transformation, or verification. Dispatch only with distinct
  contracts, parallel benefit, controlled shared state, and main-agent
  verification. Do not auto-trigger for simple lookups, single-source
  summaries, pure evidence-report research, sequential root causes, or
  complexity and file count alone.
---

# Orchestrate Workstreams

Coordinate delegation, verification, and integration across independent software-engineering, investigation, or action workstreams. The main agent owns shared contracts, dispatch state, integration, conflict resolution, whole-result verification, and the final response.

The explicit invocation is `$advance-subagent-dev:orchestrate-workstreams`.

Maintain this data flow:

`goal and success criteria → DIRECT or DISPATCH → workstream/dependency graph → shared-state safety → bounded task packets → model/role routing → spawn → terminal result normalization → evidence and ownership checks → conflict resolution → whole-result verification → final response`

## 1. Lock the Goal and Route

Fix the verifiable goal, included and excluded scope, success criteria, permissions, and source of truth. Ask only for information that could materially change the route; state other assumptions.

Choose `DISPATCH` only when every condition holds:

- at least two meaningful workstreams can start independently;
- each workstream has distinct scope, deliverable, evidence, and completion criteria;
- parallel benefit exceeds coordination and recheck cost;
- file, data, runtime, and external-state conflicts are controlled;
- the main agent can independently verify and integrate every result.

If any condition fails, choose `DIRECT`. An explicit subagent request enters the gate but does not bypass it. Software-engineering work may qualify for implicit dispatch when independent lanes are evident; implementation, review, testing, debugging, complexity, or file count alone are insufficient.

Do not automatically own pure evidence-report research. Read [Boundary Examples](references/boundary-examples.md) when selection is ambiguous.

## 2. Design the Graph and Shared State

Assign every node an objective, prerequisites, output, consumer, and evidence path. Do not start a node whose prerequisite is unfinished. The main agent must fix shared contracts and integration points first.

Default to read-only parallel work. Permit parallel writes only when the shared contract is fixed and writable ownership is fully disjoint. Never let concurrent workers mutate the same schema, config, lockfile, generated output, mutable fixture, port, database, build output, temporary directory, emulator, or external account. Downgrade to read-only or `DIRECT` when separation is unclear.

Read [Delegation Safety](references/delegation-safety.md) when the gate, graph, worker count, or shared-state decision is nontrivial.

## 3. Define Packets, Route, and Spawn

Before spawning, read [Task and Result Contracts](references/task-result-contracts.md) and [Model and Role Routing](references/model-routing.md).

- Use the minimum number of workers and group Terra xhigh work into large coherent batches.
- Put role, model and effort, objective, included and excluded scope, access mode, ownership, source of truth, no sub-subagents, no scope expansion, deliverable, result schema, evidence, validation, completion conditions, and stop conditions in every packet.
- Record read-only ownership as `none`; assign one writer per file or artifact.
- Record the returned agent ID and exact packet in dispatch state. Never infer successful spawn without tool evidence.
- If the selected model is unavailable, preserve the role contract with an available model or use `DIRECT`, then disclose the limitation.

## 4. Verify Terminal Results

Wait until every required worker reaches a terminal state. Normalize each result as `completed`, `blocked`, or `inconclusive`. Treat results as evidence, not final truth.

- Compare the original packet with claims, evidence, inspected and changed surfaces, validation, risks, and recommended action.
- Reject unsupported claims and ownership violations.
- Do not convert a read finding into write authority. Apply the gate again for new write scope.
- Allow one narrow follow-up to an existing worker only when critical assigned scope is missing and `follow_up_used` is `false`; set it to `true` immediately.
- Verify unresolved scope directly or disclose it as residual risk.

## 5. Integrate and Finish

Merge duplicate claims. Resolve conflicts using source material, artifacts, code, tests, logs, or direct execution. When no deterministic check exists and `FRONTIER_JUDGMENT` qualifies, consider a Sol worker only as a dependent audit node in a lifecycle that already passed the gate. Do not create a new dispatch for one Sol auditor.

Integrate accepted results and run the shared success criteria and whole-result verification. Disclose skipped or failed checks and limitations. Do not report success without evidence of whole-result completion.

Report the outcome, changed or inspected surfaces, verification, unresolved scope, and residual risk. Do not impose orchestration reporting on `DIRECT` work.
