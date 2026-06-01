# flow intent scenarios

이 폴더는 `flow` 경계 의도를 회귀 평가하기 위한 spec-side fixture를 소유합니다.
runtime skill 본문은 현재 spec에서 재작성되며, 이 fixture는 release surface 밖의 회귀 평가 자료입니다.

## fixture 계약

- scenario는 메시지 인터뷰, 플로우 설계, 메인 플로우, handoff 중 하나의 판단을 검증합니다.
- expected behavior는 locked brief, flow 구성, active flow, candidate, phase, handoff 중 필요한 산출물을 드러냅니다.
- forbidden behavior는 candidate 실행, phase 승격, turn-gate 권한 침범을 드러냅니다.

## 현재 fixture

- `parent-sub-flow-candidates.md`: 메시지 인터뷰 뒤 플로우 설계가 여러 candidate를 만들고 routing 전 pending 상태로 유지하는 기준.
- `phase-stays-inside-active-flow.md`: phase 이름만 있는 항목을 active flow 내부 단계로 유지하는 기준.
- `sub-flow-candidate-pending-state.md`: sub-flow candidate와 next main flow 선택을 분리하는 기준.
