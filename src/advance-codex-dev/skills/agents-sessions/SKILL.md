---
name: agents-sessions
description: Define the basic purpose and turn-gate-aligned structure of the `.agents/sessions` folder. Use when you need to explain when date-scoped agent artifacts belong under `.agents/sessions/{YYYYMMDD}`, how the `000-plan.md` and `001-*` flow record skeleton works, and how to distinguish operational session notes from repository source changes.
---

# Agents Sessions

## Overview

Use this skill to reason about the `.agents/sessions` folder itself.
It defines the folder's role as a place for session-scoped operational artifacts created while agent work is in progress.
It does not create, update, validate, or migrate specific records.

## Use When

- You need to explain what `.agents/sessions` is for.
- You need to decide whether a note, plan, change record, scratch artifact, or retrospective belongs in a session directory.
- You need to separate session-scoped operational artifacts from repository source files, specs, release surfaces, or reusable plugin assets.
- You need to understand the default `000-plan.md` and `001-*` flow record skeleton.

## Do Not Use When

- You need to create or update turn-gate flow records while an active turn is running.
- You are designing a plugin, skill, custom agent, commit workflow, or reusable tool-use policy.

## Folder Contract

- `.agents/sessions` is for date-scoped operational artifacts created during agent work.
- Session directories use the date-based shape `.agents/sessions/{YYYYMMDD}/`, matching the turn-gate convention.
- A session directory may hold operational notes, plans, change records, verification notes, retrospectives, or scratch artifacts that help continue or audit work from that date.
- The default skeleton is `000-plan.md` plus zero-padded flow records such as `001-scope-review.md` at the date directory root.
- Other session artifacts should be placed relative to that skeleton and must not introduce a parallel session record schema as a competing default.
- Session artifacts are not the source of truth for plugin behavior, skill contracts, release metadata, or product code.
- Repository source changes belong in the relevant source tree, such as `src/<plugin-name>-dev`, not only in a session directory.

## Routing

- Use `loop-kit:turn-gate` when the work needs active turn continuity, flow records, or next-flow reopening.
- Use `skill-creator`, `plugin-creator`, `tool-use-guide`, `subagent-creator`, or `git-committer` when the primary artifact is one of those reusable surfaces.
- Keep this skill at the folder-boundary level. Do not absorb turn-gate flow operation rules here.

## Output

- `Folder purpose`
- `Belongs in .agents/sessions`
- `Does not belong there`
- `Default plan/flow skeleton`
- `Next skill if active flow operations are needed`
