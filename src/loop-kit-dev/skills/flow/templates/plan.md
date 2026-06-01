---
flow_plan_active: yes
active_flow: `{count-pad3-flow-slug-or-none}`
next_action: `{next-action}`
handoff_condition: `{handoff-or-next-intake-condition}`
approval_boundary: `{approval-boundary}`
verification_expectation: `{verification-expectation}`
active_skills: []
---

<!--
File name: 000-plan.md
Location: .agents/sessions/{YYYYMMDD}/000-plan.md
Use one plan file per session date.
-->

# Flow Plan

## Recent Requests

- [current] `{compact-current-request-or-routing-signal}`

## Purpose

- [objective] `{current objective or purpose chain when it affects scope, acceptance, verification, approval, or handoff}`

## Flow Index

- [active] `{count-pad3-flow-slug-or-none}`
- [planned] `{ordered-flow-candidates-or-none}`
- [recent] `{previous-flow-slug-or-none}`
- [archive] `{older-flow-range-or-none}` are recoverable from individual flow records

## Continuity Note

- [note] `{what the next agent must know after compaction or interruption}`
