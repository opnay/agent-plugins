# Release Pattern

## 목적

이 저장소는 공개 설치용 플러그인과 로컬 개발용 플러그인을 분리해서 운영합니다.
마켓플레이스는 플러그인 목록을 제공하고, 실제 버전 단위는 각 플러그인의 `.codex-plugin/plugin.json`에 있는 `version`이 소유합니다.

## 기본 원칙

- `src/` 아래가 개발 원본입니다.
- specs는 `src/<plugin-name>-dev/specs/` 안에서만 관리합니다.
- 공개 설치용 플러그인은 release surface로 취급합니다.
- 공개 사용자는 `opnay-plugins` 마켓플레이스를 설치합니다.
- 로컬 개발자는 `src/`의 dev 플러그인을 기준으로 작업합니다.
- 공개 플러그인 이름과 로컬 개발 플러그인 이름은 충돌하지 않아야 합니다.
- dev 플러그인은 `<plugin-name>-dev` 이름을 사용합니다.
- release 플러그인은 `<plugin-name>` 이름을 사용합니다.
- `@tag` 또는 `@ref` 설치는 고정 ref 용도입니다. 일반 upgrade 채널로 사용하지 않습니다.

## 공개 설치 채널

공개 설치 채널은 GitHub source를 사용합니다.

```sh
codex plugin marketplace add opnay/agent-plugins
```

업데이트는 moving source를 갱신합니다.

```sh
codex plugin marketplace upgrade opnay-plugins
```

플러그인별 release는 해당 플러그인의 manifest version bump로 표현합니다.

```json
{
  "name": "rpg-kit",
  "version": "0.1.1"
}
```

Codex cache는 marketplace, plugin, version을 반영해 `opnay-plugins/<plugin>/<version>` 형태로 설치본을 유지합니다.

## 로컬 개발 채널

로컬 개발 채널은 `src/` 아래의 committed dev 플러그인을 사용합니다.
이 플러그인은 실제 개발 원본이며, specs도 이 안에만 둡니다.

예:

- `src/rpg-kit-dev`
- `src/loop-kit-dev`
- `src/workflow-kit-dev`
- `src/advance-codex-dev`

dev 플러그인은 공개 플러그인과 동시에 설치해도 호출 표면이 충돌하지 않아야 합니다.
따라서 dev plugin manifest의 `name`은 반드시 `<plugin-name>-dev`여야 합니다.

예:

- 공개: `$rpg-kit:subagent-role`
- 개발: `$rpg-kit-dev:subagent-role`

## Source / Release 구조

`src/`는 개발 원본입니다.

```text
src/
  rpg-kit-dev/
    .codex-plugin/plugin.json
    README.md
    specs/
    skills/
```

release surface는 사용자 설치용 산출물입니다.

```text
rpg-kit/
  .codex-plugin/plugin.json
  README.md
  skills/
```

release surface는 사용자 설치에 필요한 파일만 둡니다.
specs는 `src/` 안에서만 관리합니다.
release 과정에서 specs를 공개 산출물에 포함할지 여부는 별도 release 정책으로 명시적으로 결정합니다.

## Release 변환 규칙

dev 플러그인에서 release 플러그인을 만들 때는 다음 변환을 적용합니다.

- directory: `src/<plugin-name>-dev` -> `<plugin-name>`
- manifest `name`: `<plugin-name>-dev` -> `<plugin-name>`
- display name: `... Dev` -> `...`
- dev-only 문구와 실험 notes 제거
- release version은 release plugin manifest의 `version`으로 관리
- marketplace entry는 release plugin만 가리킴

## Release Generation Contract

release 생성은 `src/<plugin-name>-dev`를 입력으로 받고, 저장소 루트의 `<plugin-name>` release surface를 출력합니다.

### 입력

- source directory: `src/<plugin-name>-dev`
- required manifest: `src/<plugin-name>-dev/.codex-plugin/plugin.json`
- required runtime files:
  - `README.md`
  - `skills/`
- source-only files:
  - `specs/`
  - evaluation docs
  - planning docs

### 출력

- output directory: `<plugin-name>`
- required release files:
  - `.codex-plugin/plugin.json`
  - `README.md`
  - `skills/`
- optional runtime files:
  - `assets/`
  - `scripts/`
  - `.mcp.json`
  - `.app.json`

`specs/`는 기본적으로 release surface에 복사하지 않습니다.
specs는 source-of-truth인 `src/` 안에만 둡니다.
특정 spec을 공개 runtime artifact로 포함해야 한다면 release policy에서 파일 단위로 명시해야 합니다.

