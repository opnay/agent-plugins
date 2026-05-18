# turn-gate Stop hook guidance sub-spec

## 목적

이 문서는 `turn-gate` skill이 bundled Codex `Stop` hook을 어떻게 runtime backstop으로 설명하고 해석해야 하는지 소유합니다.
설치 단위 hook 자체의 위치, manifest 연결, build/release 포함 기준, hook script 계약은 plugin-level `specs/hooks/stop-hook.md`가 소유합니다.
skill 본문을 재생성하는 작업자는 Stop hook 관련 runtime guidance를 작성하기 전에 이 spec을 기준으로 삼되, packaging 계약을 skill spec으로 끌어오지 않습니다.

## 경계

- 포함:
  - `turn-gate` runtime guidance에서 Stop hook을 optional backstop으로 설명하는 기준
  - hook block을 받은 뒤 main agent가 해야 할 record refresh, reporting, next-flow reopening 의무
  - hook이 `turn-gate` conversation-level rule을 대체하지 않는다는 경계
  - hook trust/reload와 plugin hook feature enablement를 user-gated setup으로 설명하는 기준
- 제외:
  - plugin runtime hook file layout
  - `.codex-plugin/plugin.json`의 `hooks` field 연결
  - build/release surface 포함 기준
  - hook script의 입력 판단과 JSON output shape
  - global Codex config 설치
  - `/hooks` trust/reload UI 조작
  - plugin hook feature enablement
  - hook을 통한 session record write
  - turn-gate conversation-level operating rule 대체

## Plugin Hook Spec 의존 관계

Bundled Stop hook의 설치 단위 계약은 plugin-level `specs/hooks/stop-hook.md`가 소유합니다.

- canonical plugin hook config: `hooks/hooks.json`
- canonical plugin hook script: `hooks/turn_gate_stop.py`
- plugin manifest connection: `.codex-plugin/plugin.json`의 `"hooks": "./hooks/hooks.json"`
- build/release inclusion: `hooks/` runtime root

이 skill sub-spec은 위 packaging 계약을 반복 소유하지 않습니다.
plugin hook packaging, hook script behavior, manifest field, release surface를 바꿀 때는 먼저 `specs/hooks/stop-hook.md`를 갱신하고, `turn-gate` runtime guidance 의미가 달라지는 경우에만 이 파일을 함께 갱신합니다.

`turn-gate` runtime guidance를 반영할 때는 다음 표면을 함께 점검합니다.

- `skills/turn-gate/SKILL.md`
- `skills/turn-gate/references/session-records.md`가 Continuity Guard 의미를 바꿀 필요가 있는지
- `README.md`
- `.codex-plugin/plugin.json`

이 spec은 dev-side contract owner입니다.
설치된 runtime skill body가 실행 중 `specs/` 경로를 읽으라고 지시하면 안 됩니다.
필요한 runtime guidance는 `SKILL.md`나 release에 포함되는 `references/`로 흡수해야 합니다.

## Runtime guidance 작성 기준

runtime `SKILL.md`에는 Stop hook을 optional runtime backstop으로 설명합니다.
다음 의미를 반드시 유지합니다.

- Stop hook은 `turn-gate`의 conversation-level operating rule을 대체하지 않습니다.
- Stop hook이 있어도 main agent는 reporting, record refresh, next-flow reopening을 직접 수행해야 합니다.
- Stop hook이 terminal closure를 block하면 main agent는 active flow record와 `000-plan.md`를 refresh하고, block reason의 `required_next_action`으로 돌아가야 합니다.
- hook trust/reload와 plugin hook feature enablement는 user-gated setup입니다.
- Stop hook guidance는 bundled trusted plugin hook 기준입니다.
- Stop hook의 script 위치, manifest field, release inclusion, JSON output shape를 설명해야 하는 경우에는 plugin-level `specs/hooks/stop-hook.md`를 기준으로 짧게 참조합니다.

## 권한과 부작용 경계

`turn-gate` guidance는 Stop hook이 다음 권한을 만들지 않는다고 설명해야 합니다.

- destructive action approval
- external action approval
- commit/push/PR/publish/release/version bump approval
- plugin hook feature activation
- global Codex config 변경
- turn-gate record write authority

## 검토 질문

- Skill spec이 plugin hook packaging 계약을 직접 소유한다고 말하지 않는가?
- Plugin hook file layout과 manifest/build 계약은 `specs/hooks/stop-hook.md`로 라우팅되는가?
- Runtime guidance가 Stop hook을 optional backstop으로만 설명하는가?
- Stop hook이 기록을 직접 수정하지 않고, main agent가 refresh/report/next-flow를 수행한다고 설명하는가?
- Hook trust/reload와 plugin hook feature enablement가 user-gated setup으로 남는가?
- runtime body가 dev-only spec 경로를 실행 지시로 노출하지 않는가?
