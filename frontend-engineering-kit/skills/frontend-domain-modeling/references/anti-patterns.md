# Anti-Patterns

Use this guide to catch domain-modeling moves that create cleaner folders without creating clearer ownership.

## DTO Worship

Signal:

- backend response shape is treated as the frontend domain model without checking whether it matches product meaning

Safer move:

- map transport shape into frontend language where the raw contract obscures domain meaning

## Presenter As Domain

Signal:

- view models, label mappers, or display selectors are promoted into domain because they are large or reused

Safer move:

- keep presentation shaping near presentation or application unless it carries real invariants

## Page-As-Feature

Signal:

- one route folder is treated as one domain feature even though it contains several unrelated capabilities

Safer move:

- split by business capability, not by route convenience

## Fake DDD Ceremony

Signal:

- classes, entities, and use-case objects are introduced for thin CRUD data with little invariant-bearing logic

Safer move:

- keep plain data and explicit functions until meaning or policy pressure justifies more structure

## Shared Folder Laundering

Signal:

- business rules move into `shared` or generic helpers because several screens need them

Safer move:

- keep the rule near the smallest stable business owner or create a small shared kernel with explicit domain language

## Orchestration Pushed Into Domain

Signal:

- retry logic, navigation, modal sequencing, or mutation choreography is framed as domain logic

Safer move:

- keep execution flow in application and let it call domain rules explicitly

## API Grouping Mistaken For Feature Grouping

Signal:

- endpoints or resource names define the frontend slice even when the user-facing capability cuts across them

Safer move:

- slice around capability first, then hide API grouping behind adapters or application code

## Final Check

If the proposed modeling mostly improves naming neatness but does not make business ownership easier to explain, it is probably the wrong move.
