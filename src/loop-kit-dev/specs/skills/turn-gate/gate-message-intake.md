# turn-gate gate-message-intake sub-spec

## 목적

이 문서는 message intake gate의 전환 계약을 소유합니다.

message intake gate는 현재 incoming user message의 라우팅 사실을 분류합니다.

## 소유

- 현재 메시지가 explicit turn stop인지 판정한다.
- explicit stop이 아니라면 continuation input으로 분류한다.
- 새 작업, correction, review, status check, approval-sensitive request, handoff result, next-task request 같은 message kind를 파악한다.
- operation/target ambiguity, scope gap, approval boundary 가능성을 표시한다.

## 비소유

- active flow completion criteria 확정
- 파일 수정, 검증 실행, build, commit, push, PR, publish
- flow 완료나 turn closure 결정

명시적 stop이 아닌 메시지는 terminal close 근거가 아닙니다.

## 검토 질문

- message intake gate가 실행이나 flow completion을 직접 결정하지 않는가?
- explicit stop이 아닌 메시지를 continuation input으로 분류했는가?
