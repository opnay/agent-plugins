---
name: autopilot
description: Autonomous multi-phase delivery workflow for turning a brief product request into verified code using Codex tools. Use when the user wants hands-off end-to-end execution (analysis, planning, implementation, QA, and review) rather than a single focused edit.
---

# Autopilot Workflow

## Overview

Use this skill when the user wants full autonomous execution from idea to working code.
The workflow runs in gated phases with explicit verification and stop conditions.
This workflow should remain self-contained even when other skills are available for supporting tasks.

## Use When

- The user asks for end-to-end delivery from a brief idea.
- The task needs multiple phases: requirements, implementation, tests, and validation.
- The user prefers minimal back-and-forth and wants autonomous progress updates.

## Do Not Use When

- The user asks for brainstorming or options only.
- The user requests explanation or planning only.
- The task is a single small fix that does not require orchestration.
- The user explicitly asks for manual, step-by-step approval at each change.
- The main blocker is still aligning on what the user actually wants built because intent, boundaries, or tradeoffs are still materially unclear.

## Core Policy

- Complete one phase before moving to the next.
- Use parallelism only for independent work items.
- Verify immediately after each meaningful edit (format, lint, typecheck, test, build as relevant).
- Run a QA loop up to 5 cycles; if the same failure repeats 3 times, stop and report the root blocker.
- Maintain a QA issue list during the QA phase and update its status on every cycle.
- After each fix, run a short regression QA for previously closed issues before broader checks.
- Keep user-visible progress updates frequent and concrete.

## Phase Workflow

### Phase 0: Scope Lock

1. Infer goal, constraints, and success criteria from current context.
2. If a critical ambiguity blocks execution, ask a targeted question.
3. Record scope assumptions before implementation.
4. If several blocking ambiguities are narrow and concrete, batch them into a short clarification set instead of serial back-and-forth.
5. When the decision can be framed as a few realistic alternatives, prefer selection-style clarification prompts with a recommended default.
6. Ask only the minimum blocking questions, then resume execution immediately once the answers are available.
7. If the ambiguity is not narrow but about intent alignment, scope edge, or decision boundaries, stop and switch out of Autopilot instead of absorbing a full alignment interview into this workflow.

Output:

- Scope summary (goal, in-scope, out-of-scope, constraints)

### Phase 1: Read-Only Discovery

1. Inspect repository structure and relevant files.
2. Identify affected modules and risk points.
3. Choose applicable supporting skills or domain guidance for the areas involved (for example frontend architecture, backend architecture, or QA support).

Output:

- Implementation approach with impacted paths and risks

### Phase 2: Execution Plan

1. Break work into small executable tasks.
2. Use `update_plan` with a single in-progress step.
3. Define verification command(s) for each task.

Output:

- Stepwise plan with per-step verification method

### Phase 3: Implementation

1. Implement minimum viable changes first.
2. Keep execution scope explicit and stable while implementing separable work.
3. Integrate outputs and resolve conflicts safely.

Execution rules:

- Prefer `multi_tool_use.parallel` for independent reads/checks.
- Prefer `apply_patch` for focused single-file edits.
- Avoid destructive git operations.

Output:

- Working code aligned with scope and architecture constraints

### Phase 4: QA Loop

1. Run narrow checks first, then broader checks as needed, and capture failing evidence.
2. Create or update a QA issue list at `.agents/sessions/{YYYYMMDD}/autopilot-qa-issues.md`.
3. Record each issue with index, symptom, failing command, status (`open`, `fixed`, `reopened`), and owner/fix note.
4. Fix prioritized open issues one by one.
5. After each fix, run a short QA pass:
   - re-run the failing command for the current issue
   - run quick regression checks for recently fixed issues to detect reappearance
6. After each fix's short QA pass, run broader checks as needed and update issue statuses.
7. Repeat the cycle until all tracked QA issues are closed and required checks are green, but stop when the QA cycle count reaches 5.
8. Stop early if identical failure appears 3 times and report the fundamental issue.

Output:

- QA issue list history and verified test/build/lint status with command evidence

### Phase 5: Validation and Delivery

1. Review for correctness, regression risk, and missing tests.
2. Summarize changes, verification evidence, and residual risks.
3. Provide clear next options if follow-up actions are natural.

Output:

- Final delivery summary ready for user decision (merge, iterate, or pause)

## State and Resume

Track progress in thread-local files:

- Path: `.agents/sessions/{YYYYMMDD}/autopilot-state.md`
- Path: `.agents/sessions/{YYYYMMDD}/autopilot-qa-issues.md`
- Minimum state fields: current phase, qa cycle count, completed steps, open blockers, next action, open issue IDs

If interrupted, resume from the latest completed phase instead of restarting.

## Stop Conditions

- User says stop/cancel/abort.
- Critical requirement ambiguity prevents safe implementation.
- Same QA failure repeats 3 times.
- Validation fails repeatedly without new actionable fixes.

When stopping, report:

1. What completed
2. What is blocked
3. Exact decision/input needed from user

## Final Checklist

- [ ] Scope and assumptions were locked
- [ ] Plan was explicit and executable
- [ ] Implementation completed for in-scope items
- [ ] QA issue list was maintained and all tracked issues are closed
- [ ] Short regression QA confirmed no reappearance of closed issues
- [ ] Relevant verification commands passed
- [ ] Residual risks and follow-up options were reported

## Example Triggers

- "autopilot로 이 요구사항 끝까지 구현해줘"
- "이 기능 아이디어를 실제 동작 코드까지 만들어줘"
- "분석부터 테스트까지 전부 알아서 진행해줘"
