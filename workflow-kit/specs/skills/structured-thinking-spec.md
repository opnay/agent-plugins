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

- 대표 표면: `workflow-kit/skills/structured-thinking/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- ambiguity를 workflow selection에 필요한 수준까지만 정리한다.
- user intent discovery가 병목이면 `deep-interview`로 handoff한다.
- next path가 정해지면 advisory answer로 멈추지 않고 handoff를 명시한다.

## 독립성 원칙

- 이 스킬은 requirements interview를 소유하지 않는다.
- framing 결과만 읽어도 다음 workflow를 안전하게 고를 수 있어야 한다.

## 확장 원칙

- workflow ambiguity 처리 방식이 바뀌면 guide skill과 함께 갱신한다.

