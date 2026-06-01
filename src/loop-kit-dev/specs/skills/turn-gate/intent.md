# turn-gate 사용자 의도

```mermaid
graph TD
  START[사용자 메시지] --> FLOWGROUP

  subgraph TG[turn-gate / 메인]
    direction TB
    subgraph FLOWGROUP[flow skill]
      direction TB
      FLOW[flow skill: interview]
      HANDOFF[flow skill: handoff]

      FLOW -->|생략...| HANDOFF
    end
    subgraph QR[next-flow gate]
      direction TB
      SKILLREAD[사용중인 스킬 다시 읽기]
      ASK[질문 도구: 다음 플로우 선택]
      PLAN[000-plan.md 업데이트]

      SKILLREAD --> ASK
      ASK --> PLAN
    end

    HANDOFF --> SKILLREAD
    PLAN --> FLOW
  end
```

```mermaid
graph TD
  MAIN[turn-gate / 메인]
  subgraph STOPPHASE[종료 페이즈]
    direction TB
    CLEANUP[작업 중이던 플로우 정리]
    STOP[explicit-stop 기록 - active turn 종료]

    CLEANUP --> STOP
  end

  MAIN -->|모든 시점: 종료 요청| CLEANUP
```

## 핵심

- `turn-gate`는 `flow`를 의존해 사용하는 wrapper입니다.
- 사용자 메시지는 `turn-gate` wrapper 안의 `flow skill` 그룹으로 진입합니다.
- 메인 플로우는 `flow skill` 그룹과 `next-flow gate` 그룹의 루프로 봅니다.
- `flow skill` 그룹은 `flow skill: interview -> 생략... -> flow skill: handoff`를 포함합니다.
- `turn-gate`는 `flow skill` 내부를 재정의하지 않고 `flow` 스킬을 그대로 적용합니다.
- `next-flow gate`는 `사용중인 스킬 다시 읽기 -> 질문 도구: 다음 플로우 선택 -> 000-plan.md 업데이트` 순서입니다.
- `사용중인 스킬 다시 읽기`는 다음 flow 질문을 만들기 전에 현재 사용해야 할 skill context를 복구합니다.
- `질문 도구: 다음 플로우 선택`은 다음 flow 입력을 고르는 question-routing 표면입니다.
- 질문 도구는 `flow: deep-interview`와 같은 인터뷰 흐름으로 다음 flow 입력을 충분히 구체화합니다.
- `000-plan.md 업데이트`는 선택된 다음 flow 입력, 사용 skill, pending/answered question 상태, next action을 매번 반영합니다.
- `flow skill: interview` 재진입 뒤에는 구체화된 입력을 기준으로 flow design에 필요한 질문을 우선합니다.
- self-drive는 그래프 노드가 아니라 준비된 sequence gate가 질문 도구를 대체하는 핵심 경로입니다.
- 종료 플로우는 별도 그래프로 두고, `turn-gate / 메인`의 모든 시점에서 종료 요청이 오면 `종료 페이즈`로 이동합니다.
- 종료 페이즈는 `작업 중이던 플로우 정리 -> explicit-stop 기록 - active turn 종료` 순서입니다.
