# Review Rubric

Review for risk first. Do not report style preferences as important defects.

## Severity

- `P0`: immediate blocker such as data loss, serious security incident, or broad outage.
- `P1`: likely bug or major regression in normal use.
- `P2`: conditional defect, important maintainability risk, or operational problem.
- `P3`: low-risk improvement with practical value.

## Finding Format

For each finding include:

- severity
- file and location
- why it is a problem
- plausible failure mode
- smallest useful fix direction
- assumptions or confidence when uncertain

## Checklist

### Correctness

Does the code implement the requested behavior? Are invariants preserved? Are public APIs and stored data formats unchanged unless approved?

### Failure Handling

Are expected failures represented? Are errors swallowed, over-broadly caught, or converted without cause? Are cleanup and transactions safe?

### Boundaries

Are empty input, null/undefined, missing files, invalid data, large data, time zones, locale, encoding, retry, timeout, and cancellation handled where relevant?

### Data Integrity

Can partial success, concurrent writes, retries, ordering, migrations, or fallback defaults corrupt or misrepresent data?

### Security

Check injection, path traversal, unsafe HTML, authorization, authentication, secret logging, PII exposure, insecure defaults, unsafe parsing, and custom crypto/protocol code.

### Performance

Look for repeated remote calls, unbounded concurrency, unnecessary full scans, excessive allocation, blocking work on hot paths, and scale-sensitive dependency APIs.

### Concurrency

Check race conditions, ordering assumptions, shared mutable state, idempotency, cancellation, locks, resource release, and retry safety.

### Compatibility

Check runtime versions, platform support, public API compatibility, stored data compatibility, dependency compatibility, and migration needs.

### Duplicate Existing Functionality

Has the implementation ignored existing domain modules, framework/system APIs, standard libraries, installed dependencies, or public utilities that match the contract?

### Wrong Utility Reuse

Has a utility been reused because of name or shape even though domain meaning, contract, errors, ownership, or lifecycle differ?

### Unnecessary External Dependency

Does a new dependency solve a small problem at high cost? Check license, security, package size, transitive dependencies, build/deploy impact, and replacement cost.

### Risky Direct Implementation

Is the code hand-implementing crypto, password hashing, auth protocols, parsers, date/time/time-zone logic, compression, network protocols, input sanitization, or database drivers?

### Speculative Abstraction

Does a new interface, wrapper, base class, shared utility, or factory exist only for a guessed future? Does it couple unrelated concepts?

### Change Scope

Are unrelated renames, moves, formatting changes, broad refactors, or metadata churn mixed into the change?

### Test Quality

Do tests verify observable behavior and important risks? Are they over-coupled to implementation details? Is a bug fix protected by a regression test when practical?
