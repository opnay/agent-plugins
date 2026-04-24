## 사용자 스펙 의도

- session artifact를 deterministic schema로 생성하고 유지하고 싶다.
- `session_record.md`, `change_record.md`, `retrospective_record.md`를 같은 운영 규칙 아래 다루고 싶다.
- read-only 조회와 실제 write/update를 명확히 구분하고 싶다.

---

# session-manager 스킬 스펙

## 목적

`session-manager`는 `.agents/sessions/<uuid>/` 아래의 session artifact를 일관된 schema로 생성, 갱신, 검증하기 위한 운영 스킬입니다.
핵심 대상은 `session_record.md`, `change_record.md`, `retrospective_record.md`입니다.

## 경계

- 포함:
  - 세 가지 session artifact의 생성과 ensure
  - logical field 단위 갱신
  - record validation과 listing
  - multiline-safe write workflow와 schema 유지
- 제외:
  - 일반 note taking
  - `.agents/sessions` 외 레이아웃 관리
  - full skill scaffold creation

## 처리하려는 작업 형태

- 작업 시작 시 deterministic session record가 필요한 경우
- change delivery나 retrospective artifact를 현재 세션에 추가해야 하는 경우
- session artifact가 schema를 지키는지 검증해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/session-manager/SKILL.md`
- 관련 스크립트: `advance-codex-dev/skills/session-manager/scripts/controller.py`

## 핵심 처리 계약

- `session_record.md`는 numbered schema를 유지한다.
- `change_record.md`와 `retrospective_record.md`는 marker comment가 포함된 markdown schema를 유지한다.
- multiline content는 `--value-file` 경로를 우선 사용한다.
- `show`, `validate`, `list`는 read-only 작업으로 취급한다.

## 아티팩트 계약

- `session_record.md`: task context, runbook, verification evidence, residual risk
- `change_record.md`: reviewer-facing change summary와 gate notes
- `retrospective_record.md`: session close reflection과 follow-up guardrails

## 검토 질문

- 지금 다루는 대상이 세 가지 session artifact 중 무엇인지 명확한가?
- schema나 marker comment를 깨뜨리지 않고 갱신하고 있는가?
- 이번 작업이 read-only인지 write/update인지 구분이 선명한가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: session artifact schema와 조작 규칙은 hidden sibling context 없이 재사용 가능해야 하며, 다른 skill은 이 artifact를 소비할 수만 있고 schema contract 자체는 이 스펙이 닫아야 한다.

## 확장 원칙

- 새 artifact kind를 추가하려면 controller script, schema 문서, 이 스펙을 함께 갱신한다.
- existing artifact schema를 바꾸면 machine-readable parsing 영향부터 먼저 명시한다.
