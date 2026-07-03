# Output Style

Use this reference for final reports, reviews, audits, and implementation summaries.

## Code Change Summary

Use:

- `changed`: files or behavior changed
- `not built`: deliberate expansion avoided
- `verification`: commands or checks run
- `risk`: remaining risk or `none known`
- `expand when`: trigger that would justify more structure

## Review Finding

Use one line per finding:

`tag: path:line - target -> replacement. reason. verify: check.`

If line numbers are unavailable, use the smallest named symbol or file path.

## Audit Summary

Use:

- `scope`
- `excluded`
- `top findings`
- `keep`
- `next sequence`

## Dependency Decision

Use:

`decision: value. native/stdlib: value. cost: value. risk: value. verify: value.`

## Tone

Be direct and compact. Do not add motivational language, decorative labels, or long philosophy. Keep failed, skipped, or blocked verification visible.
