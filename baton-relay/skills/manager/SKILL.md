---
name: manager
description: Manage task decomposition across isolated git worktrees with fresh subagents, committed and rebased handoffs, prepared commit integration, and worktree cleanup. Use when the main agent should coordinate bounded subagent work without doing the slice work directly. worktree orchestration, subagent lifecycle, commit handoff, rebase handoff, task decomposition, baton relay
---

# Manager

## Overview

Use this skill when a task should be planned as a repository workflow and executed by fresh subagents in isolated worktrees.
The main agent acts as the manager: write a Markdown workflow plan, assign jobs, verify committed and rebased handoffs, integrate prepared commits, and clean up worktrees.

Do not use this skill for a small single edit, a purely read-only answer, or work that cannot be separated into commit-sized slices.

## Core Contract

- The main agent orchestrates; subagents perform bounded task work.
- Write a Markdown `Workflow > Jobs > Runs` todo plan before dispatching subagents.
- One subagent owns one job and one git worktree.
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
3. Write a Markdown workflow plan with jobs, runs, needs, acceptance, and handoff rules.
4. Choose parallel or sequential execution from job dependencies and parallel blockers.
5. Create a branch and worktree for each selected subagent job.
6. Spawn a fresh subagent per worktree with a complete dispatch packet derived from the plan.
7. Require each subagent to work, verify, commit, and wait for merge-prep instruction.
8. Ask the subagent to rebase onto the current integration branch when its slice is ready.
9. Confirm the subagent reports the rebase target HEAD, commit hash, verification result, changed files, and residual risk.
10. If the integration branch HEAD still matches the subagent's rebase target HEAD, import the prepared commit.
11. If the integration branch moved, request a new rebase or resequence integration.
12. Run required integration verification.
13. Clean up completed worktrees.

If orchestration does not fit, report `Orchestration fit: no`, `Spawn plan: none`, the caller-local handling path, verification expectation, and residual risk.

## Workflow Plan

The plan is a Markdown todo document, not a data file.
When writing the plan, use `templates/workflow-plan.md` as the base template.
Use frontmatter for stable workflow and job metadata.
Treat the plan as static except for the mutable allowlist: job `Status`, checklist states, evidence text appended under checklist items, and `Residual Risk`.
If the template file is unavailable, use this shape:

- `# Workflow: <name>`
- frontmatter: workflow objective, integration branch, dispatch fit, static job dependencies, worktree, write scope, parallel blockers, handoff conditions
- `## Jobs`
- `### Job N. <title>`
- `Status: planned`
- `#### Runs`
- `#### Acceptance`
- `#### Handoff`

Each `Run` must be an executable checklist item.
Each subagent job must include worktree, write scope, acceptance, and handoff.
Treat `needs` as a static dependency graph; do not update it during execution.
Start a job only after every job named in `needs` has completed its body checklist and handoff evidence.
Jobs with empty `needs` may start together unless parallel blockers make them sequential.
Update only the mutable allowlist as job status changes.

## Decomposition Rules

Use practical repository-management criteria before finalizing jobs:

- workstream: feature, bug, docs, verification, release-surface, or other practical work lane
- write scope: module, screen, API, document surface, or generated artifact this job may change
- dependency: contract, source change, setup, or verification that must finish first
- parallel blockers: shared file, shared contract, migration, generated output, or secret surface
- acceptance: verification, commit, rebase, and report conditions for job completion

Parallelize only jobs with disjoint write scopes, empty parallel blockers, and no dependency edge.
Make shared-file, shared-contract, generated-output, migration, and secret-surface jobs sequential.
Make generated-output updates sequential after source changes unless the repository policy says otherwise.
Do not make pure investigation, triage, or hypothesis elimination a subagent job until it becomes a commit-sized fix job.

## Dispatch Packet

Send each subagent:

- `job_id`
- `objective`
- `worktree_path`
- `branch`
- `base_or_integration_branch`
- `write_scope`
- `runs`
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
