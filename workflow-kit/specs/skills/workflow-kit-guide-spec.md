# workflow-kit-guide 스킬 스펙

## 목적

`workflow-kit-guide`는 들어온 요청의 현재 bottleneck을 기준으로 어떤 workflow skill이 시작점이 되어야 하는지 결정하는 global first-read router입니다.

## 경계

- 포함:
  - current bottleneck 분류
  - starting workflow와 handoff 결정
  - execution mode 선택 기준 제시
  - `turn-gate`가 turn-level loop gate로 유지되어야 하는지 판단
- 제외:
  - 각 workflow의 상세 실행
  - specialist plugin의 세부 판단

## 처리하려는 작업 형태

- 요청을 framing, interview, planning, execution, review, gate, turn-level meta flow 중 어디서 시작해야 하는지 모르는 경우
- execution mode를 autopilot, parallel-work, ralph-loop, review-loop 중에서 골라야 하는 경우
- execution mode와 별도로 `turn-gate`가 turn-level loop gate로 유지되어야 하는지 판단해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/workflow-kit-guide/SKILL.md`

## 핵심 처리 계약

- current bottleneck을 먼저 하나로 좁힌다.
- starting skill과 planned handoff를 함께 제시한다.
- specialist plugin은 workflow 이후 handoff 대상으로만 둔다.
- phase continuity가 핵심 병목이면 `turn-gate`를 명시적으로 고른다.
- repository-local rule이 non-terminal turn을 요구하면 `turn-gate`를 기본 loop gate로 유지한다.
- `turn-gate`는 execution mode가 아니라 turn-level gate contract로 분류한다.
- `turn-gate`가 활성화된 상태의 사용자 응답은 같은 턴의 다음 메시지로 이어지는 것으로 취급한다.

## 독립성 원칙

- 이 스킬은 lifecycle routing만 소유한다.
- guide output만 읽어도 왜 그 workflow가 시작점인지 이해 가능해야 한다.

## 확장 원칙

- stage model, execution mode map, 또는 turn continuity activation rule이 바뀌면 plugin spec과 함께 갱신한다.
