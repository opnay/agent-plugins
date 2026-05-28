# turn-gate date-authority 계약

## 소유 범위

turn-gate 안에서 상대 날짜 표현을 해석하고, session record나 이전 flow 기록의 날짜와 충돌할 때 clarification으로 라우팅하는 기준을 소유합니다.

이 계약은 turn-gate runtime 판단 기준입니다.
repo-wide 날짜 정책, AGENTS 전역 규칙, 외부 서비스 calendar/timezone 정책은 소유하지 않습니다.

## 기본 authority

상대 날짜 표현은 기본적으로 현재 실행 환경의 시스템 날짜와 timezone을 기준으로 해석합니다.

상대 날짜 표현에는 다음이 포함됩니다.

- 오늘
- 내일
- 어제
- 이번 주
- 지난 기록
- 마지막 기록
- 이전 flow

날짜가 결과, target, verification path, reporting scope, 또는 기록 재구성에 영향을 주면 절대 날짜를 함께 씁니다.

## 기록 기반 맥락

사용자 메시지가 session record, 이전 flow, 마지막 기록, 어제 작업, 또는 기록 기반 재개를 가리키면 시스템 날짜만으로 단정하지 않습니다.

다음 중 어느 의미인지에 따라 결과가 달라질 수 있으면 user-gated clarification으로 라우팅합니다.

- 시스템 현재일 기준의 상대 날짜
- 마지막 session record가 있는 날짜
- 이전 대화 또는 이전 flow의 날짜
- 사용자가 기억하는 업무 날짜

## 충돌 처리

session record 경로, 파일명, frontmatter 날짜, git 날짜는 기록 evidence입니다.
이 값들은 사용자 상대 날짜와 충돌해도 사용자 의도를 자동으로 대체하지 않습니다.

충돌이 target, verification path, reporting scope, 또는 기록 재구성을 바꾸지 않으면 시스템 날짜 기준으로 해석하고 사용한 기준을 짧게 기록합니다.

충돌이 결과에 영향을 줄 수 있으면 work 전에 clarification 또는 blocker routing을 엽니다.

## 보고 기준

사용자가 상대 날짜를 썼고 날짜가 결과에 영향을 주면 보고에서 절대 날짜를 함께 씁니다.

예:

- `오늘(2026-05-28)`
- `어제(2026-05-27, 시스템 날짜 기준)`
- `마지막 기록 날짜인지 시스템 기준 어제인지 확인 필요`

## 검토 기준

- 상대 날짜를 시스템 날짜/timezone 기준으로 해석했는가?
- 기록 기반 맥락이 있으면 기록 날짜가 시스템 날짜를 조용히 대체하지 않게 했는가?
- 날짜 source 차이가 target, verification path, reporting scope를 바꾸면 clarification으로 라우팅했는가?
- 결과나 기록 재구성에 영향을 주는 날짜는 절대 날짜로 남겼는가?
- repo-wide 정책으로 범위를 넓히지 않았는가?
