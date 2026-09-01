---
name: git
description: Run task-scoped Git commit, branch, force-create, and push workflows with working-tree, ref, alias, and remote verification, plus conditional branch-convention and recovery guidance. Use when the requested outcome includes committing changes, creating, resetting, or switching branches, setting an upstream, or pushing a current branch or explicit refspec; do not activate for incidental read-only Git inspection. Git commit, commit verification, post-commit message verification, commit granularity, commit type, commit subject length, branch creation, git switch -C, force-create branch, branch switch, upstream push, refspec push, branch prefix, push recovery
---

# Git Workflow

## Boundary

- Own commit, branch, and push as selectable steps in one Git workflow. Combine only the steps authorized by the user or an enclosing requested workflow.
- Read repository instructions before applying this general guidance. Resolve conflicts by instruction priority and explicit scope. Keep a repository default unless the current user clearly overrides that same rule and no higher-priority instruction blocks the override.
- Keep commit authorization separate from branch and push authorization. Do not infer push from commit, a remote destination from a local branch name, or force from a rejected update.
- Treat `git switch -C` as explicit branch force-create authority only. Do not extend it to working-tree discard, branch deletion, force push, or another history mutation.
- Do not own implementation readiness, GitHub pull requests, releases, publishing, version bumps, hosting-service APIs, or unrelated working-tree cleanup.
- Do not default to `reset --hard`, `--force`, `--discard-changes`, forced branch deletion, force push, history rewriting, or bypassing hooks and authentication.

## Preflight

Run only the checks needed to resolve the requested steps:

```sh
git status --short --branch
git branch --show-current
git branch -vv
git remote -v
git diff
git diff --staged
```

Before using a Git alias, inspect its exact definition and source:

```sh
git config --show-origin --get-regexp '^alias\.'
```

Treat an alias as its expanded sequence. If it stages, commits, and pushes, each mutation must be authorized and its intermediate state must be verified. Do not use a broad alias merely because its name sounds relevant.

## Workflow

Compose the requested steps without turning them into exclusive modes:

- Commit: scope > stage > verify staged diff > run necessary supporting check > commit > cleanup > verify stored message.
- Branch then commit: resolve start point > create, force-create, or switch > commit.
- Push: resolve remote, local source, and remote destination > push > verify remote ref.
- Branch then commit then push: preserve the authority and verification gate of every step.

If one step fails, preserve completed earlier steps, stop additional mutation, and read the recovery reference before retrying.

## Cheatsheet

### Inspect state and help

```sh
git status --short --branch
git branch --show-current
git branch -vv
git remote -v
git rev-parse HEAD
git <command> -h
```

Use installed command help as the syntax source of truth. Read-only inspection does not authorize a following mutation.

### Create or switch branches

```sh
# Create and switch from an explicit start point
git switch -c <branch> <start-point>

# Switch to an existing local branch
git switch <branch>

# Inspect local and remote-tracking branches
git branch --all --verbose --no-abbrev
```

Resolve the exact branch name and start point first. Do not substitute `--force`, `--discard-changes`, or `--ignore-other-worktrees` for an unresolved conflict. If a policy-sensitive prefix such as `codex/` or `jira/prja-000` is involved, read the branch-conventions reference before creation, reuse, reset, rename, or push.

### Force-create or reset and switch

`-C` is `--force-create`: it creates the branch when absent or resets an existing branch ref to the start point, then switches to it. Use it only when that exact force-create outcome is authorized.

Inspect and retain the relevant pre-operation state:

```sh
git status --short --branch
git show-ref --verify refs/heads/<branch>
git rev-parse --verify '<start-point>^{commit}'
git worktree list --porcelain
```

`git show-ref --verify` exits nonzero when the local branch does not exist; absence means `-C` will create rather than reset it. When the branch exists, retain its old object ID in task state before resetting it.

Run the explicit force-create:

```sh
git switch -C <branch> <start-point>
```

Do not append `--force` or `--discard-changes`. After success, verify the current branch and `HEAD`. If the branch is checked out in another worktree, the start point is unresolved, or ref-reset authority is absent, stop before mutation.

### Stage and verify a commit

Keep one commit to one related change unit. Split unrelated changes before staging. A dependency update may include only the upgrade and required fixes. Keep independently reviewable API, database, or schema changes separate, and keep later typo, omission, or index fixes in a later commit.

Stage only task-owned paths and review the exact staged result:

```sh
git add -- <path>...
git status --short
git diff --staged
```

Block the commit when staged verification is unavailable or differs from the intended change unit.

When the staged diff does not cover the material risk, run the narrowest deterministic supporting check:

- Docs: staged readback, formatting, or whitespace check.
- Code: the relevant lint, typecheck, test, or build.

Fix and rerun a failed check or report it as blocking. If a supporting check is unavailable, report the reason and residual risk; continue only when task risk permits. Skip only with user approval or when the check is disproportionate to task risk, and report the basis and residual risk. A skip is not a pass.

### Prepare and submit the commit message

Resolve commit-message conventions by instruction priority. Apply a current user convention when it clearly overrides the same rule; otherwise apply the repository convention.

Use `type: detailed subject` with a subject under 120 characters. A stricter repository limit applies instead of 120. Select the most specific supported type:

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

Use a different type set only when a higher-priority repository or user convention explicitly requires it; do not invent types. Keep the subject specific to the staged scope and avoid vague wording or unrelated concerns.

After one blank line, add a bullet body covering the concrete change and verification. Keep skipped verification or residual risk visible in the body or final report. Do not include literal `\n`, unnecessary blank lines, unrelated scope, shell syntax, heredoc or EOF delimiters, or quoting wrappers.

Use a dedicated message file. Run allocation, write, readback, final staged verification, commit, and cleanup as separate steps:

```sh
mktemp /tmp/toolkit-git-message.XXXXXX
```

Continue only when allocation exits zero and returns exactly one path matching `/tmp/toolkit-git-message.*`. Preserve that exact path. Otherwise block without guessing a cleanup target; if the current allocator invocation returned one matching path, clean only that exact path.

Write only the commit message with a file edit tool, then read it back:

```sh
sed -n '1,200p' '/tmp/toolkit-git-message.EXACT'
```

If writing or readback fails, or the message violates the contract, clean the exact allocated path and block the commit. Verify the type, subject length, staged-scope wording, blank line, bullet body, verification evidence, and absence of unintended shell text or escapes.

Recheck the scope and commit from that exact path:

```sh
git status --short
git diff --staged
git commit -F '/tmp/toolkit-git-message.EXACT'
unlink '/tmp/toolkit-git-message.EXACT'
```

Run `unlink` immediately after the commit attempt, including failure. Report commit and cleanup separately. Never construct or submit the message with heredoc/EOF, here-string, command substitution, `git commit -F -`, multiple `-m` arguments, or a combined shell script.

### Push a current branch

Resolve the remote, exact local branch, remote destination, and upstream before mutation:

```sh
git remote -v
git branch -vv
git push -u <remote> <branch>
```

After an upstream is confirmed, a later plain push may use:

```sh
git push
```

Do not use plain `git push` when the upstream or `push.default` behavior is unresolved.

### Push an explicit refspec

The refspec `<local-source>:<remote-destination>` maps two independently resolved refs:

```sh
git push <remote> <local-source>:<remote-destination>
```

Example:

```sh
git push origin wip:main
```

This sends local `wip` to remote `main`. It is not a synonym for “push the current branch.” Use it only when the exact mapping is requested or owned by an applicable repository rule. A repository rule that keeps `wip` local or protects `main` blocks this example.

Do not add `--force` or `--force-with-lease` after a rejection unless history rewriting is explicitly authorized and the exact expected remote ref is resolved.

## Post-Operation Verification

For a successful commit, run message verification after message-file cleanup:

```sh
git log -1 --format='%H%n%B'
```

Compare the stored full message with the expected message and applicable convention. Verify the actual type, subject length, staged-scope wording, blank line, bullet structure, and verification evidence. Allow a hook-added trailer or transformation only when the repository convention owns it.

Report message verification as failed when the stored message differs unexpectedly. Preserve the successful commit and do not amend, reset, or roll it back without separate authorization.

Verify the remaining state affected by the requested workflow:

```sh
git status --short --branch
git branch -vv
```

For a network-authorized push, compare local source and remote destination without repeating the mutation:

```sh
git rev-parse <local-source>
git ls-remote --heads <remote> refs/heads/<remote-destination>
```

Report commit, message verification, branch, upstream, and push results separately. Distinguish passed, failed, skipped, unavailable, and not requested checks.

## Reference Routing

- When a branch name uses `codex/`, `jira/prja-000`, or another policy-sensitive prefix, when a branch name must be derived from a ticket or task, or when force-create meets an existing-name policy, read [references/branch-conventions.md](references/branch-conventions.md).
- When switching or force-create is blocked or interrupted, HEAD is detached, commit execution is interrupted, an alias partially executes, push is rejected, authentication fails, the remote result is unclear, or refs diverge, read [references/recovery.md](references/recovery.md).
