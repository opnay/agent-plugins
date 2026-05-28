---
name: flow
description: Interpret a message, action, plan item, review finding, or handoff as a cohesive flow; decide whether it is an active flow, parent flow, finite sub-flow candidate, phase, or handoff; and define intake, framing, preparation readiness, discovery, ambiguity, contract impact, flow-local strategy, phase record checkpoints, verification expectation, and commit-readiness handoff.
---

# Flow

Use this skill to turn work into a bounded flow contract before execution.
A flow is a cohesive work unit that can be understood, reviewed, verified, reported, and handed off.
It is not merely a phase label such as `preparation`, `work`, `verification`, `reporting`, or `commit readiness`.

Each active flow has internal phases: `intake -> framing -> preparation -> work -> verification -> reporting`.
Keep those phases inside the same active flow unless a separate reviewable work unit is needed.

`interruption` is not a flow phase.
When `turn-gate` opens interruption routing for a new user message during an active flow, use this skill only to decide contract impact: inline answer, current-flow revision, background/new foreground flow, future candidate, blocker topic, or explicit-stop implication.

## Intake

Separate the user's source wording from your interpretation.
Intake owns raw input analysis, goal detection, non-goal detection, authority-sensitive signal detection, and deep-interview topics.

Identify:

- what the user appears to want now
- explicit or implied non-goals
- scope edges and tradeoffs the user may reject
- acceptance signal or completion expectation
- verification expectation if stated
- authority-sensitive language such as commit, push, PR, publish, release, version bump, destructive action, or external effect
- missing fields that require a high-leverage discovery question

Deep interview is an intake strategy, not a separate phase.
Surface the smallest question topic that would change the flow contract.
Intake does not execute a flow; it produces the input contract that framing uses.

## Framing

Before acting, classify the current item:

- `active flow`: the selected flow currently being prepared, executed, verified, or reported
- `parent flow`: a flow whose output is a finite list of `sub-flow candidates`
- `sub-flow candidate`: a proposed later flow with its own scope, non-goals, completion criteria, verification expectation, and handoff condition
- `phase`: an internal step of the active flow, not a flow by itself
- `handoff`: a readiness or routing result, not execution authority

Pure analysis, QA, consistency checks, verification reporting, evidence repair, blocker recovery, and commit-readiness reporting are not separate change-unit flows unless they create or modify a reviewable artifact.
Code, docs, fixtures, config, snapshots, operator reports, validator output, or release surfaces can be flow-owned artifacts when they are the actual reviewable result.
If a request asks for a fixed number of flows, count reviewable work units, not phase labels.
Verification and reporting become separate flows only when they own a separate reviewable artifact.

Framing owns flow decomposition, flow design, candidate-vs-selected distinction, and artifact ownership.
If the current request is too large, produce finite `sub-flow candidates` instead of starting execution.
Candidate creation is not execution.
Only the selected active flow moves into preparation.

## Flow Types

Use two default flow types:

- `operational-preparation`: locks intent, scope, non-goals, success signal, verification expectation, approval boundary, and a planned flow list or finite candidates
- `change-unit`: owns a reviewable change to code, docs, fixtures, config, release surfaces, or another artifact

An operational-preparation flow may produce change-unit candidates.
Those candidates remain inactive until selected by surrounding routing or by an explicitly prepared sequence.

## Flow Contract Output

When designing a flow, sub-flow candidate, or contract-impact decision, make these fields visible when relevant:

- flow label or slug
- flow type: `operational-preparation` or `change-unit`
- scope
- non-goals
- completion criteria
- verification expectation
- approval-sensitive checkpoint, if any
- phase start/end record checkpoint expectation
- readiness status
- missing contract fields or unresolved edges
- discovery topic or ambiguity to resolve
- recommended flow-local strategy
- handoff condition
- contract-impact result for a new message, if one is being classified
- unresolved questions or blocker
- active-flow vs sub-flow-candidate status

Keep completion criteria separate from handoff condition.
Keep discovery topics, unresolved fields, and recommended strategies separate from authority to execute.

## Preparation

Preparation is only for the selected active flow.
Do not use it to redo intake or framing.
Do not enter work until the selected flow contract is ready enough for the risk.

At minimum, check:

- user intent and expected result
- included scope and scope edges
- explicit non-goals
- completion or acceptance signal
- tradeoffs the user would reject
- verification expectation
- target and operation
- approval boundary
- handoff condition
- active-flow status, not merely sub-flow-candidate status

If an answer would change the artifact, decomposition, verification path, approval checkpoint, or handoff condition, return to intake or framing instead of beginning work.

## Discovery

Use discovery when intent, scope edge, non-goal, tradeoff, acceptance signal, verification expectation, or candidate decomposition is not locked.
Discovery usually belongs to intake; if framing or preparation reveals the missing field later, route back to the earliest stage that can lock it.

Surface the smallest high-leverage topic that would change the flow contract.
Prefer one bounded question topic at a time.
If bounded choices can lock the contract, provide the choices with tradeoffs.
If free-form input is needed, identify whether the needed answer is an example, counterexample, non-goal, rejected tradeoff, success signal, verification expectation, or decomposition preference.

