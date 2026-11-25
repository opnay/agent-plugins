# wdown inspect Markdown 구조 트리

## 변경사항 요약

- `wdown inspect <file> [--id <id>]`를 추가합니다. Markdown 한 파일을 문법 블록 기준 구조 트리로 출력합니다.
- `--id`는 자리 번호(`1.2`)와 heading slug를 받고 해당 노드의 하위 트리만 출력합니다.
- 구조 문제(`heading-skip`·`duplicate-slug`·`missing-alt`·`empty-link-text`·`table-column-mismatch`)를 노드 줄과 첫 줄 합계에 표시합니다.

## 변경 목적

위키 기록과 skill 문서를 눈으로 볼 때 heading 중첩, 목록 중첩, table 열, link·image 속성 같은 구조 문제가 본문 텍스트만으로는 드러나지 않습니다. Markdown 파서의 구조 판정을 한 줄씩 확인해 기록 구조와 문서 품질을 점검할 수단이 필요합니다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: `wdown inspect <file>`과 `wdown inspect <file> --id <id>`. 첫 줄은 `document <path> nodes=<n> flags=<n>`이고 각 줄은 `<id> <kind>` 뒤에 details 키·본문·slug·flag 순서로 붙습니다.
- 호환성 영향: 기존 명령과 설정 형식은 그대로입니다. 새 명령만 추가됩니다.
- migration 필요 여부: 없습니다. 설치된 실행 파일을 새 빌드로 교체하면 바로 사용할 수 있습니다.
- 운영자/사용자가 알아야 할 검증 결과: 아래 검증 기준의 fmt·clippy·test 명령과 release build, 실제 위키 기록과 픽스처 smoke test를 통과했습니다. 테스트는 33개입니다.

## 기록 근거

- 관련 세션 기록: `/tmp/.fx/251125-inspect/` (픽스처), 이전 세션 대화

## 변경 상세

### Markdown 구조 트리 출력

- 목적: Markdown 문법 기준으로 블록·inline 구조를 트리 한 번에 보여줍니다.
- 범위: `document`·`front matter`·`heading <level>`·`paragraph`·`list ordered|unordered`·`item`·`blockquote`·`code`·`table`·`thematic break`·`html`·`link`·`image` 노드와 details 키(`keys=`·`items=`·`task=`·`blocks=`·`lines=`·`info=`·`rows=`·`cols=`·`dest=`·`target=`·`code=`)·flag. inline code span은 노드가 아니라 부모 details의 `code=<n>` 개수입니다. 블록 시작 판정은 들여쓰기를 보지 않습니다.
- 관련 표면: `toolkit/_specs/agent-wiki/cli.md`, `scripts/src/inspect.rs`, `scripts/src/main.rs`, `scripts/tests/inspect.rs`, `toolkit/agent-wiki/README.md`, `toolkit/_specs/agent-wiki/README.md`, `plugin.md`.
- 검증: front matter keys, setext·ATX heading과 섹션 중첩, 수준 건너뛰기 flag, list marker별 분리와 task checkbox, fence·blockquote 중첩, table 열 판정, link·image·Obsidian wikilink와 `missing-alt`·`empty-link-text`, `code=<n>` 계수, 4칸 들여쓰기 줄의 블록 시작과 code 미인식을 테스트로 확인했습니다.

### `--id` 하위 트리 선택

- 목적: 큰 문서에서 관심 있는 섹션만 다시 보는 방법을 제공합니다.
- 범위: 자리 번호는 직계 자식부터 매기고 문서에서 유일합니다. heading에는 slug를 병기하고 `#` 접두어는 선택입니다. slug 검색은 모든 하위 노드를 대상으로 하며, 형제든 하위든 두 개 이상과 일치하면 실패하고 후보 자리 번호를 stderr에 출력합니다.
- 관련 표면: `toolkit/_specs/agent-wiki/cli.md`, `scripts/src/inspect.rs`, `scripts/src/main.rs`, `scripts/tests/inspect.rs`.
- 검증: 자리 번호·slug·`#` 접두어 선택, 없는 id 실패, 형제 중복 slug와 중첩 slug 모호 실패, 없는 파일·UTF-8 아닌 파일·`--id` 중복·미지정 옵션·파일 두 개 지정 실패를 테스트로 확인했습니다.

## 비목표

- Markdown 렌더링, HTML 출력, 파일 쓰기·수정, 디렉터리 재귀 점검, 여러 파일 동시 점검을 포함하지 않습니다.
- 들여쓰기 4칸 code block 인식, CommonMark reference link 정의 해석, 각주·정의 목록 같은 확장 문법, Obsidian 속성 값 검증은 이번 범위가 아닙니다.
- skill 실행 계약과 runtime `SKILL.md`는 바꾸지 않습니다. `inspect`는 CLI 사용자 표면으로만 제공합니다.
- plugin version은 0.1.0을 유지합니다.

## 호환성 및 마이그레이션

- 기존 명령(`install`·`uninstall`·`doctor`·`list`·`set`·`default`·`path`)과 설정 파일은 그대로입니다.
- `inspect`는 읽기 전용이며 설정 파일을 읽지 않고 홈 디렉터리에 의존하지 않습니다.
- 의존성 추가는 없습니다. Markdown 파싱은 `scripts/src/inspect.rs`의 자체 구현이며 Cargo 의존성은 `toml`·`tempfile` 그대로입니다.

## 검증 기준

- `toolkit/agent-wiki/`에서 `cargo fmt --manifest-path scripts/Cargo.toml -- --check`, `cargo clippy --locked --all-targets --manifest-path scripts/Cargo.toml -- -D warnings`, `cargo test --locked --manifest-path scripts/Cargo.toml`이 통과합니다.
- `sh scripts/build.sh` 뒤 `./scripts/wdown inspect <file>`과 `--id` 선택이 실제 위키 기록에서 계약대로 출력합니다.
- `cli.md`의 노드 종류·id·details·flag·실패 조건과 `inspect.rs` 출력, `scripts/tests/inspect.rs` 단언이 일치합니다.

## 정식 규칙 승격 여부

- normative spec 또는 runtime 문서로 승격할 지속 규칙: `inspect` 출력·id·flag·실패 계약과 파싱 한계는 `toolkit/_specs/agent-wiki/cli.md`가, 구현 소유는 `plugin.md`의 대표 표면이 소유합니다.
- change spec에만 남길 release/migration 이력: inline code span을 노드에서 details 개수로 바꾼 판단, 들여쓰기 code block 미인식과 들여쓰기 무시 블록 시작 판정 결정, slug 검색을 모든 하위 노드로 넓혀 중첩 중복 slug도 모호로 실패하게 바꾼 판단, 의존성 추가 없이 자체 파싱을 택한 이유.
