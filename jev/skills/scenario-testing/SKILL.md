---
name: scenario-testing
description: Predict and empirically evaluate how fresh agents act after reading reusable instructions by using fixed scenarios, scenario-specific behavior choices, optional Jev prediction, and expected-predicted-actual comparison. Use after creating or heavily changing a reusable instruction, or when repeated behavior suggests instruction ambiguity.
---

# Scenario Testing

Evaluate the behavior an instruction causes, not whether its prose sounds good.

Freeze each scenario, valid fixture, observable behavior choices, hidden expected choice, and checks before prediction or execution. Optionally use Jev to predict an action. Then run a fresh executor, classify its observed action with the same choices, and derive the verdict yourself.

This skill owns evaluation. It does not create, edit, or patch the target instruction.

## Use This Workflow

Use it for reusable agent-facing text such as skills, commands, task prompts, AGENTS or CLAUDE sections, and code-generation prompts.

Do not use it for disposable prompt polishing, known implementation defects, or tool failures unrelated to the instruction. If fresh executors are unavailable, report `structural review only` or `evaluation skipped`; do not claim an empirical evaluation.

Use an explicit user-provided scenario count. Otherwise use 40 scenarios, or 2-3 for an explicit smoke check. Keep one target and analysis theme per pass.

Follow this order:

1. Check static alignment between the target description and body.
2. Freeze the scenario catalog.
3. Validate each fixture.
4. Define scenario-specific behavior choices.
5. Lock the hidden `expected_choice`.
6. If authorized, ask Jev for `predicted_choice`.
7. Run a fresh executor.
8. Classify `actual_choice` from observed evidence.
9. Compare expected, predicted, and actual behavior.
10. Derive the caller verdict and save the evidence.

## Freeze And Validate Scenarios

Each scenario must define:

- stable ID and user prompt
- initial fixture state
- allowed mutation and forbidden expansion
- observable end state and completion condition
- critical and noncritical checks

Verify that the fixture creates the declared state and exposes enough evidence to classify behavior. Mark a missing, contradictory, or drifted setup as `invalid_fixture`; exclude it from instruction accuracy.

Use isolated fixtures for mutations. Never run destructive, publishing, or externally visible scenarios against the user's real workspace unless that exact mutation is authorized.

## Design Behavior Choices

Create a separate Choice for each scenario's decision point.

- Use 2-6 mutually exclusive observable action labels, then append `unexpected_action` as the final common label.
- Describe evidence that selects each label.
- Name actions, not quality judgments. Do not use `pass`, `partial`, or `fail` as behavior labels.
- Keep outcome states out of the Choice.

Example:

```text
commit_target_only
block_for_scope_conflict
commit_target_and_unrelated_changes
unexpected_action
```

Keep these separate from behavior:

- `insufficient_observation`: evidence cannot identify the action
- `invalid_fixture`: the test input is invalid or drifted
- `execution_environment_error`: the executor could not run because the environment failed

Set `expected_choice` before prediction or execution. Keep it out of Jev input and executor prompts.

## Use Optional Jev Prediction

Jev is an external predictive judge, not the verdict owner. Use it only when the user authorized Jev and external transmission of the sanitized evaluation input.

For one valid scenario, use a single Choice call:

```sh
jev choice \
  --if "<required target instruction, scenario, sanitized fixture, and label criteria>" \
  --conditions <scenario_labels_including_unexpected_action> \
  --threshold 0 \
  --json --pick answer,probabilities
```

For two or more valid scenarios sharing the same target instruction, use one batch request:

```sh
jev batch --state-file <sanitized-target> --input <scenario-questions.jsonl>
```

Each row uses the scenario ID as `id`, `choice` as `type`, the scenario prompt, sanitized fixture, and label criteria as `question`, and the scenario behavior labels plus `unexpected_action` as `conditions`.

Batch rules:

- Put only shared target instruction in the state file.
- Keep scenario-specific facts in the corresponding row.
- Do not combine scenarios that require different target state.
- Do not pass threshold or pick; batch always returns complete JSONL results.
- Do not expect CLI chunking, retry, resume, or partial success.

For every Jev call:

- Remove secrets, user identifiers, local absolute paths, unnecessary repository state, `expected_choice`, and caller conclusions.
- Preserve the request, labels, timestamp, exit code, stdout, stderr, parsed answer, and probabilities.
- Do not repeat a successful call for the same frozen input.
- Treat exit 1 or 2, or malformed output as prediction error, not behavior.
- If transmission is not authorized, record `prediction_skipped_no_authority` and continue.
- If Jev is unavailable, record `prediction_unavailable` and continue.

