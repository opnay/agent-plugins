## 사용자 스펙 의도

- 실제 요구사항을 advisory answer가 아니라 질문 라운드로 잠그고 싶다.
- intent, scope, tradeoff, approval boundary를 명시적으로 확인하고 싶다.
- 충분한 clarity를 얻으면 다음 workflow로 자연스럽게 handoff하고 싶다.

---

# deep-interview 스킬 스펙

## 목적

`deep-interview`는 질문과 압력 테스트를 통해 사용자의 실제 intent, scope, tradeoff, approval boundary, success criteria를 잠그는 workflow 스킬입니다.

## 경계

- 포함:
  - intent-first clarification
  - scope, non-goal, tradeoff, decision boundary 잠금
  - `request_user_input` 또는 일반 질문을 통한 requirement discovery
  - downstream workflow나 specialist plugin으로의 handoff 준비
- 제외:
  - generic workflow ambiguity 정리만 하는 일
  - read-only planning
  - implementation

## 처리하려는 작업 형태

- 요구사항이 아직 잠기지 않은 proposal, direction check, greenfield setup
- "괜찮을까?" 같은 평가형 질문에서 실제 질문 라운드가 필요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/deep-interview/SKILL.md`
- 보조 적응 문서: `workflow-kit/specs/deep-interview-adaptation.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- 질문은 intent, scope, non-goal, tradeoff, acceptance signal을 잠그는 방향으로 진행한다.
- bounded choice로 잠글 수 있으면 `request_user_input`을 우선 사용한다.
- advisory answer로 종료하지 않고 필요한 질문 라운드를 실제로 수행한다.
- 충분한 clarity를 얻으면 planner, autopilot, parallel-work, review-loop 또는 specialist plugin으로 handoff한다.

## 검토 질문

- 지금 병목이 workflow ambiguity가 아니라 실제 requirement discovery인가?
- bounded choice로 잠글 수 있는 질문인데 자유서술형으로 흐르고 있지 않은가?
- discovery 이후 handoff 대상이 충분히 선명한가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: discovery 결과는 hidden sibling context 없이 downstream workflow가 바로 시작 가능해야 하며, `structured-thinking`과의 경계도 이 스펙 자체에서 읽혀야 한다.

## 확장 원칙

- 적응 규칙이 바뀌면 `deep-interview-adaptation.md`, `workflow-kit-guide`, 이 스펙을 함께 점검한다.
