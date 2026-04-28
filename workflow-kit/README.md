# Workflow Kit

`workflow-kit`은 작업을 어떻게 질문하고, 분석하고, 정의하고, 계획하고, 실행하고, 검토하고, 마무리할지를 다루는 범용 workflow 플러그인입니다.
기본 시작점은 `workflow-kit-guide`이며, 들어온 요청을 먼저 어떤 workflow로 처리할지 정한 뒤 그 안의 skill로 라우팅합니다.
요구사항 파악, 방향 점검, sequential problem-solving, scope locking, read-only planning, end-to-end execution, bounded refinement, review handling, readiness judgment, 그리고 사용자가 명시적으로 작업을 끝낼 때까지 턴을 닫지 않는 turn-level loop gate 같은 workflow 작업을 다룹니다.

이 플러그인의 목적은 더 나은 workflow guidance로 작업의 완성도를 높이는 것입니다.
반대로 특정 도메인의 구현 가이드로 변하거나, workflow 범위를 벗어난 architecture guidance를 흡수하지 않습니다.

`workflow-kit`은 또한 `loop-kit`의 SSOT입니다.
`loop-kit`이 사용하는 turn-gate 중심 operational surface와 internal loop mode 의미는 이 플러그인의 broader workflow taxonomy와 canonical skill contract를 기준으로 유지합니다.
