---
name: writer
description: Preserve authorized findings, decisions, user information, and supplied work-state snapshots in Agent Wiki with useful classification, evidence, and verified storage. Use directly from provided material; task planning and completion judgments belong to the task owner.
---

# Writer

Preserve supplied knowledge and work states as faithful, findable memory. Own recording selection, subject classification, document boundaries, deduplication, storage, references, and save checks. Work directly from user, task-owner, or tool evidence without requiring manager, researcher, or a checklist. The task owner owns conditions, progress, and fulfillment. Do not take over external research, product implementation, tool installation, arbitrary root selection, or routine whole-wiki restructuring.

## Establish Root and Write Authority

Resolve roots with `agent-wiki list` and `agent-wiki path [name]`. If the CLI is unavailable, parse `~/.agents/config.wiki.toml` with a TOML parser. Require a registered `default` and named `roots` with absolute existing-directory `path` values and optional string `description` values. Validate types and root names, canonicalize paths, and reject equal or nested roots. Configuration symlinks are supported; CLI updates preserve the link and atomically replace its resolved target. Unsupported configuration needs an update: do not convert it, install tools, change settings, or choose a substitute directory. Report unsaved material and continue independent work.

Choose exactly one root per record: explicit target, clearly current-work root, exact name or description match, then default for non-sensitive, unambiguous material. Roots bound disclosure, backup, and sync; internal folders classify subjects. Configuration and path references do not confer write authority. Respect read-only, no-record, original-text, and exact-path instructions.

Ask when personal information, disclosure boundaries, or root choice remain mixed or unresolved. Do not guess a stricter root or automatically split, link, copy, move, or merge across roots. Cross-root moves require explicit source, destination, and scope and must leave no cross-root Markdown dependencies.

Inspect affected documents, nearby indexes, and needed references before writing. Read known exact paths directly; otherwise select an area through root and child indexes, follow useful links, and search within that area. Expand only for missing or incomplete indexes or absent material, identifying the gap and limiting expansion. Sufficient supplied evidence needs no repeated retrieval or sibling call. Whole-wiki audits need explicit scope.

## Choose What to Remember

Record established facts, user preferences or rules, adopted or rejected decisions, resolved questions, failure causes, significant attempts and direction changes, and supplied work results. Distant reuse value is not required. Preserve established milestones during long tasks rather than waiting for completion or mirroring every tool call. Distinguish facts, observations, plans, hypotheses, supplied states, and later corrections. Do not invent historical reasons or retain private reasoning, unnecessary secrets, or raw logs.

After web or other research, compare supplied results with existing wiki content and choose recording, update, or omission. Preserve new facts, methods, constraints, selection evidence, and corrections with sources, scope, and check date. Skip navigation alone, unsupported failed searches, empty browsing, and unchanged duplicates. Search uncertainty alone is not a finding; an established constraint or failure cause affecting the task can be. Keep transient information only when decision-relevant and dated. Refresh unchanged content only when its date or evidence matters. Group by resolved question and result, not query or visited page. Reuse sufficient research without a prerequisite researcher invocation.

Selection does not cancel authorized work-state recording or expand write authority. New knowledge can be absent while a supplied work record still needs preservation.

## Classify Subjects and Documents

Identify the owning subject, applicability, and document role: overview, version or migration, reusable pattern, case, or work record. Reusable concepts, methods, and cases belong in Knowledge's owning subjects; project-only designs, policies, and current state belong in the project's area. Project origin alone does not make reusable knowledge project-only. Preserve case-specific conditions rather than universalizing them.

Organize from domain to subtopic using actual content and existing useful conventions. Within the maximum folder depth below, do not impose a fixed number of levels, domain list, or repeated overview/version/pattern/case folder set. `Knowledge/agents/skills/` and `Knowledge/software/database/` are examples. Keep each body in one primary home; other areas may provide justified index or body links.

| Content relationship | Storage decision |
| --- | --- |
| Same question and applicability as an existing document | Update that document. |
| A separately useful question, version, case, or applicability | Check duplication and create a precisely named document. |
| One subject's current content supports at least two independent body documents | Immediately create or use its subject folder and split the content. |
| Different owning subjects are mixed | Separate them into their owning subject areas. |

