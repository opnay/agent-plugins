# Baton Relay Dev

`baton-relay-dev` is a development plugin for worktree-based subagent orchestration.
It keeps the main agent in a manager role: split the work, create isolated git worktrees, dispatch fresh subagents, require each subagent to commit and rebase before handoff, integrate prepared commits, and clean up worktrees.

## Skills

- `manager`: decomposes a task into worktree-safe slices, assigns each slice to a fresh subagent, controls subagent lifecycle, verifies commit/rebase handoff, integrates prepared commits, and records cleanup expectations.

## Use

Use this plugin when one task is large enough to benefit from isolated subagent execution across git worktrees.
Do not use it for a small single edit, direct code review, release publication, PR creation, or any action that needs approval outside the current task boundary.
