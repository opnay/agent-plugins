# turn-gate self-drive 계약

## 소유 범위

명시적으로 준비된 finite flow sequence 위에 적용되는 self-drive overlay. self-drive는 기본 `preparation -> work -> verification -> reporting -> next-flow` loop를 대체하지 않습니다.

## 활성화 계약

self-drive는 records에 다음 항목이 있을 때만 적용합니다.

- sequence objective
- prepared flow sequence
- active flow index와 current flow label
- progress note
- allowed and prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- acceptance signal
- verification expectation

self-drive는 이전 대화의 의욕적 표현, 긴 task list, verification success, subagent availability에서 추론할 수 없습니다. 명시적으로 요청됐거나 next-flow mode로 명시 선택돼야 합니다.

## 사이드카 계약

각 self-drive flow를 시작할 때 `000-plan.md`와 `000-self-drive.md`를 읽습니다. status, active flow index, current flow label, progress note, planned flow count, endpoint, required next action, acceptance signal, blocker state를 확인합니다.

identity 또는 index state가 누락됐거나 충돌하거나 범위를 벗어나면 user-gated routing으로 돌아갑니다. index를 wrap하거나 기억만으로 advance하지 않습니다.

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

사용자가 endpoint, order, target, approval boundary, acceptance signal을 바꾸면 autonomous advancement를 멈추고 preparation으로 라우팅합니다. 사용자가 status를 물으면 current flow, sequence position, verification state, next required action을 보고한 뒤 계속합니다.

## 종료와 승인 계약

각 flow는 endpoint exhaustion 처리 전에 verification을 수행해야 합니다. non-pass verification은 먼저 verification contract를 통해 라우팅합니다.

각 flow의 reporting 전후에는 self-drive sidecar를 현재 active flow index, current flow label, progress note, next handoff, blocker state에 맞게 갱신합니다. progress ledger는 sequence transition과 material update의 history이므로 current summary만 남기기 위해 덮어쓰지 않습니다.

open-ended self-drive도 finite current cycle과 explicit repeat policy가 필요합니다.

self-drive는 initial preparation이 exact action, target, expected effect, risk, recovery path, included/excluded scope, endpoint를 기록한 경우에만 approval-sensitive action을 실행할 수 있습니다. subagent는 approval을 대체하지 않습니다.

sequence completion은 terminal closure가 아닙니다. endpoint에 도달하면 completion을 보고하고, records를 갱신한 뒤, 사용자가 명시적으로 멈추지 않는 한 next-flow로 라우팅합니다.
