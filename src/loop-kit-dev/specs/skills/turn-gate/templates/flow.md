# turn-gate flow record template sub-spec

## 목적

이 문서는 `turn-gate`의 `skills/turn-gate/templates/flow-record-template.md` 형식 계약을 소유합니다.
Spec file 이름은 `flow.md`이지만 runtime 시작 파일 이름은 `flow-record-template.md`입니다.
runtime artifact가 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{slug}.md` flow record이기 때문에 runtime file name에 `record`가 포함됩니다.

`001+` flow record는 하나의 사용자 요청 기반 flow가 실제로 무엇을 소유했고, 무엇을 했고, 어떻게 검증됐는지를 기록하는 canonical detail artifact입니다.

## 소유

- original user request message, including raw text when needed
- task, flow type, flow scope, path-derived parent-plan relation
- current phase
- `Continuity Guard`
- flow contract
- execution log
- verification detail
- report
- next-flow options
- residual risk

## 비소유

- date-level work history
- all-day user request index
- completed flow index
- planned flow sequence 전체의 top-level 관리

위 항목은 `templates/plan.md`가 소유하는 `000-plan.md`에 둡니다.

## 필수 구조

`flow-record-template.md`는 다음 구조를 유지합니다.

1. YAML frontmatter
2. `Flow Contract`
3. `Optional Risky Actions`
4. `Execution Log`
5. `Verification`
6. `Report`
7. `Next Flow Options`
8. `Residual Risk`

## Frontmatter 규격

`flow-record-template.md`는 Markdown YAML frontmatter로 짧은 상태값과 포인터를 먼저 드러냅니다.
파일 경로에서 파생되는 `date`, `record_path`, `slug`, `parent_plan`은 frontmatter에 반복하지 않습니다.
`parent_plan`은 같은 날짜 디렉터리의 `000-plan.md`를 뜻하는 path-derived relation입니다.
일반 runtime에서는 중복 `parent_plan` frontmatter를 추가하지 않습니다. 경로를 잃는 export나 aggregation이 필요하면 base runtime template이 아니라 해당 tooling metadata가 별도로 소유해야 합니다.

frontmatter에는 다음 필드를 둡니다.

- `task`
- `flow_type`
- `flow_scope`
- `current_phase`
- `turn_gate_active`
- `question_routing_mode`
- `user_explicit_stop`
- `terminal_summary_allowed`
- `required_next_action`
- `last_refreshed_phase`
- `confirmed_closure`
- `closure_source_message`
- `closure_recorded_phase`
- `pending_question_state`
- `pending_question_id_or_summary`
- `superseded_question_id_or_summary`
- `verification_status`
- `continuity_note`
- `preparation_source`
- `scope_lock_status`

짧은 원문 요약이 필요한 경우 `user_request_summary`를 frontmatter에 둘 수 있습니다.
사용자 메시지 원문이 필요한 경우 frontmatter summary로 대체하지 말고 Markdown 본문의 `Flow Contract`에 둡니다.
긴 사용자 원문, 판단 근거, work boundary, non-goals, approval boundary, verification expectation, material judgment calls는 Markdown 본문에 둡니다.
사용자 원문을 저장할 때는 요약/해석과 분리하고, 원문 자체를 재가공, 수정, 번역, 오탈자 정정, 완곡화, 맥락 병합하지 않습니다.

## Continuity Guard 필수 필드

`Continuity Guard` 상태 필드는 frontmatter에 남아야 합니다.

- `Turn-gate active`
- `Question-routing mode`
- `User explicit stop`
- `Terminal summary allowed`
- `Required next action`
- `Last refreshed phase`
- `Confirmed closure`
- `Closure source message`
- `Closure recorded phase`
- `Pending question state`
- `Pending question id or summary`
- `Superseded question id or summary`
- `Verification status`
- `Continuity note`

`Closure source message`와 `Closure recorded phase`는 explicit stop이 source-recorded된 경우에만 terminal close 근거가 됩니다.
source 없는 closure 또는 stale `Terminal summary allowed: yes`를 발견하면 `User explicit stop: no`, `Terminal summary allowed: no`로 복구하고, `Continuity note`에 stale 상태였음을 남깁니다.

## Flow Contract 필수 필드

`Flow Contract`는 다음 필드를 소유합니다.

- user request raw, when preserving the exact message matters
- user request summary or interpretation
- preparation source/result
- boundary rationale
- current blocker
- scope lock status
- work boundary
- non-goals
- acceptance signal
- expected risky actions
- approval boundary
- user-gated checkpoints
- verification expectation
- material judgment calls

## Self-Drive Flow-Local Snapshot

Self-drive 전용 snapshot section은 기본 flow template에 노출하지 않습니다.
Self-drive가 active인 flow record는 `references/self-drive.md`의 runtime guidance를 따라 `Flow Contract`, `Execution Log`, `Report`, 또는 `Next Flow Options` 중 가장 자연스러운 위치에 sequence position, local progress note, next handoff, blocker return condition만 flow-local snapshot으로 남깁니다.
전체 sequence objective, planned flow list, active flow index, allowed/prohibited autonomous actions, approval-sensitive checkpoints, endpoint는 `000-self-drive.md`가 소유합니다.

## Verification Section 필수 필드

`Verification` section은 result status와 method를 분리해서 기록합니다.

- `Status`: `not-started`, `requested`, `pass`, `fail`, `blocked`, `insufficient` 중 하나
- `Method`: `clean-context`, `normal`, `not-required` 중 하나
- `Verifier`: clean-context verifier id, 또는 `not-used`
- `Reason`: method 선택 이유, not-required reason, 또는 blocker reason
- `Checks`: 실행했거나 검토한 command/check/source readback
- `Evidence`: pass/fail 판단에 사용한 핵심 evidence
- `Residual uncertainty`: 남은 불확실성
- `Result`: 최종 통합 결과

`not-required`는 검증 성공 상태가 아니라 method 판단입니다.
검증할 work output이 없거나 별도 검증 동작이 불필요한 이유와 남은 불확실성을 기록해야 합니다.
파일 변경, release surface, 다중 파일 계약, 실패 이력, 사용자 요청 검증, approval-sensitive action에는 `not-required`를 사용하지 않습니다.

## 중복 방지 규칙

- `Analysis`라는 broad bucket을 사용하지 않습니다. scope, routing, approval, verification expectation은 `Flow Contract`로 모읍니다.
- `Plan`과 `Work`를 별도 top-level section으로 분리하지 않고 `Execution Log` 아래에 둡니다.
- `Files changed`와 `Files or surfaces touched`를 동시에 쓰지 않습니다. 최종 표면은 `Changed surfaces`로 기록합니다.
- `Why this mode`, `Why this question-routing mode`, `Next-flow context` 같은 반복 필드는 기본 템플릿에 두지 않습니다. 필요하면 `Material judgment calls`나 `Report`에 한 줄로 남깁니다.
- `Optional Risky Actions` 섹션은 유지합니다. risky action이 없으면 `not-applicable`로 접고 checklist를 펼치지 않습니다.
- risky action이 현재 flow에서 예상되지만 아직 실행 승인이 없으면 checklist를 펼치고 `Initial agreement`를 `not-approved`, `deferred`, 또는 `handoff-required`로 기록합니다.

## Next Flow Options

- `Next Flow Options`는 flow record가 소유합니다.
- visible `request_user_input` 선택지에 turn-end option이 보이지 않아도 record에는 explicit turn-end option을 남깁니다.
- 선택 결과나 active next flow pointer만 `000-plan.md` snapshot에 반영합니다.

## 검토 질문

- flow record가 해당 flow의 canonical detail artifact로 충분한가?
- frontmatter가 explicit stop, stale closure, pending question, verification status를 판별할 수 있는가?
- source 없는 closure나 stale terminal summary allowance를 복구하는 행동이 기록되는가?
- `Flow Contract`가 work 전 scope와 approval boundary를 충분히 잠그는가?
- 사용자 원문이 필요한 flow에서 원문과 요약/해석이 분리되어 있는가?
- self-drive active flow에서 기본 template field를 억지로 채우지 않고, reference guidance에 따라 전체 sequence는 `000-self-drive.md`가 소유하고 flow record는 flow-local snapshot만 남기는가?
- `Execution Log`가 plan/work/evidence를 지나치게 흩뜨리지 않고 담는가?
- report 뒤 `Next Flow Options`가 explicit stop 없는 흐름을 다시 열 수 있는가?
