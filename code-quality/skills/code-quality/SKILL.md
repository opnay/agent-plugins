---
name: code-quality
description: Write, modify, refactor, and review production code for correctness, readability, maintainability, testability, robustness, appropriate reuse, and simplicity. Inspect existing modules, framework and system APIs, standard libraries, installed dependencies, and utilities before introducing new implementations or dependencies. Use for implementation, bug fixes, refactoring, testing, architecture decisions, dependency decisions, and risk-focused code review. Do not use for prose-only tasks, translation, generated or vendor code, formatting-only changes, style-only changes, or naming-only changes unless the user explicitly invokes `$code-quality` or asks for code-quality, correctness, security, compatibility, data-integrity, operational-risk, or behavior-change review.
---

# Code Quality

Use this skill to make production code correct, understandable, maintainable, testable, robust, and no more complex than the problem requires. Do not mechanically apply any single book, doctrine, pattern, or style rule. Prefer evidence from the user request, repository rules, existing contracts, tests, framework APIs, standard libraries, installed dependencies, and nearby code.

## Reference Loading

Read only the references needed for the current task:

- `references/principles.md`: quality tradeoffs, correctness, clarity, simplicity, robustness, compatibility, security, and operations.
- `references/reuse-and-dependencies.md`: choosing among existing domain code, framework/system APIs, standard libraries, installed dependencies, internal packages, new dependencies, existing utilities, and direct implementation.
- `references/decision-guide.md`: duplication, function boundaries, abstraction, refactor scope, comments, errors, performance, utility reuse, wrappers, and dependency decisions.
- `references/review-rubric.md`: code review severity and risk checklist.
- `references/examples.md`: concrete bad/better examples. Read when the task resembles one of the examples or the tradeoff is unclear.

## Trigger Boundary

Use this skill for production code implementation, bug fixes, refactors, tests, architecture decisions, dependency decisions, and code review where correctness, behavior, security, compatibility, data integrity, operational risk, maintainability, or validation matters.

Do not use this skill for prose-only writing, translation, generated/vendor code edits, formatting-only changes, style-only changes, or naming-only changes. A request is not in scope merely because it explicitly asks for formatting, renaming, indentation, line wrapping, table alignment, or prettier output. The exception requires explicit `$code-quality` invocation or an explicit request to evaluate code quality, correctness, behavior change, security, compatibility, data integrity, operational risk, or production readiness.

For generated or vendor code, avoid direct edits by default. Engage only for a narrow code-quality task such as risk review, security review, compatibility review, wrapper/schema/source fix direction, rollback strategy, or an explicitly requested temporary patch with risk assessment.

For docs, examples, snippets, commands, SQL, or configuration text, engage only when the request asks whether the embedded code or command is correct, safe, compatible, behavior-changing, or operationally risky. Do not engage for tone, wording, translation, or presentation-only edits.

## Priority Order

When rules conflict, follow this order:

1. User requirements and behavior that must be preserved.
2. Correctness, data integrity, security, public API compatibility, and stored data compatibility.
3. Repository `AGENTS.md`, README, design docs, lint, format, type, and test settings.
4. Adopted shared modules, framework patterns, and dependency policy.
5. Consistent local codebase structure and conventions.
6. Language and framework idioms.
7. This skill's general quality guidance.

If a higher-priority instruction appears likely to cause a bug, security issue, data loss, or serious operational risk, state the risk instead of silently following it.

## Work Procedure

Before editing code, inspect enough context to understand:

- the user's actual goal, target files, adjacent code, callers, dependencies, existing tests, public API, and compatibility constraints
- repository rules and commands for build, test, lint, format, and type checks
- whether related behavior already exists in domain modules, shared modules, framework/system APIs, standard libraries, installed dependencies, internal packages, or utilities
- the implementation, contract, tests, usages, and error behavior of any reuse candidate

Do not infer repository-wide rules from one file fragment. Do not reuse a function by name alone; inspect what it does and how callers rely on it.

Before implementation, identify the behavior and risk:

- invariants, normal flow, failure flow, input/output boundaries, external communication, state changes, side effects, ordering, concurrency, regression risk, performance limits, security, privacy, and compatibility

For complex work, state a short plan. For simple work, proceed without creating extra design text.

Choose the smallest coherent design:

- Reuse existing code only when meaning, contract, ownership, lifecycle, and long-term cost fit the current need.
- Use framework, system, standard library, or installed dependency features when they fit better than custom code.
- Add a new dependency only when the benefit exceeds security, license, compatibility, package size, transitive dependency, supply chain, build, deployment, and removal costs.
- Implement directly when the logic is small, clear, domain-specific, or cheaper than forcing reuse.
- Add abstractions only when the same concept, contract, and reason to change are actually shared.

While editing:

- Preserve requested behavior and public contracts unless the user approved a change.
- Make invalid states harder to represent where practical.
- Keep side effects, state ownership, resource cleanup, retries, cancellation, and timeouts explicit.
- Distinguish expected failures from programming errors.
- Add comments only for rationale, constraints, external quirks, or deliberate tradeoffs.
- Keep unrelated renames, moves, formatting churn, and broad refactors out of the change.

After editing:

- Run the most relevant tests, type checks, lint, format checks, builds, or integration checks that repository settings reveal.
- If full validation is too expensive or unavailable, run the closest relevant checks and report what was not run and why.
- Re-read the diff for correctness, behavior preservation, failure paths, naming, unnecessary abstraction, dead/debug code, security, performance, concurrency, data integrity, duplicate implementations, forced utility reuse, and dependency cost.

## Review Mode

When reviewing code, diffs, commits, or PRs, prioritize real defects and maintainability risks over style preferences. Respect a focused review scope such as security-only or P0/P1-only unless another serious risk would be unsafe to omit.

For each finding, include severity, file/location, why it matters, a plausible failure mode, and the smallest useful fix direction. Use:

- `P0`: immediate blocker such as data loss, serious security incident, or broad outage.
- `P1`: likely bug or major regression in normal use.
- `P2`: conditional defect, important maintainability risk, or operational problem.
- `P3`: low-risk improvement with practical value.

Check correctness, failure paths, boundary values, data integrity, security, performance, concurrency, compatibility, duplicate existing functionality, wrong utility reuse, unnecessary dependencies, risky direct implementations, speculative abstractions, change scope, and test quality. Do not invent findings. Mark assumptions and confidence when uncertain.
