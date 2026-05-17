# turn-gate session-records sub-spec

## 목적

이 문서는 `turn-gate`가 유지하는 `.agents/sessions/{YYYYMMDD}/` 기록, `000-plan.md`, optional `000-self-drive.md`, 개별 flow record, `Continuity Guard`, `Next Flow Options`의 상위 계약을 소유합니다.

구체적인 template 형식은 `templates/plan.md`, `templates/self-drive.md`, `templates/flow.md`가 소유합니다.
여기서 `templates/*.md`는 dev spec-side contract owner를 뜻합니다.
설치 후 실제 시작 파일은 runtime surface의 `skills/turn-gate/templates/plan-template.md`, `skills/turn-gate/templates/self-drive-template.md`, `skills/turn-gate/templates/flow-record-template.md`이며, runtime instructions는 dev-only spec path에 의존하지 않습니다.

## 기록 구조

- active turn-gated task마다 `.agents/sessions/{YYYYMMDD}/000-plan.md` 날짜 기준 plan과 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 상세 flow report 체계를 유지한다.
- self-drive가 active인 경우에만 `.agents/sessions/{YYYYMMDD}/000-self-drive.md` 보조 record를 유지한다.
- `000-plan.md`는 당일 작업의 히스토리, 사용자 요청 목록, flow index, 현재 snapshot, planned flow sequence, 완료 flow 요약을 소유하되 bounded date-level index로 유지한다. 세부 형식은 `templates/plan.md`가 소유한다.
- self-drive가 active인 경우 `000-plan.md`는 self-drive-specific date-level snapshot으로 self-drive active 여부와 `000-self-drive.md` pointer만 소유한다. 일반 planned flow sequence section은 date-level routing snapshot으로 남을 수 있지만, self-drive sequence-level state의 canonical owner가 아니다.
- `000-self-drive.md`는 sequence-level state를 소유한다: sequence objective, planned flow list, active flow index, allowed autonomous actions, prohibited autonomous actions, approval-sensitive checkpoints, endpoint, blocker return conditions, progress note, progress ledger.
- `active_flow_index`는 0-based machine field이며, `000-self-drive.md`에는 사람이 읽을 수 있는 current flow label도 함께 남긴다. numeric index가 planned flow numbering과 충돌하거나 모호하면 flow name/file 또는 slug를 기준으로 reconcile하고, 해소할 수 없으면 다음 autonomous flow로 진행하지 않는다.
- `000-plan.md`는 각 flow의 상세 scope, non-goal, approval boundary, evidence, verification detail을 반복하지 않는다. 필요한 경우 짧은 snapshot만 복사하고, canonical detail은 해당 `001+` flow record가 소유한다.
- `000-plan.md`의 현재 계획은 action checklist가 아니라 planned flow sequence여야 한다.
- session record에는 flow type을 구분할 수 있어야 한다: `operational-preparation` 또는 `change-unit`.
- flow type, planned flow boundary, 후속 후보와 active execution flow의 구분은 `core/flow-boundaries.md`가 소유한다.
- session record는 그 boundary 판단을 실제 `.agents/sessions/{YYYYMMDD}/000-plan.md`와 `001+` flow record에 어떻게 남기는지 소유한다.
- 사용자 메시지 기반 bootstrap을 기록할 때는 operational-preparation flow와 그 결과 planned change-unit flow list 또는 후속 후보를 구분한다.
- 실행이 아니라 판단, 설계, 범위 확인만 요청된 경우 실제 실행 flow와 구분한다. 이 경우 `operational-preparation` record는 후속 실행 후보, scope/non-goal, target ambiguity 판단, verification expectation을 기록하고 종료할 수 있다.
- 운영 flow에 남긴 후속 실행 후보는 아직 시작된 `change-unit` flow가 아니다. 실제 실행으로 이어질 때만 별도 `001+` flow record 또는 다음 count의 `change-unit` record를 만든다.
- 각 planned flow에는 flow type, flow 목적, 왜 이 경계가 필요한지, 완료 기준, 다음 flow로 넘어가는 조건이 드러나야 한다.
- concrete task에서 만든 planned flow sequence에는 preparation source, preparation result, planned flow list가 드러나야 한다.
- 실행이 아니라 판단, 설계, 범위 확인이 목적인 경우에는 `planned flow list` 대신 `follow-up change-unit candidates`로 기록할 수 있다. 이렇게 기록한 후보는 실행 승인이 있거나 다음 flow로 선택되기 전까지 완료/진행 flow로 세지 않는다.
- 기존 verbose session history는 명시적인 migration 요청이 없는 한 재작성하지 않는다. 과거 기록은 historical operational artifact이며, 현재 runtime behavior의 source of truth는 현재 spec과 template이다.
- 각 flow는 기본적으로 `preparation -> work -> verification -> reporting -> next-flow` 단계를 가진다.
- 사용자 메시지 기반 preparation이면 deep-interview result와 사용자 의도에 맞춘 flow list를 기록한다.
- 비 사용자 메시지 기반 preparation이면 수정 범위, 현재 상태, 대상 파일, stale assumption, 실행 전 조건 확인 결과를 기록한다.
- 압축했다면 어느 flow가 preparation/work/verification/reporting/next-flow를 함께 소유하는지 설명한다. 단, phase를 쪼갠 것을 flow sequence로 포장하지 않는다.
- 세부 작업 단계는 해당 `001+` flow record의 execution log와 verification에 둔다.
- `000-plan.md`는 "이 작업이 어떤 응집 변경 단위들의 흐름으로 진행되는지"를 소유하고, 각 flow record는 "그 flow 안에서 무엇을 했는지"를 소유한다.
- `000-plan.md`는 증분 갱신하고, 완료된 작업도 삭제하지 않고 요약과 flow reference를 유지한다.
- `000-plan.md`의 `Flow Index`와 `Completed Flow Summaries`는 flow당 한 줄의 compact entry로 유지한다. 상세 목적, 완료 기준, 검증 근거, next-flow 후보, residual risk는 해당 flow record로 위임한다.
- `Planned Flow Sequence`에는 현재 선택됐거나 미래에 실행할 planned flow만 둔다. 완료된 flow는 `Flow Index`와 `Completed Flow Summaries`의 compact entry로만 남긴다.
- `Open Risks`에는 active date-level risk만 둔다. 완료됐거나 flow-local인 risk는 해당 flow record에 남기고 plan에서 반복하지 않는다.
- `User Requests Today`는 최근 요청 중심의 routing context로 유지할 수 있으며, 오래된 요청의 canonical detail은 각 flow record가 소유한다.

