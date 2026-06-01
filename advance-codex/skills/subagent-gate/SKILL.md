---
name: subagent-gate
description: Prepare a subagent handoff before spawning or messaging a subagent, including return point, stop boundary, minimal context packet, ownership, output contract, approval limits, and main-thread blocked state.
---

# Subagent Gate

Use before spawning a subagent or sending a substantial new message to one. Produce a handoff packet only when the task is bounded, separable, and useful outside the main thread.

## Gate

Keep work local or return `no-handoff` when:

- user approval is the core task
- custom agent definition is unsettled
- evaluation design and judgment must stay in the main thread
- destructive, external, commit, push, PR, publish, release, or version-bump decisions are central
- return point or stop boundary cannot be stated clearly

## Packet Order

Write the exit plan before task details:

1. `Return when`: exact condition for one result.
2. `Stop if`: approval need, ambiguity, scope breach, conflict, or blocker.
3. `Close plan`: close after one result or remain for bounded follow-up.
4. `Main-thread blocked state`: what waits for the result, or what can proceed in parallel.

Then add only needed context:

- `Goal`
- `Relevant facts`
- `Assigned scope`
- `Constraints`
- `Expected output`
- `Assumptions`

If edits are possible, say the subagent is not alone in the codebase and must not revert or overwrite unrelated changes.

## Approval Limits

Do not ask a subagent to approve or execute user-gated, destructive, external, commit, push, PR, publish, release, or version-bump actions. If one is needed, it reports evidence and options to the main thread.

## Output

Require fields the main thread can inspect immediately:

- changed paths, if any
- decisions made
- assumptions
- validation run
- validation skipped and why
- residual risk
- blockers or approval needs

## Template

```text
Return when: <result condition>
Stop if: <approval boundary, scope breach, ambiguity, conflict, or blocker>
Close plan: <close after one result | remain for bounded follow-up>
Main-thread blocked state: <blocked on X | not blocked; main can do Y>

Task: <bounded task>

Context:
- You are not alone in the codebase. Do not revert or overwrite unrelated changes.
- <minimal relevant facts>

Assigned scope:
- <files, modules, questions, or responsibilities>

Constraints:
- <write limits, safety limits, style rules>
- Do not approve or execute destructive, external, commit, push, PR, publish, release, or version-bump actions.

Output:
- <required return fields>
```
