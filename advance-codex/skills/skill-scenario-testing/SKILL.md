---
name: skill-scenario-testing
description: Empirically evaluate and analyze agent-facing text instructions such as skills, slash commands, task prompts, AGENTS or CLAUDE sections, and code-generation prompts by dispatching fresh subagents on fixed scenarios and collecting both executor self-reports and caller-side metrics. Use after creating or heavily changing reusable instructions, or when agent behavior suggests ambiguity in the instructions rather than a product or tool failure.
---

# Skill Scenario Testing

## Overview

Use this skill when you need to analyze instruction quality by evidence instead of author intuition.
The writer of a prompt is a biased evaluator of that same prompt.
Your job is to expose the instruction to fresh executors, collect structured evidence from both sides, and report what the evidence shows.

This skill owns the evaluation workflow for reusable agent-facing instructions.
It does not own general skill authoring, plugin packaging, or reusable tool policy outside this workflow.

## Use When

- you just created or heavily revised a reusable skill, task prompt, slash command, or AGENTS section
- a repeated agent failure looks more like instruction ambiguity than tool misuse or product logic failure
- the instruction matters enough that you want repeatable evidence rather than one-off taste review
- you need evidence that an instruction generalizes across more than one realistic scenario

## Do Not Use When

- the prompt is disposable and the evaluation cost is not justified
- the real problem is already known to be implementation, environment, or tool access failure
- you only want to rewrite wording to match personal style preferences
- you cannot launch fresh subagents and do not want a structural-review-only fallback

## Skill Boundary

This skill owns:

- scenario-based evaluation design for reusable instructions
- fixed acceptance checklists and caller-side scoring rules
- fresh-subagent execution for bias-reduced evaluation
- dual-sided evidence collection from executor reports and runtime metadata
- evidence-backed analysis and reporting

This skill does not own:

- creating the original skill or plugin boundary from scratch
- broad tool-use guidance unrelated to this evaluation workflow
- implementation debugging when the instruction text is not the bottleneck
- editing, patching, or rewriting the target instruction

## Core Stance

- Do not trust self-rereading as validation. The author already knows what they meant.
- Freeze the evaluation scenarios and checklist before dispatching executors.
- Prefer qualitative ambiguity signals first and quantitative metrics second.
- Treat each evaluation pass as one analysis theme. Do not mix unrelated targets.
- Reuse the workflow, not the executor. Each empirical run needs a fresh subagent.
- Default to an 8-scenario baseline when the user asks for scenario testing without a count. Use 2-3 scenarios only for a smoke check or when the user explicitly asks for a small pass.

## Workflow

### Phase 0: Static Alignment Check

Before you dispatch anything, verify that the instruction advertises the same job it actually teaches.

- read the trigger and usage promises in frontmatter or adjacent usage text
- read the body and confirm it truly covers that promised scope
- report description-body drift before dispatch, because later evaluation can produce false confidence

### Phase 1: Freeze The Baseline

Prepare a fixed evaluation set before any evaluation pass.

- choose realistic scenarios before dispatch; default to 8 scenarios for a normal reusable-instruction evaluation
- use 2-3 scenarios only for smoke checks, very narrow changes, or explicit small-scope requests
- if the user requests a larger set, split it into batches of N scenarios when N is provided, or into batches sized to the available subagent capacity
- define a 3-7 item checklist for each scenario
- include at least one `[critical]` item per scenario
- keep the checklist stable across the evaluation pass unless you explicitly redesign the experiment

Use checklist items that describe observable output requirements, not vague quality labels.

### Phase 2: Dispatch Fresh Executors

Ask fresh subagents to execute the instruction as a blank reader.

- do not use your own reread as a substitute
- do not reuse the same subagent across separate evaluation passes
- when scenarios are independent, you may dispatch them in parallel
- dispatch large scenario sets in batches; do not exceed the current subagent capacity
- if you hit a subagent limit, close completed or no-longer-needed executors, then continue with the remaining batch
- give only the minimum local context needed to run the scenario
- do not leak your expected answer, suspected bug, or preferred conclusion into the evaluation prompt

