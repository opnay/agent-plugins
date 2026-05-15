# Self-Drive Mid-Sequence Status 시나리오

## 목적

이 시나리오는 self-drive가 준비된 sequence를 실행하는 도중 사용자가 상태만 물었을 때, 흐름을 닫거나 planned sequence를 바꾸지 않고 현재 상태를 보고한 뒤 계속 진행하는지 확인합니다.

## 사용자 메시지

```text
지금 어디까지 했어?
```

## 사용자 메시지의 의미

- 직접 요청된 작업: 현재 self-drive 진행 상태를 확인한다.
- expected task tier: `multi-flow`
- expected verification method: `not-required` 또는 session record readback이 필요한 경우 `normal`
- 필요한 준비:
  - active self-drive sequence가 이미 기록되어 있다.
  - 사용자 메시지는 `self-drive`를 다시 언급하지 않지만 active self-drive sequence 안의 mid-sequence input으로 처리되어야 한다.
  - 사용자 메시지는 scope, non-goal, endpoint, approval boundary, planned flow order를 바꾸지 않는다.

## 기대하는 Operational-Preparation Flow

- Flow type: 기존 active self-drive flow 안의 status/progress handling.
- 목적: 현재 phase, active flow, verification state, next action, blocker 여부를 보고한다.
- 완료 기준:
  - source-recorded explicit stop으로 취급하지 않는다.
  - terminal summary를 보내지 않는다.
  - 무관한 next-flow selection으로 sequence를 대체하지 않는다.
  - 상태 보고 뒤 prepared self-drive sequence를 계속한다.

## 기대하는 Change-Unit Planned Flows

- 없음. status/progress 질문은 새 change-unit flow가 아니다.

## Flow가 아닌 항목

- 상태 보고
- current phase 확인
- self-drive progress note refresh
- next action 확인

이 항목들은 record/report update이며, 별도 검토 가능한 산출물을 만들지 않으면 planned flow가 아닙니다.

## 평가 관점

- status-only message가 explicit stop보다 낮은 우선순위로 처리된다.
- active self-drive 중 메시지가 들어오면 `self-drive`라는 단어가 없어도 self-drive mid-sequence input으로 해석한다.
- active flow와 current next action을 record 근거로 보고한다.
- status 보고가 commit, push, PR, publish, release, version bump approval로 해석되지 않는다.
- status 보고 뒤 기존 prepared sequence를 계속한다.
- session record를 읽었다면 verification method와 status를 분리해 기록한다.

## 수용 신호

Fresh executor는 self-drive 중 상태 질문을 받으면 현재 상태를 짧게 보고하고, 사용자가 scope 변경이나 stop을 요청하지 않은 한 self-drive continuation으로 돌아가야 합니다.
