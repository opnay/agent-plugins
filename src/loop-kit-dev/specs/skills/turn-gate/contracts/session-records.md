# turn-gate session-records 계약

## 소유 범위

`turn-gate`는 active turn을 복구할 수 있을 만큼의 record 적용과 복구를 소유합니다.
record는 메인 그래프 노드가 아니며, `flow skill: interview -> flow skill: handoff -> 질문 도구` 루프를 보조합니다.
shared record template 의미와 파일명 규칙은 `flow`가 소유합니다.

## 적용 계약

- plan record: active flow pointer, next action, closure flags, self-drive pointer, active skill list, pending/answered question state를 적용합니다.
- flow record: handoff 뒤 질문 라우팅, verification status, result, risk, next action을 복구 가능하게 유지합니다.
- review record: `flow`가 소유한 회고 기록을 active turn에서 필요한 경우 적용합니다.
- self-drive record: `turn-gate`가 template과 sidecar gate를 소유합니다.

`turn-gate`는 sibling skill filesystem path를 runtime 지시로 만들지 않습니다.
`turn-gate`가 소유하는 bundled template은 self-drive sidecar뿐입니다.

## 복구 계약

- record가 없으면 현재 라우팅을 복구하는 데 필요한 최소 record만 생성합니다.
- 있어야 하는 active record가 없거나 접근 불가하면 blocker recovery로 라우팅합니다.
- stale closure state는 reset하고 recovery를 기록합니다.
- source-recorded explicit stop만 terminal closure authority입니다.
- pending question은 `answered_question`과 `pending_question`으로 복구합니다.
- `flow skill: handoff` 뒤 `next-flow gate`를 통과할 때마다 `000-plan.md`를 업데이트합니다.

## 검토 기준

- compaction 뒤 active flow, next action, question state, verification status, explicit stop state가 복구되는가?
- shared template 의미가 `turn-gate`에 중복되지 않는가?
- read-only source work가 session record 적용 금지로 오해되지 않는가?
