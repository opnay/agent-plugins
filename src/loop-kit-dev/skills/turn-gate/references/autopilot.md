# Autopilot Reference

Use this mode when the current phase is broad end-to-end execution from a brief request to a verified result.

## Core Contract

- Run a multi-phase delivery workflow from scope lock through implementation, QA, validation, and delivery.
- Ask only blocking clarification questions; otherwise keep execution moving autonomously.
- Verify after meaningful edits with the checks that fit the changed surface.
- Track QA issues when failures need iterative fixes.
- Stop and report a root blocker if critical ambiguity prevents safe execution or the same QA failure repeats.

## Mode Boundary

- Good fit: hands-off implementation, multi-phase feature work, requests that span requirements, code changes, tests, and validation.
- Not a fit: brainstorming, planning-only requests, one small bounded fix, review-finding triage, or final commit readiness.

## Review Questions

- Is the task broad enough to need end-to-end delivery rather than a bounded loop?
- Are scope assumptions clear enough to start without a full requirement interview?
- Are verification and QA loops explicit enough to support autonomous progress?
