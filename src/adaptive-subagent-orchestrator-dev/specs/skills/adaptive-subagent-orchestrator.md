# adaptive-subagent-orchestrator 스킬 스펙

## 사용자 스펙 의도

- 사용자는 `adaptive-subagent-orchestrator` skill을 원합니다.
- 이 skill은 사용자가 서브에이전트나 병렬 처리를 직접 언급하지 않아도, 복잡한 software-engineering 요청에서 독립 작업 흐름이 두 개 이상이고 병렬 위임 이점이 있으면 최소한의 bounded subagent를 생성하고 조율해야 합니다.
- 명시적 호출과 암묵적 호출을 모두 지원해야 합니다.
- skill 활성화 자체를 "서브에이전트 위임을 명시적으로 평가하고, 게이트가 통과하면 실제로 생성하라"는 지시로 취급해야 합니다.
- 위임 조건이 충족되지 않거나 사용자가 단일 에이전트 처리를 명시하면 subagent를 생성하지 않아야 합니다.
- DIRECT, PARALLEL_READ, PARALLEL_WRITE 모드를 별도로 판단해야 합니다.
- subagent task instruction과 반환 형식은 Objective, Scope, Access mode, Ownership, Inputs, Constraints, Deliverable, Evidence, Completion criteria를 포함해야 합니다.
- subagent는 다시 subagent를 생성하면 안 됩니다.
- main agent는 최종 해석, 분해, 파일 소유권, 통합, 전체 검증, 최종 응답 책임을 유지해야 합니다.
- instruction-only skill로 만들고 scripts/assets를 만들지 않아야 합니다.

---

## 목적

`adaptive-subagent-orchestrator`는 복잡한 software-engineering task를 평가하여 독립 workstream이 두 개 이상이고 병렬 위임 이점이 조정 비용보다 클 때만 최소한의 bounded subagent를 생성하도록 조율합니다.

## 경계

- 포함:
  - explicit/implicit invocation 대응
  - subagent delegation gate
  - DIRECT, PARALLEL_READ, PARALLEL_WRITE mode selection
  - 최소 agent 수 결정
  - read-only exploration과 safe parallel write 구분
  - subagent task contract와 return contract
  - evidence validation과 final integration
  - final response shape
- 제외:
  - trivial, single-scope, tightly sequential task delegation
  - 동일 파일 병렬 수정
  - 공통 설정, 잠금 파일, 생성물 동시 수정
  - nested subagent spawning
  - custom agent config, model, reasoning effort hardcoding
  - external settings mutation

## 처리하려는 작업 형태

