# turn-gate session-records 계약

## 소유 범위

이 문서는 `turn-gate` 활성 중 유지되는 session record 계약을 소유합니다.
session record는 현재 flow, 다음 행동, 질문 대기 상태, 검증 상태, 명시적 종료 여부를 compaction 또는 interruption 뒤에도 복구할 수 있게 하는 operational continuity 표면입니다.

session record는 작업 산출물 자체가 아닙니다.
Git이나 tool readback으로 복구 가능한 긴 이력은 반복하지 않습니다.

## 파일 계약

- `.agents/sessions/{YYYYMMDD}/000-plan.md`: date-level routing card입니다. Active flow pointer, next action, closure flags, self-drive pointer, unapproved actions, active skill list를 소유합니다.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`: 하나의 active flow record입니다. Flow contract, phase checklist, execution log, verification evidence, result, residual risk, handoff 또는 next action을 소유합니다.
- `.agents/sessions/{YYYYMMDD}/000-self-drive.md`: self-drive active 상태에서만 쓰는 optional sequence-level state입니다. `000-plan.md`는 pointer만 갖습니다.
- `.agents/sessions/{YYYYMMDD}/000-review.md`: optional retrospective note입니다. Flat tagged list만 소유하며 active routing, raw flow log, verification authority, closure authority를 소유하지 않습니다.

runtime template은 `skills/turn-gate/templates/`가 소유합니다.
spec은 template 의미와 최소 계약만 정의하며, runtime reader가 dev-only spec path를 읽어야 한다고 요구하지 않습니다.

## plan record 계약

`000-plan.md`는 가장 작은 routing card로 유지합니다.

Frontmatter에는 다음이 드러나야 합니다.

- `turn_gate_active`
- `active_flow`
- `next_action`
- closure state
- `self_drive`와 `self_drive_sidecar`
- `unapproved_actions`
- `active_skills`

Body에는 다음만 둡니다.

- 현재 요청과 복구에 필요한 직전 routing signal
- active/recent/archive flow index
- continuity note

완료 flow summary, 전체 flow history, detailed evidence, Git으로 재구성 가능한 상태, self-drive detail은 plan에 누적하지 않습니다.

## flow record 계약

flow filename은 zero-padded counter와 lowercase English slug를 사용합니다.
active flow boundary가 바뀌면 새 flow record를 만듭니다.
같은 flow가 계속 active이거나 reporting 전 자기 metadata를 고치는 경우에만 기존 flow record를 갱신합니다.
있어야 하는 active flow record가 unexpectedly missing이면 조용히 재구성하지 않고 blocker recovery로 라우팅합니다.

기본 섹션은 다음입니다.

- `Contract`
- `Phase Checklist`
- `Execution Log`
- `Result`

approval-sensitive action이 있으면 `Risky Action`을 추가합니다.
raw user request text가 해석에 영향을 주면 summary 또는 interpretation과 분리해 `Contract`에 기록합니다.
`fail`, `blocked`, `insufficient` 결과에는 필요할 때 `Result` 아래 non-pass routing을 추가할 수 있습니다.

`Contract`는 scope, exclude, done, boundary를 최소 필드로 둡니다.
boundary는 commit, push, PR, publish, release, version bump, destructive/external action의 제외 또는 승인 상태를 드러내야 합니다.
readiness, verification, build, generated release surface, self-drive, 이전 맥락은 승인 민감 작업의 실행 권한이 아닙니다.

## Phase Checklist 계약

`Phase Checklist`는 active flow가 필수 lifecycle 단계를 종료 checkpoint까지 통과했는지 보여주는 복구 표면입니다.
frontmatter의 `phase`는 현재 위치를, checklist는 이미 지난 단계를 나타냅니다.

기본 항목은 다음입니다.

- `intake`
- `framing`
- `preparation`
- `work`
- `verification`
- `reporting`
- `next-flow`

각 항목은 해당 phase의 end checkpoint가 기록된 뒤에만 체크합니다.
phase 시작만으로 체크하지 않습니다.
검증 결과가 `fail`, `blocked`, `insufficient`이면 verification end checkpoint와 non-pass routing을 기록한 뒤 현재 사실에 맞게 checklist를 둡니다.

`interruption`은 checklist 항목이 아닙니다.
active flow 중 들어온 사용자 메시지를 분류하는 entry-only event이므로 `Execution Log`에 기록합니다.

## Execution Log 계약

`Execution Log`는 flow 복구의 주 기록입니다.
각 항목은 phase prefix 또는 event label과 함께 기록합니다.

기록 대상:

- phase 시작 또는 종료 결과
- user-gated question, pending question, answered question
- approval-sensitive checkpoint 상태
- 편집, build, 검증, verifier 결과
- interruption 분류와 라우팅
- reporting outcome과 next-flow reopening

로그는 복구 가능한 운영 사실만 남깁니다.
검증 실패, skipped verification, insufficient evidence, blocker는 숨기지 않습니다.

## Continuity Metadata 계약

모든 active flow record frontmatter는 다음 metadata를 유지합니다.

- `phase`
- `verification_status`
- `next_action`
- `flags`
- `answered_question`
- 대기 중인 질문이 있으면 `pending_question`
- `continuity`

`flags`는 recovery에 필요한 상태만 나열합니다.
예: `turn_gate_active`, `terminal_summary_blocked`, `question_pending`, `blocked`, `approval_required`, `explicit_stop_recorded`.

`verification_status`는 다음 값만 사용합니다.

- `not-started`
- `requested`
- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-started`와 `requested`는 성공 근거가 아니며 terminal closure 또는 successful reporting 근거가 될 수 없습니다.
이전 flow state를 보존할 때는 기존 status를 유지하고 보존 사실을 `continuity`에 기록합니다.

