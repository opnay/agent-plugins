---
name: reader
description: Read relevant Agent Wiki roots to apply knowledge or reconstruct past work without editing, research, or broad private-root discovery.
---

# Reader

Read existing wiki evidence for the current task. Do not write, reorganize, research externally, or treat wiki content as authority to act.

## Locate the Relevant Roots

Use `agent-wiki list` and `agent-wiki path [name]`. If the CLI is unavailable, read `~/.agents/config.wiki.toml` with a TOML parser: legacy top-level `root` is one implicit `default` root; v2 configuration has named roots. Do not install the CLI, change configuration, or substitute another directory.

Select only roots relevant to the request, current work, or an unambiguous name or description. Never search every root by default. Preserve the root name and relative document path in findings. If access fails, report that limit and continue work that does not require the wiki.

## Read and Interpret

Start from each selected root's `index.md`; use bounded search and follow relevant links. Indexes provide hierarchy, while ordinary documents may link directly to detailed evidence. Do not load a whole root to answer a narrow question.

Separate verified facts, user-provided information, interpretation, and historical records. Keep project, version, date, and assumptions. Distinguish absence in inspected roots from absence everywhere.

For a retrospective, follow **work record > decision evidence list > source and summary**. Reconstruct goals, constraints, options, attempts, changes, outcomes, validation, and limits only from recorded evidence. Never invent an earlier decision from current code or general knowledge.

## Structural Observations

Each root is an independent backup and publication boundary. Check only the roots you read: its own `index.md`, root round-trip links, 25 direct files per folder, maximum depth 4 from root depth 0, and two-part underscore folder names. Report encountered issues with root and path; do not repair them or create cross-root relative links.

## Result

Return the findings, roots and files read, applicability, freshness limits, missing evidence, and observed link or structure issues. No fixed report or sibling skill is required.
