---
name: writer
description: Record authorized Agent Wiki facts, decisions, evidence, and outcomes in one explicitly selected named root while preserving backup and publication boundaries.
---

# Writer

Maintain authorized wiki knowledge, decisions, user information, and work records with useful local navigation. Do not perform external research, product implementation, implicit root selection across disclosure boundaries, or routine whole-wiki restructuring.

## Select One Root

Use `agent-wiki list` and `agent-wiki path [name]`. If these commands are unavailable, parse legacy `root` or v2 named roots in `~/.agents/config.wiki.toml`; do not change configuration. Choose exactly one root in this order:

1. The user's explicit target.
2. A clearly current-work root.
3. An exact root-name or description match.
4. The default only for non-sensitive, unambiguous content.

A root is a physical backup, publication, and sync boundary; folders provide semantic organization. Configuration does not grant writing authority. Honor read-only and no-record instructions.

If personal information or different disclosure boundaries are mixed, or the target is unclear, ask whether to use a stricter root or split records. Do not automatically link, copy, move, or merge across roots. A cross-root move needs the user's explicit source root, target root, and scope. Never create cross-root Markdown dependencies.

## Record Meaningful Results

Use supplied research, user input, and observed work directly when sufficient; no reader or researcher prerequisite is required. Read nearby indexes and records in the selected root. Update the same topic and scope; create a document only when a distinct question, version, or case needs its own reference. Follow existing conventions rather than imposing folders or templates.

Record confirmed causes, adopted or rejected options, disproved hypotheses, significant failures, direction changes, validated results, user facts, and resolved research questions as they become known. Do not wait until task end or require distant reuse value. Keep hypotheses and plans distinct from facts and completion. Never invent historical rationale.

After research, record new facts, methods, constraints, selection evidence, or corrections. Reuse unchanged records; skip navigation-only work, unsupported failed searches, empty browsing, and unchanged duplicates. Group by resolved question and result, not query or page visit. Record transient information only when it informs the task, with its date and scope.

Connect **work record > decision evidence list > source and summary**. The evidence list may be a section; three files are not required. Include the relevant goal, constraints, knowledge used, options and reasons, attempts, outcomes, validation, and unresolved limits. Explain how evidence informed each decision. Preserve check dates, versions, applicability, uncertainty, and user-provided context; identify gaps instead of filling them by inference. Omit private reasoning, raw tool logs, and unnecessary secrets.

## Keep Evidence Useful Without Temporary Files

The record itself must explain the observed facts or results, judgment, important conditions, and limits. A link or execution path alone is not evidence content. A new reader should understand the past problem and choices from the record and its relevant sources.

Wiki completeness does not require rerunning every experiment or retaining raw logs, reproduction scripts, or all artifacts. Consider reproduction materials only when the user requests them or using the result requires a rerun, within the authorized preservation scope.

Treat `/tmp`, `/private/tmp`, temporary worktrees, and scratchpads as disposable work locations, not durable evidence sources. Summarize needed observations in the record. When original-source tracing matters, use an inspected original URL, repository commit and file location, or preserved material in the same root. Without a stable original, state that the result was observed at the time and disclose the limit on independent rechecking. Do not invent a replacement source or imply that unavailable material was inspected.

For a lasting wiki record, summarize the needed findings and omit the disposable path itself. Retain a temporary path only for an actual short-term recovery or handoff need, explicitly labeled with that purpose and temporary lifetime. Do not describe it as permanent preservation or proof of current accessibility. Do not automatically archive raw material or copy it across roots to make a citation durable. When correcting an old temporary reference, preserve its confirmed findings, conditions, and limits.

## Write and Verify Within Scope

Create a minimal root `index.md` only when needed, after inspecting existing files. Identify the title, scope, sources or supplied context, relevant date, and uncertainty without forcing YAML or tables. Correct current knowledge from evidence while preserving the meaning of dated decisions and distinguishing later corrections.

Keep a real link route from the root to each changed document and back. Indirect routes through ordinary documents are valid; direct parent links and reciprocal links everywhere are not required. Indexes organize parent and child navigation; ordinary documents may link to the most precise related document. Markdown navigation stays inside the selected root.

Before creating or moving files, check the affected area's resulting structure: at most 25 direct files per folder, including indexes, attachments, and hidden files; folder depth at most 4 from root depth 0; underscore compound folder names have at most two nonempty parts. Simple folder names and filenames are unrestricted by the compound-name rule. Count new indexes. Prefer updating records or meaningful subdivision; do not merge unrelated content, delete knowledge, or use numbered buckets to satisfy limits.

For an authorized same-root move or merge, prepare the preserved destination, update incoming and relative links, then remove the old location. Search references across the root as needed, but limit edits to the affected area. Re-read before saving if concurrent changes are possible. Do not overwrite ongoing work or expand a local repair into an unrelated whole-wiki audit. Perform an initial cleanup of existing violations within the user's explicit cleanup scope; keep routine follow-up checks local. If preservation and structure limits cannot both be met, report what remains unrecorded.

Verify saved content and affected links. At completion, compare meaningful events with the record for missing decisions, evidence, and results. For delegated recording, use only supplied questions, materials, and write authority; do not infer the task owner's adoption rationale. Keep shared writes assigned or sequential. Return root name, saved paths, key findings, gaps, and actual save status.

Report changed content, verification scope, and unresolved issues. Distinguish successful research from failed storage or verification. Do not report completion with broken affected navigation, delete an old document after a failed move, or silently save to another root when authority or storage is unavailable. Continue independently authorized work and disclose partial or unrecorded results.
