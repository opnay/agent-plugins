# Role Templates

Use this reference when shaping a `subagent-role` packet.
These are reusable role templates for common subagent roles.
Treat them as starting points: copy the relevant axes into a task-specific role packet, then narrow objective, ownership, output, and stop condition for the actual task.

## Template Index

- `product-pm.md`: product judgment, priority, player/user impact
- `po.md`: requirements, acceptance, scope, sequencing
- `frontend-engineer.md`: UI implementation, component/state boundary, browser-facing verification
- `ux-reviewer.md`: affordance, hierarchy, interaction, material UX findings
- `qa-regression.md`: regression surface, test matrix, pass/fail signals
- `researcher.md`: bounded read-only investigation and evidence gathering

## Usage Rules

- Start from one template only when its attention axes change the expected judgment.
- Combine templates only by naming the primary role and adding a background expertise overlay.
- Remove axes that do not affect the current task.
- Keep the role packet bounded enough for the subagent to answer without rediscovering the whole project.
- Treat template output as caller-integrated evidence, not the final answer.

## Shared Axes

Every role packet should make these explicit when relevant:

- `objective`: what decision, artifact, or evidence the role should produce
- `ownership`: the read-only question or disjoint write scope the role owns
- `integration_signal`: how the caller will use the result
- `verification_signal`: what would make the result trustworthy
- `stop_condition`: when the role should stop instead of widening scope
- `learning_note`: what the caller should observe about this role's usefulness

## Packet Skeleton

Use this skeleton after choosing a template:

```yaml
role_name: <task-specific role name>
role_specialty:
  functional_role: <from template>
  responsibility_domain: <task-specific domain>
  background_expertise: <optional overlay>
  general_expertise: <from template or task>
  decision_style: <from template or task>
objective: <one bounded decision, artifact, or evidence target>
context: <only the context needed to answer>
ownership: <read-only question or disjoint write scope>
non_goals: <what this role must not widen into>
expected_output: <structured output fields>
verification_signal: <how the caller can trust or check the result>
integration_rule: <how the caller will use the answer>
stop_condition: <when the role should stop and report a gap>
```

## Background Expertise Overlays

Add a background overlay only when it changes judgment.

- `game-planning-origin`: player motivation, loop fantasy, progression, scenario coherence
- `game-PM-origin`: product metrics, retention, monetization/product tradeoffs, player segment risk
- `frontend-origin`: implementation ergonomics, UI state, rendering constraints, accessibility
- `infra-origin`: operational risk, observability, deployment boundary, failure modes
- `harness-origin`: determinism, fixture quality, debuggability, assertion clarity

Do not stack overlays just to make a role sound specific.
If the overlay does not change the expected output, omit it.
