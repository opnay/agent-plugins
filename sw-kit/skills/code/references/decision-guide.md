# Decision Guide

Use this guide when the code quality tradeoff is not obvious.

## Duplication

Remove duplication when two places express the same concept, contract, validation rule, error behavior, and reason to change. Keep small duplication when the two cases are likely to evolve independently or when a shared abstraction would hide domain meaning.

## Function Boundaries

Split a function when it mixes abstraction levels, independent reasons to change, unrelated side effects, or a condition whose meaning deserves a domain name. Keep a function together when splitting would only move obvious lines into tiny helpers.

## New Abstractions

Introduce an abstraction when it reduces real complexity for current callers and captures a stable concept. Do not add interfaces, base classes, wrappers, factories, or registries just to make code look more designed.

## Following Existing Structure

Follow established repository structure when it is consistent and supports the current behavior. Improve locally when the existing pattern is inconsistent, unsafe, deprecated, or would force the wrong contract. Keep improvements scoped to the task.

## Refactor Scope

Widen refactoring only when it is required to implement the request safely, preserve a shared contract, or remove a local hazard created by the change. Do not mix unrelated renames, moves, or broad formatting with behavior changes.

## Comments

Add comments for rationale, constraints, external system quirks, unusual tradeoffs, or why direct implementation was chosen over standard support. Remove or update stale comments. Do not repeat what the code says.

## Return Errors or Throw Exceptions

Follow the language, framework, and repository convention first. Return expected recoverable failures when callers normally branch on them. Throw or propagate exceptions for programming errors, unexpected infrastructure failures, or conventions that expect exceptions. Preserve diagnostic context either way.

## Performance Optimization

Optimize when a requirement, measurement, or obvious scale risk justifies it. Examples: repeated remote calls in a loop, unbounded concurrency, full-table scans, unnecessary parsing of large data, or hot-path allocations. Record measurement or rationale when the optimization makes code less direct.

## Existing Utility Reuse

Use an existing utility only after checking implementation, tests, usages, side effects, and error contract. Reject it when its domain meaning differs, it is private to another module, or adapters make the call more complex than local code.

## Framework API or Wrapper

Use the framework API directly when its contract matches and the repository does not already centralize the concern. Create or use a wrapper only for policy, error model, observability, testing, domain meaning, API volatility, or platform separation.

## New Dependency or Direct Implementation

Choose a new dependency for complex, security-sensitive, standard-heavy, or edge-case-heavy behavior where vetted maintenance matters. Choose direct implementation for small, clear, low-risk logic or domain-specific rules. Check license, security, runtime, package size, transitive dependencies, build impact, and replacement cost before adding a dependency.

## Standard Library or External Library

Prefer standard library support when it meets requirements safely and clearly. Prefer an external library when standard support is incomplete, error-prone, unsupported on target runtimes, or lacks required edge-case handling.

## Private Implementations

Do not depend on private internals of another module, framework, or library unless the repository already owns that contract or no stable option exists and the risk is explicitly accepted.
