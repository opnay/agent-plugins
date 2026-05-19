# flow core model

## 목적

이 문서는 `flow`의 핵심 모델인 flow, parent flow, sub-flow candidate, active flow 관계를 소유합니다.

## 계약

- flow는 phase checklist가 아니라 응집된 작업 흐름 단위입니다.
- 하나의 flow는 `preparation -> work -> verification -> reporting`을 내부 단계로 갖습니다.
- flow가 너무 크거나 여러 산출물을 만들면 parent flow는 finite `sub-flow candidates`를 만들 수 있습니다.
- `sub-flow candidate` 생성은 실행이 아닙니다.
- 후보는 `turn-gate` next-flow 질문 또는 명시적으로 준비된 self-drive sequence에 의해 선택될 때만 active flow가 됩니다.
- parent flow가 sub-flow 후보를 만드는 경우, parent flow의 산출물은 후보 목록, 각 후보의 scope/non-goals/completion criteria/verification expectation/handoff 조건, unresolved question입니다.
- 각 sub-flow는 선택되면 독립적인 flow로 취급하며, 자기 내부의 preparation/work/verification/reporting을 다시 가집니다.
- flow가 끝났다는 사실은 turn이 끝났다는 뜻이 아닙니다. turn-level next-flow reopening은 `turn-gate`가 소유합니다.

## 검토 질문

- parent flow가 finite 후보 목록을 만들었는가?
- 후보 생성과 후보 실행을 혼동하지 않았는가?
- active flow 전환 권한이 `turn-gate` 또는 준비된 self-drive sequence에 남아 있는가?
