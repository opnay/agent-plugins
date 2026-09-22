# Jev 개발

- runtime: [`../../jev/`](../../jev/)
- 플러그인 경계: [specs/plugin.md](specs/plugin.md)
- CLI·설정·출력: [specs/cli.md](specs/cli.md)
- `$jev:jev`: [specs/skills/jev.md](specs/skills/jev.md)
- 최초 버전: [changes/v0.1.0.md](changes/v0.1.0.md)
- primitive 확장: [changes/v0.2.0.md](changes/v0.2.0.md)
- JSONL batch: [changes/004-jsonl-batch.md](changes/004-jsonl-batch.md)

Go 코드는 `jev/scripts/`의 독립 모듈입니다. 명령 라우팅, primitive, 옵션, 설정 파일 I/O, HTTP API, 응답 검증, 결과 출력, 설치·제거·진단을 책임별 파일로 유지합니다. 별도 SDK나 범용 CLI 프레임워크는 사용하지 않습니다.

설치 동작은 Go의 `install`이 소유합니다. 최초 설치는 `jev/scripts/`에서 `go run . install`, 소스 갱신은 `go run . install --force`로 수행합니다. `uninstall`은 build info의 main package identity로 삭제 대상을 제한하고, `doctor`는 로컬 설치·PATH·설정·토큰 유무만 읽습니다.

```sh
cd jev/scripts
go test ./...
go vet ./...
go build -o /tmp/jev .
```

테스트는 임시 설정 경로와 가짜 HTTP transport를 사용합니다. 실제 API·사용자 설정·설치 디렉터리에 접근하지 않습니다. 설치 테스트도 임시 디렉터리만 사용합니다. 환경에서 기본 Go 캐시에 쓸 수 없으면 별도 임시 `GOCACHE`를 지정합니다.

CLI 테스트·빌드는 모델 판단 정확성이나 실제 API 인증을 증명하지 않습니다. 실제 API smoke test, 사용자 홈 설치, 플러그인 설치, 커밋·푸시는 각각 별도 실행 상태로 보고합니다.

플러그인 구조 검증과 skill frontmatter 검증에 더해, 스킬 spec과 runtime은 clean-context read-only 검증으로 대조합니다. 작성용 spec을 설치된 스킬의 필수 runtime 의존성으로 만들지 않습니다.
