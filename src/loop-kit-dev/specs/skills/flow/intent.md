## 사용자 스펙 의도

- `flow`는 하나의 메시지, 동작, 계획 항목, 검토 finding, handoff를 응집된 flow 단위로 해석해야 한다.
- flow는 phase checklist가 아니라 함께 이해하고, 검토하고, 검증하고, 필요하면 커밋할 수 있는 작업 단위여야 한다.
- 각 flow는 내부 단계로 `preparation -> work -> verification -> reporting`을 갖지만, 여러 flow 사이의 turn 지속이나 next-flow 질문은 소유하지 않는다.
- 큰 요청은 parent flow가 될 수 있고, parent flow는 finite `sub-flow candidates`를 만들 수 있어야 한다.
- `sub-flow candidate` 생성은 실행이 아니며, active flow 전환은 `turn-gate` next-flow routing 또는 준비된 self-drive sequence를 통해서만 가능해야 한다.
- flow는 `operational-preparation flow`와 `change-unit flow`를 구분해야 한다.
- 사용자 메시지를 받아 의도, scope, non-goal, acceptance signal, approval boundary, verification expectation, planned flow list를 정리하는 일 자체도 산출물을 가진 `operational-preparation flow`가 될 수 있어야 한다.
- 실제 코드, 문서, fixture, 설정, release surface 변경은 `change-unit flow`로 분리되어야 한다.
- requirement discovery, readiness, operation/target ambiguity는 `flow`의 preparation 계약이어야 한다.
- discovery는 질문 도구 실행 방식이 아니라 flow contract를 만들기 위한 질문 주제와 missing field를 산출하는 역할이어야 한다.
- review-loop, fix-verify-loop, broad-execution은 active flow 안에서 선택되는 flow-local strategy여야 한다.
- review-loop는 여러 review/QA/self-review finding 전체를 한 번에 실행하는 전략이 아니라, active flow 안의 bounded blocking finding 하나를 처리하는 전략이어야 한다.
- 여러 finding이 있으면 우선순위가 가장 높은 bounded blocking finding 하나를 선택하거나, discovery/parent flow로 finite follow-up 후보를 만들어야 한다.
- commit-readiness는 commit 실행이 아니라 flow handoff condition으로 판단해야 한다.
- flow는 질문 도구 실행, turn closure, session continuity, explicit stop handling, approval-sensitive execution authority를 소유하지 않는다.
- `turn-gate`는 flow decision을 현재 turn에 적용하고 기록하지만, flow boundary, readiness, discovery, flow-local strategy를 재정의하지 않아야 한다.
- `turn-gate`의 intent 중 flow 해석, flow 분해, readiness/discovery, review/fix/broad execution, commit-readiness 판단에 해당하는 내용은 `flow` intent와 spec이 소유해야 한다.
- 각 phase의 시작과 종료에서 `plan.md`나 flow 문서를 수정하도록 해야 한다.