## Flow Record

- `001+` record는 사용자 요청에 따른 개별 flow의 상세 보고서이며, completed flow를 기다리지 말고 각 phase가 끝날 때마다 증분 갱신한다.
- flow 기록 파일은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식을 사용한다.
- `count-pad3`는 `001`, `002`, `003`처럼 3자리 zero-padded 숫자를 사용한다.
- slug는 영어 소문자와 `-`만 사용한다.
- flow 기본 템플릿은 `skills/turn-gate/templates/flow-record-template.md`를 사용한다.
- `000-plan.md` 기본 템플릿은 `skills/turn-gate/templates/plan-template.md`를 사용한다.
- `000-self-drive.md` 기본 템플릿은 `skills/turn-gate/templates/self-drive-template.md`를 사용하며, self-drive가 active일 때만 만든다.
- 최소 flow 기록 항목은 user request message, task, flow type, flow scope, current phase, continuity guard, flow contract, execution log, verification, report, next-flow options, residual risk다. 세부 형식은 `templates/flow.md`가 소유한다.
- flow record의 parent-plan relation은 같은 날짜 디렉터리의 `000-plan.md`와 그 `Flow Index`에서 파생된다. 일반 runtime에서는 중복 `parent_plan` frontmatter를 요구하지 않는다.
- self-drive가 active인 경우 각 flow record는 기본 template에 전용 section을 상시 노출하지 않고, 전체 sequence를 반복하지 않으며, 자기 flow의 sequence position, local progress note, next handoff, blocker return condition만 가장 자연스러운 기존 section에 flow-local snapshot으로 남긴다.
- `000-plan.md`가 `self_drive_status: inactive`인데 `self_drive_record`가 남아 있거나 sidecar 파일이 존재하면 stale sidecar state로 취급한다. 이 경우 leftover sidecar를 active continuation authority로 사용하지 말고, pointer/status를 정리하거나 user-gated clarification을 기록한 뒤 진행한다. 필요하면 historical context로만 읽고, routing authority는 current `000-plan.md` snapshot을 따른다.
- 같은 정보를 `000-plan.md`와 flow record에 상세 반복하지 않는다. `000-plan.md`에 복사되는 값은 active snapshot 또는 flow index summary로만 유지한다.
- 운영 flow가 판단, 설계, 범위 확인에서 끝나는 경우 최소 기록 항목의 `planned flow list`는 `follow-up change-unit candidates`로 대체할 수 있으며, 각 후보에는 후보 type, 예상 산출물, 분리 또는 압축 근거, 예상 검증, user-gated handoff 조건을 남긴다.
- flow record는 phase 메모가 아니지만, `preparation`, `work`, `verification`, `reporting`, `next-flow` 각 phase가 끝날 때마다 현재 상태로 갱신해야 한다.

