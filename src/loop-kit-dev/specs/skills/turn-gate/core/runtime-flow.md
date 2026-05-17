# turn-gate runtime-flow sub-spec

## 목적

이 문서는 `turn-gate`가 하나의 턴을 어떻게 이어가는지에 대한 전체 phase 흐름과 phase 간 전환 조건을 소유합니다.

## 전체 흐름

`turn-gate`의 기본 flow는 아래 순서를 유지합니다.

1. preparation
2. work
3. verification
4. reporting
5. next-flow

activation과 explicit stop handling은 이 기본 flow를 둘러싼 lifecycle guard입니다.
이 lifecycle guard는 내부 gate로 적용됩니다.
flow shaping gate는 active flow와 completion criteria를 만들거나 갱신하며, task policy gate는 flow 내부 실행 정책을 정합니다.
task policy는 flow 밖의 독립 계층이 아니며, 개별 task 완료가 flow 완료나 turn closure를 결정할 수 없습니다.
verification gate와 reporting gate는 각각 검증 판정과 보고 맥락 정리를 소유합니다.
상세 gate 계약은 `gates/internal-gates.md`가 소유합니다.

deep-interview alignment, flow list design, meaning resolution, current-state inspection, target reread, scope lock, approval boundary 확인은 기본적으로 `preparation` 안의 세부 작업입니다.

`operational-preparation flow`, `change-unit flow`, planned flow boundary, 후속 후보와 active execution flow의 구분은 `core/flow-boundaries.md`가 소유합니다.
이 문서는 phase 순서와 전환 조건만 직접 소유하고, flow taxonomy 판단이 필요하면 `core/flow-boundaries.md`로 위임합니다.

이 문서는 각 phase의 순서와 전환을 소유하고, phase 내부의 세부 판단은 대응 child spec으로 위임합니다.

## Phase Start Message Prefix

`turn-gate`가 사용자에게 phase 시작을 알리거나 phase 시작과 함께 진행 상황을 말할 때, 그 사용자-facing 메시지는 `[<phase-name>(/<phase-protocol>)]` 접두사로 시작해야 합니다.

- canonical phase label은 `preparation`, `work`, `verification`, `reporting`, `next-flow`를 사용합니다.
- `(/<phase-protocol>)` segment는 optional notation이며, phase protocol 사용 시 slash suffix로 표기해야 합니다.
- 예: `[preparation]`, `[work]`, `[verification]`, `[reporting]`, `[next-flow]`, `[preparation/deep-interview]`, `[work/ralph-loop]`, `[verification/review-loop]`, `[reporting/commit-readiness-gate]`
- 실제 출력에는 literal parenthesis를 쓰지 않습니다. `(<...>)`는 optional segment notation입니다.
- 이 규칙은 phase가 시작되는 대화 메시지에 적용합니다.
- 내부 기록 파일, 최종 산출물 본문, command output 요약, 질문 선택지의 모든 문장에 기계적으로 붙이는 규칙은 아닙니다.
- phase 없이 일반 설명만 이어가는 경우에는 prefix를 억지로 붙이지 않습니다.
- activation-only에 concrete task가 없으면 scope 설정을 위해 `[preparation]`을 우선 사용하고, 실제 next-flow 선택지를 여는 메시지에서만 `[next-flow]`로 전환합니다.
- status/progress 질문은 현재 active phase label을 사용합니다. work 중이면 보통 `[work]`이고, 의도적으로 flow context를 요약하는 경우에만 `[reporting]`을 사용합니다.
- self-drive continuation에서는 상태, 검증, 보고, 자동 next-flow handoff처럼 사용자-facing phase/progress를 알리는 메시지에 현재 phase prefix를 붙입니다. 단, `000-self-drive.md`, flow record, 생성 산출물 본문, 질문 선택지 label 안으로 prefix를 전파하지 않습니다.
- session record 접근 blocker는 발견된 phase label을 사용합니다. result reporting 전이면 `[reporting]`, next-flow reopening 전이면 `[next-flow]`입니다.
- report-only evaluation은 편집 없이 evidence를 모으더라도 reporting 뒤 explicit stop이 없으면 `[next-flow]`로 이어집니다.

