# Command Usage

## Purpose
Capture best/worst/smell patterns for commit preparation, staged verification, commit message input, and supporting checks.

## Cases (Best / Worst / Smells)

### Case 0. Commit preparation and scope selection
- Best: Use `git status`, `git diff`, and `git diff --staged` to understand scope.
- Worst: Stage and commit without checking the diff.
- Smell: Unexpected files or unrelated changes appear in the commit.

### Case 1. Commit message input method
- Best: Use a prepared commit message file or stdin method that preserves exact newlines and can be reviewed before execution.
- Acceptable: Use `git commit -F -` with a heredoc when command approval and shell quoting are clear.
- Worst: Use multiple `-m` flags or `\n` escapes that introduce extra blank lines.
- Smell: The commit body shows unexpected empty lines or literal `\n` sequences.
- Example:
```
git commit -F - <<'EOF'
type: subject

- change 1
- change 2
EOF
```

### Case 2. Staged diff verification
- Best: Run `git diff --staged` right before commit.
- Worst: Commit without checking what is staged.
- Smell: The commit includes unintended files or partial changes.

### Case 3. Mixed changes handling
- Best: Split unrelated changes into separate commits using partial staging.
- Worst: Bundle unrelated changes in a single commit.
- Smell: Commit messages that mention multiple unrelated topics.

### Case 4. Supporting checks before commit
- Best: Run the narrowest supporting check when staged verification alone does not cover the risk.
- Acceptable: If a supporting check is unavailable or intentionally skipped, record the reason and residual risk before commit.
- Worst: Skip staged verification or supporting checks without reporting the skip.
- Smell: Frequent post-commit CI failures or quick follow-up fix commits.