## Continuity Guard

- 각 flow record에는 짧은 `Continuity Guard`를 두고, result reporting과 next-flow reopening 직전에 반드시 갱신한다.
- `Continuity Guard`에는 `turn-gate` 활성 여부, question-routing mode, user explicit stop 여부, terminal summary 허용 여부, required next action, last refreshed phase가 포함되어야 한다.
- explicit stop이 확인된 경우에만 `confirmed closure`를 기록한다. 이때 closure source message와 closure recorded phase를 함께 기록해야 하며, source 없는 closure 기록은 stale state로 취급한다.
- pending 또는 superseded question이 있으면 guard에 pending question state, question id 또는 요약, superseded 여부를 기록한다.
- verification이 필요한 flow에서는 guard 또는 verification section에 verification status를 `not-started`, `requested`, `pass`, `fail`, `blocked`, `insufficient` 중 하나로 남긴다.
- verification method는 `clean-context`, `normal`, `not-required` 중 하나로 verification section에 남긴다. method는 status가 아니므로 `verification_status` frontmatter 값으로 쓰지 않는다.
- verification 상세 근거, command/evidence 목록, verifier id, not-required reason, residual uncertainty는 verification section이 소유하고, guard에는 현재 status만 둔다.
- result reporting과 next-flow reopening 전에는 active flow record의 `Continuity Guard`를 먼저 읽는다.
- record recovery는 `not-yet-created`, `unexpectedly missing`, `inaccessible`, `stale` 상태를 분리한다. 새 flow를 시작하며 아직 파일이 만들어지지 않은 경우에만 template 기반 first creation을 허용한다.
- active plan, active self-drive sidecar, 현재 응답 phase, 또는 직전 flow handoff가 특정 `001+` flow record를 가리키는데 그 record가 없으면 `unexpectedly missing active record`로 취급한다. 이 경우 조용히 재구성하지 말고 blocker로 보고하거나 user-gated decision을 받아야 한다.
- active flow record가 존재하지만 읽기 실패, 권한 오류, lock, parse failure, encoding failure, partial write처럼 canonical guard를 신뢰할 수 없으면 `inaccessible record`로 취급한다. 접근 실패를 blocker로 보고하고, 접근이 복구되거나 user-gated decision이 있을 때까지 terminal summary 허용 근거로 삼지 않는다.
- 새 record first creation 또는 명시적으로 승인된 recovery를 수행한 경우, 재구성한 guard는 가능한 즉시 flow record에 다시 쓰고 `continuity_note`에 근거와 source를 남긴다.
- `000-plan.md`, `000-self-drive.md`, active flow record의 `current_phase`, 직전 handoff가 서로 다른 phase 또는 flow를 가리키면 `stale routing mismatch`로 취급한다. 최신 source와 handoff를 확인해 reconcile하고, 해소할 수 없으면 user-gated clarification으로 돌아간다.
- `stale routing mismatch`는 terminal close, successful completion, next-flow skip, self-drive continuation authority가 아니다. 해소하기 전에는 더 닫힌 상태를 임의로 선택하지 않는다.
- 기록 접근 blocker의 사용자-facing prefix는 발견 시점에 따른다. result reporting 전에 발견하면 `[reporting]`, next-flow reopening 직전에 발견하면 `[next-flow]`로 blocker routing을 연다. 관련 phase protocol을 적용 중이면 `[reporting/review-loop]`처럼 optional slash suffix를 붙일 수 있다.
- guard의 terminal summary 허용 값은 현재 incoming message 또는 source가 확인된 explicit stop 기록과 일치할 때만 유효하다. stale `terminal summary allowed: yes`나 source 없는 `confirmed closure`는 무효다.
- stale closure state를 발견하면 guard를 `user explicit stop: no`, `terminal summary allowed: no`로 갱신하고, 이전 closure state가 source-less 또는 stale이었다는 note를 남긴다.

