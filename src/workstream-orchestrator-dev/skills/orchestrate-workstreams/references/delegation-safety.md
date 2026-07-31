# Delegation Safety

## Dispatch Gate

다음 질문에 모두 `yes`일 때만 spawn하세요.

1. 독립적으로 시작 가능한 의미 있는 workstream이 두 개 이상입니까?
2. 각 workstream에 distinct scope, deliverable, evidence, completion criteria가 있습니까?
3. 병렬 이점이 coordination·merge·re-check 비용보다 큽니까?
4. file, artifact, data, runtime, external-state conflict를 통제할 수 있습니까?
5. 메인 에이전트가 각 결과를 독립적으로 검증하고 하나의 결과로 통합할 수 있습니까?

Explicit subagent 요청도 같은 질문을 통과해야 합니다. 하나라도 `no` 또는 `unknown`이면 `DIRECT`입니다.

## Workstream Graph

각 node를 다음 필드로 기록하세요.

```yaml
id: <stable id>
objective: <독립 질문 또는 산출물>
prerequisites: <완료돼야 하는 node 또는 none>
output: <bounded deliverable>
consumer: <main 또는 downstream node>
evidence_path: <검증 방법>
access_mode: read-only | write-enabled
ownership: none | <exact files/artifacts>
```

`prerequisites`가 완료되지 않은 node를 spawn하지 마세요. Shared contract, schema, interface, acceptance criteria는 메인 에이전트가 고정하세요.

## Shared-State Allowlist

- 기본: `read-only`
- `write-enabled`: shared contract가 고정되고 writable surface가 완전히 분리된 경우
- writer: file·artifact당 한 명
- main-owned: shared contract, integration file, final decision, whole-result verification

다음 shared mutable surface는 동시에 변경하거나 사용하는 task로 나누지 마세요.

- schema, interface, config, lockfile, generated output
- mutable fixture, database, port, emulator, external account
- build output, shared temp directory, live document 또는 공용 data

분리할 수 없으면 read-only 조사로 낮추거나 `DIRECT`로 처리하세요.

## Agent Count와 Cost

- 독립 lane 수와 session capacity를 넘지 마세요.
- 보통 2개 agent로 시작하고, 별도 계약을 가진 추가 lane이 있을 때만 늘리세요.
- Terra xhigh의 기계적 volume은 agent를 늘리는 대신 큰 coherent batch, strict schema, deterministic check로 처리하세요.
- Retry는 packet의 bounded stop condition 안에서만 허용하세요.
