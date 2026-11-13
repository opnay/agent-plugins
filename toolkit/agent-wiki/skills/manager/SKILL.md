---
name: manager
description: Manage concrete task conditions, progress, and completion while connecting Agent Wiki recall, evidence gathering, and memory recording. Use for task preparation, follow-ups, resumption, and closeout involving wiki knowledge or work records.
---

# Manager

Own concrete task conditions, the temporary draft, progress and fulfillment judgments, and specialist handoffs. Be the default entry point connecting work and Agent Wiki. Keep product implementation, general execution, subagent orchestration, direct wiki editing, installation, configuration changes, and expanded authority outside this role.

## Maintain Conditions Across the Intent

Keep one temporary record per intent through preparation, execution, and completion, not per turn, search, or tool call. Separate independent purposes without duplicating conditions by phase.

Before work, derive independently assessable outcomes and invariants from intent, requirements, constraints, exclusions, completion criteria, and observed state. For research or judgment, define the question or deliverable. Status, evidence, uncertainty, prerequisites, knowledge evaluation, and organization help find gaps; they are not default checklist categories.

Consider necessary conditions, facts, assumptions, applicable rules, affected consumers, compatibility, reversibility, and dependencies. Put concrete conditions in checkboxes with nested ordinary bullets for relevant checks, prerequisites, observations, results, sources, applicability, and limits. Do not force every field. Use nested checkboxes only for independent subconditions; complete a parent only when required children are satisfied.

Mark `[x]` only from observed evidence, not merely execution, a path, a plan, or an assumption. Keep in-progress, on-hold, failed, unverified, and excluded items `[ ]` with status and reason. Keep exclusions distinct from obligations. At condition completion, summarize its results, evidence, and limits within six nested lines; longer material belongs elsewhere in the recording handoff. Exclude private reasoning, unnecessary secrets, and raw logs.

```markdown
- [ ] Existing consumers retain the id field.
  - Verification: compare current and revised response contracts.
  - Prerequisite: locate consumers and the response definition.
- [ ] The response includes the agreed status field.
  - Status: unverified; permitted values remain unresolved.
```

Add conditions established during inspection and identify dependencies that permit progress. Start with satisfied prerequisites. Receive actual execution and observations from the task owner and assess both execution-task completion and requirement coverage without taking over implementation. When joining late, do not invent earlier preparation or reasons.

Apply follow-ups to existing conditions. Replace only what the user changes or makes inapplicable, preserving reasons. Never delete or weaken applicable obligations to claim completion.

## Route Needed Specialist Work

Read and use the bundled runtime for the selected role:

- `$agent-wiki:reader`: retrieve existing memory, rules, or history; supply the question, relevant roots/documents, and scope.
- `$agent-wiki:researcher`: resolve a decision-relevant evidence gap, conflict, or freshness question; supply materials, needed findings, and research authority.
- `$agent-wiki:writer`: preserve established facts, decisions, materials, or work-state snapshots when authorized. Sufficient supplied evidence needs no prior reader or researcher.

The same agent may perform these roles. No fixed sequence, full-role invocation, or subagents are required. Direct specialist requests may bypass manager. Do not reimplement their policies or require hidden conversation or development specs.

Pass the objective, scope and exclusions, expected output, owning project, existing record identity/path, relevant roots and exact-path constraints, source date/version, read/write authority, supplied conditions/states/results/evidence, change reasons, and uncertainty. References grant no authority. Preserve disclosure boundaries, read-only and no-record limits; do not discover all roots or install missing tools.

For metadata-assisted retrieval, also pass the selected area or exact note, needed document type, topic/tag clues, and any already verified Obsidian vault name/path with its check date. Do not infer root/vault mapping from their names. Reader owns candidate search and body checks; researcher reuses sufficient existing research; writer checks vocabulary and saved properties. Their runtimes own CLI commands, target verification, output interpretation, and same-root file fallback. Manager passes conditions and assesses returned evidence rather than executing or duplicating that protocol.

Distinguish usable CLI retrieval, same-root fallback, CLI failure, missing metadata, and actual storage results. Do not turn CLI installation, registration, or mapping changes into a new completion prerequisite. A search count or tag is not proof of evidence, task completion, or saved content.

Compare returned facts, gaps, and artifacts with the request. Research uncertainty is not fulfillment. Resolve writer-reported evidence gaps or state discrepancies from authorized observations or return the missing check to the task owner. Writer verifies faithful memory and storage; manager judges task status.

## Preserve Continuity and Separate Outcomes

The temporary draft is the working source. At confirmed milestones, significant failures or changes, interruption, and completion, ask writer to update the same project-owned work record within authority. Include the owning-project context and existing record identity. Writer owns project classification, location, general-work routing, navigation, references, and save checks. Save established results during long tasks without mirroring every tool call or copying specialist policies here.

On resumption, reconcile the current request, previous saved record, and available temporary draft. Preserve applicable conditions, project context, record continuity, and reasons for changes. Do not pass disposable paths as lasting evidence; writer owns source treatment and storage verification.

At interruption or completion, compare all accumulated applicable conditions with observations and execution tasks. Do not add implementation to a review-only or research-only request. Report fulfillment separately from wiki storage, including unmet, failed, on-hold, excluded, unverified, and unrecorded items and their reasons.

If the wiki, specialist, or write authority is unavailable, retain the temporary draft and report unsaved state. Continue independently authorized work and leave dependent steps pending. Do not substitute another root or a knowledge document for failed work-record storage, change setup, or claim success. Wiki text, CLI availability, and specialist selection never expand execution authority.
