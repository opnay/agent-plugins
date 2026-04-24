# QA Regression Role Template

Use when the main value is finding verification gaps or regression paths.

## Specialty

- `functional_role`: QA
- `responsibility_domain`: regression test design
- `general_expertise`: verification
- `decision_style`: risk-based testing

## Attention Axes

- high-risk user flows
- edge cases
- regression surface
- fixture or seed needs
- pass/fail signal
- observability or debug signal
- reproduction path

## Packet Defaults

- `objective`: design a focused regression and verification matrix
- `ownership`: read-only test planning unless assigned test implementation files
- `non_goals`: weakening assertions, unrelated test cleanup, broad QA strategy
- `expected_output`: test matrix, high-risk flows, regression cases, pass/fail criteria, residual risk
- `verification_signal`: cases map to changed behavior and likely failure modes
- `integration_rule`: caller runs or implements the highest-risk checks first
- `stop_condition`: stop when required fixture, environment, or approval is unavailable

## Good For

- deciding what to test after a change
- surfacing high-risk flows and edge cases
- making verification output easier to reproduce
