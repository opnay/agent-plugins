---
name: component-architecture
description: Design or review frontend component boundaries. Use when a task involves splitting large components, choosing presentational vs container responsibilities, extracting hooks, shaping component APIs, colocating state, or deciding composition, ownership, and folder structure for React or UI code.
---

# Component Architecture

## Overview

Use this skill when a frontend task is primarily about component shape and responsibility.
The goal is to decide where rendering, state, effects, domain calls, and reusable abstractions should live before implementation sprawls.
Do not use file size as the main decision signal. Use responsibility conflicts and change pressure instead.

## Workflow

1. Identify the user task or screen responsibility the component is trying to serve.
2. Map the current responsibilities and data flow.
3. Classify the work inside the component by concern:
   - rendering concern
   - orchestration concern
   - domain concern
   - async concern
   - styling concern
4. Decide whether those concerns are naturally coupled or in conflict.
5. Decide who should own the main responsibilities inside the current boundary.
6. Recommend the smallest boundary change that reduces conflict without creating abstraction theater.
7. Define public APIs, ownership boundaries, and state placement.
8. End with an implementation-ready boundary recommendation, including when not to split.

## Decision Framework

Evaluate the component through these questions:

### 1. Task Coherence

- Does the component still serve one understandable user task?
- Would splitting it make the task easier to reason about, or just spread it across files?

### 2. Concern Coupling

- Are rendering, orchestration, domain logic, async logic, and styling working together naturally?
- Or are they changing for different reasons and creating friction inside one component?

### 3. Change Pressure

- Which parts of the component are likely to change independently?
- If one area changes often without the others, that is a boundary signal.

### 4. Ownership Clarity

- Can you explain who owns rendering, state, effects, and decision-making in a few sentences?
- If ownership is hard to explain, the boundary is probably weak.

### 5. Split Cost

- Would a split reduce complexity, or only move it into props, wrapper components, or hooks?
- Prefer no split when the new boundary adds indirection without reducing reasoning cost.

## Concern Definitions

Use these definitions before classifying mixed code.

### Rendering Concern

Definition:
- decides what the UI shows and how visible structure is composed

Includes:
- JSX or template shape
- conditional rendering
- section ordering
- fallback display choice
- display-only formatting that exists for presentation

Does not include:
- business policy decisions
- async lifecycle ownership
- multi-part flow coordination

Key question:
- If the product rule stayed the same but the screen layout changed, would this code change?

### Orchestration Concern

Definition:
- coordinates multiple UI parts, state transitions, or actions into one user flow

Includes:
- event routing across child components
- tab, panel, modal, or step coordination
- sequencing user actions
- glue code that connects several parts of the screen

Does not include:
- core business rules themselves
- raw request lifecycle by itself
- purely local rendering choices

Key question:
- Is this code mainly coordinating moving pieces rather than defining one rule or one view?

### Domain Concern

Definition:
- encodes product meaning, business rules, invariants, and policy decisions

Includes:
- eligibility rules
- pricing or permission rules
- domain calculations
- invariant checks
- domain-specific transformations that should survive UI changes

Does not include:
- markup structure
- token or styling choices
- request lifecycle logic

Key question:
- If the UI changed completely but the product behavior stayed the same, would this rule still exist?

### Async Concern

Definition:
- owns the lifecycle of fetching, mutating, synchronizing, retrying, or reconciling asynchronous data

Includes:
- loading, error, and empty handling tied to data operations
- mutation progress and completion handling
- invalidation or refetch triggers
- race, stale, or synchronization handling

Does not include:
- pure domain calculations
- display-only conditionals
- high-level screen coordination unless async work is the dominant responsibility

Key question:
- Is this code about data or side-effect lifecycle rather than UI meaning?

### Styling Concern

Definition:
- controls visual treatment, tokens, emphasis, and appearance states

Includes:
- variant mapping
- token selection
- visual emphasis rules
- selected, disabled, warning, or danger appearance

Does not include:
- business rules
- request lifecycle ownership
- non-visual action sequencing

Key question:
- If behavior stayed the same but appearance changed, would this code change?

## Concern Classification Questions

Ask these when code feels ambiguous:

- Would this logic still matter if the UI were redesigned?
  - If yes, it may be domain or async rather than rendering.
- Is this code coordinating several parts of the screen?
  - If yes, it is likely orchestration.
- Is this code managing request or side-effect lifecycle?
  - If yes, it is likely async.
- Is this code deciding meaning or policy rather than display?
  - If yes, it is likely domain.
