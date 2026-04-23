## 사용자 스펙 의도

- 재사용되는 instruction을 작성자 감이 아니라 fresh executor evidence로 검증하고 싶다.
- 시나리오와 checklist를 먼저 잠그고 작은 iteration 단위로 개선하고 싶다.
- empirical evaluation이 불가능한 상황과 구조 리뷰만 가능한 상황을 구분하고 싶다.

---

# empirical-prompt-tuning 스킬 스펙

## 목적

`empirical-prompt-tuning`은 reusable agent-facing instruction을 작성자 직관이 아니라 fresh executor evidence로 검증하고 반복 개선하기 위한 메타 스킬입니다.
대상은 skill, slash command, task prompt, AGENTS section, code-generation prompt처럼 반복 사용되는 텍스트 지시입니다.

## 경계

- 포함:
  - reusable instruction의 평가 시나리오 설계
  - 고정 checklist와 `[critical]` 항목 정의
  - fresh subagent 기반 bias-reduced execution 평가
  - executor self-report와 caller-side metric의 동시 수집
  - iteration theme를 분리한 작은 수정 단위 tuning
- 제외:
  - 원본 skill이나 plugin의 경계 설계 자체
  - 일반 implementation debugging
  - reusable tool policy 전반
  - disposable prompt polishing

## 처리하려는 작업 형태

- 새로 만든 skill이나 prompt를 실제 executor 기준으로 검증하고 싶은 경우
- agent failure 원인이 implementation이 아니라 instruction ambiguity라고 의심되는 경우
- 고빈도 instruction을 반복 평가로 더 견고하게 만들고 싶은 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex/skills/empirical-prompt-tuning/SKILL.md`
- 관련 상위 라우팅: `advance-codex-guide`

## 핵심 처리 계약

- description과 body가 같은 약속을 하는지 먼저 정합성을 점검한다.
- 시나리오 2-3개와 checklist를 먼저 고정한 뒤에만 수정에 들어간다.
- 각 empirical run은 fresh subagent를 사용하고 self-reread를 대체물로 허용하지 않는다.
- 정성 신호는 ambiguity, judgment call, retry reason을 중심으로 보고 정량 신호는 accuracy, `tool_uses`, `duration_ms`를 보조 지표로 쓴다.
- 한 iteration에는 하나의 개선 theme만 적용한다.

## 환경 제약

- fresh subagent dispatch가 불가능하면 결과를 empirical evaluation로 부르지 않는다.
- 이 경우 허용되는 산출물은 `structural review only` 또는 `evaluation skipped`다.

## 검토 질문

- 시나리오와 checklist를 수정 전에 먼저 고정했는가?
- empirical run마다 fresh subagent를 실제로 사용했는가?
- 이번 iteration이 하나의 개선 theme만 다루고 있는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: empirical evaluation 절차는 hidden plugin context 없이 재사용 가능해야 하며, sibling skill은 라우팅이나 후속 설계만 맡고 평가 절차 자체는 이 스펙에서 닫혀 있어야 한다.

## 확장 원칙

- scenario template이나 report template이 반복되면 reference 문서로 분리할 수 있다.
- plugin-level routing 변경은 이 스킬 스펙이 아니라 `advance-codex-guide`와 `plugin.md`에서 먼저 반영한다.
