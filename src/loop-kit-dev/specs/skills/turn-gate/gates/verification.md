# turn-gate verification gate sub-spec

## 목적

이 문서는 verification gate의 전환 계약을 소유합니다.

verification gate는 flow work 결과가 보고 가능한 상태인지 판정합니다.

## 소유

- flow의 completion criteria와 verification expectation을 기준으로 검증 packet을 구성한다.
- clean-context verifier 결과를 `pass`, `fail`, `blocked`, `insufficient` 중 하나로 통합한다.
- `fail` 또는 `insufficient`이면 earliest safe gate로 되돌린다.
- `blocked`이면 user-gated question-routing으로 blocker를 보고하게 한다.

## 비소유

- 사용자 대신 approval boundary 승인
- next-flow 선택
- terminal closure

## 검토 질문

- verification gate가 non-pass 결과를 통과로 취급하지 않는가?
- blocked 상태를 user-gated question-routing으로 라우팅하는가?
