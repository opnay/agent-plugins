# Round State and Handoff

Keep main-owned execution records in the user-specified destination. Otherwise use the applicable project's task-scoped temporary location; if none is defined, use the environment's temporary location. Preserve source permissions and disclosure boundaries.

Use existing record formats when available. The logical fields below matter; no parser, database, script, or fixed file layout is required.

## Whole-Run State

Maintain:

- original user request, subsequent user conditions, and applicable instructions;
- authority, exclusions, and pending confirmations;
- current interpretation, assumptions, working goal, and acceptance criteria, marked known, provisional, or unresolved;
- each condition change with evidence, reason, and user-confirmation status;
- verified achievements and recoverable source and artifact locations;
- counterevidence, unresolved issues, and their impact;
- current round and actual assignments, ownership, required results, and verification state;
- user or environment allowance and remaining limits; record unspecified limits as unspecified rather than inventing an allocation;
- lifecycle-wide follow_up_used, initially false, retained through flush.

A later user message supplements existing conditions unless it changes or invalidates them. Conditions are not dropped because the user did not repeat them.

## Round Plan

Record the round purpose and its partition, alternatives, or mixed branch shape. For each node, record its local objective, prerequisites, inputs, output and consumer, domain contract, access and ownership, evidence and checks, and completion and stop conditions.

Record common comparison or integration criteria and which overall uncertainty or unmet condition the round can resolve. Global acceptance criteria need not be settled to assign a local task.

The first plan is main-authored. After flush, a fresh read-only planner proposes the next plan; the main agent validates authority, dependencies, shared state, profiles, capacity, and useful progress before spawning workers.

## Audit Record Versus Planner Input

Keep exact packets, actual agent IDs, results, source checks, rejected claims, integration actions, and status transitions in audit records. These records preserve traceability but are not forwarded wholesale to the next planner.

Prepare a minimal handoff with:

1. original request and relevant follow-up conditions;
2. applicable instructions, authority, exclusions, and confirmation dependencies;
3. current goal and criteria, assumptions, changes and verified reasons;
4. verified sources and artifacts, their locations and verification status;
5. accepted findings, retained counterevidence, chosen direction and evidence-based rationale;
6. unresolved questions, failed or omitted checks, and uncertainties;
7. remaining allowance and lifecycle recovery state;
8. the unresolved need to address next, without prescribing an unsupported conclusion.

Distinguish verified findings, hypotheses, and unknowns. Provide original sources for reinspection rather than replacing them with old workers' conclusions. Do not include private reasoning, whole conversations, unfiltered worker narratives, or a preferred answer.

## Flush Procedure

1. Collect required terminal results, or stop outstanding work and verify writes have ceased.
2. Verify and integrate findings; review current conditions and whole-request completion.
3. If complete, deliver the result and stop. No continuation flush is needed.
4. If unfinished with a useful permitted next action and allowance, persist the state and minimal handoff.
5. Spawn a new read-only planner without inherited conversation. Record the actual new agent ID.
6. Let it inspect the handoff's necessary sources and propose branches and questions. It must not spawn or edit.
7. The main agent checks the plan and starts fresh workers under new packets.

The main session remains active. Flush neither deletes artifacts nor resets the main agent. A summary or reused agent cannot substitute for a new planner. If fresh execution is unavailable, disclose that limitation and distinguish any permitted fallback output from an executed loop.

## Completion and Stop Review

| State | Required evidence or next action |
| --- | --- |
| Complete | Whole-result checks satisfy accumulated applicable conditions and the attached domain contract. |
| Continue | Identify a concrete next task that can change the result and confirm remaining allowance and authority. |
| Needs user input | Identify the material unresolved choice or permission change; hold dependent work and continue independent authorized work if useful. |
| Blocked | Identify unavailable tools, inputs, authority, capacity, or exhausted allowance and unfinished scope. |
| Inconclusive | Explain insufficient evidence or why no useful justified check remains. |

Do not treat a worker's success, an adviser recommendation, an allowance limit, or a criterion weakened without confirmation as completion.

Use a decision model only for an action-changing uncertainty. Give verified evidence, constraints, permitted choices, and relevant limits; ask for an action and a short evidence-linked rationale. Check observable facts directly, keep dependent questions sequential, and batch only independent questions with the same state. No external adviser is required; external transmission requires existing authority.

## Example: An Underspecified Performance Request

For "make search faster," the first round can inspect affected flows, measurements, and existing requirements. The integrated finding may identify a slow filter and leave the target unresolved.

The next fresh planner can propose bottleneck analysis, measurement design, and criterion clarification. The main agent obtains a necessary user decision without inventing a target.

A later round compares isolated query and index approaches. Preserve correctness regressions and valid evidence from both approaches. Another fresh plan then targets the discovered regression and final checks.

Each checkpoint carries current criteria and their origins; each round is redesigned from verified results rather than replaying the first partition.
