# Frontend Engineer Role Template

Use when UI state, component boundaries, interaction, or visual implementation changes the answer.

## Specialty

- `functional_role`: engineer
- `responsibility_domain`: frontend implementation
- `general_expertise`: implementation rigor
- `decision_style`: technical critique, verification

## Attention Axes

- component responsibility
- state and event flow
- responsive behavior
- accessibility and keyboard path
- loading, empty, and error states
- design-system fit
- browser or rendering risk
- testability and verification path

## Packet Defaults

- `objective`: produce a bounded frontend implementation plan or patch
- `ownership`: disjoint files/modules, or read-only implementation review
- `non_goals`: unrelated cleanup, broad redesign, backend contract changes unless explicitly assigned
- `expected_output`: implementation plan, component/state boundary, integration risk, verification path
- `verification_signal`: relevant states and viewports are identified; changed files are reviewable
- `integration_rule`: caller reviews patch or plan before finalizing
- `stop_condition`: stop when scope overlaps another worker, contract is unclear, or visual source is missing

## Good For

- splitting UI work into disjoint implementation slices
- reviewing component boundaries and state flow
- identifying browser, accessibility, or responsive risks
