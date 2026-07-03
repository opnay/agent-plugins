# Refactor Shrink

Use this reference for refactoring, cleanup, simplification, or behavior-preserving code shrink tasks.

## Contract

Refactor only after the behavior contract is known. Shrink concepts, not evidence.

## Safe Shrink Targets

- duplicate branches with the same effect
- wrapper functions that only rename another call
- config with one real value and no near-term owner
- interfaces with one implementation and no test seam value
- factories that never select between implementations
- hooks or callbacks that only forward values
- custom loops replaceable by clear library primitives
- local state derived from existing state without caching need

## Unsafe Shrink Targets

Do not shrink:

- validation at trust boundaries
- explicit error handling that prevents data loss
- accessibility behavior
- security checks
- public API compatibility
- concurrency, time, retry, timeout, cache, file, migration, permission, or auth safeguards
- user-requested behavior

## Sequence

1. Capture current behavior and checks.
2. Remove one concept at a time.
3. Keep names and control flow obvious.
4. Run the check after each risky step or at the smallest useful batch.
5. Stop when further shrink would hide intent or weaken safety.

## Output

Report:

- behavior kept
- concepts removed
- replacement
- verification
- rollback trigger
