---
name: turn-gate
description: Use when a task must stay inside one continuing turn until the user explicitly stops it, especially for multi-flow work that needs preparation, implementation or investigation, clean-context verification, reporting, and active next-flow reopening.
---

# turn-gate

## Important

When this skill is active, `turn-gate` is a conversation-level rule, not a checklist you may summarize and close. Keep the current turn open until the user explicitly asks to end this turn.

Do not send a terminal summary unless an explicit stop is source-recorded for the current turn. A completed task, successful command, finished report, or exhausted plan is not a turn-ending signal. If the latest user message is not an explicit turn stop, treat it as continuation input.

Every active turn-gated flow must end in one of these states:

- loop continuation into a next flow or safe handoff
- active user-gated question routing
- blocked question routing with the blocker made visible
- explicit user stop with closure source recorded for this turn

After reporting, reopen the next flow with `request_user_input` when it is available and bounded choices can be offered. Use a plain-text fallback only when that tool is unavailable; the fallback must still be an active question-routing state, not a closing phrase.

Maintain `.agents/sessions/{YYYYMMDD}/000-plan.md` and `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` records for active turn-gated work. Refresh the active flow's `Continuity Guard` before result reporting and next-flow reopening. Always record a turn-end option in `Next Flow Options`, even when visible choices are limited.

## Purpose

Use `turn-gate` to keep one user-directed session moving through coherent flows until the user explicitly stops. Each flow follows:

1. preparation
2. work
3. verification
4. reporting
5. next-flow reopening

A flow is not a phase such as "analysis", "work", "verification", "reporting", or "commit readiness". A flow is a cohesive, reviewable, or commit-sized unit. It may be an `operational-preparation` flow that owns session planning artifacts, or a `change-unit` flow that owns actual code, docs, fixtures, config, or other reviewable changes.

## Session Records

Use `.agents/sessions/{YYYYMMDD}/000-plan.md` for the date-scoped multi-flow plan. Use `templates/plan-template.md` as the local shape when creating or repairing it.

Use `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` for each flow record. `count-pad3` is `001`, `002`, `003`, and the slug uses English lowercase words joined by `-`. Use `templates/flow-record-template.md` as the local shape when creating or repairing a record.

Update records incrementally after each phase. `000-plan.md` owns the planned flow sequence and transition criteria. Each flow record owns what happened inside that flow.

Each flow record must keep a `Continuity Guard` with:

- turn-gate active status
- question-routing mode
- user explicit stop status
- terminal summary allowed status
- required next action
- last refreshed phase
- pending or superseded question state
- verification status

Only record `confirmed closure` when there is a source user message that explicitly ends the current turn. A stale closure, source-less closure, or old `terminal summary allowed: yes` does not permit terminal close.

## Flow Boundaries

Separate two flow types:

- `operational-preparation`: interprets the user request, locks intent/scope/non-goals/acceptance/verification/approval boundaries, and creates or updates session plan artifacts and planned flow candidates.
- `change-unit`: owns actual reviewable changes or artifacts.

User-message interpretation and planned flow list design may be an `operational-preparation` flow. Its output may be follow-up `change-unit` candidates, but those candidates are not active execution flows until selected or handed off.

Do not promote pure final QA, consistency checking, verification result reporting, or commit-readiness reporting into planned flows unless they create or modify a distinct reviewable artifact.

Avoid phase-only planned flows. Planned flows must be cohesive units that can be understood, reviewed, verified, and optionally committed together.

## Preparation

Before work, use the flow shaping gate to lock the active flow boundary and completion criteria. Confirm or record:

- intent
- scope
- non-goals
- acceptance signal
- verification expectation
- approval boundary
- current target state and stale assumptions

Use these preparation methods as needed:

- `references/deep-interview.md` for requirement discovery, missing scope, unclear success criteria, or unresolved approval lines.
- flow list design when the user request needs a planned sequence.
- meaning resolution when operation or target wording could point to different files, specs, phases, routing rules, release surfaces, or destructive meanings.
- current-state inspection and target rereads before editing or judging a changed surface.

If scope is empty, too broad, can produce multiple valid outputs, or can change the verification path, ask before work. Prefer `request_user_input` for bounded choices. If you infer a safe scope without asking, record the inferred work boundary and non-goals.

For user-message-driven preparation, collect enough information for the planned flow list, not just the first flow. Identify expected risky actions and whether each is approved, not approved, deferred, or handoff-required.

