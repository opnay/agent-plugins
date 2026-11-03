---
name: writer
description: Preserve authorized findings, decisions, user information, and supplied work-state snapshots in Agent Wiki with useful classification, evidence, and verified storage. Use directly from provided material; task planning and completion judgments belong to the task owner.
---

# Writer

Preserve supplied knowledge and work states as faithful, findable memory. Own selection, subject ownership, document boundaries, deduplication, storage, and save checks. The task owner owns conditions, progress, and fulfillment. Direct recording needs no manager or invented checklist. Do not take over external research, product implementation, tool installation, arbitrary root selection, or routine whole-wiki restructuring.

## Resolve the Root and Authority

Use `agent-wiki list` and `agent-wiki path [name]`. If the CLI is unavailable, parse `~/.agents/config.wiki.toml` with a TOML parser. Require a registered `default` and named `roots` with absolute existing directory `path` values and optional string `description` values. Validate types and root names, canonicalize paths, and reject equal or nested roots. Configuration symlinks are supported; CLI updates preserve the link and atomically replace its resolved target. Unsupported configuration needs an update, not automatic conversion, installation, or a substitute location. Report unsaved material and continue independent work.

Choose one root per record: explicit target, clearly current-work root, exact name or description match, then default for non-sensitive, unambiguous material. Roots bound disclosure, backup, and sync; internal folders classify subjects. Configuration and path references do not grant write authority. Respect read-only, no-record, and exact-path instructions.

Ask when personal information, disclosure boundaries, or root choice remain mixed or unresolved. Do not guess a stricter root or automatically split, link, copy, move, or merge across roots. Cross-root moves require explicit source, destination, and scope and must leave no cross-root Markdown dependencies.

Read exact known sources directly. Otherwise select the relevant area through root and child indexes, then follow useful links and search within that area. Broaden only for missing or incomplete indexes or absent material, as far as necessary. Inspect affected documents, nearby indexes, and relevant references before writing. Sufficient supplied evidence needs no repeated retrieval or sibling invocation. Whole-wiki audits require their own scope.

## Select Material to Preserve

Remember established causes, adopted or rejected choices, disproved hypotheses, significant failures, direction changes, user facts, and resolved research questions. Distant reuse value is not a prerequisite. Distinguish facts, observations, plans, hypotheses, supplied states, and later corrections. Never invent historical reasons or retain private reasoning, unnecessary secrets, or raw tool logs.

After research, compare existing content and choose new recording, update, or omission. Preserve new facts, methods, constraints, selection evidence, and corrections. Skip navigation alone, unsupported failed searches, empty browsing, and unchanged duplicates. Search uncertainty alone is not a finding; an established constraint or failure cause affecting the task can be. Group by resolved question and result, retaining sources, applicability, dates, and uncertainty. Keep transient facts only when decision-relevant and dated. Refresh an unchanged result only when its check date or evidence matters. Selection policy neither cancels authorized work-state recording nor expands authority.

## Classify Subjects and Document Boundaries

Identify the owning subject, applicability, and role: overview, version or migration, reusable pattern, case, or work record. Reusable concepts, methods, and cases belong in Knowledge's owning subjects; project-only designs, policies, and current state belong in the project's area. Origin in a project does not make reusable knowledge project-only. Preserve case conditions rather than turning them into universal rules. Work records follow their separate project layout below.

Organize from domain to subtopic using actual content and useful existing conventions. Do not impose a fixed depth, domain list, or repeated overview/version/pattern/case folder set. `Knowledge/agents/skills/` and `Knowledge/software/database/` are examples, not mandatory roots or categories. Keep one primary home for a body; other areas can point to it through justified index or body references.

| Content relationship | Action |
| --- | --- |
| Same question and applicability as an existing document | Update that document. |
| A separately useful question, version, case, or applicability | Check duplication and create a precisely named document. |
| Current content within one subject supports at least two independent body documents | Create or use the subject folder immediately and split the content. |
| Different owning subjects are mixed | Separate them into their respective subject areas. |

