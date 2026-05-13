# turn-gate preparation phase sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `preparation` phase 세부 계약을 소유합니다.

`core/runtime-flow.md`는 phase 순서와 lifecycle guard를 소유하고, 이 문서는 work로 넘어가기 전 무엇을 잠그고 기록해야 하는지 소유합니다.

## 계약

- 이 단계는 flow shaping gate를 통과해 active flow의 경계와 completion criteria를 정한다.
- work로 넘어가기 전에 이번 flow의 intent, scope, non-goal, acceptance signal, verification expectation, approval boundary가 충분한지 확인한다.
- 사용자 메시지 기반 preparation과 이미 선택된 flow의 실행 전 preparation을 구분한다.
- scope가 비어 있거나 너무 넓거나 결과/검증 경로를 바꿀 수 있으면 user-gated question-routing으로 잠근다.
- 추론한 scope로 진행하는 경우에도 work boundary와 non-goal을 flow record에 남긴다.
- planned flow boundary, `operational-preparation`, `change-unit`, 후속 후보와 active execution flow 구분은 `core/flow-boundaries.md`를 따른다.
- operation/target ambiguity는 work 전에 `core/meaning-resolution.md`로 잠근다.
- destructive, irreversible, external, commit/push/PR/publish 같은 위험 작업은 `core/approval-boundary.md`의 approval-sensitive checkpoint로 둔다.
- session plan, flow record, Continuity Guard 기록 방식은 `records/session-records.md`를 따른다.
- meaningful work가 시작되면 계획 도구를 사용해 현재 phase 상태를 유지한다.

## 검토 질문

- work로 넘어가기 전에 intent, scope, non-goal, acceptance signal, verification expectation, approval boundary가 충분히 잠겼는가?
- 사용자 메시지 기반 preparation과 기존 flow 기반 preparation을 구분했는가?
- 후속 후보와 active execution flow를 섞지 않았는가?
- 위험 작업은 user-gated checkpoint로 기록했는가?
