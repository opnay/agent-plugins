---
name: frontend-kit-guide
description: Entrypoint skill for the `frontend-kit` plugin. Use when a task may belong to this plugin and you need to decide whether the main problem is axis separation inside React-facing artifacts.
---

# Frontend Kit Guide

## Overview

Use this skill as the default entrypoint for `frontend-kit`.
Its job is to decide whether the request belongs to this plugin at all and, if it does, route the work into `react-architecture`.

## Workflow

1. Check whether the request is about React-facing artifacts such as components, hooks, utilities, pages, providers, or broader state.
2. Check whether the real problem is concern separation by cause of change rather than top-level folder choice.
3. If the task is about `system`, `design`, and `domain` separation, route to `react-architecture`.
4. If the task is still about top-level frontend structure choice or a broader task-framing question, say that this plugin is not the right first stop.

## Routing Rules

- Choose `react-architecture` when the main problem is how React-facing code should split `system`, `design`, and `domain` responsibilities.
- Keep this plugin out of top-level structure-pattern choice.
- Keep this plugin out of pure UI-polish or business-rule-definition work unless the core issue is still axis leakage inside React code.

## Output Contract

- `Task shape`
- `Plugin fit`
- `Chosen skill`
- `Why this route fits`
- `Residual risk`

## Guardrails

- Do not pretend every frontend task belongs to this plugin.
- Do not bury plugin-level routing guidance inside `react-architecture`.
- Do not classify by file name alone before checking the real cause of change.
