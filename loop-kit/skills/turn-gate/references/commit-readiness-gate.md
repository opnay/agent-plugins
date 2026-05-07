# Commit Readiness Gate Reference

Use this mode when implementation is largely done and the current phase is to decide whether the intended change unit is ready to move toward commit.

## Core Contract

- Evaluate the intended change unit, not the whole repository.
- Confirm the work boundary, unrelated-change exclusion, verification status, residual risk, and likely commit-message scope.
- Run final review and scoped verification appropriate to a readiness call.
- Report readiness, residual risk, and any minimum review recommendation together.
- Keep commit execution itself outside this mode. A readiness request is not commit approval.

## Mode Boundary

- Good fit: final readiness judgment after implementation and targeted verification.
- Not a fit: broad implementation, planning, commit execution, push, PR, publish, or commit-message finalization.

## Review Questions

- Is the intended change unit actually locked enough to judge?
- Is the scoped verification sufficient for the readiness call?
- Did this mode stop at the gate instead of drifting into commit execution?
