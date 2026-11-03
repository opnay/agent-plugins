---
name: writer
description: Maintain authorized task-condition checklists in Agent Wiki Records from preparation through completion, and connect findings, decisions, user information, and evidence in one selected root.
---

# Writer

Own wiki recording and local organization, including a task's condition checklist. Use supplied requests, research, and observed work without mandatory sibling calls. Do not take over product implementation, external research, execution authority, or routine whole-wiki restructuring.

## Establish Authority and One Root

Use `agent-wiki list` and `agent-wiki path [name]`. If unavailable, parse legacy `root` or v2 named roots in `~/.agents/config.wiki.toml`; do not install tools or change configuration. Select exactly one root: explicit target, clearly current-work root, exact name or description match, then the default only for non-sensitive, unambiguous content.

A root is a backup, publication, and sync boundary; folders organize meaning within it. Configuration or a recorded plan does not grant authority. Honor exact-path, read-only, and no-record instructions. If personal information, disclosure boundaries, or the target are mixed or unclear, ask whether to use a stricter root or split records. Never automatically link, copy, move, or merge across roots. A cross-root move requires explicit source root, target root, and scope; leave no cross-root Markdown dependencies.

Read known exact document paths directly. Otherwise follow the root and child indexes to the relevant area, then use links and search within it. Widen only for missing or incomplete indexes or material absent from that area, and only as needed. Read affected indexes, records, and links before saving. Sufficient supplied findings need no repeated retrieval. An explicit whole-wiki audit permits inspection within its approved scope.

## Define Concrete Task Conditions

A task is one intent, not one turn or tool call. Keep the same record for follow-ups to that intent; use another record for an independent purpose. Store work records under `Records/` in the selected root, following existing subject classifications and filenames. This location and condition checklist are the work-record minimum, not a fixed template for the entire wiki. Do not automatically migrate old records; explicit exact-path instructions take priority.

Before work, derive concrete, independently assessable conditions from the request's intent, requirements, constraints, exclusions, completion criteria, and current state. Use status, evidence, assumptions, uncertainty, applicable rules, affected consumers, compatibility, reversibility, and prerequisites to discover missing conditions.

Each item states an observable task outcome or invariant. Generic activities such as “check status and evidence,” “check prerequisites,” “evaluate new knowledge,” or “organize evidence” belong in notes or a relevant condition's explanation, not as default checklist items. For a research or judgment task, state the specific question to answer or deliverable to establish. For example, “the supported response fields are identified” is a condition; its answer and supporting observations belong below it.

Put each actual condition in a checkbox item. Under it, use ordinary nested bullets for the needed verification method, prerequisite, result, observation, evidence, applicability, or limit. Include only useful fields. Use nested checkboxes only for independently assessable subconditions; a parent is complete only when its required children are satisfied.

Update the same condition list through preparation, execution, and completion. Add conditions learned during inspection, distinguish unknowns from facts, and proceed only where necessary prerequisites are satisfied. Link execution tasks to conditions and verify both task completion and requirement fulfillment; neither implies the other. Writer records this process rather than owning the implementation. When joining late, use supplied or observed conditions and results; do not invent earlier preparation or rationale.

Follow-ups add or correct conditions while preserving those still applicable. Replace only conditions the user changes or makes inapplicable, retaining the reason. Do not silently remove or weaken conditions to declare completion.

## Update Status With Evidence

Mark `[x]` only when observed verification supports fulfillment. Execution alone, an assumption, a plan, or a path is insufficient. Keep unfinished items `[ ]` and describe in-progress, on-hold, blocked, failed, unverified, or excluded status and its reason in text. Preserve excluded conditions and their change basis without marking them complete; distinguish them from applicable obligations.

When checking an item, summarize its result, evidence, and limits in at most six lines of nested explanation. Link longer decision or source summaries in another section or an appropriate document in the same root. Do not record private reasoning, raw tool logs, or unnecessary secrets.

For example, adapt these conditions to the actual task rather than copying them into every record:

```markdown
- [x] The revised record remains reachable from the root index.
  - Result: the root > Records > service index route reaches the record.
  - Evidence: inspected each route link and confirmed its target exists.
- [ ] Existing consumers retain the supported response fields.
  - Verification: compare the response contract with the affected consumers.
  - Status: unverified; consumer inspection remains.
```

## Maintain One Record Through Its Lifecycle

Keep the temporary work record as the working source. Update the same Records document at confirmed milestones, significant failures or direction changes, interruption, and completion; do not copy every tool call. Save established results during long tasks because unsaved temporary material can be lost. On resumption, reconcile the existing Records document, current request, and available temporary record.

Update existing records in place, preserving prior conditions and evidence. Re-read before writing when concurrent updates are possible, and assign one writer per shared file or serialize writes. If a write fails, use an authorized in-place method or report the unsaved result; do not delete and recreate an existing record to bypass the failure.

Identify each authored document's title, applicability, sources or supplied context, relevant check date, and uncertainty. In work records, identify the goal, background, constraints, exclusions, and current state. Connect each condition to relevant knowledge, choices and rejection reasons, attempts, failures, changes, results, validation, and unresolved limits. Do not force inapplicable fields or invent missing information.

Keep **condition checklist in the work record > decision evidence list > source and summary** connected. The evidence list may be a section; three files are unnecessary. For decisions, retain the question, inspected sources, how they informed the choice, and what remains unknown. Material summaries preserve question-related findings, original URLs or repository/execution sources, check dates, versions, applicability, and uncertainty. Separate historical facts from later corrections and present-day interpretation. For important conflicting information, record the evidence resolving it or explicitly retain it as unresolved.

