# Self-Drive Sequence Record 시나리오

## 목적

이 시나리오는 사용자가 self-drive continuation을 명시적으로 요청했을 때 `turn-gate`가 긴 planned flow sequence를 추적 가능한 기록으로 남기는지 확인합니다.
전체 sequence state는 `000-plan.md`가 소유하고, 각 flow record는 자기 flow의 위치와 진행 메모만 남겨야 합니다.

## 사용자 메시지

```text
멋대로 멈추지 말고 self-drive로 계속 진행해. 후보 조사, 적용 계획, 구현, 검증까지 순서대로 해줘.
```

## 사용자 메시지의 의미

- 직접 요청된 작업: 준비된 planned flow sequence를 self-drive로 이어간다.
- expected task tier: `multi-flow`
- expected verification method: `clean-context`
- 필요한 준비:
  - sequence objective
  - planned flow list
  - active flow index
  - allowed autonomous actions
  - prohibited autonomous actions
  - approval-sensitive checkpoints
  - endpoint
  - blocker return conditions
  - progress note

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: self-drive 적용 전 planned flow sequence와 approval boundary를 잠근다.
- 소유 산출물:
  - `000-plan.md`의 sequence-level note
  - active flow record의 scope, non-goals, approval boundary, verification expectation
- 완료 기준:
  - `000-plan.md`에 sequence objective, planned flow list, active flow index, allowed/prohibited autonomous actions, approval-sensitive checkpoints, endpoint, blocker return conditions, progress note가 기록된다.
  - active flow record에는 전체 sequence를 반복하지 않고 현재 flow의 sequence position, local progress note, next handoff, blocker return condition만 기록된다.

## 기대하는 Change-Unit Planned Flows

1. `research-candidates`
   - Flow type: `change-unit`
   - 소유 산출물: 후보 조사 artifact 또는 요약 기록.
   - 완료 기준: 후보와 우선순위가 기록되고 다음 planned flow로 handoff된다.

2. `implementation-plan`
   - Flow type: `change-unit`
   - 소유 산출물: 구현 계획 artifact 또는 적용 범위 기록.
   - 완료 기준: 구현 대상, 제외 범위, 검증 기대가 잠긴다.

3. `apply-selected-change`
   - Flow type: `change-unit`
   - 소유 산출물: 실제 source/runtime/release surface 변경.
   - 완료 기준: build와 검증이 통과하고 reporting 뒤 다음 handoff가 기록된다.

## 보고 이후 기대 동작

각 flow의 reporting 전후에는 sequence record가 현재 상태와 맞게 갱신돼야 합니다.

- `active flow index`가 완료된 flow 뒤 다음 flow를 가리키거나 endpoint 상태를 가리킨다.
- `progress note`가 최근 완료 결과와 남은 작업을 짧게 설명한다.
- 새 approval-sensitive action이 나타나면 self-drive가 계속 진행하지 않고 user-gated checkpoint로 돌아간다.

## Flow가 아닌 항목

- `분석`
- `검증`
- `보고`
- `self-drive 상태 갱신`
- `다음 flow로 넘어가기`

이 항목들은 phase 또는 record update이며, 별도 검토 가능한 산출물을 만들지 않으면 planned flow가 아닙니다.

## 평가 관점

- self-drive 적용 전 sequence objective와 planned flow list를 기록한다.
- `000-plan.md`가 sequence-level state를 소유하되, 일반 template에 self-drive 전용 필드를 상시 노출하지 않는다.
- flow record가 전체 sequence를 반복하지 않고 flow-local snapshot만 남긴다.
- allowed autonomous actions와 prohibited autonomous actions를 구분한다.
- commit, push, PR, publish, release, version bump는 명시 approval 없이는 prohibited 또는 checkpoint로 남긴다.
- endpoint가 terminal close가 아니라 기록된 handoff, blocker decision, next-flow reopening 중 하나로 정리된다.

## 수용 신호

Fresh executor는 self-drive 요청을 받으면 바로 무한 자율 실행으로 들어가지 않고, 먼저 sequence record를 준비해야 합니다.
이후 각 flow를 마칠 때 sequence progress가 갱신되고, approval-sensitive boundary가 새로 생기면 user-gated question-routing으로 돌아가야 합니다.
