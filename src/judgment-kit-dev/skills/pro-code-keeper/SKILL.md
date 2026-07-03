---
name: pro-code-keeper
description: Apply lean senior developer judgment to coding, debugging, refactoring, simplification review, repository overengineering audits, dependency reduction, and lean debt tracking. Use when a task needs the smallest safe code change after understanding real code flow, reuse of existing code, standard library or platform features, deliberate avoidance of unnecessary abstraction, or review findings tagged delete, stdlib, native, yagni, and shrink. lean senior dev, smallest safe change, overengineering review, simplify code, delete code, reduce dependencies, lean debt, lean comments
---

# Pro Code Keeper

Use this skill as lean senior developer mode: understand the problem and existing flow first, then make or recommend the smallest safe change. Lean means less unnecessary code, not less investigation.

## Grounding

Before writing code, understand the real path.

- Read the relevant files and follow the actual call flow.
- When fixing a bug, look for the shared cause instead of patching only the symptom.
- When changing a function, inspect its callers and the contract they depend on.
- Keep the edit at the owning boundary; a small change in the wrong layer is still a second bug.
- If the user asked only for review, do not edit code.

## Decision Ladder

Apply this order before adding code.

1. Do not build the feature if it is not actually needed.
2. Reuse an existing helper, utility, type, or local pattern when it already fits.
3. Use the standard library when it solves the problem.
4. Use language, browser, OS, database, or framework defaults when they solve it.
5. Use an already installed dependency before adding a new one.
6. Use one clear line when one line is enough.
7. Then write the smallest working code.

Prefer deletion over addition, boring code over clever code, fewer touched files over more files, shorter diffs over longer diffs, and existing patterns over new structure. If two choices are equal size, choose the one safer for edge cases.

## Never Shrink

Do not simplify away:

- validation at trust boundaries: external input, permissions, network, files, payments, accounts, or personal data
- error handling that prevents data loss
- security
- accessibility
- features the user explicitly requested
- correction space for hardware, time, sensors, floating point, concurrency, and other real-world uncertainty
- problem understanding and code-flow inspection

If the user explicitly asks for the full version, implement the full version without arguing, while keeping the design understandable.

## Avoid

Do not create speculative structure.

- No abstraction that is not needed now.
- No interface with only one implementation.
- No layer with only one caller.
- No future-only config, factory, extension point, hook, or wrapper.
- No reimplementation of standard library or platform behavior.
- No casual new dependency.
- No explanation that hides complexity instead of removing it.
- No extra files unless they clearly reduce current risk or complexity.
- No heavy test harness for a trivial change.

## Lean Comments

Use a `lean:` comment only when a deliberately simple choice has a known limit worth tracking.

A good `lean:` comment states the current simplification, its limit, and the upgrade trigger.

Examples:

```ts
// lean: O(n) scan is fine under 1k items; switch to indexed lookup if this becomes hot.
```

```py
# lean: process-wide lock; use per-account locks if contention shows up.
```

Do not use `lean:` as an excuse for incomplete behavior.

## Testing

Match verification to risk.

- For a trivial one-line change, a new test is optional.
- For branches, loops, parsers, money, security, permissions, data transforms, or non-obvious logic, leave at least one executable check.
- Prefer one focused unit test or a small assert-style check.
- Do not add a new test framework, large fixture, or broad test suite unless the user asks or the local project already requires it.

## Review Mode

When the user asks for overengineering review, simplification review, removable code, dependency reduction, code shrinking, or technical-debt cleanup, report findings without editing.

Use these tags:

- `delete`: dead code, unused flexibility, guessed features
- `stdlib`: custom code replaceable by the standard library
- `native`: code or dependency replaceable by platform behavior
- `yagni`: abstraction, setting, layer, or extension point not needed yet
- `shrink`: same behavior expressible more simply

Write each finding in one line with location, target, replacement, and reason. Security, correctness, and performance risks are separate risks, not deletion targets. Minimal smoke tests and self-checks are not unnecessary code.

## Repository Audit Mode

When asked to audit a whole repository for overengineering, unnecessary dependencies, removable code, or technical debt:

- Scan the tree, excluding build outputs, dependency folders, `.git`, and generated artifacts.
- Rank items by reduction impact.
- Classify each item as `delete`, `stdlib`, `native`, `yagni`, or `shrink`.
- Report as: where, what, why, replacement.
- Do not edit. Apply changes only after the user asks.

## Lean Debt Ledger

When asked to list simplification comments, deferred improvements, `lean:` notes, or future expansion points:

- Search for `lean:` comments.
- Report: file, line, current simplification, known limit, upgrade condition, upgrade path.
- Mark comments without a trigger as `no-trigger`.
- Default to read-only reporting. Save a document only when the user asks.

## Reporting

For code changes, act first and explain briefly:

- what changed
- what was intentionally not built
- what condition would justify expansion later

For detailed analysis, design docs, or reports, expand only as much as the user requested.
