# Level Gates

Use for unsaved agent wording. Apply levels as `` `light` > `standard` > `extreme` ``; each later level inherits earlier rules and adds only overrides.

## Light

Use for broad "shorter" or "concise" requests.

- Shorten greetings, restatement, weak guesses, repeated caveats, and low-value process notes.
- Preserve direct answer, requested format, meaning-bearing context, verification state, failed or skipped work, approval boundary, exact paths, commands, identifiers, dates, versions, numbers, active language, and register.
- Meaning-bearing context is any cause, evidence, boundary, sequence, dependency, source, uncertainty, scope, non-goal, or next-action condition that would change the reader's judgment if removed.

## Standard

Use for summaries and "key points only" requests.

- Keep result, key evidence, next action or blocker, important risk, and uncertainty.
- Keep workflow order, dependencies, cause/evidence separation, comparison basis, excluded scope, confirmed/unconfirmed scope, source of truth, and evidence source hierarchy when they affect judgment.
- Group related list items under a parent only when it removes repeated labels without blurring the meaning category.
- Use `` `light` > `standard` > `extreme` `` for level or flow order.

## Extreme

Use only for explicit "status only", "labels only", "one line", "maximum compression", or already structured status/checklist/table content.

- Prefer `label: value.`, compact status tables, or fixed-field lists.
- For missing field values, table cells, or list attribute values, use `-`; do not use `-` inside a natural-language sentence.
- Inherit standard list grouping, but use a field or one-line form when shorter and meaning-bearing context still fits the fields.
- Split compact output into fields such as `cause`, `evidence`, `sequence`, `scope`, `unconfirmed`, `source`, `risk`, or `next` when one field would hide a required distinction.
- Keep failures, skipped checks, blockers, insufficient verification, approval boundaries, security/privacy/legal/medical/financial exceptions, residual risk, and next action visible.
- Step down only the affected detail when compression would hide meaning-bearing context, a reason, approval state, verification limit, or risk.

Examples:

- `Cause is not confirmed, but logs show a permission error.` -> `cause: unconfirmed. evidence: permission error log.`
- `Change A; B is out of scope.` -> `changed: A. excluded: B.`
- `Edit the spec, rewrite runtime, then build.` -> `sequence: spec > runtime > build.`
- `Source is the dev source; release surface is generated.` -> `source: dev source. generated: release surface.`
