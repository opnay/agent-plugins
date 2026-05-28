# turn-gate interruption 계약

## 소유 범위

active flow가 진행 중일 때 새 사용자 메시지가 도착한 경우의 entry-only routing.

## 계약

`interruption`은 lifecycle phase가 아닙니다.
새 메시지가 오면 현재 phase, scope, non-goals, approval boundary, verification status, required next action을 보존한 뒤 `flow` contract-impact 판단을 적용합니다.

결과는 하나만 기록합니다.

- `inline-answer`: 계약 변경 없이 답변하고 보존된 phase로 복귀
- `current-flow-revision`: 현재 flow 계약을 갱신하고 `framing` 또는 `preparation`으로 복귀
- `background-current-flow`: 현재 flow를 background로 보존하고 새 foreground flow 시작
- `reserve-later-analysis`: 지금 처리하지 않을 후보를 기록하고 보존된 phase로 복귀
- `supersede-current-flow`: 현재 flow를 superseded로 기록하고 새 flow 시작
- `blocker-question`: 계속하면 위험한 결정, 승인, 접근, scope gap을 질문으로 전환
- `explicit-stop`: stop source를 기록한 뒤에만 terminal closure 허용

`interruption`은 active contract 밖의 work 권한을 만들지 않습니다.
commit, push, PR, publish, release, version bump, destructive action은 별도 명시 승인 없이는 실행할 수 없습니다.

## 검토 기준

- 새 메시지가 active flow 중 도착한 interruption인가?
- 복귀할 phase와 required next action을 보존했는가?
- 계약 변경 여부와 새 flow 여부를 `flow` 판단에 맡겼는가?
- 결과를 하나만 기록하고 즉시 lifecycle로 복귀했는가?
