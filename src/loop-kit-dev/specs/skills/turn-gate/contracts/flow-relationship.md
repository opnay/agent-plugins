# turn-gate flow 관계 계약

## 소유 범위

이 문서는 `turn-gate`가 `flow skill`을 wrapper로 적용하는 경계를 소유합니다.

## 계약

- `turn-gate`는 사용자 메시지와 next turn-flow를 `flow skill`에 연결합니다.
- `flow skill` 내부의 `flow.message`, `flow.main-flows`, `flow.end` 의미는 `flow`가 소유합니다.
- `turn-gate`는 active turn continuity, record 적용, question routing, verification routing, explicit stop, self-drive gate를 소유합니다.
- `turn-gate`는 flow boundary, readiness, discovery, flow-local strategy, shared record template 의미, handoff 의미를 재정의하지 않습니다.
- plan, flow record, review template 의미와 파일명 규칙은 `flow`가 소유합니다.

## 검토 기준

- `turn-gate`가 `flow` 산출물을 적용만 하는가?
- self-drive 역방향이 `flow` 재정의가 아니라 next turn-flow gate로 처리되는가?
- 종료 요청이 `flow` 자체 판단으로 임의 종료되지 않는가?
