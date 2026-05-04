# turn-gate session-records sub-spec

## 목적

이 문서는 `turn-gate`가 유지하는 `.agents/sessions/{YYYYMMDD}/` 기록, `000-plan.md`, 개별 flow record, `Continuity Guard`, `Next Flow Options` 계약을 소유합니다.

## 기록 구조

- active turn-gated task마다 `.agents/sessions/{YYYYMMDD}/000-plan.md` 날짜 기준 plan과 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 상세 flow report 체계를 유지한다.
- `000-plan.md`는 당일 작업의 히스토리, 사용자 요청 목록, flow index, 현재 계획, 완료 flow 요약을 소유한다.
- `000-plan.md`의 현재 계획은 action checklist가 아니라 planned flow sequence여야 한다.
- 각 planned flow에는 flow 목적, 왜 이 flow가 필요한지, 완료 기준, 다음 flow로 넘어가는 조건이 드러나야 한다.
- 세부 작업 단계는 해당 `001+` flow record의 plan/work/verification에 둔다.
- `000-plan.md`는 "이 작업이 어떤 flow들의 흐름으로 진행되는지"를 소유하고, 각 flow record는 "그 flow 안에서 무엇을 했는지"를 소유한다.
- `000-plan.md`는 증분 갱신하고, 완료된 작업도 삭제하지 않고 요약과 flow reference를 유지한다.

## Flow Record

- `001+` record는 사용자 요청에 따른 개별 flow의 상세 보고서이며, completed flow를 기다리지 말고 각 phase가 끝날 때마다 증분 갱신한다.
- flow 기록 파일은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식을 사용한다.
- `count-pad3`는 `001`, `002`, `003`처럼 3자리 zero-padded 숫자를 사용한다.
- slug는 영어 소문자와 `-`만 사용한다.
- flow 기본 템플릿은 `skills/turn-gate/templates/flow-record-template.md`를 사용한다.
- `000-plan.md` 기본 템플릿은 `skills/turn-gate/templates/plan-template.md`를 사용한다.
- 최소 flow 기록 항목은 user request message, task, flow scope, current mode, question-routing mode, continuity guard, analysis, plan, work, verification, result report, next-flow options, residual risk다.
- flow record는 phase 메모가 아니지만, `analysis`, `plan`, `work`, `verification`, `result reporting` 각 phase가 끝날 때마다 현재 상태로 갱신해야 한다.

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

## Next Flow Options

- flow record의 `Next Flow Options`에는 사용자 표시 질문에 턴 종료 선택지가 보이지 않는 경우에도 명시적인 turn-end option이 포함되어야 한다.
- 완료된 작업은 삭제하지 않고 요약과 flow reference를 유지한다.
- `000-plan.md`는 날짜 기준 증분 계획과 flow-sequence artifact로, `001+`는 flow 단위 상세 보고서로 취급한다.

## 검토 질문

- `000-plan.md`가 flow sequence와 transition criteria를 소유하고 있는가?
- active flow record가 현재 phase까지 증분 갱신됐는가?
- visible choices에 turn-end option이 없어도 record에는 turn-end option이 남았는가?
- confirmed closure가 있다면 source explicit stop message가 같이 기록돼 있는가?
- pending 또는 superseded question state가 현재 required next action을 오염시키지 않는가?
