---
name: orchestrate-workstreams
description: >
  Run a fork-and-flush loop from branch design through execution, verification,
  completion review, and fresh-context replanning. Use for explicit orchestration
  or subagent requests, meaningful independent workstreams, alternative approaches,
  or fresh-context continuation, and as the execution engine for deep-research.
  Do not select for simple lookups, single-source summaries, generic implementation
  wording, complexity, or file count alone. subagent orchestration, fork-and-flush,
  parallel workstreams, fresh-context continuation
---

# Orchestrate Workstreams

Own one execution loop:

**Design branches → fork and execute → verify and integrate → review completion → if unfinished, preserve and flush → design branches again.**

The main agent remains responsible for authority, scheduling, evidence checks, integration, user questions, and the final response. There is no separate direct-execution mode inside this loop. One remaining task uses one worker; independent tasks may run in parallel.

## 1. Design This Round

Start from the original request and known instructions and permissions. The overall goal and acceptance criteria may be provisional or unknown. Understanding the request, discovering constraints, and establishing or revising completion criteria are valid branch tasks.

The main agent designs the first round. After a flush, a fresh read-only planner proposes the next round from preserved material; the main agent checks the proposal against current instructions, permissions, dependencies, and capacity.

For each task, define its question or objective, prerequisites, inputs, domain contract, access and ownership, deliverable, evidence, checks, and completion and stop conditions. Distinguish:

- **Partition:** different tasks whose accepted outputs are combined.
- **Alternatives:** different approaches to one objective, compared under common criteria.

Mix these forms when useful. Do not force parallelism or a fixed worker count. A dependent task starts only after its prerequisite is verified. Use sequential scheduling when work or shared state cannot be separated.

Read [Boundary Examples](references/boundary-examples.md) when selection or branch shape is unclear.

## 2. Fork and Execute

Before spawning, read [Execution Safety](references/delegation-safety.md), [Task and Result Contracts](references/task-result-contracts.md), and [Model Profiles and Roles](references/model-routing.md).

- Spawn planners and workers as new agents without inherited conversation. Explicitly disable conversation inheritance when supported, such as fork_turns: none. Give only the bounded packet, applicable instructions, and necessary sources or verified evidence.
- Only the main agent spawns. Planners propose plans; workers execute assigned tasks. Neither starts another orchestration loop or delegates.
- Use the minimum necessary agents within current capacity. One branch is valid.
- Default to read-only. Parallel writes require settled shared interfaces and fully disjoint ownership. Give each artifact one writer. Isolate alternative implementations; otherwise compare read-only or schedule sequentially.
- Keep shared contracts, integration surfaces, and execution state main-owned. New findings do not grant new write authority.
- Choose fast, default, or high using applicable profile definitions and task requirements. Preserve exact model/effort pairs; default omits both settings. Record the choice, reason, and exact pair or inheritance. Do not guess undisclosed inherited values.
- Record the actual returned agent ID and exact packet. Do not infer successful spawning, parallel execution, or fresh context.

If a required profile or execution capability is unavailable, use only instruction-permitted substitutions that meet the role. If none does, report the limitation and unfinished scope rather than claiming the loop ran.

## 3. Verify and Integrate

Collect every required result as completed, blocked, or inconclusive. Compare the packet with its claims, sources, inspected and changed surfaces, ownership, performed checks, skipped or failed checks, and uncertainty.

Reopen material sources and inspect actual artifacts. Worker conclusions are evidence, not established truth.

- Combine accepted partition outputs; resolve conflicts using sources, code, tests, logs, or direct execution.
- Compare alternatives by the agreed domain criteria and choose the next direction, not a truth by vote. Preserve valid evidence and counterevidence from unselected approaches.
- Keep the baseline or defer selection when evidence is insufficient.
- Allow one narrow follow-up to an existing agent across the entire run only for critical missing assigned scope. Set follow_up_used before requesting it and retain it through every flush. Do not relabel the same omission as a new task to bypass recovery limits.
- Put genuinely new work into a subsequent round within existing authority.

A decision model such as Jev may advise on uncertain branch design, approach selection, or the next action. Check observable facts directly. Give only authorized necessary information and defined choices; never use advice to invent facts, permissions, user preferences, or completion. No external service is required. Use local judgment and disclose a material fallback if the adviser is unavailable.

## 4. Review Conditions and Completion

Compare the integrated result with all accumulated applicable user conditions.

Keep original user requirements separate from current interpretations, assumptions, working goals, and acceptance criteria. Record evidence and reasons for changes. Clarifying a criterion or adding a necessary check is permitted within the request; changing goals, scope, authority, or weakening requirements needs user confirmation. Do not start dependent work before that confirmation. Continue independent authorized work while waiting.

Apply follow-ups to existing conditions, replacing only what the user changes or makes inapplicable. Do not weaken criteria to declare success.

- **Complete:** the whole result satisfies current applicable conditions and the domain checks. Report the outcome, evidence, performed and omitted checks, and material limits.
- **Continue:** a concrete action can change the result and remaining allowance permits it. Preserve and flush.
- **Needs user input, blocked, or inconclusive:** report the specific dependency, unavailable capability, evidence gap, exhausted allowance, or absence of a useful next check. These are not completion.

A worker's completion, an adviser's stop recommendation, or reaching a limit does not establish whole-request completion. Do not repeat unchanged inputs and failures without a new justified check.

## 5. Preserve and Flush

Read [Round State and Handoff](references/round-state-and-handoff.md) before the first checkpoint.

Confirm required workers have ended, or explicitly stop them and verify writes have ceased before transferring ownership. Preserve original and follow-up requests; applicable instructions and authority; current criteria, changes, reasons, and confirmation status; verified sources and artifacts; retained counterevidence; unresolved issues; remaining allowance; and lifecycle recovery state.

Keep audit records separately from the next planner's input. Give a new read-only planner only the minimal handoff and necessary original or verified material. Do not pass previous conversations, whole worker narratives, speculative claims as facts, or the main agent's preferred conclusion.

The planner returns a proposed round plan and questions, never a spawn. The main agent validates it and starts fresh workers.

Flush replaces planning and worker context, not the main session. Do not delete sources, artifacts, logs, or the current conversation. Reusing or interrupting an agent, or summarizing within the current session, is not fresh execution.

## Domain Composition

This loop does not define research methodology or code quality. Attach the task's domain contract explicitly. For research, the lead uses $advance-subagent:deep-research; assigned researchers apply its evidence and result requirements without restarting the loop. Mixed requests may attach different contracts to different branches and must preserve their dependencies and permissions.

The loop is usable without sibling skills. The research skill explicitly depends on this loop. If tools cannot provide actual fork or fresh replanning, disclose the limitation and distinguish any permitted fallback deliverable from a completed loop.