The two-document threshold counts independent bodies supported by actual content, excluding `index.md`. One existing file plus new material can qualify; do not wait for two existing files. Distinct questions and applicability establish independence, not arbitrary sections or API counts. Avoid empty future categories, placeholder overviews, speculative documents, and mechanical splitting. One question can remain a precisely named file in an appropriate existing area. This creation rule does not require dismantling existing one-child folders or overriding exact-path instructions.

Folders name subjects; ordinary filenames and titles name specific questions and content. Indexes explain the subject and actual child roles. A single coordinates/layout document might use `design/tools/figma-coordinates-and-layout.md`. Two independent Figma bodies can use `design/tools/figma/coordinates-and-layout.md` and a specifically named sibling. These are examples, not files to pre-create. General list/detail split-view design belongs in its design-pattern subject, not under Figma merely because it was drawn there.

For knowledge-discovery or missing-knowledge requests, compare supplied material and authorized sources with existing knowledge for independent subject candidates. Body enrichment does not replace that review. Report new documents, updates, duplicate omissions, and uninspected scope according to actual results; create nothing to satisfy a count.

Perform necessary local classification and reference maintenance within authority. Existing flat placement does not settle ownership. An exact-file edit does not authorize other document moves. Complete authorized work and report out-of-scope restructuring or root changes without expanding into global cleanup. Knowledge does not inherit the work-record layout or template.

## Preserve Project-Owned Work Records

Match the supplied intent and identity with existing records. Every record belongs to one project in the selected root. Resolve ownership from the user's target, existing projects, and work context; create only the project and minimal indexes needed for authorized recording. Use `general` only for work without a specific project, not uncertain ownership. For cross-project work, choose one primary responsible project without duplication; ask if the evidence cannot resolve ownership.

New records use `Projects/<project>/records/YYYY-MM/000-<slug>.md`. The month is the initial recording month; the three-digit sequence starts at 000 within that project and month. Use the next number after the existing maximum and a short descriptive slug. Recheck numbers and destination immediately before saving; resolve collisions with the next number, never overwrite an existing file. General work uses `Projects/general/records/YYYY-MM/000-<slug>.md`.

Continue the same intent in the same record across later months, preserving the original month, number, and path. Bulk renaming or relocation needs separate scope. Root `Records/` and other non-project record placements are unsupported. Report migration needs and unsaved material; without migration authority, do not update unsupported locations, move or duplicate their records, or store work observations in knowledge documents instead. An exact target conflicting with the layout remains unsaved.

Maintain downward navigation: root > project > `records/index.md` > `YYYY-MM/index.md` > record. Indexes explain actual child roles. The project index need not duplicate individual record lists or add related-project links. The month folder is depth 4 with root at 0; create no deeper folder.

Preserve supplied conditions, checkbox states, nested results and evidence, prerequisites, verification, exclusions, limits, and later corrections. Do not derive conditions, change exclusions, or reassess fulfillment. If states conflict with evidence or information is missing, distinguish the supplied claim, observation, discrepancy, and missing check; return them to the task owner without silently changing states or claiming unsupported fulfillment.

Retain nested checkboxes and ordinary evidence bullets. Summarize each condition's result, evidence, and limits within six lines; put longer material in a related section or useful same-root document while preserving its relationship. Narrative or retrospective records and user facts need no invented conditions or preparation.

Save supplied confirmed milestones, significant failures, changes, interruptions, and completion snapshots during long tasks. The task owner maintains the temporary draft and lifecycle. Re-read before writing, serialize shared-file updates, and preserve concurrent edits. Do not delete/recreate an existing record to bypass a failed update.

## Explain Findings and Use Necessary References

The body must explain the question, findings, judgments, outcomes, and verification limits. Identify the subject, scope, observation date or version, and sources. Preserve relevant supplied goals, constraints, knowledge used, choices and rejection reasons, meaningful attempts, failures, changes, and unresolved limits. Do not force irrelevant fields, invent missing reasons, or confuse task-owner status with save verification. Resolve material conflicts from evidence or label them unresolved.

