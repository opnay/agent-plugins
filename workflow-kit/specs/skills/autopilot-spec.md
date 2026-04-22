# autopilot 스킬 스펙

## 목적

`autopilot`은 brief부터 implementation, QA, review, verification까지 broad end-to-end delivery를 수행하는 execution workflow 스킬입니다.

## 경계

- 포함:
  - multi-phase execution
  - analysis, implementation, QA, review의 일관된 진행
  - broad verification과 stop condition 관리
- 제외:
  - read-only planning only
  - narrow bounded refinement
  - final commit gate only

## 처리하려는 작업 형태

- 사용자가 hands-off end-to-end delivery를 원하는 경우
- 여러 phase가 이어지는 구현 작업을 한 workflow로 끌고 가야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/autopilot/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- broad execution이 필요한지 먼저 확인하고, 맞다면 단계별 진행과 검증을 이어서 수행한다.
- 분석, 구현, QA, review를 분절되지 않은 흐름으로 관리한다.
- stop condition과 final checklist를 명시적으로 점검한다.

## 독립성 원칙

- 이 스킬은 planning-only, narrow-loop, final-gate 역할을 대체하지 않는다.
- end-to-end execution workflow만 읽어도 적용 가능해야 한다.

## 확장 원칙

- 새 규칙은 broad execution lifecycle을 더 안정적으로 운영할 때만 추가한다.

