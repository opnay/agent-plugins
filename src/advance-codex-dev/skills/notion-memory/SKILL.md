---
name: notion-memory
description: "Set up and use a Notion-backed agent work memory for Codex work history, decisions, follow-ups, verification, and reusable work knowledge. Use when the user asks to prepare, configure, verify, or record Notion memory, including setup requests such as 셋업하자, 스킬을 사용하기 위해 준비해줘, Notion 메모리 기록 쓸 수 있게 준비해줘, 작업 히스토리 기록, 결정 기록, 후속 작업 기록, Notion memory, agent memory, work history memory"
---

# Notion Memory

## Owner

Use this skill to prepare and use a Notion-backed memory DB for Codex work.
Keep the scope narrow: setup, schema, workspace rule support, capture, and verification for agent work memory.

This skill owns:

- setup for a configured Notion memory DB
- capture of work history, decisions, follow-ups, verification, reusable work knowledge, and explicit preferences
- config guidance for `~/.agents/configs/notion-memory.toml`
- allow-list first review and block-list security review
- Notion write verification and manual fallback reporting
- Notion CLI guidance for using `ntn` as the first-choice Notion I/O path for memory work

This skill does not own:

- general Notion workspace automation
- general `ntn` CLI or Notion API automation
- delegation of normal memory work to a separate `ntn` or `notion-cli` skill
- default use of a Notion plugin connector for memory I/O
- credentials, tokens, cookies, or connector auth state storage
- automatic logging of every conversation
- unsupported claims as facts
- commit, push, PR, release, or unrelated project edits

## Routing

1. If the user asks for setup, preparation, configuration, schema work, workspace rules, or verification, read `references/command-setup.md`.
2. If the user asks to record or update memory, read `references/notion-memory-contract.md`.
3. Before Notion I/O for setup, recording, or verification, read `references/ntn-cli.md` and use `ntn` first.
4. If one request includes setup and recording, complete the required setup path first, then record only after config and schema are usable.
5. If config is missing for a record request, route to `setup config` and ask for the DB URL or ID instead of guessing.

Semantic setup phrases include `setup`, `셋업하자`, `준비해줘`, `스킬을 사용하기 위해 준비해줘`, and `Notion 메모리 기록 쓸 수 있게 준비해줘`.

## Scope Map

Before writing or changing setup, state the effective scope to yourself:

- user scope: what the user asked to prepare or record
- Notion scope: the configured memory DB only
- local scope: config file, workspace rule, and temporary notes needed for setup
- excluded scope: unrelated DBs, unrelated projects, credentials, commits, pushes, PRs, and releases

## Allow-List First

Record only content that fits at least one allowed class:

- confirmed work history
- user decisions
- follow-up tasks
- verification results
- reusable work knowledge
- explicit user preferences

Then apply the block-list:

- credentials, tokens, cookies, auth state
- unnecessary personal data
- unrelated conversation
- speculation or unsupported interpretation
- private Notion DB URLs or IDs copied into plugin files

Mark uncertainty as `확인 필요` or `미검증`; do not fill gaps.

## Config

Use `~/.agents/configs/notion-memory.toml` for user-specific Notion memory setup.
Split settings into `general` and `notion` sections.
Use `general.timezone` for the timezone.
Use `general.language`, `general.locale`, and `general.register` for memory document writing preferences.
Use `notion.db_url`, `notion.db_id`, `notion.data_source_id`, `notion.view_id`, and `notion.schema_status` for Notion setup.
It must never store credentials, tokens, cookies, or connector auth state.

Use `scripts/setup_config.py` for config creation, template output, or validation when local script execution is available.
The script does not change Notion schema or write memory pages.

When writing the config requires filesystem approval, request it. If writing is not allowed, return the TOML content for the user to place there.

## Notion Workflow

Use `ntn` as the first-choice Notion I/O path:

1. Confirm auth without printing secrets: `ntn whoami` or a non-secret token presence check.
2. Inspect live syntax with `ntn --help`, `ntn api --help`, and endpoint help/spec before schema-sensitive work.
3. Fetch the configured DB or data source through `ntn` before schema-sensitive claims.
4. Add only missing required properties unless the user explicitly approves another change.
5. Create or update pages with fixed page properties and a body format chosen before writing.
6. Fetch the page or DB through `ntn` when needed to verify properties, relations, and schema.
7. Report failed `ntn` calls, skipped checks, and residual risk.

Do not use a separate CLI skill or Notion plugin connector as the default path for this skill.
If `ntn` is unavailable or blocked, report the blocker and ask before using browser UI or manual fallback.
Do not call a manual fallback a completed setup.

## Completion

For setup:

- config exists or the missing input is identified
- DB access is verified or the blocker is explicit
- required schema is verified or fallback steps are reported
- workspace rule changes are applied only inside user-approved scope

For recording:

- content passes allow-list and block-list checks
- page properties follow the fixed contract
- body format was chosen before writing and preserves required distinctions
- write/update is verified by tool output plus fetch when needed
- skipped verification and remaining risk are visible

## References

- Read `references/command-setup.md` for setup command handling.
- Read `references/notion-memory-contract.md` for schema, property, title, body format, relation, and write rules.
- Read `references/ntn-cli.md` before Notion I/O; it owns the `ntn` first path for this skill.
