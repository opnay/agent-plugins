# CLI 계약

- Rust로 macOS·Linux용 단일 실행 파일을 빌드합니다. 실행 파일과 설치 파일 이름은 `wdown`입니다. Rust/Cargo는 소스 빌드에만 필요하며 실행·설정·설치에는 언어 런타임이 필요하지 않습니다. 바이너리는 빌드한 OS·아키텍처에 맞춰 사용합니다.
- TOML 처리는 `toml`, 원자적 파일 교체는 `tempfile` crate를 사용합니다. Markdown 구조 파싱은 외부 crate 없이 자체 파서를 사용합니다. 라이브러리는 바이너리에 포함합니다.
- 설치·제거·점검 명령은 `<directory>/wdown`을 대상으로 합니다. 기본 `<directory>`는 `~/.local/bin`이며 `--dir <directory>`로 바꿉니다. `--dir`은 세 명령이 공통으로 받고 `--force`는 install만 받습니다. 각 옵션은 한 번만 지정할 수 있고, 알 수 없는 옵션은 실패합니다. `~`로 시작하는 `<directory>`는 홈 폴더를 풀어 해석하며 빈 값은 거부합니다.
- `install [--force] [--dir <directory>]`: 현재 실행 중인 바이너리 하나를 `<directory>/wdown`에 복사합니다. 상위 폴더가 없으면 만듭니다. 설치 파일 권한은 0755입니다. 대상 파일의 내용과 실행 권한이 현재 실행 파일과 같으면 재설치는 성공하며, 다르면 보존하고 실패합니다. 명시적 `--force`로 파일 교체를 허용합니다. 일반 파일 이외의 대상은 교체하지 않습니다.
- `uninstall [--dir <directory>]`: `<directory>/wdown`을 제거합니다. 실행 권한이 있는 일반 파일만 제거하고 심볼릭 링크·디렉터리·실행 권한이 없는 파일은 보존하며 실패합니다. 대상이 없으면 출력 없이 성공합니다. 설정 파일과 셸 설정은 변경하지 않습니다.
- `doctor [--dir <directory>]`: 설치·PATH·설정 상태를 읽기 전용으로 점검합니다. 항목마다 `OK`·`FAIL`·`SKIP` 한 줄과 요약 행을 stdout에 출력하고, `FAIL`이 있으면 0이 아닌 종료 코드로 끝냅니다. 파일을 생성·수정·삭제하지 않고 네트워크 요청을 하지 않습니다.
  - install: `<directory>/wdown`이 일반 실행 파일이며 현재 실행 파일과 같은 내용인지 확인합니다. 없으면 `SKIP`, 다르거나 실행 권한이 없거나 일반 파일이 아니면 `FAIL`입니다.
  - path: PATH 항목 앞에서부터 `wdown` 일반 파일을 찾아 먼저 선택되는 파일을 확인합니다. 선택 파일이 없으면 `FAIL`, 설치 파일과 다른 파일이면 `FAIL`, 같거나 설치 파일이 없을 때 다른 경로에서 선택되면 `OK`입니다. 셸 설정과 PATH는 수정하지 않고 설정 방법만 안내합니다.
  - config: `~/.agents/config.wiki.toml`을 파싱하고 root 경계를 검증한 결과와 파일 권한을 보고합니다. 설정이 없으면 `SKIP`이며 `set` 호출을 안내합니다. 손상·미지원 설정과 끊어진 링크는 `FAIL`이며 기존 파일을 보존합니다.
