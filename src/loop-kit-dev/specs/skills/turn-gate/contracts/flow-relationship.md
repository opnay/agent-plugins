# turn-gate flow 관계 계약

## 소유 범위

이 문서는 `turn-gate`가 `flow skill`을 wrapper로 적용하는 경계를 소유합니다.

## 계약

- 사용자 메시지는 `turn-gate` wrapper 안의 `flow skill` 그룹으로 진입합니다.
- `turn-gate`는 handoff 뒤 `next-flow gate` 루프를 `flow skill`에 연결합니다.
- `flow skill` 내부 의미는 `flow`가 소유합니다.
- `turn-gate`는 active turn continuity, question routing, explicit stop, self-drive gate를 소유합니다.
- record, verification, interruption, date 처리는 `turn-gate` 메인 그래프를 보조하는 라우팅 계약입니다.
- `turn-gate`는 flow boundary, readiness, discovery, flow-local strategy, shared record template 의미, handoff 의미를 재정의하지 않습니다.

## 검토 기준

- `turn-gate`가 `flow` 산출물을 적용만 하는가?
- `next-flow gate`가 skill reconfigure, next flow selection, optional self-drive sidecar update, plan update를 거쳐 `flow skill: interview`로 되돌리는가?
- self-drive 경로가 `flow` 재정의가 아니라 준비된 sequence gate로 처리되는가?
- 종료 요청이 `flow` 자체 판단으로 임의 종료되지 않는가?
