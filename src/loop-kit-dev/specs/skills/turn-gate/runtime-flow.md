# turn-gate runtime-flow sub-spec

## 목적

이 문서는 `turn-gate`가 하나의 턴을 어떻게 이어가는지에 대한 전체 phase 흐름과 phase 간 전환 조건을 소유합니다.

## 전체 흐름

`turn-gate`의 기본 flow는 아래 순서를 유지합니다.

1. preparation
2. work
3. verification
4. reporting

activation, incoming message classification, next-flow reopening, explicit stop handling은 이 기본 flow를 둘러싼 lifecycle guard입니다.
deep-interview alignment, flow list design, meaning resolution, current-state inspection, target reread, scope lock, approval boundary 확인은 기본적으로 `preparation` 안의 세부 작업입니다.

이 문서는 각 phase의 순서와 전환을 소유하고, phase 내부의 세부 판단은 대응 child spec으로 위임합니다.

## Phase 계약

- activation:
  - `turn-gate`가 호출되면 conversation-level first-class operating rule로 활성화한다.
  - concrete task 없이 activation만 요청되면 work mode를 고르지 않고 next-flow 또는 scope selection을 연다.
  - 예: "turn-gate 켜줘", "Use turn-gate", `$loop-kit:turn-gate`만 온 경우에는 activation 완료 요약으로 닫지 않고 다음 scope 또는 next-flow 선택을 연다.
- incoming message classification:
  - 모든 incoming message를 같은 loop-gated turn의 authoritative input으로 본다.
  - explicit turn stop, status/progress check, current-flow correction, current-flow priority change, next-flow priority request 중 하나로 분류한다.
  - status/progress check라면 현재 phase, blocker 또는 progress, 다음 concrete action을 짧게 보고한 뒤 active flow를 계속한다.
  - current-flow correction 또는 current-flow priority change라면 현재 analysis/plan을 즉시 조정하고 가장 이른 안전한 phase부터 이어간다.
  - correction이 target file, artifact, state를 바꾸면 이어가기 전에 해당 target을 다시 읽고 stale assumption을 재사용하지 않는다.
  - next-flow priority request라면 flow record의 next-flow 후보 중 최우선으로 등록하고 다음 safe handoff point까지 이어간다.
  - "status?", "진행 상황은?", "아니 그 파일 말고", "먼저 X", "다음엔 commit readiness 봐줘" 같은 입력은 explicit stop이 아니라 status/progress check, current-flow correction, current-flow priority change, next-flow priority request로 분류한다.
- preparation:
  - 이 flow에서 무엇을 할지, 왜 하는지, 어떤 조건에서 작업으로 넘어갈 수 있는지를 준비한다.
  - 사용자 메시지에서 시작하는 preparation은 deep-interview를 사용해 intent, scope, non-goal, success criteria, approval boundary, verification signal을 정렬한다.
  - 사용자 메시지 기반 deep-interview 결과는 단순 질문 답변이 아니라 이후 flow list로 변환되어야 한다.
  - 비 사용자 메시지에서 시작하는 preparation은 이미 준비된 flow의 실행 전 준비이며, 필요한 수정 범위, 현재 상태, 대상 파일, stale assumption, 실행 전 조건을 확인한다.
  - operation/target ambiguity가 flow list나 작업 결과를 바꿀 수 있으면 flow list design이나 work 전에 meaning resolution으로 먼저 잠근다.
  - requested intent, requested action, current blocker, likely internal mode, approval boundary, deep-interview result 또는 current-state preparation result, planned flow list를 확인한다.
  - operation/target ambiguity는 `meaning-resolution.md`가 소유한다.
  - destructive, irreversible, external, commit/publish approval boundary는 `approval-boundary.md`가 소유한다.
  - 현재 flow의 active steps를 정하고, meaningful work가 시작되면 계획 도구를 사용해 현재 상태를 유지한다.
  - multi-flow decomposition과 session plan은 `session-records.md`가 소유한다.
- work:
  - 사용자가 요청한 실제 작업을 진행한다.
  - 작업은 파일 수정, 조사, 검증 실행, 리뷰 finding 처리, 계획 작성처럼 다양한 형태일 수 있다.
  - work에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
  - mode selection과 local reference 읽기 규칙은 `mode-selection.md`가 소유한다.
- verification:
  - 현재 flow의 work 결과를 검증한다.
  - 파일 수정이라면 적용 여부, 타입 오류, 테스트/빌드/린트 같은 해당 검증 경로를 확인한다.
  - 조사나 판단 작업이라면 다양한 관점에서 논리 비판과 반례 검토를 수행한다.
  - work 뒤 reporting 전에 clean-context verification을 수행한다.
  - 검증 packet, pass/fail/blocked/insufficient 처리, non-pass return path는 `verification.md`가 소유한다.
- reporting:
  - 결과 보고는 terminal response가 아니라 다음 flow 진행을 위한 context 정리다.
  - 이번 flow에서 무엇을 준비/작업/검증했는지, 남은 불확실성과 blocker가 무엇인지 정리한다.
  - residual uncertainty와 blocker가 있으면 사용자에게 보이게 포함한다.
  - 계획된 flow가 소진되면 질문 도구를 사용해 사용자에게 다음 flow나 작업을 받는다.
  - next-flow reopening 세부는 `question-routing.md`가 소유한다.
- next-flow reopening:
  - explicit stop이 없다면 결과 보고 뒤 active question-routing으로 다음 flow를 연다.
  - `request_user_input`, fallback, visible/recorded turn-end option은 `question-routing.md`가 소유한다.
- explicit stop handling:
  - 사용자가 명시적으로 턴 종료를 요청한 경우에만 terminal summary가 가능하다.
  - "여기서 끝", "턴 종료", "이 turn은 그만", "stop the turn"처럼 현재 turn 자체를 끝내려는 의도가 분명한 입력만 explicit turn stop으로 분류한다.
  - flow record의 `confirmed closure`는 특정 explicit stop 사용자 메시지와 함께 기록된 경우에만 유효하다.
  - closure source message가 없거나 현재 incoming message와 맞지 않는 stale closure 기록은 terminal close 근거로 쓰지 않는다.
  - closure source message와 `Continuity Guard` 기록은 `session-records.md`와 함께 유지한다.

## 검토 질문

- 기본 flow가 `준비 -> 작업 -> 검증 -> 보고`로 한 번에 읽히는가?
- 각 phase의 세부 판단이 적절한 child spec으로 위임되는가?
- reporting이 terminal close가 아니라 next-flow reopening으로 이어지는가?
- explicit stop 없이 흐름이 닫히는 경로가 남아 있지 않은가?
