# flow 사용자 의도

## 전체 구조

```mermaid
graph TD
  MESSAGE[메시지] --> ENTRY_RECONFIG

  subgraph FLOW[flow]
    direction TB
    ENTRY_RECONFIG[entry skill reconfigure - 000-plan.md] --> INTERVIEW
    INTERVIEW[deep-interview - 000-plan.md] --> DESIGN[플로우 설계]
    DESIGN --> MAIN

    subgraph MAIN[메인 플로우]
      direction TB
      INTAKE[intake - flow record] --> FRAMING[framing - flow record]
      FRAMING --> PREP[preparation - flow record]
      PREP --> WORK[work - flow record]
      WORK --> VERIFY[verification - flow record]
      VERIFY --> REPORT[reporting - flow record]
      REPORT -->|다음 플로우| POST_REPORT_RECONFIG[post-reporting skill reconfigure - 000-plan.md]
      POST_REPORT_RECONFIG --> INTAKE
    end

    MAIN --> REVIEW[메인 플로우 회고 - 000-review.md]
    REVIEW --> HANDOFF[handoff condition - flow record]
  end
```

## 메시지 인터뷰

```mermaid
graph TD
  MESSAGE[메시지] --> ENTRY_RECONFIG[entry skill reconfigure]
  ENTRY_RECONFIG --> DEEP_INTERVIEW

  subgraph DEEP_INTERVIEW[deep-interview skill]
    direction TB
    SNAPSHOT[초기 의도 스냅샷 - 원하는 결과/대상/범위/제약] --> RISK[alignment risk 식별 - 000-plan.md]
    RISK --> ASK[high-leverage 질문]
    ASK --> ANSWER[답변 반영]
    ANSWER --> TEST[예시/반례/비목표/tradeoff 압력 테스트]
    TEST -->|검증 실패| RISK
    TEST --> BRIEF[locked execution brief - 목적/대상/범위/비목표/완료기준/검증기대/승인경계/근거 - 000-plan.md]
  end

  DEEP_INTERVIEW --> DESIGN[플로우 설계]
```

## 플로우 설계

```mermaid
graph TD
  BRIEF[locked execution brief] --> LIST_UP

  subgraph LIST_UP[list-up]
    direction TB
    CLASSIFY[항목 분류] --> UNIT[단일/다중 flow 판단]
    UNIT --> PARENT[parent flow 또는 단일 active flow 식별]
    PARENT --> CANDIDATES[sub-flow candidate 추출]
    CANDIDATES --> PENDING[후보 pending 상태 표시]
    PENDING --> CONTRACT[flow별 계약 작성 - flow record 작성 - 목적 사슬 필요 시]
    CONTRACT --> ORDER[진행 순서 정리 - 000-plan.md 갱신]
    ORDER -->|검증 실패| CLASSIFY
  end

  LIST_UP --> MAIN[메인 플로우 선택]
```

## 핵심

