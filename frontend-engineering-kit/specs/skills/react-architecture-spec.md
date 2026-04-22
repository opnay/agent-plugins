# react-architecture 스킬 스펙

## 목적

`react-architecture`는 선택된 frontend structure 안에서 React component, hook, context, effect, rendering boundary를 설계하거나 리뷰하는 스킬입니다.

## 경계

- 포함:
  - component와 hook boundary
  - context scope와 state ownership
  - effect discipline과 dependency handling
  - rendering/update boundary 판단
- 제외:
  - top-level pattern choice
  - local component split detail만의 문제
  - domain modeling threshold

## 처리하려는 작업 형태

- React-specific architecture 문제를 풀어야 하는 경우
- effect, rerender spread, hook API 설계가 병목인 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/react-architecture/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`

## 핵심 처리 계약

- 전체 pattern choice가 이미 정해졌다는 전제에서 React layer 문제를 푼다.
- state ownership, context scope, effect discipline을 함께 판단한다.
- utility boundary와 rendering boundary를 분리해 설명해야 한다.

## 독립성 원칙

- 이 스킬은 top-level structure나 domain model 도입 시점을 소유하지 않는다.
- React-specific architecture 판단만 따로 읽어도 적용 가능해야 한다.

## 확장 원칙

- 새로운 rule은 React update boundary와 effect discipline을 더 선명하게 만들 때만 추가한다.

