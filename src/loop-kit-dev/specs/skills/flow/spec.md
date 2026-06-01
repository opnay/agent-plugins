# flow 스킬 스펙

## 목적

`flow`는 메시지, 동작, 계획 항목, review finding, handoff를 응집된 작업 흐름 단위로 해석합니다.
직접 실행 가능한 active flow인지, parent flow가 만들 finite `sub-flow candidates`인지, 또는 flow가 아닌 phase/reporting 항목인지 판단합니다.
work 전에는 사용자 의도와 경계를 확인해 flow contract를 잠그고, active flow 안에서는 discovery, review-loop, fix-verify-loop, broad-execution, handoff 전략을 고릅니다.
next-flow 질문, 세션 지속, terminal closure는 `turn-gate`가 소유합니다.

## 경계

- 포함:
  - flow 정의와 flow-vs-phase 구분
  - parent flow와 `sub-flow candidates` 관계
  - `operational-preparation flow`와 `change-unit flow` 구분
  - active flow와 follow-up/sub-flow 후보 구분
  - flow 내부 `intake -> framing -> preparation -> work -> verification -> reporting` 계약
  - `turn-gate`의 `interruption`은 일반 flow phase가 아니라 active flow 도중 사용자 메시지를 분류하는 entry-only routing임을 구분
  - input analysis, deep interview, goal detection, non-goal detection, authority detection
  - flow decomposition, flow design, candidate-vs-selected distinction, artifact ownership
  - active flow phase start/end record checkpoint 계약
  - active flow 도중 새 사용자 메시지가 기존 flow contract를 바꾸는지 판단하는 contract-impact 기준
  - flow readiness, intent-first requirement discovery, operation/target ambiguity 판단
  - flow-local review handling, fix-verify-reassess, broad execution strategy
  - flow completion criteria와 verification expectation 산출
  - `레포지토리 목적 > 모노레포 목적 > 구조적 목적 > 변경 목적` 같은 지속 목적 계층 해석
  - `flow`만 해석하고 사용할 수 있는 목적 사슬 파일 계약
  - commit-readiness 같은 flow handoff condition 판단
  - flow가 아닌 분석, 검증, 보고, evidence repair, blocker recovery, commit-readiness 항목 판정
- 제외:
  - turn activation과 explicit stop 처리
  - 질문 도구 실행 방식과 next-flow question-routing
  - 결과 보고 뒤 next-flow 질문 도구 라우팅
  - 여러 flow 사이의 turn-level continuity
  - self-drive sequence-level continuation
  - session record의 date-level active flow pointer와 terminal closure guard
  - commit, push, PR, publish 같은 외부 실행 세부 계약

## 처리하려는 작업 형태

- 하나의 사용자 메시지 또는 동작을 flow로 해석해야 하는 작업
- 큰 요청을 parent flow로 받고 finite `sub-flow candidates`로 나눠야 하는 작업
- 어떤 항목이 flow인지 phase인지, 또는 handoff/reporting인지 판정해야 하는 작업
- flow별 scope, non-goals, completion criteria, verification expectation, handoff 조건을 설계해야 하는 작업
- 여러 flow가 이어질 때 상위 목적은 유지하고 변경 목적만 바뀌는지 판단해야 하는 작업
- 목적 사슬을 별도 파일로 드러내되, 그 파일이 flow 전용 판단 표면이어야 하는 작업
- flow contract를 만들기 위한 intent, scope, tradeoff, acceptance 질문 또는 operation/target ambiguity를 판정해야 하는 작업
- 사용자 입력을 분석하고 목표, 비목표, authority, 모호성을 탐지해야 하는 작업
- flow 후보를 분리하고 각 후보의 산출물과 실행 경계를 설계해야 하는 작업
- active flow 안에서 review finding, 작은 fix loop, broad execution 중 어떤 strategy가 필요한지 판단해야 하는 작업
- 이미 끝난 flow가 완료 조건을 만족했는지 검토해야 하는 작업

