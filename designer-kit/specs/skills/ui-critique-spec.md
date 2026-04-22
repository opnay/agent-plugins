# ui-critique 스킬 스펙

## 목적

`ui-critique`는 existing screen, concept, flow, direction을 designer/UX 관점에서 비평하고 중요한 문제를 우선순위화하는 스킬입니다.

## 경계

- 포함:
  - hierarchy, clarity, flow, tone, usability 관점의 critique
  - 가장 가치가 큰 문제의 우선순위화
  - 개선이 중요한 이유 설명
- 제외:
  - design brief 작성
  - pre-code screen specification 작성
  - code implementation guidance

## 처리하려는 작업 형태

- 기존 안을 critique해서 무엇을 먼저 고쳐야 할지 알고 싶은 경우
- implementation 전에 방향의 약점을 찾고 싶은 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `designer-kit/skills/ui-critique/SKILL.md`
- 관련 상위 라우팅: `designer-kit-guide`

## 핵심 처리 계약

- critique는 hierarchy, clarity, flow, tone, usability lens를 중심으로 한다.
- 발견 사항은 중요도와 영향 기준으로 정렬해야 한다.
- 단순 취향 평이 아니라 왜 지금 고쳐야 하는지 설명해야 한다.

## 독립성 원칙

- 이 스킬은 새 design direction 생성이나 screen contract 작성을 소유하지 않는다.
- critique 결과는 독립 문서로도 다음 의사결정에 바로 쓰일 수 있어야 한다.

## 확장 원칙

- 새로운 critique lens는 design quality 판단에 반복적으로 필요한 경우에만 추가한다.
- implementation 조언이 필요해지면 다른 specialist plugin으로 handoff한다.

