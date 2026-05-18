# bundled SessionStart hook plugin spec

## 목적

이 문서는 `loop-kit-dev`가 설치 단위로 제공하는 bundled Codex `SessionStart` hook의 packaging, runtime 파일 위치, 실행 계약을 소유합니다.
이 hook은 Codex 세션 시작 또는 재개 시 `.agents/sessions/`에 남아 있는 turn-gate plan/flow 상태를 읽어 agent에게 startup context를 제공하는 plugin-level runtime surface입니다.
자동 continuation, approval, record write를 만들지 않습니다.

## 경계

- 포함:
  - plugin runtime hook file layout
  - `hooks/hooks.json`의 `SessionStart` matcher와 handler 연결
  - plan/flow discovery order
  - additionalContext output shape
  - no-write/no-approval boundary
- 제외:
  - `turn-gate` skill body의 세부 문구 구성
  - session record 생성, 복구, 수정
  - global Codex config 설치
  - `/hooks` trust/reload UI 조작
  - plugin hook feature enablement
  - autonomous continuation authority

## Plugin Runtime Files

Bundled SessionStart hook implementation은 설치되는 plugin runtime surface 안에 둡니다.

- canonical plugin hook config: `hooks/hooks.json`
- canonical plugin hook script: `hooks/session_start_context.py`

`hooks/hooks.json`의 command는 plugin install 위치에 독립적이어야 하므로 `${PLUGIN_ROOT}`를 기준으로 script를 호출합니다.

```json
{
  "type": "command",
  "command": "/usr/bin/python3 \"${PLUGIN_ROOT}/hooks/session_start_context.py\""
}
```

## Hook Runtime Contract

SessionStart hook은 Codex `SessionStart` event payload만 처리합니다.
`source`는 `startup` 또는 `resume`일 때만 context를 출력합니다.
`clear`는 사용자가 맥락 초기화를 의도했을 수 있으므로 조용히 종료합니다.

Plan/flow discovery order:

1. payload `cwd`에서 위로 올라가며 `.agents/sessions/`를 찾습니다.
2. 오늘 날짜 `YYYYMMDD/000-plan.md`가 있으면 우선 사용합니다.
3. 오늘 plan이 없으면 가장 최근 날짜 directory의 `000-plan.md`를 사용합니다.
4. plan frontmatter에서 `active_flow`, `latest_user_request`, `latest_decision`, `required_next_action`, `pending_question_state`, `verification_status`를 읽습니다.
5. `active_flow`가 있으면 같은 날짜의 `{active_flow}.md`를 읽습니다.
6. active flow file이 없으면 `001-*` 중 번호가 가장 큰 flow file을 latest flow로 읽습니다.

Output은 `hookSpecificOutput.additionalContext`를 사용하는 JSON이어야 합니다.
Context는 1500자 이하로 제한합니다.

Context must include:

- current Codex session id when payload provides one, or `not provided`
- plan date
- active flow or latest flow file name
- latest user request summary when recorded
- latest decision when recorded
- verification status when recorded
- required next action when recorded
- pending question state when recorded
- closure state summary
- warning that the context does not grant approval or terminal closure authority

## Side-Effect Boundary

SessionStart hook은 파일을 쓰지 않습니다.
plan/flow가 없으면 기록을 만들지 않고 "no plan found" context만 출력합니다.
hook은 다음 권한을 만들지 않습니다.

- autonomous continuation approval
- destructive action approval
- external action approval
- commit/push/PR/publish/release/version bump approval
- plugin hook feature activation
- global Codex config 변경
- session record write authority

## Skill Relationship

`turn-gate` skill은 이 bundled plugin hook을 startup context backstop으로 설명할 수 있습니다.
그 skill-level guidance는 `specs/skills/turn-gate/hooks/session-start.md`가 소유합니다.
skill guidance는 이 plugin hook spec의 packaging 계약을 반복 소유하지 않고, `turn-gate`가 startup context를 어떻게 advisory context로 취급해야 하는지에 집중합니다.

## 검토 질문

- SessionStart hook code 위치가 plugin runtime `hooks/session_start_context.py`로 일관되는가?
- `hooks/hooks.json`이 `SessionStart` matcher `startup|resume`으로 script를 연결하는가?
- hook이 `clear` source에서 조용히 종료하는가?
- payload가 제공하는 current Codex session id를 advisory context로 노출하는가?
- hook이 plan/flow를 읽기만 하고 쓰지 않는가?
- context가 approval, continuation, terminal closure authority를 만들지 않는다고 명시하는가?
- release surface에 `hooks/session_start_context.py`가 포함되는가?
