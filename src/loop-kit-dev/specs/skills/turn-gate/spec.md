# turn-gate 스킬 스펙

## 목적

`turn-gate`는 사용자가 명시적으로 턴을 끝내기 전까지 현재 Codex 턴을 열린 상태로 유지합니다.
이 스킬은 `flow`가 판단한 계약을 적용하고, session record를 갱신하며, 검증과 보고 뒤 다음 행동을 라우팅합니다.

`turn-gate`는 workflow taxonomy나 구현 planner가 아닙니다.
flow identity, readiness, ambiguity, contract impact, handoff 판단은 `flow`가 소유합니다.

## 포함 범위

- active-turn continuity와 explicit-stop 처리
- `intake -> framing -> preparation -> work -> verification -> reporting -> next-flow` 운영
- session record와 Continuity Guard
- active flow 도중 사용자 메시지의 entry-only interruption routing
- reporting 뒤 next-flow, blocker, question recovery routing
- verification method/status routing
- 승인 민감 작업 guardrail
- 명시적으로 준비된 self-drive overlay 적용

## 제외 범위

- flow taxonomy, parent/sub-flow, readiness, discovery, ambiguity, handoff 정의
- dev-only `specs/`를 설치 후 runtime 지시로 남기는 일
- readiness, 검증, 이전 맥락, self-drive를 commit/push/PR/release/version bump/destructive action 승인으로 취급하는 일
- completed work, passed checks, answered questions, interrupted question tool, final-looking wording을 턴 종료로 취급하는 일

## 대표 표면

- Runtime skill: `src/loop-kit-dev/skills/turn-gate/SKILL.md`
- Runtime references/templates: `src/loop-kit-dev/skills/turn-gate/references/`, `templates/`
- User intent: `src/loop-kit-dev/specs/skills/turn-gate/intent.md`
- Regression intent fixtures: `src/loop-kit-dev/specs/skills/turn-gate/intent-scenarios/`

## 계약 맵

- `contracts/runtime.md`: lifecycle, activation, phase prefix, runtime body boundary
- `contracts/date-authority.md`: relative date와 record-date 충돌
- `contracts/question-routing.md`: next-flow, post-flow continue, question abort recovery
- `contracts/session-records.md`: records, raw request, Continuity Guard, recovery state
- `contracts/verification.md`: method/status, clean-context, non-pass routing
- `contracts/interruption.md`: active-flow 중 새 사용자 메시지 routing
- `contracts/self-drive.md`: prepared sequence overlay와 sidecar gate

## 소유권 규칙

각 지속 규칙은 하나의 contract 파일이 소유합니다.
`spec.md`는 색인과 경계만 소유하고 operational decision table을 반복하지 않습니다.
Runtime `SKILL.md`는 설치 후 실제로 존재하는 `SKILL.md`, `references/`, `templates/`만 의존해야 합니다.

## 필수 runtime 동작

- activation-only 요청도 기록된 active state와 next-flow routing을 만들어야 합니다.
- 새 flow 시작 지점에서는 필요한 skill을 다시 읽고 `000-plan.md`의 active skill list를 갱신해야 합니다.
- work 전에는 `flow`가 산출한 scope, non-goals, acceptance, verification, approval, handoff 계약을 적용해야 합니다.
- reporting 전에는 record를 갱신하고, reporting 뒤에는 next-flow, blocker, valid self-drive continuation, explicit-stop 중 하나로 라우팅해야 합니다.
- active flow 도중 새 사용자 메시지가 오면 current phase를 보존하고 `flow` contract-impact를 적용해야 합니다.
- source-recorded explicit stop만 terminal closure 근거가 됩니다.

## 검토 질문

- 현재 응답은 explicit stop 없이 열린 상태로 유지되는가?
- `turn-gate`가 flow 판단을 재정의하지 않고 적용만 하는가?
- 보고 뒤 next action 또는 blocker가 기록과 사용자-facing routing에 남아 있는가?
- question abort, completed checks, final-looking wording이 closure로 오해되지 않는가?
- 승인 민감 작업은 별도 명시 승인 없이는 실행되지 않는가?