## 엔트리포인트 / 대표 표면

- runtime skill: `src/loop-kit-dev/skills/flow/SKILL.md`
- intent: `src/loop-kit-dev/specs/skills/flow/intent.md`
- spec index: `src/loop-kit-dev/specs/skills/flow/spec.md`
- child specs: `src/loop-kit-dev/specs/skills/flow/**`
- 적용 방식: 직접 호출 또는 `turn-gate` preparation에서 flow boundary를 잠글 때 적용합니다.

## 계약 맵

- `intent.md`: 사용자 스펙 의도 기록
- `core/model.md`: flow, parent flow, sub-flow candidate, active flow 모델
- `core/types.md`: `operational-preparation flow`와 `change-unit flow` 구분
- `core/boundaries.md`: flow-vs-phase, flow가 아닌 항목, reviewable artifact 기준
- `core/output-contract.md`: flow 설계 또는 sub-flow 후보 산출물의 필수 필드
- `core/object.md`: `flow` 전용 목적 사슬 파일 계약
- `core/turn-gate-relationship.md`: `flow`와 `turn-gate`의 소유권 경계
- `core/phase-record-checkpoints.md`: active flow phase start/end에서 필요한 plan 또는 flow record checkpoint
- `intake.md`: 사용자 입력 분석, deep interview, goal/non-goal/authority 탐지
- `framing.md`: flow 분리, flow 설계, candidate-vs-selected 구분, artifact ownership
- `preparation/readiness.md`: work 진입 전 flow contract 충분성
- `preparation/discovery.md`: intent-first requirement discovery와 scope/tradeoff lock 질문 주제
- `preparation/ambiguity.md`: operation/target ambiguity
- `execution/review-loop.md`: active flow 안의 review/QA/self-review finding 처리 전략
- `execution/fix-verify-loop.md`: 작은 fix-verify-reassess cycle 전략
- `execution/broad-execution.md`: locked scope 단일 flow의 end-to-end execution 전략
- `handoff/commit-readiness.md`: commit execution이 아닌 commit-readiness handoff 판단
- `intent-scenarios/`: flow boundary 의도를 회귀 평가하기 위한 spec-side fixture

## 핵심 처리 계약

- flow는 phase checklist가 아니라 이해, 리뷰, 검증, 필요 시 커밋 가능한 작업 단위입니다.
- 하나의 flow는 `intake -> framing -> preparation -> work -> verification -> reporting`을 내부 단계로 갖습니다.
- `interruption`은 `flow` 내부 phase가 아닙니다. active flow 도중 새 사용자 메시지가 도착하면 `turn-gate`가 entry-only routing으로 처리하고, 필요할 때 flow 계약 갱신이나 새 flow 전환에 `flow` 판단을 적용합니다.
- `flow`는 interruption 자체를 운영하지 않지만, 새 메시지가 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 바꾸는지 판단하는 원천입니다.
- `intake`는 사용자 입력 분석, deep interview, 목표/비목표/authority 탐지를 소유합니다.
- `framing`은 flow 분리, flow 설계, candidate-vs-selected 구분, artifact ownership 판단을 소유합니다.
- `preparation`은 선택된 active flow의 readiness, scope/non-goals/completion/verification/approval boundary lock, work 진입 가능성 판단으로 좁힙니다.
- 각 active flow phase 시작과 종료는 `000-plan.md` 또는 active flow record 중 갱신 표면을 드러내야 합니다.
- `000-plan.md` 갱신이 필요한 경우에는 현재 flow 또는 planned sequence에서 사용할 skill 목록도 필요한 만큼 드러내야 합니다.
- flow preparation은 이미 선택된 active flow의 readiness를 잠그는 단계이며, intake/framing에서 미해결 필드가 발견되면 work로 넘어가지 않습니다.
- flow contract는 필요한 경우 목적 계층을 드러냅니다. 상위 목적은 repository/monorepo/structure 같은 지속 관점으로 유지하고, 현재 변경 목적은 active flow scope와 함께 잠급니다.
- 목적 사슬 파일이 필요한 경우 그 의미와 사용 권한은 `flow`가 소유합니다. 파일은 상태, 검증, continuity rule이 아니라 객체 사슬만 담습니다.
- flow execution은 current flow 안에서 review-loop, fix-verify-loop, broad-execution을 선택할 수 있습니다.
- verification, reporting, evidence repair, blocker recovery는 별도 reviewable artifact를 만들지 않으면 현재 flow 내부 phase 또는 handoff입니다.
- evidence 부족은 current flow verification으로, metadata mismatch는 verification mismatch 해소로, scope/target/approval/verification expectation 변경은 preparation으로, access/input/approval/external blocker는 blocked handoff로 라우팅합니다.
- review-loop는 여러 review finding 전체를 한 번에 실행하는 포괄 전략이 아니라, active flow 안의 bounded blocking finding 하나를 처리하는 전략입니다. 여러 finding이 있으면 우선순위 선택, discovery, 또는 finite follow-up 후보 설계가 먼저입니다.
- flow가 너무 크거나 여러 산출물을 만들면 parent flow는 finite `sub-flow candidates`를 만들 수 있습니다.
- `sub-flow candidate` 생성은 실행이 아닙니다.
- flow handoff는 다음 사용자 질문이나 commit execution을 직접 수행하지 않고 handoff condition을 산출합니다.
- self-drive나 turn-gate가 다음 flow identity, current flow completion, non-blocked handoff를 확인할 때도 `flow`의 output contract와 handoff condition을 원천으로 삼습니다.
- flow가 끝났다는 사실은 turn이 끝났다는 뜻이 아닙니다.

