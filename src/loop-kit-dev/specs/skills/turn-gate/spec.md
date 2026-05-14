# turn-gate 스킬 스펙

## 목적

`loop-kit-dev`의 `turn-gate`는 turn continuity를 유지하면서 기본적으로 implicit operating state로 동작하고, current-phase work에 맞는 phase protocol을 적용하는 메인 loop controller입니다.
이 skill이 활성화되면 `turn-gate` 메인 플로우는 대화 응답 자체의 1급 제어 규칙이 되며, 사용자의 explicit stop 전까지 결과 보고를 terminal response로 닫지 않습니다.

## 경계

- 포함:
  - turn-level continuity 유지
  - `preparation -> work -> verification -> reporting -> next-flow question-routing response` 구조 유지
  - explicit stop lifecycle handling
  - preparation에서 사용자 메시지의 operation 의미 해독
  - implicit default state and phase protocol selection
  - 결과 보고 뒤 explicit choice 기반 next-flow reopening
  - session record와 Continuity Guard 유지
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 여러 direct loop skill을 사용자에게 노출하는 일
  - domain-specific implementation detail 자체
  - commit execution, push, PR 같은 외부 작업의 세부 실행 계약

## 처리하려는 작업 형태

- 사용자가 턴을 종료하자고 요청하기 전까지 한 턴 안에서 여러 phase를 이어가야 하는 작업
- requirement discovery, autonomous execution, refinement, review-driven correction, readiness checking 같은 current-phase work가 번갈아 나타나는 작업
- 사용자 메시지에서 시작한 preparation과 기존 flow 실행 전 preparation을 구분하고, 준비 뒤 작업, 검증, 보고로 이어져야 하는 작업
- 결과 보고 뒤 clean stop이 아니라 다음 플로우 선택이 기본이어야 하는 작업
- 사용자 지시어가 여러 구조 단위를 가리켜 작업 전에 target 또는 operation을 잠가야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit-dev/skills/turn-gate/SKILL.md`
- 기본 skill spec index: `loop-kit-dev/specs/skills/turn-gate/spec.md`
- 사용자 스펙 의도: `loop-kit-dev/specs/skills/turn-gate/intent.md`
- sub-spec directory: `loop-kit-dev/specs/skills/turn-gate/`
- local runtime references: `loop-kit-dev/skills/turn-gate/references/*.md`
- 호출 방식: 직접 호출하거나 manifest prompt의 안내를 따른다.

## 상세 계약 구조

`turn-gate`의 사용자 스펙 의도, 전체 흐름, 세부 계약은 같은 skill spec folder 아래 ownership-based child folder로 분리합니다.
이 파일은 top-level ownership, whole-flow overview, sibling contract map을 소유합니다.

- `intent.md`: `turn-gate`의 사용자 스펙 의도 기록
- `core/runtime-flow.md`: activation부터 explicit stop까지 `turn-gate`의 전체 phase 흐름과 전환 조건
- `core/skill-contents.md`: runtime `SKILL.md` body의 필수 구성과 content boundary
- `phases/preparation.md`: work 전 intent, scope, non-goal, approval boundary, verification expectation 정렬
- `phases/work.md`: active flow 안에서 실제 작업을 수행하기 위한 work phase 계약
- `phases/verification.md`: work 이후 risk-based verification method 선택과 non-pass 처리로 이어지는 verification phase 계약
- `phases/reporting.md`: terminal close가 아니라 next-flow context를 정리하는 reporting phase 계약
- `phases/next-flow.md`: explicit stop 확인과 다음 flow reopening을 수행하는 next-flow phase 계약
- `core/flow-boundaries.md`: `operational-preparation`, `change-unit`, planned flow boundary, 후속 후보와 active execution flow 구분
- `gates/internal-gates.md`: internal gate model overview and gate detail map
- `gates/flow-shaping.md`: active flow shaping, follow-up candidate separation, completion criteria
- `gates/task-policy.md`: flow-local task sequencing, local references, target rereads, command/edit/build/test policy
- `gates/verification.md`: verification method selection, minimum-sufficient evidence or packet construction, and pass/fail/blocked/insufficient routing
- `gates/reporting.md`: result reporting as continuity context
- `core/meaning-resolution.md`: operation/target ambiguity, provenance/intent block target locking, user-gated clarification
- `modes/default.md`: implicit default operating state 계약
- `phase-protocols/routes.md`: implicit default state, phase protocol selection, local references, operating-state-vs-handoff
- `phase-protocols/deep-interview.md`: requirement discovery와 scope lock protocol 계약
- `phase-protocols/review-loop.md`: review/QA/self-review finding 처리 protocol 계약
- `phase-protocols/ralph-loop.md`: bounded fix-verify-reassess cycle protocol 계약
- `phase-protocols/autopilot.md`: locked-scope end-to-end execution protocol 계약
- `phase-protocols/commit-readiness-gate.md`: commit readiness judgment protocol 계약
- `core/approval-boundary.md`: destructive, irreversible, external-action, commit/publish approval boundary
- `records/verification.md`: risk-based verification method, minimum-sufficient evidence or packet sizing, and non-pass handling
- `records/question-routing.md`: `request_user_input`, next-flow reopening, fallback, visible/recorded turn-end option
- `records/session-records.md`: `000-plan.md`, `001+` flow records, Continuity Guard, templates, `Next Flow Options`
- `templates/plan.md`: `000-plan.md` template structure, date-level snapshot/index ownership, plan/flow deduplication
- `templates/flow.md`: `001+` flow record template structure, flow-local contract/evidence/report ownership, safety fields
- `intent-scenarios/`: runtime instruction이 아니라 flow boundary 의도를 회귀 평가하기 위한 spec-side fixture

## 핵심 처리 계약

- `turn-gate`는 대화 응답 자체를 제어하는 conversation-level first-class rule이다.
- `turn-gate`는 implicit default operating state와 phase protocol routing을 독립적으로 소유한다.
- `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `commit-readiness-gate` 같은 이름은 standalone mode가 아니라 현재 상태에서 필요한 상황별 phase protocol로 취급한다.
- phase protocol routing은 `phase-protocols/routes.md`가 소유하고, 상세 계약은 나머지 `phase-protocols/*.md`가 소유한다.
- `loop-kit-dev/skills/turn-gate/SKILL.md`는 runtime에서 읽는 운영 표면이며, 본문 구성과 runtime/spec boundary는 `core/skill-contents.md`가 소유한다.
- 전체 phase 흐름과 phase 간 전환은 `core/runtime-flow.md`가 소유한다.
- phase 시작 사용자-facing 메시지의 `[<phase-name>(/<phase-protocol>)]` prefix 계약은 `core/runtime-flow.md`와 `core/skill-contents.md`가 소유한다.
- phase 내부 세부 계약은 `phases/*` spec이 소유하고, internal gate 세부 계약은 `gates/*` spec이 소유한다.
- session record와 Continuity Guard 계약은 `records/session-records.md`가 소유한다.
- risk-based verification method 계약은 `records/verification.md`가 소유한다.

## 검토 질문

- 이번 응답이 turn continuity를 실제로 유지하고 있는가?
- skill body 구성과 runtime/spec boundary 판단이 필요하면 `core/skill-contents.md`를 확인했는가?
- skill body 앞부분에 `Important` 섹션이 있고 1급 규칙, terminal summary 금지, next-flow reopening이 먼저 드러나는가?
- `core/runtime-flow.md`만 읽어도 전체 phase 흐름과 다음 상세 spec 위치를 알 수 있는가?
- phase 시작 사용자-facing 메시지 prefix 규칙이 runtime-visible 계약으로 반영돼 있는가?
- 사용자 표현에 구조적 다의성이 있으면 mode 또는 phase protocol 선택 전에 meaning resolution 질문을 열었는가?
- deep-interview/review-loop 같은 phase protocol을 mode처럼 기록하지 않았는가?
- user-gated question routing과 계획 도구 `update_plan`를 필수 단계에서 실제로 사용했는가?
- cross-flow 작업이라면 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 planned flow sequence, 각 flow의 완료 기준, 다음 flow 전환 조건을 최신 상태로 담고 있는가?
- 사용자 메시지 해석과 flow list 설계가 필요했다면, 그 운영 준비가 별도 flow 또는 bootstrap record로 남고 결과 planned flows와 섞이지 않았는가?
- flow shaping, task policy, verification, reporting이 phase 전환 권한을 침범하지 않는가?
- task policy 결과가 flow completion이나 turn closure를 직접 승인하는 구조가 남아 있지 않은가?
- spec-side fixture 평가 규칙이 runtime skill body로 직접 누출되지 않았는가?
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`가 현재 phase까지 증분 갱신됐는가?
- `work -> verification -> result reporting` 순서를 실제로 유지했는가?
- 기본 flow를 `준비 -> 작업 -> 검증 -> 보고 -> next-flow`로 유지했는가?
- 사용자 메시지에서 시작한 preparation과 기존 flow 실행 전 preparation을 구분했는가?
- direct loop entrypoint를 사용자 표면으로 다시 열지 않았는가?
- 결과 보고 뒤 explicit next-flow choice를 실제로 열었는가?
- 결과 보고 직전에 `Continuity Guard`를 갱신했고 terminal summary 가능 여부를 확인했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `workflow-kit`의 일반 workflow skill 의미와 `loop-kit-dev`의 narrow runtime packaging을 전제로 한다. 다만 turn continuity와 operating state/phase protocol selection rule 자체는 이 index spec에서 명시적으로 읽혀야 한다.

## 확장 원칙

- `spec.md`는 top-level ownership과 routing map을 유지하고, 상세 계약이 커지면 같은 폴더 아래 sub-spec으로 내린다.
- 새 sub-spec을 추가할 때는 `spec.md`의 상세 계약 구조와 필요하면 plugin spec의 skill spec 위치 설명을 함께 갱신한다.
