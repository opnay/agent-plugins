# Adaptive Subagent Orchestrator Dev

`adaptive-subagent-orchestrator-dev`는 복잡한 software-engineering 요청을 저비용으로 선별하고, 병렬 위임 가치가 확인될 때만 bounded subagent를 생성하고 결과를 검증·통합하는 instruction-only 플러그인입니다.

세 skill이 lifecycle을 나눠 소유합니다.

- `orchestrate-subagents`: 유일한 implicit entry입니다. explicit subagent 요청 또는 명백한 독립 workstream을 `DIRECT` 또는 `DISPATCH`로 라우팅합니다.
- `dispatch-subagents`: 전체 delegation gate와 execution mode를 판단하고, complete task packet으로 최소 subagent를 실제 생성합니다.
- `integrate-subagent-results`: 완전한 dispatch manifest의 필수 결과를 기다리고 evidence, 충돌, ownership, 전체 검증을 확인합니다.

기본 시작점:

- `$adaptive-subagent-orchestrator-dev:orchestrate-subagents`
- "현재 브랜치를 보안, 정확성, 테스트, 성능 관점에서 독립적으로 검토해줘."
- "로그인, 결제, 인벤토리 실행 경로의 간헐적 실패 원인을 각각 조사해줘."

Focused skill 직접 호출:

- `$adaptive-subagent-orchestrator-dev:dispatch-subagents`: raw request 또는 `RouteDecision`이 필요하며 delegation gate를 우회하지 않습니다.
- `$adaptive-subagent-orchestrator-dev:integrate-subagent-results`: active agent IDs와 완전한 `DispatchManifest`가 필요합니다.

일반적인 구현·리뷰·테스트 표현, 높은 복잡성, 많은 파일만으로는 implicit dispatch하지 않습니다. 위임 게이트가 통과하지 않으면 main agent가 직접 처리합니다.