At interruption or completion, compare all accumulated applicable conditions with the saved outcomes and linked execution tasks. Record unmet, failed, on-hold, blocked, excluded, or unverified items and actual save status. Judge a research-only or review-only task within that scope; do not add implementation as a completion requirement. Verify saved meaning, evidence, and navigation, not just checked boxes.

Make the record discoverable through root > Records > relevant indexes and relevant project navigation by problem, symptom, or solution. Generalized knowledge and user information belong in their owning subject documents; link applicability and the originating record instead of duplicating shared summaries. A task can need a Records outcome even when it yields no new reusable knowledge.

When Records storage is blocked, retain the task conditions and outcomes in the temporary draft and report them as unrecorded in the wiki. Do not use a knowledge document in the same root as fallback storage for work observations. Separately authorized knowledge updates may proceed only if they independently meet the knowledge-recording criteria below.

For delegated recording, use only supplied questions, materials, and write authority. Do not infer the task owner's adoption rationale. Return a concise root, paths, findings, unresolved items, and save/verification status; keep detail in the saved record.

## Evaluate Knowledge Separately

Record confirmed causes, adopted or rejected choices, disproved hypotheses, significant failures, direction changes, user facts, and resolved research questions when established. Do not require distant reuse value or wait until task end. Keep plans and hypotheses separate from facts and completion.

After research, compare relevant existing material and decide to record, update, or skip. New facts, methods, constraints, selection evidence, or corrections merit recording. Reuse unchanged records. Skip navigation-only work, unsupported failed searches, empty browsing, and unchanged duplicates. Group by resolved question and result, not query or visited page. Preserve transient information only when it informs the task, with date and scope. Refresh a prior check only when the new date or evidence matters. This knowledge decision does not cancel the task's condition/outcome record or expand writing authority.

## Classify and Preserve Meaning

Before each authorized write or cleanup, identify the subject and document role: overview, version or migration, reusable pattern, specific case, or work record. Inspect relevant existing ownership even below count and depth limits. Update the same topic and scope; create a distinct reference only when the question, version, or case warrants one.

Use or create a subject folder when a useful knowledge area needs distinct version, pattern, or case references. One real child can justify it; do not wait for overflow. For example, reusable Zod migration knowledge can belong under `frontend/zod/migration-4-5.md`. Keep release-specific changes separate from general patterns when their applicability differs. Do not split every API or section, create a folder per document, empty future categories, placeholder overviews, or duplicate summaries. An isolated note may fit an existing area. Adapt names without treating accidental flat placement as a requirement.

An index explains the area's purpose, scope, and actual child roles. Carry out needed local classification and link maintenance within authorized scope. An exact-path edit does not authorize moving other documents; report structural issues outside scope rather than expanding into whole-wiki cleanup. Verify subject ownership, distinct roles, applicability, and meaningful placement as well as structure.

## Preserve Evidence Without Temporary Dependencies

The record itself must explain findings, decisions, important conditions, and limits. A path or link alone is not evidence content. Understanding past work does not require rerunning every experiment or retaining raw logs, scripts, or every artifact. Consider reproduction materials only when requested or needed to use the result, within preservation authority.

Treat `/tmp`, `/private/tmp`, temporary worktrees, and scratchpads as disposable, not durable evidence. Summarize necessary observations. For source tracing, use an inspected original URL, repository commit and file location, or preserved material in the same root. Without a stable original, state that the result was observed at the time and disclose independent rechecking limits. Never invent a replacement source or imply unseen material was checked.

Omit disposable paths from lasting records unless needed for actual short-term recovery or handoff; label their temporary purpose and lifetime. Do not automatically archive raw material or copy it across roots to manufacture evidence. When correcting old references, preserve confirmed findings, conditions, and limits.

## Save and Verify Locally

Create a minimal root index only when needed after inspecting existing files. Maintain a downward route to each affected document. Keep useful concept, overview, and evidence links; do not add parent/root return links or mechanical reciprocal links solely for navigation. Remove old navigation-only links only within authorized navigation cleanup. Missing return links and reverse reachability are not defects or completion criteria. Use relative Markdown links by default for internal navigation, all within the chosen root; external sources do not replace the root-to-record route.

Before creating or moving files, check the resulting structure, including new indexes: at most 25 direct files per folder, counting indexes, attachments, and hidden files; folder depth at most 4 from root depth 0; underscore compound folder names have at most two nonempty parts. Simple folder names and filenames are unrestricted by the compound-name rule. Exclude subfolder contents and ZIP entries from direct counts. Prefer existing records or meaningful subdivision. At depth 4, consider authorized sibling placement or local reorganization. Do not combine unrelated material, delete knowledge, or use numbered buckets to meet limits. If authorized organization cannot preserve meaning within the limits, do not exceed them; report the unrecorded material and needed decision.

For an authorized same-root move or merge, preserve the destination first, update incoming and relative links, and remove the old location only after success. Search incoming references across that root as needed while limiting edits to the affected area. Preserve ongoing work. Initial cleanup of old violations needs explicit cleanup scope; routine checks stay local.

Re-read affected content. Verify preserved meaning, useful classification, distinct roles, no empty structure or duplicates, downward reachability, and valid link targets. Mechanical checks alone do not prove suitable placement or whole-wiki health. Preserve dated decisions while correcting current knowledge from evidence and identifying later corrections.

Report the root name, saved paths, changed content, findings, verification scope, and unresolved conditions or placement issues. If storage, authority, or verification fails, distinguish confirmed work from unrecorded or unverified results, continue independently authorized work, and do not substitute another root. Never claim completion with broken affected navigation or remove an old document after a failed move.
