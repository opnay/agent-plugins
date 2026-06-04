---
name: pro-engineering
description: Apply evidence-based engineering judgment to coding, debugging, refactoring, root cause analysis, implementation scope, verification, and technical reporting. Use when a task needs disciplined problem framing, practical code quality decisions, or a clear link between cause, change, and acceptance signal. engineering judgment, problem solving, root cause analysis, technical reasoning, code quality, implementation discipline
---

# Pro Engineering

Use this skill to make engineering work explicit: define the problem, ground decisions in evidence, choose a narrow complete implementation, and report what was verified.

## Problem Frame

Start by separating the observed symptom from the expected behavior.

- Translate vague reports into observable conditions: input, action, state, output, error, timing, or user flow.
- Treat user explanations as important clues, but verify them against code, logs, tests, fixtures, configuration, or reproduction steps.
- Keep confirmed facts separate from assumptions.
- List plausible cause areas before committing to one explanation.

A good root-cause hypothesis explains the observed facts with few extra assumptions and can be challenged by a concrete counterexample.

## Scope Control

Use different scope widths at different phases.

Before direct coding, explore broadly enough to understand the problem space. This applies during research, requirement clarification, failure-path discovery, implementation sizing, alternative comparison, and risk assessment. Broad exploration does not mean reading files indiscriminately; it means keeping plausible systems, contracts, data paths, and edge conditions visible until you can justify what matters.

When implementation begins, narrow attention to the code path tied to the chosen cause and acceptance signal. Narrow implementation does not mean ignoring surrounding context; it means preserving the boundaries and non-goals learned during exploration while changing the smallest surface that connects cause to result.

Move from broad exploration to focused implementation only when:

- The symptom and expected behavior are testable.
- Relevant and irrelevant areas have been separated.
- The chosen approach can be explained against alternatives.
- The edited files, contracts, and checks map to an acceptance signal.

If these conditions are missing, continue discovery or ask the user for the product, risk, or scope decision that cannot be inferred from local evidence.

## Engineering Judgment

Prefer the repository's existing patterns, helpers, names, error handling, and test style. Introduce a new pattern only when existing patterns cannot meet the goal or are part of the problem.

Check ownership boundaries when a change may move, expose, duplicate, or blur responsibility for behavior, rules, state, validation, or fallback behavior. Put the contract where responsibility belongs: caller, callee, domain layer, adapter, storage, UI, or test harness. A nearby file is not automatically the right boundary.

Treat these as boundary-leak signals:

- A caller must know a callee's internal state, storage order, cache policy, or recovery detail.
- UI code owns domain validation or persistence rules.
- An adapter makes product or domain decisions.
- A shared helper mixes decisions from multiple ownership layers.

If ownership is unclear, narrow the contract name, call direction, state owner, and failure responsibility before implementing. Put validation at the earliest boundary that can catch the issue without duplicating the rule across layers. Use fallbacks only when the owning layer can expose when they were used and what failure they handled.

Add abstraction only when it removes real duplication, centralizes a rule, improves testability, clarifies ownership, or matches an existing design. Do not add abstraction for speculative future use.

Make contracts visible. Inputs, outputs, errors, side effects, and ownership boundaries should be clear in code or tests. For fragile boundaries such as strings, JSON, external inputs, and process or network edges, prefer structured parsing, validation, schemas, or explicit failure handling over ad hoc assumptions.

Require stronger evidence and verification when the change touches concurrency, async ordering, caching, time, timezone, randomness, retry, timeout, external services, filesystems, auth, permissions, migrations, destructive actions, shared libraries, public APIs, or broad UI workflows.

## Code Discipline

Start with the simplest complete implementation that directly addresses the selected cause. Small does not mean partial: include the normal path, relevant failure path, and meaningful verification.

Small changes still need the right owner when they touch responsibility boundaries. Do not patch the closest file if the behavior belongs in another layer or contract boundary.

Then improve within the same task:

1. Check whether names, control flow, duplication, edge handling, and failure reporting are clear.
2. Refine only the parts that are actually hard to read, risky, or inconsistent with local patterns.
3. Confirm the original acceptance signal still holds after cleanup.

Keep changes inside the ownership boundary. Do not mix unrelated cleanup, formatting churn, speculative hardening, or broad restructuring into the fix. If nearby files already contain unrelated changes, treat them as user-owned and adapt to the current state rather than reverting them.

Avoid fixes that only hide symptoms, silent fallbacks that mask failures, weakened assertions, responsibility leaks, and future-proofing that is not tied to the observed problem or an explicit contract.

## Verification

Choose the narrowest verification that proves the changed contract, then broaden only when risk demands it.

- For local logic, use focused tests or a minimal reproduction.
- For shared behavior, public APIs, cross-module contracts, or user-facing workflows, add or run representative integration or end-to-end checks.
- For configuration changes, verify parsing or loading and one affected path.
- For fixtures or sample data, run the scenario that consumes them.
- For reporting or logging changes, check both machine readability and operator usefulness when applicable.

Do not hide flaky or unclear failures with blind retries. Separate infrastructure failure, harness failure, assertion failure, and product behavior. If verification is blocked by permissions, sandboxing, network, external state, or missing tools, report the blocker and any alternative evidence separately.

## Reporting

Report the current state only. Connect the work to the confirmed problem or contract, list changed files, state what was verified, and name residual risk.

For implementation or harness work, include:

- `Scope handled`
- `Files changed`
- `Verification`
- `Residual risk`
