# Turn Gate Plan

## Routing State

- date: `{YYYYMMDD}`
- turn_gate_active: yes
- active_flow: `{flow-record-path-or-none}`
- required_next_action: `{next-action}`
- user_explicit_stop: no
- terminal_summary_allowed: no
- confirmed_closure: no
- explicit_turn_end_available: yes

## Request History

- `{timestamp}` raw: `{raw-user-request-or-activation}`
  interpretation: `{compact-interpretation}`

## Flow Index

- `{count-pad3}` `{flow-label}`: `{state}`; record `{flow-record-path}`; required next action `{next-action}`

## Planned Flow Sequence

- `{selected-current-or-future-flow}`

## Flow Skill List

- `{flow-label}`: `{skill-name}` for `{usage-point}`

## Completed Flow Summaries

- `{flow-label}`: `{compact-summary}`; verification `{status}`; residual risk `{risk}`

## Self-Drive

- status: inactive
- sidecar: none
- active_flow_index: none
- current_flow_label: none

## Active Date-Level Risks

- `{risk-or-none}`

## Continuity Note

`{what the next agent must know after compaction or interruption}`
