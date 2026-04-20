# Workflow Kit 플러그인 스펙

## 목적

`workflow-kit`은 작업 lifecycle 전반을 다루는 재사용 가능한 workflow 플러그인입니다.
핵심 역할은 작업을 시작하기 전에 먼저 task shape를 안정화하고, 올바르게 정의하고, read-only planning을 통해 실행 준비를 끝내고, 적절한 실행 모드를 고르고, 안전하게 반복 개선하고, material finding을 scope drift 없이 처리하고, 마지막으로 commit 쪽으로 넘겨도 되는 상태인지 판단하는 것입니다.

## 경계

- 포함:
  - pre-workflow framing
  - work-definition 및 alignment workflow
  - read-only planning workflow
  - end-to-end execution workflow
  - bounded refinement loop
  - review-driven fix loop
  - final readiness gate
- 제외:
  - 도메인 특화 architecture guidance
  - frontend, backend, design specialist 조언
  - teammate runtime orchestration

## 진입 표면

- 대표 엔트리포인트: `workflow-kit-guide`
- 핵심 분기: 지금 필요한 것이 framing, definition, planning, execution, review-loop handling, readiness gate 중 무엇인지 먼저 분류한다

## 스킬 구성

- `workflow-kit-guide`: 작업을 올바른 workflow stage와 starting mode로 라우팅한다
- `structured-thinking`: 아직 어떤 workflow로 들어가야 할지 불안정한 작업을 안정화하고 다음 경로를 고를 수 있게 만든다
- `deep-interview`: planning이나 execution 전에 intent, scope, tradeoff, approval boundary를 명확히 해서 작업을 제대로 정의한다
- `planner`: 실행 전에 read-only investigation과 tradeoff 분석으로 decision-complete plan을 만든다
- `autopilot`: brief부터 implementation, verification까지 broad end-to-end delivery를 수행한다
- `parallel-work`: 소수의 명확히 독립적인 lane으로 분리하고 결과를 통합한다
- `ralph-loop`: 하나의 bounded issue를 fix-verify-reassess cycle로 반복 개선한다
- `review-loop`: blocking-first 기준으로 review finding을 처리하고 bounded fix loop를 돈다
- `commit-readiness-gate`: commit-ready 판단 전에 최종 self-review, scoped verification, risk classification을 수행한다

## 확장 원칙

- 새 workflow skill은 framing, definition, planning, end-to-end delivery, bounded refinement, review handling, final gating과 분명히 다른 work mode일 때만 추가한다.
- 도메인 중립적인 workflow logic만 이 플러그인에 둔다.
- 새 skill이 stage model을 바꾸면 `workflow-kit-guide`를 같은 변경에서 함께 갱신한다.
- 편하다는 이유로 하나의 workflow가 여러 stage를 흡수하게 두지 않는다.

## 현재 의도 점검

- 현재 플러그인 표면은 framing, definition, planning, execution, review, gate lifecycle을 중심으로 일관적이어야 한다.
- 현재의 주요 리스크는 stage가 빠지거나 서로 흡수되면서 lifecycle이 다시 execution 중심으로 축소되는 것이다.
