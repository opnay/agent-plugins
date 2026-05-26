---
name: sdd-skill-creator
description: Create or update repository skills with spec-driven discipline. Use when Codex is asked to create, rewrite, regenerate, or revise a skill in this repository and the work must keep user intent, specs, runtime SKILL.md content, and validation evidence separate.
---

# SDD Skill Creator

## Overview

Use this skill as an overlay on the canonical `$skill-creator` workflow. Keep skill work in this repository spec-first: understand intent, update or confirm the owning spec, then rewrite the runtime `SKILL.md` from that spec instead of from loose conversation memory or old runtime text.

## Required Inputs

Before editing, read only the repository guidance needed for the target:

- `AGENTS.md` for repository-wide rules, language expectations, and artifact boundaries.
- `docs/SDD.md` when the target skill belongs to a spec-driven plugin or the user explicitly requests SDD behavior.
- The target plugin or skill spec when one exists.
- The current target `SKILL.md` only after the intended contract is clear.

Also load `$skill-creator` and follow its folder, frontmatter, interface metadata, validation, and forward-testing rules unless a stricter repository rule applies.

## Workflow

1. Identify the artifact type.
   - If the user asks for a standalone skill, use the path the user gives. If no path is given, ask once or use the canonical `$skill-creator` default.
   - If the skill is plugin-owned, edit the dev plugin source under `src/<plugin-name>-dev`; do not edit the root release plugin directly.

2. Separate conversation context from artifact content.
   - Treat user messages, questions, and answers as intent evidence.
   - Do not copy process notes, private reasoning, or temporary evaluation criteria into runtime skill text.
   - Put durable behavior into the owning spec first when the repository has an owning spec surface.

3. Confirm or write the spec before runtime text.
   - For plugin-owned skills, update the related skill spec and change spec before rewriting `SKILL.md`.
   - For standalone repository skills without a formal spec folder, write the durable contract directly in `SKILL.md`, but still keep transient session details out.
   - If the user says the skill should be regenerated or rewritten from spec, rewrite the runtime body from the current spec rather than patching around old wording.

4. Rewrite runtime skills from the current spec.
   - When a plugin-owned skill spec changed and runtime skill writing is needed, the main agent first reviews the current spec and locks the source of truth.
   - The main agent owns the runtime rewrite. Do not delegate runtime writing or regeneration to a clean-context subagent.
   - Use the current spec as the source of truth and treat old runtime text as optional reference, not as preserved source.
   - Do not use prior worker output, main-thread suspected findings, old deleted specs, or `git diff` as writing source.
   - Treat old spec or diff comparison as a separate read-only verification or repair flow, not as the runtime writing source.
   - For plugin-owned skills, edit dev source only. The main agent owns release build output, final integration review, and any commit workflow.

5. Generate the runtime skill.
   - Keep frontmatter to `name` and `description`.
   - Put trigger criteria in `description`; do not rely on a body-only "when to use" section.
   - Keep body instructions concise, imperative, and usable by a fresh Codex instance.
   - Reference bundled resources only when they exist in the delivered skill folder.
   - Avoid telling runtime users to read dev-only specs unless the skill is explicitly repository-local and those files are available at runtime.
   - Treat the runtime surface as the whole skill folder when needed: `SKILL.md`, `references/`, `templates/`, and metadata that ships with the skill.

6. Validate and forward-test.
   - Run the canonical `quick_validate.py` against the skill folder.
   - For nontrivial skills, forward-test with a fresh subagent using only the skill path and a realistic task prompt.
   - For spec/runtime alignment, use multiple clean-context verifier subagents with narrow read-only scopes. They verify; they do not write, rewrite, build, commit, push, PR, release, or version-bump.
   - Review generated artifacts yourself before reporting completion.

## Runtime Rewrite Rule

When a skill body is rewritten from a spec, prefer this order:

1. Read the current spec and any owned child specs.
2. Decide whether the runtime surface is only `SKILL.md` or the whole skill folder.
3. Rewrite the runtime surface from the current spec contract.
4. Delete or fully replace the old runtime surface inside the editable scope.
5. Check that no conversation-only notes, prior-worker context, `git diff` recovery framing, old deleted spec paths, or unavailable dev paths leaked into runtime text.
6. Validate the skill folder and inspect the resulting files.
7. Run read-only clean-context verifier subagents for spec/runtime alignment when the change is nontrivial or repository rules require it.

Use patch editing only for small metadata fixes or localized corrections that do not change the skill's contract.

## Report

Report the handled scope, changed files, validation commands, and remaining risks. If plugin release surfaces were intentionally not touched, say that explicitly.
