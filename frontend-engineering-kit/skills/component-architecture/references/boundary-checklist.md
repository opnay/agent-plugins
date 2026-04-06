# Boundary Checklist

Use this checklist before recommending a split, extraction, or ownership move.

## 1. Task Coherence

- Does the component still serve one understandable user task?
- If split, would the task become easier to explain or only more distributed?

## 2. Concern Classification

- Does the component mix rendering with domain logic?
- Does it mix orchestration with leaf-level rendering?
- Does async lifecycle dominate the visible rendering path?
- Are styling variants and behavior decisions tangled together?

## 3. Responsibility Ownership

- Are rendering, domain, async, formatting, and state ownership easy to explain?
- Is any responsibility owned by a component that lacks enough context to justify that ownership?
- Are siblings depending on state or handlers that live too low in the tree?
- Is a parent holding responsibilities that should remain local?

## 4. State And Effect Placement

- Are effects colocated with the state they synchronize?
- Is shared state lifted only as far as necessary?
- Are derived values being stored instead of recomputed?
- Is form or request-adjacent state isolated clearly enough?

## 5. Split Value

- Would a split reduce reasoning cost for future changes?
- Would the new boundary reduce real ownership conflict?
- Would the split mostly create wrapper components, prop tunneling, or thin pass-through layers?

## 6. Hook Extraction

- Would extracting a hook improve ownership clarity?
- Would it isolate reusable behavior or only hide complexity?
- Would the hook create a cleaner boundary or just move state and effects sideways?

## 7. Public API Quality

- Is the public API small enough to remain stable?
- Are props expressing real ownership or just compensating for a weak boundary?
- Would composition work better than adding more config props?

## 8. Reuse Reality Check

- Are reusable parts actually reused, or only hypothetically reusable?
- Would sharing this component erase domain meaning that should stay local?

## Decision Rule

Recommend a boundary change only when it makes ownership clearer, future changes safer, or the task easier to explain.
Do not recommend a boundary change when the main result is indirection, abstraction theater, or responsibility laundering.
