# turn-gate flow-boundaries delegation sub-spec

## 목적

이 문서는 `turn-gate`가 sibling `flow` skill의 flow boundary 계약을 어떻게 적용하는지 소유합니다.

flow 정의, 후보, flow-vs-phase, flow type, completion criteria, verification expectation 산출은 sibling `flow` skill과 그 spec이 소유합니다.
`turn-gate`는 이 정의를 재소유하지 않고, active turn에서 flow 없이 진행하지 못하게 하는 gate 역할만 유지합니다.

## 적용 계약

- active turn에서 work를 시작하려면 current flow가 source-recorded 되어 있거나, sibling `flow` 계약이 산출한 flow decision을 먼저 받아야 합니다.
- `turn-gate`는 `flow`가 산출한 decision을 session record와 next-flow routing에 적용합니다.
- `turn-gate`는 후보를 active flow로 전환하는 user-gated next-flow routing과 prepared self-drive continuation만 소유합니다.
- flow reporting 뒤 explicit stop이 없으면 `turn-gate`는 next-flow reopening으로 이어갑니다.

## 비소유

- flow taxonomy 자체 정의
- 후보의 내부 완료 기준
- phase와 flow의 일반 판정 규칙
- flow 설계 output contract

## Session Record와의 관계

`records/session-records.md`는 sibling `flow` decision을 `.agents/sessions/{YYYYMMDD}/000-plan.md`와 개별 flow record에 기록하는 runtime fields만 소유합니다.
어떤 항목이 flow인지, 후보인지, phase인지 판단하는 기준은 `flow`가 소유합니다.

## 검토 질문

- active turn에서 sibling `flow` decision 없이 work를 시작하지 않았는가?
- 후보를 active flow로 전환할 때 next-flow routing 또는 prepared self-drive sequence를 사용했는가?
- flow reporting 뒤 explicit stop 없이 turn을 닫지 않았는가?
