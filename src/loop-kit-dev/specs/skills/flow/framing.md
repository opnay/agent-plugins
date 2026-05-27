# flow framing 계약

## 소유 범위

intake 결과를 바탕으로 flow를 분리하고, active flow 또는 finite sub-flow candidate를 설계하는 단계.

## 계약

- framing은 현재 항목을 active flow, parent flow, sub-flow candidate, phase, handoff 중 하나로 분류한다.
- framing은 parent flow가 필요한 경우 finite sub-flow candidates를 만든다.
- framing은 candidate와 selected active flow를 구분한다. candidate 생성은 실행이 아니다.
- framing은 각 flow 또는 candidate의 artifact ownership을 드러낸다.
- framing은 flow boundary, scope edge, non-goals, completion criteria, verification expectation, handoff condition의 초안을 만든다.
- framing 결과가 선택된 active flow로 확정된 뒤에만 preparation으로 넘어간다.

## 검토 기준

- flow 후보가 phase list가 아니라 review/verification/commit 가능한 작업 단위인가?
- candidate와 selected flow를 혼동하지 않았는가?
- artifact ownership이 드러나는가?
- framing 결과가 work 실행 권한으로 확대되지 않았는가?
