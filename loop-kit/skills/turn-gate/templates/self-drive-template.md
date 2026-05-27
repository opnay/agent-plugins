---
status: active
active_flow_index: {index}
current_flow_label: {label}
active_flow_record: {record-path-or-id}
planned_flow_count: {count}
next_action: {required-next-action}
progress_note: {compact-current-progress}
blocker_state: none | {blocker-summary}
blocker_impact: none | {acceptance|verification|approval|access|external|user-input|internal-repair}
flags: [turn_gate_active, terminal_summary_blocked]
---

# Turn Gate Self-Drive

## Contract

- objective: `{objective}`
- endpoint: `{endpoint}`
- acceptance: `{acceptance-signal}`
- verification: `{verification-expectation}`
- repeat: `{cycle-boundary}; {limit-or-condition}; verification {per-cycle-verification}; stop {user-gated-stop-condition}`
- boundary: allow `{allowed-autonomous-actions}`; deny `{prohibited-actions}`
- approval: `{exact-action}; target {target}; effect {expected-effect}; risk {risk}; recovery {recovery-path}; scope {included/excluded}`
- blockers: `{blocker-return-conditions}`

## Sequence

1. `{flow-label}`: `{scope}`; endpoint `{endpoint}`; verification `{expectation}`

## Ledger

- `{timestamp}` `{flow-label}`: `{material-update}`
- `{append material updates; do not replace history with only the current summary}`
- report: `{history-preserved; new material update named}`

## Handoff

- next: `{handoff-condition-or-next-flow}`
- advance: `{verification pass; non-blocked handoff; next identity known; approval boundary matches}`
- index: `{keep current index until advance is confirmed}`
