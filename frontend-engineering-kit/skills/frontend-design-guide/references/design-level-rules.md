# Design Level Rules

## Level Detection

Ask these questions before judging hierarchy, spacing, or polish:

1. Is the main issue about a shared component, reusable pattern, token usage, or variant consistency?
2. Would the same rule likely apply across multiple screens or features?
3. Is the problem tied to one concrete user task, one section, or one feature flow?
4. Would generic cleanup erase important product meaning or task emphasis?

If the problem is mainly reusable and consistency-driven, treat it as `design-system level`.
If the problem is mainly task-driven and screen-specific, treat it as `product level`.

## Design-System-Level Rules

- Focus on token correctness before one-off polish.
- Favor constrained variants over ad hoc visual branches.
- Check spacing and typography against the shared scale, not screen-local preference.
- Treat baseline accessibility and interaction consistency as part of the shared contract.

## Product-Level Rules

- Start from the primary user task and CTA priority.
- Favor hierarchy, grouping, and state clarity over abstract consistency.
- Let the screen express feature meaning clearly instead of over-generalizing UI.
- Judge responsive behavior by preserving task priority, not only by preventing overflow.

## Cross-Level Smells

- Token cleanup used to avoid solving weak CTA emphasis.
- Shared-pattern consistency used to flatten important feature meaning.
- Product-specific urgency or state differences forced into a generic visual treatment.
- One-screen visual tweaks added where the real fix should be a shared token or component rule.
