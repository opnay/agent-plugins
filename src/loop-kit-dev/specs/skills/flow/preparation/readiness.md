# flow readiness preparation 계약

## 소유 범위

framing에서 선택된 active flow가 work로 들어가기 전에 필요한 flow contract 충분성.

## 계약

- readiness 판단은 selected active flow의 preparation 단계이며, intake나 framing을 다시 수행하는 단계가 아니다.
- intake 또는 framing에서 잠기지 않은 goal, non-goal, authority, candidate 경계가 남아 있으면 work로 넘어가지 않는다.
- work로 넘어가기 전에 최소한 scope, non-goals, completion criteria, verification expectation, handoff condition을 확인한다.
- approval-sensitive checkpoint가 예상되면 flow output에 그 필요성을 드러내되, 실행 승인 자체는 approval boundary가 소유한다.
- flow contract가 부족하면 부족한 필드와 질문 주제를 산출한다.
- 질문을 어떤 도구로 사용자에게 보여줄지는 question-routing 또는 turn-gate 적용 표면이 소유한다.
- readiness 판단은 active flow와 sub-flow candidate를 구분해야 한다. candidate는 선택되기 전까지 work로 들어가지 않는다.
- active flow 도중 새 사용자 메시지가 들어오면 그 메시지가 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 바꾸는지 판단한다. 바뀌면 current-flow revision 또는 새 flow 전환 근거를 산출하고, 바뀌지 않으면 inline answer 또는 reporting constraint 근거를 산출한다.
- self-drive가 다음 flow로 advance하려면 current flow completion, pass에 해당하는 verification expectation 충족, non-blocked handoff, known next identity가 flow output에서 확인돼야 한다.

## 검토 기준

- work에 들어갈 만큼 scope, non-goals, completion criteria가 충분한가?
- verification expectation과 handoff condition이 분리돼 있는가?
- 부족한 contract field가 질문 주제로 드러났는가?
- approval-sensitive checkpoint 필요성과 실행 승인 권한을 혼동하지 않았는가?
- interruption 또는 self-drive advance 판단에서 flow contract 변경 여부가 명시됐는가?
