# Frontend Engineering Kit 플러그인 스펙

## 목적

`frontend-engineering-kit`은 프론트엔드 작업을 올바른 decision layer에서 다루도록 안내하는 플러그인입니다.
핵심 역할은 어떤 frontend task가 overall workflow, architecture-pattern choice, React structure, local component boundary, domain modeling, UI implementation quality, test-driven execution 중 무엇이 중심인지 분류하는 것입니다.

## 경계

- 포함:
  - frontend workflow selection
  - Atomic, FSD, hybrid 같은 architecture-pattern decision
  - React layering, hook, effect, rendering boundary
  - local component boundary 및 ownership decision
  - frontend domain-modeling threshold decision
  - UI implementation quality와 designer-level frontend guidance
  - frontend TDD workflow
- 제외:
  - backend architecture
  - design-only pre-code planning
  - frontend 특화가 아닌 general workflow orchestration

## 진입 표면

- 대표 엔트리포인트: `frontend-engineering-kit-guide`
- 핵심 분기: dominant frontend concern을 분류하고 첫 specialist skill을 고른다

## 스킬 구성

- `frontend-engineering-kit-guide`: broad하거나 mixed된 frontend task를 적절한 specialist path로 라우팅한다
- `frontend-workflow-guide`: project structure를 읽고 deeper concern의 순서를 정하는 기본 frontend execution workflow다
- `frontend-architecture-patterns`: Atomic, FSD, hybrid 구조 적합성을 결정하거나 리뷰한다
- `react-architecture`: React-specific layering, hook, context, effect, rerender 구조 문제를 다룬다
- `component-architecture`: local component boundary, ownership, extraction, refactor direction을 다룬다
- `frontend-domain-modeling`: 프론트엔드에 어느 정도 modeling이 필요한지와 business meaning의 위치를 결정한다
- `frontend-design-guide`: design intent를 concrete frontend UI implementation guidance로 바꾼다
- `frontend-tdd-rgb`: 적절한 failing test와 Red-Green-Refactor loop로 frontend work를 이끈다

## 확장 원칙

- 새 skill은 workflow, pattern, React, component, domain, design, TDD 어느 곳에서도 깔끔하게 소유되지 않는 concern일 때만 추가한다.
- plugin-level routing은 `frontend-engineering-kit-guide`에 두고 specialist skill에 흩뿌리지 않는다.
- design-only planning이나 backend architecture를 이 플러그인에 섞지 않는다.
- 새 skill이 concern boundary를 바꾸면 guide skill과 이 spec을 같은 변경에서 함께 갱신한다.

## 현재 의도 점검

- 현재 플러그인 표면은 frontend decision layer를 기준으로 일관적이다.
- 현재의 주요 리스크는 React architecture, component architecture, domain-modeling 경계가 충분히 날카롭지 않으면 겹침이 생기는 것이다.
