# Code Quality Principles

Use these principles as tradeoff guides, not absolute rules.

## Correctness

Code must satisfy the requested behavior and preserve required existing behavior. Make assumptions visible through code, types, validation, or tests. Do not swallow failures or report partial success as complete success. Check whether partial writes, retries, or interrupted work can leave inconsistent data.

## Clarity

Names should communicate role and intent, not merely implementation. Make units, time zones, currencies, identifier types, and boolean meanings explicit when confusion is plausible. Avoid generic names such as `helper`, `manager`, `processor`, or `util` when they hide responsibility.

## Simplicity

Choose the smallest coherent design, not necessarily the fewest lines. Avoid abstractions, wrappers, or layers created only for a speculative future. Keep straightforward logic local when extracting it would hide the behavior.

## Cohesion and Coupling

Keep code that changes for the same reason close together. Separate responsibilities that change independently. Reuse and shared modules are useful only when they reduce real maintenance cost without binding unrelated concepts together.

## Abstraction

Introduce an abstraction when callers share the same concept, contract, and reason to change. Do not generalize one use case just because a second future use case might appear. A thin wrapper around a standard or framework API needs a real policy, domain, observability, error-model, or platform-separation responsibility.

## Reuse

Good code does not implement everything from scratch. Inspect repository modules, framework and system APIs, standard libraries, installed dependencies, internal packages, and utilities before implementing. Reuse only when meaning, contract, ownership, lifecycle, and long-term cost fit.

## State and Side Effects

Minimize shared mutable state. Make ownership and mutation points clear. Separate core calculation from files, network, database, time, randomness, and other side effects when it improves testability and failure handling. For async work, consider cancellation, timeout, ordering, races, and cleanup.

## Error Handling

Distinguish expected failures from programming errors. Add context that helps diagnosis. Avoid broad catch blocks that erase cause. Preserve source errors when converting them. Check resource cleanup and transaction boundaries.

## Testability

Test observable behavior, not private implementation shape. Prioritize important success paths, boundary values, failure paths, and regression risk. Do not contort production design only to make a trivial test easier.

## Performance

Avoid speculative micro-optimization. Do avoid obvious repeated expensive calls, unbounded parallelism, unnecessary full scans, and excessive allocation. When performance is a requirement, prefer measurement over guesswork and document non-obvious tradeoffs.

## Security

Review input validation, authorization, secrets, logging, injection, unsafe paths, unsafe HTML, and dependency supply chain. Do not implement cryptography, authentication protocols, password hashing, or protocol parsers by intuition.

## Compatibility

Treat public APIs, stored data formats, migrations, runtime versions, platform support, and dependency versions as behavior. Changing them requires explicit compatibility and migration thinking.

## Ease of Change

A future maintainer should be able to identify what the code owns, what it assumes, where failures go, and which tests protect the behavior. Avoid unrelated churn that makes review and rollback harder.

## Operational Stability

Consider logs, metrics, retries, idempotency, backpressure, resource limits, cleanup, deploy behavior, and failure diagnosis where the code touches production systems or data.
