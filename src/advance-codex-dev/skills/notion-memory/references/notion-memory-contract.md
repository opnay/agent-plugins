# Notion Memory Contract

Use this reference for recording or updating Notion memory entries.

## Config

Read `~/.agents/configs/notion-memory.toml` before Notion work.
Required runtime values are `notion.db_url` or `notion.db_id`, plus `general.timezone`.
Use `general.language`, `general.locale`, and `general.register` for memory document writing preferences.
`notion.data_source_id` and `notion.view_id` are optional discovered values.
`notion.schema_status` tracks schema verification state.

Do not store or copy credentials, tokens, cookies, connector auth state, or user-specific DB IDs into plugin files.

## Record Scope

Allowed records:

- work history
- decisions
- follow-up tasks
- verification results
- reusable work knowledge
- explicit user preferences

Excluded records:

- unrelated conversation
- unsupported speculation
- unnecessary personal data
- secrets, credentials, tokens, cookies, auth state
- unrelated Notion workspace cleanup

Reusable work knowledge should be a distilled rule or reusable fact, not a raw failure log.

## Properties

Use these property meanings:

| Property | Type | Use |
| --- | --- | --- |
| `Name` | title | `PREFIX: 짧은 의미 제목` |
| `분류` | select | `작업 히스토리`, `결정`, `요청`, `후속 작업`, `운영 규칙`, or `메모` |
| `상태` | status | `시작 전`, `진행 중`, or `완료` |
| `요약` | text | one concise sentence |
| `태그` | multi-select | narrow searchable tags |
| `기록일` | date | KST datetime |
| `출처` | select | `사용자`, `Codex`, `연결 도구`, or `파일` |
| `관련 기록` | relation | related records in the same memory DB |
| `참조된 기록` | relation | reciprocal backlink, verified by fetch |
| `후속 작업` | checkbox | follow-up remains |
| `관련 링크` | url | external URL only |

Tool adapters may require date or relation fields as expanded values. Preserve the same semantic values when field names differ.

## Title

Use `PREFIX: 짧은 의미 제목`.
Do not put date or time in `Name`.
Use `기록일` as the source of truth for ordering.

Prefix is an expandable namespace. Prefer a short uppercase repo, tool, product, or domain code that helps scanning.

## Date

Use KST datetime for `기록일`.
If multiple records would share the same displayed minute, separate them by at least one minute because Notion may normalize seconds away.

## Relations And Links

- Use `관련 기록` only for same-DB memory record links.
- Treat `참조된 기록` as a reciprocal backlink and verify it by fetch when relations matter.
- Put external URLs in `관련 링크`.
- Do not put external URLs in relation fields.

## Body Template

Use this body unless the user requests a narrower record:

```md
## 요약
한두 문장으로 핵심을 씁니다.

## 범위
- 포함:
- 제외:

## 사실
- 확인된 사실만 씁니다.

## 결정
- 결정된 내용이 없으면 `없음`으로 씁니다.

## 후속 작업
- 없으면 `없음`으로 씁니다.

## 검증
- 도구 응답, fetch 결과, 실행 명령 등 증거를 씁니다.
- 검증하지 못한 것은 `미검증`으로 씁니다.

## 출처
- 사용자 발화, 파일 경로, Notion 페이지 URL 등 근거를 씁니다.
```

## Writing Rules

- Use `general.language`, `general.locale`, and `general.register` when configured.
- Default to Korean, Seoul context, and consistent polite register.
- Keep Notion entries concise but executable for a future agent.
- Separate facts from decisions, follow-ups, and verification.
- Do not add decorative commentary or unsupported interpretation.
- Do not send optional progress commentary when a direct final record is enough.

## Write Workflow

1. Map user, Notion, local, and excluded scope.
2. Apply the allow-list first.
3. Apply the block-list.
4. Read config and fetch the DB if tools are available.
5. Normalize properties and body.
6. Create or update the page.
7. Fetch the page when property, relation, or body verification matters.
8. Report status, evidence, skipped checks, and residual risk.

## Status Words

- `WORKED`: write or update completed and required verification passed.
- `PARTIAL`: record was drafted or manual steps were provided, but write or verification was skipped or blocked.
- `BLOCKED`: required config, Notion access, user input, or approval is missing.
