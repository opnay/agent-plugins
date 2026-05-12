# Review Loop Reference

Use this phase protocol when review feedback, QA findings, or self-review findings are the current bottleneck and only material issues should block progress.

## Core Contract

- Apply a blocking threshold before choosing a finding.
- Keep one loop focused on one bounded blocking finding.
- Fix the issue and verify immediately.
- Leave low-value notes as deferred follow-up instead of widening the loop.
- If a finding points to broader ambiguous scope, a changed verification path, or a new approval boundary, return to preparation or question-routing before expanding work.
- Do not treat review feedback as approval for destructive, irreversible, external, commit, push, PR, publish, release, version bump, or similar sensitive action.

## Protocol Boundary

- Good fit: correctness, regression, reliability, or delivery-blocking review feedback that fits the current flow boundary.
- Not a fit: open-ended cleanup, speculative polish, broad implementation work, unrelated finding collection, unresolved approval boundary, commit execution, PR, or publish handoff.

## Handoff

- After the finding is handled, verify the result.
- If additional findings remain, decide whether they fit the same flow boundary or should become follow-up candidates.

## Review Questions

- Is the chosen finding truly blocking?
- Did the loop stay bounded to one finding?
- Did low-value notes remain deferred instead of expanding scope?
- Did a risky-action finding return to user-gated question-routing unless already covered by initial agreement?
