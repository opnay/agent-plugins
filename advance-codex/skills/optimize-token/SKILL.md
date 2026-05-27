---
name: optimize-token
description: Use when the user asks for shorter, clearer, lower-token responses, or when pre-action thinking, progress updates, final reports, reviews, documents, or handoffs need tighter wording without losing correctness, safety boundaries, verification status, required format, language, or tone. concise response, progress wording, pre-action thinking, response compression, token optimization, no fluff, shorter, concise, 짧게, 간결하게, 토큰 절약, 답변 압축
---

# Optimize Token

## Owner

This skill tightens response wording, pre-action judgment sentences, and progress updates. It applies one call-level strength, `light`, `standard`, or `extreme`, across response and thinking/progress wording while preserving meaning, format, language, tone, safety boundaries, verification status, and exact technical details.

It does not own context compression, session summaries, prompt rewriting, code minification, or shortening errors, APIs, identifiers, commands, or paths.

## Workflow

1. Identify the requested or implied call strength: `light`, `standard`, or `extreme`.
2. Read `references/thinking.md` before visible work updates or pre-action judgment sentences.
3. Read `references/response.md` before final or substantive user-facing responses.
4. Preserve required content before trimming: user intent, next action, verification need, approval boundary, result, evidence, required sections, exact paths, commands, numbers, and residual risk.
5. Apply the selected strength across response and thinking/progress wording; step down only when safety or clarity requires it.
6. Put result or action first, then reason or evidence, then next action or risk.
7. Ensure the shorter wording is grammatical, natural, and no more ambiguous.

## Guardrails

- Preserve active language, register, honorifics, and requested format.
- Tighten required report items instead of deleting them.
- Keep failed, skipped, or insufficient verification visible.
- Do not use `extreme` for generic shorter-response requests.
- Do not remove grammar words only to save tokens.
- Avoid external source names, external plugin names, unavailable dev paths, and character-framed style labels.
