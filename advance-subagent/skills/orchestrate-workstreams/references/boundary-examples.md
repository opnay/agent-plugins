# Boundary Examples

| Request shape | Automatic selection | Decision |
| --- | --- | --- |
| Investigate independent software modules or execution paths | Conditional `DISPATCH` | Dispatch only when each lane has its own question, scope, evidence, and immediate start. |
| Apply independent review lenses to the same code | Conditional `DISPATCH` | Dispatch distinct read-only deliverables such as correctness, security, and performance reviews. |
| Combine market research with prototype implementation | Conditional `DISPATCH` | Dispatch when both lanes can start from current inputs and have separate evidence and deliverables. If implementation depends on the research conclusion, preserve the graph dependency and run sequentially. |
| Process large schema-bound datasets | Apply the gate | Assign independent partitions to large coherent `PROCESS_STRUCTURED` batches using a permitted profile. Use `DIRECT` for one extraction lane. |
| Implement fully separate owned surfaces | Conditional `DISPATCH` | Use `IMPLEMENT_OWNED` with a permitted profile only after fixing the shared contract and proving disjoint writable ownership. |
| Resolve conflicting high-quality evidence without a deterministic check | A separate judgment audit may qualify | After the main agent checks the conflict, add `FRONTIER_JUDGMENT` with a permitted profile only as a dependent node in a lifecycle that already passed the gate. Never start a dispatch for a lone auditor. |
| Resolve one sequential root cause | `DIRECT` | Do not parallelize work that depends on one shared cause. |
| Produce a pure multi-source research report | Exclude from automatic trigger | This skill does not automatically own source methodology and evidence-traceable reporting. |
| Perform generic implementation, review, or testing | `DIRECT` | Do not dispatch software-only work unless independent workstreams are evident. |

An explicit subagent request may enter the gate for an automatically excluded shape, but it does not expand the skill boundary or guarantee spawn. The plugin's `deep-research` skill is an optional reference boundary for pure research, never a runtime prerequisite.
