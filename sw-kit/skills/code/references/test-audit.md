# Focused Test Audit

Read this for a focused audit of existing tests proposed for deletion or consolidation. Use the test value, authoring gate, junk patterns, and retention bar in [SKILL.md](../SKILL.md#test-value); they are the shared criteria, not repeated here. For a whole-subsystem cleanup, also read [test-campaign.md](test-campaign.md).

## Read-only discovery

Read root and scoped `AGENTS.md` files first. Before judging a candidate, read its complete test and production owner, entry point, callers, callees, sibling implementations, overlapping tests, CI routing, and relevant history. Inspect dependency source or types directly when the test claims dependency-backed behavior.

Report evidence before editing. For broad discovery, use parallel read-only lanes when delegation is available and permitted: core/packages, plugins, UI/apps/scripts/tooling, and a cross-cutting pattern sweep. Adapt lanes to actual ownership. Outside a campaign, prefer a few high-confidence candidates over a speculative inventory.

For every candidate, record:

- Exact test name and location.
- Failure it can actually detect.
- Non-test callers of its production or support seam.
- Stronger remaining proof at the owning boundary, or why no proof is needed.
- Relevant history and why the test or seam exists.
- Production or test-support deletion it unlocks.
- Risk and focused validation command.

A missing field means the candidate is not ready for deletion. Keep uncertain candidates as investigation items. Baseline failures in retained contracts are possible product defects, not cleanup opportunities.

## Edit one coherent boundary

Apply edits only within the requested implementation scope. Name the keeper for each meaningful contract and carry unique assertions into it before retiring duplicate tests. Move retained regressions to canonical owners; consolidate repeated package or dependency assertions at the shared owner.

Remove obsolete private test-only exports, globals, wrappers, and dead paths instead of retaining aliases. Preserve public or external contracts and meaningful dependency boundaries. Prefer net-negative production LOC without weakening proof or adding replacement tests that restate implementation. Do not turn uncertain candidates into cleanup to increase deletion counts.

## Validate

Use the target repository's testing policy, runner, resource limits, and required gates. Do not edit source or tests while validation is running in that checkout.

1. Run the smallest owner and sibling tests with the adopted runner.
2. For removed source greps or plan assertions, run the executable script or dry-run owning the real contract.
3. Run targeted formatting and `git diff --check`.
4. Classify changed paths with the repository's changed-check mechanism when present, then run required changed gates.
5. Inspect `git diff --numstat`; separate production/tooling from tests and support.
6. After final audit edits, obtain an independent review when available. Repository-required review remains a mandatory gate. Disclose unavailable review and alternative proof.

Do not report partial or unrun validation as passing. Preserve relevant failure cases even when a retained test is slow or static.

## Handoff and continuation

Report the removed low-value categories and their cause, production simplifications, retained false positives and their value, focused and full proof actually run, production versus test/support LOC, publication state, and named follow-ups.

Commit, push, PR, and merge require existing authorization and the owning Git workflow. Keep each published change coherent. After landing, refresh from the current upstream base and repeat read-only discovery for the next batch.
