# Record Formats

## session_record.md

Schema: `session-record`

Numbered sections:

1. Session UUID
2. Timestamp
3. Goal
4. Scope
5. Constraints
6. Working assumptions
7. Runbook
8. Verification evidence
9. Mistake notes
10. Residual risks

## change_record.md

Schema: `change-record`

Markdown sections:

- `<!-- session-manager:purpose -->`
- `# Purpose`
- `<!-- session-manager:key_changes -->`
- `# Key Changes`
- `<!-- session-manager:verification -->`
- `# Verification`
- `<!-- session-manager:review_gates -->`
- `# Review Gates`
- `<!-- session-manager:reviewer_notes -->`
- `# Reviewer Notes`
- `<!-- session-manager:residual_risks -->`
- `# Residual Risks`

The hidden marker comment before each visible heading is part of the format contract and is what the parser uses to delimit sections safely.

## CLI Notes

- `validate` reports `valid` and `detected_schema`, and exits nonzero for malformed records
- `list` keeps the 3-column output by default; use `--include-schema` to add schema data
- `--value-file -` reads UTF-8 content from stdin
