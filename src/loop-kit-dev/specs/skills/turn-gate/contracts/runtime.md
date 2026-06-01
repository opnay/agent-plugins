# turn-gate runtime 계약

## 소유 범위

`turn-gate` runtime은 active turn wrapper만 소유합니다.
`flow`가 정의한 work-unit lifecycle, taxonomy, readiness, discovery, ambiguity, checkpoint, template meaning, handoff meaning을 반복하지 않습니다.

## wrapper 계약

visible wrapper는 다음 순서입니다.

- `flow.intake`: 필요한 skill을 다시 읽고 `flow`가 산출한 계약을 적용합니다.
- `flow.reporting`: `flow` 결과, verification status, risk, handoff 또는 next-intake condition을 기록합니다.
- `next-flow`: next action, blocker, self-drive continuation, explicit stop 중 하나로 라우팅합니다.

task completion은 턴을 닫지 않습니다.
source-recorded explicit stop만 terminal close를 허용합니다.

## reporting 계약

`reporting`은 완료 요약이 아니라 다음 사용자 입력을 받기 위한 pre-intake surface입니다.

순서:

1. active flow record와 필요한 `000-plan.md`를 갱신합니다.
2. 결과, 검증 상태, residual risk, required next action을 짧게 보고합니다.
3. 다음 decision surface를 엽니다.

`request_user_input`이 가능하고 선택지가 좁으면 질문 도구를 사용합니다.
불가능하면 active plain-text question fallback을 씁니다.

## record 계약

shared plan, flow record, review template 의미와 파일명 규칙은 `flow`가 소유합니다.
`turn-gate`는 active turn에서 해당 기록을 적용하고 복구합니다.
`turn-gate`가 소유하는 runtime template은 self-drive sidecar뿐입니다.

## phase prefix 계약

turn-gate-owned wrapper progress에는 다음 prefix를 사용합니다.

- `[intake]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

`[framing]`과 `[preparation]`은 visible step이 명시적으로 `flow` 세부 phase일 때만 사용합니다.
prefix는 generated artifact, record, command summary, question option label에 복사하지 않습니다.

## 승인 경계

work 전에는 `flow` decision을 적용하고 approval boundary를 기록합니다.
readiness, verification, build/readback, generated release surface, self-drive, previous context, subagent output은 commit, push, PR, publish, release, version bump, destructive history rewrite, external side effect의 실행 권한을 만들 수 없습니다.

## runtime 본문 경계

Runtime `SKILL.md`는 다음만 직접 포함합니다.

- active-turn rule과 explicit stop authority
- `flow` wrapper 관계
- record 적용과 recovery entrypoint
- reporting-as-pre-intake
- question recovery
- interruption entry-only routing
- verification method/result separation
- self-drive reference discoverability
- approval-sensitive guardrail

상세 decision table은 runtime references나 owning contract로 내립니다.
Runtime `SKILL.md`는 설치된 사용자에게 dev-only `specs/` 경로를 읽으라고 지시하거나 spec-side scenario fixture를 복사하지 않습니다.
