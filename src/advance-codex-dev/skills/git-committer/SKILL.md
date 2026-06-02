---
name: git-committer
description: Prepare and execute a task-scoped git commit by separating commit preparation, commit execution authority, staged verification, message preparation, commit creation, and post-commit confirmation. Use when the user asks to finalize work toward commit; do not execute the commit without explicit commit approval, and do not use for push, PR, release, publish, or version bump.
---

# Git Committer

## Entry Gate

- Use this skill when the user asks to prepare or finalize a task-scoped commit.
- Execute the actual commit only after explicit commit approval.
- Do not infer commit approval from commit-readiness, passing checks, handoff, session records, or prior context.
- Treat commit, push, PR, release, publish, and version bump as separate approval-sensitive actions.

## Commit Preparation

1. Confirm the project-specific commit preparation step is complete.
2. Select the intended commit scope.
3. Review `git status` and `git diff --staged`; use `git diff` when unstaged changes may affect scope.
4. Split, restage, or stop if the staged diff contains unrelated or unexpected changes.

## Commit Execution Authority

1. Confirm the user explicitly approved actual commit execution.
2. Confirm the approval applies to the selected staged scope.

## Commit Execution

1. Run staged verification with `git status` and `git diff --staged`.
2. Run the narrowest supporting check when staged verification alone does not cover the risk.
3. Fix failures and rerun, or block the commit with the failed check and reason.
4. If verification is skipped or unavailable, record why and state residual risk.
5. Draft a commit message using `type: detailed subject` plus a bullet body.
6. Recheck the final staged diff immediately before committing.
7. Commit, then confirm latest commit metadata and working tree status.

## Message Rules

- Keep the subject under 120 characters.
- Choose the most specific type: `feat`, `fix`, `refactor`, `docs`, `test`, `perf`, `style`, `build`, `ci`, or `chore`.
- Always include a body with bullets for the concrete changes.
- Include verification evidence or deliberate skips in the report; do not hide skipped checks behind a pass.
- Avoid vague subjects, unrelated scopes, literal `\n` escapes, and accidental blank-line noise.

## References

- Read `references/command-usage.md` before staged verification, staging changes, commit command construction, or final diff checks.
- Read `references/usage-examples.md` when choosing commit type, granularity, message shape, or verification scope.
