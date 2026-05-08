# turn-gate phase-verification sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `verification` phase 세부 계약을 소유합니다.

## 계약

- 이 단계는 verification gate를 통과한다.
- 현재 flow의 work 결과를 검증한다.
- 파일 수정이라면 적용 여부, 타입 오류, 테스트/빌드/린트 같은 해당 검증 경로를 확인한다.
- 조사나 판단 작업이라면 다양한 관점에서 논리 비판과 반례 검토를 수행한다.
- work 뒤 reporting 전에 clean-context verification을 수행한다.
- 검증 packet, pass/fail/blocked/insufficient 처리, non-pass return path는 `verification.md`가 소유한다.

## 검토 질문

- work 결과를 reporting 전에 검증했는가?
- clean-context verification이 bounded packet으로 요청됐는가?
- fail, blocked, insufficient를 pass처럼 보고하지 않았는가?
