# turn-gate interruption 계약

## 소유 범위

active turn 중 새 사용자 메시지가 도착한 경우의 entry-only routing.

## 계약

새 메시지가 오면 active turn continuity, pending question, approval boundary, verification status, required next action을 보존합니다.
새 메시지의 기본값은 종료가 아니라 active turn 안의 열린 입력입니다.
작업 변경, 질문, 상태 확인, 방향 전환, 오류 지적, 추가 요구는 source-recorded explicit stop이 아니면 turn 종료 신호로 쓰지 않습니다.
계약 영향 판단은 `flow`에 맡기고, `turn-gate`는 routing만 적용합니다.

결과는 하나만 기록합니다.

- active-flow
- current-flow-revision
- background-current-flow
- reserve-later-analysis
- supersede-current-flow
- blocker-question
- explicit-stop

`explicit-stop`은 현재 메시지가 active turn 자체를 끝내려는 뜻을 명확히 드러낼 때만 사용합니다.
그 외에는 필요한 경우 `flow skill: interview` 또는 handoff 뒤 질문 라우팅으로 돌아가 입력을 다시 구체화합니다.
interruption은 active contract 밖 work 권한이나 approval-sensitive action 권한을 만들지 않습니다.
