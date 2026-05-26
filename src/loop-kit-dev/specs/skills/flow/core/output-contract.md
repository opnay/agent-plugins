# flow output 계약

## 소유 범위

flow 설계 또는 sub-flow 후보 산출물의 필수 필드.

## 계약

flow 설계 또는 sub-flow 후보 산출물에는 최소한 다음이 드러나야 합니다.

- flow label 또는 slug
- flow type: `operational-preparation` 또는 `change-unit`
- scope
- non-goals
- completion criteria
- verification expectation
- phase start/end record checkpoint expectation
- readiness status 또는 missing contract fields
- recommended question topics 또는 unresolved ambiguity
- recommended flow-local strategy
- approval-sensitive checkpoint가 필요한지 여부
- handoff condition
- unresolved questions 또는 blocker
- active flow인지, sub-flow candidate인지 여부

## 검토 기준

- 후보가 active flow인지 sub-flow candidate인지 명확한가?
- completion criteria와 handoff condition이 분리되어 있는가?
- verification expectation이 flow 위험도에 맞게 드러나는가?
- phase start/end에서 `000-plan.md` 또는 active flow record 중 무엇을 갱신할지 드러나는가?
- missing contract field와 recommended strategy가 execution authority처럼 쓰이지 않는가?
