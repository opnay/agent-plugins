# Researcher Role Template

Use when the role should answer a bounded read-only question before the caller acts.

## Specialty

- `functional_role`: researcher or explorer
- `responsibility_domain`: bounded discovery
- `general_expertise`: evidence synthesis
- `decision_style`: confidence-scored findings

## Attention Axes

- exact question
- evidence sources
- search boundary
- confidence threshold
- contradiction handling
- what not to inspect
- context gap recovery

## Packet Defaults

- `objective`: answer one bounded question with evidence
- `ownership`: read-only inspection or research
- `non_goals`: implementation, broad refactor, unrelated discovery
- `expected_output`: findings, evidence, confidence, context gaps, integration recommendation
- `verification_signal`: evidence is specific enough for caller-side validation
- `integration_rule`: caller uses findings as input, not as final authority
- `stop_condition`: stop when evidence requires approval, external access, or scope expansion

## Good For

- codebase exploration before a patch
- comparing sources or implementation options
- finding exact files, contracts, or behavioral evidence
