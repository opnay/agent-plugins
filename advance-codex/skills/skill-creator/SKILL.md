---
name: skill-creator
description: Extension for the canonical system `skill-creator`. Use whenever Codex would use `skill-creator`; read the base system skill and this plugin skill together so the base workflow is paired with tighter skill boundaries, stronger scope control, and conditional guidance for skills that ship inside plugins.
---

# Skill Creator

## Overview

This skill is an extension over the canonical system `skill-creator`.
Whenever the canonical `skill-creator` applies, use this skill alongside it.
Read the system skill first and treat it as the base workflow, then use this skill to make stronger design decisions about skill boundary, scope, and packaging.

## Extension Goals

Use this extension to enforce four things the base workflow does not cover strongly enough:

1. independence of each skill as its own bounded artifact
2. tighter scope control for what belongs inside one skill
3. cleaner separation between skill-level and plugin-level guidance
4. conditional rules for skills that ship inside plugins

## Added Rules

1. Every skill must remain independently understandable and usable for its own bounded job.
2. Keep the skill boundary explicit: what the skill owns, what it does not own, and what should stay outside it.
3. Do not let one skill rely on hidden behavioral context from sibling skills or nearby artifacts.
4. Keep companion skills lean. Add only the extra rules, references, or assets that materially improve behavior over the base system skill.
5. If the main problem is reusable tool-selection or tool-escalation policy, prefer a dedicated tool-use artifact over stuffing runtime-specific rules into a domain skill.

## Skills Inside Plugins

Apply this section only when the skill is meant to ship inside a plugin.

1. Follow top-down design order:
   - define the plugin boundary first
   - define the skills that belong inside that plugin second
2. Do not design a plugin-owned skill in isolation first and retrofit the plugin around it later unless that migration is explicit.
3. Place the skill under that plugin's `skills/<skill-name>` directory after the plugin boundary is clear.
4. Put cross-skill usage guidance in the plugin's `<plugin>-guide`, not inside one sibling skill.
5. If a plugin contains several user-facing skills, check whether it also needs a `<plugin>-guide` entrypoint skill before finalizing the skill set.

## Review Pass

Before considering a skill done, check:

- what the skill owns and what it does not own
- whether the skill can be invoked independently without hidden context
- whether the skill is repeating general-purpose system guidance instead of adding new value

If the skill ships inside a plugin, also check:

- whether the plugin boundary was defined before the skill boundary
- whether any guidance belongs in a `<plugin>-guide` entrypoint instead of this skill
- whether the skill is leaking plugin-level concerns into a bounded skill
- whether tool-use policy should live in a dedicated tool-use artifact instead of this skill

## Output Contract

- `Skill boundary`
- `Why this is a separate skill`
- `Plugin relationship` when applicable
- `Need for <plugin>-guide`
- `Added resources`
- `Validation path`
- `Residual risk`

## Guardrails

- Do not use plugin membership as an excuse for a weak or coupled skill boundary.
- Do not let one skill absorb responsibilities that should become a separate skill.
- Do not hide cross-skill usage guidance inside a bounded skill when it belongs in a `<plugin>-guide`.
- Do not mix domain workflow guidance with reusable runtime tool policy when those can be separated cleanly.