## Run Fresh Executors

Use a new subagent for every empirical run. Do not substitute author rereading or reuse an executor from another pass.

Dispatch large scenario sets in executable groups without changing frozen scenarios or Choices. If the executor limit is reached, close completed or unused executors and continue the remaining groups with stable scenario IDs.

Give the executor only:

- target instruction text or exact readable path
- scenario prompt and fixture
- allowed and forbidden operations
- observable checks and required evidence format

Do not reveal `expected_choice`, `predicted_choice`, suspected defects, or a preferred conclusion.

Require execution summary, commands or tools, before and after state, observed facts, ambiguity, judgment calls, retry causes, `tool_uses`, and `duration_ms`. Use `null` for unavailable metrics.

Classify `actual_choice` after execution. Use `insufficient_observation` instead of guessing.

## Attribute Unexpected Actions

Do not automatically fail `actual_choice = unexpected_action`. Attribute it first:

- `choice_set_incomplete`: valid behavior was absent from the prepared Choice; repair and rerun or reclassify only when comparability remains.
- `skill_behavior_underspecified`: the instruction permits or encourages another behavior; treat it as a rewrite candidate.
- `executor_deviation`: a one-off action is not explained by the instruction or fixture; rerun with a fresh executor.
- `fixture_drift`: runtime state differs from the frozen fixture; use `invalid_fixture`.
- `unresolved_attribution`: current evidence cannot distinguish the cause; record missing evidence.

Repeated predicted and actual `unexpected_action` strongly signals `skill_behavior_underspecified`. Any unauthorized or destructive critical behavior is `critical_fail` regardless of attribution.

## Compare And Derive Verdicts

Interpret the three choices without letting prediction determine the verdict:

- `E=P=A`: expected behavior is clear and stable.
- `E=P`, actual differs: inspect executor instability or execution conditions.
- `P=A`, expected differs: strong instruction-defect signal.
- `E=A`, prediction differs: Jev prediction or input defect.
- all differ: redesign the scenario, Choice, or fixture before a broad conclusion.

Apply verdicts in this order:

1. Unauthorized or destructive observed action -> `critical_fail`.
2. Invalid or drifted fixture -> `invalid_fixture`.
3. Executor environment failure -> `execution_environment_error`.
4. Unclassifiable observation -> `insufficient_observation`.
5. Any other critical contract violation -> `critical_fail`.
6. `unexpected_action` -> `invalid_choice_set`, `fail`, `rerun_required`, `invalid_fixture`, or `needs_attribution` from its attribution.
7. `actual_choice != expected_choice` -> `fail`.
8. Expected action with noncritical omission or documented heavy inference -> `partial`.
9. Expected action with all checks satisfied -> `pass`.

`predicted_choice` never changes the verdict.

## Preserve Evidence

Save empirical evaluation in a user-specified destination or task-scoped durable artifact. Do not put scenario execution artifacts in a plugin `changes/` directory unless the user explicitly requests inclusion in change history.

Keep one scenario-keyed evidence set:

- frozen scenario catalog with fixtures, labels, hidden expectations, and checks
- executor evidence and actual choices
- sanitized Jev single or batch inputs and results
- expected-predicted-actual comparison with state, attribution, verdict, and reason
- summary with denominators, exclusions, instruction findings, evaluation-design findings, and residual risk

Use CSV or JSONL plus a resumable executor harness for large evaluations. Preserve stable scenario IDs. The Jev batch call itself is atomic and not resumable.

Report:

- expected-actual accuracy over valid, executable, classifiable, resolved scenarios; exclude `invalid_choice_set`, `rerun_required`, and `needs_attribution`
- predicted-actual accuracy and expected-predicted agreement where Jev succeeded and actual behavior was finalized against the same frozen Choice
- verdict, relation, and unexpected-attribution distributions
- invalid fixtures, insufficient observations, environment errors, prediction errors, and prediction skips
- critical failures and repeated ambiguity, judgment, or retry causes

## Review Before Reporting

Confirm that scenarios and fixtures were frozen, every Choice includes `unexpected_action`, expected choices stayed hidden, fresh executors produced observable evidence, unexpected actions were attributed, verdict order was followed, and every scenario and Jev result links to the evidence set.

Output:

- `Target instruction and version`
- `Scenario catalog and fixture validity`
- `Behavior Choice design`
- `Expected-predicted-actual comparison`
- `Caller-derived verdicts and attribution`
- `Metrics and evidence locations`
- `Instruction issues, evaluation-design issues, and residual risk`
