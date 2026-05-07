# Loop Kit

`loop-kit`은 하나의 작업 턴을 사용자가 명시적으로 종료할 때까지 유지하기 위한 Codex 플러그인입니다.

이 플러그인의 중심 표면은 `turn-gate` 하나입니다.
여러 loop skill을 사용자에게 직접 노출하지 않고, `turn-gate`가 현재 턴의 구조를 유지하면서 기본 flow를 `준비 -> 작업 -> 검증 -> 보고`로 이어갑니다. 사용자 메시지 기반 준비에서는 deep-interview alignment로 의도를 정렬하고 이후 flow list를 만들며, 이미 선택된 flow의 준비에서는 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인합니다.
사용자가 self-drive 진행을 원하면, 초기 준비에서 planned flow list 전체에 필요한 정보와 예상 위험 작업, approval boundary를 먼저 질문한 뒤 `turn-gate-self-drive`가 여러 flow를 이어서 수행합니다. 마지막 flow는 commit 실행이 아니라 commit-ready 보고로 끝나며, 실제 commit, push, PR, publish는 별도 승인 흐름으로 남깁니다.

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

그다음 Codex에서 `/plugins`로 플러그인 목록을 열고 `Loop Kit` 항목을 찾아 설치합니다.

한 번 설치하면 어느 위치에서 Codex를 실행해도 이 플러그인을 사용할 수 있습니다.

## 업데이트 방법

마켓플레이스 source를 최신 상태로 갱신합니다.

```sh
codex plugin marketplace upgrade
```

특정 marketplace만 갱신하려면 Codex에 표시되는 marketplace 이름을 붙여 실행합니다.
그다음 `/plugins`에서 기존 `Loop Kit` 설치를 삭제하고 다시 설치하면 됩니다.

## 운영 방식

`loop-kit`은 작업 흐름을 이어가기 위해 대상 저장소에 `.agents/sessions/` 폴더를 만들어 사용할 수 있습니다.
이 기록을 Git에 포함하지 않으려면 ignore 등록이 필요합니다.

기기 전역으로 제외하려면 `~/.config/git/ignore`에 다음 항목을 추가합니다.

```gitignore
.agents/sessions/
```

## 왜 필요한가

많은 에이전트 작업은 한 번의 답변으로 깨끗하게 끝나지 않습니다.
요구사항 확인, 구현, 검증, 리뷰 수정, 커밋 준비, 후속 선택이 같은 턴 안에서 이어집니다.
`loop-kit`은 이 흐름을 명시적인 운영 계약으로 만들어, 에이전트가 상태 보고나 요약 뒤에 조용히 멈추지 않도록 합니다.

다음과 같은 작업에 적합합니다.

- 사용자가 멈추라고 할 때까지 턴을 계속 유지해야 하는 작업
- 준비, 작업, 검증, 보고, 다음 플로우 선택이 드러나야 하는 작업
- 사용자 메시지에서는 deep-interview로 의도를 정렬하고 flow list를 만들어야 하는 작업
- 초기 준비에서 필요한 정보를 모은 뒤 여러 flow를 self-drive로 이어가고 마지막에 commit-ready 보고를 받아야 하는 작업
- 이미 선택된 flow에서는 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인해야 하는 작업
- 실행, 정제, 리뷰 처리, 커밋 준비 loop를 하나의 controller 안에서 골라야 하는 작업
- 사용자 선택이 필요한 지점에서는 질문 도구를 써야 하는 작업
- 읽기 전용 clean-context verifier subagent를 flow 검증 단계의 일부로 안정적으로 사용해야 하는 작업
- 필요하면 `turn-gate-self-drive`로 bounded decision을 subagent question packet에 라우팅해야 하는 작업

## 엔트리포인트

- `turn-gate`: 실제 작업을 진행하는 메인 controller입니다.
- `turn-gate-self-drive`: `turn-gate`를 먼저 적용한 뒤 blocked question을 subagent question packet으로 라우팅하는 overlay입니다.
- `turn-gate-self-drive`: 초기 준비가 끝난 planned flow list를 협의된 경계 안에서 이어가고 마지막 flow에서 commit-ready 보고를 남기는 overlay입니다.

`turn-gate`가 호출되면, 현재 세션 동안 이 skill을 1급 운영 규칙으로 활성화한 것으로 취급합니다.
이 규칙은 skill body의 `Important` 섹션에서 먼저 드러나며, 결과 보고만으로 턴을 닫지 않고 다음 플로우 질문을 다시 여는 동작을 우선 계약으로 둡니다.

## 턴 구조

`turn-gate`는 다음 흐름을 계속 보이게 유지합니다.

1. 준비: 사용자 메시지는 deep-interview alignment와 flow list로 정렬하고, 이미 선택된 flow는 현재 상태와 작업 범위를 확인합니다.
2. 작업: 현재 flow가 소유한 실제 작업을 수행합니다.
3. 검증: 읽기 전용 bounded verifier subagent가 수정 결과, 타입/테스트/파싱 신호, 또는 조사 결과의 논리적 취약점을 확인합니다.
4. 보고: 이번 flow의 맥락을 정리하고 다음 flow 선택지를 명시적으로 다시 엽니다.
5. 사용자가 종료를 요청하지 않으면 다음 flow의 준비로 계속 진행합니다.

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
`workflow-kit`은 각 workflow skill의 일반 의미를 제공하지만, turn-gate runtime contract와 session continuity는 `loop-kit`이 직접 소유합니다.

## 질문 라우팅

`turn-gate`는 기본적으로 user-gated question routing을 사용합니다.

- `turn-gate`: 선택지, scope lock, next-flow decision을 사용자 질문 도구로 묻습니다.
- `turn-gate-self-drive`: bounded decision을 subagent question packet으로 라우팅해, 사용자 개입 없이 loop를 계속 진행합니다.

`turn-gate-self-drive`에서는 일반적인 사용자 취향 누락을 중지 조건으로 보지 않습니다.
가능한 한 안전하고 되돌릴 수 있는 기본값을 가정으로 기록하고 계속 진행합니다.
초기 준비에서 예상되는 위험 작업은 미리 질문해 승인/비승인 또는 handoff 경계를 계획합니다.
멈춰야 하는 경우는 초기 협의 범위 밖의 위험 작업, 새 명시적 승인, 파괴적이거나 비가역적인 action 승인, 외부 action 승인, platform/tool/safety policy 경계처럼 새 hard boundary가 나타나는 경우로 제한합니다.
준비된 여러 flow를 self-drive로 진행할 때는 각 flow의 work boundary, non-goal, verification expectation을 다시 확인하고, planned flow list의 마지막에는 commit-readiness gate를 실행해 commit-ready 상태를 보고합니다.

## 사용 예시

```text
$loop-kit:turn-gate 프론트엔드 리팩토링하자.
```

```text
$loop-kit:turn-gate-self-drive 마인크래프트와 비슷하게 샌드박스 RPG 게임 하나 만들자. 위치는 ~/Workspace/game으로 하자
```

## 플러그인 구조

```text
loop-kit/
  .codex-plugin/plugin.json
  README.md
  skills/
    turn-gate/
    turn-gate-self-drive/
```

## 설계 경계

`loop-kit`은 의도적으로 작은 플러그인입니다.
broader workflow taxonomy, domain-specific implementation guidance, 무관한 agent utility를 소유하지 않습니다.
이 플러그인의 책임은 turn continuity, 내부 loop mode 선택, 결과 보고 전 검증, 명시적 next-flow reopening입니다.
