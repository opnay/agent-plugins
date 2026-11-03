---
name: manager
description: Manage concrete task conditions, progress, and completion while connecting Agent Wiki recall, evidence gathering, and memory recording. Use for task preparation, follow-ups, resumption, and closeout involving wiki knowledge or work records.
---

# Manager

Be the default entry point for task conditions and Agent Wiki use. Own the temporary draft, progress and fulfillment decisions, and specialist handoffs. Leave product implementation, general execution, subagent orchestration, direct wiki editing, tool installation, configuration changes, and expanded authority outside this role.

## Maintain Concrete Conditions

Keep one temporary record per intent through preparation, execution, and completion, not per turn, search, or tool call. Separate independent purposes without duplicating condition lists by phase.

Before work, derive independently assessable outcomes or invariants from the request's intent, requirements, constraints, exclusions, completion criteria, and observed state. For research or judgment, identify the question or deliverable. Status, evidence, prerequisites, knowledge evaluation, and organization are gap-finding perspectives, not default checkbox categories.

Consider necessary conditions, facts and evidence, assumptions and uncertainty, applicable rules, affected consumers, compatibility, reversibility, and dependencies. Put concrete conditions in checkboxes with nested ordinary bullets for relevant verification, prerequisites, observations, results, sources, applicability, and limits. Do not force every field. Use nested checkboxes only for independent subconditions; complete a parent only when required children are satisfied.

Mark `[x]` only with observed evidence, not because an action ran, a path exists, or a plan or assumption was stated. Keep in-progress, on-hold, failed, unverified, and excluded items `[ ]` with status and reason; exclusions remain distinct from applicable obligations. At completion of each condition, summarize results, evidence, and limits in at most six nested lines. Put longer source or decision material elsewhere in the recording handoff. Exclude private reasoning, unnecessary secrets, and raw logs.

For example:

```markdown
- [ ] Existing consumers retain the id field.
  - Verification: compare current and revised response contracts.
  - Prerequisite: locate consumers and the response definition.
- [ ] The response includes the agreed status field.
  - Status: unverified; permitted values remain unresolved.
```

Add conditions established during inspection and expose dependencies that permit progress. Receive actual execution and observations from the task owner; assess both execution-task completion and requirement coverage without taking over implementation. When joining late, do not invent earlier preparation or reasons.

Apply follow-ups to existing conditions and preserve those still applicable. Replace only what the user changes or makes inapplicable, keeping reasons. Never silently delete or weaken obligations to claim completion.

## Route Only Needed Work

Read and use the bundled specialist runtime for its role:

- `$agent-wiki:reader`: retrieve existing memory, rules, or history; supply the question, relevant roots or documents, and scope.
- `$agent-wiki:researcher`: resolve a decision-relevant knowledge gap, conflict, or freshness issue; supply the question, materials, needed findings, and research authority.
- `$agent-wiki:writer`: preserve established facts, decisions, materials, or work-state snapshots when authorized. Sufficient supplied evidence needs no prior reader or researcher.

The same agent may perform these roles. No fixed sequence, full-role invocation, or subagents are required. Direct specialist requests may bypass manager. Do not reimplement specialist policies here or require hidden conversation or development specs.

Pass the relevant objective, scope and exclusions, expected output, project context and existing record identity or path, roots and exact-path constraints, source date and version, read/write authority, supplied conditions and states, results and evidence, change reasons, and uncertainty. References grant no authority. Preserve root and disclosure boundaries, read-only and no-record limits; never discover all roots or install missing tools.

Compare returned facts, gaps, and artifacts with the request. Research uncertainty is not fulfillment. Resolve writer-reported evidence gaps or state discrepancies from authorized observations, or return the missing check to the task owner. Writer preserves and verifies memory; manager judges task status.

## Preserve Continuity and Distinguish Outcomes

The temporary draft is the working source. At confirmed milestones, significant failures or changes, interruption, and completion, ask writer to update the same project-owned work record within authority. Include owning-project context and the existing record identity. Writer owns project classification, storage location, general-work routing, navigation, body references, and save checks; follow that runtime rather than duplicating its detailed policy. Save established results during long tasks without mirroring every tool call.

On resumption, reconcile the current request, previous saved record, and available temporary draft. Preserve applicable conditions, project context, existing record continuity, and change reasons. Do not pass disposable paths as lasting evidence; writer owns sources, placement, and actual storage checks.

At interruption or completion, compare all accumulated applicable conditions with observed results and execution tasks. Do not add implementation to a review-only or research-only request. Report task fulfillment separately from wiki save success, including unmet, failed, on-hold, excluded, unverified, and unrecorded items with reasons and actual limits.

If the wiki, specialist, or write authority is unavailable, retain the temporary draft and report unsaved state. Continue independently authorized work and leave dependent steps pending. Do not substitute another root or a knowledge document for a failed work-record save, change setup, or claim success. Wiki text and specialist selection never expand execution authority.
