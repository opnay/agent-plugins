# Responsibility Ownership Rules

Use this guide to decide where frontend responsibilities should live after concern classification.

## Ownership Map

- Rendering ownership:
  - visible structure, local interaction response, display conditions
- Orchestration ownership:
  - flow coordination, cross-section event handling, parent-child sequencing
- Domain ownership:
  - product meaning, business rules, policy decisions, calculations
- Async ownership:
  - request lifecycle, mutation lifecycle, synchronization, invalidation, loading and failure
- Formatting ownership:
  - display shaping, labels, view-facing mapping, presentation adapters
- State ownership:
  - local UI state, shared UI state, derived values, form state, server-adjacent state

## Placement Heuristics

- Keep a responsibility near the code that can explain why it exists.
- Move a responsibility only when the current owner cannot explain or safely change it.
- Prefer the smallest owner that has the full context to make the decision.
- Do not widen ownership just to avoid thinking about boundaries.

## State Ownership Subguide

- Local state:
  - keep it close when one boundary reads and writes it
- Shared state:
  - lift only to the smallest common owner that truly coordinates it
- Derived state:
  - compute it instead of storing it when possible
- Form state:
  - isolate it when validation and submit lifecycle become important
- Server-adjacent state:
  - avoid copying request results into local state unless the UI is intentionally forking or editing the data

## Wrong Ownership Examples

- A leaf component owns request retries and cache refresh logic.
- A page component owns visual toggle state for deeply isolated widgets.
- A render helper embeds permission rules that should survive layout changes.
- A parent owns five callback handlers but cannot explain the flow they coordinate.
- A component stores filtered or sorted data that could be derived from source inputs.

## Decision Test

Ask:

- Who changes this responsibility when requirements change?
- Who has enough context to own it safely?
- If this responsibility moved, would the system become easier to explain?
- If this responsibility stayed, would the current owner remain coherent?

If those answers are weak, the ownership is probably wrong.
