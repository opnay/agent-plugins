---
name: pro-engineering
description: Apply evidence-based engineering judgment to coding, debugging, refactoring, root cause analysis, implementation scope, verification, and technical reporting. Use when a task needs disciplined problem framing, practical code quality decisions, or a clear link between cause, change, and acceptance signal. engineering judgment, problem solving, root cause analysis, technical reasoning, code quality, implementation discipline
---

# Pro Engineering

Make engineering work explicit: define the problem, ground decisions in evidence, choose a narrow complete implementation, and report what was verified.

## Problem Frame

Separate the observed symptom from the expected behavior.

- Translate vague reports into observable inputs, actions, states, outputs, errors, timing, or user flows.
- Treat user explanations as clues, then verify them against code, configuration, fixtures, tests, logs, or reproduction steps.
- Keep confirmed facts separate from assumptions.
- List plausible cause areas before selecting one.
- For each cause, record supporting evidence and an observation that would falsify it.

Start with reproduction when possible. Otherwise use logs, tests, traces, or a minimal experiment to narrow the cause. Avoid broad structural changes until the failure mechanism is predictable. Distinguish product behavior from assertion, harness, infrastructure, and external-system failures.

Prefer the hypothesis that explains the observed facts with few extra assumptions and survives a concrete counterexample.

## Scope Control

Explore broadly before direct coding. Keep plausible systems, contracts, data paths, alternatives, failure conditions, and risks visible without reading files indiscriminately.

Focus implementation on the code path tied to the selected cause and acceptance signal. Preserve the boundaries and non-goals learned during exploration while changing the smallest surface that connects cause to result.

Move from exploration to implementation only when:

- The symptom and expected behavior are testable.
- At least one cause is connected to evidence.
- Relevant and irrelevant areas are separated.
- The chosen approach can be explained against alternatives.
- Edited files, contracts, and checks map to an acceptance signal.

When one implementation is too large to complete and verify coherently, split it by behavior and acceptance signal. Each deliverable must be independently executable and verifiable, include its normal path, relevant failure path and prerequisites, and leave the repository in a coherent state.

Do not treat files, layers, or types as deliverables by themselves. Keep a transaction, invariant, or required ordering in one unit when independent stages would create an invalid intermediate state.

If the entry conditions are missing, continue discovery or ask for the product, risk, deployment, or scope decision that local evidence cannot provide.

## Engineering Judgment

Prefer the repository's existing structure, helpers, names, error handling, and test style. Introduce a new pattern only when existing patterns cannot meet the goal or are part of the problem. Require an explicit reason before implementing the same behavior a second way.

### Ownership Boundaries

Put behavior, rules, state, validation, and fallback contracts with their owner: caller, callee, domain owner, adapter, storage, UI, or test harness. A nearby file is not automatically the right boundary.

Treat these as boundary-leak signals:

- A caller must know a callee's internal state, storage order, cache policy, or recovery detail.
- UI code owns domain validation or persistence rules.
- An adapter makes product or domain decisions.
- A shared helper mixes decisions from multiple ownership boundaries.

If ownership is unclear, narrow the contract name, call direction, state owner, and failure responsibility before implementing. Validate at the earliest boundary that can catch the issue without duplicating the rule.

Distinguish an intended product fallback from failure masking. State its activation condition, handled failure, owner, and observable signal; do not silently swallow a failure.

### Implementation Logic Responsibilities

Use domain, coordination, and foundation as responsibility lenses, not mandatory physical layers, folders, or files.

- Domain responsibility owns product policy, meaningful values and thresholds, invariants, eligibility, and state transitions. HTTP, database representation, JSON shape, or UI state leaking into it signals a misplaced boundary.
- Coordination responsibility maps external inputs to domain inputs, sequences calls, and translates results and failures. It may connect domain and foundation capabilities, but it must not invent thresholds, eligibility, defaults, or other product policy.
- Foundation responsibility provides domain-neutral primitives such as parsing, formatting, collection, time, or URL operations. Product API paths, domain state, or UI types make a helper product-owned rather than foundation logic.

Adapters, storage, and infrastructure do not become foundation responsibility merely because they are lower-level. A boundary that knows a product contract or external representation keeps its adapter or storage ownership; only its domain-neutral mechanics may be reused as foundation logic.

A typical flow is external or UI input to coordination, then domain decisions and state transitions, then coordination translating the result or failure back to the external boundary. Foundation capabilities may support the flow, but they must not depend upward on domain, coordination, or UI ownership.

A small cohesive unit may perform more than one responsibility when its ownership, data flow, and verification remain clear. Do not create physical boundaries solely to mirror these labels.

### Module and Reuse Boundaries

Start with the smallest cohesive implementation near its owner. Keep a small one-owner domain rule local even when it could be parameterized or generalized.

Extract a boundary only when the present maintenance or verification benefit exceeds the indirection, dependencies, parameter passing, coordination, and compatibility cost it creates. Code size and hypothetical reuse are not sufficient evidence.

Separate code when responsibility, reason to change, data flow, or ownership diverges and the new boundary can state its inputs, outputs, errors, side effects, and call direction. The boundary should support focused verification.

Keep a stable rule shared by real consumers in one reusable module owned by the domain that defines it. Separate the rule from I/O and orchestration so callers follow one contract.

