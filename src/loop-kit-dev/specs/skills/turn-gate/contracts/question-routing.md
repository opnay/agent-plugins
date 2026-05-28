# turn-gate question-routing 계약

## 소유 범위

reporting 뒤 next-flow reopening, post-flow continue, clarification, blocker decision, structured question fallback, question abort recovery.

## 계약

reporting 뒤 explicit stop이 기록되지 않았다면 `next-flow`를 엽니다.
`next-flow`는 terminal response가 아니라 다음 행동을 고르는 열린 상태입니다.
따라서 reporting은 다음 질문을 받기 위한 pre-intake surface를 만드는 과정입니다.
보고가 끝나면 사용자에게 필요한 다음 decision을 묻거나, self-drive continuation처럼 이미 준비된 다음 입력 경로로 넘겨야 합니다.

`next-flow`를 기록할 때는 다음을 복구 가능하게 남깁니다.

- lifecycle phase인지, recorded state인지, user-facing routing인지
- 선택 가능한 next action 또는 blocker
- 필요한 decision, access, approval, scope, endpoint, verification gap
- explicit turn-end option의 존재

`request_user_input`을 사용할 수 있고 선택지가 좁으면 반드시 질문 도구를 엽니다.
도구가 없으면 plain-text fallback을 쓰되 active routing임을 분명히 합니다.

`turn-gate`가 active이고 explicit stop이 기록되지 않았다면 reporting 뒤 final/terminal closeout은 금지됩니다.
final-looking wording뿐 아니라 실제 final response도 closure authority가 아닙니다.
간결 모드나 요약 요청도 next-flow 질문, blocker 질문, 또는 valid self-drive handoff를 제거할 수 없습니다.

## Post-Flow Continue

reporting 뒤 사용자가 “continue”, “계속”, “이어가”라고 하면 recorded next action을 먼저 확인합니다.

- next identity, target, scope, endpoint, approval boundary, verification expectation이 알려져 있으면 그 범위 안에서만 계속합니다.
- 불명확하면 다음 flow 선택 또는 clarification을 엽니다.
- post-flow continue는 self-drive 활성화 근거가 아니며 approval-sensitive action 권한도 만들지 않습니다.

active flow 도중의 “continue”는 `interruption`에서 처리합니다.

## Question Recovery

`request_user_input` abort, cancel, interrupt는 flow completion, explicit stop, terminal close authority가 아닙니다.
pending question state를 `aborted`, `interrupted`, `superseded` 중 하나로 기록하고 다음 사용자 메시지를 먼저 recovery로 해석합니다.

다음 메시지는 네 가지 중 하나로 라우팅합니다.

- pending question의 답변
- pending question을 대체하는 새 flow 요청
- status/progress 질문
- source-recorded explicit stop

같은 question tool call을 즉시 반복해 interrupt loop를 만들지 않습니다.
답인지 새 요청인지 모호하면 추측하지 말고 clarification을 엽니다.

## Blocker Routing

scope, target, endpoint, approval boundary, blocker state, current-flow identity가 결과나 검증 경로를 바꿀 수 있으면 user-gated routing을 사용합니다.
blocker report에는 막힌 항목, 모은 evidence, 필요한 결정이나 access, blocker 전까지 제외되는 work를 포함합니다.

## 검토 기준

- reporting 뒤 terminal close 대신 next-flow가 열렸는가?
- reporting이 다음 사용자 입력을 받기 위한 pre-intake decision surface를 만들었는가?
- explicit stop 없이 final/terminal closeout으로 끝내지 않았는가?
- post-flow continue가 recorded next action 안에서만 해석되는가?
- question abort가 closure로 처리되지 않았는가?
- blocker recovery에 필요한 decision/access/evidence가 남아 있는가?
