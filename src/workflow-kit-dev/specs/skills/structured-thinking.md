## 사용자 스펙 의도

- workflow 선택 자체가 불안정할 때 task shape를 먼저 안정화하고 싶다.
- user intent discovery와 workflow ambiguity를 구분하고 싶다.
- advisory answer로 멈추지 않고 다음 workflow handoff까지 명시하고 싶다.

---

# structured-thinking 스킬 스펙

## 목적

`structured-thinking`는 다음 workflow를 아직 안전하게 고를 수 없을 때 task shape를 안정화하는 pre-workflow framing 스킬입니다.

## 경계

- 포함:
  - ambiguity와 risky assumption 식별
  - plausible next path 비교
  - workflow choice를 위한 최소 질문 추출
- 제외:
  - intent discovery interview
  - planning
  - implementation

## 처리하려는 작업 형태

- task shape가 아직 불안정해서 workflow 선택이 어려운 경우
- ambiguity가 user intent보다 workflow choice에 더 가까운 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit-dev/skills/structured-thinking/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

- ambiguity를 workflow selection에 필요한 수준까지만 정리한다.
- user intent discovery가 병목이면 `deep-interview`로 handoff한다.
- next path가 정해지면 advisory answer로 멈추지 않고 handoff를 명시한다.

## 검토 질문

- 현재 병목이 user intent discovery가 아니라 workflow selection ambiguity가 맞는가?
- ambiguity를 workflow 선택에 필요한 수준까지만 다루고 있는가?
- next path나 handoff가 advisory answer 없이 충분히 명시됐는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: framing 결과는 hidden sibling context 없이 다음 workflow를 고를 수 있어야 하며, requirements interview는 `deep-interview`에 남겨야 한다.

## 확장 원칙

- workflow ambiguity 처리 방식이 바뀌면 guide skill과 함께 갱신한다.
