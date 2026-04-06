---
name: frontend-domain-modeling
description: Model frontend domains and feature boundaries. Use when a task involves DDD-style frontend structure, ubiquitous language, entity or value modeling, separating domain, application, and presentation logic, defining feature or slice boundaries, or deciding where business rules should live in UI-heavy codebases.
---

# Frontend Domain Modeling

## Overview

Use this skill when the real problem is domain shape, not just UI structure.
The goal is to keep business concepts explicit in the frontend so features do not collapse into ad hoc state and handler code.

## Workflow

1. Identify the main domain concepts, terms, rules, and invariants.
2. Separate domain logic from application orchestration and presentation concerns.
3. Define feature boundaries and ownership of domain concepts.
4. Decide what should be modeled explicitly and what should remain simple data.
5. End with a structure that protects domain language without overengineering.

## Output Contract

- `Domain concepts`
- `Business rules`
- `Layer boundaries`
- `Feature ownership`
- `Recommended structure`
- `Overengineering risk`

## Guardrails

- Do not import backend-style DDD complexity blindly into the frontend.
- Model only the domain concepts that materially improve clarity or correctness.
- Keep the distinction between domain logic and view-specific shaping explicit.
- Prefer language the product team would actually recognize.

## References

- Read [references/layering-rules.md](references/layering-rules.md) for boundary guidance.
