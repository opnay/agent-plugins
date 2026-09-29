---
name: jev-use
description: Use Jev for bounded choices in user questions or task execution, mixed broad and narrow judgments on uncertain evidence, and model and effort routing after checking underspecified work. quick choice, ask user or proceed, uncertain evidence, subagent model and effort
---

# Use Jev During Work

Use hosted Jev when a bounded judgment can change the next action. First check the request, applicable instructions, and readily available evidence. Act directly when they settle the answer. Jev receives the question, state, labels, and levels through TypeSafe; send only authorized, necessary content.

Jev is advisory. It cannot supply missing facts, choose a user preference, grant permission, or prove correctness. Keep the original task and acceptance conditions in force.

## Frame The Judgment

State the goal, confirmed facts, material unknowns, constraints, candidates, and decision criterion. Inspect locally when a quick check could resolve the uncertainty. Include a real recovery option such as `inspect_more` or `ask_user` when applicable.

- Use `jev choice` for a discrete next action or classification.
- Use `jev noul` for one yes/no proposition. Its answer is the probability of yes, not a verdict.
- Use `jev score` for ordered descriptive levels.
- Use `jev batch` for independent questions sharing exactly the same nonempty state. A later question that depends on an earlier answer needs a later call; questions requiring different evidence states need separate calls.

```sh
jev choice --if "Goal: ...; confirmed context: ...; constraints: ...; which next action best fits?" \
  --conditions direct,delegate,inspect_more,ask_user \
  --json --pick answer,probabilities
```

Use `$jev:jev` for full CLI, configuration, and output details when needed. This skill contains the decisions needed to use its examples without that sibling.

## Examine Uncertain Evidence

Use this pattern when identified source material exists but its relevance, support for a claim, missing premise, or counterevidence remains unclear and the answer affects the work. Keep provenance, the claim, short evidence excerpts, and known gaps in one authorized shared state. Check literal presence, dates, paths, and other deterministic facts with ordinary tools first.

- Ask several distinct broad questions when useful: overall relevance, relationship to the claim, and plausible alternative interpretations. Include `insufficient` or `not_shown` where the evidence cannot settle a question.
- Ask narrow questions for each material subclaim, counterexample, and missing premise. Phrase them so a negative answer does not turn absence of evidence into proof of the opposite.
- Batch broad and narrow questions together when each can be answered independently from that same state. Do not cap the set at two or three questions: include as many distinct, action-relevant questions as needed within the current API request and context limits and reasonable transmission cost. More paraphrases do not create independent evidence.

For example, with one state containing a release-readiness claim, observed local checks, and deployment observations:

```jsonl
{"id":"overall_relation","type":"choice","question":"How does the evidence as a whole relate to the release-readiness claim?","conditions":["supported","contradicted","insufficient"]}
{"id":"alternative","type":"choice","question":"Which alternative interpretation of the evidence remains plausible?","conditions":["local_only","deployment_unverified","no_material_alternative"]}
{"id":"local_checks","type":"choice","question":"Does the evidence show checks after the latest change?","conditions":["supported","contradicted","not_shown"]}
{"id":"deployment","type":"choice","question":"Does the evidence show the current deployment was verified?","conditions":["supported","contradicted","not_shown"]}
```

Run `jev batch --state-file evidence.md --input questions.jsonl` only if every row uses that exact state. Add rows for other decisive claims or alternatives, rather than a fixed question count. The CLI has no fixed row-count cap. [TypeSafe's model documentation](https://docs.typesafe.ai/models) currently lists 64k tokens per request and 32k for the state plus longest question; check current limits before a large batch. The CLI submits one request and emits complete result rows in input order; it does not chunk, retry, resume, or apply single-result thresholds.

Compare the broad answers with the narrow answers by ID. Do not average them or take a majority vote. A conflict, `insufficient`, or `not_shown` identifies what to inspect next in the original source. Jev responses are leads for source review, not a final evidence verdict.

## Small Decisions And User Questions

For a bounded A/B question, include the user goal and stated criteria and consider `option_a`, `option_b`, and `need_user_input`. Check the recommendation against the evidence. Ask the user when the distinction depends on their unstated preference.

Before a routine clarification, consider `proceed_with_known_default`, `inspect_more`, and `ask_user`. Ask directly when only the user can grant approval, set a preference, or define materially different scope. Skip Jev for trivial or deterministic questions and answers already present in context.

## Route Underspecified Work

1. Establish the smallest task-scoped baseline from the request and environment. Estimate necessary research, action effort, independent work, failure impact, and authority gaps.
2. If a bounded judgment remains, ask Jev to choose `direct`, `delegate`, `inspect_more`, or `ask_user`. Apply repository and tool rules before delegating.
3. If delegation is warranted and allowed, ask a separate Choice question over currently available model IDs. Include task demands, user model limits, and repository defaults. Exclude disallowed models.
4. After selecting a model, ask a separate Choice question over efforts supported by that model and allowed by the user and repository. Include the assigned subtask and applicable default; choose the least effort sufficient for it. Do not batch model and effort selection because the latter depends on the former.
5. Check the recommendations against current instructions and tool capabilities, then act. A model or effort choice applies to this task, not saved settings. Use the applicable default when Jev cannot help and explain a material fallback.

## Results And Failure

On single-command exit `0`, consume the answer once. Exit `2` means a configured result criterion was not met and stdout is empty. Exit `1` means input, configuration, authentication, API, response, or output failure. Batch exit `0` means every row returned; exit `1` means failure without a complete result. Do not lower a criterion, repeat the same call, install the executable, or change credentials automatically.

Continue from known rules and evidence when possible; ask the user only for a decision they must make. Report whether Jev answered or a local fallback was used. Confirm any selected action against the source, instructions, and execution authority before acting.
