# Safety Boundaries

Use this reference before recommending deletion, shrink, dependency removal, or implementation shortcuts.

## Never Shrink

Do not simplify away:

- validation at trust boundaries
- authentication, authorization, permissions, and privacy checks
- security controls
- accessibility behavior
- error handling that prevents data loss
- user-requested behavior
- public API compatibility
- migrations and rollback paths
- concurrency, async ordering, locks, retries, timeouts, cache invalidation, time, timezone, randomness, hardware, sensors, and floating-point correction space
- observability needed to operate a risky path
- investigation needed to understand the real flow

## Narrow Defense

Use block-lists only as a narrow defense after the valid contract is defined.

Prefer:

- allowed states over disallowed one-off cases
- structured parsing over string guessing
- explicit invalid-state errors over silent fallback
- owner-level validation over duplicated caller checks

## Full Version Request

If the user explicitly asks for the full version, implement the full version without arguing. Still keep the design understandable and avoid speculative layers outside the requested scope.

## Review Guard

Minimal smoke tests, self-checks, and safety assertions are not overengineering when they protect a real contract.
