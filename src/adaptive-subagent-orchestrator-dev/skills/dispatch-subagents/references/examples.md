# Dispatch Examples

## PARALLEL_READ

Request: review a branch for security, correctness, test gaps, and performance.

- Use distinct read-only lenses only when each produces different evidence.
- Keep final severity, deduplication, and fixes main-owned.

Request: investigate intermittent failures through auth, payment, and inventory.

- Assign independent execution paths or modules.
- Integrate causes before authorizing writes.

## DIRECT

Request: fix lint errors across the repository.

- Use `DIRECT` when one shared lint config or root type error explains the failures.
- Do not dispatch because many files are affected.

Request: change one function and then update its dependent caller.

- Use `DIRECT` because the steps are strongly sequential.

## Read Before Write

Request: unit, integration, and end-to-end suites fail.

- Investigate test bundles with `PARALLEL_READ`.
- Integrate causes.
- Re-dispatch write work only when fixes and writable files are disjoint.

## Safe PARALLEL_WRITE

- Worker A owns `packages/a/src/**` and `packages/a/test/**`.
- Worker B owns `packages/b/src/**` and `packages/b/test/**`.
- Main owns the lockfile, root config, shared schema, and integration tests.

## Unsafe PARALLEL_WRITE

Do not let different workers edit a shared API schema, lockfile, generated routes, common fixture, or unconfirmed cross-layer contract. Use `PARALLEL_READ` or one writer.
