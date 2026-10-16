---
name: deep-research
description: Conduct evidence-traceable, multi-source research for consequential questions, including parallel exploration of alternative hypotheses and fresh-context continuation when valuable. Use for market, literature, technical, policy, competitive, due-diligence, or fact-checking reports requiring primary sources, corroboration, counterevidence, freshness, and explicit uncertainty. Do not use for simple lookups, single-source summaries, unsupported brainstorming, implementation, external writes, or monitoring.
---

# Deep Research

Answer consequential research questions with an auditable report connecting claims, sources, counterevidence, and uncertainty. This is a general research workflow, not OpenAI's proprietary Deep Research product mode.

Follow this flow:

`question → scope and success criteria → research map → source plan → source inspection → claim-evidence ledger → conflict and freshness checks → synthesis → citation audit → report`

When competing approaches merit exploration, add:

`preserved evidence → fresh parallel approaches → source verification → next-direction selection → minimal handoff → fresh continuation`

## 1. Frame the Research

Fix the decision purpose, audience, included and excluded scope, definitions, as-of date and freshness needs, region and population, comparators, evaluation axes, depth, and output format. Define both the knowledge needed to answer and the evidence needed to trust the answer. Ask only about missing information that could materially change the result; otherwise disclose reasonable assumptions.

Map independently answerable subquestions, conclusion-sensitive assumptions, alternative explanations, expected evidence gaps, and suitable source types. Seek enough direct, independent, current evidence for material claims, not an arbitrary source count.

Single lookups, one-document summaries, unsupported ideation, implementation, external changes, and monitoring are outside this skill. For medical, legal, financial, or safety questions, use it when independent evidence review is needed. In mixed research-and-implementation requests, finish the research deliverable before separately authorized implementation.

## 2. Inspect Sources and Preserve Evidence

Read [Source and Evidence Policy](references/source-policy.md) before collection.

- Use snippets and search results for discovery. Open the exact source before citing it; never imply access to unavailable material.
- Prefer applicable primary evidence: official documents, original papers, data, standards, filings, or regulations. Distinguish self-report from independent corroboration and disclose reliance on secondary evidence.
- Record exact URLs, recoverable citation locations, authoring bodies, publication or as-of dates, actual access dates, versions where relevant, and scope. Use current retrieval when freshness matters; disclose tool and access limits.
- Maintain a claim-evidence ledger containing each falsifiable claim; fact, interpretation, inference, or recommendation type; supporting evidence; counterevidence or alternative explanation; source and dates; region, period, population, sample, and definition; evidence-family independence; corroboration status and impact; confidence with reasons; and remaining uncertainty.
- Seek counterevidence and independently corroborate important claims. Repeated reports of one release or dataset count as one evidence family. Explain missing corroboration and its effect.
- Resolve conflicts through definitions, samples, methods, periods, regions, interests, or revisions. Separate observed facts, source interpretations, your inferences, and recommendations. Weight directness, applicability, sound methods, independence, and freshness rather than source or worker majority.

## 3. Choose a Research Execution Shape

Read [Model and Task Routing](references/model-routing.md) before delegation. All researchers are read-only. Fix their questions, scope, dates, expected outputs, and completion conditions; delegate only when research benefit exceeds coordination cost and the lead can recheck material sources.

- **Divide subquestions:** Delegate when at least three non-overlapping research streams can start independently.
- **Explore alternatives:** For the same question, use two fresh researchers by default when distinct hypotheses, methods, or discovery paths are worth comparing. The three-stream rule does not apply to this mode. Give the same goal, constraints, and comparison criteria, with different approaches to investigate rather than conclusions to defend.

Do not delegate simple lookups or short investigations. If delegation is unavailable or not worthwhile, perform the same evidence workflow sequentially and disclose any material limitation. Sequential work or a summary inside the current session is not a fresh-context parallel run.

## 4. Explore, Select, and Continue

Use Fork-and-Flush as an internal research strategy when alternative approaches can materially reduce uncertainty.

