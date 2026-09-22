# Jev 004: JSONL Batch Evaluation

## 변경사항 요약

- 공통 state와 JSONL 질문을 한 SystemOne 요청으로 보내는 `jev batch`를 추가한다.
- batch 입력을 전체 선검증하고 성공 결과를 입력 순서 JSONL로 출력한다.
- 단건 threshold·min_score·json·pick 계약을 batch와 분리한다.

## 변경 목적

같은 target instruction을 기준으로 여러 질문을 평가할 때 state와 네트워크 요청을 반복하지 않고 SystemOne의 다중 `questions` 계약을 직접 사용한다. 서버 cache 동작은 보장하지 않으며 batch의 근거는 공통 state와 단일 요청이다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: `jev batch --state-file <path> --input <path>`와 JSONL 출력.
- 호환성 영향: 기존 단건 명령과 출력·종료 코드는 유지한다.
- migration 필요 여부: 반복 단건 호출자는 공통 state를 분리할 수 있을 때 선택적으로 batch로 바꿀 수 있다.
- 운영자/사용자가 알아야 할 검증 결과: 실제 API cache·비용 절감은 로컬 테스트가 증명하지 않는다.

## 기록 근거

- 관련 커밋: 없음.
- 관련 세션 기록: 현재 JSONL batch 설계·구현 요청.

## 변경 상세

### Batch Input And Request

- 목적: 한 state에 대한 여러 primitive 질문을 한 요청으로 평가한다.
- 범위: state file, JSONL row schema, primitive별 criteria, model·timeout.
- 관련 표면: CLI spec, `jev/scripts/`, CLI help, README.
- 검증: 전체 입력이 요청 전에 검증되고 question ID가 row ID와 일치하며 API 호출이 한 번이어야 한다.

### Atomic Output And Failure

- 목적: 부분 결과와 재개 상태를 CLI 계약에 섞지 않는다.
- 범위: 입력 순서 JSONL, 응답 완전성, stdout·stderr·exit code.
- 관련 표면: CLI spec, batch implementation과 tests.
- 검증: 전체 결과를 stdout 쓰기 전에 encode하고 성공 시 exit 0, 쓰기 전 실패는 빈 stdout과 exit 1이어야 한다.

## 비목표

- 자동 chunk, retry, resume, 부분 성공을 제공하지 않는다.
- row별 threshold, min_score, json, pick을 제공하지 않는다.
- 서로 다른 state를 자동 grouping하지 않는다.
- 서버 prompt cache 동작이나 비용 절감을 보장하지 않는다.

## 호환성 및 마이그레이션

- 기존 `noul`, `choice`, `score`, legacy Choice 호출은 변경하지 않는다.
- batch는 별도 명령과 출력 계약을 사용하며 exit 2를 반환하지 않는다.
- 반복 호출은 모든 질문이 같은 state를 공유할 때만 batch로 바꾼다.

## 검증 기준

- Go test·vet·build가 통과해야 한다.
- batch 입력·API request·response·output·failure tests가 통과해야 한다.
- CLI·skill spec과 runtime, README, manifest가 같은 batch 경계를 설명해야 한다.
- canonical skill·plugin validator, manifest JSON parse, `git diff --check`가 통과해야 한다.
- clean-context verifier가 단건 회귀와 batch spec/runtime/implementation 정합성을 독립 확인해야 한다.

## 정식 규칙 승격 여부

- 지속 CLI 계약은 `src/jev-dev/specs/cli.md`가 소유한다.
- 사용 기준은 `src/jev-dev/specs/skills/jev.md`가 소유한다.
- cache 비보장과 이번 migration 범위는 이 change spec에 남긴다.
