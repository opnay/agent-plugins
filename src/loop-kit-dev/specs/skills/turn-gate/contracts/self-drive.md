# turn-gate self-drive 계약

## 소유 범위

self-drive는 `next turn-flow / 메시지 수신`에서 사용자 메시지 없이 자체 해석으로 다시 `flow skill`에 들어갈 수 있는지 판단하는 gate입니다.

## 계약

self-drive는 명시 요청 또는 next-flow mode 선택으로만 활성화합니다.
긴 task list, pass verification, subagent availability, "continue"만으로 추론하지 않습니다.

sidecar는 다음을 복구 가능하게 둡니다.

- mode
- source-backed goal
- current identity
- active flow record
- next action
- endpoint 또는 stop condition
- acceptance signal
- verification expectation
- approval checkpoints
- blocker return condition
- ledger

advance 조건:

- current flow verification is `pass`
- `flow` handoff is not blocked
- next identity is known
- approval boundary still matches
- plan and sidecar gate pass

non-pass verification, blocker, approval need, stale sidecar, endpoint/scope/target/order/acceptance change는 autonomous advance보다 먼저 user-gated routing으로 돌아갑니다.

## 검토 기준

- self-drive가 명시 없이 시작되지 않는가?
- next turn-flow 역방향이 recorded sidecar gate를 통과하는가?
- approval-sensitive action을 자체 승인하지 않는가?
