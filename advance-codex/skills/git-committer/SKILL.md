---
name: git-committer
description: Draft and review git commit messages and commit workflows with colon-separated type, detailed subject/body, 120-character limit, and manual CI (typecheck/prettier/etc.) before commit. Use when Codex needs to finalize its change into a task-scoped commit with strong commit preparation and quality checks.
---

# Commit workflow

- Review changes
  - Understand scope and split mixed changes into separate commits
  - Use `references/command-usage.md` for concrete command patterns
- Final staged verification
  - Verify staged changes right before committing
  - Use `references/command-usage.md` for the exact command
- Run manual CI
  - Run relevant scripts and keep results green
  - Fix failures and rerun until green
  - Before committing, verify the output is complete and correct for the task
  - Use `references/usage-examples.md` for scenario-based verification examples
- Write commit message
  - Use format: `type: detailed subject`
  - Keep the first line under 120 characters and specific
  - Always write a body and list changes as bullet points
  - Avoid vague wording
  - Use `references/usage-examples.md` for commit type and message examples
- Commit and confirm
  - Use the command guidance in `references/command-usage.md` for commit message input
  - Confirm with `git log -1` that the format is correct

# Commit granularity

- Commit in small, task-scoped units.
- If a task causes follow-up fixes, commit them separately in order.
- Use `references/usage-examples.md` for concrete sequencing examples.

# References

Use these documents for command patterns and usage examples.

- `references/command-usage.md` - Command usage patterns for commit-related workflows.
- `references/usage-examples.md` - Commit types, message examples, granularity, and verification scenarios.
