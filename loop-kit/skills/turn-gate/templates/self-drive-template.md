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

# Sequence Contract

- Objective:
- Prepared flow sequence:
- Current flow:
- Acceptance signal:
- Verification expectation:

# Autonomous Boundary

- Allowed autonomous actions:
- Prohibited autonomous actions:
- Approval-sensitive checkpoints:

# Endpoint Handling

- Endpoint:
- Sequence exhaustion behavior:
- Repeat policy:
- Handoff target:
- Blocker return condition:
- Last confirmed flow or timestamp:

# Progress Ledger

-

# User-Gated Return Conditions

- Scope, endpoint, target, flow order, non-goal, or acceptance signal ambiguity.
- Approval-sensitive action outside the recorded boundary.
- Blocker, repeated critical failure, inaccessible records, or current-flow identity mismatch.
- `active_flow_index` greater than or equal to `planned_flow_count`.

# Residual Risk

-
