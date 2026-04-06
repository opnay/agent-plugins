# Layer Placement Rules

Use this guide after extracting a responsibility candidate.

## Placement Targets

- Utility:
  - framework-agnostic and business-agnostic logic
  - portable helpers that could become a tiny library without product context
- API:
  - transport, endpoints, requests, responses, DTOs, contract interpretation
- Domain:
  - product meaning, invariants, policies, calculations, validation rules, permissions
- Feature:
  - feature-specific flow coordination, user journey orchestration, cross-boundary interaction logic
- Component-local:
  - rendering structure, local interaction behavior, tightly scoped display concerns

## Decision Path

1. Name the candidate as one behavior, decision, sync rule, or presentation responsibility.
2. Ask what causes that responsibility to change.
3. Place it where that reason to change is most native.

## Quick Heuristics

- If it can survive outside this product, it may be utility.
- If it is about server contract or transport details, it belongs near api.
- If it expresses product policy or meaning, it belongs near domain.
- If it coordinates a real user flow across boundaries, it belongs near feature.
- If it only matters inside this component's rendering boundary, keep it local.

## Reuse Rule

Do not call something reusable because two code blocks look similar.

Reusable means:

- the same responsibility repeats
- the same reason to change applies
- the ownership stays stable after extraction
- the extracted API can be explained briefly

If those are weak, keep the logic closer to its current boundary.

## Failure Modes

- Product logic hidden in utility
- Server contract parsing hidden in UI code
- Feature orchestration pushed into domain
- Component-local behavior promoted too early into shared layers
- Extraction that increases indirection without clarifying ownership
