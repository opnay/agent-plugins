# component-architecture 스킬 스펙

## 목적

`component-architecture`는 기존 feature 안에서 local component boundary와 ownership을 재설계하는 스킬입니다.
핵심은 mixed component를 어디서 나누고 어떤 responsibility를 어느 component나 hook이 소유해야 하는지 결정하는 것입니다.

## 경계

- 포함:
  - component split과 extraction 방향
  - local state, rendering concern, hook ownership 판단
  - nearby module 안의 refactor boundary
- 제외:
  - architecture pattern choice
  - React runtime 전반
  - domain model strategy

## 처리하려는 작업 형태

- 큰 component를 나눠야 하는 경우
- hook extraction이나 state colocating이 필요한 경우
- feature 내부 ownership이 섞여 있는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/component-architecture/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`, `frontend-workflow-guide`

## 핵심 처리 계약

- split 기준은 responsibility, data ownership, rendering boundary를 중심으로 삼는다.
- extraction은 smallest coherent unit을 우선한다.
- local refactor direction과 expected verification path를 함께 제시해야 한다.

## 독립성 원칙

- 이 스킬은 pattern choice나 domain model depth를 소유하지 않는다.
- local boundary decision만 읽어도 refactor 방향을 이해할 수 있어야 한다.

## 확장 원칙

- 새로운 규칙은 local component ownership 판단을 더 명확하게 할 때만 추가한다.

