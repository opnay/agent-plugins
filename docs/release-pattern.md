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

## Release Checklist

1. 바뀐 plugin surface와 spec을 같은 변경 단위에서 점검합니다.
2. 변경은 먼저 `src/<plugin-name>-dev`에 적용합니다.
3. specs는 `src/<plugin-name>-dev/specs/`에서만 갱신합니다.
4. release surface를 생성하거나 갱신합니다.
5. 해당 release 플러그인의 `.codex-plugin/plugin.json` version을 올립니다.
6. `.agents/plugins/marketplace.json` 경로와 policy 필드를 검증합니다.
7. 공개 채널은 `main`에 merge/push합니다.
8. 사용자는 `codex plugin marketplace upgrade opnay-plugins`로 업데이트합니다.

## 주의점

- prod/dev를 같은 plugin name으로 동시에 설치하지 않습니다.
- dev 플러그인은 `src/` 아래에 커밋합니다.
- specs를 release root와 `src/` 양쪽에 동시에 유지하지 않습니다.
- tag는 release note나 rollback 기준으로는 사용할 수 있지만, 일반 upgrade 경로를 tag에 묶지 않습니다.
