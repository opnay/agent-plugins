---
name: pro-code-keeper
description: Apply lean senior developer judgment to coding, debugging, root-cause fixes, refactoring, simplification review, repository overengineering audits, dependency checks, refactor shrinking, and lean debt tracking. Use when a task needs the smallest safe code change after understanding real code flow, reuse of existing code, standard library or platform features, deliberate avoidance of unnecessary abstraction, or review findings tagged delete, stdlib, native, yagni, and shrink. lean senior dev, smallest safe change, root cause fix, overengineering review, simplify code, delete code, reduce dependencies, refactor shrink, lean debt, lean comments
---

# Pro Code Keeper

Use this skill for lean code stewardship: understand the real code path, then make or recommend the smallest safe change. Lean means less unnecessary code, not less investigation.

## Core Rule

Follow this order before adding code:

1. Do not build behavior that is not required.
2. Reuse an existing helper, type, module, pattern, or test shape when it fits.
3. Use the standard library when it solves the problem.
4. Use language, browser, OS, database, or framework defaults when they solve it.
5. Use an already installed dependency before adding one.
6. Use one clear line when one line is enough.
7. Then write the smallest working code.

Prefer deletion over addition, boring code over clever code, fewer touched files over more files, and existing ownership boundaries over nearby convenience.

## Reference Loading

Always start with `references/init.md`. Also load `references/safety-boundaries.md` before any deletion, simplification, implementation, dependency, or review judgment.

Load only the branch files needed for the user request:

- Implementation or small code change: `references/implement.md`
- Bug, debug, failing test, regression, or root-cause fix: `references/root-cause-fix.md` and `references/implement.md`
- Overengineering, simplification, removable-code, or dependency review: `references/code-review.md`
- Whole-repository audit: `references/repo-audit.md` and `references/impact-scoreboard.md`
- Dependency add/remove/replace decision: `references/dependency-check.md`
- Refactor, shrink, cleanup, or behavior-preserving simplification: `references/refactor-shrink.md`
- `lean:`, `ponytail:`, deferred simplification, or future expansion point listing: `references/debt-ledger.md`
- Final report or long review output: `references/output-style.md`

Use `examples/` only when you need a short calibration example:

- `examples/native-before-dependency.md`
- `examples/reuse-before-rewrite.md`
- `examples/delete-before-add.md`

## User Argument Routing

Treat arguments after the skill name as intent hints, not new skill names.

- `review`: read `code-review.md`.
- `audit`: read `repo-audit.md` and `impact-scoreboard.md`.
- `fix`, `debug`, `root cause`: read `root-cause-fix.md`.
- `dependency`, `package`, `library`: read `dependency-check.md`.
- `refactor`, `shrink`, `cleanup`: read `refactor-shrink.md`.
- `debt`, `lean comments`, `future points`: read `debt-ledger.md`.

Natural-language requests with the same meaning route the same way.

## Stop Conditions

- If the user asked only for review or audit, do not edit files.
- If the smallest change would remove security, validation, accessibility, data-loss protection, or a requested feature, stop and choose the next-smallest safe option.
- If local evidence cannot identify the owning boundary, inspect more or ask for the missing product or risk decision.
