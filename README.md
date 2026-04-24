# OPNay Agent Plugins

OPNay가 직접 관리하는 Codex 플러그인 마켓플레이스 저장소입니다.
플러그인들은 저장소 루트 바로 아래에 배치되며, `.agents/plugins/marketplace.json`이 설치 가능한 플러그인 목록의 단일 진실 공급원입니다.

## 준비

`loop-kit`처럼 사용자 선택지를 구조화된 질문 도구로 열어야 하는 플러그인을 제대로 쓰려면 Codex의 개발 중인 기능인 `default_mode_request_user_input`를 활성화해야 합니다.

```sh
codex features enable default_mode_request_user_input
```

## 마켓플레이스 등록

이 저장소를 Codex 플러그인 마켓플레이스 source로 추가합니다.

```sh
codex plugin marketplace add opnay/agent-plugins
```

Codex에서 `/plugins`를 열고 필요한 플러그인을 설치합니다.

마켓플레이스를 최신 상태로 갱신하려면 다음 명령을 사용합니다.

```sh
codex plugin marketplace upgrade
```

현재 마켓플레이스 표시명은 `OPNay Plugins`이고, 내부 id는 `opnay-plugins`입니다.

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
- 평가 설계: `rpg-kit/specs/subagent-role-packet-evaluation.md`

## 저장소 구조

```text
.
├── .agents/plugins/marketplace.json
├── advance-codex/
├── loop-kit/
├── rpg-kit/
└── workflow-kit/
```

각 플러그인은 최소한 다음 구조를 유지합니다.

```text
<plugin-name>/
  .codex-plugin/plugin.json
  README.md
  specs/plugin.md
  specs/skills/
  skills/
```

## 개발 원칙

- 플러그인은 루트 바로 아래에 둡니다. `./plugins/<plugin-name>` 경로는 사용하지 않습니다.
- 플러그인 변경은 spec-driven으로 다룹니다.
- plugin surface가 바뀌면 `README.md`, `specs/plugin.md`, 관련 skill spec, 관련 guide skill, `plugin.json`, marketplace entry를 함께 점검합니다.
- 새 skill을 추가할 때는 먼저 plugin boundary와 sibling skill 관계를 확인합니다.
- 하네스나 평가 설계는 결정론적인 fixture, 고정 시나리오, 명시적인 pass/fail 기준을 우선합니다.
