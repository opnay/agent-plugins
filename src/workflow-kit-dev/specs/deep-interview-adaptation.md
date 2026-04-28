# Deep Interview Adaptation Spec

## 참조 원본

- 원본 스킬 이름: `deep-interview`
- 원본 저장소: `Yeachan-Heo/oh-my-codex`
- 원본 경로: `skills/deep-interview/SKILL.md`
- 원본 URL: `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`
- 이 문서는 원본 스킬을 그대로 복제하는 문서가 아니라, 현재 `agent-plugins` 환경에 맞게 적응시키기 위한 스펙이다.

## 목적

이 문서는 원본 `deep-interview` skill의 핵심을 현재 `agent-plugins` 환경에 맞게 다시 정의하기 위한 적응 스펙이다.

핵심 목적은 다음과 같다.

- 원본 skill의 intent-first clarification 철학은 유지한다.
- 원본 skill의 호출 방식, CLI, 상태 파일, OMX artifact 전제를 그대로 가져오지 않는다.
- 현재 환경의 시작점인 `workflow-kit-dev-guide`와 자연스럽게 결합되도록 `deep-interview`를 재정의한다.
- 질문이 필요한 상황에서 advisory answer로 종료되지 않고, 실제 질문 흐름으로 이어지게 만든다.
- 질문 방식은 현재 환경의 tool-use 규칙에 맞게 `request_user_input` 중심으로 정렬한다.

## 문제 정의

원본 `deep-interview`는 다음 전제를 가진다.

- 사용자가 skill을 직접 호출한다.
- skill은 자체적인 CLI 인자, 상태 저장, artifact 경로, 후속 bridge를 가진다.
- 질문 단계에서 OMX 고유 도구를 필수 경로로 사용한다.

현재 환경은 다음 전제를 가진다.

- 사용자가 skill을 직접 호출하지 않는다.
- 에이전트가 요청을 분석하고 skill을 자체적으로 선택한다.
- global 시작점은 `workflow-kit-dev-guide`다.
- 질문 도구는 현재 런타임의 `request_user_input`과 일반 대화 questioning이다.
- skill은 독립 실행 artifact가 아니라 플러그인 내부 라우팅 체인의 일부다.

따라서 원본 skill의 핵심을 가져오되, 호출 방식과 운영 구조는 새로 설계해야 한다.

## 적응 원칙

### 1. 핵심만 계승

원본에서 계승할 핵심은 아래로 제한한다.

- intent-first clarification
- one question per round
- 질문보다 답의 압력 테스트가 중요하다는 관점
- scope, non-goal, tradeoff, decision boundary를 명시적으로 잠그는 방식
- requirements가 잠기기 전에는 계획이나 구현으로 너무 빨리 넘어가지 않는 태도

### 2. 호출 방식은 전면 교체

원본에서 버릴 요소는 아래다.

- CLI argument model
- OMX 전용 question command
- OMX state persistence contract
- OMX artifact path contract
- OMX-specific execution bridge

현재 환경에서는 다음으로 대체한다.

- 시작점: `workflow-kit-dev-guide`
- 질문 도구: `request_user_input` 또는 일반 대화 질문
- handoff: `workflow-kit-dev` 내부 다른 skill 또는 downstream specialist plugin

### 3. skill은 직접 호출형이 아니라 라우팅형이다

`deep-interview`는 더 이상 사용자가 직접 진입하는 primary UX가 아니다.
현재 환경에서의 primary UX는 다음과 같다.

1. 사용자가 작업이나 질문을 던진다.
2. `workflow-kit-dev-guide`가 현재 workflow를 선택한다.
3. 요구사항 파악과 방향 잠금이 병목이면 `deep-interview`를 선택한다.
4. `deep-interview`는 실제 질문을 수행하고, 그 결과를 잠근 뒤 다음 workflow로 handoff한다.

즉 `deep-interview`는 direct-entry skill이 아니라 routed workflow skill이다.

## 플러그인 내 역할 정의

### `workflow-kit-dev-guide`

이 skill은 global first-read router다.

이 skill의 책임:

- 요청을 보고 현재 workflow bottleneck을 고른다.
- `deep-interview`, `planner`, `autopilot` 등 중 시작 skill을 정한다.
- specialist plugin은 직접 시작점으로 고르지 않고, workflow 안에서 필요할 때 handoff 대상으로만 둔다.

이 skill은 다음 질문을 해결해야 한다.

- 지금 필요한 것이 workflow 선택인가
- 실제 질문 인터뷰인가
- read-only planning인가
- broad execution인가
- review correction인가

### `deep-interview`

이 skill은 requirements discovery와 direction evaluation을 실제로 수행하는 skill이다.

이 skill의 책임:

- 사용자가 실제로 원하는 결과를 질문으로 밝혀낸다.
- scope edge, non-goal, tradeoff, acceptance signal을 잠근다.
- proposal evaluation이나 “괜찮을까?”류 질문에서 빠른 advisory answer로 끝내지 않고, 필요하면 질문을 이어간다.
- execution-ready brief 또는 direction-ready brief를 만든다.

