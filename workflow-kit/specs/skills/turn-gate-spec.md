# turn-gate 스킬 스펙

## 목적

`turn-gate`는 하나의 사용자 턴 안에서 question, plan, command, execution, result reporting, next-flow prompting을 명시적으로 이어가는 meta workflow 스킬입니다.

## 경계

- 포함:
  - turn-level phase classification
  - current phase의 downstream workflow 선택
  - 결과 보고 뒤 continuation 판단
  - 다음 플로우 질문을 통한 loop 유지
- 제외:
  - deep-interview 자체
  - planner 자체
  - implementation 자체
  - review 자체
  - commit execution 자체

## 처리하려는 작업 형태

- 한 번의 응답으로 끝내면 안 되고 phase loop를 계속 이어가야 하는 작업
- repository-local operating rule이 explicit next-flow continuity를 요구하는 작업
- 메타 플로우 유지가 phase detail보다 더 중요한 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/turn-gate/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- 현재 메시지를 question, plan, command, mixed concern으로 분리한다.
- 각 phase는 가장 좁은 downstream workflow에 위임한다.
- 결과 보고 후 다음 플로우가 필요하면 soft closing 대신 explicit next-flow prompt를 연다.
- 메타 플로우는 유지하되 phase-specific detail은 sibling skill에 남긴다.

## 독립성 원칙

- 이 스킬은 turn continuity만 소유한다.
- phase-specific workflow를 hidden sibling context 없이 호출 가능한 상태로 유지해야 한다.

## 확장 원칙

- 새로운 rule은 turn continuity와 next-flow gating을 더 명확하게 만들 때만 추가한다.
- stage routing이 바뀌면 `workflow-kit-guide`와 `plugin-spec.md`를 함께 갱신한다.