Count independent body documents supported by actual content, excluding `index.md`. One existing file plus supplied material can meet the threshold; do not wait for two files or a folder limit. Use distinct questions and applicability, not arbitrary sections or API counts, to establish independence. Do not create empty future categories, placeholder overviews, or speculative documents. A single question can stay as a precisely named file in an appropriate existing area. This creation rule does not require dismantling existing one-child folders or overriding exact-path instructions.

Folders name subjects; ordinary filenames and titles name the specific question or content. Indexes explain the subject and actual child roles. For one coordinates/layout document, use a name such as `design/tools/figma-coordinates-and-layout.md`. When actual Figma content supports two independent documents, use `design/tools/figma/coordinates-and-layout.md` and a specifically named sibling. These are examples, not files to pre-create. General list/detail split-view design belongs to its design-pattern subject, not a second Figma document merely because it was drawn in Figma.

For knowledge-discovery or missing-knowledge requests, compare supplied material and authorized sources with existing knowledge for independent subject candidates. Do not substitute body enrichment for that review. Report new documents, updates, duplicate omissions, and uninspected scope according to actual results; create nothing merely to satisfy a count.

Perform necessary local classification and reference maintenance within current authority. Flat existing placement does not settle ownership. An exact-file edit does not authorize moving other documents. Complete authorized work and report out-of-scope restructuring or root changes without turning a local task into global cleanup. Knowledge does not inherit the work-record path or template.

## Preserve Project-Owned Work Records

Match the supplied intent and identity to existing records. Each record belongs to one project in the selected root. Resolve ownership from the user's target, existing projects, and task context. Create only the project and minimal indexes needed for authorized recording. Use `general` only for work with no specific project, never as a fallback for uncertain ownership. For cross-project work, choose the primary responsible project without duplication; ask if evidence cannot resolve it.

New records use `Projects/<project>/records/YYYY-MM/000-<slug>.md`: initial recording month, three-digit sequence starting at 000 within that project and month, and a short descriptive slug. Choose the number after the existing maximum and recheck numbers and destination immediately before saving. Resolve collisions with the next number, not an overwrite. General work uses `Projects/general/records/YYYY-MM/000-<slug>.md`.

Continue the same intent in the same record, preserving its original month, number, and path across later months. Bulk renaming or relocation requires separate scope. Root `Records/` and other non-project record locations are unsupported. Report migration needs and unsaved material; without migration authority, do not update unsupported locations, move or duplicate their records, or put work observations in knowledge documents instead. An exact target conflicting with the layout remains unsaved.

Maintain downward navigation: root > project > `records/index.md` > `YYYY-MM/index.md` > record. Indexes explain scope and actual child roles; the project index need not duplicate individual record lists. The month folder is depth 4 with root at 0; add no deeper folder. Its 25-direct-file limit includes the index, attachments, and hidden files. At capacity, report unsaved material and the needed decision; do not evade limits through another month, project, or numbered bucket.

Retain supplied conditions, checkbox states, nested results and evidence, prerequisites, verification, exclusions, limits, and later corrections. Do not derive conditions or reassess fulfillment. For conflicts, distinguish the supplied claim, observation, discrepancy, and missing evidence and return them to the task owner. Never silently change states or present unsupported fulfillment as verified.

Keep nested checkboxes and ordinary evidence bullets. Summarize each condition's result, evidence, and limits within six lines; put longer explanations in a related section or useful same-root document without losing the relationship. Narrative research, user facts, and retrospective records need no invented preparation or checklist.

During long tasks, save established milestones, significant failures, changes, interruptions, and completion snapshots as supplied. Do not wait for completion or mirror every tool call. The task owner maintains the temporary draft and lifecycle. Re-read before each write, serialize shared-file updates, preserve ongoing edits, and never delete/recreate an existing record to bypass a failed update.

## Write Findings and Necessary References

Explain the question, findings, judgments, outcomes, and verification limits in the body. Identify subject, observation date or version, scope, and sources. Preserve supplied goals, constraints, knowledge used, alternatives and rejection reasons, meaningful attempts, failures, changes, and unresolved limits where relevant. Do not force irrelevant fields, invent missing reasons, or confuse the task owner's status with storage checks. Resolve important conflicts from evidence or label them unresolved.

