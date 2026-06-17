## 사용자 스펙 의도

- `pro-` 시리즈는 `advance-codex`가 아니라 별도 플러그인이 소유하는 것이 더 적합한지 검토하고 싶다.
- 별도 플러그인 이름은 `judgment-kit`이 적합하다.
- `judgment-kit`은 product planning, engineering, design judgment를 역할 기반 판단 기준으로 묶어야 한다.
- `advance-codex`는 role judgment skill을 직접 제공하지 않고 Codex surface 제작/운영 도구 경계로 남아야 한다.

---

# Judgment Kit Dev 플러그인 스펙

## 플러그인 목적

`judgment-kit-dev`는 Codex가 제품 기획, 엔지니어링, 디자인 작업에서 역할 기반 판단 기준을 적용하게 하는 플러그인입니다.
핵심 책임은 product planning judgment, engineering judgment, product design judgment를 독립 skill로 제공하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - 제품, 서비스, 기능을 정의하고 좁히는 product planning judgment
  - 문제 해결 중심의 engineering judgment, root cause analysis, implementation discipline
  - UI, UX, interface content, color, tone/expression, branding, space/composition, surface/form, product quality 중심의 product design judgment
- 제외:
  - skill/plugin 제작 guidance
  - scenario testing workflow
  - commit finalization 절차
  - agent token optimization
  - 일반 제품 구현 workflow
  - 특정 도메인 기능 구현 가이드

## 처리하려는 작업 형태

- 제품, 서비스, 기능 요청에서 넓은 요청을 사용자 문제, 제품 방식 후보, 기능 영역, 디자인 시스템 브리프 같은 부가 기획 표면, 가치, 범위, 요구사항, 우선순위, acceptance criteria, handoff 계약으로 분해하는 작업
- 코드 작성과 버그 수정에서 문제 정의, 원인 분석, 작은 완전 수정, 검증과 리스크 보고 기준을 명시하는 작업
- 화면, user flow, interface content, color, tone/expression, branding, space/composition, surface/form, product quality에서 디자인 판단 기준을 명시하는 작업

## 대표 표면

- 대표 스펙: `judgment-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `judgment-kit-dev/specs/skills/<skill-name>/spec.md`
- 핵심 선택 기준: 지금 필요한 판단 역할이 product planning, engineering, design 중 무엇인가

## 내장 skill 체계

- `pro-planner`: 제품, 서비스, 기능 정의에서 넓은 요청을 사용자 문제, 제품 방식 후보, 기능 영역, 디자인 시스템 브리프 같은 부가 기획 표면, 가치, 범위, 요구사항, 우선순위, acceptance criteria, handoff 계약으로 분해한다.
  - spec: `judgment-kit-dev/specs/skills/pro-planner/spec.md`
- `pro-engineering`: 코드 작성과 문제 해결에서 엔지니어링 판단, 원인 분석, 구현 규율, 검증 기준을 제공한다.
  - spec: `judgment-kit-dev/specs/skills/pro-engineering/spec.md`
- `pro-designer`: 화면과 제품 UI 작업에서 UI, UX, Content, Color, Tone & Expression, Branding, Space & Composition, Surface & Form, Quality 축의 디자인 판단 기준을 제공한다.
  - spec: `judgment-kit-dev/specs/skills/pro-designer/spec.md`

## SDD 운영 원칙

- plugin spec은 bundle 목적, 경계, usage surface, skill composition만 소유한다.
- 각 skill의 목적, 처리 계약, 독립성 원칙은 별도 folder-based skill spec에 둔다.
- skill 책임이 바뀌면 해당 skill spec과 `plugin.md`를 같은 변경 단위로 갱신한다.
- skill 선택 기준이 바뀌면 `plugin.md`, README, manifest prompt를 함께 점검한다.

## 현재 구조 메모

- normative skill spec은 모두 `specs/skills/<skill-name>/` 아래에 둔다.
- 이 플러그인의 주요 리스크는 역할 기반 판단 기준을 일반 구현 workflow나 Codex 운영 도구로 넓히는 것이다.
