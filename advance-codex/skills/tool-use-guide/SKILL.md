---
name: tool-use-guide
description: Design or revise reusable tool-usage guidance for skills, plugins, or custom agents. Use when the main problem is how an artifact should select, sequence, constrain, or escalate tools without burying tool policy inside a domain workflow.
---

# Tool Use Guide

## Overview

Use this skill when the main deliverable is reusable guidance for tool choice and tool behavior.
Its job is to make tool-use policy explicit: what tool should be preferred, what should be avoided, when the agent should stay local, when it should ask the user, and when runtime-specific tool rules should live outside a product or domain skill.

This skill is not for designing the domain workflow itself.
It is for designing the tool layer that supports that workflow without coupling the domain artifact to accidental runtime behavior.

## Use When

- a skill, plugin, or custom agent needs clearer rules for how to choose between several tools
- tool behavior is inconsistent because the current instructions only say what outcome is wanted, not how tool choice should work
- a domain skill is starting to absorb runtime-specific tool policy that should be isolated
- the main question is when to inspect locally, browse externally, ask the user, delegate, or avoid a tool entirely
- you want reusable guidance for tool sequencing, escalation, or stop conditions rather than a new domain workflow

## Do Not Use When

- the main deliverable is a new skill boundary, plugin bundle, or custom agent definition
- the artifact already has clear tool rules and the real issue is domain scope or acceptance criteria
- the task is only about a single prompt wording tweak with no reusable tool policy behind it
- the main need is implementation, not reusable tool guidance

## Scope Boundary

This skill owns:

- tool selection policy
- tool sequencing policy
- escalation and approval rules
- when to ask the user versus infer locally
- when runtime-specific tool guidance should be isolated from a domain artifact
- how to phrase tool guidance so it is operational instead of aspirational

This skill does not own:

- the full domain workflow
- plugin packaging by itself
- custom agent file authoring by itself
- implementation of the business feature being discussed

## Core Policy

- Write tool guidance around decisions, not around tool names alone.
- Separate domain intent from runtime mechanics when they would otherwise become entangled.
- Prefer durable trigger conditions over one-off examples.
- Treat tool guidance as part of the artifact's operating contract.
- Keep the rules narrow enough that a future runtime swap is still understandable.
- Do not require a tool when the real need is only a user-facing outcome.

## Design Questions

Clarify these before drafting guidance:

- `Decision point`
  - what choice the artifact keeps getting wrong or inconsistent
- `Preferred tool behavior`
  - what tool or class of tools should normally be used first
- `Fallback behavior`
  - what should happen when the preferred tool is unavailable, insufficient, or risky
- `Ask-vs-infer boundary`
  - what must be confirmed with the user and what may be decided from local evidence
- `Local-vs-external boundary`
  - what should be learned from the workspace first and what justifies external lookup
- `Escalation boundary`
  - when approval, delegation, or a different workflow is required
- `Coupling risk`
  - what domain artifact would become too runtime-specific if this guidance were embedded there

## Workflow

### Phase 0: Identify The Real Surface

1. Identify the artifact that currently owns or is about to own the guidance:
   - skill
   - plugin guide
   - custom agent
   - adjacent runtime-facing helper artifact
2. Identify whether the problem is:
   - missing tool preference
   - missing sequencing
   - missing escalation rules
   - wrong ask-vs-infer boundary
   - wrong placement of runtime-specific guidance
3. Decide whether the guidance belongs in:
   - a bounded tool-use skill
   - a plugin entrypoint guide
   - a custom agent's behavioral instructions
   - a local patch inside an existing artifact

Output:

- `Tool-use surface`

### Phase 1: Define The Decision Rules

1. List the concrete decision points the artifact must handle.
2. For each decision point, define:
   - default action
   - exceptions
   - escalation trigger
   - explicit non-goal
3. Prefer conditions such as:
   - "ask only when the answer cannot be discovered locally without risky assumptions"
   - "prefer the structured user-input tool when the question fits 1-3 short bounded choices"
   - "read local repository context before using external search"
4. Remove tool mentions that are really product rules in disguise.

Output:

- `Decision rules`
- `Placement rationale`

### Phase 2: Draft The Guidance

Draft instructions that are:

- operational
- reusable
- tool-aware
- not overloaded with domain detail

When useful, encode guidance in this order:

1. trigger condition
2. preferred tool behavior
3. fallback behavior
4. escalation condition
5. stop condition

Output:

- `Draft tool-use guidance`

### Phase 3: Review For Coupling And Drift

Check:

- whether the guidance can be understood without hidden sibling context
- whether domain rules and runtime rules are separated cleanly enough
- whether a tool is mandated only when that mandate is actually durable
- whether the wording tells the agent when not to use a tool
- whether the same guidance would still make sense if surrounding implementation details changed

Output:

- `Reviewed guidance`
- `Residual coupling risk`

## Placement Rules

- Put tool guidance in a dedicated artifact when it would otherwise leak runtime specifics into a domain skill.
- Keep plugin-level routing guidance in the plugin guide unless the issue is specifically about tool policy.
- Keep custom-agent operational behavior in the agent spec when the tool rules are inseparable from that role.
- Do not create a new tool-use skill if one localized artifact-level patch is enough and remains cleanly bounded.

## Output Contract

- `Tool-use surface`
- `Decision points`
- `Preferred tool behaviors`
- `Fallback and escalation rules`
- `Recommended placement`
- `Why this should not live in the domain artifact`
- `Validation path`
- `Residual risk`

## Guardrails

- Do not hide tool policy inside a domain skill just because it is convenient.
- Do not create a tool-use artifact when the issue is really missing product intent.
- Do not force a specific tool without naming the decision pattern that justifies it.
- Do not turn tool guidance into a list of tool names with no trigger conditions.
- Do not let one-off runtime quirks define the whole reusable artifact.
