## 사용자 스펙 의도

- `turn-gate`는 하나의 턴에서 사용자가 턴을 종료하자고 요청할때까지 턴을 종료하지 않고 연속성을 가지도록 하는 루프 게이트입니다.
- 이 스킬을 사용한다는건, 이 세션동안 이 스킬을 1급 규칙으로 사용한다는 의미입니다.
- `turn-gate`는 연속성을 가지는 루프 게이트이면서, `종료`라는 판단을 사용자에게 넘기는 루프 게이트입니다.
- 이 의도는 결과 보고를 종결로 닫지 않고, 다음 플로우 선택권을 계속 열어두는 방향으로 해석되어야 합니다.
- 다음 플로우 선택지는 현재 결과 보고와 직접 연결된 좁은 옵션이어야 하며, 불필요하게 넓은 재프레이밍으로 흐르지 않아야 합니다.
- 기본 턴 흐름은 `사용자 메시지 -> 분석 -> 계획 -> 작업 -> 검증 -> 결과 보고 / commit-ready -> 다음 플로우 진행을 위한 question-routing 응답`으로 이어집니다.
- 분석 단계에서는 사용자 메시지를 구조 분해하여, 사용자의 요청 의도와 요청 행동을 정리해야 합니다.
- 계획 단계에서는 분석 단계에서 정리한 요청을 작업하기 위한 상세 계획을 준비해야 합니다.
- 작업 단계에서는 준비한 계획을 실행해야 합니다.
- 검증 단계에서는 작업 결과를 확인하고 남은 불확실성을 드러내야 합니다.
- 결과 보고 단계에서는 완료된 작업에 대한 결과를 보고해야 합니다.
- `turn-gate`는 각 단계가 드러나도록 유지하고, 결과 보고 뒤에는 사용자가 다음 플로우를 고를 수 있게 응답 표면을 열어야 합니다.
- 결과 보고는 종결 멘트가 아니라 다음 플로우 진입을 위한 question-routing 응답에 대한 사전 설명 형태로 보고합니다.
- 기본 user-gated question routing에서는 사용자에게 질문하는 방식으로 선택권을 주는 질문 도구를 강제해야 합니다.
- 분석 단계와 계획 단계에서는 사용자에게 질문하는 과정이 필요할 수 있습니다.
- `다음 플로우 진행을 위한 question-routing 응답` 자체는 다시 현재 메시지로 취급되어야 하며, 같은 턴 안에서 다음 루프의 입력으로 즉시 이어져야 합니다.
- 따라서 사용자가 턴을 종료하자고 요청하지 않는 한, `turn-gate`는 한 플로우가 끝날 때마다 다음 플로우로 반복 진입하는 구조를 유지해야 합니다.
- 다음 플로우 질문의 사용자 표시 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하더라도, sessions flow record의 `Next Flow Options`에는 명시적인 턴 종료 선택지가 항상 남아야 합니다.
- `turn-gate`로 진행한 작업은 `.agents/sessions` 아래에 기록이 남아야 합니다.
- 여러 플로우를 거치는 작업의 상위 계획은 `.agents/sessions/{YYYYMMDD}/000-plan.md` 경로에 증분되어야 합니다.
- 개별 플로우 기록은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식으로 남아야 합니다.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`는 phase 메모가 아니라 flow 기록으로 남아야 합니다.
- 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 필요하면 미리 설계할 수 있어야 합니다.
- 이후 loop에서 다시 분석 단계나 계획 단계로 돌아오면, 이전 flow/phase 설계를 고정값처럼 유지하지 말고 필요할 때만 다시 설계할 수 있어야 합니다.
- 검증 단계는 그 재설계를 직접 수행하는 단계라기보다, 이후 flow/phase 재설계가 필요한지 여부를 드러내는 단계여야 합니다.

예시:

- 아래 예시는 카테고리성 예시입니다.
- 실제 응답이 반드시 `분석:`, `계획:`, `작업:`, `결과 보고:` 같은 literal 라벨을 그대로 출력해야 한다는 의미는 아닙니다.

- 사용자 메시지: "`workflow-kit-dev`의 기본 시작점 뭐야?"
- 분석: 현재 저장소에서 `workflow-kit-dev`의 기본 시작점을 찾아 알려달라는 요청입니다. 시작점을 어떤 기준으로 봐야 할지 확인이 필요합니다.
- 계획:
  1. 사용자에게 기준 선택지를 엽니다.
  2. 선택된 기준에 맞는 문서를 확인합니다.
- 작업: 기준 선택을 위한 질문 도구를 준비합니다.
- 결과 보고: 지금은 기본 시작점을 하나로 단정하기보다, 어떤 기준으로 볼지 먼저 맞추는 게 안전합니다.
- 사용자 응답:
  - 시작점을 어떤 기준으로 볼까요?
  - 1. 플러그인 관점
  - 2. `AGENTS.md` 관점
  - 3. 스킬 관점

---

# turn-gate 스킬 스펙

## 목적

`turn-gate`는 하나의 사용자 턴 안에서 `분석 -> 계획 -> 작업 -> 검증 -> 결과 보고 / commit-ready -> 다음 플로우 진행을 위한 question-routing 응답`을 명시적으로 이어가고, 사용자가 턴을 종료하자고 요청할때까지 턴을 종료하지 않도록 유지하는 loop gate 스킬입니다.

## 경계

- 포함:
  - turn-level phase classification
  - current phase의 downstream workflow 선택
  - 결과 보고 뒤 다음 플로우 진행을 위한 question-routing 응답 개방
  - 선택권을 주는 user-gated question routing을 통한 loop 유지
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

- 대표 표면: `workflow-kit-dev/skills/turn-gate/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

