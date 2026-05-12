# Loop Kit Dev 플러그인 스펙

## 플러그인 목적

`loop-kit-dev`은 `turn-gate`를 메인 표면으로 삼는 loop-oriented workflow 플러그인입니다.
핵심 책임은 하나의 턴을 사용자가 턴을 종료하자고 요청할 때까지 닫지 않고 유지하면서, 각 flow를 `준비 -> 작업 -> 검증 -> 보고`로 이어가고, 기본 상태에서 현재 phase에 필요한 protocol을 선택해 실행하는 것입니다.
여기서 flow는 `분석`, `작업`, `커밋` 같은 진행 phase가 아니며, 반드시 최종 사용자에게 직접 보이는 가치 단위도 아닙니다.
flow는 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 응집된 변경 단위입니다.
예를 들어 "로그인 페이지 만들기"라는 큰 요청은 하나의 사용자 가치처럼 보일 수 있지만, planned flow는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립`처럼 커밋 단위로도 나뉠 수 있는 변경 묶음이어야 합니다.
사용자 메시지 기반 준비에서는 deep-interview alignment로 의도를 정렬하고 그 결과를 flow list로 만들며, 이미 선택된 flow의 준비에서는 수정 범위, 현재 상태, 대상 파일 또는 산출물, 검증 조건을 확인합니다.
사용자 메시지를 받아 의도를 해석하고 planned flow list를 만드는 일은 자체 산출물로 plan/session record를 소유하는 `operational-preparation flow`가 될 수 있습니다.
이 운영 flow가 만든 planned flow list의 각 항목은 실제 코드, 문서, fixture, 설정 같은 산출물을 소유하는 `change-unit flow`여야 합니다.
사용자 메시지 기반 준비는 planned flow list 전체를 실행하는 데 필요한 정보와 예상 위험 작업, approval boundary를 먼저 질문해 수집해야 합니다.
`self-drive`는 이 기본 준비와 loop surface를 수정하지 않고, 명시적으로 적용될 때 자기 계약으로 준비된 sequence의 진행 판단을 덮어쓰는 별도 overlay입니다.
commit-readiness reporting 자체는 산출물 변경을 소유하지 않는 한 planned flow boundary가 아닙니다.
phase protocol 선택 전에는 사용자 지시어의 operation 의미가 파일, skill, spec, phase, routing rule, release surface 중 무엇을 가리키는지 확인하고, 해석에 따라 작업이 달라지면 meaning resolution 질문으로 먼저 잠급니다.
이 플러그인은 `workflow-kit`의 일반 workflow skill 의미를 참조하되, turn-gate runtime contract와 session continuity는 자체 runtime-oriented surface로 소유합니다.

## 플러그인 경계와 비목표

- 포함:
  - turn-level loop gate contract
  - `preparation -> work -> verification -> reporting -> next-flow question-routing response` 구조 유지
  - flow를 phase나 direct user-value가 아니라 cohesive reviewable or commit-sized change unit으로 나누는 계약
  - 사용자 메시지 해석과 planned flow list 설계를 운영 flow로 기록하고, 그 결과 change-unit planned flows를 분리하는 계약
  - 사용자 메시지 기반 준비와 기존 flow 기반 준비의 구분 유지
  - deep-interview alignment와 flow list design을 preparation 세부 작업으로 유지
  - 사용자 메시지 기반 준비에서 planned flow list 전체에 필요한 정보, 예상 위험 작업, user-gated checkpoint 수집
  - prepared planned flow list에 적용될 수 있는 별도 self-drive overlay reference 제공
  - phase protocol 선택 전 operation meaning resolution
  - `turn-gate`의 독립적인 implicit default state 유지
  - user-gated question routing 유지
  - 읽기 전용 bounded verifier subagent를 통한 clean-context verification
  - autonomous subagent question routing을 위한 self-drive runtime reference 제공
  - `turn-gate/references/` 아래 local absorbed loop contract 유지
  - discovery, autonomous execution, refinement, review, readiness 성격의 current-phase work를 loop 안에서 처리
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 사용자에게 여러 loop skill을 직접 노출하는 구조
  - domain-specific implementation guidance
  - turn continuity가 필요 없는 일반 단발성 응답

## 처리하려는 작업 형태

- 결과 보고 뒤에도 같은 턴에서 다음 플로우를 계속 이어가야 하는 작업
- 사용자 메시지에서 이후 flow list를 도출해야 하는 작업. 이때 flow는 phase list가 아니라 검토/검증/커밋 가능한 변경 단위 list다.
- 사용자 메시지 해석, scope lock, approval boundary 정리, planned flow list 작성 자체가 plan/session record 산출물을 만드는 운영 flow로 남아야 하는 작업
- 이미 선택된 flow에서 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인해야 하는 작업
- 사용자 지시어가 여러 구조 단위를 가리켜 current-phase work를 고르기 전에 의미를 잠가야 하는 작업
- 현재 phase의 작업이 requirement discovery, refinement, review handling, readiness pass 같은 phase protocol 중 하나로 좁혀지는 작업
- loop continuity가 top-level governing contract인 작업

## 대표 표면

- 대표 실행 표면: `turn-gate`
- 대표 스펙: `loop-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `loop-kit-dev/specs/skills/*.md` 또는 복잡한 skill의 `loop-kit-dev/specs/skills/<skill-name>/spec.md`
- turn-gate local references: `loop-kit-dev/skills/turn-gate/references/*.md`
- referenced workflow skill specs: `src/workflow-kit-dev/specs/skills/deep-interview.md`, `src/workflow-kit-dev/specs/skills/autopilot.md`, `src/workflow-kit-dev/specs/skills/ralph-loop.md`, `src/workflow-kit-dev/specs/skills/review-loop.md`, `src/workflow-kit-dev/specs/skills/commit-readiness-gate.md`

