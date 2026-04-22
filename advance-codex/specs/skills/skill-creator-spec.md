# skill-creator 스킬 스펙

## 목적

`skill-creator`는 canonical system `skill-creator` 위에 더 엄격한 skill boundary 규칙을 더하는 확장 스킬입니다.
핵심은 각 skill을 independently usable한 bounded artifact로 설계하고, plugin-owned skill인 경우 plugin-level guidance와 섞이지 않게 하는 것입니다.

## 경계

- 포함:
  - skill boundary와 비책임 명시
  - hidden sibling context 제거
  - plugin 내부 skill packaging 규칙
  - runtime-specific tool policy의 분리 필요성 판단
- 제외:
  - canonical base scaffold 전체 대체
  - plugin-level routing 정의
  - domain workflow 자체의 구현

## 처리하려는 작업 형태

- 새 skill을 만들거나 existing skill boundary를 재설계하는 경우
- plugin-owned skill이 sibling context에 의존하고 있는 경우
- tool-use policy를 skill에서 분리해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex/skills/skill-creator/SKILL.md`
- 관련 상위 라우팅: `advance-codex-guide`

## 핵심 처리 계약

- 각 skill은 목적, 소유 범위, 비목표를 명확히 가져야 한다.
- plugin 내부 skill이라도 cross-skill usage guidance는 `<plugin>-guide`에 둔다.
- main bottleneck이 tool policy면 domain skill 안에 넣지 않고 별도 artifact를 고려한다.
- behavioral guidance는 skill reader 관점에서 작성한다.

## 독립성 원칙

- 이 스킬은 skill boundary 설계 extension만 소유하며 plugin packaging 전체를 소유하지 않는다.
- sibling skill 없이도 skill authoring guardrail을 이해할 수 있어야 한다.

## 확장 원칙

- skill packaging 규칙이 바뀌면 `plugin-creator`와 해당 plugin guide spec도 함께 점검한다.
- 새 규칙은 reusable skill quality를 높이는 경우에만 추가한다.

