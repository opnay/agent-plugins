# turn-gate verification phase sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `verification` phase 세부 계약을 소유합니다.

## 계약

- 이 단계는 verification gate를 통과한다.
- 현재 flow의 work 결과를 검증한다.
- 파일 수정이라면 적용 여부, 타입 오류, 테스트/빌드/린트 같은 해당 검증 경로를 확인한다.
- 조사나 판단 작업이라면 다양한 관점에서 논리 비판과 반례 검토를 수행한다.
- work 뒤 reporting 전에 risk-based verification method를 선택하고 수행한다.
- verification method는 `clean-context`, `normal`, `not-required` 중 하나다.
- `clean-context` verification은 필요한 위험 조건에서 유지하되, work 중 이미 실행해 충분히 기록된 check를 근거 없이 다시 실행하지 않는다.
- 파일 변경 없는 조사나 판단 작업에서는 `normal` method로 evidence readback, 논리 반례, 사용자 의도 부합성 검토를 수행할 수 있다.
- 검증할 work output이 없는 blocker, activation-only, next-flow selection에서는 `not-required` method를 사용할 수 있지만 이유와 residual uncertainty를 기록한다.
- 검증 packet, pass/fail/blocked/insufficient 처리, non-pass return path는 `records/verification.md`가 소유한다.

## 검토 질문

- work 결과를 reporting 전에 검증했는가?
- verification method가 현재 flow의 위험도와 verification expectation에 맞는가?
- clean-context가 필요한 작업에서 bounded packet이 요청됐는가?
- bounded packet 또는 normal evidence가 중복 검증을 요구하지 않고 현재 flow의 verification expectation에 맞게 좁혀졌는가?
- fail, blocked, insufficient를 pass처럼 보고하지 않았는가?
