# CLI 패키지 구조 정리

## 변경사항 요약

- scripts 루트는 main·모듈 파일만 유지합니다.
- 명령 조율은 cli, 설정은 config, 공통 판단 계약은 decision, 바이너리 관리는 installation으로 구분합니다.
- cli 내부는 evaluation·configcmd·maintenance 명령군으로 구분합니다.
- providers/jev의 기존 경로와 go run . install·바이너리 identity를 유지합니다.

## 변경 상세

구조 계약은 specs/cli.md의 패키지 구조와 의존성이 소유합니다. 기존 명령·설정 schema·기본값·출력·종료 코드·통신·파일 보호 정책은 바꾸지 않습니다. 새 provider 구현·설정 migration·사용자 홈 설치·푸시를 포함하지 않습니다.

기존 회귀 테스트를 해당 소유 패키지로 옮기고, 명령 테스트의 바이너리 build·go run은 모듈 루트를 대상으로 유지합니다. 검증은 전체 test·vet·build, 설치 bootstrap과 identity, 문서 경로·순환 의존성·독립 계약 대조를 포함합니다.

## 완료 기준

- [x] provider·config 선행 변경을 서명 커밋으로 보존합니다: `906ee61`.
- [x] scripts 루트는 main.go·go.mod·go.sum만 유지합니다.
- [x] cli·config·decision·providers·installation 책임과 테스트 위치를 맞춥니다.
- [x] 내장 템플릿·문서 참조를 새 경로에 맞춥니다.
- [x] 기존 명령·출력·종료 코드·파일 보호·bootstrap·identity 회귀 검증을 통과합니다.
- [x] test·vet·build 및 독립 문서·의존성 검증을 완료합니다.

사용자 홈 설정·설치 바이너리·원격 저장소는 변경하지 않습니다.

검증 결과: 기존 테스트 49개를 유지한 `go test ./...`, `go vet ./...`, 모듈 루트 build가 통과했습니다. 설치·제거·go run bootstrap 회귀와 build info의 `opnay/jev` identity를 확인했습니다. 두 clean-context read-only 검증에서 패키지 의존성·책임·문서·실행 진입점 불일치는 발견되지 않았습니다. 실제 API 호출·사용자 홈 설치는 수행하지 않았습니다.

하위 명령 분리 완료 기준:

- [x] cli 루트에 라우팅·provider 조립·최상위 도움말을 유지합니다.
- [x] evaluation·configcmd·maintenance는 상위 cli를 import하지 않습니다.
- [x] 기존 명령·도움말·오류·파일 보호·설정 우선순위·실행 진입점을 유지합니다.
- [x] 기존 회귀 테스트와 test·vet·build·독립 정적 검증을 통과합니다.

하위 명령 분리 실행 검증: `go test ./...`, `go vet ./...`, 모듈 루트 build와 `opnay/jev` build identity 확인이 통과했습니다. 설치 경로 파싱 테스트를 별도 패키지로 나누어 총 테스트는 50개이며, 기존 테스트는 모두 유지했습니다. 모델·timeout flag 우선순위는 실제 요청의 model과 deadline으로 검증합니다. 도움말 상수 13개의 본문은 분리 전과 동일합니다.

하위 명령 독립 정적 검증: 패키지 의존 방향·명령 소유권·문서 경로·설치 identity·보고서 전달은 일치했습니다. 다만 기존 `main.go`는 help 분기 전에 환경변수 토큰 세 개를 조회하며, CLI spec의 help 토큰 비조회 문구와 불일치합니다. `906ee61`에서도 같은 선조회가 존재하며 이번 패키지 분리로 발생한 회귀는 아닙니다. 동작 보존 범위에 따라 이번 변경에서 수정하지 않았습니다.

## 도움말 노출 범위

- [x] 루트는 8개 명령의 한 줄 설명·공통 판단 옵션 요약·개별 `--help` 안내를 표시합니다.
- [x] 명령 전용 옵션·schema·설정 key·출력·종료 코드 상세는 해당 명령 도움말에 둡니다.
- [x] 기존 도움말 회귀 테스트와 독립 정적 검증을 완료합니다.

사용자가 확인한 도움말 의도에 맞춰 기존 루트의 상세 중복을 제거합니다. 위의 도움말 13개 본문 보존 결과는 하위 패키지 분리 시점의 기록이며, 이번 수정은 그중 루트 도움말의 노출 범위를 변경합니다. 명령별 도움말 본문과 실행 동작은 유지합니다.

루트 공통 옵션은 `--provider`, `--model`, `--timeout`, `--long-context`로 제한하고 단건·batch 적용 범위와 doctor의 provider 지원을 안내합니다. 단건 전용 `--json`·`--pick`·판정 기준과 설치 옵션은 해당 명령 도움말에 유지합니다.

도움말 검증: 전체 `go test ./...`와 실제 `go run . --help` 출력 확인이 통과했습니다. 독립 read-only 검증에서도 공통 옵션의 적용 범위·명령별 도움말 라우팅·상세 설명 분리가 현재 계약과 일치했습니다.
