---
name: subagent-work
description: Run a reviewable work unit through a strict worker subagent lifecycle, from defining scope and spawning with a complete handoff packet through sync, verification, integration review, compact handoff, and close/dispose. worker subagent, subagent lifecycle, reviewable work unit, subagent handoff, close dispose, compact handoff
---

# Subagent Work

Use this skill when a worker subagent should own one reviewable work unit while the main thread keeps user conversation, approvals, scope decisions, and final integration judgment.

Treat each worker as disposable. Open a worker for one clear work unit, review its result, close it when that unit is reviewable, and start the next unit with a compact handoff instead of carrying accumulated worker context forward.

## Lifecycle

1. Prepare: define the reviewable work unit, main-owned decisions, worker-owned execution, approval boundary, and close criteria.
2. Spawn: send a self-contained worker packet with only the context needed to execute the unit.
3. Operate: let the worker implement and run first-pass validation while the main thread handles user-facing decisions and non-overlapping orchestration.
4. Sync: require the worker to return at checkpoints or blockers with changed paths, decisions, validation, and risk.
5. Integrate: inspect the worker output and changed files before treating the work as complete.
6. Close: close or dispose the worker once the work unit is reviewable, blocked on approval, or no longer matches the original packet.
7. Handoff: if more work remains, write a compact handoff for the next worker with completed work, remaining scope, constraints, and risks.

## Split Responsibilities

Keep these responsibilities in the main thread:

- user questions and requirement negotiation
- approval-sensitive choices
- destructive, external, commit, push, PR, publish, release, and version-bump decisions
- final review of worker output
- deciding whether to rework, ask the user, split a new unit, or stop

Give the worker only bounded execution work:

- implementation within assigned ownership
- local investigation needed for that implementation
- first-pass validation
- clear reporting of assumptions, blockers, and residual risk

If the worker can edit files, tell it that it is not alone in the codebase and must not revert or overwrite unrelated changes.

## Worker Packet Template

```text
Return when: <the exact condition for returning one reviewable result>
Stop if: <scope breach, approval need, ambiguity, conflicting edits, missing dependency, failed validation, or other blocker>
Close plan: <close after one result | close when blocked | remain only if the same work unit needs a bounded follow-up>
Main-thread blocked state: <what the main thread is blocked on, or what can proceed in parallel>

Task: <one reviewable work unit>

Context:
- You are not alone in the codebase. Do not revert or overwrite unrelated changes.
- <minimal facts needed to perform the work>

Main-owned decisions:
- <questions, approvals, or integration choices the worker must route back>

Assigned work unit:
- <implementation or verification responsibility>

Editable scope:
- <files, directories, modules, or "read-only">

Do not touch:
- <files, surfaces, workflows, or actions outside the unit>
- Do not approve or execute destructive, external, commit, push, PR, publish, release, or version-bump actions.

Validation:
- <expected checks, commands, fixtures, or review steps>

Output:
- Changed paths
- Summary of work
- Decisions or assumptions
- Validation run
- Validation not run and why
- Blockers or approval needs
- Residual risk
```

## Sync And Escalation

Set sync points before spawning when the unit is not a single short pass. Use checkpoints such as "after investigation," "after first patch," or "after validation failure."

Tell the worker to stop and return evidence when it hits any of these conditions:

- the assigned scope is too small or points at the wrong files
- a user-gated approval is needed
- requirements are ambiguous enough to change behavior
- existing edits conflict with the intended change
- validation fails and the fix is not obvious within scope
- a dependency, credential, network action, or external system is required
- the task would require destructive, commit, push, PR, publish, release, or version-bump action
- nested delegation would exceed the assigned scope or bypass an approval boundary

Nested subagents are allowed only inside the worker's assigned scope. Do not let nested delegation launder approval-sensitive work.

## Review And Integrate

Do not forward the worker result as the final answer without review. In the main thread:

- inspect the changed paths and make sure they match the editable scope
- compare the result against the task, constraints, and validation expectation
- check that assumptions and residual risk are explicit
- run or request any additional validation needed for integration confidence
- decide whether to accept, send a bounded follow-up to the same worker, ask the user, or split a new work unit

Use the same worker only for bounded follow-up on the same work unit. If the scope changes materially, close it and create a new packet.

## Compact Handoff

When the next unit needs a new worker, write a compact handoff instead of passing the whole prior conversation:

```text
Completed:
- <what is now true>

Changed paths:
- <paths from the completed unit>

Remaining scope:
- <next reviewable work unit only>

Constraints to preserve:
- <style, ownership, approval limits, known conflicts>

Validation state:
- <checks run, checks still needed, failures>

Residual risk:
- <what the next worker or main thread must watch>
```

## Close And Dispose Rules

Close or dispose the worker when:

- implementation and first-pass validation for the work unit are done
- the main thread receives a commit-ready, saved, complete, or next-unit signal
- approval, ambiguity, conflict, or missing dependency blocks progress
- the requested scope changes enough to require a new packet
- accumulated worker context is less clear than a compact handoff

Do not keep a worker alive across multiple reviewable work units. The only exception is a narrow follow-up that preserves the same goal, ownership, approval limits, and close criteria.
