---
status: active
mode: finite | infinite
active_flow_record: {record-path-or-id}
next_action: {required-next-turn-flow-or-message-action}
progress_note: {compact-current-progress}
blocker_state: none | {blocker-summary}
blocker_impact: none | {acceptance|verification|approval|access|external|user-input|internal-repair}
flags: [turn_gate_active, terminal_summary_blocked]
# finite only
active_flow_index: {index}
current_flow_label: {label}
planned_flow_count: {count}
# infinite only
loop_count: {start-at-1-and-increment-after-verified-handoff}
current_loop_label: {label}
---

# Turn Gate Self-Drive Sidecar

## Goal

```text
{source-backed-goal-or-user-message}
```

- endpoint: `{endpoint-or-stop-condition}`
- acceptance: `{acceptance-signal}`
- verification: `{verification-expectation}`
- boundary: allow `{allowed-autonomous-actions}`; checkpoint `{approval-sensitive-actions}`

## Sequence

1. `{flow-label}`: `{scope}`; endpoint `{endpoint}`; verification `{expectation}`

Omit `Sequence` for `mode: infinite` unless the current iteration needs a short local checklist. Frontmatter owns `loop_count` and `next_action`.

## Ledger

- `{timestamp}` `{identity}`: `{material-update}`
- `{append material updates; do not replace history with only the current summary}`

## Handoff

- next: `{handoff-condition-or-next-action}`
- advance: `{verification pass; not blocked; next flow identity known; approval boundary unchanged}`
- position: `{keep current index or loop_count until advance is confirmed}`
