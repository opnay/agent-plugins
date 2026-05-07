# turn-gate session-records sub-spec

## 목적

이 문서는 `turn-gate`가 유지하는 `.agents/sessions/{YYYYMMDD}/` 기록, `000-plan.md`, 개별 flow record, `Continuity Guard`, `Next Flow Options` 계약을 소유합니다.

## 기록 구조

- active turn-gated task마다 `.agents/sessions/{YYYYMMDD}/000-plan.md` 날짜 기준 plan과 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 상세 flow report 체계를 유지한다.
- `000-plan.md`는 당일 작업의 히스토리, 사용자 요청 목록, flow index, 현재 계획, 완료 flow 요약을 소유한다.
- `000-plan.md`의 현재 계획은 action checklist가 아니라 planned flow sequence여야 한다.
- session record에는 flow type을 구분할 수 있어야 한다: `operational-preparation` 또는 `change-unit`.
- `operational-preparation` flow는 사용자 메시지 해석, scope lock, approval boundary 정리, planned flow list 설계를 소유하며, 산출물은 plan/session record다.
- `change-unit` flow는 실제 코드, 문서, fixture, 설정, release surface처럼 검토 가능한 산출물 변경을 소유한다.
- 사용자 메시지 기반 bootstrap을 기록할 때는 operational-preparation flow와 그 결과 planned change-unit flow list를 구분한다.
- 실행이 아니라 판단, 설계, 범위 확인만 요청된 경우 실제 실행 flow와 구분한다. 이 경우 `operational-preparation` record는 후속 실행 후보, scope/non-goal, target ambiguity 판단, verification expectation을 기록하고 종료할 수 있다.
- 운영 flow에 남긴 후속 실행 후보는 아직 시작된 `change-unit` flow가 아니다. 실제 실행으로 이어질 때만 별도 `001+` flow record 또는 다음 count의 `change-unit` record를 만든다.
- planned flow sequence는 phase checklist가 아니며, `분석`, `작업`, `검증`, `커밋 준비` 같은 진행 단계를 별도 flow로 나열하지 않는다.
- 각 planned flow는 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 응집된 변경 단위여야 한다. 최종 사용자에게 직접 보이는 가치 단위가 아니어도 된다.
- 순수 최종 QA, 통합 검증, 정합성 점검, 검증 결과 보고, commit-readiness reporting은 별도 산출물 변경을 소유하지 않는 한 planned flow로 기록하지 않는다.
- 그런 확인과 보고는 마지막 변경 단위 flow record의 `Verification`, `Result Report`, `Next Flow Options`, 또는 user-gated handoff 상태에 기록한다.
- 회귀 테스트 fixture, snapshot baseline, 문서, 운영자 리포트 출력, validator 진단 출력처럼 검토 가능한 산출물을 만들거나 바꾸는 경우에는 그 산출물 변경을 planned flow로 기록할 수 있다.
- 각 planned flow에는 flow type, flow 목적, 왜 이 경계가 필요한지, 완료 기준, 다음 flow로 넘어가는 조건이 드러나야 한다.
- concrete task에서 만든 planned flow sequence에는 preparation source, preparation result, planned flow list가 드러나야 한다.
- 실행이 아니라 판단, 설계, 범위 확인이 목적인 경우에는 `planned flow list` 대신 `follow-up change-unit candidates`로 기록할 수 있다. 이렇게 기록한 후보는 실행 승인이 있거나 다음 flow로 선택되기 전까지 완료/진행 flow로 세지 않는다.
- 각 flow는 기본적으로 `preparation -> work -> verification -> reporting` 단계를 가진다.
- 사용자 메시지 기반 preparation이면 deep-interview result와 사용자 의도에 맞춘 flow list를 기록한다.
- 비 사용자 메시지 기반 preparation이면 수정 범위, 현재 상태, 대상 파일, stale assumption, 실행 전 조건 확인 결과를 기록한다.
- 압축했다면 어느 flow가 preparation/work/verification/reporting을 함께 소유하는지 설명한다. 단, phase를 쪼갠 것을 flow sequence로 포장하지 않는다.
- 세부 작업 단계는 해당 `001+` flow record의 plan/work/verification에 둔다.
- `000-plan.md`는 "이 작업이 어떤 응집 변경 단위들의 흐름으로 진행되는지"를 소유하고, 각 flow record는 "그 flow 안에서 무엇을 했는지"를 소유한다.
- `000-plan.md`는 증분 갱신하고, 완료된 작업도 삭제하지 않고 요약과 flow reference를 유지한다.

## Flow Record

