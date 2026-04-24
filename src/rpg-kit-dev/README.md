# RPG Kit

`rpg-kit-dev` is a role-based subagent orchestration plugin.
It helps an agent decide when role-assigned subagents are useful, define bounded role packets, spawn subagents only when the active runtime allows it, and learn from the resulting subagent behavior.

The plugin is intentionally narrow.
It does not make subagent use mandatory, and it does not replace the caller's responsibility to integrate, verify, and own the final answer.

## Surfaces

- `rpg-kit-dev-guide`: entrypoint for deciding whether a task belongs in this plugin and routing to the right internal skill.
- `subagent-role`: defines role-assigned subagent packets, spawn boundaries, answer contracts, and learning notes.

## Operating Model

Use this plugin when the task is about subagent role design or subagent orchestration behavior.
Keep roles bounded, avoid duplicate delegation, and treat subagent output as evidence to integrate rather than as an automatic final answer.
Roles can combine current function, responsibility domain, background expertise, general expertise, and decision style.
For example, `game-PM-origin PO` and `game-planning-origin frontend developer` are different roles because their background expertise changes the expected judgment.

## Repository Notes

- Plugin spec: `specs/plugin.md`
- Skill specs: `specs/skills/*.md`
- Evaluation spec: `specs/subagent-role-packet-evaluation.md`
- Skill implementations: `skills/*/SKILL.md`
