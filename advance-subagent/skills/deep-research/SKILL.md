---
name: deep-research
description: Conduct evidence-traceable, multi-source research for consequential questions, including parallel exploration of alternative hypotheses and fresh-context continuation when valuable. Use for market, literature, technical, policy, competitive, due-diligence, or fact-checking reports requiring primary sources, corroboration, counterevidence, freshness, and explicit uncertainty. Do not use for simple lookups, single-source summaries, unsupported brainstorming, implementation, external writes, or monitoring.
---

# Deep Research

Answer consequential questions with an auditable report connecting claims, sources, counterevidence, and uncertainty. This is a general research workflow, not OpenAI's proprietary Deep Research product mode.

Follow this flow:

`question → scope and success criteria → research map → source plan → source inspection → claim-evidence ledger → conflict and freshness checks → synthesis → citation audit → report`

When alternative approaches merit exploration, add:

`preserved evidence → fresh parallel approaches → source verification → next-direction selection → minimal handoff → fresh continuation`

## 1. Frame the Research

Fix the decision purpose, audience, included and excluded scope, definitions, as-of date and freshness needs, region and population, comparators, evaluation axes, depth, and output format. Define the knowledge needed to answer and the evidence needed to trust it. Ask only about missing information that could materially change the answer; otherwise disclose reasonable assumptions.

Map independent subquestions, conclusion-sensitive assumptions, alternative explanations, expected gaps, and suitable source types. Seek sufficient direct, independent, current evidence, not an arbitrary source count.

Single lookups, one-document summaries, unsupported ideation, implementation, external changes, and monitoring are outside this skill. For medical, legal, financial, or safety questions, use it when independent evidence review is needed. In mixed requests, complete the research deliverable before separately authorized implementation.

## 2. Inspect Sources and Keep Evidence

Read [Source and Evidence Policy](references/source-policy.md) before collecting evidence.

- Use search results and snippets for discovery. Open the exact source before citing it; never imply access to unavailable material.
- Prefer applicable primary evidence: official documents, original papers, data, standards, filings, or regulations. Distinguish self-report from independent corroboration and disclose secondary-only evidence.
- Record exact URLs, recoverable locations, authoring bodies, publication or as-of dates, actual access dates, relevant versions, and scope. Use current retrieval when freshness matters; disclose tool and access limits.
- Keep a claim-evidence ledger with each falsifiable claim; fact, interpretation, inference, or recommendation type; supporting evidence; counterevidence or alternatives; source and dates; region, period, population, sample, and definition; evidence-family independence; corroboration status and impact; confidence with reason; and remaining uncertainty.
- Seek counterevidence and independent corroboration for important claims. Repeated reporting of one release or dataset is one evidence family. Explain unavailable corroboration and its impact.
- Explain conflicts through definitions, samples, methods, periods, regions, interests, or revisions. Separate observed facts, source interpretations, your inferences, and recommendations. Weight directness, applicability, method quality, independence, and freshness rather than source or worker majority.

## 3. Choose Execution and Profiles

Read [Model Profiles and Research Routing](references/model-routing.md) before delegation. All researchers are read-only. Fix questions, scope, dates, expected outputs, and completion conditions. Delegate only when research benefit exceeds coordination cost and the lead can recheck material sources.

- **Divide subquestions:** Delegate at least three non-overlapping streams that can start independently.
- **Explore alternatives:** For the same question, use two fresh researchers by default when distinct hypotheses, methods, or discovery paths are worth comparing. The three-stream rule does not apply. Share goal, constraints, and criteria, but assign approaches to investigate rather than conclusions to defend.

Choose only `fast`, `default`, or `high`, separately from the research role. Applicable instructions define exact model-and-effort pairs for `fast` and `high`. For `default`, omit both spawn settings to inherit. Follow explicit profile choices and defaults first; otherwise use task demands, uncertainty, verification, and cost. Narrow directly verifiable work suits `fast`; difficult inference or consequential conflicts may suit `high`; otherwise inherit. Length alone does not justify `high`.

Keep each pair intact. Do not mix profiles, choose effort independently, or invent combinations. Match combined shorthand to a supported model ID and effort before spawning. If a profile is undefined or unsupported, use another instruction-permitted profile, inheritance, or direct work only when the role can be met. Do not replace a mandatory profile without authorization. Report when no valid route remains. Record profile, selection reason, and exact pair or inheritance without guessing hidden inherited values.

