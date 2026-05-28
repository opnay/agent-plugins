# Loop Kit Dev 플러그인 스펙

## 플러그인 목적

`loop-kit-dev`은 `flow`와 `turn-gate`를 함께 제공하는 loop-oriented workflow 플러그인입니다.
`flow`는 하나의 메시지나 동작을 flow로 해석하고, 필요하면 finite `sub-flow candidates`로 나누는 흐름 규칙을 소유합니다.
`turn-gate`는 하나의 턴을 사용자가 턴을 종료하자고 요청할 때까지 닫지 않고 유지하면서, 현재 턴의 작업이 `flow` 계약을 통과하도록 강제합니다.
여기서 flow는 `분석`, `작업`, `커밋` 같은 진행 phase가 아니며, 반드시 최종 사용자에게 직접 보이는 가치 단위도 아닙니다.
flow는 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 응집된 작업 흐름 단위입니다.
예를 들어 "로그인 페이지 만들기"라는 큰 요청은 하나의 사용자 가치처럼 보일 수 있지만, planned flow는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립`처럼 커밋 단위로도 나뉠 수 있는 변경 묶음이어야 합니다.
초기 요청은 `flow`의 intake에서 raw input과 해석, goal, non-goal, authority-sensitive signal, discovery topic을 정리합니다. 그다음 framing에서 flow 후보와 산출물 소유권을 설계하고, 이미 선택된 flow의 preparation에서는 수정 범위, 현재 상태, 대상 파일 또는 산출물, 검증 조건을 확인합니다.
초기 의도 정렬과 sub-flow 후보 설계는 자체 산출물로 plan/session record를 소유하는 `operational-preparation flow`가 될 수 있습니다.
이 운영 flow가 만든 sub-flow 후보의 각 항목은 실제 코드, 문서, fixture, 설정 같은 산출물을 소유하는 `change-unit flow`가 될 수 있습니다.
intake와 framing은 sub-flow 후보 또는 flow sequence를 실행하는 데 필요한 정보와 예상 위험 작업, approval boundary를 먼저 질문해 수집해야 합니다.
`self-drive`는 이 기본 준비와 loop surface를 수정하지 않고, 명시적으로 적용될 때 자기 계약으로 준비된 sequence의 진행 판단을 덮어쓰는 별도 overlay입니다.
commit-readiness reporting 자체는 산출물 변경을 소유하지 않는 한 planned flow boundary가 아닙니다.
work 전에는 사용자 지시어의 operation 의미가 파일, skill, spec, flow contract, routing rule, release surface 중 무엇을 가리키는지 확인하고, 해석에 따라 작업이 달라지면 meaning resolution 질문으로 먼저 잠급니다.
이 플러그인은 `workflow-kit`의 일반 workflow skill 의미를 참조하되, turn-gate runtime contract와 session continuity는 자체 runtime-oriented surface로 소유합니다.

## 플러그인 경계와 비목표

- 포함:
  - flow boundary, parent flow, finite sub-flow candidate contract
  - turn-level loop gate contract
  - `intake -> framing -> preparation -> work -> verification -> reporting -> next-flow question-routing response` 구조 유지
  - flow를 phase나 direct user-value가 아니라 cohesive reviewable or commit-sized work unit으로 나누는 계약
  - verification, reporting, evidence repair, blocker recovery를 별도 산출물 없는 flow로 과분해하지 않는 계약
  - 초기 의도 정렬과 sub-flow 후보 설계를 운영 flow로 기록하고, 그 결과 change-unit flow 후보를 분리하는 계약
  - intake, framing, preparation의 구분 유지
  - `flow` intent-first discovery, sub-flow candidate design, selected-flow readiness를 각각 intake/framing/preparation 계약으로 유지
  - intake/framing에서 flow sequence 전체에 필요한 정보, 예상 위험 작업, user-gated checkpoint 수집
  - prepared flow sequence에 적용될 수 있는 별도 self-drive overlay reference 제공
  - work 전 operation meaning resolution
  - `turn-gate`의 독립적인 implicit default state 유지
  - user-gated question routing 유지
  - active flow 도중 들어온 사용자 메시지를 `interruption` entry-only routing으로 분류
  - `clean-context`, `normal`, `not-required` verification method 선택과 중복 검증을 피하는 최소 충분 evidence 구성
  - autonomous subagent question routing을 위한 self-drive runtime reference 제공
  - `turn-gate/references/` 아래 local absorbed loop contract 유지
  - discovery, autonomous execution, refinement, review, readiness 성격의 current-phase work를 loop 안에서 처리
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 사용자에게 여러 turn controller를 직접 노출하는 구조
  - domain-specific implementation guidance
  - turn continuity가 필요 없는 일반 단발성 응답

## 처리하려는 작업 형태

- 결과 보고 뒤에도 같은 턴에서 다음 플로우를 계속 이어가야 하는 작업
- 초기 요청에서 이후 sub-flow 후보를 도출해야 하는 작업. 이때 후보는 phase list가 아니라 검토/검증/커밋 가능한 변경 단위 list다.
- 초기 의도 정렬, scope lock, approval boundary 정리, sub-flow 후보 작성 자체가 plan/session record 산출물을 만드는 운영 flow로 남아야 하는 작업
- 이미 선택된 flow에서 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인해야 하는 작업
- 사용자 지시어가 여러 구조 단위를 가리켜 current-phase work를 고르기 전에 의미를 잠가야 하는 작업
- 현재 flow의 작업이 requirement discovery, refinement, review handling, readiness pass 같은 flow-local strategy 중 하나로 좁혀지는 작업
- loop continuity가 top-level governing contract인 작업

## 대표 표면

경로 표기 규칙: `loop-kit-dev/...`는 plugin-root-relative shorthand입니다.
개발 원본의 실제 편집 경로는 `src/loop-kit-dev/...`이고, 루트 `loop-kit/`은 build command 산출 release surface이므로 직접 편집하지 않습니다.
설치된 runtime 문서는 release surface에 실제 포함되는 `references/*`, `templates/*` 같은 상대 경로만 실행 지시로 사용해야 합니다.

- 대표 실행 표면: `turn-gate`
- 대표 flow 표면: `flow`
- 대표 스펙: `loop-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `loop-kit-dev/specs/skills/*.md` 또는 복잡한 skill의 `loop-kit-dev/specs/skills/<skill-name>/spec.md`
- turn-gate local references: `loop-kit-dev/skills/turn-gate/references/*.md`

## 내장 skill 체계

- `flow`: 메시지나 동작을 active flow, parent flow, finite sub-flow candidate, `operational-preparation flow`, `change-unit flow`로 판정하고 intake, framing, selected-flow readiness, flow-vs-phase 경계, discovery, flow-local execution strategy, handoff condition을 소유한다.
  - active flow의 phase 시작/종료 record checkpoint를 산출하고, `000-plan.md`와 active flow record 중 어떤 표면이 최신화돼야 하는지 구분한다.
  - interruption이나 self-drive 중 새 입력이 기존 flow contract를 바꾸는지, 다음 flow identity와 handoff가 유효한지 판단하는 원천 계약을 소유한다.
  - spec: `loop-kit-dev/specs/skills/flow/spec.md`
- `turn-gate`: turn continuity를 유지하고, 현재 턴에서 `flow` 계약을 적용하도록 강제하며, flow reporting 뒤 next-flow 질문을 연다.
  - `flow` 판단을 재정의하지 않고 session record, question routing, explicit stop guard, verification routing, self-drive sidecar gate를 운영한다.
  - spec: `loop-kit-dev/specs/skills/turn-gate/spec.md`

## SDD 운영 원칙

- `workflow-kit`은 일반 workflow skill 의미를 제공한다.
- `loop-kit-dev`은 flow 단위 규칙, turn-gate 사용자 표면, runtime loop orchestration, session continuity contract를 소유한다.
- 복잡한 skill spec은 `specs/skills/<skill-name>/spec.md`를 기본 index로 두고, 세부 계약은 같은 folder 아래 sub-spec으로 분리할 수 있다.
- spec 문서는 현재 유지되는 계약을 짧고 직접적으로 적는다. migration 과정, 작업 지시, 중복 설명은 change spec이나 session record에 둔다.
- child spec은 `소유 범위`, `계약`, `검토 기준`처럼 반복 가능한 작은 구조를 우선 사용하고, runtime 지시와 spec-side fixture를 섞지 않는다.
- `turn-gate`의 필수 운영 도구는 기본적으로 질문 도구 `request_user_input`와 계획 도구 `update_plan`이다.
- `turn-gate`의 verification method는 `clean-context`, `normal`, `not-required`로 구분한다.
- `clean-context`는 읽기 전용 bounded verifier subagent 실행을 포함하며, 이 검증 전용 실행은 `turn-gate` 활성 중 사전 허용된 계약으로 취급한다.
- 파일 변경, release surface, 다중 파일 계약, 실패 이력, 사용자 요청 검증, approval-sensitive action에서는 `clean-context`가 기본값이다.
- `normal`은 낮은 위험 no-edit/read-only work에서 source readback, evidence checklist, command/check, 논리 반례 검토를 기록하는 방법이다.
- `not-required`는 검증할 work output이 없는 routing-only 또는 blocker-before-work 상황에서만 사용하며, reason과 residual uncertainty를 기록해야 한다.
- verification method는 result status가 아니며, `pass`, `fail`, `blocked`, `insufficient` 처리와 섞지 않는다.
- `flow`의 phase model은 `intake -> framing -> preparation -> work -> verification -> reporting`을 런타임 surface에 드러내야 하며, parent flow와 sub-flow candidate의 경계를 직접 설명해야 한다.
- `flow`의 phase model은 각 phase 시작/종료에서 `000-plan.md` 또는 active flow record 갱신이 필요한지 판단하는 checkpoint를 드러내야 한다.
- `flow`는 intake discovery, framing, selected-flow readiness, ambiguity, review-loop, fix-verify-loop, broad-execution, commit-readiness handoff 같은 flow-local strategy를 소유한다.
- `turn-gate`는 flow phase model을 재소유하지 않고, active turn에서 `flow` 계약을 적용하고 flow reporting 뒤 next-flow reopening을 강제한다.
- `turn-gate`는 active flow 도중 들어온 사용자 메시지를 `interruption`으로 먼저 열고, 계약 변경 여부와 새 flow 여부는 `flow`에 의존해 inline answer, current-flow revision, background current flow, reserved later analysis, supersede, blocker, explicit stop 중 하나로 라우팅한다.
- `turn-gate`는 reporting 뒤 `continue`를 post-flow next-action 해석으로 다루며, recorded next action이 충분하지 않으면 next-flow 질문을 유지한다.
- 초기 bootstrap은 `operational-preparation flow`로 기록할 수 있으며, 이 flow의 산출물은 session plan, sub-flow candidate list, scope/approval boundary다. 이 결과로 생성되는 product/work sub-flow candidates는 선택될 때 `change-unit flow`로 분리한다.
- self-drive는 별도 skill 표면이 아니라 명시적으로 적용될 때 준비된 sequence의 진행 판단을 덮어쓰는 독립 overlay reference로 동작한다.
- `turn-gate/contracts/self-drive.md`는 self-drive overlay의 endpoint, stop boundary, execution authority, handoff behavior에 대한 spec-side ownership을 소유한다.
- self-drive는 finite와 infinite 모두에서 another bounded batch, inventory cycle, blocker recovery를 기록 재확인과 endpoint relock 뒤에만 이어간다.
- 설치 후 실행 guidance는 `skills/turn-gate/references/self-drive.md`가 소유하며, runtime 문서는 dev-only spec 경로를 실행 지시로 사용하지 않는다.
- self-drive reference는 intake/framing/preparation을 대신하지 않고, spec-side overlay 계약을 설치 후 실행 가능한 형태로 흡수해 적용한다.
- self-drive 도중 사용자 메시지가 들어오면 멈추지 않고 현재 플로우 조정 또는 다음 플로우 우선 등록으로 처리한다.
- 새로운 explicit overlay는 기본 상태나 기존 overlay로 current flow를 소유할 수 없을 때만 추가한다.
- `loop-kit-dev`에서는 flow-local strategy를 직접 호출 가능한 사용자 엔트리포인트로 늘리지 않는다.
- `turn-gate`의 phase model, verification method, session continuity rule은 `loop-kit-dev` spec, skill body, manifest prompt가 같은 의미로 설명해야 한다.

## 구조 요약

- 이 플러그인은 intentionally narrow한 operational package다.
- `flow`가 흐름 단위 규칙을 소유하고, `turn-gate`가 메인 turn-level gate 표면이다.
- discovery, broad execution, fix-verify loop, review-loop, commit-readiness는 turn-gate protocol이 아니라 `flow`의 flow-local strategy로 둔다.
- autonomous subagent question routing은 direct skill entrypoint가 아니라 self-drive runtime reference의 책임으로 둔다.
