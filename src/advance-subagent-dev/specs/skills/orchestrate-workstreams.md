## 사용자 스펙 의도

- 목표와 성공 기준부터 최종 응답까지 하나의 lifecycle로 subagent workstream을 조율합니다.
- spawn은 독립 workstream 두 개 이상, 개별 계약, 병렬 이점, shared-state 안전, 메인 에이전트의 검증·통합 가능성을 모두 충족할 때만 허용합니다.
- 기본은 read-only이며, 병렬 write는 고정된 공유 계약과 완전히 분리된 ownership에서만 허용합니다.
- Terra xhigh worker를 `EXPLORE_READ`, `IMPLEMENT_OWNED`, `REVIEW_LENS`, `PROCESS_STRUCTURED` task packet으로 구분합니다.
- Sol xhigh는 제한된 `FRONTIER_JUDGMENT`에서만 사용하고 Luna route를 두지 않습니다.
- subagent 결과를 evidence로 취급하고, lifecycle당 narrow follow-up을 한 번만 허용합니다.
- software-engineering-only와 cross-domain workstream을 같은 gate와 lifecycle로 처리합니다.

---

# orchestrate-workstreams 스킬 스펙

## 목적

`orchestrate-workstreams`는 software-engineering 또는 cross-domain investigation-and-action 요청을 검증 가능한 workstream으로 분해하고, 필요한 경우에만 bounded subagent를 생성해 메인 에이전트가 결과를 안전하게 검증·통합하도록 안내합니다.

## 경계

- 포함:
  - 목표·성공 기준, `DIRECT`·`DISPATCH`, workstream·dependency graph
  - shared-state 안전, task packet, model·role routing, 실제 spawn
  - terminal result normalization, evidence·ownership 검사, conflict resolution
  - whole-result verification과 final response
- 제외:
  - complexity, file count, 장시간 작업만을 근거로 한 dispatch
  - 순수 조사 보고서의 출처 정책·인용 감사·evidence ledger 설계
  - nested subagent, 범위 확장, 읽기 권한의 암묵적 write 승격
  - 메인 에이전트 책임의 subagent 이전

## 처리하려는 작업 형태

