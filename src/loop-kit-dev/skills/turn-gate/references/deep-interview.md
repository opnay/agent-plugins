# Deep Interview Reference

Use this mode when the current phase is requirement discovery rather than implementation, refinement, review handling, or readiness judgment.

## Core Contract

- Ask questions that lock intent, scope, non-goals, tradeoffs, approval boundary, completion criteria, acceptance signal, and verification expectation.
- For user-message-driven preparation, collect enough information for the whole planned flow list, not only the first flow.
- Ask about expected risky actions up front when planned flows may involve destructive, irreversible, external, commit, push, PR, publish, or similar sensitive actions.
- Record expected risky actions as approved, not approved, deferred, or handoff-required, including target, expected effect, risk, rollback or recovery possibility, and included/excluded scope when known.
- Apply the scope floor before work: if scope is missing, too broad, can produce multiple valid outputs, or can change success criteria or verification, ask a user-gated scope-lock question.
- Scope-lock questions should directly cover the flow-changing subset of included scope, excluded scope, target files/surfaces/artifacts, completion criteria, and verification signal.
- Prefer `request_user_input` when bounded choices can lock the question efficiently.
- If scope is safely inferred, record the inferred work boundary and non-goals before handoff.
- Silent inference is not approval for destructive, irreversible, external, commit, push, PR, publish, or similar sensitive work.
- Convert the discovery result into a planned flow list with flow boundaries, non-goals, acceptance signal, verification expectations, expected risky actions, and user-gated checkpoints.
- If autonomous continuation is appropriate, prepare the self-drive handoff boundary: which flows can proceed under the initial agreement and which must return to user-gated question-routing.
- Do not stop at advisory guidance when an actual question round is still needed.
- Once clarity is sufficient, hand off to the next fitting mode, `self-drive`, or downstream workflow with the work boundary and non-goals intact.

## Mode Boundary

- Good fit: unclear requirements, direction checks, proposal shaping, greenfield setup alignment, missing scope boundaries, unresolved approval lines, ambiguous success criteria, or missing planned-flow acceptance signals.
- Not a fit: generic operation/target ambiguity cleanup without real requirement discovery, read-only planning, implementation itself, approval execution, or final commit execution.

## Review Questions

- Is the current bottleneck true requirement discovery?
- Would different reasonable scopes produce different outputs or verification paths?
- Could this question be locked with bounded choices instead of freeform text?
- Did initial preparation identify expected risky actions and planned approval checkpoints?
- Did the flow record capture inferred boundaries and non-goals when no question was asked?
- Is the downstream handoff clear enough after this discovery round?