Link directly to the document or section needed for a concept, claim, observation, choice, rejection, or applicability judgment. Surrounding prose can explain the relationship; no separate relationship field is required. Evidence lists and source summaries can remain within one record. Sufficient body evidence needs no additional wiki links, fixed chain, or three-file structure. An originating-work link follows the same relevance test.

Do not create related-document sections or edited-file inventories merely because files were written, read, edited, or share a task or project. Preserve useful evidence lists and downward index navigation. Judge each reference direction separately; a useful A-to-B link does not require B-to-A.

Provide a complete change inventory only when explicitly requested or essential to file-by-file verification. Reuse an adequate diff or manifest; create a separate authorized artifact only when necessary. Do not generate an audit file to replace an omitted body inventory. Lists of files never replace findings and judgments.

Keep original URLs or repository/execution identifiers, dates, versions, applicability, and uncertainty. Paths alone are not evidence summaries. Understanding a historical result does not require rerunning every experiment or preserving all raw logs, scripts, and artifacts. Preserve reproduction material only when requested or needed to use the result, within authority.

Treat `/tmp`, `/private/tmp`, scratchpads, and temporary worktrees as disposable. Summarize dated observations and cite inspected original sources or preserved same-root material. Without a stable original, state the independent rechecking limit rather than inventing a replacement source. Mention temporary paths only for short-term recovery or handoff, identifying purpose and lifetime. Do not automatically archive raw material or move it across roots to manufacture evidence. Correcting old temporary references must preserve confirmed findings and limits.

## Save and Verify Locally

Before creating a missing root index, inspect existing files. Each index explains its purpose, scope, and actual child roles to guide the next reading choice. Use same-root relative Markdown links for downward navigation and justified body references; external sources do not replace index routes. Folder structure supplies parent locations. Parent/root return links and mechanical reciprocal links are unnecessary. Remove existing navigation-only links only within authorized navigation cleanup, preserving ancestor-index links with conceptual or evidentiary value.

Before creating or moving material, check the resulting structure, including necessary indexes. Maximum folder depth is 4 with root at 0. Underscore compound folder names have exactly two nonempty parts; simple names and filenames are outside that rule. Classify and split by subject and independent content.

If a required arrangement would exceed depth 4, consider existing-content updates, meaningful classification within the allowed depth, and authorized sibling placement or local reorganization. Do not merge unrelated content, delete knowledge, or use numbered buckets to satisfy structural limits. If preservation is impossible within scope and depth, leave affected material unsaved and report the needed decision. Initial cleanup of unrelated existing violations needs explicit scope; ordinary checks remain local.

For authorized splits, renames, moves, and merges, prepare destinations while preserving content, sources, dates, conditions, corrections, and uncertainty. Update indexes, incoming links, and relative references. Remove sources of renames/moves/merges only after destinations and affected links succeed; retain them on failure. A split may retain an original overview when useful, without duplicate bodies. Incoming-reference discovery may span the selected root; edits remain within affected content.

Re-read saved content and verify supplied states, subject/project ownership, applicability, roles, precise names, missed splits or mixed subjects, unnecessary dispersion, empty structure, duplicates, downward reachability, justified valid references, depth and folder names, and actual storage. Link and structural checks do not establish semantic placement. Preserve dated decisions and corrections. Report out-of-scope issues; local inspection does not establish whole-wiki validity.

## Return the Actual Storage Result

Report the root, saved paths, new independent documents, updates, duplicate omissions, core findings, unrecorded or unverified scope, supplied-state discrepancies, and save-check coverage concisely. Keep detail in the record. Claim discovery or completeness only within inspected sources. Preserve the distinction between task fulfillment and faithful memory/storage verification.

When recording fails, report unsaved material and the reason. Do not substitute another root or a knowledge document for a failed work-record save. Independently authorized knowledge updates may proceed under their own selection criteria. Do not claim completion with broken affected navigation. The task owner judges request fulfillment; writer verifies the supplied memory and actual storage.
