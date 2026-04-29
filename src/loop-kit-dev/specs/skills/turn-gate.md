## 사용자 스펙 의도

- 하나의 턴을 사용자가 턴을 종료하자고 요청할때까지 닫지 않고 유지하고 싶다.
- 이 스킬을 사용한다는건, 이 세션동안 이 스킬을 1급 규칙으로 사용한다는 의미입니다.
- `turn-gate`의 메인 플로우는 스킬 내부 체크리스트가 아니라 대화 응답 자체를 제어하는 1급 규칙이어야 합니다.
- `turn-gate`의 1급 규칙은 일반 목적 설명보다 더 높은 우선순위로 보이도록 skill body의 앞부분에 `Important` 섹션으로 드러나야 합니다.
- `turn-gate`가 활성화된 동안 assistant의 응답은 loop continuation, question-routing, explicit user stop 처리 중 하나로 끝나야 하며, 일반적인 final summary로 턴을 닫으면 안 됩니다.
- `turn-gate`가 현재 phase의 메인 작업을 보고 적절한 내부 loop mode를 고르길 원한다.
- 필요한 경우 requirement discovery 성격의 `deep-interview`도 `turn-gate` 안의 internal mode로 흘러가길 원한다.
- `ralph-loop`, `review-loop`, readiness gate 같은 loop는 사용자가 직접 고르지 않고 `turn-gate` 안에서 흘러가길 원한다.
- 내부 loop mode의 canonical contract는 `workflow-kit`이 SSOT로 계속 소유하길 원한다.
- 내부 loop mode는 `turn-gate` 스킬의 local `references/` 아래로 흡수돼야 하고, 실행 시 그 reference를 읽는 구조이길 원한다.
- broad end-to-end delivery가 현재 phase work라면 `workflow-kit/autopilot`의 계약도 `turn-gate` internal mode로 선택되길 원한다.
- 질문, 선택지, scope lock, 다음 플로우 선택은 기본적으로 user-gated question routing으로 처리하길 원한다.
- `turn-gate`는 질문 도구와 계획 도구를 선택 사항이 아니라 필수 도구로 사용해야 한다.
- 다음 플로우 질문의 사용자 표시 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하더라도, sessions flow record의 `Next Flow Options`에는 명시적인 턴 종료 선택지가 항상 남아야 한다.
- `turn-gate`로 진행한 작업은 `.agents/sessions` 아래에 기록이 남아야 한다.
- 여러 플로우를 거치는 작업의 상위 계획은 `.agents/sessions/{YYYYMMDD}/000-plan.md` 경로에 누적되길 원한다.
- `000-plan.md`는 단순 현재 상태 로그가 아니라 여러 flow의 상위 계획과 흐름을 소유해야 한다.
- 예를 들어 "컴포넌트 오타가 있어. 수정하자" 같은 요청은 `컴포넌트 문구 점검`, `컴포넌트 문구 수정`, `commit-ready` 같은 flow sequence로 나뉘고, 각 flow는 자기 내부의 점검/수정/검증 하위 작업을 소유해야 한다.
- 개별 플로우 기록은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식으로 남고 싶다.
- `001+` 문서는 phase 메모가 아니라 flow 기록으로 남고 싶다.
- 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 필요하면 미리 설계하길 원한다.
- 이후 loop에서 다시 분석 단계나 계획 단계로 돌아오면, 이전 flow/phase 설계를 고정값처럼 취급하지 말고 필요할 때만 다시 설계하길 원한다.
- 검증 단계는 그 재설계를 직접 수행하는 단계라기보다, 이후 flow/phase 재설계가 필요한지 여부를 드러내는 단계이길 원한다.

---

# turn-gate 스킬 스펙

## 목적

`loop-kit-dev`의 `turn-gate`는 turn continuity를 유지하면서 current-phase work에 맞는 내부 loop mode를 선택해 실행하는 메인 loop controller입니다.
이 skill이 활성화되면 `turn-gate` 메인 플로우는 대화 응답 자체의 1급 제어 규칙이 되며, 사용자의 explicit stop 전까지 결과 보고를 terminal response로 닫지 않습니다.

