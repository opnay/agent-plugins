# Hook Extraction Rules

Use this guide only after:

1. naming the responsibility candidate
2. classifying its dominant concern
3. deciding its likely layer target

## Good Hook Targets

- state transition behavior
- effect synchronization behavior
- reusable interaction behavior
- UI-facing async adapter behavior

These are good hook targets because they often depend on React lifecycle and can still be explained as one behavioral contract.

## Bad Hook Targets

- pure calculations
- business policy decisions
- transport or server-contract handling
- feature-wide orchestration across multiple boundaries
- one-off render-local behavior that loses clarity when extracted

## Hook Suitability Test

Ask:

- Does this responsibility need React state, refs, effects, or lifecycle coupling?
- Does extraction reduce reasoning cost at the call site?
- Does the extracted API have a stable meaning?
- Would another target be more honest about the responsibility?

If the answers are weak, do not use a hook.

## Reuse Test

Reusable does not mean "similar code exists elsewhere."

Reusable means:

- the same behavior contract repeats
- the same reason to change repeats
- the same ownership shape still makes sense after extraction

If reuse depends on renaming everything and ignoring context, it is not real reuse.

## Alternative Targets

- Utility:
  - for pure helpers and framework-agnostic calculations
- Domain:
  - for product meaning, rules, permissions, calculations, validation policies
- API:
  - for transport and server-contract concerns
- Feature:
  - for multi-boundary flow coordination
- Component-local:
  - for logic whose meaning depends on one component's view boundary

## Smells

- the hook mirrors component props and returns nearly the same shape
- the hook is only reusable by components with the same markup assumptions
- the hook becomes a bag of handlers, effects, and state with no coherent contract
- the hook hides a better layer decision
