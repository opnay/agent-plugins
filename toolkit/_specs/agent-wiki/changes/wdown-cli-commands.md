# wdown 실행 파일명과 설치·제거·점검 명령

## 변경사항 요약

- CLI 실행 파일과 설치 파일 이름을 `agent-wiki`에서 `wdown`으로 바꿉니다. plugin 이름과 `$agent-wiki:*` skill 식별자는 유지합니다.
- `install`·`uninstall`·`doctor`를 추가합니다. 세 명령은 `--dir <directory>`를 받고 기본값은 `~/.local/bin`입니다. `--force`는 install만 받습니다.
- plugin version은 0.1.0을 유지합니다.

## 변경 목적

실행 파일명이 plugin 식별자와 같아서 CLI 호출과 plugin·skill 참조를 구분하기 어려웠습니다. 짧은 별칭으로 호출을 단순화합니다.
설치 상태 확인과 제거 수단이 없어 PATH 미설정·미설치 상태를 사용자가 직접 판별해야 했습니다. 설치·제거·점검을 한 실행 파일 안에서 제공합니다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: `wdown install`·`wdown uninstall`·`wdown doctor`와 `wdown list`·`set`·`default`·`path`. `doctor`는 항목마다 `OK`·`FAIL`·`SKIP`과 합계를 출력합니다.
- 호환성 영향: 이전 `agent-wiki` 실행 파일은 제거되지 않고 그대로 남습니다. 새 호출 이름은 `wdown`입니다.
- migration 필요 여부: `wdown install` 실행 후 기존 `~/.local/bin/agent-wiki`를 사용자가 삭제합니다. 설정 파일과 root 설정은 그대로 사용합니다.
- 운영자/사용자가 알아야 할 검증 결과: `cargo fmt --check`, `cargo clippy --locked --all-targets -- -D warnings`, `cargo test --locked` 22개, release build와 install·doctor·uninstall smoke test를 통과했습니다.

## 기록 근거

- 관련 세션 기록: `/tmp/.fx/251125-gM4NRu/wdown-rename.md`

## 변경 상세

### 실행 파일명 변경

- 목적: plugin 식별자와 CLI 호출 이름을 분리합니다.
- 범위: Cargo package명, 빌드 산출물, 설치 대상 경로, 도움말과 오류 메시지, README·spec 참조.
- 관련 표면: `toolkit/agent-wiki/scripts/Cargo.toml`, `scripts/build.sh`, `scripts/src/main.rs`, `.gitignore`, `toolkit/agent-wiki/README.md`, `toolkit/_specs/agent-wiki/README.md`, `toolkit/_specs/agent-wiki/cli.md`, `plugin.md`, `wiki.md`, `obsidian.md`, skill spec과 runtime `SKILL.md`의 CLI 호출.
- 검증: `sh scripts/build.sh`가 `scripts/wdown`을 만들고 `wdown --help`가 새 사용법을 출력합니다.

### 설치·제거·점검 명령

- 목적: 설치 파일 존재, PATH 노출, 설정 파일 상태를 한 번에 확인하고 설치 파일을 안전하게 제거합니다.
- 범위: `install [--force] [--dir <dir>]`, `uninstall [--dir <dir>]`, `doctor [--dir <dir>]`. `doctor`는 읽기 전용입니다.
- 관련 표면: `toolkit/_specs/agent-wiki/cli.md`, `scripts/src/main.rs`, `scripts/tests/cli.rs`.
- 검증: install의 기본 설치 폴더·`~` 별칭·충돌 보존·실행 권한, uninstall의 심볼릭 링크·디렉터리·비실행 파일 보존과 없는 대상 성공, doctor의 `OK`·`FAIL`·`SKIP` 판정과 종료 코드를 테스트로 확인했습니다.
- 계약 정정: 설치 파일 0755와 설정 파일 0600, install의 상위 폴더 생성, `--help`·`-h` 처리, `list`의 description 필드, `set`의 `~` 해석, doctor의 PATH 선택 기준을 `cli.md`에 명시했습니다.

## 비목표

- plugin 이름·폴더·`$agent-wiki:*` skill 식별자 변경, plugin version 상향, 이전 `agent-wiki` 실행 파일 자동 제거, 설정 파일 형식 변경, Windows 지원, 자동 업데이트를 포함하지 않습니다.

## 호환성 및 마이그레이션

- plugin과 skill 호출은 그대로 사용합니다. CLI 호출만 `agent-wiki`에서 `wdown`으로 바꿉니다.
- 이전 실행 파일은 사용자가 확인 후 삭제합니다. 두 실행 파일이 함께 있으면 PATH 앞에서 잡히는 쪽이 사용됩니다.
- `doctor`는 PATH 밖 설치와 잘못된 설정을 FAIL로 보고하지만 자동으로 고치지 않습니다.

## 검증 기준

- `toolkit/agent-wiki/`에서 `cargo fmt --manifest-path scripts/Cargo.toml -- --check`, `cargo clippy --locked --all-targets --manifest-path scripts/Cargo.toml -- -D warnings`, `cargo test --locked --manifest-path scripts/Cargo.toml`이 통과합니다.
- `sh scripts/build.sh` 뒤 `./scripts/wdown --help`, `install --dir <tmp>`, `doctor --dir <tmp>`, `uninstall --dir <tmp>`가 계약대로 동작합니다.
- 두 host manifest와 marketplace가 plugin 이름·경로·version 0.1.0을 유지합니다.

## 정식 규칙 승격 여부

- normative spec 또는 runtime 문서로 승격할 지속 규칙: 설치·제거·점검 계약과 `wdown` 실행 파일명은 `toolkit/_specs/agent-wiki/cli.md`와 `plugin.md`가 소유합니다. runtime `SKILL.md`에는 root 조회 호출 식별자만 반영했습니다.
- change spec에만 남길 release/migration 이력: 이전 `agent-wiki` 실행 파일 보존과 사용자 삭제, version 0.1.0 유지 결정.
