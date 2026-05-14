# File Change Verifier Default 시나리오

## 목적

이 시나리오는 파일 변경이 있는 작업에서 risk-based verification이 clean-context verifier 기본값을 약화하지 않는지 확인합니다.
작은 문서 fixture 수정이라도 실제 파일이 바뀌면 `micro`가 아니며, verifier 또는 동등하게 독립된 검증 경로가 필요합니다.

## 사용자 메시지

```text
turn-gate intent-scenarios에 no-edit 조사 요청 예시 하나를 추가해줘.
```

## 사용자 메시지의 의미

- 직접 요청된 작업: `intent-scenarios/` 아래 spec-side fixture 파일 추가 또는 갱신.
- expected task tier: `single-flow`
- expected verification method: `clean-context`
- 포함될 수 있는 산출물:
  - 새 scenario Markdown 파일
  - scenario README index 갱신
  - change spec의 fixture 추가 기록
- 사용자가 아직 승인하지 않은 것:
  - runtime `SKILL.md` 재작성
  - plugin README/manifest/defaultPrompt 변경
  - version bump
  - commit, push, PR, publish, release

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: 요청이 spec-side fixture 변경인지 확인하고, 단일 change-unit 후보로 충분한지 판단한다.
- 소유 산출물: active flow record의 scope/non-goal, target files, verification expectation, approval boundary.
- 완료 기준:
  - 파일 변경이 있으므로 `micro`가 아니라는 판단이 기록된다.
  - 후속 실행 후보가 하나의 scenario-fixture change-unit으로 정리된다.
  - runtime/release/version/commit 계열 작업은 제외된다.

## 기대하는 Change-Unit 실행 후보

1. `add-turn-gate-intent-scenario-fixture`
   - Flow type: `change-unit`
   - 소유 산출물: scenario Markdown 파일, scenario README entry, 관련 change spec 기록.
   - 근거: fixture 추가는 하나의 검토 가능한 spec-side 변경 단위입니다.
   - 완료 기준: 새 scenario가 기존 형식과 맞고 README가 색인하며, 관련 change spec이 변경 의도를 기록합니다.
   - 검증 기대: Markdown 구조 확인, README coverage 확인, stale wording 검색, clean-context verifier readback.

## 기대하는 Verification 기록 예시

```text
Verification
- Status: pass
- Method: clean-context
- Verifier: <subagent id>
- Evidence: README coverage check, H1 check, verifier readback of changed scenario and related spec contract
- Residual uncertainty: scenario harness empirical execution은 별도 후속 검증
```

## Flow가 아닌 항목

- `조사`
- `시나리오 설계`
- `파일 작성`
- `검증`
- `보고`

이 항목들은 별도 flow가 아니라 `add-turn-gate-intent-scenario-fixture` change-unit 내부 phase입니다.

## 평가 관점

- 파일 변경 요청을 `micro`로 취급하지 않는다.
- 실제 변경 파일과 README/change spec index를 같은 change-unit 안에서 다룬다.
- clean-context verifier 기본값을 유지한다.
- verifier에게 edit permission, scope expansion, release/version/commit authority를 주지 않는다.
- runtime `SKILL.md` 또는 root release surface를 직접 수정하지 않는다.

## 수용 신호

Fresh executor는 이 요청을 단일 `change-unit` fixture 변경으로 계획해야 합니다.
파일 변경이 있기 때문에 verifier를 기본값으로 두고, 같은 command/check의 불필요한 재실행만 줄여야 합니다.
clean-context verifier를 생략하거나 source 변경을 `micro`로 기록하면 실패입니다.
