---
name: engineering
description: Choose a sound technical direction for a software behavior contract, including architecture, data model, technology, integration, migration, and operations. Use when a material system or technology choice must be made before implementation.
---

# Engineering

Turn a behavior contract into a technical direction. Keep product policy with its owner and source-level implementation details with `$sw-kit:code`.

## Frame the decision

Establish the required behavior, non-functional constraints, current system, data ownership, external boundaries, and operating environment. Mark unknown facts as assumptions.

Compare viable options by requirement fit, complexity, change cost, reliability, security, performance, and operability. Prefer the smallest direction that meets present constraints without blocking known change.

## Define the direction

Make these decisions explicit when relevant:

- system boundaries, owners, and integration contracts
- data model, lifecycle, consistency, retention, and recovery
- technology, framework, storage, and dependency choices
- failure handling, observability, capacity, security, migration, and rollback

State the chosen direction, alternatives rejected or deferred, rationale, consequences, and validation needed. Escalate decisions that require product policy, risk tolerance, or unavailable external facts.

## Handoff

Give implementation a concise contract: the boundaries to preserve, data and integration contracts, key failure modes, and acceptance or operational signals. Do not prescribe files, local abstractions, or routine code style.
