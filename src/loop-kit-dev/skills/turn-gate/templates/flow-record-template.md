---
turn_gate_active: yes
phase: preparation
question_routing_mode: none
user_explicit_stop: no
terminal_summary_allowed: no
confirmed_closure: no
closure_source_message: none
closure_recorded_phase: none
pending_question_state: none
pending_question_id_or_summary: none
superseded_question_id_or_summary: none
verification_status: not-started
required_next_action: prepare flow contract
continuity_note: update this guard before reporting and next-flow routing
---

# `{Flow Label}`

## Flow Contract

- source request summary: `{summary}`
- raw request: `{raw-text-if-needed}`
- interpretation: `{interpretation}`
- scope: `{in-scope}`
- non-goals: `{out-of-scope}`
- acceptance signal: `{completion-or-success-condition}`
- verification expectation: `{expected-method-or-risk}`
- approval boundary: `{allowed-actions-and-approval-sensitive-checkpoints}`
- handoff condition: `{when-to-report-route-or-stop}`

## Optional Risky Actions

- action: `{none-or-action}`
- exact target: `{target}`
- expected effect: `{effect}`
- risk: `{risk}`
- recovery path: `{path}`
- included scope: `{included}`
- excluded scope: `{excluded}`
- endpoint: `{endpoint}`
- approval status: `{not-requested|required|granted|blocked}`

## Execution Log

- `[preparation]` `{event}`

## Verification

- method: `{clean-context|normal|not-required}`
- method reason: `{reason}`
- result_status: `{pass|fail|blocked|insufficient}`
- evidence: `{evidence-or-gap}`
- non-pass routing: `{repair|collect-more-evidence|blocker|none}`

## Report

- changed surfaces: `{files-artifacts-decisions}`
- verification status: `{status}`
- material judgment calls: `{calls}`
- required next action: `{action}`

## Next Flow Options

- continue: `{next-action}`
- blocked recovery: `{needed-input-or-access}`
- explicit turn-end: user may explicitly stop the turn

## Residual Risk

- `{risk-or-none}`
