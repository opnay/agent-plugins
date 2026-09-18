---
name: git
description: Run task-scoped Git commit, git-codex installation, branch, force-create, push, and recovery workflows with working-tree, ref, alias, and remote verification. Use for commits, git-codex setup, branch creation or switching, explicit git switch -C, upstream setup, current-branch or explicit refspec pushes, and partial Git workflow recovery; do not activate for incidental read-only Git inspection.
---

# Git Workflow

## Scope and workflow selection

Read repository instructions first. Select and combine only authorized workflows:

- **Commit**: stage one task-owned change unit, verify it, create a file-based message, commit, and verify the stored message.
- **Installation**: install or update `git-codex` only on an explicit user request.
- **Branch and push**: create, switch, force-create, set upstream, or push the requested refs.
- **Recovery**: inspect failed or uncertain operations while preserving completed steps.

Keep commit, installation, branch, push, and release authority separate. A commit request does not authorize push, force, destructive cleanup, version bumps, GitHub work, or publishing. Do not default to history rewriting, hook bypass, forced branch deletion, force push, `reset --hard`, or `--discard-changes`.

`git-codex` owns message-file lifecycle and commit outcome detection only. It never stages, amends, pushes, changes branches, bypasses hooks, or judges message meaning. Apply this skill without relying on sibling skills or development-only documents.

## Alias checks before execution

Before using a Git alias in any workflow, inspect its definition and compare every expanded side effect with the authorized task:

```sh
git config --show-origin --get-regexp '^alias\.'
```

Before every `git codex` invocation, including `message create` for availability and installation verification, run this read-only check in the same repository and configuration context:

```sh
git config --show-origin --get-all alias.codex
```

Exit status `1` means no alias is defined. If a definition is returned, inspect its expansion and side effects, report an alias conflict, and stop before invoking `git codex`. Any other lookup failure also blocks the invocation. Preserve Git config.

Git can dispatch to `alias.codex` when the external command is missing. Never execute that alias as an availability probe or classify an alias conflict as an installation-required result. After this check passes, use `message create` itself to determine lifecycle availability.

## Commit workflow

Inspect the working tree and staged scope:

```sh
git status --short --branch
git diff
git diff --staged
```

Keep one related change unit per commit. Preserve unrelated changes, stage only task-owned paths, and inspect the result:

```sh
git add -- <path>...
git status --short
git diff --staged
```

Block the commit if staged verification is unavailable or the scope differs from the intended change. When the diff does not establish readiness, run the narrowest relevant supporting check: readback, formatting, or whitespace for docs; lint, typecheck, test, or build for code.

Distinguish passed, failed, unavailable, and skipped checks. Fix and rerun failed checks or report the blocker. Continue past an unavailable check only when task risk permits, reporting the reason and residual risk. A skip needs user approval or a proportionality reason; report its residual risk.

### Message content

Follow instruction priority and the current user's explicit override of the same rule; otherwise preserve repository convention. Keep the subject under 120 characters or a stricter repository limit. Without another convention, use `type: detailed subject`, one blank line, and a bullet body explaining changes and verification.

Choose the most specific supported type: `feat`, `fix`, `refactor`, `docs`, `test`, `perf`, `style`, `build`, `ci`, or `chore`. Follow a different applicable type set without inventing types. Describe the staged scope specifically; keep unrelated concerns in separate commits.

Do not include literal `\n`, shell syntax, heredoc delimiters, quoting wrappers, or unnecessary blank lines. Disclose skipped verification and residual risk in the body or final report.

### Message lifecycle

Apply the alias checks before each `git codex` command, then follow this sequence:

```text
stage verification
> git codex message create
> record allocated file identity
> write expected message
> read back and compare expected message
> git codex message validate <exact-file>
> final staged verification
> git codex commit <exact-file>
> inspect stored full message
```

`message create` returns only an absolute `toolkit-git-message.*` path in the OS temp directory with mode `0600`. Record its device and inode before writing. `message validate` checks the canonical temp parent, private user-owned regular-file safety, UTF-8, NUL, lone CR, literal `\n`, subject, and body delimiter. It does not edit or delete failed files.

If writing, readback, or validation fails, block the commit. Clean only the exact allocated file when its recorded identity and current safety conditions match; otherwise preserve it and report its path.

