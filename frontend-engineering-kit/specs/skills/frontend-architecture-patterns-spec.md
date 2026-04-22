# frontend-architecture-patterns 스킬 스펙

## 목적

`frontend-architecture-patterns`는 프론트엔드 코드베이스가 Atomic Design, FSD, hybrid, 혹은 다른 custom structure 중 어떤 패턴을 채택해야 하는지 판단하는 스킬입니다.

## 경계

- 포함:
  - codebase-structure fit 판단
  - migration direction 비교
  - UI taxonomy와 feature slicing의 배치 판단
- 제외:
  - React component 내부 구조
  - local component refactor
  - design-only planning

## 처리하려는 작업 형태

- 프로젝트 구조 방향을 선택하거나 리뷰해야 하는 경우
- 기존 구조를 더 적합한 패턴으로 옮길지 판단해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/frontend-architecture-patterns/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`

## 핵심 처리 계약

- pattern 비교는 codebase fit, migration cost, boundary clarity 기준으로 해야 한다.
- selection과 review는 같은 평가축을 사용해야 한다.
- local component나 React detail은 후속 specialist skill로 넘긴다.

## 독립성 원칙

- 이 스킬은 top-level structure decision만 소유한다.
- 선택 근거는 sibling skill 없이도 codebase 구조 판단에 바로 쓸 수 있어야 한다.

## 확장 원칙

- 새로운 pattern type은 distinct structure model이 있을 때만 추가한다.

