---
summary: ""
latest_user_request: ""
latest_decision: ""
active_flow: ""
required_next_action: ""
pending_question_state: "none | pending | answered | aborted | superseded"
verification_status: "not-started | requested | pass | fail | blocked | insufficient | not-applicable"
preparation_source: "user-message | existing-flow | correction | next-flow"
scope_lock_status: "locked-by-question | inferred | not-needed | pending"
final_readiness_handoff: "commit-readiness reporting | other | not-applicable"
---

# Turn-Gate Multi-Flow Plan

This file is the date-scoped multi-flow plan for turn-gated work. Use it as an index and active snapshot, not as the canonical detail record for each flow.

Planned flows are cohesive reviewable or commit-sized change units, not phase checklists. A flow does not need to be direct user-visible value. User-message intake and planned-flow design can be an `operational-preparation` flow when they own session plan, flow list, scope, or approval-boundary artifacts.

Do not duplicate each flow's detailed scope, non-goals, approval boundary, evidence, or verification detail here. Keep those details in the `001+` flow record.

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
- Final readiness handoff: commit-readiness reporting | other | not-applicable

1. `<flow slug or title>`
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
