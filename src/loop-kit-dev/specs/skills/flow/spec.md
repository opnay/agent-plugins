# flow 스킬 스펙

## 목적

`flow`는 새 사용자 메시지를 `메시지 인터뷰 -> 플로우 설계 -> 메인 플로우 -> handoff condition`으로 해석합니다.
메시지 인터뷰는 flow 내부 alignment loop이고, 플로우 설계는 locked brief에서 진행할 flow 구성을 만들며, 메인 플로우는 `intake -> framing -> preparation -> work -> verification -> reporting`으로 실행됩니다.

## 경계

- 포함: 메시지 인터뷰, locked brief, flow 분류, flow 구성, active flow contract, phase checkpoint, reporting handoff, contract-impact 판단.
- 제외: flow 밖 실행 제어, 질문 실행, 연속 진행 제어, commit/push/PR/release/version bump 실행.

## 계약 맵

- `intent.md`: 현재 다이어그램
- `intake.md`: 메시지 인터뷰
- `framing.md`: 플로우 설계
- `core/model.md`: 메시지 인터뷰/플로우 설계/메인 플로우/handoff 관계
- `core/types.md`: operational-preparation/change-unit
- `core/boundaries.md`: active/parent/candidate/phase/handoff
- `core/output-contract.md`: 산출 필드
- `core/phase-record-checkpoints.md`: 기록 표면
- `core/object.md`: `000-plan.md` 목적 섹션
- `preparation/*`: readiness/discovery/ambiguity relock
- `execution/*`: 메인 flow 내부 전략
- `handoff/commit-readiness.md`: commit-ready handoff
- `intent-scenarios/`: 회귀 fixture

## 핵심 계약

- 메시지 인터뷰는 `snapshot -> alignment risk -> high-leverage question -> answer -> pressure test`를 반복해 locked brief를 만듭니다.
- 플로우 설계는 locked brief에서 active flow, parent flow, sub-flow candidate, phase, handoff를 구분하고 flow별 contract를 작성합니다.
- candidate는 선택되기 전 pending option입니다.
- 메인 플로우는 선택된 active flow의 lifecycle입니다.
- reporting은 result, verification, residual risk, handoff condition, 다음 intake 조건을 산출합니다.
- flow completion은 handoff condition을 남깁니다.

## 검토 질문

- alignment risk가 잠겼는가?
- active flow와 candidate가 구분되는가?
- scope, non-goals, completion, verification, approval, handoff가 있는가?
- 다음 intake 조건이 flow 산출물로 구분되는가?

## 확장 원칙

새 규칙은 메시지 인터뷰, 플로우 설계, 메인 플로우, handoff 중 하나에 귀속합니다.
목적 사슬, 실행 전략, commit-readiness는 보조 child spec에서 필요할 때 적용합니다.
