# turn-gate self-drive 계약

## 소유 범위

self-drive는 사용자가 명시적으로 맡긴 범위 안에서 다음 flow 또는 loop를 자율 진행할 수 있는지 판단하는 overlay입니다.
기본 `flow.intake -> flow.reporting -> next-flow` wrapper, 기록, 검증, 승인 checkpoint를 대체하지 않습니다.

모드는 둘입니다.

- `finite`: 준비된 flow sequence를 순서대로 진행합니다.
- `infinite`: 사용자가 멈출 때까지 반복하되 매번 하나의 bounded iteration만 준비하고 실행합니다.

self-drive는 commit, push, PR, publish, release, version bump, destructive/external action, scope expansion 권한을 만들지 않습니다.

## 활성화 계약

self-drive는 명시 요청 또는 next-flow mode 선택으로만 활성화합니다.
긴 task list, successful verification, subagent availability, 이전 대화의 의욕적 표현, 또는 “continue”만으로 추론하지 않습니다.

공통 sidecar state:

- `status`
- `mode`
- source-backed goal 또는 objective
- current flow 또는 loop identity
- active flow record identity
- `next_action`
- endpoint 또는 stop condition
- acceptance signal
- verification expectation
- allowed autonomous actions
- approval checkpoints
- blocker return conditions
- ledger

`finite`는 `active_flow_index`, `current_flow_label`, `planned_flow_count`, prepared sequence를 추가로 가집니다.
각 flow identity와 handoff condition은 `flow` output contract가 원천입니다.

`infinite`는 `loop_count`, `current_loop_label`, next bounded iteration을 가집니다.
큰 todo list로 infinity를 표현하지 않습니다.

infinite mode 첫 반복은 `loop_count: 1`로 시작합니다.
bounded target이 없으면 자동 진행하지 않고 target selection 또는 blocker routing으로 돌아갑니다.

## Sidecar Gate 계약

각 self-drive flow 또는 loop 시작 때 `000-plan.md`와 `000-self-drive.md`를 읽습니다.

확인 항목:

- plan self-drive status와 sidecar pointer
- mode
- current identity
- active flow record identity
- `next_action`
- endpoint 또는 stop condition
- verification expectation
- approval checkpoint
- blocker state

`000-plan.md`는 self-drive active 여부와 sidecar pointer만 저장합니다.
sequence, loop, ledger, approval checkpoint, endpoint detail은 `000-self-drive.md`가 소유합니다.
flow-local evidence는 active flow record가 소유합니다.

누락, 충돌, stale state, 범위 이탈이 있으면 자동 진행을 멈추고 user-gated recovery로 돌아갑니다.
숫자 index나 loop count만 맞는 상태는 충분하지 않습니다.
finite에서 `active_flow_index > planned_flow_count`이면 stale/corrupt sidecar입니다.
endpoint 판단 중 record나 sidecar를 읽을 수 없으면 access blocker로 보고하고 기억으로 재구성하지 않습니다.

`000-plan.md`가 self-drive inactive이면 sidecar 파일이 있어도 historical context입니다.

## Finite Mode 계약

finite mode는 준비된 flow sequence를 진행합니다.

advance 조건:

- current flow verification이 `pass`입니다.
- `flow` handoff condition이 blocked가 아닙니다.
- `flow` output에서 next flow identity가 기록돼 있습니다.
- approval boundary가 sidecar와 일치합니다.
- plan과 sidecar gate가 다시 통과합니다.

진행 방식:

- reporting 전에는 current index를 유지합니다.
- 위 조건이 모두 충족되어도 advance 직전 plan과 sidecar gate를 다시 확인합니다.
- advance가 확정된 뒤 next index, current flow label, progress note, `next_action`, ledger를 갱신합니다.
- repair/recheck는 current flow를 pass로 만들기 위한 work이며 advance count로 세지 않습니다.
- sequence completion은 terminal closure가 아닙니다. completion을 보고하고 next-flow routing을 다시 엽니다.

### Finite Endpoint Update

endpoint가 sequence exhaustion 뒤 self-drive stop이면 self-drive 자동 진행만 멈춥니다.

사용자가 sequence 완료 뒤 “another bounded batch”, “another inventory cycle”, “다음 묶음도 계속” 같은 요청을 하면 old endpoint에서 이어가지 않습니다.
새 bounded finite cycle로 sidecar를 refresh합니다.

새 cycle에는 다음을 다시 잠급니다.

- source-backed batch objective
- first flow identity
- planned flow count
- acceptance signal
- verification expectation
- approval checkpoints
- endpoint 또는 stop condition
- `next_action`
- ledger continuation note

새 cycle 준비가 끝나기 전에는 다음 flow로 advance하지 않습니다.

### Finite Relock And Recovery

order, endpoint, scope, non-goal, target, acceptance, approval boundary가 바뀌면 자동 진행을 멈춥니다.
해당 변경 여부와 돌아갈 earliest phase는 `flow` readiness/ambiguity 또는 handoff 판단에 의존합니다.
영향받는 earliest phase로 돌아가 relock하고, sidecar와 affected flow record를 갱신한 뒤 gate가 다시 통과할 때만 계속합니다.

external blocker 또는 access blocker가 회복돼도 이전 blocker 상태를 바로 continuation 근거로 쓰지 않습니다.
plan, sidecar, active flow record를 다시 읽고, endpoint, approval boundary, next identity, verification expectation을 relock한 뒤 계속 여부를 판단합니다.

ordinary note가 recorded boundary 안에 있고 scope/endpoint/approval을 바꾸지 않으면 자동 진행을 계속할 수 있습니다.
이 경우 active flow record `Execution Log` 또는 sidecar ledger에 짧게 기록하고, turn-level routing이 바뀔 때만 `000-plan.md`를 갱신합니다.

