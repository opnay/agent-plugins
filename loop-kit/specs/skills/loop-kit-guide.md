## 사용자 스펙 의도

- `loop-kit`이 현재 작업의 시작점이 맞는지 빠르게 판단하고 싶다.
- loop continuity가 핵심 계약이면 바로 `turn-gate`로 들어가고 싶다.
- requirement discovery가 필요해도 turn continuity가 top-level contract면 `turn-gate` 안에서 처리하고 싶다.
- `ralph-loop`, `review-loop` 같은 loop mode를 직접 고르기보다 `turn-gate`가 내부적으로 선택하게 하고 싶다.
- 사용자 개입 없이 계속 진행해야 하면 `self-drive` 질문 라우팅으로 subagent에게 질문하게 하고 싶다.

---

# loop-kit-guide 스킬 스펙

## 목적

`loop-kit-guide`는 현재 요청이 `loop-kit`의 narrow loop surface로 시작해야 하는지 판단하고, 맞다면 `turn-gate`로 진입시키는 entrypoint skill입니다.

## 경계

- 포함:
  - `loop-kit` 적합성 판단
  - `turn-gate` 시작 여부 판단
  - direct loop entrypoint를 열지 않는 라우팅
- 제외:
  - turn loop 자체의 실행
  - internal loop mode의 세부 수행
  - broad workflow taxonomy 결정 전반

## 처리하려는 작업 형태

- 사용자가 턴을 종료하자고 요청하기 전까지 turn continuity가 요구되는 작업
- current-phase work를 `turn-gate` 안의 internal mode로 처리해야 하는 작업
- requirement discovery도 같은 turn loop 안에서 이어가야 하는 작업
- `workflow-kit`의 broader workflow 대신 narrow loop package로 바로 들어가야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit/skills/loop-kit-guide/SKILL.md`
- 관련 상위 라우팅: 없음

## 핵심 처리 계약

- 먼저 loop continuity가 top-level governing contract인지 판단한다.
- 그렇다면 `turn-gate`를 main surface로 활성화한다.
- 사용자에게 `ralph-loop`, `review-loop` 같은 direct loop entrypoint를 제안하지 않는다.
- turn continuity가 핵심 계약이 아니면 `loop-kit`보다 `workflow-kit` 또는 다른 plugin이 더 적합할 수 있음을 명시한다.

## 라우팅 규칙

- repository-local rule이나 task shape가 non-terminal turn을 요구하면 `turn-gate`로 시작한다.
- current-phase work가 requirement discovery, autonomous execution, refinement, review handling, readiness pass 중 하나로 좁혀질 수 있으면 `turn-gate`로 보낸다.
- 사용자 질문을 subagent 질문으로 바꿔 자동 진행해야 하는 self-driving task라면 `turn-gate`로 보낸다.
- broad workflow selection 자체가 먼저 필요하면 `workflow-kit` 쪽 시작점을 우선 검토한다.

## 검토 질문

- 지금 진짜로 loop continuity가 top-level contract인가?
- direct loop skill을 사용자 표면으로 열지 않고 `turn-gate`로 진입시켰는가?
- `loop-kit`보다 broader workflow router가 먼저여야 하는 상황을 놓치지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `turn-gate`가 메인 실행 표면이라는 plugin context를 전제로 한다. 다만 `loop-kit`로 들어가야 하는 판단 기준 자체는 이 스펙만 읽어도 이해 가능해야 한다.

## 확장 원칙

- 라우팅 규칙을 늘릴 때는 `loop-kit`의 narrow scope를 흐리지 않는 범위에서만 추가한다.
- `turn-gate`의 책임이 바뀌면 이 guide와 plugin spec을 함께 갱신한다.
