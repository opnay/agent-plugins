## 사용자 스펙 의도

- broad end-to-end delivery를 하나의 execution workflow로 밀고 가고 싶다.
- 분석, 구현, QA, review, verification을 끊기지 않게 운영하고 싶다.
- planning-only나 final-gate를 이 스킬 안으로 흡수하지 않고 분리하고 싶다.

---

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

- 대표 표면: `workflow-kit-dev/skills/autopilot/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

- broad execution이 필요한지 먼저 확인하고, 맞다면 단계별 진행과 검증을 이어서 수행한다.
- 분석, 구현, QA, review를 분절되지 않은 흐름으로 관리한다.
- stop condition과 final checklist를 명시적으로 점검한다.

## 검토 질문

- 지금 작업이 정말 broad end-to-end execution이 필요한 범위인가?
- 단계별 진행과 검증이 끊기지 않게 설계돼 있는가?
- planning-only, narrow-loop, final-gate 역할을 이 스킬이 과도하게 흡수하고 있지 않은가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: broad execution workflow는 hidden sibling context 없이도 적용 가능해야 하며, planning이나 final gate는 후속/인접 skill로만 넘겨야 한다.

## 확장 원칙

- 새 규칙은 broad execution lifecycle을 더 안정적으로 운영할 때만 추가한다.
