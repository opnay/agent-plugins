# OPNay Agent Plugins

OPNay가 직접 관리하는 Codex 플러그인 마켓플레이스 저장소입니다.
공개 설치용 release surface는 저장소 루트 바로 아래에 배치되며, 개발 원본은 `src/` 아래에서 관리합니다.
`.agents/plugins/marketplace.json`은 공개 설치 가능한 release 플러그인 목록의 단일 진실 공급원입니다.

## 준비

`loop-kit`처럼 사용자 선택지를 구조화된 질문 도구로 열어야 하는 플러그인을 제대로 쓰려면 Codex의 개발 중인 기능인 `default_mode_request_user_input`를 활성화해야 합니다.

```sh
codex features enable default_mode_request_user_input
```

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

## 로컬 개발 마켓플레이스

공개 설치용 플러그인과 로컬 개발용 플러그인을 동시에 쓰려면 plugin name 충돌을 피해야 합니다.
로컬 개발용 플러그인은 `src/` 아래에서 관리하고, plugin name에 `-dev` suffix를 붙입니다.
일반 개발 변경은 `src/<plugin-name>-dev`에 먼저 적용하고, 루트의 공개 release surface는 build command 산출물로 갱신합니다.

## 브랜치 모델

- `main`: 공개 release 브랜치입니다.
- `next`: 개발 브랜치입니다.
- 플러그인 수정은 기본적으로 `next`의 `src/<plugin-name>-dev`에서 진행합니다.
- `main`에는 `next`의 개발 내용을 release로 승격할 때만 반영합니다.
- 마지막 `main` merge 이후 `next`에서 플러그인을 처음 수정할 때, patch/minor/major 또는 target version을 사용자 확인으로 결정합니다.
- 같은 플러그인의 이후 변경은 추가 version bump 없이 build만 수행합니다.
- 개발 버전을 쓰는 사용자는 `next` 브랜치를 marketplace source로 등록하고 `<plugin-name>-dev`를 설치합니다.
- 루트 `<plugin-name>/` release surface는 매 plugin 변경 뒤 build command로 갱신합니다.

예를 들어 공개 플러그인은 `$rpg-kit:subagent-role`, 개발 플러그인은 `$rpg-kit-dev:subagent-role`로 분리됩니다.

자세한 릴리즈/개발 분리 규칙은 `docs/release-pattern.md`를 봅니다.

## 플러그인

### Workflow Kit

`workflow-kit`은 planning, deep interview, review loop, commit readiness, turn gate 같은 일반 workflow taxonomy와 canonical workflow contract를 제공합니다.

- 경로: `workflow-kit/`
- 대표 엔트리포인트: `workflow-kit-guide`

### Advance Codex

`advance-codex`는 Codex 활용 체계를 더 깊게 관리하기 위한 플러그인입니다.
skill 작성, plugin 작성, empirical prompt tuning, session 관리, commit workflow, subagent 정의 같은 메타 작업을 다룹니다.

- 경로: `advance-codex/`
- 대표 엔트리포인트: `advance-codex-guide`

### Loop Kit

`loop-kit`은 하나의 작업 턴을 사용자가 명시적으로 종료할 때까지 유지하는 turn-gated loop 플러그인입니다.
`turn-gate`가 current-phase mode와 question-routing mode를 선택하고, 결과 보고 뒤 다음 플로우 선택지를 다시 엽니다.

- 경로: `loop-kit/`
- 대표 엔트리포인트: `loop-kit-guide`
- 주요 실행 표면: `turn-gate`
- 필수 준비: `default_mode_request_user_input`

### RPG Kit

`rpg-kit`은 역할이 할당된 subagent를 설계하고, spawn boundary와 answer contract를 명시하며, subagent 동작을 학습하기 위한 role orchestration 플러그인입니다.
역할은 단순 직함이 아니라 functional role, responsibility domain, background expertise, general expertise, decision style을 조합한 role specialty로 다룹니다.

- 경로: `rpg-kit/`
- 대표 엔트리포인트: `rpg-kit-guide`
- 주요 실행 표면: `subagent-role`
- 평가 설계: `src/rpg-kit-dev/specs/subagent-role-packet-evaluation.md`

## 저장소 구조

```text
.
├── .agents/plugins/marketplace.json
├── advance-codex/
├── loop-kit/
├── rpg-kit/
├── src/
└── workflow-kit/
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
- plugin surface가 바뀌면 `src/<plugin-name>-dev/README.md`, `src/<plugin-name>-dev/specs/plugin.md`, 관련 skill spec, 관련 guide skill, `plugin.json`, marketplace entry를 함께 점검합니다.
- 플러그인별 release version은 각 `.codex-plugin/plugin.json`의 `version`이 소유합니다.
- 새 skill을 추가할 때는 먼저 plugin boundary와 sibling skill 관계를 확인합니다.
- 하네스나 평가 설계는 결정론적인 fixture, 고정 시나리오, 명시적인 pass/fail 기준을 우선합니다.
