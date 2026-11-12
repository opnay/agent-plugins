# Branch conventions

Read this reference when a prefix, ticket, repository rule, or branch collision affects the requested branch operation.

## Apply the owning rule

Use repository rules first, then the user's exact branch request, then a documented environment convention. Preserve an exact supplied name. Do not turn a branch name seen elsewhere into a convention.

Derive a name only when the owning rule supplies its prefix, identifier, slug rule, and collision behavior. If a required input is missing, ask for it.

## Existing branches

Before switching to, reusing, or force-creating an existing branch, inspect only the relevant state:

```sh
git branch --all --verbose --no-abbrev
git status --short --branch
git worktree list --porcelain
```

- Existing local branch: confirm its commit and working-tree compatibility.
- Remote-tracking branch only: confirm the requested local name before creating a tracking branch.
- Another commit or worktree owns the name: report the conflict. Do not reset, rename, delete, or force-create it without explicit authority.
- `git switch -C`: confirm the exact branch and start point; retain the existing object ID before an authorized reset.

`git switch -C` does not authorize `--force`, `--discard-changes`, branch deletion, or force push.

## Push policy

A prefix does not choose a start point or push destination. When the user supplies a push form or exact refspec, run it directly; do not rediscover its remote, upstream, or mapping. A repository policy may prohibit the requested destination.
