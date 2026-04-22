# frontend-workflow-guide 스킬 스펙

## 목적

`frontend-workflow-guide`는 frontend implementation, refactor, fix, review 작업의 기본 execution workflow를 제공하는 스킬입니다.
핵심은 project structure를 먼저 읽고, 더 깊은 concern의 순서를 정하며, practical verification path를 확보하는 것입니다.

## 경계

- 포함:
  - frontend execution workflow framing
  - project structure inspection
  - specialist concern 순서 결정
  - TDD 우선 여부 판단
- 제외:
  - architecture pattern 자체의 최종 판단
  - React/component/domain detail 자체
  - generic repo-wide workflow orchestration

## 처리하려는 작업 형태

- frontend 작업을 시작하기 전에 어디를 읽고 어떤 specialist skill을 붙일지 정해야 하는 경우
- implementation과 refactor, fix, review 흐름을 frontend 관점에서 정리해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/frontend-workflow-guide/SKILL.md`
- 관련 상위 라우팅: `frontend-engineering-kit-guide`

## 핵심 처리 계약

- 구조를 먼저 읽고 task를 workflow/pattern/react/component/domain/design/TDD concern으로 나눈다.
- 가능한 경우 failing-test-first path를 우선 고려한다.
- 더 좁은 specialist skill이 필요하면 실행 순서를 명시한다.

## 독립성 원칙

- 이 스킬은 frontend execution framing만 소유한다.
- specialist skill 없이도 작업 시작 순서를 정할 수 있어야 한다.

## 확장 원칙

- default workflow가 바뀌면 guide skill과 plugin spec을 함께 갱신한다.

