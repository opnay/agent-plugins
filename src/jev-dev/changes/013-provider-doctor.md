# Provider별 doctor 진단

## 변경사항 요약

- Local setup·Config·Provider·Authentication으로 점검과 출력을 구분합니다.
- 설정 유효성과 선택 provider의 어댑터 지원 여부를 독립적으로 판정합니다.
- 선택된 provider·모델·timeout과 해당 키의 출처를 표시합니다.

## 변경 상세

CLI 계약은 [CLI spec](../specs/cli.md)의 설치·제거·진단이 소유합니다. 미구현 provider와 잘못된 provider flag를 Config 실패로 표시하지 않습니다. 키 형식 오류는 Authentication에서 판정하고 API를 호출하지 않습니다. 모델이 없고 내장값도 없으면 SKIP으로 표시합니다.

선택 provider의 어댑터만 인증을 담당합니다. 읽기 전용·마스킹·stdout/stderr·종료 코드·단일 조회·무추론·무재시도 정책은 유지합니다. 새 provider 구현·자동 설정 변경·원격 모델 존재 검증은 포함하지 않습니다.

## 완료 기준

- [x] 네 섹션과 provider·모델·timeout 표시를 구현합니다.
- [x] Config·provider 지원·키 오류를 구분하며 실패 시 잘못된 어댑터나 API를 호출하지 않습니다.
- [x] 기존 마스킹·인증·설치·설정 우선순위 회귀와 미구현 provider 회귀를 검증합니다.
- [x] 관련 도움말·README·CLI spec과 독립 정적 검증을 맞춥니다.

실행 검증: 전체 `go test ./...`, `go vet ./...`, 모듈 루트 build가 통과했습니다. 미구현 provider·provider override·알 수 없는 provider·환경변수 키 형식 오류를 정상 Config와 구분하는 회귀는 수정 전 실패·수정 후 통과를 확인했습니다. 모델 필드에 다른 provider의 공백 포함 키와 같은 값이 표시되는 경우도 재현하고 원문·공백 제거 값 모두 마스킹하도록 보완했습니다. 임시 OpenRouter 설정으로 Config OK·Adapter FAIL·API SKIP 출력을 확인했으며 실제 API 호출·사용자 설정 변경·설치 바이너리 갱신은 수행하지 않았습니다.

독립 정적 검증은 점검 분류·선택 provider 키 사용·읽기 전용·출력 흐름을 확인했습니다. 검토 중 제기된 앞뒤 개행은 기존 Resolve의 TrimSpace 계약으로 정규화되며, 형식 검사는 해석 후 남은 개행을 대상으로 한다고 CLI spec을 명확히 했습니다. README의 키 형식 오류 설명도 보완했습니다. 스킬은 기존 포괄적 점검·상태 해석 계약과 모순이 없어 변경하지 않았습니다.
