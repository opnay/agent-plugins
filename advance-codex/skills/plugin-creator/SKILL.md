---
name: plugin-creator
description: Extension for the canonical system `plugin-creator`. Use whenever Codex would use `plugin-creator`; read the base system skill and this plugin skill together so the base workflow is paired with top-down plugin design, plugin and skill independence, and mandatory `<plugin>-guide` entrypoint handling.
---

# Plugin Creator

## Overview

This skill is an extension over the canonical system `plugin-creator`.
Whenever the canonical `plugin-creator` applies, use this skill alongside it.
Read the system skill first and treat it as the base workflow, then use this skill to enforce stronger plugin-architecture rules for how a plugin should be shaped and how its bundled skills should relate to it.

## Extension Goals

Use this extension to enforce four things the base workflow does not cover strongly enough:

1. top-down design from plugin boundary to bundled skills
2. independence of the plugin as a coherent bundle
3. independence of each bundled skill as a bounded artifact
4. mandatory `<plugin>-guide` entrypoint handling for plugin usability

## Added Rules

1. Design plugins top-down:
   - define the plugin boundary first
   - define the bundled skills second
2. Do not start from a loose collection of skills and only later invent the plugin shape around them unless the migration is explicit and intentional.
3. Every plugin should remain independently understandable as one coherent bundle with a clear purpose.
4. Every bundled skill should remain independently usable for its own bounded job and should not depend on hidden sibling-skill context.
5. When creating a plugin, also create a `<plugin>-guide` entrypoint skill for that plugin.
6. Use `<plugin>-guide` to explain how to use the plugin well, route users to the right bundled skill, and classify multi-capability tasks before deeper execution.
7. Keep the manifest aligned with the actual bundled surfaces. Do not promise skills, MCP, or app metadata that the plugin does not ship.

## Review Pass

Before considering a plugin done, check:

- what the plugin owns as one coherent bundle
- whether the plugin boundary was defined before the bundled skill set
- why each bundled skill belongs inside this plugin
- whether a `<plugin>-guide` exists and actually improves usability
- whether plugin-level guidance is separate from skill-level guidance
- whether the manifest describes the plugin as it really exists

## Output Contract

- `Plugin boundary`
- `Why this is one plugin`
- `Bundled skills`
- `<plugin>-guide role`
- `Manifest alignment`
- `Validation path`
- `Residual risk`

## Packaging Rules

- Keep bundled guidance focused on the plugin's scope.
- Prefer narrow additions over paraphrasing general system guidance.
- Put cross-skill usage guidance in `<plugin>-guide`, not scattered across sibling skills.

## Guardrails

- Do not start from a loose pile of skills and retrofit a plugin around them unless that migration is explicit.
- Do not use one plugin to hold unrelated skills that do not form one coherent bundle.
- Do not push plugin-level routing or usage guidance down into sibling skills when it belongs in `<plugin>-guide`.
