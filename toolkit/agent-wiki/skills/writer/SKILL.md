---
name: writer
description: Preserve authorized findings, decisions, user information, and supplied work-state snapshots in Agent Wiki with useful classification, evidence, and verified storage. Use directly from provided material; task planning and completion judgments belong to the task owner.
---

# Writer

Preserve supplied knowledge and work states as faithful, findable memory. Own selection, subject classification, document boundaries, metadata, deduplication, references, and storage verification. Start from user, task-owner, or tool evidence without requiring manager, researcher, or a checklist. The task owner owns conditions, progress, and fulfillment. Keep external research, product implementation, installation, arbitrary root selection, and routine whole-wiki restructuring outside this role.

## Root, Authority, and Sources

Use `agent-wiki list` and `agent-wiki path [name]`. If unavailable, parse `~/.agents/config.wiki.toml` with a TOML parser. Require a registered `default` and named `roots` with absolute existing-directory `path` values and optional string `description` values. Validate types and root names, canonicalize paths, and reject equal or nested roots. Configuration symlinks are supported; CLI updates preserve the link and atomically replace its resolved target. Unsupported configuration requires an update: report unsaved material without conversion, installation, settings changes, or a substitute root. Continue independent work.

Choose one root per record: explicit target, clearly current-work root, exact name/description match, then default for non-sensitive, unambiguous material. Roots bound disclosure, backup, and sync; folders classify subjects. Configuration, paths, wiki text, and CLI availability never confer or expand authority. Respect read-only, no-record, exact-path, original-text, and body-only instructions.

Ask when personal information, disclosure boundaries, or root choice remain mixed or unresolved. Do not guess a stricter root or automatically split, link, copy, move, or merge across roots. Cross-root moves require explicit source, destination, and scope and must leave no cross-root Markdown dependencies.

Inspect affected documents, nearby indexes, and needed references. Read known exact paths directly; otherwise select the relevant area through root and child indexes, follow useful links, and search within that area. Expand only for missing/incomplete indexes or absent material, identifying the gap and limiting expansion. Sufficient supplied evidence needs no repeated retrieval or sibling call. Whole-wiki audits require explicit scope.

## Obsidian Vocabulary and Save Checks

When `obsidian` is available and a registered vault name/ID is already known from authorized context, run `obsidian vault="<vault>" vault info=path` and compare its canonical result with the selected Agent Wiki root. Do not map equal names automatically. Put `vault=<vault>` first on every call and use exact vault-relative `path` values. Returned candidate paths must stay within the selected root and area.

After index-based area selection, find relevant notes and inspect their structured vocabulary:

```sh
obsidian vault="<vault>" search path="<area>" query='[type:reference] [tags:obsidian]' format=json limit=10
obsidian vault="<vault>" properties path="<note.md>" format=json
obsidian vault="<vault>" read path="<note.md>"
```

Replace placeholders with verified targets, safely quote arguments, and use Search quoting for multiword values. Check installed help when uncertain. Exact-path or sufficient supplied evidence needs no new search. `outline` and `search:context` can narrow body checks; limited results are not a complete inventory. `[tags:...]` targets structured tags, while `tag:...`, `tags`, and `tag` can include inline PR/color tokens. Reuse vocabulary only after checking properties and body fit. Whole-vault tag inventories require corresponding scope.

Inspect exit status, stdout meaning, and shape together. Exit 0 or `format=json` alone does not prove success. Recognize installed-version responses such as `No matches found.`, `No frontmatter found.`, and `Error: Property "..." not found.` as scoped empty search or missing metadata. Validate search JSON as a string-path array and properties JSON as a mapping; unexpected plaintext, shapes, paths, or errors remain failed/unverified, not empty knowledge. Confirm message/format behavior against installed help and observed output. If using tags JSON, validate the actual list of tag objects and the representation of optional count values rather than assuming numeric counts. On empty filters or missing metadata, continue selected-area index/body search.

If CLI access, mapping, or commands are unavailable or mismatched, continue with files/YAML in the same authorized root and report the limitation. Do not install, register, enable, reload, change settings, or discover other vault contents to obtain a mapping. CLI calls may start the app; honor session startup restrictions. This integration uses read-only commands, not `property:set`, create/append/move, configuration commands, or `eval`. Actual writing follows the authorized file/YAML merge contract below.

