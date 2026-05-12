# Autopilot Reference

Use this phase protocol when the current phase is broad end-to-end execution from a locked request to a verified result.

## Core Contract

- Start only when scope, non-goals, verification expectation, and approval boundary are recorded well enough for execution.
- Keep execution moving inside the recorded boundary. Ask only blocking clarification questions.
- Do not start if the scope floor is unmet. Return to `deep-interview` or active question-routing when scope can change output, completion criteria, verification, or expected risky-action handling.
- Do not treat autonomous execution as approval for destructive, irreversible, external, commit, push, PR, publish, release, version bump, or similar sensitive work.
- Continue through expected risky actions only when the initial preparation explicitly approved that exact boundary and stop point. Otherwise return to user-gated question-routing.
- Verify after meaningful edits with checks that fit the changed surface.
- Track QA issues when failures need iterative fixes.
- Stop and report a root blocker if critical ambiguity prevents safe execution or the same critical QA failure repeats.

## Protocol Boundary

- Good fit: locked-scope implementation, multi-phase feature work, and requests that span code changes, QA, and validation.
- Not a fit: brainstorming, planning-only requests, one small bounded fix, review-finding triage, missing scope lock, unresolved approval boundary, final commit readiness, or commit execution.

## Handoff

- When implementation and validation are complete, hand off to reporting.
- If approval boundary changes or new risky work appears, return to user-gated question-routing.
- Treat commit-readiness reporting and commit execution as separate decisions.

## Review Questions

- Is scope locked enough to begin broad execution?
- Are expected risky actions and approval boundaries recorded?
- Are verification and QA loops explicit enough to support autonomous progress?
- Has any new approval boundary appeared?
