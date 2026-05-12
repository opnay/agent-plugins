# turn-gate plan template sub-spec

## 목적

이 문서는 `turn-gate`의 `skills/turn-gate/templates/plan-template.md` 형식 계약을 소유합니다.

`000-plan.md`는 날짜 단위 session artifact입니다. 각 flow의 상세 증거 기록이 아니라, 당일 흐름을 이어가기 위한 index와 snapshot입니다.

## 소유

- date-level summary와 latest user request/decision
- active flow pointer와 required next action
- user request history
- flow index
- planned flow sequence
- completed flow summaries
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
2. `User Requests Today`
3. `Flow Index`
4. `Planned Flow Sequence`
5. `Completed Flow Summaries`
6. `Explicit Turn-End Option`
7. `Open Risks`

## Frontmatter 규격

`plan-template.md`는 Markdown YAML frontmatter로 date-level snapshot과 routing 상태를 먼저 드러냅니다.
파일 경로가 `.agents/sessions/{YYYYMMDD}/000-plan.md`이므로 `date`와 `record_path`는 frontmatter에 반복하지 않습니다.

frontmatter에는 다음 필드를 둡니다.

- `summary`
- `latest_user_request`
- `latest_decision`
- `active_flow`
- `required_next_action`
- `pending_question_state`
- `verification_status`
- `preparation_source`
- `scope_lock_status`
- `final_readiness_handoff`

`preparation_result`, `flow_list_basis`, `flow_type_rule`, `flow_boundary_basis`처럼 길어지기 쉬운 판단 근거는 `Planned Flow Sequence` 본문에 둡니다.

## 중복 방지 규칙

- `000-plan.md`에 복사되는 값은 snapshot이어야 하며 canonical detail이 아니어야 합니다.
- `Work boundary`, `Non-goals`, `Approval boundary`, `Verification expectation`, `Evidence`의 상세값은 flow record가 소유합니다.
- plan의 `Verification status`는 상태값만 기록하고, 검사 목록과 근거는 flow record의 `Verification`에 둡니다.
- plan의 `Required next action`은 active turn routing을 위한 짧은 pointer이며, 상세 blocker/risk 설명을 반복하지 않습니다.
- plan의 `Explicit Turn-End Option`은 전체 next-flow options가 아니라 사용자가 명시적으로 turn을 끝낼 수 있다는 date-level availability snapshot입니다.
- 실행이 아니라 판단, 설계, 범위 확인만 수행한 경우에는 `Planned Flow Sequence` 안에 `Follow-up Change-Unit Candidates`로 라벨링된 후보 목록을 둘 수 있습니다. 이 후보는 선택 또는 승인 전까지 active/completed flow로 세지 않습니다.
- completed flow는 삭제하지 않고 one-line summary와 flow record link로 유지합니다.

## 검토 질문

- plan이 date-level index와 active snapshot으로 읽히는가?
- plan이 flow record의 상세 scope/evidence/verification을 반복하지 않는가?
- planned flow sequence가 phase checklist가 아니라 cohesive flow list인가?
- 후속 후보가 planned/active/completed flow와 구분되어 라벨링되는가?
- active flow pointer와 required next action이 다음 진행에 충분한가?
- completed flow summaries가 상세 보고가 아니라 링크 가능한 요약인가?
