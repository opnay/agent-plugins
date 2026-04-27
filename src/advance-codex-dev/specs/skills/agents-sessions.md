## 사용자 스펙 의도

- `.agents/sessions` 폴더의 기본 규격을 먼저 잡아두고 싶다.
- 기능이나 조작 workflow가 아니라, 이 폴더를 언제 쓰는지 정도만 정의하고 싶다.
- session artifact 관리 기능은 만들지 않고, 기본 골격을 `turn-gate`의 날짜 기반 plan/flow record 구조와 일치시키고 싶다.

---

# agents-sessions 스킬 스펙

## 목적

`agents-sessions`는 `.agents/sessions` 폴더의 기본 용도와 `turn-gate` 중심 plan/flow record 골격을 정의하는 foundation skill입니다.
이 skill은 폴더 자체가 어떤 종류의 date-scoped operational artifact를 담는지 설명하고, 실제 artifact 생성이나 갱신은 다루지 않습니다.

## 경계

- 포함:
  - `.agents/sessions` 폴더의 목적 설명
  - session-scoped artifact와 repository source artifact의 구분
  - 어떤 작업에서 이 폴더를 사용해야 하는지에 대한 기본 기준
- 제외:
  - session artifact 생성, 갱신, 검증
  - 별도 session record schema나 controller 관리
  - 구체적인 workflow state machine
  - plugin, skill, source file scaffold 생성

## 처리하려는 작업 형태

- `.agents/sessions` 폴더를 언제 쓰는지 설명해야 하는 경우
- session-scoped 기록과 repo source 변경을 구분해야 하는 경우
- 새 session artifact를 만들기 전에 `turn-gate` root 파일 골격과 어떻게 맞춰야 하는지 확인해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/agents-sessions/SKILL.md`
- 관련 기본 플로우: `loop-kit:turn-gate`
- 관련 상위 라우팅: `advance-codex-dev-guide`

## 핵심 처리 계약

- `.agents/sessions`는 agent 작업의 date-scoped operational artifacts를 두는 저장소로 설명한다.
- session directory는 `turn-gate`와 같은 날짜 기반 `.agents/sessions/{YYYYMMDD}/` 형태로 설명한다.
- 기본 골격은 날짜 디렉터리 루트의 `000-plan.md`와 `001-*` flow record 패턴이라고 설명한다.
- 다른 session artifact는 이 기본 골격을 기준으로 배치되어야 하며, 별도 session record schema를 병렬 표준으로 만들지 않는다.
- 이 폴더는 plugin source, release surface, reusable skill/spec 문서의 대체 위치가 아니다.
- session-scoped scratch artifact가 필요하면 현재 session directory 아래에 둔다.

## 검토 질문

- 지금 필요한 것이 폴더의 기본 규격 설명인가, 실제 record 조작인가?
- 기록하려는 정보가 session-scoped operational artifact인가, repo에 남아야 하는 source/spec 변경인가?
- 해당 artifact가 `000-plan.md`와 zero-padded flow record라는 기본 골격에 맞게 배치되는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 폴더의 기본 용도와 골격은 다른 workflow 없이도 이해 가능해야 한다. 다만 active turn 운영은 기본 플로우인 `turn-gate`가 소유한다.

## 확장 원칙

- 이 skill에는 폴더-level 규격만 추가한다.
- active turn 운영 규칙은 `turn-gate`로 넘기고, artifact schema나 controller command는 이 skill에 추가하지 않는다.
