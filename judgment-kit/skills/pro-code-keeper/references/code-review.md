# Code Review

Use this reference for overengineering review, simplification review, removable-code review, dependency reduction review, or code-shrinking review.

Default mode is read-only. Do not edit unless the user explicitly asks for implementation after the review.

## Survey Items

Inspect enough to judge the target:

- current call flow and ownership boundary
- callers and public contract
- existing helpers, types, schemas, platform features
- dependency usage and package boundary
- tests or checks proving current behavior
- generated/build/vendor files to exclude

## Finding Tags

- `delete`: dead code, guessed feature, unused flexibility, duplicate branch, stale path.
- `stdlib`: custom code replaceable by standard library.
- `native`: code or dependency replaceable by language, browser, OS, database, or framework behavior.
- `yagni`: abstraction, option, factory, hook, config, or extension point not needed by current requirements.
- `shrink`: same behavior expressible with less code or fewer concepts.

Do not tag security, correctness, accessibility, or data-loss protections as removable. Report those separately as risks only when they are wrong or misplaced.

## Finding Test

A finding is valid only when it names:

- location
- target to remove or replace
- current requirement that proves it is unnecessary
- safer replacement or deletion path
- verification needed after change

If any field is missing, keep inspecting or mark it as a question.

## Output

Use one line per finding:

`tag: location - target -> replacement. reason. verify: check.`

Order by impact, then confidence.
