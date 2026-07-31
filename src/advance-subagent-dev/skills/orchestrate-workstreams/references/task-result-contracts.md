# Task And Result Contracts

## Dispatch State

현재 context에 다음 상태를 유지하세요.

```yaml
route: DISPATCH
graph: <workstream/dependency nodes>
assignments:
  - agent_id: <spawn tool result>
    task_packet: <complete packet>
ownership:
  agent_owned: <per-agent files/artifacts or none>
  main_owned: <shared contracts and integration surfaces>
required_results: [<agent ids>]
follow_up_used: false
whole_result_verification: [<checks>]
```

Agent ID를 추정하지 말고 incomplete state로 integration을 시작하지 마세요.

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
deliverable: <bounded integrable output>
result_schema: <required Result Envelope fields>
evidence: <paths, URLs, symbols, commands, key outputs>
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
inspected_surfaces: [<files, artifacts, sources, systems>]
changed_surfaces: [<files/artifacts and reasons, or none>]
validation:
  performed: [<check and result>]
  skipped: [<check and reason>]
  failed: [<check and failure>]
risks_and_uncertainty: [<unknowns or none>]
recommended_integration_action: <main-agent action>
```

## 상태 판정

- `completed`: completion criteria와 evidence·validation 계약을 충족합니다.
- `blocked`: 외부 의존성, 입력, 도구, 권한이 완료를 막습니다.
- `inconclusive`: 허용 범위를 조사했지만 결론을 지지할 근거가 부족합니다.

Missing validation은 자동 성공이 아닙니다. Ownership 밖의 변경은 accepted result가 아니라 위반입니다.
