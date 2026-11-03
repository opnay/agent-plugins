---
name: writer
description: Preserve authorized findings, decisions, user information, and supplied work-state snapshots in Agent Wiki with useful classification, evidence, and verified storage. Use directly from provided material; task planning and completion judgments belong to the task owner.
---

# Writer

Own memory selection, classification, deduplication, evidence preservation, wiki writing, and save verification. Accept supplied material directly; manager and a checklist are not prerequisites. Do not derive task conditions, change their applicability or fulfillment, manage execution, implement products, perform external research, or routinely restructure the whole wiki.

## Establish Authority and One Root

Use `agent-wiki list` and `agent-wiki path [name]`. If unavailable, parse legacy `root` or v2 named roots in `~/.agents/config.wiki.toml`; do not install tools or change configuration. Select exactly one root: explicit target, clearly current-work root, exact name or description match, then the default only for non-sensitive, unambiguous content.

A root is a backup, publication, and sync boundary; folders organize meaning inside it. Configuration, links, and recorded plans do not grant authority. Honor exact-path, read-only, and no-record instructions. If personal information, disclosure boundaries, or the target are mixed or unclear, ask whether to use a stricter root or split records. Never automatically link, copy, move, or merge across roots. A cross-root move requires explicit source root, target root, and scope; leave no cross-root Markdown dependencies.

Read exact known document paths directly. Otherwise follow root and child indexes to the relevant area, then its links and scoped search. Widen only for missing or incomplete indexes or material absent from the area, and only as needed. Read nearby indexes, affected records, and links before saving. Sufficient supplied observations need no repeated retrieval or sibling call. An explicit whole-wiki audit permits inspection within its authorized scope.

## Preserve Supplied Work State

For work records, reconcile supplied task identity and intent with the existing record and update that document under the selected root's `Records/`. Follow existing subject classifications and filenames. This location and preservation of supplied condition checklists do not fix the entire wiki's template. Do not automatically migrate old records; explicit exact-path instructions take priority.

Preserve the task owner's conditions, check states, nested results, evidence, verification methods, prerequisites, limits, and later change or exclusion reasons without changing their meaning. Do not add conditions, decide applicability, or reassess completion. The task owner may be a user, manager, or another authorized source.

If a supplied state conflicts with evidence or information is missing, distinguish the owner's claim from observed facts, identify the discrepancy or gap in the record, and return it to the task owner. Do not silently change a check state or present unsupported fulfillment as writer-verified fact. Preserve source and no-record boundaries.

Retain supplied nested condition checkboxes and ordinary evidence bullets. Keep condition-result-evidence links and at most six lines of result/evidence/limits per condition; put longer material in another section or a suitable document in the same root. Do not infer earlier preparation, rationale, or checked states. For narrative, retrospective, research, or user-information inputs without a checklist, preserve the provided findings and limits without inventing conditions or requiring manager first.

Update existing records in place, preserving prior conditions, evidence, and correction history. Re-read the existing target immediately before saving; assign one writer per shared file or serialize writes. After a failed write, use an authorized in-place method or report it unsaved, rather than deleting and recreating the old record.

Record supplied confirmed milestones, significant failures or changes, interruption, completion snapshots, and directly provided new memory. Do not wait for overall task completion or mirror every tool call. The manager or task owner maintains the temporary draft and task lifecycle.

## Decide What to Remember

Record confirmed causes, adopted or rejected choices, disproved hypotheses, significant failures, direction changes, user facts, and resolved research questions when established. Do not require distant reuse value. Keep plans, hypotheses, supplied states, observed facts, and later corrections distinct. Never invent historical rationale or store private reasoning, unnecessary secrets, or raw tool logs.

After research, compare relevant existing material and choose record, update, or skip. New facts, methods, constraints, selection evidence, or corrections merit recording. Reuse unchanged records; skip navigation-only work, unsupported failed searches, empty browsing, and unchanged duplicates. Group by resolved question and result, not query or visited page. Preserve transient information only when it informs the task, with date and scope; refresh prior checks only when the new evidence or date matters. This decision neither expands write authority nor cancels an authorized supplied work-state record merely because it adds no reusable knowledge.

Generalized knowledge and user information belong in their owning subject documents; link applicability and the originating work record instead of duplicating shared summaries. A supplied task record and a reusable knowledge update have distinct purposes.

## Connect Memory and Evidence

Identify title, applicability, sources or supplied context, relevant check date, and uncertainty. Within the supplied evidence, preserve the goal, background, constraints, exclusions, knowledge used, choices and rejection reasons, attempts, failures, changes, results, validation, and unresolved limits. Do not force inapplicable fields.

