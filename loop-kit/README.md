# Loop Kit

`loop-kit`은 하나의 작업 턴을 사용자가 명시적으로 종료할 때까지 유지하기 위한 Codex 플러그인입니다.

이 플러그인의 중심 표면은 `flow`와 `turn-gate`입니다.
`flow`는 하나의 메시지나 동작을 흐름 단위로 해석하고, 필요하면 finite `sub-flow candidates`로 나눕니다.
또한 active flow의 각 phase가 시작되거나 끝날 때 `000-plan.md` 또는 active flow record 중 어느 기록을 갱신해야 하는지 판단하는 checkpoint를 둡니다.
`turn-gate`는 현재 턴의 구조를 유지하면서 모든 작업이 `flow` 계약을 통과하게 하고, flow가 끝난 뒤 다음 flow 선택지를 엽니다.
active flow 도중 새 사용자 메시지가 들어오면 `turn-gate`는 `interruption` 진입점으로 먼저 분류해 질문 답변, 현재 flow 개정, background 전환, 후속 예약, flow 대체, blocker, explicit stop 중 하나로 라우팅합니다.
초기 의도 정렬과 sub-flow 후보 설계 자체는 session plan과 approval boundary를 소유하는 운영 flow가 될 수 있고, 그 결과 만들어지는 실행 flow는 코드, 문서, fixture, 설정 같은 산출물 변경 단위로 구분합니다.
이때 flow는 `분석`, `작업`, `검증`, `커밋 준비` 같은 진행 단계가 아니라 함께 검토하고 검증하고 필요하면 커밋할 수 있는 응집된 변경 단위입니다.
또한 flow는 반드시 최종 사용자에게 직접 보이는 가치 단위일 필요가 없습니다.
예를 들어 "로그인 페이지 만들기"는 하나의 사용자 가치처럼 보일 수 있지만, 실제 sub-flow 후보는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립`처럼 보이지 않는 준비성 변경을 포함해 나뉠 수 있습니다.
`turn-gate`는 기본 turn-level gate로 독립적으로 동작합니다. 사용자가 self-drive 진행을 원하면, self-drive overlay 계약이 준비된 flow sequence의 진행 판단을 덮어씁니다.
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
- 각 phase 시작과 종료에서 plan 또는 flow record가 현재 상태를 재구성할 수 있어야 하는 작업
- intake에서 사용자 입력과 목표를 정리하고, framing에서 sub-flow 후보를 만든 뒤, preparation에서 선택된 flow의 readiness를 잠가야 하는 작업
- 초기 의도 정렬과 sub-flow 후보 설계가 session plan 산출물로 남아야 하는 작업
- sub-flow 후보가 phase checklist가 아니라 검토/검증/커밋 가능한 변경 단위로 나뉘어야 하는 작업
- intake/framing/preparation에서 필요한 정보를 모은 뒤 self-drive overlay로 여러 flow를 이어가야 하는 작업
- 이미 선택된 flow에서는 수정 범위, 현재 상태, 대상 파일, 검증 조건을 먼저 확인해야 하는 작업
- 실행, 정제, 리뷰 처리, 커밋 준비 handoff를 active flow 안의 strategy로 골라야 하는 작업
- 사용자 선택이 필요한 지점에서는 질문 도구를 써야 하는 작업
- 작업 위험도에 따라 `clean-context`, `normal`, `not-required` verification method를 구분해야 하는 작업
- 파일 변경, release surface, 다중 파일 계약, 실패 이력, approval-sensitive action에서는 clean-context verifier를 기본값으로 유지해야 하는 작업
- 필요하면 self-drive overlay로 bounded decision을 subagent question packet에 라우팅해야 하는 작업

## 엔트리포인트

- `flow`: 메시지나 동작을 flow로 해석하고, intake, framing, preparation readiness, parent flow, sub-flow candidate, operational-preparation flow, change-unit flow, flow-vs-phase 경계, discovery, flow-local strategy, handoff condition을 판정합니다.
- `flow`: phase 시작/종료 record checkpoint를 산출해 `000-plan.md`와 active flow record 중 무엇을 갱신해야 하는지 구분합니다.
- `turn-gate`: 현재 턴에서 flow 사용을 강제하고, flow reporting 뒤 다음 flow 질문을 여는 turn-level gate입니다.

`turn-gate`가 호출되면, 현재 세션 동안 이 skill을 1급 운영 규칙으로 활성화한 것으로 취급합니다.
이 규칙은 skill body의 `Important` 섹션에서 먼저 드러나며, 결과 보고만으로 턴을 닫지 않고 다음 플로우 질문을 다시 여는 동작을 우선 계약으로 둡니다.

## 턴 구조

`turn-gate`는 다음 흐름을 계속 보이게 유지합니다.

1. Intake: 초기 요청의 원문과 해석을 분리하고 goal, non-goal, authority-sensitive signal, discovery topic을 확인합니다.
2. Framing: active flow, parent flow, sub-flow candidate, phase, handoff를 구분하고 산출물 소유권과 후보/선택 상태를 정리합니다.
3. Preparation: 선택된 active flow의 scope, readiness, verification expectation, approval boundary, handoff condition을 잠급니다.
4. 작업: 현재 flow가 소유한 실제 작업을 수행합니다.
5. 검증: 작업 위험도에 맞춰 `clean-context`, `normal`, `not-required` method 중 하나로 검증합니다.
6. 보고: 이번 flow의 맥락을 정리하고 다음 flow 선택지를 명시적으로 다시 엽니다.
7. 사용자가 종료를 요청하지 않으면 다음 flow의 intake로 계속 진행합니다.

active flow 중 들어온 새 사용자 메시지는 위 lifecycle을 대체하지 않고 `interruption`으로 잠시 분류됩니다. 질문이 계약을 바꾸지 않으면 답변 후 이전 phase로 돌아가고, scope나 non-goal을 바꾸면 현재 flow를 개정한 뒤 `framing` 또는 `preparation`으로 돌아갑니다. 다른 작업이 먼저 필요하면 현재 flow를 background로 두고 새 foreground flow를 시작하며, 나중에 볼 주제는 후속 후보로 예약합니다.

각 flow는 위 1-6단계를 내부에 모두 가집니다.
따라서 `분석`, `작업`, `검증`을 서로 다른 flow로 나누지 않습니다.
flow를 나눌 때는 독립적으로 이해하고 리뷰하고 검증하고 커밋할 수 있는 변경 묶음인지 봅니다.
단, 초기 의도 정렬을 통해 sub-flow 후보를 만드는 앞단은 운영 flow로 기록할 수 있으며, 이 운영 flow는 실제 제품 변경 flow와 섞지 않습니다.
sub-flow 후보 생성은 실행이 아니며, `turn-gate` 질문 또는 준비된 self-drive sequence를 통해 선택될 때만 active flow가 됩니다.

저장소가 해당 운영 방식을 사용한다면 `.agents/sessions/{YYYYMMDD}/` 아래에 세션 기록도 유지합니다.

## 검증 방식

`turn-gate`는 검증 결과 상태와 검증 방법을 구분합니다.
결과 상태는 `pass`, `fail`, `blocked`, `insufficient`처럼 보고 가능 여부를 나타냅니다.
검증 방법은 아래 셋 중 하나입니다.

- `clean-context`: 읽기 전용 bounded verifier subagent가 독립 context에서 검증합니다. 파일 변경, release surface, manifest/template/scenario fixture/build output, 여러 파일 사이 계약, 실패 이력, 사용자 요청 검증, approval-sensitive action에서는 기본값입니다.
- `normal`: 낮은 위험의 no-edit/read-only 작업에서 command/check, source readback, evidence checklist, 논리 반례 검토를 같은 context에서 수행하고 근거를 기록합니다.
- `not-required`: activation-only, next-flow selection, blocker-before-work처럼 검증할 work output이 없을 때만 사용합니다. 이 경우에도 이유와 남은 불확실성을 기록합니다.

`not-required`는 성공 상태가 아니며, commit/push/PR/publish/release/version bump 같은 승인 민감 작업을 경량화하지 않습니다.

## Flow Strategy

사용자가 flow strategy를 직접 고를 필요는 없습니다.
`flow`는 active flow의 contract와 blocker에 맞는 strategy를 산출하고, `turn-gate`는 그 결과를 적용합니다.

- `discovery`: 요구사항 확인, 불명확한 의도, scope boundary, approval line을 다루는 flow-local strategy이며 주로 intake에서 사용
- `broad-execution`: locked scope 안에서 검증된 결과까지 이어지는 단일 flow execution strategy
- `fix-verify-loop`: 작은 수정, 즉시 검증, 재평가가 필요한 bounded cycle strategy
- `review-loop`: 리뷰 피드백이나 QA finding처럼 material issue를 좁게 처리하는 strategy
- `commit-readiness`: 변경 단위가 커밋으로 넘어갈 준비가 됐는지 확인하는 handoff strategy

Self-drive overlay의 상세 조건은 `skills/turn-gate/references/self-drive.md`가 소유합니다.

## 질문 라우팅

`turn-gate`는 기본적으로 user-gated question routing을 사용합니다.

- `turn-gate`: 선택지, scope lock, next-flow decision을 사용자 질문 도구로 묻습니다.
- self-drive overlay: bounded decision을 subagent question packet으로 라우팅해, 사용자 개입 없이 준비된 sequence를 계속 진행합니다.

Self-drive의 중지 조건, handoff, commit-readiness 이후 동작은 self-drive 계약을 따릅니다.

## 사용 예시

```text
$loop-kit:turn-gate 프론트엔드 리팩토링하자.
```

## 플러그인 구조

아래 트리는 플러그인 source shape를 보여줍니다.
이 저장소에서 실제 개발 원본은 `src/<plugin-name>-dev/`에 있고, 루트 `<plugin-name>/`은 build command가 갱신하는 release surface입니다.
일반 편집은 `src/<plugin-name>-dev/`에 적용하고, 루트 `<plugin-name>/`은 직접 편집하지 않습니다.

```text
loop-kit/
  .codex-plugin/plugin.json
  README.md
  skills/
    turn-gate/
    flow/
```

## 설계 경계

`loop-kit`은 의도적으로 작은 플러그인입니다.
broader workflow taxonomy, domain-specific implementation guidance, 무관한 agent utility를 소유하지 않습니다.
이 플러그인의 책임은 flow 경계, sub-flow 후보, flow-local strategy, turn continuity, risk-based verification method, 결과 보고 전 검증 판단, 명시적 next-flow reopening입니다.
