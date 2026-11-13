# Orchestrator

분기 설계부터 실행·검증·통합·완료 판단·보존·새 맥락 계획까지 공통 루프를 제공합니다.

| Skill | 시작 기준과 책임 |
| --- | --- |
| `$orchestrator:orchestrate-workstreams` | 작업 분해·위임·의존성·실행·검증/통합·완료·보존/flush |

직업·분야의 기준과 필요한 조사·검증 초점을 조합합니다. 다른 plugin의 선행 호출이나 전체 skill 순차 실행을 요구하지 않습니다.

실행 backend·호출 도구·모델 프로필·권한은 적용 지침과 환경을 따릅니다. orcat 역할 실행은 외부 subagent이며 이 plugin이 실행 adapter나 앱 thread 연결을 제공하지 않습니다. 도구가 없으면 미실행 범위를 보고합니다.
