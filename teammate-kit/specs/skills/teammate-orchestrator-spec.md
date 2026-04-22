# teammate-orchestrator 스킬 스펙

## 목적

`teammate-orchestrator`는 append-only event bus를 중심으로 durable multi-teammate workflow를 조정하는 orchestration 스킬입니다.

## 경계

- 포함:
  - mention routing과 action coordination
  - checkpoint, retry, DLQ, log compaction 같은 reliability rule
  - long-running collaboration의 상태 추적
- 제외:
  - bounded research/implementation/review 자체
  - direct subagent role definition
  - generic workflow design

## 처리하려는 작업 형태

- durable action history가 필요한 multi-teammate 작업
- main-agent relay 방식과 event bus coordination이 필요한 경우
- 긴 실행에서 retry와 recovery 전략이 필요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `teammate-kit/skills/teammate-orchestrator/SKILL.md`
- 관련 상위 라우팅: `teammate-kit-guide`

## 핵심 처리 계약

- orchestration은 append-only event history와 checkpoint를 유지해야 한다.
- routing rule, reliability rule, monitoring rule을 분리해서 다룬다.
- role execution detail은 각 bounded teammate skill로 위임한다.

## 독립성 원칙

- 이 스킬은 orchestration contract만 소유한다.
- event bus 규칙만 읽어도 durable collaboration 흐름을 이해할 수 있어야 한다.

## 확장 원칙

- reliability rule이 바뀌면 관련 script/reference와 함께 갱신한다.
- direct role logic은 이 스킬 안으로 흡수하지 않는다.

