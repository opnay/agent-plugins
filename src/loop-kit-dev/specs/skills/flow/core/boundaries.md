# flow boundary 계약

## 소유 범위

flow-vs-phase, flow가 아닌 항목, reviewable artifact 기준, 그리고 검증/보고/준비 상태가 독립 flow가 되는 조건.

## 계약

flow는 함께 이해하고, 검토하고, 검증하고, 필요하면 커밋할 수 있는 작업 단위입니다.
phase 이름이나 진행 상태 자체는 flow가 아닙니다.

다음 항목은 기본적으로 active flow 내부 phase 또는 handoff입니다.

- 분석
- work
- verification
- reporting
- final QA
- 통합 검증
- 정합성 점검
- commit-readiness reporting
- evidence repair
- blocker recovery

이 항목들이 별도 flow가 되려면 독립적으로 검토 가능한 산출물을 만들거나 바꾸어야 합니다.
예를 들어 회귀 테스트 fixture, snapshot baseline, 운영자 리포트 출력, validator 진단 출력, release surface 같은 산출물을 만들거나 갱신하면 해당 산출물 변경은 flow가 될 수 있습니다.
그렇지 않으면 현재 active flow의 verification, repair, reporting, handoff로 남습니다.

verification이나 reporting을 “4개 flow 중 하나”처럼 세야 하는 요청이 들어와도, 산출물이 없으면 flow가 아니라 phase로 분류합니다.
사용자가 “4개 flow를 진행”하라고 했더라도 flow 개수는 phase 개수가 아니라 reviewable work unit 개수로 계산합니다.

`interruption`은 일반 flow phase가 아닙니다.
active flow 도중 새 사용자 메시지를 분류하는 `turn-gate` entry-only routing이며, 이때 flow는 contract-impact 판단만 제공합니다.

검증이나 도구 접근이 막힌 경우에는 새 flow를 만들기 전에 가장 이른 안전한 재진입 지점을 판단합니다.

- evidence만 부족하면 현재 flow의 verification으로 돌아갑니다.
- work evidence와 verification metadata가 충돌하면 verification mismatch를 해소합니다.
- target, scope, approval boundary, verification expectation이 바뀌면 preparation으로 돌아갑니다.
- blocker가 user input, approval, access, external state를 요구하면 blocked handoff를 산출합니다.

## 검토 기준

- 이 항목이 phase 이름이나 상태 이름만 가진 것은 아닌가?
- 별도 검토 가능한 산출물을 소유하는가?
- verification/reporting/repair를 flow로 부르기 전에 reviewable artifact가 있는지 확인했는가?
- blocker나 insufficient evidence가 새 flow가 아니라 현재 flow의 earliest safe phase로 라우팅되는가?
- flow로 나누면 이해, 리뷰, 검증, 커밋 단위가 더 명확해지는가?
