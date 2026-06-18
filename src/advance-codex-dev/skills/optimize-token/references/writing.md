# Writing Axis

Use when agent token optimization touches content that will be saved in files, docs, specs, records, release notes, README files, or handoff artifacts.
For unsaved situations, apply the active level directly without a separate axis.
Stored writing is not covered by the default global level; manage it through this reference.

## Preserve

- Current state only; do not leave discarded decisions or temporary conversation context in the artifact.
- Durable contract, user intent, approval boundary, verification status, residual risk, and exact required format.
- Meaning-bearing context that a future reader needs to execute, verify, approve, or judge the artifact correctly.
- Required order, dependency, applicability condition, excluded scope, rollback condition, source of truth, evidence source, confirmed scope, and unconfirmed scope.
- Paths, commands, identifiers, dates, versions, numbers, public API names, active language, and register.
- Enough context for a future reader to execute the saved instruction without the conversation.

## Shorten

- Authoring process notes, comparisons with abandoned designs, duplicated contract text, and background that does not affect artifact use.

## Level Gates

- `light`: concise natural writing for durable docs.
- `standard`: keep contract, evidence, next action, risk, and uncertainty only.
- `extreme`: labels, compact tables, or fixed fields only when the artifact remains executable.

Step down when compression would hide meaning-bearing context, verification limits, approval boundaries, contract scope, or current-state accuracy.
