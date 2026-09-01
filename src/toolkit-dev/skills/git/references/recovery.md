# Git Workflow Recovery

Read this reference after a Git mutation fails, is interrupted, or has an unclear result. Recover from observed state; do not replay the original workflow blindly.

## Snapshot Before Recovery

Run the relevant local checks before another mutation:

```sh
git status --short --branch
git branch --show-current
git rev-parse HEAD
git branch -vv
git remote -v
git worktree list --porcelain
```

Record the failed command, its exit status and output, the last confirmed successful step, and any exact temporary path or pre-reset branch object ID retained in task state. Separate command failure from repository state.

## Branch Switch, Creation, Or Force-Create Is Blocked

- Dirty working tree: preserve staged, unstaged, and untracked changes. Do not use `--force`, `--discard-changes`, reset, or cleanup as a shortcut.
- Move current work to a new branch: when the user requested this outcome, verify the start point and create the branch without discarding the working tree.
- Name collision: inspect the existing local, remote-tracking, and worktree ownership before deciding whether reuse or explicit force-create is valid.
- `git switch -C` failure or interruption: recheck current branch, `HEAD`, and the target branch ref before deciding whether it created, reset, switched, or left the branch unchanged. Do not repeat it blindly.
- Completed force-create followed by a later failure: preserve the completed branch result. Do not restore the retained old object ID without separate authorization.
- Detached HEAD: identify the exact `HEAD` commit. Create a named branch at that commit only when the requested task or user direction requires preserving it there.

Read the branch-conventions reference when a prefix or repository branch policy affects the recovery decision.

## Commit Is Interrupted Or Fails

1. Check whether `HEAD` changed and whether the intended commit already exists.
2. Recheck staged and unstaged state; do not assume a failed client command means no hook or commit side effect occurred.
3. If the exact `/tmp/toolkit-git-message.*` path from the current allocator invocation is preserved in task state, run `unlink` once and report its result separately.
4. Do not delete a guessed path or a path whose allocator provenance is unknown.
5. Start a new commit attempt only after cleanup and current staged scope are resolved.

A successful commit followed by cleanup failure remains a successful commit with a separate cleanup failure. Do not undo the commit.

## An Alias Partially Executes

Inspect the alias definition, then determine which steps completed from repository state. For an alias that stages, commits, and pushes:

- staged changes with no new commit indicate an early commit failure;
- a new local commit with an unchanged remote indicates push did not complete;
- a matching remote ref indicates push completed even if later reporting was interrupted.

Continue only the authorized unfinished step. Never rerun the whole alias to recover.

## Push Is Rejected

- Non-fast-forward: do not retry with force. Inspect local source, remote destination, and divergence; merge, rebase, reset, or force requires its own authorized scope.
- Protected branch or server policy: preserve the local commit and report the rejected destination and server evidence.
- Hook rejection: keep the hook result distinct from authentication and divergence. Do not add `--no-verify` automatically.
- Authentication or network failure: do not rewrite the remote URL, weaken authentication, or repeat credential prompts indefinitely. Report the failed boundary and retain the local result.

## Push Result Is Unclear

Query the remote before repeating the mutation:

```sh
git rev-parse <local-source>
git ls-remote --heads <remote> refs/heads/<remote-destination>
```

- Matching object IDs: treat the remote update as complete and do not push again.
- Different object IDs: report divergence; do not infer that retrying is safe.
- Missing destination: report that the expected remote branch was not observed.
- Query unavailable: report the push as unverified rather than failed or complete.

## Partial Success Reporting

Report each requested step independently:

- branch: created, reset, switched, unchanged, failed, or unverified;
- commit: created, unchanged, failed, or unverified;
- cleanup: completed or failed with the exact retained path;
- push: updated, rejected, failed before remote mutation, or unverified;
- remaining state: staged, unstaged, untracked, ahead, behind, or divergent.

Do not automatically roll back a completed branch, force-create, or commit because a later step failed.
