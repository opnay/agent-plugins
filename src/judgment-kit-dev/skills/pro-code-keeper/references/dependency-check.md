# Dependency Check

Use this reference when adding, removing, replacing, or reviewing a dependency.

## Default

Do not add a dependency until the standard library, platform, framework, and already installed dependencies have been checked.

## Add Decision

Approve adding a dependency only when most are true:

- the problem is nontrivial and recurring
- native or standard tools are incomplete or risky
- the dependency has an active, narrow, well-understood API
- bundle/runtime/security/license cost is acceptable
- the project already accepts dependencies for similar scope
- tests can cover the integration boundary

Reject or delay when the dependency only saves a few lines, wraps one platform call, adds broad transitive risk, or solves a speculative future need.

## Remove Decision

Prefer removal when:

- usage is tiny or duplicated by native behavior
- the dependency is stale, broad, or security-sensitive
- migration is mechanical and testable
- removing it reduces build, bundle, install, or maintenance cost

Do not remove when it owns complex edge cases, compatibility, accessibility, parsing, crypto, time, i18n, or security unless the replacement is proven.

## Output

Use:

- `decision`: keep, add, remove, replace, or defer
- `native/stdlib option`
- `current usage`
- `cost`
- `risk`
- `migration path`
- `verification`
