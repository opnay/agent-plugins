---
name: review-loop
description: Blocking-first review-and-fix workflow for high-confidence issues found during code review, QA review, or self-review. Use when the task is to inspect changes, select only material findings worth fixing now, apply one bounded fix at a time, and defer speculative, stylistic, or low-impact notes that should not stall delivery.
---

# Review Loop

## Overview

Use this skill when review work should focus on material issues instead of endless polish.
The loop is intentionally strict: identify only findings that cross a meaningful threshold, fix one bounded issue at a time, verify immediately, and defer low-value notes so delivery does not stall.
This skill should be usable on its own without assuming another workflow is managing scope around it.

## Use When

- The task is to process code review feedback, QA findings, or self-review findings.
- The user wants real issues fixed without turning the review into open-ended cleanup.
- The work should prioritize correctness, regressions, safety, and workflow reliability over polish.

## Do Not Use When

- The user wants broad refactoring, cleanup, or design exploration.
- The task is mostly speculative improvement with no clear bug, regression, or operational risk.
- The user is explicitly asking for exhaustive stylistic critique.

## Core Policy

- Default to `blocking-first`.
- Default to `P0/P1-first`.
- Do not raise a finding for speculative cleanliness or marginal improvement.
- Treat non-blocking ideas as deferred notes unless the user explicitly wants polish work.
- Fix one bounded finding per loop.
- Verify immediately after each fix.
- Stop if the remaining findings do not clearly justify delay or scope growth.

## Finding Threshold

Raise a finding only when it satisfies most of these:

- it has concrete impact on correctness, regression risk, safety, reliability, or maintainability
- it is supported by high-confidence evidence or a reproducible concern
- it is specific enough to fix without broad speculative redesign

If those conditions are weak, record it as a deferred note instead of a blocking finding.

## Severity Gate

- `critical`
  - data loss, security, broken production workflow, major correctness failure
- `high`
  - likely regression, incorrect behavior, strong reliability or maintainability risk
- `medium`
  - real issue, but not worth stalling delivery unless it compounds risk or clearly affects user outcomes
- `low`
  - style, naming polish, minor cleanup, or nice-to-have clarity improvements

Default review-loop target:

- fix `critical`
- fix `high`
- consider `medium` only when the value is clearly worth the interruption
- defer `low` by default

Priority policy:

- treat `P0` and `P1` as normal loop candidates
- treat `P2` as deferred by default unless it is high-confidence, bounded, and clearly worth interrupting delivery
- treat `P3` as deferred by default and never let it block delivery

## Loop Workflow

### Loop 0: Select One Review Target

1. List the known findings or review comments.
2. Separate them into:
   - blocking findings
   - deferred notes
3. Choose exactly one bounded blocking finding as the current loop target.

Output:

- Current review target
- Why it crosses the threshold

### Loop 1: Apply The Smallest Fix

1. Make the smallest change that resolves the selected finding.
2. Do not bundle nearby cleanup unless it is necessary for correctness or verification.

Output:

- Change made

### Loop 2: Verify Immediately

1. Re-run the narrowest command, test, or check that proves the finding is fixed.
2. Run a short regression check if the fix touches nearby sensitive behavior.

Output:

- Verification result

### Loop 3: Reassess

1. Decide whether the current finding is:
   - fixed
   - partially fixed
   - reopened
2. Decide whether another blocking finding remains worth addressing now.
3. If not, stop and leave the rest as deferred notes.

Output:

- Next fix target or stop decision

## Deferred Notes Rule

Move a point to deferred notes when:

- it is mostly stylistic
- it does not materially affect correctness or regression risk
- it requires broad redesign for marginal gain
- it is plausible but not supported strongly enough to justify interruption
- it is a `P2` with weak impact or unclear evidence
- it is a `P3`

Deferred notes should not stall delivery by default.

## Review Budget Rule

- One loop should fix one blocking finding.
- Prefer closing one meaningful issue over partially addressing several.
- If two consecutive loops only produce minor polish, stop and defer the rest.

## Stop Conditions

- No blocking findings remain.
- The remaining points are mostly deferred notes.
- The next improvement would require broad redesign rather than bounded correction.
- The same finding reopens 3 times.
- The user says stop, pause, or move on.

When stopping, report:

1. Which blocking findings were fixed
2. Which notes were deferred
3. Why the loop stopped

## Output Contract

- `Blocking findings`
- `Deferred notes`
- `Current review target`
- `Why this crosses the threshold`
- `Change made`
- `Verification result`
- `Next fix target`
- `Stop reason`
- `Residual risk`

## Final Checklist

- [ ] Findings were thresholded instead of accepted blindly
- [ ] Only material issues became blocking work
- [ ] One bounded finding was fixed at a time
- [ ] Verification ran immediately after each fix
- [ ] Low-value notes were deferred instead of blocking delivery

## Example Triggers

- "리뷰에서 진짜 중요한 것만 골라서 고쳐줘"
- "requested changes 중 blocking한 것만 처리해줘"
- "사소한 거 말고 실제 문제만 review loop로 잡아줘"
