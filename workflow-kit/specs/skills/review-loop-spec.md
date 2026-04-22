# review-loop 스킬 스펙

## 목적

`review-loop`는 code review, QA review, self-review에서 나온 material finding을 blocking-first 기준으로 처리하는 workflow 스킬입니다.

## 경계

- 포함:
  - finding threshold 적용
  - blocking finding 선택
  - bounded fix와 immediate verification
  - deferred note 분리
- 제외:
  - broad cleanup
  - speculative polish
  - generic execution workflow

## 처리하려는 작업 형태

- review feedback에서 진짜 막는 이슈만 골라 처리해야 하는 경우
- correctness, regression, reliability 중심으로 빠르게 고쳐야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/review-loop/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- finding은 구체적 영향과 높은 신뢰도가 있을 때만 blocking으로 올린다.
- 한 loop에는 하나의 bounded blocking finding만 선택한다.
- low-value note는 deferred note로 남기고 delivery를 막지 않는다.

## 독립성 원칙

- 이 스킬은 broad refactor나 open-ended polish를 소유하지 않는다.
- review output만 읽어도 무엇이 blocking이고 무엇이 deferred인지 이해 가능해야 한다.

## 확장 원칙

- finding threshold가 바뀌면 severity gate와 함께 갱신한다.

