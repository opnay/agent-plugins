# turn-gate 스킬 스펙

## 목적

`turn-gate`는 하나의 사용자 턴 안에서 `분석 -> 계획 -> 작업 -> 결과 보고 / commit-ready -> 다음 플로우 진행을 위한 사용자 응답`을 명시적으로 이어가고, 사용자가 작업을 명시적으로 끝낼 때까지 턴을 종료하지 않도록 유지하는 loop gate 스킬입니다.

## 사용자 스펙 의도

- `turn-gate`는 하나의 턴에서 사용자가 작업을 명시적으로 끝낼때까지 턴을 종료하지 않고 연속성을 가지도록 하는 루프 게이트입니다.
- `turn-gate`는 연속성을 가지는 루프 게이트이면서, `종료`라는 판단을 사용자에게 넘기는 루프 게이트입니다.
- 이 의도는 결과 보고를 종결로 닫지 않고, 다음 플로우 선택권을 계속 열어두는 방향으로 해석되어야 합니다.
- 다음 플로우 선택지는 현재 결과 보고와 직접 연결된 좁은 옵션이어야 하며, 불필요하게 넓은 재프레이밍으로 흐르지 않아야 합니다.
- 기본 턴 흐름은 `사용자 메시지 -> 분석 -> 계획 -> 작업 -> 결과 보고 / commit-ready -> 다음 플로우 진행을 위한 사용자 응답`으로 이어집니다.
- `turn-gate`는 각 단계가 드러나도록 유지하고, 결과 보고 뒤에는 사용자가 다음 플로우를 고를 수 있게 응답 표면을 열어야 합니다.
- 결과 보고는 종결 멘트가 아니라 다음 플로우 진입을 위한 사용자 응답에 대한 사전 설명 형태로 보고합니다.
- 사용자에게 질문하는 방식은 사용자에게 선택권을 주는 질문 도구를 강제해야합니다.
- `다음 플로우 진행을 위한 사용자 응답` 자체는 다시 `사용자 메시지`로 취급되어야 하며, 같은 턴 안에서 다음 루프의 입력으로 즉시 이어져야 합니다.
- 따라서 사용자가 명시적으로 종료하지 않는 한, `turn-gate`는 한 플로우가 끝날 때마다 다음 플로우로 반복 진입하는 구조를 유지해야 합니다.

예시:

- 아래 예시는 카테고리성 예시입니다.
- 실제 응답이 반드시 `분석:`, `계획:`, `작업:`, `결과 보고:` 같은 literal 라벨을 그대로 출력해야 한다는 의미는 아닙니다.

- 사용자 메시지: "`workflow-kit`의 기본 시작점 뭐야?"
- 분석: 현재 저장소에서 `workflow-kit`의 기본 시작점을 찾아 알려달라는 요청입니다.
- 계획:
  1. `README.md`와 `specs/plugin-spec.md`를 확인합니다.
  2. 결과를 정리합니다.
- 작업: `README.md`와 `specs/plugin-spec.md`를 읽고 기본 시작점을 확인합니다.
- 결과 보고: 현재 저장소 기준 `workflow-kit`의 기본 시작점은 `workflow-kit-guide`입니다.
- 사용자 응답:
  - 다음 플로우는 어떤걸 진행하시나요?
  - 1. `workflow-kit-guide` 역할 점검
  - 2. `turn-gate` 동작 점검
  - 3. `plugin-spec` 라우팅 규칙 확인

## 경계

- 포함:
  - turn-level phase classification
  - current phase의 downstream workflow 선택
  - 결과 보고 뒤 다음 플로우 진행을 위한 사용자 응답 개방
  - 선택권을 주는 질문 도구를 통한 loop 유지
- 제외:
  - deep-interview 자체
  - planner 자체
  - implementation 자체
  - review 자체
  - commit execution 자체

## 처리하려는 작업 형태

- 한 번의 응답으로 끝내면 안 되고 turn loop를 계속 이어가야 하는 작업
- repository-local operating rule이 사용자의 explicit stop 전까지 turn continuity를 요구하는 작업
- loop gate 유지가 phase detail보다 더 중요한 작업
- 결과 보고 뒤 clean stop이 기본값이면 안 되는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/turn-gate/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-guide`

## 핵심 처리 계약

- 현재 메시지를 이번 턴의 분석 대상으로 받아들인다.
- 다음 플로우 진행을 위한 사용자 응답도 같은 턴의 다음 `현재 메시지`로 받아들인다.
- 각 phase는 가장 좁은 downstream workflow에 위임한다.
- 결과 보고 후에는 기본적으로 다음 플로우 진행을 위한 사용자 응답 표면을 연다.
- user explicit stop이 없는 한 clean stop을 기본 경로로 두지 않는다.
- 메타 플로우는 유지하되 phase-specific detail은 sibling skill에 남긴다.
- summary-only closing과 generic follow-up phrase를 정상 종료 형태로 취급하지 않는다.

## 독립성 원칙

- 이 스킬은 turn-level loop gate만 소유한다.
- phase-specific workflow를 hidden sibling context 없이 호출 가능한 상태로 유지해야 한다.

## 확장 원칙

- 새로운 rule은 turn continuity와 next-flow user-response gating을 더 명확하게 만들 때만 추가한다.
- stage routing이나 default loop-gate rule이 바뀌면 `workflow-kit-guide`와 `plugin-spec.md`를 함께 갱신한다.
