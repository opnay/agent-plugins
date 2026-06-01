# turn-gate 스킬 스펙

## 목적

`turn-gate`는 현재 Codex 턴을 명시적 종료 요청 전까지 열린 상태로 유지하는 wrapper입니다.
메인 플로우는 `flow skill`과 `next turn-flow / 메시지 수신`으로 압축합니다.
`flow skill` 내부 의미는 `flow.message -> flow.main-flows -> flow.end`이며, `turn-gate`는 이를 재정의하지 않고 적용합니다.

## 경계

- 포함: active-turn continuity, next turn-flow reopening, session record 적용/복구, question routing, verification routing, self-drive gate, explicit-stop 기록.
- 제외: flow taxonomy, flow lifecycle, readiness/discovery/ambiguity, shared template 의미, handoff 의미, workflow planner, commit/push/PR/release/version bump/destructive action 승인.

## 계약 맵

- `intent.md`: 현재 다이어그램
- `contracts/runtime.md`: wrapper runtime
- `contracts/flow-relationship.md`: `flow` 의존 경계
- `contracts/session-records.md`: record 적용/복구
- `contracts/question-routing.md`: next turn-flow와 질문 복구
- `contracts/interruption.md`: active flow 중 새 메시지 routing
- `contracts/verification.md`: verification method/result routing
- `contracts/self-drive.md`: self-drive 역방향 gate
- `contracts/date-authority.md`: 상대 날짜와 기록 날짜 충돌

## 핵심 계약

- `turn-gate`는 사용자 메시지를 `flow skill`로 보내고, `flow.end` 이후 `next turn-flow / 메시지 수신`을 엽니다.
- 일반 모드는 다음 사용자 메시지를 기다립니다.
- self-drive 모드는 `next turn-flow / 메시지 수신`에서 자체 해석으로 다시 `flow skill`에 들어갈 수 있습니다.
- 종료 요청은 전 과정에서 감지하며, source-recorded explicit stop이 있을 때만 현재 turn을 닫습니다.
- 완료, 검증 통과, 커밋, 보고, final-looking 문구, 질문 중단은 종료 요청이 아닙니다.

## 검토 질문

- `turn-gate`가 `flow` 의미를 재정의하지 않고 wrapper로 적용하는가?
- reporting 뒤 next turn-flow가 열리는가?
- self-drive가 명시된 gate 없이 자동 시작되지 않는가?
- 종료 요청이 source-recorded explicit stop으로만 닫히는가?
- shared record template 의미가 `flow`에 남아 있는가?
