# frontend-tdd-rgb 스킬 스펙

## 목적

`frontend-tdd-rgb`는 프론트엔드 변경을 failing test와 Red-Green-Refactor 루프로 진행하도록 이끄는 스킬입니다.

## 경계

- 포함:
  - 적절한 test level 선택
  - failing test 우선 workflow
  - refactor 가능한 검증 루프 설계
- 제외:
  - general test theory
  - UI design guidance
  - architecture pattern choice

## 처리하려는 작업 형태

- component, hook, integration, state-level 테스트 중 무엇으로 시작할지 정해야 하는 경우
- UI behavior 변경을 test-driven으로 진행해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/frontend-tdd-rgb/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`, `frontend-workflow-guide`

## 핵심 처리 계약

- 적절한 test level을 먼저 선택하고 failing test를 만든다.
- Green 이후에는 refactor와 regression check를 붙인다.
- brittle assertion을 피하고 behavior 중심 검증을 우선한다.

## 독립성 원칙

- 이 스킬은 구조 설계 전체를 소유하지 않는다.
- test-driven execution workflow만 단독으로 적용 가능해야 한다.

## 확장 원칙

- 새 규칙은 testability와 behavior verification을 개선하는 경우에만 추가한다.

