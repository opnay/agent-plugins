# turn-gate self-drive 계약

## 소유 범위

명시적으로 준비된 finite flow sequence 위에 적용되는 self-drive overlay. self-drive는 기본 `intake -> framing -> preparation -> work -> verification -> reporting -> next-flow` loop를 대체하지 않습니다.

## 활성화 계약

self-drive는 records에 다음 항목이 있을 때만 적용합니다.

- sequence objective
- prepared flow sequence
- active flow index와 current flow label
- progress note
- explicit repeat policy, if the sequence is open-ended
- allowed and prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- acceptance signal
- verification expectation

self-drive는 이전 대화의 의욕적 표현, 긴 task list, verification success, subagent availability에서 추론할 수 없습니다. 명시적으로 요청됐거나 next-flow mode로 명시 선택돼야 합니다.

## 사이드카 계약

각 self-drive flow를 시작할 때 `000-plan.md`와 `000-self-drive.md`를 읽습니다. status, sidecar pointer, active flow index, current flow label, active flow record identity, progress note, planned flow count, endpoint, required next action, acceptance signal, blocker state를 확인합니다.

identity 또는 index state가 누락됐거나 충돌하거나 범위를 벗어나면 user-gated routing으로 돌아갑니다. index를 wrap하거나 기억만으로 advance하지 않습니다.

`000-plan.md`가 self-drive inactive이거나 sidecar pointer를 갖지 않으면 `000-self-drive.md`는 historical context일 뿐이며 continuation authority가 아닙니다. 사용자가 self-drive continuation을 요청한 상태라면 user-gated recovery로 라우팅하고, 그렇지 않으면 일반 active-flow 또는 next-flow routing으로 처리합니다.

이 계약에서 user-gated recovery, user-gated routing, blocker routing은 같은 중단 계열입니다. 자동 진행을 멈추고 사용자에게 필요한 reconcile, approval, access, scope 결정을 받는 상태를 의미합니다.

active flow index, current flow label, `000-plan.md` active pointer, active flow record identity는 같은 flow를 가리켜야 합니다. 숫자 index만 맞거나 label만 맞는 상태는 충분하지 않습니다.

`000-plan.md`는 self-drive status와 sidecar pointer만 저장합니다. sequence detail은 `000-self-drive.md`가 소유하고, flow-local snapshot은 active flow record가 소유합니다.

## 중단 처리 계약

active self-drive 중 사용자 메시지가 들어오면 explicit turn stop이 아닌 한 active sequence 안에서 먼저 해석합니다.

우선순위:

1. source-recorded explicit stop
2. destructive, external, commit, push, PR, publish, release, version bump, approval-boundary-expanding request
3. scope, non-goal, endpoint, target, prepared order, acceptance-signal change
4. blocker 또는 repeated failure
5. status/progress question
6. recorded boundary 안의 ordinary note

self-drive는 질문 조건을 좁히지만 질문을 비활성화하지 않습니다.

사용자가 scope, non-goal, endpoint, order, target, approval boundary, acceptance signal을 바꾸면 autonomous advancement를 멈추고 가장 이른 영향 phase로 돌아갑니다. intent 또는 acceptance signal 변경은 `intake`, scope/non-goal/target/order/endpoint/flow boundary 변경은 `framing`, selected-flow readiness 또는 approval boundary 변경은 `preparation`으로 라우팅합니다.

사용자가 status를 물으면 current flow, sequence position, verification state, next required action을 보고합니다. verification state는 active flow record를 우선하고, sidecar는 sequence position과 handoff 상태를 보조합니다. 기록된 sequence가 여전히 허용할 때만 계속합니다.

recorded boundary 안의 ordinary note는 scope, non-goal, endpoint, target, order, approval boundary, acceptance signal, blocker state를 바꾸지 않는 한 flow boundary를 변경하지 않습니다.

## 종료와 승인 계약

각 flow는 endpoint exhaustion 처리 전에 verification을 수행해야 합니다. non-pass verification은 먼저 verification contract를 통해 라우팅합니다.

각 flow의 reporting 전후에는 self-drive sidecar를 현재 active flow index, current flow label, progress note, next handoff, blocker state에 맞게 갱신합니다. reporting 전에는 현재 index를 유지하고, handoff/advance가 확정된 뒤 다음 index로 갱신합니다. advance confirmation은 current flow verification pass, non-blocked handoff condition, next flow identity, approval boundary가 모두 기록과 일치할 때만 성립합니다. progress ledger는 sequence transition과 material update의 history이므로 current summary만 남기기 위해 덮어쓰지 않습니다. reporting에는 ledger가 append-only로 유지됐는지와 새 material update가 무엇인지 포함합니다.

사용자가 모든 쓰기나 record 생성을 명시적으로 금지하면 sidecar를 갱신하지 않고 in-memory continuity로만 보고합니다. 이 경우 self-drive continuation은 기록되지 않은 상태이므로 다음 autonomous advance의 근거가 될 수 없습니다.

open-ended self-drive도 finite current cycle과 explicit repeat policy가 필요합니다. repeat policy는 cycle boundary, repeat limit 또는 repeat condition, cycle마다 필요한 verification, user-gated stop condition을 포함합니다.

blocker state는 `none`이 아니면 영향 범주를 기록합니다. blocker가 acceptance, verification, approval boundary, access, external state, required user input에 영향을 주면 autonomous advancement를 멈추고 user-gated blocker routing으로 돌아갑니다. flow-local repair로 해결 가능한 내부 작업 실패만 verification recovery로 처리할 수 있습니다.

self-drive는 initial preparation이 exact action, target, expected effect, risk, recovery path, included/excluded scope, endpoint를 기록한 경우에만 approval-sensitive action을 실행할 수 있습니다. subagent는 approval을 대체하지 않습니다.

sequence completion은 terminal closure가 아닙니다. endpoint에 도달하면 completion을 보고하고, records를 갱신한 뒤, 사용자가 명시적으로 멈추지 않는 한 next-flow로 라우팅합니다. explicit stop은 active flow record에 source text 또는 compact source reference를 기록한 뒤에만 closure authority가 됩니다.
