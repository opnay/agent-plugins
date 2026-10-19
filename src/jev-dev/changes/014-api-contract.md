# 기존 Jev API 계약 보완

## 변경사항 요약

- 직렬화된 요청 본문 32 MiB 제한을 전송 전에 확인합니다.
- HTTP 상태를 보존하면서 검증 오류·context 초과·크기 제한·rate limit·과부하 안내를 제공합니다.
- 오류 응답에서 입력·자격증명·서버 원문을 노출하지 않습니다.

## 변경 상세

소유 계약은 [CLI spec](../specs/cli.md)의 요청 계약입니다. 단건·batch의 정상 결과·출력·종료 코드·무재시도 정책을 유지합니다. 오류 body 해석은 Jev 어댑터 안에 두고 공통 adapter 계약은 변경하지 않습니다.

입력 개수 판단: 2026-10-19 공개 OpenAPI의 ChoiceQuestion.minProperties·ScoreQuestion.minItems는 1입니다. 설명 문서의 Score는 최소 2단계를 안내합니다. 근거가 불일치하므로 단건 flags·구조화 입력·batch 모두 기존 최소 2개 제한을 유지합니다. 실제 API 유료 호출로 한계를 시험하지 않습니다.

근거: [OpenAPI](https://api.typesafe.ai/openapi.json), [API 요청 규칙·오류](https://docs.typesafe.ai/api), [Score](https://docs.typesafe.ai/primitives/score). usage·압축·이미지·Box·새 provider·모델 변경·사용자 설정 변경·설치·commit·push는 포함하지 않습니다.

## 완료 기준

- [x] 최소 입력 개수의 근거를 확인하고 유지 여부를 기록합니다.
- [x] 32 MiB 경계 허용·초과 차단·직렬화 후 바이트 기준을 검증합니다.
- [x] 422 안전한 위치·종류 및 context·413·429·529 안내와 fallback을 검증합니다.
- [x] 민감정보 비노출·무재시도·단건/batch 종료 코드·빈 stdout 회귀를 확인합니다.
- [x] 전체 test·vet·build, 문서·패키지 독립 검증을 완료합니다.

## 검증 결과

- 전체 `go test ./...`, `go vet ./...`, 모듈 루트 `go build -trimpath`가 통과했습니다.
- 수정 전 코드는 32 MiB 초과·JSON escape 후 초과 요청을 실제 fake transport에 전달했고 오류 상세 회귀도 실패했습니다. 수정 후 경계값·차단·오류 분류가 통과했습니다.
- 단건 CLI는 종료 코드 1·빈 stdout·단일 호출을, batch는 422에서도 부분 출력 없이 실패함을 확인했습니다.
- 서버 msg·input·ctx·동적 경로에 키를 넣은 fixture, 알려진 schema 단어와 키가 같은 경우, body 초과·읽기 실패·JSON 잘림·추가 JSON·미지원 코드의 fallback과 body close를 검증했습니다.
- 실제 유료 API 호출·사용자 설정 변경·로컬 설치·커밋·푸시는 수행하지 않았습니다. context 오류의 세부 JSON 위치는 문서가 완전히 고정하지 않으므로 알려진 code/type 위치만 해석하고 나머지는 안전하게 HTTP 상태로 돌아갑니다.

두 독립 read-only 검증에서 요청 크기·오류 해석·민감정보 비노출·단건/batch 실패 계약과 문서·skill·manifest 정합성의 조치 필요 불일치는 발견되지 않았습니다.
