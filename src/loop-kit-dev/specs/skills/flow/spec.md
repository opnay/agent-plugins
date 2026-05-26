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
  - flow 내부 `preparation -> work -> verification -> reporting` 계약
  - active flow phase start/end record checkpoint 계약
  - flow readiness, intent-first requirement discovery, operation/target ambiguity 판단
  - flow-local review handling, fix-verify-reassess, broad execution strategy
  - flow completion criteria와 verification expectation 산출
  - commit-readiness 같은 flow handoff condition 판단
  - flow가 아닌 분석, 검증, 보고, commit-readiness 항목 판정
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
- flow contract를 만들기 위한 intent, scope, tradeoff, acceptance 질문 또는 operation/target ambiguity를 판정해야 하는 작업
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
- `core/turn-gate-relationship.md`: `flow`와 `turn-gate`의 소유권 경계
- `core/phase-record-checkpoints.md`: active flow phase start/end에서 필요한 plan 또는 flow record checkpoint
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
- 하나의 flow는 `preparation -> work -> verification -> reporting`을 내부 단계로 갖습니다.
- 각 active flow phase 시작과 종료는 `000-plan.md` 또는 active flow record 중 갱신 표면을 드러내야 합니다.
- flow preparation은 readiness, intent-first discovery, ambiguity 판단으로 flow contract를 완성합니다.
- flow execution은 current flow 안에서 review-loop, fix-verify-loop, broad-execution을 선택할 수 있습니다.
- review-loop는 여러 review finding 전체를 한 번에 실행하는 포괄 전략이 아니라, active flow 안의 bounded blocking finding 하나를 처리하는 전략입니다. 여러 finding이 있으면 우선순위 선택, discovery, 또는 finite follow-up 후보 설계가 먼저입니다.
- flow가 너무 크거나 여러 산출물을 만들면 parent flow는 finite `sub-flow candidates`를 만들 수 있습니다.
- `sub-flow candidate` 생성은 실행이 아닙니다.
- flow handoff는 다음 사용자 질문이나 commit execution을 직접 수행하지 않고 handoff condition을 산출합니다.
- flow가 끝났다는 사실은 turn이 끝났다는 뜻이 아닙니다.

## 검토 질문

- 현재 항목이 flow인가, phase인가, handoff인가?
- flow가 너무 커서 finite sub-flow 후보로 나눠야 하는가?
- sub-flow 후보가 active execution flow처럼 실행되고 있지 않은가?
- 각 flow에 scope, non-goals, completion criteria, verification expectation, handoff 조건이 있는가?
- 각 phase 시작과 종료에서 `000-plan.md` 또는 active flow record 갱신 기준이 드러나는가?
- readiness/discovery/ambiguity가 필요한데 사용자 의도와 경계가 잠기지 않은 채 work로 넘어가지 않았는가?
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
