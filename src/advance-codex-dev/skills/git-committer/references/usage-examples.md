# Usage Examples

## Commit Types

Use the most specific type that fits the change.

- `feat`: new user-facing feature
- `fix`: bug fix
- `refactor`: behavior-preserving code restructuring
- `docs`: documentation-only change
- `test`: test addition or update
- `perf`: performance improvement
- `style`: formatting or style-only change
- `build`: build system or dependency change
- `ci`: CI configuration or script change
- `chore`: maintenance outside the above types

## Message Example

```text
docs: clarify deployment prerequisites

- document required runtime and environment configuration
- verify the staged Markdown diff and whitespace
```

## Commit Granularity

- Package updates:
  - dependency upgrade
  - upgrade-related fixes, if needed
- Login feature in a monorepo:
  - API design
  - database design
  - schema application
  - later typo or missing-index fixes

## Verification Scope

- Always confirm staged scope with `git status` and `git diff --staged` immediately before commit.
- Run the narrowest deterministic supporting check for the actual staged risk.
- For docs-only changes, read back the staged diff and check formatting or whitespace.
- For code changes, choose among lint, typecheck, test, or build based on the changed behavior.
- Block when staged verification is unavailable.
- When a supporting check is unavailable, report the reason and residual risk; continue only when task risk permits.
- Skip only with user approval or when the check is disproportionate to task risk; report the basis and residual risk.
