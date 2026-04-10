---
name: react-architecture
description: React architecture guidance for hierarchy design, component levels, design-system boundaries, prop contracts, and rendering/update boundaries after the overall pattern is already chosen. Use when a task is specifically about React component layering, Template/Page boundaries, separating business logic from reusable UI logic, or reducing unnecessary re-render spread through better structure inside the selected architecture.
---

# React Architecture

## Overview

Use this skill when the main question is about React structure rules inside an already chosen architecture pattern.
The goal is to decide how React components should be layered, what belongs in reusable UI versus product layers, how component contracts should stay readable as the system grows, and how rendering updates should stay bounded by better structure.
Do not use this skill as the first stop for choosing between Atomic Design, FSD, or a custom hybrid. Use it after that codebase-structure question is already settled or out of scope.

## What This Skill Teaches

This skill teaches React architecture as a set of structural mental models rather than a folder convention.
The focus is how components, hooks, context, external state, effects, and update propagation shape code quality.
Use it to explain why some React structures stay maintainable while others become fragile even when the folder pattern looks correct.

## Core Mental Models

- `Ownership before extraction`
  - Decide who owns state, side effects, orchestration, and rendering before extracting hooks or wrapper components.
  - A neat extraction on top of weak ownership usually preserves the real problem.
- `Render propagation is architectural`
  - Rerender spread is not only a performance detail; it is a sign of where state and providers are placed.
  - Structural fixes usually outperform memoization-first fixes.
- `Composition is a boundary tool`
  - Composition should narrow responsibility, not create layers of indirection with no ownership clarity.
  - Good composition makes data flow easier to explain.
- `Effects are for synchronization`
  - Effects should synchronize React with external systems, subscriptions, browser APIs, or async lifecycles.
  - Effects should not become a second render pipeline for pure derivation or ordinary event handling.
- `Shared UI is not product logic`
  - Reusable UI can expose interaction surfaces and display rules.
  - Product rules, workflow decisions, and feature-specific policy should stay in product-facing layers.

## Workflow

1. Confirm the repository's chosen architecture context:
   - Atomic Design
   - FSD
   - custom or hybrid
2. Confirm the component level being discussed:
   - primitive or foundation component
   - composed UI component
   - section or screen-level composition
   - Template or Page boundary
3. Review the architecture by axis:
   - components
   - hooks
   - context
   - utility categories
   - external state
   - rendering boundaries
   - effects
4. Separate reusable UI logic from product-specific business logic.
5. Decide which layer should own composition, data access, and product decisions.
6. End with a React-specific structure recommendation that fits the existing repository rules.

## Architecture By Axis

### Component Architecture

- `Component roles`
  - `Shared UI component`
    - owns reusable display structure and interaction contracts
    - should stay free of product workflow and business policy
  - `Feature-aware component`
    - understands one feature's local meaning and can bind product-facing behavior to UI
    - should not pretend to be generic shared UI
  - `Composition component`
    - arranges several child parts into one coherent section, panel, or flow boundary
    - should own composition decisions rather than low-level reusable contracts
  - `Screen or route boundary`
    - owns route-level orchestration, data wiring, and feature assembly
    - should not collapse into generic reusable UI language
- `Preferred patterns`
  - Use composition to narrow responsibility instead of widening one component with configuration flags.
  - Keep shared UI responsible for display rules and interaction surfaces, not product policy.
  - Keep component contracts small, explicit, and semantically clear.
- `When to split`
  - split when one component mixes shared UI responsibility with feature policy
  - split when composition responsibility and reusable contract responsibility change for different reasons
  - split when the prop contract begins acting like mode switching for several conceptual components
  - split when one user-facing area contains several owners with different change pressure
- `When not to split`
  - do not split only because the file is long if ownership is still coherent
  - do not split when extraction only moves local context into props and wrappers
  - do not split when a utility extraction solves the readability problem without creating a new component owner
  - do not split when the extracted part would have no stable contract outside its current local context
- `Contract smells`
  - several boolean props that combine into hidden modes
  - mutually exclusive props that signal several concepts are sharing one API
  - callbacks that expose parent implementation details instead of user-facing events
  - generic-looking props that actually carry feature policy or workflow branching
- `Shared UI guardrails`
  - shared UI may expose slots, variants, and interaction hooks
  - shared UI should not decide permissions, workflow branching, or product-specific policy
- `Anti-patterns`
  - boolean-heavy prop contracts that simulate several components
  - wrapper layers that add indirection without ownership
  - shared UI components that contain business workflow or feature policy

