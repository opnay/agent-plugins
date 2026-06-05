---
name: memory-vault
description: Manage a user-provided folder itself as a memory vault with knowledge repository documents and AGENTS.md rules when the user asks for repository memory, persistent project notes, folder knowledge base, memory vault setup, or AGENTS.md memory rules. memory vault, repository memory, project memory, persistent notes, knowledge repository, folder knowledge base, AGENTS.md, 지식 저장소, 메모리 저장소, 폴더 지식, 프로젝트 메모리
---

# Memory Vault

## Purpose

Use this skill when the user provides a folder that should itself hold persistent project knowledge.
Create or maintain the vault documents, category `INDEX.md` files, and a bounded `AGENTS.md` rule section directly in that folder.

## Workflow

1. Identify the target folder.
   - If the user gave a path, use that path.
   - If the user says current folder or this repository, use the current working directory.
   - If the target is unclear, ask one short question before editing.
2. Lock the write allowlist:
   - `<target>/README.md`
   - `<target>/INDEX.md`
   - `<target>/decisions.md`
   - `<target>/glossary.md`
   - `<target>/open-questions.md`
   - `<target>/AGENTS.md`
   - `<target>/<category>/INDEX.md`
   - `<target>/<category>/<subcategory>/INDEX.md`
3. Keep categories to one or two levels.
   - Example: `Programming`
   - Example: `Programming/React`
4. Resolve the helper script path relative to the installed plugin root that contains this skill.
5. Preview changes with the helper script:

```bash
python3 <plugin-root>/scripts/memv.py <target-folder> --category Programming/React --dry-run
```

6. If the preview matches the request and approval boundaries, run:

```bash
python3 <plugin-root>/scripts/memv.py <target-folder> --category Programming/React
```

7. Verify the files exist and report created, updated, and preserved files.

## Existing Vault Maintenance

When the target folder already has vault documents:

- Read `INDEX.md`, `decisions.md`, `glossary.md`, and `open-questions.md` before changing durable knowledge.
- Update `INDEX.md` for current high-level context.
- Move durable decisions to `decisions.md`.
- Move local terms to `glossary.md`.
- Move unresolved issues to `open-questions.md`.
- Use one-level category folders for broad areas and two-level folders for focused topics.
- Keep each category folder's `INDEX.md` as that folder's local map.
- Preserve source wording when facts are uncertain, and mark uncertainty instead of inventing facts.

## AGENTS.md Rules

The helper script owns only the section between:

```text
<!-- memory-vault:start -->
<!-- memory-vault:end -->
```

Keep unmarked `AGENTS.md` content unchanged.
Do not rewrite or delete existing vault documents.
If the user asks to change the rule wording, update the marked section only unless they explicitly approve a broader edit.

## Vault Use

When working in a memory vault folder:

- Read `INDEX.md` before relying on remembered project facts.
- Record durable project decisions in `decisions.md`.
- Record terms and local vocabulary in `glossary.md`.
- Record unresolved items in `open-questions.md`.
- Keep the 1-2 level category map in the root `AGENTS.md`.
- Keep transient task logs out of the vault unless they become durable project knowledge.

## Guardrails

- Do not modify files outside the target folder.
- Do not create a nested `memory-vault/` folder inside the provided target.
- Do not create a vault in a guessed parent directory.
- Do not treat generated vault files as proof that their contents are true.
- Do not commit, push, publish, or release from this skill without separate user approval.
