# teammate-implementer 스킬 스펙

## 목적

`teammate-implementer`는 bounded code 또는 content change를 end-to-end로 소유하는 teammate execution role 스킬입니다.

## 경계

- 포함:
  - scoped change 구현
  - local verification
  - concrete edits와 residual risk 보고
- 제외:
  - orchestration
  - broad research
  - independent review

## 처리하려는 작업 형태

- 하나의 teammate가 한 변경 단위를 맡아야 하는 경우
- 구현과 검증 결과를 role-based handoff로 남겨야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `teammate-kit/skills/teammate-implementer/SKILL.md`
- 관련 상위 라우팅: `teammate-kit-guide`, `teammate-orchestrator`

## 핵심 처리 계약

- scope를 벗어나지 않는 bounded change를 소유한다.
- 구현 결과와 verification evidence, residual risk를 함께 보고한다.
- 다른 teammate가 이어받을 수 있게 handoff-friendly output을 남긴다.

## 독립성 원칙

- 이 스킬은 orchestration이나 review judgement를 소유하지 않는다.
- implementer output만 읽어도 변경 범위와 검증 상태를 이해할 수 있어야 한다.

## 확장 원칙

- 새 규칙은 bounded execution role의 전달 품질을 높일 때만 추가한다.