## Phase 계약

- activation:
  - `turn-gate`가 호출되면 conversation-level first-class operating rule로 활성화한다.
  - concrete task 없이 activation만 요청되면 work로 들어가지 않고 next-flow 또는 scope selection을 연다.
  - activation-only의 첫 사용자-facing 응답은 기본적으로 `[preparation]` scope setup이며, 바로 선택지를 여는 별도 메시지는 `[next-flow]`를 사용할 수 있다.
  - 예: "turn-gate 켜줘", "Use turn-gate", `$loop-kit:turn-gate`만 온 경우에는 activation 완료 요약으로 닫지 않고 다음 scope 또는 next-flow 선택을 연다.
- preparation:
  - 이 단계는 flow shaping gate를 통과해 active flow의 경계와 completion criteria를 정한다.
  - work로 넘어가기 전 intent, scope, non-goal, acceptance signal, verification expectation, approval boundary를 정렬한다.
  - 세부 계약은 `phases/preparation.md`가 소유한다.
- work:
  - 이 단계는 task policy gate를 통과해 현재 flow 내부 실행 정책을 정한다.
  - 사용자가 요청한 실제 작업을 진행한다.
  - implicit default state, phase protocol 선택, local reference 읽기 규칙은 `phase-protocols/routes.md`가 소유한다.
  - 세부 계약은 `phases/work.md`가 소유한다.
- verification:
  - 이 단계는 verification gate를 통과한다.
  - 현재 flow의 work 결과를 검증한다.
  - 검증 packet, pass/fail/blocked/insufficient 처리, non-pass return path는 `records/verification.md`가 소유한다.
  - 세부 계약은 `phases/verification.md`가 소유한다.
- reporting:
  - 이 단계는 reporting gate를 통과한다.
  - 결과 보고는 terminal response가 아니라 다음 flow 진행을 위한 context 정리다.
  - next-flow reopening 세부는 `records/question-routing.md`가 소유한다.
  - 세부 계약은 `phases/reporting.md`가 소유한다.

- next-flow reopening:
  - 이 단계는 `next-flow` phase다.
  - explicit stop이 없다면 결과 보고 뒤 active question-routing으로 다음 flow를 연다.
  - `request_user_input`, fallback, visible/recorded turn-end option은 `records/question-routing.md`가 소유한다.
  - 세부 계약은 `phases/next-flow.md`가 소유한다.
- explicit stop handling:
  - 사용자가 명시적으로 턴 종료를 요청한 경우에만 terminal summary가 가능하다.
  - "여기서 끝", "턴 종료", "이 turn은 그만", "stop the turn"처럼 현재 turn 자체를 끝내려는 의도가 분명한 입력만 explicit turn stop으로 분류한다.
  - 명시적 종료 의도가 불분명하면 종료로 추정하지 말고 continuation input으로 처리하거나 user-gated clarification을 연다.
  - flow record의 `confirmed closure`는 특정 explicit stop 사용자 메시지와 함께 기록된 경우에만 유효하다.
  - closure source message가 없거나 현재 incoming message와 맞지 않는 stale closure 기록은 terminal close 근거로 쓰지 않는다.
  - closure source message와 `Continuity Guard` 기록은 `records/session-records.md`와 함께 유지한다.

## 검토 질문

- 기본 flow가 `준비 -> 작업 -> 검증 -> 보고 -> next-flow`로 한 번에 읽히는가?
- flow taxonomy 판단이 `core/flow-boundaries.md`로 위임되는가?
- phase 세부 계약이 `phases/preparation.md`, `phases/work.md`, `phases/verification.md`, `phases/reporting.md`, `phases/next-flow.md`로 위임되는가?
- 각 phase의 세부 판단이 적절한 child spec으로 위임되는가?
- reporting이 terminal close가 아니라 next-flow reopening으로 이어지는가?
- explicit stop 없이 흐름이 닫히는 경로가 남아 있지 않은가?
