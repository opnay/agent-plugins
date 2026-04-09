---
name: advance-codex-guide
description: Entrypoint skill for the `advance-codex` plugin. Use when a task involves creating or revising reusable Codex artifacts and Codex should first classify whether the primary deliverable is a skill, a plugin bundle, or a subagent workflow or custom agent definition.
---

# Advance Codex Guide

## Overview

Use this skill as the default entrypoint for the `advance-codex` plugin.
Its job is to classify the artifact being designed before implementation starts and route to the right creator skill.
If the task spans several artifact types, choose the sequence explicitly instead of mixing them together.

## Workflow

1. Identify the primary reusable artifact:
   - a skill
   - a plugin
   - a subagent workflow or custom agent
2. Identify whether the task is:
   - new creation
   - revision of an existing artifact
   - packaging several artifacts into one bundle
3. Decide the ownership surface:
   - skill-only under a skill folder
   - plugin-level bundle under a plugin directory
   - custom-agent definition or usage guidance
4. Route to the narrowest creator skill that owns the main deliverable.
5. If the request spans several artifact types, define the execution order before editing files.

## Routing Rules

- Choose `skill-creator` when the main output is a reusable skill folder with `SKILL.md` and optional references, scripts, or assets.
- Treat the local `skill-creator` as an independent extension over the canonical system `skill-creator`.
- Choose `plugin-creator` when the main output is an installable plugin bundle with `.codex-plugin/plugin.json` and optional bundled skills.
- Treat the local `plugin-creator` as an independent extension over the canonical system `plugin-creator`.
- Choose `subagent-creator` when the main output is a custom Codex agent definition: how to define `.codex/agents/*.toml`, shape the agent's role, and document how that custom agent should be used from Codex workflows.
- When building a new multi-skill plugin, prefer this order:
  1. plugin structure
  2. contained skills
  3. subagent guidance if the plugin needs it

## Decision Rules

- Treat plugin packaging as the top-level concern when the user wants discoverability, installable packaging, or several bundled skills.
- Treat skill design as the top-level concern when the user wants a reusable workflow but no plugin packaging.
- Treat subagent design as the top-level concern when the hard part is custom agent behavior, agent reuse, or multi-agent runtime control rather than local file structure.
- If the task is mostly about choosing between several bundled skills, add or update an entrypoint skill.

## Output Contract

- `Primary artifact`
- `Task shape`
- `Chosen creator skill`
- `Execution order`
- `Main risk`
- `Validation path`
- `Residual risk`

## Guardrails

- Do not start by writing files before classifying the artifact.
- Do not bury "when to use" logic in the body of a skill; keep it in the description.
- Do not widen custom-agent work into unrelated orchestration design.
- Do not create a multi-skill plugin without deciding whether it needs an entrypoint skill.
