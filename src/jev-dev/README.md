# Jev 개발

- runtime: [`../../jev/`](../../jev/)
- 플러그인 경계: [specs/plugin.md](specs/plugin.md)
- CLI·설정·출력: [specs/cli.md](specs/cli.md)
- `$jev:jev`: [specs/skills/jev.md](specs/skills/jev.md) — CLI·기본 사용 패턴·내장 유즈케이스
- `$jev:scenario-testing`: [specs/skills/scenario-testing.md](specs/skills/scenario-testing.md)
- 최초 버전: [changes/v0.1.0.md](changes/v0.1.0.md)
- primitive 확장: [changes/v0.2.0.md](changes/v0.2.0.md)
- JSONL batch: [changes/004-jsonl-batch.md](changes/004-jsonl-batch.md)
- scenario testing 이전: [changes/005-scenario-testing-migration.md](changes/005-scenario-testing-migration.md)
- 작업 중 Jev 사용: [changes/006-jev-use.md](changes/006-jev-use.md)
- 기본 패턴 통합·호출 마이그레이션: [changes/008-jev-use-integration.md](changes/008-jev-use-integration.md)
- CLI 패키지 구조: [changes/012-scripts-packages.md](changes/012-scripts-packages.md)
- provider 설정·초기 템플릿: [changes/011-provider-config.md](changes/011-provider-config.md)
- provider 어댑터: [changes/010-provider-adapter.md](changes/010-provider-adapter.md)
- 구조화 입력·장문 모드: [changes/009-structured-input.md](changes/009-structured-input.md)
- doctor 인증·출력: [changes/007-doctor-auth.md](changes/007-doctor-auth.md)

Go 코드는 `jev/scripts/`의 독립 모듈입니다. 루트 `main.go`는 실행 환경과 signal context를 준비하고 `cli.App.Run`에 위임합니다. 별도 SDK나 범용 CLI 프레임워크는 사용하지 않습니다.

| 패키지 | 책임 |
| --- | --- |
| `cli` | 최상위 명령 라우팅·오류 보고·종료 코드·provider 조립 |
| `cli/evaluation` | 단건·batch 입력·질문 검증·요청 구성·결과 기준·출력·도움말 |
| `cli/configcmd` | config 명령 조율·stdin·TOML 출력·도움말 |
| `cli/maintenance` | 설치·제거·doctor 명령 조율·옵션·보고서·도움말 |
| `config` | 설정 schema·검증·우선순위·편집·마스킹·파일 저장·내장 템플릿 |
| `decision` | Primitive·Request·Question·Result, JSON 값 처리와 출력 필드 검증 |
| `providers` | Adapter·인증 진단 계약 |
| `providers/jev` | TypeSafe 요청·응답·인증 |
| `installation` | 바이너리 복사·제거·identity와 대상 파일 보호 |

최상위 라우팅과 provider 조립을 함께 검증하는 명령 통합 테스트는 `cli`, 설치 경로 파싱 테스트는 `cli/maintenance`, 직접 저장소 테스트는 `config`에 둡니다. 하위 명령 패키지는 상위 `cli`를 import하지 않으며 입출력·설정 경로·설정 해석·어댑터 생성 함수를 전달받습니다. 바이너리 build·bootstrap 테스트는 모듈 루트를 실행 대상으로 사용합니다.

설치 동작은 Go의 `install`이 소유합니다. 최초 설치는 `jev/scripts/`에서 `go run . install`, 소스 갱신은 `go run . install --force`로 수행합니다. `uninstall`은 build info의 main package identity로 삭제 대상을 제한하고, `doctor`는 로컬 설치·PATH·설정·키 유무·출처와 인증된 models 조회를 점검합니다. 조회는 저장된 timeout으로 한 번만 호출하며 추론·redirect·retry·자동 수정은 없습니다.

```sh
cd jev/scripts
go test ./...
go vet ./...
go build -o /tmp/jev .
```

테스트는 임시 설정 경로와 가짜 HTTP transport를 사용합니다. 실제 API·사용자 설정·설치 디렉터리에 접근하지 않습니다. 설치 테스트도 임시 디렉터리만 사용합니다. 환경에서 기본 Go 캐시에 쓸 수 없으면 별도 임시 `GOCACHE`를 지정합니다.

루트 `jev --help`는 명령별 한 줄 설명과 공통 판단 옵션의 적용 범위를 표시합니다. 옵션·입력 형식·출력·종료 코드와 설정 상세는 `jev <command> --help`에서 확인합니다.

CLI 테스트·빌드는 모델 판단 정확성이나 실제 API 인증을 증명하지 않습니다. 실제 API smoke test, 사용자 홈 설치, 플러그인 설치, 커밋·푸시는 각각 별도 실행 상태로 보고합니다.

플러그인 구조 검증과 skill frontmatter 검증에 더해, 스킬 spec과 runtime은 clean-context read-only 검증으로 대조합니다. 작성용 spec을 설치된 스킬의 필수 runtime 의존성으로 만들지 않습니다. `jev`는 CLI와 기본 사용 패턴·내장 유즈케이스를 안내하고, `scenario-testing`은 prediction·fresh executor·comparison을 조율하고, `jev batch`는 공통 state 질문 전송만 소유합니다.

단건 구조화 입력은 `jev <primitive> --input request.json`, batch JSON state는 `--state-format json`, 요청별 장문 모드는 `--long-context`로 사용합니다. 신규 기준과 Score legend의 JSON 구조는 CLI가 보존합니다. 기존 명령·JSONL은 유지하며 이미지·Box·usage는 이번 확장에 포함하지 않습니다.

Provider 경계: `scripts/providers/provider.go`의 `Adapter`가 `decision.Request`를 받아 question id별 정규화 `decision.Result`를 반환합니다. `scripts/providers/jev/`의 client·response·authentication이 TypeSafe 전용 코드입니다. `scripts/cli/provider.go`는 현재 Jev 구현을 선택하는 조립 지점이며 단건·batch·doctor가 이를 공유합니다. 공통 JSON 값 처리는 `decision/content.go`에 둡니다. 새 provider는 자체 wire 요청·응답을 어댑터 안에서 변환해야 합니다.

설정 schema는 defaults와 providers.<name>으로 구분합니다. `config/template.toml`은 go:embed로 포함하며 config init·최초 set의 생성 기준입니다. 기존 flat 설정은 수동 이동을 요구하며 사용자 홈 migration은 구현 검증에서 수행하지 않습니다.
