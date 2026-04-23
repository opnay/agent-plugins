# frontend-kit-guide 스킬 스펙

## 목적

`frontend-kit-guide`는 요청이 `frontend-kit` 범위에 속하는지 판단하고, 축 분리 중심의 작업을 `react-architecture`로 라우팅하는 엔트리포인트 스킬입니다.

## 경계

- 포함:
  - plugin 범위 판별
  - axis-separation 작업의 시작점 지정
  - broad request와 narrow request의 진입 판단
- 제외:
  - 실제 axis classification 수행
  - top-level frontend pattern choice

## 처리하려는 작업 형태

- 이 플러그인을 써야 하는지 판단해야 하는 경우
- broad한 프론트엔드 요청에서 축 분리 작업이 핵심인지 확인해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-kit/skills/frontend-kit-guide/SKILL.md`

## 핵심 처리 계약

- 요청이 React-facing artifact의 축 분리, 의존 방향, 구현 순서 문제라면 `react-architecture`로 보낸다.
- 요청이 top-level structure choice나 generic task-framing 문제라면 이 플러그인 범위를 벗어났다고 명시한다.

## 독립성 원칙

- 이 스킬은 라우팅만 소유하고 세부 구조 규칙은 소유하지 않는다.

## 확장 원칙

- 새 규칙은 plugin entrypoint로서의 분류 정확도를 높일 때만 추가한다.
