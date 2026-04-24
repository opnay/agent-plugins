## 사용자 스펙 의도

- 구현 전에 decision-complete plan을 먼저 만들고 싶다.
- read-only 조사와 tradeoff analysis를 선행하고 싶다.
- open question과 working assumption을 섞지 않고 정리하고 싶다.

---

# planner 스킬 스펙

## 목적

`planner`는 read-only investigation을 통해 decision-complete, execution-ready plan을 만드는 planning workflow 스킬입니다.

## 경계

- 포함:
  - scope locking
  - read-only 조사
  - step-wise execution plan
  - risks, dependencies, open question 정리
- 제외:
  - implementation
  - state-changing verification
  - final readiness gate

## 처리하려는 작업 형태

- 구현 전 승인 가능한 계획이 필요한 경우
- read-only 조사와 tradeoff analysis가 먼저 필요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit-dev/skills/planner/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

- planning 단계에서는 시스템 상태를 바꾸지 않는다.
- 각 step은 `Action`, `Expected Output`, `Verification Method`를 포함해야 한다.
- open question과 working assumption을 명확히 구분한다.

## 검토 질문

- 이번 작업이 구현이 아니라 read-only planning 단계로 유지되고 있는가?
- 각 step이 `Action`, `Expected Output`, `Verification Method`를 포함하는가?
- open question과 working assumption이 구분돼 있는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: planning artifact는 hidden sibling context 없이 실행 준비 상태를 설명해야 하며, execution이나 refinement는 별도 skill로 넘겨야 한다.

## 확장 원칙

- 새 규칙은 decision-complete planning의 품질을 높일 때만 추가한다.
