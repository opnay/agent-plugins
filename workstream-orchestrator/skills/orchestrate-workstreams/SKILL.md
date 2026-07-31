---
name: orchestrate-workstreams
description: >
  Orchestrate bounded subagents across cross-domain or mixed
  investigation-and-action work. Use for explicit subagent requests or goals
  with at least two meaningful, independently startable workstreams such as
  research plus implementation, transformation, review, or verification.
  Dispatch only with distinct contracts, parallel benefit, controlled shared
  state, and main-agent verification. Do not auto-trigger for simple lookups,
  single-source summaries, pure evidence-report research, ordinary
  software-engineering-only requests, sequential root causes, or complexity
  alone.
---

# Orchestrate Workstreams

조사와 실행이 섞인 독립 workstream의 위임·검증·통합을 하나의 lifecycle로 수행하세요. 메인 에이전트가 공유 계약, dispatch state, 통합, 충돌 해결, 전체 검증, 최종 응답을 소유하세요.

명시 호출 식별자는 `$workstream-orchestrator:orchestrate-workstreams`입니다.

다음 data flow를 유지하세요.

`goal and success criteria → DIRECT or DISPATCH → workstream/dependency graph → shared-state safety → bounded task packets → model/role routing → spawn → terminal result normalization → evidence and ownership checks → conflict resolution → whole-result verification → final response`

## 1. 목표와 Route 고정

검증 가능한 최종 goal, 포함·제외 범위, success criteria, 권한, source of truth를 먼저 고정하세요. 결론을 실질적으로 바꾸는 정보만 질문하고, 나머지는 가정을 공개하세요.

다음을 모두 충족할 때만 `DISPATCH`하세요.

- 의미 있는 workstream이 두 개 이상이며 지금 독립적으로 시작할 수 있습니다.
- 각 workstream에 distinct scope, deliverable, evidence, completion criteria가 있습니다.
- parallel benefit가 coordination·re-check cost보다 큽니다.
- file, data, runtime, external-state conflict를 통제할 수 있습니다.
- 메인 에이전트가 결과를 독립적으로 검증·통합할 수 있습니다.

하나라도 실패하면 `DIRECT`로 실행하세요. Explicit subagent 요청은 gate에 진입하지만 우회하지 않습니다. Complexity, 많은 파일, 긴 작업, “깊게 작업” 요청만으로 dispatch하지 마세요.

Trigger 또는 인접 plugin 경계가 불명확하면 [boundary-examples.md](references/boundary-examples.md)를 읽으세요.

## 2. Graph와 Shared State 설계

각 workstream node에 objective, prerequisites, output, consumer, evidence path를 지정하세요. 미완성 결과를 prerequisite로 요구하는 node는 동시에 시작하지 마세요. 공유 계약과 integration point는 메인 에이전트가 먼저 고정하세요.

기본은 read-only parallel입니다. Parallel write는 공유 계약이 고정되고 writable ownership이 완전히 분리된 경우에만 허용하세요. Shared schema, config, lockfile, generated output, mutable fixture, port, database, build output, temp directory, emulator, external account를 동시에 변경하지 마세요. 불명확하면 read-only 또는 `DIRECT`로 낮추세요.

Gate, graph, agent 수, shared-state 판단이 단순하지 않으면 [delegation-safety.md](references/delegation-safety.md)를 읽으세요.

## 3. Packet, Routing, Spawn

Spawn 전에 [task-result-contracts.md](references/task-result-contracts.md)와 [model-routing.md](references/model-routing.md)를 읽으세요.

- 최소 agent만 사용하고 Terra xhigh 작업은 큰 coherent batch로 묶으세요.
- 모든 packet에 role·model/effort, objective, included·excluded scope, access mode, ownership, source of truth, no sub-subagents, no scope expansion, deliverable, result schema, evidence·validation, completion·stop conditions를 넣으세요.
- Read-only ownership은 `none`으로 기록하세요. 한 파일·artifact에는 writer 한 명만 두세요.
- 반환된 agent ID와 exact packet을 dispatch state에 기록하세요. Tool evidence 없이 spawn 성공을 추정하지 마세요.
- 명시 모델을 사용할 수 없으면 role contract를 보존할 수 있는 모델로 대체하거나 `DIRECT`로 실행하고 제한을 공개하세요.

## 4. Terminal Result 검증

모든 required agent가 terminal state가 될 때까지 기다리고 결과를 `completed`, `blocked`, `inconclusive`로 정규화하세요. 결과는 final truth가 아니라 evidence입니다.

- Original packet과 summary·claims·evidence·inspected/changed surfaces·validation·risks·recommended action을 대조하세요.
- Unsupported claim과 ownership 위반을 거부하세요.
- Read finding을 write authority로 승격하지 마세요. 새 write scope는 다시 gate에 적용하세요.
- Critical assigned scope가 누락됐고 `follow_up_used`가 `false`일 때만 기존 agent에 narrow follow-up 한 번을 허용하고 즉시 `true`로 바꾸세요.
- 확인하지 못한 범위는 직접 검증하거나 residual risk로 공개하세요.

## 5. 통합과 완료

중복 claim은 병합하세요. 충돌은 메인 에이전트가 source, artifact, code, test, log, direct execution으로 해결하세요. Deterministic check가 없고 `FRONTIER_JUDGMENT` 조건을 충족하면, gate를 이미 통과한 lifecycle의 dependent audit node에서만 Sol worker를 고려하세요. Sol 한 명만을 위한 새 dispatch를 만들지 마세요.

Accepted result를 메인 에이전트가 통합하고 공통 success criteria와 whole-result verification을 실행하세요. Skipped·failed check와 limitation을 숨기지 마세요. 전체 완료를 입증하지 못하면 성공으로 보고하지 마세요.

최종 응답에는 결과, 변경·검사 surface, 검증, 미확인 범위, residual risk를 포함하세요. `DIRECT` 작업에는 orchestration 보고 형식을 강제하지 마세요.
