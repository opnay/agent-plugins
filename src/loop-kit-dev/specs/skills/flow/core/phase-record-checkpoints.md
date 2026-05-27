# flow phase record checkpoint 계약

## 소유 범위

active flow의 `intake -> framing -> preparation -> work -> verification -> reporting` phase 시작과 종료에 필요한 기록 checkpoint.

## 계약

- flow는 phase checklist가 아니지만, active flow의 각 phase는 시작과 종료 시점에 기록 checkpoint를 가져야 합니다.
- phase 시작 checkpoint는 현재 phase, scope boundary, required next action, pending question/blocker 상태가 `000-plan.md` 또는 active flow record에 반영돼야 하는지 판단하게 해야 합니다.
- phase 종료 checkpoint는 phase 결과, 다음 phase, verification status 변화, residual risk, handoff 또는 next-flow 조건이 `000-plan.md` 또는 active flow record에 반영돼야 하는지 판단하게 해야 합니다.
- active flow가 바뀌거나 turn-level required next action이 바뀌면 `000-plan.md` 갱신이 필요합니다.
- active flow 또는 planned sequence가 특정 skill 적용을 요구하면 `000-plan.md`에는 사용할 skill 목록이 드러나야 합니다.
- turn-gate-managed 사용자 메시지 flow가 preparation으로 들어가면 `turn-gate`와 `flow`를 다시 읽고 active skill list에 포함해야 합니다.
- skill 목록은 현재 선택된 flow와 준비된 future flow에 필요한 것만 담습니다. 후보 단계에서 나온 가능성만으로 active skill처럼 기록하지 않습니다.
- 같은 active flow 내부의 phase 상태, execution log, verification evidence, report outcome, residual risk가 바뀌면 active flow record 갱신이 필요합니다.
- record checkpoint는 실제 기록 적용을 위한 계약입니다. `flow`는 어떤 기록이 필요한지 산출하고, `turn-gate`는 active turn 안에서 그 기록을 유지합니다.
- checkpoint는 phase를 별도 flow로 만들지 않습니다. `intake`, `framing`, `preparation`, `work`, `verification`, `reporting`은 여전히 같은 active flow 내부의 phase입니다.
- trivial read-only judgment처럼 기록 변경이 필요 없다고 판단하는 경우에도 이유가 active flow record나 report에 남아야 합니다.

## 검토 기준

- phase 시작 때 현재 phase와 required next action이 기록에서 재구성 가능한가?
- phase 종료 때 결과와 다음 phase가 기록에서 재구성 가능한가?
- `000-plan.md`를 바꿔야 하는 turn-level pointer 변화와 active flow record만 바꾸면 되는 flow-local 상태 변화를 구분했는가?
- `000-plan.md`에 현재 flow 또는 planned sequence에서 사용할 skill 목록이 필요한 만큼만 드러나는가?
- checkpoint를 이유로 phase 자체를 새 flow로 오해하지 않았는가?
