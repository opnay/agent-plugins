---
name: optimize-token
description: Use when agent output, progress wording, status wording, verification notes, approval wording, or other agent text should use fewer tokens without losing correctness, meaning-bearing context, safety boundaries, verification status, required format, language, tone, current-state accuracy, or durable contract meaning. agent token optimization, concise response, progress wording, token optimization, no fluff, shorter, concise, 짧게, 간결하게, 토큰 절약, 답변 압축
---

# Optimize Token

## Owner

This skill reduces token use in agent-generated wording across responses, progress updates, status notes, verification notes, approval boundaries, and other operational text.
It preserves correctness, meaning-bearing context, requested format, safety boundaries, verification status, exact technical details, active language, and register.

It does not own context/session compression, prompt rewriting, code minification, or shortening errors, APIs, identifiers, commands, paths, and safety notices.

## Workflow

1. Choose the active level: `light`, `standard`, or `extreme`.
2. Apply the level by default to every situation except `writing`.
3. Read `references/levels.md` when level choice, level behavior, or level comparison affects the output.
4. Read `references/writing.md` only when writing content that will be saved in files, docs, specs, records, or other durable artifacts.
   Manage writing separately through that reference.
5. Preserve meaning-bearing context before trimming: any cause, evidence, boundary, sequence, dependency, source, verification state, uncertainty, approval state, risk, or next-action condition that would change the reader's judgment if removed.
6. Preserve required literals and structure: requested sections, paths, commands, identifiers, dates, versions, numbers, active language, and register.
7. Step down for the affected detail when meaning, safety, verification, approval, or clarity needs more words.
8. Put result or action first, then evidence, then next action or risk.
9. Keep the shorter wording grammatical, natural, and no more ambiguous.

## Guardrails

- Preserve active language, register, honorifics, and requested format.
- Tighten required report items instead of deleting them.
- Keep failed, skipped, blocked, or insufficient verification visible.
- Keep cause separate from evidence, sequence separate from outcome, and delegated evidence separate from main judgment when that distinction changes meaning.
- Do not infer `extreme` from generic shorter-response requests.
- Do not remove grammar words only to save tokens.
- Avoid external source names, external plugin names, unavailable dev paths, and character-framed style labels.
