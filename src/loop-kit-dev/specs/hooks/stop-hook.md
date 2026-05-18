# bundled Stop hook plugin spec

## 목적

이 문서는 `loop-kit-dev`가 설치 단위로 제공하는 bundled Codex `Stop` hook의 packaging, manifest 연결, runtime 파일 위치, 실행 계약을 소유합니다.
이 hook은 `turn-gate` skill을 보강하는 plugin-level runtime surface입니다.
skill spec은 hook을 어떻게 안내하고 해석할지만 소유하고, 설치 단위 hook 자체의 위치와 packaging 계약은 이 spec을 기준으로 합니다.

## 경계

- 포함:
  - plugin runtime hook file layout
  - `.codex-plugin/plugin.json`의 `hooks` field 연결
  - build/release surface 포함 기준
  - Codex `Stop` event payload 처리 경계
  - block output shape
  - no-write/no-approval boundary
- 제외:
  - `turn-gate` skill body의 문구 구성
  - `turn-gate` phase, record, next-flow lifecycle 자체
  - global Codex config 설치
  - `/hooks` trust/reload UI 조작
  - plugin hook feature enablement
  - repo-local `.codex/hooks.*` experiment 또는 local override 관리

## Plugin Runtime Files

Bundled Stop hook implementation은 설치되는 plugin runtime surface 안에 둡니다.

- canonical plugin hook config: `hooks/hooks.json`
- canonical plugin hook script: `hooks/turn_gate_stop.py`

`hooks/hooks.json`의 command는 plugin install 위치에 독립적이어야 하므로 `${PLUGIN_ROOT}`를 기준으로 script를 호출합니다.

```json
{
  "type": "command",
  "command": "/usr/bin/python3 \"${PLUGIN_ROOT}/hooks/turn_gate_stop.py\""
}
```

`.codex/hooks.json`와 `.codex/hooks/*`는 repo-local experiment 또는 local override 위치일 수 있지만, plugin install surface가 아닙니다.
플러그인으로 배포되는 Stop hook의 canonical 위치로 쓰면 안 됩니다.

## Manifest And Build Contract

`.codex-plugin/plugin.json`은 top-level `hooks` field로 bundled hook config를 가리켜야 합니다.

```json
{
  "hooks": "./hooks/hooks.json"
}
```

build/release script는 `hooks/` runtime root를 release surface에 포함해야 합니다.
release `loop-kit/`에는 `hooks/hooks.json`와 `hooks/turn_gate_stop.py`가 포함되어야 하고, `specs/`는 포함되면 안 됩니다.

Plugin hooks는 Codex에서 별도 feature/trust 조건을 갖습니다.
`[features].plugin_hooks = true`와 plugin hook trust/review가 끝나지 않으면 bundled hook이 설치되어 있어도 실행되지 않을 수 있습니다.

## Hook Runtime Contract

Stop hook은 Codex `Stop` event payload만 처리합니다.

다음 경우에는 조용히 종료해야 합니다.

- `hook_event_name`이 `Stop`이 아닌 경우
- `stop_hook_active`가 true인 재진입 payload
- payload에서 current Codex session id를 확인할 수 없는 경우
- 현재 cwd에서 같은 날짜 `.agents/sessions/{YYYYMMDD}/000-plan.md`를 찾을 수 없는 경우
- plan에 `active_flow`가 없는 경우
- active flow record를 읽을 수 없거나 frontmatter가 없는 경우
- active flow record 또는 plan에 `turn_gate_session_id`가 없는 경우
- payload의 current Codex session id가 record의 `turn_gate_session_id`와 일치하지 않는 경우
- `turn_gate_active`가 true가 아닌 경우
- `terminal_summary_allowed`, `user_explicit_stop`, 또는 `confirmed_closure`가 true인 경우
- `required_next_action`이 비어 있는 경우
- 마지막 assistant message가 이미 next-flow나 user-gated question으로 라우팅하는 경우

위 quiet-exit 조건에 걸리지 않으면 terminal closure를 차단합니다.
즉, plugin 설치 자체는 Stop block authority가 아니며, 현재 main turn이 `turn-gate`로 armed된 session id와 일치해야만 차단할 수 있습니다.

출력은 JSON 한 줄이어야 합니다.

```json
{"decision":"block","reason":"turn-gate is active and terminal closure is not allowed. Refresh the active flow record, then continue to the required next action: <required_next_action>"}
```

`reason`은 다음 내용을 포함해야 합니다.

- turn-gate가 active이고 terminal closure가 허용되지 않았다는 점
- current Codex session id가 active turn-gate record의 `turn_gate_session_id`와 일치한다는 전제
- active flow record를 refresh하라는 지시
- active flow의 `required_next_action`

## Side-Effect Boundary

Stop hook은 파일을 쓰지 않습니다.
hook은 current state를 읽고 block 여부만 결정합니다.
hook이 block한 뒤 active flow record와 plan을 갱신하는 책임은 main agent에게 있습니다.

Stop hook은 다음 권한을 만들지 않습니다.

- destructive action approval
- external action approval
- commit/push/PR/publish/release/version bump approval
- plugin hook feature activation
- global Codex config 변경

## Skill Relationship

`turn-gate` skill은 이 bundled plugin hook을 optional runtime backstop으로 설명할 수 있습니다.
그 skill-level guidance는 `specs/skills/turn-gate/hooks/stop-hook.md`가 소유합니다.
skill guidance는 이 plugin hook spec의 packaging 계약을 반복 소유하지 않고, `turn-gate`가 hook block을 어떻게 해석하고 reporting/record refresh/next-flow reopening 의무를 유지해야 하는지에 집중합니다.

## 검토 질문

- Stop hook code 위치가 plugin runtime `hooks/turn_gate_stop.py`로 일관되는가?
- manifest가 `hooks` field로 `./hooks/hooks.json`를 가리키는가?
- build output release surface에 `hooks/hooks.json`와 `hooks/turn_gate_stop.py`가 포함되는가?
- release surface에서 `specs/`가 빠지는가?
- Stop hook이 기록을 직접 수정하지 않는가?
- block 조건이 Continuity Guard 필드와 `required_next_action`을 기준으로 하는가?
- block 조건이 payload session id와 record `turn_gate_session_id` 일치를 요구하는가?
- payload나 record에 session id가 없을 때 fail-open으로 조용히 종료하는가?
- 재진입 payload에서 조용히 종료하는가?
