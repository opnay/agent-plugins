# flow record template 계약

## 소유 범위

`flow`는 active flow record가 하나의 reviewable work unit을 어떻게 복구 가능하게 남기는지 정의합니다.
runtime template은 `skills/flow/templates/flow-record.md`가 제공합니다.
새 active flow boundary가 생기면 새 record가 필요합니다.

## 계약

- 파일명: `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`.
- 생성 규칙: active flow boundary가 바뀔 때 새 파일을 만들고, counter는 session date 안에서 증가시킵니다.
- frontmatter: `phase`, `verification_status`, `next_action`, `flags`, `answered_question`, optional `pending_question`, `continuity`.
- 기본 섹션: `Contract`, `Verification Todo`, `Phase Checklist`, `Execution Log`, `Result`.
- `Contract`: `scope`, `exclude`, `done`, `boundary`, `handoff`.
- `Verification Todo`: `Requirement Verification`, `Implementation Verification`.
- `Requirement Verification`: 사용자 요구사항과 완료 기준을 결과물이 충족해야 할 세부 observable requirement 목록.
- `Implementation Verification`: 타입, 테스트, 빌드, lint, import path, release surface, 코드베이스 관례 같은 구현 정합성 목록.
- `Phase Checklist`: `intake`, `framing`, `preparation`, `work`, `verification`, `reporting`, `next-flow`.
- `Execution Log`: phase start/end, question state, approval checkpoint, edits, build, verification, interruption routing, reporting outcome.
- `Result`: status, evidence, residual risk, next action 또는 handoff.

## 조건부 섹션

- `Risky Action`: approval-sensitive action이 있을 때만 추가합니다.
- raw request text: 정확한 원문이 해석에 영향을 줄 때만 summary와 분리해 기록합니다.
- non-pass routing: `fail`, `blocked`, `insufficient`일 때 필요한 만큼 둡니다.

## 검토 기준

- checklist가 phase 시작이 아니라 종료 checkpoint 통과를 뜻하는가?
- `interruption`이 checklist가 아니라 log event로 남는가?
- requirement verification이 행위 수행 여부가 아니라 사용자 요구사항 충족 여부를 검증하는가?
- implementation verification이 requirement verification을 대체하지 않는가?
- `verification_status`와 `Result`가 success evidence를 과장하지 않는가?
- boundary가 commit, push, PR, publish, release, version bump, destructive/external action의 승인 상태를 드러내는가?
