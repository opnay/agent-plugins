# teammate-reviewer 스킬 스펙

## 목적

`teammate-reviewer`는 correctness, regression, scope creep, weak assumption, missing test를 찾는 independent review role 스킬입니다.

## 경계

- 포함:
  - independent risk-focused review
  - correctness/regression/scope/test 관점 점검
  - reviewer-style handoff 작성
- 제외:
  - implementation
  - orchestration
  - broad research planning

## 처리하려는 작업 형태

- 다른 teammate의 결과물에 독립적인 검토가 필요한 경우
- risk-focused review output이 필요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `teammate-kit/skills/teammate-reviewer/SKILL.md`
- 관련 상위 라우팅: `teammate-kit-guide`, `teammate-orchestrator`

## 핵심 처리 계약

- review는 correctness, regression, scope, test 관점을 우선한다.
- 발견 사항은 근거와 함께 전달해야 한다.
- reviewer는 implementer 역할과 섞이지 않도록 독립 시각을 유지한다.

## 독립성 원칙

- 이 스킬은 직접 구현을 소유하지 않는다.
- review output만 읽어도 후속 조치 여부를 판단할 수 있어야 한다.

## 확장 원칙

- 새 규칙은 review signal의 신뢰도를 높일 때만 추가한다.