## Approval Boundaries

Meaning resolution is not approval. Readiness is not approval.

Before destructive, irreversible, external, commit, push, PR, publish, release, promote, version bump, or similar sensitive action, get user approval for the exact target, expected effect, risk, recovery path, included scope, and excluded scope.

Commit-readiness reporting may say whether a change unit appears ready. It must not stage, commit, push, open a PR, publish, release, promote, bump a version, or imply approval for those actions. Those are separate user-gated handoffs.

If target ambiguity and approval-sensitive action appear together, lock the target/meaning first, then ask for action approval. One question may ask both only if the two decisions are clearly separated.

## Work

Before work, choose one internal mode for the current phase and read the matching local reference:

- `references/deep-interview.md`: requirement discovery or missing scope boundaries.
- `references/review-loop.md`: review feedback, QA finding, or self-review finding is the current blocker.
- `references/ralph-loop.md`: one bounded fix-verify-reassess loop is enough.
- `references/autopilot.md`: broad end-to-end delivery from locked scope to verified result.
- `references/self-drive.md`: a prepared planned flow sequence can continue through bounded subagent decisions without asking the user at every step.
- `references/commit-readiness-gate.md`: the change unit is mostly complete and readiness judgment is the current work.

If several modes fit, prefer the earliest bottleneck in this order: `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `self-drive`, `commit-readiness-gate`.

External actions are not internal modes. Commit, push, PR, publish, release, promote, and version bump work require user-gated handoff.

## Internal Gates

Use these internal gates to control transitions:

- Message intake: classify the latest user message, detect explicit turn stop, and mark ambiguity, scope gaps, or possible approval boundaries. It does not execute work or close flows.
- Flow shaping: decide whether to create a new flow, update the active flow, report, or ask a question. It owns active flow boundary, completion criteria, verification expectation, and follow-up candidate separation.
- Task policy: choose commands, edits, builds, tests, rereads, local references, and handoffs inside the active flow only. It cannot close a flow, skip reporting, choose the next flow, or end the turn.
- Verification: integrate work evidence and verifier output as `pass`, `fail`, `blocked`, or `insufficient`.
- Reporting: summarize preparation, work, verification, blockers, residual risk, approval boundary state, and material routing judgments as continuity context.
- Continuation: after reporting, check for source-recorded explicit stop; otherwise reopen next-flow routing, continue a planned loop, hand off self-drive, or ask about a blocker.

No gate may infer terminal closure without an explicit current-turn stop.

## Verification

After work and before reporting, run clean-context verification. This must be a read-only bounded verifier subagent, not a full-history fork and not same-context self-certification.

Send only the bounded packet needed for verification:

- verifier identity or request id
- verification target
- expected user intent
- changed files or artifacts
- commands or checks to run
- pass/fail criteria
- edit permission: none
- forbidden actions: scope expansion, destructive/external work, commit, push, PR, publish, release, promote, version bump
- stop condition

Classify the result as:

- `pass`: reportable completion evidence exists.
- `fail`: return to the earliest safe phase and fix before reporting success.
- `blocked`: route the blocker to the user; do not report success.
- `insufficient`: request better verification or return to the earliest safe phase.

If the verifier tool or subagent is unavailable, treat verification as `blocked` or `insufficient`, not pass.

## Reporting

Reporting is not turn closure. Before reporting, reread or refresh the active flow record's `Continuity Guard`.

Report concisely:

- what was prepared
- what work was done
- what verification ran and its classification
- changed paths or artifacts when relevant
- blockers and residual risk
- approval boundaries still requiring user action
- material routing judgments that affected the flow

If verification did not pass, make that visible and route the next action through blocker question-routing or the earliest safe phase. Do not phrase non-pass work as successful completion.

## Next-Flow Reopening

After reporting, if there is no source-recorded explicit stop for the current turn, reopen the next flow.

Use `request_user_input` when available and bounded choices are possible. Offer narrow choices tied to the just-reported result. If the visible choices cannot include a turn-end choice, still state that the user can explicitly end the turn, and record that option in `Next Flow Options`.

Use fallback only when `request_user_input` is unavailable. In fallback, say the tool is unavailable, list the open choices, and record the required next action. The fallback must leave the turn in active question-routing.

User follow-up questions, reviews, corrections, priority changes, status requests, or next-task requests are continuation inputs. Classify them through message intake and continue the same turn unless they explicitly stop the turn.