## Record Recovery Decision Table

| 상태 | 예시 | 허용 동작 | 금지 동작 |
| --- | --- | --- | --- |
| `not-yet-created plan` | 당일 첫 turn-gated work이고 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 없음 | plan template로 first creation | 이전 날짜 기록을 active authority로 승격 |
| `not-yet-created flow` | 새 flow count/slug가 방금 선택됐고 아직 `001+` record를 만들기 전 | flow template로 first creation | source 없는 closure나 pass status를 추정 |
| `unexpectedly missing active record` | `000-plan.md`가 `022-*.md`를 가리키는데 파일이 없음 | blocker 보고 또는 user-gated recovery 선택 | 같은 slug로 조용히 재구성하고 reporting/next-flow 진행 |
| `inaccessible active record` | 권한 오류, lock, parse failure, partial write, encoding failure | blocker 보고, 복구 후 재시도 또는 사용자 결정 | missing record처럼 재구성하거나 terminal summary 허용 |
| `stale closure state` | `terminal summary allowed: yes`인데 source message가 없거나 현재 메시지와 불일치 | stop/summary 허용 값을 `no`로 reset하고 note 기록 | stale closure를 explicit stop으로 사용 |
| `stale self-drive sidecar` | `000-plan.md`는 self-drive inactive인데 sidecar가 남아 있음 | historical context로만 읽고 current plan 기준으로 routing | sidecar의 active_flow_index를 continuation authority로 사용 |
| `stale routing mismatch` | plan은 `next-flow`, flow record는 `reporting`처럼 phase/handoff가 불일치 | 최신 source/handoff로 reconcile 또는 clarification | 더 닫힌 상태를 임의 선택하거나 next-flow skip |

## Next Flow Options

- flow record의 `Next Flow Options`에는 사용자 표시 질문에 턴 종료 선택지가 보이지 않는 경우에도 명시적인 turn-end option이 포함되어야 한다.
- `Next Flow Options`는 flow record가 소유한다. `000-plan.md`는 선택 결과나 active next flow만 snapshot으로 반영한다.
- 완료된 작업은 삭제하지 않고 한 줄 요약과 flow reference를 유지한다.
- `000-plan.md`는 날짜 기준 증분 계획과 flow-sequence artifact로, `001+`는 flow 단위 상세 보고서로 취급한다.

## Read-Only Write Boundary

- 일반적인 `read-only`, `no-edit`, `파일만 읽고`, `source 수정 금지`, `코드는 건드리지 마` 요청은 기본적으로 target/source/spec/runtime/release surface 변경 금지로 해석한다. 이 경우 `.agents/sessions/{YYYYMMDD}/` session record는 turn continuity를 위한 운영 기록으로 작성할 수 있다.
- 사용자가 `아무 파일도 쓰지 마`, `어떤 파일도 만들지 마`, `기록 파일도 쓰지 마`, `세션 기록 남기지 마`, `무기록으로 답만 해`처럼 workspace-wide no-write 또는 no-record를 명시하면 session record 작성도 금지된다.
- read-only 요청을 받으면 work boundary에 target/source 변경 금지와 session record 운영 기록 작성 여부를 분리해 남긴다.
- no-write/no-record 제약 때문에 session record를 쓸 수 없으면, 기록을 만들기 전에 user-gated clarification 또는 blocker로 라우팅한다. 필요한 경우 in-memory continuity는 질문 또는 blocker 보고를 끝내는 데 필요한 최소 범위로만 유지한다.
- `세션 기록 남기지 마`, `무기록으로 답만 해`처럼 write만 금지하는지 기존 record read까지 금지하는지 애매하면 session record를 읽기 전에 확인한다. 단, 사용자가 기존 상태 보고를 요청했고 record read를 명시적으로 금지하지 않았다면 read-only record inspection은 가능하지만 write는 금지된다.
- `수정이 필요하면 멈춰`는 target/source 수정 필요성이 생기면 멈추라는 의미로 해석한다. session record 작성도 금지하는지 애매하면 쓰기 전에 확인한다.
- clean-context verifier나 read-only subagent의 `read-only`는 검증 대상 수정 금지와 subagent edit permission 금지를 뜻한다. 별도 no-write/no-record 지시가 없는 한 main turn-gate session record 운영 기록까지 자동 금지하지 않는다.

