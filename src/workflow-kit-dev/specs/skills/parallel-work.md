## 사용자 스펙 의도

- 독립적인 작업을 병렬 lane으로 나눠 시간을 줄이고 싶다.
- lane ownership과 integration 규칙을 먼저 잠그고 싶다.
- broad execution이나 review-only correction과는 분리하고 싶다.

---

# parallel-work 스킬 스펙

## 목적

`parallel-work`는 몇 개의 clearly independent task를 병렬 lane으로 나누고 결과를 통합하는 execution workflow 스킬입니다.

## 경계

- 포함:
  - lane 분해와 ownership 정의
  - 병렬 실행 중 coordination
  - integration과 shared verification
- 제외:
  - broad multi-phase delivery 전반
  - sequential planning only
  - review-only correction

## 처리하려는 작업 형태

- 서로 독립적으로 시작 가능한 bounded work item이 여러 개 있는 경우
- main challenge가 safe split과 integration인 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit-dev/skills/parallel-work/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

- lane은 unanswered design decision이나 file overlap이 없어야 한다.
- 각 lane은 독립 검증 경로를 가져야 한다.
- 통합 단계에서 whole-result verification을 다시 수행한다.

## 검토 질문

- lane 사이에 unresolved design decision이나 file overlap이 없는가?
- 각 lane이 자기 검증 경로를 갖고 있는가?
- integration 뒤 whole-result verification이 다시 잡혀 있는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 병렬 분해와 통합 규칙은 hidden sibling context 없이 재사용 가능해야 하며, broad execution이나 refinement loop는 별도 skill에 남겨야 한다.

## 확장 원칙

- independence 판단 기준이 바뀌면 guide skill과 함께 갱신한다.
