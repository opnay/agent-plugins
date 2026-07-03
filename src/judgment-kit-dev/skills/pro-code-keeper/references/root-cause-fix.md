# Root-Cause Fix

Use this reference for bugs, failing tests, regressions, crashes, broken UI behavior, or debugging requests.

## Rule

Fix the owner of the cause, not the closest symptom.

## Sequence

1. Reproduce or inspect the failing layer.
2. Separate observed symptom from expected behavior.
3. List plausible owners: caller, callee, state, data, adapter, config, build, environment, test.
4. Check the call flow and the data shape at the boundary.
5. Choose the cause that explains the evidence with the fewest extra assumptions.
6. Patch the owning boundary with the smallest complete change.
7. Verify the original failure and one relevant guard path when risk warrants it.

## Symptom Patch Signals

Stop and inspect deeper when the proposed fix:

- hides an exception without handling the bad state
- weakens a test to match broken behavior
- adds fallback data without explaining the missing source
- duplicates validation in a caller because the callee is unclear
- adds a block-list case without defining valid states
- fixes only one instance of repeated failure logic

## Minimal Complete Fix

A root-cause fix should include:

- corrected owner
- preserved public contract or explicit migration
- relevant error or invalid-state handling
- focused verification tied to the observed failure

If the root cause is not proven, say so and report evidence, likely cause, and next check.
