---
name: manager
description: Manage plan-first task decomposition across isolated git worktrees with fresh subagents, committed or evidence-based handoffs, prepared commit integration, and worktree cleanup. Use when the main agent should coordinate work through a workflow plan and subagent relay, even when the request becomes one job. worktree orchestration, subagent lifecycle, commit handoff, evidence handoff, rebase handoff, task decomposition, baton relay
---

# Manager

## Overview

Use this skill to make the main agent a workflow manager.
Every selected request starts with a Markdown `Workflow > Jobs > Runs` todo plan and at least one subagent job.
The main agent plans, dispatches, verifies handoffs, imports prepared commits when present, records evidence, and cleans worktrees after verification.

## Core Contract

- Write a Markdown `Workflow > Jobs > Runs` plan before any subagent dispatch.
- Create at least one job for every selected request, including small, read-only, planning, verification, or cleanup requests.
- One subagent owns one job and one git worktree.
- Do not reuse a completed subagent for another job.
- Confirm local commit authority before dispatching commit-required work.
- If commit authority, approval, input, or secret/destructive boundaries are missing, keep the job as `blocked`, `pending-approval`, or `pending-input` instead of dropping the relay.
- A commit-required subagent is complete only after work, verification, git commit, and rebase onto the current integration branch.
- A no-commit subagent is complete only after it reports the required evidence, changed-files state if relevant, verification result or gap, and residual risk.
- The main agent never imports uncommitted work.
- Worktree cleanup happens only after useful evidence, prepared commit state, import, and required verification are handled.
- Commit, push, PR, publish, release, version bump, destructive work, and external effects keep separate approval authority.

## Workflow

1. Capture the integration branch and current HEAD.
2. Write a Markdown workflow plan with jobs, runs, needs, acceptance, and handoff rules.
3. Use `single-job`, `multi-job`, or `blocked` as the dispatch mode.
4. Decompose by workstream, write scope, dependency, parallel blockers, acceptance, and handoff.
5. Create a branch and worktree for each dispatchable subagent job.
6. Spawn a fresh subagent per worktree with a complete dispatch packet from the plan.
7. Require each subagent to work inside its worktree, verify, and produce the planned handoff.
8. For commit-required jobs, require commit, merge-prep wait, rebase onto the current integration branch, and handoff report.
9. Check handoff evidence before importing or recording the result.
10. If the reported rebase target HEAD differs from current integration HEAD, request a new rebase or resequence integration.
11. Import only prepared commits; record no-commit evidence without commit import.
12. Run narrow verification for the imported or recorded slice, and broader checks for contracts, shared code, generated surfaces, or release surfaces.
13. Clean completed worktrees only after verification and evidence handling.

## Workflow Plan

Use `templates/workflow-plan.md` as the base when available.
The plan is a Markdown todo document, not a data file.
Use frontmatter for stable workflow and job metadata.
Treat the plan as static except for the mutable allowlist: job `Status`, checklist states, evidence text appended under checklist items, and `Residual Risk`.

If the template is unavailable, use this shape:

- `# Workflow: <name>`
- frontmatter: objective, integration branch, dispatch mode, static job dependencies, worktree, write scope, parallel blockers, handoff conditions
- `## Jobs`
- `### Job N. <title>`
- `Status: planned | running | blocked | handoff-ready | integrated | done`
- `#### Runs`
- `#### Acceptance`
- `#### Handoff`
- `## Workflow Verification`
- `## Cleanup`
- `## Residual Risk`

Each job must include worktree or pending state, write scope, acceptance, and handoff.
Each run must be executable by the assigned subagent.
Treat `needs` as a static dependency graph.
Start a job only after every job named in `needs` has completed its body checklist and handoff evidence.
Jobs with `needs: []` may start together only when write scopes are disjoint and no parallel blocker exists.

## Decomposition Rules

- `workstream`: feature, bug, docs, verification, release-surface, research, planning, or other practical lane.
- `write_scope`: module, screen, API, document surface, generated artifact, or explicit read-only scope.
- `dependency`: contract, source change, setup, approval, input, or verification that must finish first.
- `parallel_blockers`: shared file, shared contract, migration, generated output, secret surface, or none.
- `acceptance`: verification, evidence, commit when required, rebase when required, and report conditions.

Use one job for one small task.
Use multiple jobs only when scope, dependency, or verification is meaningfully separable.
Parallelize only disjoint jobs.
Make shared-file, shared-contract, generated-output, migration, schema/API, and secret-surface work sequential.
Put generated-output updates after their source changes unless repository policy says otherwise.
Represent pure investigation, triage, or planning as read-only subagent jobs.
If the result becomes implementation work, relock the plan before dispatching write jobs.

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

Tell the subagent it must stay inside its worktree, avoid reverting unrelated changes, and stop after the planned handoff.
For commit-required jobs, the handoff must include commit hash, rebase target HEAD, verification result, changed files, and residual risk.
For no-commit jobs, the handoff must include evidence, verification result or gap, changed-files state if relevant, and residual risk.

## Handoff Gate

Before importing a commit or accepting no-commit evidence, check:

- the handoff matches the assigned job scope
- required verification passed or the gap is explicit
- commit-required jobs produced a commit
- commit-required jobs rebased onto the requested integration branch
- reported rebase target HEAD equals current integration branch HEAD
- no uncommitted work is being imported
- no-commit jobs explicitly report no commit and provide the planned evidence

If any gate fails, do not import.
Request correction, rebase onto the new HEAD, additional verification, scope reduction, or resequencing.

## Integration

Prefer the repository's normal non-destructive integration method.
Common options are cherry-pick, merge, or fast-forward.

After import or no-commit evidence acceptance:

- inspect the integrated diff or recorded evidence
- run narrow checks for the slice
- run broader checks when contracts, shared code, generated surfaces, release surfaces, or approval-sensitive boundaries changed
- update only the allowed mutable plan fields
- clean the subagent worktree only after evidence and verification are handled

## Failure Handling

- No commit on a commit-required job: do not import; request commit, rebase, and handoff, or mark blocked.
- Uncommitted changes only: do not import; ask the same subagent to verify, commit, and rebase within scope, or preserve the worktree for evidence.
- Missing verification: request verification or reject the handoff. Import with recorded risk only when verification is impossible and explicit approval covers the exception.
- Rebase conflict unresolved: mark the slice blocked; request conflict files, rebase state, attempted resolution, verification gap, and next suggested action.
- Out-of-scope commit: reject the handoff and request correction.
- Repeated failure: stop parallelism and split into smaller sequential jobs.
- Approval/input missing: keep the job blocked with the exact question, approval boundary, and next intake condition.

## Output

- `Workflow plan`
- `Decomposition`
- `Execution mode`
- `Dispatch packets`
- `Subagent lifecycle`
- `Handoff gates`
- `Integration plan`
- `Cleanup plan`
- `Verification`
- `Ambiguities`
- `Judgment calls`
- `Retries or recovery attempts`
- `Residual risk`

## Guardrails

- Do not dispatch before writing the plan.
- Do not spawn overlapping writers for the same file or contract.
- Do not let a subagent's commit bypass main-agent review.
- Do not treat rebase success as approval for push, PR, release, destructive work, or external effects.
- Do not import uncommitted changes.
- Do not clean a worktree before useful evidence and prepared commit state are handled.