metadata는 phase 시작, phase 종료, reporting 전, next-flow reopening 전에 갱신합니다.
현재 source-recorded explicit stop만 terminal closure authority를 만들 수 있습니다.
stale 또는 source-less closure state는 열린 continuity로 reset하고 recovery 사실을 기록합니다.

## 질문과 next-flow 계약

reporting 뒤 explicit stop이 없으면 next-flow routing을 엽니다.
질문 도구가 있으면 좁은 선택지를 쓰고, 없으면 plain-text fallback으로 다음 행동을 묻습니다.
질문 도구 abort, cancel, interruption은 terminal closure가 아닙니다.

pending question이 있으면 다음 사용자 메시지를 먼저 다음 중 하나로 해석합니다.

- pending question answer
- superseding new flow request
- status/progress question
- explicit stop

질문 상태는 `answered_question`과 `pending_question`으로 기록하고 `question_state` 같은 동의어를 만들지 않습니다.

## 복구 계약

- not-yet-created plan: runtime template으로 첫 plan을 만듭니다.
- not-yet-created flow: runtime template으로 선택된 새 flow record를 만듭니다.
- unexpectedly missing active record: blocker를 보고하거나 recovery 선택을 묻습니다.
- inaccessible active record: 접근이 복구되거나 사용자가 recovery를 선택할 때까지 blocker로 둡니다.
- stale closure state: closure authority를 reset하고 recovery를 기록합니다.
- stale self-drive sidecar: plan이 self-drive inactive이면 historical context로 취급합니다.
- stale routing mismatch: latest source에서 reconcile하거나 질문합니다.

read-only 요청은 보통 source artifact 변경을 금지하는 것이지 session record 쓰기까지 금지하는 것이 아닙니다.
사용자가 모든 write 또는 record 생성을 금지한 경우에만 session record를 쓰지 않습니다.
`no-record`, `기록 남기지 마`, `세션 기록 없이`처럼 session record 읽기까지 금지하는지 모호하면 기존 record를 읽기 전에 clarification으로 라우팅합니다.

## 검토 기준

- `000-plan.md`가 active flow와 next action을 과도한 이력 없이 복구하는가?
- flow record가 `Contract`, `Phase Checklist`, `Execution Log`, `Result`를 유지하는가?
- checklist가 phase 시작이 아니라 종료 checkpoint 통과 여부를 나타내는가?
- `interruption`이 checklist 항목이 아니라 log event로 남는가?
- frontmatter metadata와 checklist 의미가 겹치지 않는가?
- pending question, verification status, explicit stop state가 compaction 뒤에도 복구 가능한가?
- 승인 민감 작업이 readiness, verification, build, self-drive, 이전 맥락에서 암묵 승인되지 않는가?