## 내장 skill 체계

- `turn-gate`: turn continuity를 유지하고 각 flow를 `준비 -> 작업 -> 검증 -> 보고`로 진행하며 기본 상태의 phase protocol을 적용한다.
  - spec: `loop-kit-dev/specs/skills/turn-gate/spec.md`

## SDD 운영 원칙

- `workflow-kit`은 일반 workflow skill 의미를 제공한다.
- `loop-kit-dev`은 turn-gate 사용자 표면, runtime loop orchestration, session continuity contract를 소유한다.
- `turn-gate`는 runtime phase protocol을 local `references/`로 흡수해 사용하되, 그 reference는 관련 workflow skill spec과 동기화해 유지한다.
- 복잡한 skill spec은 `specs/skills/<skill-name>/spec.md`를 기본 index로 두고, 세부 계약은 같은 folder 아래 sub-spec으로 분리할 수 있다.
- `turn-gate`의 phase protocol 선택 기준과 routing은 `specs/skills/turn-gate/phase-protocols/routes.md`가 소유하고, phase protocol 상세 계약은 같은 폴더의 개별 protocol spec이 소유한다.
- `turn-gate`의 필수 운영 도구는 기본적으로 질문 도구 `request_user_input`와 계획 도구 `update_plan`이다.
- `turn-gate`의 clean-context verification은 읽기 전용 bounded verifier subagent 실행을 포함하며, 이 검증 전용 실행은 `turn-gate` 활성 중 사전 허용된 계약으로 취급한다.
- `turn-gate`의 phase model은 `준비 -> 작업 -> 검증 -> 보고`를 런타임 surface에 드러내야 하며, deep-interview alignment, flow list design, meaning resolution, current-state inspection은 preparation 세부 작업으로 설명해야 한다.
- 사용자 메시지에서 시작하는 bootstrap은 `operational-preparation flow`로 기록할 수 있으며, 이 flow의 산출물은 session plan, flow list, scope/approval boundary다. 이 결과로 생성되는 product/work planned flows는 `change-unit flow`로 분리한다.
- self-drive는 별도 skill 표면이 아니라 명시적으로 적용될 때 준비된 sequence의 진행 판단을 덮어쓰는 독립 overlay reference로 동작한다.
- self-drive reference는 초기 preparation을 대신하지 않고, self-drive mode 계약을 runtime에서 실행 가능한 형태로 적용한다.
- self-drive의 autonomous continuation과 handoff behavior는 `modes/self-drive.md`가 소유한다.
- self-drive 도중 사용자 메시지가 들어오면 멈추지 않고 현재 플로우 조정 또는 다음 플로우 우선 등록으로 처리한다.
- 새로운 phase protocol이 필요하면 해당 workflow skill의 일반 의미를 `workflow-kit`에 정의하거나 갱신한 뒤 `loop-kit-dev`의 runtime reference에 반영한다.
- 새로운 explicit overlay는 기본 상태나 기존 overlay로 current flow를 소유할 수 없을 때만 추가한다.
- `loop-kit-dev`에서는 `autopilot`, `ralph-loop`, `review-loop`, `commit-readiness-gate`를 직접 호출 가능한 사용자 엔트리포인트로 늘리지 않는다.
- `turn-gate`의 phase model이나 session continuity rule이 바뀌면 `loop-kit-dev` spec, skill body, manifest prompt를 같은 변경 단위에서 점검한다. Phase protocol 의미가 바뀐 경우에만 관련 `workflow-kit` skill spec을 함께 점검한다.

## 현재 구조 메모

- 이 플러그인은 intentionally narrow한 operational package다.
- `turn-gate`가 메인 실행 표면이다.
- deep-interview, autopilot, ralph-loop, review-loop, commit-readiness-gate의 일반 의미는 mode가 아니라 phase protocol로 보고, `turn-gate/references/`에는 그 실행용 absorbed contract를 둔다.
- autonomous subagent question routing은 direct skill entrypoint가 아니라 self-drive runtime reference의 책임으로 둔다.
