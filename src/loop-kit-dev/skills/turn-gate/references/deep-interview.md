# Deep Interview Reference

Use this mode when the current phase is requirement discovery rather than implementation, refinement, review handling, or readiness judgment.

## Core Contract

- Ask questions that lock intent, scope, non-goals, tradeoffs, approval boundary, completion criteria, and verification signal.
- Apply the scope floor before work: if scope is missing, too broad, can produce multiple valid outputs, or can change success criteria or verification, ask a user-gated scope-lock question.
- Scope-lock questions should directly cover the flow-changing subset of included scope, excluded scope, target files/surfaces/artifacts, completion criteria, and verification signal.
- Prefer `request_user_input` when bounded choices can lock the question efficiently.
- If scope is safely inferred, record the inferred work boundary and non-goals before handoff.
- Silent inference is not approval for destructive, irreversible, external, commit, push, PR, publish, or similar sensitive work.
- Do not stop at advisory guidance when an actual question round is still needed.
- Once clarity is sufficient, hand off to the next fitting mode or downstream workflow with the work boundary and non-goals intact.

## Mode Boundary

- Good fit: unclear requirements, direction checks, proposal shaping, greenfield setup alignment, missing scope boundaries, unresolved approval lines, ambiguous success criteria.
- Not a fit: generic operation/target ambiguity cleanup without real requirement discovery, read-only planning, implementation itself, or approval execution.

## Review Questions

- Is the current bottleneck true requirement discovery?
- Would different reasonable scopes produce different outputs or verification paths?
- Could this question be locked with bounded choices instead of freeform text?
- Did the flow record capture inferred boundaries and non-goals when no question was asked?
- Is the downstream handoff clear enough after this discovery round?
