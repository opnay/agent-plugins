# Level Gates

Use for unsaved agent wording. Apply levels as `` `light` > `standard` > `extreme` ``; each later level inherits earlier rules and adds only overrides.

## Light

Use for broad "shorter" or "concise" requests.

- Shorten greetings, restatement, weak guesses, repeated caveats, and low-value process notes.
- Preserve direct answer, requested format, verification state, failed or skipped work, approval boundary, exact paths, commands, identifiers, dates, versions, numbers, active language, and register.

## Standard

Use for summaries and "key points only" requests.

- Keep result, key evidence, next action or blocker, important risk, and uncertainty.
- Group related list items under a parent when it removes repeated labels or prefixes.
- Use `` `light` > `standard` > `extreme` `` for level or flow order.

## Extreme

Use only for explicit "status only", "labels only", "one line", "maximum compression", or already structured status/checklist/table content.

- Prefer `label: value.`, compact status tables, or fixed-field lists.
- For missing field values, table cells, or list attribute values, use `-`; do not use `-` inside a natural-language sentence.
- Inherit standard list grouping, but use a field or one-line form when shorter.
- Keep failures, skipped checks, blockers, insufficient verification, approval boundaries, security/privacy/legal/medical/financial exceptions, residual risk, and next action visible.
- Step down only the affected detail when compression would hide a reason, approval state, verification limit, or risk.
