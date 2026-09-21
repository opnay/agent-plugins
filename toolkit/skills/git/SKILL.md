---
name: git
description: Run task-scoped Git commit, alias.codex maintenance, branch, force-create, push, and recovery workflows. Use for commits, alias.codex install, uninstall, or doctor, branch creation or switching, explicit git switch -C, current-branch or exact refspec pushes, and partial Git workflow recovery; not for incidental read-only Git inspection.
---

# Git Workflow

Read repository instructions first. Commit, alias maintenance, branch, push, and recovery are separate requested steps. Do not infer another step, force, destructive cleanup, GitHub work, release, or publishing.

## Commit

Inspect and stage one task-owned change unit:

```sh
git status --short --branch
git diff
git diff --staged
git add -- <path>...
git diff --staged
```

Keep unrelated changes unstaged. Stop when the staged diff is unavailable or outside scope. Run the narrowest needed check: docs readback, formatting, or whitespace; code lint, typecheck, test, or build. Report failed, unavailable, or skipped checks and residual risk.

Follow repository message rules. Otherwise use a specific supported type, a subject below 120 characters, and a bullet body. Do not include literal `\n`, shell syntax, delimiters, or unnecessary blank lines.

Create the message after staged verification:

```sh
git codex message create --stdin <<'EOF'
fix: describe the staged change

- Describe the change and verification evidence.
EOF
git codex commit <returned-MSG-ID>
```

`message create` validates stdin input and prints an ID only on success. Do not repeat validation unless the message changes or the result is unclear. If it is unavailable, report that installation is needed and ask for install authority or a manual workflow. Read [references/message-lifecycle.md](references/message-lifecycle.md) for manual allocation, retained files, metadata paths, or commit result codes.

## Alias codex

For an explicit install, uninstall, or doctor request, read [references/alias-codex.md](references/alias-codex.md). Use the command output as the result; do not add a normal post-command query.

## Branch and push

```sh
git switch -c <branch> <start-point>
git switch <branch>
```

Use `git switch -C <branch> <start-point>` only with explicit force-create authority. First confirm the branch, start point, existing target ref, working tree, and other worktree use. It does not permit discarding changes, deleting branches, or force-pushing.

Run the requested push form directly:

```sh
git push -u <remote> <branch>
git push <remote> <local-source>:<remote-destination>
```

Do not infer a refspec from a general push request. `git push origin wip:main` sends local `wip` to remote `main`. Do not pre-read remote, upstream, or ref mapping. Use clear command output as the result.

Read [references/branch-conventions.md](references/branch-conventions.md) for policy-sensitive prefixes, derived names, or force-create conflicts.

## Recovery

After a failed, interrupted, conflicting, or unclear result, inspect only the state needed to recover. Preserve completed work; do not retry blindly, roll back, or force-push. Read [references/recovery.md](references/recovery.md).