### 이름 변환

`<plugin-name>-dev`에서 `<plugin-name>`을 계산합니다.

- plugin directory: `src/rpg-kit-dev` -> `rpg-kit`
- manifest `name`: `rpg-kit-dev` -> `rpg-kit`
- display name: `RPG Kit Dev` -> `RPG Kit`
- skill identifier prefix:
  - `rpg-kit-dev-guide` -> `rpg-kit-guide`
  - `workflow-kit-dev-guide` -> `workflow-kit-guide`
- skill names that do not encode the plugin name are kept as-is:
  - `subagent-role` -> `subagent-role`
  - `turn-gate` -> `turn-gate`

### Text Rewrite

release 생성은 runtime 문서와 skill frontmatter에서 다음 token을 바꿉니다.

- `$<plugin-name>-dev:` -> `$<plugin-name>:`
- `<plugin-name>-dev` -> `<plugin-name>`
- `<Display Name> Dev` -> `<Display Name>`
- 남은 `" Dev"` suffix 제거
- `<plugin-name>-dev-guide` -> `<plugin-name>-guide`

광범위한 자연어 rewrite는 하지 않습니다.
release-only 문구가 필요하면 source 문서에 명시적인 release marker를 두고 변환 규칙으로 처리합니다.

### Version Rule

release version은 자동 추론하지 않습니다.
release 생성 명령은 다음 둘 중 하나를 요구해야 합니다.

- 명시적인 새 version
- 기존 release manifest version 유지 옵션

이유:

- marketplace upgrade는 moving source를 갱신합니다.
- 사용자가 설치하는 실제 plugin version은 release manifest의 `version`으로 구분됩니다.
- 여러 플러그인이 한 repo에 있으므로 repo tag 하나로 plugin별 release 상태를 대신할 수 없습니다.

### Safety Rules

- release 생성은 먼저 output directory를 검증합니다.
- output directory가 수동 수정된 release surface라면 덮어쓰기 전에 diff를 보여줘야 합니다.
- source `specs/`가 release root에 남아 있으면 release 검증은 실패해야 합니다.
- release manifest name과 marketplace entry name이 다르면 실패해야 합니다.
- release marketplace path는 항상 `./<plugin-name>`이어야 합니다.

## Release Checklist

1. 바뀐 plugin surface와 spec을 같은 변경 단위에서 점검합니다.
2. 변경은 먼저 `src/<plugin-name>-dev`에 적용합니다.
3. specs는 `src/<plugin-name>-dev/specs/`에서만 갱신합니다.
4. release generation contract에 따라 release surface를 생성하거나 갱신합니다.
5. 해당 release 플러그인의 `.codex-plugin/plugin.json` version을 명시적으로 올리거나 유지합니다.
6. release root에 `specs/`가 남지 않았는지 확인합니다.
7. `.agents/plugins/marketplace.json` 경로와 policy 필드를 검증합니다.
8. 공개 채널은 `main`에 merge/push합니다.
9. 사용자는 `codex plugin marketplace upgrade opnay-plugins`로 업데이트합니다.

## Release Generation Command

release surface 생성 명령은 다음 형태를 사용합니다.

```sh
node scripts/generate-release.mjs <plugin-name> --version <version>
```

기존 release version을 의도적으로 유지하려면 다음처럼 명시합니다.

```sh
node scripts/generate-release.mjs <plugin-name> --keep-version
```

기존 release directory를 덮어쓸 때는 `--force`를 붙여야 합니다.

```sh
node scripts/generate-release.mjs rpg-kit --version 0.1.1 --force
```

스크립트는 현재 다음을 강제합니다.

- `src/<plugin-name>-dev/.codex-plugin/plugin.json`이 있어야 한다.
- source manifest name은 `<plugin-name>-dev`여야 한다.
- release version은 `--version` 또는 `--keep-version`으로 명시해야 한다.
- 기존 release directory는 `--force` 없이는 덮어쓰지 않는다.
- release output에는 `specs/`가 없어야 한다.

## 주의점

- prod/dev를 같은 plugin name으로 동시에 설치하지 않습니다.
- dev 플러그인은 `src/` 아래에 커밋합니다.
- specs를 release root와 `src/` 양쪽에 동시에 유지하지 않습니다.
- tag는 release note나 rollback 기준으로는 사용할 수 있지만, 일반 upgrade 경로를 tag에 묶지 않습니다.
