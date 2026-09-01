# Branch Conventions

Read this reference only when a branch name or workflow is governed by a prefix, ticket key, repository policy, user-specific convention, or force-create collision rule.

## Resolve The Owning Rule

Resolve conflicts by instruction priority and explicit scope. Keep a repository default unless the current user clearly overrides that same rule and no higher-priority instruction blocks the override. When no conflict exists, discover rules in this order:

1. Repository instructions such as `AGENTS.md`, contribution guidance, or an owning workflow document.
2. The current user's exact branch, base, remote, and destination request.
3. A reusable convention explicitly documented for the current environment.
4. General Git behavior from `SKILL.md`.

Do not turn a branch name observed in one repository into a universal convention.

## Prefixes Are Conditions, Not Complete Policies

Treat prefixes such as `codex/` and `jira/prja-000` as signals to look for an owning rule. The string alone does not determine:

- the start point;
- whether the branch is local-only or publishable;
- whether an existing branch may be reset with `git switch -C`;
- the push remote or destination;
- whether the branch may be reused, renamed, deleted, or force-updated; or
- the pull-request base and cleanup timing.

Preserve an exact user-provided name, including case and separators. Do not infer that `codex/` means agent-owned or that a `jira/` branch may be derived without a known ticket key.

## Deriving A Branch Name

Derive a name only when the owning rule provides every required input:

- prefix or namespace;
- ticket or task identifier;
- optional slug rules;
- case normalization;
- allowed characters and length; and
- collision behavior.

If the ticket, suffix, base branch, or collision rule is missing or ambiguous, stop before creating the branch and request the missing decision. Do not use a placeholder ticket such as `prja-000` as a real identifier unless the user supplied it as the exact branch name.

## Existing Branches And Collisions

Before create, reuse, reset, or switch:

```sh
git branch --all --verbose --no-abbrev
git status --short --branch
git worktree list --porcelain
```

- Existing local branch: verify its commit and working-tree compatibility before switching.
- Remote-tracking branch only: verify the intended local name and upstream before creating a tracking branch.
- Same name at another commit: do not use `git switch -C`, forced rename, or deletion merely to resolve the collision.
- Explicit force-create: resolve the exact branch and start point, retain the existing branch object ID, verify other worktree ownership, and use `git switch -C` only when resetting that ref is authorized.
- Branch used by another worktree: preserve that worktree and report the conflict.

`git switch -C` authority does not authorize `--force`, `--discard-changes`, branch deletion, or force push.

## Push Mapping

A prefix never implies a remote destination. Resolve the mapping explicitly:

```text
local source -> remote -> remote destination
```

Repository rules override cheatsheet examples. If a repository defines `wip` as local-only, `git push origin wip:main` is prohibited there even though the refspec syntax is valid.

## Durable Policy Placement

Keep repository-specific names and branch lifecycle rules in that repository's owning instructions. Keep cross-repository personal conventions in an explicitly owned user-level policy. This reference only defines how the runtime discovers, prioritizes, and applies those policies.
