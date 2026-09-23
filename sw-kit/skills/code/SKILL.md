---
name: code
description: Implement, modify, refactor, test, and review production code for correctness, clarity, maintainability, and safe change. Use for source-level work after a technical direction is known; surface material architecture decisions instead of making them implicitly.
---

# Code

Implement the requested behavior with clear, maintainable code within the chosen technical direction.

## Work

Read the relevant code, callers, tests, contracts, and repository conventions. Follow an existing engineering direction. Surface material system, data, or technology choices for `$sw-kit:engineering` when the direction is still open.

Make the smallest coherent change. Keep control flow, ownership, side effects, errors, and boundary conditions understandable. Reuse code only when its meaning, contract, ownership, and lifecycle fit. Avoid speculative abstractions and unrelated cleanup.

Use blank lines to group a function by purpose and reading effort. Treat an entry guard at the start of a function as the first reading unit; leave a blank line after it before the main flow. Keep consecutive one-line `if` checks together when they serve one checking step. Keep preparation and immediate checks together unless a dense multiline declaration or expression needs its own reading space.

Name the result of a meaningful call before passing it to another call when nesting hides a step. Direct nesting is fine when the flow stays clear.

Preserve public contracts unless their change is approved. Check existing framework, system, standard-library, and installed dependency capabilities before adding a new implementation or dependency.

## Verify and review

Run checks proportionate to the change and its risks. Re-read the diff for behavior, failure paths, compatibility, data integrity, security, concurrency, and unnecessary complexity.

For review-only work, report actionable findings by severity, location, failure mode, and a useful fix direction. Do not edit code or present style preferences as defects.

## References

- Read `references/principles.md` for quality tradeoffs.
- Read `references/reuse-and-dependencies.md` before a reuse or dependency decision.
- Read `references/decision-guide.md` for a local structure or abstraction tradeoff.
- Read `references/review-rubric.md` for review severity and coverage.
- Read `references/examples.md` when a concrete comparison helps.
