# Concern Classification

Use this guide when one block of frontend code could plausibly belong to more than one concern.

## Fast Classification

- Rendering:
  - what the user sees and how visible structure is composed
- Orchestration:
  - how multiple UI parts, events, or state transitions are coordinated
- Domain:
  - what the product means, permits, calculates, or guarantees
- Async:
  - how requests, mutations, synchronization, loading, and failure are managed
- Styling:
  - how emphasis, tokens, variants, and visual states are expressed

## Common Ambiguous Cases

### `if user.role === "admin"` inside render

- Domain when:
  - the real meaning is permission or policy
- Rendering when:
  - the rule is already decided elsewhere and this branch only decides what to show

Heuristic:
- If the condition is the source of truth for access meaning, treat it as domain.
- If it is only consuming a precomputed permission, treat it as rendering.

### `useEffect` that fetches data and writes local state

- Async when:
  - the main job is request lifecycle and synchronization
- Orchestration when:
  - the main job is coordinating several UI transitions or downstream actions after the data arrives

Heuristic:
- Classify by the dominant responsibility, not by the hook type.

### Formatter or mapper functions

- Domain when:
  - the transformation carries product meaning that should survive UI changes
- Rendering when:
  - the transformation exists only to support display text or display structure

### Derived flags like `isEligible`, `canSubmit`, `isDanger`

- Domain when:
  - the flag represents product policy or invariant logic
- Rendering or styling when:
  - the flag only controls display emphasis and the business decision is already made elsewhere

### `onClick` handlers

- Orchestration when:
  - the handler coordinates several effects, panels, child actions, or navigation steps
- Domain when:
  - the handler performs a business rule decision directly
- Rendering when:
  - the handler only toggles local view behavior with no wider coordination

## Dominant Responsibility Rule

One block may touch multiple concerns.
Classify it by the responsibility that would make future changes risky.

Examples:

- A render branch that embeds pricing policy:
  - dominant concern is domain
- A fetch callback that controls several panels and focus changes:
  - dominant concern is orchestration
- A variant map that only changes appearance:
  - dominant concern is styling

## Split Trigger

Consider a boundary change when:

- one concern is obscuring another concern's intent
- product rules are being hidden inside rendering code
- request lifecycle code is overwhelming a screen-level flow
- orchestration code is forcing unrelated children to share ownership

Do not split only because a file contains several concerns.
Split when their coexistence makes ownership, explanation, or change safety materially worse.
