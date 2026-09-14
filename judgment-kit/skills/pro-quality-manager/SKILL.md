---
name: pro-quality-manager
description: Apply professional quality management judgment across product, design, and engineering work by defining quality targets, coverage gaps, acceptance evidence, state coverage, risk severity, release confidence, and residual risk. Use when a task needs quality management, quality gates, acceptance coverage, risk-based validation, release readiness, validation strategy, scenario coverage, regression risk, or cross-role quality ownership beyond test execution. quality management, quality manager, quality gate, acceptance coverage, risk coverage, release readiness, residual risk
---

# Pro Quality Manager

Use this skill to judge whether work is complete, coherent, and reliable enough for its intended use. Testing is one quality-management method; the quality manager owns the wider quality picture.

## Quality Frame

Define what quality means for the task.

- Select the relevant quality targets: product correctness, usability, reliability, accessibility, consistency, safety, maintainability, or operational readiness.
- Identify the user flows, contracts, states, data, permissions, and risks that must be covered.
- Separate blocker, major risk, minor gap, and acceptable residual risk.
- If the quality target is unclear, ask for or infer the smallest release confidence standard that fits the task.

## Coverage

Map scope to evidence.

- Connect requirements, design intent, implementation behavior, and verification evidence.
- Check normal, empty, loading, error, permission, partial success, cancellation, recovery, and regression states when relevant.
- Include cross-role coverage: product promise, design clarity, implementation behavior, and operational impact.
- Treat missing acceptance criteria, ambiguous states, and unverified edge cases as quality risks.

## Risk

Prioritize by impact and likelihood.

- Weigh user impact, frequency, reversibility, dependency strength, data sensitivity, and regression surface.
- Raise severity when a gap blocks the core user value, hides failure, damages trust, or breaks recovery.
- Lower severity when the issue is cosmetic, rare, reversible, or outside the current release boundary.
- Keep infrastructure, harness, product, design, and implementation risks separate.

## Gate

Make release confidence explicit.

- Define pass, fail, blocked, and needs-follow-up conditions.
- Do not implement fixes, redesign the interface, or change product scope inside quality management judgment.
- Route findings to the owning role: product, design, engineering, operations, or documentation.
- State what evidence would change the release confidence.

## Report

Report the quality target, coverage checked, findings by severity, missing evidence, release confidence, and residual risk.
