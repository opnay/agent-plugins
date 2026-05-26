# flow broad-execution 계약

## 소유 범위

scope가 잠긴 단일 active flow를 end-to-end로 수행하는 flow-local 실행 전략.

## 계약

- broad execution은 scope, non-goals, completion criteria, verification expectation, approval boundary가 충분히 잠긴 단일 flow 안에서만 사용한다.
- implementation, QA, validation이 같은 flow boundary 안에 있을 때 적합하다.
- blocking clarification이 없고 기록된 boundary 안에서 계속 진행할 수 있어야 한다.
- 의미 있는 수정 뒤에는 변경 표면에 맞는 검증을 수행한다.
- QA issue가 생기면 review-loop 또는 fix-verify-loop로 좁혀 처리한다.
- 여러 flow를 자동으로 이어가는 sequence-level continuation은 self-drive가 소유한다.
- destructive, external, commit, push, PR, publish, release, version bump 실행 승인은 approval boundary가 소유한다.

## 검토 기준

- 단일 active flow의 scope가 broad execution을 시작할 만큼 잠겼는가?
- sequence-level continuation을 flow-local execution과 혼동하지 않았는가?
- 검증과 QA 처리 경로가 명시적인가?
