---
name: turn-gate
description: Meta workflow for repositories that require explicit turn continuity across question, plan, command, execution, result reporting, and next-flow prompting. Use when the main job is to keep the phase loop alive instead of ending after a single answer or execution pass.
---

# Turn Gate

## Overview

Use this skill when the repository or working agreement requires one user turn to continue across several internal phases.
Its job is not to replace downstream workflow skills.
Its job is to keep the loop explicit:

1. classify the current message
2. choose the right downstream owner for this phase
3. execute that phase
4. report the result
5. decide whether another flow should start immediately
6. if so, ask for the narrowest next-step input and continue

This skill is a meta workflow.
It owns turn continuity, not the domain work inside each phase.

## Use When

- the repository requires question / plan / command style phase separation inside one ongoing turn
- the task should not naturally stop after one answer because another explicit flow is expected right away
- result reporting must be followed by an explicit next-flow decision instead of a soft closing
- the main risk is premature turn termination rather than lack of a phase-specific workflow

## Do Not Use When

- the task is a normal single-phase request that can end cleanly after one answer
- a narrower workflow already fully owns the problem and no explicit next-flow loop is needed
- the main blocker is still choosing between clarification, planning, or execution rather than managing turn continuity itself

## Scope Boundary

This skill owns:

- turn-level phase classification
- downstream workflow selection for the current phase
- result reporting with explicit loop assessment
- next-flow prompting when another phase should begin immediately

This skill does not own:

- requirements interviewing itself
- read-only planning itself
- implementation itself
- review-loop handling itself
- commit finalization itself

## Core Policy

- Treat each incoming message as at least one of: question, plan, command.
- Choose the narrowest downstream workflow that owns the current phase.
- Do not let result reporting become a soft stop when the work obviously expects another phase.
- Ask for the next flow only when another phase is genuinely needed.
- Prefer the structured user-input tool when the next-flow decision fits 1-3 bounded choices.
- Keep the loop moving; do not reopen broad framing once the next phase is already clear.

## Phase Loop

### Phase 0: Analyze The Current Message

1. Separate the message into question, plan, command, or mixed concerns.
2. Decide which concern is blocking the next useful action.
3. Choose the downstream workflow for that concern.

Typical downstream choices:

- `structured-thinking`
- `deep-interview`
- `planner`
- `autopilot`
- `parallel-work`
- `ralph-loop`
- `review-loop`
- `commit-readiness-gate`
- a specialist plugin after the workflow phase is clear

Output:

- `Current phase type`
- `Chosen downstream owner`

### Phase 1: Run The Current Phase

1. Hand off to the selected downstream workflow.
2. Keep the current phase bounded.
3. Do not mix the next phase into the current one unless the downstream workflow itself requires it.

Output:

- `Phase result`
- `Phase residual risk`

### Phase 2: Report And Assess Continuation

1. Report what changed, what was decided, or what remains blocked.
2. Decide whether another explicit flow is expected immediately.
3. If not, stop cleanly.
4. If yes, define the narrowest next-flow choice.

Continuation signals:

- the current phase completed but naturally hands off to another workflow
- the user-local operating rule requires explicit next-flow choice instead of soft closure
- the result is actionable only if a next phase starts right away

Output:

- `Current result`
- `Next-flow need`

### Phase 3: Open The Next Flow

1. If another flow is needed, ask for it explicitly.
2. Prefer `request_user_input` when the next step fits bounded choices.
3. Route the answer back into Phase 0 instead of ending the turn conceptually.

Output:

- `Next-flow question`
- `Planned loop re-entry`

## Output Contract

- `Current phase type`
- `Chosen downstream owner`
- `Phase result`
- `Current result`
- `Next-flow need`
- `Next-flow question`
- `Residual risk`

## Guardrails

- Do not replace a phase-specific workflow with vague meta commentary.
- Do not keep the loop open when the work can end cleanly.
- Do not ask a next-flow question when the next phase is already obvious and can be started directly.
- Do not let this skill absorb domain execution, planning, or review detail.
- Do not confuse turn continuity with endless conversation.
