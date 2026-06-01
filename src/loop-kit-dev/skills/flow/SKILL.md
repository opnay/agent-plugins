---
name: flow
description: Interpret a message, action, plan item, review finding, or handoff as a cohesive flow; decide active flow vs parent flow vs finite sub-flow candidate vs phase vs handoff; define intake, framing, readiness, purpose continuity, discovery, ambiguity, contract impact, flow-local strategy, phase record checkpoints, verification expectation, and commit-readiness handoff.
---

# Flow

Use this skill to turn work into a bounded flow contract before execution.
A flow is a cohesive work unit that can be understood, reviewed, verified, reported, and handed off.
It is not a lifecycle phase, status label, final QA pass, verification report, evidence repair step, blocker recovery, or commit-readiness report unless that item creates or changes its own reviewable artifact.

Each active flow uses internal phases:

`intake -> framing -> preparation -> work -> verification -> reporting`

`interruption` is not a flow phase.
When `turn-gate` opens interruption routing for a new user message during an active flow, use this skill only to decide contract impact.

## Flow Model

Classify the current item before work:

- `active flow`: selected flow currently being prepared, executed, verified, or reported
- `parent flow`: flow whose output is a finite list of `sub-flow candidates`
- `sub-flow candidate`: proposed later flow with its own scope, non-goals, completion criteria, verification expectation, and handoff condition
- `phase`: internal step of the active flow, not a separate flow
- `handoff`: readiness or routing result, not execution authority

Use two default flow types:

- `operational-preparation`: locks intent, scope, non-goals, success signal, verification expectation, approval boundary, and planned flow list or finite candidates
- `change-unit`: owns a reviewable change to code, docs, fixtures, config, release surfaces, or another artifact

Candidate creation is not execution.
A sub-flow candidate becomes active only through surrounding routing, such as `turn-gate` next-flow selection or an explicitly prepared self-drive sequence.
Flow completion does not close the turn or authorize the next flow.

## Intake

Separate the user's source wording from your interpretation.
Intake owns raw input analysis, goal detection, non-goal detection, authority-sensitive signal detection, and deep-interview topics.

Identify:

- what the user appears to want now
- persistent purpose and current change purpose, when exposed
- explicit or implied non-goals
- scope edges and tradeoffs the user may reject
- acceptance signal or completion expectation
- verification expectation if stated
- authority-sensitive language such as commit, push, PR, publish, release, version bump, destructive action, or external effect
- missing fields that require a high-leverage discovery question

Deep interview is an intake strategy, not a separate phase.
Intake does not execute work; it produces the input contract that framing uses.

## Framing

Framing owns flow decomposition, candidate-vs-selected distinction, artifact ownership, and draft flow contract.
If the current request is too large or owns multiple reviewable artifacts, produce finite `sub-flow candidates` instead of starting execution.

For each flow or candidate, draft:

- label or slug
- flow type
- purpose continuity, only when it changes continuity across flows
- scope and non-goals
- completion criteria
- verification expectation
- approval-sensitive checkpoint, if any
- phase checkpoint expectation
- readiness gaps or missing fields
- recommended question topics or unresolved ambiguity
- recommended flow-local strategy
- handoff condition
- active-flow vs sub-flow-candidate status

Keep completion criteria separate from handoff condition.
Keep missing fields, discovery topics, and recommended strategies separate from execution authority.

## Purpose Continuity

Use purpose continuity when a sequence of flows should keep a stable viewpoint while each flow changes only the current work target.
Treat this as a supporting part of flow framing and readiness, not as a separate flow type or primary entrypoint.

Default shape:

`repository purpose` > `aggregate or package purpose` > `structural or skill purpose` > `change purpose`

When a purpose chain file such as `000-object.md` is explicitly used:

- only `flow` interprets and uses its meaning
- keep it to optional frontmatter plus object-name-and-kind bullets from broadest purpose to current change
- exclude current flow state, verification state, continuity rules, phase logs, next actions, and `turn-gate` routing or closure state
- do not let `turn-gate` redefine it or use it as routing, closure, verification, or approval authority

If purpose continuity changes scope, acceptance, verification path, approval boundary, or handoff, return to intake or framing before work.

## Preparation

Preparation is only for the selected active flow.
Do not use it to redo intake or framing.
Do not enter work until the selected flow contract is ready enough for the risk.

At minimum, check:

- selected active-flow status, not merely candidate status
- user intent and expected result
- included scope and scope edges
- explicit non-goals
- completion or acceptance signal
- tradeoffs the user would reject
- verification expectation
- target and operation
- approval boundary
- handoff condition
- purpose continuity, if present

