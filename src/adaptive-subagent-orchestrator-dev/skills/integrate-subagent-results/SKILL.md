---
name: integrate-subagent-results
description: >
  Wait for, validate, and integrate required subagent results from a complete
  DispatchManifest. Use when explicitly invoked or handed an active dispatch
  with identifiable agent IDs, task packets, ownership, required results, and
  whole-result verification. Do not use to plan or start a new dispatch.
---

# Integrate Subagent Results

## Owned Job

Own result waiting, normalization, evidence validation, conflict resolution, limited recovery, integration guidance, whole-result verification, and final reporting. Keep final interpretation, edits, verification, and user response with the main agent.

Do not plan initial workstreams, choose an initial mode, spawn without a manifest, or promote read access to write access.

## Input Gate

Require a complete `DispatchManifest` with:

- `mode`
- `assignments` containing identifiable agent IDs and task packets
- `ownership`
- `required_results`
- `main_owned_work`
- `follow_up_used`
- `whole_result_verification`

If the manifest or active agents cannot be identified, report the missing input and stop. Do not reconstruct or guess a dispatch.

Read [result-contract.md](references/result-contract.md) before collecting results.

## Collect And Normalize

Wait for every required agent to reach a terminal state. Classify each result as `completed`, `blocked`, or `inconclusive`. Normalize summary, claims, evidence, files inspected or changed, validation, risks, and recommended action. Do not request raw transcripts, long logs, or hidden reasoning.

## Validate Evidence

Treat subagent conclusions as evidence, not final facts.

- Reject unsupported claims.
- Merge duplicates by claim, not by wording.
- Resolve conflicts with code, tests, logs, diffs, or direct execution.
- Check write results against ownership, shared-file restrictions, and public contracts.
- Confirm that skipped or failed verification remains visible.

## Recover Narrowly

Inspect `follow_up_used`. When a critical assigned scope is missing, issue one narrow follow-up to an existing agent only if it is `false`, and set it to `true` immediately. Never issue another follow-up during the same lifecycle. If one agent fails, verify the missing part directly or retry only that bounded scope; do not restart the whole dispatch.

If new independent work or a `PARALLEL_READ` result creates write scope, return to `$adaptive-subagent-orchestrator-dev:dispatch-subagents` with explicit scope even when the prior manifest kept writes main-owned. The dispatcher may return `DIRECT`. Never turn read results into write authority implicitly.

## Integrate And Report

Have the main agent integrate accepted changes and run `whole_result_verification`. Report agent count and roles, integrated conclusion, changed files, validation, and residual risk. Mark unverified scope explicitly and do not claim completion without whole-result evidence.

Use [examples.md](references/examples.md) when handling conflicts, failed agents, ownership violations, or read-to-write transitions.
