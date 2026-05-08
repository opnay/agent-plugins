# turn-gate gate-continuation sub-spec

## 목적

이 문서는 continuation gate의 전환 계약을 소유합니다.

continuation gate는 reporting 이후의 전환을 결정합니다.

## 소유

- source-recorded explicit stop이 있는지 확인한다.
- explicit stop이 없으면 active question-routing, loop continuation, planned self-drive handoff, blocker decision 중 하나로 이어간다.
- `request_user_input` 사용 가능성과 fallback 필요성을 판단한다.
- session record의 `Next Flow Options`에 explicit turn-end option을 남긴다.

## 비소유

- 새로운 work를 질문 없이 확장
- commit/push/PR/publish 승인 추론
- stale closure 또는 source-less closure를 terminal close 근거로 사용

continuation gate는 task 결과와 reporting 결과를 turn 종료로 해석하지 않습니다.

## 검토 질문

- continuation gate가 explicit stop 없는 흐름을 next-flow reopening 또는 active question-routing으로 이어가는가?
- stale closure를 terminal close 근거로 사용하지 않는가?
