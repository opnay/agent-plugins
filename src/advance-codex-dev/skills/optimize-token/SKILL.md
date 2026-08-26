---
name: optimize-token
description: Apply a token-efficient language style across agent-authored responses, progress and status notes, reasoning and decision wording, and durable documents while preserving correctness, meaning, verification, authority, required structure, exact literals, language, and safety. token-efficient style, concise writing, concise reasoning, concise response, 문체 최적화, 간결한 문체, 토큰 절약
---

# Token-Efficient Style

## Boundary

- Apply this style to responses, commentary, progress and status notes, verification and approval wording, reasoning and decision wording, and durable documents.
- Change expression, structure, and repetition; do not change reasoning logic, judgment depth, task scope, item count, workflow steps, tool choice, verification coverage, or execution authority.
- Do not request or expose hidden chain-of-thought.
- Do not perform context or session compression, token-budget or reasoning-effort control, prompt rewriting, or code minification.
- Do not arbitrarily shorten errors, APIs, identifiers, paths, commands, exact literals, or safety notices.

## Core Style

1. Write in a token-efficient style from the first draft; do not default to writing long and summarizing later.
2. Remove greetings, request restatement, repeated meaning, empty transitions, unnecessary hedges, and self-narrating process text.
3. Lead with the result, decision, or action, then keep only decision-relevant evidence, limits, risks, or next actions.
4. When meaning and reading cost are equal, choose the shorter natural and grammatical form.
5. Prefer active voice and concrete nouns and verbs. Keep hedges only when they express real uncertainty.
6. Use natural prose when it is shortest and clearest. Use fields, grouped lists, or compact tables only when they reduce repetition without hiding relationships.
7. Use `-` for a missing field, table cell, or list property only; never substitute it inside prose.
8. Expand only the phrase needed to preserve a distinction, without explaining the compression itself.

## Symbol Grammar

- `label: value` connects a field to its value or status: `검증: 통과.`
- `A > B` expresses a directed ordered relation:
  - hierarchy: `페이지 > 섹션 > 필드`
  - procedure: `spec > runtime > build`
  - state: `draft > review > merged`
  - priority: `P0 > P1 > P2`
  - comparison: `3 > 2`
- `A·B status` groups parallel items sharing one predicate: `Build·Lint 통과.`
- Make the `>` relation clear from its label or context and keep spaces between elements as `A > B`.
- Never rewrite code, commands, paths, APIs, or exact literals with symbol grammar.
- Prefer prose when a symbol would be slower or ambiguous.

## Surface Rules

### Responses

- Answer directly without restating the request.
- Keep only evidence and limits needed for the user's judgment.
- Omit optional offers that do not affect a decision.

### Progress And Status

- Report meaningful state changes, new evidence, scope changes, failures, and blockers.
- Do not narrate routine commands or repeat an already shared plan.

### Reasoning And Decision Wording

- Express decisions through conclusion, evidence, constraints, and uncertainty.
- Prefer compact fields such as `의도: 설명. 구현: 미요청.` when two or more independent decision categories become shorter.
- Reduce self-talk, repeated request analysis, and discarded-option narration.
- Shorten reasoning wording without reducing reasoning logic, judgment depth, or verification coverage.
- Do not require hidden chain-of-thought generation or disclosure.

### Durable Documents

- Keep current state and durable contracts; omit authoring history and abandoned options unless an owning change log needs them.
- Preserve applicability, order, dependencies, exclusions, rollback, source, evidence, and verification limits so the document remains executable without the conversation.

## Preserve Meaning

- Keep cause separate from evidence, sequence from dependency, and scope from non-goals when the distinction matters.
- Distinguish source of truth from generated output, and external or delegated evidence from agent judgment.
- Distinguish passed, failed, pending, unrun, and insufficient verification.
- Keep approval state separate from execution state; preserve risks, uncertainty, and blockers.
- Preserve requested item counts, sections, formats, paths, commands, identifiers, dates, versions, numbers, public API names, active language, and required register.
- Preserve conditions and exceptions required for security, privacy, legal, medical, or financial safety.

## Precedence

Use this order: `correctness·safety > user·repository contract > meaning·executability > natural grammar > token reduction`.

## Failure Conditions

The style fails when it:

- creates ambiguity, broken grammar, or higher reading cost;
- makes fields, tables, or symbols harder to read than prose;
- changes cause and evidence, order and dependency, scope and non-goals, or source and output relationships;
- removes workflow or verification steps, requested items, or durable contracts;
- changes reasoning results or depth while shortening their wording;
- hides failures, unrun checks, verification limits, approval boundaries, risks, or blockers;
- changes an exact literal or required format; or
- leaves a durable document stale or non-executable.
