# Turn-Gate Multi-Flow Plan Template

This file is the date-scoped multi-flow plan for turn-gated work.
Use it to show how the user's request decomposes into planned flows, why each flow exists, and when the work should move to the next flow.
Planned flows are cohesive reviewable or commit-sized change units, not phase checklists.
A flow does not need to be direct user-visible value; supporting component, logic, or integration work can be a flow when it is a coherent change unit.
사용자 메시지 intake와 planned-flow design은 session plan, flow-list, scope, approval-boundary artifact를 소유할 때 operational-preparation flow가 될 수 있습니다.
operational-preparation flow는 code, docs, fixtures, config, release-surface 변경을 소유하는 change-unit flow와 구분합니다.
Final QA, consistency checking, verification-result reporting, and commit-readiness reporting are not planned flows unless they create or change a distinct reviewable artifact/change unit.
Keep it incremental. Do not delete completed work; summarize completed flows and keep their links.
Do not duplicate each flow's detailed scope, non-goals, approval boundary, evidence, or verification detail here. Keep those details in the `001+` flow record and keep this plan as a date-level index and snapshot.

## Date

- YYYY-MM-DD

## Daily Snapshot

- Summary:
- Latest user request:
- Latest decision:
- Active flow:
- Required next action:
- Pending question state: none | pending | answered | aborted | superseded
- Verification status: not-started | requested | pass | fail | blocked | insufficient | not-applicable

## User Requests Today

1. <user request summary and time/order>

## Flow Index

1. `.agents/sessions/YYYYMMDD/001-english-lower-slug.md`
   - User request:
   - Flow type: operational-preparation | change-unit
   - Flow purpose:
   - Status: planned | active | complete | blocked
   - Current phase: preparation | work | verification | reporting | next-flow
   - Completion criteria:
   - Next-flow trigger:
   - Verification status: not-started | requested | pass | fail | blocked | insufficient | not-applicable
   - Summary:

## Planned Flow Sequence

- Preparation source: user-message | existing-flow | correction | next-flow
- Preparation result:
- Flow-list basis:
- Flow-type rule: operational-preparation owns session/plan artifacts; change-unit owns reviewable code/doc/fixture/config/release changes
- Flow-boundary basis: cohesive reviewable or commit-sized change unit; not phase checklist; direct user-visible value not required; pure final QA/readiness/reporting is not a flow without a distinct artifact/change unit
- Scope lock status: locked-by-question | inferred | not-needed | pending
- Self-drive eligibility: eligible | partial | not-eligible | undecided
- Final readiness handoff: commit-readiness reporting | other | not-applicable

1. `<flow slug or title>`
   - Flow type: operational-preparation | change-unit
   - Purpose:
   - Why this flow boundary:
   - Owns:
   - Core phase coverage: preparation | work | verification | reporting | next-flow
   - Completion criteria:
   - Next-flow trigger:
2. `<flow slug or title>`
   - Flow type: operational-preparation | change-unit
   - Purpose:
   - Why this flow boundary:
   - Owns:
   - Core phase coverage: preparation | work | verification | reporting | next-flow
   - Completion criteria:
   - Next-flow trigger:
3. `<flow slug or title>`
   - Flow type: operational-preparation | change-unit
   - Purpose:
   - Why this flow boundary:
   - Owns:
   - Core phase coverage: preparation | work | verification | reporting | next-flow
   - Completion criteria:
   - Next-flow trigger:

### Follow-up Change-Unit Candidates

Use this subsection only when the current flow produced possible future implementation units that are not selected or approved yet. These candidates are not active or completed flows.

1. `<candidate title>`
   - Candidate type: change-unit
   - Expected artifact:
   - Separation or compression rationale:
   - Expected verification:
   - User-gated handoff condition:

## Completed Flow Summaries

- `<flow filename>`: <short retained summary>

## Explicit Turn-End Option

- Record only the date-level availability that the user can explicitly stop the turn even when the visible question UI cannot show a stop option. Keep detailed next-flow options in the active flow record.
- Recorded turn-end option:

## Open Risks

- <risk or uncertainty>
