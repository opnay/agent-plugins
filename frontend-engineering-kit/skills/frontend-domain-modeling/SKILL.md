---
name: frontend-domain-modeling
description: Decide when frontend code should stay thin and when it needs explicit domain modeling. Use when a task is about business rules, feature boundaries, ubiquitous language, view models versus domain logic, or whether to keep server contracts close versus introducing policy, value, or entity-like models.
---

# Frontend Domain Modeling

## Overview

Use this skill when the real question is not "how should this component be structured?" but "how much modeling does this frontend actually need?"

This skill treats frontend domain modeling as a threshold decision, not a default architecture style.
Its job is to decide:

- when raw server contracts are good enough
- when a thin mapper or view model is enough
- when a policy or domain function is needed
- when a value-like or entity-like model is justified
- where feature boundaries should follow business capability rather than route or component layout

The default stance is `thin-by-default`.
Do not model because the data exists.
Model because meaning, invariants, or repeated business decisions need protection.

## What This Skill Is For

Use this skill when the main problem is one or more of these:

- business rules are leaking into JSX, handlers, selectors, or generic helpers
- the team is debating whether to wrap backend responses in frontend models
- view-model logic and domain logic are getting mixed together
- feature or slice boundaries are unclear because the code follows pages or endpoints instead of product capability
- the code needs shared product language rather than more React refactors

## What This Skill Is Not For

Do not use this skill when the real question is:

- Atomic Design versus FSD versus hybrid structure
- hook API shape, provider scope, rerender spread, or effect discipline
- one local component split after the business meaning is already clear
- visual hierarchy, spacing, accessibility, or design-system quality

Those belong to sibling skills.
This skill should stop once the modeling threshold and boundary decision are clear, then hand implementation structure to `react-architecture` or `component-architecture`.

## Default Strategy

Start with the thinnest thing that preserves meaning.

The usual escalation order is:

1. `Typed server contract`
   - use the server shape close to the UI when it already matches product meaning and no important invariant is being protected
2. `Thin mapper or view model`
   - add one local mapping layer when the UI needs flattening, labels, grouping, or screen-specific shaping
3. `Policy or domain function`
   - extract explicit business decisions when the same rule repeats across screens, flows, or handlers
4. `Value-like model`
   - introduce a small explicit model when validation, normalization, comparison, or domain-safe operations need one owner
5. `Entity-like model`
   - use only when identity, lifecycle, or clustered rule-bearing behavior matters to frontend reasoning

Do not skip directly to entity-like modeling because the code feels important.
Most frontend code should stop at step 1, 2, or 3.

## Core Rules

- `Server contract is not the enemy`
  - Using server responses close to the UI is often the right default.
  - A backend DTO only becomes a problem when transport shape starts distorting product meaning or scattering business decisions.
- `View shaping is not domain modeling`
  - Labels, grouping, formatting, badge variants, and empty-state flags are often presentation work.
  - Large presentation logic is still presentation logic unless it protects a real invariant.
- `Policy before entity`
  - Frontend business logic more often needs explicit decisions than rich objects.
  - If the real need is `canEdit`, `canSubmit`, or `nextAllowedStatus`, start with policy functions.
- `Model to protect meaning`
  - If a rule should survive screen rewrites, route moves, and layout changes, it deserves a stable owner.
  - If a wrapper only mirrors shape, it is probably ceremony.
- `Feature boundaries follow capability`
  - A feature is a user-facing capability or coherent product responsibility.
  - Routes, modals, and page sections may compose several features and should not automatically define them.

## Decision Framework

Ask these questions in order.

### 1. Can raw data stay raw?

Keep the server contract close when:

- the data is mostly read-only
- the UI only needs direct display plus light formatting
- the response shape already matches product language
- no repeated business rule depends on interpreting the fields

Do not keep it raw when:

- the UI repeats the same field interpretation in several places
- transport quirks are leaking into product logic
- raw enums or nullable fields are driving business decisions all over the app
- the backend shape does not match how the product talks about the concept

### 2. Is this just view shaping?

Use a thin mapper or view model when the main need is:

- flattening nested response data
- grouping or sorting for one screen
- producing labels, display states, or badge variants
- screen-specific convenience flags

Do not call this domain modeling unless the mapping itself protects a business invariant.

