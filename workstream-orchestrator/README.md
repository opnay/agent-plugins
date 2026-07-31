# Workstream Orchestrator

독립적으로 시작할 수 있는 조사·실행 작업 흐름을 bounded subagent에 위임하고, 메인 에이전트가 근거와 변경을 검증·통합하도록 안내하는 instruction-only 플러그인입니다.

## 목적

`workstream-orchestrator`는 검색, browsing, 자료 회수, 코드·문서 조사, 구현, 변환, 검토, 검증이 한 목표 안에 섞인 cross-domain 또는 mixed investigation-and-action 작업을 다룹니다. 복잡성이나 작업량이 아니라 실제 독립성과 병렬 이점으로 위임을 결정합니다.

하나의 `orchestrate-workstreams` skill이 다음 수명주기를 소유합니다.

`목표·성공 기준 → DIRECT 또는 DISPATCH → workstream·의존성 graph → shared-state 안전 → task packet → 모델·역할 routing → spawn → 결과 정규화 → 근거·ownership 검사 → 충돌 해결 → 전체 검증 → 최종 응답`

메인 에이전트는 공유 계약, dispatch 상태, 통합, 충돌 해결, 전체 검증, 최종 응답을 계속 소유합니다.

## 적합한 상황

- 시장 조사와 프로토타입 구현처럼 조사와 실행 lane이 독립적으로 시작될 때
- 문서 변환, 데이터 추출, 구현, 품질 검토를 서로 분리된 산출물로 병렬화할 때
- 별도 코드·문서·데이터 표면을 소유하는 실행 lane이 공통 계약을 공유할 때
- 사용자가 subagent orchestration을 명시하고 위임 gate를 충족할 때

## 자동 적용하지 않는 상황

- 단순 조회, 단일 출처 요약, 하나의 순차적 원인 분석
- 여러 출처의 근거와 인용을 추적하는 순수 조사 보고서
- 명시적 orchestration 요청이 없는 일반 software-engineering-only 작업
- 복잡성, 많은 파일, 긴 실행 시간만 있는 작업

순수 조사 보고서의 source methodology는 `deep-research`의 경계이며, 일반 software-engineering-only 정밀 routing과 3단 dispatch lifecycle은 `adaptive-subagent-orchestrator`의 경계입니다. 이 플러그인은 어느 sibling도 runtime 의존성으로 요구하지 않습니다. 명시 호출은 현재 설치 상태와 관계없이 자체 gate와 계약으로 동작합니다.

## 사용 방법

명시 호출:

```text
$workstream-orchestrator:orchestrate-workstreams 시장 조사를 수행하고 결과를 바탕으로 독립적인 프로토타입 두 개를 구현해 주세요.
```

자동 트리거가 적합한 요청:

```text
경쟁 제품 조사, 기존 문서 구조 분석, 데모 구현을 각각 진행한 뒤 하나의 제안으로 검증·통합해 주세요.
```

명시 호출이나 subagent 요청도 위임 gate를 우회하지 않습니다. 두 개 이상의 의미 있는 workstream이 독립적으로 시작할 수 없거나 coordination·재검사 비용이 더 크면 `DIRECT`로 수행합니다.

## 모델·역할

기본 worker는 `gpt-5.6-terra`, `xhigh`입니다.

| 역할 | 용도 |
| --- | --- |
| `EXPLORE_READ` | 조사, discovery, log, source·dependency mapping |
| `IMPLEMENT_OWNED` | 완전히 분리된 owned surface의 구현·실행 |
| `REVIEW_LENS` | correctness, security, performance, quality, counterevidence 검토 |
| `PROCESS_STRUCTURED` | schema-bound 추출, 변환, 분류, test 생성, 기계적 반복 작업 |
| `FRONTIER_JUDGMENT` | 제한 조건을 충족할 때만 `gpt-5.6-sol`, `xhigh`로 수행하는 고난도 판단 |

Terra xhigh의 기계적 volume 비용을 줄이기 위해 agent 수를 최소화하고, 큰 coherent batch, 엄격한 schema, deterministic check, bounded retry, 명시적 stop condition을 사용합니다. Luna route는 제공하지 않습니다.

## 안전과 결과

- 기본은 읽기 전용 병렬 실행입니다.
- 병렬 write는 공유 계약이 고정되고 writable ownership이 완전히 분리된 경우에만 허용합니다.
- shared schema, config, lockfile, generated output, mutable fixture, port, database, external account를 동시에 변경하지 않습니다.
- 모든 결과를 `completed`, `blocked`, `inconclusive`로 정규화하고 근거·ownership·검증 상태를 메인 에이전트가 직접 확인합니다.
- critical scope 누락에는 lifecycle 전체에서 좁은 follow-up 한 번만 허용합니다.
