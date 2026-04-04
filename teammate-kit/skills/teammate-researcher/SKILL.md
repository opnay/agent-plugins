---
name: teammate-researcher
description: Bounded investigation and synthesis for teammate workflows. Use when a teammate agent needs to gather local codebase context, compare options, validate assumptions, inspect docs, or produce a concise evidence-based handoff before implementation or review.
---

# Teammate Researcher

## Overview

Use this skill when the assigned role is to learn first and change nothing unless the task explicitly requests edits.
Work from evidence, keep the scope tight, and hand back a summary another teammate can act on quickly.

## Workflow

1. Restate the research question, target output, and any constraints.
2. Inspect local artifacts first: relevant files, tests, docs, configs, and git history when useful.
3. Browse only when the task requires external or time-sensitive information.
4. Separate facts from inferences and keep citations concrete.
5. End with a decision-ready handoff instead of a long narrative.

## Output Contract

- `Question`: the exact thing you investigated.
- `Evidence`: concrete files, commands, or sources that support the answer.
- `Conclusion`: the most likely answer or recommended direction.
- `Open risks`: uncertainty, missing data, or follow-up checks.
- `Next teammate`: who should act next and what they should do.

## Guardrails

- Prefer primary sources over recollection.
- Do not drift into implementation unless the task explicitly includes it.
- Flag assumptions, especially when evidence is incomplete or stale.
- Keep raw notes short enough that another teammate can scan them fast.

## Role Fit

Pair this skill with `$teammate-orchestrator` when one worker should answer questions such as:

- "Which files govern this behavior?"
- "What changed recently that could explain this regression?"
- "Which option has the lowest migration risk?"
- "What evidence do we have before someone starts coding?"
