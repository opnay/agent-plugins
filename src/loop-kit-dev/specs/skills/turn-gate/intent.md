# turn-gate 사용자 의도

- turn-gate를 사용할 때 flow 흐름이 `next-flow gate`를 사용하지 않는 문제가 있다. `flow skill: handoff` 뒤에는 일반 flow 종료처럼 닫히지 않고 반드시 `next-flow gate`로 들어가야 한다.
- 스킬 작성에서는 `<gate:next-flow>` 같은 태그 작성법으로 gate를 감싸는 방식이 더 강하다. next-flow gate는 태그로 감싼 실행 계약으로 드러나야 한다.
- `skill reconfigure` 같은 하위 gate도 태그 방식으로 표현할 수 있다. next-flow gate 안에서 `skill reconfigure` 경계도 태그로 감싸야 한다.

## 전체 루프

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
      SKILLRECONFIG[skill reconfigure]
      NEXTFLOW[다음 플로우 선택]
      SELFDRIVE[000-self-drive.md 업데이트]
      PLAN[000-plan.md 업데이트]

      SKILLRECONFIG -->|질문 도구| NEXTFLOW
      SKILLRECONFIG -->|self-drive| NEXTFLOW
      NEXTFLOW --> PLAN
      NEXTFLOW -->|self-drive| SELFDRIVE
      SELFDRIVE --> PLAN
    end

    HANDOFF --> SKILLRECONFIG
    PLAN --> FLOW
  end
```

## skill reconfigure

```mermaid
graph TD
  HANDOFF[flow skill: handoff] --> IDENTIFY
  subgraph SKILLRECONFIG[skill reconfigure]
    direction TB
    IDENTIFY[세션에서 사용중인 전체 skill 목록 식별]
    REREAD[각 skill 본문 새로 읽기]
    ACCEPT[새 active skill set으로 수용]
    RECORD[000-plan.md에 active skills 기록 준비]

    IDENTIFY --> REREAD
    REREAD --> ACCEPT
    ACCEPT --> RECORD
  end
  RECORD --> NEXTFLOW[다음 플로우 선택]
```

## 종료 플로우

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

### wrapper 경계

- `turn-gate`는 `flow`를 의존해 사용하는 wrapper입니다.
- 사용자 메시지는 `turn-gate` wrapper 안의 `flow skill` 그룹으로 진입합니다.
- `flow skill` 그룹은 `flow skill: interview -> 생략... -> flow skill: handoff`를 포함합니다.
- `flow skill: handoff`는 terminal closure가 아니라 `next-flow gate` 진입 조건입니다.
- `turn-gate`는 `flow skill` 내부를 재정의하지 않고 `flow` 스킬을 그대로 적용합니다.

### next-flow gate

- `next-flow gate`는 `flow skill: handoff -> skill reconfigure 그룹 -> 다음 플로우 선택 -> 000-plan.md 업데이트`를 기본 경로로 둡니다.
- `next-flow gate`는 handoff 직후에 먼저 실행되며, final-looking 보고나 flow-only 종료보다 우선합니다.
- runtime skill은 `next-flow gate` 실행 계약을 `<gate:next-flow>...</gate:next-flow>`로 감싸 다음 gate 경계를 강하게 드러냅니다.
- `skill reconfigure` 그룹은 `<gate:skill-reconfigure>...</gate:skill-reconfigure>`로 감싸며, 다음 flow 질문을 만들기 전에 세션에서 사용중인 전체 skill 목록을 식별하고, 각 skill 본문을 새로 읽고, 이전 대화의 stale skill context가 아니라 새 active skill set으로 수용합니다.
- `다음 플로우 선택`은 질문 도구 또는 self-drive로 다음 flow 입력을 고르는 question-routing 표면입니다.
- 질문 도구는 `flow: deep-interview`와 같은 인터뷰 흐름으로 다음 flow 입력을 충분히 구체화합니다.
- self-drive 진입은 질문 도구를 사용하지 않습니다.
- self-drive 경로는 `다음 플로우 선택` 뒤 `000-self-drive.md 업데이트`를 거쳐 통합 `000-plan.md 업데이트`로 들어갑니다.
- `flow skill: interview` 재진입 뒤에는 구체화된 입력을 기준으로 flow design에 필요한 질문을 우선합니다.

### 표시와 종료

- phase label은 사용자-facing phase 시작 또는 의미 있는 진행 메시지에 붙입니다.
- `turn-gate`는 `flow`가 산출한 phase label을 재정의하지 않고 적용합니다. `next-flow gate`를 열 때는 `turn-gate`가 `[next-flow]` label을 소유합니다.
- phase label은 진행 표시이며 artifact 본문, record 본문, command output summary, 질문 option label에 전파하지 않습니다.
- self-drive는 그래프 노드가 아니라 준비된 sequence gate가 질문 도구를 대체하는 핵심 경로입니다.
- 종료 플로우는 별도 그래프로 두고, `turn-gate / 메인`의 모든 시점에서 종료 요청이 오면 `종료 페이즈`로 이동합니다.
- 종료 페이즈는 `작업 중이던 플로우 정리 -> explicit-stop 기록 - active turn 종료` 순서입니다.