## Read-Only Boundary Decision Table

| 사용자 제약 | session record write | 기대 라우팅 |
| --- | --- | --- |
| `파일만 읽고 정리해줘`, `수정하지 마` | 허용 가능 | target/source는 변경하지 않고 운영 기록에 boundary를 남김 |
| `코드는 건드리지 마`, `source 수정 금지` | 허용 가능 | source change는 후속 후보로만 남김 |
| `read-only로 검증만 해줘` | 허용 가능 | 검증 대상과 verifier는 read-only, 운영 기록은 유지 |
| `수정이 필요하면 멈춰` | 보통 허용, 애매하면 확인 | target/source 수정 필요 시 blocker/question |
| `아무 파일도 쓰지 마`, `어떤 파일도 만들지 마` | 금지 | session record 생성 전 clarification/blocker |
| `세션 기록도 남기지 마`, `무기록으로 답만 해` | 금지 | 기록 없이 최소 in-memory blocker/report 또는 clarification |
| `세션 기록 남기지 말고 상태만 알려줘` | write 금지, read는 명시 금지 여부 확인 | 기존 record read가 필요한지 먼저 확인하거나 read-only 상태 보고 |

## 검토 질문

- `000-plan.md`가 flow sequence와 transition criteria를 소유하고 있는가?
- self-drive가 active라면 `000-plan.md`가 self-drive active 여부와 sidecar pointer만 self-drive-specific snapshot으로 소유하고 `000-self-drive.md`가 sequence-level state를 소유하는가?
- `000-plan.md`가 상세 flow contract와 verification evidence를 반복하지 않고 snapshot/index만 유지하는가?
- `000-plan.md`의 index와 completed summaries가 flow당 한 줄 compact entry로 유지되는가?
- `Planned Flow Sequence`에 완료된 flow의 stale 계획이 남지 않는가?
- `Open Risks`가 active date-level risk만 담고 flow-local risk를 반복하지 않는가?
- flow record가 scope, non-goals, approval boundary, evidence, verification detail의 canonical owner인가?
- template 형식 세부 판단이 `templates/plan.md`와 `templates/flow.md`로 위임되는가?
- flow sequence가 preparation 결과에서 파생됐고 각 flow가 preparation/work/verification/reporting/next-flow 구조를 유지하는가?
- 사용자 메시지 해석과 flow list 설계가 operational-preparation flow로 기록되고, 실행용 planned flows와 섞이지 않았는가?
- 판단, 설계, 범위 확인용 운영 flow 기록이 실제 실행 flow record처럼 보이지 않도록 후속 후보와 handoff 조건을 구분했는가?
- flow sequence가 `core/flow-boundaries.md`의 planned flow boundary 규칙에 맞게 기록됐는가?
- 기존 session history를 자동 migration하지 않고 향후 기록에 새 template을 적용하는가?
- 최종 QA/readiness/reporting만 수행하는 항목이 산출물 변경 없이 flow sequence에 들어가지 않았는가?
- active flow record가 현재 phase까지 증분 갱신됐는가?
- self-drive가 active라면 `000-self-drive.md`가 sequence-level state를 소유하고 active flow record가 flow-local sequence snapshot만 소유하는가?
- active flow index와 current flow label이 함께 있어 numeric index ambiguity 없이 다음 flow를 식별할 수 있는가?
- visible choices에 turn-end option이 없어도 record에는 turn-end option이 남았는가?
- confirmed closure가 있다면 source explicit stop message가 같이 기록돼 있는가?
- pending 또는 superseded question state가 현재 required next action을 오염시키지 않는가?
- read-only/no-edit 요청에서 target/source 변경 금지와 session record 운영 기록 작성 여부를 분리했는가?
- workspace-wide no-write 또는 no-record 요청에서 session record write를 수행하지 않고 blocker/clarification으로 라우팅했는가?
