## 사용자 스펙 의도

- adaptive subagent 플러그인의 스킬을 `경량 진입 → 위임·생성 → 결과 검증·통합`으로 세분화합니다.
  - 위임·생성 단계는 전체 gate, 계획, 실제 subagent 생성을 함께 소유합니다.
  - 후속 스킬은 implicit 호출하지 않으며 입력 게이트를 갖춘 직접 호출을 지원합니다.

---

# dispatch-subagents 스킬 스펙

## 목적

`dispatch-subagents`는 raw request 또는 `RouteDecision`을 평가하여 전체 delegation gate를 통과할 때만 최소 bounded subagent를 실제 생성합니다.

## 경계

- 포함:
  - 전체 delegation gate
  - DIRECT, PARALLEL_READ, PARALLEL_WRITE 선택
  - workstream, agent 수, 역할, ownership 확정
  - task packet 작성과 실제 spawn
  - `DispatchManifest` 생성
- 제외:
  - implicit invocation
  - 결과 대기와 evidence 통합
  - 최종 구현 방향, 코드 통합, 전체 검증, 최종 응답
  - nested subagent 생성

## 처리하려는 작업 형태

- 독립 모듈, 서비스, 실행 경로 조사
- 독립 관점의 cross-cutting review
- 테스트 묶음별 실패 분석
- migration 영향, 플랫폼, 기술 대안 비교
- 공통 계약과 disjoint ownership이 확정된 제한적 병렬 구현

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/dispatch-subagents/SKILL.md`
- 상세 판단: `references/delegation-rubric.md`
- task packet: `references/task-contract.md`
- 실행 예시: `references/examples.md`
- 호출 방식: `$adaptive-subagent-orchestrator-dev:dispatch-subagents` 또는 유효한 `RouteDecision` handoff
- implicit policy: `allow_implicit_invocation: false`

## 입력 게이트

- raw user request 또는 사용자 목표·제약을 보존한 `RouteDecision`을 요구합니다.
- direct call은 entry를 우회할 수 있지만 delegation gate를 우회하지 않습니다.
- 불완전하거나 stale한 handoff는 raw request와 현재 상태로 다시 평가합니다.
- 사용자 금지 조건이 있으면 agent를 생성하지 않고 `DIRECT`를 반환합니다.

## 위임 게이트

Subagent를 생성하려면 모두 참이어야 합니다.

- 의미 있는 독립 workstream이 두 개 이상입니다.
- 각 workstream에 명확한 scope와 deliverable이 있습니다.
- 각 workstream이 다른 미완성 결과 없이 시작할 수 있습니다.
- main agent가 결과를 비교, 검증, 통합할 수 있습니다.
- 병렬 이점이 coordination cost보다 큽니다.
- 파일, runtime, data, shared state 충돌을 통제할 수 있습니다.

하나라도 거짓이면 `DIRECT`이며 spawn하지 않습니다.

## 실행 모드

- `DIRECT`: 하나의 workstream, 작은 범위, 강한 순차성, 공통 원인, 높은 coordination cost입니다.
- `PARALLEL_READ`: 독립 조사·검토가 가능하거나 write overlap과 shared-state 위험이 불확실합니다.
- `PARALLEL_WRITE`: 파일 집합이 완전히 분리되고 공통 계약이 확정됐으며 shared mutable surface가 없습니다.
- 불확실하면 `PARALLEL_READ`를 사용합니다.

## Agent 수와 역할

- 필요한 최소 수만 사용합니다.
- 작은 병렬 작업은 2개, 일반 복합 작업은 2~3개, 명확한 대형 작업은 session capacity 안에서 최대 4개를 기본으로 합니다.
- 사용자가 더 많은 수를 명시해도 session capacity와 의미 있는 lane 수를 넘지 않습니다.
- explorer는 read-only 조사·리뷰·원인 분석, worker는 disjoint implementation과 제한 검증에 사용합니다.
- 모든 packet에 nested subagent 금지를 명시합니다.

## Task Packet 계약

각 packet은 다음을 포함합니다.

- Objective
- Scope와 out-of-scope
- Access mode
- Ownership
- Inputs
- Constraints와 shared-state limits
- Deliverable
- Evidence requirements
- Completion criteria

## DispatchManifest 계약

spawn 후 현재 대화 안의 구조화된 manifest에 다음을 기록합니다.

- `mode`
- `assignments`: agent ID, role, task packet
- `ownership`: writable surface와 main-owned shared surface
- `required_results`: integration 전에 필요한 agent ID
- `main_owned_work`: 공통 계약, shared file, 최종 통합 책임
- `whole_result_verification`: 통합 후 실행할 검증
- `follow_up_used`: lifecycle follow-up 사용 여부, 초기값 `false`

manifest가 완전하지 않으면 integrator로 handoff하지 않습니다.

## 병렬 안전 규칙

- 한 파일에는 writer 한 명만 배정합니다.
- shared config, schema, lockfile, generated output, mutable fixture는 parallel write하지 않습니다.
- DB, port, temp directory, build output, emulator, external account를 공유하는 테스트는 동시에 실행하지 않습니다.
- 공통 인터페이스가 확정되기 전에는 cross-layer parallel write하지 않습니다.
- subagent는 commit, push, PR, release를 수행하지 않습니다.

## 검토 질문

- 모든 gate 조건에 evidence가 있는가?
- packet 없이 시작해야 하는 agent가 있는가?
- ownership과 shared state가 완전히 분리됐는가?
- agent 수가 의미 있는 lane 수를 초과하지 않는가?
- DispatchManifest가 integrator 입력으로 완전한가?

## 독립성 원칙

- 이 skill은 raw request만으로 전체 gate와 dispatch를 수행할 수 있어야 합니다.
- `RouteDecision`은 선택적이고 명시적인 handoff이며 hidden prerequisite가 아닙니다.

## 확장 원칙

- 새로운 mode를 추가하기보다 기존 세 mode의 allowlist를 먼저 명확히 합니다.
- detailed rubric, packet template, 사례는 runtime references가 소유합니다.
- 결과 처리 규칙을 dispatcher에 추가하지 않습니다.
