# flow output 계약

## 소유 범위

flow 설계 또는 sub-flow 후보 산출물의 필수 필드.

## 계약

flow 설계 또는 sub-flow 후보 산출물에는 최소한 다음이 드러나야 합니다.

- flow label 또는 slug
- flow type: `operational-preparation` 또는 `change-unit`
- 연속 flow의 지속 관점에 영향을 주는 경우 목적 계층: 레포지토리 목적 > 모노레포 목적 > 구조적 목적 > 변경 목적
- 목적 사슬 파일을 쓰는 경우 `flow` 전용 object chain
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
- contract-impact result when a new message may revise an active flow
- unresolved questions 또는 blocker
- active flow인지, sub-flow candidate인지 여부

## 검토 기준

- 후보가 active flow인지 sub-flow candidate인지 명확한가?
- completion criteria와 handoff condition이 분리되어 있는가?
- 목적 계층이 필요한 경우 상위 목적과 현재 변경 목적이 구분되는가?
- 목적 사슬 파일이 필요한 경우 상태나 라우팅 정보 없이 object chain만 담는가?
- verification expectation이 flow 위험도에 맞게 드러나는가?
- phase start/end에서 `000-plan.md` 또는 active flow record 중 무엇을 갱신할지 드러나는가?
- missing contract field와 recommended strategy가 execution authority처럼 쓰이지 않는가?
- 새 메시지나 self-drive advance가 기존 flow의 completion, handoff, verification expectation을 바꾸는지 판단할 근거가 있는가?
