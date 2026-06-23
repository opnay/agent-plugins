# Reuse and Dependencies

Reuse is a judgment about meaning, contract, ownership, lifecycle, risk, and long-term cost. The source of a candidate is not an absolute priority.

## Candidate Sources

Consider these before writing new code:

1. repository domain modules or public shared modules
2. adopted platform or wrapper layers
3. framework APIs
4. system APIs
5. standard libraries
6. installed external dependencies
7. organization-owned internal packages
8. well-maintained new external libraries
9. direct implementation

Choose the option that best fits the current requirement. Do not force an earlier source if its contract is wrong.

## Repository Domain Code

Prefer existing domain modules when they represent the same business concept, state model, terminology, and reason to change. Verify implementation, tests, callers, and documented contract. Do not merge unrelated concepts only because the code shape is similar.

## Framework APIs

Before implementing authentication, request validation, serialization, caching, job scheduling, lifecycle management, logging, tracing, permissions, transactions, resource management, cancellation, or timeouts, check the framework's supported APIs and the repository's adopted patterns.

Do not use a framework API when it does not match requirements, supported versions, testability, observability, platform constraints, security, performance, or an existing repository abstraction.

## System APIs

Use operating system or platform APIs for files, networking, secure storage, background execution, permissions, process lifecycle, and resource limits when they provide the right contract. Avoid platform coupling when portability is a requirement.

## Standard Libraries

Check standard libraries before implementing dates, time zones, paths, URLs, encoding, sorting, search, concurrency, hashing, cryptography primitives, data structures, numeric conversion, I/O, resource management, regex, compression, archives, UUIDs, and identifiers.

Standard library availability is not enough. Check supported runtime version, API safety, errors, performance, and cross-platform behavior.

## Installed Dependencies

If a dependency is already installed, check whether it is used for similar work, whether the API is public and stable, whether it ships in production, and whether using it avoids more complexity than it adds. Do not use dev-only or test-only dependencies in production code unless the project explicitly supports that.

## Internal Packages

Treat organization-owned packages like external dependencies. Confirm ownership, support scope, version policy, deprecation policy, incident responsibility, and release compatibility.

## New External Libraries

Add a dependency only when the benefit exceeds long-term ownership cost. Review fit, maintenance, release cadence, security history, license, runtime support, API stability, package size, transitive dependencies, performance, observability, build/deploy impact, supply chain risk, and removal cost.

Prefer vetted libraries for cryptography, authentication protocols, time zones, file format parsing, network protocols, compression formats, internationalization, sanitization, database drivers, and complex parsing.

Do not add a large dependency for a few lines of clear, low-risk logic.

## Existing Utilities

Do not reuse utilities by name alone. Inspect implementation, tests, callers, side effects, boundaries, and error behavior. Reuse when the utility's meaning and contract match the current requirement and it is genuinely maintained as shared functionality.

If a utility is a poor abstraction, avoid increasing its usage. Prefer local logic, a domain-owned module, a compatible improvement, an adapter with clear migration path, or a follow-up refactor note.

Avoid adding vague functions to `utils`, `common`, `shared`, or `helpers`. Prefer ownership near the domain or module that owns the concept.

## Direct Implementation

Direct implementation can be best when the logic is small and clear, the domain meaning differs from existing code, error contracts differ, reuse adapters would be more complex, performance or version constraints matter, the existing code is private implementation detail, or coupling would make independent change harder.

If choosing direct implementation, re-check that standard, framework, or existing dependency support was not missed.

## Framework Wrappers

Create a wrapper only for real responsibility:

- applying application-wide policy
- isolating external API volatility
- unifying an error model
- adding logging or observability
- providing meaningful test seams
- expressing domain meaning
- separating platform-specific implementations

Do not wrap only because a future replacement is imaginable.

## Reuse vs Abstraction

Using an existing, tested capability is reuse. Generalizing one local implementation for imagined future callers is speculative abstraction. Similar code is not enough; require shared concept, contract, and reason to change.

## Wrong DRY

Do not combine two rules only because both normalize strings, parse objects, call an API, or share a loop shape. A display name and a payment identifier can both be strings and still require separate rules.

## Long-Term Cost

Every shared utility or dependency creates ownership, compatibility, documentation, testing, security, and replacement cost. Count that cost before choosing reuse or adding dependencies.

## Security, License, and Platform

For dependencies and copied code, check vulnerability handling, license compatibility, supported runtimes, platform behavior, and whether production builds include the code as intended.

## Removal and Replacement

Prefer choices that can be removed or replaced without spreading private implementation details across the codebase.
