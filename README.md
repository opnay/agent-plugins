# OPNay Agent Plugins

OPNay가 직접 관리하는 에이전트 플러그인 마켓플레이스 저장소입니다.
`.agents/plugins/marketplace.json`은 플러그인 목록·순서의 단일 진실 공급원입니다.

## 마켓플레이스 등록

마켓 이름은 `opnay-plugins`입니다. 저장소 루트에서 로컬 마켓을 등록하고 필요한 플러그인을 선택합니다.

| CLI | 마켓 등록 | 플러그인 설치·선택 |
| --- | --- | --- |
| Codex | `codex plugin marketplace add .` | `/plugins`에서 선택 |
| Claude Code | `claude plugin marketplace add .` | `claude plugin install sw-kit@opnay-plugins` |

호출 표기는 Codex의 `$sw-kit:spec`, Claude Code의 `/sw-kit:spec`처럼 사용합니다.
GitHub 등록·갱신과 매니페스트 형식은 [마켓플레이스](docs/marketplace.md)를 확인해 주세요.

## 로컬 개발

- 레이아웃·운영 규칙: [AGENTS.md](AGENTS.md)
- 스펙 작성·검증: [SDD.md](docs/SDD.md)
- 브랜치·릴리즈: [release-pattern.md](docs/release-pattern.md)

## 플러그인

### Codex Optimize

`codex-optimize`는 토큰 효율적인 문체와 기술 지시의 명확성을 개선하는 플러그인입니다.

- 경로: `codex/codex-optimize/`
- 주요 실행 표면: `optimize-token`, `asd-ste100`

### Codex Creator

`codex-creator`는 canonical system creator를 확장하여 plugin·skill 제작을 다룹니다.

- 경로: `codex/codex-creator/`
- 주요 실행 표면: `plugin-creator`, `skill-creator`
- 호출: `$codex-creator:plugin-creator`, `$codex-creator:skill-creator`

### Judgment Kit

`judgment-kit`은 리서치, 제품 기획, 품질 관리 판단 기준을 제공하는 플러그인입니다.

- 경로: `agent/judgment-kit/`
- 주요 실행 표면: `pro-researcher`, `pro-planner`, `pro-quality-manager`

### Design Kit

`design-kit`은 디자인 기반, 웹·프로덕트·시각 콘텐츠 설계, 아트·크리에이티브 디렉션, 프로젝트별 디자인 규칙을 제공합니다.

- 플러그인: `designer/design-kit/`
- 개발 문서: `designer/_specs/design-kit/`
- 주요 실행 표면: `design-base`, `web-designer`, `product-designer`, `visual-content-designer`, `art-director`, `creative-director`, `project-design-rules`
- marketplace: `.agents/plugins/marketplace.json`의 `design-kit` 항목이 `./designer/design-kit`을 가리킵니다.

### Advance Subagent

`advance-subagent`는 분기 설계부터 검증·완료·새 맥락 인계까지의 공통 fork-and-flush 루프와 근거 중심 조사 계약을 제공합니다.

- 경로: `agent/advance-subagent/`
- 주요 실행 표면: `deep-research`, `orchestrate-workstreams`

### Quality

`quality`는 모든 분야의 산출물에서 특성에 맞는 검증 기준과 절차를 도출하고, 실제 근거로 판정하는 공통 행동 패턴입니다.

- 플러그인: [agent/quality](agent/quality/README.md)
- 개발 문서: [agent/_specs/quality](agent/_specs/quality/README.md)
- 주요 실행 표면: `$quality:verify`

### SW Kit

`sw-kit`는 요구 행동 계약, 기술 방향, production code 품질, 코드베이스 유지보수를 제공합니다.

- 경로: `engineer/sw-kit/`
- 주요 실행 표면: `spec`, `engineering`, `code`, `maintenance`

## 저장소 구조

```text
.
├── .agents/plugins/marketplace.json
├── .claude-plugin/marketplace.json
├── agent/       # advance-subagent, agent-wiki, judgment-kit, quality
├── codex/       # codex-creator, codex-optimize
├── toolkit/     # toolkit, app-extensions
├── engineer/    # sw-kit
├── designer/    # design-kit
└── docs/
```

각 분류는 실제 플러그인과 개발 문서를 다음처럼 관리합니다. 분류는 저장 위치이며 설치 이름·skill 호출 식별자는 바뀌지 않습니다.

```text
<group>/
  <plugin-name>/
    .codex-plugin/plugin.json
    .claude-plugin/plugin.json
    README.md
    skills/
  _specs/<plugin-name>/
    README.md
    plugin.md
    skills/<skill-name>/
      intent.md
      spec.md
    changes/
```

`_specs`는 작성·검증 문서이며 설치 대상이 아닙니다. 플러그인 복사본이나 별도 복사 빌드를 두지 않습니다.

## 개발 원칙

- 플러그인 변경은 spec-driven으로 다룹니다.
- plugin surface가 바뀌면 `<group>/_specs/<plugin-name>/README.md`, `<group>/_specs/<plugin-name>/plugin.md`, 관련 skill spec, `plugin.json`, marketplace entry를 함께 점검합니다.
- 플러그인별 release version은 각 `.codex-plugin/plugin.json`의 `version`이 소유합니다.
- 매니페스트 간 `version`을 일치시키고 목록 이름·순서·경로도 공통 목록과 맞춥니다.
- 새 skill을 추가할 때는 먼저 plugin boundary와 sibling skill 관계를 확인합니다.
- 하네스나 평가 설계는 결정론적인 fixture, 고정 시나리오, 명시적인 pass/fail 기준을 우선합니다.
