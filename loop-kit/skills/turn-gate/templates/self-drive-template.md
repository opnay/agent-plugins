# Turn Gate Self-Drive

## Sequence Contract

- status: active
- objective: `{objective}`
- endpoint: `{endpoint}`
- acceptance signal: `{acceptance-signal}`
- verification expectation: `{verification-expectation}`
- allowed autonomous actions: `{allowed-actions}`
- prohibited autonomous actions: `{prohibited-actions}`
- approval-sensitive checkpoints: `{checkpoints}`
- blocker return conditions: `{conditions}`

## Current Position

- active_flow_index: `{zero-based-or-one-based-index}`
- current_flow_label: `{label}`
- planned_flow_count: `{count}`
- required_next_action: `{action}`
- blocker_state: none

## Prepared Flow Sequence

1. `{flow-label}`: `{scope}`; endpoint `{endpoint}`; verification `{expectation}`

## Progress Ledger

- `{timestamp}` `{flow-label}`: `{material-update}`

## Next Handoff

- `{handoff-condition-or-next-flow}`
