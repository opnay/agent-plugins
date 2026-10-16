# Model Profiles and Role Routing

## Model Profiles

Choose only `fast`, `default`, or `high`. A profile is an indivisible model-and-effort setting, separate from role and authority.

| Profile | Binding | Use when instructions do not specify a choice |
| --- | --- | --- |
| `fast` | Exact model and effort pair from applicable instructions | Narrow, repeatable work with direct verification |
| `default` | Inherit; omit both model and reasoning effort from spawn | No justified adjustment to the inherited setting |
| `high` | Exact model and effort pair from applicable instructions | Complex reasoning, conflicting evidence, or high error cost |

Follow applicable task or repository profile selections and defaults first. Otherwise select by task demands, uncertainty, verification, and cost. Length or file count alone does not justify `high`. Do not bind roles permanently to profiles.

Resolve concrete pairs from current user or project instructions; do not maintain model IDs, family rankings, or default effort values here. If instructions use combined shorthand, match it to a supported model ID and effort. Do not send the combined shorthand as a model ID or guess an ambiguous split.

Before spawning, check that the selected pair is defined and supported by the current tool. Keep both values together. Do not mix profiles, adjust effort independently, or choose an unconfigured combination. A decision model may advise only among defined, available profiles; a new pair requires an explicit instruction redefining a profile.

If a pair is undefined, unavailable, or has unsupported effort, do not silently repair it. Use another instruction-permitted profile, inheritance, or direct execution only when it meets the role contract. A mandatory profile must not be substituted without authorization. If no valid route remains, report the missing definition or capability. Do not install tools or edit configuration to create a route.

Record the selected profile, why it fits, and either the explicit pair or `inherited`. Do not guess undisclosed inherited values. Report substitutions, unexecuted work, and validation limits.

## Role Contracts

| Role | Responsibility | Access |
| --- | --- | --- |
| `EXPLORE_READ` | Research, discovery, logs, source and dependency mapping | Read-only |
| `IMPLEMENT_OWNED` | Implementation or action inside a disjoint owned surface | Write-enabled |
| `REVIEW_LENS` | Correctness, security, performance, quality, or counterevidence review | Read-only |
| `PROCESS_STRUCTURED` | Schema-bound extraction, transformation, classification, or repetitive mechanical work | Read-only unless disjoint write ownership is explicit |
| `FRONTIER_JUDGMENT` | Ambiguous framing, shared architecture, strong evidence conflicts, consequential judgment without deterministic verification, or justified independent final audit | Read-only unless separately authorized |

Roles describe responsibility, not model families. Deterministic verification remains the lead's job when sufficient. Neither `high` nor `FRONTIER_JUDGMENT` bypasses the dispatch gate. A separate judgment audit is a dependent node in an already-qualified lifecycle; wait for its prerequisites. Do not start a new lifecycle for a lone auditor.

## Coordination Cost

Use few workers and coherent batches. Apply strict schemas, deterministic checks, bounded retries, and explicit stop conditions. Do not compensate for mechanical volume by adding workers. Preserve role requirements when selecting or substituting a profile.
