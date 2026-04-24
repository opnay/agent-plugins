## 사용자 스펙 의도

- skill boundary와 비책임을 더 엄격하게 잠그고 싶다.
- plugin-owned skill이 hidden sibling context에 기대지 않게 만들고 싶다.
- tool policy와 domain workflow가 섞일 때 이를 분리할 기준이 필요하다.

---

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

- 대표 표면: `advance-codex-dev/skills/skill-creator/SKILL.md`
- 관련 상위 라우팅: `advance-codex-dev-guide`

## 핵심 처리 계약

- 각 skill은 목적, 소유 범위, 비목표를 명확히 가져야 한다.
- plugin 내부 skill이라도 cross-skill usage guidance는 `<plugin>-guide`에 둔다.
- main bottleneck이 tool policy면 domain skill 안에 넣지 않고 별도 artifact를 고려한다.
- behavioral guidance는 skill reader 관점에서 작성한다.

## 검토 질문

- 이 skill이 hidden sibling context 없이도 이해 가능한 경계를 갖고 있는가?
- cross-skill usage guidance가 개별 skill 안으로 잘못 들어오지 않았는가?
- tool policy를 domain skill 안에 묻어 두고 있지 않은가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 스킬 자체가 hidden sibling context 제거를 다루므로, skill boundary guardrail은 독립적으로 읽히게 강제하는 편이 맞다.

## 확장 원칙

- skill packaging 규칙이 바뀌면 `plugin-creator`와 해당 plugin guide spec도 함께 점검한다.
- 새 규칙은 reusable skill quality를 높이는 경우에만 추가한다.
