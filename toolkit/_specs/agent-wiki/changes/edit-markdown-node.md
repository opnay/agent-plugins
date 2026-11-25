# wdown edit `--id` 노드 편집

## 변경사항 요약

- `wdown edit <file> --id <id> [--write]`를 추가합니다. 표준 입력 본문으로 `inspect`의 id가 고른 노드 하나만 대체합니다.
- 기본은 dry-run이고 `-`/`+` diff를 출력합니다. `--write`가 있어야 파일을 교체합니다.
- 편집 가능한 kind는 `heading`·`paragraph`·`item`·`code`이고, 나머지 kind와 link·image 인라인 노드는 실패합니다.

## 변경 목적

`inspect`로 구조를 확인한 뒤 본문을 고칠 때 에이전트는 줄 번호 없이 문자열 일치로 파일을 편집합니다. 긴 위키 기록에서는 같은 문장이 여러 곳에 나타나고 들여쓰기·marker를 다시 맞춰야 해 오편집 위험이 큽니다. 구조 트리에서 확인한 id 하나로 대상 범위를 정하고, 저장 전에 무엇을 바꿀지 보여줄 수단이 필요합니다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: `wdown edit note.md --id 2.1 <<'EOF' … EOF`. stdout은 `edit <path>`, 선택 노드 줄, `-` 원본 줄, `+` 대체 줄, 마지막 줄에 dry-run 안내 또는 `저장했습니다: <path>`입니다.
- 호환성 영향: 기존 명령과 설정 형식은 그대로입니다. 새 명령만 추가됩니다.
- migration 필요 여부: 없습니다. 설치된 실행 파일을 새 빌드로 교체하면 바로 사용할 수 있습니다.
- 운영자/사용자가 알아야 할 검증 결과: 아래 검증 기준의 fmt·clippy·test 명령과 release build, 실제 위키 기록과 픽스처 smoke test를 통과했습니다. 테스트는 52개입니다.

## 기록 근거

- 관련 세션 기록: `/tmp/.fx/251125-edit-mPs2Zg/` (픽스처·smoke test), 이전 세션 대화

## 변경 상세

### 노드 줄 범위 기록

- 목적: id로 고른 노드가 파일의 몇 번째 줄인지 확정합니다.
- 범위: `Parser`에 원본 오프셋을 두어 blockquote와 list continuation처럼 다시 파싱하는 영역도 절대 줄 번호를 유지합니다. `block()`이 감싼 블록과 list 항목은 시작·끝 줄을 기록하고, inline code span과 link·image 노드는 줄 범위를 가지지 않습니다. `front matter`는 `---` 포함 범위를 가집니다.
- 관련 표면: `scripts/src/inspect.rs`, `scripts/tests/edit.rs`.
- 검증: 중첩 list 항목과 blockquote 안 paragraph의 줄 범위가 원본 줄 번호와 일치하는지, `apply`가 나머지 본문과 마지막 줄 바꿈 유지를 보존하는지 테스트로 확인했습니다.

### `--id` 노드 대체

- 목적: 구조 트리에서 확인한 id 하나로 편집 범위를 정합니다.
- 범위: heading은 제목 줄, paragraph는 본문 줄 전체, item은 항목 첫 줄(하위 항목·중첩 블록 보존), code는 fence 안 본문만 대체합니다. 들여쓰기·blockquote 마커는 노드의 첫 원본 줄에서 이어받고, heading marker·item의 list marker와 task checkbox·code의 fence와 info는 원본을 유지합니다. setext heading은 `#` marker로 바뀝니다.
- 관련 표면: `toolkit/_specs/agent-wiki/cli.md`, `scripts/src/edit.rs`, `scripts/src/main.rs`, `scripts/tests/edit.rs`, `scripts/tests/cli.rs`, `toolkit/agent-wiki/README.md`, `toolkit/_specs/agent-wiki/README.md`, `plugin.md`.
- 검증: kind별 대체와 marker·checkbox·fence 보존, setext 전환, item은 항목 자기 줄만 대체하고 하위 항목·중첩 블록 보존, blockquote 안 heading·item·code의 `>` 접두어 유지, 한 줄 제한, paragraph 빈 줄 거부, 빈 입력 거부, 닫지 않은 fence 거부, 편집 불가 kind와 인라인 노드 거부, 없는 id·모호 id 실패, 옵션 해석(`--id` 필수·중복 거부, `--write`, 미지정 옵션, 파일 두 개)을 테스트로 확인했습니다.

