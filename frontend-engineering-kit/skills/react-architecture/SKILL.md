---
name: react-architecture
description: React architecture guidance for component and hook boundaries, context scope, state ownership, effect discipline, prop contracts, and rendering/update boundaries after the overall pattern is already chosen. Use when a task is specifically about React component layering, hook API design, effect/dependency issues, separating business logic from reusable UI logic, or reducing unnecessary re-render spread through better structure inside the selected architecture.
---

# React Architecture

## Overview

Use this skill when the main question is about React structure rules inside an already chosen architecture pattern.
The goal is to decide how React components should be layered, what belongs in reusable UI versus product layers, how component contracts should stay readable as the system grows, and how rendering updates should stay bounded by better structure.
Do not use this skill as the first stop for choosing between Atomic Design, FSD, or a custom hybrid. Use it after that codebase-structure question is already settled or out of scope.

## What This Skill Teaches

This skill teaches React architecture as a set of structural mental models rather than a folder convention.
The focus is how components, hooks, context, utilities, external logic boundaries, state ownership, effects, and update propagation shape code quality.
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
   - external logic boundaries
   - state ownership
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

- `Hook roles`
  - `State hook`
    - owns one local stateful behavior or interaction model
    - should keep one coherent owner for state transitions
  - `Synchronization hook`
    - owns one synchronization boundary with browser APIs, subscriptions, async lifecycle, or external systems
    - should not absorb unrelated product workflow
  - `Orchestration hook`
    - coordinates one feature-local interaction flow or state machine
    - should stay bounded to one meaningful workflow instead of becoming a page dump
  - `Adapter hook`
    - wraps one external store, query library, or platform integration in a React-friendly contract
    - should clarify the boundary instead of hiding several concerns behind one name
- `Preferred patterns`
  - Extract a hook only when it owns a real reusable behavior, sync boundary, or orchestration unit.
  - Keep the hook boundary explicit: what state it owns, what it synchronizes, and what contract it exposes.
  - Prefer ordinary functions for pure computation and shaping logic when React lifecycle is not needed.
- `When to extract`
  - extract when one reusable lifecycle, synchronization concern, or interaction model needs an explicit owner
  - extract when the hook contract makes the caller's responsibility easier to explain
  - extract when several call sites truly share one behavioral boundary rather than one implementation detail
- `When not to extract`
  - do not extract when the logic is only pure computation or shaping that belongs in a utility
  - do not extract when the real issue is a weak component boundary and the hook would only relocate the same mixed responsibilities
  - do not extract when the hook would have no stable contract beyond one local component's private sequencing
- `Contract rules`
  - return a contract that reflects one coherent owner, not a bag of unrelated state and handlers
  - expose user-meaningful actions and states rather than leaking internal implementation plumbing
  - keep read values, write actions, and status signals understandable at the call site
- `Contract smells`
  - one hook returning many unrelated booleans, callbacks, refs, and data slices
  - hook names that sound like page replacement rather than one boundary
  - hook APIs that mirror a component's entire internal state shape
- `Implementation rules`
  - obey Rules of Hooks and keep hook calls at stable top-level positions
  - avoid derived state duplication; compute pure derived values in render or memo
  - use effects only for synchronization, not for ordinary computation
  - prefer explicit status or phase models over piles of booleans
- `Dependency and closure rules`
  - make dependency arrays complete and predictable
  - treat stale closures as correctness problems, not as minor optimization details
  - refactor ownership or logic placement before suppressing dependency issues
  - stabilize returned references only when consumers actually depend on referential equality
- `Async and cleanup rules`
  - add cancellation, latest-only guards, or cleanup when the hook owns async or subscription flows
  - always return cleanup for subscriptions, observers, or long-lived external connections
  - keep async ownership explicit so the caller can explain what happens on rerender, restart, and unmount
- `Anti-patterns`
  - hooks used as dumping grounds for logic that has no clear ownership
  - effect-heavy hooks that hide fragile state machines
  - extracting a hook before deciding whether the responsibility should stay local
  - hiding a component responsibility problem inside a large `useXxx` file instead of fixing the component boundary

### Context Architecture

- `Context roles`
  - `Read-only context`
    - provides stable read-mostly values such as theme, locale, or static configuration
    - should avoid unnecessary write churn across broad trees
  - `Coordination context`
    - shares one interaction owner across several related consumers inside one subtree
    - should stay scoped to the feature, panel, or section that actually coordinates the interaction
  - `Capability context`
    - exposes one capability surface such as analytics, feature capability, or service access
    - should stay explicit about what consumers can do with it
  - `Boundary context`
    - defines one local ownership and update boundary for a feature, widget, or screen subtree
    - should not silently expand into application-wide shared state
- `Preferred patterns`
  - Treat each provider as one intentional ownership and update boundary.
  - Scope providers to the smallest subtree that actually needs the shared value.
  - Keep context values focused and stable enough that consumers depend on one coherent concern.
- `When to use context`
  - use context when several consumers inside one meaningful subtree truly need one shared owner
  - use context when prop passing would obscure a real shared boundary rather than simply add one explicit prop hop
  - use context when the provider itself explains one clear feature, widget, or screen-level boundary
- `When not to use context`
  - do not use context only because prop drilling feels inconvenient
  - do not use context when the state is still local to one owner and can stay there cleanly
  - do not use context to bundle unrelated concerns under one broad provider
  - do not use context when a store, hook, service adapter, or explicit props describe the ownership boundary more clearly
