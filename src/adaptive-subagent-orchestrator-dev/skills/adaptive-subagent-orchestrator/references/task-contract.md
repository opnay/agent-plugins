# Task Contract

Use this reference before spawning subagents or checking their returned output.

## Subagent Prompt Template

```text
Objective:
- <exact question or implementation goal>

Scope:
- Include: <directories, modules, packages, files, feature area>
- Exclude: <out-of-scope items>

Access mode:
- read-only | write-enabled

Ownership:
- <none for read-only>
- <explicit writable file set for write-enabled>

Inputs:
- User request: <summary or quoted requirement>
- Required artifacts: <diff, logs, failing tests, code paths, contracts>

Constraints:
- Do not spawn subagents.
- Do not expand beyond Scope.
- Do not change excluded areas.
- Preserve <compatibility, performance, security, API, behavior>.
- Commands allowed: <commands or none>
- Shared-state limits: <ports, DBs, temp dirs, lockfiles, generated files>

Deliverable:
- <specific result the main agent can integrate>

Evidence:
- Cite file paths, symbols, stable line numbers when useful, commands, and key outputs.
- Separate confirmed facts from inference or assumptions.

Completion criteria:
- <observable pass/block/inconclusive criteria>
```

## Return Template

```text
Status:
- completed | blocked | inconclusive

Summary:
- <concise conclusion>

Evidence:
- <paths, symbols, line numbers, command results>
- <separate confirmed facts from inference>

Files inspected:
- <main files>

Files changed:
- <files and reasons, or none>

Validation:
- <tests/checks/repro commands and pass/fail>

Risks and uncertainties:
- <remaining assumptions, unknowns, regression risks>

Recommended action:
- <next action for main agent>
```

Do not request raw full logs, long command output, work diaries, hidden reasoning, or unsupported confidence.

## Read-Only Investigation Example

```text
Objective:
- Identify likely causes of intermittent login request failures in the auth module.

Scope:
- Include: src/auth, tests/auth, logs mentioning auth request IDs.
- Exclude: payment and inventory modules except for call boundaries into auth.

Access mode:
- read-only

Ownership:
- none

Inputs:
- User report: login, payment, and inventory requests fail intermittently.
- Required artifacts: current branch diff, auth tests, recent auth logs.

Constraints:
- Do not spawn subagents.
- Do not edit files.
- Do not inspect unrelated modules beyond boundary calls.

Deliverable:
- Ranked auth-module causes with evidence and a recommended next check.

Evidence:
- Cite files, symbols, stable line numbers, and log snippets only when necessary.
- Mark assumptions explicitly.

Completion criteria:
- At least one supported cause or an explicit inconclusive result with missing evidence.
```

## Write-Enabled Implementation Example

```text
Objective:
- Fix the isolated parser bug in src/parser/date.ts and add focused tests.

Scope:
- Include: src/parser/date.ts, tests/parser/date.test.ts.
- Exclude: shared parser config, package manager files, generated files.

Access mode:
- write-enabled

Ownership:
- src/parser/date.ts
- tests/parser/date.test.ts

Inputs:
- Confirmed bug: timezone offset is ignored for ISO strings with explicit offset.
- Contract: keep existing local-date parsing behavior unchanged.

Constraints:
- Do not spawn subagents.
- Do not edit files outside Ownership.
- Do not change public API names.
- Allowed command: pnpm test -- tests/parser/date.test.ts

Deliverable:
- Minimal patch plus validation result.

Evidence:
- Changed file paths and test output summary.

Completion criteria:
- Focused parser tests pass or blocker is clearly reported.
```

## Field Guidance

- Scope: name real directories or modules. Avoid "check the whole repo" unless the task is a repo-wide audit and the workstream owns a specific lens.
- Ownership: list exact writable files or directories. One writer per file.
- Evidence: prefer source paths, symbols, stable line numbers, commands, and summarized outputs.
- Completion criteria: define what `completed`, `blocked`, or `inconclusive` means before launch.
