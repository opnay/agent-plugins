# turn-gate autopilot phase protocol spec

## 목적

`autopilot` phase protocol은 scope가 잠긴 요청을 end-to-end로 구현, QA, 검증, 보고까지 밀고 가야 할 때 적용하는 세부 규격입니다.
이 protocol은 mode가 아니며, implicit default state 안에서 broad execution을 다루는 phase-level 계약입니다.

## 적합 기준

- scope, non-goal, verification expectation, approval boundary가 이미 충분히 잠겼다.
- 작업이 단일 작은 fix보다 넓고 implementation, QA, validation을 모두 포함한다.
- blocking clarification 외에는 계속 진행할 수 있다.

## 부적합 기준

- brainstorming, planning-only, requirement discovery가 주목적이다.
- scope floor가 충족되지 않았다.
- 하나의 좁은 fix-verify-reassess cycle이면 충분하다.
- review finding 하나만 처리하면 된다.
- final commit readiness 판단이나 commit execution이 목적이다.

## 핵심 계약

- 기록된 scope, non-goal, verification expectation, approval boundary 안에서 자율적으로 진행한다.
- scope floor가 충족되지 않으면 deep-interview protocol 또는 question-routing으로 돌아간다.
- 예상 위험 작업은 초기 preparation에서 정확히 승인된 경계 안에서만 진행한다.
- 의미 있는 수정 뒤에는 변경 표면에 맞는 검증을 수행한다.
- QA issue가 생기면 bounded loop로 다루고, 같은 critical failure가 반복되면 root blocker로 보고한다.
- autonomous execution을 destructive, external, commit, push, PR, publish, release, version bump 승인으로 해석하지 않는다.

## Handoff

- implementation과 검증이 끝나면 reporting phase로 넘긴다.
- approval boundary가 새로 생기면 user-gated question-routing으로 돌아간다.
- 마지막 change-unit 이후 commit-readiness는 reporting 또는 별도 handoff로 다루며, commit execution과 구분한다.

## 검토 질문

- scope가 broad execution을 시작할 만큼 잠겼는가?
- blocking clarification 없이 진행 가능한가?
- 검증과 QA loop가 명시적으로 남았는가?