Discovery identifies question topics only.
The mechanism for asking the user, choosing a question tool, continuing the turn, or routing the next flow is owned outside this skill.

## Ambiguity

Use ambiguity handling when the operation or target can point to multiple structures and the interpretation changes scope, output, verification path, approval sensitivity, reporting scope, record reconstruction, or handoff.

Record:

- interpreted operation and target
- alternate interpretations
- why the difference matters
- which contract field must be locked before work

Relative dates, record dates, previous-flow references, "current target" wording, and structural verbs such as `move`, `promote`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, or `contract` are ambiguity triggers when they can change the flow contract.

If ambiguity affects the result, return to intake discovery, framing, or a parent-flow candidate output instead of beginning work.

## Contract Impact

Use contract-impact classification when a new message arrives while an active flow already exists.
This does not replace `turn-gate` interruption routing; it gives that routing a flow-owned decision.

Decide whether the message:

- leaves the current flow contract unchanged and can be answered inline
- changes scope, non-goals, completion criteria, verification expectation, approval boundary, or handoff condition
- should become a new foreground flow while preserving or backgrounding the current flow
- should be reserved as a future candidate
- supersedes the current flow
- reveals a blocker question
- explicitly stops the turn

For self-drive or prepared sequences, also decide whether current flow completion, pass-aligned verification expectation, non-blocked handoff, and next flow identity are present.
Self-drive may use those outputs; this skill does not advance the sequence.

## Phase Record Checkpoints

At the start and end of each active-flow phase, decide whether `000-plan.md` or the active flow record needs an update.
This skill defines the checkpoint expectation; the active turn controller applies record updates when it owns that runtime surface.

When an active flow or planned sequence requires specific skills, expect `000-plan.md` to expose the active skill list.
Keep that list limited to the current selected flow and prepared future flows.
Do not record speculative candidate skills as active.
For turn-gate-managed user-message flows entering preparation, expect `turn-gate` and `flow` to be reread and listed.

Use these checkpoint expectations:

- `intake` start: expose raw request source, current interpretation boundary, pending input-analysis fields, and next intake action
- `intake` end: expose goal, non-goals, authority-sensitive signals, missing discovery topic, and whether framing may begin
- `framing` start: expose item to classify, candidate boundary, artifact ownership question, and decomposition risk
- `framing` end: expose selected active flow or finite sub-flow candidates, candidate-vs-selected status, draft verification expectation, and readiness gaps
- `preparation` start: expose flow label, type, scope boundary, pending contract fields, and next preparation action
- `preparation` end: expose readiness, locked scope and non-goals, missing questions or blockers, selected strategy, and whether work may begin
- `work` start: expose active flow boundary, next action, approval-sensitive checkpoint status, and expected artifact
- `work` end: expose changed artifact or work result, issue found, next phase, and any changed verification expectation
- `verification` start: expose verification method, evidence needed, target surface, and known limitation
- `verification` end: expose pass, fail, blocked, or insufficient status; evidence gathered; residual risk; and the earliest safe next phase if verification did not pass
- `reporting` start: expose reportable result, verification status, handoff condition, and unresolved question or blocker
- `reporting` end: expose reported outcome, residual risk, handoff condition, and any next-flow candidate without executing it

Update `000-plan.md` expectation when the active flow, turn-level required next action, or active skill list changes.
Update active flow record expectation when phase state, execution evidence, verification evidence, report outcome, or residual risk changes.
If a trivial read-only judgment needs no record change, make that reason visible in the report or active flow record.

## Flow-Local Strategies

Choose a strategy inside the active flow only after readiness is sufficient:

- `review-loop`: handle one bounded blocking review, QA, or self-review finding that materially affects correctness, regression risk, reliability, or delivery
- `fix-verify-loop`: use one small fix or confirmation action to test one primary issue, verify immediately, then reassess whether another loop is justified
- `broad-execution`: execute a single locked active flow end to end when scope, non-goals, completion criteria, verification expectation, and approval boundary are clear

Do not use a flow-local strategy as authority to continue through multiple flows.
If the issue introduces new scope, a new approval boundary, destructive or external effects, or changed completion criteria, return to intake, framing, preparation, or handoff.

## Verification And Handoff

Set verification expectation from the flow risk and changed surfaces.
After work, verify against that expectation or mark what evidence is missing.
A verification failure or insufficient evidence sends the flow back to the earliest safe phase instead of success reporting.

For verification blockers, choose the earliest safe phase explicitly:

- missing evidence returns to current flow verification
- work evidence vs metadata mismatch returns to verification reconciliation
- changed target, scope, approval boundary, or verification expectation returns to preparation
- user input, approval, access, or external state need becomes blocked handoff

Do not create a new flow for evidence repair or blocker recovery unless that repair creates or changes its own reviewable artifact.

For commit-readiness, judge only whether handoff conditions are ready: intended change unit, diff scope, unrelated-change exclusion, verification evidence, and residual risk.
Commit-readiness is not commit execution, staging, pushing, PR creation, publishing, release, or version bump authority.

Flow completion means the active flow met its completion criteria and produced its reporting or handoff condition.
It does not mean the surrounding turn is closed, and it does not automatically execute the next flow.
