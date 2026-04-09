---
name: subagent-creator
description: Create or revise custom Codex agents. Use when Codex needs to define `.codex/agents/*.toml` files, choose personal or project scope for a custom agent, write `name`, `description`, and `developer_instructions`, set optional inherited defaults such as model or sandbox, or document how that custom agent should be used from normal Codex workflows.
---

# Subagent Creator

## Overview

Use this skill to design reusable custom Codex agents.
The output can be:

- a custom agent spec under `.codex/agents/*.toml`
- a custom-agent template or starter file
- a custom-agent review checklist
- guidance for how that custom agent should be invoked from normal Codex work

Reference basis:

- OpenAI Codex subagents docs: https://developers.openai.com/codex/subagents

## Workflow

1. Identify the stable role the custom agent should own.
2. Decide the scope of installation:
   - personal scope under `~/.codex/agents/`
   - project scope under `.codex/agents/`
3. Define the minimum required TOML fields:
   - `name`
   - `description`
   - `developer_instructions`
4. Define optional fields only when they materially improve the role:
   - `nickname_candidates`
   - `model`
   - `model_reasoning_effort`
   - `sandbox_mode`
   - `mcp_servers`
   - `skills.config`
5. Write `description` as user-facing guidance for when this custom agent should be used.
6. Write `developer_instructions` as the main behavioral contract for the custom agent.
7. Shape `developer_instructions` around:
   - role identity
   - primary objective
   - scope boundaries
   - preferred working style
   - output expectations
   - escalation or stop conditions
8. Review the instructions for naming clarity, role overlap, and unnecessary config overrides.
9. Document how the custom agent should be used from normal Codex work, including the name that will be passed when spawning it.

Read [references/custom-agent-template.md](references/custom-agent-template.md) while drafting the custom agent file.

## Core Facts

- Built-in agents are `default`, `worker`, and `explorer`.
- Custom agents live in standalone TOML files under `~/.codex/agents/` for personal use or `.codex/agents/` for project scope.
- The required custom-agent fields are:
  - `name`
  - `description`
  - `developer_instructions`
- Optional fields such as `nickname_candidates`, `model`, `model_reasoning_effort`, `sandbox_mode`, `mcp_servers`, and `skills.config` inherit from the parent session when omitted.
- Child agents inherit the current sandbox and approval state from the parent run, even when the custom agent file sets defaults.
- A custom agent is used by name when spawning it from normal Codex workflows.

## Design Rules

- Keep custom agents narrow and opinionated. Give each one a clear job and instructions that keep it from drifting into adjacent work.
- Prefer creating one strong custom agent for one recurring role over one broad agent that tries to cover several unrelated jobs.
- Match the tool surface and sandbox mode to the role:
  - exploration and review style agents should usually be read-only
  - implementation agents should have the smallest write surface that still lets them finish the task
- Use optional config only when the role genuinely benefits from a stable override.
- Keep `description` human-facing and routing-friendly.
- Treat `developer_instructions` as the most important field after `name`.
- Keep `developer_instructions` procedural, role-specific, and explicit about what good execution looks like.
- Use `developer_instructions` to define:
  - what the agent is for
  - what it should prioritize
  - what it should avoid widening into
  - what kind of output it should produce
- Prefer concrete behavioral guidance over abstract mission statements.
- If a custom agent name collides with a built-in agent name, the custom agent takes precedence.

Read [references/custom-agent-checklist.md](references/custom-agent-checklist.md) before finalizing the custom agent.

## Custom Agent File Rules

- Match the filename to the agent name when practical, but treat the `name` field as the source of truth.
- Write the `description` as human-facing guidance for when the agent should be used.
- Write `developer_instructions` to keep the role narrow, stable, and resistant to scope drift.
- Structure strong `developer_instructions` around these elements when relevant:
  - identity: who the agent is
  - objective: what success means
  - scope: what belongs to the role
  - non-goals: what should stay outside the role
  - method: how the agent should approach the work
  - output: what shape the result should take
- If the role is review-oriented, tell the agent what kinds of findings matter.
- If the role is implementation-oriented, tell the agent what boundaries it owns and what it must preserve.
- If the role is research-oriented, tell the agent what evidence quality and synthesis depth are expected.
- Use `nickname_candidates` only for display clarity; they do not affect spawning identity.
- Keep `nickname_candidates` unique and human-readable.

## Deliverable Shapes

Common outputs include:

- a custom agent TOML spec
- a custom-agent starter template
- a custom-agent review checklist
- usage notes for how to call the custom agent from Codex workflows

## Validation

- Verify that required TOML fields are present.
- Verify that the role is narrower than a general-purpose fallback.
- Verify that optional config fields are justified rather than copied by habit.
- Verify that the selected scope, personal or project, matches the intended reuse pattern.
- Verify that the agent can be named and invoked clearly from Codex workflows.
- Verify that `developer_instructions` contain enough operational detail to distinguish the agent from a vague prompt wrapper.
- Verify that `developer_instructions` define both what the agent should do and what it should stay away from.

## Guardrails

- Do not define a custom agent that is so broad it becomes another general-purpose fallback agent.
- Do not use optional config as decoration.
- Do not hide the real role of the agent behind a vague `description`.
- Do not let one custom agent silently absorb several unrelated recurring jobs.
- Do not write `developer_instructions` as a slogan, persona blurb, or vague aspiration.
