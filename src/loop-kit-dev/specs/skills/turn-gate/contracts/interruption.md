# turn-gate interruption 계약

## 소유 범위

active turn 중 새 사용자 메시지가 도착한 경우의 entry-only routing.

## 계약

새 메시지가 오면 active turn continuity, pending question, approval boundary, verification status, required next action을 보존합니다.
계약 영향 판단은 `flow`에 맡기고, `turn-gate`는 routing만 적용합니다.

결과는 하나만 기록합니다.

- inline-answer
- current-flow-revision
- background-current-flow
- reserve-later-analysis
- supersede-current-flow
- blocker-question
- explicit-stop

interruption은 active contract 밖 work 권한이나 approval-sensitive action 권한을 만들지 않습니다.
