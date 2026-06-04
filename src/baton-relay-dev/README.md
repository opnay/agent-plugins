# Baton Relay Dev

`baton-relay-dev` is a development plugin for worktree-based subagent orchestration.
It keeps the main agent in a manager role: split the work, create isolated git worktrees, dispatch fresh subagents, require each subagent to commit and rebase before handoff, integrate prepared commits, and clean up worktrees.

## Skills

- `manager`: decomposes a task into worktree-safe slices, assigns each slice to a fresh subagent, controls subagent lifecycle, verifies commit/rebase handoff, integrates prepared commits, and records cleanup expectations.

## Use

Use this plugin when the main agent should manage work through a plan-first subagent relay, even for a single small task.
Do not use it to bypass approval for release publication, PR creation, version bumps, destructive work, or external actions.
