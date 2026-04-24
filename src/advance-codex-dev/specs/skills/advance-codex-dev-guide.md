## 사용자 스펙 의도

- 먼저 지금 다루는 산출물이 무엇인지 분류하고 싶다.
- artifact type이 섞여 있을 때 가장 좁은 creator skill과 실행 순서를 정하고 싶다.
- guide가 직접 상세 설계를 흡수하지 않고 적절한 sibling skill로 보내길 원한다.

---

# advance-codex-dev-guide 스킬 스펙

## 목적

`advance-codex-dev-guide`는 `advance-codex-dev` 플러그인의 첫 라우터로서, Codex 활용 개선 작업의 주된 artifact type을 분류하고 가장 좁은 built-in skill과 실행 순서를 정하는 역할을 맡습니다.

## 경계

- 포함:
  - skill, empirical evaluation, tool-use guidance, plugin, session workflow, commit workflow, custom agent 중 주된 artifact 분류
  - revision, creation, packaging 여부 판단
  - multi-artifact 작업의 실행 순서 결정
- 제외:
  - 각 artifact의 상세 구현
  - plugin 외부 workflow 선택
  - domain-specific execution guidance

## 처리하려는 작업 형태

- "어떤 creator skill을 먼저 써야 하는가"가 병목인 작업
- 여러 Codex-facing artifact가 섞여 있어 우선순위를 정해야 하는 작업
- plugin 내부 guide가 먼저 작업 형태를 분류해야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/advance-codex-dev-guide/SKILL.md`
- 상위 맥락: `workflow-kit`이 execution workflow를 고른 뒤, creator-oriented specialist routing이 필요할 때 사용한다

## 핵심 처리 계약

- 먼저 주된 reusable artifact를 하나로 고정한다.
- creation, revision, packaging 중 현재 task shape를 판별한다.
- plugin-level concern이면 plugin packaging을 우선하고, 하위 skill은 그 다음 순서로 둔다.
- multi-artifact 작업이면 실행 순서를 명시하고 섞지 않는다.
- 출력은 artifact 분류, 선택된 creator skill, 실행 순서, 주요 리스크를 포함해야 한다.

## 검토 질문

- 지금의 main artifact가 skill, plugin, tool policy, session, commit, custom agent 중 무엇인가?
- plugin-level concern을 먼저 잠가야 하는데 하위 skill로 성급히 내려가고 있지 않은가?
- 선택한 실행 순서가 여러 artifact를 한 단계에 섞지 않도록 유지하는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 스킬은 guide이므로 sibling creator skill map을 전제로 해도 되지만, 라우팅 기준 자체는 이 스펙만 읽어도 이해 가능해야 한다.

## 확장 원칙

- 새 artifact type이 plugin boundary 안에서 독립 concern이 될 때만 라우팅 분기를 추가한다.
- routing 분기가 바뀌면 plugin spec과 관련 skill spec을 함께 갱신한다.
