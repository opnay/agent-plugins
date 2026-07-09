# ntn CLI First Path

Use this reference before Notion I/O for `notion-memory` setup, recording, or verification.
`ntn` is the first-choice path for this skill.

## Priority

1. Use `ntn` directly for Notion memory DB fetches, page writes, and read-back verification.
2. Do not route normal memory work to a separate `ntn` or `notion-cli` skill.
3. Do not use a Notion plugin connector as the default DB fetch or page write path.
4. Use local CLI help before relying on examples:
   - `ntn --help`
   - `ntn <command> --help`
   - `ntn api ls`
   - `ntn api <path> --help`
   - `ntn api <path> --docs`
   - `ntn api <path> --spec`
5. If `ntn` is unavailable or blocked, report the blocker and ask before using browser UI or manual fallback.

## Authentication

- Check whether `NOTION_API_TOKEN` is set without printing the value:

```bash
if printenv NOTION_API_TOKEN >/dev/null; then echo NOTION_API_TOKEN=set; else echo NOTION_API_TOKEN=unset; fi
```

- If `NOTION_API_TOKEN` is set, prefer it.
- If no token is available, `ntn login` may require the user to open a browser URL.
- Never store credentials, tokens, cookies, keychain state, or connector auth state in `~/.agents/configs/notion-memory.toml` or plugin files.

## `ntn api`

Use `ntn api` for the public Notion API surface when higher-level commands do not cover the task.
The CLI is self-documenting; inspect the endpoint before calling it.

Common syntax:

```bash
ntn api v1/users page_size==100
ntn api v1/pages parent[page_id]=abc123
ntn api v1/pages -d '{"parent":{"page_id":"abc123"}}'
ntn api v1/pages -d @body.json
ntn api v1/pages -X POST --spec
```

Input rules:

- `name==value`: query parameter.
- `path=value`: body string value.
- `path:=json`: body typed JSON value for numbers, booleans, arrays, objects, and null.
- `-d <JSON|@PATH|@->`: JSON body source.
- Method defaults to `GET`; body input usually infers `POST`; `-X` wins.
- Use exactly one JSON body source: stdin, `--data`, or inline body inputs.

## Pages

Prefer `ntn pages` for Markdown page content:

```bash
ntn pages get <page-id>
ntn pages get <page-id> --json
ntn pages create --parent data-source:<data-source-id> --content '# Title\n\nBody'
ntn pages edit <page-id> --content '# Updated body'
```

Rules:

- Use `--json` when machine parsing matters.
- Use `ntn api v1/pages` when fixed database properties, templates, or full Pages API details are needed.
- Keep `notion-memory` properties and body-format contract from `references/notion-memory-contract.md`.

## Data Sources

Use data source IDs for query commands:

```bash
ntn datasources resolve <database-id>
ntn datasources query <data-source-id> --limit 50 --json
ntn datasources query <data-source-id> --filter '{"property":"Done","checkbox":{"equals":true}}'
```

Rules:

- A Notion database can have multiple data sources.
- Resolve a database ID before using query commands when only a database ID is known.
- Fetch or inspect schema before schema-sensitive claims.

## Files

Use `ntn files` for file-upload API tasks:

```bash
ntn files create < file.png
ntn files create --filename photo.png --content-type image/png < /tmp/blob
ntn files create --external-url https://example.com/photo.png
ntn files get <upload-id>
ntn files list
```

Use file uploads only when the memory task genuinely needs a Notion file object.

## Workers Out Of Scope

Worker commands are outside `notion-memory` setup, recording, and verification.
If the user asks for worker work, stop using this reference for that task and ask before proceeding because that work is outside this skill's memory scope.

## Safety

- Do not run destructive commands such as delete, trash, move, schema rename, or schema narrowing without explicit user approval and a clear target.
- Do not call manual instructions a completed setup.
- Do not print secrets, pass tokens in command text, or save auth state in repo files.
- Record skipped checks, failed CLI calls, and residual risk in the final memory note.
