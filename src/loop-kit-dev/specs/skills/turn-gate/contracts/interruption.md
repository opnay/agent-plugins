# turn-gate interruption 계약

## 소유 범위

active flow가 이미 진행 중일 때 새 사용자 메시지가 도착한 경우의 entry-only routing.

## 계약

`interruption`은 일반 lifecycle phase가 아닙니다. 새 사용자 메시지가 active flow 도중 도착했을 때만 열리고, 메시지를 분류한 뒤 기존 phase, 개정된 phase, background 상태, 예약된 후보, 새 foreground flow, blocker, explicit stop 중 하나로 빠져나갑니다.

`interruption`이 시작되면 현재 foreground flow의 phase, scope, non-goals, approval boundary, verification status, required next action을 보존합니다. 그런 다음 새 메시지가 active flow의 계약 또는 turn-level routing에 미치는 영향을 판단합니다.

결과는 하나만 선택합니다.

- `inline-answer`: 현재 flow 계약을 바꾸지 않는 질문입니다. 답변 후 보존된 phase로 돌아갑니다.
- `current-flow-revision`: scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 바꾸는 메시지입니다. 현재 flow record를 갱신하고 `framing` 또는 `preparation`으로 돌아갑니다.
- `background-current-flow`: 현재 flow를 유지하되 다른 foreground flow를 먼저 처리해야 합니다. 현재 flow를 background로 기록하고 새 flow를 시작합니다.
- `reserve-later-analysis`: 지금 처리하지 않을 관련 주제입니다. `000-plan.md` 또는 적절한 routing 표면에 future candidate로 예약하고 보존된 phase로 돌아갑니다.
- `supersede-current-flow`: 사용자가 현재 flow를 취소하거나 새 요청으로 대체합니다. 현재 flow를 superseded로 기록하고 새 flow를 시작합니다.
- `blocker-question`: 결정, 승인, 접근, scope gap 없이는 계속하면 위험합니다. 현재 flow를 blocked로 두고 user-gated question을 엽니다.
- `explicit-stop`: 사용자가 턴 종료를 명시합니다. closure source를 기록한 뒤에만 terminal closure를 허용합니다.

`interruption` 결과는 active flow 계약 밖의 work 권한을 만들지 않습니다. 특히 논의 flow는 구현 flow로 자동 전환되지 않고, commit, push, PR, publish, release, version bump, destructive action은 별도 명시 승인 없이는 실행할 수 없습니다.

짧은 자연어 지시는 효과 기준으로 분류합니다.

- "요약만", "상태만", "왜 멈췄어"처럼 flow 계약을 바꾸지 않는 질문이나 보고 방식 제한은 `inline-answer` 또는 reporting constraint입니다.
- "나중에 봐", "기억해"처럼 현재 scope를 바꾸지 않는 관련 주제는 `reserve-later-analysis`입니다.
- "계속", "알아서 계속"은 기존 contract 안의 진행 허용일 수 있지만 self-drive 활성화 근거는 아닙니다.
- 같은 말이 scope, endpoint, acceptance, verification expectation, approval boundary를 바꾸면 `current-flow-revision`으로 승격합니다.
- 현재 flow를 대체하는 새 작업이면 `supersede-current-flow` 또는 `background-current-flow`를 선택합니다.

## 검토 기준

- 새 사용자 메시지가 active flow 도중 들어온 메시지인가, 새 flow 시작 메시지인가?
- 현재 foreground flow의 복귀 phase와 required next action을 보존했는가?
- 질문 답변이 flow 계약을 바꾸지 않는 `inline-answer`인지 확인했는가?
- 정정이나 금지가 있으면 `current-flow-revision`으로 계약을 먼저 갱신했는가?
- 현재 flow를 background로 둘지, 후속 후보로 예약할지, supersede할지 구분했는가?
- `interruption`을 일반 lifecycle phase처럼 오래 유지하거나 work/verification/reporting 대체물로 사용하지 않았는가?
