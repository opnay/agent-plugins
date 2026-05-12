# Commit Readiness Gate Reference

Use this phase protocol when the intended change unit is largely complete and the current phase is to decide whether it is ready to move toward commit.

## Core Contract

- Evaluate the intended change unit, not the whole repository.
- Confirm the work boundary, unrelated-change exclusion, verification status, residual risk, and likely commit-message scope.
- Run final review and scoped verification appropriate to a readiness call.
- Report readiness, residual risk, intended diff scope, unrelated changes to exclude, verification evidence, and any minimum review recommendation together.
- Keep readiness judgment separate from execution authority. A readiness request provides readiness evidence only.
- Commit-readiness reporting happens as reporting or handoff state after the intended change-unit work. Create a new planned flow only when the readiness work owns a distinct artifact change.
- Stage, commit, push, PR creation, publish, release, and version-bump execution require recorded execution authority. Use approval-boundary and user-gated routing to decide whether to proceed or reopen a question.

## Protocol Boundary

- Good fit: final readiness judgment after implementation and targeted verification.
- Not a fit: broad implementation, unclear intended change unit, readiness without enough verification evidence, or execution steps that lack recorded execution authority.

## Handoff

- If ready, report the execution handoff options allowed by the approval boundary.
- If not ready, return to the earliest safe phase for corrective work or user-gated clarification.

## Review Questions

- Is the intended change unit locked enough to judge?
- Is scoped verification sufficient for the readiness call?
- Did the report distinguish commit readiness from execution authority?
- If execution follows, is execution authority recorded?
