---
name: ralph-loop
description: Iterative fix-verify-reassess workflow for bounded frontend or code quality improvements. Use when the task should be improved through short repeated loops instead of one large pass, such as UI polish, refactor stabilization, flaky issue reduction, QA issue burn-down, or repeated bug-fix verification.
---

# Ralph Loop

## Overview

Use this skill when the work should improve through short bounded loops rather than a large end-to-end delivery pass.
The loop is simple: pick one problem, apply the smallest useful fix, verify immediately, reassess, and then decide whether another loop is worth running.

## Use When

- The task needs repeated improvement cycles on one bounded problem at a time.
- The user wants progressive polish rather than one large refactor pass.
- QA or review work has surfaced several issues that should be burned down one by one.
- The task involves UI polish, refactor stabilization, flaky issue reduction, or repeated verification.

## Do Not Use When

- The task is a broad end-to-end feature delivery better suited to a phased workflow.
- The work is mostly planning, brainstorming, or architecture exploration without immediate fix-verify loops.
- The task needs one decisive implementation pass rather than repeated bounded refinement.

## Relationship To Larger Workflows

- Use `ralph-loop` as a bounded refinement loop when a larger task needs repeated fix-verify-reassess cycles.
- Do not let `ralph-loop` replace broader scope management when the task really needs end-to-end orchestration.

## Core Loop Policy

- One loop should focus on one primary issue.
- Prefer the smallest fix that can prove or disprove the current hypothesis.
- Verify immediately after each change.
- Reassess after every verification result instead of queueing several speculative fixes.
- Stop when the loop stops producing meaningful improvement.

## Loop Workflow

### Loop 0: Select The Current Target

1. Identify the current highest-value bounded issue.
2. Name the issue in one sentence.
3. State the current hypothesis for why it happens or how it should improve.

Output:

- Current target and loop hypothesis

### Loop 1: Apply The Smallest Fix

1. Make the smallest change that tests the current hypothesis.
2. Avoid bundling unrelated cleanup into the same loop.

Output:

- Concrete change made

### Loop 2: Verify Immediately

1. Re-run the narrowest command, check, or manual verification that can validate the change.
2. Record what passed, what failed, and what remains ambiguous.

Output:

- Verification result with evidence

### Loop 3: Reassess

1. Decide whether the issue is fixed, partially improved, unchanged, or has regressed.
2. If still open, decide whether another loop is justified.
3. If another loop is justified, choose the next smallest hypothesis to test.

Output:

- Next loop decision

## Issue Selection Rules

- Prefer high-signal issues over broad vague dissatisfaction.
- Prefer issues with a clear verification path.
- Prefer one dominant issue over several loosely related complaints.
- If several issues are present, choose the one most likely to unblock or simplify the rest.

## Verification Rules

- Run the narrowest meaningful check first.
- If a local fix could affect nearby behavior, run a short regression check before starting the next loop.
- Keep verification proportional to the loop size.
- Record repeated failures explicitly instead of rephrasing them as new issues.

## Stop Conditions

- The bounded target is fixed well enough.
- The same failure repeats 3 times.
- The next improvement would require a larger architectural change than the loop should own.
- The task should escalate into a broader workflow.
- The user says stop, pause, or redirect.

When stopping, report:

1. What improved
2. What remains open
3. Whether a broader workflow is now needed

## Output Contract

- `Current target`
- `Loop hypothesis`
- `Change made`
- `Verification result`
- `Next loop decision`
- `Stop reason`
- `Residual risk`

## Final Checklist

- [ ] One bounded issue was selected
- [ ] The loop used the smallest meaningful fix
- [ ] Verification ran immediately after the change
- [ ] The next step was chosen from evidence, not momentum
- [ ] The loop stopped when improvement was no longer local or meaningful

## Example Triggers

- "이거 한 번에 크게 말고 루프로 조금씩 다듬어줘"
- "QA 이슈를 하나씩 줄여가면서 고쳐줘"
- "UI를 반복 검증하면서 폴리시해줘"
- "플레이키한 문제를 좁은 루프로 안정화해줘"