Immediately before committing, inspect `git status --short` and `git diff --staged` again. `git codex commit <exact-file>` repeats mechanical validation, records HEAD, runs only `git commit -F <file>`, and deletes the same safe file only after confirming command success and a new HEAD.

Interpret its status before taking another action:

- `0`: commit confirmed and file deleted.
- `1`: definitive precondition, validation, or commit failure. Read the final `git-codex: commit_result: reason=<code> attempted=<true|false>` diagnostic. Clean `validation_failed` only when identity remains safe; preserve `head_unavailable` and `commit_failed` files. If `attempted=true`, inspect state before recovery or fallback.
- `2`: commit confirmed; cleanup failed. Inspect HEAD, file identity, and stored message, then resolve only the remaining file.
- `3`: commit attempted with an unknown result. Inspect HEAD, file identity, and stored message before any retry or manual fallback.

After a confirmed commit and message-file cleanup, read the stored hash, full message, and local state:

```sh
git log -1 --format='%H%n%B'
git status --short --branch
```

Compare the subject, body, and any hook-added trailers with expected content and applicable convention. Report a mismatch as failed message verification; do not automatically amend, reset, or roll back the commit.

## Installation workflow

Enter only for an explicit `git-codex` installation or update request.

1. Locate the matching bundled `scripts/bin/<os>-<arch>/git-codex` executable.
2. Run that executable's `install` command to install at `~/.local/bin/git-codex`.
3. Apply the alias checks, then run `git codex install --check` to verify the installed target and dispatch.
4. Apply the alias checks, then run `git codex --version` and compare `toolkit-git-codex <version>` with the Toolkit manifest version.

The installer preserves PATH, shell profiles, and Git config. `--force` requires explicit replacement authority. Report installation and verification separately if a later check fails.

## Error and recovery workflow

- If `git codex message create` is unavailable after alias checks pass, report that installation is required and end the current commit lifecycle. A later explicit request selects installation.
- An alias conflict, alias lookup error, or other create failure is not proof of a missing installation. Report the observed cause; read [references/recovery.md](references/recovery.md) when recovery is needed.
- Preserve completed steps after failure. Re-observe current branch, working tree, HEAD, upstream, and required remote refs before another mutation. Apply the status-specific message-file rules above before cleanup or retry.
- Read [references/recovery.md](references/recovery.md) for blocked branch changes, detached HEAD, interrupted or uncertain commits, partial alias execution, authentication failures, rejected or uncertain pushes, or divergent refs.
- A non-fast-forward rejection does not authorize force push. Query an uncertain remote result before repeating a push.

A manual message-file workflow is available when the user selects it or `git-codex` is unsupported. Preserve the same validation, readback, identity-safe cleanup, commit-failure preservation, and stored-message verification contract.

## Branch and push workflow

Create a normal branch from a confirmed start point:

```sh
git switch -c <branch> <start-point>
```

Switch to an existing local branch without resetting its ref:

```sh
git switch <branch>
```

Use `git switch -C <branch> <start-point>` only when the user or owning repository workflow explicitly authorizes force-create. First confirm the exact branch, start point, existing target ref, working tree, and `git worktree list --porcelain`. This authority does not include discarding changes, branch deletion, or force push.

Before pushing, resolve remote URL, local source, remote destination, and upstream separately:

```sh
git branch --show-current
git branch -vv
git remote -v
git push -u <remote> <branch>
git push <remote> <local-source>:<remote-destination>
```

`git push origin wip:main` maps local `wip` to remote `main`. Never infer this mapping from a general push request; repository restrictions and authorization for the exact destination still apply. After mutation, verify local branch or commit state and the required remote ref without repeating the mutation. Report branch, commit, and push outcomes separately.

After an authorized push, compare the source commit with the destination ref:

```sh
git rev-parse <local-source>
git ls-remote --heads <remote> refs/heads/<remote-destination>
```

Read [references/branch-conventions.md](references/branch-conventions.md) for policy-sensitive prefixes such as `codex/` or `jira/prja-000`, derived names, or force-create naming conflicts. Preserve exact user-provided names; a prefix alone does not determine start point, remote destination, or push permission.
