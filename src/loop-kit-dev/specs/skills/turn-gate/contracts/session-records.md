# turn-gate session-records 계약

## 목적

이 계약은 활성 `turn-gate` work에서 사용하는 operational continuity record를 소유합니다.

## 파일

- `.agents/sessions/{YYYYMMDD}/000-plan.md`: date-level routing context, active flow pointer, required next action, request history, compact flow index, planned current/future sequence, completed summaries, explicit turn-end availability, active date-level risks를 소유합니다.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`: 하나의 active flow contract, 필요 시 raw request, interpretation, scope, non-goals, approval boundary, execution log, verification, report, next-flow options, residual risk를 소유합니다.
- `.agents/sessions/{YYYYMMDD}/000-self-drive.md`: self-drive가 active일 때만 사용하는 optional self-drive sequence state입니다.

runtime template은 `skills/turn-gate/templates/`의 파일을 사용합니다.

flow filename은 zero-padded counter와 lowercase English slug를 사용합니다. active flow boundary가 바뀌면 새 flow record가 필요합니다. 같은 flow가 계속 active이거나 reporting 전 자기 Continuity Guard를 고치는 경우에만 이전 flow를 갱신할 수 있습니다.

flow record는 completed flow를 기다리지 않고 각 phase 시작과 종료마다 현재 상태로 증분 갱신합니다.
`flow`가 산출한 phase start/end record checkpoint expectation을 적용해 `000-plan.md`와 active flow record 중 어느 표면이 갱신돼야 하는지 구분합니다.
active flow pointer, date-level required next action, planned/current sequence가 바뀌면 `000-plan.md`를 갱신합니다.
같은 active flow 내부의 current phase, execution log, verification evidence, report outcome, residual risk, handoff condition이 바뀌면 active flow record를 갱신합니다.

flow record는 compact contract를 유지하더라도 최소한 다음 섹션을 가져야 합니다.

- `Flow Contract`
- `Optional Risky Actions`
- `Execution Log`
- `Verification`
- `Report`
- `Next Flow Options`
- `Residual Risk`

flow record frontmatter 또는 Continuity Guard에는 현재 phase, required next action, closure 관련 필드, pending/superseded question state, verification status, continuity note가 드러나야 합니다. 정확한 template 문구는 runtime template이 소유하지만, 이 최소 정보가 빠지면 compaction 또는 interruption 뒤 다음 행동을 복구하기 어렵습니다.

## 중복 방지

`000-plan.md`는 compact하게 유지합니다. detailed scope, evidence, verification, residual risk, self-drive sequence detail은 plan에 반복하지 않고 active flow record 또는 self-drive sidecar에 둡니다.

`000-plan.md`의 `Flow Index`와 `Completed Flow Summaries`는 flow당 한 줄의 compact entry로 유지하고 완료된 flow 요약을 삭제하지 않습니다. `Planned Flow Sequence`에는 현재 또는 미래 selected flow만 두며, 판단/설계/범위 확인에서 나온 후속 후보는 선택 전까지 active 또는 completed flow가 아닌 candidate handoff로 구분합니다.

self-drive가 active이면 `000-plan.md`는 status와 sidecar pointer만 저장합니다. `000-self-drive.md`가 sequence-level state를 소유합니다.

raw user request text를 기록할 때는 interpretation 또는 summary와 분리해야 합니다. raw request field 안에서는 normalize, translate, correct, soften, merge, infer missing words를 하지 않습니다. summary와 interpretation은 별도로 작성할 수 있습니다.

## Continuity Guard

모든 flow record는 Continuity Guard를 둡니다.

- turn-gate active state
- question-routing mode
- user explicit stop
- terminal summary allowed
- required next action
- last refreshed phase
- confirmed closure
- closure source message
- closure recorded phase
- pending question state
- pending question id or summary
- superseded question id or summary
- verification status
- continuity note

guard는 각 phase 시작과 종료, reporting, next-flow reopening 전에 갱신합니다.

현재 source-recorded explicit stop만 terminal closure authority를 설정할 수 있습니다. stale 또는 source-less closure state는 열린 continuity 상태로 reset해야 합니다.

guard는 compaction 또는 interruption 뒤에도 다음 행동을 알 수 있게 해야 합니다. 최소한 `turn-gate` active 여부, pending question 여부, verification status, closure 허용 여부, 다음 required action을 보여야 합니다.

`verification status`는 기록 진행 상태와 결과 상태를 모두 표현할 수 있습니다. work 전 또는 검증 전에는 `not-started`, verifier 요청 후 결과 전에는 `requested`, 검증 결과가 있으면 `pass`, `fail`, `blocked`, `insufficient` 중 하나를 사용합니다. `not-started`와 `requested`는 terminal close나 successful reporting 근거가 아닙니다.

## 복구

missing/not-created record와 unexpectedly missing 또는 inaccessible active record를 구분합니다. 있어야 하는 active record를 조용히 재구성하지 않습니다. inaccessible active record는 blocker recovery로 라우팅합니다.

recovery case:

- not-yet-created plan: runtime template으로 첫 plan을 만듭니다.
- not-yet-created flow: runtime template으로 선택된 새 flow record를 만듭니다.
- unexpectedly missing active record: blocker를 보고하거나 recovery를 질문합니다.
- inaccessible active record: access가 복구되거나 사용자가 recovery를 선택할 때까지 blocker를 보고합니다.
- stale closure state: closure authority를 reset하고 recovery를 기록합니다.
- stale self-drive sidecar: plan이 self-drive inactive라고 하면 historical로 취급합니다.
- stale routing mismatch: latest source에서 reconcile하거나 질문합니다.

read-only 요청은 보통 target/source 변경을 금지하는 것이지 session record를 금지하는 것이 아닙니다. 사용자가 모든 write 또는 record 생성을 금지한 경우에만 record를 쓰지 않습니다.

사용자 표현이 `no-record`, `기록 남기지 마`, `세션 기록 없이`처럼 session record의 읽기까지 금지하는지, 새 기록 생성/수정만 금지하는지 모호하면 기존 session record를 읽기 전에 clarification으로 라우팅합니다. 모호성이 풀리기 전에는 최소한의 in-memory continuity만 유지합니다.
