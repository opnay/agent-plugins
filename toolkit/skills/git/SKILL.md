---
name: git
description: Run task-scoped Git commit, branch, force-create, and push workflows with working-tree, ref, alias, and remote verification. Use for commits, branch creation or switching, explicit git switch -C, upstream setup, current-branch or explicit refspec pushes, and partial Git workflow recovery; do not activate for incidental read-only Git inspection.
---

# Git Workflow

## Boundary

- Own commit, branch, and push as selectable workflow steps. Combine only user-authorized steps.
- Read repository instructions first. Keep commit, branch, push, installation, and release authority separate.
- Do not infer push, force, destructive cleanup, version bumps, GitHub work, or publishing from a commit request.
- Do not default to history rewriting, hook bypass, forced branch deletion, force push, `reset --hard`, `--discard-changes`, or unrelated working-tree cleanup.
- `git-codex` owns only message-file lifecycle and commit outcome detection. It never stages, amends, pushes, changes branches, bypasses hooks, or judges message meaning.

## Preflight

Run only the checks needed for the requested steps:

```sh
git status --short --branch
git branch --show-current
git branch -vv
git remote -v
git diff
git diff --staged
git config --show-origin --get-regexp '^alias\.'
```

Inspect an alias before using it. Treat every expanded side effect as a separate mutation requiring authority. Do not use a broad alias because its name sounds relevant.

## Commit

Keep one related change unit per commit. Stage only task-owned paths, then inspect the exact staged result:

```sh
git add -- <path>...
git status --short
git diff --staged
```

Block the commit when staged scope is unavailable or differs from the intended change. When the diff does not cover material risk, run the narrowest deterministic check:

- Docs: staged readback, formatting, or whitespace.
- Code: relevant lint, typecheck, test, or build.

Distinguish passed, failed, unavailable, and skipped checks. A skip needs user approval or a proportionality reason; report its residual risk.

### Commit message

Apply a repository convention unless a higher-priority instruction changes it. Otherwise use `type: detailed subject`, subject under 120 characters, and a bullet body after one blank line. Select the most specific type: `feat`, `fix`, `refactor`, `docs`, `test`, `perf`, `style`, `build`, `ci`, or `chore`.

Keep the message specific to staged scope. Do not include literal `\n`, shell syntax, heredoc delimiters, quoting wrappers, unnecessary blank lines, or hidden skipped checks and residual risk.

### Preferred `git-codex` lifecycle

Use the bundled command only when all checks pass:

1. `alias.codex` is absent.
2. The Git external-command dispatch target is the same canonical executable returned by `git codex install --check`.
3. `git codex --version` is `toolkit-git-codex <current Toolkit manifest version>`.

Do not install or update it without separate authority. `git codex install` copies its bundled platform binary to `~/.local/bin/git-codex`; it never changes PATH, shell profiles, or Git config. Use `install --check` for read-only verification and `install --force` only when replacement is authorized.

When available, run this order:

```text
stage verification
> git codex message create
> write expected message
> read back and compare expected message
> git codex message validate <exact-file>
> final staged verification
> git codex commit <exact-file>
> inspect stored full message
```

`message create` prints one absolute `toolkit-git-message.*` path in the OS temp directory with mode `0600`. Record its device and inode before writing. `message validate` checks the canonical temp parent, private user-owned regular-file safety, UTF-8, NUL, lone CR, literal `\n`, subject, and body delimiter. It never edits or deletes the file.

If write, readback, or validation fails, clean only that exact file when its recorded identity and safety conditions still match; otherwise preserve it and report the path. Block the commit.

`git codex commit <file>` reruns mechanical validation, records HEAD, invokes only `git commit -F <file>`, confirms the new HEAD, and deletes the same safe file on success. Interpret outcomes as follows:

- `0`: commit confirmed and file deleted.
- `1`: definitive precondition, validation, or commit failure. Read final `git-codex: commit_result: reason=<code> attempted=<true|false>` before deciding cleanup or recovery.
- `2`: commit confirmed; cleanup failed. Check HEAD and resolve only the remaining file.
- `3`: commit was attempted but its result is unknown. Re-observe HEAD and the message before any retry or manual fallback.

For status `1`, clean a `validation_failed` file only when identity remains safe; preserve `head_unavailable` and `commit_failed` files. If `attempted=true`, inspect state before any fallback.

### Manual fallback

If `git-codex` is unavailable, conflicting, stale, or unsupported, preserve the same lifecycle with a dedicated OS-temp message file. On write/readback/validation failure, perform identity-safe cleanup. On a failed `git commit -F <exact-file>`, preserve the message file for recovery. Do not use heredocs, `git commit -F -`, multiple `-m` arguments, or a combined shell script.

After a confirmed commit and message-file cleanup, verify the stored message:

```sh
git log -1 --format='%H%n%B'
git status --short --branch
```

Compare the full stored message with expected content and applicable convention. Report a mismatch without automatic amend, reset, or rollback.

## Branches and push

Create a normal branch from an explicit start point:

```sh
git switch -c <branch> <start-point>
```

Switch to an existing local branch without changing its ref:

```sh
git switch <branch>
```

`git switch -C <branch> <start-point>` is force-create only. Before it, inspect working tree, current ref, target ref, start point, and `git worktree list --porcelain`. It does not authorize discard changes, branch deletion, or force push.

Before push, resolve remote, local source, remote destination, and upstream separately:

```sh
git push -u <remote> <branch>
git push <remote> <local-source>:<remote-destination>
```

`git push origin wip:main` maps local `wip` to remote `main`; never infer that mapping from a general push request. After an authorized push, compare the requested local and remote refs without repeating the mutation.

## Recovery

Preserve completed steps and stop additional mutation after a failure. Re-observe the current branch, working tree, HEAD, upstream, and required remote refs before retrying. Do not transform a rejection into force push.

Read [references/branch-conventions.md](references/branch-conventions.md) for policy-sensitive prefixes such as `codex/` or `jira/prja-000`, derived branch names, or force-create naming conflicts.

Read [references/recovery.md](references/recovery.md) when switching or force-create is blocked, HEAD is detached, a commit is interrupted or uncertain, an alias partially executes, authentication fails, push is rejected, or refs diverge.
