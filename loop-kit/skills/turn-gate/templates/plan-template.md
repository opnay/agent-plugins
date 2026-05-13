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

Keep this section focused on recent routing context. Older request detail is canonically owned by the relevant `001+` flow records.

1. <user request summary and time/order>

## Flow Index

Use one compact entry per flow. Do not repeat detailed scope, completion criteria, evidence, verification detail, or next-flow options here.

- `.agents/sessions/YYYYMMDD/001-english-lower-slug.md` | operational-preparation | complete | verification: not-applicable | <short outcome>

## Planned Flow Sequence

Keep only current or future selected flows here. Do not retain completed flow detail in this section after the flow completes; move it to `Flow Index` and `Completed Flow Summaries` as compact entries.

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

Keep every completed flow as a one-line link/outcome entry. Do not expand this into detailed reports; the linked `001+` flow record owns detailed scope, evidence, verification, and residual risk.

- `.agents/sessions/YYYYMMDD/001-english-lower-slug.md`: <short outcome>

## Explicit Turn-End Option

- Record only the date-level availability that the user can explicitly stop the turn even when the visible question UI cannot show a stop option. Keep detailed next-flow options in the active flow record.
- Recorded turn-end option:

## Open Risks

Keep only active date-level risks here. Completed or flow-local risks belong in the relevant `001+` flow record.

- <active date-level risk or uncertainty>
