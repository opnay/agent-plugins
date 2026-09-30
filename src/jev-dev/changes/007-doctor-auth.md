# Jev doctor 인증·출력 Change Spec

## 변경사항 요약

- 기본 doctor는 키 유무·출처와 API 인증을 점검한다.
- 결과는 상태별 항목·해결 안내·요약으로 표시한다.

## 변경 목적

키 존재만으로 정상 처리돼 잘못된 키를 발견하지 못하는 진단 공백을 해결한다.

## 변경 상세

### 인증 조회

- 설정·키가 유효하면 Bearer 인증으로 `GET /v1/models`를 한 번 호출한다. 저장된 timeout을 적용한다.
- 401·403을 키 거부·접근 거부로 안내하며 환경변수 우선순위에 맞는 조치를 표시한다.
- 네트워크·timeout·취소·429·서버 오류는 인증 미확인으로 구분한다.
- 근거: [공개 OpenAPI](https://api.typesafe.ai/openapi.json), [공식 SDK models resource](https://docs.typesafe.ai/sdk/python/api/clients/sync#models-resource), 확인일 2026-09-30.

### 결과 표시

- Local setup·Authentication 섹션, `OK / FAIL / SKIP`, 문제별 Fix, 상태별 개수 요약을 제공한다.
- 키 값·서버 body·원문 transport 오류는 출력하지 않는다.
- 소유 표면: CLI spec, jev skill spec/runtime, README, help, manifest.

## 비목표

- 자동 수정·추론·모델 정확성 검증·redirect·retry·설치 바이너리 자동 교체·릴리즈.

## 호환성 및 마이그레이션

- `jev doctor` 기본 동작은 API 인증 조회를 포함한다. API 접근이 불가능하면 실패한다.
- 별도 offline 진단은 제공하지 않는다.
- 종료·stream 계약은 유지한다: 정상 stdout·0, 문제 빈 stdout·stderr·1.

## 검증 기준

- 키 출처·환경변수 우선순위·키 누락·설정 오류·권한·설치·PATH 회귀.
- GET·Bearer·단일 요청·timeout·취소·redirect 거부·401·403·429·서버 오류·토큰 비노출.
- 설정 오류·키 누락 시 API 미호출, 로컬 상태 보존·help·요약·문제별 해결 안내.
- Go test/vet/build, skill validator, clean-context spec/runtime 검증.

## 정식 규칙 승격 여부

지속 CLI 계약은 `specs/cli.md`, 스킬 계약은 `specs/skills/jev.md`, 사용 표면은 plugin spec·README·manifest에 반영한다.

## 검증 결과

- Go test·vet·build, skill validator, manifest JSON parsing, diff whitespace 검증 통과.
- 두 clean-context read-only 검토에서 CLI·skill·plugin 사용 표면 정합성 통과.
- 임시 빌드로 현재 설정된 키의 실제 인증 조회 성공: `5 OK, 0 FAIL, 0 SKIP`.
- `go run . install --force`로 `~/.local/bin/jev`를 교체하고 설치된 `jev doctor`의 실제 인증·PATH·설정 점검 통과를 확인했다.
