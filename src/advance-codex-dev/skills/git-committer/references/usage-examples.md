# Usage Examples

## Purpose
Provide concrete examples for commit types, messages, granularity, and verification checks.

## Commit types
Use the most specific type that fits the change.

- feat: new user-facing feature
- fix: bug fix
- refactor: code change that neither fixes a bug nor adds a feature
- docs: documentation-only changes
- test: adding or updating tests
- perf: performance improvement
- style: formatting or style-only changes (no behavior change)
- build: build system or dependencies
- ci: CI configuration or scripts
- chore: maintenance tasks not in the above

## Message example

```
feat: login page for google oauth 2.0

- add `pages/login`
- add login api for backend `/api/auth/google`
- connect pages and api for google login
```

## Commit granularity examples

- Package updates:
  - `honojs` upgrade commit
  - fix commit for `honojs` upgrade errors (if any)
  - `pg` upgrade commit
- Login feature in a monorepo:
  - REST API design doc commit
  - DB design doc commit
  - DB schema apply commit
  - table name typo fix commit
  - missing index fix commit

## Staged verification example (scenario-based)

- Always confirm staged scope with `git status` and `git diff --staged` before commit.
- Example only: if staged verification does not cover risk for modified Vite components, run a supporting check such as lint, typecheck, or build.
- Choose supporting checks based on the actual project and staged scope.
- If staged verification or a supporting check is skipped or unavailable, report the skip reason and residual risk; do not call it pass.
