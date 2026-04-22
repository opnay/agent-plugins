# screen-spec 스킬 스펙

## 목적

`screen-spec`은 aligned design direction을 pre-code screen specification으로 전환하는 스킬입니다.
핵심 산출물은 screen section, content intent, interaction expectation, hierarchy, scope boundary가 포함된 화면 계약입니다.

## 경계

- 포함:
  - screen section 구조 정의
  - content intent와 interaction expectation 명시
  - hierarchy와 scope boundary 정리
- 제외:
  - design direction 자체를 새로 정의하는 일
  - implementation detail
  - Figma canvas execution

## 처리하려는 작업 형태

- aligned brief를 실제 화면 수준으로 내려야 하는 경우
- code나 Figma 작업 전에 section-level contract가 필요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `designer-kit/skills/screen-spec/SKILL.md`
- 관련 상위 라우팅: `designer-kit-guide`

## 핵심 처리 계약

- spec은 section, intent, interaction, hierarchy를 빠짐없이 포함해야 한다.
- pre-code artifact이므로 component API나 engineering structure를 끌어오지 않는다.
- design direction과 화면 계약 사이의 누락된 가정을 드러내야 한다.

## 독립성 원칙

- 이 스킬은 design brief나 critique를 대체하지 않는다.
- screen spec만 읽어도 구현 전에 필요한 화면 의도가 이해 가능해야 한다.

## 확장 원칙

- interaction depth가 늘어나도 engineering detail은 별도 플러그인으로 넘긴다.
- screen spec format이 바뀌면 guide와 plugin spec을 함께 점검한다.

