# empirical-prompt-tuning 스킬 스펙

## 목적

`empirical-prompt-tuning`은 재사용 가능한 agent-facing instruction을 작성자 직관이 아니라 fresh executor evidence로 검증하고 개선하기 위한 메타 스킬입니다.
대상은 skill, slash command, task prompt, AGENTS 또는 CLAUDE 섹션, 코드 생성 프롬프트처럼 반복 사용되는 텍스트 지시입니다.

## 경계

- 포함:
  - 재사용 가능한 instruction의 평가 시나리오 설계
  - 시나리오별 고정 checklist와 `[critical]` 항목 정의
  - fresh subagent를 사용한 bias-reduced execution 평가
  - executor self-report와 caller-side metric을 함께 쓰는 양면 평가
  - 작은 수정 단위의 iterative prompt tuning
  - empirical evaluation 불가 환경에서 structural review only 또는 skipped 판정
- 제외:
  - 원본 skill 또는 plugin의 경계 설계 자체
  - 일반적인 implementation debugging
  - reusable tool policy 설계 전반
  - 한 번 쓰고 버릴 disposable prompt 손질

## 처리하려는 작업 형태

- 새로 만든 skill 또는 프롬프트를 실제 executor 기준으로 검증하고 싶은 경우
- agent 동작이 기대와 다를 때 원인을 instruction ambiguity에서 찾고 싶은 경우
- 고빈도 지시문을 반복 평가로 더 견고하게 만들고 싶은 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex/skills/empirical-prompt-tuning/SKILL.md`
- 핵심 진입 질문:
  - 이 작업의 병목이 skill 생성이 아니라 instruction validation인가
  - author self-reread가 아니라 fresh executor evidence가 필요한가
  - 2개 이상 현실 시나리오로 generalization을 확인해야 하는가

## 워크플로 계약

1. description/body 또는 trigger/body의 정합성을 먼저 본다.
2. 시나리오 2-3개와 checklist를 먼저 고정한다.
3. fresh subagent에게 동일 instruction을 실행하게 한다.
4. self-report와 caller-side metric을 함께 수집한다.
5. 한 iteration에 한 theme만 수정한다.
6. fresh subagent로 다시 실행한다.
7. 수렴하거나 설계 자체가 틀렸다고 판단될 때까지 반복한다.

## 평가 계약

- 성공/실패는 `[critical]` 항목 기준으로 가른다.
- 정확도는 checklist 충족률로 계산한다.
- `tool_uses`, `duration_ms`, retry는 보조 지표로 사용한다.
- 질적 신호:
  - ambiguous wording
  - judgment calls
  - retry reasons
- 양적 신호:
  - accuracy
  - steps
  - duration

## 환경 제약

- fresh subagent dispatch가 불가능하면 empirical evaluation이라고 부르지 않는다.
- 이런 경우 허용되는 결과는 둘 중 하나다.
  - structural review only
  - evaluation skipped
- self-reread를 empirical의 대체물로 허용하지 않는다.

## 독립성 원칙

- 이 스킬은 skill 생성법을 소유하지 않는다.
- 이 스킬은 tool-use policy 전반을 소유하지 않는다.
- 이 스킬은 오직 reusable instruction evaluation workflow만 소유한다.
- sibling skill의 숨은 맥락 없이 단독으로 이해 가능해야 한다.

## 확장 원칙

- 새 reference나 asset은 empirical workflow를 더 명확하게 만드는 경우에만 추가한다.
- 반복적으로 쓰는 scenario template이나 report template이 생기면 별도 reference로 분리할 수 있다.
- plugin-level routing 변경은 이 스킬 수정과 별도 변경 단위로 다룬다.

## 출처

- 참고 원문:
  - `mizchi/chezmoi-dotfiles`의 `dot_claude/skills/empirical-prompt-tuning/SKILL.md`
  - 링크: `https://github.com/mizchi/chezmoi-dotfiles/blob/main/dot_claude/skills/empirical-prompt-tuning/SKILL.md`
- 이 스펙과 로컬 스킬은 위 원문의 핵심 관점인 empirical evaluation, fresh executor, dual-sided evidence, iterative tuning을 참고해 이 저장소의 규칙과 경계에 맞게 재구성했다.
