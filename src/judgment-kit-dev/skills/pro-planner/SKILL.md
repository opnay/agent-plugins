---
name: pro-planner
description: Apply product and service planning judgment to decompose broad product or feature requests into user problems, product modes, feature maps, supplemental planning surfaces, MVP scope, requirements, priorities, acceptance criteria, tradeoffs, and designer/engineer handoff contracts before design or implementation. Use when a task needs PO/PM judgment, feature planning, service planning, product requirements, scope narrowing, product strategy, user problem framing, product decomposition, design system brief, prioritization, release slicing, or handoff clarity. product planning, service planning, feature planning, product requirements, product decomposition, design system brief, MVP scope, acceptance criteria, product strategy, user problem, prioritization, product tradeoff
---

# Pro Planner

Use this skill to decide what a product, service, or feature should solve, for whom, which product mode fits, how the problem breaks into feature areas and supporting surfaces, how far the current scope should go, and what evidence will show the result is done. Keep planning separate from visual design and implementation detail: define the product contract that design and engineering can use.

## Problem

Start with the user problem before accepting the requested feature shape.

- Separate the user's goal, current workaround, blocked moment, frequency, impact, and risk.
- For broad requests such as "build an expense tracker", separate the product name from the user's desired outcome and possible product modes.
- Treat internal requests, competitor features, and technical possibilities as clues, not the problem itself.
- Define the problem in observable terms: situation, user action, expected result, current failure, and consequence.
- If the problem is vague, narrow it with one concrete use case before listing features.

## Decomposition

Break broad product or feature requests before planning screens or implementation.

- Identify product mode options when the same product label can mean different behavior, policy, success criteria, or risks.
- Compare product modes by user action, default policy, required data, core feedback, and acceptance signal.
- Create a feature map from user-recognized work areas, not from code modules. For each area, name the user goal, core action, required state, data, and failure or empty case.
- Group requirements into functional, policy, state, data, exception, and non-goal groups.
- Identify supplemental planning surfaces such as information architecture, screen inventory, design system brief, settings, onboarding, notifications, data management, analytics, and operational surfaces.
- Treat decomposition as planning judgment. Hand off design and engineering decisions; do not run the design or development loop inside this skill.

## Supplemental Surfaces

Capture product decisions that core feature lists often miss.

- Design system brief: product tone, density, trust level, state-language needs, component candidates, accessibility, and responsive priorities. Do not choose final visual style, palette, or component design here.
- Information architecture: navigation groups, screen inventory, each screen's purpose, primary information, and required states.
- Supporting surfaces: settings, onboarding, notifications, data import/export, help, admin, analytics, and other non-core surfaces that affect product completeness.
- Deferred surfaces: useful but out-of-scope surfaces with the reason they are deferred and the revisit condition.

## User

Planning depends on who is using the product and under what constraints.

- Identify primary users, secondary users, operators, administrators, and affected non-users when relevant.
- Check user skill level, use frequency, environment, urgency, permission level, and risk tolerance.
- Distinguish one-time setup, repeated work, high-risk decisions, and exploratory use.
- Prefer user language over internal system names when framing goals and outcomes.

## Value

State why the feature or service should exist.

- Express value as a user result, reduced effort, avoided risk, faster decision, clearer control, or stronger confidence.
- Check whether each product mode changes the value promise, user behavior, or definition of success.
- Connect each major feature candidate to a user value or a required product constraint.
- Drop, defer, or question items that do not support the problem, value, or acceptance signal.
- If value is still unclear, return to problem and user framing before expanding scope.

## Scope

Define the smallest complete release slice that proves the core value.

- MVP means the smallest complete user value or product hypothesis test, not a thin pile of partial features.
- For a broad product request, choose or question the product mode before defining MVP scope.
- Separate current scope, next scope, and explicit non-goals.
- Use non-goals to protect the current decision boundary, not to erase future ideas.
- Keep scope tied to a user journey, policy boundary, or acceptance signal.

## Requirements

Turn planning decisions into testable product contracts.

- Functional requirements: what the user can do and how the system responds.
- Policy requirements: permissions, limits, state transitions, risk actions, retention, notifications, and operational rules.
- State requirements: empty, loading, partial success, error, permission loss, cancellation, delay, and concurrent changes.
- Data requirements: required fields, source, ownership, visibility, freshness, and audit needs.
- Feature-area requirements: the user goal, core action, required state, data, exception, and acceptance signal for each area.
- Supplemental requirements: design-system inputs, IA, screen inventory, settings, analytics, onboarding, and operational support that downstream roles need.
- Avoid locking UI layout or code architecture unless the product contract truly requires it.

## Flow

Describe the large user flow and decision points.

- Include entry condition, goal recognition, key choice, action, feedback, result, and recovery.
- Name the states users must understand before design begins.
- Identify points where users need confidence, confirmation, explanation, undo, or escalation.
- Keep flow at product level; leave component layout and visual hierarchy to design work.

## Acceptance

Make done observable.

- Write acceptance criteria with input, action, state, output, and failure condition where possible.
- Include success conditions and explicit failure or cannot-handle conditions.
- Define what must be true for users, operators, and the system after completion.
- If acceptance is subjective, translate it into what a user can see, choose, complete, avoid, or recover from.

## Priority

Rank work by product value and decision quality.

- Weigh user value, risk reduction, dependencies, implementation and operation cost, learning value, and reversibility.
- Prioritize questions that decide product mode and dependencies between feature areas.
- Raise priority for frequent user blockers, high-risk moments, strong dependencies, or decisions that unlock later work.
- Lower priority for cosmetic detail, speculative expansion, or requirements without a clear user result.
- When uncertainty is high, prioritize the smallest release slice or experiment that can teach the next decision.

## Tradeoff

Make product choices explicit.

- Compare speed vs completeness, automation vs control, simplicity vs extensibility, personalization vs consistency, and guided flow vs expert efficiency.
- Record the chosen option, rejected alternative, reason, residual risk, and revisit condition.
- Do not hide tradeoffs inside vague wording such as "better UX" or "more flexible"; name the actual cost.

## Handoff

Prepare the next role without taking over its job.

- Designer handoff: product mode, user goal, feature map, information architecture, screen inventory, design system brief, main flow, information priority, required states, risky moments, and UX concerns.
- Engineer handoff: feature-area requirements, data and state contracts, policy rules, exceptions, acceptance criteria, and non-goals.
- Keep handoff concrete enough to prevent ambiguity and open enough for design and engineering judgment.

## Working Rule

Use only the sections that matter to the request. For small tasks, produce a compact planning contract: problem, user, scope, requirements, acceptance, and handoff. For broad or ambiguous tasks, start by decomposing the product problem: product mode options, feature map, supplemental planning surfaces, requirement groups, scope, non-goals, and handoff needs.

Report planning work by naming the planning basis, product mode decision or open question, feature map, supplemental surfaces, chosen scope, tradeoff, handoff, and remaining ambiguity when important.