After saving, compare CLI `properties` and `read` with actual saved type, tags, and body. Retain raw YAML duplicate-key, type, content, and reference checks. If the cache and file disagree or CLI verification fails, distinguish actual storage from CLI confirmation and continue independent file verification. Do not overwrite content to make a cache check pass or silently claim CLI confirmation.

## Select Memory

Preserve established facts, user information, adopted/rejected choices, resolved questions, disproved hypotheses, failure causes, significant attempts and direction changes, and supplied work results. Distant reuse value is not required. Save established milestones during long tasks rather than waiting for completion or mirroring each tool call. Separate facts, observations, plans, hypotheses, supplied states, and later corrections. Do not invent historical reasons or retain private reasoning, raw logs, or unnecessary secrets.

After web or other research, compare supplied findings with existing content and choose new recording, update, or omission. Preserve new facts, methods, constraints, selection evidence, and corrections with sources, applicability, and dates. Skip navigation alone, unsupported failed searches, empty browsing, and unchanged duplicates. Search uncertainty alone is not a finding; established constraints or failure causes affecting the task can be. Keep transient information only when decision-relevant and dated. Refresh unchanged content only when its date or evidence matters. Group by resolved question and result, not query or visited page. Reuse sufficient research without requiring researcher execution. Selection neither expands authority nor cancels authorized work-state recording; supplied records may need preservation without new reusable knowledge.

## Classify Subjects and Bodies

Identify the owning subject, applicability, and role: overview, version/migration, reusable pattern, case, or work record. Reusable concepts, methods, and cases belong in Knowledge's owning subjects; project-only designs, policies, and current state belong in the project's area. Project origin alone does not make reusable knowledge project-only. Preserve case conditions rather than universalizing them.

Organize domain to subtopic from actual content and useful conventions. Within the maximum depth below, do not impose a fixed number of levels, domain list, or repeated overview/version/pattern/case folder set. `Knowledge/agents/skills/` and `Knowledge/software/database/` are examples. Keep one primary home per body; other areas can provide justified index/body links.

| Relationship | Decision |
| --- | --- |
| Same question and applicability | Update the existing document. |
| Separately useful question, version, case, or applicability | Check duplication and create a precisely named document. |
| One subject's current content supports two independent bodies | Immediately create/use its subject folder and split content. |
| Different owning subjects are mixed | Separate into their owning subject areas. |

Count independent bodies supported by current content, excluding `index.md`. One existing file plus new material can qualify; do not wait for two existing files. Distinct questions and applicability establish independence, not arbitrary sections or API counts. Avoid empty future categories, placeholders, speculative documents, and mechanical splitting. One question can remain a precise file in an appropriate area. This rule does not require dismantling existing one-child folders or overriding exact paths.

Folders name subjects; ordinary names and titles identify specific questions/content. Indexes explain the subject and actual child roles. One coordinates/layout document might use `design/tools/figma-coordinates-and-layout.md`; two independent Figma bodies can use `design/tools/figma/coordinates-and-layout.md` and a specifically named sibling. These are examples, not files to pre-create. General list/detail split-view design belongs to its design-pattern subject, not under Figma merely because it was drawn there.

For discovery or missing-knowledge requests, compare supplied material and authorized sources with existing knowledge for independent subject candidates. Body enrichment does not replace that review. Report actual new documents, updates, duplicate omissions, and uninspected scope; create nothing to satisfy a count. Existing flat placement does not settle ownership. Perform necessary local classification/reference maintenance within authority; an exact-file edit does not authorize other moves. Complete permitted work and report out-of-scope restructuring/root changes without global cleanup. Knowledge does not inherit work-record paths or body templates.

## Document Properties

For ordinary knowledge notes and project work records being created/updated, maintain one YAML frontmatter mapping at the beginning. `type` carries primary character, `tags` searchable topics/technologies, folders ownership, and body links conceptual/evidentiary/decision relationships. Indexes own downward navigation and do not require these fields; attachments are outside this metadata contract.

`type` is one nonempty string from the main purpose. Reuse equivalent existing vocabulary; examples include `reference`, `comparison`, `design`, `record`, `diagnosis`, `benchmark`, and `guide`. They are not a closed enum; check duplication before introducing a needed type. A measurement-focused work record can be `benchmark`, a mixed conditions/results record `record`. Type does not establish success, completion, or evidence strength.

