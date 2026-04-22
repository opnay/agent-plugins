# Designer Kit 플러그인 스펙

## 플러그인 목적

`designer-kit`은 구현 이전 단계의 design-only workflow를 다루는 플러그인입니다.
핵심 책임은 지금 필요한 design work가 방향 정의인지, 기존 방향 critique인지, pre-code screen specification인지 분류하고, 코드나 Figma 실행으로 넘어가기 전에 design artifact를 정교하게 만드는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - design brief 작성
  - UI/UX critique
  - pre-code screen specification
  - 위 세 design stage 사이의 routing
- 제외:
  - frontend engineering structure
  - code implementation guidance
  - Figma execution detail

## 처리하려는 작업 형태

- vague product/interface idea를 design-ready brief로 정리하는 작업
- 이미 나온 concept, screen, flow를 critique하는 작업
- aligned direction을 pre-code screen contract로 정리하는 작업

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `designer-kit-guide`
- 대표 스펙: `designer-kit/specs/plugin-spec.md`
- skill 상세 스펙 위치: `designer-kit/specs/skills/*.md`

## 내장 skill 체계

- `designer-kit-guide`: design task를 적절한 design stage로 라우팅한다.
  - spec: `designer-kit/specs/skills/designer-kit-guide-spec.md`
- `design-brief`: vague idea를 aligned design direction으로 바꾼다.
  - spec: `designer-kit/specs/skills/design-brief-spec.md`
- `ui-critique`: existing direction을 critique하고 핵심 문제를 우선순위화한다.
  - spec: `designer-kit/specs/skills/ui-critique-spec.md`
- `screen-spec`: aligned direction을 pre-code screen specification으로 전환한다.
  - spec: `designer-kit/specs/skills/screen-spec-spec.md`

## SDD 운영 원칙

- plugin spec은 design bundle의 목적, stage model, routing surface만 소유한다.
- 각 skill의 처리 계약은 `specs/skills/` 아래 독립 문서로 분리한다.
- design stage의 집합이나 handoff 순서가 바뀌면 `designer-kit-guide`와 `plugin-spec.md`를 함께 갱신한다.
- implementation이나 Figma execution guidance는 이 플러그인 안으로 끌어오지 않는다.

## 현재 구조 메모

- 이 플러그인의 주요 리스크는 critique나 screen-spec이 code-facing guidance를 흡수하면서 design-only 경계가 흐려지는 것이다.
