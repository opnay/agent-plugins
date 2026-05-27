---
turn_gate_active: yes
active_flow: `{count-pad3-flow-slug-or-none}`
next_action: `{next-action}`
terminal_summary_allowed: no
explicit_turn_end_available: yes
self_drive: inactive
self_drive_sidecar: none
unapproved_actions: []
active_skills: []
---

# Turn Gate Plan

## Recent Requests

- [current] `{compact-current-request-or-routing-signal}`

## Flow Index

- [active] `{count-pad3-flow-slug-or-none}`
- [recent] `{previous-flow-slug-or-none}`
- [archive] `{older-flow-range-or-none}` are recoverable from individual flow records

## Continuity Note

- [note] `{what the next agent must know after compaction or interruption}`