- `list`: 등록된 root 이름·canonical 경로·default 여부·선택 description을 탭 구분 한 줄로 조회합니다. description이 없으면 `-`를 출력합니다.
- `set <name> <directory>`: named root를 기존 디렉터리의 canonical 절대 경로로 설정합니다. `<directory>`가 `~`로 시작하면 홈 폴더를 풀어 해석합니다. 위키 폴더·문서는 생성하지 않습니다.
- `default <name>`: 기존 named root를 default로 지정합니다. `path [name]`은 default 또는 지정 root의 절대 경로를 한 줄로 출력합니다.
- `inspect <file> [--id <id>]`: Markdown 한 파일을 문법 블록 기준으로 파싱해 구조 트리로 출력합니다. 읽기 전용이며 파일을 생성·수정하지 않고 네트워크 요청을 하지 않습니다. `--id`는 한 번만 받고 알 수 없는 옵션은 실패합니다. 옵션은 `--` 접두어만 인식하고 그 외 인자는 파일로 취급합니다. 파일은 하나만 받습니다.
  - 노드 종류는 `document`·`front matter`·`heading <level>`·`paragraph`·`list ordered|unordered`·`item`·`blockquote`·`code`·`table`·`thematic break`·`html`·`link`·`image`입니다. heading은 같은 수준 이상의 다음 heading까지 오는 블록을 하위에 둡니다. inline code span은 노드로 만들지 않고 부모 노드 details에 `code=<n>` 개수로 표시합니다.
  - id는 직계 자식부터 `1`·`1.2`처럼 자리 번호로 매깁니다. 자리 번호는 문서에서 유일합니다. heading에는 본문 텍스트에서 만든 slug를 병기하며 `--id`는 자리 번호와 slug를 모두 받습니다. slug의 `#` 접두어는 선택이고 검색은 모든 하위 노드를 대상으로 합니다. slug가 두 개 이상 노드와 일치하면 실패하고 후보 자리 번호를 stderr에 출력합니다.
  - 각 줄은 2칸 들여쓰기로 계층을 나타내고 `<id> <kind>` 뒤에 details 키와 본문 따옴표, heading slug, `flag=<name>` 순서로 붙입니다. 첫 줄은 `document <path> nodes=<n> flags=<n>`이며 `--id` 사용 때는 `subtree <id> nodes=<n> flags=<n>` 뒤에 선택 노드 줄을 다시 출력합니다. `nodes`는 document 자기 자신을 포함하고 `flags`는 하위 전부의 합계이며 `<path>`는 지정받은 문자열을 그대로 표시합니다.
  - 본문 details는 첫 줄을 60자로 잘라 따옴표로 표시하고 잘림을 `…`로 나타냅니다. table은 `rows=<n> cols=<n>`(rows는 본문 행만 셈), list는 `items=<n>`, item의 task checkbox는 `task=done|open`, front matter는 `keys=<n>`, blockquote는 `blocks=<n>`, code와 html은 `lines=<n>`, code의 fence info는 `info=<info>`, link·image의 일반 대상은 `dest=<대상>`, Obsidian wikilink는 `target=<본문>`을 details로 표시합니다.
  - flag는 heading 수준 건너뛰기(`heading-skip`)·중복 slug(`duplicate-slug`)·alt 없는 image(`missing-alt`)·텍스트 없는 link(`empty-link-text`)·table 열 개수 불일치(`table-column-mismatch`)입니다. 노드 줄에 `flag=<name>`으로 붙이고 첫 줄의 `flags`에 합산합니다.
  - `code`는 fenced block(``` · ~~~)만 대상으로 합니다. 들여쓰기 4칸 code block은 code로 인식하지 않고 본문으로 취급합니다. 블록 시작 판정에는 들여쓰기를 쓰지 않으므로 4칸 들여쓰기 줄도 heading·thematic break·fence·table·html·blockquote로 인식합니다. front matter는 파일 첫 줄이 `---`일 때만 인식하며 `---` 또는 `...`로 끝냅니다.
  - 파일이 없거나 읽을 수 없거나 UTF-8이 아니거나 `--id`를 찾지 못하면 stderr와 0이 아닌 종료 코드로 실패합니다.
- 설정은 `default`, `[roots.<name>] path`와 선택 `description`을 사용합니다. 최상위 `root` 설정과 `set <directory>`는 실패하며 기존 설정 바이트를 보존합니다. 자동 변환은 하지 않습니다.
- 설정 파싱이나 갱신 결과의 root 검증에 실패하면 기존 파일을 보존합니다. 유효한 `set` 호출로 같은 root의 경로를 정정할 수 있습니다. 심볼릭 링크 설정 파일은 조회·갱신을 지원합니다. 갱신은 링크의 실제 대상 파일을 원자적으로 교체하고 링크 자체는 보존합니다. 끊어진 링크는 대상·링크를 만들거나 교체하지 않고 실패합니다.
- 설정 갱신은 임시 파일 완성 후 원자적으로 교체하며 주석·서식은 정규화됩니다. 새로 만드는 설정 파일 권한은 0600입니다.
- 실패는 stderr와 0이 아닌 종료 코드로 보고합니다. 도움말·조회·점검은 파일을 만들지 않습니다.
- `--help`와 `-h`는 사용법을 stdout에 출력하고 0으로 끝냅니다. 명령 뒤에 붙은 도움말 옵션도 같은 처리를 합니다.
- 셸 설정이나 PATH는 수정하지 않습니다. PATH 설정 방법은 `doctor`의 `FAIL` 메시지와 README가 안내합니다.

## 구조와 검증

- `toolkit/agent-wiki/scripts/tests/`는 CLI 검증을 소유합니다. `scripts/src/inspect.rs`는 Markdown 구조 파싱과 트리 출력을 소유하고 `scripts/src/main.rs`는 명령 해석과 stdout·stderr 연결을 담당합니다.
- 설치·재설치·제거·점검, 옵션 해석, 기본 설치 폴더와 `~` 별칭, 설정·조회, 경로의 공백·유니코드, 잘못된 입력과 기존 파일 보존을 임시 홈에서 검증합니다. 점검 검증은 PATH 입력을 인자로 받아 실제 프로세스 환경에 의존하지 않게 합니다.
- `inspect` 검증은 front matter, heading 중첩과 수준 건너뛰기, list 중첩과 task checkbox, table 열 판정, code fence와 blockquote 중첩, link·image·Obsidian wikilink, `--id`의 자리 번호와 slug 선택, 형제·중첩 중복 slug 모호 실패, 4칸 들여쓰기 블록 시작과 code 미인식, 없는 파일과 없는 id 실패를 임시 파일에서 확인합니다.
- `toolkit/agent-wiki/`에서 `cargo test --locked --manifest-path scripts/Cargo.toml`을 실행합니다.
- CLI 바이너리는 `toolkit/agent-wiki/`에서 `sh scripts/build.sh`로 빌드합니다. 생성 바이너리 `scripts/wdown`과 Cargo target은 Git에 포함하지 않으며 실행 파일의 권한은 유지합니다.
