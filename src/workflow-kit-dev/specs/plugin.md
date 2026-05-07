# Workflow Kit Dev 플러그인 스펙

## 플러그인 목적

`workflow-kit-dev`은 작업 lifecycle 전반을 다루는 workflow 플러그인입니다.
핵심 책임은 requirement discovery, sequential analysis, planning, execution, refinement, review, final gating을 각각의 workflow skill로 제공하는 것입니다.
workflow 선택 전에는 사용자 지시어의 operation 의미가 skill, spec, plugin, phase, routing rule, release surface 중 어디를 가리키는지 확인해야 하며, 해석에 따라 작업이 달라지면 meaning resolution 질문으로 먼저 잠급니다.
각 workflow skill은 직접 호출 가능한 작업 표면으로 유지합니다.

## 플러그인 경계와 비목표

- 포함:
  - sequential and reflective problem solving
  - operation meaning resolution before workflow selection
  - requirement discovery와 direction evaluation
  - read-only planning
  - broad execution workflow
  - bounded refinement loop
  - review-driven correction
  - final readiness gate
- 제외:
  - 도메인 특화 architecture guidance
  - frontend, design, teammate specialist advice
  - runtime-specific teammate orchestration

## 처리하려는 작업 형태

- 요청을 어떤 workflow로 먼저 처리해야 하는지 결정하는 작업
- 사용자의 구조 변경 지시어가 여러 작업 단위를 가리킬 수 있어 의미를 먼저 잠가야 하는 작업
- 복잡한 분석, 설계, 계획, 디버깅에서 단계적 사고, revision, branch, hypothesis verification이 필요한 작업
- 구현 전에 alignment나 planning이 필요한 작업
- broad execution부터 review와 final gate까지 이어지는 lifecycle 작업
- 여러 workflow skill 중 어떤 표면을 직접 사용할지 정해야 하는 작업

## 대표 표면

- 대표 스펙: `workflow-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `workflow-kit-dev/specs/skills/*.md`
- 보조 적응 문서: `workflow-kit-dev/specs/deep-interview-adaptation.md`

## 내장 skill 체계

- `sequential-thinking`: 복잡한 분석, 설계, 계획, 디버깅 문제를 단계적으로 풀고 필요하면 revision, branch, hypothesis verification을 수행한다.
  - spec: `workflow-kit-dev/specs/skills/sequential-thinking.md`
- `deep-interview`: intent, scope, tradeoff, approval boundary를 질문으로 잠근다.
  - spec: `workflow-kit-dev/specs/skills/deep-interview.md`
- `planner`: read-only investigation을 통해 decision-complete plan을 만든다.
  - spec: `workflow-kit-dev/specs/skills/planner.md`
- `autopilot`: brief부터 implementation, verification까지 broad end-to-end delivery를 수행한다.
  - spec: `workflow-kit-dev/specs/skills/autopilot.md`
- `parallel-work`: 소수의 독립 lane으로 분리하고 결과를 통합한다.
  - spec: `workflow-kit-dev/specs/skills/parallel-work.md`
- `ralph-loop`: 하나의 bounded issue를 iterative fix-verify-reassess loop로 개선한다.
  - spec: `workflow-kit-dev/specs/skills/ralph-loop.md`
- `review-loop`: blocking-first 기준으로 review finding을 처리한다.
  - spec: `workflow-kit-dev/specs/skills/review-loop.md`
- `commit-readiness-gate`: final self-review와 scoped verification으로 commit-ready 여부를 판단한다.
  - spec: `workflow-kit-dev/specs/skills/commit-readiness-gate.md`

## SDD 운영 원칙

- plugin spec은 lifecycle stage model과 skill composition만 소유한다.
- 각 workflow의 처리 계약은 `specs/skills/` 아래 독립 문서로 분리한다.
- stage model이나 handoff 규칙이 바뀌면 해당 skill spec, `plugin.md`, manifest prompt를 같은 변경에서 갱신한다.
- specialist plugin이 일반 workflow 책임을 흡수하지 않도록 global workflow boundary를 유지한다.

## 현재 구조 메모

- `deep-interview-adaptation.md`는 적응 배경 문서로 유지하되 normative skill contract는 `specs/skills/deep-interview.md`가 소유한다.
- 이 플러그인의 주요 리스크는 lifecycle stage가 execution 중심으로 납작해지거나, 특정 workflow skill이 다른 skill의 책임을 과도하게 흡수하는 것이다.
