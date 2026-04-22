# Teammate Kit 플러그인 스펙

## 플러그인 목적

`teammate-kit`은 teammate-style collaboration을 구조화하는 플러그인입니다.
핵심 책임은 durable orchestration이 필요한 협업인지, 아니면 하나의 bounded teammate role만으로 충분한지 먼저 분류하고, 그에 맞는 role surface를 제공하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - collaboration mode 분류
  - event-bus 기반 durable orchestration
  - bounded research, implementation, review role
- 제외:
  - generic execution workflow selection
  - skill/plugin/subagent 설계 자체
  - domain-specific implementation standard

## 처리하려는 작업 형태

- 여러 teammate role을 조합해야 하는 협업 작업
- research, implementation, review 중 하나의 역할만 명확히 맡기면 되는 작업
- durable history, checkpoint, retry가 필요한 협업 실행

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `teammate-kit-guide`
- 대표 스펙: `teammate-kit/specs/plugin-spec.md`
- skill 상세 스펙 위치: `teammate-kit/specs/skills/*.md`

## 내장 skill 체계

- `teammate-kit-guide`: 협업 작업을 orchestration 또는 direct role로 라우팅한다.
  - spec: `teammate-kit/specs/skills/teammate-kit-guide-spec.md`
- `teammate-orchestrator`: durable event-bus 기반 multi-teammate workflow를 실행한다.
  - spec: `teammate-kit/specs/skills/teammate-orchestrator-spec.md`
- `teammate-researcher`: evidence 수집과 decision-ready handoff를 만든다.
  - spec: `teammate-kit/specs/skills/teammate-researcher-spec.md`
- `teammate-implementer`: bounded change를 소유하고 구현과 검증을 마친다.
  - spec: `teammate-kit/specs/skills/teammate-implementer-spec.md`
- `teammate-reviewer`: independent risk-focused review를 수행한다.
  - spec: `teammate-kit/specs/skills/teammate-reviewer-spec.md`

## SDD 운영 원칙

- plugin spec은 collaboration mode와 role map만 소유한다.
- orchestration과 bounded role skill의 처리 계약은 각 skill spec으로 분리한다.
- collaboration mode boundary가 바뀌면 guide와 plugin spec을 함께 갱신한다.
- generic workflow guidance는 `workflow-kit`에 남기고 이 플러그인 안으로 흡수하지 않는다.

## 현재 구조 메모

- 이 플러그인의 주요 리스크는 direct role이면 충분한 상황에도 orchestration을 과도하게 쓰는 것이다.
