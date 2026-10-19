# Built-In Jev Use Cases

These examples apply Jev's basic judgment patterns. The current task and applicable instructions define execution, permission, and model policy. Use only the example relevant to the pending judgment.

## Small Choices

For an A/B question, supply the user's goal, known constraints, evidence, and comparison criteria. Include `need_user_input` when an unstated preference could decide the answer. Check the recommendation against the evidence before presenting it.

```sh
jev choice --if "Goal: ...; confirmed facts: ...; constraints: ...; comparison criterion: ...; which option fits?" \
  --conditions option_a,option_b,need_user_input \
  --json --pick answer,probabilities
```

## Before Asking The User

When current context may suffice, compare `proceed_with_known_default`, `inspect_more`, and `ask_user`. First inspect readily available evidence. Ask the user directly when only they can provide a preference, approval, or materially different scope; Jev cannot provide these on their behalf. Skip Jev when deterministic rules or existing evidence settle the answer.

## Uncertain Evidence

Keep inspected source locations, the claim, short excerpts, and known gaps in one authorized state. Check literal presence, dates, paths, and artifact existence directly before asking Jev about relevance or interpretation.

Combine distinct broad questions about overall support, relevance, or alternative interpretations with narrow questions about material subclaims, counterexamples, and missing premises. Include `insufficient` or `not_shown` to distinguish missing evidence from contradiction. Batch only questions answerable independently from that exact state.

For a release-readiness claim with local validation and deployment observations:

```jsonl
{"id":"overall_relation","type":"choice","question":"How does the evidence as a whole relate to the release-readiness claim?","conditions":["supported","contradicted","insufficient"]}
{"id":"alternative","type":"choice","question":"Which alternative interpretation remains plausible?","conditions":["local_only","deployment_unverified","no_material_alternative"]}
{"id":"local_checks","type":"choice","question":"Does the evidence show checks after the latest change?","conditions":["supported","contradicted","not_shown"]}
{"id":"deployment","type":"choice","question":"Does the evidence show the current deployment was verified?","conditions":["supported","contradicted","not_shown"]}
```

Run `jev batch --state-file evidence.md --input questions.jsonl`. Add distinct action-relevant questions as needed rather than imposing a small fixed count. Check current request and context limits in [TypeSafe's model documentation](https://docs.typesafe.ai/models) before large requests. Split questions with different states; ask dependent follow-ups after inspecting earlier answers.

Compare answers by ID. A conflict, `insufficient`, or `not_shown` identifies the next source check. Do not average answers, count votes, or treat repeated paraphrases as independent support. Return to original evidence; Jev advice is not a factual verdict.

## Model Selection

Model selection illustrates permitted-candidate construction and dependent questions. It does not establish an orchestration workflow or authorize spawning agents.

1. Inspect the task's minimum scope, research and execution effort, independent work, failure impact, and authority boundaries. Use applicable defaults when they settle the choice.
2. If the next action is still uncertain, Jev can compare `direct`, `delegate`, `inspect_more`, and `ask_user` using those facts. Check the recommendation against current delegation rules before considering models.
3. Build candidates from the current instructions and available tool capabilities:
   - If policy defines model-and-effort pairs as profiles, ask Choice over permitted profiles and preserve each pair intact.
   - If policy permits individual selection, ask Choice over currently available, permitted model IDs. After selecting a model, ask a separate Choice question over its supported, permitted efforts. Include the assigned work, applicable defaults, and the criterion of sufficient capability at reasonable cost. Do not batch these dependent questions.
4. Check the recommendation against the task, policy, and tools. Model selection is advice for this task, not a saved setting or new execution authority. The owning task workflow handles decomposition, agent execution, and integration.

If Jev cannot help, follow an applicable default or local judgment within existing authority and disclose a material fallback. Do not invent model IDs, replace mandatory profiles, or alter settings to make the example executable.