### dry-run 기본과 원자적 쓰기

- 목적: 저장 전에 변경 내용을 확인하게 합니다.
- 범위: dry-run은 파일과 표준 입력을 읽어서 diff를 출력하지만 파일은 쓰지 않습니다. 원본과 대체가 같으면 선택 노드 줄 뒤에 `변경 없음`을 출력하고 쓰지 않습니다. `--write`로 실제 변경이 있을 때만 원본 권한을 유지한 채 `atomic_write`로 교체하며, 심볼릭 링크 파일은 저장을 거부하고 링크와 대상은 그대로 둡니다. CRLF 파일은 dry-run에서도 거부합니다.
- 관련 표면: `toolkit/_specs/agent-wiki/cli.md`, `scripts/src/edit.rs`, `scripts/src/main.rs`, `scripts/tests/edit.rs`, `scripts/tests/cli.rs`.
- 검증: dry-run 출력 문자열, `변경 없음` 경로, CRLF 거부, 옵션 해석을 계획 단위 테스트로 확인했습니다. 파일 단위는 `edit_target`을 임시 파일로 실행해 dry-run 파일 불변, `--write` 저장과 원본 권한 유지, 마지막 줄 바꿈 유지·부재, 심볼릭 링크 저장 거부와 dry-run 허용, 링크·대상 보존을 테스트와 release build smoke test로 확인했습니다.

## 비목표

- 노드 삭제·이동·삽입, 섹션(heading 하위 전체) 대체, 여러 노드 동시 편집, front matter·table·list·blockquote·html 편집을 포함하지 않습니다.
- 인라인 텍스트 일부 치환, Markdown 정규화·재정렬, 줄 번호 출력, 백업 파일 생성, 파일 생성(신규 문서 작성)은 이번 범위가 아닙니다.
- skill 실행 계약과 runtime `SKILL.md`는 바꾸지 않습니다. `edit`는 CLI 사용자 표면으로만 제공합니다.
- plugin version은 0.1.0을 유지합니다.

## 호환성 및 마이그레이션

- 기존 명령(`install`·`uninstall`·`doctor`·`list`·`set`·`default`·`path`·`inspect`)과 설정 파일은 그대로입니다.
- `inspect`의 출력·id 계약은 그대로이고, parser에는 노드 줄 범위만 추가됩니다.
- 의존성 추가는 없습니다. 쓰기는 기존 `tempfile` 기반 `atomic_write`를 재사용합니다.

## 검증 기준

- `toolkit/agent-wiki/`에서 `cargo fmt --manifest-path scripts/Cargo.toml -- --check`, `cargo clippy --locked --all-targets --manifest-path scripts/Cargo.toml -- -D warnings`, `cargo test --locked --manifest-path scripts/Cargo.toml`이 통과합니다.
- `sh scripts/build.sh` 뒤 `./scripts/wdown edit <file> --id <id>`의 dry-run과 `--write`, 실제 위키 기록 복사본에 대한 heading·paragraph·item·code·blockquote 편집, 권한 유지, 심볼릭 링크 거부가 계약대로 동작합니다.
- `cli.md`의 편집 가능 kind·대체 범위·출력·쓰기 실패 조건과 `edit.rs` 구현, `scripts/tests/edit.rs` 단언이 일치합니다.

## 정식 규칙 승격 여부

- normative spec 또는 runtime 문서로 승격할 지속 규칙: `edit`의 편집 가능 kind·대체 범위·dry-run 기본·쓰기 실패 조건은 `toolkit/_specs/agent-wiki/cli.md`가, 구현 소유는 `plugin.md`의 대표 표면이 소유합니다.
- change spec에만 남길 release/migration 이력: 섹션 전체가 아니라 노드 자기 자신만 대체하기로 정한 판단, dry-run 기본으로 정한 판단, 심볼릭 링크·CRLF 파일을 거부한 판단, setext heading을 ATX로 바꾸는 동작, 노드 삭제 미지원, 줄 범위를 parser 오프셋으로 기록한 구현 판단.
