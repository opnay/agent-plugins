# teammate-researcher 스킬 스펙

## 목적

`teammate-researcher`는 bounded investigation과 evidence-based synthesis를 담당하는 teammate role 스킬입니다.

## 경계

- 포함:
  - local codebase context 조사
  - option 비교와 assumption 검증
  - implementation 전 decision-ready handoff 작성
- 제외:
  - implementation
  - independent review
  - orchestration

## 처리하려는 작업 형태

- 조사와 선택지 정리가 먼저 필요한 경우
- 구현 전에 근거 기반 handoff가 필요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `teammate-kit/skills/teammate-researcher/SKILL.md`
- 관련 상위 라우팅: `teammate-kit-guide`, `teammate-orchestrator`

## 핵심 처리 계약

- 조사 결과는 evidence와 interpretation을 구분해서 남긴다.
- handoff는 concise하지만 decision-ready해야 한다.
- speculation보다 검증 가능한 관찰을 우선한다.

## 독립성 원칙

- 이 스킬은 구현이나 review judgement를 소유하지 않는다.
- research output만 읽어도 다음 실행자가 바로 이어받을 수 있어야 한다.

## 확장 원칙

- 새 규칙은 evidence quality와 handoff clarity를 높일 때만 추가한다.

