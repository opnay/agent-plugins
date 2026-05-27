---
phase: intake
verification_status: not-started
next_action: prepare_flow_contract
flags: [turn_gate_active, terminal_summary_blocked]
answered_question: none
# pending_question: `{omit unless a question is waiting}`
continuity: update_before_reporting_or_next_flow
---

# `{Flow Label}`

## Contract

- scope: `{in-scope}`
- exclude: `{out-of-scope}`
- done: `{completion-or-success-condition}`
- boundary: `{approval-and-handoff-boundary}`

## Risky Action

- `{omit this section unless approval-sensitive action exists}`
- `{action}`: target `{target}`; effect `{effect}`; risk `{risk}`; recovery `{path}`; approval `{not-requested|required|granted|blocked}`.
- `{readiness, verification, build, or generated release surface update is not commit, release, publish, version bump, destructive, or external-action authority}`

## Execution Log

- `[intake]` `{event}`

## Result

- status: `{pass|fail|blocked|insufficient|not-required}`
- evidence: `{verification-method-and-key-evidence-or-gap}`
- risk: `{risk-or-none}`
- next: `{next-action-or-explicit-stop-routing}`
