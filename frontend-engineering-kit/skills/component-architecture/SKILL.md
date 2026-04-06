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
3. Extract the responsibility candidates that are actually worth evaluating.
4. Classify the work inside the component by concern:
   - rendering concern
   - orchestration concern
   - domain concern
   - async concern
   - styling concern
5. Decide whether those concerns are naturally coupled or in conflict.
6. Decide who should own the main responsibilities inside the current boundary.
7. Place extracted responsibilities into the right target layer or boundary.
8. Decide whether extraction is needed and whether a hook is actually the right extraction target.
9. Decide whether the component should be treated as library-level or product-level before evaluating its API.
10. Recommend the smallest boundary change that reduces conflict without creating abstraction theater.
11. Define public APIs, ownership boundaries, and state placement.
12. End with an implementation-ready boundary recommendation, including when not to split.

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

## Candidate Extraction Rules

Do not evaluate random code fragments.
First extract responsibility candidates: units that have one meaningful reason to change and can be placed somewhere intentionally.

### What Counts As A Candidate

- a behavior unit:
  - one user-visible action or flow, such as submit, select, dismiss, navigate, confirm
- a decision unit:
  - one eligibility, permission, pricing, validation, or policy rule
- a sync unit:
  - one synchronization responsibility, such as URL sync, request lifecycle sync, effect sync, or state mirroring
- a presentation unit:
  - one display responsibility, such as view shaping, status mapping, or state-specific rendering

### How To Extract Candidates

- Start from the user task, not from file boundaries.
- Group code by shared reason to change, not by physical proximity.
- Pull out only the units where boundary choice matters.
- Ignore code that is already obviously well-placed and coherent.

### Candidate Extraction Questions

- What user action or system response is this code supporting?
- If requirements changed, which part of this code would change first?
- Does this code contain more than one reason to change?
- Can this unit be named as one behavior, one decision, one sync rule, or one presentation responsibility?

### Bad Candidates

- arbitrary line ranges
- entire files with several unrelated responsibilities
- tiny helpers with no real boundary decision
- "whatever feels messy" without naming the responsibility

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

## Layer Placement Rules

After extracting and classifying a responsibility candidate, place it by meaning, not by convenience.

### Layer Targets

- utility:
  - framework-agnostic, business-agnostic, transport-agnostic logic that could reasonably live as a small library helper
- api:
  - transport, request, response, DTO, endpoint, and server contract handling
- domain:
  - product meaning, business rules, calculations, permissions, validation policies, invariants
- feature:
  - user-facing flow orchestration that combines UI behavior, async work, and domain/application decisions for one feature or slice
- component-local:
  - rendering, local interaction, and tightly scoped view behavior that should stay near the component

### Placement Questions

- Would this still make sense in another project with no product context?
  - if yes, it may be utility
- Is this responsibility mainly about server contract or transport lifecycle?
  - if yes, it may be api
- Would this responsibility survive a complete UI redesign of the same product?
  - if yes, it may be domain
- Is this responsibility coordinating a real user flow for this feature?
  - if yes, it may be feature
- Is this responsibility only meaningful inside this component's view boundary?
  - if yes, keep it component-local

### Placement Guardrails

- Do not call something reusable just because code looks similar.
- Reusable means repeated responsibility with stable ownership and the same reason to change.
- Do not move feature flow into domain just to make files look cleaner.
- Do not move product meaning into utility.
- Do not hide API response handling inside view code when it is part of transport interpretation.

## Hook Extraction Rules

Do not treat hook extraction as default cleanup.
Evaluate hook extraction only after the responsibility candidate is named and its target layer is understood.

### Hook-Suitable Candidates

A candidate is hook-suitable when it is mainly:

- a state transition unit
- an effect synchronization unit
- a reusable interaction behavior
- an async view-adapter behavior that is still UI-facing

### Hook Decision Questions

- Does this candidate depend on React state or effect lifecycle in a meaningful way?
- Would extracting it make ownership clearer, not just move code elsewhere?
- Would the calling component become easier to read after extraction?
- Is the extracted API explainable in a few lines?
- Is this still UI-facing behavior rather than domain, api, or feature flow logic?

### Prefer Other Targets When

- the candidate is a pure calculation:
  - prefer utility or domain
- the candidate is transport or server-contract handling:
  - prefer api
- the candidate coordinates a full user journey across several boundaries:
  - prefer feature
- the candidate only makes sense inside one component's render boundary:
  - keep it component-local

