---
name: spec
description: Clarify a user or product request as a lightweight software behavior contract. Use when the required behavior, important exceptions, completion signal, or unresolved product or technical decision is unclear before implementation.
---

# Spec

Clarify what must happen without deciding how to build it. Keep the result short.

- `behavior`: what the user does and what result follows.
- `exceptions`: only failures, permissions, or preserved behavior that materially change that result.
- `acceptance`: observable evidence that the behavior is complete.
- `open`: product or technical decisions that cannot be inferred.

Do not decide product value, priority, MVP scope, architecture, technology, storage, or source-level implementation. Leave product decisions with their owner and technical decisions to `$sw-kit:engineering`.

Add states, integration details, non-goals, or other structure only when the request actually needs them.
