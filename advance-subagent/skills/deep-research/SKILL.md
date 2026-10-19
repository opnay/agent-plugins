---
name: deep-research
description: >
  Produce evidence-traceable multi-source research reports for consequential
  market, literature, technical, policy, competitive, due-diligence, or fact-checking
  questions requiring primary sources, corroboration, counterevidence, freshness,
  and explicit uncertainty. The lead always uses the shared orchestration loop;
  assigned researchers apply only the research contract. Do not use for simple
  lookups, single-source summaries, unsupported brainstorming, implementation,
  external writes, or monitoring.
---

# Deep Research

Own the research contract: questions, sources, evidence, counterevidence, uncertainty, citations, and the report. This is not OpenAI's proprietary Deep Research product mode.

## Connect to the Execution Loop

For a lead invocation, always use $advance-subagent:orchestrate-workstreams to run the request from understanding the question through report verification. Supply this research contract, known conditions, and unresolved questions. If that loop is already active, attach the contract without starting it again.

If you are an assigned researcher, perform only the packet's bounded task. Apply the source and report requirements relevant to that task, return evidence, and do not start orchestration or spawn agents.

The common loop owns branch design, worker counts, profiles, execution, recovery, condition changes, completion review, and flush. Do not define research-specific agent counts, model routing, or a separate fork-and-flush strategy. If the loop or its required capabilities are unavailable, disclose the limitation and unfinished scope; do not claim execution that did not occur.

## Define and Refine the Research Contract

Record the decision purpose, audience, scope, definitions, as-of date, freshness, region and population, comparators, evaluation axes, depth, and requested format as known, assumed, or unresolved.

Do not require these to be settled before entering the loop. Clarifying the question, finding constraints, and forming acceptance criteria are valid research tasks. Ask only when missing information materially changes the answer and cannot be resolved from evidence. The main agent records changes and confirmation status under the shared loop contract.

Map subquestions, conclusion-sensitive assumptions, alternative explanations, evidence gaps, and suitable source types. Seek decision-relevant evidence, not an arbitrary source count.

## Inspect Sources and Preserve Evidence

Read [Source and Evidence Policy](references/source-policy.md) before collecting evidence.

- Use search results and snippets for discovery. Open the exact source before citing it; disclose inaccessible originals and secondary-only evidence.
- Prefer applicable primary evidence such as official documents, original papers, data, laws, standards, filings, or regulations. Distinguish self-report from independent corroboration.
- Record exact URLs or original artifact paths, recoverable locations, authoring bodies, publication or as-of dates, actual access dates, versions, and scope. Distinguish publication, revision, represented period, and access.
- Keep a claim-evidence ledger for each material falsifiable claim: fact, source interpretation, inference, or recommendation; supporting evidence; counterevidence and alternatives; source and dates; region, period, population, sample, definition, and method; evidence-family independence; corroboration status and impact; confidence with reason; and remaining uncertainty.
- Seek counterevidence and independent corroboration. Repetition of one release or dataset is one evidence family. Record unavailable corroboration and its effect.
- Explain conflicts through definitions, samples, methods, periods, regions, interests, or revisions. Weight directness, applicability, independence, method quality, and freshness rather than votes.
- Preserve verified findings and counterevidence from unselected approaches. Selecting a direction allocates effort; it does not prove other explanations false.

## Assigned Research Results

All research workers are read-only. Their packet supplies a bounded question, scope, dates, preferred sources, necessary original material, and evidence requirements under the common task contract.

Return a provisional answer, material findings, exact source locations and dates, supporting passages, counterevidence, method and scope limits, confidence and uncertainty, and unanswered items with reasons. Apply the full report format only when assigned report synthesis.

Structured extractors follow fixed fields and missing-value rules, quote ambiguous passages, and do not decide causality, final confidence, or recommendations outside their assignment. Escalate failed sample validation.

The main agent rechecks material claims and citations, resolves conflicts, and owns synthesis and the final report. In mixed research-and-action requests, keep research read-only and assign implementation to separately authorized domain branches; honor prerequisites when implementation depends on findings.

## Synthesize and Supply Completion Evidence

Read [Report Contracts](references/report-contracts.md) before drafting. Answer the question in the opening paragraph and scale detail to the requested depth.

Expose each material judgment's type, evidence, counterevidence or limitation, scope, corroboration status and impact, confidence with reason, exact source, publication or as-of date, access date, recoverable location, and uncertainty. Use a compact table or equivalent judgment sections. The full working ledger may remain separate unless requested.

Supply these checks to the common loop's completion review:

- every material subquestion is answered or has an explicit reason and impact for being unanswered;
- important claims have inspected primary evidence or a disclosed limitation;
- counterevidence, conflicts, freshness, and independent corroboration were checked;
- further collection adds repetition rather than information that could change the conclusion;
- citations support the claims and language strength matches evidence strength;
- the report serves the current user purpose and accumulated conditions.

These are domain checks, not a separate stop command. If time, allowance, access, or tools force an earlier stop, identify partial scope, unresolved questions, impact, and the most useful next check.
