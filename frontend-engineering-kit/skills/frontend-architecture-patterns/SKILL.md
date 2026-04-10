---
name: frontend-architecture-patterns
description: Frontend architecture-pattern guidance for choosing and reviewing Atomic Design, FSD, or a coherent custom hybrid. Use when a task is about codebase-structure fit, migration direction, exact Atomic/FSD interpretation, or deciding whether the frontend needs UI taxonomy, feature slicing, or both.
---

# Frontend Architecture Patterns

## Overview

Use this skill when the main question is which frontend codebase-structure pattern fits the repository and how to review that choice rigorously.
Its job is to explain Atomic Design and FSD accurately, decide when each pattern fits, and review whether the current repository is using the chosen structure coherently.
Do not use this skill as the first stop for React rendering semantics, one messy component split, or one hook extraction decision inside an already chosen structure.

## Pattern Definitions

- `Atomic Design`
  - A UI composition methodology that organizes components into five levels: atoms, molecules, organisms, templates, and pages.
  - It is strongest as a design-system or UI-composition model for reusable interface parts and visual consistency.
  - It is not, by itself, a complete frontend application architecture for domain ownership, feature boundaries, API placement, or routing structure.
- `Feature-Sliced Design (FSD)`
  - A frontend application architecture that organizes code by layers and slices, typically `app`, `pages`, `widgets`, `features`, `entities`, and `shared`.
  - It is designed to structure product code by business capability, ownership, dependency direction, and public API boundaries.
  - It usually covers where routing, API interaction, domain logic, and feature-facing UI should live in the app.
  - It is not a deployment or infrastructure architecture.
- `Custom or Hybrid`
  - A valid option when the repository already has a coherent structure that should not be forced into Atomic or FSD labels.
  - In many real projects, Atomic Design works well inside shared UI or design-system space while FSD shapes the product application around it.

## Workflow

1. Identify the decision shape:
   - greenfield architecture choice
   - review of an existing architecture
   - migration from one pattern to another
2. Read the current repository shape:
   - routes and page structure
   - feature boundaries
   - shared UI and design-system scope
   - API and model placement
   - team or ownership boundaries
3. Decide the dominant architectural problem:
   - reusable UI taxonomy
   - product feature scaling
   - domain ownership clarity
   - mixed growth with both design-system and app-structure pressure
4. Compare Atomic Design, FSD, and custom or hybrid options against the actual repository pressures.
5. Choose the pattern or combination that fits the current codebase and growth path with the smallest structural distortion.
6. Use rendering and update behavior only as a secondary signal when it reveals that the chosen structure is grouping unrelated responsibilities too broadly.
7. End with an implementation-ready architecture recommendation, review verdict, or migration direction.

## Selection Criteria

Choose `Atomic Design` when:

- the main challenge is building or cleaning up a reusable UI system
- the repository needs a clearer component taxonomy more than a clearer feature taxonomy
- visual consistency, composition rules, and design-system reuse are the main pain points
- business logic and domain ownership are relatively light compared with UI reuse pressure

Choose `FSD` when:

- the main challenge is structuring a growing product application rather than only a UI library
- the repository needs clear boundaries for routes, features, entities, API calls, and business logic
- several features or teams need clearer ownership and dependency direction
- cross-feature coupling and unclear slice boundaries are causing maintenance problems

Choose `Custom or Hybrid` when:

- the repository already has a coherent structure that should be evolved rather than renamed
- one pattern fits the shared UI space while another fits the product app space
- forcing all code into Atomic or all code into FSD would introduce artificial layers or migration churn
- the team needs explicit rules, but those rules are better expressed as a constrained hybrid than as a pure pattern

## Review Criteria

Review `Atomic Design` with these checks:

- are the five levels used consistently without ad hoc extra levels
- are atoms and molecules kept reusable and context-light
- are organisms, templates, and pages handling composition rather than hidden business ownership
- is the naming system component-oriented rather than feature-locked
- is product logic leaking downward into reusable UI layers

Review `FSD` with these checks:

- are the standard layers and dependency direction respected
- are slices treated as ownership units rather than folders of convenience
- are public APIs used instead of deep imports into slice internals
- is business logic staying out of `shared`
- are routing, API, and model concerns placed in layers that match FSD intent instead of drifting randomly

Review `Custom or Hybrid` with these checks:

- can the structure be explained with a small set of stable rules
- are the mixed rules intentional rather than accidental leftovers
- does each area of the codebase have a clear ownership model
- is the hybrid solving a real repository need instead of avoiding a hard decision

## Decision Rules

- Do not present Atomic Design and FSD as identical problem-solvers; Atomic is primarily a UI-composition model, while FSD is primarily an application-structure model.
- Do not force a product application into pure Atomic Design when the main problem is feature or domain ownership.
- Do not force a design-system-heavy repository into FSD vocabulary when the main problem is reusable UI taxonomy.
- Prefer the repository's current coherent pattern over a cleaner theoretical pattern imported from elsewhere.
- If the app needs both strong shared UI taxonomy and strong product feature boundaries, prefer an explicit hybrid over pretending one pattern fully covers both jobs.
- Treat render churn only as a supporting review lens; use it to notice structure problems, not to turn this skill into React rendering guidance.

## Output Contract

- `Decision shape`
- `Current repository pressures`
- `Pattern definitions applied`
- `Selection criteria result`
- `Review verdict`
- `Recommended pattern`
- `Migration direction` when applicable
- `Residual risk`

## Guardrails

- Do not reduce the choice to trend-following or personal preference.
- Do not pretend Atomic Design is a full replacement for application architecture.
- Do not treat FSD as if it governs deployment, infrastructure, or backend system design.
- Do not create a hybrid that is just "everything goes everywhere."
- Do not use this skill when the real problem is React rendering patterns, a local component split, hook extraction, or one feature-boundary surgery inside an already accepted structure.
