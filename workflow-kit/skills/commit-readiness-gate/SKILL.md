---
name: commit-readiness-gate
description: Final readiness gate for deciding whether an intended change unit is ready to move toward commit. Use when implementation is nearly done and Codex needs to finish self-review, scoped verification, risk classification, and the minimum necessary review recommendation before telling the user the change is commit-ready.
---

# Commit Readiness Gate

## Overview

Use this skill to decide whether an intended change unit has reached commit-ready status.
The goal is not to ask for a commit by itself. The goal is to finish the final self-review, verification, and risk classification needed to tell the user whether it is reasonable to commit now.
Use this skill after the main implementation work is complete or nearly complete.

## Use When

- a change is implemented and needs a final commit-readiness check
- you want a repeatable gate before review or commit
- you need to confirm the change unit is isolated, verified, and scoped correctly
- you want the minimum necessary review recommendation before telling the user the change is ready

## Do Not Use When

- the task is still mostly exploratory
- implementation is still clearly incomplete
- the main need is to process review findings one by one
- a repository-specific release or deployment process should be used instead

## Workflow

### 1. Lock The Intended Change Unit

- Confirm the intended commit scope is a single-purpose change set.
- Exclude unrelated workspace changes before final review begins.
- If the diff mixes concerns, split the work before continuing.

### 2. Verify During Final Editing

- Run the narrowest relevant verification immediately after each final edit.
- Prefer deterministic local checks first.
- Do not defer all validation to the very end if a fix is still being made.

### 3. Perform Commit-Readiness Self-Review

- Inspect the intended diff directly, whether staged already or still being prepared for staging.
- Confirm the change matches the requested scope and contains no unrelated edits.
- Check naming, comments, docs, and control flow for obvious cleanup opportunities.
- Confirm workflow records or delivery notes are updated when the surrounding workflow requires them.

### 4. Run Final Relevant Verification

- Rerun the final relevant checks for the actual intended change unit.
- If a required check fails, fix it before calling the change commit-ready.
- If a relevant check is intentionally skipped, record the skip and residual risk explicitly.

### 5. Classify Change Risk

- `docs-only`
  - wording, structure, or documentation-only changes with no runtime effect
- `low-risk code`
  - isolated implementation changes with clear local behavior and narrow verification
- `behavioral`
  - changes that alter existing behavior, compatibility, fallback logic, or user-visible outcomes
- `structural`
  - changes that affect responsibilities, abstractions, boundaries, or ownership

### 6. Recommend The Minimum Necessary Review

- `docs-only`
  - usually no extra review lens when the change unit is isolated and final verification is complete
- `low-risk code`
  - recommend `correctness`
- `behavioral`
  - add `regression`
- `structural`
  - add `design integrity`
- `abstraction-heavy` or repeated-logic risk
  - add `clean architecture hygiene`
- `scope drift` risk
  - add `scope boundary`

### 7. Gate Outcome

- If the change unit is isolated, verified, and risk-classified, say it is commit-ready.
- If not, say exactly what blocks commit-readiness and what must be fixed first.
- Do not blur the result; the output should clearly pass or fail the gate.

## Commit-Readiness Checklist

- the intended diff is single-purpose
- unrelated files are not part of the change unit
- final relevant verification is green
- skip decisions and residual risks are explicit when needed
- recommended review lenses match the staged or intended risk
- the eventual commit message scope can match the real change unit

## Output Contract

- `Commit-readiness status`
- `Intended change unit`
- `Verification run`
- `Deliberate skips`
- `Risk class`
- `Recommended review lenses`
- `Blocking issues`
- `Residual risk`

## Guardrails

- Do not treat this gate as a substitute for unfinished implementation.
- Do not declare commit-readiness when the change unit still mixes concerns.
- Do not recommend broad review sets without a risk reason.
- Do not hide failed or skipped verification behind a soft-ready answer.
