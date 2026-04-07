# Refactor Patterns

Use these patterns after diagnosing the component problem shape.

## Direction Rule

- For page or screen refactoring, work top-down:
  - page or feature intent
  - section or composite boundary
  - lower-level component split only when still needed
- For component contract design, work bottom-up:
  - primitive or lower-level contract
  - composite contract
  - page recomposition

Do not start a page refactor from primitive cleanup.
Do not leave composites as accidental page-only arrangements with no stable lower-level contract.

## God Component Split

Best for:

- components with too many responsibility candidates
- mixed concerns that keep changing for different reasons

Keep:

- one clear owner for the user task

Avoid:

- fragmenting the task until no boundary owns the full behavior

## Page / Section / Widget Split

Best for:

- large screens with several distinct visual and ownership regions

Keep:

- page-level orchestration and cross-region coordination

Avoid:

- slicing only by layout blocks without real ownership differences

## List / Item Split

Best for:

- lists where item behavior and collection behavior are competing

Keep:

- list-level collection semantics

Avoid:

- moving markup out while the list still owns every meaningful item behavior

## View-Model / Adapter Split

Best for:

- render-heavy mapping, formatting, and UI-shaping logic

Keep:

- the owner that consumes the UI-facing shape

Avoid:

- creating a mapper layer with no stable contract

## Effect Isolation

Best for:

- effects that obscure rendering or ownership clarity

Keep:

- the boundary that has real context for effect lifecycle decisions

Avoid:

- extracting effects mechanically into hooks with no stronger ownership story

## Selection Rule

Pick the first pattern that removes the dominant source of confusion.
Do not stack several patterns at once unless there is a clear dependency between them.
