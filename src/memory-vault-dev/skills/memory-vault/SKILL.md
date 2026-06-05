---
name: memory-vault
description: Manage a personal agent memory vault for durable cross-task knowledge, user preferences, environment facts, workflows, terminology, and unresolved questions; use when the user asks to remember something, save knowledge for future work, maintain agent memory, or update long-term notes. personal agent memory, remember this, save for future, user preferences, durable notes, long-term memory, memory vault, persistent knowledge, 에이전트 기억, 장기 기억, 사용자 선호, 기억해, 저장해
---

# Memory Vault

## Purpose

Use this skill to read, initialize, or update a personal agent memory vault.
The vault stores durable knowledge that should apply across tasks, not one project's temporary state.

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
   - `<target>/<category>/<subcategory>/INDEX.md`
3. If the vault exists, read `INDEX.md` first, then only the documents relevant to the current task.
4. Classify memory candidates before writing.
5. Ask one short question when a candidate is useful but not confirmed enough to store.
6. Preview structure changes with the helper script:

```bash
python3 <plugin-root>/scripts/memv.py <target-folder> --category Agents/Prompting --dry-run
```

7. If the preview matches the request and approval boundary, run the same command without `--dry-run`.
8. Verify files exist and report created, updated, and preserved files.

## Memory Rules

Store only reusable, durable knowledge:

- user preferences for language, tone, response shape, question style, and tool use
- long-lived decisions and operating rules
- local environment facts such as paths, runtimes, package managers, and repeated commands
- recurring workflows, validation routines, and problem-solving patterns
- user-defined terminology, abbreviations, and naming rules
- unresolved questions that should be asked later

Do not store:

- one-off progress logs
- temporary errors or transient state
- guesses, canceled directions, or unverified claims
- secrets, credentials, or sensitive personal data
- anything the user says not to remember

If a candidate is unclear, record it in `open-questions.md` only when it is worth resolving later; otherwise skip it.

## Document Routing

- `preferences.md`: user preferences and interaction defaults
- `decisions.md`: durable decisions and policies
- `environment.md`: paths, tools, runtimes, package managers, and local setup
- `workflows.md`: repeated procedures and verification patterns
- `glossary.md`: terms, aliases, and abbreviations
- `open-questions.md`: unresolved memory candidates and future questions
- `INDEX.md`: high-level vault summary, document map, and category map
- `<category>/INDEX.md`: one or two level topic maps, for example `Agents/Prompting`

## AGENTS.md Rules

The helper script owns only the section between:

```text
<!-- memory-vault:start -->
<!-- memory-vault:end -->
```

Keep unmarked `AGENTS.md` content unchanged.
Do not delete existing vault documents.
Do not create a nested `memory-vault/` folder inside the target.

## Guardrails

- Treat generated vault files as structure, not proof that their contents are true.
- Keep updates small and source-grounded.
- Prefer allowlisted durable memory over blocklisting every bad case.
- Do not commit, push, publish, release, or version-bump from this skill without separate user approval.