- Is this code mostly changing emphasis, tokens, or visual state?
  - If yes, it is likely styling.

## Responsibility Ownership Rules

After classifying concerns, decide where each responsibility should live.
Do not treat state as the only ownership problem. State is one subcase of responsibility placement.

### Ownership Types

- rendering ownership:
  - visible structure, local display conditions, local interaction feedback
- orchestration ownership:
  - coordinating child components, flows, transitions, and event sequencing
- domain ownership:
  - product rules, calculations, permissions, invariants, policy decisions
- async ownership:
  - fetch and mutation lifecycle, synchronization, invalidation, loading and error handling
- formatting ownership:
  - display mapping, labels, presentation-specific adapters, view-facing shaping
- state ownership:
  - local UI state, lifted shared state, derived state, form state, server-adjacent state

### Placement Questions

- Which responsibility changes when the UI changes?
  - keep that responsibility closer to rendering or formatting
- Which responsibility survives UI redesign?
  - move that responsibility toward domain or async ownership
- Which responsibility coordinates several children or view regions?
  - keep that responsibility at the smallest boundary that actually owns the flow
- Which responsibility is only computed from other values?
  - prefer deriving it instead of storing it
- Which responsibility is shared by multiple siblings?
  - move it to the smallest common owner, not automatically to a global layer

### Wrong Ownership Signals

- product rules are hidden inside render branches
- request lifecycle is managed deep inside leaf UI with no clear ownership
- parent components own state or handlers they cannot explain
- siblings depend on state that lives too low in the tree
- derived values are stored and synchronized manually
- formatting logic with product meaning is mixed directly into view code
- one component owns orchestration, async lifecycle, and domain decisions without a clear reason

### State As A Subcase

Use these rules for state-specific ownership decisions:

- keep local state local when only one boundary reads and writes it
- lift state only when multiple consumers truly need shared ownership
- derive values instead of storing them when they can be recomputed safely
- isolate form state when validation, touched, dirty, or submit lifecycle semantics appear
- avoid copying async results into local state unless there is a clear ownership break or editing buffer

### Ownership Decision Rule

Recommend a move only when ownership becomes easier to explain and future changes become safer.

Do not move a responsibility when the main result is:

- broader but vaguer ownership
- extra prop tunneling
- global or shared state without a real coordination need
- duplicated responsibility across layers
- a component that delegates everything but still conceptually owns the task

## Boundary Signals

Signals that a split is likely justified:

- domain rules are mixed directly into view rendering
- async orchestration is overwhelming the rendering path
- a parent component coordinates multiple distinct UI regions with different change patterns
- styling variants and behavioral variants are tangled together
- state ownership is hard to explain or leaks across unrelated subtrees

Signals that a split is probably premature:

- the component is large but still serves one coherent task
- extracted pieces would only forward props without gaining real ownership
- a custom hook would hide complexity without reducing it
- the proposed split follows a pattern mechanically instead of solving a concrete conflict
- the current structure matches the project well enough and the real issue is local clarity

## Split Decision Rule

A split should be recommended only when at least one of these is true:

- two or more concerns are changing independently and fighting for ownership
- a boundary would make state, effects, or decisions materially easier to explain
- the new structure would reduce reasoning cost for future changes

A split should not be recommended when the main result is:

- more indirection
- prop tunneling
- wrapper-only components
- hook extraction without clearer ownership

## Output Contract

- `Current responsibility map`
- `Concern classification`
- `Responsibility ownership decision`
- `Boundary problems`
- `Split decision`
- `Recommended split`
- `Public API shape`
- `State and hook placement`
- `Why not split yet`
- `Migration risk`

## Guardrails

- Do not split components just because they are large.
- Do not treat a hook extraction as a successful boundary by default.
- Do not apply container/presentational patterns mechanically.
- Do not widen ownership upward unless shared responsibility is real.
- Do not store values that should be derived.
- Prefer composition over inheritance-style abstraction.
- Prefer existing project patterns unless they are actively causing confusion.
- Prefer boundaries that match the current project structure unless that structure is the source of the problem.
- Keep boundaries understandable to the next maintainer.

## References

- Read [references/boundary-checklist.md](references/boundary-checklist.md) for split criteria.
- Read [references/concern-classification.md](references/concern-classification.md) when rendering, orchestration, domain, and async responsibilities are easy to confuse.
- Read [references/responsibility-ownership-rules.md](references/responsibility-ownership-rules.md) when deciding where rendering, domain, async, formatting, and state responsibilities should live.
