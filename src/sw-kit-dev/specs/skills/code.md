## 사용자 스펙 의도

- 코드의 구조와 최적화는 code 레이어가 맡아야 한다.
- 엔지니어링 방향과 source-level 구현 판단을 분리하고 싶다.

---

# Code Skill Spec

## 목적

`code`는 선택된 방향 안에서 정확하고 이해하기 쉬우며 유지 가능한 production code를 구현, 수정, 리팩터링, 테스트, 리뷰한다.

## 경계

- 포함: 파일·module 구조, control flow, error handling, local reuse, abstraction, tests, source-level dependency use, risk-focused review.
- 제외: system architecture와 technology/storage 선택, product requirement 결정, repository-wide lifecycle audit, Git workflow.

## 처리 계약

- 관련 코드, callers, tests, repository conventions를 확인한다.
- 현재 요구에 맞는 가장 작은 일관된 변경을 만든다. 미래 가능성만으로 abstraction을 만들지 않는다.
- 재사용은 의미, contract, ownership, lifecycle이 맞을 때만 선택한다.
- 오류와 boundary condition을 명시적으로 처리하고 중요한 동작을 검증한다.
- 리뷰에서는 결함과 유지보수 위험을 영향 순으로 제시하며, 수정 권한이 없는 요청에는 findings만 낸다.

## 검토 질문

- 요구와 기존 contract를 보존하는가?
- 코드 흐름, 책임, 실패 처리가 이해 가능한가?
- 재사용 또는 abstraction이 실제 복잡도를 줄이는가?
- 중요한 회귀 위험을 검증했는가?

## 독립성 원칙

기술 방향이 주어지면 따른다. 주어지지 않으면 local implementation 결정에 머물고 material architecture choice는 `engineering`으로 드러낸다.
