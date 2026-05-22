# turn-gate question-routing 계약

## 목적

이 계약은 reporting 뒤 user-gated routing, clarification, blocker decision, structured `request_user_input` 사용, fallback, question abort recovery를 소유합니다.

## 다음 Flow 라우팅

reporting 뒤 현재 사용자 메시지가 턴을 명시적으로 끝내지 않았다면 다음 flow를 다시 열어야 합니다.

next-flow reopening은 final response가 아닙니다. 턴을 ongoing conversation channel에 열린 상태로 두고 필요한 다음 행동이나 선택지를 보여줍니다. terminal/final closeout을 next-flow routing의 대체물로 사용하지 않습니다.

structured choices가 가능하고 도구가 사용 가능하면 `request_user_input`을 사용합니다. visible choice는 좁고 보고된 결과와 연결되어야 합니다. visible choice에 stop option을 표시하지 못하더라도 flow record의 `Next Flow Options`에는 explicit turn-end option을 남깁니다. 도구 UI 제약으로 stop choice를 직접 넣지 못하는 경우에도 user-facing prompt나 fallback text에서 사용자가 명시적으로 턴을 종료할 수 있음을 드러내야 합니다.

도구가 없으면 active plain-text fallback을 사용합니다. 도구가 없다고 밝히고, 열린 선택지를 나열하며, required next action을 기록합니다. fallback도 active routing이며 terminal summary가 아닙니다.

다음 항목이 실행을 바꿀 수 있으면 question routing이 필요합니다.

- 다음 flow 선택
- scope, target, endpoint, acceptance signal clarification
- blocker recovery
- approval-sensitive action boundary
- pending question state와 latest user message의 충돌

질문은 지금 필요한 결정만 물어야 합니다. 가능하다는 이유만으로 관계없는 미래 작업 선택지를 묶지 않습니다.

## 중단 복구

`request_user_input`이 사용자 interrupt로 abort 또는 cancel되더라도, 그것은 flow completion, explicit stop, terminal close authority가 아닙니다.

recovery state를 다음처럼 기록합니다.

- `user_explicit_stop: no`
- `terminal_summary_allowed: no`
- `confirmed_closure: no`
- pending question state는 `aborted`, `interrupted`, `superseded` 중 하나
- 알 수 있으면 pending question id 또는 summary

다음 사용자 메시지는 먼저 question-routing recovery로 해석합니다.

- pending question의 답변: 그 답변에서 계속합니다.
- 새로운 flow 요청: pending question을 superseded로 표시하고 새 flow를 준비합니다.
- status/progress 질문: active flow, pending question, verification state, required next action을 보고한 뒤 routing을 다시 엽니다.
- explicit stop: source를 기록한 뒤에만 턴을 닫습니다.

같은 question tool call을 즉시 반복해 interrupt loop를 만들지 않습니다. recovery state를 텍스트로 설명하거나 사용자가 제공한 선택 또는 요청에서 진행합니다.

다음 사용자 메시지가 답변인지 새 요청인지 모호하면 추측보다 clarification을 우선합니다. 메시지가 visible label과 일치하지 않더라도 free-form answer가 명확하면 답변으로 취급합니다.

pending question이 superseded되면 superseded question id 또는 summary와 새 flow source를 기록합니다. 질문이 superseded되더라도 이전 flow의 report나 verification result는 지우지 않습니다.

## 명확화와 차단 상태

scope, target, endpoint, approval boundary, blocker state, current-flow identity가 성공 조건이나 verification path를 바꿀 수 있으면 user-gated routing을 사용합니다. 이 경계를 추측으로 넘지 않습니다.

blocker routing은 무엇이 막혔는지, 어떤 evidence를 모았는지, 어떤 결정이나 access가 필요한지, blocker가 해결되기 전까지 어떤 work가 제외되는지 말해야 합니다.
