## 사용자 스펙 의도

- custom agent 한 개의 역할과 경계를 명확한 contract로 만들고 싶다.
- TOML field shape와 usage guidance를 함께 잠그고 싶다.
- runtime orchestration 전체가 아니라 agent definition 자체에 집중하고 싶다.

---

# subagent-creator 스킬 스펙

## 목적

`subagent-creator`는 `.codex/agents/*.toml` 기반 custom agent를 정의하거나 개편하기 위한 스킬입니다.
핵심은 agent role, scope, defaults, usage guidance를 하나의 bounded agent contract로 정리하는 것입니다.

## 경계

- 포함:
  - custom agent role 정의
  - TOML file shape와 required fields 설계
  - project scope와 personal scope 선택
  - agent 사용 가이드와 validation 기준
- 제외:
  - multi-agent orchestration strategy 전반
  - plugin packaging
  - generic workflow selection

## 처리하려는 작업 형태

- 새로운 custom agent를 만들거나 기존 agent를 정리하는 경우
- `.codex/agents/*.toml`의 name, description, developer_instructions를 설계해야 하는 경우
- agent role과 reuse boundary를 명확히 해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/subagent-creator/SKILL.md`
- 관련 상위 라우팅: `advance-codex-dev-guide`

## 핵심 처리 계약

- agent의 주 책임과 비책임을 먼저 고정한다.
- TOML 필수 필드와 선택적 inherited defaults를 명시한다.
- role description은 실제 호출자가 언제 이 agent를 써야 하는지까지 설명해야 한다.
- validation은 file shape와 behavioral coherence를 함께 본다.

## 검토 질문

- 이 agent의 주 책임과 비책임이 충분히 잠겨 있는가?
- TOML 필수 필드와 선택적 defaults가 구분돼 있는가?
- usage guidance가 “언제 이 agent를 써야 하는가”까지 설명하는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: custom agent 한 개의 contract는 hidden sibling artifact 없이도 읽혀야 하며, orchestration 문맥은 별도 artifact로 넘겨야 한다.

## 확장 원칙

- custom agent field contract가 바뀌면 template reference와 validation 기준을 함께 갱신한다.
- plugin 수준 orchestration이 필요해지면 별도 artifact로 분리한다.