When consumers disagree, first determine whether the behavior should be one policy or an intentional variation. If local evidence cannot establish the product policy, stop before extraction and ask its owner; do not propose or implement a default.

Reuse foundation logic when current consumers share the same domain-neutral contract. Do not combine unrelated primitives merely because they fit under a `utils` name.

Choose a separate package or shared library only when consumers, distribution, versioning, and compatibility responsibilities are genuinely independent. Prefer an internal module when ownership and release cadence remain shared.

Do not split cohesive product policy merely to make it look smaller. Do not add interfaces, layers, packages, or stages for hypothetical reuse. A processing stage is justified only when it has an independent contract and preserves transactions, invariants, and required ordering.

### Contracts and Abstraction

Define valid inputs, states, transitions, effects, and failure conditions before expanding blocklist cases. Route unknown or invalid values to explicit failure and use blocklists only as narrow supplemental safeguards. If exceptions grow larger than the allowed contract, revisit the contract owner or data model.

Add abstraction only when it removes real duplication, centralizes a rule, improves testability, clarifies ownership, or matches an existing design. A one-use function or wrapper still needs a concrete readability, testability, or failure-boundary benefit.

Make inputs, outputs, errors, side effects, and ownership visible in code or tests. At strings, JSON, external inputs, network, filesystem, and process boundaries, prefer structured parsing, validation, schemas, and explicit failures over ad hoc assumptions.

Require stronger evidence and verification for concurrency, async ordering, caching, time, timezone, randomness, retry, timeout, backoff, external services, network, filesystems, process boundaries, auth, permissions, migrations, destructive actions, shared libraries, public APIs, or broad UI workflows.

Ask for user direction when product policy, risk tolerance, deployment strategy, or essential external state cannot be inferred. Always ask before breaking a public contract or requiring a migration. Otherwise continue from local evidence.

## Code Discipline

Start with the simplest complete implementation tied directly to the selected cause and place it near the responsible owner. Include the normal path, relevant failure path, and meaningful verification. Extraction is not an improvement by itself; require a concrete reduction in current reading, change, or verification cost.

Keep even small changes with the correct owner. Do not accept a deliverable that leaves modules, transactions, invariants, or ordering in an incoherent intermediate state.

If a public contract changes, update or assess affected callers, documentation, tests, and migration requirements in the same scope.

Then improve within the same task:

1. Check names, control flow, duplication, edge handling, and failure reporting.
2. Refine only what is hard to read, risky, or inconsistent with local patterns.
3. Confirm the original acceptance signal still holds.

Use names that reveal the domain concept, state, or failure. Prefer early returns for guards, validation failures, and cannot-handle states when they reduce nesting and clarify the normal path. Keep a single exit when cleanup, transactions, locks, or `finally`-style safety makes it clearer or safer.

Separate error cause from recoverability. Add comments only for non-obvious constraints, decisions, or failure intent. Handle edge cases when they follow from a known failure mechanism or explicit contract.

Keep changes inside the selected ownership boundary. Avoid unrelated cleanup, formatting churn, speculative hardening, or broad restructuring. Preserve nearby user-owned changes and adapt to the current worktree.

Avoid symptom-hiding patches, silent fallbacks, weakened assertions, responsibility leaks, blocklist-only fixes, speculative extension points, one-owner domain rules extracted only for possible generalization, and shared packages without real consumers or independent release responsibilities.

## Verification

Choose the narrowest verification that proves the changed contract, then broaden when risk demands it.

- When regression risk exists, consider adding or updating tests.
- For local logic, use focused tests or a minimal reproduction.
- For core harness logic, shared behavior, public APIs, cross-module contracts, or user-facing workflows, add or run representative integration or end-to-end checks.
- For a new module or processing-stage boundary, verify its inputs, outputs, failures, side effects, and one representative end-to-end data flow.
- When responsibilities move, verify the policy owner, boundary mapping, dependency direction, and representative flow. Domain logic must not learn external representations, coordination must not choose product policy, and foundation logic must not depend on product or UI ownership.
- For a reusable boundary, identify its current consumers, run focused contract checks, and verify representative callers. For a one-consumer extraction, confirm that its current maintenance or verification benefit exceeds its added cost.
- For configuration changes, verify parsing or loading and one affected path.
- For fixtures or sample data, run a consuming scenario.
- For reporting or logging changes, check machine readability and operator usefulness when applicable.

Do not hide flaky or unclear failures with blind retries. Separate infrastructure, harness, assertion, and product failures. For an unreproduced failure, state what was checked and what remains uncertain. If verification is blocked by permissions, sandboxing, network, external state, or missing tools, report the blocker and alternative evidence separately.

## Completion and Reporting

Call work complete only when:

- The original symptom and expected behavior are explained.
- The change maps to a confirmed cause or explicit contract.
- Verification is proportional to the affected risk.
- Unrun checks and residual risk remain visible.
- The reported state matches current files and command results.

Report the current state, handled scope, changed files, verification, unrun checks, and residual risk. Mention abandoned attempts only when needed to understand the result.

For implementation or harness work, include:

- `Scope handled`
- `Files changed`
- `Verification`
- `Residual risk`
