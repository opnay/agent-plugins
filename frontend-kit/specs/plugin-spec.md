# Frontend Kit 플러그인 스펙

## 플러그인 목적

`frontend-kit`은 React-facing artifact를 `system`, `design`, `domain` 축으로 분리해서 다루는 플러그인입니다.
핵심 책임은 component, hook, utility, page, global state가 어떤 축에 속하는지 판단하고, dependency direction과 구현 순서, refactor 순서를 정하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - React-facing artifact의 축 분류
  - `system -> design -> domain` 구조 중요도 판단
  - `design -> domain -> system` 구현 순서 판단
  - `domain -> design -> system` 의존 방향 판단
  - mixed component, hook, provider, broader state 정리
- 제외:
  - top-level frontend pattern choice
  - UI 시각 품질 자체 설계
  - business rule 자체 설계
  - generic task orchestration

## 처리하려는 작업 형태

- React 코드에서 system/design/domain 책임이 섞여 있는 작업
- component, hook, utility, page, global state의 ownership 기준이 필요한 작업
- 구현 순서와 refactor 순서를 분리해서 다뤄야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `frontend-kit-guide`
- 대표 스펙: `frontend-kit/specs/plugin-spec.md`
- skill 상세 스펙 위치: `frontend-kit/specs/skills/*.md`

## 내장 skill 체계

- `frontend-kit-guide`: 이 플러그인의 범위인지 판단하고, 축 분리 작업을 `react-architecture`로 라우팅한다.
  - spec: `frontend-kit/specs/skills/frontend-kit-guide-spec.md`
- `react-architecture`: React-facing artifact를 `system`, `design`, `domain` 축으로 분리하고 dependency direction과 순서를 정한다.
  - spec: `frontend-kit/specs/skills/react-architecture-spec.md`

## SDD 운영 원칙

- plugin spec은 플러그인 경계와 routing surface만 소유한다.
- 각 skill의 처리 계약은 별도 `specs/skills/` 문서로 분리한다.
- 축 분리 기준이 바뀌면 `frontend-kit-guide`, `react-architecture`, 관련 spec을 같은 변경에서 갱신한다.
- artifact 이름이나 기존 레이어 이름보다 실제 존재 이유를 우선하는 원칙을 유지한다.

## 현재 구조 메모

- 이 플러그인의 주요 리스크는 `system`, `design`, `domain` 축을 파일 종류나 기존 레이어 이름으로 성급히 고정해 다시 섞어버리는 것이다.
