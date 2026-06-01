# turn-gate question-routing 계약

## 소유 범위

이 문서는 handoff 뒤 질문 라우팅과 질문 복구를 소유합니다.

## 계약

`flow skill: handoff` 이후 explicit stop이 없으면 `next-flow gate`를 엽니다.
이 상태는 terminal response가 아니라 `질문 도구: 다음 플로우 선택`으로 다음 flow 입력을 고르는 열린 routing입니다.

`next-flow gate` 순서:

- 사용중인 스킬 다시 읽기
- 질문 도구 응답 또는 메시지 수신
- `000-plan.md` 업데이트
- `flow: deep-interview`와 같은 인터뷰 흐름으로 다음 flow 입력 구체화
- 구체화된 입력으로 `flow skill: interview` 재진입

사용중인 스킬 다시 읽기는 다음 flow 질문을 만들기 전에 항상 수행합니다.
`000-plan.md` 업데이트는 선택된 다음 flow 입력, 사용 skill, pending/answered question 상태, next action을 매번 반영합니다.
`request_user_input`을 사용할 수 있고 선택지가 좁으면 `질문 도구: 다음 플로우 선택`으로 표시합니다.
도구가 없으면 active plain-text question fallback을 씁니다.
blocker, approval, explicit stop은 메인 그래프 안의 세부 노드가 아니라 전역 라우팅과 승인 경계에서 처리합니다.

## question recovery

question abort, cancel, interrupt는 flow completion이나 terminal closure가 아닙니다.
pending question을 기록하고 다음 사용자 메시지를 먼저 다음 중 하나로 해석합니다.

- pending question answer
- superseding new flow request
- status/progress question
- explicit stop

## continue

사용자가 "continue", "계속", "이어가"라고 하면 다음 flow 입력이 이미 충분히 구체화됐는지 먼저 확인합니다.
identity, target, scope, endpoint, approval boundary, verification expectation이 알려져 있을 때만 `flow skill: interview`로 되돌립니다.
그 외에는 handoff 뒤 질문 라우팅을 다시 엽니다.
