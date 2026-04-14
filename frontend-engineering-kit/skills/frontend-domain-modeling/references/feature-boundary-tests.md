# Feature Boundary Tests

Use this guide when the main question is whether code belongs in one feature, several features, or a smaller shared kernel.

## Boundary Test

Ask these questions before drawing a feature or slice boundary:

1. Can the boundary be named as one user-facing capability?
2. Do the main rules inside it change for the same business reasons?
3. Would one product owner explain this area as one thing or several things?
4. If a screen disappeared, would the capability still make sense?

If those answers are weak, the proposed feature boundary is probably wrong.

## Good Boundary Signals

- one capability or workflow explains the slice
- one public API can expose the main useful surface
- most internal changes happen for related business reasons
- neighboring features can depend on it without learning its internals

## Bad Boundary Signals

- the feature exists only because files were colocated on one route
- one folder contains unrelated policies, tables, forms, and API calls
- the main shared trait is visual similarity rather than capability
- the boundary is mostly transport or component taxonomy with no business meaning

## Shared Kernel Test

Promote logic into a broader shared kernel only when:

- the same product meaning repeats
- the same reason to change applies
- the same terms are used by several features
- centralizing the logic clarifies ownership instead of hiding it

Do not promote logic only because:

- two screens look similar
- the code can technically be imported in many places
- a feature folder feels too large

## Page Versus Feature

- A page can compose several features.
- A modal can belong to a feature without defining one.
- A feature can span several pages if the same capability and rules persist.
- Route boundaries are useful structure, but they are not automatically domain boundaries.

## Entity-Centered Boundaries

An entity-centered feature can make sense when:

- one concept has meaningful lifecycle or permissions
- several flows touch the same invariant set
- the concept remains recognizable across screens

Do not force entity-centered slicing when:

- the app is mostly thin CRUD with little invariant-bearing logic
- the same entity appears in unrelated capabilities with different meanings

## Failure Modes

- Page folder mistaken for feature boundary
- Shared folder becoming a trash can for half-domain logic
- One feature absorbing several capabilities because they share a route
- Technical API grouping mistaken for product feature grouping
