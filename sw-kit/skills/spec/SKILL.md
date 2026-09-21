---
name: spec
description: Turn a user or product request into a technology-neutral software behavior contract. Use when actors, states, rules, exceptions, integration boundaries, non-goals, or acceptance criteria are unclear before engineering or implementation.
---

# Spec

Define what the software must do without choosing how to build it.

## Build the behavior contract

Separate confirmed requirements, assumptions, and open questions. Describe the actors, triggers, inputs, states, rules, outputs, and observable outcomes. Include normal flow, permissions, empty or invalid input, failure, recovery, and any state transition that changes behavior.

Name preserved behavior, external integration points, and non-goals. Write acceptance criteria as observable results, not architecture, framework, storage, or file instructions.

Do not decide product value, priority, or MVP scope; request or use that handoff when it matters. Do not make technical design choices; leave those to `$sw-kit:engineering`.

## Finish well

Produce a compact behavior contract with open decisions visible. A useful handoff lets engineering decide the system direction and lets implementation verify the result without rediscovering the intended behavior.
