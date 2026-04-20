# Designer Kit 플러그인 스펙

## 목적

`designer-kit`은 구현 이전 단계의 design-only workflow를 지원하기 위한 플러그인입니다.
핵심 역할은 지금 필요한 것이 directional definition인지, existing direction에 대한 critique인지, pre-code screen-level specification인지 분류하는 것입니다.

## 경계

- 포함:
  - design brief
  - UI/UX critique
  - pre-code screen specification
  - 위 세 design stage 사이의 routing
- 제외:
  - frontend engineering structure
  - code implementation guidance
  - Figma execution detail

## 진입 표면

- 대표 엔트리포인트: `designer-kit-guide`
- 핵심 분기: 시작점을 design brief, critique, screen specification 중 어디에 둘지 고른다

## 스킬 구성

- `designer-kit-guide`: mixed되었거나 아직 불명확한 design task를 적절한 design stage로 라우팅한다
- `design-brief`: vague idea를 goal, hierarchy, tone, non-goal이 정리된 aligned design direction으로 바꾼다
- `ui-critique`: existing concept, screen, flow를 비평하고 가장 가치가 큰 문제를 우선순위화한다
- `screen-spec`: aligned direction을 sections, content intent, interaction expectation이 있는 pre-code screen specification으로 바꾼다

## 확장 원칙

- 새 skill은 briefing, critique, pre-code specification과 다른 distinct design stage일 때만 추가한다.
- 이 플러그인은 design-first로 유지하고, implementation question을 직접 다루는 방향으로 넓히지 않는다.
- design stage의 집합이나 순서가 바뀌면 `designer-kit-guide`를 함께 갱신한다.
- critique나 screen-spec이 code-facing guidance를 흡수하지 않게 한다.

## 현재 의도 점검

- 현재 플러그인 표면은 design-only pre-code workflow를 중심으로 일관적이다.
- 현재의 주요 리스크는 frontend implementation이나 Figma execution detail 쪽으로 범위가 흐려지는 것이다.