Link directly to the document or section needed for a concept, claim, observation, choice, rejection, or applicability judgment. Let surrounding prose explain the relationship. Evidence lists and source summaries may stay inside a work record; no fixed link chain or three-file structure is required. Sufficient body summaries and source identifiers need no extra wiki links. An originating-work reference must pass the same relevance test.

Do not create edited-file inventories or related-document sections merely because files were written, read, edited, or share a project or task. Preserve useful evidence lists and downward index routes. Judge each reference direction separately; useful A-to-B does not require B-to-A.

Provide a complete change inventory only when explicitly requested or essential to file-by-file verification. Reuse an adequate diff or manifest; create a separate authorized artifact only when necessary. Do not replace an omitted body inventory with an automatic audit file. An inventory never replaces findings and judgments.

Keep original URLs or repository/execution identifiers, dates, versions, applicability, and uncertainty. Paths cannot replace findings or evidence summaries. Understanding the historical result does not require rerunning every experiment or retaining all raw logs, scripts, and artifacts. Preserve reproduction material only when requested or needed to use the result and within authority.

Treat `/tmp`, `/private/tmp`, scratchpads, and temporary worktrees as disposable. Summarize dated observations and trace inspected original sources or preserved same-root material. If no stable original exists, state the independent rechecking limit without inventing a replacement source. Mention temporary paths only for genuine short-term recovery or handoff, with purpose and lifetime. Do not archive raw material or copy across roots to manufacture evidence. Corrections to old references must preserve confirmed findings and limits.

## Save, Verify, and Report

Inspect existing files before creating a missing root index. Every index explains its purpose, scope, and actual child roles to guide the next reading choice. Use same-root relative Markdown links for downward navigation and justified body references. External sources do not replace index routes. Do not add navigation-only parent/root links or mechanical reciprocal links. Missing reverse reachability is not a defect. Remove existing navigation-only links only during authorized navigation cleanup; preserve ancestor-index links with conceptual or evidentiary value.

Before creating or moving anything, check the resulting structure, including new indexes: at most 25 direct files per folder and maximum folder depth 4 with root at 0. Count indexes, attachments, and hidden files, not subfolder contents or ZIP entries. Underscore compound folder names must have exactly two nonempty parts; simple folder names and filenames are outside that rule.

If limits prevent a required subject split, consider meaningful subdivisions, existing-content updates, or authorized sibling placement/local reorganization at depth 4. Do not merge unrelated content, delete knowledge, or use numbered buckets to meet limits. If preservation is impossible within scope and limits, leave the affected material unsaved and report the needed decision. Initial cleanup of unrelated existing violations needs explicit scope; ordinary checks remain local.

For authorized splits, renames, moves, and merges, prepare destinations first while preserving content, sources, dates, conditions, corrections, and uncertainty. Update indexes, incoming links, and relative references. For renames/moves/merges, remove the source only after destinations and affected links succeed. For splits, retain an original overview only when useful to its role, without duplicating bodies. Incoming-reference discovery may span the selected root; edits stay within affected content. Preserve ongoing edits and retain the source on failure.

Re-read saved content and verify supplied states, subject/project ownership, applicability and roles, precise names, missed splits or mixed subjects, unnecessary dispersion, empty structure, duplicates, downward reachability, valid justified references, structural limits, and actual storage. Semantic placement is not proved by link or count checks. Preserve dated decisions and corrections. Report out-of-scope issues; local checks do not establish whole-wiki validity.

Return the root, saved paths, new independent documents, updates, duplicate omissions, findings, unrecorded/unverified scope, supplied-state discrepancies, and actual save-check coverage concisely. Keep detail in the record. Claim discovery or completeness only within inspected sources. If recording fails, report unsaved material and reasons rather than substituting another root or a knowledge document for work observations. Independently authorized knowledge updates may still proceed. Do not claim completion with broken affected navigation. The task owner judges task fulfillment; writer verifies faithful memory and storage.
