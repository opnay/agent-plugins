# Turn-Gate Multi-Flow Plan Template

This file is the date-scoped multi-flow plan for turn-gated work.
Use it to show how the user's request decomposes into planned flows, why each flow exists, and when the work should move to the next flow.
Keep it incremental. Do not delete completed work; summarize completed flows and keep their links.

## Date

- YYYY-MM-DD

## Daily Work History

- Summary:
- Latest user request:
- Latest decision:

## User Requests Today

1. <user request summary and time/order>

## Flow Index

1. `.agents/sessions/YYYYMMDD/001-english-lower-slug.md`
   - User request:
   - Flow purpose:
   - Status: planned | active | complete | blocked
   - Completion criteria:
   - Next-flow trigger:
   - Summary:

## Planned Flow Sequence

- Preparation source: user-message | existing-flow | correction | next-flow
- Preparation result:
- Flow-list basis:

1. `<flow slug or title>`
   - Purpose:
   - Why this flow:
   - Owns:
   - Core phase coverage: preparation | work | verification | reporting
   - Completion criteria:
   - Next-flow trigger:
2. `<flow slug or title>`
   - Purpose:
   - Why this flow:
   - Owns:
   - Core phase coverage: preparation | work | verification | reporting
   - Completion criteria:
   - Next-flow trigger:
3. `<flow slug or title>`
   - Purpose:
   - Why this flow:
   - Owns:
   - Core phase coverage: preparation | work | verification | reporting
   - Completion criteria:
   - Next-flow trigger:

## Current Status

- Active flow:
- Current core phase: preparation | work | verification | reporting
- Latest decision:
- Required next action:
- Work boundary:
- Verification expectation:
- Next planned flow:

## Completed Flow Summaries

- `<flow filename>`: <short retained summary>

## Explicit Turn-End Option

- Record that the user can explicitly stop the turn even when the visible question UI cannot show a stop option.

## Open Risks

- <risk or uncertainty>
