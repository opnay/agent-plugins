# Turn-Gate Flow Report Template

This file is the detailed report for one user-request-driven flow.
Update it incrementally after each completed phase. Do not wait until the end of the flow.
Keep the date-level history and flow index in `000-plan.md`; keep the detailed evidence and phase report here.

## Metadata

- Date: YYYY-MM-DD
- Record path: .agents/sessions/YYYYMMDD/001-english-lower-slug.md
- Slug: english-lower-slug
- User request message: <original user message for this flow>
- Task: <short task statement>
- Flow scope: <what this flow owns>
- Parent plan: .agents/sessions/YYYYMMDD/000-plan.md
- Current mode: deep-interview | autopilot | review-loop | ralph-loop | commit-readiness-gate | undecided
- Question-routing mode: user-gated | self-drive-handoff | undecided
- Current core phase: preparation | work | verification | reporting

## Continuity Guard

- Turn-gate active: yes | no
- Question-routing mode: user-gated | self-drive-handoff | undecided
- User explicit stop: yes | no
- Terminal summary allowed: yes | no
- Required next action:
- Last refreshed phase: preparation | work | verification | reporting
- Confirmed closure: yes | no
- Closure source message: <required only when confirmed closure is recorded>
- Closure recorded phase: <required only when confirmed closure is recorded>
- Pending question state: none | pending | answered | aborted | superseded
- Pending question id or summary:
- Superseded question id or summary:
- Verification status: not-started | requested | pass | fail | blocked | insufficient | not-applicable
- Continuity note:

## Analysis

- User request:
- Preparation source: user-message | existing-flow | correction | next-flow
- Preparation result:
- Planned flow list:
- Current blocker:
- Why this flow:
- Why this mode:
- Why this question-routing mode:
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

## Expected Risky Actions

1. <destructive | irreversible | external | commit | push | PR | publish | other>
   - Exact target:
   - Expected effect:
   - Risk:
   - Rollback or recovery:
   - Included scope:
   - Excluded scope:
   - Initial agreement: approved | not-approved | deferred | handoff-required | not-applicable
   - Self-drive handling: covered-by-initial-agreement | return-to-user-gated-question-routing

## Plan

1. <first action>
2. <next action>
3. <next action>

## Work

- Change or action:
- Files or surfaces touched:
- Evidence or command output summary:

## Verification

- Verification packet or request id:
- What was checked:
- Result: not-started | requested | pass | fail | blocked | insufficient | not-applicable
- Remaining uncertainty:
- Contrary evidence or multi-angle critique:

## Result Report

- Outcome:
- Files changed:
- Commit-readiness report:
- Commit execution approval: not-requested | requested-separately | not-applicable
- Next-flow context:
- Blocker, if any:

## Next Flow Options

1. <next flow choice>
2. <alternative>
3. <defer or alternate path>
4. Turn end option: <record explicit user-stop option even when not shown in the visible question UI>

## Residual Risk

- <risk or uncertainty>
