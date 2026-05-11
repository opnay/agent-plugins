# Turn-Gate Flow Report Template

This file is the detailed report for one user-request-driven flow.
Update it incrementally after each completed phase. Do not wait until the end of the flow.
Keep the date-level history, active snapshot, planned flow sequence, and flow index in `000-plan.md`; keep the detailed flow contract, evidence, verification, and report here.
One flow is a cohesive reviewable or commit-sized change unit. It is not an analysis/work/verification/commit phase slice, and it does not need to be direct user-visible value.
사용자 메시지 intake와 planned-flow design은 session plan, flow-list, scope, approval-boundary artifact를 소유할 때 operational-preparation flow가 될 수 있습니다.
operational-preparation flow는 code, docs, fixtures, config, release-surface 변경을 소유하는 change-unit flow와 구분해야 합니다.
Final QA, consistency checking, verification-result reporting, and commit-readiness reporting belong in this flow's verification/reporting sections unless they create or change a distinct reviewable artifact/change unit.
Do not duplicate detailed flow contract fields in both `000-plan.md` and this flow record. If a value is copied into `000-plan.md`, keep it as a short snapshot and treat this flow record as canonical.

## Metadata

- Date: YYYY-MM-DD
- Record path: .agents/sessions/YYYYMMDD/001-english-lower-slug.md
- Slug: english-lower-slug
- User request message: <original user message for this flow>
- Task: <short task statement>
- Flow type: operational-preparation | change-unit
- Flow scope: <what this flow owns>
- Parent plan: .agents/sessions/YYYYMMDD/000-plan.md
- Current phase: preparation | work | verification | reporting | next-flow

## Continuity Guard

- Turn-gate active: yes | no
- Question-routing mode: user-gated | self-drive-handoff | undecided
- User explicit stop: yes | no
- Terminal summary allowed: yes | no
- Required next action:
- Last refreshed phase: preparation | work | verification | reporting | next-flow
- Confirmed closure: yes | no
- Closure source message: <required only when confirmed closure is recorded>
- Closure recorded phase: <required only when confirmed closure is recorded>
- Pending question state: none | pending | answered | aborted | superseded
- Pending question id or summary:
- Superseded question id or summary:
- Verification status: not-started | requested | pass | fail | blocked | insufficient | not-applicable
- Continuity note:

If closure is source-less or stale, reset `User explicit stop: no` and `Terminal summary allowed: no`, then note the stale state in `Continuity note`.

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
- Self-drive eligibility:
- Verification expectation:
- Material judgment calls:

## Optional Risky Actions

1. <destructive | irreversible | external | commit | push | PR | publish | other>
   - Exact target:
   - Expected effect:
   - Risk:
   - Rollback or recovery:
   - Included scope:
   - Excluded scope:
   - Initial agreement: approved | not-approved | deferred | handoff-required | not-applicable
   - Self-drive handling: covered-by-initial-agreement | return-to-user-gated-question-routing

Keep this section. If no risky action applies, write `not-applicable` and do not expand the checklist. If a risky action is possible but not approved, expand the checklist and record `Initial agreement` as `not-approved`, `deferred`, or `handoff-required`.

## Execution Log

- Plan:
  1. <first action>
  2. <next action>
  3. <next action>
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
