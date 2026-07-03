# Implement

Use this reference when the task requires code changes.

## Sequence

1. Read the relevant files and callers before editing.
2. State the required behavior in one sentence.
3. Find the owning boundary.
4. Apply the decision ladder from `SKILL.md`.
5. Edit the smallest surface that completes the behavior.
6. Run the narrowest useful verification.
7. Report changed scope, skipped expansion, verification, and remaining risk.

## Smallest Safe Change

A change is small enough when it:

- touches the owner of the behavior
- keeps caller contracts intact or updates all affected callers
- handles the relevant failure path
- avoids speculative options, wrappers, factories, hooks, config, or layers
- uses local patterns before new structure
- leaves unrelated cleanup out

Small is not partial. Include the normal path, required edge path, and evidence that the contract works.

## Reuse Check

Before writing new code, search for:

- existing helper or utility
- existing type, schema, parser, formatter, validator, or error shape
- standard library function
- framework/platform feature
- installed dependency already used for the same job
- nearby test style or fixture pattern

If a local pattern is flawed, fix the owned flaw instead of cloning it.

## Verification

Match checks to risk.

- Trivial expression or copy change: targeted command or existing check may be enough.
- Branch, loop, parser, transform, permission, money, data loss, auth, async, time, cache, public API: add or run executable checks.
- Shared or UI workflow: run representative integration, E2E, or manual evidence when available.

Do not add a new test framework or broad harness for a narrow change unless the project already requires it.
