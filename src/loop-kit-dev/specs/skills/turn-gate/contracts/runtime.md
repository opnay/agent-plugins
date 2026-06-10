# turn-gate runtime 계약

## 소유 범위

`turn-gate` runtime은 graph wrapper를 소유합니다.
`flow skill` 내부 의미는 `flow`가 소유하며, `turn-gate` runtime은 이를 재정의하지 않고 적용합니다.

## wrapper 계약

- entry: 사용자 메시지는 `turn-gate` wrapper 안의 `flow skill` 그룹으로 진입합니다.
- exit: `flow skill: handoff` 이후 `next-flow gate`를 엽니다.
- handoff priority: `flow skill: handoff`는 terminal closure가 아니라 `next-flow gate`의 입력입니다. `turn-gate` 활성 중에는 final-looking 보고, status-only 답변, verification pass, commit completion, flow reporting이 handoff 뒤 routing을 닫을 수 없습니다.
- gate tag: Runtime `SKILL.md`는 handoff 뒤 필수 gate 경계를 `<gate:next-flow>...</gate:next-flow>` 태그로 감쌉니다.
- loop: 일반 모드는 `next-flow gate`에서 `skill reconfigure` 그룹을 거쳐 질문 도구로 `다음 플로우 선택 -> 000-plan.md 업데이트`를 수행하고, `flow: deep-interview`와 같은 인터뷰 흐름으로 충분히 구체화한 뒤 `flow skill: interview`에 들어갑니다.
- self-drive: 명시적으로 준비된 sequence gate가 통과한 경우에만 질문 도구를 대체해 `다음 플로우 선택`에 진입하고, `000-self-drive.md 업데이트 -> 000-plan.md 업데이트`를 거칩니다.
- stop: 종료 요청은 `turn-gate / 메인`의 모든 시점에서 감지하고 종료 페이즈로 이동합니다.
- stop phase: `작업 중이던 플로우 정리 -> explicit-stop 기록 - active turn 종료` 순서로 처리합니다.
- stop authority: source-recorded explicit stop으로만 active turn을 닫습니다.
- default continuation: 새 사용자 메시지, 질문, 상태 확인, 작업 변경, 방향 전환, 오류 지적, 추가 요구는 기본적으로 active turn 안의 열린 입력입니다. 현재 메시지가 명시적 종료 요청이 아니면 terminal closeout으로 닫지 않습니다.
- wrapper precedence: `flow`는 wrapper 안에서만 실행되며, `turn-gate`가 활성인 동안 `flow` handoff 결과는 항상 `next-flow gate`에서 소비된 뒤 다음 routing 상태로 기록됩니다.

## runtime 본문 경계

Runtime `SKILL.md`는 다음만 직접 포함합니다.

- active-turn rule
- `flow` wrapper dependency
- handoff routing reopening
- explicit-stop authority
- question recovery
- default continuation for non-stop messages
- record recovery entrypoint
- non-pass routing
- self-drive gate
- approval-sensitive guardrail

Runtime은 flow taxonomy, flow lifecycle, shared template shape, readiness/discovery/ambiguity, handoff meaning을 반복하지 않습니다.

## phase prefix 계약

Visible progress label은 wrapper 상태를 돕는 표시일 뿐 메인 그래프 노드가 아닙니다.
사용자-facing phase 시작 또는 의미 있는 진행 메시지는 현재 단계 prefix로 시작합니다.

`turn-gate`는 `flow`가 산출한 phase prefix를 재정의하지 않고 적용합니다.
`turn-gate`가 직접 소유하는 prefix는 `[next-flow]`이며, handoff 뒤 다음 flow 선택, 질문 도구, self-drive continuation을 여는 메시지에 사용합니다.

prefix나 lifecycle label을 artifact body, record body, raw command output, command output summary, question option label에 기계적으로 복사하지 않습니다.
이미 prefix가 붙은 사용자-facing 메시지 안의 모든 문장이나 bullet에 prefix를 반복하지 않습니다.
