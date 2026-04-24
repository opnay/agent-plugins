## 사용자 스펙 의도

- review feedback 중 실제로 delivery를 막는 finding만 골라 처리하고 싶다.
- bounded fix와 immediate verification 중심으로 빠르게 대응하고 싶다.
- low-value note는 deferred로 남기고 broad cleanup으로 번지지 않게 하고 싶다.

---

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

- 대표 표면: `workflow-kit-dev/skills/review-loop/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

- finding은 구체적 영향과 높은 신뢰도가 있을 때만 blocking으로 올린다.
- 한 loop에는 하나의 bounded blocking finding만 선택한다.
- low-value note는 deferred note로 남기고 delivery를 막지 않는다.

## 검토 질문

- 지금 선택한 finding이 정말 blocking threshold를 넘는가?
- 한 loop에 하나의 bounded finding만 남겨 두었는가?
- deferred note와 blocking issue가 섞이지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: review finding 판단 기준은 hidden sibling context 없이 읽혀야 하며, broad cleanup과 polish는 별도 workflow로 분리돼야 한다.

## 확장 원칙

- finding threshold가 바뀌면 severity gate와 함께 갱신한다.
