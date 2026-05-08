# turn-gate gate-flow-shaping sub-spec

## 목적

이 문서는 flow shaping gate의 전환 계약을 소유합니다.

flow shaping gate는 message intake 결과를 active flow에 반영합니다.

## 소유

- 새 flow를 만들지, 기존 flow를 갱신할지, reporting이나 question-routing으로 이어갈지 결정한다.
- flow kind를 `operational-preparation`, `change-unit`, user-gated handoff, reporting context, question-routing state 중 적절한 형태로 정한다.
- flow boundary, completion criteria, verification expectation, next-flow reopening 조건을 기록한다.
- 후속 `change-unit` 후보와 active execution flow를 구분한다.

## 비소유

- flow 내부 command sequence 실행
- verification pass 판정
- explicit stop 없는 turn closure

flow shaping gate는 task 목록을 phase checklist나 planned flow list로 오해하지 않게 막아야 합니다.

## 검토 질문

- flow shaping gate가 active flow와 후속 후보를 분리하는가?
- flow boundary와 completion criteria를 기록했는가?
