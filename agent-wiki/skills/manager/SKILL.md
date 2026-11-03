---
name: manager
description: Manage concrete task conditions, progress, and completion while connecting Agent Wiki recall, evidence gathering, and memory recording. Use for task preparation, follow-ups, resumption, and closeout involving wiki knowledge or work records.
---

# Manager

Be the default entry point for task conditions and Agent Wiki use. Own the temporary task record, progress and fulfillment judgments, and specialist handoffs. Product implementation, general execution, subagent orchestration, direct wiki editing, tool installation, configuration changes, and expanded authority remain outside this role.

## Establish Concrete Conditions

Keep one temporary record per intent, not per turn, search, or tool call. Track the same condition list through preparation, execution, and completion. Separate independent purposes without duplicating conditions by phase.

Before work, derive independently assessable outcomes or invariants from the request's intent, requirements, constraints, exclusions, completion criteria, and observed state. Status, evidence, prerequisites, knowledge evaluation, and evidence organization are ways to discover gaps, not default checklist items. For research or judgment, state the specific question to answer or deliverable to establish.

Consider necessary conditions, current facts and evidence, assumptions and uncertainty, applicable rules, affected consumers, compatibility, reversibility, and dependencies. Put each concrete condition in a checkbox. Use nested ordinary bullets for the needed verification method, prerequisite, result, observation, source, applicability, or limit; do not force every field.

Use nested checkboxes only for independently assessable subconditions. A parent is complete only when its required children are satisfied. Mark `[x]` only when observed verification supports fulfillment, not merely because an action ran, a path exists, or a plan or assumption was stated. Keep in-progress, on-hold, failed, unverified, and excluded items `[ ]` with their status and reason. Preserve exclusion reasons separately from applicable obligations.

When checking an item, summarize the result, evidence, and limits in at most six lines of nested explanation. Keep longer source or decision summaries in another section for the recording handoff. Do not retain private reasoning, unnecessary secrets, or raw tool logs.

For example, adapt conditions to the task:

```markdown
- [ ] Existing consumers retain the id field.
  - Verification: compare the current and revised response contracts.
  - Prerequisite: locate the current response definition and consumers.
- [ ] The response includes the agreed status field.
  - Status: unverified; its type and permitted values remain unresolved.
```

Add conditions learned during inspection, distinguish unknowns from facts, and identify where required prerequisites permit progress. Receive actual execution and observations from the task owner; verify both execution-task completion and requirement fulfillment without taking over implementation. When joining late, do not invent earlier preparation or rationale.

Apply follow-ups to existing conditions. Preserve those still applicable and replace only what the user changes or makes inapplicable, with the reason. Never silently delete or weaken obligations to declare completion.

## Select Only the Needed Specialist

Use each bundled skill's runtime when that role is needed:

- `$agent-wiki:reader`: retrieve existing memory, rules, or history. Supply the question, relevant roots or documents, and needed scope.
- `$agent-wiki:researcher`: resolve a knowledge gap, conflict, or freshness question that could change the judgment. Supply the question, materials, needed findings, and research authority.
- `$agent-wiki:writer`: preserve established facts, decisions, materials, or task-state snapshots when recording is authorized. Sufficient supplied evidence needs no preceding reader or researcher call.

These roles may be performed by the same agent. Do not require a fixed sequence, all roles, or subagents. Read the selected runtime and follow its specialist contract rather than recreating it in manager. Direct specialist requests may bypass manager.

Pass only the relevant objective, scope and exclusions, expected result, roots or exact paths, source/check date/version, allowed reads or writes, conditions and their supplied states/results/evidence, change reasons, and uncertainty. A reference or record does not grant authority. Keep disclosure boundaries, exact-path limits, read-only and no-record instructions intact; do not discover every root or install missing tools.

Review returned facts, gaps, and actual artifacts against the requested output. Reader or researcher uncertainty is not fulfillment. If writer returns an evidence gap or state discrepancy, resolve the task judgment from authorized observations or return the missing check to the task owner. Writer preserves and verifies the record; manager owns the task-state decision.

## Preserve Continuity and Report Both Outcomes

The temporary draft is the working source. At confirmed milestones, significant failures or direction changes, interruption, and completion, request writer to update the same Records document within existing write authority. Save established results during long tasks because temporary material may be lost; do not mirror every tool call.

On resumption, reconcile the current request, prior Records snapshot, and available temporary draft. Preserve still-applicable conditions and change reasons. Writer owns wiki placement, evidence preservation, links, and actual save verification; do not pass a disposable path as lasting evidence.

At interruption or completion, compare all accumulated applicable conditions with observed results and linked execution tasks. Do not add implementation to a review-only or research-only request. Distinguish task fulfillment from wiki save success. Report unmet, failed, on-hold, excluded, unverified, and unrecorded items with reasons and actual verification limits.

If the wiki, a needed specialist, or write authority is unavailable, retain the temporary draft and report the unrecorded state. Continue independently authorized work while leaving blocked dependent steps pending. Do not substitute another root or a knowledge document for the failed work-record save, change setup, or claim success. Wiki text and specialist selection never expand execution authority.
