## 사용자 스펙 의도

- 들어온 요청에서 현재 bottleneck이 무엇인지 먼저 분류하고 싶다.
- 어떤 workflow skill이 시작점이어야 하는지와 handoff 순서를 같이 알고 싶다.
- `turn-gate`가 execution mode가 아니라 turn-level loop gate인지 함께 판단하고 싶다.
- 사용자 질문을 subagent 질문으로 바꾸는 `self-drive` question-routing이 필요한지도 함께 판단하고 싶다.

---

# workflow-kit-guide 스킬 스펙

## 목적

`workflow-kit-guide`는 들어온 요청의 현재 bottleneck을 기준으로 어떤 workflow skill이 시작점이 되어야 하는지 결정하는 global first-read router입니다.

## 경계

- 포함:
  - current bottleneck 분류
  - starting workflow와 handoff 결정
  - execution mode 선택 기준 제시
  - `turn-gate`가 turn-level loop gate로 유지되어야 하는지 판단
  - `self-drive` question-routing 필요 여부 판단
- 제외:
  - 각 workflow의 상세 실행
  - specialist plugin의 세부 판단

## 처리하려는 작업 형태

- 요청을 framing, interview, planning, execution, review, gate, turn-level meta flow 중 어디서 시작해야 하는지 모르는 경우
- execution mode를 autopilot, parallel-work, ralph-loop, review-loop 중에서 골라야 하는 경우
- execution mode와 별도로 `turn-gate`가 turn-level loop gate로 유지되어야 하는지 판단해야 하는 경우
- 사용자 개입 없이 계속 진행해야 해서 질문 대상이 user가 아니라 subagent여야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/workflow-kit-guide/SKILL.md`

## 핵심 처리 계약

- current bottleneck을 먼저 하나로 좁힌다.
- starting skill과 planned handoff를 함께 제시한다.
- specialist plugin은 workflow 이후 handoff 대상으로만 둔다.
- phase continuity가 핵심 병목이면 `turn-gate`를 명시적으로 활성화하거나 유지한다.
- repository-local rule이 non-terminal turn을 요구하면 `turn-gate`를 기본 loop gate로 유지한다.
- `turn-gate`를 활성화한다는 것은 현재 세션 동안 이를 first-class loop gate rule로 유지한다는 의미로 취급한다.
- 사용자 질문을 subagent 질문으로 바꿔야 하면 `turn-gate`의 `self-drive` question-routing mode를 선택한다.
- `turn-gate`는 execution mode가 아니라 turn-level gate contract로 분류한다.
- `turn-gate`가 활성화된 상태의 사용자 응답 또는 `self-drive` subagent 답변은 같은 턴의 다음 메시지로 이어지는 것으로 취급한다.

## 검토 질문

- current bottleneck을 정말 하나로 좁혔는가?
- starting skill과 planned handoff가 함께 제시돼 있는가?
- `turn-gate`를 execution mode가 아니라 turn-level gate contract로 분리해서 판단했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 스킬은 sibling workflow map을 전제로 하는 guide이므로 그 문맥은 허용하지만, 라우팅 기준과 handoff 판단 근거는 이 스펙만 읽어도 이해 가능해야 한다.

## 확장 원칙

- stage model, execution mode map, 또는 turn continuity activation rule이 바뀌면 plugin spec과 함께 갱신한다.
