---
name: react-architecture
description: React architecture guidance for separating React-facing code into system, design, and domain axes, choosing dependency direction, and deciding implementation versus refactor order after the overall pattern is already chosen.
---

# React Architecture

## Overview

Use this skill when the main question is how React-facing artifacts should separate `system`, `design`, and `domain` concerns inside an already chosen architecture pattern.
The goal is not to defend an existing folder layout. The goal is to classify responsibilities by why they exist, keep dependency direction coherent, and make implementation order different from cleanup order when that produces cleaner code.
Do not use this skill as the first stop for top-level structure-pattern choice. Use it after that question is settled or out of scope.

## Use When

- the task is about component, hook, utility, page, or global-state placement inside React code
- responsibilities are mixed and you need a real separation rule instead of local cleanup instincts
- rerender spread, effect misuse, hook sprawl, or provider growth are symptoms of mixed axes
- you need to decide which concern should be extracted first and in what order code should be built or cleaned up

## Do Not Use When

- the main problem is still top-level frontend structure choice
- the design rules themselves are still undefined
- the domain policy itself is still undefined
- the task is only about a tiny local split and does not need an axis model

## Workflow

1. Confirm the repository's current structure context only enough to respect existing boundaries while doing axis separation.
2. Enumerate the responsibilities being discussed instead of starting from file names.
3. Classify each responsibility by cause of change:
   - `system`: React lifecycle, browser APIs, runtime synchronization, rendering/update boundaries
   - `design`: reusable UI contracts, interaction models, design-system semantics, visual state contracts
   - `domain`: page meaning, feature workflow, business policy, product state, orchestration
4. Check dependency direction:
   - `domain -> design -> system`
   - `system` must not know `design` or `domain`
   - `design` may use `system`, but must not absorb `domain`
5. Decide the target owner for each responsibility:
   - component
   - hook
   - utility
   - page or route boundary
   - context or broader state owner
6. Choose the sequence:
   - for new spec or feature work: `design -> domain -> system`
   - for refactor or cleanup: `system -> design -> domain`
7. End with a structure recommendation that explains the axis map, dependency direction, and why the chosen owner fits.

## Core Model

- `System axis`
  - owns React mechanics, browser coupling, effect synchronization, render/update boundaries, provider churn, platform adapters
  - examples:
    - `useResizeObserver`
    - URL sync adapter
    - focus trap
    - external store subscription wrapper
- `Design axis`
  - owns reusable UI contracts, reusable interaction semantics, component API shape, variant rules, visual state contracts
  - examples:
    - `Card` API
    - `Button` variants
    - selection interaction model shared across reusable cards
    - slot and token mapping
- `Domain axis`
  - owns page meaning, feature workflow, product state, business-specific orchestration, query composition for one feature
  - examples:
    - checkout flow coordinator
    - product list filter state
    - permission-aware view logic
    - feature-level query shaping

## Classification Tests

- `Cause-of-change test`
  - Ask what kind of change would force this code to change first.
  - React or browser change means `system`.
  - reusable UI contract change means `design`.
  - product meaning or workflow change means `domain`.
- `Deletion test`
  - If the feature meaning disappears but the code still makes sense, it is not `domain`.
  - If the reusable UI contract disappears but the runtime synchronization still makes sense, it is `system`.
- `Leak test`
  - If a supposedly reusable artifact mentions feature policy, route meaning, permissions, or page workflow, `domain` leaked downward.
  - If a supposedly domain artifact is mostly event wiring, subscriptions, or observers, `system` leaked upward.
- `Naming warning`
  - Do not classify by file or artifact name alone.
  - `page`, `global state`, and custom hook names are hints, not proof.

## Artifact Guidance

### Components

- Reusable primitives and reusable composites usually live on the `design` axis.
- Page sections, feature panels, and route assemblies usually live on the `domain` axis.
- Browser-coupled wrappers, measurement shells, and synchronization-only containers usually live on the `system` axis.
- Split when one component changes for more than one axis reason.

## Component Contracts

- `Event prop naming`
  - Function props should be treated as event listeners and use the `on-` prefix shape such as `onConfirm`, `onClose`, or `onSelect`.
  - Name the prop from the point of view of the component that owns the prop contract, not from the parent or feature that happens to use it.
  - A modal that exposes a confirm action should usually prefer `onConfirm` over `onProjectConfirm` because the contract belongs to the modal, not the surrounding feature.
- `Contract viewpoint rule`
  - The prop name should describe what happens at the boundary of this component.
  - Do not smuggle parent workflow meaning into a reusable component contract unless that meaning is truly part of the component itself.
- `Transport choice rule`
  - Do not assume every state or data handoff must be solved by props alone.
  - Evaluate `context`, prop drilling, broader external state, backend connectors, and internal hooks as separate options with different ownership consequences.
  - Internal hooks are a valid first-class option when the component needs local coordination, local data shaping, or local integration ownership without widening the public prop contract.
- `Prop pressure smell`
  - If a component keeps gaining props just to thread state, backend data, or orchestration concerns through it, stop and check whether a local hook or another owner would express the boundary more honestly.
- `Good default`
  - Keep the public prop contract focused on the component's external API.
  - Move local coordination and local integration detail behind internal hooks when that reduces axis leakage and keeps the component contract readable.

### Hooks

- `System hook`
  - owns one synchronization or platform boundary
  - examples:
    - observers
    - storage sync
    - media-query sync
    - focus or scroll management