### 활성화와 phase 흐름

- 현재 메시지를 이번 턴의 분석 대상으로 받아들인다.
- 중간 사용자 메시지를 stop, completion, approval-boundary pause로 해석하지 않고 authoritative loop input으로 받아들인다.
- 중간 사용자 메시지는 explicit turn stop, current-flow correction, current-flow priority change, next-flow priority request 중 하나로 분류한다.
- current-flow correction 또는 current-flow priority change라면 현재 analysis/plan을 즉시 조정하고 가장 이른 안전한 phase부터 이어간다.
- next-flow priority request라면 flow record의 next-flow 후보 중 최우선으로 등록하고 다음 safe handoff point까지 이어간다.
- 이 skill이 사용되면 현재 세션 동안 `turn-gate`를 first-class loop gate rule로 활성화한 것으로 취급한다.
- 분석 단계에서는 사용자 메시지를 구조 분해해 요청 의도와 요청 행동을 정리한다.
- 분석 단계는 필요하면 이후 이어질 future flow/phase 후보까지 미리 설계할 수 있다.
- 계획 단계에서는 분석 단계에서 정리한 요청을 작업하기 위한 상세 계획과, 필요하면 이후 flow/phase를 위한 provisional 설계를 준비한다.
- 작업 단계에서는 준비한 계획을 실행한다.
- 검증 단계에서는 작업 결과를 확인하고 남은 불확실성을 드러내며, 이후 flow/phase 재설계가 필요한지 여부를 surface한다.
- 결과 보고 단계에서는 완료된 작업의 결과를 보고한다.

### 세션 기록과 Continuity Guard

- `.agents/sessions/{YYYYMMDD}/000-plan.md`는 날짜 기준 plan artifact이며, 당일 turn-gated 작업의 히스토리, 사용자 요청 목록, flow index, 현재 계획, 완료 flow 요약을 소유한다.
- `000-plan.md`는 증분 갱신되어야 하고, 완료된 작업도 삭제하지 않고 요약과 flow reference를 유지한다.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` record는 사용자 요청에 따른 개별 flow의 상세 보고서이며, 해당 flow의 analysis, plan, work, verification, result report를 구체적으로 남긴다.
- 상세 flow record는 completed flow를 기다리지 말고 각 phase가 끝날 때마다 증분 갱신되어야 한다.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` record의 최소 항목에는 user request message와 question-routing mode가 포함되어야 한다.
- flow record의 `Next Flow Options`에는 사용자 표시 질문에 턴 종료 선택지가 보이지 않는 경우에도 명시적인 turn-end option이 포함되어야 한다.
- 각 flow record에는 짧은 `Continuity Guard`가 있어야 하며 result reporting과 next-flow reopening 직전에 갱신되어야 한다.
- `Continuity Guard`에는 `turn-gate` 활성 여부, question-routing mode, user explicit stop 여부, terminal summary 허용 여부, required next action이 포함되어야 한다.
- default flow-record template는 `skills/turn-gate/templates/flow-record-template.md`여야 한다.
- default `000-plan.md` template는 `skills/turn-gate/templates/plan-template.md`여야 한다.
- flow record는 phase 메모가 아니라 상세 보고서이지만 각 phase 종료 시점의 현재 상태를 증분 반영해야 한다.

