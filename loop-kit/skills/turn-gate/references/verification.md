# Verification

Use this reference to select and record verification before reporting.

## Method Is Separate From Result

Record both:

- `method`: `clean-context`, `normal`, or `not-required`
- `result_status`: `pass`, `fail`, `blocked`, or `insufficient`

`not-required` is not a pass. It means no separate verification action is justified. Record the reason and residual uncertainty.

Progress states such as `not-started` and `requested` may appear in compact continuity metadata before a result exists. They are not `Result.status` values and cannot support success reporting. When preserving prior flow state, keep the existing `verification_status` value and record the preservation in `continuity`; do not use `verification_status: preserved`.

## Methods

Use `clean-context` for a bounded read-only verifier packet. It must not be a full-history fork.

Use `normal` for main-thread checks, readback, evidence review, and logical counterexample review.

Use `not-required` only when a separate verification action would not add meaningful confidence for the recorded risk.

## Clean-Context Default

Default to `clean-context` for:

- file changes
- generated release surface changes
- multi-file contract changes
- prior check failures
- user-requested verification, review, QA, or commit-readiness
- approval-sensitive action boundaries

The verifier packet must include:

- target
- user intent
- changed files or artifacts
- checks or evidence to inspect
- pass/fail criteria
- no edit permission
- no scope expansion
- no destructive or external work
- no commit, push, PR, publish, release, or version-bump action

Generated release surface build/readback and commit-readiness judgment are verification or preparation evidence only. They do not authorize commit, publish, release, version bump, destructive work, or external side effects. Skip clean-context verification only when the flow record explains why `normal` or `not-required` covers the actual risk.

## Result Routing

Before success reporting:

- `pass`: report evidence that supports the acceptance signal.
- `fail`: return to the earliest safe repair or work point.
- `insufficient`: collect more evidence or strengthen verification.
- `blocked`: open blocker routing for the needed input, access, approval, or external state change.

Non-pass status takes priority over self-drive continuation, endpoint exhaustion, release readiness, commit-readiness, and next-flow continuation.
