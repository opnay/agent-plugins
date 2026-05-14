# Loop Kit Dev

`loop-kit-dev`은 하나의 작업 턴을 사용자가 명시적으로 종료할 때까지 유지하기 위한 Codex 플러그인입니다.

이 플러그인의 중심 표면은 `turn-gate` 하나입니다.
여러 loop skill을 사용자에게 직접 노출하지 않고, `turn-gate`가 현재 턴의 구조를 유지하면서 기본 flow를 `준비 -> 작업 -> 검증 -> 보고`로 이어갑니다. 초기 준비에서는 deep-interview alignment로 의도를 정렬하고 이후 flow list를 만들며, 이미 선택된 flow의 준비에서는 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인합니다.
초기 의도 정렬과 flow list 설계 자체는 session plan과 approval boundary를 소유하는 운영 flow가 될 수 있고, 그 결과 만들어지는 실행 flow는 코드, 문서, fixture, 설정 같은 산출물 변경 단위로 구분합니다.
이때 flow는 `분석`, `작업`, `검증`, `커밋 준비` 같은 진행 단계가 아니라 함께 검토하고 검증하고 필요하면 커밋할 수 있는 응집된 변경 단위입니다.
또한 flow는 반드시 최종 사용자에게 직접 보이는 가치 단위일 필요가 없습니다.
예를 들어 "로그인 페이지 만들기"는 하나의 사용자 가치처럼 보일 수 있지만, 실제 planned flow는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립`처럼 보이지 않는 준비성 변경을 포함해 나뉠 수 있습니다.
`turn-gate`는 기본 loop controller로 독립적으로 동작합니다. 사용자가 self-drive 진행을 원하면, self-drive overlay 계약이 준비된 sequence의 진행 판단을 덮어씁니다.
최종 QA, 정합성 점검, 검증 결과 보고, commit-ready 보고는 별도 산출물 변경이 없다면 flow가 아니라 각 flow의 검증/보고 또는 handoff입니다.

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
- 준비, 작업, 검증, 보고, 다음 플로우 선택이 드러나야 하는 작업
- 초기 준비에서 deep-interview로 의도를 정렬하고 flow list를 만들어야 하는 작업
- 초기 의도 정렬과 planned flow list 설계가 session plan 산출물로 남아야 하는 작업
- flow list가 phase checklist가 아니라 검토/검증/커밋 가능한 변경 단위로 나뉘어야 하는 작업
- 초기 준비에서 필요한 정보를 모은 뒤 self-drive overlay로 여러 flow를 이어가야 하는 작업
- 이미 선택된 flow에서는 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인해야 하는 작업
- 실행, 정제, 리뷰 처리, 커밋 준비 loop를 하나의 controller 안에서 골라야 하는 작업
- 사용자 선택이 필요한 지점에서는 질문 도구를 써야 하는 작업
- 작업 위험도에 따라 `clean-context`, `normal`, `not-required` verification method를 구분해야 하는 작업
- 파일 변경, release surface, 다중 파일 계약, 실패 이력, approval-sensitive action에서는 clean-context verifier를 기본값으로 유지해야 하는 작업
- 필요하면 self-drive overlay로 bounded decision을 subagent question packet에 라우팅해야 하는 작업

## 엔트리포인트

- `turn-gate`: 실제 작업을 진행하는 메인 controller입니다.

`turn-gate`가 호출되면, 현재 세션 동안 이 skill을 1급 운영 규칙으로 활성화한 것으로 취급합니다.
이 규칙은 skill body의 `Important` 섹션에서 먼저 드러나며, 결과 보고만으로 턴을 닫지 않고 다음 플로우 질문을 다시 여는 동작을 우선 계약으로 둡니다.

## 턴 구조

`turn-gate`는 다음 흐름을 계속 보이게 유지합니다.

1. 준비: 초기 요청은 deep-interview alignment와 flow list로 정렬하고, 이미 선택된 flow는 현재 상태와 작업 범위를 확인합니다.
2. 작업: 현재 flow가 소유한 실제 작업을 수행합니다.
3. 검증: 작업 위험도에 맞춰 `clean-context`, `normal`, `not-required` method 중 하나로 검증합니다.
4. 보고: 이번 flow의 맥락을 정리하고 다음 flow 선택지를 명시적으로 다시 엽니다.
5. 사용자가 종료를 요청하지 않으면 다음 flow의 준비로 계속 진행합니다.

각 flow는 위 1-4단계를 내부에 모두 가집니다.
따라서 `분석`, `작업`, `검증`을 서로 다른 flow로 나누지 않습니다.
flow를 나눌 때는 독립적으로 이해하고 리뷰하고 검증하고 커밋할 수 있는 변경 묶음인지 봅니다.
단, 초기 의도 정렬을 통해 planned flow list를 만드는 앞단은 운영 flow로 기록할 수 있으며, 이 운영 flow는 실제 제품 변경 flow와 섞지 않습니다.

저장소가 해당 운영 방식을 사용한다면 `.agents/sessions/{YYYYMMDD}/` 아래에 세션 기록도 유지합니다.

## 검증 방식

`turn-gate`는 검증 결과 상태와 검증 방법을 구분합니다.
결과 상태는 `pass`, `fail`, `blocked`, `insufficient`처럼 보고 가능 여부를 나타냅니다.
검증 방법은 아래 셋 중 하나입니다.

- `clean-context`: 읽기 전용 bounded verifier subagent가 독립 context에서 검증합니다. 파일 변경, release surface, manifest/template/scenario fixture/build output, 여러 파일 사이 계약, 실패 이력, 사용자 요청 검증, approval-sensitive action에서는 기본값입니다.
- `normal`: 낮은 위험의 no-edit/read-only 작업에서 command/check, source readback, evidence checklist, 논리 반례 검토를 같은 context에서 수행하고 근거를 기록합니다.
- `not-required`: activation-only, next-flow selection, blocker-before-work처럼 검증할 work output이 없을 때만 사용합니다. 이 경우에도 이유와 남은 불확실성을 기록합니다.

`not-required`는 성공 상태가 아니며, commit/push/PR/publish/release/version bump 같은 승인 민감 작업을 경량화하지 않습니다.

## Phase Protocol

사용자가 phase protocol을 직접 고를 필요는 없습니다.
`turn-gate`는 기본 상태로 동작하고, 현재 blocker에 맞는 phase protocol을 적용합니다.

- `deep-interview`: 요구사항 확인, 불명확한 의도, scope boundary, approval line을 다루는 phase protocol
- `autopilot`: 검증된 결과까지 이어지는 broad end-to-end delivery phase protocol
- `ralph-loop`: 작은 수정, 즉시 검증, 재평가가 필요한 bounded cycle phase protocol
- `review-loop`: 리뷰 피드백이나 QA finding처럼 material issue를 좁게 처리하는 phase protocol
- `commit-readiness-gate`: 변경 단위가 커밋으로 넘어갈 준비가 됐는지 확인하는 phase protocol

이 계약들의 실행용 absorbed contract는 `skills/turn-gate/references/` 아래에 있습니다.
`workflow-kit`은 각 workflow skill의 일반 의미를 제공하지만, turn-gate runtime contract와 session continuity는 `loop-kit`이 직접 소유합니다.
Self-drive overlay의 상세 조건은 `skills/turn-gate/references/self-drive.md`가 소유합니다.

## 질문 라우팅

`turn-gate`는 기본적으로 user-gated question routing을 사용합니다.

- `turn-gate`: 선택지, scope lock, next-flow decision을 사용자 질문 도구로 묻습니다.
- self-drive overlay: bounded decision을 subagent question packet으로 라우팅해, 사용자 개입 없이 준비된 sequence를 계속 진행합니다.

Self-drive의 중지 조건, handoff, commit-readiness 이후 동작은 self-drive 계약을 따릅니다.

## 사용 예시

```text
$loop-kit-dev:turn-gate 프론트엔드 리팩토링하자.
```

## 플러그인 구조

```text
loop-kit-dev/
  .codex-plugin/plugin.json
  README.md
  specs/plugin.md
  specs/skills/
  skills/
    turn-gate/
```

## 설계 경계

`loop-kit-dev`은 의도적으로 작은 플러그인입니다.
broader workflow taxonomy, domain-specific implementation guidance, 무관한 agent utility를 소유하지 않습니다.
이 플러그인의 책임은 turn continuity, phase protocol 선택, risk-based verification method, 결과 보고 전 검증 판단, 명시적 next-flow reopening입니다.
