# turn-gate lifecycle sub-spec

## 목적

이 문서는 `turn-gate` 활성화, 응답 종료 상태, 중간 사용자 입력, explicit stop 처리를 소유합니다.

## 핵심 계약

- 각 incoming message를 같은 loop-gated turn의 현재 입력으로 취급한다.
- 중간 사용자 메시지를 stop, completion, approval-boundary pause로 해석하지 않고 같은 loop-gated turn의 authoritative input으로 취급한다.
- 중간 사용자 메시지는 explicit turn stop, status/progress check, current-flow correction, current-flow priority change, next-flow priority request 중 하나로 분류한다.
- status/progress check라면 현재 phase, blocker 또는 progress, 다음 concrete action을 짧게 보고한 뒤 active flow를 계속한다.
- current-flow correction 또는 current-flow priority change라면 현재 analysis/plan을 즉시 조정하고 가장 이른 안전한 phase부터 이어간다.
- next-flow priority request라면 flow record의 next-flow 후보 중 최우선으로 등록하고 다음 safe handoff point까지 이어간다.
- 이 skill이 사용되면 현재 세션 동안 `turn-gate`를 conversation-level first-class operating rule로 활성화한 것으로 취급한다.
- 이 규칙은 skill 내부 체크리스트가 아니라 assistant response lifecycle 자체에 적용한다.
- concrete task 없이 activation만 요청된 경우에도 `turn-gate`를 활성화하고 session record를 생성 또는 갱신한다. 이때 `user_explicit_stop=false`로 보고, work mode를 성급히 고르지 않고 user-gated next-flow 또는 scope selection을 연다.
- `user_explicit_stop`이 false인 동안 result reporting은 terminal response가 아니며, 반드시 next-flow reopening 또는 active question-routing으로 이어져야 한다.
- 사용자가 명시적으로 턴 종료를 요청했거나 flow record에 confirmed closure가 기록된 경우가 아니라면 일반적인 final summary로 턴을 닫지 않는다.
- 사용자가 explicit stop을 요청하면 active flow record와 `Continuity Guard`에 confirmed closure를 기록하고, `terminal summary allowed`를 허용 상태로 갱신하며, next-flow choices를 다시 열지 않는다.

## Phase Shape

- `analysis`, `plan`, `work`, `verification`, `result reporting`, `question-routing reopening`을 응답 shape에 계속 드러낸다.
- 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 미리 설계할 수 있다.
- future flow/phase 설계는 provisional하며, 이후 loop에서 새 증거, changed intent, 새 blocker가 생겼을 때만 다시 설계한다.
- 계획 이후 current-flow correction이나 target file/state 변경이 들어오면 affected files 또는 state를 다시 읽고 reconcile한 뒤 가장 이른 안전한 phase부터 재개한다.
- status/progress check를 처리한 뒤에도 진행 상태가 바뀌었으면 active flow record의 phase, blocker, required next action을 갱신한다.

## 검토 질문

- 이 응답은 loop continuation, active question-routing, explicit user stop handling 중 하나로 끝나는가?
- activation-only 요청을 work 완료 요약으로 닫지 않고 scope 또는 next-flow 질문으로 열었는가?
- explicit stop이 아닌 사용자 개입을 중단이나 승인 경계로 오해하지 않았는가?
