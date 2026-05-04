# turn-gate question-routing sub-spec

## 목적

이 문서는 `turn-gate`의 user-gated question routing, `request_user_input`, next-flow reopening, fallback, turn-end option 기록 계약을 소유합니다.

## 핵심 계약

- clarification, choices, scope locks, mode narrowing, next-flow decisions는 user-gated question routing으로 처리한다.
- 구조적 선택지를 줄 수 있고 도구가 사용 가능하면 `request_user_input`을 사용한다.
- result reporting 뒤 explicit stop이 없다면 active question-routing으로 next-flow를 다시 연다.
- next-flow choices는 방금 보고한 결과와 직접 연결된 좁은 선택지여야 한다.
- plain text follow-up이나 generic closing phrase는 next-flow reopening을 대체하지 못한다.
- result report에는 routing 또는 phase selection에 영향을 준 material judgment call이 있으면 함께 드러낸다.

## Fallback

- `request_user_input`이 unavailable한 경우에만 fallback을 사용한다.
- fallback result report에는 도구가 unavailable인 사실, 열린 선택지, record에 남긴 required next action을 명시한다.
- fallback은 terminal summary가 아니라 active question-routing 상태로 남아야 한다.

## Turn-End Option

- 사용자가 명시적으로 턴을 종료하자고 요청하지 않으면 clean stop을 기본 경로로 두지 않는다.
- 사용자에게 보이는 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하는 경우에도 visible prompt에는 명시적 stop 요청이 가능하다는 사실을 드러낸다.
- flow record의 `Next Flow Options`에는 별도 turn-end option을 항상 기록한다.
- explicit stop 처리와 closure source 기록은 `runtime-flow.md`와 `session-records.md`가 함께 소유한다.

## 검토 질문

- 결과 보고 뒤 `request_user_input` 또는 active question-routing으로 next-flow를 열었는가?
- fallback을 썼다면 도구 unavailable 사실과 required next action을 기록했는가?
- visible choices가 result report와 직접 연결되어 있는가?
- visible choices에 stop이 없더라도 flow record에는 turn-end option이 남아 있는가?
