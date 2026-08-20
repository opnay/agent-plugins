---
name: git-committer
description: Finalize a task-scoped git commit with staged verification, a mandatory message-file gate, cleanup, and post-commit confirmation. Use when commit is part of the requested work, including a broader PR workflow; this skill does not perform push, PR, release, publish, or version bump.
---

# Git Committer

## Boundary

- Own task-scoped commit finalization when commit is part of the requested work, directly or within a broader workflow.
- Do not add a separate commit-specific approval or confirmation gate.
- Own only the commit step; do not perform push, PR, release, publish, or version bump.
- Do not own implementation, general cleanup, readiness judgment, or unrelated change removal.

## Commit Preparation

1. Confirm the project-specific commit preparation step is complete.
2. Select the intended commit scope.
3. Review `git status`, `git diff --staged`, and relevant unstaged or untracked state.
4. Split, restage, or stop if the staged diff contains unrelated or unexpected changes.

## Commit Execution

1. Run staged verification with `git status` and `git diff --staged`.
2. Run the narrowest deterministic supporting check when staged verification does not cover the risk.
3. Fix and rerun failures, or block the commit with the failed check and reason.
4. If staged verification is unavailable, record the reason and residual risk, then block the commit. If a supporting check is unavailable, record the reason and residual risk, and continue only when task risk permits.
5. Skip a supporting check only when the user approved the skip or the check is disproportionate to task risk; record the basis and residual risk.
6. Draft a message using `type: detailed subject`, one blank line, and a bullet body covering concrete changes and verification evidence.
7. Complete the Message File Gate.
8. After a successful commit and cleanup attempt, confirm the latest commit hash, subject, message shape, and working tree status.

## Message File Gate

1. On resume, run the Cleanup Gate for a leftover path only when task state proves it came from this gate's allocator and matches the fixed template.
2. Run one trusted temporary-file allocator with a fixed safe template.
   - Proceed only when it exits zero, returns exactly one path, that path matches the template, and it names a regular non-symlink file.
   - On nonzero exit without a path, block the commit; no cleanup target exists.
   - On nonzero exit with a path, never use it for commit. Run the Cleanup Gate only if provenance, template, and file type can be verified; otherwise report the path and residual risk without deleting it.
3. Preserve the exact successful allocator path in task state.
4. Write only the commit message with a filesystem write or edit tool. Do not construct the content with shell multiline input or redirection. On write failure, run the Cleanup Gate and block the commit.
5. Read the file back and verify the subject, blank line, bullet body, verification evidence, and absence of unintended shell text or escapes. On readback failure or mismatch, run the Cleanup Gate and block the commit.
6. Recheck final `git status` and `git diff --staged`. If either check is unavailable or the staged scope differs from the selected scope, run the Cleanup Gate and block the commit.
7. Run `git commit -F <exact-allocated-file-path>` as a separate command.
8. Run the Cleanup Gate immediately after the commit attempt, or before any controllable stop after allocation.
9. If forced interruption prevents cleanup, preserve the exact allocator-returned path in task state. On resume, perform step 1 before any new commit attempt.
10. Report the commit result and cleanup result separately. On cleanup failure, report the remaining exact path and residual risk; do not undo a successful commit.

Never use Bash heredoc/EOF, here-string, command substitution, `git commit -F -`, multiple `-m` arguments, or a combined shell script to construct and submit the message.

## Cleanup Gate

1. Accept only the exact allocator-returned path whose provenance and fixed template are verified. Never accept an arbitrary or user-provided path.
2. If separate file and symlink absence checks both pass, cleanup is complete.
3. Otherwise, immediately before deletion, verify the path still names the expected regular non-symlink file. If either check fails, do not delete it; report the exact path and residual risk.
4. Delete only that exact path.
5. Re-run separate file and symlink absence checks. Both must pass.

## Message Rules

- Keep the subject under 120 characters and describe the staged scope precisely.
- Choose the most specific type: `feat`, `fix`, `refactor`, `docs`, `test`, `perf`, `style`, `build`, `ci`, or `chore`.
- Always include bullet body items for concrete changes and verification evidence.
- Avoid unrelated scope, literal `\n` escapes, accidental blank-line noise, shell syntax, and delimiter text.

## Reporting

- Distinguish passed, failed, skipped, and unavailable verification.
- Report the commit hash, subject, and message shape only after confirming them.
- Report message-file cleanup and remaining working tree changes.
- Do not report push, PR, release, publish, or version bump as completed by this skill.

## References

- Read `references/command-usage.md` before staged verification, staging changes, commit command construction, message-file handling, or final diff checks.
- Read `references/usage-examples.md` when choosing commit type, granularity, message shape, or verification scope.
