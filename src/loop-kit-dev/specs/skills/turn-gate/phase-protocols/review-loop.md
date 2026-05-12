# turn-gate review-loop phase protocol spec

## 목적

`review-loop` phase protocol은 review feedback, QA finding, self-review finding이 현재 flow의 주요 병목일 때 적용하는 세부 규격입니다.
이 protocol은 mode가 아니며, material finding을 선별하고 처리하는 phase-level 계약입니다.

## 적합 기준

- 이미 나온 review, QA, self-review finding 중 실제 진행을 막는 문제가 있다.
- finding 하나가 correctness, regression, reliability, delivery risk에 직접 연결된다.
- 수정과 검증을 같은 flow boundary 안에서 처리할 수 있다.

## 부적합 기준

- finding 수집 자체가 목적이다.
- open-ended cleanup이나 speculative polish가 중심이다.
- finding이 새 scope, 새 approval boundary, destructive/external action을 요구한다.
- commit execution이나 PR/publish 같은 외부 handoff가 필요하다.

## 핵심 계약

- finding을 모두 처리하려고 하지 말고 blocking threshold를 먼저 적용한다.
- 한 loop는 하나의 bounded blocking finding에 집중한다.
- 수정 뒤 즉시 해당 finding을 검증한다.
- low-value note는 현재 flow를 넓히지 말고 follow-up candidate로 둔다.
- finding이 broader ambiguous scope나 새 approval boundary를 만들면 preparation 또는 question-routing으로 되돌린다.
- review feedback을 위험 작업 승인으로 해석하지 않는다.

## Handoff

- finding 처리가 끝나면 verification phase에서 결과를 확인한다.
- 추가 finding 처리가 필요하면 같은 flow boundary 안에서 정당한지 재평가하고, 넓어지면 새 flow 후보로 분리한다.

## 검토 질문

- 이 finding은 실제로 blocking인가?
- 한 번에 하나의 bounded finding만 처리했는가?
- 낮은 가치의 note가 scope를 넓히지 않았는가?
