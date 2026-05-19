# flow intent scenarios

이 폴더는 `flow` 경계 의도를 회귀 평가하기 위한 spec-side fixture를 소유합니다.
runtime skill 본문이 직접 읽는 실행 지시가 아니며, release surface에 포함되지 않습니다.

## 작성 규칙

- 하나의 scenario는 하나의 flow boundary 판단을 검증합니다.
- expected behavior에는 active flow, parent flow, sub-flow candidate, phase/handoff 중 어떤 분류가 맞는지 드러냅니다.
- sub-flow 후보를 만드는 scenario는 후보 생성이 실행이 아니라는 forbidden behavior를 포함합니다.
- turn-level next-flow 질문, terminal closure, self-drive continuation은 `turn-gate` 또는 self-drive scenario가 소유하므로 여기서는 flow 산출물의 handoff 조건으로만 다룹니다.

## Scenario Index

- `parent-sub-flow-candidates.md`: 큰 요청을 parent flow로 받고 finite sub-flow 후보를 산출하는 기준.
- `phase-is-not-flow.md`: phase 이름만 있는 항목을 flow로 승격하지 않는 기준.
- `sub-flow-candidate-not-execution.md`: sub-flow 후보 생성과 실행을 분리하는 기준.
