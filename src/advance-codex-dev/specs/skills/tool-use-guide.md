## 사용자 스펙 의도

- tool 선택, 순서, escalation 기준을 reusable policy로 분리하고 싶다.
- domain workflow 설명과 runtime tool guidance를 섞지 않고 싶다.
- ask-vs-infer, local-vs-external 같은 판단 기준을 artifact 바깥으로 꺼내고 싶다.

---

# tool-use-guide 스킬 스펙

## 목적

`tool-use-guide`는 skill, plugin, custom agent가 도구를 어떻게 선택하고 순서화하며 언제 escalate해야 하는지에 대한 reusable policy를 설계하는 스킬입니다.
핵심은 domain intent와 runtime mechanics를 분리해 tool contract를 독립 surface로 유지하는 것입니다.

## 경계

- 포함:
  - tool selection policy
  - tool sequencing과 fallback 설계
  - ask-vs-infer, local-vs-external, escalation boundary 정의
  - runtime-specific guidance의 적절한 배치 판단
- 제외:
  - domain workflow 전체 설계
  - plugin packaging 자체
  - business feature implementation

## 처리하려는 작업 형태

- artifact의 tool behavior가 일관되지 않은 경우
- domain skill이 runtime-specific policy를 과도하게 흡수한 경우
- approval, delegation, browsing, local inspection의 기준을 명확히 해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/tool-use-guide/SKILL.md`
- 관련 상위 라우팅: `advance-codex-dev-guide`

## 핵심 처리 계약

- decision point를 기준으로 default action, fallback, escalation, non-goal을 정의한다.
- tool 이름만 나열하지 않고 trigger condition과 stop condition을 함께 적는다.
- local evidence로 확인 가능한 것은 먼저 로컬에서 확인한다.
- domain artifact에 묻어들어간 runtime-specific guidance는 필요한 경우 별도 tool-use artifact로 분리한다.

## 검토 질문

- 지금 정의하는 것이 domain workflow가 아니라 reusable tool policy가 맞는가?
- trigger, fallback, escalation, stop condition이 함께 적혀 있는가?
- local evidence로 먼저 확인할 수 있는 판단을 불필요하게 외부 의존으로 넘기고 있지 않은가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: tool policy는 domain workflow 바깥에서 재사용되는 surface이므로 hidden sibling context 없이 해석 가능해야 한다.

## 확장 원칙

- 새로운 tool rule은 durable decision pattern이 있을 때만 추가한다.
- artifact placement가 바뀌면 plugin guide나 custom agent spec과 함께 반영한다.
