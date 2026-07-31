# Task and Result Contracts

## Dispatch State

Maintain:

```yaml
route: DISPATCH
graph: <workstream and dependency nodes>
assignments:
  - agent_id: <spawn tool result>
    task_packet: <complete packet>
ownership:
  agent_owned: <per-agent files or artifacts, or none>
  main_owned: <shared contracts and integration surfaces>
required_results: [<agent ids>]
follow_up_used: false
whole_result_verification: [<checks>]
```

Never infer an agent ID or begin integration with incomplete dispatch state.

## Task Packet

```yaml
role: EXPLORE_READ | IMPLEMENT_OWNED | REVIEW_LENS | PROCESS_STRUCTURED | FRONTIER_JUDGMENT
model: gpt-5.6-terra | gpt-5.6-sol
reasoning_effort: xhigh
objective: <exact question or action>
scope:
  included: <exact surfaces>
  excluded: <explicit non-goals>
access_mode: read-only | write-enabled
owned_files_or_artifacts: none | <exact list>
inputs: <artifacts and required context>
source_of_truth: <authoritative source>
constraints:
  - Do not spawn sub-subagents.
  - Do not expand scope.
  - Do not write outside ownership.
  - Preserve explicit user permissions and shared-state limits.
deliverable: <bounded, integrable output>
result_schema: <required Result Envelope fields>
evidence: <paths, URLs, symbols, commands, and key outputs>
validation: <deterministic or review checks>
completion_conditions: <observable success>
stop_conditions: <blocked, inconclusive, retry, time, or evidence limit>
```

## Result Envelope

```yaml
status: completed | blocked | inconclusive
summary_and_claims:
  - claim: <concise claim>
    basis: confirmed | inference
evidence:
  - <source, path, symbol, command, or summarized output>
inspected_surfaces: [<files, artifacts, sources, or systems>]
changed_surfaces: [<files or artifacts and reasons, or none>]
validation:
  performed: [<check and result>]
  skipped: [<check and reason>]
  failed: [<check and failure>]
risks_and_uncertainty: [<unknowns or none>]
recommended_integration_action: <main-agent action>
```

## Status Rules

- `completed`: Meets completion criteria and the evidence and validation contract.
- `blocked`: An external dependency, input, tool, or permission prevents completion.
- `inconclusive`: The allowed scope was inspected but evidence remains insufficient.

Missing validation is not success. A change outside ownership is a violation, not an accepted result.
