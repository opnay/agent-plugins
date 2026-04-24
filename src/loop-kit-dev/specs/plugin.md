# Loop Kit Dev 플러그인 스펙

## 플러그인 목적

`loop-kit-dev`은 `turn-gate`를 메인 표면으로 삼는 loop-oriented workflow 플러그인입니다.
핵심 책임은 하나의 턴을 사용자가 턴을 종료하자고 요청할때까지 닫지 않고 유지하면서, 현재 phase의 메인 작업에 맞는 내부 loop mode를 `turn-gate` 안에서 선택해 실행하는 것입니다.
이 플러그인은 `workflow-kit`이 정의한 broader workflow taxonomy와 canonical loop-mode contract를 runtime-oriented surface로 묶어 제공합니다.

## 플러그인 경계와 비목표

- 포함:
  - turn-level loop gate contract
  - `analysis -> plan -> work -> verification -> result reporting -> next-flow question-routing response` 구조 유지
  - `turn-gate` 내부의 loop mode 선택
  - user-gated 또는 self-drive question-routing mode 선택
  - `turn-gate/references/` 아래 local absorbed loop contract 유지
  - discovery, autonomous execution, refinement, review, readiness 성격의 current-phase work를 loop 안에서 처리
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 사용자에게 여러 loop skill을 직접 노출하는 구조
  - domain-specific implementation guidance
  - turn continuity가 필요 없는 일반 단발성 응답

## 처리하려는 작업 형태

- 결과 보고 뒤에도 같은 턴에서 다음 플로우를 계속 이어가야 하는 작업
- 현재 phase의 작업이 requirement discovery, autonomous execution, refinement, review handling, readiness pass 중 하나로 좁혀지는 작업
- loop continuity가 top-level governing contract인 작업

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `loop-kit-dev-guide`
- 대표 실행 표면: `turn-gate`
- 대표 스펙: `loop-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `loop-kit-dev/specs/skills/*.md`
- turn-gate local references: `loop-kit-dev/skills/turn-gate/references/*.md`
- canonical upstream SSOT: `workflow-kit/specs/plugin.md`, `workflow-kit/specs/skills/*.md`

## 내장 skill 체계

- `loop-kit-dev-guide`: `loop-kit-dev`가 현재 작업의 적절한 시작점인지 판단하고 `turn-gate`로 진입시킨다.
  - spec: `loop-kit-dev/specs/skills/loop-kit-dev-guide.md`
- `turn-gate`: turn continuity를 유지하고 current-phase work에 맞는 내부 loop mode를 고른다.
  - spec: `loop-kit-dev/specs/skills/turn-gate.md`

## SDD 운영 원칙

- `workflow-kit`이 broader workflow map과 canonical loop contract의 SSOT를 계속 소유한다.
- `loop-kit-dev`은 사용자 표면과 runtime loop orchestration만 소유한다.
- `turn-gate`는 internal mode와 question-routing mode를 local `references/`로 흡수해 사용하되, 그 reference는 upstream SSOT와 동기화해 유지한다.
- `turn-gate`의 필수 운영 도구는 기본적으로 질문 도구 `request_user_input`와 계획 도구 `update_plan`이다.
- `self-drive` question-routing mode가 활성화되면 사용자 질문 대신 subagent 질문으로 결정 입력을 얻고, 명시적 승인 경계가 필요한 경우에는 사용자 승인 경계를 유지한다.
- `self-drive` 도중 사용자 메시지가 들어오면 멈추지 않고 현재 플로우 조정 또는 다음 플로우 우선 등록으로 처리한다.
- 새로운 내부 loop mode가 필요하면 먼저 `workflow-kit`의 canonical contract를 정의하거나 갱신한 뒤 `loop-kit-dev`에 반영한다.
- `loop-kit-dev`에서는 `autopilot`, `ralph-loop`, `review-loop`, `commit-readiness-gate`를 직접 호출 가능한 사용자 엔트리포인트로 늘리지 않는다.
- `turn-gate`의 phase model이나 internal mode selection rule이 바뀌면 `workflow-kit` upstream spec과 `loop-kit-dev` spec을 같은 변경 단위에서 점검한다.

## 현재 구조 메모

- 이 플러그인은 intentionally narrow한 operational package다.
- `turn-gate`가 메인 실행 표면이고, `loop-kit-dev-guide`는 진입 분류만 담당한다.
- 내부 loop mode의 canonical 의미는 `workflow-kit/specs/skills/deep-interview.md`, `workflow-kit/specs/skills/autopilot.md`, `workflow-kit/specs/skills/ralph-loop.md`, `workflow-kit/specs/skills/review-loop.md`, `workflow-kit/specs/skills/commit-readiness-gate.md`를 기준으로 보고, `turn-gate/references/`에는 그 실행용 absorbed contract를 둔다.
- `self-drive`는 current-phase mode가 아니라 질문 대상 축의 mode로 보고, `turn-gate/references/self-drive.md`에 실행용 absorbed contract를 둔다.
