---
name: memory-vault
description: Manage a personal agent memory vault for durable cross-task knowledge, indexed knowledge documents, user preferences, environment facts, workflows, terminology, unresolved questions, and reusable knowledge learned while working; use when the user asks to remember something, save knowledge for future work, maintain agent memory, initialize or update memory vault indexes, or when an agent discovers durable knowledge from mistakes, command failures, web research, user instructions, code inspection, or tool behavior. personal agent memory, remember this, save for future, user preferences, durable notes, long-term memory, indexed knowledge docs, agent mistakes, command failure, web research knowledge, persistent knowledge, 에이전트 기억, 장기 기억, 사용자 선호, 지식 문서, 작업 중 습득, 실수 기록, 기억해, 저장해
---

# Memory Vault

## Purpose

Use this skill to read, initialize, or update a personal agent memory vault.
The vault stores durable knowledge learned across agent work, not one project's temporary state.
It is not limited to explicit user save requests. Use it when work reveals reusable facts, failures, constraints, or instructions.

Default vault root: `~/Workspace/Memory-vault`.
Use another target only when the user explicitly provides one.

## Workflow

1. Identify the vault root.
   - If the user gave a path, use that path.
   - Otherwise use `~/Workspace/Memory-vault`.
   - Do not guess a project folder as the vault root.
2. Lock the write allowlist:
   - `<target>/README.md`
   - `<target>/INDEX.md`
   - `<target>/preferences.md`
   - `<target>/decisions.md`
   - `<target>/environment.md`
   - `<target>/workflows.md`
   - `<target>/glossary.md`
   - `<target>/open-questions.md`
   - `<target>/AGENTS.md`
   - `<target>/<category>/INDEX.md`
   - `<target>/<category>/<index>-<slug>.md`
   - `<target>/<category>/<subcategory>/INDEX.md`
   - `<target>/<category>/<subcategory>/<index>-<slug>.md`
3. If the vault exists, read `INDEX.md` first, then only the documents relevant to the current task.
4. Classify memory candidates before writing.
5. Store confirmed durable knowledge even if the user did not explicitly say "remember this."
6. Ask one short question only when the candidate is useful but not confirmed enough to store.
7. Preview structure changes with the helper script:

```bash
python3 <plugin-root>/scripts/memv.py <target-folder> --category Agents/Prompting --dry-run
python3 <plugin-root>/scripts/memv.py <target-folder> --knowledge Programming/React/hook-rules --dry-run
```

8. If the preview matches the request and approval boundary, run the same command without `--dry-run`.
9. Verify files exist and report created, updated, and preserved files.

## Memory Rules

Store only reusable, durable knowledge:

- user preferences for language, tone, response shape, question style, and tool use
- long-lived decisions and operating rules
- local environment facts such as paths, runtimes, package managers, and repeated commands
- recurring workflows, validation routines, and problem-solving patterns
- user-defined terminology, abbreviations, and naming rules
- verified mistakes and command failures with the future action they imply
- recurring permission, sandbox, network, package-manager, or tool constraints
- reusable knowledge from web research, official docs, code inspection, or tests
- reusable knowledge notes that belong under a 1-2 depth category
- unresolved questions that should be asked later

Do not store:

- one-off progress logs
- raw temporary errors or transient state without a reusable rule
- guesses, canceled directions, or unverified claims
- web research that lacks a reliable source or may be too time-sensitive
- secrets, credentials, or sensitive personal data
- anything the user says not to remember

If a candidate is unclear, ask once when a direct answer would make it durable.
Use `open-questions.md` only when the unresolved point is worth revisiting later.

## Document Routing

- `preferences.md`: user preferences and interaction defaults
- `decisions.md`: durable decisions and policies
- `environment.md`: paths, tools, runtimes, package managers, and local setup
- `workflows.md`: repeated procedures, verification patterns, and commands that require escalation
- `glossary.md`: terms, aliases, and abbreviations
- `open-questions.md`: unresolved memory candidates and future questions
- `INDEX.md`: vault root index of core documents and category indexes
- `<category>/INDEX.md`: links to that folder's knowledge documents and child category indexes
- `<category>/<index>-<slug>.md`: first-level category knowledge document
- `<category>/<subcategory>/<index>-<slug>.md`: second-level category knowledge document

Knowledge document filenames must use a 3-digit index plus lowercase hyphen-case slug, for example:

```text
Programming/001-some-knowledge.md
Programming/React/001-hook-rules.md
```

## Index Rules

- Treat each `INDEX.md` as an index, not as the place to store knowledge content.
- Root `INDEX.md` lists core documents and 1-2 depth category indexes.
- Category `INDEX.md` files list knowledge document links and child category index links.
- Keep category paths relative and limited to one or two levels.
- Let `memv.py --knowledge` assign the next 3-digit index in the target folder.
- Preserve existing knowledge documents. If the same slug already exists in a folder, reuse that file instead of creating another numbered copy.

## AGENTS.md Rules

The helper script owns only the section between:

```text
<!-- memory-vault:start -->
<!-- memory-vault:end -->
```

Keep unmarked `AGENTS.md` content unchanged.
The managed section must tell agents to:

- read `INDEX.md` and relevant durable memory documents before related work
- identify reusable memory candidates while working, including mistakes, web research, and user instructions
- route long-term memory candidates by document type
- avoid guesses, temporary logs, canceled directions, and sensitive information
- keep a 1-2 depth category map for indexed knowledge documents

Preserve ordinary memory documents and knowledge documents unless the user asks to edit them directly.
Do not delete existing vault documents.
Do not create a nested `memory-vault/` folder inside the target.

## Guardrails

- Treat generated vault files as structure, not proof that their contents are true.
- Keep updates small and source-grounded.
- Store the durable rule, not the whole failure log. Example: `pnpm install` failed for permissions, so future `pnpm install` runs need an escalation request.
- Prefer allowlisted durable memory over blocklisting every bad case.
- Do not write outside the vault root.
- Do not create category or knowledge paths deeper than two category levels.
- Do not commit, push, publish, release, or version-bump from this skill without separate user approval.