Do not delegate simple lookups or short investigations. If delegation is unavailable or not worthwhile, perform the same evidence workflow sequentially and disclose material limitations. Sequential work or a summary in the current session is not a fresh-context parallel run.

## 4. Explore, Select, and Continue

Use Fork-and-Flush as an internal strategy when alternatives can materially reduce uncertainty.

1. Preserve existing sources, artifacts, and the ledger. Fix goal, user constraints, authority, completion criteria, common comparison axes, per-candidate allowance, and overall round limit.
2. Start fresh researchers without inherited conversation. Give the question, necessary original sources and verified evidence, constraints, criteria, profile selection, and output contract. Do not supply prior worker outputs, the lead's preferred answer, speculative conclusions as facts, or private reasoning. Each approach must be able to challenge its starting hypothesis.
3. Reopen material sources and verify claims. Compare approaches by directness, applicability, independence, freshness, unresolved uncertainty, and further investigation cost. Choose where to spend the next research effort, not which account must be true.
4. Keep all verified evidence and counterevidence in the shared ledger, including unselected approaches. Do not discard inconvenient findings or decide truth by votes, model confidence, or ranking.
5. Give a fresh continuation researcher a minimal handoff: goal, constraints, authority, completion criteria; preserved source and ledger locations with verification status; selected direction and verified rationale; conflicts and unknowns; remaining allowance and next action. Select its profile under the same policy. It inspects preserved material and continues; the lead still owns source checks, synthesis, and the report.
6. Keep the baseline if no approach improves the direction. If all fail or decisive evidence is missing, defer selection and identify the next useful check or unresolved limit. Fork again only with a new question and remaining allowance.

Use new agents without inherited conversation for exploration and continuation. Reusing or interrupting an old worker is not a reset. Do not delete the current conversation or workspace. If fresh execution is unavailable, continue sequentially within scope and disclose the limitation.

## 5. Ask Action-Changing Decision Questions

Use a decision model such as Jev when a bounded, uncertain choice changes the next action. Check source access, dates, artifact existence, and remaining allowance directly. Ask only needed questions:

| Decision | Question | Example choices |
| --- | --- | --- |
| Start exploration | Is current investigation or alternative exploration more valuable for reducing uncertainty? | `continue_current`, `fork`, `inspect_more` |
| Allocate effort | Which approach warrants further research given verified evidence, gaps, and cost? | Candidate IDs, `keep_baseline`, `insufficient_evidence` |
| Resolve uncertainty | What affordable check could change the choice? | Concrete research actions, `explore_alternative`, `ask_user` |
| Next round | Should we deepen the selected direction, fork again, or end research? | `continue_selected`, `fork_again`, `end_research` |

Supply goal, constraints, authority, completion criteria, verified evidence, conflicts, gaps, and remaining allowance. Request the action, short evidence-linked rationale, and unknowns that could change it. Do not send the full conversation or a preferred answer. For model-profile advice, offer only defined, available profiles; never ask the judge to invent a model/effort pair.

Treat responses as advice. They cannot invent facts, determine truth, supply user preferences, grant permission, or waive stopping criteria. The lead checks evidence, authority, and completion before acting. Ask dependent questions sequentially; batch only independent questions with the same state.

No service or CLI is required. Send only authorized necessary information externally. If the decision model is unavailable or fails, apply the same criteria locally and disclose the fallback. Do not install tools or edit settings.

## 6. Synthesize, Audit, and Report

Read [Report Contracts](references/report-contracts.md) before drafting. Answer directly in the opening paragraph, scale length to the requested depth, and place exact citations near material claims. Present counterevidence and conflicts alongside conclusions.

Expose each material claim's type, evidence, counterevidence or limitation, scope, corroboration status and impact, confidence with reason, exact URL, publication or as-of date, access date, recoverable location, and uncertainty. Use a compact table or equivalent judgment sections. The full working ledger can remain internal unless requested.

Before claiming completion, confirm:

- every material subquestion is answered or has an explicit reason it cannot be answered;
- every important claim has inspected primary evidence or a disclosed limitation;
- counterevidence, conflicts, freshness, and independent corroboration have been checked;
- further sources add repetition rather than decision-changing information;
- citations support claims, language strength matches evidence strength, and the report serves the user's decision.

An `end_research` recommendation does not establish completion. If allowance, time, access, or tools force an earlier stop, report partial scope, unresolved questions, impact, and the highest-value next step. Distinguish facts, interpretations, inferences, and recommendations, and state what could change the conclusion.
