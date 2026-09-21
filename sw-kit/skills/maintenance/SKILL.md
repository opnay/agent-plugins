---
name: maintenance
description: Audit and plan safe codebase maintenance for dependencies, deprecations, dead code, drift, and technical debt. Use for lifecycle work rather than routine feature implementation, architecture selection, or Git workflow.
---

# Maintenance

Find and reduce lifecycle cost without weakening real contracts. This skill owns maintenance analysis and plans; `$sw-kit:code` owns a requested source-level implementation and `$toolkit:git` owns Git workflow.

## Investigate before recommending change

Confirm actual usage, public or external contracts, generated/vendor boundaries, dependency ownership, compatibility, and migration cost. Treat uncertain consumers as investigation, not proof that code can be removed.

Prefer deletion, consolidation, replacement, or deprecation only when behavior, security, accessibility, data integrity, operational safeguards, and supported compatibility remain protected. Break broad cleanups into independently verifiable units with a migration and rollback path.

## Report actionable maintenance work

For every finding, give its location, evidence, cost or risk, recommended action, compatibility impact, verification, and rollback condition. Rank broad audit findings by impact and confidence. Review and audit requests remain read-only unless the user also requests implementation.

## References

- Read `references/repo-audit.md` for a repository audit.
- Read `references/dependency-check.md` for add, remove, replace, or deprecation decisions.
- Read `references/refactor-shrink.md` before simplifying or removing behavior-preserving code.
- Read `references/debt-ledger.md` for debt-marker or deferred-work analysis.
- Read `references/impact-scoreboard.md` to rank many findings.
- Read `references/safety-boundaries.md` before deletion, migration, or reduction.
- Read `references/output-style.md` for a compact maintenance report.
