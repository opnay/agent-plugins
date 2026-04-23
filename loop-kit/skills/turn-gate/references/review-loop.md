# Review Loop Reference

Use this mode when the current phase is driven by review feedback, QA findings, or self-review findings and only material issues should block progress.

## Core Contract

- Apply a blocking threshold before choosing a finding.
- Keep one loop focused on one bounded blocking finding.
- Fix the issue and verify immediately.
- Leave low-value notes as deferred follow-up instead of widening the loop.

## Mode Boundary

- Good fit: correctness, regression, reliability, or delivery-blocking review feedback.
- Not a fit: open-ended cleanup, speculative polish, or broad implementation work.

## Review Questions

- Is the chosen finding truly blocking?
- Did the loop stay bounded to one finding?
- Did low-value notes remain deferred instead of expanding scope?
