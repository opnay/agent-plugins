# CLI 계약

- Rust로 macOS·Linux용 단일 실행 파일을 빌드합니다. Rust/Cargo는 소스 빌드에만 필요하며 실행·설정·설치에는 언어 런타임이 필요하지 않습니다. 바이너리는 빌드한 OS·아키텍처에 맞춰 사용합니다.
- TOML 처리는 `toml`, 원자적 파일 교체는 `tempfile` crate를 사용합니다. 라이브러리는 바이너리에 포함합니다.
- `install`: 현재 실행 중인 바이너리 하나를 `~/.local/bin/agent-wiki`에 복사합니다. 설치 결과가 같으면 재설치는 성공하며, 다른 파일이 있으면 보존하고 실패합니다. 명시적 `--force`로 파일 교체를 허용합니다. 일반 파일 이외의 대상은 교체하지 않습니다.
- `list`: 등록된 root 이름·경로와 default 여부를 조회합니다.
- `set <name> <directory>`: named root를 기존 디렉터리의 canonical 절대 경로로 설정합니다. 위키 폴더·문서는 생성하지 않습니다.
- `default <name>`: 기존 named root를 default로 지정합니다. `path [name]`은 default 또는 지정 root의 절대 경로를 한 줄로 출력합니다.
- 설정은 `default`, `[roots.<name>] path`와 선택 `description`을 사용합니다. 최상위 `root` 설정과 `set <directory>`는 실패하며 기존 설정 바이트를 보존합니다. 자동 변환은 하지 않습니다.
- 설정 파싱이나 갱신 결과의 root 검증에 실패하면 기존 파일을 보존합니다. 유효한 `set` 호출로 같은 root의 경로를 정정할 수 있습니다. 심볼릭 링크 설정 파일은 조회·갱신을 지원합니다. 갱신은 링크의 실제 대상 파일을 원자적으로 교체하고 링크 자체는 보존합니다. 끊어진 링크는 대상·링크를 만들거나 교체하지 않고 실패합니다.
- 설정 갱신은 임시 파일 완성 후 원자적으로 교체하며 주석·서식은 정규화됩니다.
- 실패는 stderr와 0이 아닌 종료 코드로 보고합니다. 도움말·조회는 파일을 만들지 않습니다.
- 셸 설정이나 PATH는 수정하지 않습니다. 설치 후 PATH 설정 방법만 안내합니다.

## 구조와 검증

- `agent-wiki/scripts/tests/`는 CLI 검증을 소유합니다.
- 설치·재설치, 설정·조회, 경로 특수문자, 잘못된 입력과 기존 파일 보존을 임시 홈에서 검증합니다.
- `agent-wiki/`에서 `cargo test --locked --manifest-path scripts/Cargo.toml`을 실행합니다.
- CLI 바이너리는 `agent-wiki/`에서 `sh scripts/build.sh`로 빌드합니다. 생성 바이너리와 Cargo target은 Git에 포함하지 않으며 실행 파일의 권한은 유지합니다.
