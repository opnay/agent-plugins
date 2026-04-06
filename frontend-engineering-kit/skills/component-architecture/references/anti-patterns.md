# Anti-Patterns

Use this guide to catch architectural moves that look clean but weaken ownership.

## Early Abstraction

Signal:
- shared extraction before repeated responsibility is proven

Safer move:
- keep the logic near the feature or product boundary until real reuse appears

## Mechanical Presentational/Container Split

Signal:
- splitting by pattern habit instead of actual ownership conflict

Safer move:
- split only when rendering and orchestration need different owners

## Complexity Laundering Into Hooks

Signal:
- moving complexity into a hook without making the ownership story clearer

Safer move:
- compare hook extraction honestly against utility, feature, domain, api, and component-local placement

## Shared Extraction That Erases Domain Meaning

Signal:
- product-level components losing semantic naming to look reusable

Safer move:
- delay promotion until repeated responsibility and repeated change reason are both visible

## File Split Without Responsibility Split

Signal:
- more files, same mixed responsibilities

Safer move:
- split only when owners and reasons to change become easier to explain

## Generic Naming That Hides Ownership

Signal:
- generic names replacing real product meaning

Safer move:
- keep domain language where product-level ownership still matters

## Variant Inflation

Signal:
- one component absorbing too many behavior or workflow modes through variants or booleans

Safer move:
- keep variants constrained and split when the meaning becomes feature-specific

## Final Check

If the refactor mainly improves visual neatness but not ownership clarity, it is probably the wrong move.
