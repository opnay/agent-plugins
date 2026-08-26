---
name: git-committer
description: Finalize a task-scoped git commit with staged verification, a mandatory message-file gate, cleanup, and post-commit confirmation. Use when commit is part of the requested work, including a broader PR workflow; this skill does not perform push, PR, release, publish, or version bump.
---

# Git Committer

## Boundary

- Own task-scoped commit finalization when commit is part of the requested work, directly or within a broader workflow.
- Do not add a separate commit-specific approval or confirmation gate.
- Own only the commit step; do not perform push, PR, release, publish, or version bump.
- Do not own implementation, readiness judgment, unrelated change removal, or general cleanup.

## Commit Preparation And Verification

1. Confirm the project-specific commit preparation step is complete.
2. Select the intended commit scope with `git status`, `git diff`, and `git diff --staged`; inspect relevant unstaged and untracked state.
3. Split or restage when the scope contains unrelated, unexpected, or partial changes.
4. Immediately before commit, run these staged checks separately:

   ```sh
   git status --short
   ```

   ```sh
   git diff --staged
   ```

5. Block the commit when staged verification is unavailable or differs from the selected scope.
6. When staged verification does not cover the risk, run the narrowest deterministic supporting check:
   - Docs: read back the staged diff and check formatting or whitespace.
   - Code: choose lint, typecheck, test, or build based on the changed behavior.
7. Fix and rerun failures, or block with the failed check and reason.
8. If a supporting check is unavailable, report the reason and residual risk; continue only when task risk permits.
9. Skip a supporting check only with user approval or when it is disproportionate to task risk; report the basis and residual risk. A skip is not a pass.

## Commit Message

- Use `type: detailed subject`; keep the subject under 120 characters and specific to the staged scope.
- Choose the most specific type:
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
- After one blank line, add a bullet body covering concrete changes and verification evidence.
- Keep skipped verification or residual risk visible in the body or final report.
- Avoid vague subjects, unrelated scope, literal `\n`, extra blank lines, shell syntax, or delimiter text.

Example:

```text
docs: clarify deployment prerequisites

- document required runtime and environment configuration
- verify the staged Markdown diff and whitespace
```

## Message File Gate

Run each step separately. Use only the exact path returned by this gate's fixed allocator; never use a user-provided or arbitrary stored path.

1. If a forced interruption preserved an allocator path, confirm its task-state provenance and fixed template, then run the Cleanup Gate before a new attempt.
2. Allocate one message file:

   ```sh
   mktemp /tmp/git-committer-message.XXXXXX
   ```

   Proceed only when the command exits zero and returns exactly one path matching the template.
   - Nonzero without a path: block; no cleanup target exists.
   - Nonzero with exactly one matching path from the current invocation: never use it for commit; run the Cleanup Gate.
   - Any other output: report it and the residual risk without deleting anything, then block.
3. Preserve the exact successful allocator path in task state.
4. Write only the commit message with a filesystem write or edit tool. Do not use shell multiline input or redirection. On failure, run the Cleanup Gate and block.
5. Read the file back:

   ```sh
   sed -n '1,200p' '/tmp/git-committer-message.EXACT'
   ```

   Verify the subject, blank line, bullet body, verification evidence, and absence of unintended shell text or escapes. On failure or mismatch, run the Cleanup Gate and block.
6. Recheck `git status --short` and `git diff --staged` separately. If unavailable or the scope changed, run the Cleanup Gate and block.
7. Commit from the exact allocated path as a separate command:

   ```sh
   git commit -F '/tmp/git-committer-message.EXACT'
   ```

8. Run the Cleanup Gate immediately after the commit attempt, including a failed commit, or before any controllable stop after allocation.
9. If forced interruption prevents cleanup, preserve the exact allocator path in task state for step 1 on resume.

Never use Bash heredoc/EOF, here-string, command substitution, `git commit -F -`, multiple `-m` arguments, or a combined shell script to construct and submit the message.

## Cleanup Gate

1. Accept only the fixed-template exact path returned by the current allocator invocation or preserved in task state after forced interruption.
2. Run once:

   ```sh
   unlink '/tmp/git-committer-message.EXACT'
   ```

3. Exit zero completes cleanup. On failure, report the exact path and residual risk; do not undo a successful commit.

## Commit Granularity

- Keep one commit to one related change unit; split unrelated changes before staging.
- A package update may include the dependency upgrade and only its required fixes.
- For a monorepo login feature, keep API design, database design, and schema application independently reviewable; keep later typo or missing-index fixes separate.

## Reporting

- Distinguish passed, failed, skipped, and unavailable verification.
- Report commit and cleanup results separately.
- After success, confirm the latest commit hash, subject, message shape, and working tree status.
- Report remaining unstaged, untracked, or unrelated changes.
- Do not report push, PR, release, publish, or version bump as completed by this skill.
