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

Coordinate independent software-engineering, investigation, or action workstreams. The main agent owns the goal, shared contracts, dispatch state, integration, conflict resolution, whole-result verification, and final response.

The explicit invocation is `$advance-subagent:orchestrate-workstreams`.

Follow this flow:

`goal and success criteria → DIRECT or DISPATCH → dependency graph → shared-state safety → task packets → profile and role selection → spawn → terminal results → evidence and ownership checks → integration → whole-result verification → final response`

## 1. Fix the Goal and Route

Establish the verifiable goal, included and excluded scope, success criteria, permissions, and source of truth. Ask only about information that can materially change the route; disclose other assumptions.

Choose `DISPATCH` only when all conditions hold:

- at least two meaningful workstreams can start independently;
- each has distinct scope, deliverable, evidence, and completion criteria;
- parallel benefit exceeds coordination and recheck cost;
- conflicts over files, data, runtime, and external state are controlled;
- the main agent can independently verify and integrate every result.

Otherwise use `DIRECT`. Explicit requests enter this gate without bypassing it. Complexity, file count, implementation, review, testing, debugging, and duration alone do not qualify. Do not automatically own pure evidence-report research. Read [Boundary Examples](references/boundary-examples.md) when selection is ambiguous.

For `DIRECT`, use the applicable task workflow without imposing orchestration reporting.

## 2. Define Dependencies and Safe Ownership

Give every node an objective, prerequisites, output, consumer, and evidence path. Fix shared contracts and integration points before dispatch. Wait for prerequisites before starting dependent work.

Default to read-only parallel work. Allow parallel writes only with a fixed shared contract and completely disjoint ownership. Assign one writer per artifact. Never concurrently mutate the same schema, config, lockfile, generated output, mutable fixture, port, database, build output, temporary directory, emulator, or external account. Use read-only work or `DIRECT` when separation is unclear.

Read [Delegation Safety](references/delegation-safety.md) when the gate, graph, worker count, or ownership requires more analysis.

## 3. Select Profiles and Dispatch Packets

Read [Task and Result Contracts](references/task-result-contracts.md) and [Model Profiles and Role Routing](references/model-routing.md) before spawning.

Choose only `fast`, `default`, or `high`. Applicable instructions define the exact model-and-effort pairs for `fast` and `high`. For `default`, omit both spawn settings to inherit. Follow explicit task and repository profile choices first; otherwise use the role's demands, uncertainty, verification, and cost. Profiles do not change roles, permissions, or the dispatch gate.

Treat each explicit pair as indivisible. Do not mix profiles, tune effort separately, or introduce another combination. Resolve combined shorthand against supported model IDs and efforts; do not pass the shorthand as a model ID. Check the chosen pair against the current tool. If undefined or unsupported, use another permitted profile, inheritance, or `DIRECT` only when the role can still be met. Do not substitute a mandatory profile without authority. Report when no valid route exists.

- Use the minimum workers and coherent batches; do not add workers for mechanical volume.
- Include role, selected profile and reason, explicit model/effort pair or inheritance, objective, included and excluded scope, access, ownership, source of truth, no nested delegation, no scope expansion, deliverable, result schema, evidence, validation, completion, and stop conditions in each packet.
- Record read-only ownership as `none`. Confirm capacity before spawn.
- Record the returned agent ID and exact packet in dispatch state; do not infer spawn success.
- Record undisclosed inherited values as inherited, not guessed IDs. Disclose substitutions and omitted work.

## 4. Verify Required Results

Wait for every required worker to reach a terminal state: `completed`, `blocked`, or `inconclusive`. Treat results as evidence rather than final truth.

Compare the assigned packet with claims, inspected and changed surfaces, ownership, evidence, checks, risks, and recommended integration. Reject unsupported claims and ownership violations. Read findings do not grant write authority; apply the gate again to new write scope.

Allow one narrow follow-up across the lifecycle only when critical assigned scope is missing and `follow_up_used` is false. Set it true immediately. Verify remaining scope directly or disclose it as residual risk.

## 5. Integrate and Finish

Merge duplicate claims. Resolve conflicts using original sources, artifacts, code, tests, logs, or direct execution. When deterministic verification is insufficient and `FRONTIER_JUDGMENT` is justified, a separate judgment audit may be a dependent node in a lifecycle that already passed the gate. Select its profile under the same policy. Do not create a new dispatch for one auditor.

Integrate accepted results and check common success criteria and the whole result. Report outcome, changed or inspected surfaces, checks, unresolved scope, and residual risk. Disclose skipped or failed checks and do not claim completion without whole-result evidence.
