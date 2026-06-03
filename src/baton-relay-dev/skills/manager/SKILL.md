---
name: manager
description: Manage task decomposition across isolated git worktrees with fresh subagents, committed and rebased handoffs, prepared commit integration, and worktree cleanup. Use when the main agent should coordinate bounded subagent work without doing the slice work directly. worktree orchestration, subagent lifecycle, commit handoff, rebase handoff, task decomposition, baton relay
---

# Manager

## Overview

Use this skill when a task should be split into isolated worktree tasks and executed by fresh subagents.
The main agent acts as the manager: decompose the work, assign each slice, verify committed and rebased handoffs, integrate prepared commits, and clean up worktrees.

Do not use this skill for a small single edit, a purely read-only answer, or work that cannot be separated into commit-sized slices.

## Core Contract

- The main agent orchestrates; subagents perform bounded task work.
- One subagent owns one task slice and one git worktree.
- Do not reuse a completed subagent for the next slice; create a fresh subagent.
- This skill does not grant commit authority by itself; confirm local subagent commit authority before dispatch.
- A subagent is complete only after work, verification, git commit, and rebase onto the main agent's current integration branch.
- The main agent must not import uncommitted changes from a subagent worktree.
- The main agent integrates only prepared commits.
- Worktree cleanup happens only after commit import and required verification.
- Commit, push, PR, publish, release, version bump, destructive work, and external effects still require their own approval authority.

## Workflow

1. Decide whether worktree orchestration is warranted.
2. Capture the integration branch and current HEAD.
3. Decompose the task into commit-sized slices.
4. Choose parallel or sequential execution.
5. Create a branch and worktree for each selected slice.
6. Spawn a fresh subagent per worktree with a complete dispatch packet.
7. Require each subagent to work, verify, commit, and wait for merge-prep instruction.
8. Ask the subagent to rebase onto the current integration branch when its slice is ready.
9. Confirm the subagent reports the rebase target HEAD, commit hash, verification result, changed files, and residual risk.
10. If the integration branch HEAD still matches the subagent's rebase target HEAD, import the prepared commit.
11. If the integration branch moved, request a new rebase or resequence integration.
12. Run required integration verification.
13. Clean up completed worktrees.

If orchestration does not fit, report `Orchestration fit: no`, `Spawn plan: none`, the caller-local handling path, verification expectation, and residual risk.

## Decomposition Axes

Use more than one axis before finalizing slices:

- problem: root-cause hypothesis, requirement, risk, blocker
- code: package, module, layer, frontend, backend, docs, tests, infra
- workflow: discovery, implementation, verification, documentation, refactor
- contract: schema, API, service, client, UI, test contract
- verification: lint, typecheck, unit test, integration test, build, smoke
- conflict: shared files, shared contracts, generated outputs, migrations
- generated output: treat generated release surfaces, generated clients, and generated migration outputs as integration/build artifacts unless the repository explicitly owns them as source
- security: keep secrets, token rotation, live credentials, external auth calls, and unredacted logs or fixtures out of subagent scope unless separately approved
- read-only: do not make pure investigation, triage, or hypothesis elimination a manager slice until it becomes a commit-sized fix slice

Parallelize only disjoint slices.
Make shared-file or shared-contract slices sequential.
Make generated-output updates sequential after source changes unless the repository policy says otherwise.

## Dispatch Packet

Send each subagent:

- `objective`
- `worktree_path`
- `branch`
- `base_or_integration_branch`
- `write_scope`
- `non_goals`
- `verification_required`
- `commit_approval_state`
- `commit_expectation`
- `rebase_target`
- `handoff_output`
- `stop_condition`

Tell the subagent it must stay inside its worktree, avoid reverting unrelated changes, and stop after reporting a committed and rebased handoff.
If local commit authority is missing, do not dispatch a worker that must commit.

## Subagent Lifecycle

Start:

- create task slice
- create branch and worktree
- spawn fresh subagent
- provide dispatch packet

Work:

- inspect only needed context
- implement or investigate inside the worktree
- verify the slice
- create a git commit

Merge prep:

- wait for the main agent's merge-prep request
- rebase onto the main agent's current integration branch
- resolve rebase conflicts inside the subagent worktree
- report commit hash, rebase target HEAD, verification, changed files, and residual risk

Stop:

- after successful rebase handoff, the subagent terminates
- the next task uses a new subagent

## Handoff Gate

Before importing a subagent commit, the main agent checks:

- the subagent produced a commit
- the commit is within the assigned scope
- required slice verification passed or the gap is explicit
- the subagent rebased onto the requested integration branch
- the reported rebase target HEAD equals the current integration branch HEAD
- no uncommitted work is being imported

If the HEAD check fails, do not import the commit.
Request a rebase onto the new HEAD or redesign the integration order.

## Integration

Prefer the repository's normal non-destructive integration method.
Common options are cherry-pick, merge, or fast-forward where appropriate.

After import:

- inspect the integrated diff
- run narrow checks for the imported slice
- run broader checks when contracts, shared code, or generated surfaces changed
- clean the subagent worktree only after the imported result is verified enough for the active task

## Failure Handling

- No commit: keep or discard the worktree based on evidence value; do not import changes.
- Uncommitted changes only: do not import; ask the same subagent to verify, commit, and rebase within scope, or preserve/discard the worktree based on evidence value.
- Rebase conflict unresolved: mark the slice blocked and decide whether to shrink scope, resequence, or handle manually with explicit approval.
- For unresolved conflict, request conflict files, rebase state, attempted resolution, verification gap, and next suggested action; do not clean up the worktree.
- Out-of-scope commit: reject the handoff and request correction.
- Missing verification: default to requesting verification or rejecting the handoff. Import with recorded risk only when verification is impossible and explicit approval covers the exception.
- Repeated failure: stop parallelism and split into smaller sequential slices.

## Output

- `Orchestration fit`
- `Decomposition`
- `Execution mode`
- `Dispatch packets`
- `Subagent lifecycle`
- `Handoff gates`
- `Integration plan`
- `Cleanup plan`
- `Verification`
- `Residual risk`

## Guardrails

- Do not make subagent use automatic.
- Do not spawn overlapping writers for the same file or contract.
- Do not let a subagent's commit bypass main-agent review.
- Do not treat rebase success as approval for push, PR, release, destructive work, or external effects.
- Do not clean a worktree before its useful evidence and prepared commit state have been handled.
