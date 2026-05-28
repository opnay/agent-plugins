# turn-gate self-drive 계약

## 소유 범위

self-drive는 사용자가 명시적으로 맡긴 범위 안에서 다음 행동을 자율 선택하게 하는 overlay입니다. 기본 `intake -> framing -> preparation -> work -> verification -> reporting -> next-flow` loop를 대체하지 않고, 그 loop 위에서 언제 사용자에게 되묻지 않고 계속할 수 있는지만 좁힙니다.

self-drive는 두 모드를 가질 수 있습니다.

- `finite`: 준비된 flow sequence를 순서대로 진행합니다.
- `infinite`: 사용자가 멈추라고 할 때까지 반복하되, 매번 하나의 bounded iteration만 준비하고 실행합니다.

두 모드 모두 위험작업 실행 권한을 만들지 않습니다. commit, push, PR, publish, release, version bump, destructive/external action, scope expansion은 self-drive 중에도 명시 승인 checkpoint로 돌아갑니다.

## 활성화 계약

self-drive는 명시 요청 또는 next-flow mode 선택으로만 활성화합니다. 긴 task list, verification success, subagent availability, 이전 대화의 의욕적 표현으로 추론하지 않습니다.

공통 sidecar state는 다음을 가져야 합니다.

- `status`
- objective 또는 source-backed goal
- mode: `finite` 또는 `infinite`
- current flow/iteration identity
- active flow record identity
- next action
- progress note
- endpoint 또는 stop condition
- acceptance signal
- verification expectation
- allowed autonomous actions
- approval checkpoints
- blocker return conditions

`finite` mode는 prepared flow sequence, active flow index, current flow label, planned flow count를 추가로 기록합니다.

`infinite` mode는 큰 todo list나 speculative sequence를 만들지 않습니다. `loop_count`, current loop label, next bounded iteration만 기록합니다. "강제로 종료할 때까지", "계속", "무한히" 같은 요청은 무제한 실행 권한이 아니라 counted bounded iteration으로 해석합니다.

infinite mode를 새로 준비할 때 첫 반복은 `loop_count: 1`로 시작합니다. target이 없으면 자동 진행하지 않고 user-gated target selection 또는 blocker routing으로 돌아갑니다. 기록된 범위 안에서 하나의 bounded target을 고를 수 있을 때만 첫 iteration을 시작합니다.

## Sidecar 계약

각 self-drive flow 또는 iteration 시작 때 `000-plan.md`와 `000-self-drive.md`를 읽습니다.

공통 확인 항목:

- plan의 self-drive status와 sidecar pointer
- mode
- current identity와 active flow record identity
- next action
- endpoint 또는 stop condition
- acceptance signal
- verification expectation
- approval checkpoint
- blocker state

`000-plan.md`는 self-drive active 여부와 sidecar pointer만 저장합니다. sequence, loop, ledger, approval checkpoint, endpoint detail은 `000-self-drive.md`가 소유합니다. flow-local snapshot은 active flow record가 소유합니다.

state가 누락, 충돌, 범위 이탈, stale이면 자동 진행을 멈추고 user-gated recovery로 돌아갑니다. 숫자 index나 loop count만 맞는 상태는 충분하지 않습니다. finite mode에서 `active_flow_index`가 `planned_flow_count`를 초과하면 stale/corrupt sidecar로 보고 reconcile을 요청합니다. endpoint 판단 중 record 또는 sidecar를 읽을 수 없으면 access blocker로 보고하고, 기억으로 endpoint를 재구성하지 않습니다.

## Finite Mode

finite mode는 준비된 flow sequence를 소유합니다.

진행 조건:

- 현재 flow verification이 pass입니다.
- handoff가 blocked가 아닙니다.
- 다음 flow identity가 기록돼 있습니다.
- approval boundary가 기록과 일치합니다.

진행 방식:

