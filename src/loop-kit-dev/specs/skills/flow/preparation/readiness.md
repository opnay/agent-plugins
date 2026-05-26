# flow readiness preparation 계약

## 소유 범위

active flow가 work로 들어가기 전에 필요한 flow contract 충분성.

## 계약

- readiness 판단은 flow-local preparation의 일부이며, turn-level next-flow routing이나 explicit stop 판단이 아니다.
- work로 넘어가기 전에 최소한 scope, non-goals, completion criteria, verification expectation, handoff condition을 확인한다.
- approval-sensitive checkpoint가 예상되면 flow output에 그 필요성을 드러내되, 실행 승인 자체는 approval boundary가 소유한다.
- flow contract가 부족하면 부족한 필드와 질문 주제를 산출한다.
- 질문을 어떤 도구로 사용자에게 보여줄지는 question-routing 또는 turn-gate 적용 표면이 소유한다.
- readiness 판단은 active flow와 sub-flow candidate를 구분해야 한다. candidate는 선택되기 전까지 work로 들어가지 않는다.

## 검토 기준

- work에 들어갈 만큼 scope, non-goals, completion criteria가 충분한가?
- verification expectation과 handoff condition이 분리돼 있는가?
- 부족한 contract field가 질문 주제로 드러났는가?
- approval-sensitive checkpoint 필요성과 실행 승인 권한을 혼동하지 않았는가?
