# Command Usage

## Purpose
Capture best/worst/smell patterns for git commit commands and related checks.

## Cases (Best / Worst / Smells)

### Case 0. Review changes before staging
- Best: Use `git status`, `git diff`, and `git diff --staged` to understand scope.
- Worst: Stage and commit without checking the diff.
- Smell: Unexpected files or unrelated changes appear in the commit.

### Case 1. Commit message input method
- Best: Use `git commit -F -` with a heredoc for multi-line bodies.
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

### Case 4. Manual CI before commit
- Best: Run relevant `typecheck/lint/test` scripts and keep results green.
- Worst: Skip checks and rely on CI to catch issues later.
- Smell: Frequent post-commit CI failures or quick follow-up fix commits.
