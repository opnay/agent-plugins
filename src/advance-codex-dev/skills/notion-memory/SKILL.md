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

This skill does not own:

- general Notion workspace automation
- credentials, tokens, cookies, or connector auth state storage
- automatic logging of every conversation
- unsupported claims as facts
- commit, push, PR, release, or unrelated project edits

## Routing

1. If the user asks for setup, preparation, configuration, schema work, workspace rules, or verification, read `references/command-setup.md`.
2. If the user asks to record or update memory, read `references/notion-memory-contract.md`.
3. If one request includes setup and recording, complete the required setup path first, then record only after config and schema are usable.
4. If config is missing for a record request, route to `setup config` and ask for the DB URL or ID instead of guessing.

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

Use connected Notion tools when available:

1. Fetch the configured DB before schema-sensitive work.
2. Add only missing required properties unless the user explicitly approves another change.
3. Create or update pages with the property and body contract.
4. Fetch the page or DB when needed to verify properties, relations, and schema.
5. Report failed tool calls, skipped checks, and residual risk.

If connector writes fail or Notion tools are unavailable, use browser UI when available, or provide exact manual setup steps. Do not call a manual fallback a completed setup.

## Completion

For setup:

- config exists or the missing input is identified
- DB access is verified or the blocker is explicit
- required schema is verified or fallback steps are reported
- workspace rule changes are applied only inside user-approved scope

For recording:

- content passes allow-list and block-list checks
- page properties and body follow the contract
- write/update is verified by tool output plus fetch when needed
- skipped verification and remaining risk are visible

## References

- Read `references/command-setup.md` for setup command handling.
- Read `references/notion-memory-contract.md` for schema, property, title, body, relation, and write rules.