### Hook Architecture

- `Preferred patterns`
  - Extract a hook only when it owns a real reusable behavior, sync boundary, or orchestration unit.
  - Keep the hook boundary explicit: what state it owns, what it synchronizes, and what contract it exposes.
  - Prefer ordinary functions for pure computation and shaping logic when React lifecycle is not needed.
- `Anti-patterns`
  - hooks used as dumping grounds for logic that has no clear ownership
  - effect-heavy hooks that hide fragile state machines
  - extracting a hook before deciding whether the responsibility should stay local

### Context Architecture

- `Preferred patterns`
  - Treat each provider as one intentional ownership and update boundary.
  - Scope providers to the smallest subtree that actually needs the shared value.
  - Keep context values focused and stable enough that consumers depend on one coherent concern.
- `Anti-patterns`
  - one broad provider for many unrelated concerns
  - using context as a reflex instead of deciding real ownership
  - provider values that churn because ownership and update boundaries are vague

### Utility Categories

- `React utility`
  - Use for pure helper logic that supports React component contracts without needing lifecycle or ownership of its own.
  - Good fits include slot interpretation, prop shaping, render-time helper decisions, and small composition helpers.
  - Keep these near the component or feature boundary they support unless they become truly shared across several UI surfaces.
- `Library utility`
  - Use for framework-agnostic pure functions such as formatting, parsing, normalization, sorting, grouping, or mapping.
  - Prefer this category when the logic is plain input-output transformation with no React or browser dependency.
  - Shared placement is justified only when the meaning is also shared, not just the implementation shape.
- `Browser utility`
  - Use for helpers around URL, DOM, storage, clipboard, media queries, observers, or other web platform APIs.
  - Treat many of these as platform adapters rather than generic utilities when they hide browser coupling.
  - Do not blur browser integration into ordinary pure helpers.
- `Style utility`
  - Use for variant mapping, token selection, class composition, and other appearance-focused helper logic.
  - Keep style helpers close to shared UI or design-system space.
  - Do not let style utilities absorb product policy or feature workflow.
- `Feature-local utility`
  - Use when the logic is pure enough to be a helper but still expresses one feature's private meaning.
  - Keep it local when moving it to shared space would erase ownership clarity.
  - Reuse alone is not enough reason to globalize a feature-local utility.

## Utility Boundary Rules

- `Utility extraction is not one thing`
  - Distinguish React, library, browser, style, and feature-local utilities before extracting.
  - A generic `utils` bucket weakens ownership because these categories have different coupling and placement rules.
- `Lifecycle test`
  - If logic needs React lifecycle, subscriptions, or synchronization, it is not a plain utility.
  - Consider a hook, adapter, or owner component instead.
- `Ownership test`
  - If extraction changes who owns state, orchestration, or business meaning, it is not just utility extraction.
  - Solve the ownership question before moving files.
- `Platform test`
  - If logic touches browser APIs, treat it as platform-coupled first and utility-shaped second.
  - Avoid presenting browser integration as an innocent shared helper.
- `Meaning test`
  - If the logic expresses feature policy or domain meaning, keep it feature-local or move it to a more explicit boundary.
  - Do not globalize code just because it is pure.

### External State Architecture

- `Preferred patterns`
  - Keep the distinction between server state, local UI state, and shared client state explicit.
  - Put data in an external store only when several distant owners need one durable client-side source of truth.
  - Keep store shape focused on shared ownership, not convenience dumping.
- `Anti-patterns`
  - putting local transient UI state into a global store by default
  - storing pure derived values that can be computed from existing state
  - letting the external store absorb product orchestration that should stay closer to the feature boundary

### Rendering Architecture

- `Preferred patterns`
  - Put state near the narrowest subtree that actually coordinates it.
  - Treat context providers as update boundaries and keep them narrow.
  - Fix ownership and composition before reaching for memoization.
- `Anti-patterns`
  - lifted state by default
  - parent churn from overly broad owners
  - prop identity churn caused by vague ownership
  - memoization-first architecture

### Effect Architecture

- `Preferred patterns`
  - Use an effect only when React must synchronize with something outside render.
  - Let event handlers own user-driven flow and let render own pure derivation.
  - Keep effects small, purpose-specific, and tied to one synchronization concern.
- `Anti-patterns`
  - effect for pure derivation
  - effect for basic event-triggered state transitions
  - effect as a hidden control-flow layer
  - resetting state through effect when ownership or keys should solve it

## Rendering Boundary Rules

