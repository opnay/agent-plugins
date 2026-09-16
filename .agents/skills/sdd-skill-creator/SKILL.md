---
name: sdd-skill-creator
description: Create or update repository skills with spec-driven discipline. Use when Codex is asked to create, rewrite, regenerate, or revise a skill in this repository and the work must keep user intent, specs, runtime SKILL.md content, and validation evidence separate.
---

# SDD Skill Creator

Use this repository overlay with the canonical $skill-creator workflow. Confirm the owning contract before writing runtime instructions.

## Source of Truth

Read AGENTS.md for repository paths and operating rules, docs/SDD.md for spec and change-record contracts, and the target plugin and skill specs before editing. Load $skill-creator for folder structure, frontmatter, interface metadata, validation, and forward-testing rules.

- Standalone skills: use the user-provided path or the canonical skill-creator default; ask if the target is unclear.

## Spec-First Workflow

1. Identify the artifact, owning plugin, affected siblings, and editable scope.
2. Read or update the owning spec before runtime text.
   - Record durable behavior and responsibility boundaries in the owning spec.
   - Keep user intent separate from implementation notes and temporary evaluation criteria.
   - Keep folder-based user intent in intent.md, not the index or child specs.
   - Apply the Change Spec criteria in docs/SDD.md when recording changes.
   - For repository-local standalone skills without a spec folder, put the durable contract directly in SKILL.md.
3. Write runtime from the current contract.
   - The main agent owns runtime writing and rewriting.
   - When a skill spec changes, rewrite the runtime body from that spec rather than preserving or patching old wording.
   - Old runtime may be reference material, not the source of truth.
   - Do not use prior worker output, suspected findings, deleted specs, conversation memory, or git diff as the writing contract.
   - Treat historical comparisons and recovery checks as separate read-only verification.
4. Update related surfaces.
   - Check plugin boundary, sibling responsibilities, README files, manifest description and prompts, and affected specs.

## Runtime Contract

- Write concise English instructions usable by a fresh agent.
- Use name and description frontmatter. Put selection criteria in description.
- Keep the runtime contract in SKILL.md and bundled resources that actually exist.
- Do not require installed users to read repository-only specs. Repository-local authoring skills may read available owning documents.
- Refer to other installed skills and their resources by plugin-scoped identifiers, following AGENTS.md.
- Keep process notes, private reasoning, temporary context, and unavailable paths out of runtime text.
- Use localized patches only for small metadata fixes or corrections that do not change the skill contract.

## Validation

Run the canonical quick_validate.py against the actual skill folder. Inspect frontmatter, bundled resource targets, identifiers, and alignment with the owning spec.

For nontrivial skills, forward-test with a fresh subagent using the skill path and a realistic task. For spec/runtime alignment, use multiple clean-context read-only verifiers with narrow responsibilities.

Each verifier receives only source-of-truth files, target files, scope, constraints, validation criteria, and the output contract. Do not provide previous conclusions or failure narratives. Verifiers may not rewrite, edit, build, commit, push, release, version-bump, or open PRs.

The main agent reviews verifier evidence, resolves discrepancies, and checks the whole result. Keep writing authority separate from validation and publication authority.

## Report

Report the handled scope, changed locations, checks and results, and unresolved risks. Distinguish local validation from installation, commit, push, and release.
