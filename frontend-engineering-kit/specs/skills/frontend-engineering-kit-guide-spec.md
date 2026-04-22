# frontend-engineering-kit-guide 스킬 스펙

## 목적

`frontend-engineering-kit-guide`는 `frontend-engineering-kit`의 엔트리포인트로서 현재 frontend task의 dominant concern을 분류하고 첫 specialist skill을 정하는 스킬입니다.

## 경계

- 포함:
  - workflow, pattern, React, component, domain, UI, TDD concern 분류
  - first specialist handoff 선택
- 제외:
  - 각 specialist concern의 상세 해결
  - generic workflow routing

## 처리하려는 작업 형태

- frontend task가 여러 concern으로 섞여 있는 경우
- 어떤 specialist skill이 먼저 문제를 소유해야 하는지 정해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-engineering-kit/skills/frontend-engineering-kit-guide/SKILL.md`

## 핵심 처리 계약

- dominant concern을 하나로 좁히고 이유를 명시한다.
- 필요한 경우 handoff sequence를 제안하되 specialist detail은 중복하지 않는다.
- React/component/domain boundary overlap이 생기면 어느 layer가 먼저인지 명확히 해야 한다.

## 독립성 원칙

- 이 스킬은 specialist routing만 소유한다.
- guide만 읽어도 어떤 frontend concern이 어떤 skill로 가야 하는지 이해 가능해야 한다.

## 확장 원칙

- concern map이 바뀌면 plugin spec과 함께 갱신한다.