### Phase 3: Collect Dual-Sided Evidence

For each scenario, collect:

- executor self-report:
  - ambiguous wording
  - places where judgment had to be filled in
  - retries and why they happened
- caller-side metrics:
  - binary success/failure gated by the `[critical]` items
  - checklist accuracy percentage
  - `tool_uses`
  - `duration_ms`
  - retry count if the executor reported it

Treat qualitative evidence as the primary signal.
Treat time and step counts as secondary indicators of cognitive load or structural gaps.

### Phase 4: Interpret The Evidence

Interpret the evidence without editing the target instruction.

- identify which checklist items failed, partially passed, or required heavy inference
- group ambiguity reports by common cause
- distinguish critical failures from wording ambiguity and expected judgment calls
- describe follow-up options as analysis output, not as applied changes

### Phase 5: Optional Follow-Up Evaluation

If the user separately changes the instruction later, evaluate the changed instruction with fresh subagents.

- keep the scenario/checklist stable when comparing against the earlier run
- use fresh executors for the follow-up run
- do not treat the previous analysis as proof that the changed instruction works

### Phase 6: Decide What To Report

Report whether the evidence is sufficient, whether more scenarios are needed, or whether the evaluation design should change.

Reasonable sufficient-evidence signals:

- repeated evaluation passes with no new ambiguity reports
- stable accuracy, steps, and duration across comparable passes
- hold-out coverage does not collapse on one unseen scenario

Reasonable evaluation-redesign signals:

- ambiguity reports remain unclear after several evaluation passes
- the instruction only works on evaluated scenarios and fails on a hold-out scenario
- metrics look acceptable only because the executor is reconstructing missing structure by heavy reference-searching

## Evaluation Heuristics

- `tool_uses` is a structural signal, not only a cost metric. Large per-scenario skew often means the instruction is under-specified for one path.
- Accuracy alone can hide a brittle instruction if the executor had to infer too much.
- Faster is not automatically better. Missing explanation can reduce time while making the instruction less reliable.
- If the observed ambiguity does not map to any checklist item, the scenario or checklist may be aimed at the wrong abstraction level.

## Subagent Contract

When you dispatch a fresh executor, structure the request so the evidence comes back in a stable shape.

Include:

- the target instruction text or an exact path to read
- the scenario description
- the fixed checklist, including at least one `[critical]` item
- a task that says to execute the instruction and then report in a fixed format

Ask the executor to report:

- output or execution summary
- checklist status per item: pass, fail, or partial with reason
- ambiguous points
- judgment calls they had to make
- retry count and causes

Do not embed your own conclusions in this prompt.

## Environment Constraint

If you cannot launch fresh subagents, do not pretend the result is empirical evaluation.

Use one of these outcomes explicitly:

- structural review only: review the text for drift, clarity, and missing contracts without claiming empirical evidence
- evaluation skipped: report that empirical evaluation could not run because fresh dispatch was unavailable

Never replace fresh-executor evaluation with author self-rereading.

## Reporting Template

Use a compact per-evaluation report:

```md
## Evaluation Pass N

### Analysis Theme
- <one-line summary>

### Scenario Results
| Scenario | Success | Accuracy | Steps | Duration | Retries |
|---|---|---:|---:|---:|---:|
| A | pass | 100% | 4 | 18s | 0 |

### New Ambiguities
- <scenario>: <one-line ambiguity>

### New Judgment Calls
- <scenario>: <one-line judgment call>

### Follow-Up Options
- <additional evaluation or clarification needed>
```

## Review Pass

Before you stop, check:

- whether the scenarios still reflect real usage rather than overfitting to a prior finding
- whether `[critical]` items actually represent the minimum acceptable behavior
- whether a "clean" result was achieved by leaking answer context
- whether the findings are specific enough for the user to decide what to do next

## Output Contract

- `Target instruction`
- `Scenario set`
- `Checklist design`
- `Findings`
- `Sufficient evidence or redesign decision`
- `Residual risk`
