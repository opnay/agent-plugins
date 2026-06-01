---
name: subagent-work
description: Run a reviewable work unit through a strict worker subagent lifecycle, from defining scope and spawning with a complete handoff packet through sync, verification, integration review, compact handoff, and close/dispose. worker subagent lifecycle, reviewable work unit, subagent handoff, close dispose, compact handoff
---

# Subagent Work

Use when one worker subagent should own one reviewable work unit while the main thread keeps user conversation, approvals, scope decisions, and final integration judgment.

Treat workers as disposable. Open one worker for one clear unit, review its result, then close it. Start the next unit with a compact handoff instead of accumulated worker context.

## Lifecycle

1. `Prepare`: define work unit, main-owned decisions, worker-owned execution, approval boundary, close criteria.
2. `Spawn`: send a self-contained packet with only needed context.
3. `Operate`: worker implements and runs first-pass validation; main thread handles user-facing decisions and non-overlapping orchestration.
4. `Sync`: worker returns at checkpoints or blockers with changed paths, decisions, validation, and risk.
5. `Integrate`: main thread inspects output and changed files before accepting.
6. `Close`: close/dispose when reviewable, approval-blocked, or outside the original packet.
7. `Handoff`: for a next unit, write a compact handoff with completed work, remaining scope, constraints, and risks.

## Responsibility Split

Main thread owns:

- user questions and requirement negotiation
- approval-sensitive choices
- destructive, external, commit, push, PR, publish, release, and version-bump decisions
- final review of worker output
- rework, user question, new unit, or stop decisions

Worker owns only:

- implementation inside assigned scope
- local investigation needed for that implementation
- first-pass validation
- assumptions, blockers, and residual risk reporting

If edits are possible, tell the worker it is not alone in the codebase and must not revert or overwrite unrelated changes.

## Worker Packet

```text
Return when: <condition for one reviewable result>
Stop if: <scope breach, approval need, ambiguity, conflicting edits, missing dependency, failed validation, or blocker>
Close plan: <close after one result | close when blocked | remain only for same-unit bounded follow-up>
Main-thread blocked state: <blocked on X | not blocked; main can do Y>

Task: <one reviewable work unit>

Context:
- You are not alone in the codebase. Do not revert or overwrite unrelated changes.
- <minimal facts needed for the unit>

Main-owned decisions:
- <questions, approvals, or integration choices to route back>

Assigned work unit:
- <implementation or verification responsibility>

Editable scope:
- <files, directories, modules, or "read-only">

Do not touch:
- <excluded files, surfaces, workflows, or actions>
- Do not approve or execute destructive, external, commit, push, PR, publish, release, or version-bump actions.

Validation:
- <expected checks, commands, fixtures, or review steps>

Output:
- Changed paths
- Summary of work
- Decisions or assumptions
- Validation run
- Validation skipped and why
- Blockers or approval needs
- Residual risk
```

## Sync And Stop

Set checkpoints before spawning when the unit is not a single short pass, such as after investigation, first patch, or validation failure.

Worker stops and returns evidence when:

- scope is wrong or too small
- approval is needed
- requirements are behavior-changing ambiguous
- existing edits conflict
- validation fails and the fix is not obvious in scope
- dependency, credential, network, or external system is required
- destructive, commit, push, PR, publish, release, or version-bump action is needed
- nested delegation would exceed assigned scope or bypass approval

Nested subagents are allowed only inside the worker's assigned scope. They cannot launder approval-sensitive work.

## Integrate

Never forward worker output as final without review. In the main thread:

- inspect changed paths against editable scope
- compare result with task, constraints, and validation expectation
- check assumptions and residual risk
- run or request added validation when needed
- accept, send same-unit follow-up, ask the user, or split a new unit

Reuse the same worker only for bounded follow-up with the same goal, ownership, approval limits, and close criteria.

## Compact Handoff

```text
Completed:
- <what is now true>

Changed paths:
- <paths from completed unit>

Remaining scope:
- <next reviewable work unit only>

Constraints to preserve:
- <style, ownership, approval limits, known conflicts>

Validation state:
- <checks run, checks still needed, failures>

Residual risk:
- <what the next worker or main thread must watch>
```

## Close

Close or dispose the worker when implementation and first-pass validation are done, a commit-ready/saved/complete/next-unit signal arrives, approval or ambiguity blocks progress, scope changes, or compact handoff is clearer than accumulated worker context.
