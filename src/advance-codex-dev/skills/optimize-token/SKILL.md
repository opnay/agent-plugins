---
name: optimize-token
description: Use when the user asks for shorter, clearer, lower-token responses, or when pre-action thinking, progress updates, final reports, reviews, documents, or user-facing handoffs need tighter wording without losing correctness, safety boundaries, verification status, required format, language, or tone. concise response, progress wording, pre-action thinking, response compression, token optimization, no fluff, shorter, concise, 짧게, 간결하게, 토큰 절약, 답변 압축
---

# Optimize Token

## Owner

This skill compresses user-facing wording for responses, pre-action judgment sentences, and progress updates.
It preserves correctness, requested format, safety boundaries, verification status, exact technical details, active language, and register.

It does not own context/session compression, prompt rewriting, code minification, or shortening errors, APIs, identifiers, commands, paths, and safety notices.

## Workflow

1. Choose the active level: `light`, `standard`, or `extreme`.
2. Read `references/thinking.md` before visible progress or pre-action wording.
3. Read `references/response.md` before final or substantive user-facing responses.
4. Apply levels as `` `light` > `standard` > `extreme` ``: each level inherits the previous level and applies only its overrides.
5. Preserve required content before trimming: intent, result, evidence, next action, verification, approval boundary, required sections, paths, commands, numbers, and residual risk.
6. Step down for the affected detail when safety, verification, approval, or clarity needs more words.
7. Put result or action first, then evidence, then next action or risk.
8. Keep the shorter wording grammatical, natural, and no more ambiguous.

## Guardrails

- Preserve active language, register, honorifics, and requested format.
- Tighten required report items instead of deleting them.
- Keep failed, skipped, blocked, or insufficient verification visible.
- Do not infer `extreme` from generic shorter-response requests.
- Do not remove grammar words only to save tokens.
- Avoid external source names, external plugin names, unavailable dev paths, and character-framed style labels.
