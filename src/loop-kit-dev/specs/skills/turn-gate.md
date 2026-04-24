## 사용자 스펙 의도

- 하나의 턴을 사용자가 턴을 종료하자고 요청할때까지 닫지 않고 유지하고 싶다.
- 이 스킬을 사용한다는건, 이 세션동안 이 스킬을 1급 규칙으로 사용한다는 의미입니다.
- `turn-gate`가 현재 phase의 메인 작업을 보고 적절한 내부 loop mode를 고르길 원한다.
- 필요한 경우 requirement discovery 성격의 `deep-interview`도 `turn-gate` 안의 internal mode로 흘러가길 원한다.
- `ralph-loop`, `review-loop`, readiness gate 같은 loop는 사용자가 직접 고르지 않고 `turn-gate` 안에서 흘러가길 원한다.
- 내부 loop mode의 canonical contract는 `workflow-kit`이 SSOT로 계속 소유하길 원한다.
- 내부 loop mode는 `turn-gate` 스킬의 local `references/` 아래로 흡수돼야 하고, 실행 시 그 reference를 읽는 구조이길 원한다.
- broad end-to-end delivery가 현재 phase work라면 `workflow-kit/autopilot`의 계약도 `turn-gate` internal mode로 선택되길 원한다.
- 지금까지의 current-phase mode 축과 별도로, 질문 대상을 고르는 question-routing 축을 두고 싶다.
- `self-drive` mode에서는 사용자에게 묻던 질문을 subagent에게 물어 사용자 개입 없이 계속 자동 진행하고 싶다.
- `self-drive` 진행 중 사용자가 중간 개입하면 멈추지 말고, 사용자 메시지를 현재 loop 입력으로 받아 현재 플로우를 조정하거나 다음 플로우 우선순위로 등록하길 원한다.
- `turn-gate`는 질문 도구와 계획 도구를 선택 사항이 아니라 필수 도구로 사용해야 한다.
- `turn-gate`로 진행한 작업은 `.agents/sessions` 아래에 기록이 남아야 한다.
- 여러 플로우를 거치는 작업의 상위 계획은 `.agents/sessions/{YYYYMMDD}/000-plan.md` 경로에 누적되길 원한다.
- 개별 플로우 기록은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식으로 남고 싶다.
- `001+` 문서는 phase 메모가 아니라 flow 기록으로 남고 싶다.
- 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 필요하면 미리 설계하길 원한다.
- 이후 loop에서 다시 분석 단계나 계획 단계로 돌아오면, 이전 flow/phase 설계를 고정값처럼 취급하지 말고 필요할 때만 다시 설계하길 원한다.
- 검증 단계는 그 재설계를 직접 수행하는 단계라기보다, 이후 flow/phase 재설계가 필요한지 여부를 드러내는 단계이길 원한다.

---

# turn-gate 스킬 스펙

## 목적

`loop-kit-dev`의 `turn-gate`는 turn continuity를 유지하면서 current-phase work에 맞는 내부 loop mode를 선택해 실행하는 메인 loop controller입니다.

## 경계

- 포함:
  - turn-level continuity 유지
  - `analysis -> plan -> work -> verification -> result reporting -> next-flow question-routing response` 구조 유지
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

### 활성화와 응답 흐름

- 각 incoming message를 같은 loop-gated turn의 현재 입력으로 취급한다.
- `self-drive`가 활성화된 중간에 사용자 메시지가 들어와도 이를 stop, completion, approval-boundary pause로 해석하지 않고 같은 loop-gated turn의 authoritative input으로 취급한다.
- 중간 사용자 메시지는 explicit turn stop, current-flow correction, current-flow priority change, next-flow priority request 중 하나로 분류한다.
- current-flow correction 또는 current-flow priority change라면 현재 analysis/plan을 즉시 조정하고 가장 이른 안전한 phase부터 이어간다.
- next-flow priority request라면 flow record의 next-flow 후보 중 최우선으로 등록하고 다음 safe handoff point까지 이어간다.
- 새 사용자 입력과 충돌하는 self-drive subagent 답변은 stale answer로 보고 다음 phase 입력으로 쓰지 않는다.
- 이 skill이 사용되면 현재 세션 동안 `turn-gate`를 first-class operating rule로 활성화한 것으로 취급한다.
- `analysis`, `plan`, `work`, `verification`, `result reporting`, `question-routing reopening`을 응답 shape에 계속 드러낸다.
- 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 미리 설계할 수 있다.
- 그 future flow/phase 설계는 provisional하며, 이후 loop에서 새 증거, changed intent, 새 blocker가 생겼을 때만 다시 설계한다.

### 세션 기록과 Continuity Guard

