# Commit Readiness Gate Reference

Use this phase protocol when implementation is largely done and the current phase is to decide whether the intended change unit is ready to move toward commit.

## Core Contract

- Evaluate the intended change unit, not the whole repository.
- Confirm the work boundary, unrelated-change exclusion, verification status, residual risk, and likely commit-message scope.
- Run final review and scoped verification appropriate to a readiness call.
- Report readiness, residual risk, intended diff scope, unrelated changes to exclude, verification evidence, and any minimum review recommendation together.
- Keep readiness judgment separate from execution authority. A readiness request provides readiness evidence only.
- Commit-readiness reporting happens as reporting or handoff state after the intended change-unit work. Create a new planned flow only when the readiness work owns a distinct artifact change.
- Stage, commit, push, PR creation, and publish execution require recorded execution authority. Use approval-boundary and user-gated routing to decide whether to proceed or reopen a question.

## Protocol Boundary

- Good fit: final readiness judgment after implementation and targeted verification.
- Use another protocol for broad implementation, planning, or execution steps that lack recorded execution authority.

## Review Questions

- Is the intended change unit actually locked enough to judge?
- Is the scoped verification sufficient for the readiness call?
- Did the report distinguish commit readiness from execution authority?
- If execution follows, is the execution authority recorded?
