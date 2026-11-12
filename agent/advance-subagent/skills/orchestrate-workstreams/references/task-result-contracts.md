# Task and Result Contracts

These are logical records, not a required serialization format. Use the same field meanings in Markdown, YAML, or an equivalent existing task system.

## Task Packet

~~~yaml
round_id: <current round>
task_id: <stable task id>
role: PLAN_ROUND | EXPLORE_READ | IMPLEMENT_OWNED | REVIEW_LENS | PROCESS_STRUCTURED | FRONTIER_JUDGMENT
model_profile: fast | default | high
profile_reason: <applicable instruction or task need>
model: <resolved profile model; omit for default>
reasoning_effort: <resolved effort; omit for default>
objective: <bounded question, plan, or deliverable>
scope:
  included: <assigned surfaces>
  excluded: <non-goals>
prerequisites: <verified inputs or none>
inputs: <necessary original or verified material>
source_of_truth: <authoritative sources and applicable instructions>
domain_contract: <relevant installed contract or explicit task rules>
access_mode: read-only | write-enabled
owned_files_or_artifacts: none | <exact writable ownership>
constraints:
  - No nested orchestration or delegation.
  - No scope or permission expansion.
  - No writes outside ownership.
deliverable: <integrable output>
evidence: <recoverable sources and verification locations>
validation: <observable checks>
completion_conditions: <local task completion, even if global criteria are unresolved>
stop_conditions: <dependencies, evidence, retry, allowance, or capability limits>
~~~

For planners, require a proposed round plan and unresolved questions, set read-only ownership to none, and allow only the minimal handoff and its necessary source materials. Do not supply previous worker narratives or preferred answers.

For researchers, attach $advance-subagent:deep-research as an assigned domain contract, not a request to launch another loop. Other domains supply their own applicable requirements.

Record the actual spawn's returned agent ID and exact packet in the main-owned audit record. For default, omit model and effort from actual spawn arguments. Do not pass placeholder values.

## Result Envelope

~~~yaml
status: completed | blocked | inconclusive
summary_and_claims:
  - claim: <bounded statement>
    basis: confirmed | inference
evidence: <exact sources, paths, locations, commands, and key results>
inspected_surfaces: <inspected files, sources, artifacts, or systems>
changed_surfaces: <actual changes and reasons, or none>
validation:
  performed: <checks and outcomes>
  skipped: <checks and reasons>
  failed: <checks and failures>
risks_and_uncertainty: <remaining limits>
recommended_integration_action: <suggestion for the main agent>
~~~

A planner also returns the proposed round plan. A research result also supplies the source dates, supporting passages, counterevidence, and scope limits required by its domain contract.

Completed means the local packet's evidence and validation conditions are met. Blocked means an external dependency, input, tool, authority, or allowance prevents completion. Inconclusive means permitted work was done but evidence does not settle the question.

Missing validation is not success. Ownership violations are not accepted changes. The main agent verifies every required result and decides whole-request completion separately.