이 skill의 비책임:

- broad architecture planning 전체
- implementation 자체
- final readiness gate

## 라우팅 규칙

### `workflow-kit-dev-guide -> deep-interview`로 가야 하는 경우

다음 조건이면 단순 workflow 선택보다 `deep-interview`를 우선한다.

- 요청이 proposal, direction check, “괜찮을까?”, “어떻게 가는 게 맞나?” 같은 평가형 질문이다.
- 신규 프로젝트, 신규 구조, 신규 시스템 제안인데 사용자의 실제 우선순위가 아직 드러나지 않았다.
- 사용자가 어느 정도 solution wording은 줬지만, success criteria나 non-goal이 아직 잠기지 않았다.
- 답을 주기 전에 먼저 질문으로 intent를 드러내야 더 나은 판단이 가능하다.
- high-level recommendation 뒤에 바로 “원하시면 다음 턴에서…”로 미루지 말고 지금 질문을 이어가야 한다.

### handoff 규칙

- `deep-interview`가 enough clarity를 확보했다면 `planner`, `autopilot`, `parallel-work`, `review-loop` 또는 specialist plugin handoff를 추천한다.
- specialist plugin handoff는 requirements discovery 이후에 붙는다.

## 질문 도구 규격

현재 환경에서 `deep-interview`는 질문 도구를 아래 원칙으로 사용한다.

### 1. 기본 원칙

- 질문은 one-question-per-round를 유지한다.
- 사용자 effort를 줄이는 방향으로 질문을 구성한다.
- repository facts는 먼저 로컬에서 확인하고, 확인할 수 없는 것만 사용자에게 묻는다.

### 2. `request_user_input` 사용 조건

다음 조건이면 `request_user_input`을 우선 사용한다.

- 질문이 1-3개의 짧은 bounded choice로 표현 가능하다.
- 각 선택지의 tradeoff를 짧게 설명할 수 있다.
- 답이 다음 workflow나 requirement lock에 직접적으로 영향을 준다.

예:

- 첫 번째 showcase를 무엇으로 둘지
- palette strategy를 어디까지 열어둘지
- monorepo 초기 범위를 어디까지 잠글지

### 3. 일반 대화 질문 사용 조건

다음 조건이면 일반 질문을 사용한다.

- 사용자의 답이 자유 서술이어야 한다.
- choice를 미리 잘라내는 것이 오히려 intent discovery를 망친다.
- example, counterexample, non-goal, tradeoff rejection 같은 서술형 답이 필요하다.

### 4. 금지

- 추가 user input이 필요하다고 판단했는데도 high-level advisory answer로 종료하지 않는다.
- 질문이 bounded choice로 충분한데 장문의 자유 질문으로 사용자 effort를 늘리지 않는다.
- tool availability가 있는데도 ask-vs-infer 원칙 없이 임의로 plain question만 반복하지 않는다.

## 출력 계약

`deep-interview`의 출력은 원본의 artifact-writing contract가 아니라, 현재 대화와 handoff를 위한 contract다.

최소 출력 항목:

- `Initial alignment snapshot`
- `Main alignment risks`
- `Alignment progress`
- `Locked brief`
- `What is in scope`
- `What is out of scope`
- `What may be decided autonomously`
- `What still needs confirmation`
- `Recommended handoff`
- `Residual risk`

## 성공 조건

다음이 만족되면 적응이 성공한 것으로 본다.

- greenfield proposal 질문이 들어왔을 때 단순 routing answer에서 끝나지 않는다.
- `deep-interview`가 필요한 경우 실제 질문 라운드가 시작된다.
- bounded choice 질문은 `request_user_input`으로 묻는다.
- specialist plugin은 requirements lock 이후에 붙는다.
- “원하시면 다음 턴에서…” 같은 deferral closing이 기본 종료 패턴이 되지 않는다.

## 비목표

- 원본 skill의 OMX runtime, CLI, artifact bridge를 재현하지 않는다.
- ambiguity scoring system 전체를 그대로 이식하지 않는다.
- deep-interview를 planning, autopilot, review-loop의 상위 meta-system으로 만들지 않는다.
- global 단계에서 specialist plugin을 병렬 first stop으로 복원하지 않는다.

## 구현 시 반영 대상

이 스펙을 구현할 때 같이 봐야 하는 파일:

- `workflow-kit-dev/skills/workflow-kit-dev-guide/SKILL.md`
- `workflow-kit-dev/skills/deep-interview/SKILL.md`
- `workflow-kit-dev/specs/plugin.md`
- 필요 시 `advance-codex/skills/tool-use-guide/SKILL.md`

## 현재 구조에서 우선 고쳐야 하는 것

1. `workflow-kit-dev-guide`가 proposal / direction evaluation / greenfield setup에서 `deep-interview`를 더 우선 고르게 할 것
2. guide가 handoff 전 advisory answer로 끝나지 않게 할 것
3. `deep-interview`에 `request_user_input` 우선 사용 규칙을 명시할 것
4. specialist plugin은 downstream handoff라는 점을 유지할 것
