# turn-gate 스킬 스펙

## 목적

`turn-gate`는 사용자가 명시적으로 현재 턴을 끝내기 전까지 활성 Codex 턴을 열린 상태로 유지합니다.
각 active flow에 필수 `flow` 계약을 적용하고, session record를 갱신하며, 보고 전에 검증을 처리하고, 보고 뒤에는 다음 flow 선택을 다시 엽니다.

이 폴더에서는 `intent.md`가 사용자 의도 기록을 소유하고, `intent-scenarios/`가 회귀 의도 fixture를 소유합니다. 지속 실행 계약은 `contracts/`가 소유합니다.

`turn-gate`는 workflow taxonomy가 아니며 두 번째 구현 planner도 아닙니다. 책임은 더 좁습니다. 현재 응답을 운영상 열린 상태로 유지하고, 실행되는 항목마다 기록된 flow 계약을 요구하며, 우발적인 terminal closure를 막고, 보고 뒤 필요한 다음 행동을 라우팅합니다.

## 경계

- 포함:

  - 대화 턴 단위 continuity
  - `intake -> framing -> preparation -> work -> verification -> reporting -> next-flow`
  - source-recorded explicit stop 처리
  - session record와 Continuity Guard
  - next-flow question routing과 question abort recovery
  - active flow 도중 들어온 사용자 메시지의 `interruption` routing
  - 위험 기반 verification routing
  - approval-sensitive checkpoint 경계
  - 명시적으로 준비된 overlay로서의 self-drive

- 제외:

  - flow taxonomy, parent flow, sub-flow candidate, readiness, discovery, flow-local strategy 정의
  - 설치 후 runtime reader가 dev-only spec 경로를 읽도록 지시하는 일
  - readiness, verification, self-drive, 이전 맥락을 commit, push, PR, publish, release, version bump 승인으로 취급하는 일
  - 완료된 작업, 성공한 검증, 답변된 질문, 중단된 질문 도구, stale record, final처럼 보이는 문구를 turn closure로 취급하는 일

## 대표 표면

- Runtime skill: `src/loop-kit-dev/skills/turn-gate/SKILL.md`
- Runtime references: `src/loop-kit-dev/skills/turn-gate/references/*.md`
- Runtime templates: `src/loop-kit-dev/skills/turn-gate/templates/*-template.md`
- 사용자 의도: `src/loop-kit-dev/specs/skills/turn-gate/intent.md`
- 회귀 의도 fixture: `src/loop-kit-dev/specs/skills/turn-gate/intent-scenarios/`

## 계약 맵

- `contracts/runtime.md`: active-turn lifecycle, activation, explicit stop, phase prefix, runtime body 경계, approval checkpoint
- `contracts/question-routing.md`: next-flow reopening, structured question 사용, fallback, pending question recovery, `request_user_input` abort 처리
- `contracts/session-records.md`: `000-plan.md`, flow record, self-drive sidecar pointer, raw request, Continuity Guard, recovery case
- `contracts/verification.md`: verification method 선택, result status, clean-context verifier 경계, non-pass routing
- `contracts/interruption.md`: active flow 도중 들어온 사용자 메시지의 entry-only routing, foreground/background/reserved/superseded/blocked 상태 전환
- `contracts/self-drive.md`: prepared sequence overlay, sidecar gate, interruption handling, endpoint 처리, approval boundary

## 소유권 규칙

- 지속 규칙은 정확히 하나의 contract 파일이 소유합니다.
- 다른 contract 파일은 소유 파일을 가리킬 수 있지만, 상세 decision table을 반복하지 않습니다.
- `spec.md`는 index와 boundary map으로 유지하며, operational decision table을 축적하지 않습니다.
- `intent.md`는 사용자 의도를 기록하고 migration step을 기록하지 않습니다.
- `intent-scenarios/`는 회귀 예시를 보관하고 runtime instruction을 보관하지 않습니다.
- Runtime `SKILL.md`는 dev-only spec 경로를 참조하지 않고, 이 계약들을 실행 가능한 지시로 압축해야 합니다.
- Runtime references와 templates에는 설치 후 실제로 존재하는 guidance만 둡니다. 설치 후 실행이 필요한 계약은 `specs/`를 가리키지 말고 `SKILL.md`, `references/`, `templates/`로 흡수합니다.

## 필수 runtime 동작

- activation-only 요청은 terminal activation summary가 아니라 기록된 활성 상태와 next-flow routing을 만들어야 합니다.
- 각 새 flow 시작 지점에서는 현재 flow에 필요한 skill을 다시 읽고, `000-plan.md`에 current/planned flow skill list를 compact하게 남겨야 합니다.
- intake는 raw input과 해석을 분리하고, goal, non-goals, authority-sensitive signal, discovery topic을 드러내야 합니다.
- framing은 필수 `flow` contract를 적용해 active flow, parent flow, sub-flow candidate, phase, handoff를 구분하고 selected flow와 candidate를 혼동하지 않아야 합니다.
- work는 active flow boundary, non-goals, acceptance signal, verification expectation, approval boundary, handoff condition이 알려졌거나 user-gated된 뒤에만 시작할 수 있습니다.
- reporting은 먼저 기록을 갱신한 뒤 changed surfaces, verification status, material judgment calls, residual risk, required next action을 보고해야 합니다.
- reporting 뒤에는 `next-flow`, `blocked`, 유효한 self-drive continuation, source-recorded explicit stop 중 하나로 라우팅해야 합니다.
- active flow 도중 새 사용자 메시지가 들어오면 `interruption`으로 먼저 분류하고, 일반 lifecycle phase 또는 새 foreground flow로 빠져나가야 합니다.
- source-recorded explicit stop만 terminal closure의 근거가 됩니다.

## 검토 질문

- 현재 응답은 현재 source-recorded explicit stop이 없는 한 계속 열린 상태인가?
- active flow는 필수 `flow` decision 또는 충분히 기록된 flow contract에 기반하는가?
- 새 flow 시작 때 필요한 skill을 다시 읽고 `000-plan.md` skill list가 최신인지 확인했는가?
- reporting 뒤 next-flow routing, blocker routing, 유효한 self-drive continuation, source-recorded explicit stop 중 하나로 이어졌는가?
- question-tool abort를 turn closure로 보거나 같은 질문 도구 호출을 무작정 반복하지 않고 recoverable routing으로 다루는가?
- active flow 도중 사용자 메시지를 `interruption`으로 분류하고, inline answer와 flow 계약 변경, background 전환, 후속 예약, supersede, blocker, explicit stop을 구분하는가?
- approval-sensitive action은 exact target, effect, risk, recovery path, included/excluded scope, endpoint에 기반하는가?
- raw user message가 중요할 때 summary 또는 interpretation과 분리해 기록하는가?
