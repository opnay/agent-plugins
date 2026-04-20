# Teammate Kit 플러그인 스펙

## 목적

`teammate-kit`은 teammate-style collaboration을 구조화하기 위한 플러그인입니다.
핵심 역할은 durable multi-teammate orchestration이 필요한지, 아니면 하나의 bounded teammate role만 있으면 되는지 판단한 뒤 research, implementation, review에 맞는 역할 표면을 제공하는 것입니다.

## 경계

- 포함:
  - teammate workflow 분류
  - event bus 기반 durable orchestration
  - bounded research role
  - bounded implementation role
  - bounded review role
- 제외:
  - generic execution workflow selection
  - skill/plugin/subagent 설계
  - domain-specific implementation standard

## 진입 표면

- 대표 엔트리포인트: `teammate-kit-guide`
- 핵심 분기: orchestration이 필요한지, 아니면 하나의 bounded teammate role이면 되는지 결정한다

## 스킬 구성

- `teammate-kit-guide`: 협업 작업을 orchestration 또는 하나의 direct role로 라우팅한다
- `teammate-orchestrator`: checkpoint와 retry를 갖춘 durable event-bus 중심 teammate workflow를 실행한다
- `teammate-researcher`: evidence를 수집하고 concise한 decision-ready handoff를 만든다
- `teammate-implementer`: 하나의 scoped change를 소유하고 정확한 verification과 residual risk를 보고한다
- `teammate-reviewer`: 하나의 independent risk-focused review pass를 수행한다

## 확장 원칙

- 새 role skill은 research, implementation, review, orchestration과 다른 고유한 teammate responsibility를 가질 때만 추가한다.
- orchestration과 bounded role skill은 분리한다.
- role 설명이 generic workflow guidance로 흐르지 않게 한다.
- collaboration mode boundary가 바뀌면 `teammate-kit-guide`와 이 spec을 함께 갱신한다.

## 현재 의도 점검

- 현재 플러그인 표면은 teammate collaboration mode를 중심으로 일관적이다.
- 현재의 주요 리스크는 하나의 bounded teammate role이면 충분한 상황에도 orchestration을 과도하게 쓰는 것이다.