- 조사와 prototype, 문서, 코드, 데이터 transformation 같은 action lane이 함께 있는 작업
- 독립 모듈·서비스·실행 경로·테스트 묶음의 software-engineering 조사와 구현
- 보안·정확성·성능·품질 같은 독립 review lens
- independent source retrieval, schema-bound processing, disjoint implementation, review lens를 결합하는 작업
- 각 lane의 산출물과 근거를 공통 성공 기준으로 검증·통합하는 작업
- explicit subagent 요청 중 dispatch gate를 모두 충족하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/orchestrate-workstreams/SKILL.md`
- 호출 방식: `$advance-subagent-dev:orchestrate-workstreams` 또는 narrow implicit trigger
- implicit policy: `allow_implicit_invocation: true`
- 자동 포함: software-engineering 또는 cross-domain 요청에서 의미 있고 독립적인 workstream 두 개 이상이 명백한 경우
- 자동 제외: simple lookup, single-source summary, pure evidence-report research, sequential root cause, complexity·file count·일반 engineering 표현만 있는 요청

## 핵심 처리 계약

다음 data flow를 순서대로 유지합니다.

`goal and success criteria → DIRECT or DISPATCH → workstream/dependency graph → shared-state safety → bounded task packets → model/role routing → spawn → terminal result normalization → evidence and ownership checks → conflict resolution → whole-result verification → final response`

- 메인 에이전트는 goal, success criteria, shared contracts, dispatch state, main-owned surface, integration, conflict resolution, final verification, final response를 소유합니다.
- 각 단계의 산출물을 다음 단계 입력으로 사용하고, 누락된 계약을 추정해 spawn하지 않습니다.
- `DIRECT`이면 orchestration 형식을 사용자 응답에 강제하지 않고 메인 에이전트가 기존 도구·skill 계약으로 처리합니다.
- `DISPATCH`이면 agent ID, task packet, ownership, required result, follow-up 상태, whole-result verification을 현재 context에 유지합니다.

## Dispatch Gate

다음을 모두 충족할 때만 spawn합니다.

- 의미 있는 workstream이 두 개 이상이며 독립적으로 시작할 수 있습니다.
- 각 workstream에 서로 구분되는 scope, deliverable, evidence, completion criteria가 있습니다.
- 병렬 이점이 coordination과 재검사 비용보다 큽니다.
- file, data, runtime, external-state conflict를 통제할 수 있습니다.
- 메인 에이전트가 결과를 독립적으로 검증하고 통합할 수 있습니다.

하나라도 실패하면 `DIRECT`입니다. Complexity, 많은 파일, “깊게 작업” 요청은 근거가 아닙니다. explicit subagent 요청은 gate에 진입하지만 우회하지 않습니다.

## Workstream·Dependency Graph

- 각 node에 objective, prerequisites, output, consumer, evidence path를 둡니다.
- unfinished result를 prerequisite로 요구하는 node는 동시에 시작하지 않습니다.
- shared contract와 integration point는 메인 에이전트가 먼저 고정합니다.
- 독립 lane 수보다 agent 수를 늘리지 않습니다.
- Terra xhigh의 비용을 고려해 적은 agent, 큰 coherent batch, strict schema, deterministic check, bounded retry, explicit stop condition을 사용합니다.

## Shared-State Safety

- 기본 execution mode는 read-only parallel입니다.
- parallel write는 shared contract가 고정되고 writable ownership이 완전히 분리된 경우에만 허용합니다.
- 한 파일·artifact에는 writer 한 명만 둡니다.
- shared schema, config, lockfile, generated output, mutable fixture, port, database, build output, temp directory, emulator, external account를 동시에 변경하지 않습니다.
- 안전 분리가 불명확하면 read-only 또는 `DIRECT`로 낮춥니다.
- read result를 write authority로 암묵 승격하지 않습니다. 새 write scope는 메인 에이전트가 다시 gate에 적용합니다.

## Task Packet·Spawn

모든 task packet은 다음을 포함합니다.

- role과 model/effort
- objective
- included·excluded scope
- access mode
- owned files/artifacts 또는 `none`
- inputs와 source of truth
- constraints: no sub-subagents, no scope expansion, ownership·shared-state 제한
- deliverable과 result schema
- evidence와 validation
- completion과 stop conditions

packet이 완전하고 session capacity가 확인된 뒤 최소 agent만 spawn합니다. 반환된 agent ID와 exact packet을 기록하며 spawn 성공을 추정하지 않습니다.

## Model·Role Routing

기본 worker는 `gpt-5.6-terra`, `xhigh`입니다.

- `EXPLORE_READ`: research, discovery, logs, source·dependency mapping
- `IMPLEMENT_OWNED`: disjoint owned surface 안의 implementation·action
- `REVIEW_LENS`: correctness, security, performance, quality, counterevidence review
- `PROCESS_STRUCTURED`: schema-bound extraction, transformation, classification, test generation, repetitive mechanical work

`gpt-5.6-sol`, `xhigh`의 `FRONTIER_JUDGMENT`는 다음 중 하나 이상일 때만 허용합니다.

- ambiguous goal framing이 필요합니다.
- shared contract 또는 architecture를 설계해야 합니다.
- 강한 evidence가 충돌합니다.
- error cost가 높고 deterministic verification이 없습니다.
- 독립 final frontier-level audit가 정당화됩니다.

작업이 길다는 이유로 Sol을 사용하지 않습니다. Luna route는 두지 않습니다. 명시 선택 모델을 사용할 수 없으면 가능한 모델로 role contract를 보존하거나 `DIRECT`로 수행하고 제한을 공개합니다.

Sol worker도 dispatch gate의 예외가 아닙니다. 이미 gate를 통과한 lifecycle에서 strong-evidence conflict 같은 조건부 dependent audit node로 graph에 포함하고, prerequisite가 완료된 뒤에만 spawn할 수 있습니다. 독립 workstream이 하나뿐인 새 lifecycle에서 Sol audit 한 명만 spawn하지 않고 `DIRECT`로 판단합니다.

## Terminal Result·Integration

모든 결과를 다음 필드로 정규화합니다.

- `status`: `completed`, `blocked`, `inconclusive`
- `summary and claims`
- `evidence`
- `inspected and changed surfaces`
- `validation performed, skipped, or failed`
- `risks and uncertainty`
- `recommended integration action`

- 결과를 final truth가 아닌 evidence로 취급합니다.
- task packet, ownership, evidence, validation을 대조하고 unsupported claim과 ownership 위반을 거부합니다.
- 중복 claim은 병합하고 conflict는 메인 에이전트가 source, code, test, log, direct execution으로 판정합니다.
- critical assigned scope가 누락된 경우 lifecycle 전체에서 기존 agent에 narrow follow-up 한 번만 허용합니다.
- follow-up 뒤에도 확인되지 않은 범위는 직접 검증하거나 residual risk로 공개합니다.
- accepted result를 통합한 뒤 공통 성공 기준과 whole-result verification을 실행합니다.
- 최종 응답은 통합 결론, 변경·검사 surface, 검증 결과, skipped·failed check, limitation, residual risk를 포함합니다.

## Completion·Stop Conditions

- `completed`: packet의 completion criteria와 evidence·validation 계약을 충족합니다.
- `blocked`: 외부 의존성, 도구, 권한, 입력 부족으로 완료할 수 없습니다.
- `inconclusive`: 허용 범위를 조사했지만 결론을 지지할 근거가 부족합니다.
- agent는 scope 완료, 명시된 한계 도달, stop condition 발생 시 종료합니다.
- bounded retry와 follow-up 한도를 넘기지 않습니다.
- 전체 성공을 검증하지 못하면 완료로 보고하지 않습니다.

## 검토 질문

- goal과 success criteria가 검증 가능한가?
- gate의 모든 조건에 근거가 있는가?
- graph에서 독립 node와 dependency가 구분되는가?
- shared contract, ownership, mutable state가 통제되는가?
- 모든 packet이 role·model·scope·evidence·stop contract를 갖는가?
- Sol route가 허용 조건을 충족하고 Luna route가 없는가?
- 모든 terminal result와 ownership을 메인 에이전트가 확인했는가?
- conflict와 whole-result verification이 해결되거나 risk로 공개됐는가?
- pure evidence-report research의 자동 trigger 경계를 침범하지 않는가?

## 독립성 원칙

- 이 skill은 sibling skill과 dev-only spec 없이 독립 실행 가능해야 합니다.
- 설치된 delegation 도구와 범용 조사·구현 도구만 사용하며 별도 MCP, app, hook, script를 요구하지 않습니다.
- pure research 전문 계약이 필요해도 sibling skill을 숨은 prerequisite로 취급하지 않습니다. 현재 요청이 이 skill의 경계 안에 있으면 설치된 도구와 사용자 지침으로 `DIRECT` 또는 제한된 명시 dispatch를 수행하고 한계를 공개합니다.

## 확장 원칙

- cohesive lifecycle을 유지하고 세부 계약은 one-level runtime reference로만 분리합니다.
- role, result field, execution mode는 메인 에이전트의 routing·검증 판단을 실제로 바꿀 때만 추가합니다.
- broad trigger보다 boundary example과 negative trigger를 우선 보강합니다.
