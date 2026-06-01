# plan template 계약

## 소유 범위

`flow`는 plan record가 flow routing card로 어떤 정보를 담아야 하는지 정의합니다.
runtime template은 `skills/flow/templates/plan.md`가 제공합니다.
`turn-gate`는 active turn에서 필요한 경우 이 계약을 적용하고 실제 파일 갱신을 라우팅할 수 있습니다.

## 계약

- 파일명: `.agents/sessions/{YYYYMMDD}/000-plan.md`.
- 생성 규칙: session date마다 하나만 둡니다.
- frontmatter: `flow_plan_active`, `active_flow`, `next_action`, `handoff_condition`, `approval_boundary`, `verification_expectation`, `active_skills`.
- body: current request 또는 routing signal, purpose, active/planned/recent/archive flow index, continuity note.
- 목적 사슬은 scope, acceptance, verification, approval, handoff에 영향을 줄 때만 compact purpose로 둡니다.
- active skills는 다음 flow 복구에 필요한 skill만 기록합니다.

## 제외

- 완료된 flow 전체 요약
- full conversation history
- detailed verification evidence
- Git이나 tool readback으로 복구 가능한 상태
- turn-level closure state나 self-drive sequence detail
- commit, push, PR, release, version bump, destructive action authority

## 검토 기준

- 현재 active flow와 next action을 compaction 뒤 복구할 수 있는가?
- 기록이 routing card 크기를 넘지 않는가?
- 승인 민감 작업이 readiness나 verification에서 암묵 승인되지 않는가?