- 메시지가 들어오면 `flow`는 먼저 entry `skill reconfigure`로 현재 flow에 필요한 skill 본문을 source에서 다시 읽고, 그다음 메시지를 해석하고 실제 진행할 flow를 함께 설계합니다.
- 다음 main flow가 있으면 `flow`는 `reporting` 직후 post-reporting `skill reconfigure`로 다음 main flow에 필요한 skill context를 복구합니다.
- `skill reconfigure`는 루프가 이어지는 동안 잊혔거나 오래된 skill context를 복구하기 위해 flow entry와 post-reporting continuation boundary에서 실행합니다.
- 실제 진행할 플로우가 정해지면 각 메인 플로우는 `intake -> framing -> preparation -> work -> verification -> reporting`으로 진행합니다.
- 다음 flow가 있으면 `reporting` 직후 skill reconfigure를 수행하고 다음 `intake`로 라우팅합니다.
- 메인 플로우 그룹 이후 결과는 `메인 플로우 회고 -> handoff condition`입니다.
- `handoff condition`은 메인 플로우 종료 뒤 산출되는 종료 조건입니다.
- 여러 플로우가 필요하면 리스트업 결과가 여러 메인 플로우가 될 수 있습니다.
- 메시지 인터뷰는 별도 `deep-interview` skill을 필수로 적용합니다.
- `deep-interview`는 alignment risk를 식별하고, 필요한 high-leverage 질문으로 답변을 압력 테스트합니다.
- 답변이 여전히 모호하면 같은 alignment risk를 다시 좁힙니다.
- 초기 의도 스냅샷은 원하는 결과, 대상, 범위, 제약을 드러냅니다.
- alignment risk는 locked execution brief를 실행 입력으로 쓰기 어렵게 만드는 가장 큰 불확실성입니다.
- high-leverage 질문은 하나의 alignment risk를 좁히기 위해 사용합니다.
- locked execution brief는 목적, 대상, 범위, 비목표, 완료 기준, 검증 기대, 승인 경계, 근거, 해소된 alignment risk, 남은 모호성을 현재 확정 상태로 남깁니다.
- 모든 사용자 메시지는 같은 entry skill reconfigure, `deep-interview`, 플로우 설계 경로를 거쳐 메인 플로우로 처리합니다.
- 메시지 인터뷰가 충분히 잠긴 brief를 만들면 사용자 질문 없이 플로우 설계로 진행합니다.
- 플로우 설계는 진행할 flow 구성을 만들고 바로 메인 flow `intake`로 들어갑니다.
- `000-plan.md`와 flow record는 각 그래프 노드의 업데이트 시점으로 표시합니다.

## 플로우 설계 핵심

- 항목 분류는 active flow, parent flow, sub-flow candidate, phase, handoff를 구분합니다.
- flow 분해는 단일/다중 flow 판단, parent flow 또는 단일 active flow 식별, sub-flow candidate 추출, 후보 pending 상태 표시로 나눕니다.
- 후보는 다음 메인 flow로 선택되기 전까지 실행 권한을 만들지 않습니다.
- flow별 계약 작성은 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 정리하고, 선택된 active flow의 flow record를 작성합니다.
- 진행 순서 정리는 다음에 들어갈 메인 flow와 이후 후보를 구분하고 `000-plan.md` 갱신 시점입니다.
- flow record는 flow별 계약 작성 시점에 만들어지고, 이후 메인 플로우 phase, evidence, verification, reporting 시점에서 갱신합니다.
- `000-review.md`는 메인 플로우 그룹 이후, `handoff condition` 직전에 항상 갱신하고, active routing이나 handoff authority로 쓰지 않습니다.
- 회고 finding이 없으면 no-finding 결과로 짧게 기록합니다.
- 목적 사슬은 contract에 영향을 줄 때 `000-plan.md`의 목적 섹션에 흡수합니다.

## Phase 단위

- phase는 선택된 active flow를 실행하기 위한 내부 진행 단위입니다.
- `intake`, `framing`, `preparation`, `work`, `verification`, `reporting`은 고정된 phase 집합입니다.
- phase는 active flow의 현재 위치, 다음 행동, 기록 갱신 시점을 드러냅니다.
- phase는 독립 실행 단위가 아니며 flow 후보, handoff, 회고, next-flow 선택지를 대체하지 않습니다.
- phase가 별도 검토 가능한 산출물, 완료 기준, 승인 경계, handoff 조건을 갖기 시작하면 flow 설계에서 새 flow나 sub-flow candidate로 다시 분류합니다.
- `framing`은 사용자 요구사항을 requirement verification todo로 분해합니다.
- `preparation`은 requirement verification todo와 implementation verification todo를 확정합니다.
- requirement verification은 결과물이 사용자 요구사항을 충족하는지 확인하는 primary axis입니다.
- implementation verification은 타입, 테스트, 빌드, lint, import path, release surface, 코드베이스 관례를 확인하는 supporting axis입니다.
- 사용자에게 보이는 메시지가 특정 phase에서 나오면 `[<phase-name>]` 형식의 메시지 앞 라벨을 붙입니다.
- phase 라벨은 진행 표시입니다. 산출물 본문, 기록 본문, 명령 요약, 질문 선택지 라벨에 기계적으로 전파하지 않습니다.
