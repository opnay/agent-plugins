# frontend-domain-modeling 스킬 스펙

## 목적

`frontend-domain-modeling`은 프론트엔드 코드가 thin contract layer로 남아야 하는지, 아니면 explicit domain model을 도입해야 하는지 판단하는 스킬입니다.

## 경계

- 포함:
  - business rule의 위치 판단
  - view model과 domain logic 분리
  - feature boundary와 ubiquitous language 정리
  - contract-near code와 explicit model의 tradeoff 판단
- 제외:
  - backend domain modeling
  - UI hierarchy guidance
  - local component extraction

## 처리하려는 작업 형태

- 프론트엔드에 business meaning이 누적되어 경계가 흐려진 경우
- server contract를 그대로 둘지, policy/value/entity-like model을 둘지 고민하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/frontend-domain-modeling/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`

## 핵심 처리 계약

- default는 thin frontend를 기준으로 시작하고, 실제 business rule 밀도가 높을 때만 modeling을 늘린다.
- domain model 도입 여부는 boundary clarity와 maintenance cost로 판단한다.
- model depth, placement, review questions를 함께 제시해야 한다.

## 독립성 원칙

- 이 스킬은 React/component 구조나 backend domain design을 소유하지 않는다.
- frontend domain threshold 판단만 단독으로 이해 가능해야 한다.

## 확장 원칙

- 새 rule은 frontend에 domain object를 둘 시점 판단을 더 명확하게 만들 때만 추가한다.

