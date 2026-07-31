# Model and Task Routing

Use this reference only when parallel research streams justify delegation.

## Route by Role

Define the responsibility before selecting an available model.

| Role | Suitable work | Required traits | Current family example |
| --- | --- | --- | --- |
| Research lead and synthesizer | Scope, conflict resolution, sensitive inference, final report | Strongest reasoning and long-context synthesis | Sol |
| Source-stream researcher | Retrieve, inspect, and summarize primary sources for one subquestion | Balanced reasoning, tool use, reading throughput | Terra |
| Structured extractor | Repeat fixed-field extraction across many documents | Fast processing and strict schema adherence | Luna |

Treat model names as examples. Check which models and capabilities are available, then preserve the role contract with the closest valid substitute.

## Delegation Gate

Delegate only when all conditions hold:

- at least three non-overlapping read-only research streams exist;
- each stream can start without another stream's result;
- question, scope, as-of date, preferred source types, and output schema are fixed;
- time or coverage benefit exceeds coordination and synthesis cost;
- the research lead can recheck material claims and citations.

Do not delegate:

- one- or two-lookup questions;
- subquestions that depend strongly on a shared definition or evidence base;
- work requiring external writes, messages, purchases, or change authority;
- decisions that would be resolved by worker majority.

## Research Packet

Provide each researcher only:

- one independent subquestion;
- research purpose and decision context;
- included and excluded scope;
- as-of date, region, and key definitions;
- preferred source types;
- a read-only boundary;
- completion conditions;
- the output schema below.

Do not pass another researcher's conclusion, the lead's suspicion, a preferred answer, or other narrative that could bias discovery.

## Researcher Result

Require:

1. provisional answer to the subquestion;
2. material findings;
3. exact primary-source URLs and recoverable locations;
4. publication or as-of dates and access dates;
5. evidence that directly supports each claim;
6. counterevidence or alternative explanations;
7. scope and methodological limitations;
8. confidence and remaining uncertainty;
9. unanswered items and reasons.

A claim list without URLs or a summary of search results is incomplete.

## Role Limits

### Research lead and synthesizer

- Own definitions and success criteria.
- Manage overlap and evidence gaps across streams.
- Do not average worker conclusions.
- Reopen material sources and citations.
- Own conflict resolution, language strength, and final uncertainty.

### Source-stream researcher

- Focus on one bounded subquestion and source family.
- Open primary sources instead of relying on snippets.
- Return findings and counterevidence in the required schema.
- Do not synthesize other streams or produce the final recommendation.

### Structured extractor

- Accept only tasks with fixed fields and missing-value rules.
- Flag ambiguous text with its source passage for review.
- Do not own legal interpretation, causality, conflict resolution, or final confidence.
- Stop and escalate when sample validation fails.

## Execution Order

1. The research lead fixes the research map and shared definitions.
2. Assign independent streams to source researchers.
3. Use structured extraction only for repetitive work inside a stream.
4. Validate each result's URL, dates, scope, and counterevidence.
5. Reopen material sources and merge accepted evidence into the ledger.
6. Run only the follow-up research needed to resolve conflicts or gaps.
7. Complete synthesis and the citation audit.

## Sequential Fallback

If delegation is unavailable or fails the gate, one model must execute the same stages sequentially. Keep stream notes separate until synthesis to preserve role separation.
