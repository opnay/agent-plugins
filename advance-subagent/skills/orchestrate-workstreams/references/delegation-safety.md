# Delegation Safety

## Dispatch Gate

Spawn only when every answer is `yes`:

1. Are there at least two meaningful workstreams that can start independently?
2. Does each workstream have distinct scope, deliverable, evidence, and completion criteria?
3. Does the parallel benefit exceed coordination, merge, and recheck cost?
4. Can file, artifact, data, runtime, and external-state conflicts be controlled?
5. Can the main agent independently verify each result and integrate one outcome?

Explicit subagent requests use the same gate. Any `no` or `unknown` means `DIRECT`.

## Workstream Graph

Record each node as:

```yaml
id: <stable id>
objective: <independent question or deliverable>
prerequisites: <required completed nodes or none>
output: <bounded deliverable>
consumer: <main agent or downstream node>
evidence_path: <verification method>
access_mode: read-only | write-enabled
ownership: none | <exact files or artifacts>
```

Do not spawn a node before its prerequisites complete. The main agent must fix shared contracts, schemas, interfaces, and acceptance criteria.

## Shared-State Allowlist

- Default: `read-only`
- `write-enabled`: only when the shared contract is fixed and writable surfaces are fully disjoint
- Writer count: one per file or artifact
- Main-owned: shared contracts, integration files, final decisions, and whole-result verification

Do not split work into concurrent tasks that mutate or consume the same:

- schema, interface, config, lockfile, or generated output;
- mutable fixture, database, port, emulator, or external account;
- build output, shared temporary directory, live document, or shared data.

If separation is impossible, downgrade to read-only investigation or use `DIRECT`.

## Worker Count and Cost

- Do not exceed the number of independent lanes or available session capacity.
- Start with two workers; add another only for a lane with a distinct contract.
- Handle mechanical Terra xhigh volume through large coherent batches, strict schemas, and deterministic checks instead of more workers.
- Retry only within the packet's bounded stop condition.
