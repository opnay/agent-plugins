# Phase Protocols

This reference is the runtime contract for choosing and applying phase protocols inside `turn-gate`. Protocols are not modes and not standalone user entrypoints. They are phase-local operating contracts under the implicit default state.

## Route First

Before selecting a protocol:

- confirm the active flow is source-recorded or a sibling `flow` decision has been applied
- resolve operation/target ambiguity when it can change files, scope, routing, deletion, approval, or handoff
- confirm approval boundaries before destructive, irreversible, external, commit, push, PR, publish, release, or version-bump actions
- choose the earliest blocker rather than the most convenient protocol

Default priority when several protocols seem plausible:

1. `deep-interview`
2. `review-loop`
3. `ralph-loop`
4. `autopilot`
5. `commit-readiness-gate`

## deep-interview

Use during `preparation` when requirement discovery or scope lock blocks work.

Apply when intent, scope, non-goals, acceptance signal, verification expectation, approval boundary, or expected risky action is not sufficient to proceed. Prefer `request_user_input` for bounded choices.

Do not use it for a simple operation/target ambiguity that can be locked with one narrow clarification.

Handoff: once sufficiently locked, return to protocol routing. If still broad or unsafe, stay in user-gated question routing.

## review-loop

Use when review feedback, QA findings, or self-review findings are the current blocker.

Apply one bounded blocking finding at a time. Fix, verify the finding, and reassess. Low-value notes stay as follow-up candidates unless they block correctness, reliability, or delivery.

Handoff: return to verification after the finding is addressed. If the finding expands scope or approval boundary, return to preparation or user-gated routing.

## ralph-loop

Use for a narrow fix-verify-reassess cycle.

Apply when one primary issue is clear and a small change can test the hypothesis. Keep each loop small, verify immediately, and reassess whether another loop is justified.

Handoff: continue only while the next cycle remains inside the active flow boundary. If scope grows, split or ask.

## autopilot

Use when scope is locked and the current flow needs broad end-to-end execution.

Apply inside recorded scope, non-goals, verification expectation, and approval boundary. Continue through implementation, QA, validation, and reporting unless blocked.

Autonomous execution is not authority for destructive, external, commit, push, PR, publish, release, or version-bump actions.

Handoff: report after implementation and verification. If a new approval boundary appears, return to user-gated routing.

## commit-readiness-gate

Use when judging whether an intended change unit is ready to commit.

Evaluate intended diff scope, unrelated changes to exclude, verification evidence, residual risk, and likely commit-message scope. This is readiness reporting, not commit execution approval.

Handoff: if ready, open approval-boundary handoff or next-flow routing. If not ready, return to the earliest safe repair or verification phase.

## Local Reference Rule

If a protocol's meaning would change the active flow or approval boundary, ask before work. If it only chooses how to perform the current phase inside an already locked flow, record the protocol and continue.
