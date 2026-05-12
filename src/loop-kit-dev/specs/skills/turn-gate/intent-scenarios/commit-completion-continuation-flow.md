# 커밋 완료 후 Continuation Flow Boundary 시나리오

## 목적

이 시나리오는 커밋 실행이 끝난 뒤 `turn-gate`가 자동으로 턴을 닫지 않는지 확인합니다.
커밋 완료 보고는 explicit stop이 아니며, 보고 이후에는 next-flow question-routing으로 이어져야 합니다.

## 사용자 메시지

```text
일단 커밋하자.
```

## 사용자 메시지의 의미

- 직접 요청된 작업: 현재 승인된 변경 범위를 커밋한다.
- 커밋 전 필요한 것:
  - staged/final status 확인
  - intended diff 확인
  - unrelated change exclusion 확인
  - commit message scope 확인
  - 관련 검증 결과 확인
- 커밋 후 필요한 것:
  - 커밋 hash와 메시지 확인
  - working tree 상태 확인
  - 다음 flow 선택 열기

커밋 완료는 곧 turn stop이 아닙니다.
사용자가 `턴 종료`, `여기서 끝`처럼 명시적으로 말하지 않았다면 `user_explicit_stop=false`를 유지합니다.

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: 커밋 요청의 approval boundary와 커밋 전 확인 조건을 정리한다.
- 소유 산출물: active flow record의 commit scope, staged/final status, unrelated-change exclusion, verification status, commit message draft.
- 완료 기준:
  - 커밋 실행이 승인된 handoff인지 기록한다.
  - commit/push/PR/publish를 서로 다른 approval-sensitive boundary로 구분한다.
  - 커밋 후 terminal summary가 아니라 next-flow question-routing이 필요하다는 점을 기록한다.

## 기대하는 Change-Unit 실행 후보

1. `commit-current-approved-diff`
   - Flow type: `change-unit`
   - 소유 산출물: Git commit object.
   - 근거: 사용자가 명시적으로 커밋 실행을 요청했고, 커밋 전 scope와 staged 상태가 확인된 경우 하나의 실행 단위로 충분하다.
   - 완료 기준: commit hash와 commit message가 확인되고, working tree 상태가 보고된다.
   - 검증 기대: `git log -1`과 `git status --short` 확인.

## 보고 이후 기대 동작

커밋이 완료되면 결과 보고는 다음 맥락을 포함해야 합니다.

- commit hash
- commit subject
- 포함된 변경 범위 요약
- 실행한 검증
- working tree 상태
- 남은 승인 경계: push, PR, publish 등
- 다음 flow 선택지

보고 마지막 상태는 terminal summary가 아니라 active question-routing이어야 합니다.

## Flow가 아닌 항목

- `커밋 완료 보고`
- `최종 요약`
- `턴 종료`
- `push`
- `PR 생성`

이 항목들은 사용자가 별도 요청하거나 승인하지 않는 한 자동 실행 flow로 이어지지 않습니다.
특히 `커밋 완료 보고`와 `최종 요약`은 explicit stop이 없으면 turn closure 근거가 아닙니다.

## 평가 관점

- 커밋 실행 전 staged/final status와 intended diff를 확인한다.
- commit approval을 push/PR/publish approval로 확장하지 않는다.
- 커밋 완료 후 `user_explicit_stop=false`를 유지한다.
- 커밋 완료 보고 뒤 next-flow question-routing을 연다.
- `git status --short`가 깨끗하더라도 턴 종료로 해석하지 않는다.
- visible choices에 턴 종료가 없더라도 flow record의 `Next Flow Options`에는 explicit turn-end option을 기록한다.

## 수용 신호

Fresh executor는 커밋 완료 후 terminal close를 하지 않고 다음 flow 선택을 열어야 합니다.
커밋 완료는 “작업 하나가 끝났다”는 신호이지 “현재 turn이 끝났다”는 신호가 아닙니다.
