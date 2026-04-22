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

- 대표 표면: `workflow-kit/skills/parallel-work/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- lane은 unanswered design decision이나 file overlap이 없어야 한다.
- 각 lane은 독립 검증 경로를 가져야 한다.
- 통합 단계에서 whole-result verification을 다시 수행한다.

## 독립성 원칙

- 이 스킬은 broad autopilot이나 bounded refinement loop를 대체하지 않는다.
- parallel split과 integration 규칙만 읽어도 적용 가능해야 한다.

## 확장 원칙

- independence 판단 기준이 바뀌면 guide skill과 함께 갱신한다.

