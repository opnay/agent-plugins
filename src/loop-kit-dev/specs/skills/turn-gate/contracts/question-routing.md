# turn-gate question-routing 계약

## 소유 범위

이 문서는 handoff 뒤 질문 라우팅과 질문 복구를 소유합니다.

## 계약

`flow skill: handoff` 이후 explicit stop이 없으면 즉시 `next-flow gate`를 엽니다.
이 상태는 terminal response가 아니라 `다음 플로우 선택`으로 다음 flow 입력을 고르는 열린 routing입니다.
handoff result, final-looking summary, status answer, verification pass는 이 gate를 건너뛰는 closure authority가 아닙니다.
Runtime에서는 이 경계를 `<gate:next-flow>...</gate:next-flow>` 태그로 감싸고, 내부 `skill reconfigure` 경계를 `<gate:skill-reconfigure>...</gate:skill-reconfigure>` 태그로 감쌉니다.
일반 진입은 질문 도구를 사용하고, prepared self-drive 진입은 질문 도구를 대체합니다.

`next-flow gate` 순서:

- `skill reconfigure`
- `다음 플로우 선택`: 질문 도구 응답, 메시지 수신, 또는 prepared self-drive gate
- self-drive일 경우 `000-self-drive.md` 업데이트
- `000-plan.md` 업데이트
- `flow: deep-interview`와 같은 인터뷰 흐름으로 다음 flow 입력 구체화
- 구체화된 입력으로 `flow skill: interview` 재진입

`skill reconfigure` 그룹은 `flow skill: handoff`에서 시작해 다음 flow 질문을 만들기 전에 세션에서 사용중인 전체 skill 목록을 식별하고, 기존에 읽은 skill context를 폐기하고, 각 skill 본문을 source에서 새로 읽고, freshly read bodies만 새 active skill set으로 수용하는 과정입니다.
Runtime에서는 이 과정을 `<gate:skill-reconfigure>` 태그 안에 둡니다.
`000-plan.md` 업데이트는 선택된 다음 flow 입력, 사용 skill, pending/answered question 상태, next action을 매번 반영합니다.
`request_user_input`을 사용할 수 있고 선택지가 좁으면 질문 도구로 `다음 플로우 선택`을 표시합니다.
도구가 없으면 active plain-text question fallback을 씁니다.
blocker, approval, explicit stop은 메인 그래프 안의 세부 노드가 아니라 전역 라우팅과 승인 경계에서 처리합니다.

## question recovery

question abort, cancel, interrupt는 flow completion이나 terminal closure가 아닙니다.
pending question을 기록하고 다음 사용자 메시지를 active turn 안의 열린 입력으로 해석합니다.
현재 메시지가 source-recorded explicit stop이 아니면 terminal closure로 닫지 않습니다.
필요할 때만 다음 효과 중 하나로 정리합니다.

- pending question answer
- superseding new flow request
- status/progress question
- explicit stop

방향 전환, 작업 변경, 추가 질문, 오류 지적은 pending question을 supersede할 수 있지만 turn 종료 권한은 만들지 않습니다.
다음 입력이 충분히 구체적이지 않으면 같은 인터뷰 흐름으로 다시 구체화합니다.

## continue

사용자가 "continue", "계속", "이어가"라고 하면 다음 flow 입력이 이미 충분히 구체화됐는지 먼저 확인합니다.
identity, target, scope, endpoint, approval boundary, verification expectation이 알려져 있을 때만 `flow skill: interview`로 되돌립니다.
그 외에는 handoff 뒤 질문 라우팅을 다시 엽니다.
