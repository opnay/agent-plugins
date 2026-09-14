## 사용자 스펙 의도

- 재사용되는 instruction을 작성자 감이 아니라 fresh executor evidence로 테스트하고 싶다.
- 시나리오와 checklist를 먼저 잠그고 분석 결과를 evidence 중심으로 보고하고 싶다.
- empirical evaluation이 불가능한 상황과 구조 리뷰만 가능한 상황을 구분하고 싶다.
- 기본 평가 세트는 8개 시나리오를 기준으로 삼고, 큰 평가 요청은 N개 단위 batch로 나눠 진행하고 싶다.
- subagent 한도에 걸리면 완료됐거나 더 쓰지 않는 executor를 닫고 남은 batch를 이어가고 싶다.

---

# skill-scenario-testing 스킬 스펙

## 목적

`skill-scenario-testing`은 reusable agent-facing instruction을 작성자 직관이 아니라 fresh executor evidence로 테스트하고 분석 보고하기 위한 메타 스킬입니다.
대상은 skill, slash command, task prompt, AGENTS section, code-generation prompt처럼 반복 사용되는 텍스트 지시입니다.

## 경계

- 포함:
  - reusable instruction의 평가 시나리오 설계
  - 고정 checklist와 `[critical]` 항목 정의
  - fresh subagent 기반 bias-reduced execution 평가
  - executor self-report와 caller-side metric의 동시 수집
  - ambiguity, judgment call, retry reason 중심의 분석 보고
- 제외:
  - 원본 skill이나 plugin의 경계 설계 자체
  - 일반 implementation debugging
  - reusable tool policy 전반
  - disposable prompt polishing
  - instruction 직접 수정 또는 patch 적용

## 처리하려는 작업 형태

- 새로 만든 skill이나 prompt를 실제 executor 기준으로 테스트하고 싶은 경우
- agent failure 원인이 implementation이 아니라 instruction ambiguity라고 의심되는 경우
- 고빈도 instruction을 실제 executor 관점에서 분석하고 싶은 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/skill-scenario-testing/SKILL.md`
- 호출 방식: 직접 호출하거나 manifest prompt의 안내를 따른다.

## 핵심 처리 계약

- description과 body가 같은 약속을 하는지 먼저 정합성을 점검한다.
- 시나리오와 checklist를 먼저 고정한 뒤 평가를 시작한다.
- 별도 요청이 없으면 기본 scenario set은 8개로 잡는다. smoke check만 필요한 경우에는 2-3개를 사용할 수 있고, 사용자가 더 큰 수를 요청하면 batch로 나눠 진행한다.
- 각 empirical run은 fresh subagent를 사용하고 self-reread를 대체물로 허용하지 않는다.
- 정성 신호는 ambiguity, judgment call, retry reason을 중심으로 보고 정량 신호는 accuracy, `tool_uses`, `duration_ms`를 보조 지표로 쓴다.
- 한 evaluation pass에는 하나의 분석 target/theme만 다룬다.
- independent scenario는 batch로 병렬 dispatch할 수 있다. subagent 한도에 걸리면 완료됐거나 더 이상 사용하지 않을 executor를 닫고 다음 batch를 이어간다.

## 환경 제약

- fresh subagent dispatch가 불가능하면 결과를 empirical evaluation로 부르지 않는다.
- 이 경우 허용되는 산출물은 `structural review only` 또는 `evaluation skipped`다.

## 검토 질문

- 시나리오와 checklist를 평가 전에 먼저 고정했는가?
- empirical run마다 fresh subagent를 실제로 사용했는가?
- 기본 요청에서 8개 시나리오를 고려했고, 더 큰 요청은 N개 batch로 나눴는가?
- subagent 한도에 걸렸을 때 완료된 executor를 닫고 남은 scenario를 이어갔는가?
- 이번 evaluation pass가 하나의 분석 target/theme만 다루고 있는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: empirical evaluation 절차는 hidden plugin context 없이 재사용 가능해야 하며, sibling skill은 라우팅이나 후속 설계만 맡고 평가 절차 자체는 이 스펙에서 닫혀 있어야 한다.

## 확장 원칙

- scenario template이나 report template이 반복되면 reference 문서로 분리할 수 있다.
- plugin 사용 기준 변경은 이 스킬 스펙이 아니라 `plugin.md`, README, manifest prompt에서 먼저 반영한다.
