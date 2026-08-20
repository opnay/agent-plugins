# Command Usage

## Purpose

Define safe command patterns for commit scope, staged verification, message-file handling, supporting checks, and cleanup.

## Message File Sequence

Run each step separately. Use only the exact path returned by the fixed allocator command; never substitute a user-provided or arbitrary stored path.

1. If an interrupted attempt recorded an allocator path, validate its task-state provenance and fixed template, then run the Cleanup Gate before continuing.
2. Allocate one message file:

   ```sh
   mktemp /tmp/git-committer-message.XXXXXX
   ```

   Treat allocation as successful only when the command exits zero, returns exactly one path matching `/tmp/git-committer-message.` plus the allocator suffix, and these separate checks pass:

   ```sh
   test -f '/tmp/git-committer-message.EXACT'
   ```

   ```sh
   test ! -L '/tmp/git-committer-message.EXACT'
   ```

   On nonzero exit without a path, block the commit and run no cleanup. On nonzero exit with a path, never use it for commit; run the Cleanup Gate only after provenance, template, and file type are verified. Otherwise report the path and residual risk without deleting it.

3. Preserve the exact successful allocator path in task state.
4. Use a filesystem write or edit tool to place only the commit message in that exact file. Do not use shell redirection, heredoc/EOF, here-string, command substitution, or stdin. On write failure, run the Cleanup Gate and block the commit.
5. Read the file back:

   ```sh
   sed -n '1,200p' '/tmp/git-committer-message.EXACT'
   ```

   On readback failure or message mismatch, run the Cleanup Gate and block the commit.

6. Recheck staged state with separate commands:

   ```sh
   git status --short
   ```

   ```sh
   git diff --staged
   ```

   If either check is unavailable or the staged scope differs from the selected scope, run the Cleanup Gate and block the commit.

7. Commit from the exact allocated file:

   ```sh
   git commit -F '/tmp/git-committer-message.EXACT'
   ```

8. Run the Cleanup Gate after the commit command returns, including a failed commit. Also run it before any controllable stop after allocation.

If forced interruption prevents cleanup, preserve the exact allocator-returned path and start the next attempt at step 1. Report commit and cleanup results separately. Never combine this sequence into one shell script.

## Cleanup Gate

1. Accept only an exact allocator-returned path with verified task-state provenance and fixed template.
2. Check whether the path is already absent:

   ```sh
   test ! -e '/tmp/git-committer-message.EXACT'
   ```

   ```sh
   test ! -L '/tmp/git-committer-message.EXACT'
   ```

   If both pass, cleanup is complete.

3. Otherwise, immediately before deletion, verify the expected regular non-symlink file:

   ```sh
   test -f '/tmp/git-committer-message.EXACT'
   ```

   ```sh
   test ! -L '/tmp/git-committer-message.EXACT'
   ```

   If either fails, do not delete the path; report the exact path and residual risk.

4. Delete only the verified exact path:

   ```sh
   unlink '/tmp/git-committer-message.EXACT'
   ```

5. Repeat both absence checks from step 2. Both must pass. On any cleanup failure, report the remaining exact path and residual risk; do not undo a successful commit.

## Cases

### Scope Selection

- Best: Use `git status`, `git diff`, and `git diff --staged` to understand scope.
- Worst: Stage and commit without checking the diff.
- Smell: Unexpected files or unrelated changes appear in the commit.

### Message Input

- Best: Allocate, verify, write, read back, submit with `git commit -F <file>`, then run the Cleanup Gate.
- Worst: Use a user-provided path, `git commit -F -`, heredoc/EOF, here-string, command substitution, multiple `-m` arguments, or one combined shell script.
- Smell: The body lacks verification evidence or contains shell text, delimiter markers, literal `\n`, or unexpected blank lines.

### Staged Verification

- Best: Run `git status` and `git diff --staged` immediately before `git commit -F <file>`.
- Worst: Commit without checking what is staged.
- Smell: The commit includes unintended files or partial changes.

### Mixed Changes

- Best: Split unrelated changes with partial staging.
- Worst: Bundle unrelated changes in one commit.
- Smell: The message mentions multiple unrelated topics.

### Supporting Checks

- Best: Run the narrowest deterministic supporting check when staged verification does not cover the risk.
- Unavailable staged verification: Record the reason and residual risk, then block the commit.
- Unavailable supporting check: Record the reason and residual risk; continue only when task risk permits.
- Deliberate skip: Require user approval or disproportionate task risk; record the basis and residual risk.
- Worst: Hide unavailable, skipped, or failed checks behind a pass.
- Smell: CI repeatedly fails immediately after commits.