`tags` is a string list of actual topic/technology keywords. Reuse equivalent spellings from the authorized relevant area before adding a keyword. Use `tags: []` when none is justified. Keep vocabulary open without fixed lists or counts and without inspecting unrelated roots/whole-wiki contents just to choose tags. For new English tags, use lowercase and needed hyphens. Preserve established Korean/proper-term spellings. Values have no `#`, category prefixes, spaces, or slashes and contain a non-numerical character. Avoid case-only duplicates and redundant synonyms. Select useful retrieval topics, not every mentioned noun/version/work state. A project name qualifies as an actual topic, not an automatic copy of folder ownership.

```yaml
---
type: benchmark
tags:
  - llm
  - uzu
  - performance
---
```

This is a measurement-document example, not fixed vocabulary or a body template. Read existing frontmatter and merge within its mapping, preserving unrelated fields/values, body, and references. Reuse valid, content-fitting type/tags. Do not add a second frontmatter block or duplicate keys. Quote strings that could parse as another YAML type.

Apply metadata within authorized edits. Whole-wiki tagging, normalization, moves, and settings are separate work. Respect body-only/original-text boundaries and continue independent authorized body work. Check authority before converting established hierarchical tags or other user conventions. Malformed YAML, field-type conflicts, or uncertain meaning require preserving affected metadata, continuing independently authorized body work, and reporting unapplied fields rather than silently overwriting them. Verify the saved mapping, key uniqueness, nonempty string type, string-list tags, keyword syntax, duplicates, content fit, and existing field/body/evidence preservation. Metadata complements placement, navigation, and permission checks.

## Project-Owned Records

Match supplied intent/identity with existing records. Every record belongs to one project in the chosen root. Resolve ownership from the user's target, existing projects, and work context; create only the project/minimal indexes needed for authorized recording. Use `general` only for work without a specific project, not uncertain ownership. For cross-project work, choose one primary responsible project without duplication; ask if unresolved.

New records use `Projects/<project>/records/YYYY-MM/000-<slug>.md`: initial recording month, three-digit sequence from 000 within that project/month, and a short descriptive slug. Use the next number after the existing maximum. Recheck numbers/destination immediately before saving; resolve collisions with the next number, not overwrite. General work uses `Projects/general/records/YYYY-MM/000-<slug>.md`.

Continue the same intent in the same record across later months, preserving month, number, and path. Bulk renaming/relocation requires separate scope. Root `Records/` and other non-project placements are unsupported. Report migration needs and unsaved material; without migration authority, do not update unsupported locations, move/duplicate their records, or substitute knowledge documents for work observations. An exact target conflicting with layout remains unsaved.

Maintain root > project > `records/index.md` > `YYYY-MM/index.md` > record. Indexes explain actual child roles; the project index need not duplicate record lists or add related-project links. The month folder is depth 4 with root at 0; add no deeper folder.

Preserve supplied conditions, checkboxes, nested results/evidence, prerequisites, verification, exclusions, limits, and corrections. Do not derive conditions, change exclusions, or reassess fulfillment. For evidence conflicts or gaps, distinguish supplied claim, observation, discrepancy, and missing check and return them to the task owner without silently changing states or claiming unsupported fulfillment. Keep nested checkboxes and ordinary evidence bullets; summarize each condition's results/evidence/limits within six lines. Put longer material in a related section/useful same-root document without losing its relation. Narrative/retrospective records and user facts need no invented conditions/preparation.

Save supplied confirmed milestones, significant failures/changes, interruptions, and completion snapshots during long tasks. The task owner keeps the temporary draft/lifecycle. Re-read before writing, serialize shared-file updates, preserve concurrent edits, and do not delete/recreate a record to bypass update failure.

## Body and References

Explain the question, findings, judgments, outcomes, and verification limits in the body. Identify subject, scope, observation date/version, and sources. Preserve relevant supplied goals, constraints, knowledge, choices/rejection reasons, meaningful attempts, failures, changes, and unresolved limits. Do not force irrelevant fields, invent reasons, or confuse task status with save checks. Resolve material conflicts from evidence or label them unresolved.

