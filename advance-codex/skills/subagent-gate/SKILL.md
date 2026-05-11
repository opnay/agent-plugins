---
name: subagent-gate
description: Prepare a subagent handoff before spawning or messaging a subagent, including return point, stop boundary, minimal context packet, ownership, output contract, approval limits, and main-thread blocked state.
---

# Subagent Gate

Use this skill before spawning a subagent or sending a substantial new message to an existing subagent. Produce a handoff packet the subagent can act on without inheriting unnecessary main-thread context.

## Decide Whether To Hand Off

Use a subagent only when the work is bounded, separable, and useful to run outside the main thread. Keep the task local when the next main-thread step is immediately blocked, the scope cannot be cleanly owned, or the task depends on user-gated decisions.

For unsafe or out-of-scope delegation, return a no-handoff note instead of a prompt. Common no-handoff cases:

- the core work is user-gated approval
- custom agent definition is still unsettled
- scenario-based instruction evaluation must be designed and judged in the same thread
- destructive, external, commit, push, PR, publish, release, or version-bump decisions are central
- the return point or stop boundary cannot be stated clearly

## Start With The Exit Plan

Before task details, write:

- `Return when`: the exact condition for sending one result back
- `Stop if`: conditions that require returning without continuing
- `Close plan`: whether to close after one result or remain available for follow-up
- `Main-thread blocked state`: what the main thread cannot do until the result arrives, or what it can safely do in parallel

## Build The Minimal Context Packet

Include only what the subagent needs from its point of view:

- `Goal`: the concrete objective
- `Relevant facts`: facts the subagent must preserve
- `Assigned scope`: files, modules, questions, or responsibilities it owns
- `Constraints`: write limits, style rules, safety rules, and coordination notes
- `Expected output`: exact fields or artifact shape to return
- `Assumptions`: assumptions to preserve, verify, or report

Tell the subagent it is not alone in the codebase when edits are possible. It must not revert or overwrite unrelated changes, and it should adapt to changes made by others.

## Set Approval Limits

Do not ask a subagent to approve or execute user-gated, destructive, external, commit, push, PR, publish, release, or version-bump actions. If one becomes necessary, the subagent should report the need, evidence, and options back to the main thread.

## Output Contract

Request an output the main thread can immediately inspect:

- changed paths, if any
- decisions made
- assumptions
- validation run
- validation not run and why
- residual risk
- blockers or approval needs

## Handoff Template

```text
Return when: <result condition>
Stop if: <approval boundary, scope breach, ambiguity, or blocker>
Close plan: <close after one result | remain available for follow-up>
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
