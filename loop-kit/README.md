# Loop Kit

`loop-kit`는 Codex 작업 턴을 flow 단위로 유지하는 플러그인입니다.

- `flow`: `메시지 인터뷰 -> 플로우 설계 -> 메인 플로우 -> handoff condition`
- `turn-gate`: active turn continuity, phase prefix, verification, next-flow routing

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

Codex에서 `/plugins`를 열고 `Loop Kit`를 설치합니다.

## 업데이트

```sh
codex plugin marketplace upgrade
```

그다음 `/plugins`에서 기존 `Loop Kit` 설치를 삭제하고 다시 설치합니다.

## Flow

`flow`는 새 사용자 메시지를 실행 가능한 flow 구성으로 바꿉니다.

1. 메시지 인터뷰: intent snapshot, alignment risk, high-leverage question, answer pressure test, locked brief
2. 플로우 설계: active flow, parent flow, candidate, phase, handoff, artifact ownership, flow contract
3. 메인 플로우: `intake -> framing -> preparation -> work -> verification -> reporting`
4. handoff condition: result, verification, residual risk, next intake condition, commit-readiness

여러 flow가 필요하면 플로우 설계가 메인 플로우 후보를 만들고, 선택된 flow가 `intake`로 들어갑니다.
`reporting`에서 다음 flow가 준비되면 다음 `intake`로 라우팅합니다.

## Turn Gate

`turn-gate`는 현재 턴에서 `flow` 판단을 적용하고, 사용자의 explicit stop까지 턴을 유지합니다.

사용자-facing phase 시작 또는 의미 있는 진행 메시지는 다음 prefix를 사용합니다.

- `[intake]`
- `[framing]`
- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

Artifact, record, command output, question option label은 각 표면의 원래 형식을 유지합니다.
`reporting`은 다음 사용자 결정의 pre-intake 표면입니다.
Self-drive가 명시되면 준비된 sequence의 다음 flow 준비를 자체적으로 진행합니다.

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
$loop-kit:turn-gate 프론트엔드 리팩토링하자.
```

## 구조

```text
loop-kit/
  .codex-plugin/plugin.json
  README.md
  skills/
    flow/
    turn-gate/
```
