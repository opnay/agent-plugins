## 사용자 스펙 의도

- skill boundary와 비책임을 더 엄격하게 잠그고 싶다.
- plugin-owned skill이 hidden sibling context에 기대지 않게 만들고 싶다.
- tool policy와 domain workflow가 섞일 때 이를 분리할 기준이 필요하다.
- 스킬의 `description`은 스킬 사용 여부를 결정하는 트리거 조건이므로, passive skill은 설명 끝에 토큰 매칭에 걸릴 수 있는 plain token 목록을 쉼표로 나열해야 한다.
  - 이 목록은 literal hashtag가 아니며 `#`를 붙이지 않는다.
  - 이 방식은 사용자가 명시적으로 스킬을 요청하지 않아도 적용되어야 하는 passive skill에 사용한다.

---

# skill-creator 스킬 스펙

## 목적

`skill-creator`는 canonical system `skill-creator` 위에 더 엄격한 skill boundary 규칙을 더하는 확장 스킬입니다.
핵심은 각 skill을 independently usable한 bounded artifact로 설계하고, plugin-owned skill인 경우 plugin usage guidance와 섞이지 않게 하는 것입니다.

## 경계

- 포함:
  - skill boundary와 비책임 명시
  - hidden sibling context 제거
  - plugin 내부 skill packaging 규칙
  - runtime-specific tool policy의 분리 필요성 판단
  - skill `description`의 trigger metadata 작성 규칙
- 제외:
  - canonical base scaffold 전체 대체
  - plugin 사용 기준 정의
  - domain workflow 자체의 구현

## 처리하려는 작업 형태

- 새 skill을 만들거나 existing skill boundary를 재설계하는 경우
- plugin-owned skill이 sibling context에 의존하고 있는 경우
- tool-use policy를 skill에서 분리해야 하는 경우
- passive skill이 명시적 호출 없이도 필요한 상황에서 선택되도록 `description` trigger token을 설계해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/skill-creator/SKILL.md`
- 호출 방식: 직접 호출하거나 manifest prompt의 안내를 따른다.

## 핵심 처리 계약

- 각 skill은 목적, 소유 범위, 비목표를 명확히 가져야 한다.
- plugin 내부 skill이라도 cross-skill usage guidance는 manifest prompt, README, plugin spec에 둔다.
- main bottleneck이 tool policy면 domain skill 안에 넣지 않고 별도 artifact를 고려한다.
- behavioral guidance는 skill reader 관점에서 작성한다.
- skill frontmatter `description`은 본문을 읽기 전에 사용되는 trigger metadata로 취급한다.
- `description`은 먼저 사람이 읽을 수 있는 기본 설명과 사용 조건을 짧게 쓰고, passive skill인 경우 마지막에 token matching을 돕는 plain token 목록을 쉼표로 나열한다.
- passive skill token 목록은 `#`가 붙은 hashtag가 아니라 `skill creation, skill update, passive trigger, description metadata` 같은 쉼표 구분 plain token 형식이어야 한다.
- 이 token 목록은 사용자가 명시적으로 skill을 호출하지 않아도 적용되어야 하는 passive skill에만 사용한다.
- 사용자가 명시적으로 호출하는 active skill이나 이미 충분히 명확한 narrow trigger skill에는 불필요한 token 목록을 붙이지 않는다.
- token 목록은 skill이 실제로 적용되어야 하는 입력 표현, artifact 이름, workflow 이름, 흔한 동의어를 포함하되 unrelated broad token으로 과도하게 넓히지 않는다.

## 검토 질문

- 이 skill이 hidden sibling context 없이도 이해 가능한 경계를 갖고 있는가?
- cross-skill usage guidance가 개별 skill 안으로 잘못 들어오지 않았는가?
- tool policy를 domain skill 안에 묻어 두고 있지 않은가?
- `description`이 body를 읽기 전 trigger metadata로 충분히 작동하는가?
- passive skill이라면 `description` 끝에 `#` 없는 쉼표 구분 plain token 목록이 있는가?
- token 목록이 실제 적용 조건을 넓히되, 무관한 broad match를 만들지 않는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 스킬 자체가 hidden sibling context 제거를 다루므로, skill boundary guardrail은 독립적으로 읽히게 강제하는 편이 맞다.

## 확장 원칙

- skill packaging 규칙이 바뀌면 `plugin-creator`, plugin spec, manifest prompt도 함께 점검한다.
- 새 규칙은 reusable skill quality를 높이는 경우에만 추가한다.
