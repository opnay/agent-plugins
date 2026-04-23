# Deep Interview Reference

Use this mode when the current phase is requirement discovery rather than implementation, refinement, review handling, or readiness judgment.

## Core Contract

- Ask questions that lock intent, scope, non-goal, tradeoff, approval boundary, and acceptance signal.
- Prefer `request_user_input` when bounded choices can lock the question efficiently.
- Do not stop at advisory guidance when an actual question round is still needed.
- Once clarity is sufficient, hand off to the next fitting mode or downstream workflow.

## Mode Boundary

- Good fit: unclear requirements, direction checks, proposal shaping, greenfield setup alignment, ambiguous success criteria.
- Not a fit: generic ambiguity cleanup without real requirement discovery, read-only planning, or implementation itself.

## Review Questions

- Is the current bottleneck true requirement discovery?
- Could this question be locked with bounded choices instead of freeform text?
- Is the downstream handoff clear enough after this discovery round?
