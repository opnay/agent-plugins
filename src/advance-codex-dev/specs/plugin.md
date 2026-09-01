# Advance Codex Dev 플러그인 스펙

## 플러그인 목적

`advance-codex-dev`는 Codex 활용 방식을 더 명시적이고 재사용 가능하게 설계하는 플러그인입니다.
핵심 책임은 skill, plugin bundle, skill scenario testing, agent token optimization 같은 Codex 활용 산출물을 각각의 좁은 skill 표면으로 제공하는 것입니다.
`.agents/sessions/{YYYYMMDD}` session folder convention은 호출 가능한 skill이 아니라 문서 수준의 참고 규칙으로만 유지합니다.

## 플러그인 경계와 비목표

- 포함:
  - skill 설계와 개편을 위한 creator-oriented guidance
  - reusable instruction을 fresh executor와 고정 시나리오로 테스트하고 분석 보고하는 workflow
  - installable plugin boundary와 bundled skill coherence 설계
  - 응답, 진행·상태 문구, reasoning·decision wording, 저장 문서에 정확성, 의미, 검증, 승인, 안전을 보존하는 token-efficient style 적용
- 제외:
  - 일반 제품 구현 workflow
  - 특정 도메인 기능 구현 가이드
  - `advance-codex-dev` 목적과 무관한 generic utility accumulation

## 처리하려는 작업 형태

- 새 skill이나 plugin을 만들거나 기존 것을 재설계하는 작업
- reusable instruction을 clean-context scenario로 테스트하고 evidence 중심으로 분석하는 작업
- agent-authored language 전반에 token-efficient style을 적용하되 판단 논리, workflow, 검증 범위는 바꾸지 않는 작업

## 대표 표면

- 대표 스펙: `advance-codex-dev/specs/plugin.md`
- skill 상세 스펙 위치: `advance-codex-dev/specs/skills/*.md` 또는 `advance-codex-dev/specs/skills/<skill-name>/spec.md`
- 핵심 선택 기준: 지금 개선하려는 주된 reusable artifact가 무엇인가

## 내장 skill 체계

- `skill-creator`: canonical `skill-creator` 위에 bounded skill 설계, plugin-owned skill 규칙, passive skill description trigger metadata 규칙을 덧붙인다.
  - spec: `advance-codex-dev/specs/skills/skill-creator.md`
- `skill-scenario-testing`: reusable instruction을 fresh subagent와 고정 시나리오로 테스트하고 evidence 중심으로 분석 보고한다.
  - spec: `advance-codex-dev/specs/skills/skill-scenario-testing.md`
- `plugin-creator`: top-down plugin boundary와 manifest-aligned packaging 규칙을 강화한다.
  - spec: `advance-codex-dev/specs/skills/plugin-creator.md`
- `optimize-token`: 응답, 진행·상태 문구, reasoning·decision wording, 저장 문서에 token-efficient style과 제한된 symbol grammar를 적용하되 판단 논리, workflow, 검증, 승인, exact literal, 안전 계약을 보존한다.
  - spec: `advance-codex-dev/specs/skills/optimize-token/spec.md`

## SDD 운영 원칙

- plugin spec은 bundle 목적, 경계, usage surface, skill composition만 소유한다.
- 각 skill의 목적, 처리 계약, 독립성 원칙은 반드시 별도 `specs/skills/<skill-name>.md` 또는 folder-based `specs/skills/<skill-name>/spec.md`에 둔다.
- skill 책임이 바뀌면 해당 skill spec과 `plugin.md`를 같은 변경 단위로 갱신한다.
- skill 선택 기준이 바뀌면 `plugin.md`, manifest prompt, 관련 creator skill spec을 함께 점검한다.
- scenario testing workflow처럼 독립 관심사로 분리된 계약은 다시 sibling skill 안으로 흡수하지 않는다.

## 현재 구조 메모

- normative skill spec은 모두 `specs/skills/` 아래에 둔다.
- 이 플러그인의 주요 리스크는 일반 workflow guidance나 unrelated convenience feature로 범위가 흐려지는 것이다.
