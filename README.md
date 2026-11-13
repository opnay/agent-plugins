# OPNay Agent Plugins

직업·전문 분야·작업 초점에 맞는 에이전트 플러그인을 관리하는 로컬 마켓플레이스입니다. 공통 목록과 순서는 `.agents/plugins/marketplace.json`이 소유합니다.

## 설치와 사용

저장소 루트에서 로컬 마켓을 등록하고 필요한 plugin을 설치합니다.

```sh
codex plugin marketplace add .
codex plugin add engineer@opnay-plugins
```

Claude Code는 `claude plugin marketplace add .`, `claude plugin install engineer@opnay-plugins`를 사용합니다. 호출은 Codex의 `$engineer:spec`, Claude의 `/engineer:spec`처럼 plugin 이름을 포함합니다. 등록만으로 모든 plugin이 설치되지는 않습니다. 자세한 형식은 [마켓플레이스](docs/marketplace.md)를 확인합니다.

## 플러그인 구성

| 분류 | Plugin | Skills |
| --- | --- | --- |
| 직업 | [planner](agents/planner/README.md) | plan |
| 직업 | [engineer](agents/engineer/README.md) | spec, engineering, code, maintenance |
| 직업 | [designer](agents/designer/README.md) | design-base, product-designer, web-designer, visual-content-designer, art-director, creative-director, project-design-rules |
| 실행 역할·조정 | [orchestrator](agents/orchestrator/README.md) | orchestrate-workstreams |
| 공통 기능·조사 | [researcher](capabilities/researcher/README.md) | research |
| 공통 기능·검증 | [quality](capabilities/quality/README.md) | verify |
| 공통 기반 | [agent-wiki](toolkit/agent-wiki/README.md) | manager, reader, researcher, writer |
| 도구 | [toolkit](toolkit/toolkit/README.md) | git, apple-container |
| 외부 앱 확장 | [app-extensions](toolkit/app-extensions/README.md) | figma-ext |
| Codex 제작 | [codex-creator](codex/codex-creator/README.md) | plugin-creator, skill-creator |
| Codex 문체 | [codex-optimize](codex/codex-optimize/README.md) | optimize-token, asd-ste100 |

`agents`는 직업과 실행 조정을 포함한 독립 실행 역할을, `capabilities`는 역할에 조합할 조사·검증 기능을 제공합니다. 직업 plugin은 공통 작업과 분야 기준을 함께 소유합니다. 숙련도는 직업 plugin의 설치된 `assets/proficiency.md`에서 선택해 역할 지침에 적용합니다. 단순 설치가 수준을 자동 적용하거나 전문성·권한을 보장하지 않습니다. [에이전트 구성](docs/agent-composition.md)에서 조합·실행 환경·이전 호출 migration을 확인합니다.

역할별 Codex 홈에는 필요한 plugin과 지침을 준비합니다. orcat 외부 역할 실행을 서브에이전트로 취급하되 이 저장소가 호출 adapter·자동 셋업·앱 native 내역 연결을 제공하지는 않습니다.

## 저장소 구조와 개발

```text
.
├── .agents/plugins/marketplace.json
├── .claude-plugin/marketplace.json
├── agents/       # planner, engineer, designer, orchestrator
├── capabilities/ # researcher, quality
├── toolkit/      # toolkit, app-extensions, agent-wiki
├── codex/        # codex-creator, codex-optimize
└── docs/
```

각 분류에는 실제 plugin과 `_specs/<plugin>/` 작성 계약이 함께 있습니다. 실제 plugin은 두 호스트 manifest·README·skills와 필요한 assets/scripts를 소유합니다. `_specs`에는 plugin.md, README, skills/<skill>/{spec.md,intent.md}, changes가 있으며 runtime에 포함하지 않습니다. 원본 복사본이나 별도 복사 빌드를 만들지 않습니다.

[AGENTS.md](AGENTS.md)는 운영 규칙, [SDD](docs/SDD.md)는 작성·검증, [브랜치·릴리즈](docs/release-pattern.md)는 배포 경계를 소유합니다. 사용 표면 변경은 관련 spec·README·두 manifest·marketplace를 함께 검증합니다. 과거 change/실험 기록은 당시 이름과 경로를 보존합니다.
