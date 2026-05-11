# Advance Codex Dev 플러그인 스펙

## 플러그인 목적

`advance-codex-dev`는 Codex 활용 방식을 더 명시적이고 재사용 가능하게 설계하는 플러그인입니다.
핵심 책임은 skill, reusable tool policy, plugin bundle, custom agent, skill scenario testing, session folder convention, change finalization 같은 Codex 활용 산출물을 각각의 좁은 skill 표면으로 제공하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - skill 설계와 개편을 위한 creator-oriented guidance
  - reusable instruction을 fresh executor와 고정 시나리오로 테스트하고 분석 보고하는 workflow
  - tool selection, sequencing, ask-vs-infer, escalation policy 분리
  - installable plugin boundary와 bundled skill coherence 설계
  - custom agent 정의와 usage guidance
  - subagent runtime handoff의 종료 시점, 최소 맥락, 위임 경계 설계
  - reviewable work unit 동안 worker subagent를 운영하고 작업 단위 종료 시 close/dispose하는 생애주기 설계
  - `.agents/sessions` 폴더의 기본 용도와 `turn-gate` 정렬 plan/flow record 골격 정의
  - task-scoped commit finalization discipline
- 제외:
  - 일반 제품 구현 workflow
  - 특정 도메인 기능 구현 가이드
  - `advance-codex-dev` 목적과 무관한 generic utility accumulation

## 처리하려는 작업 형태

- 새 skill, plugin, custom agent를 만들거나 기존 것을 재설계하는 작업
- subagent를 호출하기 전 종료 조건과 context packet을 gate로 잠그는 작업
- reviewable work unit 동안 worker subagent를 spawn, operate, verify, close/dispose하는 작업
- reusable instruction을 clean-context scenario로 테스트하고 evidence 중심으로 분석하는 작업
- domain workflow와 분리된 tool-use policy를 설계하는 작업
- `.agents/sessions` 폴더 경계나 commit workflow처럼 Codex 사용 자체의 운영 품질을 안정화하는 작업

## 대표 표면

- 대표 스펙: `advance-codex-dev/specs/plugin.md`
- skill 상세 스펙 위치: `advance-codex-dev/specs/skills/*.md`
- 핵심 선택 기준: 지금 개선하려는 주된 reusable artifact가 무엇인가

## 내장 skill 체계

- `skill-creator`: canonical `skill-creator` 위에 bounded skill 설계와 plugin-owned skill 규칙을 덧붙인다.
  - spec: `advance-codex-dev/specs/skills/skill-creator.md`
- `skill-scenario-testing`: reusable instruction을 fresh subagent와 고정 시나리오로 테스트하고 evidence 중심으로 분석 보고한다.
  - spec: `advance-codex-dev/specs/skills/skill-scenario-testing.md`
- `tool-use-guide`: domain artifact에서 분리되어야 하는 reusable tool policy를 설계한다.
  - spec: `advance-codex-dev/specs/skills/tool-use-guide.md`
- `plugin-creator`: top-down plugin boundary와 manifest-aligned packaging 규칙을 강화한다.
  - spec: `advance-codex-dev/specs/skills/plugin-creator.md`
- `subagent-creator`: `.codex/agents/*.toml`과 custom agent usage guidance를 정의한다.
  - spec: `advance-codex-dev/specs/skills/subagent-creator.md`
- `subagent-gate`: subagent 호출 전 종료 시점, 최소 context packet, 위임 경계, 결과 계약을 잠근다.
  - spec: `advance-codex-dev/specs/skills/subagent-gate.md`
- `subagent-work`: reviewable work unit 동안 worker subagent를 생성, 운영, 검증, 종료하는 엄격한 lifecycle을 제공한다.
  - spec: `advance-codex-dev/specs/skills/subagent-work.md`
- `agents-sessions`: `.agents/sessions` 폴더의 기본 용도와 `turn-gate` 정렬 plan/flow record 골격을 정의한다.
  - spec: `advance-codex-dev/specs/skills/agents-sessions.md`
- `git-committer`: 검증 가능한 task-scoped commit finalization 규율을 제공한다.
  - spec: `advance-codex-dev/specs/skills/git-committer.md`

## SDD 운영 원칙

- plugin spec은 bundle 목적, 경계, usage surface, skill composition만 소유한다.
- 각 skill의 목적, 처리 계약, 독립성 원칙은 반드시 별도 `specs/skills/<skill-name>.md`에 둔다.
- skill 책임이 바뀌면 해당 skill spec과 `plugin.md`를 같은 변경 단위로 갱신한다.
- skill 선택 기준이 바뀌면 `plugin.md`, manifest prompt, 관련 creator skill spec을 함께 점검한다.
- scenario testing workflow나 tool-use policy처럼 독립 관심사로 분리된 계약은 다시 sibling skill 안으로 흡수하지 않는다.

## 현재 구조 메모

- normative skill spec은 모두 `specs/skills/` 아래에 둔다.
- 이 플러그인의 주요 리스크는 일반 workflow guidance나 unrelated convenience feature로 범위가 흐려지는 것이다.
