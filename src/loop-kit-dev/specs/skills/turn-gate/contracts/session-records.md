# turn-gate session-records 계약

## 소유 범위

이 문서는 `turn-gate` 활성 중 유지되는 session record 적용과 복구 계약을 소유합니다.
shared record template 의미와 파일명 규칙은 `flow`가 소유합니다.

session record는 현재 flow, 다음 행동, 질문 대기 상태, 검증 상태, 명시적 종료 여부를 compaction 또는 interruption 뒤에도 복구하게 하는 operational continuity 표면입니다.
작업 산출물 자체가 아니며, Git이나 tool readback으로 복구 가능한 긴 이력을 반복하지 않습니다.

## 파일 계약

- plan record: date-level routing card입니다. Active flow pointer, next action, closure flags, self-drive pointer, unapproved actions, active skill list를 적용합니다.
- flow record: 하나의 active flow record입니다. `flow` contract, phase state, execution log, verification evidence, result, residual risk, handoff 또는 next action을 적용합니다.
- self-drive record: self-drive active 상태에서만 쓰는 optional sequence-level state입니다. `turn-gate`가 template과 sidecar gate를 소유합니다.
- review record: retrospective note입니다. `flow`가 template 의미를 소유하고, `turn-gate`는 active turn에서 필요한 경우 갱신만 적용합니다.

shared runtime templates for plan, flow record, and review are owned by the `flow` skill's bundled templates.
`turn-gate` must not reference those templates through sibling filesystem paths at runtime.
`turn-gate` owns only its bundled self-drive runtime template.

## 적용 계약

`000-plan.md`는 routing card 크기로 유지합니다.
완료 flow summary, full history, detailed evidence, Git으로 재구성 가능한 상태, self-drive detail은 누적하지 않습니다.

active flow boundary가 바뀌면 새 flow record를 만듭니다.
같은 flow가 계속 active이거나 reporting 전 자기 metadata를 고치는 경우에만 기존 flow record를 갱신합니다.

`000-review.md`는 retrospective note만 담습니다.
active routing, raw flow log, verification authority, closure authority, commit/release authority를 소유하지 않습니다.

## recovery 계약

- not-yet-created plan: flow-owned runtime template으로 첫 plan을 만듭니다.
- not-yet-created flow: flow-owned runtime template으로 선택된 새 flow record를 만듭니다.
- unexpectedly missing active record: blocker를 보고하거나 recovery 선택을 묻습니다.
- inaccessible active record: 접근이 복구되거나 사용자가 recovery를 선택할 때까지 blocker로 둡니다.
- stale closure state: closure authority를 reset하고 recovery를 기록합니다.
- stale self-drive sidecar: plan이 self-drive inactive이면 historical context로 취급합니다.
- stale routing mismatch: latest source에서 reconcile하거나 질문합니다.

read-only 요청은 보통 source artifact 변경을 금지하는 것이지 session record 쓰기까지 금지하는 것이 아닙니다.
사용자가 모든 write 또는 record 생성을 금지한 경우에만 session record를 쓰지 않습니다.

## 질문 계약

reporting 뒤 explicit stop이 없으면 next-flow routing을 엽니다.
질문 상태는 `answered_question`과 `pending_question`으로 기록하고 `question_state` 같은 동의어를 만들지 않습니다.

pending question이 있으면 다음 사용자 메시지를 먼저 다음 중 하나로 해석합니다.

- pending question answer
- superseding new flow request
- status/progress question
- explicit stop

## 검토 기준

- `turn-gate`가 shared template 의미를 재정의하지 않는가?
- active flow와 next action을 compaction 뒤 복구할 수 있는가?
- pending question, verification status, explicit stop state가 복구 가능한가?
- 승인 민감 작업이 readiness, verification, build, self-drive, 이전 맥락에서 암묵 승인되지 않는가?