### mode selection과 phase 위임

- current phase의 downstream workflow 선택에는 최소한 `deep-interview`, `autopilot`, `review-loop`, `ralph-loop`, `commit-readiness-gate` 구분 신호가 드러나야 한다.
- 분석 단계와 계획 단계에서는 필요하면 user-gated question routing으로 질문을 열 수 있다.
- future flow/phase 설계는 고정값이 아니며, 이후 loop에서 새 증거, changed intent, 새 blocker가 생겼을 때만 다시 설계한다.
- 각 phase는 가장 좁은 downstream workflow에 위임한다.
- 메타 플로우는 유지하되 phase-specific detail은 sibling skill에 남긴다.

### question routing

- 기본 question routing은 `user-gated`이고, 사용자 선택지, scope lock, next-flow decision은 질문 도구로 묻는다.
- explicit user, tool, platform, safety, destructive, irreversible, external-action approval boundary는 반드시 사용자 질문으로 유지한다.

### next-flow reopening

- 다음 플로우 진행을 위한 `user-gated` 사용자 응답도 같은 턴의 다음 `현재 메시지`로 받아들인다.
- 결과 보고 전에는 `Continuity Guard`를 읽거나 재구성하고, 사용자가 명시적으로 종료하지 않았으면 terminal summary가 invalid임을 확인한다.
- 결과 보고 후에는 기본적으로 다음 플로우 진행을 위한 question-routing 응답 표면을 연다.
- 사용자에게 보이는 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하는 경우에도, flow record의 `Next Flow Options`에는 별도 turn-end option을 기록한다.
- user explicit stop이 없는 한 clean stop을 기본 경로로 두지 않는다.
- summary-only closing과 generic follow-up phrase를 정상 종료 형태로 취급하지 않는다.

## deep-interview 원본 관계

- 참고한 원본 skill source는 `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`다.
- 원본은 requirement discovery를 ambiguity gating, OMX tooling, artifact handoff까지 포함한 큰 workflow로 다룬다.
- 여기서 `deep-interview`는 그 full workflow 전체가 아니라 requirement-discovery lane을 가리키는 workflow 신호이며, turn-gate 라우팅에 필요한 해석만 소유한다.

## 검토 질문

- 이번 응답이 `분석 -> 계획 -> 작업 -> 검증 -> 결과 보고`를 visible shape로 유지하고 있는가?
- cross-flow task라면 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 최신 상태인가?
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` record가 현재 phase까지 증분 갱신됐는가?
- 결과 보고 뒤에 explicit choice가 있는 다음 플로우 질문을 실제로 열었는가?
- 사용자 표시 선택지에 턴 종료 option이 없더라도 flow record의 `Next Flow Options`에 명시적인 turn-end option을 남겼는가?
- 결과 보고 직전에 `Continuity Guard`를 갱신했고 terminal summary 가능 여부를 확인했는가?
- clean stop, summary-only closing, generic follow-up phrase로 턴을 닫고 있지 않은가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `workflow-kit-dev-guide` 또는 repository-local operating rule이 이미 turn continuity를 활성화한 문맥을 허용한다. 다만 turn-level loop gate contract 자체와 phase-specific workflow에 무엇을 위임하는지는 이 스펙에서 계속 명시적으로 읽혀야 한다.

## 확장 원칙

- 새로운 rule은 turn continuity와 next-flow question-routing gating을 더 명확하게 만들 때만 추가한다.
- stage routing이나 default loop-gate rule이 바뀌면 `workflow-kit-dev-guide`와 `plugin.md`를 함께 갱신한다.
