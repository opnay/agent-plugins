# AGENTS.md

---
## Part 1. Repository And Marketplace Rules
Applies to repository structure, plugin placement, and marketplace metadata changes.
---

## Repository Role

This repository is the root for a local plugin marketplace and a working area for harness engineering.
Treat this file as the operating guide for both repository layout and harness-quality expectations.

## Repository Layout

- Each local plugin lives directly under the repository root.
- Current example:
  - `./teammate-kit`
  - `./frontend-engineering-kit`
  - `./workflow-kit`
- Do not create or use `./plugins/<plugin-name>` in this repository unless the repository structure is intentionally changed later.

## Required Plugin Structure

Each plugin should include:

- `./<plugin-name>/.codex-plugin/plugin.json`
- Optional `./<plugin-name>/skills/`
- Optional `./<plugin-name>/assets/`
- Optional `./<plugin-name>/scripts/`
- Optional `./<plugin-name>/.mcp.json`
- Optional `./<plugin-name>/.app.json`

The plugin folder name and `plugin.json` `"name"` must match.

## Marketplace Source Of Truth

- Marketplace file: `./.agents/plugins/marketplace.json`
- Every plugin added to this repository should have a matching entry in that file.
- In this repository, marketplace `source.path` values should point to repo-root plugin folders:
  - `./teammate-kit`
  - `./frontend-engineering-kit`
  - `./workflow-kit`
- Do not register repo-local plugins as `./plugins/<plugin-name>` in this marketplace.

## Plugin Change Workflow

1. Create or move the plugin folder at the repository root.
2. Ensure `.codex-plugin/plugin.json` exists and is valid JSON.
3. Add or update the matching entry in `./.agents/plugins/marketplace.json`.
4. Keep `policy.installation`, `policy.authentication`, and `category` present on every marketplace entry.
5. Validate edited JSON files after changes.

## Plugin Entry Skill Guidance

For plugins with multiple user-facing skills, prefer adding one entrypoint skill.

Use an entrypoint skill when:

- the plugin contains two or more skills with different roles
- users or agents may need help choosing the right workflow or domain skill
- the plugin benefits from task-shape classification before narrower skill selection

Entrypoint skill expectations:

- classify the task before deeper execution
- guide toward the right mode, domain, or workflow without hard-coding fragile dependencies
- remain independently usable as a first-stop skill
- avoid requiring every plugin to have one when the plugin is intentionally single-purpose

Rule of thumb:

- if a plugin grows beyond one clear skill, add an entrypoint skill unless there is a strong reason not to

## Repository Editing Rules

- Do not silently introduce a second layout convention.
- Preserve existing marketplace ordering unless reordering is explicitly requested.
- Prefer small, direct edits over scaffold regeneration when only metadata or paths need adjustment.
- If a plugin is moved, update marketplace paths in the same change.
- If a scaffold tool generates `./plugins/<plugin-name>`, move it to `./<plugin-name>` before finishing.

---
## Part 2. Harness Engineering Rules
Applies to harness design, implementation, validation, and failure handling.
---

## Harness Engineering Purpose

Use this repository with a harness-engineering mindset.
Optimize for reproducibility, diagnosability, explicit contracts, and safe iteration rather than clever one-off fixes.

## Harness Working Principles

- Prefer deterministic fixtures over live external dependencies.
- Make failures easier to debug, not easier to hide.
- Keep changes narrow, attributable, and easy to verify.
- Treat observability as part of the harness, not optional follow-up work.
- Prefer explicit contracts for inputs, outputs, and failure modes.

## Harness Change Workflow

1. Identify the exact harness behavior, contract, or workflow being changed.
2. Read the relevant scripts, configs, fixtures, docs, and tests before editing.
3. Make the smallest complete change that resolves the target problem.
4. Run the narrowest meaningful verification that matches the risk.
5. Report what changed, how it was validated, and what remains uncertain.

## Harness Design Rules

- Prefer deterministic inputs, fixtures, and seeded randomness.
- Avoid hidden state between runs.
- Make retries, timeouts, and backoff explicit and justified.
- Separate infra failure, harness failure, and assertion failure clearly.
- Prefer structured logs or machine-readable outputs over ad hoc prints when the harness is expected to be replayed or triaged.
- Keep setup and teardown explicit.
- If a workflow cannot be made deterministic, document the unstable boundary and constrain its blast radius.

## Validation Rules

- Config-only change:
  - Validate parsing or load behavior and run one representative path if the config affects execution.
- Fixture or sample-data change:
  - Re-run only the scenarios that depend on those fixtures.
- Core harness logic change:
  - Run targeted tests plus one representative end-to-end or integration path when available.
- Logging or reporting change:
  - Verify both machine-readable output shape and operator readability.

## Failure Handling

- Do not patch flaky behavior with blind retries unless the root cause is understood and the retry is part of the intended contract.
- Do not depend on wall-clock sleeps when a state check, event, or assertion can be used instead.
- When a failure cannot be reproduced locally, record the exact boundary of uncertainty.
- If external systems are involved, isolate which part is nondeterministic before changing harness logic.

## Do Not

- Do not introduce randomness without a seed or a clear reason.
- Do not add silent fallbacks that conceal broken harness behavior.
- Do not mix unrelated cleanup into harness-fix changes.
- Do not weaken assertions just to get a passing run.
- Do not change output contracts casually without updating the callers or documenting the break.

## Agent Output Expectations

When finishing harness-related work, include:

- `Scope handled`
- `Files changed`
- `Verification`
- `Residual risk`
