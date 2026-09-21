# OPNay Agent Plugins

OPNay가 직접 관리하는 Codex 플러그인 마켓플레이스 저장소입니다.
`.agents/plugins/marketplace.json`은 플러그인 목록의 단일 진실 공급원입니다.

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

## 로컬 개발

- 레이아웃·운영 규칙: [AGENTS.md](AGENTS.md)
- 스펙 작성·검증: [SDD.md](docs/SDD.md)
- 브랜치·릴리즈: [release-pattern.md](docs/release-pattern.md)

## 플러그인

### Advance Codex

`advance-codex`는 Codex 활용 체계를 더 깊게 관리하기 위한 플러그인입니다.
skill 작성, plugin 작성, skill scenario testing, session 관리, commit workflow, subagent 정의 같은 메타 작업을 다룹니다.

- 경로: `advance-codex/`
- 주요 실행 표면: `plugin-creator`, `skill-creator`, `skill-scenario-testing`, `agents-sessions`, `git-committer`, `tool-use-guide`, `subagent-gate`, `subagent-creator`

### Judgment Kit

`judgment-kit`은 리서치, 제품 기획, 품질 관리 판단 기준을 제공하는 플러그인입니다.

- 경로: `judgment-kit/`
- 주요 실행 표면: `pro-researcher`, `pro-planner`, `pro-quality-manager`

### Design Kit

`design-kit`은 디자인 기반, 웹·프로덕트·시각 콘텐츠 설계, 아트·크리에이티브 디렉션, 프로젝트별 디자인 규칙을 제공합니다.

- 플러그인: `design-kit/`
- 개발 문서: `src/design-kit-dev/`
- 주요 실행 표면: `design-base`, `web-designer`, `product-designer`, `visual-content-designer`, `art-director`, `creative-director`, `project-design-rules`
- marketplace: `.agents/plugins/marketplace.json`의 `design-kit` 항목이 `./design-kit`을 가리킵니다.

### Advance Subagent

`advance-subagent`는 서브에이전트를 활용한 근거 중심 조사와 독립 workstream 위임·검증·통합의 심화 실행 방법을 제공합니다.

- 경로: `advance-subagent/`
- 주요 실행 표면: `deep-research`, `orchestrate-workstreams`

### SW Kit

`sw-kit`는 요구 행동 계약, 기술 방향, production code 품질, 코드베이스 유지보수를 제공합니다.

- 경로: `sw-kit/`
- 주요 실행 표면: `spec`, `engineering`, `code`, `maintenance`

## 저장소 구조

```text
.
├── .agents/plugins/marketplace.json
├── advance-codex/
├── advance-subagent/
├── agent-wiki/
├── app-extensions/
├── sw-kit/
├── design-kit/
├── judgment-kit/
├── toolkit/
├── src/<plugin-name>-dev/  # 스펙·변경 기록·개발 README
└── docs/
```

실행 플러그인과 개발 문서는 다음 위치에서 관리합니다.

```text
<plugin-name>/
  .codex-plugin/plugin.json
  README.md
  skills/

src/<plugin-name>-dev/
  README.md
  specs/plugin.md
  specs/skills/
  changes/
```

## 개발 원칙

- 플러그인 변경은 spec-driven으로 다룹니다.
- plugin surface가 바뀌면 `src/<plugin-name>-dev/README.md`, `src/<plugin-name>-dev/specs/plugin.md`, 관련 skill spec, `plugin.json`, marketplace entry를 함께 점검합니다.
- 플러그인별 release version은 각 `.codex-plugin/plugin.json`의 `version`이 소유합니다.
- 새 skill을 추가할 때는 먼저 plugin boundary와 sibling skill 관계를 확인합니다.
- 하네스나 평가 설계는 결정론적인 fixture, 고정 시나리오, 명시적인 pass/fail 기준을 우선합니다.
