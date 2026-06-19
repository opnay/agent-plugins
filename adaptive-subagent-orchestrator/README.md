# Adaptive Subagent Orchestrator

`adaptive-subagent-orchestrator`는 복잡한 소프트웨어 엔지니어링 요청을 독립 작업 흐름으로 나눌 수 있는지 판단하고, 병렬 위임 이점이 분명할 때만 최소한의 서브에이전트를 조율하는 플러그인입니다.

이 플러그인은 단일 instruction-only skill을 제공합니다.
사용자가 서브에이전트를 직접 언급하지 않아도 다중 모듈 조사, 교차 관점 리뷰, 독립 테스트 실패 분석, 마이그레이션 영향 분석, 기술 대안 비교처럼 안전하게 나눌 수 있는 작업이면 활성화될 수 있습니다.

이 플러그인은 서브에이전트 사용을 항상 강제하지 않습니다.
위임 게이트가 통과하지 않으면 메인 에이전트가 직접 처리합니다.

대표 호출:

- `$adaptive-subagent-orchestrator:adaptive-subagent-orchestrator`
- "현재 브랜치를 main과 비교해서 보안 문제, 실제 버그, 테스트 누락, 성능 위험을 검토해줘."
- "unit, integration, end-to-end 테스트 묶음이 각각 실패하고 있어. 원인을 찾고 최소 수정안을 적용해줘."
