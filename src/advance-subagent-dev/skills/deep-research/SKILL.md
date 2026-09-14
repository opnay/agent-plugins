---
name: deep-research
description: Conduct evidence-traceable, multi-source research for consequential questions. Use for market, literature, technical, policy, competitive, due-diligence, or fact-checking work that requires primary-source retrieval, independent corroboration, counterevidence, freshness checks, explicit uncertainty, and a cited report. Do not use for simple lookups, single-source summaries, unsupported brainstorming, implementation, external writes, or ongoing monitoring.
---

# Deep Research

Investigate consequential questions across multiple sources and produce a report that traces each important claim to evidence, counterevidence, and uncertainty. Prioritize reproducibility and calibrated language over confidence.

This skill is a general research workflow. Do not claim that it reproduces OpenAI's proprietary `Deep Research` product mode.

Maintain this data flow:

`question and context → scope and success criteria → research map → source plan → discovery → primary-source review → claim-evidence ledger → conflict and freshness checks → synthesis → citation and uncertainty audit → report`

Use each stage's output as the next stage's input. Do not jump from search results directly to prose.

## 1. Lock the Question

Establish:

- the decision or use case the research supports;
- included and excluded scope and key definitions;
- the as-of date and required freshness;
- region, population, comparators, and evaluation axes;
- audience, depth, and output format.

Ask only when missing information could materially change the conclusion. Otherwise state a reasonable assumption and proceed. Define success as both the knowledge required to answer and the evidence required to trust that answer.

Do not use this workflow for a single lookup, a summary of one user-provided document, unsupported ideation, implementation, external changes, or monitoring. For high-stakes medical, legal, financial, or safety questions, apply the workflow when independent evidence review is required.

If a request mixes research with implementation, finish the research deliverable first. Treat implementation as separate work governed by its own permissions and instructions.

## 2. Build the Research Map

Split the question into non-overlapping subquestions. Record:

- required facts, definitions, comparisons, causal claims, or forecasts;
- assumptions that could reverse the conclusion;
- plausible counter-explanations and expected evidence gaps;
- the best source types for each subquestion.

Do not target an arbitrary source count. Gather enough direct, independent, and current evidence to support the material claims.

Read and apply [Source and Evidence Policy](references/source-policy.md) before collecting evidence.

## 3. Retrieve and Inspect Primary Sources

- Use search results and snippets only to discover candidates.
- Open the exact webpage, document, paper, dataset, or PDF before citing it.
- Verify material claims in primary sources such as official documents, original papers, original data, standards, filings, or regulations when possible.
- Distinguish a primary source's self-report from independent corroboration.
- Record publication or as-of date, access date, authoring body, scope, and a recoverable citation location.
- Never imply access to a source that could not be opened or inspected.

Use current browsing or retrieval tools when freshness matters. If tools or sources are unavailable, disclose the access limit and its effect instead of filling the gap from memory.

## 4. Maintain a Claim-Evidence Ledger

Track at least:

| Field | Required content |
| --- | --- |
| Claim | One falsifiable statement |
| Type | Fact / interpretation / inference / recommendation |
| Supporting evidence | What the source directly supports |
| Counterevidence | Contradiction, exception, or alternative explanation |
| Source | Exact URL and recoverable location |
| Dates | Publication or as-of date and access date |
| Scope | Region, period, population, sample, and definition |
| Independence | Whether it is materially independent of other evidence |
| Corroboration | Complete / unavailable / insufficient, plus impact |
| Confidence | High / medium / low, with reason |
| Remaining uncertainty | What remains unverified |

Repeated reporting of the same press release or dataset is one evidence family, not independent corroboration.

## 5. Challenge the Evidence

For every material conclusion:

- seek counterevidence or a simpler explanation;
- corroborate important claims with an independent source;
- record why corroboration is unavailable and how that affects the conclusion;
- explain conflicting numbers through definitions, samples, periods, regions, methods, incentives, or revision dates;
- separate source facts, source interpretations, your inferences, and recommendations.

Do not decide by source majority. Give more weight to evidence that is more direct, applicable, methodologically sound, independent, and current, and explain the weighting.

## 6. Delegate Only When Valuable

Delegate only when there are at least three non-overlapping, independently startable, read-only research streams with fixed questions, scopes, and output schemas. Do not delegate simple lookups or short investigations.

Before delegating, read [Model and Task Routing](references/model-routing.md). Route by responsibility and verification needs, not model name alone. The lead researcher must reopen the important sources and verify material claims and citations.

If delegation tools are unavailable or the streams are not independent, execute the same contract sequentially.

## 7. Synthesize and Report

Read [Report Contracts](references/report-contracts.md) before drafting.

- Answer the question in the opening paragraph.
- Place evidence and exact citations close to each material claim.
- Present counterevidence and conflicts with the conclusion, not as an afterthought.
- Match claim strength to evidence strength.
- State what remains unknown and what evidence could change the conclusion.
- Scale report length to the requested depth.

The full working ledger may remain internal. The report must still expose, for every material claim, its type, evidence, counterevidence or limitation, scope, corroboration status, confidence, exact URL, publication or as-of date, access date, citation location, and remaining uncertainty. Provide the full ledger as an appendix when requested.

## Stop Conditions

Stop when all of the following hold:

- each material subquestion is answered or has an explicit reason it cannot be answered;
- every important claim has primary evidence or a disclosed limitation;
- new sources add repetition rather than decision-changing information;
- counterevidence, conflicts, and freshness risks have been checked;
- citations support the claims and language strength matches the evidence.

If time, access, or tooling forces an earlier stop, report the inspected scope, unresolved areas, and the highest-value next research step.

## Final Audit

Before responding, confirm:

- every cited source was opened and inspected;
- each material claim has a source, date, and scope;
- no search snippet, memory, or secondary repetition is presented as direct evidence;
- counterevidence and conflicting values are visible;
- facts, interpretations, inferences, and recommendations remain distinguishable;
- uncertainty and conclusion-changing conditions are explicit;
- the report directly serves the question and decision purpose.

Revise any claim or conclusion that fails this audit.
