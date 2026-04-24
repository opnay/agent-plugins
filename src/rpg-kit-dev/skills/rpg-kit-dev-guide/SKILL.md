---
name: rpg-kit-dev-guide
description: Entrypoint skill for `rpg-kit-dev`. Use it to decide whether a task is about role-based subagent orchestration and route to `subagent-role` when role packets, spawn boundaries, answer contracts, or learning notes are needed.
---

# RPG Kit Guide

## Overview

Use this skill as the default entrypoint for `rpg-kit-dev`.
Its job is to decide whether the current task belongs in role-based subagent orchestration.
If the task is about assigning specialty-rich roles to subagents, defining role packets, deciding whether to spawn role-specific agents, or learning from subagent behavior, route into `subagent-role`.

## Routing Rules

- Use `subagent-role` when the user wants role-assigned subagents.
- Use `subagent-role` when the role needs a specific specialty mix such as function, responsibility domain, background expertise, or general expertise.
- Use `subagent-role` when the main work is role packet design, spawn boundaries, answer contracts, or integration rules.
- Use `subagent-role` when the user wants to learn from how role-assigned subagents behave.
- Do not use `rpg-kit-dev` for ordinary implementation, review, or planning when subagent role orchestration is not the main issue.
- Do not spawn subagents from this guide. Route to the narrower skill and preserve runtime/tool approval boundaries.

## Handoff Summary

Before routing to `subagent-role`, summarize:

- orchestration goal
- likely role shape
- specialty mix
- learning focus
- known approval or spawn boundary

## Output

- `RPG fit`
- `Why rpg-kit-dev or not`
- `Chosen entrypoint`
- `Handoff summary`
- `Residual risk`

## Guardrails

- Do not treat subagents as mandatory.
- Do not hide a broader workflow-selection problem inside `rpg-kit-dev`.
- Do not let subagent output bypass caller verification and integration.
