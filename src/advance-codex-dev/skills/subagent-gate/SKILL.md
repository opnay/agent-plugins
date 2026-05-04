---
name: subagent-gate
description: Prepare a subagent handoff before spawning or messaging a subagent. Use when you need to decide the subagent's return point, stop boundary, minimal context packet, ownership, output contract, and approval limits from the subagent's point of view.
---

# Subagent Gate

## Overview

Use this skill before you spawn, message, or redirect a subagent.
The job is to make the handoff bounded enough that the subagent knows when to return, what context matters, what not to do, and what answer shape the main thread needs.

This skill is not for defining `.codex/agents/*.toml` custom agents. Use `subagent-creator` for custom agent definitions.
It is also not an empirical evaluation loop. Use `empirical-prompt-tuning` when the goal is to test reusable instructions across fixed scenarios.

## Workflow

1. Decide why a subagent is needed:
   - parallel sidecar work
   - bounded implementation in a clear write scope
   - read-only exploration
   - clean-context verification
   - role-specific review
2. Write the exit plan before the task prompt:
   - when the subagent should return
   - when it should stop instead of expanding scope
   - what output is required
   - whether it may edit files
   - whether it should be closed after the result
   - whether the main thread is blocked or can continue non-overlapping work
3. Decide whether the main thread is blocked:
   - If the subagent is on the critical path, state exactly what is blocked.
   - If it is sidecar work, state what non-overlapping work the main thread will continue.
4. Build a context packet from the subagent's point of view.
5. Check approval boundaries before delegation.
6. Spawn or message the subagent only after the handoff packet is complete.
7. When the subagent returns, compare the result against the exit plan before integrating it.

## Exit Plan First

Put the exit plan near the top of the subagent prompt.
Include:

- `Return when`: the concrete completion condition.
- `Stop if`: blockers, ambiguity, permission boundaries, or scope expansion triggers.
- `Output`: the exact answer shape needed by the main thread.
- `Edit permission`: whether file changes are allowed.
- `Close plan`: whether the subagent should be closed after one result or reused for a related follow-up.

Do not treat "research this" or "look into this" as a sufficient exit plan.
If the subagent cannot know when to stop, narrow the task before delegating.

If the request is outside this skill's boundary, do not force it into a handoff packet.
Return a no-handoff routing note that names the right skill or workflow, explains why no handoff should be prepared, and states the main thread's next action.

## Context Packet

Rewrite context for the subagent instead of forwarding the main thread's whole context.
Include only what the subagent needs to act:

- User-visible goal.
- Assigned files, directories, or responsibility boundary.
- Relevant facts already verified by commands or file reads.
- Current constraints, repository rules, and approval boundaries.
- Expected output shape.
- Pass/fail criteria or review focus.
- Known non-goals and actions to avoid.

Leave out:

- Unrelated conversation history.
- Main-thread internal reasoning that does not affect the task.
- Guesses presented as facts.
- Decisions that require the user or main thread to approve.
- Broad permission to change unrelated files.

## Delegation Rules

- Delegate concrete, bounded work with a clear material contribution to the current task.
- Prefer sidecar tasks that can run while the main thread does non-overlapping work.
- Keep urgent blocking work local unless waiting for the subagent is still the clearest path.
- For implementation work, assign explicit ownership of files or modules.
- Tell implementation subagents they are not alone in the codebase and must not revert others' edits.
- For verification work, keep the subagent read-only and give pass/fail criteria.
- Do not ask subagents to make user-gated approvals, destructive decisions, external publish decisions, or safety decisions.
- Do not leak the intended answer to a validation subagent unless that is part of the test.

## Prompt Packet Shape

Use this compact shape when preparing a handoff:

```markdown
Return when: <completion condition>
Stop if: <blockers, ambiguity, or approval boundaries>
Output: <required result shape>
Edit permission: <none | specific paths | specific responsibility>
Close plan: <close after result | keep available for related follow-up>
Main-thread blocked state: <blocked on this result | can continue non-overlapping work | no handoff because routing/approval is needed>

Task: <one bounded task>

Context:
- <verified fact or relevant constraint>
- <file/path/surface the subagent owns or should inspect>

Constraints:
- <non-goal or forbidden action>
- <approval boundary>

Pass/fail criteria:
- <what must be true for the result to be usable>
```

## No-Handoff Routing Note

Use this shape when a request is not delegation-ready or is outside this skill's boundary:

```markdown
No handoff prepared.

Reason: <why a subagent prompt would be unsafe, premature, or out of scope>
Right skill or workflow: <subagent-creator | empirical-prompt-tuning | user-gated clarification | main-thread decision | other specific route>
Main-thread next action: <the next action before delegation can happen>
Main-thread blocked state: <blocked on routing/clarification/approval | can continue non-overlapping work>
Re-entry condition: <what must become true before subagent-gate can prepare a handoff>
```

Do not add ownership, edit permission, or pass/fail criteria for a subagent when no subagent should be handed off yet.
For no-handoff cases, the integration-ready output is the routing note itself.

## Review Questions

- Is the return point stated before the task details?
- Does the prompt say when to stop, not only what to do?
- Does it explicitly state whether the main thread is blocked or can continue non-overlapping work?
- If no handoff is prepared, does the routing note state reason, route, next action, blocked state, and re-entry condition?
- Is the context rewritten for the subagent's point of view?
- Are verified facts separated from assumptions?
- Are write scope and ownership explicit when edits are allowed?
- Are user-gated, destructive, irreversible, external, or safety decisions kept with the main thread?
- Is the expected output specific enough to integrate without follow-up?

## Output Contract

- `Subagent purpose`
- `Exit plan`
- `Context packet`
- `Delegation boundary`
- `Main-thread blocked state`
- `Subagent prompt`
- `Residual risk`

## Guardrails

- Do not spawn first and figure out the stop condition later.
- Do not pass the whole parent context when a task-local packet is enough.
- Do not leave the main thread's blocked state implicit.
- Do not ask a subagent to approve what only the user can approve.
- Do not use a subagent to hide uncertainty that should be resolved with a user-gated question.
- Do not let a validation subagent implement fixes or expand scope.
- Do not keep a subagent open without a reason tied to the next bounded handoff.