- `State locality rule`
  - Put state near the narrowest subtree that actually coordinates it.
  - Treat every upward state move as a tradeoff that must justify broader updates.
- `Provider scope rule`
  - Treat context providers as update boundaries.
  - If unrelated consumers rerender from one provider update, the scope is probably too broad.
- `Parent churn rule`
  - If a frequently updating parent owns too much rendering responsibility, split ownership before optimizing leaves.
- `Prop identity rule`
  - New object, array, or callback identities are often a symptom of overly broad composition or poor ownership, not the primary disease.
- `Memoization rule`
  - Reach for memoization after the ownership and update boundaries are coherent, not before.

## Effect Discipline

- Use an effect when React must synchronize with something outside render:
  - subscriptions
  - DOM or browser APIs
  - async lifecycle coordination
  - timers or observers
- Do not use an effect for:
  - pure derivation from props and state
  - basic event-triggered state transitions
  - resetting state that should be keyed or owned differently
  - encoding business logic that belongs in normal control flow

## Decision Rules

- Use the project's current architecture pattern as the source of truth; do not reopen Atomic versus FSD choice unless the current pattern is clearly failing.
- Keep low-level reusable components free of data fetching, business policy, and feature-specific copy.
- Keep reusable UI logic in design-system or shared UI layers; keep product-specific behavior in Pages, features, widgets, or app-specific composition layers as the chosen architecture requires.
- Treat Templates as composition or example boundaries only; do not let Templates become hidden business-logic containers.
- Choose Page, feature, or widget ownership for product decisions, route data, and workflow orchestration according to the existing repository rules.
- Keep state as local as practical to reduce unnecessary render spread; do not lift state higher unless coordination actually requires it.
- Treat context boundaries as architectural boundaries; do not use one broad provider when unrelated consumers can be split.
- Reduce prop identity churn through clearer ownership and narrower contracts before reaching for memoization.
- Prefer structural fixes over memoization-first fixes when render churn is caused by poor ownership, broad state placement, or mixed responsibilities.
- Use memoization only when the boundary is already coherent and the remaining rerender cost is still material.
- Prefer clear component contracts over flexible but ambiguous prop surfaces.
- Keep props small and explicit; use semantic literal unions for visual or interaction states when a boolean state flag would blur meaning.
- Preserve dependency direction: lower-level React UI must not import higher-level product layers.

## Output Contract

- `Architecture context`
- `Mental model`
- `Current hierarchy fit`
- `Component level`
- `Business vs UI split`
- `Rendering boundary`
- `Prop contract decision`
- `Dependency direction`
- `Recommended structure`
- `Residual risk`

## Review Questions

- `Components`
  - Does this component own one coherent responsibility?
  - Does this shared UI component expose a clean contract, or does it smuggle product policy downward?
- `Hooks`
  - Does this hook own a real reusable lifecycle or orchestration concern?
  - Is this hook clarifying ownership, or hiding a weak boundary?
- `Context`
  - Is this provider scoped to one concern, or is it grouping unrelated consumers?
  - Is context being used because ownership is truly shared, or because prop flow feels inconvenient?
- `External state`
  - Does this value need a shared durable owner, or is it only local UI state?
  - Is this store holding source-of-truth state or a pile of derived convenience values?
- `Utilities`
  - Is this logic a React utility, library utility, browser utility, style utility, or feature-local utility?
  - Does extraction improve clarity without erasing ownership or hiding platform coupling?
- `Rendering`
  - Who owns this state, and why is that owner the narrowest meaningful owner?
  - If this state changes, which subtree rerenders, and is that blast radius justified?
  - Is this memoization protecting a coherent boundary, or hiding a weak one?
- `Effects`
  - Could this value be derived during render instead of stored through state and effect?
  - Is this effect synchronizing with the outside world, or acting as hidden control flow?

## Guardrails

- Do not turn a local refactor question into a full architecture rewrite unless the repo is clearly blocked by the current structure.
- Do not treat Template, Page, feature, widget, and shared UI as interchangeable labels.
- Do not move product logic into reusable UI just to reduce file count.
- Do not widen one component contract until it acts as several components behind variants and booleans.
- Do not treat `memo`, `useMemo`, or `useCallback` as the default fix before checking state placement, provider scope, and ownership.
- Do not widen provider scope just to avoid prop passing if it causes avoidable rerender spread.
- Do not use this skill when the real problem is choosing a folder or layer pattern, planning an architecture migration, or evaluating whether the repository needs Atomic, FSD, or a hybrid.
