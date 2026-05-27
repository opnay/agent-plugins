# flow core model 계약

## 소유 범위

flow, parent flow, sub-flow candidate, active flow 관계.

## 계약

- flow는 phase checklist가 아니라 응집된 작업 흐름 단위입니다.
- 하나의 flow는 `intake -> framing -> preparation -> work -> verification -> reporting`을 내부 단계로 갖습니다.
- 각 phase의 시작과 종료는 기록 checkpoint를 가지며, `000-plan.md` 또는 active flow record 중 어떤 표면이 최신화돼야 하는지 판단해야 합니다.
- flow intake는 사용자 입력, 목표, 비목표, authority-sensitive signal, discovery topic을 정리합니다.
- flow framing은 flow 분리, candidate-vs-selected 구분, artifact ownership 판단으로 flow contract 초안을 만듭니다.
- flow preparation은 선택된 active flow의 readiness와 ambiguity 판단으로 work 진입 조건을 완성합니다.
- flow work는 current flow boundary 안에서 review-loop, fix-verify-loop, broad-execution 같은 flow-local strategy를 선택할 수 있습니다.
- flow가 너무 크거나 여러 산출물을 만들면 parent flow는 finite `sub-flow candidates`를 만들 수 있습니다.
- `sub-flow candidate` 생성은 실행이 아닙니다.
- 후보는 `turn-gate` next-flow 질문 또는 명시적으로 준비된 self-drive sequence에 의해 선택될 때만 active flow가 됩니다.
- parent flow가 sub-flow 후보를 만드는 경우, parent flow의 산출물은 후보 목록, 각 후보의 scope/non-goals/completion criteria/verification expectation/handoff 조건, unresolved question입니다.
- 각 sub-flow는 선택되면 독립적인 flow로 취급하며, 자기 내부의 intake/framing/preparation/work/verification/reporting을 다시 가집니다.
- flow가 끝났다는 사실은 turn이 끝났다는 뜻이 아닙니다. turn-level next-flow reopening은 `turn-gate`가 소유합니다.
- flow-local strategy는 여러 flow를 자동으로 이어가는 self-drive sequence authority가 아닙니다.

## 검토 기준

- parent flow가 finite 후보 목록을 만들었는가?
- 후보 생성과 후보 실행을 혼동하지 않았는가?
- active flow 전환 권한이 `turn-gate` 또는 준비된 self-drive sequence에 남아 있는가?
- flow-local strategy와 turn-level continuation authority를 혼동하지 않았는가?
- 각 phase 시작/종료의 기록 checkpoint가 flow-local 상태와 turn-level active pointer를 구분하는가?
