# Execution Safety

## Per-Task Readiness

Before executing a task, verify:

- a bounded objective or question, output, source of truth, checks, and stop conditions;
- required inputs and verified prerequisites;
- authority and explicit access and ownership;
- controlled file, data, runtime, and external-state conflicts;
- main-agent ability to verify and integrate the result;
- available capacity and a supported instruction-permitted profile.

The overall goal and completion criteria may still be provisional. The local task can establish them. One task is sufficient; readiness is not a two-worker gate.

If a task is not ready, resolve its missing input, schedule an independent safe task, or report the dependency. Do not invent permissions or restore a separate direct-execution mode.

## Dependencies and Parallel Benefit

Run independent tasks in parallel only when their benefit exceeds coordination and recheck cost. Use sequential scheduling inside the same loop for prerequisites, one remaining task, or shared state.

A round plan records objective, prerequisites, inputs, output, consumer, evidence, access mode, ownership, and domain contract for each node. Verify the prerequisite before starting a downstream node. Settle shared interfaces before writes that depend on them.

Use coherent batches for repetitive structured work instead of adding workers for volume alone.

## Shared-State Ownership

Default to read-only. Parallel writes require completely disjoint writable ownership and settled shared contracts. Use one writer per file or artifact.

The main agent owns execution records, shared contracts, integration surfaces, final decisions, and whole-result verification. A planner is read-only and cannot expand permissions.

Avoid concurrent mutation, or consumption of unstable state, across schemas, interfaces, configs, lockfiles, generated output, mutable fixtures, databases, ports, emulators, build output, shared temporary directories, live documents, and external accounts.

Isolate alternative implementations and their mutable runtime resources. If isolation is unavailable, investigate read-only or schedule alternatives sequentially. Read results never promote a task to write access.

## Capacity and Ownership Transfer

Count active agents before spawning, including planners. Use only necessary workers; do not exceed independent ready lanes or available capacity. Plan execution after the planner returns; it does not spawn workers.

Record returned agent IDs. Use supported termination or cleanup operations when needed; do not assume an interrupted agent is gone or a completed agent's slot is reusable without checking the tool.

Collect required terminal results, or explicitly stop outstanding work and confirm mutations have ceased, before a checkpoint or ownership transfer.

## Execution Limits

Preserve user and environment limits across rounds. A mandatory unavailable profile, forbidden delegation, missing tool, denied authority, or exhausted capacity is an execution limitation, not permission to simulate a successful fork.

One lifecycle-wide narrow follow-up may recover critical missing assigned scope. Preserve its used state across flush. A new round must have a genuinely new justified task, not the same failed packet with a different name.
