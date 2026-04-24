# Workflow Kit 플러그인 스펙

## 플러그인 목적

`workflow-kit`은 작업 lifecycle 전반을 다루는 workflow 플러그인입니다.
핵심 책임은 들어온 요청에 대해 requirement discovery, framing, planning, execution, refinement, review, final gating, 그리고 사용자가 턴을 종료하자고 요청할때까지 턴을 닫지 않는 loop-gated continuity 중 현재 병목이 무엇인지 판단하고, 가장 맞는 workflow skill로 연결하는 것입니다.
repository-local operating rule이 non-terminal turn을 요구하면, `turn-gate`를 turn-level loop gate로 유지한 채 현재 phase owner를 선택합니다.
`turn-gate`가 활성화되면 현재 세션 동안 first-class loop gate rule로 취급합니다.
사용자 질문을 subagent 질문으로 바꿔 자동 진행해야 하는 경우 `self-drive`를 `turn-gate`의 question-routing mode로 선택합니다.
이 플러그인은 `loop-kit`이 사용하는 broader workflow taxonomy와 canonical loop contract의 SSOT이기도 합니다.

## 플러그인 경계와 비목표

- 포함:
  - pre-workflow framing
  - requirement discovery와 direction evaluation
  - read-only planning
  - broad execution workflow
  - turn-level loop gate contract
  - bounded refinement loop
  - review-driven correction
  - final readiness gate
- 제외:
  - 도메인 특화 architecture guidance
  - frontend, design, teammate specialist advice
  - runtime-specific teammate orchestration

## 처리하려는 작업 형태

- 요청을 어떤 workflow로 먼저 처리해야 하는지 결정하는 작업
- 구현 전에 alignment나 planning이 필요한 작업
- broad execution부터 review와 final gate까지 이어지는 lifecycle 작업
- 결과 보고 뒤 다음 플로우 진행을 위한 question-routing 응답을 열어야 하는 지속적 turn workflow

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `workflow-kit-guide`
- 대표 스펙: `workflow-kit/specs/plugin.md`
- skill 상세 스펙 위치: `workflow-kit/specs/skills/*.md`
- 보조 적응 문서: `workflow-kit/specs/deep-interview-adaptation.md`

## 내장 skill 체계

- `workflow-kit-guide`: current bottleneck에 맞는 starting workflow와 handoff를 정한다.
  - spec: `workflow-kit/specs/skills/workflow-kit-guide.md`
- `structured-thinking`: workflow 선택이 아직 불안정한 task를 안정화한다.
  - spec: `workflow-kit/specs/skills/structured-thinking.md`
- `deep-interview`: intent, scope, tradeoff, approval boundary를 질문으로 잠근다.
  - spec: `workflow-kit/specs/skills/deep-interview.md`
- `planner`: read-only investigation을 통해 decision-complete plan을 만든다.
  - spec: `workflow-kit/specs/skills/planner.md`
- `turn-gate`: `분석 -> 계획 -> 작업 -> 결과 보고 / commit-ready -> 다음 플로우 진행을 위한 question-routing 응답` 구조를 유지하고, repository rule이 요구하면 사용자가 턴을 종료하자고 요청할때까지 턴을 닫지 않는 loop gate를 관리한다.
  - spec: `workflow-kit/specs/skills/turn-gate.md`
- `autopilot`: brief부터 implementation, verification까지 broad end-to-end delivery를 수행한다.
  - spec: `workflow-kit/specs/skills/autopilot.md`
- `parallel-work`: 소수의 독립 lane으로 분리하고 결과를 통합한다.
  - spec: `workflow-kit/specs/skills/parallel-work.md`
- `ralph-loop`: 하나의 bounded issue를 iterative fix-verify-reassess loop로 개선한다.
  - spec: `workflow-kit/specs/skills/ralph-loop.md`
- `review-loop`: blocking-first 기준으로 review finding을 처리한다.
  - spec: `workflow-kit/specs/skills/review-loop.md`
- `commit-readiness-gate`: final self-review와 scoped verification으로 commit-ready 여부를 판단한다.
  - spec: `workflow-kit/specs/skills/commit-readiness-gate.md`

## SDD 운영 원칙

- plugin spec은 lifecycle stage model과 routing surface만 소유한다.
- 각 workflow의 처리 계약은 `specs/skills/` 아래 독립 문서로 분리한다.
- stage model이나 handoff 규칙이 바뀌면 `workflow-kit-guide`와 해당 skill spec, `plugin.md`를 같은 변경에서 갱신한다.
- specialist plugin이 first stop이 되지 않도록 global routing boundary를 유지한다.

## 현재 구조 메모

- `deep-interview-adaptation.md`는 적응 배경 문서로 유지하되 normative skill contract는 `specs/skills/deep-interview.md`가 소유한다.
- 이 플러그인의 주요 리스크는 lifecycle stage와 turn-level loop gate가 서로 흡수되면서 workflow가 execution 중심으로 납작해지거나, 반대로 loop gate 규칙이 phase skill을 과도하게 오염시키는 것이다.
- `loop-kit`은 이 플러그인의 narrower operational package로 두고, `turn-gate` 중심 표면과 internal loop mode orchestration을 별도 플러그인으로 노출한다.
- `loop-kit`의 internal mode 의미는 `workflow-kit/specs/skills/ralph-loop.md`, `workflow-kit/specs/skills/review-loop.md`, `workflow-kit/specs/skills/commit-readiness-gate.md`를 canonical upstream으로 본다.
