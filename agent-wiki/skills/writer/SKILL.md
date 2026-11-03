---
name: writer
description: Record authorized Agent Wiki facts, decisions, evidence, and outcomes in one explicitly selected named root while preserving backup and publication boundaries.
---

# Writer

Maintain authorized wiki knowledge, decisions, user information, and work records with useful local navigation. Do not perform external research, product implementation, implicit root selection across disclosure boundaries, or routine whole-wiki restructuring.

## Select One Root

Use `agent-wiki list` and `agent-wiki path [name]`. If unavailable, parse legacy `root` or v2 named roots in `~/.agents/config.wiki.toml`; do not change configuration. Select exactly one root in this order:

1. The user's explicit target.
2. A clearly current-work root.
3. An exact root-name or description match.
4. The default only for non-sensitive, unambiguous content.

A root is a physical backup, publication, and sync boundary. Folders organize meaning inside that boundary. Configuration does not grant writing authority. Honor read-only and no-record instructions.

If personal information or different disclosure boundaries are mixed, or the target is unclear, ask whether to use a stricter root or split records. Do not automatically link, copy, move, or merge across roots. A cross-root move requires the user's explicit source root, target root, and scope. Never create cross-root Markdown dependencies.

## Read Context and Record Meaningful Results

Use sufficient supplied research, user input, and observed work directly without repeating retrieval or requiring reader or researcher calls. Read known exact document paths directly. Otherwise select an area through root and child indexes, then follow relevant links and search within that area. Expand only for missing or incomplete indexes or material not found there, and only as far as needed. Explicit whole-wiki cleanup or audit requests permit inspection within that approved scope.

Read nearby indexes, target records, and affected links before saving. Update the same topic and scope; create a document when a distinct question, version, or case needs its own reference. Use sound local conventions without treating an accidental existing placement as a requirement. Do not impose a fixed folder template.

Record confirmed causes, adopted or rejected options, disproved hypotheses, significant failures, direction changes, validated results, user facts, and resolved research questions as they become known. Do not wait until task end or require distant reuse value. Keep hypotheses and plans separate from facts and completion. Never invent historical rationale.

After research, compare with relevant existing wiki material and decide whether to record, update, or skip. Record new facts, methods, constraints, selection evidence, or corrections. Reuse unchanged records; skip navigation-only work, unsupported failed searches, empty browsing, and unchanged duplicates. Group by resolved question and result, not query or page visit. Record transient information only when it informs the task, with its date and scope.

Connect **work record > decision evidence list > source and summary**. The evidence list may be a section; three files are not required. Include the relevant goal, constraints, knowledge used, options and reasons, attempts, outcomes, validation, and unresolved limits. For each decision, retain the question, relevant source links, how the evidence informed it, and unresolved items. Material summaries retain the question-related findings. Preserve supplied source and check dates, versions, applicability, uncertainty, and user-provided context; disclose gaps instead of filling them by inference. Retain inspected original URLs or repository/source locations in material summaries, or state the missing-source limit. Omit private reasoning, raw tool logs, and unnecessary secrets.

Make work records discoverable through relevant project navigation by problem, symptom, or solution. Link generalized knowledge to its originating work record and applicability, and reuse shared summaries rather than duplicating them.

Update the same work record as meaningful events occur. At completion, compare those events with saved decisions, evidence, and results. For delegated recording, use only supplied questions, materials, and write authority; do not infer the task owner's adoption rationale. Assign one writer per shared file or serialize writes, and re-read before saving when concurrent changes are possible.

## Classify the Subject and Document Role

Before each authorized write or cleanup, identify the owning subject and the document's role: overview, version or migration, reusable pattern, specific case, or work record. Inspect relevant existing content for split or scattered ownership. Make this judgment even below file-count and depth limits; those limits do not establish a suitable location.

Use or create a subject folder when the content establishes an independently useful knowledge area and its version notes, patterns, or cases need distinct references. A concrete version document can be the first child; do not wait for a second file or an overflow. For example, a Zod migration belongs under `frontend/zod/migration-4-5.md` when Zod owns the reusable validation knowledge. Adapt names to the wiki. Keep release-specific changes separate from general patterns when their questions or applicability differ; do not split every API or heading mechanically.

