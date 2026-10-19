# Jev provider 어댑터 분리

## 변경사항 요약

- 공통 decision 요청·결과와 provider 인터페이스를 `scripts/providers`에 둡니다.
- TypeSafe HTTP·응답 검증·인증 진단을 `scripts/providers/jev`에 모읍니다.
- CLI 단건·batch·doctor는 어댑터를 사용하며 기존 동작을 유지합니다.

## 변경 상세

OpenRouter·OpenAI 등의 decision 모델 연결을 위한 내부 구조 변경입니다. 이번에는 Jev 어댑터만 구현합니다. provider 선택 flag·새 인증 설정·다른 provider 구현·설치·커밋·푸시는 포함하지 않습니다.

CLI는 입력 검증·설정 해석·결과 기준·출력·종료 코드를 소유하며, Jev 어댑터는 provider 요청과 검증된 결과 변환을 소유합니다. 실제 API 응답 타입은 어댑터 내부에 유지합니다. 단건·batch가 같은 정규화 결과를 사용하며 실패 응답은 stdout 전에 차단합니다.

검증: 기존 가짜 HTTP 기반 CLI·batch·doctor 회귀, 공통 어댑터를 통한 요청·결과 경계, 전체 test·vet·build를 확인합니다. 실제 API 추론·새 provider 동작은 검증 범위가 아닙니다.
