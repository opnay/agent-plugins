# turn-gate plan template sub-spec

## 목적

이 문서는 `turn-gate`의 `skills/turn-gate/templates/plan-template.md` 형식 계약을 소유합니다.

`000-plan.md`는 날짜 단위 session artifact입니다. 각 flow의 상세 증거 기록이 아니라, 당일 흐름을 이어가기 위한 bounded index와 snapshot입니다.

## 소유

- date-level summary와 latest user request/decision
- active flow pointer와 required next action
- current recovery state
- compact flow table
- planned/current flow sequence only when needed
- open risks
- visible UI에 turn-end option이 없더라도 날짜 단위 availability snapshot으로 남기는 explicit turn-end option

## 비소유

- 각 flow의 상세 scope, non-goals, approval boundary
- command output, evidence, verification detail
- per-flow residual risk의 상세 근거
- `Continuity Guard`의 canonical state
- flow-local work log
- `Next Flow Options`의 상세 후보 목록과 turn-end option 문구

상세 flow contract와 evidence는 `templates/flow.md`가 소유하는 `001+` flow record에 둡니다.

## 필수 구조

`plan-template.md`는 다음 구조를 유지합니다.

1. YAML frontmatter
2. `Current State`
3. `Flow Table`
4. `Open Risks`
5. `Turn-End Rule`

## Frontmatter 규격

`plan-template.md`는 Markdown YAML frontmatter로 date-level snapshot과 routing 상태를 먼저 드러냅니다.
파일 경로가 `.agents/sessions/{YYYYMMDD}/000-plan.md`이므로 `date`와 `record_path`는 frontmatter에 반복하지 않습니다.

frontmatter에는 다음 필드를 둡니다.

- `summary`
- `turn_gate_session_id`
- `latest_user_request`
- `latest_decision`
- `active_flow`
- `required_next_action`
- `pending_question_state`
- `verification_status`
- `self_drive_status`
- `self_drive_record`
- `preparation_source`
- `scope_lock_status`
- `final_readiness_handoff`

`turn_gate_session_id`는 current main Codex session id가 known일 때만 기록합니다.
이 값은 plugin-installed Stop hook이 다른 session이나 subagent의 Stop event를 막지 않도록 하는 arming marker입니다.
unknown이면 빈 문자열로 두고, Stop hook은 fail-open으로 조용히 종료해야 합니다.

`preparation_result`, `flow_list_basis`, `flow_type_rule`, `flow_boundary_basis`처럼 길어지기 쉬운 판단 근거는 `Planned Flow Sequence` 본문에 둡니다.

## 중복 방지 규칙

- `000-plan.md`에 복사되는 값은 snapshot이어야 하며 canonical detail이 아니어야 합니다.
- `Work boundary`, `Non-goals`, `Approval boundary`, `Verification expectation`, `Evidence`의 상세값은 flow record가 소유합니다.
- plan의 `Verification status`는 상태값만 기록하고, 검사 목록과 근거는 flow record의 `Verification`에 둡니다.
- plan의 `Required next action`은 active turn routing을 위한 짧은 pointer이며, 상세 blocker/risk 설명을 반복하지 않습니다.
- plan의 `Explicit Turn-End Option`은 전체 next-flow options가 아니라 사용자가 명시적으로 turn을 끝낼 수 있다는 date-level availability snapshot입니다.
- 실행이 아니라 판단, 설계, 범위 확인만 수행한 경우에는 `Planned Flow Sequence` 안에 `Follow-up Change-Unit Candidates`로 라벨링된 후보 목록을 둘 수 있습니다. 이 후보는 선택 또는 승인 전까지 active/completed flow로 세지 않습니다.
- self-drive 전용 sequence shape는 기본 plan template에 노출하지 않습니다. `000-plan.md`는 self-drive-specific fields로 `self_drive_status`와 `self_drive_record` pointer만 유지하고, self-drive sequence state는 optional `000-self-drive.md`가 소유합니다. 일반 `Planned Flow Sequence` section은 date-level routing snapshot일 수 있지만 self-drive sequence-level state의 canonical record가 아닙니다.
- `Current State`는 active flow, current task, work boundary, latest material outcome, and current required next action을 짧게 담습니다.
- `Flow Table`은 flow당 한 줄 compact entry로 유지합니다. 각 entry는 flow number/path and current recovery value만 포함합니다.
- `Planned Flow Sequence`가 필요한 경우 current/future selected flow만 담고, completed flow의 stale plan detail을 남기지 않습니다.
- 오래된 `User Requests Today` chronology와 `Completed Flow Summaries` 중복은 기본 plan template에 두지 않습니다. 최근 요청은 frontmatter snapshot이나 active flow record로 충분해야 합니다.
- `Open Risks`는 active date-level risk만 담고, completed 또는 flow-local risk는 해당 flow record에 둡니다.
- `000-plan.md`는 active context로 자주 읽히므로 강하게 capping합니다. detailed history, verbose request chronology, stale next-flow options, per-flow evidence, and duplicate completed summaries는 flow record로 위임합니다.

## 검토 질문

- plan이 date-level index와 active snapshot으로 읽히는가?
- plan이 flow record의 상세 scope/evidence/verification을 반복하지 않는가?
- flow table이 flow당 한 줄 compact recovery entry로 유지되는가?
- planned flow sequence가 phase checklist가 아니라 cohesive flow list인가?
- planned flow sequence에 completed flow의 stale detail이 남지 않는가?
- 후속 후보가 planned/active/completed flow와 구분되어 라벨링되는가?
- self-drive가 active일 때 `000-plan.md`에는 active status와 sidecar pointer만 있고, sequence-level state는 `000-self-drive.md`로 분리됐는가?
- active flow pointer와 required next action이 다음 진행에 충분한가?
- `turn_gate_session_id`가 known session에서만 기록되고 unknown일 때 빈 값으로 남아 Stop hook이 fail-open 할 수 있는가?
- 오래된 user request chronology나 completed summaries 중복이 active context를 오염시키지 않는가?