Keep an isolated note in a suitable existing area when no separate subject boundary is useful. Do not create one folder per document, empty future categories, placeholder overview files, or duplicated summaries. A subject index explains its scope and existing children; one real child can be sufficient. Preserve useful structure and content rather than retaining a flat layout solely because it already exists.

Carry out needed local classification and link maintenance within the authorized writing or cleanup scope without a separate routine confirmation. An exact-path-only edit does not authorize moving other documents. If coherent placement requires changes outside that scope or across a root boundary, complete the authorized part and report the remaining structural issue. Do not turn a local write into a whole-wiki reorganization.

## Keep Evidence Useful Without Temporary Files

The record itself must explain observed facts or results, judgment, important conditions, and limits. A link or execution path alone is not evidence content. A new reader should understand the past problem and choices from the record and relevant sources.

Wiki completeness does not require rerunning every experiment or retaining raw logs, reproduction scripts, or all artifacts. Consider reproduction materials only when the user requests them or using the result requires a rerun, within the authorized preservation scope.

Treat `/tmp`, `/private/tmp`, temporary worktrees, and scratchpads as disposable work locations, not durable evidence sources. Summarize needed observations in the record. When original-source tracing matters, use an inspected original URL, repository commit and file location, or preserved material in the same root. Without a stable original, state that the result was observed at the time and disclose the limit on independent rechecking. Do not invent a replacement source or imply unavailable material was inspected.

For a lasting wiki record, omit the disposable path itself. Retain a temporary path only for an actual short-term recovery or handoff need, explicitly labeled with its purpose and temporary lifetime. Do not describe it as permanent preservation or proof of current accessibility. Do not automatically archive raw material or copy it across roots to make a citation durable. When correcting an old temporary reference, preserve confirmed findings, conditions, and limits.

## Maintain Top-Down Navigation

The folder tree supplies parent location. An index explains its area's purpose, scope, and the roles of child topics and documents so the reader can choose what to read next; do not merely append an unstructured file list. Maintain a real downward route from the root index through relevant areas to each affected document. Ordinary documents may link to precise related concepts, overviews, and evidence.

Do not add parent, root, or mechanically reciprocal links solely for returning to an index. Remove existing navigation-only links only when navigation cleanup is authorized, not as an extra during an unrelated content update. Judge a link by its label and context, not by whether it targets a parent index; preserve meaningful overview and evidence references. Keep local repairs within scope. Verify downward reachability and valid link targets; reverse reachability and return links are not completion criteria. Markdown navigation stays inside the selected root.

## Save and Verify Within Scope

Create a minimal root `index.md` only when needed, after inspecting existing files. Identify title, scope, sources or supplied context, relevant date, and uncertainty without forcing YAML or tables. Correct current knowledge from evidence while preserving dated decisions and distinguishing later corrections.

Before creating or moving files, check the resulting structure: at most 25 direct files per folder, including indexes, attachments, and hidden files; folder depth at most 4 from root depth 0; underscore compound folder names have at most two nonempty parts. Simple folder names and filenames are unrestricted by the compound-name rule. Count new indexes; do not count subfolder contents or ZIP entries as direct files. Prefer updating records or meaningful subdivision. At depth 4, consider existing sibling areas or reorganization within the authorized area before concluding the content cannot be stored. Do not merge unrelated content, delete knowledge, or use numbered buckets to satisfy limits.

For an authorized same-root move or merge, prepare the preserved destination, update incoming and relative links, then remove the old location. Search incoming references across the root as needed, but limit edits to the affected area. Do not overwrite ongoing work. Perform an initial cleanup of existing violations within explicit cleanup scope; keep routine follow-up checks local. If preservation and structure limits cannot both be met, report what remains unrecorded.

Verify meaning as well as mechanics: each affected document belongs to the chosen subject, its role and applicability are distinct, related knowledge is not needlessly scattered, and indexes explain the existing children. Confirm that no empty structure or duplicate content was introduced. Check preserved content, downward routes, and file targets. Passing link, count, and depth checks alone is insufficient. Report any placement issue left outside the authorized scope; do not claim the whole wiki is valid from a local check.

Return the root name, saved paths, changed content, key findings, verification scope, unresolved gaps, and actual save status. For delegated recording, keep this handoff concise and leave the detailed summary in the saved document. Distinguish successful research from failed storage or verification. Do not report completion with broken affected navigation, delete an old document after a failed move, or silently save to another root when authority or storage is unavailable. Continue independently authorized work and disclose partial or unrecorded results.
