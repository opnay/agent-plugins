---
name: empirical-prompt-tuning
description: Empirically evaluate and iteratively improve agent-facing text instructions such as skills, slash commands, task prompts, AGENTS or CLAUDE sections, and code-generation prompts by dispatching fresh subagents on fixed scenarios, collecting both executor self-reports and caller-side metrics, and repeating until improvement converges. Use after creating or heavily revising reusable instructions, or when agent behavior suggests ambiguity in the instructions rather than a product or tool failure.
---

# Empirical Prompt Tuning

## Overview

Use this skill when you need to improve instruction quality by evidence instead of author intuition.
The writer of a prompt is a biased evaluator of that same prompt.
Your job is to expose the instruction to fresh executors, collect structured evidence from both sides, and make the smallest justified revision each round.

This skill owns the evaluation workflow for reusable agent-facing instructions.
It does not own general skill authoring, plugin packaging, or reusable tool policy outside this workflow.

## Use When

- you just created or heavily revised a reusable skill, task prompt, slash command, or AGENTS section
- a repeated agent failure looks more like instruction ambiguity than tool misuse or product logic failure
- the instruction matters enough that you want a repeatable evaluation loop rather than one-off taste edits
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
- minimal revision loops and convergence rules

This skill does not own:

- creating the original skill or plugin boundary from scratch
- broad tool-use guidance unrelated to this evaluation workflow
- implementation debugging when the instruction text is not the bottleneck

## Core Stance

- Do not trust self-rereading as validation. The author already knows what they meant.
- Freeze the evaluation scenarios and checklist before revising the instruction.
- Prefer qualitative ambiguity signals first and quantitative metrics second.
- Treat each iteration as one improvement theme. Do not mix unrelated edits.
- Reuse the workflow, not the executor. Each empirical run needs a fresh subagent.

## Workflow

### Phase 0: Static Alignment Check

Before you dispatch anything, verify that the instruction advertises the same job it actually teaches.

- read the trigger and usage promises in frontmatter or adjacent routing text
- read the body and confirm it truly covers that promised scope
- fix description-body drift first, or the later evaluation can produce false confidence

### Phase 1: Freeze The Baseline

Prepare a fixed evaluation set before any tuning pass.

- choose 2-3 realistic scenarios: one common path and 1-2 edge cases
- define a 3-7 item checklist for each scenario
- include at least one `[critical]` item per scenario
- keep the checklist stable across the iteration unless you explicitly redesign the experiment

Use checklist items that describe observable output requirements, not vague quality labels.

### Phase 2: Dispatch Fresh Executors

Ask fresh subagents to execute the instruction as a blank reader.

- do not use your own reread as a substitute
- do not reuse the same subagent across iterations
- when scenarios are independent, you may dispatch them in parallel
- give only the minimum local context needed to run the scenario
- do not leak your expected answer, suspected bug, or intended fix into the evaluation prompt

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

### Phase 4: Apply The Smallest Justified Patch

Revise the instruction only after stating what the patch is expected to improve.

- name the single improvement theme for the iteration
- map the patch to the exact checklist items or decision rules it should affect
- keep related micro-edits together when they serve the same theme
- defer unrelated improvements to later iterations

### Phase 5: Re-Run With Fresh Executors

Run the same scenarios again with fresh subagents.

- compare both qualitative and quantitative changes
- do not declare success from one clean run
- if accuracy stays high but one scenario is much more expensive in `tool_uses`, treat that as a structural gap worth fixing

### Phase 6: Decide Whether To Stop Or Redesign

Prefer stopping when the loop has genuinely converged, not when you are tired of editing.

Reasonable stop signals:

- two consecutive iterations with no new ambiguity reports
- only marginal improvement in accuracy, steps, and duration
- hold-out coverage does not collapse on one unseen scenario

Reasonable redesign signals:

- ambiguity reports do not shrink after several iterations
- the instruction only works on tuned scenarios and fails on a hold-out scenario
- metrics look acceptable only because the executor is reconstructing missing structure by heavy reference-searching

## Evaluation Heuristics

- `tool_uses` is a structural signal, not only a cost metric. Large per-scenario skew often means the instruction is under-specified for one path.
- Accuracy alone can hide a brittle instruction if the executor had to infer too much.
- Faster is not automatically better. Missing explanation can reduce time while making the instruction less reliable.
- If the improvement estimate did not move the expected checklist items, your patch was probably aimed at the wrong abstraction level.

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
- evaluation skipped: report that empirical tuning could not run because fresh dispatch was unavailable

Never replace fresh-executor evaluation with author self-rereading.

## Reporting Template

Use a compact per-iteration report:

```md
## Iteration N

### Change Theme
- <one-line summary>

### Scenario Results
| Scenario | Success | Accuracy | Steps | Duration | Retries |
|---|---|---:|---:|---:|---:|
| A | pass | 100% | 4 | 18s | 0 |

### New Ambiguities
- <scenario>: <one-line ambiguity>

### New Judgment Calls
- <scenario>: <one-line judgment call>

### Next Patch
- <smallest justified revision>
```

## Review Pass

Before you stop, check:

- whether the scenarios still reflect real usage rather than the latest patch
- whether `[critical]` items actually represent the minimum acceptable behavior
- whether a "clean" result was achieved by leaking answer context
- whether the latest patch improved clarity without making the instruction bloated

## Output Contract

- `Target instruction`
- `Scenario set`
- `Checklist design`
- `Latest findings`
- `Recommended patch`
- `Stop or redesign decision`
- `Residual risk`
