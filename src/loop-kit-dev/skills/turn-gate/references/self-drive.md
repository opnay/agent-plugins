# turn-gate self-drive overlay

Self-drive is an overlay for a prepared finite flow sequence. It is not a separate skill entrypoint and it does not remove the core `preparation -> work -> verification -> reporting -> next-flow` loop.

Use this reference only when self-drive has been explicitly requested or is already active in current records.

## Activation Requirements

Apply self-drive only after records contain:

- sequence objective
- prepared flow sequence
- active flow index
- current flow label
- allowed autonomous actions
- prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- acceptance signal
- verification expectation

`000-plan.md` stores self-drive status and sidecar pointer. `000-self-drive.md` stores sequence-level state. Each active flow record stores only its flow-local sequence position, local progress note, next handoff, and blocker return condition.

## Flow-Start Sidecar Gate

At the start of each self-drive flow, read the plan pointer and `000-self-drive.md`. Confirm:

- `status`
- `active_flow_index`
- `current_flow_label`
- `planned_flow_count`
- `endpoint`
- `required_next_action`
- acceptance signal
- blocker state

If values are missing, conflicting, or do not identify the current active flow, reconcile by flow name/file/slug or return to user-gated routing. If `active_flow_index >= planned_flow_count`, treat the sidecar as stale or corrupt. Do not use modulo, wraparound, or a remembered next label.

## Interruption Handling

When a user message arrives during active self-drive, interpret it inside the active sequence first unless it is an explicit turn stop.

Priority:

1. Source-recorded explicit stop: record closure and report toward stop.
2. Destructive, external, commit, push, PR, publish, release, version bump, or approval-boundary-expanding request: stop self-drive and ask for approval.
3. Scope, non-goal, endpoint, target, prepared order, or acceptance-signal change: stop self-drive and relock the sequence. A clear future endpoint constraint that does not change the current active flow boundary, such as "stop when the listed items are exhausted", may be recorded as a source-backed endpoint update and the current boundary may continue.
4. Blocker or repeated failure: route to earliest safe repair phase or user-gated blocker decision.
5. Status/progress question only: report current phase, active flow, verification state, and next action; continue if no higher rule applies.
6. Ordinary note inside the recorded boundary: record material details and continue.

Self-drive narrows question conditions; it does not disable questions.

## Verification And Endpoint

Each flow must still verify before reporting. Before endpoint exhaustion handling, route non-pass verification first:

- `fail`: repair or return to work.
- `insufficient`: gather evidence or repair verification.
- `blocked`: open user-gated blocker routing.

Only after verification passes may you evaluate sequence exhaustion. Read the sidecar endpoint and endpoint handling body again. Follow only recorded behavior: self-drive stop, handoff, repeat cycle, blocker decision, or next-flow reopening.

Open-ended self-drive still needs a finite current cycle. Do not treat "forever" or "until stopped" alone as autonomous continuation authority. Repeat inventory loops are allowed only when the endpoint explicitly permits bounded repeats and the sidecar is refreshed for the new cycle.

## Approval Boundary

Self-drive can execute approval-sensitive actions only when the initial preparation recorded exact action, target, expected effect, risk, recovery path, included/excluded scope, and endpoint. Otherwise return to user-gated approval routing.

Subagents may support evidence readback, status synthesis, or low-risk local judgment inside recorded boundaries. They do not replace approval for scope, endpoint, or approval-sensitive execution.
