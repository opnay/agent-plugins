---
task: "<short task statement>"
flow_type: "operational-preparation | change-unit"
flow_scope: "<what this flow owns>"
current_phase: "preparation | work | verification | reporting | next-flow"
turn_gate_active: "yes | no"
question_routing_mode: "user-gated | pending-question | blocked | undecided"
user_explicit_stop: "yes | no"
terminal_summary_allowed: "yes | no"
required_next_action: ""
last_refreshed_phase: "preparation | work | verification | reporting | next-flow"
confirmed_closure: "yes | no"
closure_source_message: ""
closure_recorded_phase: ""
pending_question_state: "none | pending | answered | aborted | superseded"
pending_question_id_or_summary: ""
superseded_question_id_or_summary: ""
verification_status: "not-started | requested | pass | fail | blocked | insufficient | not-applicable"
continuity_note: ""
preparation_source: "user-message | existing-flow | correction | next-flow"
scope_lock_status: "locked-by-question | inferred | not-needed | pending"
user_request_summary: ""
---

# Turn-Gate Flow Report

This file is the canonical detail record for one user-request-driven flow. Update it incrementally after each completed phase.

Keep date-level history, active snapshot, planned flow sequence, and flow index in `000-plan.md`. Keep detailed flow contract, evidence, verification, report, `Continuity Guard`, next-flow options, and residual risk here.

One flow is a cohesive reviewable or commit-sized change unit. It is not an analysis/work/verification/commit phase slice, and it does not need to be direct user-visible value. User-message intake and planned-flow design can be an `operational-preparation` flow when they own session plan, flow list, scope, or approval-boundary artifacts.

## Flow Contract

- User request:
- Preparation source: user-message | existing-flow | correction | next-flow
- Preparation result:
- Boundary rationale:
- Current blocker:
- Scope lock status: locked-by-question | inferred | not-needed | pending
- Work boundary:
- Non-goals:
- Acceptance signal:
- Expected risky actions:
- Approval boundary:
- User-gated checkpoints:
- Verification expectation:
- Material judgment calls:

## Optional Risky Actions

not-applicable

If a risky action is possible but not approved, replace `not-applicable` with:

1. <destructive | irreversible | external | commit | push | PR | publish | release | version-bump | other>
   - Exact target:
   - Expected effect:
   - Risk:
   - Rollback or recovery:
   - Included scope:
   - Excluded scope:
   - Initial agreement: approved | not-approved | deferred | handoff-required | not-applicable
   - Question-routing handling: covered-by-initial-agreement | return-to-user-gated-question-routing | handoff-required

## Execution Log

- Plan:
  1. <first action>
  2. <next action>
- Work:
  - Change or action:
  - Changed surfaces:
- Evidence:
  - Command or output summary:

## Verification

- Verification packet or request id:
- What was checked:
- Result: not-started | requested | pass | fail | blocked | insufficient | not-applicable
- Remaining uncertainty:
- Contrary evidence or multi-angle critique:

## Report

- Outcome:
- Commit-readiness report:
- Commit execution approval: not-requested | requested-separately | not-applicable
- Blocker, if any:

## Next Flow Options

1. <next flow choice>
2. <alternative>
3. <defer or alternate path>
4. Turn end option: <record explicit user-stop option even when not shown in the visible question UI>

## Residual Risk

- <risk or uncertainty>
