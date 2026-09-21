---
name: code
description: Implement, modify, refactor, test, and review production code for correctness, clarity, maintainability, and safe change. Use for source-level work after a technical direction is known; surface material architecture decisions instead of making them implicitly.
---

# Code

Build correct, understandable code without making it more complex than the current problem requires.

## Work

Inspect the target behavior, nearby code, callers, tests, repository rules, and relevant contracts before changing code. Follow an explicit engineering direction when one exists. If a material system, data, or technology choice is still open, surface it for `$sw-kit:engineering` instead of hiding the choice in implementation.

Choose the smallest coherent change. Reuse existing code only when its meaning, contract, ownership, and lifecycle fit. Prefer existing framework, system, standard-library, or installed dependency capabilities when they fit; add a dependency only when its benefit exceeds its cost and risk.

Keep control flow, ownership, side effects, errors, and boundary conditions clear. Avoid speculative abstractions, unrelated cleanup, and behavior changes outside the requested scope. Preserve public contracts unless their change is approved.

## Verify and review

Run proportionate checks from the repository: focused tests first, then type, lint, build, or integration checks when the risk requires them. Re-read the diff for behavior preservation, failure paths, compatibility, data integrity, security, concurrency, and unnecessary complexity.

For review-only work, do not edit. Report real findings by severity, location, failure mode, and smallest useful fix direction. Do not invent style findings.

## References

- Read `references/principles.md` for quality tradeoffs.
- Read `references/reuse-and-dependencies.md` before a reuse or dependency decision.
- Read `references/decision-guide.md` for a local structure or abstraction tradeoff.
- Read `references/review-rubric.md` for review severity and coverage.
- Read `references/examples.md` only when a concrete comparison helps.