If the answer to a question would change the artifact, decomposition, verification path, approval checkpoint, purpose chain, or handoff condition, return to intake or framing instead of beginning work.

## Discovery And Ambiguity

Use discovery when intent, scope edge, non-goal, tradeoff, acceptance signal, verification expectation, purpose continuity, or candidate decomposition is not locked.
Discovery identifies question topics only; question UI and next-flow routing are owned outside this skill.

Ask for the smallest high-leverage decision that can lock the contract.
If bounded choices can lock it, provide choices with tradeoffs.
If free-form input is needed, identify whether the needed answer is an example, counterexample, non-goal, rejected tradeoff, success signal, verification expectation, purpose chain, or decomposition preference.

Use ambiguity handling when an operation or target can point to multiple structures and the interpretation changes scope, output, verification path, approval sensitivity, reporting scope, record reconstruction, purpose continuity, or handoff.
Structural verbs such as `move`, `promote`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, and `contract` are ambiguity triggers when they can change the flow contract.

Record interpreted operation, target, alternate interpretations, why the difference matters, and which contract field must be locked before work.

## Contract Impact

When a new message arrives during an active flow, decide whether it:

- leaves the current flow contract unchanged and can be answered inline
- changes scope, non-goals, completion criteria, verification expectation, approval boundary, purpose continuity, or handoff condition
- should become a new foreground flow while preserving or backgrounding the current flow
- should be reserved as a future candidate
- supersedes the current flow
- reveals a blocker question
- explicitly stops the turn

For self-drive or prepared sequences, also decide whether current flow completion, pass-aligned verification expectation, non-blocked handoff, and next flow identity are present.
Self-drive may use those outputs; this skill does not advance the sequence.

## Phase Record Checkpoints

At the start and end of each active-flow phase, decide whether `000-plan.md` or the active flow record needs an update.
This skill defines checkpoint expectations; the active turn controller applies record updates when it owns that runtime surface.

Use these checkpoint expectations:

- `intake` start: raw request source, interpretation boundary, pending input-analysis fields, next intake action
- `intake` end: goal, non-goals, authority-sensitive signals, missing discovery topic, whether framing may begin
- `framing` start: item to classify, candidate boundary, artifact ownership question, decomposition risk
- `framing` end: selected active flow or finite candidates, candidate-vs-selected status, draft verification expectation, readiness gaps
- `preparation` start: flow label, type, scope boundary, pending contract fields, next preparation action
- `preparation` end: readiness, locked scope and non-goals, missing questions or blockers, selected strategy, whether work may begin
- `work` start: active flow boundary, next action, approval-sensitive checkpoint status, expected artifact
- `work` end: changed artifact or work result, issue found, next phase, changed verification expectation if any
- `verification` start: method, evidence needed, target surface, known limitation
- `verification` end: pass/fail/blocked/insufficient status, evidence gathered, residual risk, earliest safe next phase for non-pass
- `reporting` start: reportable result, verification status, handoff condition, unresolved question or blocker
- `reporting` end: reported outcome, residual risk, handoff condition, next-flow candidate without executing it

Update `000-plan.md` when the active flow, turn-level required next action, or active skill list changes.
Update the active flow record when phase state, execution evidence, verification evidence, report outcome, or residual risk changes.
If a trivial read-only judgment needs no record change, make that reason visible in the report or active flow record.

## Flow-Local Strategies

Choose a strategy inside the active flow only after readiness is sufficient:

- `review-loop`: handle one bounded blocking review, QA, or self-review finding tied to correctness, regression risk, reliability, or delivery
- `fix-verify-loop`: test one primary issue with the smallest useful fix or confirmation action, verify immediately, then reassess
- `broad-execution`: execute one locked active flow end to end when scope, non-goals, completion criteria, verification expectation, and approval boundary are clear

Do not use a flow-local strategy as authority to continue through multiple flows.
If a finding or fix changes scope, approval boundary, destructive/external action, completion criteria, or verification expectation, return to preparation or handoff.

## Verification And Handoff

Set verification expectation from the flow risk and changed surfaces.
After work, verify against that expectation or mark what evidence is missing.

For non-pass verification:

- missing evidence returns to current flow verification
- work evidence vs metadata mismatch returns to verification reconciliation
- changed target, scope, approval boundary, purpose continuity, or verification expectation returns to preparation
- user input, approval, access, or external state need becomes blocked handoff

Do not create a new flow for verification, reporting, evidence repair, or blocker recovery unless that repair creates or changes its own reviewable artifact.

For commit-readiness, judge only whether handoff conditions are ready: intended change unit, diff scope, unrelated-change exclusion, verification evidence, and residual risk.
Commit-readiness is not commit execution, staging, pushing, PR creation, publishing, release, or version bump authority.
