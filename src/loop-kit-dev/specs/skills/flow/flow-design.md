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
-> 진행 순서 정리
-> 메인 플로우 선택
```

진행 순서 정리가 실패하면 항목 분류로 돌아갑니다.

## 계약

- 플로우 설계는 locked execution brief를 실제 진행할 flow 구성으로 바꿉니다.
- 항목 분류는 active flow, parent flow, sub-flow candidate, phase, handoff를 구분합니다.
- flow 분해는 단일/다중 flow 판단, parent flow 또는 단일 active flow 식별, sub-flow candidate 추출, 후보 pending 상태 표시로 나눕니다.
- 후보 pending 상태 표시는 후보가 다음 메인 flow로 선택되기 전까지 실행 권한을 만들지 않음을 드러냅니다.
- flow별 계약 작성은 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 정리하고, 선택된 active flow의 flow record를 작성합니다.
- verification expectation은 requirement verification과 implementation verification의 축을 구분할 수 있어야 합니다.
- 진행 순서 정리는 다음에 들어갈 메인 flow와 이후 후보를 구분합니다.
- 이후 후보는 다음 메인 flow로 선택되기 전까지 대기 상태이며 실행 권한을 만들지 않습니다.
- 목적 사슬은 contract에 영향을 줄 때 `000-plan.md`의 목적 섹션에 흡수합니다.

## 산출

- 진행할 메인 flow
- 이후 후보
- parent flow 또는 단일 active flow 판단
- 후보 pending 상태
- flow별 계약
- 선택된 active flow의 flow record
- 진행 순서
- `000-plan.md` 갱신 필요 여부
