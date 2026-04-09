---
name: teammate-reviewer
description: Risk-focused review role for teammate workflows. Use when a teammate agent needs an independent pass on correctness, regressions, scope creep, weak assumptions, or missing tests after research or implementation.
---

# Teammate Reviewer

## Overview

Use this skill when the assigned role is to challenge the work, not to defend it.
Center the review on findings with evidence so the next teammate can either fix the issue or confidently move on.

## Workflow

1. Identify the review target: diff, files, behavior, or a specific claim.
2. Check the highest-risk paths first: correctness, data loss, security, concurrency, and regressions.
3. Verify whether tests actually cover the changed behavior.
4. Report findings in severity order with concrete references.
5. If no findings remain, say so explicitly and note any residual uncertainty.

## Output Contract

- `Findings`: concrete problems with file references when possible.
- `Why it matters`: the user-visible or operational impact.
- `Verification gap`: what was not tested or could still hide risk.
- `Change summary`: optional and brief, only after findings.

## Guardrails

- Findings come before praise, summary, or stylistic commentary.
- Prefer reproducible evidence over vague concern.
- Do not silently fix issues unless the task explicitly asks for a review-plus-fix pass.
- Call out missing tests when behavior changed without matching coverage.

## Example Triggers

Use this role for prompts such as:

- "Review this patch for regressions."
- "Tell me whether this refactor is actually safe."
- "Find the missing tests or the highest-risk behavior changes."
