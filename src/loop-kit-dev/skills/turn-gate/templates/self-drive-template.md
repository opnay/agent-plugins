---
sequence_objective: ""
active_flow_index: 0
current_flow_label: ""
progress_note: ""
planned_flow_count: 0
endpoint: ""
status: "active"
last_updated_flow: ""
required_next_action: ""
---

# Self-Drive Sequence

## Sequence Contract

- Objective:
- Planned flows:
- Active flow index: 0-based machine field
- Current flow: must match `current_flow_label`; use a human-readable number/name/file or slug
- Progress note: current sequence summary; refresh before reporting and before next planned flow handoff
- Endpoint:
- Blocker return conditions:

## Autonomous Boundary

- Allowed autonomous actions:
- Prohibited autonomous actions:
- Approval-sensitive checkpoints:
- Approval notes:
- Recovery path:
- Stop boundary:

## Progress Ledger

- Append-only history of sequence transitions and material updates.

## User-Gated Return Conditions

- New risky action:
- Scope or non-goal change:
- Unclear endpoint:
- Current-flow identity ambiguity:
- Record access failure:
- Repeated critical failure or root blocker:
- Commit, push, PR, publish, release, or version bump outside exact approved boundary:

## Residual Risk

- TBD