### 3. Is there a repeated business decision?

Introduce a policy or domain function when:

- the same permission, eligibility, pricing, validation, or transition logic repeats
- the same rule appears in JSX, event handlers, selectors, and submit flows
- the same business question is answered in more than one place

Examples:

- `canEditOrder`
- `canPublishDraft`
- `isUpgradeEligible`
- `getNextAllowedStatuses`

This is the most common place where explicit frontend domain logic should begin.

### 4. Does a small explicit model buy real safety?

Introduce a value-like model when:

- normalization matters
- invalid states are easy to create accidentally
- comparisons or calculations need one safe representation
- repeated conversions are hiding the real concept

Common fits:

- money
- date range
- filter criteria
- permission set
- validated form input

### 5. Does identity or lifecycle matter?

Introduce an entity-like model only when:

- the concept has stable identity in frontend reasoning
- lifecycle transitions matter
- several behaviors cluster around one durable concept
- several features touch the same invariant-bearing object

Common fits:

- draft with save, submit, publish, archive transitions
- order with status transition rules
- account with eligibility and entitlement rules

If identity does not matter, prefer policy or value modeling instead.

## Feature Boundary Strategy

Feature boundaries should answer this question:

`What capability does this part of the product own?`

Good feature boundaries:

- can be named in product language
- change for related business reasons
- expose a small understandable public surface
- keep shared meaning local until it truly repeats elsewhere

Bad feature boundaries:

- route-shaped folders pretending to be capabilities
- endpoint-shaped folders pretending to be product slices
- generic shared folders holding half-domain logic from many features
- one feature absorbing several capabilities because they appear on one page

Remember:

- a page can compose several features
- a modal can belong to a feature without defining one
- one feature can span several routes if the same capability persists

## Layer Split

Use these working distinctions:

- `Transport or adapter`
  - raw API contracts, endpoint details, protocol quirks, DTO mapping
- `Presentation`
  - rendering, view models, labels, grouping, display-only states
- `Application`
  - command flow, orchestration, retries, submit sequencing, navigation coordination
- `Domain`
  - business meaning, policies, invariants, validation rules, lifecycle rules

The main mistake to avoid is pretending large presentation or orchestration code is domain logic just because it became important.

## Workflow

1. Name the user-facing capability or workflow.
2. Write down the product terms that should stay stable.
3. Mark each important piece of logic as one of:
   - transport
   - presentation
   - application orchestration
   - domain rule
4. Decide the lowest layer that can safely own each rule.
5. Start from `raw data` and climb the escalation ladder only when the lower step fails.
6. Stop as soon as the meaning is protected and the ownership is explainable.
7. Hand React-specific implementation structure to the appropriate sibling skill.

## Output Contract

- `Capability`
- `Product language`
- `Keep raw vs map vs model decision`
- `Policy candidates`
- `Value-like model candidates`
- `Entity-like model candidates`
- `Feature boundary decision`
- `Layer split`
- `React handoff`
- `Overmodeling risk`

## Review Questions

- Is this wrapper protecting meaning, or only renaming fields?
- Is this rule business logic, or only screen logic?
- Does this concept need identity, or only validation and normalization?
- Would this logic survive a screen rewrite?
- Is this feature boundary capability-shaped or route-shaped?
- Are we introducing a model because the product needs it, or because the architecture looks nicer with one?

## Guardrails

- Do not treat all backend responses as domain entities.
- Do not force DDD vocabulary onto thin CRUD or read-only surfaces.
- Do not promote view models into domain models because they became large.
- Do not hide business rules in shared UI, generic utilities, or selectors.
- Do not use `domain` as a prestige folder for whatever feels important.
- Do not keep climbing the modeling ladder once a lower step already protects meaning well enough.

## References

- Read [references/modeling-thresholds.md](references/modeling-thresholds.md) when deciding how far up the modeling ladder to climb.
- Read [references/layering-rules.md](references/layering-rules.md) when domain, application, presentation, and transport responsibilities are getting mixed together.
- Read [references/feature-boundary-tests.md](references/feature-boundary-tests.md) when deciding whether a slice is a real capability or only a route-shaped grouping.
- Read [references/anti-patterns.md](references/anti-patterns.md) when a proposed cleanup looks neat but may just be ceremony or domain confusion in disguise.
