---
name: verify
description: Verify artifacts or design their verification by deriving relevant criteria and methods from their purpose, characteristics, and risks, then checking actual evidence. Use for artifact validation, quality review, acceptance checks, or verification planning within an authorized task. Do not turn ordinary creation into a mandatory broad audit. artifact verification, quality review, validation plan, acceptance evidence, verification procedure
---

# Verify

Determine what an artifact needs to satisfy, choose suitable ways to check it, and perform authorized verification against actual evidence. Apply across domains and artifact types; a domain label or fixed checklist does not define the boundary.

## Establish The Task

Inspect the artifact and original requirements before selecting checks. Establish its purpose, consumers, use context, constraints, relevant contracts, changed scope, environment, and the decision the user needs. Separate the author's explanation and success claims from your observations.

Identify the version or state you actually check using appropriate evidence: a commit, build, file version, rendered state, or data date. Do not require hashes or a separate ledger for every task.

Use available context to fill ordinary gaps. Ask only when unresolved information materially changes judgment, scope, or authority; continue independent checks whose prerequisites are satisfied.

A request for a verification plan produces a plan, not execution. Perform actual checks when requested or needed within the authorized task. A review does not authorize editing the artifact. Verification fixtures, temporary output, external operations, credentials, and costs stay within existing permissions.

## Derive Criteria And Methods

Derive quality dimensions from the artifact's properties, claims, relationships, change effects, and credible failure modes. Keep original requirements and mandatory acceptance conditions in scope. Add a new dimension when the artifact needs it; do not force every example below onto every task.

- Behavior, states, failure, and recovery may need real-path execution, boundary cases, and regression checks.
- Structure, ownership, dependencies, maintainability, and compatibility may need source inspection, static analysis, consumer checks, and change-impact analysis.
- Visual, spatial, perceptual, and interaction properties may need actual rendering or artifact observation across relevant content, sizes, and states.
- Claims, logic, decisions, and communication may need requirement comparison, consistency checks, counterexamples, or reader-task review.
- Numbers, data, and measurements may need independent recalculation and checks of definitions, samples, missingness, periods, error, and reproducibility.
- Evidence and provenance may need original-source inspection, claim support, independence, counterevidence, freshness, and applicability checks.
- Permissions, information flows, integrations, and operations may need threat, interface, failure, and recovery analysis with authorized real-path checks.

For each important claim, establish a criterion and what evidence could support or contradict it. Distinguish explicit requirements, owning contracts, inspected standards, and observations from derived criteria, assumptions, and unresolved policy. Do not invent a settled standard or decide a material policy question for its owner.

Use available specialist knowledge, references, tools, and role skills when relevant. No particular role plugin, orchestrator, service, or development spec is required. Missing expertise, evidence, or access remains a limit; installing this skill does not supply it.

Choose the least costly method that directly checks the claim. Add another method only for a material risk the existing evidence cannot reach. Define relevant inputs, states, boundaries, observable results, and failure conditions. Match depth to uncertainty, impact, and the requested scope rather than an arbitrary check count, score, or pass rate.

## Execute And Examine Evidence

Perform the selected checks within authority. Inspect their meaningful results: a command invocation, exit code, completed checklist, installed skill, or another agent's conclusion alone does not establish the artifact's quality.

Prefer evidence that can be compared independently with the original criterion. Do not treat the artifact's own generated expected values, repeated claims, or the author's pass declaration as independent support. Check whether a credible contract violation would produce a different observation.

Keep the proof boundaries visible. Source inspection does not establish visual appearance or runtime behavior. A local check does not establish an external integration or deployed state. When a needed observation is unavailable, state what the remaining evidence does establish.

Distinguish defects from preferences, improvement proposals, and untested hypotheses. For subjective or inferred judgments, state the basis, context, and uncertainty rather than presenting taste as an objective violation.

If the artifact, environment, or original input changes, reevaluate affected evidence. Do not apply an earlier version's success to a new snapshot without checking its relevance.

## Judge And Report

Connect each important result to its checked scope, criterion and source, method, actual evidence, judgment, and remaining uncertainty. Use the simplest useful prose, list, or table; no mandatory ledger or numeric score.

| Status | Meaning |
| --- | --- |
| `pass` | Sufficient evidence supports the criterion within the identified scope. |
| `fail` | An actual criterion violation is established. Give the location or state, conditions, and impact. |
| `inconclusive` | Evidence, criteria, or expertise are insufficient or conflicting. |
| `not-run` | The check was not executed. State why and how that limits the conclusion. |

Separate artifact failures from tool, environment, fixture, and verification-procedure failures. Do not classify an unavailable tool as an artifact defect, or an unperformed check as a pass.

For the overall result, distinguish **confirmed defects**, **checked scope satisfied**, and **judgment deferred**. Do not average away an important defect. Do not declare an overall pass when mandatory or important verification is unperformed or inconclusive. Explain other residual risks, exclusions, and recommendations with their scope and reason. A passing check does not establish universal quality.

Return the evidence needed to act: findings and impact, missing or conflicting evidence, relevant owner, and the next useful check. A plan-only response instead gives proposed dimensions, criteria, procedures, prerequisites, observable evidence, success/failure conditions, and open questions; never label these as observed results.

## Stop Or Reverify

Finish when important claims in the requested scope are judged with evidence and remaining risks or unknowns have an explained effect. Do not keep checking every possible quality dimension.

Preserve original obligations. Do not drop important conditions or shrink the scope to claim completion. If permissions, time, tools, or cost prevent necessary checks, return the partial result and unmet scope with reasons.

Expand a check when new evidence exposes a material omission or invalid assumption. Avoid repeating checks that add no decision-relevant evidence. Explain corrections to faulty criteria or fixtures; do not relax them after a failure merely to obtain a pass. Refer changes to goals, policy, or authority to the owner.

Verification findings do not authorize repairs. If a repair is separately authorized, recheck affected claims and relevant regressions against the applicable criteria. State which evidence remains valid and which was rerun.

## Work With Other Agents

Use this contract directly or within an authorized orchestration. Scheduling, agent creation, model profiles, and permission policy stay with their owners; an assigned verifier does not gain delegation authority.

A verification packet should identify the artifact/state, original requirements, purpose, constraints, permitted tools/actions, and evidence format. Specialist skills can supply professional criteria. Do not rely on hidden sibling context or use the author's desired conclusion as the answer.

Treat returned agent judgments as evidence to assess. Recheck material claims against the artifact, original criterion, and actual results. A separate agent name alone does not establish independent evidence or expertise. Reporting a verification result does not grant release approval or external-action authority.
