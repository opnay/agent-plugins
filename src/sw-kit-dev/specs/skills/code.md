## 사용자 스펙 의도

- 코드의 구조와 최적화는 code 레이어가 맡아야 한다.
- 엔지니어링 방향과 source-level 구현 판단을 분리하고 싶다.
- sw-kit 플러그인 수정하자. 코드 작성시 한 함수안에서도 맥락이 바뀌는 지점에 빈줄을 남기는 방식의 readability 개선하려고
- prevent로 기본 흐름 방지 맥락이고, 다음이 함수에서 사용할 변수 초기화여서 맥락이 다른데 왜 붙어있지?
  초기화 + 초기화검증은 붙는게 맞긴한데
- 코드작성 논리를 설명하는건데 js 예제를 넣는건 일반화하기 어렵지 않아?
- 그리고 input에 대한 맥락도 같이 있어야되서 설명이 장황해지잖아
- handleImport에서 initial 함수 진입시 필수 검토 if 문 다음에 맥락 바뀌는거 아냐?
  변수 선언시 n 줄 넘어가면 그것만으로도 읽는데 피로도가 생겨 맥락이 바뀌는걸로 취급될 정도인거 같아. 이걸 뭐라 설명해야할지 모르겠네
- oneline if 절이 연속되면 그냥 붙여도 되는걸로하자.
- 매개변수로 함수실행 넣는건 되도록이면 피하는게 맞을거 같아. 이게 188L에서 그런식으로 표현되어있어.

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
- 작성·수정하는 함수에서는 목적과 독해 부담을 기준으로 빈줄을 둔다. 함수 시작에 있는 진입 가드는 첫 읽기 단위로 보고, 그 뒤를 빈줄로 구분한다. 같은 검사 흐름의 연속된 한 줄 `if`는 붙여 쓸 수 있다. 준비와 즉시 검증은 보통 붙이지만, 긴 선언·표현식은 그 자체로 하나의 읽기 단위가 될 수 있다.
- 의미 있는 작업을 다른 호출의 인수 안에서 실행하면 결과에 이름을 붙여 전달한다. 중첩이 짧고 흐름이 분명한 경우에는 직접 전달할 수 있다.
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
