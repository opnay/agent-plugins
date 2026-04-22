# ralph-loop 스킬 스펙

## 목적

`ralph-loop`는 하나의 bounded issue를 짧은 fix-verify-reassess cycle로 반복 개선하는 workflow 스킬입니다.

## 경계

- 포함:
  - bounded issue selection
  - smallest useful fix
  - immediate verification
  - next loop decision
- 제외:
  - broad end-to-end delivery
  - review finding triage 전반
  - planning-only work

## 처리하려는 작업 형태

- UI polish, refactor stabilization, flaky issue reduction 같은 bounded refinement
- progressive improvement가 더 적합한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/ralph-loop/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- 한 loop는 하나의 primary issue만 다룬다.
- 가장 작은 fix로 가설을 검증한다.
- 각 loop 후에 개선 여부와 다음 loop 정당성을 다시 판단한다.

## 독립성 원칙

- 이 스킬은 broad execution이나 blocking-first review handling을 소유하지 않는다.
- loop 기록만 읽어도 무엇을 시험했고 무엇이 남았는지 이해 가능해야 한다.

## 확장 원칙

- stop condition이나 issue threshold가 바뀌면 guide skill과 함께 갱신한다.

