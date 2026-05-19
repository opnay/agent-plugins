# flow 스킬 스펙

## 목적

`flow`는 하나의 메시지나 동작을 작업 흐름 단위로 해석하고, 그 flow가 직접 실행 가능한지 또는 finite `sub-flow candidates`로 나뉘어야 하는지 판단하는 skill입니다.
각 flow는 자체 `preparation -> work -> verification -> reporting` 흐름을 가지며, flow가 끝나면 그 flow는 종료됩니다.
다음 flow 진행 여부와 사용자 질문 라우팅은 `turn-gate`가 소유합니다.

## 경계

- 포함:
  - flow 정의와 flow-vs-phase 구분
  - parent flow와 `sub-flow candidates` 관계
  - `operational-preparation flow`와 `change-unit flow` 구분
  - active flow와 follow-up/sub-flow 후보 구분
  - flow 내부 `preparation -> work -> verification -> reporting` 계약
  - flow completion criteria와 verification expectation 산출
  - flow가 아닌 분석, 검증, 보고, commit-readiness 항목 판정
- 제외:
  - turn activation과 explicit stop 처리
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
- 이미 끝난 flow가 완료 조건을 만족했는지 검토해야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit-dev/skills/flow/SKILL.md`
- 사용자 스펙 의도: `loop-kit-dev/specs/skills/flow/intent.md`
- skill spec index: `loop-kit-dev/specs/skills/flow/spec.md`
- sub-spec directory: `loop-kit-dev/specs/skills/flow/`
- 호출 방식: 직접 호출하거나, `turn-gate` preparation이 flow boundary를 잠글 때 이 계약을 적용한다.

## 상세 계약 구조

- `intent.md`: 사용자 스펙 의도 기록
- `core/model.md`: flow, parent flow, sub-flow candidate, active flow의 핵심 모델
- `core/types.md`: `operational-preparation flow`와 `change-unit flow` 구분
- `core/boundaries.md`: flow-vs-phase, flow가 아닌 항목, reviewable artifact 기준
- `core/output-contract.md`: flow 설계 또는 sub-flow 후보 산출물의 필수 필드
- `core/turn-gate-relationship.md`: `flow`와 `turn-gate`의 소유권 경계
- `intent-scenarios/`: flow boundary 의도를 회귀 평가하기 위한 spec-side fixture

## 핵심 처리 계약

- flow는 phase checklist가 아니라 응집된 작업 흐름 단위입니다.
- 하나의 flow는 `preparation -> work -> verification -> reporting`을 내부 단계로 갖습니다.
- flow가 너무 크거나 여러 산출물을 만들면 parent flow는 finite `sub-flow candidates`를 만들 수 있습니다.
- `sub-flow candidate` 생성은 실행이 아닙니다.
- flow가 끝났다는 사실은 turn이 끝났다는 뜻이 아닙니다.

## 검토 질문

- 현재 항목이 flow인가, phase인가, handoff인가?
- flow가 너무 커서 finite sub-flow 후보로 나눠야 하는가?
- sub-flow 후보가 active execution flow처럼 실행되고 있지 않은가?
- 각 flow에 scope, non-goals, completion criteria, verification expectation, handoff 조건이 있는가?
- flow 완료와 turn 종료를 혼동하지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `flow`는 `turn-gate` 없이도 flow boundary, sub-flow 후보, flow completion을 판단할 수 있어야 합니다. 다만 `turn-gate`가 활성인 경우에는 `turn-gate`가 next-flow 질문, session continuity, terminal closure guard를 소유합니다.

## 확장 원칙

- flow 분해 규칙이 커지면 `core/*` child spec 또는 `intent-scenarios/*` fixture로 내립니다.
- 새로운 flow type은 기존 `operational-preparation`과 `change-unit`으로 표현할 수 없을 때만 추가합니다.
- runtime skill은 finite sub-flow 후보와 실행 금지 경계를 직접 설명해야 합니다.