- 다중 모듈 탐색, 서비스별 장애 조사, 실행 경로별 원인 분석
- 보안, 정확성, 테스트, 성능, 동시성 관점의 cross-cutting review
- 독립 테스트 묶음별 실패 조사와 최소 수정
- 플랫폼, 런타임, 설계 대안 비교
- migration 영향 범위 분석
- 로그, 코드 영역, 실행 경로 분리 조사
- disjoint write scope가 확정된 제한적 병렬 구현

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/adaptive-subagent-orchestrator/SKILL.md`
- 상세 판단표: `references/delegation-rubric.md`
- task contract: `references/task-contract.md`
- examples: `references/examples.md`
- 호출 방식: 직접 `$adaptive-subagent-orchestrator` 호출, subagent/parallel/delegate 표현, 또는 description 기반 implicit invocation

## 핵심 처리 계약

- skill이 활성화되면 subagent 위임 가능성을 명시적으로 평가합니다.
- delegation gate가 통과하면 추천만 하지 않고 최소한의 bounded subagent를 실제로 생성하도록 지시합니다.
- gate가 통과하지 않으면 DIRECT로 처리합니다.
- 사용자가 subagent를 쓰지 말라고 하면 DIRECT로 처리합니다.
- main agent는 모든 필수 subagent 결과를 기다리고, evidence를 검증하고, 충돌을 직접 판정하고, 최종 통합과 전체 검증을 수행합니다.

## 스킬 발견 판단

- description 앞부분은 automatic delegation, parallel subagents, independent workstreams, multi-module exploration, cross-cutting review를 포함합니다.
- 한국어 발견을 위해 자동 위임, 병렬 서브에이전트, 독립 작업, 다중 모듈 분석을 포함합니다.
- trivial, tightly sequential, single-scope, overlapping-write task에는 호출되지 않도록 제외 조건을 description에 둡니다.

## 위임 게이트

Subagent를 생성하려면 모두 참이어야 합니다.

- 독립적으로 진행 가능한 의미 있는 workstream이 두 개 이상입니다.
- 각 workstream에 명확한 scope와 deliverable을 줄 수 있습니다.
- 각 workstream이 다른 미완성 결과 없이 시작될 수 있습니다.
- main agent가 결과를 비교, 검증, 통합할 수 있습니다.
- 병렬 실행 이점이 조정 비용보다 큽니다.
- 파일, 실행 환경, 데이터, 공유 상태 충돌을 통제할 수 있습니다.

## 실행 모드

- DIRECT: 독립 workstream이 하나뿐이거나, 작고 국소적이거나, 강하게 순차적이거나, 분할 비용이 더 큽니다.
- PARALLEL_READ: 독립 조사/검토가 두 개 이상이고, 구현 전 여러 증거가 필요하거나, 수정 파일 overlap 가능성이 있습니다. 불확실하면 이 모드를 선택합니다.
- PARALLEL_WRITE: 수정 대상 파일 집합이 완전히 분리되고, 공통 인터페이스와 변경 계약이 확정됐고, 공유 설정/스키마/잠금 파일/생성물을 동시에 수정하지 않을 때만 사용합니다.

## Agent 수와 역할

- 필요한 최소 수만 생성합니다.
- 기본은 2개, 일반 복합 작업은 2~3개, 독립 영역이 명확한 큰 작업은 최대 4개입니다.
- 사용자가 더 많은 수를 명시하지 않으면 4개를 넘지 않습니다.
- explorer는 조사, 리뷰, 로그/테스트 원인 분석에 우선 사용합니다.
- worker는 disjoint write scope가 명확한 구현이나 제한된 검증에 사용합니다.
- default는 explorer/worker로 분류하기 어려울 때만 사용합니다.

## Main Agent 책임

- 사용자 요구 최종 해석, delegation 여부, workstream 분해, agent 수와 역할, 파일 소유권, 공통 계약, 충돌 판정, 최종 구현 방향, 코드 통합, 전체 검증, 최종 응답은 main agent가 책임집니다.
- subagent 결론은 evidence로 취급하고, 중요한 결론은 코드, 테스트, 로그, 직접 실행 결과로 검증합니다.

## Task Contract

- 모든 subagent 지시는 Objective, Scope, Access mode, Ownership, Inputs, Constraints, Deliverable, Evidence, Completion criteria를 포함합니다.
- read-only 작업은 Ownership을 `none`으로 둡니다.
- write-enabled 작업은 수정 가능한 파일 집합을 명시하고 한 파일을 한 writer만 소유합니다.
- 모든 subagent에게 추가 subagent를 생성하지 말라고 지시합니다.

## 병렬 안전 규칙

- 동일 파일을 두 agent가 수정하지 않습니다.
- 공통 인터페이스가 확정되기 전에는 병렬 구현하지 않습니다.
- 공유 설정, 공통 스키마, 생성 파일, 잠금 파일은 동시에 수정하지 않습니다.
- 테스트가 DB, 포트, 임시 디렉터리, 빌드 산출물, 에뮬레이터, mutable fixture, 외부 계정을 공유하면 병렬 실행하지 않습니다.

## 결과 수집과 응답

- 모든 필수 결과를 기다립니다.
- completed, blocked, inconclusive를 구분합니다.
- 중복을 합치고, 주장별 evidence를 비교하고, 충돌은 직접 검증으로 판정합니다.
- 후속 지시는 중요한 범위 누락 때만 제한적으로 한 차례 사용합니다.
- subagent를 사용한 최종 응답은 agent 수/역할, 통합 결론, 변경 파일, 검증, 남은 위험 순서를 기본으로 합니다.
- subagent를 사용하지 않은 경우 DIRECT 판단을 장황하게 설명하지 않습니다.

## 검토 질문

- 의미 있는 독립 workstream이 둘 이상인가?
- 각 workstream은 미완성 결과 없이 시작 가능한가?
- 수정 파일과 공유 상태 충돌을 통제할 수 있는가?
- 병렬 이점이 조정 비용보다 큰가?
- main agent가 결과를 검증하고 통합할 수 있는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 이유: 이 skill은 단일 skill로 설치되어도 delegation 판단, task packet, result integration을 수행할 수 있어야 합니다. plugin-level 사용 안내 없이도 runtime references만으로 동작해야 합니다.

## 확장 원칙

- 새 예시는 `references/examples.md`에 추가합니다.
- 새 판단 기준은 먼저 skill spec의 현재 계약과 충돌하지 않는지 확인한 뒤 `SKILL.md` 또는 `references/delegation-rubric.md`로 배치합니다.
- runtime에는 dev-only spec 경로를 실행 지시로 남기지 않습니다.
