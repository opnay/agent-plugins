# flow 플로우 설계 스펙

## 기준 그래프

```text
locked execution brief
-> 항목 분류
-> 단일/다중 flow 판단
-> parent flow 또는 단일 active flow 식별
-> sub-flow candidate 추출
-> 후보 pending 상태 표시
-> flow별 계약 작성
-> flow 계약 검증
-> 메인 플로우 선택
```

flow 계약 검증이 실패하면 항목 분류로 돌아갑니다.

## 계약

- 플로우 설계는 locked execution brief를 실제 진행할 flow 구성으로 바꿉니다.
- 항목 분류는 active flow, parent flow, sub-flow candidate, phase, handoff를 구분합니다.
- flow 분해는 단일/다중 flow 판단, parent flow 또는 단일 active flow 식별, sub-flow candidate 추출, 후보 pending 상태 표시로 나눕니다.
- 후보 pending 상태 표시는 후보가 다음 메인 flow로 선택되기 전까지 실행 권한을 만들지 않음을 드러냅니다.
- flow별 계약 작성은 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 정리하고, 선택된 active flow의 flow record를 작성합니다.
- verification expectation은 requirement verification과 implementation verification의 축을 구분할 수 있어야 합니다.
- flow 계약 검증은 선택된 flow가 별도 완료 기준, 검증 결과, 승인 경계, handoff 조건, 독립 산출물을 갖는 독립 계약 단위인지 확인합니다.
- flow 계약 검증은 내부 todo나 phase 수준의 작업이 별도 flow로 과하게 분해되지 않았는지 확인합니다.
- flow 계약 검증은 artifact 기준, 승인 기준, 검증 기준, 사용자 목적 기준 중 어떤 분해 축이 현재 계약에 더 맞는지 확인하고 필요하면 조정합니다.
- flow 계약 검증은 다음에 들어갈 메인 flow와 이후 후보의 순서가 의존관계, 승인 경계, 검증 기대와 맞는지 확인합니다.
- flow 계약 검증은 사용자 계약 의도가 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition에 빠짐없이 반영됐는지 확인합니다.
- 이후 후보는 다음 메인 flow로 선택되기 전까지 대기 상태이며 실행 권한을 만들지 않습니다.
- 목적 사슬은 contract에 영향을 줄 때 `000-plan.md`의 목적 섹션에 흡수합니다.

## 산출

- 진행할 메인 flow
- 이후 후보
- parent flow 또는 단일 active flow 판단
- 후보 pending 상태
- flow별 계약
- 선택된 active flow의 flow record
- flow 계약 검증 결과
- 과분해 여부
- 분해 축 조정 여부
- 진행 순서
- `000-plan.md` 갱신 필요 여부