- active turn-gated task마다 `.agents/sessions/{YYYYMMDD}/000-plan.md` 상위 계획과 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` record 체계를 유지한다.
- `000-plan.md`는 사용자 요청 종료 이후에도 더 큰 작업이 이어지면 계속 증분 갱신한다.
- 해당 `001+` record는 completed flow를 기다리지 말고 각 phase가 끝날 때마다 증분 갱신한다.
- 각 flow record에는 짧은 `Continuity Guard`를 두고, result reporting과 next-flow reopening 직전에 반드시 갱신한다.
- `Continuity Guard`에는 `turn-gate` 활성 여부, question-routing mode, user explicit stop 여부, terminal summary 허용 여부, required next action이 포함되어야 한다.

### 도구와 internal mode 선택

- 실질적인 작업이 시작되면 계획 도구 `update_plan`을 필수로 사용하고 현재 active step 상태를 유지한다.
- `work`에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
- `loop-kit-dev`에서는 사용자가 internal mode를 직접 호출하는 대신 `turn-gate`가 이를 선택한다.
- `turn-gate`는 선택된 internal mode에 대응하는 local `references/` 문서를 먼저 읽고 그 계약을 적용한다.
- local `references/`는 `workflow-kit` upstream spec과 동기화된 absorbed operational contract로 유지한다.

### question-routing mode

- 질문, 선택지 제시, scope lock, next-flow reopening에는 active question-routing mode를 필수로 사용한다.
- `turn-gate`는 current-phase internal mode와 별개로 question-routing mode를 선택한다.
- 기본 question-routing mode는 `user-gated`이며, scope lock, 선택지, next-flow reopening을 `request_user_input`으로 사용자에게 묻는다.
- `self-drive` question-routing mode가 활성화되면 `references/self-drive.md`를 읽고, 사용자에게 묻던 phase 질문을 self-drive question packet으로 구성해 subagent에게 물어 그 답을 다음 결정 입력으로 사용한다.
- self-drive question packet에는 최소한 phase, current mode, question type, decision needed, options, context, continuity guard, constraints, fallback, expected answer를 포함한다.
- self-drive 도중 사용자 메시지가 들어온 경우 question packet에는 user interventions, classification, superseded assumptions/answers를 포함한다.
- self-drive subagent answer에는 최소한 question id, selected option, decision, rationale, evidence, assumptions, confidence, blockers, approval boundary, continuity check, next action을 포함한다.
- self-drive subagent answer는 더 최신 사용자 입력과 충돌하면 superseded로 취급되어야 한다.

### self-drive pause와 recovery

- self-drive answer는 user explicit stop 또는 hard approval boundary가 없는 한 terminal summary를 허용하지 않고 계속 이어질 next action을 제시해야 한다.
- subagent answer가 `context_gap`을 반환하면 메인 에이전트가 repo search, file read, log, deterministic check, policy-allowed web research로 회복 가능한지 먼저 판단하고, 회복 가능하면 evidence를 보강한 packet으로 다시 질문한다.
- 사용자 취향을 확정할 수 없으면 명시적 manual preference lock 요청이 없는 한 가장 안전하고 되돌릴 수 있는 기본값을 가정으로 기록하고 계속 진행한다.
- `low` confidence는 명시적 승인, 파괴적/비가역적/외부 action 승인, platform/tool/safety policy 경계처럼 self-drive가 회복하면 안 되는 경우에만 approval-boundary pause로 취급한다.
- `self-drive`에서 approval-boundary pause는 자율 라우팅의 일시 중지만 의미하며, 턴을 종료하지 않고 `user-gated`로 전환해 `request_user_input`을 열어야 한다.
- `self-drive`는 mode selection, criteria, scope assumption, verification choice, next-flow decision을 subagent 질문으로 처리할 수 있다.
- 사용자 중간 개입은 `self-drive`를 자동으로 해제하지 않는다. 사용자가 manual control을 요구하거나 실제 승인 경계가 생긴 경우에만 `user-gated`로 전환한다.
- runtime, tool, safety policy가 명시적 사용자 승인을 요구하는 경계는 `self-drive`가 대신 동의한 것으로 처리하지 않는다.

### 검증과 next-flow reopening

- `work` 뒤에는 결과 보고 전에 명시적 검증 단계를 두고, 그 검증은 이후 flow/phase 재설계 필요 여부를 드러내는 단계로 취급한다.
- 결과 보고 전에는 `Continuity Guard`를 읽거나 재구성하고, 사용자가 명시적으로 종료하지 않았으면 terminal summary가 invalid임을 확인한다.
- 결과 보고 뒤에는 explicit choice를 주는 active question-routing mode로 다음 플로우를 다시 연다.
- 사용자가 턴을 종료하자고 요청하지 않으면 clean stop을 기본 경로로 두지 않는다.

## 내부 loop mode 선택 규칙

- 실제 requirement discovery가 병목이면 `references/deep-interview.md`를 따른다.
- broad end-to-end delivery가 필요한 current phase이면 `references/autopilot.md`를 따른다.
- bounded issue를 작은 fix-verify-reassess cycle로 다루는 경우 `references/ralph-loop.md`를 따른다.
- review feedback이나 material finding 처리인 경우 `references/review-loop.md`를 따른다.
- 현재 변경 단위가 거의 끝났고 readiness 판단이 핵심이면 `references/commit-readiness-gate.md`를 따른다.
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

## question-routing axis

- `user-gated`: 기본 mode. 선택지, scope lock, next-flow decision을 사용자에게 질문 도구로 묻는다.
- `self-drive`: 사용자 질문을 subagent 질문으로 바꾸고, subagent 답변과 가정을 flow record에 남긴 뒤 자동으로 다음 phase를 진행한다.
- `self-drive`는 current-phase mode가 아니라 질문 대상 축의 mode이므로 `deep-interview`, `autopilot`, `ralph-loop`, `review-loop`, `commit-readiness-gate`와 함께 활성화될 수 있다.

## next-flow reopening 규칙

- 다음 플로우 질문은 현재 결과에 직접 연결된 좁은 선택지여야 한다.
- generic follow-up phrase나 자유형 마무리 질문으로 턴을 닫지 않는다.
- `user-gated` 사용자 응답 또는 `self-drive` subagent 답변은 같은 턴의 다음 메시지로 즉시 이어진다.
- `self-drive` 중간 사용자 메시지도 같은 턴의 다음 메시지로 즉시 이어지며, 현재 플로우 조정 또는 다음 플로우 우선 등록 중 하나로 처리된다.

## session record 규칙

- 여러 플로우를 거치는 작업이면 `.agents/sessions/{YYYYMMDD}/000-plan.md`를 먼저 둔다.
- `000-plan.md`는 상위 multi-flow plan artifact로 유지하고, 요청 종료 뒤에도 같은 큰 작업이 이어지면 계속 증분한다.
- flow 기록 파일은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식을 사용한다.
- `count-pad3`는 `001`, `002`, `003`처럼 3자리 zero-padded 숫자를 사용한다.
- slug는 영어 소문자와 `-`만 사용한다.
- flow 기본 템플릿은 `skills/turn-gate/templates/flow-record-template.md`를 사용한다.
- `000-plan.md` 기본 템플릿은 `skills/turn-gate/templates/plan-template.md`를 사용한다.
- 최소 flow 기록 항목은 user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, residual risk다.
- `000-plan.md`는 장기 증분 계획 artifact로, `001+`는 flow 단위 운영 artifact로 취급한다.
- flow record는 phase 메모가 아니지만, `analysis`, `plan`, `work`, `verification`, `result reporting` 각 phase가 끝날 때마다 현재 상태로 갱신해야 한다.

## SSOT 동기화 규칙

- internal mode contract 변경은 먼저 `workflow-kit` upstream spec에서 정리한다.
- `loop-kit-dev`은 runtime orchestration 관점의 차이와 local absorbed references를 별도로 소유한다.
- upstream contract와 `turn-gate` references의 문구가 어긋나면 같은 변경 단위에서 함께 갱신한다.

## 검토 질문

- 이번 응답이 turn continuity를 실제로 유지하고 있는가?
- current-phase work에 맞는 internal mode를 하나로 좁혔는가?
- active question-routing mode와 계획 도구 `update_plan`를 필수 단계에서 실제로 사용했는가?
- cross-flow 작업이라면 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 최신 상태인가?
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`가 현재 phase까지 증분 갱신됐는가?
- `work -> verification -> result reporting` 순서를 실제로 유지했는가?
- direct loop entrypoint를 사용자 표면으로 다시 열지 않았는가?
- 결과 보고 뒤 explicit next-flow choice를 실제로 열었는가?
- 결과 보고 직전에 `Continuity Guard`를 갱신했고 terminal summary 가능 여부를 확인했는가?
- `self-drive`가 hard boundary에 도달했을 때 자율 라우팅을 일시 중지하고, 턴을 종료하지 않은 채 `user-gated` 질문 도구로 전환했는가?
- `self-drive` 중간 사용자 메시지를 stop으로 오해하지 않고 현재 플로우 조정 또는 다음 플로우 우선 등록으로 처리했는가?
- 새 사용자 입력과 충돌하는 self-drive subagent 답변을 stale answer로 폐기했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `workflow-kit`의 canonical loop-mode contract와 `loop-kit-dev`의 narrow runtime packaging을 전제로 한다. 다만 turn continuity와 mode selection rule 자체는 이 스펙에서 명시적으로 읽혀야 한다.

## 확장 원칙

- 새로운 internal mode는 기존 mode로 현재 phase work를 소유할 수 없을 때만 추가한다.
- internal mode set이나 mandatory tool rule이 바뀌면 `workflow-kit` upstream spec, `loop-kit-dev` plugin spec, `loop-kit-dev-guide`, `turn-gate`, `turn-gate/references/`를 함께 갱신한다.
