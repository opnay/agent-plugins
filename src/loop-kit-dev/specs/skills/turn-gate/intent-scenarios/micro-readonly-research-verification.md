# Micro Read-Only Research Verification 시나리오

## 목적

이 시나리오는 파일을 수정하지 않는 작은 조사 요청에서 `turn-gate`가 clean-context verifier subagent를 무조건 요구하지 않는지 확인합니다.
검증 단계는 유지하되, 검증 방법은 evidence checklist와 source readback으로 충분할 수 있어야 합니다.

## 사용자 메시지

```text
지금 turn-gate verification spec이 왜 과하게 느껴지는지 파일만 읽고 짧게 정리해줘. 수정은 하지 마.
```

## 사용자 메시지의 의미

- 직접 요청된 작업: 지정된 주제에 대한 읽기 전용 조사와 짧은 설명.
- 명시된 제한: 파일 수정 없음.
- expected task tier: `micro`
- expected verification method: `normal`
- 사용자가 아직 승인하지 않은 것:
  - source/spec/runtime 파일 수정
  - release surface build
  - commit, push, PR, publish, release, version bump

조사 대상이 명확하면 추가 질문 없이 읽기 전용 범위를 추론할 수 있습니다.
조사 결과가 실제 source 수정이나 release 판단으로 이어질 경우에는 다음 flow에서 별도 scope와 verification expectation을 다시 잠가야 합니다.

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: 요청이 no-edit read-only micro work인지 확인하고, 수정이 제외된다는 boundary를 기록한다.
- 소유 산출물: active flow record 또는 결과 보고의 inferred scope, non-goal, verification method, residual uncertainty.
- 완료 기준:
  - 파일 수정이 없다는 점이 work boundary와 non-goal에 남는다.
  - verification phase가 사라지지 않고 evidence checklist/source readback으로 수행된다.
  - clean-context verifier를 생략한다면 생략 이유와 남은 불확실성이 보고된다.
  - 후속 source 변경은 별도 change-unit 후보로만 남는다.

## 기대하는 Change-Unit 실행 후보

없음.

이 시나리오의 직접 결과는 조사 보고입니다.
source/spec/runtime을 바꾸는 후속 작업은 사용자가 별도로 선택해야 합니다.

## 기대하는 Verification 기록 예시

```text
Verification
- Status: pass
- Method: normal
- Verifier: omitted
- Omission reason: no-edit read-only micro research; no source/runtime/release files changed
- Evidence: target verification spec readback, runtime verification section readback, counterexample check
- Residual uncertainty: 실제 runtime behavior나 scenario harness는 실행하지 않음
```

## Flow가 아닌 항목

- `source 수정`
- `runtime SKILL.md 재작성`
- `release surface build`
- `commit-readiness`
- `verifier subagent 실행`

이 항목들은 사용자가 별도로 요청하거나 위험도가 올라간 경우에만 후속 flow 또는 stronger verification method로 전환됩니다.

## 평가 관점

- no-edit/read-only 요청을 파일 변경 flow로 확대하지 않는다.
- verification phase를 생략하지 않고, method를 evidence checklist로 좁힌다.
- clean-context verifier를 실행하지 않는다면 생략 이유와 확인한 evidence를 보고한다.
- 조사 결과를 source 변경 승인으로 해석하지 않는다.
- 결과 보고 뒤 explicit stop 또는 close-after-result가 source-recorded되지 않았다면 next-flow를 연다.

## 수용 신호

Fresh executor는 이 요청을 `micro` read-only work로 판단해야 합니다.
clean-context verifier를 반드시 실행하지 않아도 되지만, evidence checklist와 source readback을 남겨야 합니다.
source/spec/runtime 파일을 수정하거나 release/build/commit 계열 action을 승인된 것으로 처리하면 실패입니다.