Link directly to the needed document/section for a concept, claim, observation, choice, rejection, or applicability. Prose can explain the relationship without a separate field. Evidence lists/summaries may remain in one record; sufficient body evidence needs no extra links, fixed chain, or three-file structure. Originating-work links use the same relevance test. Do not create related-document sections or edited-file inventories merely because files were used/edited or share a task/project. Preserve useful evidence lists/downward indexes. Judge each direction independently; A-to-B does not require B-to-A.

In authored/edited prose, join a known pull request's kind and number as `PR#150` and wrap the whole identifier in inline code, with grammatical particles outside. A known real URL uses the same code-formatted link label and unchanged target. Do not guess kind, number, or URL. For repository ambiguity, include repository identity in the same code span.

```md
`PR#150`은 참고로 남깁니다.
[`PR#150`](https://github.com/opnay/blog-server/pull/150)에서 확인했습니다.
`blog-server PR#150`과 `another-repo PR#150`을 구분합니다.
```

Preserve quotations, code blocks, URLs, frontmatter, intentional tags, and original-text instructions. Apply notation within authorized prose edits rather than bulk-replacing other documents. Verify kind/number, code boundaries, particles, and actual URL after saving.

Provide complete change inventories only when explicitly requested or essential to file-by-file verification. Reuse adequate diffs/manifests; create separate authorized artifacts only if necessary. Do not automatically generate audit files to replace body inventories. File lists never replace findings/judgments.

Keep original URLs or repository/execution identifiers, dates, versions, applicability, and uncertainty. Paths alone are not evidence summaries. Historical understanding does not require rerunning every experiment or retaining all logs/scripts/artifacts. Keep reproduction material only when requested or needed to use the result, within authority. Treat `/tmp`, `/private/tmp`, scratchpads, and temporary worktrees as disposable. Summarize dated observations and cite inspected originals or preserved same-root material. Without a stable original, state rechecking limits rather than inventing one. Temporary paths belong only in short-term recovery/handoff with purpose/lifetime. Do not automatically archive or cross-root-copy raw material to manufacture evidence. Correcting old temporary references must preserve confirmed findings/limits.

## Save and Report

Inspect existing files before creating a missing root index. Every index explains its purpose, scope, and actual child roles. Use same-root relative Markdown links for downward navigation and justified body references; external sources do not replace index routes. Folder structure supplies parents. Parent/root return links and mechanical reciprocity are unnecessary. Remove existing navigation-only links only within authorized cleanup, preserving ancestor-index links with conceptual/evidentiary value.

Before creating/moving material, check the resulting structure including indexes. Maximum folder depth is 4 with root at 0. Underscore compound folders have exactly two nonempty parts; simple names/filenames are outside this rule. Classify/split by subject and independent content. If an arrangement exceeds depth 4, consider existing-content updates, meaningful classification within allowed depth, and authorized sibling/local reorganization. Do not merge unrelated content, delete knowledge, or use numbered buckets to meet limits. If preservation is impossible within scope/depth, leave affected material unsaved and report the needed decision. Unrelated existing violations need explicit cleanup scope; routine checks stay local.

For authorized splits/renames/moves/merges, prepare destinations preserving content, sources, dates, conditions, corrections, and uncertainty; update indexes/incoming/relative links. Remove rename/move/merge sources only after destinations and affected links succeed; retain them on failure. Splits may retain a useful original overview without duplicate bodies. Incoming-reference discovery may span the selected root; edits stay within affected content.

Re-read saved content and verify supplied states, ownership, applicability/roles, precise names, metadata, missed splits/mixed subjects, dispersion, empty structures, duplicates, downward reachability, justified valid references, depth/names, and actual storage. Link/structure checks do not establish semantic placement. Preserve dated decisions/corrections. Report out-of-scope issues; local inspection does not establish whole-wiki validity.

Return root, saved paths, new documents, updates, duplicate omissions, findings, unapplied metadata, unrecorded/unverified scope, CLI confirmation limits, state discrepancies, and actual save-check coverage concisely. Keep detail in the record; claim discovery/completeness only within inspected sources. The task owner judges fulfillment; writer verifies faithful memory/storage. On failure, report unsaved material/reasons without substituting another root or knowledge document for a work-record save. Independently authorized knowledge updates may continue under their criteria. Do not claim completion with broken affected navigation. Distinguish completed body work, verified storage, and metadata/CLI checks that remain unapplied or unverified.
