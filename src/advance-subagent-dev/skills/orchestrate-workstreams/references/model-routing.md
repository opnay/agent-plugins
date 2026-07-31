# Model and Role Routing

## Default Route

Use this default spawned worker:

```yaml
model: gpt-5.6-terra
reasoning_effort: xhigh
```

Assign a task-packet role:

| Role | Responsibility | Access |
| --- | --- | --- |
| `EXPLORE_READ` | Research, discovery, logs, source and dependency mapping | Read-only |
| `IMPLEMENT_OWNED` | Implementation or action inside a disjoint owned surface | Write-enabled |
| `REVIEW_LENS` | Correctness, security, performance, quality, or counterevidence review | Read-only |
| `PROCESS_STRUCTURED` | Schema-bound extraction, transformation, classification, test generation, or repetitive mechanical work | Usually read-only; write-enabled only with disjoint output ownership |

Use Terra xhigh for `PROCESS_STRUCTURED`. Do not define a Luna route or claim that different model families are equivalent.

## Frontier Judgment

Use `gpt-5.6-sol`, `xhigh` with `FRONTIER_JUDGMENT` only when at least one condition holds:

- the goal requires ambiguous framing;
- a shared contract or architecture must be designed;
- strong evidence conflicts;
- error cost is high and no deterministic verification exists;
- an independent frontier-level final audit is justified.

Do not select Sol because work is long or large. When deterministic verification is sufficient, the main agent must verify Terra results directly.

Sol does not bypass the dispatch gate. Add it only as a conditional dependent audit node in a lifecycle that already passed the gate, after its prerequisites complete. Do not start a new lifecycle to dispatch one Sol auditor.

## Cost and Availability

- Use few workers and large coherent batches.
- Put strict schemas, deterministic checks, bounded retries, and explicit stop conditions in packets.
- Do not compensate for mechanical volume with more workers.
- If the selected model is unavailable, preserve the role contract with an available model or use `DIRECT`.
- Disclose substitute models, omitted audits, and verification limits.
