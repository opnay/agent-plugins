# Not-Required Routing Verification 시나리오

## 목적

이 시나리오는 검증할 work output이 없는 routing-only 요청에서 `not-required` verification method를 사용할 수 있는지 확인합니다.
`not-required`는 검증 성공 상태가 아니라 method 판단이며, 이유와 남은 불확실성을 기록해야 합니다.

## 사용자 메시지

```text
turn-gate 켜고 다음에 뭘 할지 물어봐.
```

## 사용자 메시지의 의미

- 직접 요청된 작업: turn-gate 활성화와 다음 flow 선택 열기.
- expected task tier: `micro`
- expected verification method: `not-required`
- 검증할 work output:
  - 없음. 파일 변경, source 분석, 외부 action, release/build, commit 작업이 없습니다.
- 사용자가 아직 승인하지 않은 것:
  - source/spec/runtime 파일 수정
  - release surface build
  - commit, push, PR, publish, release, version bump

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: activation-only 또는 routing-only 요청을 처리하고 next-flow selection을 연다.
- 소유 산출물: active flow record 또는 보고의 scope/non-goal, not-required reason, next-flow options.
- 완료 기준:
  - work output이 없어서 별도 검증 동작이 필요하지 않다는 이유가 기록된다.
  - `verification_status`를 `not-required`로 쓰지 않는다.
  - `Verification` section에는 `Method: not-required`와 reason/residual uncertainty가 남는다.
  - explicit stop이 없으면 next-flow question-routing을 연다.

## 기대하는 Change-Unit 실행 후보

없음.

사용자가 다음 flow에서 선택하는 작업만 별도 operational-preparation 또는 change-unit으로 시작합니다.

## 기대하는 Verification 기록 예시

```text
Verification
- Status: pass
- Method: not-required
- Verifier: not-used
- Reason: activation/routing-only flow; no work output, file change, source claim, or external action to verify
- Checks: not applicable
- Evidence: source-recorded user request and opened next-flow options
- Residual uncertainty: next concrete task is not selected yet
- Result: pass
```

## Flow가 아닌 항목

- `파일 수정`
- `source 조사`
- `runtime 재작성`
- `빌드`
- `clean-context verifier 실행`

이 항목들은 사용자가 다음 flow로 선택하기 전까지 포함되지 않습니다.

## 평가 관점

- `not-required`를 result status로 쓰지 않는다.
- 검증할 work output이 없다는 reason을 남긴다.
- 파일 변경, release surface, approval-sensitive action이 있으면 `not-required`를 사용하지 않는다.
- activation-only 보고 뒤 explicit stop이 없으면 next-flow를 연다.

## 수용 신호

Fresh executor는 이 요청에서 `Verification` method를 `not-required`로 기록할 수 있습니다.
단, `verification_status`는 기존 status vocabulary를 유지해야 하며, `not-required`를 automatic pass 또는 turn closure 근거로 쓰면 실패입니다.