- reporting 전에는 현재 index를 유지합니다.
- advance가 확정된 뒤 다음 index와 current flow label을 갱신합니다.
- planned sequence가 소진돼도 terminal closure가 아닙니다. completion을 보고하고 next-flow routing으로 돌아갑니다.

endpoint가 "sequence exhausted -> stop self-drive"이면 self-drive 자동 진행만 멈춥니다. 이것은 turn terminal closure가 아닙니다. endpoint가 "sequence exhausted -> create next inventory cycle"이면 sidecar를 새 bounded finite cycle로 갱신한 뒤 첫 flow identity, planned count, acceptance, verification을 다시 잠급니다.

## Infinite Mode

infinite mode는 "계속 작업" 요청을 하나의 무한 todo로 보지 않습니다. 현재 반복만 구체화하고, 반복이 끝날 때 다음 반복을 다시 고릅니다.

frontmatter는 최소한 다음을 가져야 합니다.

- `mode: infinite`
- `loop_count`
- `current_loop_label`
- `active_flow_record`
- `next_action`

본문은 `Goal`과 `Ledger` 중심으로 유지합니다. 별도 `Todo` 또는 `Sequence` 섹션은 필요할 때만 두며, 무한성을 표현하기 위해 큰 목록을 만들지 않습니다.

각 iteration은 다음 순서를 따릅니다.

1. 현재 범위 안에서 하나의 bounded target을 고릅니다.
2. 가장 작은 완전한 변경 또는 확인을 수행합니다.
3. 검증합니다.
4. 결과와 ledger를 기록합니다.
5. approval checkpoint, blocker, endpoint를 확인합니다.
6. 계속 가능하면 `loop_count`를 증가시키고 `next_action`을 다음 bounded iteration으로 갱신합니다.

무의미한 반복, 검증 불충분, 반복 실패, 대상 부재, 승인 필요, access/external blocker, 사용자 입력 필요가 생기면 자동 진행을 멈춥니다.

## Interruption 계약

active self-drive 중 사용자 메시지는 explicit stop이 아닌 한 현재 self-drive 안에서 먼저 해석합니다.

우선순위:

1. source-recorded explicit stop
2. approval-sensitive action 또는 approval boundary 확장
3. scope, non-goal, endpoint, target, order, acceptance signal 변경
4. blocker 또는 repeated failure
5. status/progress question
6. recorded boundary 안의 ordinary note

scope, endpoint, target, order, acceptance signal, approval boundary가 바뀌면 자동 진행을 멈추고 가장 이른 영향 phase로 돌아갑니다. endpoint 변경은 relock/update event입니다. 바뀐 endpoint를 sidecar에 기록하고, affected flow record의 next action과 ledger를 갱신한 뒤에만 계속 여부를 판단합니다. status 질문은 current identity, sequence/loop position, verification state, next required action을 보고한 뒤 기록된 self-drive가 여전히 허용할 때만 계속합니다.

## 종료와 승인 계약

각 flow 또는 iteration은 endpoint 처리 전에 verification을 가져야 합니다. `fail`, `blocked`, `insufficient`, `not-started`, `requested` 같은 non-pass 또는 미완료 verification은 endpoint exhaustion, loop advance, next-flow continuation보다 먼저 repair, evidence collection, blocker routing으로 처리합니다.

approval-sensitive action은 initial preparation 또는 sidecar가 exact action, target, expected effect, risk, recovery path, included/excluded scope, endpoint를 기록한 경우에만 실행 checkpoint로 들어갈 수 있습니다. self-drive, subagent output, readiness, verification은 승인을 대체하지 않습니다.

사용자가 모든 쓰기나 record 생성을 금지하면 sidecar를 갱신하지 않고 in-memory continuity로만 보고합니다. 그 상태는 다음 autonomous advance 근거가 될 수 없습니다.

explicit stop은 active flow record에 source text 또는 compact source reference를 기록한 뒤에만 terminal closure authority가 됩니다.
