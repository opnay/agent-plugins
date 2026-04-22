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

- 대표 표면: `workflow-kit/skills/planner/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- planning 단계에서는 시스템 상태를 바꾸지 않는다.
- 각 step은 `Action`, `Expected Output`, `Verification Method`를 포함해야 한다.
- open question과 working assumption을 명확히 구분한다.

## 독립성 원칙

- 이 스킬은 execution이나 refinement loop를 소유하지 않는다.
- plan artifact만 읽어도 실행 준비 상태를 판단할 수 있어야 한다.

## 확장 원칙

- 새 규칙은 decision-complete planning의 품질을 높일 때만 추가한다.

