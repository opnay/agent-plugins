# Repo Audit

Use this reference for whole-repository overengineering, dependency, removable-code, or lean debt audits.

Default mode is read-only.

## Scan Scope

Start from the repository root unless the user names a narrower path.

Include:

- source code
- tests
- scripts
- config
- package manifests
- local docs that define behavior

Exclude:

- `.git`
- dependency folders
- build outputs
- generated artifacts
- cache directories
- lockfiles unless dependency choice is in scope
- release copies generated from dev source unless release drift is in scope

## Audit Passes

1. Tree map: identify modules, entry points, and generated/excluded areas.
2. Dependency pass: inspect manifests and actual imports.
3. Abstraction pass: find one-implementation interfaces, wrapper-only layers, speculative config, and unused extension points.
4. Duplication pass: find repeated local logic that can reuse existing helpers or platform behavior.
5. Debt marker pass: search `lean:` and `ponytail:` if debt is in scope.
6. Verification pass: identify the check needed for each proposed reduction.

## Ranking

Rank by reduction impact:

- user-visible risk reduced
- files or dependencies removed
- concepts removed
- future maintenance avoided
- confidence that behavior is unchanged
- verification cost

Use `impact-scoreboard.md` for scoring when there are many findings.

## Output

Report:

- scope scanned
- exclusions
- findings ordered by impact
- no-change areas worth keeping
- follow-up sequence if the user wants implementation