### Hook Smells

- props go into the hook and come back out almost unchanged
- the hook hides complexity without reducing it
- the hook mixes domain rules, async lifecycle, and UI orchestration with no clear center
- the hook is reusable only in theory
- the call site becomes harder to understand
- the hook still requires component-specific DOM or markup assumptions to make sense

### Reusable Hook Rule

Do not call a hook reusable because similar code exists twice.
Treat a hook as reusable only when:

- the same behavior contract repeats
- the same reason to change applies
- the ownership stays stable after extraction
- the extracted inputs and outputs are explainable without the original component's full context

### Hook Extraction Decision Rule

Recommend hook extraction only when at least one of these is true:

- the candidate is a stable UI behavior that can be isolated cleanly
- the candidate repeats across components with the same contract
- the candidate is tightly tied to React lifecycle but not to one render structure
- the extraction makes testing or reasoning materially easier

Do not recommend hook extraction when the main result is:

- file splitting without ownership improvement
- stronger coupling between the component and an opaque hook
- hiding a better utility, domain, api, or feature placement
- abstracting a one-off screen behavior too early

## Component API Rules By Level

Do not evaluate a component API in the abstract.
First decide whether the component is library-level or product-level, because the right API shape depends on that level.

### Component Level Detection

Treat a component as library-level when most of these are true:

- it has little or no domain language in its contract
- it is intended for repeated use across many features or screens
- its main job is structural, visual, or interaction reuse
- its API should express constrained flexibility rather than one feature's exact workflow

Treat a component as product-level when most of these are true:

- it carries feature or domain language in its name or props
- it exists mainly for one product workflow or feature slice
- its API should express domain intent more than general flexibility
- extracting it into shared space would weaken meaning or ownership clarity

### Library-Level API Rules

- prefer constrained variants over open-ended config dumping
- use composition when the structural flexibility is real and stable
- keep tokens, slots, and interaction contracts explicit
- avoid leaking product or feature meaning into the API
- prefer a small set of allowed extension points over unlimited escape hatches

### Product-Level API Rules

- prefer semantic prop names that reflect domain or feature meaning
- do not generalize an API just because another feature looks similar
- prefer explicit feature intent over reusable-looking abstraction
- expose semantic events rather than low-level internal state changes
- keep the API close to the current workflow unless repeated evidence justifies promotion

### Cross-Level API Smells

- prop explosion caused by trying to satisfy unrelated use cases
- boolean props representing hidden state machines or domain meaning
- generic names like `data`, `config`, `item`, or `meta` where domain language should exist
- children or slot patterns used to avoid deciding ownership
- callback surfaces that expose implementation details instead of meaningful events
- variant systems that hide product-level behavioral differences

### Promotion Rule

Do not promote a product-level component into shared or library space unless:

- the same responsibility clearly repeats across features
- the same reason to change applies across those usages
- the extracted API remains easy to explain
- domain meaning is not erased by generalization

If those conditions are weak, keep the component product-level and optimize for clarity instead of reuse theater.

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
- `Responsibility candidates`
- `Concern classification`
- `Responsibility ownership decision`
- `Layer placement decision`
- `Hook extraction decision`
- `Component level detection`
- `API surface assessment`
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
- Do not extract a hook before naming the responsibility candidate and its target layer.
- Do not evaluate a component API before deciding whether the component is library-level or product-level.
- Prefer composition over inheritance-style abstraction.
- Prefer existing project patterns unless they are actively causing confusion.
- Prefer boundaries that match the current project structure unless that structure is the source of the problem.
- Do not choose a target layer before naming the responsibility candidate being moved.
- Keep boundaries understandable to the next maintainer.

## References

- Read [references/boundary-checklist.md](references/boundary-checklist.md) for split criteria.
- Read [references/concern-classification.md](references/concern-classification.md) when rendering, orchestration, domain, and async responsibilities are easy to confuse.
- Read [references/responsibility-ownership-rules.md](references/responsibility-ownership-rules.md) when deciding where rendering, domain, async, formatting, and state responsibilities should live.
- Read [references/layer-placement-rules.md](references/layer-placement-rules.md) when deciding between utility, api, domain, feature, and component-local placement.
- Read [references/hook-extraction-rules.md](references/hook-extraction-rules.md) when deciding whether a candidate should stay local, become a hook, or move to another layer.
- Read [references/component-api-smells.md](references/component-api-smells.md) when evaluating library-level versus product-level component APIs.
