# turn-gate session-records 계약

## 소유 범위

활성 `turn-gate` work에서 사용하는 operational continuity record.

## 파일 계약

- `.agents/sessions/{YYYYMMDD}/000-plan.md`: date-level routing card입니다. Frontmatter가 active flow pointer, next action, closure flags, self-drive status와 sidecar pointer, unapproved actions, active skill list를 소유합니다. Body는 compact recent request list, active/recent/archive flow index, continuity note만 소유합니다.
- `.agents/sessions/{YYYYMMDD}/000-review.md`: optional date-level retrospective notes를 소유합니다. 이 파일은 flat tagged list로 reusable lesson, process correction, follow-up candidate를 기록하며 active routing state, raw flow log, verification authority, closure authority를 소유하지 않습니다.
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`: 하나의 active flow의 compact `Contract`, `Execution Log`, `Result`를 소유하고, 필요할 때만 raw request와 `Risky Action`을 추가합니다.
- `.agents/sessions/{YYYYMMDD}/000-self-drive.md`: self-drive가 active일 때만 사용하는 optional self-drive sequence state입니다.

runtime template은 `skills/turn-gate/templates/`의 파일을 사용합니다.

flow filename은 zero-padded counter와 lowercase English slug를 사용합니다. active flow boundary가 바뀌면 새 flow record가 필요합니다. 같은 flow가 계속 active이거나 reporting 전 자기 continuity metadata를 고치는 경우에만 이전 flow를 갱신할 수 있습니다.

flow record는 completed flow를 기다리지 않고 각 phase 시작과 종료마다 현재 상태로 증분 갱신합니다.
`flow`가 산출한 phase start/end record checkpoint expectation을 적용해 `000-plan.md`와 active flow record 중 어느 표면이 갱신돼야 하는지 구분합니다.
active flow pointer, date-level required next action, active skill list, self-drive status, unapproved action state, turn-level routing이 바뀌면 `000-plan.md` frontmatter를 갱신합니다.
같은 active flow 내부의 current phase, execution log, verification evidence, report outcome, residual risk, handoff condition이 바뀌면 active flow record를 갱신합니다.

flow record는 기본적으로 compact formal style을 사용합니다. 기본 섹션은 다음입니다.

- `Contract`
- `Execution Log`
- `Result`

approval-sensitive action이 있으면 `Risky Action` 섹션을 추가합니다. readiness, verification, build, generated release surface 갱신은 그 자체로 commit, publish, release, version bump, destructive/external action 실행 권한을 만들지 않습니다. 실패, blocked, insufficient 결과에는 `Result` 아래에 non-pass routing을 추가할 수 있습니다. raw user text가 중요할 때만 `Contract`에 raw request를 추가합니다.

flow record frontmatter는 긴 boolean 나열보다 formal metadata를 우선합니다. 현재 phase, verification status, next action, turn-gate/terminal-summary/pending-question 같은 flags, question state, continuity note가 드러나야 합니다. 기본 필드는 `answered_question`이고, 대기 중인 질문이 있으면 `pending_question`을 추가합니다. `question_state`처럼 새 동의어를 만들지 않습니다. 정확한 template 문구는 runtime template이 소유하지만, 이 최소 정보가 빠지면 compaction 또는 interruption 뒤 다음 행동을 복구하기 어렵습니다.

## 중복 방지 계약

`000-plan.md`는 compact routing card로 유지합니다. Git으로 재구성 가능한 branch/latest commit, detailed scope, evidence, verification, residual risk, self-drive sequence detail은 plan에 반복하지 않고 active flow record, self-drive sidecar, 또는 실제 tool readback에 둡니다.
skill list는 frontmatter의 `active_skills`에 skill 이름만 기록합니다. 사용 지점 설명, 전체 사용 가능 skill catalog, 후보 단계의 가능성, 이미 끝난 flow의 상세 사용 내역은 plan에 반복하지 않습니다.

`000-plan.md`의 `Flow Index`는 active, recent, archive만 둡니다. `active`는 현재 flow, `recent`는 바로 이전 handoff flow, `archive`는 오래된 flow range와 individual flow record로의 복구 가능성을 가리킵니다. 완료 flow 전체 목록, completed summaries, planned old flows는 plan에 누적하지 않습니다.

`Recent Requests`는 현재 요청과 복구에 필요한 직전 routing signal만 짧게 둡니다. `user requested ...` 같은 문장형 이력보다 `[current] compact note` 형식을 우선합니다. 오래된 raw request와 interpretation은 개별 flow record가 소유합니다.

self-drive가 active이면 `000-plan.md`는 status와 sidecar pointer만 저장합니다. `000-self-drive.md`가 sequence-level state를 소유합니다.

`000-review.md`는 flow 순서나 카테고리별 헤더가 아니라 flat tagged list를 사용합니다. 각 항목은 `[conversation]`, `[records]`, `[docs]`, `[code-structure]`, `[verification]`, `[git]`, `[release]` 같은 bracketed axis tag 하나로 시작합니다. tag는 open-ended이며 그날 의미 있는 축에 맞게 추가하거나 바꿀 수 있습니다. 항목에는 필요한 경우 invalid/correct 예시, evidence, follow-up candidate를 짧은 sub-bullet으로 둡니다.

`000-review.md`는 복구 표면이 아닙니다. 현재 active flow, required next action, pending question, verification status 같은 continuity state는 `000-plan.md`와 active flow record에 남깁니다. review item은 다음 세션이나 spec 개선에 재사용할 수 있는 관찰과 교정만 남깁니다.

raw user request text를 기록할 때는 interpretation 또는 summary와 분리해야 합니다. raw request field 안에서는 normalize, translate, correct, soften, merge, infer missing words를 하지 않습니다. summary와 interpretation은 별도로 작성할 수 있습니다.

## Continuity Metadata 계약

모든 flow record는 compact continuity metadata를 둡니다.

- `phase`
- `verification_status`
- `next_action`
- `flags`
- `answered_question`, 그리고 대기 중인 질문이 있으면 `pending_question`
- `continuity`

`flags`는 `turn_gate_active`, `terminal_summary_blocked`, `question_pending`, `blocked`, `approval_required`, `explicit_stop_recorded`처럼 recovery에 필요한 상태만 나열합니다. source-recorded explicit stop이 있으면 closure source를 `continuity` 또는 별도 compact field에 남깁니다.

metadata는 각 phase 시작과 종료, reporting, next-flow reopening 전에 갱신합니다.

현재 source-recorded explicit stop만 terminal closure authority를 설정할 수 있습니다. stale 또는 source-less closure state는 열린 continuity 상태로 reset해야 합니다.

metadata는 compaction 또는 interruption 뒤에도 다음 행동을 알 수 있게 해야 합니다. 최소한 `turn-gate` active 여부, pending question 여부, verification status, closure 허용 여부, 다음 required action을 보여야 합니다.
flow start recovery 때는 `000-plan.md`의 current/planned skill list를 기준으로 필요한 skill을 다시 읽고, 이전 runtime context에 남은 skill 본문만 신뢰하지 않습니다.

`verification_status`는 기록 진행 상태와 결과 상태를 모두 표현할 수 있습니다. work 전 또는 검증 전에는 `not-started`, verifier 요청 후 결과 전에는 `requested`, 검증 결과가 있으면 `pass`, `fail`, `blocked`, `insufficient` 중 하나를 사용합니다. 이전 flow state를 보존할 때 `preserved` 같은 새 status 값을 만들지 말고 기존 값을 그대로 유지한 뒤 보존 사실을 `continuity`에 기록합니다. `not-started`와 `requested`는 terminal close나 successful reporting 근거가 아닙니다.

## 복구 계약

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
