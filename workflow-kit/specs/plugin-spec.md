# Workflow Kit 플러그인 스펙

## 플러그인 목적

`workflow-kit`은 작업 lifecycle 전반을 다루는 workflow 플러그인입니다.
핵심 책임은 들어온 요청에 대해 requirement discovery, framing, planning, execution, refinement, review, final gating, phase-loop continuity 중 현재 병목이 무엇인지 판단하고, 가장 맞는 workflow skill로 연결하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - pre-workflow framing
  - requirement discovery와 direction evaluation
  - read-only planning
  - broad execution workflow
  - phase-loop continuity를 관리하는 meta workflow
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
- 결과 보고 뒤 다음 플로우를 명시적으로 열어야 하는 지속적 turn workflow

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `workflow-kit-guide`
- 대표 스펙: `workflow-kit/specs/plugin-spec.md`
- skill 상세 스펙 위치: `workflow-kit/specs/skills/*.md`
- 보조 적응 문서: `workflow-kit/specs/deep-interview-adaptation-spec.md`

## 내장 skill 체계

- `workflow-kit-guide`: current bottleneck에 맞는 starting workflow와 handoff를 정한다.
  - spec: `workflow-kit/specs/skills/workflow-kit-guide-spec.md`
- `structured-thinking`: workflow 선택이 아직 불안정한 task를 안정화한다.
  - spec: `workflow-kit/specs/skills/structured-thinking-spec.md`
- `deep-interview`: intent, scope, tradeoff, approval boundary를 질문으로 잠근다.
  - spec: `workflow-kit/specs/skills/deep-interview-spec.md`
- `planner`: read-only investigation을 통해 decision-complete plan을 만든다.
  - spec: `workflow-kit/specs/skills/planner-spec.md`
- `turn-gate`: 질문/계획/명령 -> 작업 -> 결과 보고 -> 다음 플로우 질문을 루프로 이어가는 meta workflow를 관리한다.
  - spec: `workflow-kit/specs/skills/turn-gate-spec.md`
- `autopilot`: brief부터 implementation, verification까지 broad end-to-end delivery를 수행한다.
  - spec: `workflow-kit/specs/skills/autopilot-spec.md`
- `parallel-work`: 소수의 독립 lane으로 분리하고 결과를 통합한다.
  - spec: `workflow-kit/specs/skills/parallel-work-spec.md`
- `ralph-loop`: 하나의 bounded issue를 iterative fix-verify-reassess loop로 개선한다.
  - spec: `workflow-kit/specs/skills/ralph-loop-spec.md`
- `review-loop`: blocking-first 기준으로 review finding을 처리한다.
  - spec: `workflow-kit/specs/skills/review-loop-spec.md`
- `commit-readiness-gate`: final self-review와 scoped verification으로 commit-ready 여부를 판단한다.
  - spec: `workflow-kit/specs/skills/commit-readiness-gate-spec.md`

## SDD 운영 원칙

- plugin spec은 lifecycle stage model과 routing surface만 소유한다.
- 각 workflow의 처리 계약은 `specs/skills/` 아래 독립 문서로 분리한다.
- stage model이나 handoff 규칙이 바뀌면 `workflow-kit-guide`와 해당 skill spec, `plugin-spec.md`를 같은 변경에서 갱신한다.
- specialist plugin이 first stop이 되지 않도록 global routing boundary를 유지한다.

## 현재 구조 메모

- `deep-interview-adaptation-spec.md`는 적응 배경 문서로 유지하되 normative skill contract는 `specs/skills/deep-interview-spec.md`가 소유한다.
- 이 플러그인의 주요 리스크는 lifecycle stage와 meta-flow가 서로 흡수되면서 workflow가 execution 중심으로 납작해지거나, 반대로 메타 운영 규칙이 phase skill을 과도하게 오염시키는 것이다.
