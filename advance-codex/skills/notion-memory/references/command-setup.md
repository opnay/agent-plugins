# Notion Memory Setup Commands

Use this reference for setup, preparation, config, schema, workspace-rule, or verification requests.

## Command Router

Map semantic setup requests to the same command family:

| Request | Command |
| --- | --- |
| `setup`, `셋업하자`, `준비해줘`, `스킬을 사용하기 위해 준비해줘` | `setup` |
| config path, DB URL/ID, machine setup | `setup config` |
| property, schema, data source, view | `setup schema` |
| AGENTS, workspace rule, folder rule | `setup workspace-rule` |
| verify, check, test setup | `setup verify` |

`setup` means:

1. `setup config`
2. `setup schema`
3. `setup workspace-rule`
4. `setup verify`

Stop at the first blocker that prevents the next step. Report the blocker and the exact next input or approval needed.

## `setup config`

Goal: create or update `~/.agents/configs/notion-memory.toml` without storing secrets.

Preferred script:

```bash
python3 scripts/setup_config.py --db-url "<notion database url>"
python3 scripts/setup_config.py --check
python3 scripts/setup_config.py --print-template
```

Use the script when it is available and local file writes are allowed.
Use `--dry-run` or `--print-template` when the user only wants the TOML content.
If the script fails, report the command, error, and manual TOML fallback.
Do not use the script for Notion schema changes, DB fetches, or page writes.

Allowed config fields:

```toml
[general]
timezone = "Asia/Seoul"
language = "ko"
locale = "ko-KR"
register = "polite"

[notion]
db_url = ""
db_id = ""
data_source_id = ""
view_id = ""
schema_status = "unverified"
```

Rules:

- If config exists, read it before asking for input.
- If `notion.db_url` and `notion.db_id` are both missing, ask the user for one of them.
- Derive DB ID from URL only when the format is clear.
- Never store credentials, tokens, cookies, workspace auth, or connector state.
- Request filesystem approval if writing under `~/.agents/configs/` requires it.
- If config write is blocked or denied, provide the TOML block and mark setup incomplete.
- Do not write the user's DB URL/ID into plugin files, specs, README, or examples.
- After script writes, run `python3 scripts/setup_config.py --check` when feasible.

Completion evidence:

- config path checked
- script write/check result or manual fallback reason
- `notion.db_url` or `notion.db_id` present
- `general.timezone` present
- `general.language` present
- `general.locale` present
- `general.register` present
- `notion.schema_status` present

## `setup schema`

Goal: make the configured Notion DB capable of storing agent work memory records.

Required properties:

| Property | Type | Required use |
| --- | --- | --- |
| `Name` | title | `PREFIX: 짧은 의미 제목` |
| `분류` | select | record class |
| `상태` | status | lifecycle |
| `요약` | text | one-sentence summary |
| `태그` | multi-select | search tags |
| `기록일` | date | KST datetime source of truth |
| `출처` | select | source class |
| `관련 기록` | relation | same-DB memory links |
| `참조된 기록` | relation | reciprocal backlink |
| `후속 작업` | checkbox | whether follow-up remains |
| `관련 링크` | url | external URLs |

Required option values:

| Property | Options |
| --- | --- |
| `분류` | `작업 히스토리`, `결정`, `요청`, `후속 작업`, `운영 규칙`, `메모` |
| `상태` | `시작 전`, `진행 중`, `완료` |
| `출처` | `사용자`, `Codex`, `연결 도구`, `파일` |

Rules:

- Use `ntn` first for DB/data source fetches and schema-sensitive checks.
- Fetch the DB or data source before making schema-sensitive claims.
- Add missing properties only. Do not delete, rename, or narrow existing properties without explicit approval.
- If an existing property has an incompatible type, report it before changing anything.
- Read `references/ntn-cli.md` before running Notion I/O through `ntn`.
- Do not route normal setup work to a separate `ntn`/`notion-cli` skill or Notion plugin connector.
- If `ntn` schema checks or updates fail, report the command and blocker, then ask before using browser UI or manual steps.
- If relation setup is not possible through tools, record the blocker and keep relation status incomplete.
- Update `notion.schema_status` only after verification evidence exists.

Completion evidence:

- DB fetch result or explicit fetch failure
- required properties present, added, or blocked
- `notion.schema_status` updated if config write is available

## `setup workspace-rule`

Goal: make the current workspace route future Notion memory requests to this skill.

Rules:

- Do not edit repository or global instruction files unless the user asked for workspace-rule setup or approved it during `setup`.
- Prefer the narrowest local rule that applies to the target workspace.
- Keep the rule short:

```md
Notion memory setup or recording requests should use `$advance-codex:notion-memory`.
Store workspace-specific DB values in `~/.agents/configs/notion-memory.toml`, not in repository files.
```

- If the workspace has an existing AGENTS-style instruction file, inspect it before editing.
- If editing is outside the sandbox or requires approval, request approval.
- If approval is not available, provide the proposed rule and mark this command incomplete.

Completion evidence:

- target instruction file identified
- edit applied or proposed rule reported
- scope excludes unrelated repository policy changes

## `setup verify`

Goal: prove the configured Notion memory setup is usable.

Checks:

- config file exists and contains `notion.db_url` or `notion.db_id`
- DB can be fetched through `ntn`, or access blocker is explicit
- `ntn` command syntax was confirmed with local help/docs/spec output
- required schema exists or missing fields are listed
- `general.timezone` is `Asia/Seoul` unless the user chose another value
- `general.language`, `general.locale`, and `general.register` are present or defaulted
- optional write test is performed only with explicit user approval

Verification reporting:

- `WORKED`: config, DB fetch, schema check, and approved write test if requested all pass
- `PARTIAL`: config exists but DB/schema/write verification is blocked or skipped
- `BLOCKED`: missing DB URL/ID, missing Notion access, or denied required approval

Do not claim setup is complete when only manual instructions were provided.
