# Provider 설정·초기 템플릿

## 변경사항 요약

- 공통 설정을 defaults, 모델·인증을 providers.<name>으로 구분합니다.
- 요청별 provider 선택과 provider별 환경변수·키 마스킹을 처리합니다.
- config init과 최초 set에 같은 내장 TOML 템플릿을 사용합니다.

## 변경 상세

Provider 어댑터 확장에 맞춘 설정 구조입니다. 현재 실행은 Jev만 제공하며 OpenRouter·OpenAI 어댑터 구현은 포함하지 않습니다. schema·우선순위·생성·실패 계약은 specs/cli.md가 소유합니다.

기존 flat 설정은 defaults 및 providers.jev로 수동 이동해야 합니다. 읽기·init·set은 기존 파일을 자동 migration하지 않습니다. 사용자 홈 설정 변경·설치·commit·push는 이번 구현에 포함하지 않습니다.

검증: 내장값과 템플릿 일치, init/최초 set의 내용·0600·기존 파일 보존, dotted set/unset, provider 우선순위·키 격리·전체 마스킹, 구형/혼합 설정 거부, 읽기 비생성, 단건·batch·doctor 회귀를 확인합니다.
