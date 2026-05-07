# Loop Kit Dev 플러그인 스펙

## 플러그인 목적

`loop-kit-dev`은 `turn-gate`를 메인 표면으로 삼는 loop-oriented workflow 플러그인입니다.
핵심 책임은 하나의 턴을 사용자가 턴을 종료하자고 요청할때까지 닫지 않고 유지하면서, 각 flow를 `준비 -> 작업 -> 검증 -> 보고`로 이어가고, 현재 phase의 메인 작업에 맞는 내부 loop mode를 `turn-gate` 안에서 선택해 실행하는 것입니다.
사용자 메시지 기반 준비에서는 deep-interview alignment로 의도를 정렬하고 그 결과를 flow list로 만들며, 이미 선택된 flow의 준비에서는 수정 범위, 현재 상태, 대상 파일 또는 산출물, 검증 조건을 확인합니다.
내부 mode 선택 전에는 사용자 지시어의 operation 의미가 파일, skill, spec, phase, routing rule, release surface 중 무엇을 가리키는지 확인하고, 해석에 따라 작업이 달라지면 meaning resolution 질문으로 먼저 잠급니다.
이 플러그인은 `workflow-kit`의 일반 workflow skill 의미를 참조하되, turn-gate runtime contract와 session continuity는 자체 runtime-oriented surface로 소유합니다.

## 플러그인 경계와 비목표

- 포함:
  - turn-level loop gate contract
  - `preparation -> work -> verification -> reporting -> next-flow question-routing response` 구조 유지
  - 사용자 메시지 기반 준비와 기존 flow 기반 준비의 구분 유지
  - deep-interview alignment와 flow list design을 preparation 세부 작업으로 유지
  - internal mode 선택 전 operation meaning resolution
  - `turn-gate` 내부의 loop mode 선택
  - user-gated question routing 유지
  - 읽기 전용 bounded verifier subagent를 통한 clean-context verification
  - autonomous subagent question routing을 위한 별도 overlay skill 제공
  - `turn-gate/references/` 아래 local absorbed loop contract 유지
  - discovery, autonomous execution, refinement, review, readiness 성격의 current-phase work를 loop 안에서 처리
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 사용자에게 여러 loop skill을 직접 노출하는 구조
  - domain-specific implementation guidance
  - turn continuity가 필요 없는 일반 단발성 응답

## 처리하려는 작업 형태

- 결과 보고 뒤에도 같은 턴에서 다음 플로우를 계속 이어가야 하는 작업
- 사용자 메시지에서 이후 flow list를 도출해야 하는 작업
- 이미 선택된 flow에서 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인해야 하는 작업
- 사용자 지시어가 여러 구조 단위를 가리켜 current-phase work를 고르기 전에 의미를 잠가야 하는 작업
- 현재 phase의 작업이 requirement discovery, autonomous execution, refinement, review handling, readiness pass 중 하나로 좁혀지는 작업
- loop continuity가 top-level governing contract인 작업

## 대표 표면

- 대표 실행 표면: `turn-gate`
- 대표 스펙: `loop-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `loop-kit-dev/specs/skills/*.md` 또는 복잡한 skill의 `loop-kit-dev/specs/skills/<skill-name>/spec.md`
- turn-gate local references: `loop-kit-dev/skills/turn-gate/references/*.md`
- referenced workflow skill specs: `src/workflow-kit-dev/specs/skills/deep-interview.md`, `src/workflow-kit-dev/specs/skills/autopilot.md`, `src/workflow-kit-dev/specs/skills/ralph-loop.md`, `src/workflow-kit-dev/specs/skills/review-loop.md`, `src/workflow-kit-dev/specs/skills/commit-readiness-gate.md`

## 내장 skill 체계

- `turn-gate`: turn continuity를 유지하고 각 flow를 `준비 -> 작업 -> 검증 -> 보고`로 진행하며 current-phase work에 맞는 내부 loop mode를 고른다.
  - spec: `loop-kit-dev/specs/skills/turn-gate/spec.md`
- `turn-gate-self-drive`: `turn-gate`를 base contract로 적용하고, bounded decision을 subagent question packet으로 라우팅해 자동 진행한다.
  - spec: `loop-kit-dev/specs/skills/turn-gate-self-drive.md`

## SDD 운영 원칙

- `workflow-kit`은 일반 workflow skill 의미를 제공한다.
- `loop-kit-dev`은 turn-gate 사용자 표면, runtime loop orchestration, session continuity contract를 소유한다.
- `turn-gate`는 internal mode를 local `references/`로 흡수해 사용하되, 그 reference는 관련 workflow skill spec과 동기화해 유지한다.
- 복잡한 skill spec은 `specs/skills/<skill-name>/spec.md`를 기본 index로 두고, 세부 계약은 같은 folder 아래 sub-spec으로 분리할 수 있다.
- `turn-gate`의 필수 운영 도구는 기본적으로 질문 도구 `request_user_input`와 계획 도구 `update_plan`이다.
- `turn-gate`의 clean-context verification은 읽기 전용 bounded verifier subagent 실행을 포함하며, 이 검증 전용 실행은 `turn-gate` 활성 중 사전 허용된 계약으로 취급한다.
- `turn-gate`의 phase model은 `준비 -> 작업 -> 검증 -> 보고`를 런타임 surface에 드러내야 하며, deep-interview alignment, flow list design, meaning resolution, current-state inspection은 preparation 세부 작업으로 설명해야 한다.
- `turn-gate-self-drive`는 self-drive overlay로만 동작하며, base loop gate 자체는 `turn-gate`를 직접 따른다.
- `turn-gate-self-drive` 도중 사용자 메시지가 들어오면 멈추지 않고 현재 플로우 조정 또는 다음 플로우 우선 등록으로 처리한다.
- 새로운 내부 loop mode가 필요하면 해당 workflow skill의 일반 의미를 `workflow-kit`에 정의하거나 갱신한 뒤 `loop-kit-dev`의 runtime reference에 반영한다.
- `loop-kit-dev`에서는 `autopilot`, `ralph-loop`, `review-loop`, `commit-readiness-gate`를 직접 호출 가능한 사용자 엔트리포인트로 늘리지 않는다.
- `turn-gate`의 phase model이나 session continuity rule이 바뀌면 `loop-kit-dev` spec, skill body, manifest prompt를 같은 변경 단위에서 점검한다. Internal mode 의미가 바뀐 경우에만 관련 `workflow-kit` skill spec을 함께 점검한다.

## 현재 구조 메모

- 이 플러그인은 intentionally narrow한 operational package다.
- `turn-gate`가 메인 실행 표면이다.
- 내부 loop mode의 일반 의미는 `src/workflow-kit-dev/specs/skills/deep-interview.md`, `src/workflow-kit-dev/specs/skills/autopilot.md`, `src/workflow-kit-dev/specs/skills/ralph-loop.md`, `src/workflow-kit-dev/specs/skills/review-loop.md`, `src/workflow-kit-dev/specs/skills/commit-readiness-gate.md`를 기준으로 보고, `turn-gate/references/`에는 그 실행용 absorbed contract를 둔다.
- autonomous subagent question routing은 current-phase mode가 아니라 `turn-gate-self-drive` overlay skill의 책임으로 둔다.
