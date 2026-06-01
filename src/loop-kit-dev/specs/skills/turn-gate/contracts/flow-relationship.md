# turn-gate flow 관계 계약

- `turn-gate`는 `flow`가 산출한 active flow, candidate, phase, handoff, next intake condition을 적용합니다.
- `turn-gate`는 active turn 유지, 기록, 질문 라우팅, explicit stop, self-drive gate를 소유합니다.
- `turn-gate`는 flow boundary, readiness, discovery, flow-local strategy, handoff 의미를 재정의하지 않습니다.
- active flow 도중 새 사용자 메시지가 들어오면 `turn-gate`는 entry를 열고, contract impact 판단은 `flow` 산출물에 따릅니다.
