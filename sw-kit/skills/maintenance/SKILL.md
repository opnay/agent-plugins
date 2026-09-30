---
name: maintenance
description: Audit and plan safe codebase maintenance for dependencies, deprecations, dead code, drift, technical debt, and test debt. Use for lifecycle analysis rather than feature implementation, architecture selection, or Git workflow.
---

# Maintenance

Reduce lifecycle cost while preserving real contracts. Own investigation and maintenance plans; `$sw-kit:code` owns requested source-level implementation and `$toolkit:git` owns Git workflow.

## Investigate

Confirm actual usage, public and external contracts, generated/vendor boundaries, dependency ownership, compatibility, and migration cost. Uncertain consumers remain investigation items, not proof that code can be removed.

Recommend deletion, consolidation, replacement, or deprecation only while behavior, security, accessibility, data integrity, operational safeguards, and supported compatibility remain protected. Split broad cleanup into independently verifiable units with migration and rollback conditions.

For test-debt deletion or consolidation analysis, read the test criteria in `$sw-kit:code` and `$sw-kit:code/references/test-audit.md`. For an explicitly requested whole-subsystem test audit, also use the read-only discovery and planning stages in `$sw-kit:code/references/test-campaign.md`. These bundled resources share the owning contract; no prior execution of a sibling skill is required. Load them only for that test-debt task. Keep investigation and recommendations read-only; distinguish them from requested implementation.

## Report

For each finding, give location, evidence, cost or risk, recommended action, compatibility impact, verification, and rollback condition. Rank broad findings by impact and confidence. Audit and review requests remain read-only unless implementation is also requested.

## References

- `references/repo-audit.md`: repository audit.
- `references/dependency-check.md`: dependency addition, removal, replacement, or deprecation.
- `references/refactor-shrink.md`: behavior-preserving simplification or removal.
- `references/debt-ledger.md`: debt markers and deferred work.
- `references/impact-scoreboard.md`: ranking many findings.
- `references/safety-boundaries.md`: deletion, migration, or reduction.
- `references/output-style.md`: compact maintenance reports.
