# OPNay Agent Plugins

OPNay가 직접 관리하는 Codex 플러그인 마켓플레이스 저장소입니다.
공개 설치용 release surface는 저장소 루트 바로 아래에 배치되며, 개발 원본은 `src/` 아래에서 관리합니다.
`.agents/plugins/marketplace.json`은 공개 설치 가능한 release 플러그인 목록의 단일 진실 공급원입니다.

## 마켓플레이스 등록

공개 설치용 마켓플레이스는 GitHub source를 사용합니다.

```sh
codex plugin marketplace add opnay/agent-plugins
```

Codex에서 `/plugins`를 열고 필요한 플러그인을 설치합니다.

마켓플레이스를 최신 상태로 갱신하려면 다음 명령을 사용합니다.

```sh
codex plugin marketplace upgrade
```

현재 마켓플레이스 표시명은 `OPNay Plugins`이고, 내부 id는 `opnay-plugins`입니다.

## 로컬 개발 원본

로컬 개발 원본은 `src/` 아래에서 관리하고 plugin name에 `-dev` suffix를 붙입니다.
이 원본은 직접 설치하거나 marketplace에 노출할 필요가 없습니다.
일반 개발 변경은 `src/<plugin-name>-dev`에 먼저 적용하고, 루트의 공개 release surface는 build command 산출물로 갱신합니다.

## 브랜치 모델

- `main`: 공개 release 브랜치입니다.
- `next`: 개발 브랜치입니다.
- 플러그인 수정은 기본적으로 `next`의 `src/<plugin-name>-dev`에서 진행합니다.
- `main`에는 `next`의 개발 내용을 release로 승격할 때만 반영합니다.
- 마지막 `main` merge 이후 `next`에서 플러그인을 처음 수정할 때, patch/minor/major 또는 target version을 사용자 확인으로 결정합니다.
- 같은 플러그인의 이후 변경은 추가 version bump 없이 build만 수행합니다.
- `src/<plugin-name>-dev`의 직접 설치 가능성은 개발 완료 조건이 아닙니다.
- 루트 `<plugin-name>/` release surface는 매 plugin 변경 뒤 build command로 갱신합니다.

`-dev` suffix는 설치 호출 표면이 아니라 개발 원본과 release 산출물을 구분하는 식별자입니다.

자세한 릴리즈/개발 분리 규칙은 `docs/release-pattern.md`를 봅니다.

## 플러그인

### Advance Codex

`advance-codex`는 Codex 활용 체계를 더 깊게 관리하기 위한 플러그인입니다.
skill 작성, plugin 작성, skill scenario testing, session 관리, commit workflow, subagent 정의 같은 메타 작업을 다룹니다.

- 경로: `advance-codex/`
- 주요 실행 표면: `plugin-creator`, `skill-creator`, `skill-scenario-testing`, `agents-sessions`, `git-committer`, `tool-use-guide`, `subagent-gate`, `subagent-creator`

### Judgment Kit

`judgment-kit`은 리서치, 기획, 엔지니어링, lean code, 디자인, 품질 관리 판단 기준을 제공하는 플러그인입니다.

- 경로: `judgment-kit/`
- 주요 실행 표면: `pro-researcher`, `pro-planner`, `pro-engineering`, `pro-code-keeper`, `pro-designer`, `pro-quality-manager`

### Advance Subagent

`advance-subagent`는 서브에이전트를 활용한 근거 중심 조사와 독립 workstream 위임·검증·통합의 심화 실행 방법을 제공합니다.

- 경로: `advance-subagent/`
- 주요 실행 표면: `deep-research`, `orchestrate-workstreams`

### Code Quality

`code-quality`는 production code 변경과 리뷰에서 correctness, maintainability, testability, robustness 기준을 적용합니다.

- 경로: `code-quality/`
- 주요 실행 표면: `code-quality`

## 저장소 구조

```text
.
├── .agents/plugins/marketplace.json
├── advance-codex/
├── advance-subagent/
├── code-quality/
├── judgment-kit/
├── src/
└── docs/
```

개발 원본 플러그인은 `src/` 아래에서 최소한 다음 구조를 유지합니다.

```text
src/<plugin-name>-dev/
  .codex-plugin/plugin.json
  README.md
  specs/plugin.md
  specs/skills/
  skills/
```

## 개발 원칙

- 플러그인은 루트 바로 아래에 둡니다. `./plugins/<plugin-name>` 경로는 사용하지 않습니다.
- 개발 원본은 `src/<plugin-name>-dev`에 둡니다.
- specs는 `src/` 안에서만 관리합니다.
- 일반 개발 변경은 `src/<plugin-name>-dev`에 먼저 적용합니다.
- 루트 release surface는 build command 산출물로만 갱신합니다.
- 플러그인 변경은 spec-driven으로 다룹니다.
- plugin surface가 바뀌면 `src/<plugin-name>-dev/README.md`, `src/<plugin-name>-dev/specs/plugin.md`, 관련 skill spec, `plugin.json`, marketplace entry를 함께 점검합니다.
- 플러그인별 release version은 각 `.codex-plugin/plugin.json`의 `version`이 소유합니다.
- 새 skill을 추가할 때는 먼저 plugin boundary와 sibling skill 관계를 확인합니다.
- 하네스나 평가 설계는 결정론적인 fixture, 고정 시나리오, 명시적인 pass/fail 기준을 우선합니다.
