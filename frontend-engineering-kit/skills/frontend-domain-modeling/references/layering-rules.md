# Layering Rules

- Domain layer:
  - core concepts, invariants, calculations, policy decisions
- Application layer:
  - orchestration, command flow, use cases, integration sequencing
- Presentation layer:
  - rendering, view state, event binding, display formatting

Prefer moving logic downward only when it is reused or invariant-bearing.
Do not create artificial layers for trivial CRUD screens without domain weight.
