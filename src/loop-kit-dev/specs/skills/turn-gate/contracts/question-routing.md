# turn-gate question-routing 계약

## 소유 범위

이 문서는 `next turn-flow / 메시지 수신`과 질문 복구를 소유합니다.

## 계약

`flow.end` 이후 explicit stop이 없으면 `next turn-flow / 메시지 수신`을 엽니다.
이 상태는 terminal response가 아니라 다음 입력을 받기 위한 열린 routing입니다.

다음 입력 경로:

- 사용자 메시지 수신
- self-drive 모드의 자체 해석
- blocker decision
- approval decision
- explicit stop

`request_user_input`을 사용할 수 있고 선택지가 좁으면 질문 도구를 씁니다.
도구가 없으면 active plain-text question fallback을 씁니다.

## question recovery

question abort, cancel, interrupt는 flow completion이나 terminal closure가 아닙니다.
pending question을 기록하고 다음 사용자 메시지를 먼저 다음 중 하나로 해석합니다.

- pending question answer
- superseding new flow request
- status/progress question
- explicit stop

## continue

사용자가 "continue", "계속", "이어가"라고 하면 recorded next action을 먼저 확인합니다.
identity, target, scope, endpoint, approval boundary, verification expectation이 알려져 있을 때만 계속합니다.
그 외에는 next turn-flow 질문을 다시 엽니다.
