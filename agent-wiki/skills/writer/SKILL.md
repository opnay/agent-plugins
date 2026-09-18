---
name: writer
description: Record authorized Agent Wiki facts, decisions, evidence, and outcomes in one explicitly selected named root while preserving backup and publication boundaries.
---

# Writer

Write authorized, traceable wiki records and local links. Do not perform external research, product implementation, or choose a different wiki implicitly.

## Select One Root First

Use `agent-wiki list` and `agent-wiki path [name]`. If the CLI is unavailable, read legacy `root` or v2 named roots from `~/.agents/config.wiki.toml` with a TOML parser; do not change it. A root is a physical backup, publication, and sync boundary; folders inside it are semantic organization. Select exactly one root in this order:

1. An explicit user target.
2. A clearly current-work root.
3. An exact root-name or description match.
4. The default only for non-sensitive, unambiguous content.

If personal information, work history, or different disclosure boundaries are mixed or the target is unclear, ask the user whether to use a stricter root or split records. Do not automatically link, copy, move, or merge across roots. A cross-root move requires the user to explicitly name the source root, target root, and scope; never create a cross-root Markdown dependency. Configuration never grants writing authority.

## Record Meaningful Events

Within the selected root, record confirmed causes, adopted or rejected options, disproved hypotheses, significant failures, direction changes, validated results, user facts, and resolved research questions. Reuse unchanged records. Keep hypotheses distinct from facts and omit private reasoning, raw logs, and unnecessary secrets.

Connect **work record > decision evidence list > source and summary**. Preserve original URLs or file/execution locations, check date, version, scope, uncertainty, user-provided context, and the actual reason evidence informed a decision. Do not invent historical rationale.

After research, record new facts, constraints, selection evidence, or corrections; skip navigation-only work, unsupported failed searches, and unchanged duplicates. Group by resolved question and result, not query or page visit.

## Write and Verify

Read nearby indexes and documents before editing. Create a minimal `index.md` only in the selected root when needed. Keep a real route from that root to every changed document and back. Markdown links must stay inside that root.

Before saving, check the affected root: 25 direct files per folder, maximum depth 4 from root depth 0, and two-part underscore compound folder names. Preserve incoming and relative links when moving or merging within that root. Do not expand an affected-area repair into a whole-wiki audit.

Report saved paths with root names, recorded content, verification scope, and unresolved issues. If authority, storage, or verification fails, report the partial or unrecorded state plainly.
