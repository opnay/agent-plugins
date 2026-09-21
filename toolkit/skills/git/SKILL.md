---
name: git
description: Run task-scoped Git commit, managed alias.codex maintenance, branch, force-create, push, and recovery workflows. Use for heredoc commit messages, commits, alias.codex install, uninstall, or doctor, branch creation or switching, explicit git switch -C, upstream setup, current-branch or explicit refspec pushes, and partial Git workflow recovery; do not activate for incidental read-only Git inspection.
---

# Git Workflow

## Select the authorized work

Read repository instructions first. Commit, alias maintenance, branch, push, and recovery are composable steps with separate authority. A commit request does not authorize alias changes, push, force, destructive cleanup, version bumps, GitHub work, or publishing.

`git-codex` owns message-file lifecycle and commit outcome detection. It does not stage, amend, push, change branches, bypass hooks, or judge message meaning. Use this skill without sibling skills or development-only documents. Do not default to history rewriting, forced branch deletion, force push, `reset --hard`, or `--discard-changes`.

## Alias codex maintenance

`git codex` uses the managed global `alias.codex` dispatch to `$HOME/.local/bin/git-codex`. For an explicit `alias.codex` install, uninstall, or doctor request, read [references/alias-codex.md](references/alias-codex.md). That reference owns the expected value, binary installation, verification, force, and preservation rules.

## Commit scope and message

Inspect the working tree, then stage one task-owned related change unit:

```sh
git status --short --branch
git diff
git diff --staged
git add -- <path>...
git status --short
git diff --staged
```

Preserve unrelated changes. Block a commit if staged verification is unavailable or differs from the intended scope. If the diff leaves material risk unresolved, run the narrowest supporting check: readback, formatting, or whitespace for docs; relevant lint, typecheck, test, or build for code.

Distinguish passed, failed, unavailable, and skipped checks. Fix and rerun failures or report the blocker. Continue past unavailable checks only when task risk permits, reporting the cause and residual risk. A skip needs user approval or a proportionality reason, with residual risk disclosed.

Apply instruction priority and a current user's explicit override of the same rule; otherwise preserve repository convention. Keep the subject under 120 characters or a stricter repository limit. Without another convention, use `type: detailed subject`, one blank line, and a bullet body explaining changes and verification.

Use the most specific supported type: `feat`, `fix`, `refactor`, `docs`, `test`, `perf`, `style`, `build`, `ci`, or `chore`. Follow an applicable alternative type set without inventing types. Keep wording specific to the staged change. Do not embed literal `\n`, shell syntax, heredoc delimiters, quoting wrappers, or unnecessary blank lines in the message. Report skipped checks and residual risk in the body or final response.

## Preferred message lifecycle

After staged verification, create the complete message with a quoted heredoc:

```sh
git codex message create --stdin <<'EOF'
fix: describe the staged change

- Describe the change and verification evidence.
EOF
```

Use a delimiter absent from the message. Quoting it prevents variable and command expansion; the shell consumes the delimiter, not the message file. This path is preferred over allocating an empty file and editing it with a file tool.

Create reads through EOF, writes a private `0600` file, and automatically validates it. Only success prints one short `MSG-…` ID. Use that exact returned ID for the later authorized commit:

```sh
git codex commit <returned-MSG-ID>
```

Create never commits. Confirm the staged scope before commit; recheck it if changes or intervening operations make the prior observation stale. Do not repeat readback or `message validate` after successful stdin creation unless the message was subsequently edited or the result is uncertain. Commit still validates immediately before its mutation.

### Git metadata paths

Files live in the current repository or worktree's actual Git directory, resolved by `git rev-parse --absolute-git-dir`. A regular checkout stores `.git/MSG-…`; nested directories, linked worktrees, separate Git directories, and bare repositories follow Git's own discovery. Do not assume `.git` is a directory or create a nested `.git`.

Pass the short ID from any directory in the same repository/worktree. `validate` and `commit` accept only `MSG-…` files in that Git directory, by ID or real relative or absolute path. Other files are rejected. The Git directory must be discoverable, user-owned, and not group/other writable.

### Empty allocation and validation

Plain `git codex message create` allocates an empty file. Use it only when incremental editing is needed. Record the allocated file's device and inode before writing, read back against expected content, and run:

```sh
git codex message validate <returned-MSG-ID>
```

Validation checks the allowed canonical parent, a private user-owned regular file, UTF-8, NUL, lone CR, literal `\n`, a nonempty subject, and the subject/body delimiter. Standalone validation never edits or deletes its input.

If stdin reading, writing, or automatic validation fails, create removes only its own still-safe file with matching identity and returns no success ID. If cleanup cannot be confirmed, it reports the retained path. For manual write, readback, or validation failure, block commit and apply the same identity-safe cleanup rule; preserve uncertain files.

## Commit outcomes

`git codex commit <message>` validates, records HEAD, invokes only `git commit -F <resolved-file>`, and confirms command success plus a new HEAD before deleting the same safe message file.

- `0`: commit confirmed; message deleted.
- `1`: definitive precondition, validation, or commit failure; the command preserves its input. Read `git-codex: commit_result: reason=<code> attempted=<true|false>`. The caller cleans `validation_failed` only when the recorded allocation identity remains safe; preserve `head_unavailable` and `commit_failed` files. An attempted commit requires state inspection before recovery or fallback.
- `2`: commit confirmed; cleanup failed. Inspect HEAD, file identity, and stored message, then address only the remaining file.
- `3`: attempted commit with an unknown result. Inspect HEAD, file identity, and stored message before retrying or using a manual fallback.

## Branch and push

Create from a confirmed start point or switch to an existing local branch:

```sh
git switch -c <branch> <start-point>
git switch <branch>
```

Use `git switch -C <branch> <start-point>` only with explicit force-create authority from the user or owning repository workflow. Confirm exact branch, start point, existing target ref, working tree, and `git worktree list --porcelain` first. This does not authorize discarding changes, branch deletion, or force push.

Before push, resolve the remote URL, local source, remote destination, and upstream:

```sh
git branch --show-current
git branch -vv
git remote -v
git push -u <remote> <branch>
git push <remote> <local-source>:<remote-destination>
```

`git push origin wip:main` maps local `wip` to remote `main`; do not infer this from a general push request. Repository restrictions and authorization for the exact destination still apply.

Interpret exit status and result text together. When the output clearly identifies the expected destination and success, rejection, or no-op, report that result without another status, fetch, or remote-ref query. Query only the state needed to resolve a wrong invocation, interruption, missing or conflicting output, or an uncertain outcome; do not repeat a mutation to verify it.

Read [references/branch-conventions.md](references/branch-conventions.md) for policy-sensitive prefixes such as `codex/` or `jira/prja-000`, derived names, or force-create conflicts. Preserve exact supplied names; a prefix alone does not determine base, remote destination, or permission.

## Errors and recovery

- Use `message create` itself for lifecycle availability. Report an observed failure without changing alias configuration.
- Preserve completed work. Before another mutation after failure, inspect only the local or remote state needed to resolve that failure. A rejected push does not undo a successful commit or authorize force push.
- Read [references/recovery.md](references/recovery.md) for blocked branch changes, detached HEAD, interrupted or uncertain commits, partial alias execution, authentication failures, rejected or uncertain pushes, or divergent refs. Apply the message-file status rules above before cleanup.
- A user-selected manual workflow or unsupported platform preserves the same message validation, identity-safe cleanup, commit-failure preservation, and full-message verification contract.
