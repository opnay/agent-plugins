---
name: subagent-role
description: Define role-assigned subagent packets, spawn boundaries, answer contracts, integration rules, and learning notes for understanding how subagents behave under bounded roles.
---

# Subagent Role

## Overview

Use this skill when the task is specifically about role-assigned subagents.
The goal is to make subagent use explicit and learnable: decide whether a subagent is useful, define a bounded role packet with real specialty, preserve spawn boundaries, require a structured answer, and record what the caller learned from the result.

## Role Fit

Use a role-assigned subagent only when at least one condition is true:

- the role is a sidecar task that can run while the caller does non-overlapping work
- the role has a concrete, bounded output contract
- the role owns a disjoint write scope or a specific read-only question
- the role has a specialty mix that changes the expected judgment
- the caller wants to compare subagent behavior across roles

Prefer local caller work when the task is the immediate blocker, too vague to delegate, or tightly coupled to the caller's next step.

## Role Specialty

Define the role as a specialty mix, not as a loose persona.
Use only the axes that matter for the task:

- `functional_role`: PM, PO, engineer, designer, reviewer, researcher
- `responsibility_domain`: game scenario, frontend implementation, prompt evaluation, harness diagnostics
- `background_expertise`: game-planning-origin, game-PM-origin, frontend-origin, infra-origin
- `general_expertise`: product strategy, UX writing, implementation rigor, verification
- `decision_style`: product sense, technical critique, verification

Examples:

- `PM role for game scenario ownership`: PM judgment focused on game scenario responsibility
- `engineer role for frontend development`: engineering judgment focused on frontend implementation
- `game-planning-origin frontend developer`: frontend implementation with game-planning background
- `game-PM-origin PO`: PO judgment with game PM background and product strategy expertise

## Role Pattern Catalog

Use these as starting points and tighten them for the actual task:

- `game-PM-origin PO`: product tradeoffs, player impact, retention risk, prioritization tradeoff, recommended decision. Caller verifies the final product decision against product goals, player data, and product assumptions.
- `PM role for game scenario ownership`: scenario intent, coherence risk, player-facing issue, revision direction
- `engineer role for frontend development`: implementation plan, file ownership, integration risk, verification path. If coding is involved, require disjoint write scope, concurrent changes guard, and caller review of staged changes or final patch.
- `game-planning-origin frontend developer`: interaction interpretation, implementation constraints, UX risk, fallback option
- `designer role for UX review`: UX findings, severity, proposed adjustment, acceptance signal. Keep it read-only unless explicitly assigned implementation; caller integrates only material issues.
- `QA role for regression test design`: test matrix, high-risk flows, regression cases, pass/fail signal

## Role Packet

Define the packet before spawning:

- `role_name`
- `role_specialty`
- `objective`
- `context`
- `ownership`
- `non_goals`
- `expected_output`
- `verification_signal`
- `integration_rule`
- `stop_condition`

For coding work, include disjoint file or module ownership and tell the subagent it is not alone in the codebase. It should not revert edits made by others and should adapt to concurrent changes.

## Spawn Boundary

- Follow active runtime and tool policy for whether subagents may be spawned.
- Do not treat this skill as user approval for destructive, external, or approval-boundary actions.
- Do not delegate duplicate versions of the same unresolved task.
- Do not hand off the caller's immediate critical-path blocker if the caller would only wait for the answer.

## Answer Contract

Ask the subagent to return:

- `findings`
- `changed_files` or `inspected_surfaces`
- `verification_performed`
- `confidence`
- `context_gaps`
- `integration_recommendation`
- `learning_note`

Treat the answer as evidence. The caller remains responsible for verification, integration, and final reporting.

## Learning Note

After the subagent returns, record:

- whether the role was narrow enough
- whether the specialty mix affected the answer in a useful way
- whether the output was easy to integrate
- whether the caller had to redo discovery
- what should change in the next role packet

## Output

- `Role fit`
- `Role specialty`
- `Role packet`
- `Spawn boundary`
- `Answer contract`
- `Integration rule`
- `Learning note`
- `Residual risk`

## Guardrails

- Do not make subagent use automatic.
- Do not bypass user, runtime, tool, or safety approval boundaries.
- Do not let a subagent's answer become the final answer without caller-side review.