- `Design hook`
  - owns one reusable interaction model or UI behavior contract
  - examples:
    - card selection model shared across reusable card patterns
    - disclosure behavior used by several reusable surfaces
- `Domain hook`
  - owns one feature workflow, page orchestration, or product-facing state machine
  - examples:
    - checkout flow
    - dashboard filter orchestration
    - search result coordination
- `Hook ownership rule`
  - a custom hook is not its own axis
  - classify the hook by the first change that should force it to move
  - browser/runtime synchronization belongs in `system`
  - reusable interaction semantics belong in `design`
  - feature workflow and product meaning belong in `domain`
- `Hook composition rule`
  - `domain` hooks may compose `design` and `system` hooks
  - `design` hooks may compose `system` hooks
  - `system` hooks must stay terminal and must not compose `design` or `domain` knowledge
- `Hook refactor rule`
  - when one hook mixes axes, pull out `system` work first:
    - effects
    - subscriptions
    - observers
    - storage or URL sync
  - then extract reusable interaction or UI behavior into `design` hooks
  - leave feature workflow, query shaping, and product decisions in the `domain` hook or feature boundary
- `Local hook warning`
  - being internal to one component does not make a hook automatically `design` or `domain`
  - local hooks can still be `system` if they mostly own synchronization

### Utilities

- `System utility`
  - runtime helpers, adapter shaping, platform-normalization helpers
- `Design utility`
  - slot interpretation, variant resolution, token selection, reusable visual mapping
- `Domain utility`
  - feature mapping, domain-specific normalization, policy-shaped derivation

### Pages and Global State

- A page file is not automatically `domain`. It may temporarily host `design` and `system` responsibilities during implementation, but those should remain identifiable.
- Global state is not automatically `domain`. Shared visibility or interaction state for a reusable surface can still be `design`, and runtime-backed synchronization state can still be `system`.
- Treat broader state as `domain` only when the source of truth carries real product meaning.

## Ordering Rules

- `Implementation order`
  - Use `design -> domain -> system` for new spec or feature work.
  - First define the reusable UI contract.
  - Then bind product meaning and workflow.
  - Then attach React and browser synchronization.
- `Refactor order`
  - Use `system -> design -> domain` for cleanup.
  - First pull runtime mechanics down to stable owners.
  - Then extract reusable UI contracts.
  - Leave product meaning closest to the feature boundary.
  - For mixed hooks, split synchronization first, reusable interaction second, and feature orchestration last.
- `Structural priority`
  - For components, hooks, and utilities, preserve `system` cleanliness first, `design` clarity second, and `domain` locality last.
  - This does not mean `system` is more important to the product. It means lower-level contamination is more expensive to clean up later.

## React-Specific Rules

- Effects are for synchronization, so they usually belong to the `system` axis unless they are just misplaced domain flow.
- Component event props should read like component-owned event listeners, not feature-owned workflow labels.
- Before widening a prop surface, check whether an internal hook would keep the contract smaller while preserving the right owner.
- Provider scope is an axis question:
  - reusable surface coordination is usually `design`
  - feature or route orchestration is usually `domain`
  - subscription and external synchronization layers are usually `system`
- Rerender spread is a boundary smell before it is an optimization problem.
- Hook extraction is not the default answer. First decide whether the mixed responsibility should stay together at all.
- When a hook feels unclear, inspect its effects and subscriptions before renaming or relocating it. Mixed hooks usually become clearer after `system` work is peeled away.

## Stress-Test Examples

- `ProductCard`
  - card API, slots, hover/selection pattern: `design`
  - stock badge meaning, navigation decision, pricing policy: `domain`
  - measurement, observer, focus sync: `system`
- `CreateProjectModal`
  - prefer `onConfirm` and `onClose` at the modal boundary
  - avoid `onProjectConfirm` when the "project" meaning belongs to the parent flow rather than the modal contract itself
  - if the modal needs local form state, validation wiring, or submit coordination, consider an internal hook before expanding the public prop surface
- `useCheckoutFlow`
  - checkout steps, guard rules, submission phases: `domain`
  - URL sync, storage sync, pending subscription cleanup: `system`
- `useProductCardController`
  - hover, selection, disclosure semantics reused across card surfaces: `design`
  - product navigation choice and inventory policy: `domain`
  - resize observer or focus sync: `system`
- `ModalStackStore`
  - reusable overlay stacking and open-close coordination: `design`
  - checkout-confirmation modal sequence and business guardrails: `domain`
  - portal or focus synchronization: `system`

## Output Contract

- `Artifact map`
- `Axis classification`
- `Dependency direction`
- `Implementation order`
- `Refactor order`
- `Main leak risk`
- `Residual risk`

## Guardrails

- Do not classify by filename, hook prefix, or page location.
- Do not let `domain` leak into reusable `design` surfaces.
- Do not let `design` or `domain` concerns contaminate `system` adapters and synchronization layers.
- Do not treat broader state as automatically `domain`.
- Do not use extraction as a substitute for deciding ownership.
- Do not keep mixed axes in one unit after the change reasons clearly diverge.
- Do not turn a local cleanup into a fake axis rewrite when the problem is still just one local owner.
- Do not let page-level convenience justify leaking product meaning into reusable UI.
- Do not name event props from the parent's workflow point of view when the contract belongs to the child component.
- Do not keep adding props to solve local coordination if an internal hook would keep the boundary smaller and clearer.
- Do not use `memo`, `useMemo`, or `useCallback` as the first fix for mixed axes.
- Do not treat provider widening as a harmless shortcut.
