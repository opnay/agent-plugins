# Review Loop Reference

Use this mode when the current phase is driven by review feedback, QA findings, or self-review findings and only material issues should block progress.

## Core Contract

- Apply a blocking threshold before choosing a finding.
- Keep one loop focused on one bounded blocking finding.
- Fix the issue and verify immediately.
- Leave low-value notes as deferred follow-up instead of widening the loop.
- If a finding points to a broader ambiguous scope, changed verification path, or new approval boundary, return to preparation or question-routing before expanding work.
- Do not treat review feedback as approval for destructive, irreversible, external, commit, push, PR, publish, or similar sensitive action.

## Mode Boundary

- Good fit: correctness, regression, reliability, or delivery-blocking review feedback.
- Not a fit: open-ended cleanup, speculative polish, broad implementation work, unrelated finding collection, unresolved approval boundary, or commit execution.

## Review Questions

- Is the chosen finding truly blocking?
- Did the loop stay bounded to one finding?
- Did low-value notes remain deferred instead of expanding scope?
- Did a risky-action finding return to user-gated question-routing unless already covered by the initial agreement?