## Infinite Mode 계약

infinite mode는 "무한 todo list"가 아니라 counted bounded iteration입니다.
현재 iteration만 구체화하고, iteration이 끝날 때 다음 bounded target을 다시 고릅니다.

최소 state:

- `mode: infinite`
- `loop_count`
- `current_loop_label`
- `active_flow_record`
- `next_action`
- `Goal`
- `Ledger`

각 iteration:

1. 기록된 scope 안에서 하나의 bounded target을 고릅니다.
2. 가장 작은 완전한 변경 또는 확인을 수행합니다.
3. `flow` verification expectation에 맞춰 검증합니다.
4. 결과와 ledger를 기록합니다.
5. approval checkpoint, blocker state, endpoint를 확인합니다.
6. 계속 가능하면 `loop_count`를 증가시키고 `next_action`을 다음 bounded iteration으로 갱신합니다.

`loop_count`는 continuation gate가 통과한 뒤에만 증가합니다.
gate 충돌 전에는 이전 count를 유지합니다.
endpoint가 이번 batch의 마지막 iteration이면 다음 loop용 count를 올리지 않고 endpoint completion을 보고합니다.

자동 진행 중단 조건:

- missing target
- `fail`, `blocked`, `insufficient`, `not-started`, `requested`
- same bounded target과 same cause의 2회 연속 non-pass
- approval need
- access 또는 external blocker
- user input need
- no useful bounded work left
- sidecar/record conflict
- endpoint, scope, target, order, acceptance, approval boundary 변경

`insufficient`는 같은 bounded target 안에서 evidence를 보강하고 재검증할 수 있습니다.
보강 결과가 pass가 되기 전에는 다음 loop로 advance하지 않습니다.

external blocker가 회복되면 이전 blocker 상태를 바로 continuation 근거로 쓰지 않습니다.
sidecar gate, endpoint, approval boundary, active flow record를 다시 확인하고 relock한 뒤 계속 여부를 판단합니다.

### Infinite Endpoint Update

사용자가 "또 8개", "another 8", "another bounded batch", "another inventory cycle"처럼 추가 bounded batch나 cycle을 요청하면 현재 infinite state 안의 endpoint/order update로 처리합니다.
무한 todo list를 만들지 않습니다.

기존 `loop_count`와 ledger history는 보존합니다.
새 batch objective, stop condition, next bounded target, acceptance signal, verification expectation, approval checkpoint를 sidecar에 다시 잠급니다.
현재 loop를 포함하는지 추가 loop인지 모호하면 user-gated clarification으로 라우팅합니다.

infinite 중 endpoint가 finite stop condition으로 바뀌면 기존 infinite autonomous advance를 멈춥니다.
endpoint 변경을 sidecar와 ledger에 기록하고, finite sequence가 필요하면 새 finite preparation으로 전환합니다.

## Interruption 계약

active self-drive 중 사용자 메시지는 explicit stop이 아닌 한 self-drive boundary 안에서 먼저 해석합니다.

우선순위:

1. source-recorded explicit stop
2. approval-sensitive action 또는 approval boundary 확장
3. scope, non-goal, endpoint, target, order, acceptance 변경
4. blocker 또는 repeated failure
5. status/progress question
6. recorded boundary 안의 ordinary note

status 질문에는 current identity, finite index 또는 loop count, verification state, next required action을 보고합니다.
기록된 self-drive가 여전히 허용할 때만 계속합니다.

explicit stop은 active flow record에 source text 또는 compact source reference를 기록한 뒤에만 terminal closure authority가 됩니다.

## 승인과 검증 계약

각 flow 또는 iteration은 endpoint 처리 전에 verification을 가져야 합니다.
non-pass 또는 미완료 verification은 endpoint exhaustion, loop advance, next-flow continuation보다 먼저 repair, evidence collection, blocker routing으로 처리합니다.
작업 증거가 pass처럼 보이더라도 active flow record 또는 sidecar metadata의 `verification_status`가 `requested`, `not-started`, `fail`, `blocked`, `insufficient`이면 pass 상태로 갱신되기 전까지 advance할 수 없습니다.
work evidence와 verification metadata가 충돌하면 verification mismatch로 라우팅하고, mismatch가 해소되기 전에는 continuation 또는 advance할 수 없습니다.

approval-sensitive action은 exact action, target, expected effect, risk, recovery path, included/excluded scope, endpoint가 기록된 경우에만 실행 checkpoint로 들어갈 수 있습니다.
self-drive, subagent output, readiness, verification은 승인을 대체하지 않습니다.

사용자가 모든 write 또는 record 생성을 금지하면 sidecar를 갱신하지 않고 in-memory continuity로만 보고합니다.
그 상태는 다음 autonomous advance 근거가 될 수 없습니다.

## 검토 기준

- self-drive가 명시 요청 또는 선택 없이 활성화되지 않는가?
- 매 advance/loop 전에 plan과 sidecar gate를 확인하는가?
- pass verification, non-blocked handoff, known next identity, matching approval boundary가 모두 있을 때만 advance하는가?
- handoff와 next identity 판단이 turn-gate 자체 추정이 아니라 `flow` output에 근거하는가?
- non-pass, blocker, approval need, stale state, endpoint/order/scope change가 advance보다 먼저 처리되는가?
- finite sequence completion과 self-drive stop을 terminal closure로 오해하지 않는가?
- finite blocker recovery와 another bounded batch가 sidecar refresh 또는 relock을 거치는가?
- infinite가 counted bounded iteration으로만 진행되는가?
- infinite another bounded batch 또는 inventory cycle이 endpoint/order update로 직접 처리되는가?
- ordinary note, blocker recovery, insufficient repair, another batch, mode/endpoint transition이 기록과 relock을 거치는가?
