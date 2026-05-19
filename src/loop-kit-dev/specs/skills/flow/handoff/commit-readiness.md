# flow commit-readiness handoff sub-spec

## 목적

이 문서는 active flow 또는 change unit이 commit handoff로 이동할 준비가 됐는지 판단하는 flow-local handoff 계약을 소유합니다.

## 계약

- commit-readiness는 commit execution 권한이 아니라 readiness 판단과 handoff 조건을 소유한다.
- intended change unit, diff scope, unrelated-change exclusion, verification evidence, residual risk가 충분한지 확인한다.
- readiness 판단은 flow reporting 또는 handoff condition으로 표현한다.
- staging, commit, push, PR, publish, release, version bump 실행은 approval boundary와 해당 실행 workflow가 소유한다.
- verification evidence가 부족하면 readiness를 성공으로 보고하지 않고 flow의 earliest safe phase로 되돌린다.
- commit-readiness 자체가 별도 산출물 변경을 소유하지 않으면 독립 change-unit flow가 아니다.

## 검토 질문

- intended change unit과 unrelated-change exclusion이 분명한가?
- verification evidence와 residual risk가 commit handoff에 충분한가?
- readiness 판단과 execution authority를 분리했는가?
