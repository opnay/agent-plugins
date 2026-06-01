# flow 스킬 스펙

## 목적

`flow`는 새 사용자 메시지를 `메시지 인터뷰 -> 플로우 설계 -> 메인 플로우 -> 메인 플로우 회고 -> handoff condition`으로 해석합니다.
메시지 인터뷰는 flow 내부 alignment loop이고, 플로우 설계는 locked brief에서 진행할 flow 구성을 만들며, 메인 플로우는 `intake -> framing -> preparation -> work -> verification -> reporting`으로 실행됩니다.

## 경계

- 포함: 메시지 인터뷰, locked brief, flow 분류, flow 구성, active flow contract, phase checkpoint, reporting loop, 메인 플로우 회고, handoff, contract-impact 판단.
- 제외: flow 밖 실행 제어, 질문 실행, 연속 진행 제어, commit/push/PR/release/version bump 실행.

## 계약 맵

- `intent.md`: 현재 다이어그램
- `intake.md`: 메시지 인터뷰
- `framing.md`: 플로우 설계
- `core/model.md`: 메시지 인터뷰/플로우 설계/메인 플로우/회고/handoff 관계
- `core/types.md`: operational-preparation/change-unit
- `core/boundaries.md`: active/parent/candidate/phase/handoff
- `core/output-contract.md`: 산출 필드
- `core/phase-record-checkpoints.md`: 기록 표면
- `templates/plan.md`: flow plan 템플릿 계약
- `templates/flow-record.md`: active flow record 템플릿 계약
- `templates/review.md`: retrospective review 템플릿 계약
- `core/object.md`: `000-plan.md` 목적 섹션
- `preparation/*`: readiness/discovery/ambiguity relock
- `execution/*`: 메인 flow 내부 전략
- `handoff/commit-readiness.md`: commit-ready handoff
- `intent-scenarios/`: 회귀 fixture

## 핵심 계약

- 메시지 인터뷰는 `snapshot -> alignment risk -> high-leverage question -> answer -> pressure test`를 반복해 locked brief를 만듭니다.
- 플로우 설계는 locked brief에서 active flow, parent flow, sub-flow candidate, phase, handoff를 구분하고 flow별 contract를 작성합니다.
- candidate는 선택되기 전 pending option입니다.
- 메인 플로우는 선택된 active flow의 lifecycle이며 `reporting -> intake` loop로 다음 flow를 연결할 수 있습니다.
- 메인 플로우 그룹 이후에는 항상 메인 플로우 회고를 수행하고 `000-review.md`를 갱신한 뒤 handoff condition을 산출합니다.
- 회고 finding이 없으면 no-finding 결과로 짧게 기록합니다.
- reporting은 result, verification, residual risk, 다음 intake 조건을 산출합니다.
- flow completion은 handoff condition을 남깁니다.
- 기록 템플릿 계약은 `flow`가 소유하고, active turn 적용과 실제 기록 갱신은 `turn-gate`가 수행합니다.
- 사용자-facing phase 시작 또는 의미 있는 진행 메시지에는 현재 phase label을 산출합니다. 기본 label은 `[intake]`, `[framing]`, `[preparation]`, `[work]`, `[verification]`, `[reporting]`입니다.
- phase label은 사용자-facing 진행 표시이며, artifact 본문, record 본문, command output summary, question option label에 전파하지 않습니다.

## 검토 질문

- alignment risk가 잠겼는가?
- active flow와 candidate가 구분되는가?
- scope, non-goals, completion, verification, approval, handoff가 있는가?
- 다음 intake 조건이 flow 산출물로 구분되는가?
- 사용자-facing phase/progress label이 phase model에서 산출되고, artifact/record/command/question option 표면에는 전파되지 않는가?
- 메인 플로우 회고가 active routing이나 handoff authority를 소유하지 않는가?
- 메인 플로우 회고가 항상 수행되고 `000-review.md`에 반영되는가?
- plan, flow-record, review 템플릿이 각자 소유할 정보와 금지할 정보를 구분하는가?

## 확장 원칙

새 규칙은 메시지 인터뷰, 플로우 설계, 메인 플로우, 메인 플로우 회고, handoff 중 하나에 귀속합니다.
목적 사슬, 실행 전략, commit-readiness는 보조 child spec에서 필요할 때 적용합니다.
기록 템플릿은 flow 산출물과 복구 표면의 의미를 설명하고, runtime 적용 제어는 turn-gate로 둡니다.
