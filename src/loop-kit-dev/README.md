# Loop Kit Dev

`loop-kit-dev`는 Codex 작업 턴을 flow 단위로 유지하는 플러그인입니다.

- `deep-interview`: intent, scope, tradeoff, approval boundary를 질문과 압력 테스트로 잠그는 인터뷰 스킬
- `flow`: `entry skill reconfigure -> deep-interview -> 플로우 설계 -> 메인 플로우 -> 메인 플로우 회고 -> handoff condition`
- `turn-gate`: active turn continuity, flow wrapper, handoff question routing, self-drive gate, explicit stop

의존 방향은 `deep-interview -> flow -> turn-gate`입니다. 상위 skill은 하위 skill을 적용하지만, 하위 skill은 상위 skill 맥락을 전제로 하지 않습니다.

> [!WARNING]
> Codex의 개발 중인 기능인 `default_mode_request_user_input`를 활성화해야 합니다.
>
> ```sh
> codex features enable default_mode_request_user_input
> ```

## 설치

```sh
codex plugin marketplace add opnay/agent-plugins
```

Codex에서 `/plugins`를 열고 `Loop Kit Dev`를 설치합니다.

## 업데이트

```sh
codex plugin marketplace upgrade
```

그다음 `/plugins`에서 기존 `Loop Kit Dev` 설치를 삭제하고 다시 설치합니다.

## Flow

`flow`는 새 사용자 메시지를 실행 가능한 flow 구성으로 바꿉니다.
모든 사용자 메시지는 같은 경로를 탑니다.

1. skill reconfigure: flow entry와 post-reporting continuation boundary에서 필요한 skill 본문을 source에서 다시 읽어 루프 중 잊힌 skill context를 복구
2. deep-interview: locked execution brief
3. 플로우 설계: active flow, parent flow, candidate, phase, handoff, artifact ownership, flow 계약 검증
4. 메인 플로우: `intake -> framing -> preparation -> work -> verification -> reporting`
5. 메인 플로우 회고: 항상 `000-review.md`를 갱신하고, finding이 없으면 no-finding 결과로 짧게 남깁니다.
6. handoff condition: result, verification, residual risk, next intake condition, commit-readiness

사용자-facing 진행 메시지는 현재 phase label을 사용할 수 있고, artifact, 기록, command summary, 질문 option label에는 label을 전파하지 않습니다.
여러 flow가 필요하면 플로우 설계가 메인 플로우 후보를 만들고, 선택된 flow가 `intake`로 들어갑니다.
`reporting`과 필요한 회고 뒤 다음 flow가 준비되면 `reporting` 직후 skill reconfigure를 수행하고 다음 `intake`로 라우팅합니다.

## Turn Gate

`turn-gate`는 active turn에 `flow` 판단을 적용하고, 사용자의 explicit stop까지 턴을 유지합니다.

`flow skill: handoff` 뒤에는 `질문 도구: 다음 플로우 선택`으로 다음 flow 입력을 고릅니다.
`next-flow gate`에서 예약 플로우를 재검토하고 다음 flow를 선택한 뒤 `000-plan.md`를 매번 업데이트합니다. 필요한 skill reread는 flow entry 또는 post-reporting continuation boundary의 skill reconfigure가 수행합니다.
질문 도구는 `deep-interview`로 입력을 구체화한 뒤 다시 `flow`로 들어갑니다.
Self-drive가 명시되면 그래프 노드가 아니라 준비된 sequence gate가 질문 도구를 대체합니다.
Record, verification, interruption, date 처리는 메인 그래프 노드가 아니라 active turn을 복구하고 안전하게 라우팅하기 위한 지원 계약입니다.
사용자-facing 진행 메시지는 source skill이 소유한 phase prefix로 현재 단계를 드러냅니다. `turn-gate`는 `flow` phase label을 재정의하지 않고 적용하며, `next-flow gate`에서는 `[next-flow]`를 소유합니다. Artifact, 기록, command summary, 질문 option label에는 prefix를 전파하지 않습니다.

## 검증

`turn-gate`는 검증 방법과 결과 상태를 구분합니다.

- `clean-context`: 파일 변경, release surface, 다중 파일 계약, 실패 이력, approval-sensitive action
- `normal`: 낮은 위험의 read-only 또는 no-edit work
- `not-required`: activation-only, routing-only, blocker-before-work

결과 상태는 `pass`, `fail`, `blocked`, `insufficient`로 보고합니다.

## 운영 기록

대상 저장소에 `.agents/sessions/` 기록을 만들 수 있습니다.
Git 전역 ignore를 쓰려면 `~/.config/git/ignore`에 추가합니다.

```gitignore
.agents/sessions/
```

## 사용 예시

```text
$loop-kit-dev:turn-gate 프론트엔드 리팩토링하자.
```

## 구조

```text
loop-kit-dev/
  .codex-plugin/plugin.json
  README.md
  specs/plugin.md
  specs/skills/
  skills/
    deep-interview/
    flow/
    turn-gate/
```
