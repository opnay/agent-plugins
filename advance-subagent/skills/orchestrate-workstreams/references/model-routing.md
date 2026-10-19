# Model Profiles and Roles

## Profile Policy

Choose only fast, default, or high. Profiles are indivisible model-and-effort settings, separate from responsibility and permission.

| Profile | Binding | When no applicable instruction selects one |
| --- | --- | --- |
| fast | Exact model and effort pair from applicable instructions | Narrow, repeatable work with direct verification |
| default | Inherit; omit both spawn settings | No justified adjustment |
| high | Exact model and effort pair from applicable instructions | Difficult inference, conflicting evidence, or high error cost |

Follow explicit task or repository profile choices and defaults first. Length and file count alone do not justify high. Do not hard-code model families or bind roles permanently to a profile.

Resolve combined shorthand against the current tool's supported model IDs and efforts. Never send the combined shorthand as a model ID, split an ambiguous value by guess, mix profiles, or tune effort separately.

Check availability before spawning. When a pair is undefined or unsupported, use another instruction-permitted profile or inheritance only if the role can still be met. Do not replace a mandatory profile without authority. If no valid route remains, report the unavailable capability and unfinished work.

Record the chosen profile, reason, and exact pair or inheritance. Do not guess undisclosed inherited values. Disclose substitutions and material omitted work.

## Roles

| Role | Responsibility | Access |
| --- | --- | --- |
| PLAN_ROUND | Propose a round from the minimal preserved handoff | Read-only |
| EXPLORE_READ | Research, discovery, sources, logs, dependency mapping | Read-only |
| IMPLEMENT_OWNED | Implementation within assigned disjoint ownership | Explicit write access |
| REVIEW_LENS | Correctness, security, performance, quality, counterevidence | Read-only |
| PROCESS_STRUCTURED | Schema-bound extraction, transformation, classification | Read-only unless isolated write access is explicit |
| FRONTIER_JUDGMENT | Ambiguous framing, conflicts, consequential judgment, independent audit | Read-only unless separately authorized |

All roles return evidence under the task contract. Only the main agent spawns and integrates. A planner or worker never delegates, expands scope, or overrides user confirmation.

A single judgment or verification task can be a branch. Check its prerequisites and value; do not manufacture another task to meet a worker-count rule.

A decision adviser may compare only currently defined, available profiles. Advice cannot define a new pair, supply permission, or waive role requirements.
