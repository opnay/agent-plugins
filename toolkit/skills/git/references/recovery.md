# Git workflow recovery

Read this reference after a failed, interrupted, conflicting, or unclear mutation. Use command output first, then inspect only the state needed to decide the unfinished authorized step. Do not retry blindly or undo completed work.

## Branch

For a blocked switch or `git switch -C`, inspect the current branch, `HEAD`, target ref, working tree, and other worktree use. Preserve dirty changes. Do not use `--force`, `--discard-changes`, reset, or cleanup as a shortcut.

If a force-create may have completed, determine whether it created, reset, or switched before another mutation. Keep a completed branch result; restoring an old ref needs separate authority.

## Commit

Check whether `HEAD` changed, then inspect staged and unstaged scope. Preserve a retained `MSG-…` file after a definitive failure or unknown outcome. Read [message-lifecycle.md](message-lifecycle.md) for file identity, cleanup, and result codes.

A successful commit with cleanup failure needs cleanup, not another commit.

## Aliases

For an alias that stages, commits, and pushes, determine the completed requested steps from staged state and `HEAD`. Query the remote only when the push result is unclear. Continue with the authorized unfinished step; do not rerun the alias.

## Push

Treat destination-specific rejection output as a clear rejection. Do not query only to reconfirm it.

- Non-fast-forward: inspect divergence before an authorized merge, rebase, reset, or force decision. Never switch to force push automatically.
- Protected branch, hook, authentication, or network failure: preserve the local commit and report the failed boundary. Do not bypass hooks, alter the remote URL, or weaken authentication.
- Unclear result: before retrying, compare the attempted local source with the attempted remote destination:

```sh
git rev-parse <local-source>
git ls-remote --heads <remote> refs/heads/<remote-destination>
```

Matching IDs mean the update completed. Different IDs mean divergence. A missing destination or unavailable query leaves the result unverified.

## Report

Report each requested step separately: branch, commit, cleanup, push, and remaining staged or unstaged work.
