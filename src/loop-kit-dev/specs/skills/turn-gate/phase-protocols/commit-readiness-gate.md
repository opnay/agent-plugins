# turn-gate commit-readiness-gate phase protocol spec

## 목적

`commit-readiness-gate` phase protocol은 의도한 change unit이 commit으로 이동할 준비가 됐는지 판단할 때 적용하는 세부 규격입니다.
이 protocol은 readiness 판단과 commit execution approval 근거를 분리해 기록하는 reporting 계약입니다.

## 적합 기준

- 구현이나 문서 변경이 대부분 끝났고 intended change unit을 판단할 수 있다.
- diff scope, unrelated-change exclusion, verification evidence, residual risk를 정리해야 한다.
- 사용자가 commit 준비, readiness, 최종 점검을 요청했다.

## 부적합 기준

- broad implementation 자체가 남아 있다.
- intended change unit이 아직 불명확하다.
- readiness 판단 없이 commit execution, staging, push, PR, publish, release, version bump 자체가 목적이다.
- verification evidence가 부족해 readiness 판단 전에 보완 work가 필요하다.

## 핵심 계약

- 전체 저장소가 아니라 intended change unit만 평가한다.
- work boundary, unrelated changes to exclude, verification status, residual risk, likely commit-message scope를 확인한다.
- 필요한 scoped verification과 final review를 수행한다.
- readiness, residual risk, intended diff scope, excluded unrelated changes, verification evidence, minimum review recommendation을 함께 보고한다.
- readiness request 자체는 readiness evidence만 제공한다.
- stage, commit, push, PR, publish, release, version bump 실행 근거는 approval boundary에서 확인한다.

## Handoff

- readiness가 `ready`면 approval boundary에 따라 execution handoff 또는 next-flow question을 연다.
- readiness가 부족하면 earliest safe phase로 돌아가 보완 work 또는 question-routing을 연다.

## 검토 질문

- intended change unit이 충분히 잠겨 있는가?
- readiness 판단에 필요한 verification evidence가 있는가?
- commit readiness와 execution authority 근거를 구분했는가?
