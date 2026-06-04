# flow 플로우 설계 스펙

## 기준 그래프

```text
locked execution brief -> 항목 분류 -> flow 분해 -> flow별 계약 작성 -> 진행 순서 정리 -> 메인 플로우 선택
```

진행 순서 정리가 실패하면 항목 분류로 돌아갑니다.

## 계약

- 플로우 설계는 locked execution brief를 실제 진행할 flow 구성으로 바꿉니다.
- 항목 분류는 active flow, parent flow, sub-flow candidate, phase, handoff를 구분합니다.
- flow 분해는 단일 메인 flow로 충분한지, 여러 메인 flow가 필요한지 정리합니다.
- 질문 답변, 상태 확인, 설명 요청도 active flow 후보로 분류할 수 있습니다.
- sub-flow candidate는 선택 전 대기 항목이며 실행 권한을 갖지 않습니다.
- phase는 active flow 내부 단계이며 별도 flow가 아닙니다.
- handoff는 메인 플로우 종료 뒤 산출되는 조건이며 실행 권한을 만들지 않습니다.
- 목적 사슬이 계약에 영향을 주면 `000-plan.md` 목적 섹션에 흡수합니다.

## flow별 계약

각 flow 계약은 다음 항목을 가집니다.

- identity
- scope
- non-goals
- completion criteria
- verification expectation
- approval boundary
- handoff condition

## 산출

- 진행할 메인 flow
- 이후 후보 목록
- flow별 계약
- `000-plan.md` 갱신 내용
- unresolved blocker
