---
name: advance-codex-guide
description: Entrypoint skill for the `advance-codex` plugin. Use when a task is about improving how Codex is used and you should first classify whether the primary deliverable is a skill, empirical prompt-evaluation workflow, tool-use guidance layer, plugin bundle, session folder convention, commit workflow, subagent handoff gate, or subagent/custom agent definition.
---

# Advance Codex Guide

## Overview

Use this skill as the default entrypoint for the `advance-codex` plugin.
Its job is to classify what part of Codex usage is being improved before implementation starts and route to the right built-in skill.
If the task spans several deliverable types, choose the sequence explicitly instead of mixing them together.

## Workflow

1. Identify the primary reusable artifact:
   - a skill
   - an empirical prompt-evaluation workflow
   - a tool-use guidance layer
   - a plugin
   - a session folder convention
   - a commit workflow
   - a subagent workflow or custom agent
   - a subagent handoff gate
2. Identify whether the task is:
   - new creation
   - revision of an existing artifact
   - packaging several artifacts into one bundle
3. Decide the ownership boundary:
   - skill-only under a skill folder
   - plugin-level bundle under a plugin directory
   - custom-agent definition or usage guidance
4. Route to the narrowest built-in skill that owns the main deliverable.
5. If the request spans several artifact types, define the execution order before editing files.

## Routing Rules

- Choose `skill-creator` when the main output is a reusable skill folder with `SKILL.md` and optional references, scripts, or assets.
- When `skill-creator` applies, pair the canonical system `skill-creator` with the local `advance-codex:skill-creator` extension.
- Choose `empirical-prompt-tuning` when the main output is evidence-based evaluation and revision of a reusable instruction, skill, AGENTS section, or prompt using fresh subagents, fixed scenarios, and caller-side metrics.
- Choose `tool-use-guide` when the main output is reusable guidance for how an artifact should choose, sequence, constrain, or escalate tools without burying tool policy inside a domain workflow.
- Choose `plugin-creator` when the main output is an installable plugin bundle with `.codex-plugin/plugin.json` and optional bundled skills.
- When `plugin-creator` applies, pair the canonical system `plugin-creator` with the local `advance-codex:plugin-creator` extension.
- Choose `agents-sessions` when the main output is defining or explaining the basic purpose and default `turn-gate`-aligned structure of the `.agents/sessions` folder.
- Choose `git-committer` when the main output is a disciplined task-scoped commit workflow for changes you made rather than a new artifact shape.
- Choose `subagent-creator` when the main output is a custom agent definition: how to define `.codex/agents/*.toml`, shape the agent's role, and document how you should use that custom agent from normal work.
- Choose `subagent-gate` when the main output is preparing a subagent handoff before spawn or message: exit plan, minimal context packet, ownership, output contract, and approval limits.
- When building a new multi-skill plugin, prefer this order:
  1. plugin structure
  2. contained skills
  3. tool-use guidance if the plugin needs a reusable tool-policy layer
  4. session folder convention if the plugin needs `.agents/sessions` boundary guidance
  5. commit finalization if the plugin needs stable change finalization rules
  6. subagent guidance if the plugin needs it

## Decision Rules

- Treat plugin packaging as the top-level concern when the user wants discoverability, installable packaging, or several bundled skills.
- Treat skill design as the top-level concern when the user wants a reusable workflow but no plugin packaging.
- Treat empirical prompt tuning as the top-level concern when the hard part is validating or improving instruction quality through repeatable fresh-executor evidence rather than author intuition.
- Treat tool-use design as the top-level concern when the hard part is not the domain workflow itself, but how artifacts should select and use tools consistently.
- Treat agents-sessions as the top-level concern when the hard part is whether an artifact belongs under `.agents/sessions` at all.
- Treat commit finalization as the top-level concern when the hard part is how you should review, verify, message, and split your changes into stable commits.
- Treat subagent design as the top-level concern when the hard part is custom agent behavior, agent reuse, or multi-agent runtime control rather than local file structure.
- Treat subagent handoff gating as the top-level concern when the hard part is deciding what a subagent should receive, when it should return, and which decisions must stay with the main thread.
- If the task is mostly about choosing between several bundled skills, add or update an entrypoint skill.

## Output Contract

- `Primary artifact`
- `Task shape`
- `Chosen creator or gate skill`
- `Execution order`
- `Main risk`
- `Validation path`
- `Residual risk`

## Guardrails

- Do not start by writing files before classifying the artifact.
- Do not bury "when to use" logic in the body of a skill; keep it in the description.
- Do not bury reusable tool policy inside a domain skill when that policy should be isolated.
- Do not treat self-rereading as empirical validation when the real requirement is fresh-executor evidence.
- Do not treat session folder conventions as generic repository hygiene when the real concern is the `turn-gate`-aligned plan/flow record skeleton.
- Do not treat change finalization as generic git hygiene when the real concern is commit-scoped output finalization.
- Do not widen custom-agent work into unrelated orchestration design.
- Do not create a multi-skill plugin without deciding whether it needs an entrypoint skill.
