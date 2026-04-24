# Loop Kit Dev

`loop-kit-dev`은 하나의 작업 턴을 사용자가 명시적으로 종료할 때까지 유지하기 위한 Codex 플러그인입니다.

이 플러그인의 중심 표면은 `turn-gate` 하나입니다.
여러 loop skill을 사용자에게 직접 노출하지 않고, `turn-gate`가 현재 턴의 구조를 유지하면서 current phase에 맞는 내부 loop mode를 고르고, 결과 보고 전에 검증을 수행하고, 다음 플로우 선택지를 다시 엽니다.

> [!WARNING]
> Codex의 개발 중인 기능인 `default_mode_request_user_input`를 활성화해야 합니다.
> shell에서 다음 명령으로 활성화할 수 있습니다.
>
> ```sh
> codex features enable default_mode_request_user_input
> ```

## 설치 방법

먼저 이 저장소를 플러그인 마켓플레이스 source로 추가합니다.

```sh
codex plugin marketplace add opnay/agent-plugins
```

그다음 Codex에서 `/plugins`로 플러그인 목록을 열고 `Loop Kit Dev` 항목을 찾아 설치합니다.

한 번 설치하면 어느 위치에서 Codex를 실행해도 이 플러그인을 사용할 수 있습니다.

## 업데이트 방법

마켓플레이스 source를 최신 상태로 갱신합니다.

```sh
codex plugin marketplace upgrade
```

특정 marketplace만 갱신하려면 Codex에 표시되는 marketplace 이름을 붙여 실행합니다.
그다음 `/plugins`에서 기존 `Loop Kit Dev` 설치를 삭제하고 다시 설치하면 됩니다.

## 운영 방식

`loop-kit-dev`은 작업 흐름을 이어가기 위해 대상 저장소에 `.agents/sessions/` 폴더를 만들어 사용할 수 있습니다.
이 기록을 Git에 포함하지 않으려면 ignore 등록이 필요합니다.

기기 전역으로 제외하려면 `~/.config/git/ignore`에 다음 항목을 추가합니다.

```gitignore
.agents/sessions/
```

## 왜 필요한가

많은 에이전트 작업은 한 번의 답변으로 깨끗하게 끝나지 않습니다.
요구사항 확인, 구현, 검증, 리뷰 수정, 커밋 준비, 후속 선택이 같은 턴 안에서 이어집니다.
`loop-kit-dev`은 이 흐름을 명시적인 운영 계약으로 만들어, 에이전트가 상태 보고나 요약 뒤에 조용히 멈추지 않도록 합니다.

다음과 같은 작업에 적합합니다.

- 사용자가 멈추라고 할 때까지 턴을 계속 유지해야 하는 작업
- 분석, 계획, 작업, 검증, 결과 보고, 다음 플로우 선택이 드러나야 하는 작업
- 실행, 정제, 리뷰 처리, 커밋 준비 loop를 하나의 controller 안에서 골라야 하는 작업
- 사용자 선택이 필요한 지점에서는 질문 도구를 써야 하는 작업
- 필요하면 `self-drive`로 사용자 질문을 subagent 질문으로 바꿔 계속 진행해야 하는 작업

## 엔트리포인트

- `loop-kit-dev-guide`: 현재 요청이 `loop-kit-dev`으로 시작할 작업인지 판단합니다.
- `turn-gate`: 실제 작업을 진행하는 메인 controller입니다.

`turn-gate`가 호출되면, 현재 세션 동안 이 skill을 1급 운영 규칙으로 활성화한 것으로 취급합니다.

## 턴 구조

`turn-gate`는 다음 흐름을 계속 보이게 유지합니다.

1. 현재 입력을 분석합니다.
2. 실행 가능한 계획을 유지합니다.
3. 현재 작업을 수행합니다.
4. 작업 결과를 검증합니다.
5. 결과 또는 readiness 상태를 보고합니다.
6. 다음 플로우 선택지를 명시적으로 다시 엽니다.
7. 사용자가 종료를 요청하지 않으면 계속 진행합니다.

저장소가 해당 운영 방식을 사용한다면 `.agents/sessions/{YYYYMMDD}/` 아래에 세션 기록도 유지합니다.

## 내부 Loop Mode

사용자가 내부 loop mode를 직접 고를 필요는 없습니다.
`turn-gate`가 현재 blocker를 보고 다음 mode 중 하나를 선택합니다.

- `deep-interview`: 요구사항 확인, 불명확한 의도, scope boundary, approval line
- `autopilot`: 짧은 요청에서 검증된 결과까지 이어지는 broad end-to-end delivery
- `ralph-loop`: 작은 수정, 즉시 검증, 재평가가 필요한 bounded cycle
- `review-loop`: 리뷰 피드백이나 QA finding처럼 material issue만 좁게 처리해야 하는 흐름
- `commit-readiness-gate`: 변경 단위가 커밋으로 넘어갈 준비가 됐는지 확인하는 gate

이 mode들의 실행용 absorbed contract는 `skills/turn-gate/references/` 아래에 있습니다.
더 넓은 workflow taxonomy와 canonical loop-mode contract의 upstream SSOT는 `workflow-kit`이 계속 소유합니다.

## 질문 라우팅

`turn-gate`에는 current-phase mode와 별도로 질문 대상을 고르는 question-routing 축이 있습니다.

- `user-gated`: 기본값입니다. 선택지, scope lock, next-flow decision을 사용자 질문 도구로 묻습니다.
- `self-drive`: 사용자에게 묻던 질문을 subagent 질문 packet으로 바꿔, 사용자 개입 없이 loop를 계속 진행합니다.

`self-drive`에서는 일반적인 사용자 취향 누락을 중지 조건으로 보지 않습니다.
가능한 한 안전하고 되돌릴 수 있는 기본값을 가정으로 기록하고 계속 진행합니다.
멈춰야 하는 경우는 명시적 승인, 파괴적이거나 비가역적인 action 승인, 외부 action 승인, platform/tool/safety policy 경계처럼 hard boundary가 있는 경우로 제한합니다.

## 사용 예시

```text
$loop-kit-dev:turn-gate 프론트엔드 리팩토링하자.
```

```text
$loop-kit-dev:turn-gate self-drive; 마인크래프트와 비슷하게 샌드박스 RPG 게임 하나 만들자. 위치는 ~/Workspace/game으로 하자
```

## 플러그인 구조

```text
loop-kit-dev/
  .codex-plugin/plugin.json
  README.md
  specs/plugin.md
  specs/skills/
  skills/
    loop-kit-dev-guide/
    turn-gate/
```

## 설계 경계

`loop-kit-dev`은 의도적으로 작은 플러그인입니다.
broader workflow taxonomy, domain-specific implementation guidance, 무관한 agent utility를 소유하지 않습니다.
이 플러그인의 책임은 turn continuity, 내부 loop mode 선택, 결과 보고 전 검증, 명시적 next-flow reopening입니다.
