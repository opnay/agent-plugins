---
name: session-manager
description: Manage repository-local session artifacts under .agents/sessions/session-id/, including session_record.md for task context and change_record.md for change-delivery notes. Use when Codex needs to create, ensure, inspect, validate, update, or list session-scoped records.
---

# Session Manager

## Overview

Use this skill to manage session-scoped records in `.agents/sessions/<uuid>/`.

The skill supports two artifacts:

- `session_record.md` for task context, runbook, verification, retrospective notes, and recurrence-prevention guardrails when relevant
- `change_record.md` for change-delivery purpose, key changes, verification, review gates, reviewer notes, and residual risks

Read [references/record-formats.md](references/record-formats.md) when you need the exact section layouts or CLI behavior details.

## Use When

- you are starting or resuming a task and need a deterministic session record
- you need to ensure a session-scoped artifact exists before writing to it
- you need to inspect, validate, or list session artifacts under `.agents/sessions`
- you need to write multiline record content safely via `--value-file`
- you are in change-delivery flow and need to maintain `change_record.md`

## Do Not Use When

- the task only needs a normal answer and no working record
- the repository uses a different artifact layout than `.agents/sessions/<uuid>/`
- you need a full skill scaffold rather than a session artifact workflow

## Quick Start

```bash
python3 scripts/controller.py init --artifact session-record
python3 scripts/controller.py ensure --artifact change-record
python3 scripts/controller.py validate --artifact session-record
python3 scripts/controller.py write --artifact change-record --field purpose --value-file ./purpose.txt
python3 scripts/controller.py list --artifact session-record --include-schema
```

## Workflow

1. Run `python3 scripts/controller.py ensure --artifact session-record` at task start.
2. Use `write`, `show`, and `validate` to keep the session record current while working.
3. When change-delivery starts, run `python3 scripts/controller.py ensure --artifact change-record`.
4. Update `change_record.md` with purpose, verification, and review-gate notes as work progresses.
5. Use `list` to inspect existing session artifacts and their detected schemas.

## Artifact Contract

### `session_record.md`

- schema: `session-record`
- format: numbered sections in English

### `change_record.md`

- schema: `change-record`
- format: markdown heading sections in English
- purpose: PR-style session change record and review-gate notes for change delivery

## Script Contract

`scripts/controller.py` supports:

- `init`: create a new artifact template
- `ensure`: return an artifact path, creating it when missing
- `write`: update one logical field with `--value` or `--value-file`
- `show`: print the full artifact or one field value
- `validate`: report `valid` and `detected_schema`, and exit nonzero for malformed records
- `list`: list artifact files under `.agents/sessions`

Key flags:

- `--artifact`: `session-record` or `change-record`
- `--record-path`: operate on an explicit file path
- `--value-file`: read multiline content from a UTF-8 file, or `-` for stdin
- `--include-schema`: include the detected schema column in `list` output

## Mandatory Rules

- Keep generated record content in English.
- Keep `session_record.md` numbered sections stable for machine and human scanning.
- Keep `change_record.md` heading order stable for reviewer readability.
- Prefer `--value-file` for multiline content instead of shell-quoted inline text.
- Treat `show`, `validate`, and `list` as read-only commands.

## Example Triggers

- "Start a session record for this task."
- "Ensure the current session has `change_record.md`."
- "Validate the current session artifacts."
- "Write reviewer notes into `change_record.md`."
- "List session records and show which ones are available."
