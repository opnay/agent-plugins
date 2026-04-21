# AGENTS.global.md

## 지속 대화 플로우

- 다음 스텝을 사용자의 자유 채팅 입력에 의존하지 않습니다.
- 추가 입력이 필요하면 질문 도구를 사용합니다.
- 질문이 1-3개의 bounded choice로 표현 가능하면 질문 도구를 우선 사용합니다.
- `"원하면 제가..."`, `"읽고 다음 행동 정해주세요"` 같은 종료 문구를 사용하지 않습니다.
- planning, 설계, 구현, 검증 단계에서는 질문 흐름을 유지합니다.
- 자유 채팅 입력으로 자연스럽게 대화가 닫히는 시점은 기본적으로 `commit ready`뿐입니다.

## workflow-kit

- 사용자 요청을 분석한 뒤 `workflow-kit` skill을 직접 선택해 사용합니다.
- 사용자가 직접 `workflow-kit`을 호출하는 방식으로 보지 않습니다.
- 현재 작업 흐름에 맞는 시작점은 `workflow-kit-guide`에서 고릅니다.
- `structured-thinking`, `deep-interview`, `planner`, `autopilot`, `parallel-work`, `ralph-loop`, `review-loop`, `commit-readiness-gate` 중 현재 병목을 해결하는 skill로 이어갑니다.
- high-level 조언만 주고 멈추지 않습니다.
- `deep-interview`가 맞으면 실제 질문으로 이어갑니다.
- `planner`가 맞으면 계획으로 이어갑니다.
- `autopilot`이 맞으면 실행으로 이어갑니다.

## advance-codex

- creator 계열 작업 전에는 `advance-codex` skill을 먼저 검토합니다.
- `skill-creator`를 사용할 때는 `advance-codex:skill-creator`를 함께 검토합니다.
- `plugin-creator`를 사용할 때는 `advance-codex:plugin-creator`를 함께 검토합니다.
- tool 사용 방식이 결과 품질에 직접 영향을 주는 작업이면 `advance-codex:tool-use-guide`를 먼저 검토합니다.
- `tool-use-guide`는 터미널 사용, 파일 수정, 질문 도구, ask-vs-infer, escalation 같은 도구 사용 규칙을 보강하는 용도로 사용합니다.

## 결합 규칙

- 먼저 `workflow-kit`으로 작업 흐름을 고릅니다.
- creator 계열 작업에서는 대응되는 `advance-codex` 보강 skill을 같이 검토합니다.
- workflow 진행 중 도구 사용 규칙이 중요해지면 `advance-codex:tool-use-guide`를 같이 검토합니다.
- specialist plugin은 선택된 workflow 안에서 필요할 때 사용합니다.

## 금지

- 저장소 소개 문구를 적지 않습니다.
- 플러그인 소개 문구를 적지 않습니다.
- `workflow-kit`을 사용자 직접 호출형으로 적지 않습니다.
- creator 계열 작업을 `advance-codex` 보강 없이 단독 처리하지 않습니다.
- 질문이 필요한데 advisory answer로 대화를 종료하지 않습니다.
