# turn-gate flow record template sub-spec

## 목적

이 문서는 `turn-gate`의 `skills/turn-gate/templates/flow-record-template.md` 형식 계약을 소유합니다.

`001+` flow record는 하나의 사용자 요청 기반 flow가 실제로 무엇을 소유했고, 무엇을 했고, 어떻게 검증됐는지를 기록하는 canonical detail artifact입니다.

## 소유

- original user request message
- task, flow type, flow scope, parent plan
- current phase
- `Continuity Guard`
- flow contract
- execution log
- verification detail
- report
- next-flow options
- residual risk

## 비소유

- date-level work history
- all-day user request index
- completed flow index
- planned flow sequence 전체의 top-level 관리

위 항목은 `templates/plan.md`가 소유하는 `000-plan.md`에 둡니다.

## 필수 구조

`flow-record-template.md`는 다음 구조를 유지합니다.

1. `Metadata`
2. `Continuity Guard`
3. `Flow Contract`
4. `Optional Risky Actions`
5. `Execution Log`
6. `Verification`
7. `Report`
8. `Next Flow Options`
9. `Residual Risk`

## Continuity Guard 필수 필드

`Continuity Guard`에는 다음 필드가 남아야 합니다.

- `Turn-gate active`
- `Question-routing mode`
- `User explicit stop`
- `Terminal summary allowed`
- `Required next action`
- `Last refreshed phase`
- `Confirmed closure`
- `Closure source message`
- `Closure recorded phase`
- `Pending question state`
- `Pending question id or summary`
- `Superseded question id or summary`
- `Verification status`
- `Continuity note`

`Closure source message`와 `Closure recorded phase`는 explicit stop이 source-recorded된 경우에만 terminal close 근거가 됩니다.
source 없는 closure 또는 stale `Terminal summary allowed: yes`를 발견하면 `User explicit stop: no`, `Terminal summary allowed: no`로 복구하고, `Continuity note`에 stale 상태였음을 남깁니다.

## Flow Contract 필수 필드

`Flow Contract`는 다음 필드를 소유합니다.

- user request
- preparation source/result
- boundary rationale
- current blocker
- scope lock status
- work boundary
- non-goals
- acceptance signal
- expected risky actions
- approval boundary
- user-gated checkpoints
- self-drive eligibility
- verification expectation
- material judgment calls

## 중복 방지 규칙

- `Analysis`라는 broad bucket을 사용하지 않습니다. scope, routing, approval, verification expectation은 `Flow Contract`로 모읍니다.
- `Plan`과 `Work`를 별도 top-level section으로 분리하지 않고 `Execution Log` 아래에 둡니다.
- `Files changed`와 `Files or surfaces touched`를 동시에 쓰지 않습니다. 최종 표면은 `Changed surfaces`로 기록합니다.
- `Why this mode`, `Why this question-routing mode`, `Next-flow context` 같은 반복 필드는 기본 템플릿에 두지 않습니다. 필요하면 `Material judgment calls`나 `Report`에 한 줄로 남깁니다.
- `Optional Risky Actions` 섹션은 유지합니다. risky action이 없으면 `not-applicable`로 접고 checklist를 펼치지 않습니다.
- risky action이 현재 flow에서 예상되지만 아직 실행 승인이 없으면 checklist를 펼치고 `Initial agreement`를 `not-approved`, `deferred`, 또는 `handoff-required`로 기록합니다.

## Next Flow Options

- `Next Flow Options`는 flow record가 소유합니다.
- visible `request_user_input` 선택지에 turn-end option이 보이지 않아도 record에는 explicit turn-end option을 남깁니다.
- 선택 결과나 active next flow pointer만 `000-plan.md` snapshot에 반영합니다.

## 검토 질문

- flow record가 해당 flow의 canonical detail artifact로 충분한가?
- `Continuity Guard`가 explicit stop, stale closure, pending question, verification status를 판별할 수 있는가?
- source 없는 closure나 stale terminal summary allowance를 복구하는 행동이 기록되는가?
- `Flow Contract`가 work 전 scope와 approval boundary를 충분히 잠그는가?
- `Execution Log`가 plan/work/evidence를 지나치게 흩뜨리지 않고 담는가?
- report 뒤 `Next Flow Options`가 explicit stop 없는 흐름을 다시 열 수 있는가?
