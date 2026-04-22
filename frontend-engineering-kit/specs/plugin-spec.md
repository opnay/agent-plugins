# Frontend Engineering Kit 플러그인 스펙

## 플러그인 목적

`frontend-engineering-kit`은 프론트엔드 문제를 올바른 decision layer에서 다루도록 돕는 specialist 플러그인입니다.
핵심 책임은 현재 병목이 frontend workflow인지, architecture pattern인지, React structure인지, local component boundary인지, domain modeling인지, UI implementation quality인지, test-driven execution인지 먼저 분류하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - frontend workflow selection과 execution framing
  - architecture pattern choice
  - React layering, hook, effect, rendering boundary
  - local component boundary와 ownership decision
  - frontend domain-modeling threshold
  - UI implementation quality와 design-to-code guidance
  - frontend TDD workflow
- 제외:
  - backend architecture
  - design-only pre-code planning
  - frontend 특화가 아닌 generic workflow orchestration

## 처리하려는 작업 형태

- 구조 개편이나 feature 작업에서 어떤 frontend concern이 먼저 풀려야 하는지 정하는 작업
- architecture pattern과 React/component 구조를 분리해서 판단해야 하는 작업
- UI 품질과 테스트 전략까지 포함한 frontend 구현 개선 작업

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `frontend-engineering-kit-guide`
- 대표 스펙: `frontend-engineering-kit/specs/plugin-spec.md`
- skill 상세 스펙 위치: `frontend-engineering-kit/specs/skills/*.md`

## 내장 skill 체계

- `frontend-engineering-kit-guide`: dominant frontend concern을 분류하고 첫 specialist path를 고른다.
  - spec: `frontend-engineering-kit/specs/skills/frontend-engineering-kit-guide-spec.md`
- `frontend-workflow-guide`: project structure를 읽고 deeper concern의 순서와 실행 흐름을 정한다.
  - spec: `frontend-engineering-kit/specs/skills/frontend-workflow-guide-spec.md`
- `frontend-architecture-patterns`: Atomic, FSD, hybrid 같은 구조 패턴 적합성을 판단한다.
  - spec: `frontend-engineering-kit/specs/skills/frontend-architecture-patterns-spec.md`
- `react-architecture`: React-specific layering, hook, context, effect, rerender 문제를 다룬다.
  - spec: `frontend-engineering-kit/specs/skills/react-architecture-spec.md`
- `component-architecture`: local component boundary, extraction, ownership, refactor direction을 다룬다.
  - spec: `frontend-engineering-kit/specs/skills/component-architecture-spec.md`
- `frontend-domain-modeling`: 프론트엔드에 필요한 domain model의 깊이와 위치를 정한다.
  - spec: `frontend-engineering-kit/specs/skills/frontend-domain-modeling-spec.md`
- `frontend-design-guide`: design intent를 구현 가능한 UI guidance로 전환한다.
  - spec: `frontend-engineering-kit/specs/skills/frontend-design-guide-spec.md`
- `frontend-tdd-rgb`: frontend change를 failing test와 Red-Green-Refactor 루프로 이끈다.
  - spec: `frontend-engineering-kit/specs/skills/frontend-tdd-rgb-spec.md`

## SDD 운영 원칙

- plugin spec은 concern map과 routing surface만 소유한다.
- 각 skill의 처리 계약은 별도 `specs/skills/` 문서로 분리한다.
- concern boundary가 바뀌면 `frontend-engineering-kit-guide`와 관련 skill spec, `plugin-spec.md`를 같은 변경에서 갱신한다.
- React, component, domain-modeling 경계는 서로 대체 관계가 아니라 다른 decision layer임을 유지한다.

## 현재 구조 메모

- 이 플러그인의 주요 리스크는 React architecture, component architecture, domain modeling 경계가 흐려지면서 같은 문제를 여러 skill이 동시에 소유하게 되는 것이다.