1. Preserve existing sources, artifacts, and the ledger. Fix the goal, user constraints, authority, completion criteria, common comparison axes, per-candidate exploration allowance, and overall round limit.
2. Start fresh researchers without inherited conversation. Supply the question, necessary original sources and verified evidence, constraints, criteria, and output contract. Do not supply prior worker outputs, the lead's preferred answer, speculative conclusions as facts, or private reasoning. Each approach must be able to challenge its starting hypothesis.
3. Reopen material sources and verify claims. Compare approaches by evidence directness, applicability, independence, freshness, unresolved uncertainty, and cost of further investigation. Choose where to spend the next research effort, not which account must be true.
4. Retain verified evidence and counterevidence from every approach in the shared ledger, including unselected approaches. Never discard inconvenient findings or decide truth by votes, model confidence, or candidate ranking.
5. Give a fresh continuation researcher a minimal handoff: goal, constraints, authority, completion criteria; preserved sources and ledger locations with verification status; chosen direction and verified rationale; conflicts and unknowns; remaining allowance and next action. They inspect the preserved material and continue the selected investigation. The lead retains responsibility for source checks, synthesis, and the report.
6. Preserve the baseline if no approach offers a better direction. If all fail or decisive evidence is missing, defer selection and identify the next useful check or unresolved limit. Fork again only with a new research question and remaining allowance.

Use new agents without conversation inheritance for exploration and continuation. Reusing or interrupting an old worker is not a context reset. Do not delete the current conversation or workspace. If fresh execution is unavailable, continue sequentially within scope and report that fresh-context continuation was not performed.

## 5. Ask Action-Changing Decision Questions

Use a decision model such as Jev when a bounded, uncertain judgment changes the next research action. Check deterministic facts such as source access, dates, artifact existence, and remaining allowance directly. Ask only the questions needed now:

| Decision | Question | Example choices |
| --- | --- | --- |
| Start exploration | Is continuing the current investigation or exploring alternative approaches more valuable for reducing uncertainty? | `continue_current`, `fork`, `inspect_more` |
| Allocate effort | Which approach warrants the next research effort given verified evidence, gaps, and cost? | Actual candidate IDs, `keep_baseline`, `insufficient_evidence` |
| Resolve uncertainty | What affordable check could change the choice? | Concrete research actions, `explore_alternative`, `ask_user` |
| Next round | Should we deepen the selected direction, fork again, or end research given current evidence and remaining allowance? | `continue_selected`, `fork_again`, `end_research` |

Supply the goal, constraints, authority, completion criteria, verified evidence, conflicts, gaps, and remaining allowance. Request the action, a short evidence-linked rationale, and unknowns that could change it. Do not send the full conversation or a preferred answer.

Treat the response as advice. It cannot invent facts, determine truth, supply user preferences, grant permission, or waive stopping criteria. The lead checks evidence, authority, and completion before acting. Ask dependent questions sequentially; batch only independent questions sharing the same state.

No service or CLI is required. Send only authorized, necessary information to an external model. If unavailable or unsuccessful, apply the same criteria locally and disclose the fallback rather than installing tools or changing settings.

## 6. Synthesize, Audit, and Report

Read [Report Contracts](references/report-contracts.md) before drafting. Answer directly in the opening paragraph, scale length to the requested depth, and place exact citations close to material claims. Make counterevidence and conflicts visible alongside conclusions.

The report must expose each material claim's type, evidence, counterevidence or limitation, scope, corroboration status and impact, confidence with reason, exact source URL, publication or as-of date, access date, recoverable location, and remaining uncertainty. A compact table or equivalent judgment sections may expose these fields; the full working ledger can remain internal unless requested.

Before reporting completion, confirm that:

- every material subquestion is answered or has an explicit reason it cannot be answered;
- every important claim has inspected primary evidence or a disclosed limitation;
- counterevidence, conflicts, freshness, and independent corroboration have been checked;
- further sources add repetition rather than decision-changing information;
- citations support the claims, language strength matches evidence strength, and the report serves the user's decision.

An `end_research` recommendation does not establish completion. If allowance, time, source access, or tools force an earlier stop, report partial scope, unresolved questions, their impact, and the highest-value next research step. Distinguish source facts, interpretations, inferences, and recommendations, and state what could change the conclusion.
