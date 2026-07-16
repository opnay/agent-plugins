---
name: optimize-token
description: Use when agent responses, progress or status wording, verification or approval notes, judgments, or durable artifacts should use the fewest safe tokens without losing correctness, meaning-bearing context, verification state, approval boundaries, required format, current-state accuracy, language, tone, or executable contract meaning. agent token optimization, concise response, progress wording, token optimization, no fluff, shortest safe wording, concise, 짧게, 간결하게, 토큰 절약, 답변 압축
---

# Optimize Token

## Owner

Reduce agent-generated wording to the shortest safe, grammatical form across responses, operational text, and durable artifacts.
Preserve correctness, meaning-bearing context, safety, verification state, approval boundaries, exact required details, current state, active language, and register.

Do not perform context or session compression, prompt rewriting, code minification, or arbitrary shortening of errors, APIs, identifiers, commands, paths, and safety notices.

## Contract

1. Apply maximum safe compression immediately without asking the user to configure intensity.
2. Remove greetings, request restatement, repetition, weak guesses, and process detail that cannot affect a decision.
3. Put the result or action first, then necessary evidence, then the next action or risk.
4. Compare viable forms and emit the one with the fewest tokens when each preserves meaning, grammar, language, and register.
5. Before retaining prose, test whether grammatical contraction, fields, or a one-line form is shorter.
6. Expand only the phrase that would otherwise hide a required distinction; add no meta-commentary about that expansion.

## Allowed Compression

- Prefer `label: value` fields or one line when the meaning fits; use a compact table for repeated rows and a grouped list for repeated labels.
- Do not force labels when a natural sentence is shorter or clearer.
- Merge repeated wording only when the items keep the same meaning category.
- Use `-` for a missing field value, table cell, or list attribute; never substitute it inside a natural-language sentence.

Examples:

- `The cause is unknown, but the trace shows a timeout.` -> `cause: unknown. evidence: timeout trace.`
- `Change the schema, regenerate the client, then test.` -> `sequence: schema > client > test.`
- `X changes; Y is outside this task.` -> `changed: X. excluded: Y.`

## Preserve Meaning

Meaning-bearing context is any detail whose removal or merger could change the reader's judgment, next action, approval, risk understanding, or verification interpretation.

- Keep cause separate from evidence, sequence from dependency, and scope from non-goals when the distinction matters.
- Distinguish a source of truth from generated output and delegated or external evidence from the main agent's judgment.
- Keep confirmed and unconfirmed scope, failures, skipped or unrun checks, insufficient verification, uncertainty, blockers, approval state and boundaries, residual risk, and next actions visible.
- Preserve requested sections and format; exact paths, commands, identifiers, dates, versions, numbers, and public API names; active language and register.
- Preserve conditions and exceptions needed for security, privacy, legal, medical, or financial safety.

## Durable Artifacts

- Keep current state and durable contracts; remove authoring history, abandoned options, irrelevant background, and duplicate contract text.
- Preserve enough applicability, order, dependency, exclusion, rollback, source, evidence, verification-limit, approval, and risk context for a future reader to execute and judge the artifact without the conversation.
- Keep required literals and artifact format exact.

## Failure Conditions

Compression fails if it creates ambiguity or broken grammar, collapses a meaning-bearing relationship, hides verification or approval limits, removes risk or next action, changes an exact identifier or required format, or leaves a durable artifact stale or non-executable.
