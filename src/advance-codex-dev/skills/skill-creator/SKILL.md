---
name: skill-creator
description: Extension for the canonical system `skill-creator`. Use whenever you would use `skill-creator`; read the base system skill and this plugin skill together so the base workflow is paired with stricter skill boundaries, plugin-owned skill guidance, and trigger description metadata rules. skill creation, skill update, skill boundary, plugin-owned skill, description metadata, passive trigger
---

# Skill Creator

## Overview

This skill is an extension over the canonical system `skill-creator`.
Whenever the canonical `skill-creator` applies, use this skill alongside it.
Read the system skill first and treat it as the base workflow, then use this skill to enforce stricter boundaries for reusable skills, especially skills owned by a plugin.

## Extension Goals

Use this extension to enforce five things the base workflow does not cover strongly enough:

1. explicit skill ownership and non-goals
2. removal of hidden sibling-skill context
3. clean separation between plugin usage guidance and skill-level behavior
4. separation of runtime-specific tool policy from domain workflow guidance
5. trigger-oriented frontmatter descriptions, especially for passive skills

## Boundary Rules

When creating or revising a skill:

- Define the skill's owned job before writing workflow steps.
- State what the skill does not own when that boundary could be confused.
- Keep each skill understandable as a bounded artifact without assuming hidden context from sibling skills.
- Put cross-skill or plugin-wide usage guidance in the plugin's usage surfaces, not inside a sibling skill unless it directly affects that skill's own behavior.
- Do not use a skill body to compensate for unclear plugin packaging or manifest guidance.
- If the main issue is tool selection, tool sequencing, or escalation policy, consider a separate tool-use guidance artifact instead of burying tool policy in a domain skill.

## Plugin-Owned Skills

For a skill bundled inside a plugin, design the plugin boundary first and the individual skill boundary second.
The skill may mention how it fits the plugin only when that helps the reader apply the skill correctly.
Avoid instructions that require the reader to know sibling skills, development specs, repository-only paths, or unpublished context at runtime.

Before finalizing a plugin-owned skill, check:

- why this behavior belongs in this skill rather than another plugin surface
- whether the skill remains useful when read by itself
- whether plugin-level guidance has been kept out of the skill body
- whether sibling-skill references are optional navigation rather than hidden prerequisites
- whether any runtime-only tool policy should be extracted

## Description Metadata

Treat frontmatter `description` as trigger metadata that is read before the skill body.
The description should first give a short human-readable explanation of what the skill does and when it should be used.
Write it as selection guidance for the skill reader, not as project history or session context.

For passive skills that should apply even when the user does not explicitly request the skill, append a comma-separated list of plain tokens at the end of the `description`.
These are matching tokens, not literal hashtags, so do not prefix them with `#`.

Use passive trigger tokens for:

- inputs where the user is unlikely to name the skill directly
- workflows that should be selected from artifact names, task phrases, or common synonyms
- reusable guardrails that must activate from natural user wording

Do not force passive trigger tokens onto:

- active or direct-call skills the user is expected to request by name
- skills with already narrow and reliable trigger wording
- descriptions where extra tokens would broaden selection into unrelated work

Good passive token lists include the expressions that should actually select the skill: likely user phrases, artifact names, workflow names, and common synonyms.
Keep them narrow.
For example, use a tail like `skill creation, skill update, passive trigger, description metadata` rather than broad words that would match unrelated tasks.

## Review Pass

Before considering a skill done, check:

- whether the `description` can select the skill before the body is loaded
- whether a passive skill has useful comma-separated plain tokens at the end of the `description`
- whether those passive tokens omit `#`
- whether active or already-narrow skills avoid unnecessary token tails
- whether the skill body is written for a runtime reader rather than for a repository change session
- whether development-only specs or unpublished paths have been kept out of runtime instructions
- whether ownership, non-goals, and plugin boundaries are clear

## Output Contract

When this extension materially changes the result, include:

- `Skill boundary`
- `Non-goals`
- `Plugin ownership`
- `Description trigger behavior`
- `Runtime-only guidance removed or avoided`
- `Validation path`
- `Residual risk`

## Guardrails

- Do not replace the canonical `skill-creator` workflow; extend it.
- Do not write skill bodies that depend on hidden sibling context.
- Do not put plugin usage guidance into a skill unless the skill itself owns that behavior.
- Do not add `#` before passive description matching tokens.
- Do not use broad passive tokens to make a skill trigger everywhere.