## 경계

- 포함:
  - turn-level continuity 유지
  - `analysis -> plan -> work -> verification -> result reporting -> next-flow question-routing response` 구조 유지
  - 사용자 메시지의 operation 의미 해독
  - current-phase work의 internal mode selection
  - 결과 보고 뒤 explicit choice 기반 next-flow reopening
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 여러 direct loop skill을 사용자에게 노출하는 일
  - domain-specific implementation detail 자체

## 처리하려는 작업 형태

- 사용자가 턴을 종료하자고 요청하기 전까지 한 턴 안에서 여러 phase를 이어가야 하는 작업
- requirement discovery, autonomous execution, refinement, review-driven correction, readiness checking 같은 current-phase work가 번갈아 나타나는 작업
- 결과 보고 뒤 clean stop이 아니라 다음 플로우 선택이 기본이어야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit-dev/skills/turn-gate/SKILL.md`
- local reference surface: `loop-kit-dev/skills/turn-gate/references/*.md`
- 관련 상위 라우팅: `loop-kit-dev-guide`

## 핵심 처리 계약

### Skill body 작성 계약

- `loop-kit-dev/skills/turn-gate/SKILL.md`는 이 스펙의 단순 요약본이 아니라 runtime에서 읽는 운영 표면이다.
- skill body는 대화 응답 자체를 제어하는 conversation-level first-class rule을 `## Important` 섹션으로 앞부분에서 명시해야 한다.
- `## Important` 섹션은 `Purpose`보다 먼저 위치해야 하며, 최소한 session-level activation, terminal summary 금지, required ending states, `request_user_input` 기반 next-flow reopening, session record 유지 의무를 포함해야 한다.
- `## Important` 섹션은 긴 설명이 아니라 실행 중 우선 확인할 수 있는 짧은 규칙 목록이어야 한다.
- skill body에는 `Core Loop` 또는 이에 준하는 단계별 실행 섹션이 있어야 하며, 최소한 analysis, plan, work, verification, result reporting, question-routing reopening을 각각 구분해 설명해야 한다.
- skill body에는 internal mode selection과 local `references/` 읽기 규칙이 직접 남아 있어야 한다.
- skill body에는 terminal summary 금지, next-flow reopening, Continuity Guard 확인, user-gated question routing, explicit turn-end option 기록 규칙이 직접 남아 있어야 한다.
- skill body를 짧게 다듬더라도 위 단계와 금지 규칙을 한 문단으로 뭉개지 말고, 실행 중 빠르게 확인 가능한 형태로 유지한다.

### 활성화와 응답 흐름

- 각 incoming message를 같은 loop-gated turn의 현재 입력으로 취급한다.
- 중간 사용자 메시지를 stop, completion, approval-boundary pause로 해석하지 않고 같은 loop-gated turn의 authoritative input으로 취급한다.
- 중간 사용자 메시지는 explicit turn stop, status/progress check, current-flow correction, current-flow priority change, next-flow priority request 중 하나로 분류한다.
- status/progress check라면 현재 phase, blocker 또는 progress, 다음 concrete action을 짧게 보고한 뒤 active flow를 계속한다.
- current-flow correction 또는 current-flow priority change라면 현재 analysis/plan을 즉시 조정하고 가장 이른 안전한 phase부터 이어간다.
- next-flow priority request라면 flow record의 next-flow 후보 중 최우선으로 등록하고 다음 safe handoff point까지 이어간다.
- 이 skill이 사용되면 현재 세션 동안 `turn-gate`를 conversation-level first-class operating rule로 활성화한 것으로 취급한다.
- 이 규칙은 skill 내부 체크리스트가 아니라 assistant response lifecycle 자체에 적용한다.
- 이 규칙은 skill body의 일반 설명보다 앞선 `Important` 섹션에서 다시 확인 가능해야 한다.
- concrete task 없이 activation만 요청된 경우에도 `turn-gate`를 활성화하고 session record를 생성 또는 갱신한다. 이때 `user_explicit_stop=false`로 보고, work mode를 성급히 고르지 않고 user-gated next-flow 또는 scope selection을 연다.
- `user_explicit_stop`이 false인 동안 result reporting은 terminal response가 아니며, 반드시 next-flow reopening 또는 active question-routing으로 이어져야 한다.
- 사용자가 명시적으로 턴 종료를 요청했거나 flow record에 confirmed closure가 기록된 경우가 아니라면 일반적인 final summary로 턴을 닫지 않는다.
- 사용자가 explicit stop을 요청하면 active flow record와 `Continuity Guard`에 confirmed closure를 기록하고, `terminal summary allowed`를 허용 상태로 갱신하며, next-flow choices를 다시 열지 않는다.
- `analysis`, `plan`, `work`, `verification`, `result reporting`, `question-routing reopening`을 응답 shape에 계속 드러낸다.
- 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 미리 설계할 수 있다.
- 그 future flow/phase 설계는 provisional하며, 이후 loop에서 새 증거, changed intent, 새 blocker가 생겼을 때만 다시 설계한다.
- 계획 이후 current-flow correction이나 target file/state 변경이 들어오면 affected files 또는 state를 다시 읽고 reconcile한 뒤 가장 이른 안전한 phase부터 재개한다.
- status/progress check를 처리한 뒤에도 진행 상태가 바뀌었으면 active flow record의 phase, blocker, required next action을 갱신한다.

### meaning resolution

- analysis 단계에서는 internal mode 선택이나 작업 실행 전에 사용자 메시지의 operation 의미를 먼저 해독한다.
- `merge`, `absorb`, `remove`, `delete`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, `contract` 또는 이에 대응되는 한국어 표현처럼 여러 구조 단위를 가리킬 수 있는 표현은 바로 하나의 작업으로 단정하지 않는다.
- `그`, `그 밑`, `그건`, `그거`, `위`, `아래`, `현재 것`처럼 주변 문맥의 여러 대상을 가리킬 수 있는 지시 표현도 해석에 따라 작업이 달라지면 meaning resolution 대상으로 본다.
- source URL, provenance note, `사용자 스펙 의도` 또는 spec intent block은 대화 맥락처럼 버릴 수 있는 텍스트가 아니라 작업 target이 될 수 있다. `출처`, `원본`, `의도`, `그 밑`이 provenance, intent block, normative spec body 중 무엇을 가리키는지에 따라 작업이 달라지면 먼저 target을 잠근다.
- 해석 후보에 따라 파일 범위, 삭제 여부, phase 설계, routing rule, migration 의미, commit scope가 달라지면 active question-routing으로 의미를 먼저 잠근다.
- meaning resolution 질문도 user-gated이며, 구조적 선택지를 줄 수 있으면 `request_user_input`으로 잠근다.
- meaning resolution 질문은 `deep-interview`가 소유하는 requirement discovery가 아니라, 현재 지시어의 operation 또는 target을 잠그는 current-flow clarification이다.
- 질문은 넓은 freeform 질문이 아니라 "여기서 병합은 skill/spec surface를 합치는 뜻인가, `turn-gate` phase로 흡수하는 뜻인가"처럼 다의어가 가리키는 구조 단위를 직접 잠그는 형태여야 한다.
- meaning resolution이 필요한 경우 flow record의 analysis에는 literal wording, interpreted operation, operation target, alternate interpretations, impact of ambiguity를 남긴다.

### 세션 기록과 Continuity Guard

- active turn-gated task마다 `.agents/sessions/{YYYYMMDD}/000-plan.md` 날짜 기준 plan과 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 상세 flow report 체계를 유지한다.
- `000-plan.md`는 당일 작업의 히스토리, 사용자 요청 목록, flow index, 현재 계획, 완료 flow 요약을 소유한다.
- `000-plan.md`의 현재 계획은 action checklist가 아니라 planned flow sequence여야 한다. 각 planned flow에는 flow 목적, 왜 이 flow가 필요한지, 완료 기준, 다음 flow로 넘어가는 조건이 드러나야 한다.
- 세부 작업 단계는 해당 `001+` flow record의 plan/work/verification에 둔다. `000-plan.md`는 "이 작업이 어떤 flow들의 흐름으로 진행되는지"를 소유하고, 각 flow record는 "그 flow 안에서 무엇을 했는지"를 소유한다.
- `000-plan.md`는 증분 갱신하고, 완료된 작업도 삭제하지 않고 요약과 flow reference를 유지한다.
- 해당 `001+` record는 사용자 요청에 따른 개별 flow의 상세 보고서이며, completed flow를 기다리지 말고 각 phase가 끝날 때마다 증분 갱신한다.
- flow record의 `Next Flow Options`에는 사용자 표시 질문에 턴 종료 선택지가 보이지 않는 경우에도 명시적인 turn-end option이 포함되어야 한다.
- 각 flow record에는 짧은 `Continuity Guard`를 두고, result reporting과 next-flow reopening 직전에 반드시 갱신한다.
- `Continuity Guard`에는 `turn-gate` 활성 여부, question-routing mode, user explicit stop 여부, terminal summary 허용 여부, required next action이 포함되어야 한다.
- result reporting과 next-flow reopening 전에는 active flow record의 `Continuity Guard`를 먼저 읽는다. 기록이 없거나 접근할 수 없을 때만 재구성하고, 재구성한 guard는 가능한 즉시 flow record에 다시 쓴다.

### 도구와 internal mode 선택

- 실질적인 작업이 시작되면 계획 도구 `update_plan`을 필수로 사용하고 현재 active step 상태를 유지한다.
- `work`에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
- `loop-kit-dev`에서는 사용자가 internal mode를 직접 호출하는 대신 `turn-gate`가 이를 선택한다.
- `turn-gate`는 선택된 internal mode에 대응하는 local `references/` 문서를 먼저 읽고 그 계약을 적용한다.
- local `references/`는 `workflow-kit` upstream spec과 동기화된 absorbed operational contract로 유지한다.

### question routing

- 질문, 선택지 제시, scope lock, next-flow reopening에는 user-gated question routing을 필수로 사용한다.
- scope lock, 선택지, next-flow reopening은 `request_user_input`으로 사용자에게 묻는다.
- explicit user, tool, platform, safety, destructive, irreversible, external-action approval boundary는 반드시 사용자 질문으로 유지한다.
- destructive, irreversible, external action 전에는 현재 상태를 확인해 대상과 위험을 말할 수 있어야 한다.
- publish, push, pull request 생성처럼 외부 시스템에 영향을 주는 action은 branch, remote, scope, risk를 먼저 확인하고 사용자 승인을 받은 뒤 GitHub 또는 해당 external-action workflow로 넘긴다.
- 사용자가 subagent로 막힌 질문을 처리하라고 하면 autonomous question routing은 `turn-gate-self-drive`로 넘기고, approval, destructive, irreversible, external-action, safety 결정은 user-gated로 유지한다.

### 검증과 next-flow reopening

- `work` 뒤에는 결과 보고 전에 명시적 검증 단계를 두고, 그 검증은 이후 flow/phase 재설계 필요 여부를 드러내는 단계로 취급한다.
- 결과 보고 전에는 `Continuity Guard`를 읽거나 재구성하고, 사용자가 명시적으로 종료하지 않았으면 terminal summary가 invalid임을 확인한다.
- 결과 보고 뒤에는 explicit choice를 주는 active question-routing mode로 다음 플로우를 다시 연다.
- 결과 보고 뒤 visible next-flow choice는 도구 기반 질문이어야 하며, 가능한 경우 `request_user_input`으로 직접 연다. loop continuation도 사용자에게 현재 phase, required next action, 열린 선택지를 볼 수 있게 해야 한다.
- assistant final-answer channel이 사용되더라도 그것만으로 loop termination을 의미하지 않는다. explicit stop이 없다면 next-flow question이 여전히 필수다.
- 사용자에게 보이는 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하는 경우에도, flow record의 `Next Flow Options`에는 별도 turn-end option을 기록한다.
- 사용자가 턴을 종료하자고 요청하지 않으면 clean stop을 기본 경로로 두지 않는다.

## 내부 loop mode 선택 규칙

- 실제 requirement discovery가 병목이면 `references/deep-interview.md`를 따른다.
- broad end-to-end delivery가 필요한 current phase이면 `references/autopilot.md`를 따른다.
- bounded issue를 작은 fix-verify-reassess cycle로 다루는 경우 `references/ralph-loop.md`를 따른다.
- review feedback이나 material finding 처리인 경우 `references/review-loop.md`를 따른다.
- 현재 변경 단위가 거의 끝났고 readiness 판단이 핵심이면 `references/commit-readiness-gate.md`를 따른다.
- 사용자가 명시적으로 commit 실행을 요청한 경우에도 먼저 intended change unit의 scope, staged/final status, readiness를 확인한 뒤 commit execution workflow로 넘긴다. readiness 점검 요청은 commit 승인으로 해석하지 않는다.
- 아직 어떤 internal mode가 맞는지 확정되지 않았다면 active question-routing mode나 좁은 분석으로 먼저 mode selection을 잠근다.

## deep-interview 원본 관계

- 원본 skill source는 `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`다.
- 원본은 ambiguity gating, OMX tooling, artifact handoff까지 포함한 더 큰 workflow다.
- `loop-kit-dev`의 `references/deep-interview.md`는 그 full workflow를 그대로 복제한 것이 아니라, `turn-gate`가 requirement-discovery phase에서 필요한 boundary만 흡수한 derived reference다.

## mode selection matrix

- blocker가 requirement discovery, intent ambiguity, scope boundary, approval line이라면 `deep-interview`를 고른다.
- input이 review feedback, QA finding, self-review finding이고 한 번에 하나의 material issue만 처리해야 한다면 `review-loop`를 고른다.
- blocker가 하나의 bounded improvement cycle이고 작은 fix 뒤 즉시 검증하는 흐름이 맞다면 `ralph-loop`를 고른다.
- brief request부터 implementation, QA, validation까지 이어지는 broad end-to-end delivery가 현재 phase work라면 `autopilot`을 고른다.
- 구현이 거의 끝났고 intended change unit의 commit readiness 판단이 핵심이면 `commit-readiness-gate`를 고른다.
- 여러 mode가 겹쳐 보이면 `deep-interview -> review-loop -> ralph-loop -> autopilot -> commit-readiness-gate` 순으로 더 이른 병목을 우선한다.
- 그래도 mode를 못 잠그면 active question-routing mode로 선택지를 좁힌 뒤 work phase로 들어간다.

## question routing

- `user-gated`: 선택지, scope lock, next-flow decision을 사용자에게 질문 도구로 묻는다.
- 이 스펙은 user-gated base loop gate만 소유한다.

## next-flow reopening 규칙

- 다음 플로우 질문은 현재 결과에 직접 연결된 좁은 선택지여야 한다.
- generic follow-up phrase나 자유형 마무리 질문으로 턴을 닫지 않는다.
- `user-gated` 사용자 응답은 같은 턴의 다음 메시지로 즉시 이어진다.

## session record 규칙

- 여러 플로우를 거치는 작업이면 `.agents/sessions/{YYYYMMDD}/000-plan.md`를 먼저 둔다.
- `000-plan.md`는 날짜 기준 plan artifact로 유지하고, 당일 작업 히스토리, 사용자 요청 목록, flow index, planned flow sequence, 완료 flow 요약을 계속 증분한다.
- `000-plan.md`의 planned flow sequence는 여러 flow를 위에서 아래로 설계한다. 예: `컴포넌트 문구 점검` flow가 위치 확인, 문구 확인, 맥락 파악, 리스트업을 소유하고, `컴포넌트 문구 수정` flow가 리스트 기반 수정과 빌드/린트를 소유하며, `commit-ready` flow가 사용자 의도 대비 commit diff 확인과 review를 소유한다.
- `000-plan.md`에는 각 flow의 목적, 상태, 완료 기준, 다음 flow 전환 조건을 요약하고, flow 내부의 하위 단계는 해당 `001+` record에 둔다.
- 완료된 작업은 삭제하지 않고 요약과 flow reference를 유지한다.
- flow 기록 파일은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식을 사용한다.
- `count-pad3`는 `001`, `002`, `003`처럼 3자리 zero-padded 숫자를 사용한다.
- slug는 영어 소문자와 `-`만 사용한다.
- flow 기본 템플릿은 `skills/turn-gate/templates/flow-record-template.md`를 사용한다.
- `000-plan.md` 기본 템플릿은 `skills/turn-gate/templates/plan-template.md`를 사용한다.
- 최소 flow 기록 항목은 user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, residual risk다.
- `000-plan.md`는 날짜 기준 증분 계획과 flow-sequence artifact로, `001+`는 flow 단위 상세 보고서로 취급한다.
- flow record는 phase 메모가 아니지만, `analysis`, `plan`, `work`, `verification`, `result reporting` 각 phase가 끝날 때마다 현재 상태로 갱신해야 한다.

## SSOT 동기화 규칙

- internal mode contract 변경은 먼저 `workflow-kit` upstream spec에서 정리한다.
- `loop-kit-dev`은 runtime orchestration 관점의 차이와 local absorbed references를 별도로 소유한다.
- upstream contract와 `turn-gate` references의 문구가 어긋나면 같은 변경 단위에서 함께 갱신한다.

## 검토 질문

- 이번 응답이 turn continuity를 실제로 유지하고 있는가?
- skill body 앞부분에 `Important` 섹션이 있고 1급 규칙, terminal summary 금지, next-flow reopening이 먼저 드러나는가?
- 사용자 표현에 구조적 다의성이 있으면 internal mode 선택 전에 meaning resolution 질문을 열었는가?
- current-phase work에 맞는 internal mode를 하나로 좁혔는가?
- user-gated question routing과 계획 도구 `update_plan`를 필수 단계에서 실제로 사용했는가?
- cross-flow 작업이라면 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 planned flow sequence, 각 flow의 완료 기준, 다음 flow 전환 조건을 최신 상태로 담고 있는가?
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`가 현재 phase까지 증분 갱신됐는가?
- `work -> verification -> result reporting` 순서를 실제로 유지했는가?
- direct loop entrypoint를 사용자 표면으로 다시 열지 않았는가?
- 결과 보고 뒤 explicit next-flow choice를 실제로 열었는가?
- 사용자 표시 선택지에 턴 종료 option이 없더라도 flow record의 `Next Flow Options`에 명시적인 turn-end option을 남겼는가?
- 결과 보고 직전에 `Continuity Guard`를 갱신했고 terminal summary 가능 여부를 확인했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `workflow-kit`의 canonical loop-mode contract와 `loop-kit-dev`의 narrow runtime packaging을 전제로 한다. 다만 turn continuity와 mode selection rule 자체는 이 스펙에서 명시적으로 읽혀야 한다.

## 확장 원칙

- 새로운 internal mode는 기존 mode로 현재 phase work를 소유할 수 없을 때만 추가한다.
- internal mode set이나 mandatory tool rule이 바뀌면 `workflow-kit` upstream spec, `loop-kit-dev` plugin spec, `loop-kit-dev-guide`, `turn-gate`, `turn-gate/references/`를 함께 갱신한다.