- `001+` record는 사용자 요청에 따른 개별 flow의 상세 보고서이며, completed flow를 기다리지 말고 각 phase가 끝날 때마다 증분 갱신한다.
- flow 기록 파일은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식을 사용한다.
- `count-pad3`는 `001`, `002`, `003`처럼 3자리 zero-padded 숫자를 사용한다.
- slug는 영어 소문자와 `-`만 사용한다.
- flow 기본 템플릿은 `skills/turn-gate/templates/flow-record-template.md`를 사용한다.
- `000-plan.md` 기본 템플릿은 `skills/turn-gate/templates/plan-template.md`를 사용한다.
- 최소 flow 기록 항목은 user request message, task, flow type, flow scope, current mode, question-routing mode, current core phase, preparation source/result, planned flow list, continuity guard, work, verification, report, next-flow options, residual risk다.
- 운영 flow가 판단, 설계, 범위 확인에서 끝나는 경우 최소 기록 항목의 `planned flow list`는 `follow-up change-unit candidates`로 대체할 수 있으며, 각 후보에는 후보 type, 예상 산출물, 분리 또는 압축 근거, 예상 검증, user-gated handoff 조건을 남긴다.
- flow record는 phase 메모가 아니지만, `preparation`, `work`, `verification`, `reporting` 각 phase가 끝날 때마다 현재 상태로 갱신해야 한다.

## Continuity Guard

- 각 flow record에는 짧은 `Continuity Guard`를 두고, result reporting과 next-flow reopening 직전에 반드시 갱신한다.
- `Continuity Guard`에는 `turn-gate` 활성 여부, question-routing mode, user explicit stop 여부, terminal summary 허용 여부, required next action, last refreshed phase가 포함되어야 한다.
- explicit stop이 확인된 경우에만 `confirmed closure`를 기록한다. 이때 closure source message와 closure recorded phase를 함께 기록해야 하며, source 없는 closure 기록은 stale state로 취급한다.
- pending 또는 superseded question이 있으면 guard에 pending question state, question id 또는 요약, superseded 여부를 기록한다.
- verification이 필요한 flow에서는 guard 또는 verification section에 verification status를 `not-started`, `requested`, `pass`, `fail`, `blocked`, `insufficient` 중 하나로 남긴다.
- result reporting과 next-flow reopening 전에는 active flow record의 `Continuity Guard`를 먼저 읽는다.
- 기록이 없을 때만 재구성하고, 재구성한 guard는 가능한 즉시 flow record에 다시 쓴다.
- 기록이 접근 불가인 경우 missing record처럼 조용히 재구성하지 않는다. 접근 실패를 blocker로 보고하고, 접근이 복구되거나 user-gated decision이 있을 때까지 terminal summary 허용 근거로 삼지 않는다.
- guard의 terminal summary 허용 값은 현재 incoming message 또는 source가 확인된 explicit stop 기록과 일치할 때만 유효하다. stale `terminal summary allowed: yes`나 source 없는 `confirmed closure`는 무효다.
- stale closure state를 발견하면 guard를 `user explicit stop: no`, `terminal summary allowed: no`로 갱신하고, 이전 closure state가 source-less 또는 stale이었다는 note를 남긴다.

## Next Flow Options

- flow record의 `Next Flow Options`에는 사용자 표시 질문에 턴 종료 선택지가 보이지 않는 경우에도 명시적인 turn-end option이 포함되어야 한다.
- 완료된 작업은 삭제하지 않고 요약과 flow reference를 유지한다.
- `000-plan.md`는 날짜 기준 증분 계획과 flow-sequence artifact로, `001+`는 flow 단위 상세 보고서로 취급한다.

## 검토 질문

- `000-plan.md`가 flow sequence와 transition criteria를 소유하고 있는가?
- flow sequence가 preparation 결과에서 파생됐고 각 flow가 preparation/work/verification/reporting 구조를 유지하는가?
- 사용자 메시지 해석과 flow list 설계가 operational-preparation flow로 기록되고, 실행용 planned flows와 섞이지 않았는가?
- 판단, 설계, 범위 확인용 운영 flow 기록이 실제 실행 flow record처럼 보이지 않도록 후속 후보와 handoff 조건을 구분했는가?
- flow sequence가 phase list나 direct user-value list가 아니라 reviewable or commit-sized change-unit list인가?
- 최종 QA/readiness/reporting만 수행하는 항목이 산출물 변경 없이 flow sequence에 들어가지 않았는가?
- active flow record가 현재 phase까지 증분 갱신됐는가?
- visible choices에 turn-end option이 없어도 record에는 turn-end option이 남았는가?
- confirmed closure가 있다면 source explicit stop message가 같이 기록돼 있는가?
- pending 또는 superseded question state가 현재 required next action을 오염시키지 않는가?
