# turn-gate runtime 계약

## 소유 범위

`turn-gate` runtime은 graph wrapper를 소유합니다.
`flow skill` 내부 의미인 `flow.message -> flow.main-flows -> flow.end`는 `flow`가 소유합니다.

## wrapper 계약

- entry: 사용자 메시지를 `flow skill`에 적용합니다.
- exit: `flow.end` 이후 `next turn-flow / 메시지 수신`을 엽니다.
- loop: 일반 모드는 다음 메시지를 기다리고, self-drive 모드는 자체 해석으로 다시 `flow skill`에 들어갑니다.
- stop: 종료 요청은 전 과정에서 감지하고 source-recorded explicit stop으로만 현재 turn을 닫습니다.

## runtime 본문 경계

Runtime `SKILL.md`는 다음만 직접 포함합니다.

- active-turn rule
- `flow` wrapper dependency
- next turn-flow reopening
- explicit-stop authority
- question recovery
- record application/recovery entrypoint
- verification method/result separation
- self-drive gate
- approval-sensitive guardrail

Runtime은 flow taxonomy, flow lifecycle, shared template shape, readiness/discovery/ambiguity, handoff meaning을 반복하지 않습니다.

## phase prefix 계약

turn-gate-owned wrapper progress에는 `[intake]`, `[work]`, `[verification]`, `[reporting]`, `[next-flow]`를 사용합니다.
`[framing]`과 `[preparation]`은 visible step이 명시적으로 `flow` 세부 phase일 때만 사용합니다.
prefix는 generated artifact, record, command summary, question option label에 복사하지 않습니다.
