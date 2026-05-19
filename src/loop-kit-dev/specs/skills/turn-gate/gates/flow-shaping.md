# turn-gate flow-shaping gate sub-spec

## 목적

이 문서는 `turn-gate` 안에서 flow shaping gate가 sibling `flow` 계약을 적용하는 전환 계약을 소유합니다.

flow shaping gate는 현재 사용자 요청과 flow 상태를 active flow에 반영합니다.

## 소유

- current turn에 source-recorded active flow가 있는지 확인한다.
- active flow가 없거나 current request가 active flow 경계를 바꿀 수 있으면 sibling `flow` contract decision을 요구한다.
- `flow` decision 결과를 current active flow, next-flow routing, blocker, 또는 report-only handoff로 연결한다.
- 후보를 active flow로 바꾸기 전에는 `turn-gate` next-flow question routing 또는 준비된 self-drive sequence를 요구한다.

## 비소유

- flow boundary/type/completion 판단 자체
- flow 내부 command sequence 실행
- verification pass 판정
- explicit stop 없는 turn closure

flow shaping gate는 flow logic을 재판정하지 않고, sibling `flow` decision 없이 실행으로 넘어가지 않게 막아야 합니다.

## 검토 질문

- active flow가 source-recorded 되어 있거나 sibling `flow` decision이 있는가?
- 후보를 active flow로 바꾸는 결정을 next-flow routing 또는 prepared self-drive sequence로 처리했는가?
