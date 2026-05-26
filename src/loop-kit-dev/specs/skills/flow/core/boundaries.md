# flow boundary 계약

## 소유 범위

flow-vs-phase, flow가 아닌 항목, reviewable artifact 기준.

## 계약

- `analysis`, `work`, `verification`, `reporting`, `commit readiness` 같은 phase 이름은 그 자체로 flow가 아닙니다.
- 순수 최종 QA, 통합 검증, 정합성 점검, 검증 결과 보고, commit-readiness reporting은 별도 산출물 변경을 소유하지 않는 한 flow가 아닙니다.
- 회귀 테스트 fixture, snapshot baseline, 운영자 리포트 출력, validator 진단 출력처럼 검토 가능한 산출물을 만들거나 바꾸면 그 산출물 변경은 flow가 될 수 있습니다.
- flow는 최종 사용자에게 직접 보이는 가치 단위일 필요는 없지만, 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 단위여야 합니다.

## 검토 기준

- 이 항목이 phase 이름만 가진 것은 아닌가?
- 별도 검토 가능한 산출물을 소유하는가?
- flow로 나누면 이해, 리뷰, 검증, 커밋 단위가 더 명확해지는가?