- `Contract rules`
  - keep provider values organized around one coherent concern
  - separate stable read values from write actions when that improves consumer clarity
  - design provider values so consumers can explain why they subscribe in one sentence
  - treat provider value growth as a signal to revisit the boundary rather than keep adding fields
- `Contract smells`
  - a provider value that keeps accumulating unrelated booleans, callbacks, and data
  - consumers that read the same context for unrelated reasons
  - provider updates that fan out to large areas of the tree without one shared interaction boundary
- `Anti-patterns`
  - one broad provider for many unrelated concerns
  - using context as a global store replacement by default
  - using context as a reflex instead of deciding real ownership
  - provider values that churn because ownership and update boundaries are vague

### Utility Categories

- `React utility`
  - Use for pure helper logic that supports React component contracts without needing lifecycle or ownership of its own.
  - Good fits include slot interpretation, prop shaping, render-time helper decisions, and small composition helpers.
  - Keep these near the component or feature boundary they support unless they become truly shared across several UI surfaces.
- `Pure utility`
  - Use for framework-agnostic pure functions such as formatting, parsing, normalization, sorting, grouping, or mapping.
  - Prefer this category when the logic is plain input-output transformation with no React or platform dependency.
  - Shared placement is justified only when the meaning is also shared, not just the implementation shape.
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
  - Distinguish React, pure, style, and feature-local utilities before extracting.
  - A generic `utils` bucket weakens ownership because these categories have different coupling and placement rules.
- `Lifecycle test`
  - If logic needs React lifecycle, subscriptions, or synchronization, it is not a plain utility.
  - Consider a hook, adapter, or owner component instead.
- `Ownership test`
  - If extraction changes who owns state, orchestration, or business meaning, it is not just utility extraction.
  - Solve the ownership question before moving files.
- `Platform test`
  - If logic touches browser APIs, treat it as platform-coupled first and utility-shaped second.
  - Move that logic to an external logic boundary instead of presenting it as an innocent shared helper.
- `Meaning test`
  - If the logic expresses feature policy or domain meaning, keep it feature-local or move it to a more explicit boundary.
  - Do not globalize code just because it is pure.

### External Logic Boundaries

- `Library integration`
  - Treat third-party library interaction as an integration boundary with its own contract.
  - Prefer wrapping library-specific details behind feature-local adapters, helper modules, or adapter hooks when that improves ownership clarity.
  - Do not let raw library usage leak everywhere by default.
- `Browser and DOM logic`
  - Treat URL, DOM measurement, observers, storage, clipboard, focus management, media queries, and browser APIs as platform-coupled logic.
  - Keep browser logic near the synchronization or adapter boundary that owns it.
  - Do not disguise DOM or browser coupling as innocent shared utilities.
- `Platform adapter mindset`
  - When logic depends on a specific runtime capability, describe it as an adapter or integration boundary before calling it a utility.
  - Keep the React-facing contract small and purpose-specific.
- `Anti-patterns`
  - scattering raw browser API usage across many components
  - letting third-party library details shape component APIs directly
  - hiding platform coupling inside vaguely named helpers or hooks

### State Ownership

- `State categories`
  - `Local UI state`
    - belongs to one component or one narrow subtree
    - should stay local unless several owners truly need the same source of truth
  - `Shared client state`
    - belongs to several distant consumers that need one durable client-side owner
    - should be introduced only when local state or scoped context no longer explains ownership cleanly
  - `Server state`
    - originates from the server and follows fetch, cache, invalidate, and refetch semantics
    - should not be treated as ordinary client-owned state
  - `Derived state`
    - can be computed from existing state and inputs
    - should not become a second source of truth without strong reason

- `Preferred patterns`
  - Keep the distinction between server state, local UI state, shared client state, and derived state explicit.
  - Introduce broader client-owned state only when several owners truly need one durable source of truth.
  - Keep state placement focused on ownership clarity rather than access convenience.
- `When broader shared state is justified`
  - when several distant consumers need the same client-owned source of truth across subtree boundaries
  - when route or subtree replacement should not destroy the current client-owned value
  - when local state or scoped context would obscure the real owner rather than clarify it
- `When broader shared state is a mistake`
  - when the value is still local UI state
  - when the value is server state being mirrored unnecessarily
  - when the value is derived from existing state and can be recomputed
  - when broad access is the only reason for centralization
- `Anti-patterns`
  - putting local transient UI state into broader shared state by default
  - mirroring server state into client-owned state without a true ownership reason
  - storing pure derived values that can be computed from existing state
  - letting one broad owner absorb product orchestration that should stay closer to the feature boundary

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
- `External logic`
  - Is this library or browser dependency isolated behind a clear boundary, or leaking into general component logic?
  - Is this really a utility, or is it a platform or library adapter with ownership implications?
- `State ownership`
  - Does this value need a broader durable owner, or is it only local UI state?
  - Is this server state, shared client state, or derived state pretending to be source of truth?
- `Utilities`
  - Is this logic a React utility, pure utility, style utility, or feature-local utility?
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

## References

- Read [references/react-18-19-breaking-changes.md](references/react-18-19-breaking-changes.md) when a React architecture task also involves version upgrade risk, deprecated APIs, or behavior changes between React 18 and React 19.