## 검토 질문

- 현재 항목이 flow인가, phase인가, handoff인가?
- 사용자 입력 분석과 목표/비목표/authority 탐지가 intake에서 끝났는가?
- flow 분리와 후보/선택 구분이 framing에서 끝났는가?
- flow가 너무 커서 finite sub-flow 후보로 나눠야 하는가?
- sub-flow 후보가 active execution flow처럼 실행되고 있지 않은가?
- 각 flow에 scope, non-goals, completion criteria, verification expectation, handoff 조건이 있는가?
- 연속 flow에서 상위 목적과 현재 변경 목적을 구분했는가?
- 목적 사슬 파일을 `turn-gate` 라우팅이나 세션 상태 파일처럼 사용하지 않았는가?
- verification/reporting/repair를 새 flow로 분리하기 전에 별도 reviewable artifact가 있는지 확인했는가?
- 각 phase 시작과 종료에서 `000-plan.md` 또는 active flow record 갱신 기준이 드러나는가?
- `000-plan.md`에 사용할 skill 목록이 필요한 flow 또는 planned sequence 기준으로 유지되는가?
- intake/framing/preparation 중 어느 단계가 부족한지 구분하지 않은 채 work로 넘어가지 않았는가?
- flow-local strategy를 turn-level next-flow routing이나 self-drive sequence authority와 혼동하지 않았는가?
- review-loop를 여러 finding 묶음 실행으로 넓히지 않고 bounded finding 하나 또는 후보 설계로 처리했는가?
- flow 완료와 turn 종료를 혼동하지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `flow`는 `turn-gate` 없이도 flow boundary, sub-flow 후보, flow readiness, flow-local strategy, flow completion을 판단할 수 있어야 합니다. 다만 `turn-gate`가 활성인 경우에는 `turn-gate`가 next-flow 질문, session continuity, terminal closure guard를 소유합니다.

## 확장 원칙

- flow 분해 규칙이 커지면 `core/*` child spec 또는 `intent-scenarios/*` fixture로 내립니다.
- 새로운 flow type은 기존 `operational-preparation`과 `change-unit`으로 표현할 수 없을 때만 추가합니다.
- runtime skill은 finite sub-flow 후보와 실행 금지 경계를 직접 설명합니다.
