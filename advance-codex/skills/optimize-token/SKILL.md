---
name: optimize-token
description: Use when the user asks for shorter, clearer, lower-token responses, or when pre-action thinking, progress updates, final reports, reviews, documents, or handoffs need tighter wording without losing correctness, safety boundaries, verification status, required format, language, or tone. concise response, progress wording, pre-action thinking, response compression, token optimization, no fluff, shorter, concise, 짧게, 간결하게, 토큰 절약, 답변 압축
---

# Optimize Token

## Owner

This skill owns wording optimization for response surfaces, pre-action judgment sentences, and progress updates: shorter, clearer wording that preserves meaning, required format, language, tone, safety boundaries, verification status, and exact technical details.

It does not own context compression, session summaries, prompt rewriting, code minification, or shortening of error messages, APIs, identifiers, commands, and file paths.

## Workflow

1. Read `references/thinking.md` before planning visible work updates or pre-action judgment sentences.
2. Read `references/response.md` before final or substantive user-facing responses.
3. Identify the required content before trimming: user intent, next action, verification need, approval boundary, result, evidence, required sections, exact paths, commands, numbers, and residual risk.
4. Select the weakest compression strength that satisfies the request: `light`, `standard`, or `dense`.
5. Put the result or action first, then the reason or evidence, then the next action or risk.
6. Check that the shorter wording is still grammatical, natural, and no more ambiguous than the longer version.

## Guardrails

- Preserve the active language, register, honorifics, and user-requested output format.
- Keep repository-required report items; make them denser instead of deleting them.
- Do not hide failed, skipped, or insufficient verification.
- Do not use `dense` for a generic shorter-response request.
- Do not remove grammar words only to save tokens.
- Do not use external source names, external plugin names, unavailable development paths, or character-framed style labels.
