# turn-gate SessionStart hook guidance sub-spec

## 목적

이 문서는 `turn-gate` skill이 bundled Codex `SessionStart` hook의 startup context를 어떻게 해석해야 하는지 소유합니다.
설치 단위 hook 자체의 위치, hook config, script 계약은 plugin-level `specs/hooks/session-start.md`가 소유합니다.

## 경계

- 포함:
  - SessionStart context를 advisory startup context로 취급하는 기준
  - plan/flow 발견 결과를 active work로 재개할지 질문으로 라우팅할지 판단하는 기준
  - context가 approval 또는 terminal closure authority가 아니라는 경계
- 제외:
  - plugin runtime hook file layout
  - hook script의 discovery algorithm 세부 구현
  - hook trust/reload UI 조작
  - plugin hook feature enablement
  - session record write

## Runtime Guidance 작성 기준

runtime `SKILL.md`에는 SessionStart hook을 optional startup context backstop으로 짧게 설명할 수 있습니다.
다음 의미를 유지합니다.

- SessionStart hook은 `.agents/sessions/`의 plan/flow 상태를 알려주는 context loader입니다.
- SessionStart hook이 제공한 context는 자동 continuation 지시가 아닙니다.
- Context에 `required_next_action`이 있어도 main agent는 현재 사용자 메시지와 active flow record를 다시 확인해야 합니다.
- Context는 commit, push, PR, release, version bump, deletion, external action, terminal closure 승인을 만들지 않습니다.
- 이전 flow가 closed 상태면 historical context로만 취급합니다.
- plan/flow가 없다는 context는 정상 preparation으로 시작하라는 신호입니다.

## 검토 질문

- Runtime guidance가 SessionStart hook을 advisory context로만 설명하는가?
- Hook-provided context가 approval 또는 continuation authority로 오해되지 않게 설명하는가?
- Plugin hook packaging 계약은 `specs/hooks/session-start.md`로 라우팅되는가?
- Runtime body가 dev-only spec 경로를 실행 지시로 노출하지 않는가?
