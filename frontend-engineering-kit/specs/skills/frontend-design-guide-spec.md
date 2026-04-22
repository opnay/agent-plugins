# frontend-design-guide 스킬 스펙

## 목적

`frontend-design-guide`는 design intent를 실제 프론트엔드 UI 구현 규칙으로 전환하는 스킬입니다.
핵심은 hierarchy, layout, spacing, typography, states, responsiveness, accessibility를 구현 가능한 기준으로 정리하는 것입니다.

## 경계

- 포함:
  - hierarchy와 layout 판단
  - state와 interaction 표현 규칙
  - responsiveness, accessibility, design token usage
  - reusable UI pattern guidance
- 제외:
  - design-only brief
  - Figma execution
  - component ownership refactor

## 처리하려는 작업 형태

- product intent를 concrete UI implementation guidance로 바꿔야 하는 경우
- 화면이나 컴포넌트의 시각적/상태적 품질 기준을 정해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/frontend-design-guide/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`, `frontend-workflow-guide`

## 핵심 처리 계약

- guidance는 hierarchy, state, responsiveness, accessibility를 함께 다뤄야 한다.
- design token과 reusable pattern을 우선 고려한다.
- pure design planning이 아니라 implementation-ready guidance를 산출해야 한다.

## 독립성 원칙

- 이 스킬은 pre-code design brief를 소유하지 않는다.
- UI implementation guidance만 읽어도 frontend 적용 기준이 이해 가능해야 한다.

## 확장 원칙

- 새 규칙은 reusable UI quality를 높이는 경우에만 추가한다.
- 구조 리팩터링 이슈는 component/react skill로 넘긴다.