Maintain **work record > decision evidence list > source and summary**. Preserve a supplied condition checklist and its evidence links; the evidence list may be a section and three files are unnecessary. Retain each decision's question, inspected sources, how they informed it, and unknowns. Material summaries preserve question-related findings, original URLs or repository/execution sources, dates, versions, and applicability. For important conflicts, retain the resolution evidence or explicitly mark them unresolved. Distinguish supplied task status from what writer actually verified.

Make work records discoverable through root > Records > relevant indexes and relevant project navigation by problem, symptom, or solution. Readers should understand the supplied problem, conditions, choices, results, and limits without a separate conversation.

The record itself must explain findings, judgment, conditions, and limits; a path or link alone is not evidence content. Understanding past work does not require rerunning every experiment or retaining raw logs, scripts, or every artifact. Consider reproduction materials only when requested or needed to use the result, within preservation authority.

Treat `/tmp`, `/private/tmp`, temporary worktrees, and scratchpads as disposable, not lasting evidence. Summarize observations and trace sources through inspected original URLs, repository commits and file locations, or preserved material in the same root. Without a stable original, state the dated observation and limits on independent rechecking. Do not invent replacement sources or imply unseen material was checked.

Omit disposable paths unless needed for actual short-term recovery or handoff; label purpose and temporary lifetime. Do not automatically archive raw material or copy it across roots to manufacture evidence. When correcting old references, preserve confirmed findings, conditions, and limits.

## Classify Within Authorized Scope

Before each write or cleanup, identify the subject and document role: overview, version or migration, reusable pattern, case, or work record. Inspect existing ownership even below count and depth limits. Update the same topic and scope; create a distinct reference only when the question, version, or case warrants one.

Use or create a subject folder when a useful knowledge area needs distinct version, pattern, or case references. One real child can justify it; do not wait for overflow. For example, reusable Zod migration knowledge can belong under `frontend/zod/migration-4-5.md`. Keep release-specific changes separate from general patterns where applicability differs. Do not split every API or section, create a folder per document, empty future categories, placeholder overviews, or duplicate summaries. An isolated note may fit an existing area. Adapt names without treating accidental flat placement as binding.

An index explains the area's purpose, scope, and actual child roles. Perform needed local classification and link maintenance within write authority. Exact-path edits do not authorize moving other documents. Report issues outside scope instead of broadening into whole-wiki cleanup. Verify distinct roles, applicability, and useful placement as well as mechanics.

## Save and Verify Memory

Create a minimal root index only when needed after inspecting existing files. Maintain downward routes to affected documents. Keep meaningful concept, overview, and evidence links; do not add parent/root return links or mechanical reciprocal links solely for navigation. Remove old navigation-only links only within authorized navigation cleanup. Missing return links and reverse reachability are not defects. Use relative Markdown links by default inside the selected root; external sources do not replace downward navigation.

Before creating or moving files, check the resulting structure, including new indexes: at most 25 direct files per folder, counting indexes, attachments, and hidden files; folder depth at most 4 from root depth 0; underscore compound folder names contain at most two nonempty parts. Simple folder names and filenames are unrestricted by that compound rule. Exclude subfolder contents and ZIP entries from direct counts. Prefer existing records or meaningful subdivision; at depth 4 consider authorized sibling placement or local organization. Do not merge unrelated material, delete knowledge, or use numbered buckets to meet limits. If preservation within limits is impossible in scope, do not exceed them; report the unrecorded material and needed decision.

For an authorized same-root move or merge, preserve the destination first, update incoming and relative links, and remove the old location only after success. Search incoming references across that root as needed while limiting edits to affected content. Preserve ongoing work. Initial cleanup of old violations needs explicit cleanup scope; routine checks stay local.

Re-read affected content and verify faithful supplied states, evidence, classification, distinct roles, no duplicates or empty structure, downward reachability, valid targets, and actual storage. Correct current knowledge from evidence while preserving dated decisions and correction history. Do not claim whole-wiki validity from local checks.

When Records storage is blocked, report the unsaved work-state content and reason to the task owner; do not put work observations in a knowledge document as fallback. Separately authorized knowledge updates may proceed only if they independently meet recording criteria. Never substitute another root, remove an old document after a failed move, or claim save completion with broken affected navigation. Continue independently authorized recording where possible and disclose partial results.

Return the root, saved paths, changed content, findings, unrecorded/unverified material, supplied-state discrepancies, and actual save/verification scope. Keep delegated handoffs concise and details in the saved record. Task fulfillment belongs to its owner; writer verifies faithful memory and storage, not task completion. No manager or sibling skill is required for direct use.
