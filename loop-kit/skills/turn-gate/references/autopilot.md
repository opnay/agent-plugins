# Autopilot Reference

Use this phase protocol when the current phase is broad end-to-end execution from a locked request to a verified result.

## Core Contract

- Run a multi-phase delivery workflow from locked scope through implementation, QA, validation, and delivery.
- Ask only blocking clarification questions; otherwise keep execution moving autonomously within the recorded scope, non-goals, verification expectation, and approval boundaries.
- Do not start if the scope floor is unmet. Return to `deep-interview` or active question-routing when scope can change output, completion criteria, verification, or expected risky-action handling.
- Do not treat autonomous execution as approval for destructive, irreversible, external, commit, push, PR, publish, or similar sensitive work.
- Continue through expected risky actions only when the initial preparation explicitly approved that exact boundary. Otherwise return to user-gated question-routing.
- Verify after meaningful edits with the checks that fit the changed surface.
- Track QA issues when failures need iterative fixes.
- Stop and report a root blocker if critical ambiguity prevents safe execution or the same QA failure repeats.

## Protocol Boundary

- Good fit: hands-off implementation, multi-phase feature work, requests that span requirements, code changes, tests, and validation after scope is locked.
- Not a fit: brainstorming, planning-only requests, one small bounded fix, review-finding triage, missing scope lock, unresolved approval boundary, or final commit readiness.

## Review Questions

- Is the task broad enough to need end-to-end delivery rather than a bounded loop?
- Are scope assumptions clear enough to start without a full requirement interview?
- Are expected risky actions and approval boundaries recorded?
- Are verification and QA loops explicit enough to support autonomous progress?
